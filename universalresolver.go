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
	"fmt"
	"math/big"
	"net/http"

	"github.com/ensdomains/go-ens/v5/ccipread"
	"github.com/ensdomains/go-ens/v5/contracts/v2/universalresolver"
	"github.com/ensdomains/go-ens/v5/internal/chainid"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// CoinTypeETH is SLIP-0044 coin type 60: native Ethereum addresses.
const CoinTypeETH = uint64(60)

// addrSelector is keccak256("addr(bytes32)")[:4]: the legacy ETH-address
// resolver profile.
var addrSelector = [4]byte{0x3b, 0x3b, 0x57, 0xde}

// Parameter is a functional option for NewUniversalResolver. It is the same
// type as Option, so WithHTTPClient, WithMaxHops and WithGatewayURLs apply to
// both constructors.
type Parameter = Option

// WithAddress binds the resolver to an explicit Universal Resolver address
// instead of the canonical proxy. Use it on chains without a baked-in
// configuration, or to target UniversalResolverV2 directly before the
// canonical proxy is repointed; it also skips the chain-ID guardrail.
func WithAddress(address common.Address) Parameter {
	return func(o *clientOptions) { o.ur = &address; o.skipProbe = true }
}

// WithGateways overrides the batch gateway URLs the Universal Resolver uses
// for CCIP-Read. It is an alias of WithGatewayURLs.
func WithGateways(urls []string) Parameter { return WithGatewayURLs(urls) }

// UniversalResolver wraps the ENS Universal Resolver and exposes resolution
// helpers that transparently follow ERC-3668 OffchainLookup reverts via the
// ccipread package. It speaks the UniversalResolverV2 ABI, whose resolve and
// reverse entry points are compatible with the v1 implementation that the
// canonical proxy serves until the ENS DAO completes the v2 cutover.
type UniversalResolver struct {
	backend    bind.ContractBackend
	addr       common.Address
	c          *universalresolver.UniversalResolverV2
	httpClient *http.Client
	maxHops    int
	gateways   []string
}

// NewUniversalResolver returns a UniversalResolver for the chain reached
// through backend. It honours ctx for the chain-ID probe.
//
// By default it binds to the canonical proxy address and, when the backend
// exposes a chain ID, verifies the chain has a baked-in configuration,
// returning *UnknownChainError otherwise. Supply WithAddress to target any
// other deployment, which also skips the guardrail.
func NewUniversalResolver(ctx context.Context, backend bind.ContractBackend, params ...Parameter) (*UniversalResolver, error) {
	var o clientOptions
	for _, param := range params {
		param(&o)
	}
	addr := common.HexToAddress(UniversalResolverContractAddress)
	switch {
	case o.ur != nil:
		addr = *o.ur
	case o.cfg != nil:
		addr = o.cfg.V2.UniversalResolver
	case o.chainID != nil:
		cfg, ok := ConfigForChain(*o.chainID)
		if !ok {
			return nil, &UnknownChainError{ChainID: *o.chainID}
		}
		addr = cfg.V2.UniversalResolver
	default:
		id, ok, err := chainid.Probe(ctx, backend)
		if err != nil && ctx.Err() != nil {
			return nil, ctx.Err()
		}
		if err == nil && ok && id != nil {
			if _, known := ConfigForChain(id.Uint64()); !known {
				return nil, &UnknownChainError{ChainID: id.Uint64()}
			}
		}
	}
	return newUniversalResolver(backend, addr, o.http, o.maxHops, o.gateways), nil
}

func newUniversalResolver(backend bind.ContractBackend, addr common.Address, httpClient *http.Client, maxHops int, gateways []string) *UniversalResolver {
	return &UniversalResolver{
		backend:    backend,
		addr:       addr,
		c:          universalresolver.NewUniversalResolverV2(),
		httpClient: httpClient,
		maxHops:    maxHops,
		gateways:   gateways,
	}
}

// callOptions builds the ccipread options from the resolver's configuration,
// returning nil when nothing has been customised so ccipread uses its defaults.
func (u *UniversalResolver) callOptions() *ccipread.Options {
	if u.httpClient == nil && u.maxHops == 0 {
		return nil
	}
	return &ccipread.Options{HTTPClient: u.httpClient, MaxRedirects: u.maxHops}
}

// Address returns the on-chain address of the bound Universal Resolver.
func (u *UniversalResolver) Address() common.Address {
	return u.addr
}

// call runs a CCIP-Read aware eth_call against the resolver and translates
// its custom errors.
func (u *UniversalResolver) call(ctx context.Context, input []byte) ([]byte, error) {
	out, err := ccipread.Call(ctx, u.backend, u.addr, input, u.callOptions())
	if err != nil {
		return nil, translateURRevert(err)
	}
	return out, nil
}

// view runs a plain eth_call (no CCIP-Read) for view functions.
func (u *UniversalResolver) view(ctx context.Context, input []byte) ([]byte, error) {
	out, err := u.backend.CallContract(ctx, ethereum.CallMsg{To: &u.addr, Data: input}, nil)
	if err != nil {
		return nil, translateURRevert(err)
	}
	return out, nil
}

