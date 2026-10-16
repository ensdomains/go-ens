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
	"math/big"
	"testing"

	"github.com/ensdomains/go-ens/v5/internal/revert"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

// fakeRPCErr is a stand-in for a go-ethereum RPC error with revert data.
// It implements the rpc.DataError interface that ContractBackend.CallContract
// surfaces on a contract revert.
type fakeRPCErr struct {
	msg  string
	data any
}

func (e *fakeRPCErr) Error() string  { return e.msg }
func (e *fakeRPCErr) ErrorData() any { return e.data }

// revertErr builds an RPC error whose payload is selector(sig) + abi.encode(args).
func revertErr(t *testing.T, sig string, args ...any) error {
	t.Helper()
	return &fakeRPCErr{msg: "execution reverted", data: "0x" + hex.EncodeToString(revertData(t, sig, args...))}
}

func revertData(t *testing.T, sig string, args ...any) []byte {
	t.Helper()
	sel := revert.Selector(sig)
	out := append([]byte{}, sel[:]...)
	if len(args) == 0 {
		return out
	}
	types := sig[len(sig[:len(sig)-1]) : len(sig)-1] // unused; kept for clarity
	_ = types
	var fields abi.Arguments
	for _, typ := range argTypes(sig) {
		ty, err := abi.NewType(typ, "", nil)
		require.NoError(t, err)
		fields = append(fields, abi.Argument{Type: ty})
	}
	body, err := fields.Pack(args...)
	require.NoError(t, err)
	return append(out, body...)
}

func argTypes(sig string) []string {
	open := 0
	for i := range sig {
		if sig[i] == '(' {
			open = i
			break
		}
	}
	list := sig[open+1 : len(sig)-1]
	if list == "" {
		return nil
	}
	var out []string
	start := 0
	for i := 0; i <= len(list); i++ {
		if i == len(list) || list[i] == ',' {
			out = append(out, list[start:i])
			start = i + 1
		}
	}
	return out
}

func TestTranslateURRevert_ResolverNotFound(t *testing.T) {
	got := translateURRevert(revertErr(t, "ResolverNotFound(bytes)", []byte{0x07, 'e', 'x', 'a', 'm', 'p', 'l', 'e', 0x03, 'e', 't', 'h', 0x00}))
	var typed *ResolverNotFoundError
	require.True(t, errors.As(got, &typed), "got %T: %v", got, got)
	require.True(t, errors.Is(got, ErrUnregistered))
	require.Equal(t, "unregistered name", got.Error())
	require.NotEmpty(t, typed.DNSName)
}

func TestTranslateURRevert_EmptyAddress(t *testing.T) {
	got := translateURRevert(revertErr(t, "EmptyAddress()"))
	var typed *EmptyAddressError
	require.True(t, errors.As(got, &typed))
	require.True(t, errors.Is(got, ErrNoAddress))
}

func TestTranslateURRevert_UnsupportedProfile(t *testing.T) {
	got := translateURRevert(revertErr(t, "UnsupportedResolverProfile(bytes4)", [4]byte{0x3b, 0x3b, 0x57, 0xde}))
	var typed *UnsupportedResolverProfileError
	require.True(t, errors.As(got, &typed))
	require.Equal(t, [4]byte{0x3b, 0x3b, 0x57, 0xde}, typed.Selector)
	require.True(t, errors.Is(got, ErrUnsupportedProfile))
}

func TestTranslateURRevert_ResolverNotContract(t *testing.T) {
	addr := common.HexToAddress("0x000000000000000000000000000000000000dEaD")
	got := translateURRevert(revertErr(t, "ResolverNotContract(bytes,address)", []byte{0x03, 'e', 't', 'h', 0x00}, addr))
	var typed *ResolverNotContractError
	require.True(t, errors.As(got, &typed))
	require.Equal(t, addr, typed.Resolver)
}

func TestTranslateURRevert_ResolverErrorWrapsInner(t *testing.T) {
	inner := revertData(t, "UnsupportedResolverProfile(bytes4)", [4]byte{0xc8, 0x69, 0x02, 0x33})
	got := translateURRevert(revertErr(t, "ResolverError(bytes)", inner))
	var typed *ResolverRevertError
	require.True(t, errors.As(got, &typed))
	require.Equal(t, inner, typed.Data)
	// The nested error is reachable through Unwrap.
	require.True(t, errors.Is(got, ErrUnsupportedProfile))
	var profile *UnsupportedResolverProfileError
	require.True(t, errors.As(got, &profile))
}

