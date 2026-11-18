// Copyright 2026 Weald Technology Trading.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0

package v2

import (
	"math/big"

	ens "github.com/ensdomains/go-ens/v5"
)

// MaxExpiry is the largest expiry a registry accepts; subnames registered
// with it never expire.
const MaxExpiry = ^uint64(0)

var low32 = new(big.Int).SetUint64(0xffffffff)

// LabelID returns the id a registry derives from a label: the labelhash as
// a uint256 (LibLabel.id). The label is ENSIP-15 normalised first. Registry
// functions that take an "anyId" accept this value, a token id or a
// resource, and mask the low 32 bits themselves.
func LabelID(label string) (*big.Int, error) {
	h, err := ens.LabelHash(label)
	if err != nil {
		return nil, err
	}
	return new(big.Int).SetBytes(h[:]), nil
}

// CanonicalID strips the version from any registry id: the low 32 bits
// are cleared (LibLabel.withVersion(id, 0)). Token ids and resources of the
// same name share a canonical id.
func CanonicalID(id *big.Int) *big.Int {
	return new(big.Int).AndNot(id, low32)
}

// WithVersion returns the id with its low 32 bits replaced by version.
func WithVersion(id *big.Int, version uint32) *big.Int {
	return new(big.Int).Or(CanonicalID(id), new(big.Int).SetUint64(uint64(version)))
}

// VersionOf returns the version held in the low 32 bits of a token id or
// resource.
func VersionOf(id *big.Int) uint32 {
	return uint32(new(big.Int).And(id, low32).Uint64())
}

// SameName reports whether two ids refer to the same name regardless of
// version.
func SameName(a, b *big.Int) bool {
	return CanonicalID(a).Cmp(CanonicalID(b)) == 0
}
