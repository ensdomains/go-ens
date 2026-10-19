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
	"math/big"
	"strings"
	"testing"

	v2 "github.com/ensdomains/go-ens/v5/v2"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

// TestRoleConstants pins every role to the numeric value in
// RegistryRolesLib.sol and PermissionedResolverLib.sol so an upstream
// renumbering fails loudly.
func TestRoleConstants(t *testing.T) {
	one := big.NewInt(1)
	cases := []struct {
		role v2.Role
		bit  *big.Int
	}{
		{v2.RoleRegistrar, new(big.Int).Lsh(one, 0)},
		{v2.RoleRegisterReserved, new(big.Int).Lsh(one, 4)},
		{v2.RoleSetParent, new(big.Int).Lsh(one, 8)},
		{v2.RoleUnregister, new(big.Int).Lsh(one, 12)},
		{v2.RoleRenew, new(big.Int).Lsh(one, 16)},
		{v2.RoleSetSubregistry, new(big.Int).Lsh(one, 20)},
		{v2.RoleSetResolver, new(big.Int).Lsh(one, 24)},
		{v2.RoleCanTransfer.Admin(), new(big.Int).Lsh(one, 28+128)},
		{v2.RoleWasReserved, new(big.Int).Lsh(one, 32)},
		{v2.RoleSetURI, new(big.Int).Lsh(one, 36)},
		{v2.RoleCanName, new(big.Int).Lsh(one, 120)},
		{v2.RoleUpgrade, new(big.Int).Lsh(one, 124)},
		{v2.ResolverRoleSetAddress, new(big.Int).Lsh(one, 0)},
		{v2.ResolverRoleSetText, new(big.Int).Lsh(one, 4)},
		{v2.ResolverRoleSetContenthash, new(big.Int).Lsh(one, 8)},
		{v2.ResolverRoleSetABI, new(big.Int).Lsh(one, 12)},
		{v2.ResolverRoleSetInterface, new(big.Int).Lsh(one, 16)},
		{v2.ResolverRoleSetName, new(big.Int).Lsh(one, 20)},
		{v2.ResolverRoleSetData, new(big.Int).Lsh(one, 24)},
		{v2.ResolverRoleLink, new(big.Int).Lsh(one, 28)},
		{v2.ResolverRoleCanName, new(big.Int).Lsh(one, 120)},
		{v2.ResolverRoleUpgrade, new(big.Int).Lsh(one, 124)},
	}
	for _, tc := range cases {
		require.Equal(t, 0, tc.bit.Cmp(tc.role.Bit()), "role %d", tc.role)
		require.Equal(t, 0, tc.bit.Cmp(v2.NewRoles(tc.role).Big()))
	}
	// Admin role = role << 128.
	require.Equal(t, 0, new(big.Int).Lsh(v2.RoleSetResolver.Bit(), 128).Cmp(v2.RoleSetResolver.Admin().Bit()))
	require.True(t, v2.RoleSetResolver.Admin().IsAdmin())
	require.Equal(t, v2.RoleSetResolver, v2.RoleSetResolver.Admin().Base())
}

func TestRolesSetOperations(t *testing.T) {
	all := v2.AllRoles.Big()
	require.Equal(t, strings.Repeat("1", 64), all.Text(16), "ALL_ROLES is 0x1111…1111")

	r := v2.NewRoles(v2.RoleSetResolver, v2.RoleRenew)
	require.True(t, r.Has(v2.RoleSetResolver))
	require.True(t, r.Has(v2.RoleRenew))
	require.False(t, r.Has(v2.RoleUnregister))
	require.True(t, r.HasAll(v2.NewRoles(v2.RoleRenew)))
	require.False(t, r.HasAll(v2.NewRoles(v2.RoleRenew, v2.RoleUnregister)))
	require.Equal(t, []v2.Role{v2.RoleRenew, v2.RoleSetResolver}, r.List())
	require.Equal(t, "RENEW|SET_RESOLVER", r.String())

	withAdmins := r.WithAdmins()
	require.True(t, withAdmins.Has(v2.RoleRenew.Admin()))
	require.True(t, withAdmins.Has(v2.RoleSetResolver.Admin()))
	require.Equal(t, "RENEW|SET_RESOLVER|RENEW_ADMIN|SET_RESOLVER_ADMIN", withAdmins.String())

	require.False(t, r.Without(v2.RoleRenew).Has(v2.RoleRenew))
	require.True(t, v2.Roles{}.IsZero())
	require.Equal(t, "none", v2.Roles{}.String())

	// Round trip through the on-chain representation.
	back, err := v2.RolesFromBig(withAdmins.Big())
	require.NoError(t, err)
	require.True(t, back.HasAll(withAdmins))
	_, err = v2.RolesFromBig(big.NewInt(2))
	require.Error(t, err, "bit 1 is not a role nybble's low bit")

	// What the registrar grants: resolver and subregistry control with
	// admins, and the transfer admin, but no renew role.
	require.True(t, v2.RegistrationRoles.Has(v2.RoleSetResolver.Admin()))
	require.True(t, v2.RegistrationRoles.Has(v2.RoleCanTransfer.Admin()))
	require.False(t, v2.RegistrationRoles.Has(v2.RoleRenew))
	require.Equal(t, "SET_SUBREGISTRY|SET_RESOLVER|SET_SUBREGISTRY_ADMIN|SET_RESOLVER_ADMIN|CAN_TRANSFER_ADMIN", v2.RegistrationRoles.String())
}

func TestDecodeRoleCounts(t *testing.T) {
	// Two holders of SET_RESOLVER (nybble 6) and one of RENEW (nybble 4).
	counts := new(big.Int).Lsh(big.NewInt(2), 24)
	counts.Or(counts, new(big.Int).Lsh(big.NewInt(1), 16))
	rc := v2.DecodeRoleCounts(counts)
	require.Equal(t, uint8(2), rc[v2.RoleSetResolver])
	require.Equal(t, uint8(1), rc[v2.RoleRenew])
	require.Len(t, rc, 2)
}

func TestSetterScopeResource(t *testing.T) {
	// coin type: keccak256(uint256 as 32-byte word)
	coin := v2.SetterScope{Kind: v2.SetterAddress, CoinType: 60}
	var word [32]byte
	word[31] = 60
	require.Equal(t, hex.EncodeToString(crypto.Keccak256(word[:])), coin.Resource().Text(16))
	require.Equal(t, v2.ResolverRoleSetAddress, coin.Role())

	// text key: keccak256(utf8)
	text := v2.SetterScope{Kind: v2.SetterText, Key: "com.twitter"}
	require.Equal(t, new(big.Int).SetBytes(crypto.Keccak256([]byte("com.twitter"))).Text(16), text.Resource().Text(16))
	require.Equal(t, v2.ResolverRoleSetText, text.Role())

	// interface id: keccak256(4 bytes)
	iface := v2.SetterScope{Kind: v2.SetterInterface, InterfaceID: [4]byte{0x3b, 0x3b, 0x57, 0xde}}
	require.Equal(t, new(big.Int).SetBytes(crypto.Keccak256([]byte{0x3b, 0x3b, 0x57, 0xde})).Text(16), iface.Resource().Text(16))

	// The setter payload starts with the setter selector and a 0x00 name.
	setter, err := coin.EncodeSetter()
	require.NoError(t, err)
	require.Equal(t, "b4436dde", hex.EncodeToString(setter[:4]), "setAddress(bytes,uint256,bytes) selector")
}
