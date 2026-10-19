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
	"math/big"

	ens "github.com/ensdomains/go-ens/v5"
	"github.com/ensdomains/go-ens/v5/contracts/v2/publicresolver"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

var publicResolverBinding = publicresolver.NewPublicResolverV2()

// PublicResolverV2 is the shared, node-keyed v2 resolver. It keeps the v1
// PublicResolver setters and getters, and authorises writes through the v1
// NameWrapper's record of the name plus the v2 owner, so it serves wrapped
// names migrated with CANNOT_SET_RESOLVER rather than native v2 names,
// which use a PermissionedResolver.
type PublicResolverV2 struct {
	client *Client
	addr   common.Address
}

func newPublicResolverV2(c *Client, addr common.Address) *PublicResolverV2 {
	return &PublicResolverV2{client: c, addr: addr}
}

// Address returns the resolver contract address.
func (r *PublicResolverV2) Address() common.Address { return r.addr }

func (r *PublicResolverV2) view(ctx context.Context, data []byte) ([]byte, error) {
	return r.client.view(ctx, r.addr, data)
}

func (r *PublicResolverV2) call(data []byte) ens.Call { return r.client.call(r.addr, data) }

func node(name string) ([32]byte, error) { return ens.NameHash(name) }

// CanModifyName reports whether an account may write records for a name.
func (r *PublicResolverV2) CanModifyName(ctx context.Context, name string, account common.Address) (bool, error) {
	n, err := node(name)
	if err != nil {
		return false, err
	}
	out, err := r.view(ctx, publicResolverBinding.PackCanModifyName(n, account))
	if err != nil {
		return false, err
	}
	return publicResolverBinding.UnpackCanModifyName(out)
}

// IsApprovedFor reports whether delegate may write records for owner's name.
func (r *PublicResolverV2) IsApprovedFor(ctx context.Context, owner common.Address, name string, delegate common.Address) (bool, error) {
	n, err := node(name)
	if err != nil {
		return false, err
	}
	out, err := r.view(ctx, publicResolverBinding.PackIsApprovedFor(owner, n, delegate))
	if err != nil {
		return false, err
	}
	return publicResolverBinding.UnpackIsApprovedFor(out)
}

// IsApprovedForAll reports whether operator may write records for all of
// owner's names.
func (r *PublicResolverV2) IsApprovedForAll(ctx context.Context, owner, operator common.Address) (bool, error) {
	out, err := r.view(ctx, publicResolverBinding.PackIsApprovedForAll(owner, operator))
	if err != nil {
		return false, err
	}
	return publicResolverBinding.UnpackIsApprovedForAll(out)
}

// Approve lets delegate write records for one of the sender's names.
func (r *PublicResolverV2) Approve(opts *bind.TransactOpts, name string, delegate common.Address, approved bool) (*types.Transaction, error) {
	n, err := node(name)
	if err != nil {
		return nil, err
	}
	return r.client.send(opts, r.call(publicResolverBinding.PackApprove(n, delegate, approved)))
}

// SetApprovalForAll lets operator write records for all of the sender's names.
func (r *PublicResolverV2) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return r.client.send(opts, r.call(publicResolverBinding.PackSetApprovalForAll(operator, approved)))
}

// Direct reads (no wildcard or CCIP-Read handling).

// Addr returns the Ethereum address record.
func (r *PublicResolverV2) Addr(ctx context.Context, name string) (common.Address, error) {
	n, err := node(name)
	if err != nil {
		return common.Address{}, err
	}
	out, err := r.view(ctx, publicResolverBinding.PackAddr(n))
	if err != nil {
		return common.Address{}, err
	}
	return publicResolverBinding.UnpackAddr(out)
}

// MultiAddr returns the address record of a coin type.
func (r *PublicResolverV2) MultiAddr(ctx context.Context, name string, coinType uint64) ([]byte, error) {
	n, err := node(name)
	if err != nil {
		return nil, err
	}
	out, err := r.view(ctx, publicResolverBinding.PackAddr0(n, new(big.Int).SetUint64(coinType)))
	if err != nil {
		return nil, err
	}
	return publicResolverBinding.UnpackAddr0(out)
}

