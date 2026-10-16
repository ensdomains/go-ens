// Copyright 2026 Weald Technology Trading.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0

package v2

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	ens "github.com/ensdomains/go-ens/v5"
	"github.com/ensdomains/go-ens/v5/contracts/v2/permissionedregistry"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

var registryBinding = permissionedregistry.NewPermissionedRegistry()

// Status is a name's registration status in a registry.
type Status uint8

// Registry statuses.
const (
	StatusAvailable Status = iota
	StatusReserved
	StatusRegistered
)

func (s Status) String() string {
	switch s {
	case StatusReserved:
		return "reserved"
	case StatusRegistered:
		return "registered"
	case StatusAvailable:
		return "available"
	default:
		return "available"
	}
}

// State is a registry's record of a name.
type State struct {
	Status Status
	// Expiry is zero when the registry has no record of the name.
	Expiry time.Time
	// LatestOwner is the most recent holder, even if the name expired.
	LatestOwner common.Address
	// TokenID is the current ERC-1155 id. It changes whenever roles on the
	// name change, so fetch it fresh before any transfer.
	TokenID *big.Int
	// Resource is the EnhancedAccessControl resource the name's roles key on.
	Resource *big.Int
}

// Registry is a PermissionedRegistry: the root registry, the .eth registry,
// a user registry or a wrapper registry. Reads are keyed by label; the
// *ByID variants accept a labelhash, token id or resource.
type Registry struct {
	client *Client
	addr   common.Address
}

func newRegistry(c *Client, addr common.Address) *Registry {
	return &Registry{client: c, addr: addr}
}

// Address returns the registry contract address.
func (r *Registry) Address() common.Address { return r.addr }

func (r *Registry) view(ctx context.Context, data []byte) ([]byte, error) {
	return r.client.view(ctx, r.addr, data)
}

func (r *Registry) call(data []byte) ens.Call { return r.client.call(r.addr, data) }

func labelID(label string) (*big.Int, error) {
	return LabelID(label)
}

// State returns the registry's record of a label.
func (r *Registry) State(ctx context.Context, label string) (*State, error) {
	id, err := labelID(label)
	if err != nil {
		return nil, err
	}
	return r.StateByID(ctx, id)
}

// StateByID returns the record of any id (labelhash, token id or resource).
func (r *Registry) StateByID(ctx context.Context, anyID *big.Int) (*State, error) {
	out, err := r.view(ctx, registryBinding.PackGetState(anyID))
	if err != nil {
		return nil, err
	}
	s, err := registryBinding.UnpackGetState(out)
	if err != nil {
		return nil, fmt.Errorf("unpack getState: %w", err)
	}
	st := &State{Status: Status(s.Status), LatestOwner: s.LatestOwner, TokenID: s.TokenId, Resource: s.Resource}
	if s.Expiry != 0 {
		st.Expiry = time.Unix(int64(s.Expiry), 0).UTC()
	}
	return st, nil
}

// Owner returns the owner of a label, or the zero address once it expired.
func (r *Registry) Owner(ctx context.Context, label string) (common.Address, error) {
	id, err := labelID(label)
	if err != nil {
		return common.Address{}, err
	}
	out, err := r.view(ctx, registryBinding.PackGetOwner(id))
	if err != nil {
		return common.Address{}, err
	}
	return registryBinding.UnpackGetOwner(out)
}

// LatestOwner returns the most recent holder of a label, even if expired.
func (r *Registry) LatestOwner(ctx context.Context, label string) (common.Address, error) {
	st, err := r.State(ctx, label)
	if err != nil {
		return common.Address{}, err
	}
	return st.LatestOwner, nil
}

// Expiry returns the expiry of a label; zero when unknown.
func (r *Registry) Expiry(ctx context.Context, label string) (time.Time, error) {
	st, err := r.State(ctx, label)
	if err != nil {
		return time.Time{}, err
	}
	return st.Expiry, nil
}

