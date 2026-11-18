// Copyright 2026 Weald Technology Trading.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0

package ens

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ensdomains/go-ens/v5/contracts/v1/baseregistrar"
	"github.com/ensdomains/go-ens/v5/contracts/v1/namewrapper"
	"github.com/ensdomains/go-ens/v5/contracts/v2/ethregistrar"
	"github.com/ensdomains/go-ens/v5/contracts/v2/permissionedregistry"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// MigrationState is where a second-level .eth name stands in the v1 to v2
// migration.
type MigrationState int

const (
	// MigrationStateUnknown: the v2 registry has no record of the label
	// (expiry 0). It was never seeded; consult the v1 registrar.
	MigrationStateUnknown MigrationState = iota
	// MigrationStateAvailable: registrable on v2 once the registrar is enabled.
	MigrationStateAvailable
	// MigrationStateV1Reserved: an unmigrated v1 name. Its owner is in the v1
	// BaseRegistrar or NameWrapper and it renews through ETHRenewerV1.
	MigrationStateV1Reserved
	// MigrationStateV1ReservedGrace: the reservation lapsed less than the
	// grace period ago; still claimable only by its v1 owner.
	MigrationStateV1ReservedGrace
	// MigrationStateMigratedFromV1: registered on v2 and carrying the
	// WAS_RESERVED role, that is migrated from v1.
	MigrationStateMigratedFromV1
	// MigrationStateNativeV2: registered directly on v2.
	MigrationStateNativeV2
	// MigrationStateV2Expired: a v2 registration that lapsed; its former
	// owner may renew within the grace period.
	MigrationStateV2Expired
)

func (s MigrationState) String() string {
	switch s {
	case MigrationStateAvailable:
		return "available"
	case MigrationStateV1Reserved:
		return "v1-reserved"
	case MigrationStateV1ReservedGrace:
		return "v1-reserved-grace"
	case MigrationStateMigratedFromV1:
		return "migrated-from-v1"
	case MigrationStateNativeV2:
		return "native-v2"
	case MigrationStateV2Expired:
		return "v2-expired"
	case MigrationStateUnknown:
		return "unknown"
	default:
		return "unknown"
	}
}

// MigrationStatus describes a second-level .eth name's migration state.
type MigrationStatus struct {
	Label string
	State MigrationState
	// Expiry is the v2 registry expiry (for reserved names the v1 expiry
	// plus the migration bonus); zero when unknown.
	Expiry time.Time
	// GraceEnds is Expiry plus the registrar grace period, when relevant.
	GraceEnds time.Time
	// Owner is the current owner: the v2 token owner for registered names,
	// the v1 owner for reserved ones.
	Owner common.Address
	// LatestOwner is the v2 registry's latest owner, even if expired.
	LatestOwner common.Address
	// TokenID and Resource are the current v2 identifiers; nil unless the
	// name is registered on v2.
	TokenID  *big.Int
	Resource *big.Int

	// V1Owner is BaseRegistrar.ownerOf for reserved names; it equals the
	// NameWrapper when the name is wrapped.
	V1Owner common.Address
	Wrapped bool
	Fuses   uint32
	// Locked is true when the wrapped name has CANNOT_UNWRAP burned, which
	// routes its migration through the LockedMigrationController.
	Locked bool
}

// Registry status values of IPermissionedRegistry.
const (
	registryStatusAvailable  = 0
	registryStatusReserved   = 1
	registryStatusRegistered = 2
)

// roleWasReserved is RegistryRolesLib.ROLE_WAS_RESERVED. The v2 package
// exports the same value; an external test keeps them in sync.
var roleWasReserved = new(big.Int).Lsh(big.NewInt(1), 32)

const fuseCannotUnwrap = 1