// Text returns a text record.
func (r *PublicResolverV2) Text(ctx context.Context, name, key string) (string, error) {
	n, err := node(name)
	if err != nil {
		return "", err
	}
	out, err := r.view(ctx, publicResolverBinding.PackText(n, key))
	if err != nil {
		return "", err
	}
	return publicResolverBinding.UnpackText(out)
}

// Contenthash returns the contenthash record.
func (r *PublicResolverV2) Contenthash(ctx context.Context, name string) ([]byte, error) {
	n, err := node(name)
	if err != nil {
		return nil, err
	}
	out, err := r.view(ctx, publicResolverBinding.PackContenthash(n))
	if err != nil {
		return nil, err
	}
	return publicResolverBinding.UnpackContenthash(out)
}

// Pubkey returns the public key record.
func (r *PublicResolverV2) Pubkey(ctx context.Context, name string) ([32]byte, [32]byte, error) {
	var x, y [32]byte
	n, err := node(name)
	if err != nil {
		return x, y, err
	}
	out, err := r.view(ctx, publicResolverBinding.PackPubkey(n))
	if err != nil {
		return x, y, err
	}
	pk, err := publicResolverBinding.UnpackPubkey(out)
	if err != nil {
		return x, y, err
	}
	return pk.X, pk.Y, nil
}

// RecordVersion returns the record version, which clearRecords bumps.
func (r *PublicResolverV2) RecordVersion(ctx context.Context, name string) (uint64, error) {
	n, err := node(name)
	if err != nil {
		return 0, err
	}
	out, err := r.view(ctx, publicResolverBinding.PackRecordVersions(n))
	if err != nil {
		return 0, err
	}
	return publicResolverBinding.UnpackRecordVersions(out)
}

// Node-keyed setters.

// BuildSetAddr packs setAddr(node, address).
func (r *PublicResolverV2) BuildSetAddr(name string, address common.Address) (ens.Call, error) {
	n, err := node(name)
	if err != nil {
		return ens.Call{}, err
	}
	return r.call(publicResolverBinding.PackSetAddr0(n, address)), nil
}

// SetAddr sets the Ethereum address record.
func (r *PublicResolverV2) SetAddr(opts *bind.TransactOpts, name string, address common.Address) (*types.Transaction, error) {
	c, err := r.BuildSetAddr(name, address)
	if err != nil {
		return nil, err
	}
	return r.client.send(opts, c)
}

// BuildSetMultiAddr packs setAddr(node, coinType, address).
func (r *PublicResolverV2) BuildSetMultiAddr(name string, coinType uint64, address []byte) (ens.Call, error) {
	n, err := node(name)
	if err != nil {
		return ens.Call{}, err
	}
	if address == nil {
		address = []byte{}
	}
	data, err := publicResolverBinding.TryPackSetAddr(n, new(big.Int).SetUint64(coinType), address)
	if err != nil {
		return ens.Call{}, err
	}
	return r.call(data), nil
}

// SetMultiAddr sets the address record of a coin type.
func (r *PublicResolverV2) SetMultiAddr(opts *bind.TransactOpts, name string, coinType uint64, address []byte) (*types.Transaction, error) {
	c, err := r.BuildSetMultiAddr(name, coinType, address)
	if err != nil {
		return nil, err
	}
	return r.client.send(opts, c)
}

// BuildSetText packs setText(node, key, value).
func (r *PublicResolverV2) BuildSetText(name, key, value string) (ens.Call, error) {
	n, err := node(name)
	if err != nil {
		return ens.Call{}, err
	}
	data, err := publicResolverBinding.TryPackSetText(n, key, value)
	if err != nil {
		return ens.Call{}, err
	}
	return r.call(data), nil
}

// SetText sets a text record.
func (r *PublicResolverV2) SetText(opts *bind.TransactOpts, name, key, value string) (*types.Transaction, error) {
	c, err := r.BuildSetText(name, key, value)
	if err != nil {
		return nil, err
	}
	return r.client.send(opts, c)
}

// BuildSetContenthash packs setContenthash(node, hash).
func (r *PublicResolverV2) BuildSetContenthash(name string, hash []byte) (ens.Call, error) {
	n, err := node(name)
	if err != nil {
		return ens.Call{}, err
	}
	if hash == nil {
		hash = []byte{}
	}
	data, err := publicResolverBinding.TryPackSetContenthash(n, hash)
	if err != nil {
		return ens.Call{}, err
	}
	return r.call(data), nil
}