func TestTranslateURRevert_HTTPGateway(t *testing.T) {
	got := translateURRevert(revertErr(t, "HttpError(uint16,string)", uint16(502), "bad gateway"))
	var typed *HTTPGatewayError
	require.True(t, errors.As(got, &typed))
	require.Equal(t, uint16(502), typed.Status)
	require.Equal(t, "bad gateway", typed.Message)
}

func TestTranslateURRevert_ReverseAddressMismatch(t *testing.T) {
	got := translateURRevert(revertErr(t, "ReverseAddressMismatch(string,bytes)", "vitalik.eth", []byte{1, 2, 3}))
	var typed *ReverseAddressMismatchError
	require.True(t, errors.As(got, &typed))
	require.Equal(t, "vitalik.eth", typed.Name)
	require.Equal(t, []byte{1, 2, 3}, typed.Address)
}

func TestTranslateURRevert_V2Errors(t *testing.T) {
	cases := []struct {
		sig  string
		args []any
		want error
		is   error
	}{
		{"DNSDecodingFailed(bytes)", []any{[]byte{0xff}}, &DNSDecodingFailedError{}, nil},
		{"DNSEncodingFailed(string)", []any{"bad..name"}, &DNSEncodingFailedError{}, nil},
		{"LabelIsEmpty()", nil, &LabelIsEmptyError{}, ErrEmptyLabel},
		{"LabelIsTooLong(string)", []any{"x"}, &LabelTooLongError{}, ErrLabelTooLong},
		{"InvalidBatchGatewayResponse()", nil, &InvalidBatchGatewayResponseError{}, nil},
		{"UnsafeBatchGatewayResponse(bytes)", []any{[]byte{1}}, &UnsafeBatchGatewayResponseError{}, nil},
		{"OffsetOutOfBoundsError(uint256,uint256)", []any{big.NewInt(5), big.NewInt(3)}, &OffsetOutOfBoundsError{}, nil},
		{"NormalizationChangedName(bytes,bytes,address)", []any{[]byte{1}, []byte{2}, common.Address{}}, &NormalizationChangedNameError{}, nil},
		{"PrimaryNameNotNormalized(string)", []any{"Vitalik.eth"}, &PrimaryNameNotNormalizedError{}, nil},
	}
	for _, tc := range cases {
		t.Run(tc.sig, func(t *testing.T) {
			got := translateURRevert(revertErr(t, tc.sig, tc.args...))
			require.NotNil(t, got)
			target := tc.want
			require.ErrorAs(t, got, &target)
			if tc.is != nil {
				require.True(t, errors.Is(got, tc.is))
			}
		})
	}
}

func TestTranslateURRevert_StringAndPanic(t *testing.T) {
	got := translateURRevert(revertErr(t, "Error(string)", "ERC20: transfer amount exceeds allowance"))
	var s *StringRevertError
	require.True(t, errors.As(got, &s))
	require.Equal(t, "ERC20: transfer amount exceeds allowance", s.Reason)

	got = translateURRevert(revertErr(t, "Panic(uint256)", big.NewInt(0x11)))
	var p *PanicError
	require.True(t, errors.As(got, &p))
	require.Equal(t, uint64(0x11), p.Code)
}

func TestTranslateURRevert_UnknownPassThrough(t *testing.T) {
	rpcErr := &fakeRPCErr{msg: "execution reverted", data: "0xdeadbeef" + hex.EncodeToString(make([]byte, 32))}
	got := translateURRevert(rpcErr)
	require.Same(t, rpcErr, got, "unknown selectors must pass through unchanged")
}

func TestTranslateURRevert_NonRPCError(t *testing.T) {
	plain := errors.New("dial tcp: connection refused")
	require.Same(t, plain, translateURRevert(plain))
	require.NoError(t, translateURRevert(nil))
}

func TestTranslateURRevert_AcceptsByteSliceErrorData(t *testing.T) {
	data := revertData(t, "EmptyAddress()")
	got := translateURRevert(&fakeRPCErr{msg: "execution reverted", data: data})
	var typed *EmptyAddressError
	require.True(t, errors.As(got, &typed))
}