// TokenID returns the current ERC-1155 id of a label. Never cache it across
// role changes.
func (r *Registry) TokenID(ctx context.Context, label string) (*big.Int, error) {
	id, err := labelID(label)
	if err != nil {
		return nil, err
	}
	out, err := r.view(ctx, registryBinding.PackGetTokenId(id))
	if err != nil {
		return nil, err
	}
	return registryBinding.UnpackGetTokenId(out)
}

// Resource returns the EAC resource id of a label.
func (r *Registry) Resource(ctx context.Context, label string) (*big.Int, error) {
	id, err := labelID(label)
	if err != nil {
		return nil, err
	}
	out, err := r.view(ctx, registryBinding.PackGetResource(id))
	if err != nil {
		return nil, err
	}
	return registryBinding.UnpackGetResource(out)
}

// Resolver returns the resolver set for a label in this registry, or zero.
// It is the parent-level record and does not follow wildcards; use
// ens.Client.WritableResolver for the resolver that serves a name.
func (r *Registry) Resolver(ctx context.Context, label string) (common.Address, error) {
	label, err := ens.Normalize(label)
	if err != nil {
		return common.Address{}, err
	}
	out, err := r.view(ctx, registryBinding.PackGetResolver(label))
	if err != nil {
		return common.Address{}, err
	}
	return registryBinding.UnpackGetResolver(out)
}

// Subregistry returns the child registry of a label, or zero.
func (r *Registry) Subregistry(ctx context.Context, label string) (common.Address, error) {
	label, err := ens.Normalize(label)
	if err != nil {
		return common.Address{}, err
	}
	out, err := r.view(ctx, registryBinding.PackGetSubregistry(label))
	if err != nil {
		return common.Address{}, err
	}
	return registryBinding.UnpackGetSubregistry(out)
}

// Parent returns the registry's canonical parent and the label it holds there.
func (r *Registry) Parent(ctx context.Context) (common.Address, string, error) {
	out, err := r.view(ctx, registryBinding.PackGetParent())
	if err != nil {
		return common.Address{}, "", err
	}
	p, err := registryBinding.UnpackGetParent(out)
	if err != nil {
		return common.Address{}, "", err
	}
	return p.Parent, p.Label, nil
}

// IsEmancipated reports whether no root-level role can reach the names in
// this registry.
func (r *Registry) IsEmancipated(ctx context.Context) (bool, error) {
	out, err := r.view(ctx, registryBinding.PackIsEmancipated())
	if err != nil {
		return false, err
	}
	return registryBinding.UnpackIsEmancipated(out)
}

// OwnerOf returns the holder of an exact token id, or zero when the id is
// stale or the name expired.
func (r *Registry) OwnerOf(ctx context.Context, tokenID *big.Int) (common.Address, error) {
	out, err := r.view(ctx, registryBinding.PackOwnerOf(tokenID))
	if err != nil {
		return common.Address{}, err
	}
	return registryBinding.UnpackOwnerOf(out)
}

// IsApprovedForAll reports whether operator may act on all of owner's names.
func (r *Registry) IsApprovedForAll(ctx context.Context, owner, operator common.Address) (bool, error) {
	out, err := r.view(ctx, registryBinding.PackIsApprovedForAll(owner, operator))
	if err != nil {
		return false, err
	}
	return registryBinding.UnpackIsApprovedForAll(out)
}

// Roles returns the roles an account holds on a label (including those
// inherited from the root resource and from operator approval).
func (r *Registry) Roles(ctx context.Context, label string, account common.Address) (Roles, error) {
	id, err := labelID(label)
	if err != nil {
		return Roles{}, err
	}
	out, err := r.view(ctx, registryBinding.PackRoles(id, account))
	if err != nil {
		return Roles{}, err
	}
	b, err := registryBinding.UnpackRoles(out)
	if err != nil {
		return Roles{}, err
	}
	return mustRoles(b), nil
}

