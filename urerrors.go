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
	"math/big"

	"github.com/ensdomains/go-ens/v5/internal/revert"
	"github.com/ethereum/go-ethereum/common"
)

// ResolverNotFoundError corresponds to the Universal Resolver custom error
// ResolverNotFound(bytes): no resolver is set anywhere up the parent chain
// of the requested name.
type ResolverNotFoundError struct {
	// DNSName is the DNS-encoded name as passed to resolve().
	DNSName []byte
}

func (e *ResolverNotFoundError) Error() string { return "unregistered name" }

// Unwrap returns ErrUnregistered.
func (e *ResolverNotFoundError) Unwrap() error { return ErrUnregistered }

// ResolverNotContractError corresponds to ResolverNotContract(bytes,address):
// the registry pointed at an address that holds no code.
type ResolverNotContractError struct {
	DNSName  []byte
	Resolver common.Address
}

func (e *ResolverNotContractError) Error() string {
	return fmt.Sprintf("resolver address has no code: %s", e.Resolver.Hex())
}

// UnsupportedResolverProfileError corresponds to
// UnsupportedResolverProfile(bytes4): the resolver did not implement the
// requested function selector.
type UnsupportedResolverProfileError struct {
	Selector [4]byte
}

func (e *UnsupportedResolverProfileError) Error() string {
	return fmt.Sprintf("resolver does not implement profile 0x%s", hex.EncodeToString(e.Selector[:]))
}

// Unwrap returns ErrUnsupportedProfile.
func (e *UnsupportedResolverProfileError) Unwrap() error { return ErrUnsupportedProfile }

// ResolverRevertError corresponds to ResolverError(bytes): the downstream
// resolver itself reverted; the inner revert bytes are preserved verbatim.
type ResolverRevertError struct {
	Data []byte
}

func (e *ResolverRevertError) Error() string {
	return fmt.Sprintf("resolver reverted: 0x%s", hex.EncodeToString(e.Data))
}

// Inner decodes the nested revert payload with the Universal Resolver
// decoder. It returns nil when the payload is empty.
func (e *ResolverRevertError) Inner() error {
	if len(e.Data) == 0 {
		return nil
	}
	inner, _ := urDecoder.TranslateData(e.Data)
	return inner
}

// Unwrap exposes the decoded inner error to errors.Is and errors.As.
func (e *ResolverRevertError) Unwrap() error { return e.Inner() }

// EmptyAddressError corresponds to EmptyAddress(): reverse() was given an
// empty address.
type EmptyAddressError struct{}

func (e *EmptyAddressError) Error() string { return "no address" }

// Unwrap returns ErrNoAddress.
func (e *EmptyAddressError) Unwrap() error { return ErrNoAddress }

// HTTPGatewayError corresponds to HttpError(uint16,string): a CCIP-Read
// gateway returned a non-success status whose body the resolver forwarded.
type HTTPGatewayError struct {
	Status  uint16
	Message string
}

func (e *HTTPGatewayError) Error() string {
	return fmt.Sprintf("ccip-read gateway HTTP %d: %s", e.Status, e.Message)
}

// ReverseAddressMismatchError corresponds to
// ReverseAddressMismatch(string,bytes): the reverse-resolved name's forward
// record does not point back at the queried address.
type ReverseAddressMismatchError struct {
	Name    string
	Address []byte
}

func (e *ReverseAddressMismatchError) Error() string {
	return fmt.Sprintf("reverse name %q does not resolve back to 0x%s", e.Name, hex.EncodeToString(e.Address))
}

// DNSDecodingFailedError corresponds to DNSDecodingFailed(bytes): the
// DNS-encoded name was malformed.
type DNSDecodingFailedError struct{ Data []byte }

func (e *DNSDecodingFailedError) Error() string {
	return fmt.Sprintf("dns decoding failed: 0x%s", hex.EncodeToString(e.Data))
}

// DNSEncodingFailedError corresponds to DNSEncodingFailed(string): a name,
// typically a primary name during reverse resolution, could not be
// DNS-encoded on-chain.
type DNSEncodingFailedError struct{ Name string }

func (e *DNSEncodingFailedError) Error() string {
	return fmt.Sprintf("dns encoding failed: %q", e.Name)
}

// LabelIsEmptyError corresponds to LabelIsEmpty().
type LabelIsEmptyError struct{}

func (e *LabelIsEmptyError) Error() string { return "empty label" }

// Unwrap returns ErrEmptyLabel.
func (e *LabelIsEmptyError) Unwrap() error { return ErrEmptyLabel }

// InvalidBatchGatewayResponseError corresponds to InvalidBatchGatewayResponse().
type InvalidBatchGatewayResponseError struct{}

func (e *InvalidBatchGatewayResponseError) Error() string { return "invalid batch gateway response" }

// UnsafeBatchGatewayResponseError corresponds to UnsafeBatchGatewayResponse(bytes).
type UnsafeBatchGatewayResponseError struct{ Data []byte }

func (e *UnsafeBatchGatewayResponseError) Error() string { return "unsafe batch gateway response" }

// OffsetOutOfBoundsError corresponds to OffsetOutOfBoundsError(uint256,uint256).
type OffsetOutOfBoundsError struct{ Offset, Length *big.Int }

func (e *OffsetOutOfBoundsError) Error() string {
	return fmt.Sprintf("offset %s out of bounds (length %s)", e.Offset, e.Length)
}

// NormalizationChangedNameError corresponds to
// NormalizationChangedName(bytes,bytes,address), raised by the normalising
// resolve variants when the input was not already normalised. It carries the
// successful result for the normalised name.
type NormalizationChangedNameError struct {
	NormalizedName []byte
	Result         []byte
	Resolver       common.Address
}

