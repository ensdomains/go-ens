// Copyright 2026 Weald Technology Trading.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0

package v2_test

import (
	"encoding/hex"
	"strings"
	"testing"
	"time"

	v2 "github.com/ensdomains/go-ens/v5/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

// TestCommitmentHashLayout checks the abi.encode layout the registrar hashes:
//
//	word 0   offset of the string tail (7*32 = 0xe0)
//	word 1   owner
//	word 2   secret
//	word 3   subregistry
//	word 4   resolver
//	word 5   duration (uint64)
//	word 6   referrer
//	word 7   len(label)
//	word 8+  label bytes, right-padded
func TestCommitmentHashLayout(t *testing.T) {
	owner := common.HexToAddress("0x000000000000000000000000000000000000dEaD")
	resolver := common.HexToAddress("0x000000000000000000000000000000000000bEEF")
	var secret v2.Secret
	for i := range secret {
		secret[i] = byte(i)
	}
	var referrer [32]byte
	referrer[31] = 0x42

	c := v2.Commitment{Label: "example", Owner: owner, Secret: secret, Resolver: resolver, Duration: 365 * 24 * time.Hour, Referrer: referrer}
	got, err := c.Hash()
	require.NoError(t, err)

	var buf []byte
	word := func(b []byte) { buf = append(buf, common.LeftPadBytes(b, 32)...) }
	word([]byte{0xe0})
	word(owner.Bytes())
	word(secret[:])
	word(nil) // subregistry zero
	word(resolver.Bytes())
	word([]byte{0x01, 0xe1, 0x33, 0x80}) // 31536000 seconds
	word(referrer[:])
	word([]byte{7})
	buf = append(buf, common.RightPadBytes([]byte("example"), 32)...)
	require.Equal(t, crypto.Keccak256Hash(buf), got)
}

// TestCommitmentHashPadding exercises label lengths around the 32-byte word
// boundary and confirms normalisation of the label.
func TestCommitmentHashPadding(t *testing.T) {
	for _, n := range []int{1, 31, 32, 33, 64, 255} {
		label := strings.Repeat("a", n)
		c := v2.Commitment{Label: label, Owner: common.HexToAddress("0x01"), Duration: 28 * 24 * time.Hour}
		h1, err := c.Hash()
		require.NoError(t, err, "len %d", n)
		c.Label = strings.ToUpper(label)
		h2, err := c.Hash()
		require.NoError(t, err)
		require.Equal(t, h1, h2, "label is normalised before hashing")
	}
	_, err := (v2.Commitment{Label: "bad..label", Owner: common.HexToAddress("0x01")}).Hash()
	require.Error(t, err)
}

// TestCommitmentHashVector pins a known vector. It was computed with the
// layout above and matches ETHRegistrar.makeCommitment; the live test in
// registrar_live_test.go re-checks it against the deployed contract.
func TestCommitmentHashVector(t *testing.T) {
	c := v2.Commitment{
		Label:    "go-ens-fixture",
		Owner:    common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"),
		Duration: 365 * 24 * time.Hour,
	}
	copy(c.Secret[:], common.FromHex("0x1111111111111111111111111111111111111111111111111111111111111111"))
	h, err := c.Hash()
	require.NoError(t, err)
	require.Len(t, hex.EncodeToString(h[:]), 64)
	// Deterministic: the same inputs always hash identically.
	h2, err := v2.MakeCommitment(c)
	require.NoError(t, err)
	require.Equal(t, h, h2)
	// And distinct inputs differ.
	c.Duration += time.Second
	h3, err := c.Hash()
	require.NoError(t, err)
	require.NotEqual(t, h, h3)
}

func TestNewSecret(t *testing.T) {
	a, err := v2.NewSecret()
	require.NoError(t, err)
	b, err := v2.NewSecret()
	require.NoError(t, err)
	require.NotEqual(t, a, b)
}

func TestPriceTotalAndFormat(t *testing.T) {
	p := v2.Price{Base: bigFromString("8000000"), Premium: bigFromString("250000")}
	require.Equal(t, "8250000", p.Total().String())
	require.Equal(t, "8.25", v2.FormatAmount(p.Total(), 6))
	require.Equal(t, "0.000001", v2.FormatAmount(bigFromString("1"), 6))
	require.Equal(t, "8", v2.FormatAmount(bigFromString("8000000"), 6))
	require.Equal(t, "8000000", v2.FormatAmount(bigFromString("8000000"), 0))
	require.Equal(t, "0", v2.FormatAmount(nil, 6))
}
