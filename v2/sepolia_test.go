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
	"crypto/rand"
	"errors"
	"encoding/hex"
	"math/big"
	"os"
	"testing"
	"time"

	ens "github.com/ensdomains/go-ens/v5"
	"github.com/ensdomains/go-ens/v5/internal/livetest"
	v1 "github.com/ensdomains/go-ens/v5/v1"
	v2 "github.com/ensdomains/go-ens/v5/v2"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

// Sepolia fixtures from contracts-v2 deployments/sepolia/fixtures.md: names
// the pre-migration seeded as RESERVED on the v2 .eth registry.
var sepoliaReservedFixtures = []string{
	"bnmig-0601-fe-address-001-r04", // unwrapped v1 name
	"bnmig-0665-fe-perm-001-r04",    // wrapped, locked v1 name
}

func TestSepoliaCutoverStatus(t *testing.T) {
	c := livetest.Sepolia(t)
	ok, err := c.IsENSv2(context.Background())
	require.NoError(t, err)
	t.Logf("sepolia %s isENSv2=%v", c.Config().V2.UniversalResolver.Hex(), ok)
	v2c := v2.NewClient(c)
	open, err := v2c.RegistrationOpen(context.Background())
	require.NoError(t, err)
	t.Logf("sepolia v2 registration open: %v", open)
}

func TestSepoliaV2Reads(t *testing.T) {
	c := livetest.SepoliaV2(t)
	ctx := context.Background()
	ok, err := c.IsENSv2(ctx)
	require.NoError(t, err)
	require.True(t, ok)
	v2c := v2.NewClient(c)
	for _, label := range sepoliaReservedFixtures {
		st, err := v2c.ETHRegistry().State(ctx, label)
		require.NoError(t, err)
		t.Logf("%s.eth: %s expiry=%s", label, st.Status, st.Expiry)
		ms, err := c.MigrationStatus(ctx, label+".eth")
		require.NoError(t, err)
		t.Logf("%s.eth migration: %s wrapped=%v locked=%v", label, ms.State, ms.Wrapped, ms.Locked)
	}
	parent, label, err := v2c.ETHRegistry().Parent(ctx)
	require.NoError(t, err)
	require.Equal(t, c.Config().V2.RootRegistry, parent)
	require.Equal(t, "eth", label)
}

