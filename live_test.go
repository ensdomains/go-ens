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
	"errors"
	"strings"
	"testing"

	ens "github.com/ensdomains/go-ens/v5"
	"github.com/ensdomains/go-ens/v5/internal/livetest"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

// Mainnet fixtures that hold under both the legacy and the v2 Universal
// Resolver.
var (
	vitalikAddr = common.HexToAddress("0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045")
	nickAddr    = common.HexToAddress("0xb8c2C29ee19D8307cb7255e1Cd9CbDE883A267d5")
)

// TestLiveClientReads exercises the version-agnostic read path through the
// canonical proxy, asserting only facts true before and after the cutover.
func TestLiveClientReads(t *testing.T) {
	c := livetest.Mainnet(t)
	ctx := context.Background()

	addr, err := c.ResolveAddress(ctx, "vitalik.eth")
	require.NoError(t, err)
	require.Equal(t, vitalikAddr, addr)

	name, err := c.ReverseResolve(ctx, nickAddr)
	require.NoError(t, err)
	require.Equal(t, "nick.eth", name)

	_, err = c.ReverseResolve(ctx, common.Address{})
	require.ErrorIs(t, err, ens.ErrNoPrimaryName)

	require.Equal(t, "nick.eth", c.Format(ctx, nickAddr))

	// CCIP-Read through an offchain resolver.
	addr, err = c.ResolveAddress(ctx, "test.offchaindemo.eth")
	require.NoError(t, err)
	require.Equal(t, strings.ToLower("0x779981590E7Ccc0CFAe8040Ce7151324747cDb97"), strings.ToLower(addr.Hex()))

	_, err = c.ResolveAddress(ctx, "this-name-is-not-registered-9f8e7d6c.eth")
	require.ErrorIs(t, err, ens.ErrUnregistered)

	recs, err := c.Records(ctx, "vitalik.eth", ens.RecordsRequest{Texts: []string{"url"}, Coins: []uint64{ens.CoinTypeETH}, Contenthash: true})
	require.NoError(t, err)
	require.Equal(t, vitalikAddr.Bytes(), recs.Coins[ens.CoinTypeETH])
	require.Empty(t, recs.Errors)
}

// TestLiveClientV2Direct exercises UniversalResolverV2 directly, which works
// before and after the DAO repoints the canonical proxy.
func TestLiveClientV2Direct(t *testing.T) {
	c := livetest.MainnetV2(t)
	ctx := context.Background()

	v2, err := c.IsENSv2(ctx)
	require.NoError(t, err)
	require.True(t, v2, "UniversalResolverV2 must report isENSv2")

	addr, err := c.ResolveAddress(ctx, "vitalik.eth")
	require.NoError(t, err)
	require.Equal(t, vitalikAddr, addr)

	// An unmigrated name is served through the v1 mirror; the writable
	// resolver is the one behind it.
	info, err := c.WritableResolver(ctx, "vitalik.eth")
	require.NoError(t, err)
	require.Equal(t, c.Config().V2.ENSV1Resolver, info.Resolver)
	require.True(t, info.Composite)
	require.NotEqual(t, info.Resolver, info.Underlying)
	require.Equal(t, ens.ResolverFamilyNodeKeyed, info.Family)
	require.True(t, info.Writable)

	res, err := c.Reverse(ctx, nickAddr.Bytes(), ens.CoinTypeETH)
	require.NoError(t, err)
	require.Equal(t, "nick.eth", res.Name)

	// The bare TLD has no resolver on v2.
	_, err = c.FindResolver(ctx, "eth")
	require.ErrorIs(t, err, ens.ErrUnregistered)

	// Default-record reverse lookups are accepted (and may be empty).
	_, err = c.Reverse(ctx, nickAddr.Bytes(), ens.CoinTypeDefault)
	require.NoError(t, err)
}

// TestLiveCutoverStatus records whether the canonical proxy serves v2 yet.
// It never fails; the result is printed so CI summaries show the day the
// cutover lands.
func TestLiveCutoverStatus(t *testing.T) {
	c := livetest.Mainnet(t)
	v2, err := c.IsENSv2(context.Background())
	require.NoError(t, err)
	t.Logf("mainnet %s isENSv2=%v", c.Config().V2.UniversalResolver.Hex(), v2)
}

// TestLiveMigrationStatus reads the migration state of a few well-known
// names from the v2 .eth registry.
func TestLiveMigrationStatus(t *testing.T) {
	c := livetest.Mainnet(t)
	ctx := context.Background()

	st, err := c.MigrationStatus(ctx, "vitalik.eth")
	require.NoError(t, err)
	t.Logf("vitalik.eth: %s owner=%s v1owner=%s wrapped=%v expiry=%s", st.State, st.Owner.Hex(), st.V1Owner.Hex(), st.Wrapped, st.Expiry)
	switch st.State {
	case ens.MigrationStateUnknown:
		t.Skip("pre-migration has not seeded vitalik.eth on the v2 registry yet")
	case ens.MigrationStateV1Reserved, ens.MigrationStateV1ReservedGrace:
		require.NotEqual(t, common.Address{}, st.V1Owner)
	case ens.MigrationStateMigratedFromV1, ens.MigrationStateNativeV2:
		require.NotNil(t, st.TokenID)
	default:
		t.Fatalf("unexpected state %s", st.State)
	}

	_, err = c.MigrationStatus(ctx, "sub.vitalik.eth")
	require.ErrorIs(t, err, ens.ErrNot2LD)
}

// TestLiveOwner checks the UniversalHelper owner lookup.
func TestLiveOwner(t *testing.T) {
	c := livetest.Mainnet(t)
	owner, err := c.Owner(context.Background(), "vitalik.eth")
	require.NoError(t, err)
	t.Logf("vitalik.eth v2 owner: %s", owner.Hex())
	// Before migration the v2 owner is zero (reserved); afterwards it is
	// the migrated owner. Either way the call must succeed.
	_ = owner
}

func TestLiveUnknownChainErrorMessage(t *testing.T) {
	var e error = &ens.UnknownChainError{ChainID: 8453}
	var typed *ens.UnknownChainError
	require.True(t, errors.As(e, &typed))
}
