// Copyright 2026 Weald Technology Trading.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0

package ens

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/ensdomains/go-ens/v5/ccipread"
	"github.com/ensdomains/go-ens/v5/internal/chainid"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// Client is the entry point of the library. It binds a contract backend to
// the configuration of the chain it is connected to and exposes the
// version-agnostic read surface: resolution through the Universal Resolver,
// reverse resolution, record reads, resolver and owner discovery and
// migration status. The v1 and v2 packages build their clients on top of it.
type Client struct {
	backend  bind.ContractBackend
	cfg      ChainConfig
	ur       *UniversalResolver
	http     *http.Client
	maxHops  int
	gateways []string

	graceOnce sync.Once
	grace     time.Duration
	graceErr  error
}

// Option configures NewClient.
type Option func(*clientOptions)

type clientOptions struct {
	cfg       *ChainConfig
	chainID   *uint64
	ur        *common.Address
	http      *http.Client
	maxHops   int
	gateways  []string
	skipProbe bool
}

// WithChainConfig supplies a full configuration and skips the chain-ID probe.
// Use it for chains without a baked-in table, or to override addresses.
func WithChainConfig(cfg ChainConfig) Option {
	return func(o *clientOptions) { o.cfg = &cfg; o.skipProbe = true }
}

// WithChainID selects a baked-in configuration without probing the backend.
func WithChainID(chainID uint64) Option {
	return func(o *clientOptions) { o.chainID = &chainID; o.skipProbe = true }
}

// WithUniversalResolver overrides the Universal Resolver address the client
// resolves through, for example cfg.V2.UniversalResolverV2 to exercise v2
// resolution before the canonical proxy is repointed.
func WithUniversalResolver(addr common.Address) Option {
	return func(o *clientOptions) { o.ur = &addr }
}

// WithHTTPClient sets the HTTP client used for ERC-3668 CCIP-Read gateway
// requests.
func WithHTTPClient(c *http.Client) Option {
	return func(o *clientOptions) { o.http = c }
}

// WithMaxHops caps the number of OffchainLookup redirects a single
// resolution may chain.
func WithMaxHops(n int) Option {
	return func(o *clientOptions) { o.maxHops = n }
}

// WithGatewayURLs overrides the batch gateway URLs the Universal Resolver
// uses, routing reads through resolveWithGateways and reverseWithGateways.
func WithGatewayURLs(urls []string) Option {
	return func(o *clientOptions) { o.gateways = urls }
}

// NewClient creates a client for the chain reached through backend. Unless
// WithChainConfig or WithChainID is given, it probes the backend's chain ID
// (honouring ctx) and selects the baked-in configuration, returning
// *UnknownChainError when there is none.
func NewClient(ctx context.Context, backend bind.ContractBackend, opts ...Option) (*Client, error) {
	var o clientOptions
	for _, opt := range opts {
		opt(&o)
	}
	var cfg ChainConfig
	switch {
	case o.cfg != nil:
		cfg = *o.cfg
	case o.chainID != nil:
		c, ok := ConfigForChain(*o.chainID)
		if !ok {
			return nil, &UnknownChainError{ChainID: *o.chainID}
		}
		cfg = c
	default:
		id, ok, err := chainid.Probe(ctx, backend)
		if err != nil {
			return nil, err
		}
		if !ok || id == nil {
			// Backends without a chain ID (mocks) default to mainnet addresses.
			cfg = MainnetConfig()
			break
		}
		c, known := ConfigForChain(id.Uint64())
		if !known {
			return nil, &UnknownChainError{ChainID: id.Uint64()}
		}
		cfg = c
	}
	if o.ur != nil {
		cfg.V2.UniversalResolver = *o.ur
	}
	if err := cfg.Validate(); err != nil {
		return nil, err
	}
	return &Client{
		backend:  backend,
		cfg:      cfg,
		ur:       newUniversalResolver(backend, cfg.V2.UniversalResolver, o.http, o.maxHops, o.gateways),
		http:     o.http,
		maxHops:  o.maxHops,
		gateways: o.gateways,
	}, nil
}

