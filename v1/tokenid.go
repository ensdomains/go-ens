// Copyright 2019-2023 Weald Technology Trading.
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
	"errors"
	"fmt"
	"math/big"

	ens "github.com/ensdomains/go-ens/v5"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

// DeriveTokenID derive tokenID from the ENS domain.
//
// The tokenID of the ENS name is simply the uint256 representation of the tokenID of ERC721.
func DeriveTokenID(ctx context.Context, backend bind.ContractBackend, domain string) (string, error) {
	if domain == "" {
		return "", errors.New("empty domain")
	}
	_, err := Resolve(ctx, backend, domain)
	if err != nil {
		return "", err
	}
	domain, err = ens.NormaliseDomain(domain)
	if err != nil {
		return "", err
	}

	domain, err = ens.DomainPart(domain, 1)
	if err != nil {
		return "", err
	}
	labelHash, err := ens.LabelHash(domain)
	if err != nil {
		return "", err
	}
	hash := fmt.Sprintf("%#x", labelHash)
	tokenID, ok := new(big.Int).SetString(hash, 0)
	if !ok {
		return "", errors.New("failed to parse token ID")
	}
	return tokenID.String(), nil
}
