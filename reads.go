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

	"github.com/ensdomains/go-ens/v5/contracts/v1/resolver"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

// resolverABI is the ABI of the legacy PublicResolver, used to pack the
// calldata for individual resolver profiles that the Universal Resolver
// forwards to the discovered resolver. The profile selectors are stable
// across resolver generations: PublicResolverV2 keeps them verbatim and the
// v2 PermissionedResolver dispatches on them inside resolve(bytes,bytes),
// re-deriving the node from the DNS-encoded name.
var resolverABI abi.ABI

func init() {
	parsed, err := resolver.ContractMetaData.GetAbi()
	if err != nil {
		panic(fmt.Sprintf("ens: parse resolver abi: %v", err))
	}
	resolverABI = *parsed
}

// profileCall packs a resolver profile call for name and runs it through the
// Universal Resolver. It returns the raw result (empty when the record is
// unset) and the resolver address.
func (u *UniversalResolver) profileCall(ctx context.Context, name, method string, args ...any) ([]byte, error) {
	callData, err := resolverABI.Pack(method, args...)
	if err != nil {
		return nil, fmt.Errorf("pack %s: %w", method, err)
	}
	out, _, err := u.Resolve(ctx, name, callData)
	return out, err
}

func unpackOne[T any](method string, data []byte) (T, error) {
	var zero T
	values, err := resolverABI.Unpack(method, data)
	if err != nil {
		return zero, fmt.Errorf("unpack %s: %w", method, err)
	}
	if len(values) != 1 {
		return zero, fmt.Errorf("universal resolver: unexpected %s output arity", method)
	}
	out, ok := values[0].(T)
	if !ok {
		return zero, fmt.Errorf("universal resolver: %s result has unexpected type %T", method, values[0])
	}
	return out, nil
}

// Text resolves a text record for a name. It follows ENSIP-10 wildcards and
// ERC-3668 CCIP-Read. An unset record yields "" with no error.
func (u *UniversalResolver) Text(ctx context.Context, name, key string) (string, error) {
	nameHash, err := NameHash(name)
	if err != nil {
		return "", err
	}
	result, err := u.profileCall(ctx, name, "text", nameHash, key)
	if err != nil || len(result) == 0 {
		return "", err
	}
	return unpackOne[string]("text", result)
}

// Contenthash resolves the contenthash record for a name. An unset record
// yields nil with no error.
func (u *UniversalResolver) Contenthash(ctx context.Context, name string) ([]byte, error) {
	nameHash, err := NameHash(name)
	if err != nil {
		return nil, err
	}
	result, err := u.profileCall(ctx, name, "contenthash", nameHash)
	if err != nil || len(result) == 0 {
		return nil, err
	}
	out, err := unpackOne[[]byte]("contenthash", result)
	if err != nil {
		return nil, err
	}
	if len(out) == 0 {
		return nil, nil
	}
	return out, nil
}

// MultiAddress resolves the address of a name for an arbitrary SLIP-0044 coin
// type through the addr(bytes32,uint256) profile. It returns ErrNoAddress
// when the record is unset. For coin type 60 prefer ResolveAddress.
func (u *UniversalResolver) MultiAddress(ctx context.Context, name string, coinType uint64) ([]byte, error) {
	nameHash, err := NameHash(name)
	if err != nil {
		return nil, err
	}
	result, err := u.profileCall(ctx, name, "addr0", nameHash, new(big.Int).SetUint64(coinType))
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, ErrNoAddress
	}
	out, err := unpackOne[[]byte]("addr0", result)
	if err != nil {
		return nil, err
	}
	if len(out) == 0 {
		return nil, ErrNoAddress
	}
	return out, nil
}

// PubKey resolves the SECP256k1 public key record. Unset records yield zero
// values. v2 PermissionedResolvers do not implement the profile and surface
// ErrUnsupportedProfile.
func (u *UniversalResolver) PubKey(ctx context.Context, name string) ([32]byte, [32]byte, error) {
	var x, y [32]byte
	nameHash, err := NameHash(name)
	if err != nil {
		return x, y, err
	}
	result, err := u.profileCall(ctx, name, "pubkey", nameHash)
	if err != nil || len(result) == 0 {
		return x, y, err
	}
	values, err := resolverABI.Unpack("pubkey", result)
	if err != nil {
		return x, y, fmt.Errorf("unpack pubkey: %w", err)
	}
	if len(values) != 2 {
		return x, y, errors.New("universal resolver: unexpected pubkey output arity")
	}
	xv, okx := values[0].([32]byte)
	yv, oky := values[1].([32]byte)
	if !okx || !oky {
		return x, y, errors.New("universal resolver: pubkey result has unexpected type")
	}
	return xv, yv, nil
}

// ABI content types as defined by EIP-205.
const (
	ABIContentTypeJSON       = uint64(1)
	ABIContentTypeZlibJSON   = uint64(2)
	ABIContentTypeCBOR       = uint64(4)
	ABIContentTypeURI        = uint64(8)
	abiContentTypesAnyJSON   = ABIContentTypeJSON | ABIContentTypeZlibJSON
	abiMaxDecompressedLength = 1 << 20
)

