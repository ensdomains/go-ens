// Copyright 2026 Weald Technology Trading.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0

package ens

import (
	"fmt"
	"strings"
)

// maxDNSLabelBytes is the largest label the DNS wire format can carry in a
// single length octet. ENS does not use RFC 1035 compression pointers, so the
// full 0–255 range is available (unlike classic DNS, which caps labels at 63).
const maxDNSLabelBytes = 255

// DNSEncode converts an ENS name into the DNS wire format used by EIP-137 and
// the Universal Resolver: each label is preceded by a single byte holding its
// length, and the encoding is terminated with a zero byte representing the
// root label. The empty string encodes to a single zero byte.
//
// The name is normalised per ENSIP-15 first. A single trailing dot is
// trimmed so that DNS-style "foo.eth." still encodes. Labels longer than 255
// bytes return *LabelTooLongError: UniversalResolverV2 cannot decode the
// bracketed labelhash form that earlier releases substituted for them.
// ENSIP-10 removes RFC 1035's 255-octet limit on the total encoded length.
func DNSEncode(name string) ([]byte, error) {
	if name == "" {
		return []byte{0}, nil
	}
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
			return nil, fmt.Errorf("empty label at position %d in %q: %w", i, name, ErrEmptyLabel)
		}
		if len(label) > maxDNSLabelBytes {
			return nil, &LabelTooLongError{Label: label}
		}
		out = append(out, byte(len(label)))
		out = append(out, label...)
	}
	out = append(out, 0)
	return out, nil
}

// DNSDecode converts a DNS wire-format name back into dotted form. It is the
// inverse of DNSEncode for well-formed input and does not normalise.
func DNSDecode(dnsName []byte) (string, error) {
	var labels []string
	i := 0
	for {
		if i >= len(dnsName) {
			return "", fmt.Errorf("dns name is not terminated")
		}
		n := int(dnsName[i])
		i++
		if n == 0 {
			break
		}
		if i+n > len(dnsName) {
			return "", fmt.Errorf("dns label at offset %d runs past the end", i-1)
		}
		labels = append(labels, string(dnsName[i:i+n]))
		i += n
	}
	return strings.Join(labels, "."), nil
}
