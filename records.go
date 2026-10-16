// Copyright 2026 Weald Technology Trading.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0

package ens

import (
	"bytes"
	"compress/zlib"
	"context"
	"errors"
	"fmt"
	"io"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// inflateJSON decompresses a zlib-compressed ABI record.
func inflateJSON(data []byte) (string, error) {
	r, err := zlib.NewReader(bytes.NewReader(data))
	if err != nil {
		return "", fmt.Errorf("decompress ABI: %w", err)
	}
	defer r.Close()
	out, err := io.ReadAll(io.LimitReader(r, abiMaxDecompressedLength))
	if err != nil {
		return "", fmt.Errorf("decompress ABI: %w", err)
	}
	return string(out), nil
}

// Selector returns the 4-byte function or error selector of a canonical
// signature such as "addr(bytes32)".
func Selector(sig string) [4]byte { return selectorOf(sig) }

func selectorOf(sig string) [4]byte {
	var s [4]byte
	copy(s[:], crypto.Keccak256([]byte(sig))[:4])
	return s
}

// multicallSelector is keccak256("multicall(bytes[])")[:4].
var multicallSelector = selectorOf("multicall(bytes[])")

var bytesArrayArgs abi.Arguments

func init() {
	t, _ := abi.NewType("bytes[]", "", nil)
	bytesArrayArgs = abi.Arguments{{Type: t}}
}

// PackMulticall wraps several resolver profile calls into a single
// multicall(bytes[]) payload suitable for Resolve.
func PackMulticall(calls [][]byte) ([]byte, error) {
	body, err := bytesArrayArgs.Pack(calls)
	if err != nil {
		return nil, err
	}
	return append(multicallSelector[:], body...), nil
}

// UnpackMulticall decodes a multicall(bytes[]) result into its per-call
// return data. Failed items carry their raw revert data (or are empty).
func UnpackMulticall(data []byte) ([][]byte, error) {
	values, err := bytesArrayArgs.Unpack(data)
	if err != nil {
		return nil, fmt.Errorf("unpack multicall: %w", err)
	}
	out, ok := values[0].([][]byte)
	if !ok {
		return nil, errors.New("multicall result has unexpected type")
	}
	return out, nil
}

// RecordsRequest selects the records Records fetches in one call.
type RecordsRequest struct {
	Texts       []string
	Coins       []uint64
	Contenthash bool
	ABI         bool // JSON or zlib-JSON ABI, decoded to text
	Name        bool // name(bytes32)
}

// Records is the result of a batched record read. Maps only contain keys
// whose record is set. Errors holds per-item failures keyed "text:<key>",
// "coin:<type>", "contenthash", "abi" or "name"; an item that is merely
// unset is not an error.
type Records struct {
	Resolver    common.Address
	Texts       map[string]string
	Coins       map[uint64][]byte
	Contenthash []byte
	ABI         string
	Name        string
	Errors      map[string]error
}

type recordItem struct {
	key    string
	decode func(data []byte) error
}

// Records fetches several records for a name in one Universal Resolver call
// by wrapping the profile calls in multicall(bytes[]). It follows ENSIP-10
// wildcards and ERC-3668 CCIP-Read like every other read.
func (u *UniversalResolver) Records(ctx context.Context, name string, req RecordsRequest) (*Records, error) {
	nameHash, err := NameHash(name)
	if err != nil {
		return nil, err
	}
	out := &Records{Texts: map[string]string{}, Coins: map[uint64][]byte{}, Errors: map[string]error{}}
	var calls [][]byte
	var items []recordItem
	add := func(key string, calldata []byte, decode func([]byte) error) {
		calls = append(calls, calldata)
		items = append(items, recordItem{key: key, decode: decode})
	}
	for _, key := range req.Texts {
		cd, err := resolverABI.Pack("text", nameHash, key)
		if err != nil {
			return nil, err
		}
		k := key
		add("text:"+k, cd, func(data []byte) error {
			v, err := unpackOne[string]("text", data)
			if err != nil {
				return err
			}
			if v != "" {
				out.Texts[k] = v
			}
			return nil
		})
	}
	for _, coin := range req.Coins {
		cd, err := resolverABI.Pack("addr0", nameHash, new(big.Int).SetUint64(coin))
		if err != nil {
			return nil, err
		}
		c := coin
		add("coin:"+strconv.FormatUint(c, 10), cd, func(data []byte) error {
			v, err := unpackOne[[]byte]("addr0", data)
			if err != nil {
				return err
			}
			if len(v) > 0 {
				out.Coins[c] = v
			}
			return nil
		})
	}
	if req.Contenthash {
		cd, err := resolverABI.Pack("contenthash", nameHash)
		if err != nil {
			return nil, err
		}
		add("contenthash", cd, func(data []byte) error {
			v, err := unpackOne[[]byte]("contenthash", data)
			if err != nil {
				return err
			}
			if len(v) > 0 {
				out.Contenthash = v
			}
			return nil
		})
	}
	if req.ABI {
		cd, err := resolverABI.Pack("ABI", nameHash, new(big.Int).SetUint64(abiContentTypesAnyJSON))
		if err != nil {
			return nil, err
		}
		add("abi", cd, func(data []byte) error {
			values, err := resolverABI.Unpack("ABI", data)
			if err != nil {
				return err
			}
			if len(values) != 2 {
				return errors.New("unexpected ABI output arity")
			}
			ct, _ := values[0].(*big.Int)
			raw, _ := values[1].([]byte)
			if len(raw) == 0 {
				return nil
			}
			if ct != nil && ct.Uint64() == ABIContentTypeZlibJSON {
				s, err := inflateJSON(raw)
				if err != nil {
					return err
				}
				out.ABI = s
				return nil
			}
			out.ABI = string(raw)
			return nil
		})
	}
	if req.Name {
		cd, err := resolverABI.Pack("name", nameHash)
		if err != nil {
			return nil, err
		}
		add("name", cd, func(data []byte) error {
			v, err := unpackOne[string]("name", data)
			if err != nil {
				return err
			}
			out.Name = v
			return nil
		})
	}
	if len(calls) == 0 {
		return out, nil
	}
	payload, err := PackMulticall(calls)
	if err != nil {
		return nil, err
	}
	result, resolverAddr, err := u.Resolve(ctx, name, payload)
	if err != nil {
		return nil, err
	}
	out.Resolver = resolverAddr
	results, err := UnpackMulticall(result)
	if err != nil {
		return nil, err
	}
	if len(results) != len(items) {
		return nil, fmt.Errorf("multicall returned %d items for %d calls", len(results), len(items))
	}
	for i, item := range items {
		data := results[i]
		if len(data) == 0 {
			continue
		}
		if len(data)%32 != 0 {
			// Reverts are selector-prefixed and never word-aligned; decode
			// them with the Universal Resolver error table when possible.
			inner, _ := urDecoder.TranslateData(data)
			out.Errors[item.key] = inner
			continue
		}
		if err := item.decode(data); err != nil {
			out.Errors[item.key] = err
		}
	}
	return out, nil
}