// Backend returns the contract backend.
func (c *Client) Backend() bind.ContractBackend { return c.backend }

// Config returns the active chain configuration.
func (c *Client) Config() ChainConfig { return c.cfg }

// ChainID returns the configured chain ID.
func (c *Client) ChainID() uint64 { return c.cfg.ChainID }

// UniversalResolver returns the low-level Universal Resolver wrapper.
func (c *Client) UniversalResolver() *UniversalResolver { return c.ur }

// CCIPOptions returns the ccipread options the client uses, or nil for the
// package defaults.
func (c *Client) CCIPOptions() *ccipread.Options { return c.ur.callOptions() }

// Resolve runs resolve(name, callData) through the Universal Resolver. See
// UniversalResolver.Resolve.
func (c *Client) Resolve(ctx context.Context, name string, callData []byte) ([]byte, common.Address, error) {
	return c.ur.Resolve(ctx, name, callData)
}

// ResolveAddress resolves a name to its Ethereum address (coin type 60).
func (c *Client) ResolveAddress(ctx context.Context, name string) (common.Address, error) {
	return c.ur.ResolveAddress(ctx, name)
}

// MultiAddress resolves the address record for a SLIP-0044 coin type.
func (c *Client) MultiAddress(ctx context.Context, name string, coinType uint64) ([]byte, error) {
	return c.ur.MultiAddress(ctx, name, coinType)
}

// HasAddr reports whether an address record exists for the coin type.
func (c *Client) HasAddr(ctx context.Context, name string, coinType uint64) (bool, error) {
	return c.ur.HasAddr(ctx, name, coinType)
}

// Text resolves a text record.
func (c *Client) Text(ctx context.Context, name, key string) (string, error) {
	return c.ur.Text(ctx, name, key)
}

// Contenthash resolves the raw contenthash record.
func (c *Client) Contenthash(ctx context.Context, name string) ([]byte, error) {
	return c.ur.Contenthash(ctx, name)
}

// ContenthashString resolves the contenthash record as a URI such as
// ipfs://… or an empty string when unset.
func (c *Client) ContenthashString(ctx context.Context, name string) (string, error) {
	raw, err := c.ur.Contenthash(ctx, name)
	if err != nil || len(raw) == 0 {
		return "", err
	}
	return ContenthashToString(raw)
}

// PubKey resolves the public key record.
func (c *Client) PubKey(ctx context.Context, name string) ([32]byte, [32]byte, error) {
	return c.ur.PubKey(ctx, name)
}

// ABI resolves the ABI record for the given content types bitmask.
func (c *Client) ABI(ctx context.Context, name string, contentTypes uint64) (uint64, []byte, error) {
	return c.ur.ABI(ctx, name, contentTypes)
}

// ABIJSON resolves the ABI record as JSON text.
func (c *Client) ABIJSON(ctx context.Context, name string) (string, error) {
	return c.ur.ABIJSON(ctx, name)
}

// InterfaceImplementer resolves the EIP-165 interface implementer record.
func (c *Client) InterfaceImplementer(ctx context.Context, name string, interfaceID [4]byte) (common.Address, error) {
	return c.ur.InterfaceImplementer(ctx, name, interfaceID)
}

// NameRecord resolves the name(bytes32) record.
func (c *Client) NameRecord(ctx context.Context, name string) (string, error) {
	return c.ur.NameRecord(ctx, name)
}

// Data resolves an ENSIP-24 data record.
func (c *Client) Data(ctx context.Context, name, key string) ([]byte, error) {
	return c.ur.Data(ctx, name, key)
}

// Records fetches several records in one Universal Resolver call.
func (c *Client) Records(ctx context.Context, name string, req RecordsRequest) (*Records, error) {
	return c.ur.Records(ctx, name, req)
}

// IsENSv2 reports whether the configured Universal Resolver is
// UniversalResolverV2, that is whether the v2 resolution cutover has reached
// the address the client uses.
func (c *Client) IsENSv2(ctx context.Context) (bool, error) {
	return c.ur.IsENSv2(ctx)
}
