// Copyright 2026 Weald Technology Trading.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0

// Package livetest holds helpers for tests that talk to real networks. Every
// helper skips the test under `go test -short`.
package livetest

import (
	"context"
	"os"
	"testing"

	ens "github.com/ensdomains/go-ens/v5"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
)

// Environment variables that select RPC endpoints.
const (
	MainnetRPCEnv = "GO_ENS_TEST_RPC"
	SepoliaRPCEnv = "GO_ENS_TEST_SEPOLIA_RPC"
	SepoliaKeyEnv = "GO_ENS_TEST_SEPOLIA_KEY"
)

// Public fallbacks that surface revert data, which CCIP-Read requires.
const (
	defaultMainnetRPC = "https://ethereum.publicnode.com"
	defaultSepoliaRPC = "https://ethereum-sepolia-rpc.publicnode.com"
)

func dial(t *testing.T, env, fallback string) *ethclient.Client {
	t.Helper()
	if testing.Short() {
		t.Skip("skipping live-RPC test under -short")
	}
	url := os.Getenv(env)
	if url == "" && env == SepoliaRPCEnv {
		url = os.Getenv("GO_ENS_TEST_SEPOLIA_URL") // accepted alias
	}
	if url == "" {
		url = fallback
	}
	client, err := ethclient.Dial(url)
	require.NoError(t, err, "failed to dial %s", url)
	t.Cleanup(client.Close)
	return client
}

// MainnetRPC dials the mainnet endpoint from GO_ENS_TEST_RPC or a public fallback.
func MainnetRPC(t *testing.T) *ethclient.Client {
	t.Helper()
	return dial(t, MainnetRPCEnv, defaultMainnetRPC)
}

// SepoliaRPC dials the Sepolia endpoint from GO_ENS_TEST_SEPOLIA_RPC or a public fallback.
func SepoliaRPC(t *testing.T) *ethclient.Client {
	t.Helper()
	return dial(t, SepoliaRPCEnv, defaultSepoliaRPC)
}

// Mainnet returns a client for mainnet.
func Mainnet(t *testing.T, opts ...ens.Option) *ens.Client {
	t.Helper()
	c, err := ens.NewClient(context.Background(), MainnetRPC(t), opts...)
	require.NoError(t, err)
	return c
}

// Sepolia returns a client for Sepolia.
func Sepolia(t *testing.T, opts ...ens.Option) *ens.Client {
	t.Helper()
	c, err := ens.NewClient(context.Background(), SepoliaRPC(t), opts...)
	require.NoError(t, err)
	return c
}

// MainnetV2 returns a mainnet client that resolves through UniversalResolverV2
// directly, so v2 behaviour can be exercised before the canonical proxy is
// repointed.
func MainnetV2(t *testing.T, opts ...ens.Option) *ens.Client {
	t.Helper()
	cfg := ens.MainnetConfig()
	return Mainnet(t, append([]ens.Option{ens.WithUniversalResolver(cfg.V2.UniversalResolverV2)}, opts...)...)
}

// SepoliaV2 is MainnetV2 for Sepolia.
func SepoliaV2(t *testing.T, opts ...ens.Option) *ens.Client {
	t.Helper()
	cfg := ens.SepoliaConfig()
	return Sepolia(t, append([]ens.Option{ens.WithUniversalResolver(cfg.V2.UniversalResolverV2)}, opts...)...)
}

// RequireENSv2 skips the test unless the client's Universal Resolver is
// UniversalResolverV2.
func RequireENSv2(t *testing.T, c *ens.Client) {
	t.Helper()
	ok, err := c.IsENSv2(context.Background())
	require.NoError(t, err)
	if !ok {
		t.Skipf("universal resolver at %s is not ENSv2 yet (cutover pending)", c.Config().V2.UniversalResolver.Hex())
	}
}