func (e *NormalizationChangedNameError) Error() string { return "name was changed by normalization" }

// PrimaryNameNotNormalizedError corresponds to PrimaryNameNotNormalized(string).
type PrimaryNameNotNormalizedError struct{ Name string }

func (e *PrimaryNameNotNormalizedError) Error() string {
	return fmt.Sprintf("primary name %q is not normalized", e.Name)
}

// RegistryRevertError corresponds to ResolutionError(address,bytes): a
// registry reverted while the Universal Resolver, helper or mirror walked
// the v2 tree. The inner revert bytes are preserved verbatim.
type RegistryRevertError struct {
	Registry common.Address
	Data     []byte
}

func (e *RegistryRevertError) Error() string {
	return fmt.Sprintf("registry %s reverted: 0x%s", e.Registry.Hex(), hex.EncodeToString(e.Data))
}

// Inner decodes the nested revert payload; nil when empty.
func (e *RegistryRevertError) Inner() error {
	if len(e.Data) == 0 {
		return nil
	}
	inner, _ := urDecoder.TranslateData(e.Data)
	return inner
}

// Unwrap exposes the decoded inner error to errors.Is and errors.As.
func (e *RegistryRevertError) Unwrap() error { return e.Inner() }

// InvalidOffchainLookupError corresponds to InvalidOffchainLookup(): a
// resolver produced a malformed ERC-3668 OffchainLookup revert.
type InvalidOffchainLookupError struct{}

func (e *InvalidOffchainLookupError) Error() string {
	return "resolver produced an invalid OffchainLookup"
}

// urDecoder decodes the Universal Resolver custom errors.
var urDecoder = revert.NewDecoder().
	Register("ResolverNotFound(bytes)", func(a []any) error {
		return &ResolverNotFoundError{DNSName: bytesArg(a, 0)}
	}).
	Register("ResolverNotContract(bytes,address)", func(a []any) error {
		return &ResolverNotContractError{DNSName: bytesArg(a, 0), Resolver: addressArg(a, 1)}
	}).
	Register("UnsupportedResolverProfile(bytes4)", func(a []any) error {
		out := &UnsupportedResolverProfileError{}
		if s, ok := a[0].([4]byte); ok {
			out.Selector = s
		}
		return out
	}).
	Register("ResolverError(bytes)", func(a []any) error {
		return &ResolverRevertError{Data: bytesArg(a, 0)}
	}).
	Register("EmptyAddress()", func([]any) error { return &EmptyAddressError{} }).
	Register("HttpError(uint16,string)", func(a []any) error {
		out := &HTTPGatewayError{Message: stringArg(a, 1)}
		if s, ok := a[0].(uint16); ok {
			out.Status = s
		}
		return out
	}).
	Register("ReverseAddressMismatch(string,bytes)", func(a []any) error {
		return &ReverseAddressMismatchError{Name: stringArg(a, 0), Address: bytesArg(a, 1)}
	}).
	Register("DNSDecodingFailed(bytes)", func(a []any) error {
		return &DNSDecodingFailedError{Data: bytesArg(a, 0)}
	}).
	Register("DNSEncodingFailed(string)", func(a []any) error {
		return &DNSEncodingFailedError{Name: stringArg(a, 0)}
	}).
	Register("LabelIsEmpty()", func([]any) error { return &LabelIsEmptyError{} }).
	Register("LabelIsTooLong(string)", func(a []any) error {
		return &LabelTooLongError{Label: stringArg(a, 0)}
	}).
	Register("InvalidBatchGatewayResponse()", func([]any) error { return &InvalidBatchGatewayResponseError{} }).
	Register("UnsafeBatchGatewayResponse(bytes)", func(a []any) error {
		return &UnsafeBatchGatewayResponseError{Data: bytesArg(a, 0)}
	}).
	Register("OffsetOutOfBoundsError(uint256,uint256)", func(a []any) error {
		return &OffsetOutOfBoundsError{Offset: bigArg(a, 0), Length: bigArg(a, 1)}
	}).
	Register("NormalizationChangedName(bytes,bytes,address)", func(a []any) error {
		return &NormalizationChangedNameError{NormalizedName: bytesArg(a, 0), Result: bytesArg(a, 1), Resolver: addressArg(a, 2)}
	}).
	Register("PrimaryNameNotNormalized(string)", func(a []any) error {
		return &PrimaryNameNotNormalizedError{Name: stringArg(a, 0)}
	}).
	Register("ResolutionError(address,bytes)", func(a []any) error {
		return &RegistryRevertError{Registry: addressArg(a, 0), Data: bytesArg(a, 1)}
	}).
	Register("InvalidOffchainLookup()", func([]any) error { return &InvalidOffchainLookupError{} })

// URDecoder returns a copy of the Universal Resolver error decoder so other
// packages can merge it into their own.
func URDecoder() *revert.Decoder { return revert.NewDecoder().Merge(urDecoder) }

// translateURRevert decodes a CallContract error as a Universal Resolver
// custom error, returning err unchanged when it is not one.
func translateURRevert(err error) error { return urDecoder.Translate(err) }

func bytesArg(a []any, i int) []byte {
	if i < len(a) {
		if b, ok := a[i].([]byte); ok {
			return b
		}
	}
	return nil
}

func stringArg(a []any, i int) string {
	if i < len(a) {
		if s, ok := a[i].(string); ok {
			return s
		}
	}
	return ""
}

func addressArg(a []any, i int) common.Address {
	if i < len(a) {
		if x, ok := a[i].(common.Address); ok {
			return x
		}
	}
	return common.Address{}
}

func bigArg(a []any, i int) *big.Int {
	if i < len(a) {
		if x, ok := a[i].(*big.Int); ok {
			return x
		}
	}
	return new(big.Int)
}
