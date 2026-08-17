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
	"fmt"
	"strings"

	"golang.org/x/crypto/sha3"
)

// maxDNSLabelBytes is the largest label the Universal Resolver can store in
// a single DNS length octet. ENS does not use RFC 1035 compression pointers,
// so the full 0–255 range is available (unlike classic DNS, which caps labels
// at 63). Labels longer than this are replaced with an encoded labelhash, the
// same substitution used by viem's packetToBytes and ENSjs.
const maxDNSLabelBytes = 255

// DNSEncode converts an ENS name into the DNS wire format used by
// EIP-137 / the Universal Resolver: each label is preceded by a
// single byte holding its length, and the encoding is terminated
// with a zero byte representing the root label. The empty string
// encodes to a single zero byte.
//
// Labels of at most 255 bytes are encoded as-is. Labels longer than
// 255 bytes are replaced with "[" + hex(keccak256(label)) + "]",
// which the Universal Resolver treats as an unknown-label hash.
// ENSIP-10 removes RFC 1035's 255-octet limit on the total encoded
// name length.
func DNSEncode(name string) ([]byte, error) {
	if name == "" {
		return []byte{0}, nil
	}
	// Trim a single trailing dot before normalisation — ENSIP-15 rejects
	// empty labels, but DNS-style "foo.eth." input is convenient. Multiple
	// trailing dots still produce an empty label and are rejected.
	name = strings.TrimSuffix(name, ".")
	if name == "" {
		return []byte{0}, nil
	}
	normalised, err := Normalize(name)
	if err != nil {
		return nil, err
	}
	parts := strings.Split(normalised, ".")
	out := make([]byte, 0, len(normalised)+len(parts)+1)
	for i, label := range parts {
		if label == "" {
			return nil, fmt.Errorf("empty label at position %d in %q", i, name)
		}
		labelBytes := []byte(label)
		if len(labelBytes) > maxDNSLabelBytes {
			hashed, err := encodeHashedLabel(labelBytes)
			if err != nil {
				return nil, err
			}
			labelBytes = hashed
		}
		out = append(out, byte(len(labelBytes)))
		out = append(out, labelBytes...)
	}
	out = append(out, 0)
	return out, nil
}

// encodeHashedLabel returns the Universal Resolver unknown-label form
// "[" + lowercase hex(keccak256(label)) + "]", matching viem encodeLabelhash.
func encodeHashedLabel(label []byte) ([]byte, error) {
	sha := sha3.NewLegacyKeccak256()
	if _, err := sha.Write(label); err != nil {
		return nil, err
	}
	var hash [32]byte
	sha.Sum(hash[:0])
	return []byte("[" + hex.EncodeToString(hash[:]) + "]"), nil
}
