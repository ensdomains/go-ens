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
	"errors"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

// fakeRPCErr is a stand-in for a go-ethereum RPC error with revert data.
// It implements the rpc.DataError interface that ContractBackend.CallContract
// surfaces on a contract revert.
type fakeRPCErr struct {
	msg  string
	data string
}

func (e *fakeRPCErr) Error() string             { return e.msg }
func (e *fakeRPCErr) ErrorData() interface{}    { return e.data }

// build a revert payload by concatenating selector + abi-encoded body.
func revert(t *testing.T, sel [4]byte, args ...byte) string {
	t.Helper()
	out := append([]byte{}, sel[:]...)
	out = append(out, args...)
	return "0x" + hex.EncodeToString(out)
}

func TestTranslateURRevert_ResolverNotFound(t *testing.T) {
	// ResolverNotFound(bytes name) with name = 0x07example03eth00.
	body := mustPackArgs(t, argsResolverNotFound, []byte{0x07, 'e', 'x', 'a', 'm', 'p', 'l', 'e', 0x03, 'e', 't', 'h', 0x00})
	rpcErr := &fakeRPCErr{msg: "execution reverted", data: revert(t, selResolverNotFound, body...)}

	got := translateURRevert(rpcErr)
	var typed *ResolverNotFoundError
	require.True(t, errors.As(got, &typed), "expected ResolverNotFoundError, got %T: %v", got, got)
	require.True(t, errors.Is(got, ErrUnregistered))
	require.Equal(t, "unregistered name", got.Error())
	require.NotEmpty(t, typed.DNSName)
}

func TestTranslateURRevert_EmptyAddress(t *testing.T) {
	rpcErr := &fakeRPCErr{msg: "execution reverted", data: revert(t, selEmptyAddress)}

	got := translateURRevert(rpcErr)
	var typed *EmptyAddressError
	require.True(t, errors.As(got, &typed))
	require.True(t, errors.Is(got, ErrNoAddress))
}

func TestTranslateURRevert_UnsupportedProfile(t *testing.T) {
	body := mustPackArgs(t, argsUnsupportedResolverProfile, [4]byte{0x3b, 0x3b, 0x57, 0xde})
	rpcErr := &fakeRPCErr{msg: "execution reverted", data: revert(t, selUnsupportedResolverProfile, body...)}

	got := translateURRevert(rpcErr)
	var typed *UnsupportedResolverProfileError
	require.True(t, errors.As(got, &typed))
	require.Equal(t, [4]byte{0x3b, 0x3b, 0x57, 0xde}, typed.Selector)
}

func TestTranslateURRevert_ResolverNotContract(t *testing.T) {
	dnsName := []byte{0x03, 'e', 't', 'h', 0x00}
	addr := common.HexToAddress("0x000000000000000000000000000000000000dEaD")
	body := mustPackArgs(t, argsResolverNotContract, dnsName, addr)
	rpcErr := &fakeRPCErr{msg: "execution reverted", data: revert(t, selResolverNotContract, body...)}

	got := translateURRevert(rpcErr)
	var typed *ResolverNotContractError
	require.True(t, errors.As(got, &typed))
	require.Equal(t, addr, typed.Resolver)
	require.Equal(t, dnsName, typed.DNSName)
}

func TestTranslateURRevert_UnknownPassThrough(t *testing.T) {
	// A revert that is not a known UR custom error must pass through unchanged.
	rpcErr := &fakeRPCErr{msg: "execution reverted", data: "0xdeadbeef"}
	got := translateURRevert(rpcErr)
	require.Same(t, rpcErr, got, "unknown selectors must not be translated")
}

func TestTranslateURRevert_NonRPCError(t *testing.T) {
	plain := errors.New("boom")
	require.Same(t, plain, translateURRevert(plain), "non-RPC errors must pass through")
	require.Nil(t, translateURRevert(nil))
}

// mustPackArgs is a tiny test helper: pack args with the given Arguments
// definition or fail the test.
func mustPackArgs(t *testing.T, args interface{ Pack(...interface{}) ([]byte, error) }, vals ...interface{}) []byte {
	t.Helper()
	out, err := args.Pack(vals...)
	require.NoError(t, err)
	return out
}
