// Copyright 2026 Weald Technology Trading.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0

package v2_test

import (
	"encoding/hex"
	"errors"
	"math/big"
	"strings"
	"testing"
	"time"

	ens "github.com/ensdomains/go-ens/v5"
	"github.com/ensdomains/go-ens/v5/contracts/v2/ethregistrar"
	"github.com/ensdomains/go-ens/v5/contracts/v2/permissionedregistry"
	v2 "github.com/ensdomains/go-ens/v5/v2"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

type rpcErr struct{ data string }

func (e *rpcErr) Error() string  { return "execution reverted" }
func (e *rpcErr) ErrorData() any { return e.data }

func bigFromString(s string) *big.Int {
	b, ok := new(big.Int).SetString(s, 10)
	if !ok {
		panic(s)
	}
	return b
}

// revertOf packs selector(sig) + abi.encode(args) using the ABI types in sig.
func revertOf(t *testing.T, sig string, args ...any) error {
	t.Helper()
	sel := crypto.Keccak256([]byte(sig))[:4]
	open := strings.IndexByte(sig, '(')
	list := sig[open+1 : len(sig)-1]
	var fields abi.Arguments
	if list != "" {
		for _, typ := range strings.Split(list, ",") {
			ty, err := abi.NewType(typ, "", nil)
			require.NoError(t, err)
			fields = append(fields, abi.Argument{Type: ty})
		}
	}
	body, err := fields.Pack(args...)
	require.NoError(t, err)
	return &rpcErr{data: "0x" + hex.EncodeToString(append(sel, body...))}
}

func TestDecodeRevert_RegistrarErrors(t *testing.T) {
	var c [32]byte
	c[0] = 1
	err := v2.DecodeRevert(revertOf(t, "CommitmentTooNew(bytes32,uint64,uint64)", c, uint64(1_700_000_060), uint64(1_700_000_000)))
	var tooNew *v2.CommitmentTooNewError
	require.ErrorAs(t, err, &tooNew)
	require.Equal(t, c, tooNew.Commitment)
	require.Equal(t, time.Unix(1_700_000_060, 0).UTC(), tooNew.ValidFrom)
	require.True(t, errors.Is(err, ens.ErrCommitmentTooNew))

	err = v2.DecodeRevert(revertOf(t, "NameNotAvailable(string)", "vitalik"))
	require.True(t, errors.Is(err, ens.ErrNameNotAvailable))
	var na *v2.NameNotAvailableError
	require.ErrorAs(t, err, &na)
	require.Equal(t, "vitalik", na.Label)

	err = v2.DecodeRevert(revertOf(t, "DurationTooShort(uint64,uint64)", uint64(60), uint64(2419200)))
	var short *v2.DurationTooShortError
	require.ErrorAs(t, err, &short)
	require.Equal(t, 28*24*time.Hour, short.Min)
	require.True(t, errors.Is(err, ens.ErrDurationTooShort))

	err = v2.DecodeRevert(revertOf(t, "PaymentTokenNotSupported(address)", common.HexToAddress("0x01")))
	require.True(t, errors.Is(err, v2.ErrPaymentTokenNotSupported))

	err = v2.DecodeRevert(revertOf(t, "UnexpiredCommitmentExists(bytes32)", c))
	require.True(t, errors.Is(err, ens.ErrUnexpiredCommitment))
}

// TestDecodeRevert_UsesGeneratedSelectors builds payloads with the abigen
// error ids so the hand-registered signatures cannot drift from the ABIs.
func TestDecodeRevert_UsesGeneratedSelectors(t *testing.T) {
	regABI := ethregistrar.NewETHRegistrar().GetABI()
	e := regABI.Errors["NameNotRenewable"]
	body, err := e.Inputs.Pack("foo")
	require.NoError(t, err)
	payload := append(append([]byte{}, e.ID.Bytes()[:4]...), body...)
	got := v2.DecodeRevert(&rpcErr{data: "0x" + hex.EncodeToString(payload)})
	require.True(t, errors.Is(got, ens.ErrNameNotRenewable), "%v", got)

	pr := permissionedregistry.NewPermissionedRegistry().GetABI()
	for name := range pr.Errors {
		e := pr.Errors[name]
		args := make([]any, len(e.Inputs))
		for i, in := range e.Inputs {
			switch in.Type.T {
			case abi.UintTy:
				switch {
				case in.Type.Size <= 8:
					args[i] = uint8(1)
				case in.Type.Size <= 16:
					args[i] = uint16(1)
				case in.Type.Size <= 32:
					args[i] = uint32(1)
				case in.Type.Size <= 64:
					args[i] = uint64(1)
				default:
					args[i] = big.NewInt(1)
				}
			case abi.AddressTy:
				args[i] = common.HexToAddress("0x02")
			case abi.StringTy:
				args[i] = "x"
			case abi.FixedBytesTy:
				args[i] = [32]byte{}
			default:
				t.Fatalf("unhandled input type %s in %s", in.Type, name)
			}
		}
		body, err := e.Inputs.Pack(args...)
		require.NoError(t, err)
		payload := append(append([]byte{}, e.ID.Bytes()[:4]...), body...)
		got := v2.DecodeRevert(&rpcErr{data: "0x" + hex.EncodeToString(payload)})
		var unknown *ens.UnknownRevertError
		require.False(t, errors.As(got, &unknown), "registry error %s is not decoded", name)
	}
}

func TestDecodeRevert_EACAndTransfer(t *testing.T) {
	roles := v2.NewRoles(v2.RoleSetResolver).Big()
	err := v2.DecodeRevert(revertOf(t, "EACUnauthorizedAccountRoles(uint256,uint256,address)", big.NewInt(0), roles, common.HexToAddress("0x03")))
	var unauth *v2.UnauthorizedRolesError
	require.ErrorAs(t, err, &unauth)
	require.True(t, unauth.Roles.Has(v2.RoleSetResolver))
	require.True(t, errors.Is(err, v2.ErrUnauthorized))
	require.Contains(t, err.Error(), "SET_RESOLVER")

	err = v2.DecodeRevert(revertOf(t, "TransferUnsafeWithMultipleAssignees(uint256,address)", big.NewInt(9), common.HexToAddress("0x04")))
	var unsafe *v2.TransferUnsafeError
	require.ErrorAs(t, err, &unsafe)
	require.Equal(t, v2.TransferUnsafeMultipleAssignees, unsafe.Reason)
	require.True(t, errors.Is(err, v2.ErrTransferUnsafe))

	err = v2.DecodeRevert(revertOf(t, "TransferDisallowed(uint256,address)", big.NewInt(9), common.HexToAddress("0x04")))
	require.True(t, errors.Is(err, v2.ErrTransferDisallowed))
}

func TestDecodeRevert_TokensAndStrings(t *testing.T) {
	err := v2.DecodeRevert(revertOf(t, "Error(string)", "ERC20: transfer amount exceeds allowance"))
	require.True(t, errors.Is(err, v2.ErrInsufficientAllowance), "%v", err)
	err = v2.DecodeRevert(revertOf(t, "Error(string)", "Dai/insufficient-balance"))
	require.True(t, errors.Is(err, v2.ErrInsufficientBalance))
	err = v2.DecodeRevert(revertOf(t, "Error(string)", "something else"))
	var s *ens.StringRevertError
	require.ErrorAs(t, err, &s)

	err = v2.DecodeRevert(revertOf(t, "ERC20InsufficientAllowance(address,uint256,uint256)", common.HexToAddress("0x05"), big.NewInt(1), big.NewInt(2)))
	var ia *v2.InsufficientAllowanceError
	require.ErrorAs(t, err, &ia)
	require.Equal(t, "2", ia.Required.String())
	require.True(t, errors.Is(err, v2.ErrInsufficientAllowance))

	// Universal Resolver errors are merged into the v2 decoder.
	err = v2.DecodeRevert(revertOf(t, "ResolverNotFound(bytes)", []byte{0}))
	require.True(t, errors.Is(err, ens.ErrUnregistered))

	// Unknown payloads pass through unchanged.
	plain := errors.New("dial error")
	require.Same(t, plain, v2.DecodeRevert(plain))
	require.NoError(t, v2.DecodeRevert(nil))
}