// ABI resolves the ABI record, asking for any of the given content types
// (a bitmask). It returns the content type served and the raw data; an
// unset record yields 0 and nil.
func (u *UniversalResolver) ABI(ctx context.Context, name string, contentTypes uint64) (uint64, []byte, error) {
	nameHash, err := NameHash(name)
	if err != nil {
		return 0, nil, err
	}
	result, err := u.profileCall(ctx, name, "ABI", nameHash, new(big.Int).SetUint64(contentTypes))
	if err != nil || len(result) == 0 {
		return 0, nil, err
	}
	values, err := resolverABI.Unpack("ABI", result)
	if err != nil {
		return 0, nil, fmt.Errorf("unpack ABI: %w", err)
	}
	if len(values) != 2 {
		return 0, nil, errors.New("universal resolver: unexpected ABI output arity")
	}
	ct, ok := values[0].(*big.Int)
	if !ok {
		return 0, nil, errors.New("universal resolver: ABI content type has unexpected type")
	}
	data, ok := values[1].([]byte)
	if !ok {
		return 0, nil, errors.New("universal resolver: ABI data has unexpected type")
	}
	if len(data) == 0 {
		return 0, nil, nil
	}
	return ct.Uint64(), data, nil
}

// ABIJSON resolves the ABI record as JSON text, accepting plain or
// zlib-compressed JSON. An unset record yields "".
func (u *UniversalResolver) ABIJSON(ctx context.Context, name string) (string, error) {
	ct, data, err := u.ABI(ctx, name, abiContentTypesAnyJSON)
	if err != nil || len(data) == 0 {
		return "", err
	}
	if ct == ABIContentTypeZlibJSON {
		return inflateJSON(data)
	}
	return string(data), nil
}

// InterfaceImplementer resolves the EIP-165 interface implementer record.
// An unset record yields the zero address with no error.
func (u *UniversalResolver) InterfaceImplementer(ctx context.Context, name string, interfaceID [4]byte) (common.Address, error) {
	nameHash, err := NameHash(name)
	if err != nil {
		return common.Address{}, err
	}
	result, err := u.profileCall(ctx, name, "interfaceImplementer", nameHash, interfaceID)
	if err != nil || len(result) == 0 {
		return common.Address{}, err
	}
	return unpackOne[common.Address]("interfaceImplementer", result)
}

// NameRecord resolves the name(bytes32) record, which reverse nodes use to
// store a primary name. An unset record yields "".
func (u *UniversalResolver) NameRecord(ctx context.Context, name string) (string, error) {
	nameHash, err := NameHash(name)
	if err != nil {
		return "", err
	}
	result, err := u.profileCall(ctx, name, "name", nameHash)
	if err != nil || len(result) == 0 {
		return "", err
	}
	return unpackOne[string]("name", result)
}

// dataSelector is keccak256("data(bytes32,string)")[:4] (ENSIP-24), which the
// legacy PublicResolver ABI does not carry.
var dataArgs, hasAddrArgs abi.Arguments

func init() {
	bytes32T, _ := abi.NewType("bytes32", "", nil)
	stringT, _ := abi.NewType("string", "", nil)
	uint256T, _ := abi.NewType("uint256", "", nil)
	dataArgs = abi.Arguments{{Type: bytes32T}, {Type: stringT}}
	hasAddrArgs = abi.Arguments{{Type: bytes32T}, {Type: uint256T}}
}

var (
	dataSelector    = selectorOf("data(bytes32,string)")
	hasAddrSelector = selectorOf("hasAddr(bytes32,uint256)")
)

// Data resolves an ENSIP-24 data(bytes32,string) record. An unset record
// yields nil.
func (u *UniversalResolver) Data(ctx context.Context, name, key string) ([]byte, error) {
	nameHash, err := NameHash(name)
	if err != nil {
		return nil, err
	}
	body, err := dataArgs.Pack(nameHash, key)
	if err != nil {
		return nil, err
	}
	result, _, err := u.Resolve(ctx, name, append(dataSelector[:], body...))
	if err != nil || len(result) == 0 {
		return nil, err
	}
	bytesT, _ := abi.NewType("bytes", "", nil)
	values, err := abi.Arguments{{Type: bytesT}}.Unpack(result)
	if err != nil {
		return nil, fmt.Errorf("unpack data: %w", err)
	}
	out, _ := values[0].([]byte)
	if len(out) == 0 {
		return nil, nil
	}
	return out, nil
}

// HasAddr reports whether an address record exists for the coin type,
// through the v2 hasAddr(bytes32,uint256) profile. Resolvers without the
// profile yield false with no error.
func (u *UniversalResolver) HasAddr(ctx context.Context, name string, coinType uint64) (bool, error) {
	nameHash, err := NameHash(name)
	if err != nil {
		return false, err
	}
	body, err := hasAddrArgs.Pack(nameHash, new(big.Int).SetUint64(coinType))
	if err != nil {
		return false, err
	}
	result, _, err := u.Resolve(ctx, name, append(hasAddrSelector[:], body...))
	if err != nil {
		if errors.Is(err, ErrUnsupportedProfile) {
			return false, nil
		}
		return false, err
	}
	if len(result) < 32 {
		return false, nil
	}
	return result[31] == 1, nil
}
