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

	"github.com/ensdomains/go-ens/v5/contracts/v1/defaultreverseregistrar"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// DefaultReverseRegistrar is the ENSIP-19 default.reverse registrar. A name
// set here is the primary name for every EVM chain that has no
// chain-specific reverse record.
type DefaultReverseRegistrar struct {
	backend      bind.ContractBackend
	Contract     *defaultreverseregistrar.Contract
	ContractAddr common.Address
}

// NewDefaultReverseRegistrarAt binds the registrar at an address.
func NewDefaultReverseRegistrarAt(backend bind.ContractBackend, address common.Address) (*DefaultReverseRegistrar, error) {
	contract, err := defaultreverseregistrar.NewContract(address, backend)
	if err != nil {
		return nil, err
	}
	return &DefaultReverseRegistrar{backend: backend, Contract: contract, ContractAddr: address}, nil
}

// SetName sets the default primary name of the sender.
func (r *DefaultReverseRegistrar) SetName(opts *bind.TransactOpts, name string) (*types.Transaction, error) {
	return r.Contract.SetName(opts, name)
}

// SetNameForAddr sets the default primary name of an address the sender is
// authorised to name (itself, an Ownable it owns, or through a controller).
func (r *DefaultReverseRegistrar) SetNameForAddr(opts *bind.TransactOpts, addr common.Address, name string) (*types.Transaction, error) {
	return r.Contract.SetNameForAddr(opts, addr, name)
}

// SetNameForAddrWithSignature sets the default primary name of an address
// with that address's signature, so a relayer can pay the gas.
func (r *DefaultReverseRegistrar) SetNameForAddrWithSignature(opts *bind.TransactOpts, addr common.Address, signatureExpiry *big.Int, name string, signature []byte) (*types.Transaction, error) {
	return r.Contract.SetNameForAddrWithSignature(opts, addr, signatureExpiry, name, signature)
}

// NameOf returns the default primary name stored for an address, or "".
func (r *DefaultReverseRegistrar) NameOf(ctx context.Context, addr common.Address) (string, error) {
	return r.Contract.NameForAddr(&bind.CallOpts{Context: ctx}, addr)
}

// IsController reports whether an account is an authorised controller of
// the registrar, which is how the v2 adapters are enabled.
func (r *DefaultReverseRegistrar) IsController(ctx context.Context, account common.Address) (bool, error) {
	return r.Contract.Controllers(&bind.CallOpts{Context: ctx}, account)
}
