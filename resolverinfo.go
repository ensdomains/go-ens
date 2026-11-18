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

	"github.com/ensdomains/go-ens/v5/contracts/v2/compositeresolver"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

// ERC-165 interface IDs used to classify resolvers.
var (
	InterfaceIDPermissionedResolver = [4]byte{0x8c, 0x24, 0x27, 0xcc}
	InterfaceIDCompositeResolver    = [4]byte{0xee, 0xa3, 0x30, 0xf9}
	InterfaceIDExtendedResolver     = [4]byte{0x90, 0x61, 0xb9, 0x23}
	InterfaceIDAddr                 = [4]byte{0x3b, 0x3b, 0x57, 0xde}
)

// ResolverFamily classifies a resolver contract by how it is written to.
type ResolverFamily int

const (
	// ResolverFamilyUnknown is a contract whose family could not be determined.
	ResolverFamilyUnknown ResolverFamily = iota
	// ResolverFamilyNodeKeyed is a v1 PublicResolver or PublicResolverV2:
	// setters take a bytes32 node.
	ResolverFamilyNodeKeyed
	// ResolverFamilyPermissioned is a v2 PermissionedResolver: setters take
	// a DNS-encoded name and authorisation is by roles on the resolver.
	ResolverFamilyPermissioned
	// ResolverFamilyMirror is a read-only composite resolver such as
	// ENSV1Resolver, ENSV2Resolver or DNSTLDResolver.
	ResolverFamilyMirror
)

func (f ResolverFamily) String() string {
	switch f {
	case ResolverFamilyNodeKeyed:
		return "node-keyed"
	case ResolverFamilyPermissioned:
		return "permissioned"
	case ResolverFamilyMirror:
		return "mirror"
	case ResolverFamilyUnknown:
		return "unknown"
	default:
		return "unknown"
	}
}

// ResolverInfo describes the resolver the Universal Resolver uses for a name.
type ResolverInfo struct {
	// Resolver is what findResolver returned. After the v2 cutover it may be
	// a mirror such as ENSV1Resolver for names still held in the v1 registry.
	Resolver common.Address
	// Node is the namehash the resolver was found for.
	Node [32]byte
	// Offset is the byte offset into the DNS-encoded name at which the
	// resolver was found; non-zero means a parent's resolver (wildcard).
	Offset uint64
	// Composite reports whether Resolver implements ICompositeResolver and
	// Underlying was obtained by unwrapping it.
	Composite bool
	// Underlying is the resolver that actually holds the records: the
	// unwrapped target for composite resolvers, else Resolver.
	Underlying common.Address
	// Offchain is the composite resolver's report that Underlying lives
	// offchain or on another chain.
	Offchain bool
	// Family is the detected family of Underlying.
	Family ResolverFamily
	// Writable is true when Underlying is non-zero and not a mirror.
	Writable bool
}

// FindResolver returns the resolver the Universal Resolver uses for a name
// without unwrapping composites. It returns ErrUnregistered when no resolver
// is set anywhere up the parent chain.
func (c *Client) FindResolver(ctx context.Context, name string) (*ResolverInfo, error) {
	resolver, node, offset, err := c.ur.FindResolver(ctx, name)
	if err != nil {
		return nil, err
	}
	if resolver == (common.Address{}) {
		return nil, &ResolverNotFoundError{}
	}
	return &ResolverInfo{Resolver: resolver, Node: node, Offset: offset, Underlying: resolver}, nil
}

// WritableResolver finds the resolver for a name, unwraps composite mirrors
// to the contract that holds the records, and classifies it. This is the
// resolver to use for writes; record reads should go through Resolve.
func (c *Client) WritableResolver(ctx context.Context, name string) (*ResolverInfo, error) {
	info, err := c.FindResolver(ctx, name)
	if err != nil {
		return nil, err
	}
	if c.supportsInterface(ctx, info.Resolver, InterfaceIDCompositeResolver) {
		info.Composite = true
		dnsName, err := DNSEncode(name)
		if err != nil {
			return nil, err
		}
		cr := compositeresolver.NewCompositeResolver()
		to := info.Resolver
		out, err := c.backend.CallContract(ctx, ethereum.CallMsg{To: &to, Data: cr.PackGetResolver(dnsName)}, nil)
		if err == nil {
			if res, uerr := cr.UnpackGetResolver(out); uerr == nil {
				info.Underlying = res.Arg0
				info.Offchain = res.Arg1
			}
		}
	}
	info.Family = c.resolverFamily(ctx, info.Underlying, info.Composite)
	info.Writable = info.Underlying != (common.Address{}) && info.Family != ResolverFamilyMirror
	return info, nil
}

// resolverFamily classifies a resolver with ERC-165 probes. Probe failures
// count as unsupported, as ensjs does.
func (c *Client) resolverFamily(ctx context.Context, addr common.Address, unwrapped bool) ResolverFamily {
	if addr == (common.Address{}) {
		return ResolverFamilyUnknown
	}
	if c.supportsInterface(ctx, addr, InterfaceIDPermissionedResolver) {
		return ResolverFamilyPermissioned
	}
	if c.supportsInterface(ctx, addr, InterfaceIDCompositeResolver) {
		// A composite that unwrapped to another composite, or a mirror
		// reached directly: either way nothing is writable here.
		return ResolverFamilyMirror
	}
	if addr == c.cfg.V2.ENSV1Resolver || addr == c.cfg.V2.ENSV2Resolver || addr == c.cfg.V2.DNSTLDResolver {
		return ResolverFamilyMirror
	}
	if addr == c.cfg.V2.PublicResolverV2 || addr == c.cfg.V1.PublicResolver || c.supportsInterface(ctx, addr, InterfaceIDAddr) {
		return ResolverFamilyNodeKeyed
	}
	_ = unwrapped
	return ResolverFamilyUnknown
}

var supportsInterfaceSelector = selectorOf("supportsInterface(bytes4)")

// supportsInterface runs an ERC-165 probe, treating any failure as false.
func (c *Client) supportsInterface(ctx context.Context, addr common.Address, id [4]byte) bool {
	data := make([]byte, 4+32)
	copy(data[:4], supportsInterfaceSelector[:])
	copy(data[4:8], id[:])
	out, err := c.backend.CallContract(ctx, ethereum.CallMsg{To: &addr, Data: data}, nil)
	if err != nil || len(out) < 32 {
		return false
	}
	return out[31] == 1
}