// RootRoles returns the registry-wide roles an account holds.
func (r *Registry) RootRoles(ctx context.Context, account common.Address) (Roles, error) {
	out, err := r.view(ctx, registryBinding.PackRoles(RootResource, account))
	if err != nil {
		return Roles{}, err
	}
	b, err := registryBinding.UnpackRoles(out)
	if err != nil {
		return Roles{}, err
	}
	return mustRoles(b), nil
}

// HasRoles reports whether an account holds every given role on a label.
func (r *Registry) HasRoles(ctx context.Context, label string, roles Roles, account common.Address) (bool, error) {
	id, err := labelID(label)
	if err != nil {
		return false, err
	}
	out, err := r.view(ctx, registryBinding.PackHasRoles(id, roles.Big(), account))
	if err != nil {
		return false, err
	}
	return registryBinding.UnpackHasRoles(out)
}

// HasRootRoles reports whether an account holds every given role registry-wide.
func (r *Registry) HasRootRoles(ctx context.Context, roles Roles, account common.Address) (bool, error) {
	out, err := r.view(ctx, registryBinding.PackHasRootRoles(roles.Big(), account))
	if err != nil {
		return false, err
	}
	return registryBinding.UnpackHasRootRoles(out)
}

// RoleCounts returns how many accounts hold each role on a label.
func (r *Registry) RoleCounts(ctx context.Context, label string) (RoleCount, error) {
	id, err := labelID(label)
	if err != nil {
		return nil, err
	}
	out, err := r.view(ctx, registryBinding.PackRoleCount(id))
	if err != nil {
		return nil, err
	}
	counts, err := registryBinding.UnpackRoleCount(out)
	if err != nil {
		return nil, err
	}
	return DecodeRoleCounts(counts), nil
}

// IsOnlyAssignee reports whether account is the only holder of every given
// role on the label.
func (r *Registry) IsOnlyAssignee(ctx context.Context, label string, roles Roles, account common.Address) (bool, error) {
	id, err := labelID(label)
	if err != nil {
		return false, err
	}
	out, err := r.view(ctx, registryBinding.PackIsOnlyAssignee(id, roles.Big(), account))
	if err != nil {
		return false, err
	}
	return registryBinding.UnpackIsOnlyAssignee(out)
}

// Writes.

// RegisterParams are the inputs of IStandardRegistry.register: creating a
// name directly in a registry. On user registries this needs the registrar
// role; .eth names are registered through the Registrar instead.
type RegisterParams struct {
	Label       string
	Owner       common.Address // zero with empty Roles creates a reservation
	Subregistry common.Address
	Resolver    common.Address
	Roles       Roles
	// Expiry zero means MaxExpiry (never expires).
	Expiry time.Time
}

func expiryArg(t time.Time) uint64 {
	if t.IsZero() {
		return MaxExpiry
	}
	return uint64(t.Unix())
}

// BuildRegister packs register(label, owner, subregistry, resolver, roles, expiry).
func (r *Registry) BuildRegister(p RegisterParams) (ens.Call, error) {
	label, err := ens.Normalize(p.Label)
	if err != nil {
		return ens.Call{}, err
	}
	data, err := registryBinding.TryPackRegister(label, p.Owner, p.Subregistry, p.Resolver, p.Roles.Big(), expiryArg(p.Expiry))
	if err != nil {
		return ens.Call{}, err
	}
	return r.call(data), nil
}

// Register sends register(...).
func (r *Registry) Register(opts *bind.TransactOpts, p RegisterParams) (*types.Transaction, error) {
	c, err := r.BuildRegister(p)
	if err != nil {
		return nil, err
	}
	return r.client.send(opts, c)
}

// BuildSetResolver packs setResolver(labelhash, resolver).
func (r *Registry) BuildSetResolver(label string, resolver common.Address) (ens.Call, error) {
	id, err := labelID(label)
	if err != nil {
		return ens.Call{}, err
	}
	return r.call(registryBinding.PackSetResolver(id, resolver)), nil
}

