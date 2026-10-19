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

	"github.com/ensdomains/go-ens/v5/internal/revert"
	"github.com/ethereum/go-ethereum/common"
)

// UnknownAddress is the zero address.
var UnknownAddress = common.Address{}

// Sentinel errors. Compare with errors.Is; the typed errors in this package
// and in the v1 and v2 packages unwrap to them.
//
//nolint:godoclint // a var block has no single symbol name
var (
	// ErrUnregistered is returned when a name has no resolver anywhere up its
	// parent chain.
	ErrUnregistered = errors.New("unregistered name")
	// ErrNoAddress is returned when a resolver is found but the address
	// record is missing or zero.
	ErrNoAddress = errors.New("no address")
	// ErrUnsupportedProfile is returned when the resolver does not implement
	// the requested record profile (for example pubkey on a v2
	// PermissionedResolver).
	ErrUnsupportedProfile = errors.New("resolver profile not supported")
	// ErrNoPrimaryName is returned by reverse resolution when the address has
	// no primary name. The message is kept from earlier releases.
	ErrNoPrimaryName = errors.New("no resolution")
	// ErrLabelTooLong is returned when a label exceeds 255 bytes, which the
	// DNS wire format used by the Universal Resolver cannot carry.
	ErrLabelTooLong = errors.New("label exceeds 255 bytes")
	// ErrEmptyLabel is returned when a name contains an empty label.
	ErrEmptyLabel = errors.New("empty label")
	// ErrNot2LD is returned when an operation that only applies to
	// second-level .eth names is given another name.
	ErrNot2LD = errors.New("name is not a second-level .eth name")
	// ErrNotENSv2 is returned when the configured Universal Resolver is not
	// UniversalResolverV2.
	ErrNotENSv2 = errors.New("universal resolver is not ENSv2")

	// Registrar sentinels shared by the v1 and v2 packages.
	ErrNameNotAvailable    = errors.New("name not available")
	ErrCommitmentTooNew    = errors.New("commitment too new")
	ErrCommitmentTooOld    = errors.New("commitment too old")
	ErrUnexpiredCommitment = errors.New("unexpired commitment exists")
	ErrDurationTooShort    = errors.New("duration too short")
	ErrNameNotRenewable    = errors.New("name not renewable")

	errAddressTooLong     = errors.New("address too long")
	errAddressUnparseable = errors.New("could not parse address")
)

// Re-exported revert error types, so callers do not need the internal package.
type (
	// StringRevertError is a revert carrying Error(string).
	StringRevertError = revert.StringRevertError
	// PanicError is a revert carrying Panic(uint256).
	PanicError = revert.PanicError
	// EmptyRevertError is a revert with no payload.
	EmptyRevertError = revert.EmptyRevertError
	// UnknownRevertError is a revert whose selector no decoder recognises.
	UnknownRevertError = revert.UnknownRevertError
)

// RevertData extracts the raw revert payload from an RPC error, if any.
func RevertData(err error) ([]byte, bool) { return revert.Data(err) }

// LabelTooLongError reports a label longer than 255 bytes.
type LabelTooLongError struct{ Label string }

func (e *LabelTooLongError) Error() string {
	return "label " + truncateLabel(e.Label) + " exceeds 255 bytes"
}

// Unwrap returns ErrLabelTooLong.
func (e *LabelTooLongError) Unwrap() error { return ErrLabelTooLong }

func truncateLabel(s string) string {
	if len(s) <= 32 {
		return `"` + s + `"`
	}
	return `"` + s[:32] + `..."`
}
