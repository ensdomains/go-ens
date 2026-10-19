// Copyright 2026 Weald Technology Trading.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0

package v2

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/holiman/uint256"
)

// Role is one EnhancedAccessControl role: a nybble index in the 256-bit
// role bitmap. Indexes 0 to 31 are regular roles, 32 to 63 their admin
// roles (regular role shifted left by 128 bits).
type Role uint8

// Bit returns the role's bit in the bitmap: 1 << (4 * r).
func (r Role) Bit() *big.Int { return new(big.Int).Lsh(big.NewInt(1), uint(4*r)) }

// Admin returns the admin role that may grant and revoke r.
func (r Role) Admin() Role { return r + 32 }

// IsAdmin reports whether r is an admin role.
func (r Role) IsAdmin() bool { return r >= 32 }

// Base returns the regular role of an admin role, or r itself.
func (r Role) Base() Role {
	if r.IsAdmin() {
		return r - 32
	}
	return r
}

// Registry roles (RegistryRolesLib).
const (
	RoleRegistrar        Role = 0  // register and reserve names; root only
	RoleRegisterReserved Role = 1  // promote a reserved name; root only
	RoleSetParent        Role = 2  // set the registry's canonical parent; root only
	RoleUnregister       Role = 3  // unregister a name
	RoleRenew            Role = 4  // extend a name's expiry
	RoleSetSubregistry   Role = 5  // change a name's child registry
	RoleSetResolver      Role = 6  // change a name's resolver
	RoleCanTransfer      Role = 7  // only RoleCanTransfer.Admin() exists; auto-granted to the owner
	RoleWasReserved      Role = 8  // marks a name registered from a reservation (migrated from v1)
	RoleSetURI           Role = 9  // set the token URI
	RoleCanName          Role = 30 // name the contract through the ContractNamer
	RoleUpgrade          Role = 31 // UUPS upgrades
)

// Resolver roles (PermissionedResolverLib).
const (
	ResolverRoleSetAddress     Role = 0
	ResolverRoleSetText        Role = 1
	ResolverRoleSetContenthash Role = 2
	ResolverRoleSetABI         Role = 3
	ResolverRoleSetInterface   Role = 4
	ResolverRoleSetName        Role = 5
	ResolverRoleSetData        Role = 6
	ResolverRoleLink           Role = 7
	ResolverRoleCanName        Role = 30
	ResolverRoleUpgrade        Role = 31
)

// Roles is an immutable set of roles: a 256-bit bitmap with one nybble per
// role. The zero value is the empty set.
type Roles struct{ bits uint256.Int }

// NewRoles builds a set from roles.
func NewRoles(roles ...Role) Roles {
	var r Roles
	for _, role := range roles {
		r.bits.Or(&r.bits, bitOf(role))
	}
	return r
}

func bitOf(role Role) *uint256.Int {
	return new(uint256.Int).Lsh(uint256.NewInt(1), uint(4*role))
}

// RolesFromBig converts a bitmap read from a contract. It rejects bitmaps
// with bits outside the role nybbles' low bit, which the contracts also
// refuse (EACInvalidRoleBitmap).
func RolesFromBig(b *big.Int) (Roles, error) {
	var r Roles
	if b == nil {
		return r, nil
	}
	if b.Sign() < 0 || b.BitLen() > 256 {
		return r, errors.New("role bitmap out of range")
	}
	r.bits.SetFromBig(b)
	if !new(uint256.Int).And(&r.bits, &invalidRoleMask).IsZero() {
		return r, fmt.Errorf("invalid role bitmap 0x%s", b.Text(16))
	}
	return r, nil
}

// mustRoles converts a bitmap, keeping invalid bits rather than failing;
// used for values read back from contracts.
func mustRoles(b *big.Int) Roles {
	var r Roles
	if b != nil {
		r.bits.SetFromBig(b)
	}
	return r
}

// Big returns a fresh copy of the bitmap.
func (r Roles) Big() *big.Int { return r.bits.ToBig() }

// Has reports whether the set contains the role.
func (r Roles) Has(role Role) bool {
	return !new(uint256.Int).And(&r.bits, bitOf(role)).IsZero()
}

// HasAll reports whether the set contains every role of o.
func (r Roles) HasAll(o Roles) bool {
	return new(uint256.Int).And(&r.bits, &o.bits).Eq(&o.bits)
}

