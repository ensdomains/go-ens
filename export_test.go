package ens

import "math/big"

// RoleWasReservedForTest exposes the private ROLE_WAS_RESERVED copy to the
// external test package.
func RoleWasReservedForTest() *big.Int { return new(big.Int).Set(roleWasReserved) }