// TestSepoliaWritePath registers a throwaway name with mock USDC when a
// funded key is supplied through GO_ENS_TEST_SEPOLIA_KEY and v2
// registration has been enabled on Sepolia. It deploys a resolver with
// records, registers, reads back through UniversalResolverV2, creates a
// subname and unregisters it again.
func TestSepoliaWritePath(t *testing.T) {
	keyHex := os.Getenv(livetest.SepoliaKeyEnv)
	if keyHex == "" {
		t.Skipf("%s not set", livetest.SepoliaKeyEnv)
	}
	key, err := crypto.HexToECDSA(keyHex)
	require.NoError(t, err)
	root := livetest.SepoliaV2(t)
	ctx := context.Background()
	c := v2.NewClient(root)
	open, err := c.RegistrationOpen(ctx)
	require.NoError(t, err)
	if !open {
		t.Skip("v2 registration is not enabled on Sepolia yet")
	}
	opts, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(int64(ens.ChainIDSepolia)))
	require.NoError(t, err)
	me := opts.From
	wait := c.DefaultWait()

	var rnd [6]byte
	_, err = rand.Read(rnd[:])
	require.NoError(t, err)
	label := "go-ens-" + hex.EncodeToString(rnd[:])
	name := label + ".eth"

	// Resolver with initial records.
	salt, err := v2.RandomSalt()
	require.NoError(t, err)
	resolverAddr, err := c.Factory().PredictProxyAddress(ctx, me, salt)
	require.NoError(t, err)
	pr := c.PermissionedResolver(resolverAddr)
	setAddr, err := pr.BuildSetAddress(name, ens.CoinTypeETH, me.Bytes())
	require.NoError(t, err)
	tx, _, err := c.DeployPermissionedResolver(ctx, opts, v2.DeployResolverParams{Salt: salt, InitCalls: []ens.Call{setAddr}})
	require.NoError(t, err)
	_, err = wait(ctx, tx)
	require.NoError(t, err)

	// Mock USDC: mint, approve.
	price, err := c.ETHRegistrar().RegisterPrice(ctx, label, 28*24*time.Hour, common.Address{})
	require.NoError(t, err)
	tx, err = c.Token(price.Token).Mint(opts, me, price.Total())
	require.NoError(t, err)
	_, err = wait(ctx, tx)
	require.NoError(t, err)
	tx, err = c.EnsureAllowance(ctx, opts, price.Token, c.Config().V2.ETHRegistrar, price.Total())
	require.NoError(t, err)
	if tx != nil {
		_, err = wait(ctx, tx)
		require.NoError(t, err)
	}

	// Commit, wait, register.
	secret, err := v2.NewSecret()
	require.NoError(t, err)
	params := v2.RegistrationParams{Commitment: v2.Commitment{Label: label, Owner: me, Secret: secret, Resolver: resolverAddr, Duration: 28 * 24 * time.Hour}}
	tx, h, err := c.ETHRegistrar().Commit(opts, params.Commitment)
	require.NoError(t, err)
	_, err = wait(ctx, tx)
	require.NoError(t, err)
	require.NoError(t, c.ETHRegistrar().WaitForCommitment(ctx, h))
	tx, err = c.ETHRegistrar().Register(ctx, opts, params)
	require.NoError(t, err)
	res, err := c.ETHRegistrar().WaitRegistered(ctx, tx)
	require.NoError(t, err)
	require.Equal(t, label, res.Label)

	addr, err := root.ResolveAddress(ctx, name)
	require.NoError(t, err)
	require.Equal(t, me, addr)
	ms, err := root.MigrationStatus(ctx, name)
	require.NoError(t, err)
	require.Equal(t, ens.MigrationStateNativeV2, ms.State)

	// Primary name through the v1 reverse registrar (the EOA path).
	rr, err := v1.NewClient(root).ReverseRegistrar()
	require.NoError(t, err)
	tx, err = rr.SetName(opts, name)
	require.NoError(t, err)
	_, err = wait(ctx, tx)
	require.NoError(t, err)
	eventually(t, func() bool {
		primary, err := root.ReverseResolve(ctx, me)
		return err == nil && primary == name
	}, "primary name via v1 registrar")

	// Primary name through the v2 adapters, stored in the PermissionedResolver.
	claimOK, _, err := c.ReverseAdapters().IsEnabled(ctx)
	require.NoError(t, err)
	if claimOK {
		txs, err := c.SetPrimaryNameWithPermissionedResolver(ctx, opts, me, resolverAddr, name, wait)
		require.NoError(t, err)
		require.Len(t, txs, 2)
		eventually(t, func() bool {
			primary, err := root.ReverseResolve(ctx, me)
			return err == nil && primary == name
		}, "primary name via adapters")
		res, err := root.Reverse(ctx, me.Bytes(), ens.CoinTypeETH)
		require.NoError(t, err)
		t.Logf("reverse via adapters: resolver=%s reverseResolver=%s", res.Resolver.Hex(), res.ReverseResolver.Hex())
	} else {
		t.Log("reverse adapters not enabled on Sepolia; skipped the adapter path")
	}

	// Scoped resolver roles: a delegate may set one text key only.
	delegate := common.HexToAddress("0x000000000000000000000000000000000000d00d")
	scope := v2.SetterScope{Kind: v2.SetterText, Key: "com.example"}
	tx, err = pr.GrantSetterRoles(opts, scope, delegate)
	require.NoError(t, err)
	_, err = wait(ctx, tx)
	require.NoError(t, err)
	has, err := pr.HasRoles(ctx, scope.Resource(), v2.NewRoles(scope.Role()), delegate)
	require.NoError(t, err)
	require.True(t, has)
	hasRoot, err := pr.HasRootRoles(ctx, v2.NewRoles(scope.Role()), delegate)
	require.NoError(t, err)
	require.False(t, hasRoot)
	tx, err = pr.RevokeScopedRoles(opts, scope, delegate)
	require.NoError(t, err)
	_, err = wait(ctx, tx)
	require.NoError(t, err)
	has, err = pr.HasRoles(ctx, scope.Resource(), v2.NewRoles(scope.Role()), delegate)
	require.NoError(t, err)
	require.False(t, has)

	// Records: set, link to the default record (clears), set again.
	tx, err = pr.SetRecords(opts, name, v2.RecordSet{Texts: map[string]string{"description": "go-ens sepolia"}})
	require.NoError(t, err)
	_, err = wait(ctx, tx)
	require.NoError(t, err)
	eventually(t, func() bool {
		desc, err := root.Text(ctx, name, "description")
		return err == nil && desc == "go-ens sepolia"
	}, "text record visible after SetRecords")
	tx, err = pr.LinkToRecord(opts, name, nil)
	require.NoError(t, err)
	_, err = wait(ctx, tx)
	require.NoError(t, err)
	eventually(t, func() bool {
		desc, err := root.Text(ctx, name, "description")
		return err == nil && desc == ""
	}, "linking to record 0 clears the name's records")
	eventually(t, func() bool {
		_, err := root.ResolveAddress(ctx, name)
		return errors.Is(err, ens.ErrNoAddress)
	}, "address cleared with the record")
	tx, err = pr.SetAddr(opts, name, me)
	require.NoError(t, err)
	_, err = wait(ctx, tx)
	require.NoError(t, err)
	eventually(t, func() bool {
		addr, err := root.ResolveAddress(ctx, name)
		return err == nil && addr == me
	}, "address set again")

	// Subname, then clean up.
	receipts, err := c.CreateSubname(ctx, opts, v2.SubnameParams{Name: "sub." + name, Owner: me, Resolver: resolverAddr}, wait)
	require.NoError(t, err)
	require.NotEmpty(t, receipts)
	n, err := c.Name(ctx, "sub."+name)
	require.NoError(t, err)
	tx, err = n.Unregister(opts)
	require.NoError(t, err)
	_, err = wait(ctx, tx)
	require.NoError(t, err)
	t.Logf("registered and exercised %s", name)
}

