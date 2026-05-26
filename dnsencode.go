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
	"strings"
)

// DNSEncode converts an ENS name into the DNS wire format used by
// EIP-137 / the Universal Resolver: each label is preceded by a
// single byte holding its length, and the encoding is terminated
// with a zero byte representing the root label. The empty string
// encodes to a single zero byte.
//
// Per RFC 1035, a single label may be at most 63 bytes (the top two
// bits of the length octet are reserved for compression pointers) and
// the full encoded name may be at most 255 octets including length
// bytes and the terminating null.
func DNSEncode(name string) ([]byte, error) {
	if name == "" {
		return []byte{0}, nil
	}
	normalised, err := Normalize(name)
	if err != nil {
		return nil, err
	}
	// Trailing dots produce empty labels which are not meaningful in ENS;
	// trim a single trailing dot to be lenient with input.
	normalised = strings.TrimSuffix(normalised, ".")
	parts := strings.Split(normalised, ".")
	out := make([]byte, 0, len(normalised)+len(parts)+1)
	for i, label := range parts {
		if label == "" {
			return nil, fmt.Errorf("empty label at position %d in %q", i, name)
		}
		labelBytes := []byte(label)
		if len(labelBytes) > 63 {
			return nil, errors.New("label exceeds 63 bytes")
		}
		out = append(out, byte(len(labelBytes)))
		out = append(out, labelBytes...)
	}
	// +1 for the terminating null we are about to append.
	if len(out)+1 > 255 {
		return nil, errors.New("encoded name exceeds 255 bytes")
	}
	out = append(out, 0)
	return out, nil
}