// SetResolver sets the resolver of a label. Needs ROLE_SET_RESOLVER.
func (r *Registry) SetResolver(opts *bind.TransactOpts, label string, resolver common.Address) (*types.Transaction, error) {
	c, err := r.BuildSetResolver(label, resolver)
	if err != nil {
		return nil, err
	}
	return r.client.send(opts, c)
}

// BuildSetSubregistry packs setSubregistry(labelhash, registry).
func (r *Registry) BuildSetSubregistry(label string, sub common.Address) (ens.Call, error) {
	id, err := labelID(label)
	if err != nil {
		return ens.Call{}, err
	}
	return r.call(registryBinding.PackSetSubregistry(id, sub)), nil
}

// SetSubregistry sets the child registry of a label. Needs ROLE_SET_SUBREGISTRY.
func (r *Registry) SetSubregistry(opts *bind.TransactOpts, label string, sub common.Address) (*types.Transaction, error) {
	c, err := r.BuildSetSubregistry(label, sub)
	if err != nil {
		return nil, err
	}
	return r.client.send(opts, c)
}

// BuildSetParent packs setParent(parent, label).
func (r *Registry) BuildSetParent(parent common.Address, label string) (ens.Call, error) {
	label, err := ens.Normalize(label)
	if err != nil {
		return ens.Call{}, err
	}
	return r.call(registryBinding.PackSetParent(parent, label)), nil
}

// SetParent records the registry's canonical parent. Needs root ROLE_SET_PARENT.
func (r *Registry) SetParent(opts *bind.TransactOpts, parent common.Address, label string) (*types.Transaction, error) {
	c, err := r.BuildSetParent(parent, label)
	if err != nil {
		return nil, err
	}
	return r.client.send(opts, c)
}

// BuildRenew packs renew(labelhash, newExpiry). Expiry is absolute.
func (r *Registry) BuildRenew(label string, newExpiry time.Time) (ens.Call, error) {
	id, err := labelID(label)
	if err != nil {
		return ens.Call{}, err
	}
	return r.call(registryBinding.PackRenew(id, expiryArg(newExpiry))), nil
}

// Renew sets a new absolute expiry. Needs ROLE_RENEW, which the .eth
// registrar does not grant to owners: renew .eth names with Client.Renew.
func (r *Registry) Renew(opts *bind.TransactOpts, label string, newExpiry time.Time) (*types.Transaction, error) {
	c, err := r.BuildRenew(label, newExpiry)
	if err != nil {
		return nil, err
	}
	return r.client.send(opts, c)
}

// BuildUnregister packs unregister(labelhash).
func (r *Registry) BuildUnregister(label string) (ens.Call, error) {
	id, err := labelID(label)
	if err != nil {
		return ens.Call{}, err
	}
	return r.call(registryBinding.PackUnregister(id)), nil
}

// Unregister burns a name. Needs ROLE_UNREGISTER.
func (r *Registry) Unregister(opts *bind.TransactOpts, label string) (*types.Transaction, error) {
	c, err := r.BuildUnregister(label)
	if err != nil {
		return nil, err
	}
	return r.client.send(opts, c)
}

// BuildGrantRoles packs grantRoles(labelhash, roles, account). Granting
// regenerates the token id.
func (r *Registry) BuildGrantRoles(label string, roles Roles, account common.Address) (ens.Call, error) {
	id, err := labelID(label)
	if err != nil {
		return ens.Call{}, err
	}
	return r.call(registryBinding.PackGrantRoles(id, roles.Big(), account)), nil
}

// GrantRoles grants roles on a label. The caller must hold the admin role
// of each granted role; only the name owner can hold admin roles.
func (r *Registry) GrantRoles(opts *bind.TransactOpts, label string, roles Roles, account common.Address) (*types.Transaction, error) {
	c, err := r.BuildGrantRoles(label, roles, account)
	if err != nil {
		return nil, err
	}
	return r.client.send(opts, c)
}

