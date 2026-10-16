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

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// ReverseResult is the outcome of a reverse lookup.
type ReverseResult struct {
	// Name is the primary name, or "" when none is set.
	Name string
	// Resolver is the resolver that answered the forward check. After the
	// v2 cutover it may be a read-only mirror; do not write to it.
	Resolver common.Address
	// ReverseResolver is the resolver of the reverse node.
	ReverseResolver common.Address
}

// Reverse looks up the primary name of an address for a coin type through
// the Universal Resolver, which verifies that the name resolves forward to
// the same address and otherwise returns *ReverseAddressMismatchError. Use
// EVMCoinType to derive the coin type of an EVM chain.
func (c *Client) Reverse(ctx context.Context, address []byte, coinType uint64) (*ReverseResult, error) {
	name, resolver, reverseResolver, err := c.ur.Reverse(ctx, address, coinType)
	if err != nil {
		return nil, err
	}
	return &ReverseResult{Name: name, Resolver: resolver, ReverseResolver: reverseResolver}, nil
}

// ReverseResolve returns the primary name of an Ethereum address, or
// ErrNoPrimaryName when none is set.
func (c *Client) ReverseResolve(ctx context.Context, address common.Address) (string, error) {
	return c.ReverseResolveCoin(ctx, address.Bytes(), CoinTypeETH)
}

// ReverseResolveCoin returns the primary name of an address for a coin
// type, or ErrNoPrimaryName when none is set.
func (c *Client) ReverseResolveCoin(ctx context.Context, address []byte, coinType uint64) (string, error) {
	res, err := c.Reverse(ctx, address, coinType)
	if err != nil {
		return "", err
	}
	if res.Name == "" {
		return "", ErrNoPrimaryName
	}
	return res.Name, nil
}

// Format returns the primary name of an address if it has one, otherwise
// the checksummed address.
func (c *Client) Format(ctx context.Context, address common.Address) string {
	name, err := c.ReverseResolve(ctx, address)
	if err != nil {
		return address.Hex()
	}
	return name
}

// FormatCoin returns the primary name of an address for a coin type if it
// has one, otherwise fallback.
func (c *Client) FormatCoin(ctx context.Context, address []byte, coinType uint64, fallback string) string {
	name, err := c.ReverseResolveCoin(ctx, address, coinType)
	if err != nil {
		return fallback
	}
	return name
}

// Package-level shortcuts. Each constructs a throwaway Client, which costs
// one eth_chainId call; servers should hold a Client instead.

// Resolve resolves a name to an Ethereum address through the Universal
// Resolver, following ENSIP-10 wildcards and ERC-3668 CCIP-Read. If input
// contains no dot it is treated as a literal hex address and returned after
// validation without any network call.
func Resolve(ctx context.Context, backend bind.ContractBackend, input string) (common.Address, error) {
	if !containsDot(input) {
		return parseAddress(input)
	}
	if _, err := NameHash(input); err != nil {
		return UnknownAddress, err
	}
	c, err := NewClient(ctx, backend)
	if err != nil {
		return UnknownAddress, err
	}
	return c.ResolveAddress(ctx, input)
}

// ReverseResolve returns the primary name of an address, or ErrNoPrimaryName.
func ReverseResolve(ctx context.Context, backend bind.ContractBackend, address common.Address) (string, error) {
	c, err := NewClient(ctx, backend)
	if err != nil {
		return "", err
	}
	return c.ReverseResolve(ctx, address)
}

// Format returns the primary name of an address if it has one, otherwise
// the checksummed address.
func Format(ctx context.Context, backend bind.ContractBackend, address common.Address) string {
	name, err := ReverseResolve(ctx, backend, address)
	if err != nil {
		return address.Hex()
	}
	return name
}

func containsDot(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == '.' {
			return true
		}
	}
	return false
}

func parseAddress(input string) (common.Address, error) {
	hasPrefix := len(input) >= 2 && input[0] == '0' && (input[1] == 'x' || input[1] == 'X')
	if (hasPrefix && len(input) > 42) || (!hasPrefix && len(input) > 40) {
		return UnknownAddress, errAddressTooLong
	}
	address := common.HexToAddress(input)
	if address == UnknownAddress {
		return UnknownAddress, errAddressUnparseable
	}
	return address, nil
}