// IsZero reports whether the set is empty.
func (r Roles) IsZero() bool { return r.bits.IsZero() }

// With returns the set plus the given roles.
func (r Roles) With(roles ...Role) Roles {
	out := r
	for _, role := range roles {
		out.bits.Or(&out.bits, bitOf(role))
	}
	return out
}

// Union returns the union of the sets.
func (r Roles) Union(o Roles) Roles {
	var out Roles
	out.bits.Or(&r.bits, &o.bits)
	return out
}

// Without returns the set minus the given roles.
func (r Roles) Without(roles ...Role) Roles {
	out := r
	for _, role := range roles {
		out.bits.And(&out.bits, new(uint256.Int).Not(bitOf(role)))
	}
	return out
}

// WithAdmins returns the set plus the admin role of every regular role in it.
func (r Roles) WithAdmins() Roles {
	out := r
	for _, role := range r.List() {
		if !role.IsAdmin() {
			out = out.With(role.Admin())
		}
	}
	return out
}

// List returns the roles in the set, ascending.
func (r Roles) List() []Role {
	var out []Role
	for i := 0; i < 64; i++ {
		if r.Has(Role(i)) {
			out = append(out, Role(i))
		}
	}
	return out
}

// String renders the set with registry role names; use ResolverString for
// resolver roles.
func (r Roles) String() string { return r.format(registryRoleNames) }

// ResolverString renders the set with resolver role names.
func (r Roles) ResolverString() string { return r.format(resolverRoleNames) }

func (r Roles) format(names map[Role]string) string {
	roles := r.List()
	parts := make([]string, 0, len(roles))
	for _, role := range roles {
		name := names[role.Base()]
		if name == "" {
			name = fmt.Sprintf("ROLE_%d", role.Base())
		}
		if role.IsAdmin() {
			name += "_ADMIN"
		}
		parts = append(parts, name)
	}
	if len(parts) == 0 {
		return "none"
	}
	return strings.Join(parts, "|")
}

var registryRoleNames = map[Role]string{
	RoleRegistrar: "REGISTRAR", RoleRegisterReserved: "REGISTER_RESERVED", RoleSetParent: "SET_PARENT",
	RoleUnregister: "UNREGISTER", RoleRenew: "RENEW", RoleSetSubregistry: "SET_SUBREGISTRY",
	RoleSetResolver: "SET_RESOLVER", RoleCanTransfer: "CAN_TRANSFER", RoleWasReserved: "WAS_RESERVED",
	RoleSetURI: "SET_URI", RoleCanName: "CAN_NAME", RoleUpgrade: "UPGRADE",
}

var resolverRoleNames = map[Role]string{
	ResolverRoleSetAddress: "SET_ADDRESS", ResolverRoleSetText: "SET_TEXT", ResolverRoleSetContenthash: "SET_CONTENTHASH",
	ResolverRoleSetABI: "SET_ABI", ResolverRoleSetInterface: "SET_INTERFACE", ResolverRoleSetName: "SET_NAME",
	ResolverRoleSetData: "SET_DATA", ResolverRoleLink: "LINK", ResolverRoleCanName: "CAN_NAME", ResolverRoleUpgrade: "UPGRADE",
}

// Well-known sets.
var (
	// AllRoles is every regular and admin role (EACBaseRolesLib.ALL_ROLES).
	AllRoles Roles
	// AdminRoles is every admin role.
	AdminRoles Roles
	// RegistrationRoles is what ETHRegistrar grants a new .eth owner. It
	// does not include RoleRenew: .eth names renew only through the registrar.
	RegistrationRoles = NewRoles(RoleSetSubregistry, RoleSetSubregistry.Admin(), RoleSetResolver, RoleSetResolver.Admin(), RoleCanTransfer.Admin())
	// ResolverAllRoles is AllRoles, named for readability at resolver call sites.
	ResolverAllRoles Roles

	invalidRoleMask uint256.Int
)

func init() {
	for i := 0; i < 64; i++ {
		AllRoles = AllRoles.With(Role(i))
		if i >= 32 {
			AdminRoles = AdminRoles.With(Role(i))
		}
	}
	ResolverAllRoles = AllRoles
	// Bits that are not the low bit of a nybble are never valid.
	invalidRoleMask.Not(&AllRoles.bits)
}

