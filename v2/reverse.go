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
	"fmt"

	ens "github.com/ensdomains/go-ens/v5"
	"github.com/ensdomains/go-ens/v5/contracts/v2/defaultreverseregistraradapter"
	"github.com/ensdomains/go-ens/v5/contracts/v2/reverseregistraradapter"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

var (
	reverseAdapterBinding        = reverseregistraradapter.NewReverseRegistrarAdapter()
	defaultReverseAdapterBinding = defaultreverseregistraradapter.NewDefaultReverseRegistrarAdapter()
	controllersSelector          = ens.Selector("controllers(address)")
)

// ReverseAdapters are the v2 adapters in front of the v1 reverse registrars.
// Primary names are still stored by the v1 addr.reverse and default.reverse
// registrars; the adapters add authorisation for contracts (Ownable owners
// and IContractNamer) and let a PermissionedResolver hold the reverse record.
// An externally owned account can keep using v1.ReverseRegistrar.SetName.
type ReverseAdapters struct {
	client *Client
}

func newReverseAdapters(c *Client) *ReverseAdapters { return &ReverseAdapters{client: c} }

// BuildClaim packs ReverseRegistrarAdapter.claim(account, resolver): the
// sender becomes the owner of <account>.addr.reverse with the given resolver.
func (a *ReverseAdapters) BuildClaim(account, resolver common.Address) ens.Call {
	return a.client.call(a.client.cfg.V2.ReverseRegistrarAdapter, reverseAdapterBinding.PackClaim(account, resolver))
}

// Claim claims the reverse node of an account the sender may name: itself,
// an Ownable it owns, or a contract that names it through IContractNamer.
func (a *ReverseAdapters) Claim(opts *bind.TransactOpts, account, resolver common.Address) (*types.Transaction, error) {
	return a.client.send(opts, a.BuildClaim(account, resolver))
}

// BuildSetDefaultName packs DefaultReverseRegistrarAdapter.setName(account, name).
func (a *ReverseAdapters) BuildSetDefaultName(account common.Address, name string) (ens.Call, error) {
	data, err := defaultReverseAdapterBinding.TryPackSetName(account, name)
	if err != nil {
		return ens.Call{}, err
	}
	return a.client.call(a.client.cfg.V2.DefaultReverseRegistrarAdapter, data), nil
}

// SetDefaultName sets the ENSIP-19 default primary name of an account the
// sender may name.
func (a *ReverseAdapters) SetDefaultName(opts *bind.TransactOpts, account common.Address, name string) (*types.Transaction, error) {
	c, err := a.BuildSetDefaultName(account, name)
	if err != nil {
		return nil, err
	}
	return a.client.send(opts, c)
}

// IsEnabled reports whether the v1 reverse registrars have authorised the
// adapters as controllers. This requires the ENS DAO's migration phase-1
// transactions and is false on mainnet until they execute.
func (a *ReverseAdapters) IsEnabled(ctx context.Context) (bool, bool, error) {
	claim, err := a.isController(ctx, a.client.cfg.V1.ReverseRegistrar, a.client.cfg.V2.ReverseRegistrarAdapter)
	if err != nil {
		return false, false, err
	}
	defaultName, err := a.isController(ctx, a.client.cfg.V1.DefaultReverseRegistrar, a.client.cfg.V2.DefaultReverseRegistrarAdapter)
	return claim, defaultName, err
}

func (a *ReverseAdapters) isController(ctx context.Context, registrar, adapter common.Address) (bool, error) {
	if registrar == (common.Address{}) || adapter == (common.Address{}) {
		return false, nil
	}
	data := make([]byte, 4+32)
	copy(data[:4], controllersSelector[:])
	copy(data[4+12:], adapter.Bytes())
	out, err := a.client.view(ctx, registrar, data)
	if err != nil {
		return false, err
	}
	return len(out) == 32 && out[31] == 1, nil
}

// SetPrimaryNameWithPermissionedResolver claims the reverse node of account
// with resolver, then stores the primary name in that resolver's
// <account>.addr.reverse record, waiting for each transaction. The sender
// must be allowed to name account and hold ROLE_SET_NAME on the resolver.
// It verifies afterwards that the name resolves forward to account, since
// the Universal Resolver rejects mismatched reverse records.
func (c *Client) SetPrimaryNameWithPermissionedResolver(ctx context.Context, opts *bind.TransactOpts, account, resolver common.Address, name string, wait WaitFunc) ([]*types.Transaction, error) {
	if wait == nil {
		wait = c.DefaultWait()
	}
	forward, err := c.ens.ResolveAddress(ctx, name)
	if err != nil {
		return nil, fmt.Errorf("forward resolution of %q: %w", name, err)
	}
	if forward != account {
		return nil, fmt.Errorf("%q resolves to %s, not %s: the reverse record would be rejected", name, forward.Hex(), account.Hex())
	}
	var txs []*types.Transaction
	tx, err := c.ReverseAdapters().Claim(opts, account, resolver)
	if err != nil {
		return nil, err
	}
	txs = append(txs, tx)
	if _, err := wait(ctx, tx); err != nil {
		return txs, err
	}
	tx, err = c.PermissionedResolver(resolver).SetName(opts, ens.ReverseName(account.Bytes(), ens.CoinTypeETH), name)
	if err != nil {
		return txs, err
	}
	txs = append(txs, tx)
	if _, err := wait(ctx, tx); err != nil {
		return txs, err
	}
	return txs, nil
}