// BuildRevokeRoles packs revokeRoles(labelhash, roles, account).
func (r *Registry) BuildRevokeRoles(label string, roles Roles, account common.Address) (ens.Call, error) {
	id, err := labelID(label)
	if err != nil {
		return ens.Call{}, err
	}
	return r.call(registryBinding.PackRevokeRoles(id, roles.Big(), account)), nil
}

// RevokeRoles revokes roles on a label.
func (r *Registry) RevokeRoles(opts *bind.TransactOpts, label string, roles Roles, account common.Address) (*types.Transaction, error) {
	c, err := r.BuildRevokeRoles(label, roles, account)
	if err != nil {
		return nil, err
	}
	return r.client.send(opts, c)
}

// BuildGrantRootRoles packs grantRootRoles(roles, account).
func (r *Registry) BuildGrantRootRoles(roles Roles, account common.Address) ens.Call {
	return r.call(registryBinding.PackGrantRootRoles(roles.Big(), account))
}

// GrantRootRoles grants registry-wide roles.
func (r *Registry) GrantRootRoles(opts *bind.TransactOpts, roles Roles, account common.Address) (*types.Transaction, error) {
	return r.client.send(opts, r.BuildGrantRootRoles(roles, account))
}

// BuildRevokeRootRoles packs revokeRootRoles(roles, account).
func (r *Registry) BuildRevokeRootRoles(roles Roles, account common.Address) ens.Call {
	return r.call(registryBinding.PackRevokeRootRoles(roles.Big(), account))
}

// RevokeRootRoles revokes registry-wide roles.
func (r *Registry) RevokeRootRoles(opts *bind.TransactOpts, roles Roles, account common.Address) (*types.Transaction, error) {
	return r.client.send(opts, r.BuildRevokeRootRoles(roles, account))
}

// BuildSetApprovalForAll packs setApprovalForAll(operator, approved).
func (r *Registry) BuildSetApprovalForAll(operator common.Address, approved bool) ens.Call {
	return r.call(registryBinding.PackSetApprovalForAll(operator, approved))
}

// SetApprovalForAll grants or revokes an operator over all of the sender's
// names in this registry. Operators inherit the owner's roles.
func (r *Registry) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return r.client.send(opts, r.BuildSetApprovalForAll(operator, approved))
}

// Transfers.

// TransferCheck is the result of a pre-flight for a safe transfer.
type TransferCheck struct {
	TokenID        *big.Int
	Owner          common.Address
	Emancipated    bool // registry.isEmancipated()
	OnlyAssignee   bool // nobody else holds roles on the name
	CanTransfer    bool // owner holds ROLE_CAN_TRANSFER_ADMIN
	SafeTransferOK bool // all of the above
}

// CheckTransfer runs the checks safeTransferFrom will apply.
func (r *Registry) CheckTransfer(ctx context.Context, label string) (*TransferCheck, error) {
	st, err := r.State(ctx, label)
	if err != nil {
		return nil, err
	}
	owner, err := r.Owner(ctx, label)
	if err != nil {
		return nil, err
	}
	chk := &TransferCheck{TokenID: st.TokenID, Owner: owner}
	if chk.Emancipated, err = r.IsEmancipated(ctx); err != nil {
		return nil, err
	}
	if chk.OnlyAssignee, err = r.IsOnlyAssignee(ctx, label, AllRoles, owner); err != nil {
		return nil, err
	}
	roles, err := r.Roles(ctx, label, owner)
	if err != nil {
		return nil, err
	}
	chk.CanTransfer = roles.Has(RoleCanTransfer.Admin())
	chk.SafeTransferOK = chk.Emancipated && chk.OnlyAssignee && chk.CanTransfer
	return chk, nil
}

