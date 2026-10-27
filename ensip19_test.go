// Copyright 2026 Weald Technology Trading.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0

package ens_test

import (
	"testing"

	ens "github.com/ensdomains/go-ens/v5"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestEVMCoinType(t *testing.T) {
	require.Equal(t, uint64(60), ens.EVMCoinType(1))
	require.Equal(t, uint64(0x8000000a), ens.EVMCoinType(10))       // Optimism
	require.Equal(t, uint64(0x80002105), ens.EVMCoinType(8453))     // Base
	require.Equal(t, uint64(0x80aa36a7), ens.EVMCoinType(11155111)) // Sepolia

	require.True(t, ens.IsEVMCoinType(60))
	require.True(t, ens.IsEVMCoinType(ens.CoinTypeDefault))
	require.True(t, ens.IsEVMCoinType(0x8000000a))
	require.False(t, ens.IsEVMCoinType(0))   // BTC
	require.False(t, ens.IsEVMCoinType(501)) // SOL

	id, ok := ens.ChainIDFromCoinType(60)
	require.True(t, ok)
	require.Equal(t, uint64(1), id)
	id, ok = ens.ChainIDFromCoinType(0x8000000a)
	require.True(t, ok)
	require.Equal(t, uint64(10), id)
	_, ok = ens.ChainIDFromCoinType(ens.CoinTypeDefault)
	require.False(t, ok)
	_, ok = ens.ChainIDFromCoinType(0)
	require.False(t, ok)
}

func TestReverseName(t *testing.T) {
	addr := common.HexToAddress("0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045")
	const hex = "d8da6bf26964af9d7eed9e03e53415d37aa96045"
	require.Equal(t, hex+".addr.reverse", ens.ReverseName(addr.Bytes(), ens.CoinTypeETH))
	require.Equal(t, hex+".default.reverse", ens.ReverseName(addr.Bytes(), ens.CoinTypeDefault))
	require.Equal(t, hex+".8000000a.reverse", ens.ReverseName(addr.Bytes(), ens.EVMCoinType(10)))
	require.Equal(t, hex+".0.reverse", ens.ReverseName(addr.Bytes(), 0))

	// The reverse name must namehash cleanly so it can be resolved.
	_, err := ens.NameHash(ens.ReverseName(addr.Bytes(), ens.EVMCoinType(10)))
	require.NoError(t, err)
}

func TestDNSDecode(t *testing.T) {
	enc, err := ens.DNSEncode("foo.example.eth")
	require.NoError(t, err)
	dec, err := ens.DNSDecode(enc)
	require.NoError(t, err)
	require.Equal(t, "foo.example.eth", dec)

	dec, err = ens.DNSDecode([]byte{0})
	require.NoError(t, err)
	require.Equal(t, "", dec)

	_, err = ens.DNSDecode([]byte{3, 'e', 't'})
	require.Error(t, err)
	_, err = ens.DNSDecode([]byte{3, 'e', 't', 'h'})
	require.Error(t, err)
}
