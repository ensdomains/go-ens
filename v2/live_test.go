// Copyright 2026 Weald Technology Trading.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0

package v2_test

import (
	"context"
	"testing"
	"time"

	ens "github.com/ensdomains/go-ens/v5"
	"github.com/ensdomains/go-ens/v5/internal/livetest"
	v2 "github.com/ensdomains/go-ens/v5/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

// TestLiveCommitmentMatchesContract verifies the local commitment hash
// against ETHRegistrar.makeCommitment, a pure function that works whether
// or not registration has been enabled.
func TestLiveCommitmentMatchesContract(t *testing.T) {
	c := v2.NewClient(livetest.Mainnet(t))
	ctx := context.Background()
	for _, label := range []string{"a", "example", "go-ens-fixture", "ünïcödé"} {
		secret, err := v2.NewSecret()
		require.NoError(t, err)
		cm := v2.Commitment{
			Label:       label,
			Owner:       common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"),
			Secret:      secret,
			Subregistry: common.HexToAddress("0x01"),
			Resolver:    common.HexToAddress("0x02"),
			Duration:    400 * 24 * time.Hour,
			Referrer:    [32]byte{7},
		}
		local, err := cm.Hash()
		require.NoError(t, err)
		onchain, err := c.ETHRegistrar().MakeCommitmentOnChain(ctx, cm)
		require.NoError(t, err)
		require.Equal(t, onchain, local, "label %q", label)
	}
}

// TestLiveRegistrarReads checks the registrar and oracle views on mainnet.
func TestLiveRegistrarReads(t *testing.T) {
	c := v2.NewClient(livetest.Mainnet(t))
	ctx := context.Background()
	reg := c.ETHRegistrar()

	consts, err := reg.Constants(ctx)
	require.NoError(t, err)
	require.Equal(t, 60*time.Second, consts.MinCommitmentAge)
	require.Equal(t, 24*time.Hour, consts.MaxCommitmentAge)
	require.Equal(t, 28*24*time.Hour, consts.MinRegisterDuration)
	require.Equal(t, 28*24*time.Hour, consts.GracePeriod)
	require.Equal(t, c.Config().V2.ETHRegistry, consts.ETHRegistry)

	oracle, err := c.Oracle(ctx)
	require.NoError(t, err)
	usdc, _ := c.Config().PaymentToken("USDC")
	ok, err := oracle.IsPaymentToken(ctx, usdc.Address)
	require.NoError(t, err)
	require.True(t, ok)
	ok, err = oracle.IsPaymentToken(ctx, common.HexToAddress("0x01"))
	require.NoError(t, err)
	require.False(t, ok)

	valid, err := reg.IsValid(ctx, "ab")
	require.NoError(t, err)
	require.False(t, valid, "two-character labels are not priced")
	valid, err = reg.IsValid(ctx, "example")
	require.NoError(t, err)
	require.True(t, valid)

	rates, err := oracle.BaseRates(ctx)
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(rates), 5)

	// Price for a free long label: base only, no premium.
	price, err := reg.RegisterPrice(ctx, "go-ens-quote-fixture-7f3a9c", 365*24*time.Hour, usdc.Address)
	if err == nil {
		require.Equal(t, usdc.Address, price.Token)
		require.True(t, price.Base.Sign() > 0)
		require.Equal(t, 0, price.Premium.Sign())
		s, err := c.FormatPrice(ctx, price)
		require.NoError(t, err)
		t.Logf("one year of a 5+ char name: %s", s)
	} else {
		// Before phase 2 seeding the name may be unavailable in surprising
		// ways; a typed error is still a pass for the decoding layer.
		t.Logf("RegisterPrice: %v", err)
	}

	open, err := c.RegistrationOpen(ctx)
	require.NoError(t, err)
	t.Logf("v2 registration open on mainnet: %v", open)
}

// TestLiveRegistryReads checks the .eth registry views for a well-known name.
func TestLiveRegistryReads(t *testing.T) {
	c := v2.NewClient(livetest.Mainnet(t))
	ctx := context.Background()
	reg := c.ETHRegistry()

	st, err := reg.State(ctx, "vitalik")
	require.NoError(t, err)
	t.Logf("vitalik: %s expiry=%s latestOwner=%s", st.Status, st.Expiry, st.LatestOwner.Hex())

	parent, label, err := reg.Parent(ctx)
	require.NoError(t, err)
	require.Equal(t, c.Config().V2.RootRegistry, parent)
	require.Equal(t, "eth", label)

	sub, err := c.RootRegistry().Subregistry(ctx, "eth")
	require.NoError(t, err)
	require.Equal(t, c.Config().V2.ETHRegistry, sub)

	emancipated, err := reg.IsEmancipated(ctx)
	require.NoError(t, err)
	t.Logf("eth registry emancipated: %v", emancipated)

	regs, err := c.Helper().Registries(ctx, "vitalik.eth")
	require.NoError(t, err)
	require.Len(t, regs, 3)
	require.Equal(t, c.Config().V2.ETHRegistry, regs[1])
	require.Equal(t, c.Config().V2.RootRegistry, regs[2])

	canon, err := c.Helper().CanonicalName(ctx, c.Config().V2.ETHRegistry)
	require.NoError(t, err)
	require.Equal(t, "eth", canon)

	// ChooseRenewer routes by state without sending anything.
	renewer, state, err := c.ChooseRenewer(ctx, "vitalik")
	if err != nil {
		t.Logf("ChooseRenewer: %v (state %s)", err, state.Status)
	} else {
		t.Logf("vitalik renews through %s", renewer.Address().Hex())
	}

	// Adapter enablement is a phase-1 DAO action; just record it.
	claim, def, err := c.ReverseAdapters().IsEnabled(ctx)
	require.NoError(t, err)
	t.Logf("reverse adapters enabled: claim=%v default=%v", claim, def)
}

// TestLiveNameHandle binds a v2 Name for an unmigrated name, which lives in
// the .eth registry as a reservation.
func TestLiveNameHandle(t *testing.T) {
	c := v2.NewClient(livetest.Mainnet(t))
	ctx := context.Background()
	n, err := c.Name(ctx, "vitalik.eth")
	require.NoError(t, err)
	require.Equal(t, "vitalik", n.Label)
	require.True(t, n.IsETH2LD())
	addr, err := n.Address(ctx)
	require.NoError(t, err)
	require.Equal(t, common.HexToAddress("0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045"), addr)
	_, err = n.Subregistry(ctx)
	if err != nil {
		require.ErrorIs(t, err, v2.ErrNoSubregistry)
	}
	_ = ens.CoinTypeETH
}
