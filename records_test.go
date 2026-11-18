// Copyright 2026 Weald Technology Trading.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0

package ens_test

import (
	"context"
	"encoding/hex"
	"math/big"
	"testing"

	ens "github.com/ensdomains/go-ens/v5"
	"github.com/ensdomains/go-ens/v5/contracts/v1/resolver"
	"github.com/ensdomains/go-ens/v5/contracts/v2/universalresolver"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

// fakeBackend answers eth_call with a function of the calldata and records
// what it was asked. It satisfies bind.ContractBackend; transaction methods
// are unused.
type fakeBackend struct {
	bind.ContractBackend
	calls  []ethereum.CallMsg
	answer func(msg ethereum.CallMsg) ([]byte, error)
}

func (f *fakeBackend) CallContract(_ context.Context, msg ethereum.CallMsg, _ *big.Int) ([]byte, error) {
	f.calls = append(f.calls, msg)
	return f.answer(msg)
}

func (f *fakeBackend) CodeAt(context.Context, common.Address, *big.Int) ([]byte, error) {
	return []byte{1}, nil
}

func (f *fakeBackend) HeaderByNumber(context.Context, *big.Int) (*types.Header, error) {
	return &types.Header{Number: big.NewInt(1)}, nil
}

func mustABI(t *testing.T, typ string) abi.Type {
	t.Helper()
	ty, err := abi.NewType(typ, "", nil)
	require.NoError(t, err)
	return ty
}

// TestRecordsMulticall checks that Records wraps the profile calls in
// multicall(bytes[]) and decodes the per-item results, treating empty
// items as unset and revert payloads as per-item errors.
func TestRecordsMulticall(t *testing.T) {
	resolverAddr := common.HexToAddress("0x1111111111111111111111111111111111111111")
	resABI, err := resolver.ContractMetaData.GetAbi()
	require.NoError(t, err)
	ur := universalresolver.NewUniversalResolverV2()

	backend := &fakeBackend{}
	backend.answer = func(msg ethereum.CallMsg) ([]byte, error) {
		// The resolve(bytes,bytes) request carries multicall calldata.
		require.Equal(t, ens.Selector("resolve(bytes,bytes)"), [4]byte(msg.Data[:4]))
		args, err := ur.GetABI().Methods["resolve"].Inputs.Unpack(msg.Data[4:])
		require.NoError(t, err)
		inner, ok := args[1].([]byte)
		require.True(t, ok)
		require.Equal(t, ens.Selector("multicall(bytes[])"), [4]byte(inner[:4]))
		calls, err := abi.Arguments{{Type: mustABI(t, "bytes[]")}}.Unpack(inner[4:])
		require.NoError(t, err)
		items, ok := calls[0].([][]byte)
		require.True(t, ok)
		require.Len(t, items, 4)

		textOut, err := resABI.Methods["text"].Outputs.Pack("hello")
		require.NoError(t, err)
		coinOut, err := resABI.Methods["addr0"].Outputs.Pack(common.HexToAddress("0x02").Bytes())
		require.NoError(t, err)
		// Item 2 (contenthash) is unset; item 3 (name) reverted with
		// UnsupportedResolverProfile(bytes4).
		revertPayload := append(append([]byte{}, crypto.Keccak256([]byte("UnsupportedResolverProfile(bytes4)"))[:4]...), common.LeftPadBytes([]byte{0x69, 0x1f, 0x34, 0x31}, 32)...)
		revertPayload = revertPayload[:4+32]
		copy(revertPayload[4:8], []byte{0x69, 0x1f, 0x34, 0x31})
		for i := 8; i < 36; i++ {
			revertPayload[i] = 0
		}
		results, err := abi.Arguments{{Type: mustABI(t, "bytes[]")}}.Pack([][]byte{textOut, coinOut, {}, revertPayload})
		require.NoError(t, err)
		out, err := ur.GetABI().Methods["resolve"].Outputs.Pack(results, resolverAddr)
		require.NoError(t, err)
		return out, nil
	}

	c, err := ens.NewClient(context.Background(), backend, ens.WithChainID(ens.ChainIDMainnet))
	require.NoError(t, err)
	recs, err := c.Records(context.Background(), "example.eth", ens.RecordsRequest{
		Texts: []string{"description"}, Coins: []uint64{60}, Contenthash: true, Name: true,
	})
	require.NoError(t, err)
	require.Equal(t, resolverAddr, recs.Resolver)
	require.Equal(t, "hello", recs.Texts["description"])
	require.Equal(t, common.HexToAddress("0x02").Bytes(), recs.Coins[60])
	require.Nil(t, recs.Contenthash)
	require.Empty(t, recs.Name)
	require.Len(t, recs.Errors, 1)
	require.ErrorIs(t, recs.Errors["name"], ens.ErrUnsupportedProfile)
	require.Len(t, backend.calls, 1, "one eth_call for the whole batch")
}

// TestResolveAddressThroughFakeBackend checks the addr(bytes32) path and
// the ErrNoAddress mapping for empty and zero results.
func TestResolveAddressThroughFakeBackend(t *testing.T) {
	ur := universalresolver.NewUniversalResolverV2()
	want := common.HexToAddress("0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045")
	var answer []byte
	backend := &fakeBackend{answer: func(ethereum.CallMsg) ([]byte, error) { return answer, nil }}
	c, err := ens.NewClient(context.Background(), backend, ens.WithChainID(ens.ChainIDMainnet))
	require.NoError(t, err)

	pack := func(result []byte) []byte {
		out, err := ur.GetABI().Methods["resolve"].Outputs.Pack(result, common.HexToAddress("0x01"))
		require.NoError(t, err)
		return out
	}
	answer = pack(common.LeftPadBytes(want.Bytes(), 32))
	got, err := c.ResolveAddress(context.Background(), "vitalik.eth")
	require.NoError(t, err)
	require.Equal(t, want, got)

	answer = pack([]byte{})
	_, err = c.ResolveAddress(context.Background(), "vitalik.eth")
	require.ErrorIs(t, err, ens.ErrNoAddress)

	answer = pack(make([]byte, 32))
	_, err = c.ResolveAddress(context.Background(), "vitalik.eth")
	require.ErrorIs(t, err, ens.ErrNoAddress)

	// The DNS-encoded name in the request is normalised.
	last := backend.calls[len(backend.calls)-1]
	args, err := ur.GetABI().Methods["resolve"].Inputs.Unpack(last.Data[4:])
	require.NoError(t, err)
	require.Equal(t, "07766974616c696b0365746800", hex.EncodeToString(args[0].([]byte)))
}