// SetContenthash sets the contenthash record.
func (r *PublicResolverV2) SetContenthash(opts *bind.TransactOpts, name string, hash []byte) (*types.Transaction, error) {
	c, err := r.BuildSetContenthash(name, hash)
	if err != nil {
		return nil, err
	}
	return r.client.send(opts, c)
}

// BuildSetABI packs setABI(node, contentType, data).
func (r *PublicResolverV2) BuildSetABI(name string, contentType uint64, abiData []byte) (ens.Call, error) {
	n, err := node(name)
	if err != nil {
		return ens.Call{}, err
	}
	if abiData == nil {
		abiData = []byte{}
	}
	data, err := publicResolverBinding.TryPackSetABI(n, new(big.Int).SetUint64(contentType), abiData)
	if err != nil {
		return ens.Call{}, err
	}
	return r.call(data), nil
}

// SetABI sets the ABI record.
func (r *PublicResolverV2) SetABI(opts *bind.TransactOpts, name string, contentType uint64, abiData []byte) (*types.Transaction, error) {
	c, err := r.BuildSetABI(name, contentType, abiData)
	if err != nil {
		return nil, err
	}
	return r.client.send(opts, c)
}

// SetPubkey sets the public key record.
func (r *PublicResolverV2) SetPubkey(opts *bind.TransactOpts, name string, x, y [32]byte) (*types.Transaction, error) {
	n, err := node(name)
	if err != nil {
		return nil, err
	}
	return r.client.send(opts, r.call(publicResolverBinding.PackSetPubkey(n, x, y)))
}

// SetName sets the name record.
func (r *PublicResolverV2) SetName(opts *bind.TransactOpts, name, primaryName string) (*types.Transaction, error) {
	n, err := node(name)
	if err != nil {
		return nil, err
	}
	data, err := publicResolverBinding.TryPackSetName(n, primaryName)
	if err != nil {
		return nil, err
	}
	return r.client.send(opts, r.call(data))
}

// SetInterface sets an interface implementer record.
func (r *PublicResolverV2) SetInterface(opts *bind.TransactOpts, name string, interfaceID [4]byte, implementer common.Address) (*types.Transaction, error) {
	n, err := node(name)
	if err != nil {
		return nil, err
	}
	return r.client.send(opts, r.call(publicResolverBinding.PackSetInterface(n, interfaceID, implementer)))
}

// SetData sets a data record.
func (r *PublicResolverV2) SetData(opts *bind.TransactOpts, name, key string, value []byte) (*types.Transaction, error) {
	n, err := node(name)
	if err != nil {
		return nil, err
	}
	if value == nil {
		value = []byte{}
	}
	data, err := publicResolverBinding.TryPackSetData(n, key, value)
	if err != nil {
		return nil, err
	}
	return r.client.send(opts, r.call(data))
}

// ClearRecords bumps the record version, invalidating every record of the name.
func (r *PublicResolverV2) ClearRecords(opts *bind.TransactOpts, name string) (*types.Transaction, error) {
	n, err := node(name)
	if err != nil {
		return nil, err
	}
	return r.client.send(opts, r.call(publicResolverBinding.PackClearRecords(n)))
}

// BuildMulticall packs multicallWithNodeCheck(node, calls), which rejects
// any inner call for a different node.
func (r *PublicResolverV2) BuildMulticall(name string, calls []ens.Call) (ens.Call, error) {
	n, err := node(name)
	if err != nil {
		return ens.Call{}, err
	}
	datas := make([][]byte, 0, len(calls))
	for _, c := range calls {
		datas = append(datas, c.Data)
	}
	data, err := publicResolverBinding.TryPackMulticallWithNodeCheck(n, datas)
	if err != nil {
		return ens.Call{}, err
	}
	return r.call(data), nil
}

// Multicall sends several setters for one name in one transaction.
func (r *PublicResolverV2) Multicall(opts *bind.TransactOpts, name string, calls []ens.Call) (*types.Transaction, error) {
	c, err := r.BuildMulticall(name, calls)
	if err != nil {
		return nil, err
	}
	return r.client.send(opts, c)
}
