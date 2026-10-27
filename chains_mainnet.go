// Copyright 2026 Weald Technology Trading.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0

package ens

import "github.com/ethereum/go-ethereum/common"

// MainnetConfig returns the Ethereum mainnet configuration.
//
// v2 addresses come from contracts-v2 deployments/mainnet (pinned revision in
// contracts/v2/manifest.json; several contracts were replaced on 2026-10-18 and 2026-10-26; v2 went live on 2026-10-26) and are cross-checked by chains_test.go; v1
// addresses are the long-standing ENS v1 deployment.
func MainnetConfig() ChainConfig {
	usdc := common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48")
	return ChainConfig{
		ChainID: ChainIDMainnet,
		Name:    "mainnet",
		V1: V1Contracts{
			Registry:                common.HexToAddress("0x00000000000C2E074eC69A0dFb2997BA6C7d2e1e"),
			BaseRegistrar:           common.HexToAddress("0x57f1887a8BF19b14fC0dF6Fd9B2acc9Af147eA85"),
			ETHRegistrarController:  common.HexToAddress("0x253553366Da8546fC250F225fe3d25d0C782303b"),
			NameWrapper:             common.HexToAddress("0xD4416b13d2b3a9aBae7AcD5D6C2BbDBE25686401"),
			PublicResolver:          common.HexToAddress("0x231b0Ee14048e9dCcD1d247744d114a4EB5E8E63"),
			ReverseRegistrar:        common.HexToAddress("0xa58E81fe9b61B5c3fE2AFD33CF304c454AbFc7Cb"),
			DefaultReverseRegistrar: common.HexToAddress("0x283F227c4Bd38ecE252C4Ae7ECE650B0e913f1f9"),
			DefaultReverseResolver:  common.HexToAddress("0xA7d635c8de9a58a228AA69353a1699C7Cc240DCF"),
			BulkRenewal:             common.HexToAddress("0xa12159e5131b1eEf6B4857EEE3e1954744b5033A"),
			DNSRegistrar:            common.HexToAddress("0xB32cB5677a7C971689228EC835800432B339bA2B"),
			DNSSECImpl:              common.HexToAddress("0x0fc3152971714E5ed7723FAFa650F86A4BaF30C5"),
		},
		V2: V2Contracts{
			UniversalResolver:        common.HexToAddress(UniversalResolverContractAddress),
			ManagedUniversalResolver: common.HexToAddress("0xcffd2ae680cd110d28739a70922636b5937b022b"),
			UniversalResolverV2:      common.HexToAddress("0x06fdd07f007adac667177b46ab2eb9a75132cfd5"),
			UniversalHelper:          common.HexToAddress("0x46ae127d14683ac74ef151e46ae2117a1a25c676"),

			RootRegistry:    common.HexToAddress("0x2553c09be6447beffdb86fa622ad41fb8f77a91e"),
			ETHRegistry:     common.HexToAddress("0xf60aba43e1f7f3dab1186ff03b561401c9df85e2"),
			ETHRegistrar:    common.HexToAddress("0xd6d8369e55c94ae8afa714e6cf7ab5be54d5cc81"),
			ETHRenewerV1:    common.HexToAddress("0x56c447f308f463a01121c41e72b0e7e7557fcb68"),
			RentPriceOracle: common.HexToAddress("0x76f29609a363a8a1a2a6abe2e859c0f8916af469"),

			VerifiableFactory:        common.HexToAddress("0x77510ba162dfea3d8989e3f0928a691b345730d2"),
			UserRegistryImpl:         common.HexToAddress("0xfaa23592aca10a6d8875b18ea86954e8eb6d8d72"),
			WrapperRegistryImpl:      common.HexToAddress("0xe86a4ef4d7a210aaa1d6a9594fc7add7cd2cf888"),
			PermissionedResolverImpl: common.HexToAddress("0x565b0f456a04ac9b48758190e90030623cff57c5"),
			PublicResolverV2:         common.HexToAddress("0x66f9f566ff2f26309c1bd9134388478a801da3e1"),

			ENSV1Resolver:  common.HexToAddress("0x394d8eb4a4fdf599384deab25f31d111d755c398"),
			ENSV2Resolver:  common.HexToAddress("0x3071ad643ce4ed8b3d67b289845b45db82b068a4"),
			DNSTLDResolver: common.HexToAddress("0xab48304a2520ce420e636aacd48a8bb3e7e485ee"),
			LabelStore:     common.HexToAddress("0x31e5e081fe38d15b3e8214ba5cb08dc4708b70ba"),

			ReverseRegistrarAdapter:        common.HexToAddress("0x349151735790269a9c14668b58d7c83b5ecb9607"),
			DefaultReverseRegistrarAdapter: common.HexToAddress("0xd86e32499ecafe9c181dd7410f93a1a9bb61ef0d"),

			BatchRegistrar:              common.HexToAddress("0xf0ccfbf6dba1e01ccb833ced6d551a8f000e9721"),
			MigrationHelper:             common.HexToAddress("0x6c8452afc6043b5af32eb345a8135108906f07ee"),
			LockedMigrationController:   common.HexToAddress("0x3cea18620883a69ed349b23328a3b6302598e51a"),
			UnlockedMigrationController: common.HexToAddress("0xab821ca676cf2e0790fa273526a165cfa86aeb64"),
			Graveyard:                   common.HexToAddress("0xda021a2c1c2f99574d861ef8c6fa745851beb7f6"),
		},
		PaymentTokens: []PaymentToken{
			{Symbol: "USDC", Address: usdc, Decimals: 6},
			{Symbol: "DAI", Address: common.HexToAddress("0x6B175474E89094C44Da98b954EedeAC495271d0F"), Decimals: 18},
		},
		DefaultPaymentToken: usdc,
	}
}
