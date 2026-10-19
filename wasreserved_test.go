// Copyright 2026 Weald Technology Trading.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0

package ens_test

import (
	"testing"

	ens "github.com/ensdomains/go-ens/v5"
	v2 "github.com/ensdomains/go-ens/v5/v2"
	"github.com/stretchr/testify/require"
)

// TestRoleWasReservedMatchesV2 keeps the root package's private copy of
// ROLE_WAS_RESERVED (used by MigrationStatus) equal to the v2 constant.
func TestRoleWasReservedMatchesV2(t *testing.T) {
	require.Equal(t, 0, ens.RoleWasReservedForTest().Cmp(v2.RoleWasReserved.Bit()))
}

func TestMigrationStateString(t *testing.T) {
	require.Equal(t, "v1-reserved", ens.MigrationStateV1Reserved.String())
	require.Equal(t, "migrated-from-v1", ens.MigrationStateMigratedFromV1.String())
	require.Equal(t, "unknown", ens.MigrationStateUnknown.String())
}