// Resolve performs resolve(name, callData), DNS-encoding and normalising the
// name first and transparently following any OffchainLookup reverts. The
// returned bytes are the ABI-encoded return value of the downstream resolver
// call (for addr(bytes32) they decode as address); resolver is the contract
// the Universal Resolver called, which after the v2 cutover may be a
// read-only mirror such as ENSV1Resolver rather than a writable resolver.
func (u *UniversalResolver) Resolve(ctx context.Context, name string, callData []byte) ([]byte, common.Address, error) {
	dnsName, err := DNSEncode(name)
	if err != nil {
		return nil, common.Address{}, fmt.Errorf("dns-encode %q: %w", name, err)
	}
	return u.ResolveDNS(ctx, dnsName, callData)
}

// ResolveDNS is Resolve for a name that is already DNS-encoded.
func (u *UniversalResolver) ResolveDNS(ctx context.Context, dnsName, callData []byte) ([]byte, common.Address, error) {
	var input []byte
	var err error
	if len(u.gateways) > 0 {
		input, err = u.c.TryPackResolveWithGateways(dnsName, callData, u.gateways)
	} else {
		input, err = u.c.TryPackResolve(dnsName, callData)
	}
	if err != nil {
		return nil, common.Address{}, fmt.Errorf("pack resolve: %w", err)
	}
	out, err := u.call(ctx, input)
	if err != nil {
		return nil, common.Address{}, err
	}
	res, err := u.c.UnpackResolve(out)
	if err != nil {
		return nil, common.Address{}, fmt.Errorf("unpack resolve: %w", err)
	}
	return res.Arg0, res.Arg1, nil
}

// ResolveAddress resolves a name to its primary Ethereum address through the
// legacy addr(bytes32) profile. It returns ErrNoAddress when the record is
// missing or zero.
func (u *UniversalResolver) ResolveAddress(ctx context.Context, name string) (common.Address, error) {
	nameHash, err := NameHash(name)
	if err != nil {
		return UnknownAddress, err
	}
	callData := make([]byte, 4+32)
	copy(callData[:4], addrSelector[:])
	copy(callData[4:], nameHash[:])
	result, _, err := u.Resolve(ctx, name, callData)
	if err != nil {
		return UnknownAddress, err
	}
	if len(result) == 0 {
		return UnknownAddress, ErrNoAddress
	}
	if len(result) != 32 {
		return UnknownAddress, fmt.Errorf("addr result has unexpected length %d", len(result))
	}
	addr := common.BytesToAddress(result)
	if addr == UnknownAddress {
		return UnknownAddress, ErrNoAddress
	}
	return addr, nil
}

// Reverse performs reverse(lookupAddress, coinType) and returns the primary
// name, the resolver that answered the forward check and the reverse
// resolver. An address with no primary name yields an empty string and no
// error. CCIP-Read is followed transparently.
func (u *UniversalResolver) Reverse(ctx context.Context, address []byte, coinType uint64) (string, common.Address, common.Address, error) {
	ct := new(big.Int).SetUint64(coinType)
	var input []byte
	var err error
	if len(u.gateways) > 0 {
		input, err = u.c.TryPackReverseWithGateways(address, ct, u.gateways)
	} else {
		input, err = u.c.TryPackReverse(address, ct)
	}
	if err != nil {
		return "", common.Address{}, common.Address{}, fmt.Errorf("pack reverse: %w", err)
	}
	out, err := u.call(ctx, input)
	if err != nil {
		return "", common.Address{}, common.Address{}, err
	}
	res, err := u.c.UnpackReverse(out)
	if err != nil {
		return "", common.Address{}, common.Address{}, fmt.Errorf("unpack reverse: %w", err)
	}
	return res.Arg0, res.Arg1, res.Arg2, nil
}

// FindResolver performs findResolver(name): the resolver the Universal
// Resolver would use for the name, the node it matched and the byte offset
// into the DNS-encoded name at which the match was made (non-zero means a
// parent's resolver, that is a wildcard). The resolver is zero when none is
// set anywhere up the chain.
func (u *UniversalResolver) FindResolver(ctx context.Context, name string) (common.Address, [32]byte, uint64, error) {
	dnsName, err := DNSEncode(name)
	if err != nil {
		return common.Address{}, [32]byte{}, 0, fmt.Errorf("dns-encode %q: %w", name, err)
	}
	out, err := u.view(ctx, u.c.PackFindResolver(dnsName))
	if err != nil {
		return common.Address{}, [32]byte{}, 0, err
	}
	res, err := u.c.UnpackFindResolver(out)
	if err != nil {
		return common.Address{}, [32]byte{}, 0, fmt.Errorf("unpack findResolver: %w", err)
	}
	return res.Resolver, res.Node, res.Offset.Uint64(), nil
}

// IsENSv2 reports whether the bound resolver is UniversalResolverV2. The
// legacy v1 implementation has no isENSv2() and reverts, and the normalising
// UniversalResolverV1 returns false; both map to false with a nil error.
func (u *UniversalResolver) IsENSv2(ctx context.Context) (bool, error) {
	out, err := u.backend.CallContract(ctx, ethereum.CallMsg{To: &u.addr, Data: u.c.PackIsENSv2()}, nil)
	if err != nil {
		if ctx.Err() != nil {
			return false, ctx.Err()
		}
		return false, nil // legacy resolver reverts: not v2
	}
	v, err := u.c.UnpackIsENSv2(out)
	if err != nil {
		return false, nil //nolint:nilerr // undecodable answer: not v2
	}
	return v, nil
}
