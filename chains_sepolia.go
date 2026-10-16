// Copyright 2026 Weald Technology Trading.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0

package ens

import "github.com/ethereum/go-ethereum/common"

// SepoliaConfig returns the Sepolia testnet configuration.
//
// v2 addresses come from contracts-v2 deployments/sepolia (the current
// namespace; the dated directories beside it are immutable archives) and are
// cross-checked by chains_test.go. The v1 set is taken from ensjs
// clients/chain.ts, which is the only published source for it. Payment
// tokens are the mock USDC and DAI the Sepolia oracle accepts, plus Circle USDC.
func SepoliaConfig() ChainConfig {
	usdc := common.HexToAddress("0x240b0316df57887dbbe58b586508b19e633a14aa")
	return ChainConfig{
		ChainID: ChainIDSepolia,
		Name:    "sepolia",
		V1: V1Contracts{
			Registry:                common.HexToAddress("0x00000000000C2E074eC69A0dFb2997BA6C7d2e1e"),
			BaseRegistrar:           common.HexToAddress("0x57f1887a8BF19b14fC0dF6Fd9B2acc9Af147eA85"),
			ETHRegistrarController:  common.HexToAddress("0xfed6a969aaa60e4961fcd3ebf1a2e8913ac65b72"), // the wrapped controller ETHRenewerV1 syncs with; ensjs lists an older one
			NameWrapper:             common.HexToAddress("0x0635513f179D50A207757E05759CbD106d7dFcE8"),
			PublicResolver:          common.HexToAddress("0x8FADE66B79cC9f707aB26799354482EB93a5B7dD"),
			ReverseRegistrar:        common.HexToAddress("0xA0a1AbcDAe1a2a4A2EF8e9113Ff0e02DD81DC0C6"),
			DefaultReverseRegistrar: common.HexToAddress("0x4F382928805ba0e23B30cFB75fC9E848e82DFD47"),
			DefaultReverseResolver:  common.HexToAddress("0x9DC60e7BD81cCc96774C55214fF389D42Ae5E9Ac"),
			BulkRenewal:             common.HexToAddress("0x7f86d816165BaF4fD68bFd9A0706601cDD666ac4"),
			DNSRegistrar:            common.HexToAddress("0x5a07C75Ae469Bf3ee2657B588e8E6ABAC6741b4f"),
			DNSSECImpl:              common.HexToAddress("0xe62E4b6cE018Ad6e916fcC24545e20a33b9d8653"),
		},
		V2: V2Contracts{
			UniversalResolver:        common.HexToAddress(UniversalResolverContractAddress),
			ManagedUniversalResolver: common.HexToAddress("0x6d80f2172cfdec5730fe683860c33d26fc42e6f1"),
			UniversalResolverV2:      common.HexToAddress("0x24e1d8e068620b647ca097f961a61055f4f42d72"),
			UniversalHelper:          common.HexToAddress("0xd453e5bdb62cc3bea84341b1e306319c8ffd7dfe"),

			RootRegistry:    common.HexToAddress("0xb458d6a3a77919449d03e7a6903c26827c1ec43f"),
			ETHRegistry:     common.HexToAddress("0xd4ebcbbdf463c9c45784603db0ddd499bc44a8b4"),
			ETHRegistrar:    common.HexToAddress("0xf633e7fc17e2bbe0d0965d18ec1821dcb754a3d3"),
			ETHRenewerV1:    common.HexToAddress("0xf2ece44980778966b8a0fccb3a9e339440f6e045"),
			RentPriceOracle: common.HexToAddress("0x8196665d4ca7488b6474a9ec8e7d2719fb42263a"),

			VerifiableFactory:        common.HexToAddress("0xda70306c98e97ece36f997a21368e53298572991"),
			UserRegistryImpl:         common.HexToAddress("0x9bd8a88719068d09ecee662f36c0e3856708366a"),
			WrapperRegistryImpl:      common.HexToAddress("0xbe768b63e5fbbfbb0ae97e9064e0002df8001880"),
			PermissionedResolverImpl: common.HexToAddress("0x115eb53f0c60696633855f90b138178fb40b2b2c"),
			PublicResolverV2:         common.HexToAddress("0xdc4a563d00f5c3012b699794eb9e13a561be386f"),

			ENSV1Resolver:  common.HexToAddress("0x322b7581ca210a69c6d0e0d7c88a7688d2789cb0"),
			ENSV2Resolver:  common.HexToAddress("0x1cf3989ed3e5ec3cb1d731fc3777323813b61acf"),
			DNSTLDResolver: common.HexToAddress("0x0c9f5e9ae61165140b49919f0df13c0a6642e80d"),
			LabelStore:     common.HexToAddress("0xed8246ff02203a4d4cb262bd78beaa7408a57cae"),

			ReverseRegistrarAdapter:        common.HexToAddress("0x56bce5e727faa9d237341bb5b9e8a03d5919779d"),
			DefaultReverseRegistrarAdapter: common.HexToAddress("0x36f97328e843e37520cbf530e9402791c2754066"),

			BatchRegistrar:              common.HexToAddress("0x4a4c8b7cdab6b19dc2cdb417cdb53a2ccbaf5322"),
			MigrationHelper:             common.HexToAddress("0xa8f86ee5cdd28703bd876f3a8c10b1de70f36899"),
			LockedMigrationController:   common.HexToAddress("0x6029a063d69b09d23c52a754a90e4fe43adac3a8"),
			UnlockedMigrationController: common.HexToAddress("0x2a35b94df22cc7354570be2284655e2cdc0e64a2"),
			Graveyard:                   common.HexToAddress("0xb58a90a39d13cce1d0e192b5da5c47640855b04d"),
		},
		PaymentTokens: []PaymentToken{
			{Symbol: "USDC", Address: usdc, Decimals: 6},
			{Symbol: "DAI", Address: common.HexToAddress("0xf6fac8a58a0be13b9197f27c41b73162fe32572b"), Decimals: 18},
			// Circle's Sepolia USDC is accepted too; the mock is the default
			// because it exposes mint() for tests.
			{Symbol: "CircleUSDC", Address: common.HexToAddress("0x1c7D4B196Cb0C7B01d743Fbc6116a902379C7238"), Decimals: 6},
		},
		DefaultPaymentToken: usdc,
	}
}