// RoleCount maps roles to how many accounts hold them (the contracts cap
// this at 15 per role per resource).
type RoleCount map[Role]uint8

// DecodeRoleCounts converts a roleCount() bitmap, where each nybble holds
// the assignee count of its role.
func DecodeRoleCounts(counts *big.Int) RoleCount {
	out := RoleCount{}
	if counts == nil {
		return out
	}
	var v uint256.Int
	v.SetFromBig(counts)
	for i := 0; i < 64; i++ {
		n := new(uint256.Int).Rsh(&v, uint(4*i))
		n.And(n, uint256.NewInt(0xf))
		if c := n.Uint64(); c > 0 {
			out[Role(i)] = uint8(c)
		}
	}
	return out
}

// RootResource is the resource id of contract-wide (root) roles.
var RootResource = new(big.Int)

// SetterKind identifies which PermissionedResolver setter a scope covers.
type SetterKind int

// Setter kinds that accept scoped (per-argument) roles.
const (
	SetterAddress   SetterKind = iota // setAddress(coinType)
	SetterText                        // setText(key)
	SetterData                        // setData(key)
	SetterABI                         // setABI(contentType)
	SetterInterface                   // setInterface(interfaceId)
)

// SetterScope describes a scoped resolver permission: a setter plus the
// argument it is limited to. Contenthash, name and link roles have no
// scope and must be granted on the root resource.
type SetterScope struct {
	Kind        SetterKind
	CoinType    uint64 // SetterAddress
	Key         string // SetterText, SetterData
	ContentType uint64 // SetterABI
	InterfaceID [4]byte
}

// Role returns the role the scope corresponds to.
func (s SetterScope) Role() Role {
	switch s.Kind {
	case SetterAddress:
		return ResolverRoleSetAddress
	case SetterText:
		return ResolverRoleSetText
	case SetterData:
		return ResolverRoleSetData
	case SetterABI:
		return ResolverRoleSetABI
	case SetterInterface:
		return ResolverRoleSetInterface
	default:
		return ResolverRoleSetInterface
	}
}

// Resource returns the EAC resource id of the scope, as
// PermissionedResolverLib derives it: keccak256 of the 32-byte word for
// numeric arguments, of the UTF-8 key for strings, of the 4 bytes for
// interface ids.
func (s SetterScope) Resource() *big.Int {
	var h []byte
	switch s.Kind {
	case SetterAddress:
		h = crypto.Keccak256(common.LeftPadBytes(new(big.Int).SetUint64(s.CoinType).Bytes(), 32))
	case SetterABI:
		h = crypto.Keccak256(common.LeftPadBytes(new(big.Int).SetUint64(s.ContentType).Bytes(), 32))
	case SetterText, SetterData:
		h = crypto.Keccak256([]byte(s.Key))
	case SetterInterface:
		h = crypto.Keccak256(s.InterfaceID[:])
	default:
		h = crypto.Keccak256(s.InterfaceID[:])
	}
	return new(big.Int).SetBytes(h)
}

// EncodeSetter returns the calldata grantSetterRoles expects: the setter
// with a placeholder name (0x00) and an empty value.
func (s SetterScope) EncodeSetter() ([]byte, error) {
	placeholder := []byte{0}
	switch s.Kind {
	case SetterAddress:
		return resolverBinding.TryPackSetAddress(placeholder, new(big.Int).SetUint64(s.CoinType), []byte{})
	case SetterText:
		return resolverBinding.TryPackSetText(placeholder, s.Key, "")
	case SetterData:
		return resolverBinding.TryPackSetData(placeholder, s.Key, []byte{})
	case SetterABI:
		return resolverBinding.TryPackSetABI(placeholder, new(big.Int).SetUint64(s.ContentType), []byte{})
	case SetterInterface:
		return resolverBinding.TryPackSetInterface(placeholder, s.InterfaceID, common.Address{})
	default:
		return resolverBinding.TryPackSetInterface(placeholder, s.InterfaceID, common.Address{})
	}
}

// Grant pairs an account with the roles it receives at proxy initialisation.
type Grant struct {
	Account common.Address
	Roles   Roles
}
