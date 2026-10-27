// Copyright 2026 Weald Technology Trading.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0

package v2_test

import (
	"math/big"
	"testing"

	ens "github.com/ensdomains/go-ens/v5"
	v2 "github.com/ensdomains/go-ens/v5/v2"
	"github.com/stretchr/testify/require"
)

func TestTokenIDHelpers(t *testing.T) {
	id, err := v2.LabelID("vitalik")
	require.NoError(t, err)
	h, err := ens.LabelHash("vitalik")
	require.NoError(t, err)
	require.Equal(t, 0, id.Cmp(new(big.Int).SetBytes(h[:])), "LabelID is the labelhash as uint256")

	// Normalisation applies.
	upper, err := v2.LabelID("VITALIK")
	require.NoError(t, err)
	require.Equal(t, 0, id.Cmp(upper))

	canon := v2.CanonicalID(id)
	require.Equal(t, uint32(0), v2.VersionOf(canon))
	require.Equal(t, 0, new(big.Int).Rsh(canon, 32).Cmp(new(big.Int).Rsh(id, 32)), "high bits untouched")

	v7 := v2.WithVersion(id, 7)
	require.Equal(t, uint32(7), v2.VersionOf(v7))
	require.True(t, v2.SameName(id, v7))
	require.True(t, v2.SameName(canon, v7))
	require.Equal(t, 0, v2.CanonicalID(v7).Cmp(canon))

	// LibLabel.withVersion: anyId ^ uint32(anyId) ^ versionId.
	low := new(big.Int).And(id, big.NewInt(0xffffffff))
	expect := new(big.Int).Xor(id, low)
	expect.Xor(expect, big.NewInt(7))
	require.Equal(t, 0, expect.Cmp(v7))

	other, err := v2.LabelID("nick")
	require.NoError(t, err)
	require.False(t, v2.SameName(id, other))
}