// BuildTransfer packs safeTransferFrom(from, to, tokenID, 1, data) for a
// token id the caller has just fetched.
func (r *Registry) BuildTransfer(from, to common.Address, tokenID *big.Int, data []byte) ens.Call {
	return r.call(registryBinding.PackSafeTransferFrom(from, to, tokenID, big.NewInt(1), data))
}

// Transfer moves a name to another account with safeTransferFrom. It fetches
// the current token id first and runs CheckTransfer, returning
// *TransferDisallowedError or *TransferUnsafeError before spending gas.
func (r *Registry) Transfer(ctx context.Context, opts *bind.TransactOpts, label string, to common.Address, data []byte) (*types.Transaction, error) {
	chk, err := r.CheckTransfer(ctx, label)
	if err != nil {
		return nil, err
	}
	if !chk.CanTransfer {
		return nil, &TransferDisallowedError{TokenID: chk.TokenID, From: chk.Owner}
	}
	if !chk.Emancipated {
		return nil, &TransferUnsafeError{Reason: TransferUnsafeNotEmancipated, TokenID: chk.TokenID, From: chk.Owner}
	}
	if !chk.OnlyAssignee {
		return nil, &TransferUnsafeError{Reason: TransferUnsafeMultipleAssignees, TokenID: chk.TokenID, From: chk.Owner}
	}
	if data == nil {
		data = []byte{}
	}
	return r.client.send(opts, r.BuildTransfer(chk.Owner, to, chk.TokenID, data))
}

// BuildUnsafeTransfer packs unsafeTransfer(to, tokenID, data).
func (r *Registry) BuildUnsafeTransfer(to common.Address, tokenID *big.Int, data []byte) ens.Call {
	return r.call(registryBinding.PackUnsafeTransfer(to, tokenID, data))
}

// UnsafeTransfer moves a name without the emancipation and sole-assignee
// checks. Roles move with the token, but other accounts keep any roles they
// were granted on the name.
func (r *Registry) UnsafeTransfer(ctx context.Context, opts *bind.TransactOpts, label string, to common.Address, data []byte) (*types.Transaction, error) {
	tokenID, err := r.TokenID(ctx, label)
	if err != nil {
		return nil, err
	}
	if data == nil {
		data = []byte{}
	}
	return r.client.send(opts, r.BuildUnsafeTransfer(to, tokenID, data))
}

// Events.

// LabelRegistered is the registry's registration event.
type LabelRegistered struct {
	TokenID   *big.Int
	LabelHash [32]byte
	Label     string
	Owner     common.Address
	Expiry    time.Time
}

// ParseRegistered finds the LabelRegistered event of this registry in a receipt.
func (r *Registry) ParseRegistered(receipt *types.Receipt) (*LabelRegistered, error) {
	for _, log := range receipt.Logs {
		if log.Address != r.addr {
			continue
		}
		ev, err := registryBinding.UnpackLabelRegisteredEvent(log)
		if err != nil {
			continue
		}
		return &LabelRegistered{TokenID: ev.TokenId, LabelHash: ev.LabelHash, Label: ev.Label, Owner: ev.Owner, Expiry: time.Unix(int64(ev.Expiry), 0).UTC()}, nil
	}
	return nil, errors.New("no LabelRegistered event in receipt")
}

// TokenRegenerated records a token id change.
type TokenRegenerated struct {
	OldTokenID, NewTokenID *big.Int
}

// ParseTokenRegenerated finds every TokenRegenerated event of this registry
// in a receipt.
func (r *Registry) ParseTokenRegenerated(receipt *types.Receipt) []TokenRegenerated {
	var out []TokenRegenerated
	for _, log := range receipt.Logs {
		if log.Address != r.addr {
			continue
		}
		ev, err := registryBinding.UnpackTokenRegeneratedEvent(log)
		if err != nil {
			continue
		}
		out = append(out, TokenRegenerated{OldTokenID: ev.OldTokenId, NewTokenID: ev.NewTokenId})
	}
	return out
}
