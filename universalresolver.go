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
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/wealdtech/go-ens/v3/ccipread"
	"github.com/wealdtech/go-ens/v3/contracts/universalresolver"
)

// UniversalResolverContractAddress is the proxy address of the ENS
// UniversalResolver. The same address is used on Ethereum mainnet and Sepolia
// (the proxy is deployed at the vanity address by the ENS DAO).
const UniversalResolverContractAddress = "0xeEeEEEeE14D718C2B47D9923Deab1335E144EeEe"

// CoinTypeETH is SLIP-0044 coin type 60: native Ethereum addresses.
const CoinTypeETH = uint64(60)

// addrSelector is keccak256("addr(bytes32)")[:4]: the legacy ETH-address
// resolver profile. UniversalResolver.resolve forwards calldata starting with
// this selector to the discovered resolver and returns the abi-encoded result.
var addrSelector = [4]byte{0x3b, 0x3b, 0x57, 0xde}

// knownURChains lists the chain IDs on which the canonical UR proxy address
// is deployed by the ENS DAO. Callers on other chains must pass an explicit
// address via NewUniversalResolverAt.
var knownURChains = map[uint64]struct{}{
	1:        {}, // Ethereum mainnet
	11155111: {}, // Sepolia testnet
}

// UnknownChainError is returned by NewUniversalResolver when the backend
// reports a chain ID without a known UR deployment at the canonical address.
type UnknownChainError struct {
	ChainID *big.Int
}

func (e *UnknownChainError) Error() string {
	return fmt.Sprintf("universal resolver: no canonical deployment on chain ID %s; pass an explicit address via NewUniversalResolverAt", e.ChainID)
}

// UniversalResolver wraps the ENS Universal Resolver and exposes resolution
// helpers that transparently follow ERC-3668 OffchainLookup reverts via the
// ccipread package.
type UniversalResolver struct {
	backend bind.ContractBackend
	addr    common.Address
	abi     abi.ABI
}

// NewUniversalResolver returns a UniversalResolver bound to the canonical
// proxy address on whichever chain the backend is connected to.
//
// When the backend exposes a chain ID (ethclient.Client and most production
// backends do) this verifies the chain is one where the canonical UR is
// deployed and returns *UnknownChainError otherwise. Backends that don't
// expose a chain ID skip the check and the canonical address is used as-is —
// callers on devnets / forks / alt-L1s should prefer NewUniversalResolverAt.
func NewUniversalResolver(backend bind.ContractBackend) (*UniversalResolver, error) {
	if cid, ok := backend.(interface {
		ChainID(context.Context) (*big.Int, error)
	}); ok {
		chainID, err := cid.ChainID(context.Background())
		if err == nil && chainID != nil {
			if _, known := knownURChains[chainID.Uint64()]; !known {
				return nil, &UnknownChainError{ChainID: chainID}
			}
		}
	}
	return NewUniversalResolverAt(backend, common.HexToAddress(UniversalResolverContractAddress))
}

// NewUniversalResolverAt returns a UniversalResolver bound to a specific
// address. Use this when targeting a non-standard deployment, e.g. a local
// fork.
func NewUniversalResolverAt(backend bind.ContractBackend, addr common.Address) (*UniversalResolver, error) {
	parsed, err := universalresolver.ContractMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("parse universal resolver abi: %w", err)
	}
	return &UniversalResolver{backend: backend, addr: addr, abi: *parsed}, nil
}

// Address returns the on-chain address of the bound UniversalResolver.
func (u *UniversalResolver) Address() common.Address {
	return u.addr
}

// Resolve performs UniversalResolver.resolve(name, callData) using the chain
// connected via the backend, transparently following any OffchainLookup
// reverts. The returned bytes are the abi-encoded return value of the
// downstream resolver call (e.g. for `addr(bytes32)` the bytes decode as
// `address`); resolverAddr is the resolver contract that produced it.
func (u *UniversalResolver) Resolve(ctx context.Context, name string, callData []byte) ([]byte, common.Address, error) {
	dnsName, err := DNSEncode(name)
	if err != nil {
		return nil, common.Address{}, fmt.Errorf("dns-encode %q: %w", name, err)
	}
	input, err := u.abi.Pack("resolve", dnsName, callData)
	if err != nil {
		return nil, common.Address{}, fmt.Errorf("pack resolve: %w", err)
	}
	out, err := ccipread.Call(ctx, u.backend, u.addr, input, nil)
	if err != nil {
		return nil, common.Address{}, translateURRevert(err)
	}
	values, err := u.abi.Unpack("resolve", out)
	if err != nil {
		return nil, common.Address{}, fmt.Errorf("unpack resolve: %w", err)
	}
	if len(values) != 2 {
		return nil, common.Address{}, errors.New("universal resolver: unexpected resolve output arity")
	}
	result, _ := values[0].([]byte)
	resolver, _ := values[1].(common.Address)
	return result, resolver, nil
}

// ResolveAddress resolves an ENS name to its primary Ethereum address by
// invoking the legacy `addr(bytes32)` resolver profile through the
// UniversalResolver.
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

// Reverse performs UniversalResolver.reverse(lookupAddress, coinType) and
// returns the primary name, the resolver address, and the reverse resolver
// address. CCIP-Read is followed transparently.
func (u *UniversalResolver) Reverse(ctx context.Context, address []byte, coinType uint64) (string, common.Address, common.Address, error) {
	input, err := u.abi.Pack("reverse", address, new(big.Int).SetUint64(coinType))
	if err != nil {
		return "", common.Address{}, common.Address{}, fmt.Errorf("pack reverse: %w", err)
	}
	out, err := ccipread.Call(ctx, u.backend, u.addr, input, nil)
	if err != nil {
		return "", common.Address{}, common.Address{}, translateURRevert(err)
	}
	values, err := u.abi.Unpack("reverse", out)
	if err != nil {
		return "", common.Address{}, common.Address{}, fmt.Errorf("unpack reverse: %w", err)
	}
	if len(values) != 3 {
		return "", common.Address{}, common.Address{}, errors.New("universal resolver: unexpected reverse output arity")
	}
	name, _ := values[0].(string)
	resolver, _ := values[1].(common.Address)
	reverseResolver, _ := values[2].(common.Address)
	return name, resolver, reverseResolver, nil
}
