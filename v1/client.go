// Copyright 2026 Weald Technology Trading.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0

// Package v1 is the ENS v1 surface of go-ens: the registry, the .eth
// BaseRegistrar and controller, the legacy PublicResolver, the reverse
// registrars and the DNS, auction and deed helpers.
//
// Reads of unmigrated names and primary-name writes keep working under
// ENS v2, which stores primary names in the v1 reverse registrars. v1
// registration and renewal stop working once the ENS DAO disables the v1
// controllers; the v2 package is the replacement. Record reads should go
// through ens.Client, which follows wildcards and CCIP-Read for every name.
package v1

import (
	"context"

	ens "github.com/ensdomains/go-ens/v5"
	"github.com/ensdomains/go-ens/v5/internal/chainid"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// Client binds the v1 contracts of a chain configuration to a backend.
type Client struct {
	ens *ens.Client
	cfg ens.ChainConfig
}

// NewClient returns a v1 client sharing the backend and configuration of c.
func NewClient(c *ens.Client) *Client {
	return &Client{ens: c, cfg: c.Config()}
}

// ENS returns the underlying root client.
func (c *Client) ENS() *ens.Client { return c.ens }

// Config returns the chain configuration.
func (c *Client) Config() ens.ChainConfig { return c.cfg }

// Registry returns the v1 ENSRegistry.
func (c *Client) Registry() (*Registry, error) {
	return NewRegistryAt(c.ens.Backend(), c.cfg.V1.Registry)
}

// BaseRegistrar returns the .eth BaseRegistrar.
func (c *Client) BaseRegistrar() (*BaseRegistrar, error) {
	return newBaseRegistrarAt(c.ens.Backend(), "eth", c.cfg.V1.BaseRegistrar)
}

// Controller returns the v1 .eth registrar controller. Its registration and
// renewal functions are disabled by the ENS DAO at migration phase 3.
func (c *Client) Controller() (*ETHController, error) {
	if c.cfg.V1.ETHRegistrarController == (common.Address{}) {
		return nil, ens.ErrNotConfigured
	}
	return NewETHControllerAt(c.ens.Backend(), "eth", c.cfg.V1.ETHRegistrarController)
}

// ReverseRegistrar returns the addr.reverse registrar.
func (c *Client) ReverseRegistrar() (*ReverseRegistrar, error) {
	if c.cfg.V1.ReverseRegistrar == (common.Address{}) {
		return nil, ens.ErrNotConfigured
	}
	return NewReverseRegistrarAt(c.ens.Backend(), c.cfg.V1.ReverseRegistrar)
}

// DefaultReverseRegistrar returns the ENSIP-19 default.reverse registrar.
func (c *Client) DefaultReverseRegistrar() (*DefaultReverseRegistrar, error) {
	if c.cfg.V1.DefaultReverseRegistrar == (common.Address{}) {
		return nil, ens.ErrNotConfigured
	}
	return NewDefaultReverseRegistrarAt(c.ens.Backend(), c.cfg.V1.DefaultReverseRegistrar)
}

// NameWrapper returns the v1 NameWrapper.
func (c *Client) NameWrapper() (*NameWrapper, error) {
	if c.cfg.V1.NameWrapper == (common.Address{}) {
		return nil, ens.ErrNotConfigured
	}
	return NewNameWrapperAt(c.ens.Backend(), c.cfg.V1.NameWrapper)
}

// PublicResolver returns the v1 PublicResolver bound to a name. Writes
// through it only affect names whose registry resolver is that contract.
func (c *Client) PublicResolver(name string) (*Resolver, error) {
	if c.cfg.V1.PublicResolver == (common.Address{}) {
		return nil, ens.ErrNotConfigured
	}
	return NewResolverAt(c.ens.Backend(), name, c.cfg.V1.PublicResolver)
}

// Name binds a .eth name to the v1 registrar and controller.
//
// Deprecated: see NewName.
func (c *Client) Name(name string) (*Name, error) {
	return NewName(c.ens.Backend(), name)
}

// configFor returns the baked-in configuration for the backend's chain, if
// the backend exposes a chain ID with one.
func configFor(backend bind.ContractBackend) (ens.ChainConfig, bool) {
	id, ok, err := chainid.Probe(context.Background(), backend)
	if err != nil || !ok || id == nil {
		return ens.ChainConfig{}, false
	}
	return ens.ConfigForChain(id.Uint64())
}
