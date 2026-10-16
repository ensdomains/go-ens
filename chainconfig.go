// Copyright 2026 Weald Technology Trading.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0

package ens

import (
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

// Chain IDs with baked-in configuration.
const (
	ChainIDMainnet uint64 = 1
	ChainIDSepolia uint64 = 11155111
)

// UniversalResolverContractAddress is the canonical ENS Universal Resolver
// proxy. The same vanity address is used on Ethereum mainnet and Sepolia and
// it is the DAO-designated entry point for resolution before and after the
// v2 cutover.
const UniversalResolverContractAddress = "0xeEeEEEeE14D718C2B47D9923Deab1335E144EeEe"

// V1Contracts holds the ENS v1 contracts a client may still need: for reads
// of unmigrated names, for primary-name writes (which stay on the v1 reverse
// registrars under v2), and for the v1 package.
type V1Contracts struct {
	Registry                common.Address // ENSRegistry
	BaseRegistrar           common.Address // .eth ERC-721 registrar
	ETHRegistrarController  common.Address // v1 controller; disabled at migration phase 3
	NameWrapper             common.Address
	PublicResolver          common.Address
	ReverseRegistrar        common.Address // addr.reverse
	DefaultReverseRegistrar common.Address // default.reverse (ENSIP-19)
	DefaultReverseResolver  common.Address
	BulkRenewal             common.Address
	DNSRegistrar            common.Address
	DNSSECImpl              common.Address
}

// V2Contracts holds the ENS v2 deployment.
type V2Contracts struct {
	// UniversalResolver is the address a Client resolves through. By default
	// it is the canonical proxy; ManagedUniversalResolver (the intermediate
	// proxy the security council administers) and UniversalResolverV2 (the
	// bare implementation) are exposed so callers can exercise v2 resolution
	// before the DAO repoints the canonical proxy.
	UniversalResolver        common.Address
	ManagedUniversalResolver common.Address
	UniversalResolverV2      common.Address
	UniversalHelper          common.Address

	RootRegistry    common.Address
	ETHRegistry     common.Address
	ETHRegistrar    common.Address
	ETHRenewerV1    common.Address
	RentPriceOracle common.Address // StandardRentPriceOracle at deploy time; ETHRegistrar.rentPriceOracle() is authoritative

	VerifiableFactory        common.Address
	UserRegistryImpl         common.Address
	WrapperRegistryImpl      common.Address
	PermissionedResolverImpl common.Address
	PublicResolverV2         common.Address

	ENSV1Resolver  common.Address // read-only mirror into the v1 registry; never a write target
	ENSV2Resolver  common.Address // read-only mirror into the v2 registry tree
	DNSTLDResolver common.Address
	LabelStore     common.Address

	ReverseRegistrarAdapter        common.Address
	DefaultReverseRegistrarAdapter common.Address

	BatchRegistrar              common.Address
	MigrationHelper             common.Address
	LockedMigrationController   common.Address
	UnlockedMigrationController common.Address
	Graveyard                   common.Address
}

// PaymentToken is an ERC-20 accepted by the v2 rent price oracle.
type PaymentToken struct {
	Symbol   string
	Address  common.Address
	Decimals uint8
}

// ChainConfig is the full address table a Client needs on one chain.
type ChainConfig struct {
	ChainID uint64
	Name    string
	V1      V1Contracts
	V2      V2Contracts
	// PaymentTokens lists the ERC-20s the registrar accepts for registration
	// and renewal. DefaultPaymentToken is used when a caller passes the zero
	// address.
	PaymentTokens       []PaymentToken
	DefaultPaymentToken common.Address
}

// ErrNotConfigured is returned when an operation needs a contract whose
// address is zero in the active ChainConfig.
var ErrNotConfigured = errors.New("contract not configured for this chain")

// Validate reports the fields every Client needs.
func (c ChainConfig) Validate() error {
	if c.ChainID == 0 {
		return errors.New("chain config: chain id is zero")
	}
	if c.V2.UniversalResolver == (common.Address{}) {
		return errors.New("chain config: universal resolver address is zero")
	}
	return nil
}

// PaymentToken looks a payment token up by symbol, case-insensitively.
func (c ChainConfig) PaymentToken(symbol string) (PaymentToken, bool) {
	for _, t := range c.PaymentTokens {
		if strings.EqualFold(t.Symbol, symbol) {
			return t, true
		}
	}
	return PaymentToken{}, false
}

// PaymentTokenByAddress looks a payment token up by address.
func (c ChainConfig) PaymentTokenByAddress(addr common.Address) (PaymentToken, bool) {
	for _, t := range c.PaymentTokens {
		if t.Address == addr {
			return t, true
		}
	}
	return PaymentToken{}, false
}

// ConfigForChain returns the baked-in configuration for a chain ID.
func ConfigForChain(chainID uint64) (ChainConfig, bool) {
	switch chainID {
	case ChainIDMainnet:
		return MainnetConfig(), true
	case ChainIDSepolia:
		return SepoliaConfig(), true
	default:
		return ChainConfig{}, false
	}
}

// KnownChains lists the chain IDs with baked-in configuration, ascending.
func KnownChains() []uint64 {
	ids := []uint64{ChainIDMainnet, ChainIDSepolia}
	sort.Slice(ids, func(i, j int) bool { return ids[i] < ids[j] })
	return ids
}

// UnknownChainError is returned when a backend reports a chain ID without a
// baked-in configuration and the caller supplied none.
type UnknownChainError struct {
	ChainID uint64
}

func (e *UnknownChainError) Error() string {
	return fmt.Sprintf("ens: no configuration for chain ID %d; supply one with WithChainConfig", e.ChainID)
}
