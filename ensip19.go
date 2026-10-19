// Copyright 2026 Weald Technology Trading.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0

package ens

import (
	"encoding/hex"
	"strconv"
)

// CoinTypeDefault is the ENSIP-19 coin type of the default EVM address
// record, which v2 resolvers fall back to for any EVM chain.
const CoinTypeDefault = uint64(0x80000000)

// evmCoinTypeFlag marks ENSIP-11 / ENSIP-19 EVM coin types.
const evmCoinTypeFlag = uint64(0x80000000)

// EVMCoinType returns the SLIP-0044 coin type ENSIP-11 assigns to an EVM
// chain: 60 for Ethereum mainnet, 0x80000000 | chainID otherwise.
func EVMCoinType(chainID uint64) uint64 {
	if chainID == ChainIDMainnet {
		return CoinTypeETH
	}
	return evmCoinTypeFlag | chainID
}

// IsEVMCoinType reports whether a coin type denotes an EVM chain (60, the
// default record, or any value with the ENSIP-11 flag set).
func IsEVMCoinType(coinType uint64) bool {
	if coinType == CoinTypeETH {
		return true
	}
	return coinType&evmCoinTypeFlag != 0 && coinType <= 0xffffffff
}

// ChainIDFromCoinType returns the EVM chain ID a coin type denotes. ok is
// false for non-EVM coin types and for the default record.
func ChainIDFromCoinType(coinType uint64) (uint64, bool) {
	switch {
	case coinType == CoinTypeETH:
		return ChainIDMainnet, true
	case coinType == CoinTypeDefault:
		return 0, false
	case IsEVMCoinType(coinType):
		return coinType &^ evmCoinTypeFlag, true
	default:
		return 0, false
	}
}

// ReverseName returns the ENSIP-19 reverse name for an address and coin type:
// "<hex>.addr.reverse" for coin type 60, "<hex>.default.reverse" for the
// default record and "<hex>.<hex coin type>.reverse" otherwise. The address
// is lowercase hex without a 0x prefix.
func ReverseName(address []byte, coinType uint64) string {
	h := hex.EncodeToString(address)
	switch coinType {
	case CoinTypeETH:
		return h + ".addr.reverse"
	case CoinTypeDefault:
		return h + ".default.reverse"
	default:
		return h + "." + strconv.FormatUint(coinType, 16) + ".reverse"
	}
}
