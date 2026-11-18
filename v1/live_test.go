// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0

package v1

import (
	"context"
	"testing"

	ens "github.com/ensdomains/go-ens/v5"
	"github.com/ensdomains/go-ens/v5/internal/livetest"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
)

// TestLiveClientContracts checks that the configured v1 contracts answer the
// ABIs their wrappers speak.
func TestLiveClientContracts(t *testing.T) {
	ctx := context.Background()
	root, err := ens.NewClient(ctx, mainnetClient(t))
	require.NoError(t, err)
	c := NewClient(root)

	// ETHController speaks the legacy controller ABI; the wrapped
	// controller rejects its makeCommitment.
	controller, err := c.Controller()
	require.NoError(t, err)
	require.Equal(t, root.Config().V1.LegacyETHRegistrarController, controller.ContractAddr)
	hash, err := controller.CommitmentHash("go-ens-test.eth", common.HexToAddress("0x1"), [32]byte{1})
	require.NoError(t, err)
	require.NotEqual(t, common.Hash{}, hash)

	// Client.Name binds the configured contracts.
	name, err := c.Name("go-ens-test.eth")
	require.NoError(t, err)
	require.Equal(t, controller.ContractAddr, name.controller.ContractAddr)
	require.Equal(t, root.Config().V1.Registry, name.registry.ContractAddr)

	wrapper, err := c.NameWrapper()
	require.NoError(t, err)
	wrapped, err := wrapper.IsWrapped(ctx, "vitalik.eth")
	require.NoError(t, err)
	require.False(t, wrapped)
}

// mainnetClient dials mainnet; see internal/livetest for the endpoint rules.
func mainnetClient(t *testing.T) *ethclient.Client {
	t.Helper()
	return livetest.MainnetRPC(t)
}
