// Copyright 2026 Weald Technology Trading.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0

package v1

import (
	"context"
	"math/big"
	"strings"
	"time"

	ens "github.com/ensdomains/go-ens/v5"
	"github.com/ensdomains/go-ens/v5/contracts/v1/namewrapper"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Fuses is the NameWrapper fuse bitmap of a wrapped name.
type Fuses uint32

// NameWrapper fuses.
const (
	FuseCannotUnwrap          Fuses = 1 << 0
	FuseCannotBurnFuses       Fuses = 1 << 1
	FuseCannotTransfer        Fuses = 1 << 2
	FuseCannotSetResolver     Fuses = 1 << 3
	FuseCannotSetTTL          Fuses = 1 << 4
	FuseCannotCreateSubdomain Fuses = 1 << 5
	FuseCannotApprove         Fuses = 1 << 6
	FuseParentCannotControl   Fuses = 1 << 16
	FuseIsDotETH              Fuses = 1 << 17
	FuseCanExtendExpiry       Fuses = 1 << 18
)

var fuseNames = []struct {
	f    Fuses
	name string
}{
	{FuseCannotUnwrap, "CANNOT_UNWRAP"},
	{FuseCannotBurnFuses, "CANNOT_BURN_FUSES"},
	{FuseCannotTransfer, "CANNOT_TRANSFER"},
	{FuseCannotSetResolver, "CANNOT_SET_RESOLVER"},
	{FuseCannotSetTTL, "CANNOT_SET_TTL"},
	{FuseCannotCreateSubdomain, "CANNOT_CREATE_SUBDOMAIN"},
	{FuseCannotApprove, "CANNOT_APPROVE"},
	{FuseParentCannotControl, "PARENT_CANNOT_CONTROL"},
	{FuseIsDotETH, "IS_DOT_ETH"},
	{FuseCanExtendExpiry, "CAN_EXTEND_EXPIRY"},
}

// Has reports whether every fuse in f2 is burned.
func (f Fuses) Has(f2 Fuses) bool { return f&f2 == f2 }

// Locked reports whether CANNOT_UNWRAP is burned. Locked names migrate to
// v2 through the LockedMigrationController.
func (f Fuses) Locked() bool { return f.Has(FuseCannotUnwrap) }

// Emancipated reports whether PARENT_CANNOT_CONTROL is burned.
func (f Fuses) Emancipated() bool { return f.Has(FuseParentCannotControl) }

// Names lists the burned fuses by name.
func (f Fuses) Names() []string {
	var out []string
	for _, fn := range fuseNames {
		if f.Has(fn.f) {
			out = append(out, fn.name)
		}
	}
	return out
}

func (f Fuses) String() string { return strings.Join(f.Names(), "|") }

// WrapperData is the NameWrapper's record of a wrapped name.
type WrapperData struct {
	Owner  common.Address
	Fuses  Fuses
	Expiry time.Time
}

// NameWrapper is a read handle on the v1 NameWrapper, plus the approval
// setter the v2 migration helper needs.
type NameWrapper struct {
	backend      bind.ContractBackend
	Contract     *namewrapper.Contract
	ContractAddr common.Address
}

// NewNameWrapperAt binds the NameWrapper at an address.
func NewNameWrapperAt(backend bind.ContractBackend, address common.Address) (*NameWrapper, error) {
	contract, err := namewrapper.NewContract(address, backend)
	if err != nil {
		return nil, err
	}
	return &NameWrapper{backend: backend, Contract: contract, ContractAddr: address}, nil
}

func wrapperID(name string) (*big.Int, error) {
	node, err := ens.NameHash(name)
	if err != nil {
		return nil, err
	}
	return new(big.Int).SetBytes(node[:]), nil
}

// OwnerOf returns the wrapped owner of a name, or the zero address when the
// name is not wrapped.
func (w *NameWrapper) OwnerOf(ctx context.Context, name string) (common.Address, error) {
	id, err := wrapperID(name)
	if err != nil {
		return common.Address{}, err
	}
	return w.Contract.OwnerOf(&bind.CallOpts{Context: ctx}, id)
}

// IsWrapped reports whether a name is held by the NameWrapper.
func (w *NameWrapper) IsWrapped(ctx context.Context, name string) (bool, error) {
	owner, err := w.OwnerOf(ctx, name)
	if err != nil {
		return false, err
	}
	return owner != (common.Address{}), nil
}

// Data returns the owner, fuses and expiry of a wrapped name, or nil when
// the name is not wrapped.
func (w *NameWrapper) Data(ctx context.Context, name string) (*WrapperData, error) {
	id, err := wrapperID(name)
	if err != nil {
		return nil, err
	}
	data, err := w.Contract.GetData(&bind.CallOpts{Context: ctx}, id)
	if err != nil {
		return nil, err
	}
	if data.Owner == (common.Address{}) {
		return nil, nil
	}
	return &WrapperData{Owner: data.Owner, Fuses: Fuses(data.Fuses), Expiry: time.Unix(int64(data.Expiry), 0).UTC()}, nil
}

// CanModifyName reports whether an account may modify a wrapped name.
func (w *NameWrapper) CanModifyName(ctx context.Context, name string, account common.Address) (bool, error) {
	node, err := ens.NameHash(name)
	if err != nil {
		return false, err
	}
	return w.Contract.CanModifyName(&bind.CallOpts{Context: ctx}, node, account)
}

// IsApprovedForAll reports whether operator may act on all of owner's
// wrapped names.
func (w *NameWrapper) IsApprovedForAll(ctx context.Context, owner, operator common.Address) (bool, error) {
	return w.Contract.IsApprovedForAll(&bind.CallOpts{Context: ctx}, owner, operator)
}

// SetApprovalForAll grants or revokes an operator over all of the sender's
// wrapped names. The v2 MigrationHelper needs this approval.
func (w *NameWrapper) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return w.Contract.SetApprovalForAll(opts, operator, approved)
}
