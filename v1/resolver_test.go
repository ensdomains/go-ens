// Copyright 2017 Weald Technology Trading
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1

import (
	"context"
	"encoding/hex"
	"testing"

	ens "github.com/ensdomains/go-ens/v5"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestResolveEmpty(t *testing.T) {
	client := mainnetClient(t)
	_, err := Resolve(context.Background(), client, "")
	assert.NotNil(t, err, "Resolved empty name")
}

func TestResolveZero(t *testing.T) {
	client := mainnetClient(t)
	_, err := Resolve(context.Background(), client, "0")
	assert.NotNil(t, err, "Resolved empty name")
}

func TestResolveNotPresent(t *testing.T) {
	client := mainnetClient(t)
	_, err := Resolve(context.Background(), client, "sirnotappearinginthisregistry.eth")
	require.NotNil(t, err, "Resolved name that does not exist")
	assert.Equal(t, "unregistered name", err.Error(), "Unexpected error")
}

// func TestResolveNoResolver(t *testing.T) {
// 	client := mainnetClient(t)
// 	_, err := Resolve(context.Background(), client, "noresolver.eth")
// 	require.NotNil(t, err, "Resolved name without a resolver")
// 	assert.Equal(t, "no resolver", err.Error(), "Unexpected error")
// }

// resolvestozero.eth and test.enstest.eth, the fixtures earlier releases
// used here, expired on v1 (2020 and January 2026) and were never reserved
// on v2, so since the cutover they are simply unregistered. The zero-address
// mapping is covered offline (records_test.go and the simulated chain).
func TestResolveExpiredNameIsUnregistered(t *testing.T) {
	client := mainnetClient(t)
	_, err := ens.Resolve(context.Background(), client, "resolvestozero.eth")
	require.ErrorIs(t, err, ens.ErrUnregistered)
}

func TestResolveSubname(t *testing.T) {
	client := mainnetClient(t)
	expected := "d8da6bf26964af9d7eed9e03e53415d37aa96045"
	actual, err := ens.Resolve(context.Background(), client, "vitalik.eth")
	require.Nil(t, err, "Error resolving name")
	assert.Equal(t, expected, hex.EncodeToString(actual[:]), "Did not receive expected result")
}

func TestResolveResolverEth(t *testing.T) {
	client := mainnetClient(t)
	expected := "231b0ee14048e9dccd1d247744d114a4eb5e8e63"
	actual, err := Resolve(context.Background(), client, "resolver.eth")
	require.Nil(t, err, "Error resolving name")
	assert.Equal(t, expected, hex.EncodeToString(actual[:]), "Did not receive expected result")
}

func TestResolveAddress(t *testing.T) {
	client := mainnetClient(t)
	expected := "b8c2c29ee19d8307cb7255e1cd9cbde883a267d5"
	actual, err := Resolve(context.Background(), client, "0xb8c2C29ee19D8307cb7255e1Cd9CbDE883A267d5")
	require.Nil(t, err, "Error resolving address")
	assert.Equal(t, expected, hex.EncodeToString(actual[:]), "Did not receive expected result")
}

func TestResolveShortAddress(t *testing.T) {
	client := mainnetClient(t)
	expected := "0000000000000000000000000000000000000001"
	actual, err := Resolve(context.Background(), client, "0x1")
	require.Nil(t, err, "Error resolving address")
	assert.Equal(t, expected, hex.EncodeToString(actual[:]), "Did not receive expected result")
}

func TestResolveHexString(t *testing.T) {
	client := mainnetClient(t)
	_, err := Resolve(context.Background(), client, "0xe32c6d1a964749b6de2130e20daed821a45b9e7261118801ff5319d0ffc6b54a")
	assert.NotNil(t, err, "Resolved too-long hex string")
}

func TestReverseResolveTestEnsTest(t *testing.T) {
	client := mainnetClient(t)
	expected := "nick.eth"
	address := common.HexToAddress("b8c2C29ee19D8307cb7255e1Cd9CbDE883A267d5")
	actual, err := ReverseResolve(context.Background(), client, address)
	require.Nil(t, err, "Error resolving address")
	assert.Equal(t, expected, actual, "Did not receive expected result")
}
