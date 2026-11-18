// Copyright 2026 Weald Technology Trading.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0

// Package revert extracts EVM revert payloads from go-ethereum RPC errors and
// decodes them into typed Go errors through a selector-keyed registry that
// the ens, v1 and v2 packages populate at init.
package revert

import (
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto"
)

// Data extracts the raw revert payload from an error returned by
// ContractBackend.CallContract or by gas estimation. go-ethereum surfaces the
// JSON-RPC "data" field through the rpc.DataError interface, as a 0x-prefixed
// hex string for live backends or as raw bytes for wrapped and mocked ones.
func Data(err error) ([]byte, bool) {
	if err == nil {
		return nil, false
	}
	type dataError interface{ ErrorData() any }
	var de dataError
	if !errors.As(err, &de) {
		return nil, false
	}
	return normalize(de.ErrorData())
}

func normalize(v any) ([]byte, bool) {
	switch d := v.(type) {
	case string:
		raw, err := hex.DecodeString(strings.TrimPrefix(d, "0x"))
		if err != nil {
			return nil, false
		}
		return raw, true
	case []byte:
		return d, true
	default:
		return nil, false
	}
}

// Selector returns the first four bytes of keccak256(sig).
func Selector(sig string) [4]byte {
	var s [4]byte
	copy(s[:], crypto.Keccak256([]byte(sig))[:4])
	return s
}

// Builder turns the ABI-decoded arguments of a custom error into a Go error.
type Builder func(args []any) error

type entry struct {
	sig   string
	args  abi.Arguments
	build Builder
}

// Decoder maps custom-error selectors to builders.
type Decoder struct {
	entries map[[4]byte]entry
}

// NewDecoder returns an empty decoder.
func NewDecoder() *Decoder {
	return &Decoder{entries: map[[4]byte]entry{}}
}

// Register adds a custom error by canonical signature, for example
// "CommitmentTooNew(bytes32,uint64,uint64)". It panics on a malformed
// signature, which is a programming error.
func (d *Decoder) Register(sig string, build Builder) *Decoder {
	args, err := parseArgs(sig)
	if err != nil {
		panic(fmt.Sprintf("revert: register %q: %v", sig, err))
	}
	d.entries[Selector(sig)] = entry{sig: sig, args: args, build: build}
	return d
}

// Merge copies every entry of the given decoders into d. Later decoders win
// on selector collisions.
func (d *Decoder) Merge(others ...*Decoder) *Decoder {
	for _, o := range others {
		if o == nil {
			continue
		}
		for sel, e := range o.entries {
			d.entries[sel] = e
		}
	}
	return d
}

// Translate decodes err if it carries a revert payload the decoder
// recognises, otherwise it returns err unchanged. Error(string) and
// Panic(uint256) reverts are decoded into *StringRevertError and
// *PanicError; an empty payload becomes *EmptyRevertError.
func (d *Decoder) Translate(err error) error {
	if err == nil {
		return nil
	}
	data, ok := Data(err)
	if !ok {
		return err
	}
	if decoded, ok := d.TranslateData(data); ok {
		return decoded
	}
	return err
}

// TranslateData decodes a raw revert payload. The boolean is false when the
// payload is non-empty but the selector is unknown; the returned error is
// then an *UnknownRevertError the caller may still use.
func (d *Decoder) TranslateData(data []byte) (error, bool) {
	if len(data) == 0 {
		return &EmptyRevertError{}, true
	}
	if len(data) < 4 {
		return &UnknownRevertError{Data: data}, false
	}
	var sel [4]byte
	copy(sel[:], data[:4])
	body := data[4:]
	switch sel {
	case selError:
		if vals, err := argsString.Unpack(body); err == nil && len(vals) == 1 {
			if s, ok := vals[0].(string); ok {
				return &StringRevertError{Reason: s}, true
			}
		}
	case selPanic:
		if vals, err := argsUint256.Unpack(body); err == nil && len(vals) == 1 {
			if c, ok := vals[0].(*big.Int); ok {
				return &PanicError{Code: c.Uint64()}, true
			}
		}
	}
	if e, ok := d.entries[sel]; ok {
		vals, err := e.args.Unpack(body)
		if err != nil {
			return &UnknownRevertError{Data: data, Selector: sel}, false
		}
		return e.build(vals), true
	}
	return &UnknownRevertError{Data: data, Selector: sel}, false
}

// Known reports whether the decoder has an entry for the selector.
func (d *Decoder) Known(sel [4]byte) bool {
	_, ok := d.entries[sel]
	return ok
}

var (
	selError    = Selector("Error(string)")
	selPanic    = Selector("Panic(uint256)")
	argsString  abi.Arguments
	argsUint256 abi.Arguments
)

func init() {
	stringT, _ := abi.NewType("string", "", nil)
	uintT, _ := abi.NewType("uint256", "", nil)
	argsString = abi.Arguments{{Type: stringT}}
	argsUint256 = abi.Arguments{{Type: uintT}}
}

// parseArgs builds abi.Arguments from the parameter list of a canonical
// signature. Tuples are not supported; no ENS custom error uses them.
func parseArgs(sig string) (abi.Arguments, error) {
	open := strings.IndexByte(sig, '(')
	if open < 0 || !strings.HasSuffix(sig, ")") {
		return nil, errors.New("not a canonical signature")
	}
	list := sig[open+1 : len(sig)-1]
	if list == "" {
		return abi.Arguments{}, nil
	}
	var args abi.Arguments
	for _, t := range strings.Split(list, ",") {
		typ, err := abi.NewType(strings.TrimSpace(t), "", nil)
		if err != nil {
			return nil, err
		}
		args = append(args, abi.Argument{Type: typ})
	}
	return args, nil
}

// StringRevertError is a revert carrying Error(string).
type StringRevertError struct{ Reason string }

func (e *StringRevertError) Error() string { return "reverted: " + e.Reason }

// PanicError is a revert carrying Panic(uint256).
type PanicError struct{ Code uint64 }

func (e *PanicError) Error() string {
	name := panicNames[e.Code]
	if name == "" {
		name = "unknown"
	}
	return fmt.Sprintf("panic 0x%02x (%s)", e.Code, name)
}

var panicNames = map[uint64]string{
	0x00: "generic compiler panic",
	0x01: "assertion failed",
	0x11: "arithmetic overflow or underflow",
	0x12: "division or modulo by zero",
	0x21: "invalid enum conversion",
	0x22: "invalid storage byte array access",
	0x31: "pop on empty array",
	0x32: "array index out of bounds",
	0x41: "out of memory",
	0x51: "uninitialised function pointer call",
}

// EmptyRevertError is a revert with no payload.
type EmptyRevertError struct{}

func (e *EmptyRevertError) Error() string { return "reverted without data" }

// UnknownRevertError is a revert whose selector is not registered.
type UnknownRevertError struct {
	Selector [4]byte
	Data     []byte
}

func (e *UnknownRevertError) Error() string {
	return fmt.Sprintf("reverted with unknown error 0x%s", hex.EncodeToString(e.Data))
}