// TestSepoliaRenewReservedV1 renews a reserved (unmigrated) v1 fixture through
// ETHRenewerV1, which must extend both the v2 reservation and the v1 registrar
// expiry. Anyone may pay for a renewal, so the fixture's owner is irrelevant.
func TestSepoliaRenewReservedV1(t *testing.T) {
	keyHex := os.Getenv(livetest.SepoliaKeyEnv)
	if keyHex == "" {
		t.Skipf("%s not set", livetest.SepoliaKeyEnv)
	}
	key, err := crypto.HexToECDSA(keyHex)
	require.NoError(t, err)
	root := livetest.SepoliaV2(t)
	ctx := context.Background()
	c := v2.NewClient(root)
	opts, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(int64(ens.ChainIDSepolia)))
	require.NoError(t, err)
	wait := c.DefaultWait()

	const label = "bnmig-0091-leg-rn-003-r04" // "ETHRenewerV1 renews active reserved name in both systems"
	renewer, st, err := c.ChooseRenewer(ctx, label)
	require.NoError(t, err)
	require.Equal(t, v2.StatusReserved, st.Status)
	require.Equal(t, c.Config().V2.ETHRenewerV1, renewer.Address())

	v1c := v1.NewClient(root)
	br, err := v1c.BaseRegistrar()
	require.NoError(t, err)
	v1Before, err := br.Expiry(label + ".eth")
	require.NoError(t, err)

	duration := 24 * time.Hour
	amount, err := renewer.RenewPrice(ctx, label, duration, common.Address{})
	require.NoError(t, err)
	tx, err := c.Token(c.Config().DefaultPaymentToken).Mint(opts, opts.From, amount)
	require.NoError(t, err)
	_, err = wait(ctx, tx)
	require.NoError(t, err)
	tx, err = c.EnsureAllowance(ctx, opts, common.Address{}, renewer.Address(), amount)
	require.NoError(t, err)
	if tx != nil {
		_, err = wait(ctx, tx)
		require.NoError(t, err)
	}
	tx, err = c.Renew(ctx, opts, v2.RenewData{Label: label, Duration: duration}, common.Address{})
	require.NoError(t, err)
	receipt, err := wait(ctx, tx)
	require.NoError(t, err)
	renewals, err := renewer.ParseRenewed(receipt)
	require.NoError(t, err)
	require.Len(t, renewals, 1)
	require.Equal(t, st.Expiry.Add(duration), renewals[0].NewExpiry)

	after, err := c.ETHRegistry().State(ctx, label)
	require.NoError(t, err)
	require.Equal(t, v2.StatusReserved, after.Status, "renewal keeps the reservation")
	require.Equal(t, st.Expiry.Add(duration), after.Expiry)
	v1After, err := br.Expiry(label + ".eth")
	require.NoError(t, err)
	require.Equal(t, 0, new(big.Int).Sub(v1After, v1Before).Cmp(big.NewInt(int64(duration/time.Second))), "v1 registrar expiry extended too")
	t.Logf("renewed %s.eth by %s through ETHRenewerV1 for %s USDC base units", label, duration, amount)
}

// eventually retries a read for a short while: public Sepolia RPCs are
// load-balanced and a node may lag a block behind the one that returned the
// receipt.
func eventually(t *testing.T, cond func() bool, what string) {
	t.Helper()
	deadline := time.Now().Add(90 * time.Second)
	for {
		if cond() {
			return
		}
		if time.Now().After(deadline) {
			t.Fatalf("timed out waiting for: %s", what)
		}
		time.Sleep(3 * time.Second)
	}
}
