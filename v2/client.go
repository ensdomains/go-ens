// Copyright 2026 Weald Technology Trading.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0

// Package v2 is the ENS v2 surface of go-ens: the hierarchical registries,
// the .eth registrar and renewers, the rent price oracle, the
// PermissionedResolver and PublicResolverV2, the VerifiableFactory, the
// UniversalHelper and the reverse-registrar adapters.
//
// Record reads stay on ens.Client, which resolves through the Universal
// Resolver for v1 and v2 names alike. Every write here comes as a pair: a
// Build* method that returns an ens.Call without touching the network, and
// a method that signs and sends it.
package v2

import (
	"context"
	"sync"

	ens "github.com/ensdomains/go-ens/v5"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Client binds the v2 contracts of a chain configuration to a backend.
type Client struct {
	ens *ens.Client
	cfg ens.ChainConfig

	oracleOnce sync.Once
	oracle     *Oracle
	oracleErr  error
}

// NewClient returns a v2 client sharing the backend and configuration of c.
func NewClient(c *ens.Client) *Client {
	return &Client{ens: c, cfg: c.Config()}
}

// ENS returns the underlying root client.
func (c *Client) ENS() *ens.Client { return c.ens }

// Config returns the chain configuration.
func (c *Client) Config() ens.ChainConfig { return c.cfg }

// Backend returns the contract backend.
func (c *Client) Backend() bind.ContractBackend { return c.ens.Backend() }

// RootRegistry returns the root registry.
func (c *Client) RootRegistry() *Registry { return c.Registry(c.cfg.V2.RootRegistry) }

// ETHRegistry returns the .eth registry.
func (c *Client) ETHRegistry() *Registry { return c.Registry(c.cfg.V2.ETHRegistry) }

// Registry binds any PermissionedRegistry (root, .eth, user or wrapper).
func (c *Client) Registry(addr common.Address) *Registry { return newRegistry(c, addr) }

// ETHRegistrar returns the .eth registrar.
func (c *Client) ETHRegistrar() *Registrar { return newRegistrar(c, c.cfg.V2.ETHRegistrar) }

// RenewerV1 returns the renewer for unmigrated v1 names.
func (c *Client) RenewerV1() *RenewerV1 { return newRenewerV1(c, c.cfg.V2.ETHRenewerV1) }

// Oracle returns the rent price oracle the registrar currently uses. The
// address is read from ETHRegistrar.rentPriceOracle() once and cached.
func (c *Client) Oracle(ctx context.Context) (*Oracle, error) {
	c.oracleOnce.Do(func() {
		addr, err := c.ETHRegistrar().OracleAddress(ctx)
		if err != nil {
			c.oracleErr = err
			return
		}
		c.oracle = newOracle(c, addr)
	})
	return c.oracle, c.oracleErr
}

// OracleAt binds a rent price oracle at an address.
func (c *Client) OracleAt(addr common.Address) *Oracle { return newOracle(c, addr) }

// Token binds an ERC-20.
func (c *Client) Token(addr common.Address) *ERC20 { return newERC20(c, addr) }

// Factory returns the VerifiableFactory.
func (c *Client) Factory() *Factory { return newFactory(c, c.cfg.V2.VerifiableFactory) }

// Helper returns the UniversalHelper.
func (c *Client) Helper() *UniversalHelper { return newHelper(c, c.cfg.V2.UniversalHelper) }

// PermissionedResolver binds a PermissionedResolver proxy for writes.
func (c *Client) PermissionedResolver(addr common.Address) *PermissionedResolver {
	return newPermissionedResolver(c, addr)
}

// PublicResolverV2 returns the shared PublicResolverV2.
func (c *Client) PublicResolverV2() *PublicResolverV2 {
	return newPublicResolverV2(c, c.cfg.V2.PublicResolverV2)
}

// ReverseAdapters returns the reverse-registrar adapters.
func (c *Client) ReverseAdapters() *ReverseAdapters { return newReverseAdapters(c) }

// RegistrationOpen reports whether the ETHRegistrar holds the registrar
// role on the .eth registry, that is whether v2 registration has been
// enabled by the ENS DAO (migration phase 6).
func (c *Client) RegistrationOpen(ctx context.Context) (bool, error) {
	return c.ETHRegistry().HasRootRoles(ctx, NewRoles(RoleRegistrar), c.cfg.V2.ETHRegistrar)
}

// view runs an eth_call and decodes reverts into v2 typed errors.
func (c *Client) view(ctx context.Context, to common.Address, data []byte) ([]byte, error) {
	if to == (common.Address{}) {
		return nil, ens.ErrNotConfigured
	}
	out, err := c.ens.Backend().CallContract(ctx, ethereum.CallMsg{To: &to, Data: data}, nil)
	if err != nil {
		return nil, DecodeRevert(err)
	}
	return out, nil
}

// call builds an ens.Call that decodes reverts into v2 typed errors.
func (c *Client) call(to common.Address, data []byte) ens.Call {
	return ens.NewCall(to, data, nil, decoder)
}

// send signs and submits a call.
func (c *Client) send(opts *bind.TransactOpts, call ens.Call) (*types.Transaction, error) {
	if call.To == (common.Address{}) {
		return nil, ens.ErrNotConfigured
	}
	return c.ens.Send(opts, call)
}

// CodeAt reports whether a contract exists at an address.
func (c *Client) hasCode(ctx context.Context, addr common.Address) (bool, error) {
	code, err := c.ens.Backend().CodeAt(ctx, addr, nil)
	if err != nil {
		return false, err
	}
	return len(code) > 0, nil
}

// WaitFunc waits for a transaction and returns its receipt; ens.WaitMinedOK
// is the usual implementation.
type WaitFunc func(ctx context.Context, tx *types.Transaction) (*types.Receipt, error)

// DefaultWait waits for the transaction with the client's backend.
func (c *Client) DefaultWait() WaitFunc {
	return func(ctx context.Context, tx *types.Transaction) (*types.Receipt, error) {
		return ens.WaitMinedOK(ctx, c.ens.Backend(), tx)
	}
}
