// Copyright 2026 Weald Technology Trading.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0

package v2

import (
	"bytes"
	"compress/zlib"
	"context"
	"errors"
	"fmt"
	"math/big"
	"math/bits"

	ens "github.com/ensdomains/go-ens/v5"
	"github.com/ensdomains/go-ens/v5/contracts/v2/permissionedresolver"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

var resolverBinding = permissionedresolver.NewPermissionedResolver()

// PermissionedResolver is a write and role handle on a v2 PermissionedResolver
// proxy. Record reads go through ens.Client, which resolves through the
// Universal Resolver; this contract has no direct record getters.
//
// Setters take a DNS-encoded name, which the resolver namehashes itself, so
// a single resolver can hold records for any number of names. Writes are
// authorised by roles on the resolver (see Roles), not by name ownership.
type PermissionedResolver struct {
	client *Client
	addr   common.Address
}

func newPermissionedResolver(c *Client, addr common.Address) *PermissionedResolver {
	return &PermissionedResolver{client: c, addr: addr}
}

// Address returns the resolver contract address.
func (r *PermissionedResolver) Address() common.Address { return r.addr }

func (r *PermissionedResolver) view(ctx context.Context, data []byte) ([]byte, error) {
	return r.client.view(ctx, r.addr, data)
}

func (r *PermissionedResolver) call(data []byte) ens.Call { return r.client.call(r.addr, data) }

// Resolve calls resolve(name, data) on the resolver directly, bypassing the
// Universal Resolver. It does not follow CCIP-Read.
func (r *PermissionedResolver) Resolve(ctx context.Context, name string, callData []byte) ([]byte, error) {
	dnsName, err := ens.DNSEncode(name)
	if err != nil {
		return nil, err
	}
	out, err := r.view(ctx, resolverBinding.PackResolve(dnsName, callData))
	if err != nil {
		return nil, err
	}
	return resolverBinding.UnpackResolve(out)
}

// RecordID returns the record a node is linked to (0 means the default).
func (r *PermissionedResolver) RecordID(ctx context.Context, node [32]byte) (*big.Int, error) {
	out, err := r.view(ctx, resolverBinding.PackGetRecordId(node))
	if err != nil {
		return nil, err
	}
	return resolverBinding.UnpackGetRecordId(out)
}

// RecordCount returns how many records the resolver holds.
func (r *PermissionedResolver) RecordCount(ctx context.Context) (*big.Int, error) {
	out, err := r.view(ctx, resolverBinding.PackGetRecordCount())
	if err != nil {
		return nil, err
	}
	return resolverBinding.UnpackGetRecordCount(out)
}

// Setters.

// BuildSetAddress packs setAddress(name, coinType, addressBytes). For EVM
// coin types the address must be empty (to clear) or 20 bytes.
func (r *PermissionedResolver) BuildSetAddress(name string, coinType uint64, address []byte) (ens.Call, error) {
	if ens.IsEVMCoinType(coinType) && len(address) != 0 && len(address) != 20 {
		return ens.Call{}, &InvalidEVMAddressError{Data: address}
	}
	dnsName, err := ens.DNSEncode(name)
	if err != nil {
		return ens.Call{}, err
	}
	if address == nil {
		address = []byte{}
	}
	data, err := resolverBinding.TryPackSetAddress(dnsName, new(big.Int).SetUint64(coinType), address)
	if err != nil {
		return ens.Call{}, err
	}
	return r.call(data), nil
}

// SetAddress sets the address record of a name for a coin type.
func (r *PermissionedResolver) SetAddress(opts *bind.TransactOpts, name string, coinType uint64, address []byte) (*types.Transaction, error) {
	c, err := r.BuildSetAddress(name, coinType, address)
	if err != nil {
		return nil, err
	}
	return r.client.send(opts, c)
}

// SetAddr sets the Ethereum address (coin type 60) of a name.
func (r *PermissionedResolver) SetAddr(opts *bind.TransactOpts, name string, address common.Address) (*types.Transaction, error) {
	return r.SetAddress(opts, name, ens.CoinTypeETH, address.Bytes())
}

// ClearAddress removes the address record of a name for a coin type.
func (r *PermissionedResolver) ClearAddress(opts *bind.TransactOpts, name string, coinType uint64) (*types.Transaction, error) {
	return r.SetAddress(opts, name, coinType, nil)
}

// BuildSetText packs setText(name, key, value).
func (r *PermissionedResolver) BuildSetText(name, key, value string) (ens.Call, error) {
	dnsName, err := ens.DNSEncode(name)
	if err != nil {
		return ens.Call{}, err
	}
	data, err := resolverBinding.TryPackSetText(dnsName, key, value)
	if err != nil {
		return ens.Call{}, err
	}
	return r.call(data), nil
}

// SetText sets a text record.
func (r *PermissionedResolver) SetText(opts *bind.TransactOpts, name, key, value string) (*types.Transaction, error) {
	c, err := r.BuildSetText(name, key, value)
	if err != nil {
		return nil, err
	}
	return r.client.send(opts, c)
}

// BuildSetContenthash packs setContenthash(name, hash). Use
// ens.StringToContenthash to encode an ipfs:// or similar URI.
func (r *PermissionedResolver) BuildSetContenthash(name string, hash []byte) (ens.Call, error) {
	dnsName, err := ens.DNSEncode(name)
	if err != nil {
		return ens.Call{}, err
	}
	if hash == nil {
		hash = []byte{}
	}
	data, err := resolverBinding.TryPackSetContenthash(dnsName, hash)
	if err != nil {
		return ens.Call{}, err
	}
	return r.call(data), nil
}

// SetContenthash sets the contenthash record.
func (r *PermissionedResolver) SetContenthash(opts *bind.TransactOpts, name string, hash []byte) (*types.Transaction, error) {
	c, err := r.BuildSetContenthash(name, hash)
	if err != nil {
		return nil, err
	}
	return r.client.send(opts, c)
}

// BuildSetABI packs setABI(name, contentType, data). contentType must be a
// single bit (ens.ABIContentTypeJSON and friends).
func (r *PermissionedResolver) BuildSetABI(name string, contentType uint64, abiData []byte) (ens.Call, error) {
	if contentType == 0 || bits.OnesCount64(contentType) != 1 {
		return ens.Call{}, &InvalidContentTypeError{ContentType: new(big.Int).SetUint64(contentType)}
	}
	dnsName, err := ens.DNSEncode(name)
	if err != nil {
		return ens.Call{}, err
	}
	if abiData == nil {
		abiData = []byte{}
	}
	data, err := resolverBinding.TryPackSetABI(dnsName, new(big.Int).SetUint64(contentType), abiData)
	if err != nil {
		return ens.Call{}, err
	}
	return r.call(data), nil
}

// SetABI sets the ABI record.
func (r *PermissionedResolver) SetABI(opts *bind.TransactOpts, name string, contentType uint64, abiData []byte) (*types.Transaction, error) {
	c, err := r.BuildSetABI(name, contentType, abiData)
	if err != nil {
		return nil, err
	}
	return r.client.send(opts, c)
}

// SetABIJSON sets a JSON ABI, zlib-compressed when compress is true.
func (r *PermissionedResolver) SetABIJSON(opts *bind.TransactOpts, name, abiJSON string, compress bool) (*types.Transaction, error) {
	if !compress {
		return r.SetABI(opts, name, ens.ABIContentTypeJSON, []byte(abiJSON))
	}
	var buf bytes.Buffer
	w := zlib.NewWriter(&buf)
	if _, err := w.Write([]byte(abiJSON)); err != nil {
		return nil, err
	}
	if err := w.Close(); err != nil {
		return nil, err
	}
	return r.SetABI(opts, name, ens.ABIContentTypeZlibJSON, buf.Bytes())
}

// BuildSetName packs setName(name, primaryName): the reverse record stored
// on <hex>.addr.reverse names.
func (r *PermissionedResolver) BuildSetName(name, primaryName string) (ens.Call, error) {
	dnsName, err := ens.DNSEncode(name)
	if err != nil {
		return ens.Call{}, err
	}
	data, err := resolverBinding.TryPackSetName(dnsName, primaryName)
	if err != nil {
		return ens.Call{}, err
	}
	return r.call(data), nil
}

// SetName sets the name record.
func (r *PermissionedResolver) SetName(opts *bind.TransactOpts, name, primaryName string) (*types.Transaction, error) {
	c, err := r.BuildSetName(name, primaryName)
	if err != nil {
		return nil, err
	}
	return r.client.send(opts, c)
}

// BuildSetInterface packs setInterface(name, interfaceId, implementer).
func (r *PermissionedResolver) BuildSetInterface(name string, interfaceID [4]byte, implementer common.Address) (ens.Call, error) {
	dnsName, err := ens.DNSEncode(name)
	if err != nil {
		return ens.Call{}, err
	}
	data, err := resolverBinding.TryPackSetInterface(dnsName, interfaceID, implementer)
	if err != nil {
		return ens.Call{}, err
	}
	return r.call(data), nil
}

// SetInterface sets an interface implementer record.
func (r *PermissionedResolver) SetInterface(opts *bind.TransactOpts, name string, interfaceID [4]byte, implementer common.Address) (*types.Transaction, error) {
	c, err := r.BuildSetInterface(name, interfaceID, implementer)
	if err != nil {
		return nil, err
	}
	return r.client.send(opts, c)
}

// BuildSetData packs setData(name, key, value) (ENSIP-24).
func (r *PermissionedResolver) BuildSetData(name, key string, value []byte) (ens.Call, error) {
	dnsName, err := ens.DNSEncode(name)
	if err != nil {
		return ens.Call{}, err
	}
	if value == nil {
		value = []byte{}
	}
	data, err := resolverBinding.TryPackSetData(dnsName, key, value)
	if err != nil {
		return ens.Call{}, err
	}
	return r.call(data), nil
}

// SetData sets a data record.
func (r *PermissionedResolver) SetData(opts *bind.TransactOpts, name, key string, value []byte) (*types.Transaction, error) {
	c, err := r.BuildSetData(name, key, value)
	if err != nil {
		return nil, err
	}
	return r.client.send(opts, c)
}

// BuildLinkToRecord packs linkToRecord(name, recordId). Record 0 unlinks
// the name so it falls back to the resolver's default record; this is the
// v2 replacement for clearRecords.
func (r *PermissionedResolver) BuildLinkToRecord(name string, recordID *big.Int) (ens.Call, error) {
	dnsName, err := ens.DNSEncode(name)
	if err != nil {
		return ens.Call{}, err
	}
	if recordID == nil {
		recordID = new(big.Int)
	}
	data, err := resolverBinding.TryPackLinkToRecord(dnsName, recordID)
	if err != nil {
		return ens.Call{}, err
	}
	return r.call(data), nil
}

// LinkToRecord points a name at an existing record. Needs ROLE_LINK.
func (r *PermissionedResolver) LinkToRecord(opts *bind.TransactOpts, name string, recordID *big.Int) (*types.Transaction, error) {
	c, err := r.BuildLinkToRecord(name, recordID)
	if err != nil {
		return nil, err
	}
	return r.client.send(opts, c)
}

// BuildLinkToNode packs linkToNode(name, targetNode): the name shares the
// record of the target node.
func (r *PermissionedResolver) BuildLinkToNode(name string, target [32]byte) (ens.Call, error) {
	dnsName, err := ens.DNSEncode(name)
	if err != nil {
		return ens.Call{}, err
	}
	data, err := resolverBinding.TryPackLinkToNode(dnsName, target)
	if err != nil {
		return ens.Call{}, err
	}
	return r.call(data), nil
}

// LinkToNode makes a name share another name's record. Needs ROLE_LINK.
func (r *PermissionedResolver) LinkToNode(opts *bind.TransactOpts, name string, target [32]byte) (*types.Transaction, error) {
	c, err := r.BuildLinkToNode(name, target)
	if err != nil {
		return nil, err
	}
	return r.client.send(opts, c)
}

// BuildMulticall packs multicall(bytes[]) from calls that all target this
// resolver. The resolver reverts the whole batch on the first failure.
func (r *PermissionedResolver) BuildMulticall(calls []ens.Call) (ens.Call, error) {
	datas := make([][]byte, 0, len(calls))
	for _, c := range calls {
		if c.To != r.addr {
			return ens.Call{}, fmt.Errorf("multicall item targets %s, not resolver %s", c.To.Hex(), r.addr.Hex())
		}
		datas = append(datas, c.Data)
	}
	data, err := resolverBinding.TryPackMulticall(datas)
	if err != nil {
		return ens.Call{}, err
	}
	return r.call(data), nil
}

// Multicall sends several setter calls in one transaction.
func (r *PermissionedResolver) Multicall(opts *bind.TransactOpts, calls []ens.Call) (*types.Transaction, error) {
	c, err := r.BuildMulticall(calls)
	if err != nil {
		return nil, err
	}
	return r.client.send(opts, c)
}

// RecordSet is a batch of records for SetRecords.
type RecordSet struct {
	// ETHAddress, when non-nil, sets the coin type 60 address; a pointer to
	// the zero address clears it.
	ETHAddress *common.Address
	// Coins sets multicoin addresses by coin type; empty values clear.
	Coins map[uint64][]byte
	// Texts sets text records; empty values clear.
	Texts map[string]string
	// Contenthash, when non-nil, sets the contenthash; empty clears.
	Contenthash []byte
	// Name, when non-nil, sets the name record.
	Name *string
}

// BuildSetRecords packs the record set as one setter or a multicall.
func (r *PermissionedResolver) BuildSetRecords(name string, recs RecordSet) (ens.Call, error) {
	var calls []ens.Call
	add := func(c ens.Call, err error) error {
		if err != nil {
			return err
		}
		calls = append(calls, c)
		return nil
	}
	if recs.Contenthash != nil {
		if err := add(r.BuildSetContenthash(name, recs.Contenthash)); err != nil {
			return ens.Call{}, err
		}
	}
	for key, value := range recs.Texts {
		if err := add(r.BuildSetText(name, key, value)); err != nil {
			return ens.Call{}, err
		}
	}
	if recs.ETHAddress != nil {
		addr := recs.ETHAddress.Bytes()
		if *recs.ETHAddress == (common.Address{}) {
			addr = nil
		}
		if err := add(r.BuildSetAddress(name, ens.CoinTypeETH, addr)); err != nil {
			return ens.Call{}, err
		}
	}
	for coin, value := range recs.Coins {
		if err := add(r.BuildSetAddress(name, coin, value)); err != nil {
			return ens.Call{}, err
		}
	}
	if recs.Name != nil {
		if err := add(r.BuildSetName(name, *recs.Name)); err != nil {
			return ens.Call{}, err
		}
	}
	switch len(calls) {
	case 0:
		return ens.Call{}, errors.New("no records to set")
	case 1:
		return calls[0], nil
	default:
		return r.BuildMulticall(calls)
	}
}

// SetRecords sets several records in one transaction.
func (r *PermissionedResolver) SetRecords(opts *bind.TransactOpts, name string, recs RecordSet) (*types.Transaction, error) {
	c, err := r.BuildSetRecords(name, recs)
	if err != nil {
		return nil, err
	}
	return r.client.send(opts, c)
}

// Roles.

// Roles returns the roles an account holds on a resource (RootResource or
// a SetterScope.Resource()).
func (r *PermissionedResolver) Roles(ctx context.Context, resource *big.Int, account common.Address) (Roles, error) {
	out, err := r.view(ctx, resolverBinding.PackRoles(resource, account))
	if err != nil {
		return Roles{}, err
	}
	b, err := resolverBinding.UnpackRoles(out)
	if err != nil {
		return Roles{}, err
	}
	return mustRoles(b), nil
}

// RootRoles returns the resolver-wide roles of an account.
func (r *PermissionedResolver) RootRoles(ctx context.Context, account common.Address) (Roles, error) {
	return r.Roles(ctx, RootResource, account)
}

// HasRoles reports whether an account holds every role on a resource.
func (r *PermissionedResolver) HasRoles(ctx context.Context, resource *big.Int, roles Roles, account common.Address) (bool, error) {
	out, err := r.view(ctx, resolverBinding.PackHasRoles(resource, roles.Big(), account))
	if err != nil {
		return false, err
	}
	return resolverBinding.UnpackHasRoles(out)
}

// HasRootRoles reports whether an account holds every role resolver-wide.
func (r *PermissionedResolver) HasRootRoles(ctx context.Context, roles Roles, account common.Address) (bool, error) {
	out, err := r.view(ctx, resolverBinding.PackHasRootRoles(roles.Big(), account))
	if err != nil {
		return false, err
	}
	return resolverBinding.UnpackHasRootRoles(out)
}

// BuildGrantRootRoles packs grantRootRoles(roles, account).
func (r *PermissionedResolver) BuildGrantRootRoles(roles Roles, account common.Address) ens.Call {
	return r.call(resolverBinding.PackGrantRootRoles(roles.Big(), account))
}

// GrantRootRoles grants resolver-wide roles; the sender needs the admin
// role of each. grantRoles(resource, …) is disabled on this contract.
func (r *PermissionedResolver) GrantRootRoles(opts *bind.TransactOpts, roles Roles, account common.Address) (*types.Transaction, error) {
	return r.client.send(opts, r.BuildGrantRootRoles(roles, account))
}

// BuildRevokeRootRoles packs revokeRootRoles(roles, account).
func (r *PermissionedResolver) BuildRevokeRootRoles(roles Roles, account common.Address) ens.Call {
	return r.call(resolverBinding.PackRevokeRootRoles(roles.Big(), account))
}

// RevokeRootRoles revokes resolver-wide roles.
func (r *PermissionedResolver) RevokeRootRoles(opts *bind.TransactOpts, roles Roles, account common.Address) (*types.Transaction, error) {
	return r.client.send(opts, r.BuildRevokeRootRoles(roles, account))
}

// BuildGrantSetterRoles packs grantSetterRoles(setter, account): the role
// for one setter, limited to one argument (a coin type, a text key, …).
func (r *PermissionedResolver) BuildGrantSetterRoles(scope SetterScope, account common.Address) (ens.Call, error) {
	setter, err := scope.EncodeSetter()
	if err != nil {
		return ens.Call{}, err
	}
	data, err := resolverBinding.TryPackGrantSetterRoles(setter, account)
	if err != nil {
		return ens.Call{}, err
	}
	return r.call(data), nil
}

// GrantSetterRoles grants a scoped setter role.
func (r *PermissionedResolver) GrantSetterRoles(opts *bind.TransactOpts, scope SetterScope, account common.Address) (*types.Transaction, error) {
	c, err := r.BuildGrantSetterRoles(scope, account)
	if err != nil {
		return nil, err
	}
	return r.client.send(opts, c)
}

// BuildRevokeRoles packs revokeRoles(resource, roles, account), which is
// how scoped setter roles are revoked.
func (r *PermissionedResolver) BuildRevokeRoles(resource *big.Int, roles Roles, account common.Address) ens.Call {
	return r.call(resolverBinding.PackRevokeRoles(resource, roles.Big(), account))
}

// RevokeScopedRoles revokes the role of a setter scope.
func (r *PermissionedResolver) RevokeScopedRoles(opts *bind.TransactOpts, scope SetterScope, account common.Address) (*types.Transaction, error) {
	return r.client.send(opts, r.BuildRevokeRoles(scope.Resource(), NewRoles(scope.Role()), account))
}

// EncodePermissionedResolverInit packs initialize(grants, calls) for deploying a
// PermissionedResolver proxy through the factory. calls run inside the
// initializer without role checks.
func EncodePermissionedResolverInit(grants []Grant, calls [][]byte) ([]byte, error) {
	gs := make([]permissionedresolver.Grant, 0, len(grants))
	for _, g := range grants {
		gs = append(gs, permissionedresolver.Grant{Account: g.Account, RoleBitmap: g.Roles.Big()})
	}
	if calls == nil {
		calls = [][]byte{}
	}
	return resolverBinding.TryPackInitialize(gs, calls)
}