// MigrationStatus reports where a second-level .eth name stands in the
// migration. name may be "label" or "label.eth"; anything else returns
// ErrNot2LD.
func (c *Client) MigrationStatus(ctx context.Context, name string) (*MigrationStatus, error) {
	label, err := ethLabel(name)
	if err != nil {
		return nil, err
	}
	if c.cfg.V2.ETHRegistry == (common.Address{}) {
		return nil, fmt.Errorf("eth registry: %w", ErrNotConfigured)
	}
	labelHash, err := LabelHash(label)
	if err != nil {
		return nil, err
	}
	id := new(big.Int).SetBytes(labelHash[:])
	reg := permissionedregistry.NewPermissionedRegistry()
	out, err := c.ethCall(ctx, c.cfg.V2.ETHRegistry, reg.PackGetState(id))
	if err != nil {
		return nil, err
	}
	state, err := reg.UnpackGetState(out)
	if err != nil {
		return nil, fmt.Errorf("unpack getState: %w", err)
	}
	st := &MigrationStatus{Label: label, LatestOwner: state.LatestOwner}
	if state.Expiry != 0 {
		st.Expiry = time.Unix(int64(state.Expiry), 0).UTC()
	}
	grace, err := c.gracePeriod(ctx)
	if err != nil {
		return nil, err
	}
	if !st.Expiry.IsZero() {
		st.GraceEnds = st.Expiry.Add(grace)
	}
	now := time.Now()

	switch state.Status {
	case registryStatusRegistered:
		st.TokenID = state.TokenId
		st.Resource = state.Resource
		ownerOut, err := c.ethCall(ctx, c.cfg.V2.ETHRegistry, reg.PackGetOwner(id))
		if err != nil {
			return nil, err
		}
		owner, err := reg.UnpackGetOwner(ownerOut)
		if err != nil {
			return nil, err
		}
		st.Owner = owner
		rolesOut, err := c.ethCall(ctx, c.cfg.V2.ETHRegistry, reg.PackHasRoles(state.Resource, roleWasReserved, owner))
		if err != nil {
			return nil, err
		}
		was, err := reg.UnpackHasRoles(rolesOut)
		if err != nil {
			return nil, err
		}
		if was {
			st.State = MigrationStateMigratedFromV1
		} else {
			st.State = MigrationStateNativeV2
		}
		return st, nil
	case registryStatusReserved:
		st.State = MigrationStateV1Reserved
		if err := c.fillV1Owner(ctx, st, label); err != nil {
			return nil, err
		}
		return st, nil
	default: // available
		if state.Expiry == 0 {
			st.State = MigrationStateUnknown
			return st, nil
		}
		if state.LatestOwner != (common.Address{}) {
			st.State = MigrationStateV2Expired
			if now.After(st.GraceEnds) {
				st.State = MigrationStateAvailable
			}
			return st, nil
		}
		if now.Before(st.GraceEnds) {
			st.State = MigrationStateV1ReservedGrace
			if err := c.fillV1Owner(ctx, st, label); err != nil {
				return nil, err
			}
			return st, nil
		}
		st.State = MigrationStateAvailable
		return st, nil
	}
}

// fillV1Owner looks the name up in the v1 BaseRegistrar and NameWrapper.
func (c *Client) fillV1Owner(ctx context.Context, st *MigrationStatus, label string) error {
	if c.cfg.V1.BaseRegistrar == (common.Address{}) {
		return nil
	}
	labelHash, _ := LabelHash(label)
	br, err := baseregistrar.NewContract(c.cfg.V1.BaseRegistrar, c.backend)
	if err != nil {
		return err
	}
	opts := &bind.CallOpts{Context: ctx}
	owner, err := br.OwnerOf(opts, new(big.Int).SetBytes(labelHash[:]))
	if err != nil {
		st.V1Owner = common.Address{}
		return nil //nolint:nilerr // ownerOf reverts for unknown tokens; treat as unowned
	}
	st.V1Owner = owner
	st.Owner = owner
	if owner != c.cfg.V1.NameWrapper || c.cfg.V1.NameWrapper == (common.Address{}) {
		return nil
	}
	st.Wrapped = true
	node, err := NameHash(label + ".eth")
	if err != nil {
		return err
	}
	nw, err := namewrapper.NewContract(c.cfg.V1.NameWrapper, c.backend)
	if err != nil {
		return err
	}
	data, err := nw.GetData(opts, new(big.Int).SetBytes(node[:]))
	if err != nil {
		return err
	}
	st.Owner = data.Owner
	st.Fuses = data.Fuses
	st.Locked = data.Fuses&fuseCannotUnwrap != 0
	return nil
}

// gracePeriod reads ETHRegistrar.GRACE_PERIOD once and caches it, falling
// back to 28 days when the registrar is not configured.
func (c *Client) gracePeriod(ctx context.Context) (time.Duration, error) {
	c.graceOnce.Do(func() {
		c.grace = 28 * 24 * time.Hour
		if c.cfg.V2.ETHRegistrar == (common.Address{}) {
			return
		}
		r := ethregistrar.NewETHRegistrar()
		out, err := c.ethCall(ctx, c.cfg.V2.ETHRegistrar, r.PackGRACEPERIOD())
		if err != nil {
			c.graceErr = err
			return
		}
		secs, err := r.UnpackGRACEPERIOD(out)
		if err != nil {
			c.graceErr = err
			return
		}
		c.grace = time.Duration(secs) * time.Second
	})
	return c.grace, c.graceErr
}

func (c *Client) ethCall(ctx context.Context, to common.Address, data []byte) ([]byte, error) {
	out, err := c.backend.CallContract(ctx, ethereum.CallMsg{To: &to, Data: data}, nil)
	if err != nil {
		return nil, urDecoder.Translate(err)
	}
	return out, nil
}

// ethLabel extracts the label of a second-level .eth name ("label" or
// "label.eth"), normalised.
func ethLabel(name string) (string, error) {
	n, err := Normalize(strings.TrimSuffix(name, "."))
	if err != nil {
		return "", err
	}
	parts := strings.Split(n, ".")
	switch {
	case len(parts) == 1 && parts[0] != "":
		return parts[0], nil
	case len(parts) == 2 && parts[1] == "eth" && parts[0] != "":
		return parts[0], nil
	default:
		return "", fmt.Errorf("%q: %w", name, ErrNot2LD)
	}
}
