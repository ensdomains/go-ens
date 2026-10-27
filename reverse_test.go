// Copyright 2022 Weald Technology Trading
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

package ens_test

import (
	"context"
	"testing"

	ens "github.com/ensdomains/go-ens/v5"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

// TestReverseResolve tests the reverse resolution functionality.
func TestReverseResolve(t *testing.T) {
	tests := []struct {
		name    string
		address common.Address
		res     string
		anyOf   []string
		err     string
	}{
		{
			name:    "NoResolver",
			address: common.Address{},
			// ens.UniversalResolver returns an empty primary name for the zero
			// address; ReverseResolve translates that to "no resolution".
			err: "no resolution",
		},
		{
			name:    "NoReverseRecord",
			address: common.HexToAddress("0x388ea662ef2c223ec0b047d41bf3c0f362142ad5"),
			err:     "no resolution",
		},
		{
			name:    "Exists",
			address: common.HexToAddress("0x3AF5CFc7C7f29Da9d7758339b8eaB6a78550Cbed"),
			// The integration fixture answers with a different sentinel
			// depending on which Universal Resolver implementation asks
			// (legacy v1 vs v2), so either is a pass; what matters is that
			// a forward-verified primary name comes back.
			anyOf: []string{"ur-reverse.integration-tests.eth", "v1-reverse.integration-tests.eth"},
		},
	}

	client := mainnetClient(t)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res, err := ens.ReverseResolve(context.Background(), client, test.address)
			if test.err != "" {
				require.EqualError(t, err, test.err)
			} else {
				require.NoError(t, err)
				if len(test.anyOf) > 0 {
					require.Contains(t, test.anyOf, res)
				} else {
					require.Equal(t, test.res, res)
				}
			}
		})
	}
}
