// Copyright 2026 Weald Technology Trading.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0

package ens_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	ens "github.com/ensdomains/go-ens/v5"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

type deploymentsFile struct {
	ChainID   uint64 `json:"chainId"` //nolint:tagliatelle // matches the vendored artifacts
	Contracts map[string]struct {
		Address         common.Address `json:"address"`
		ConstructorArgs map[string]any `json:"constructorArgs"` //nolint:tagliatelle // matches the vendored artifacts
	} `json:"contracts"`
}

func loadDeployments(t *testing.T, network string) deploymentsFile {
	t.Helper()
	raw, err := os.ReadFile(filepath.Join("contracts", "v2", "deployments", network+".json"))
	require.NoError(t, err)
	var d deploymentsFile
	require.NoError(t, json.Unmarshal(raw, &d))
	return d
}

// TestChainConfigMatchesVendoredDeployments pins every V2 address in the
// baked-in tables to the vendored contracts-v2 deployment artifacts, and the
// v1 addresses to the constructor arguments the v2 contracts were deployed
// with, so neither can drift silently.
func TestChainConfigMatchesVendoredDeployments(t *testing.T) {
	for _, tc := range []struct {
		network string
		cfg     ens.ChainConfig
	}{
		{"mainnet", ens.MainnetConfig()},
		{"sepolia", ens.SepoliaConfig()},
	} {
		t.Run(tc.network, func(t *testing.T) {
			d := loadDeployments(t, tc.network)
			require.Equal(t, tc.cfg.ChainID, d.ChainID)
			v2 := map[string]common.Address{
				"UpgradableUniversalResolverProxy": tc.cfg.V2.UniversalResolver,
				"ManagedUniversalResolverProxy":    tc.cfg.V2.ManagedUniversalResolver,
				"UniversalResolverV2":              tc.cfg.V2.UniversalResolverV2,
				"UniversalHelper":                  tc.cfg.V2.UniversalHelper,
				"RootRegistry":                     tc.cfg.V2.RootRegistry,
				"ETHRegistry":                      tc.cfg.V2.ETHRegistry,
				"ETHRegistrar":                     tc.cfg.V2.ETHRegistrar,
				"ETHRenewerV1":                     tc.cfg.V2.ETHRenewerV1,
				"StandardRentPriceOracle":          tc.cfg.V2.RentPriceOracle,
				"VerifiableFactory":                tc.cfg.V2.VerifiableFactory,
				"UserRegistryImpl":                 tc.cfg.V2.UserRegistryImpl,
				"WrapperRegistryImpl":              tc.cfg.V2.WrapperRegistryImpl,
				"PermissionedResolverImpl":         tc.cfg.V2.PermissionedResolverImpl,
				"PublicResolverV2":                 tc.cfg.V2.PublicResolverV2,
				"ENSV1Resolver":                    tc.cfg.V2.ENSV1Resolver,
				"ENSV2Resolver":                    tc.cfg.V2.ENSV2Resolver,
				"DNSTLDResolver":                   tc.cfg.V2.DNSTLDResolver,
				"LabelStore":                       tc.cfg.V2.LabelStore,
				"ReverseRegistrarAdapter":          tc.cfg.V2.ReverseRegistrarAdapter,
				"DefaultReverseRegistrarAdapter":   tc.cfg.V2.DefaultReverseRegistrarAdapter,
				"BatchRegistrar":                   tc.cfg.V2.BatchRegistrar,
				"MigrationHelper":                  tc.cfg.V2.MigrationHelper,
				"LockedMigrationController":        tc.cfg.V2.LockedMigrationController,
				"UnlockedMigrationController":      tc.cfg.V2.UnlockedMigrationController,
				"Graveyard":                        tc.cfg.V2.Graveyard,
			}
			for name, want := range v2 {
				entry, ok := d.Contracts[name]
				require.True(t, ok, "%s missing from vendored deployments", name)
				require.Equal(t, entry.Address, want, "%s address drifted", name)
			}

			// v1 contracts cross-checked against constructor arguments.
			arg := func(contract, field string) common.Address {
				entry, ok := d.Contracts[contract]
				require.True(t, ok, "%s missing", contract)
				v, ok := entry.ConstructorArgs[field]
				require.True(t, ok, "%s has no constructor arg %s", contract, field)
				s, ok := v.(string)
				require.True(t, ok)
				return common.HexToAddress(s)
			}
			require.Equal(t, arg("ETHRenewerV1", "nameWrapper"), tc.cfg.V1.NameWrapper, "NameWrapper")
			require.Equal(t, arg("ETHRenewerV1", "wrappedController"), tc.cfg.V1.ETHRegistrarController, "ETHRegistrarController")
			require.Equal(t, arg("ReverseRegistrarAdapter", "reverseRegistrar"), tc.cfg.V1.ReverseRegistrar, "ReverseRegistrar")
			require.Equal(t, arg("DefaultReverseRegistrarAdapter", "defaultReverseRegistrar"), tc.cfg.V1.DefaultReverseRegistrar, "DefaultReverseRegistrar")
			require.Equal(t, arg("ETHRegistry", "labelStore"), tc.cfg.V2.LabelStore, "LabelStore via ETHRegistry")
			require.Equal(t, arg("UniversalResolverV2", "rootRegistry"), tc.cfg.V2.RootRegistry, "RootRegistry via UR")
			require.Equal(t, arg("ETHRegistrar", "oracle"), tc.cfg.V2.RentPriceOracle, "oracle via ETHRegistrar")
			require.Equal(t, arg("ETHRegistrar", "ethRegistry"), tc.cfg.V2.ETHRegistry, "ETHRegistry via ETHRegistrar")

			// Payment tokens are the oracle's configured ratios.
			entry := d.Contracts["StandardRentPriceOracle"]
			ratios, ok := entry.ConstructorArgs["paymentRatios"].([]any)
			require.True(t, ok)
			require.Len(t, tc.cfg.PaymentTokens, len(ratios))
			for _, r := range ratios {
				m, ok := r.(map[string]any)
				require.True(t, ok)
				tok := common.HexToAddress(m["paymentToken"].(string))
				_, found := tc.cfg.PaymentTokenByAddress(tok)
				require.True(t, found, "payment token %s not in config", tok.Hex())
			}
			_, ok = tc.cfg.PaymentTokenByAddress(tc.cfg.DefaultPaymentToken)
			require.True(t, ok, "default payment token must be a payment token")
			require.NoError(t, tc.cfg.Validate())
		})
	}
}

func TestConfigForChain(t *testing.T) {
	cfg, ok := ens.ConfigForChain(ens.ChainIDMainnet)
	require.True(t, ok)
	require.Equal(t, common.HexToAddress(ens.UniversalResolverContractAddress), cfg.V2.UniversalResolver)
	_, ok = ens.ConfigForChain(8453)
	require.False(t, ok)
	require.Equal(t, []uint64{1, 11155111}, ens.KnownChains())

	usdc, ok := cfg.PaymentToken("usdc")
	require.True(t, ok)
	require.Equal(t, uint8(6), usdc.Decimals)
	_, ok = cfg.PaymentToken("WETH")
	require.False(t, ok)
}
