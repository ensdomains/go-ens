// Copyright 2026 Weald Technology Trading.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0

package v2

import (
	"context"

	ens "github.com/ensdomains/go-ens/v5"
	"github.com/ensdomains/go-ens/v5/contracts/v2/universalhelper"
	"github.com/ethereum/go-ethereum/common"
)

var helperBinding = universalhelper.NewUniversalHelper()

// UniversalHelper walks the v2 registry tree on-chain: owners, registries
// and canonical names. It sees only the v2 tree; names still held in v1
// appear through their reservation.
type UniversalHelper struct {
	client *Client
	addr   common.Address
}

func newHelper(c *Client, addr common.Address) *UniversalHelper {
	return &UniversalHelper{client: c, addr: addr}
}

// Address returns the helper contract address.
func (h *UniversalHelper) Address() common.Address { return h.addr }

func (h *UniversalHelper) dnsView(ctx context.Context, name string, pack func([]byte) []byte) ([]byte, error) {
	dnsName, err := ens.DNSEncode(name)
	if err != nil {
		return nil, err
	}
	return h.client.view(ctx, h.addr, pack(dnsName))
}

// ExactOwner returns the owner of exactly this name, or zero.
func (h *UniversalHelper) ExactOwner(ctx context.Context, name string) (common.Address, error) {
	out, err := h.dnsView(ctx, name, helperBinding.PackFindExactOwner)
	if err != nil {
		return common.Address{}, err
	}
	return helperBinding.UnpackFindExactOwner(out)
}

// NearestOwner returns the owner of the nearest owned ancestor and the DNS
// offset at which it was found.
func (h *UniversalHelper) NearestOwner(ctx context.Context, name string) (common.Address, uint64, error) {
	out, err := h.dnsView(ctx, name, helperBinding.PackFindNearestOwner)
	if err != nil {
		return common.Address{}, 0, err
	}
	res, err := helperBinding.UnpackFindNearestOwner(out)
	if err != nil {
		return common.Address{}, 0, err
	}
	return res.Owner, res.Offset.Uint64(), nil
}

// ExactRegistry returns the registry that is this name (holds its
// children), or zero when the name has no subregistry.
func (h *UniversalHelper) ExactRegistry(ctx context.Context, name string) (common.Address, error) {
	out, err := h.dnsView(ctx, name, helperBinding.PackFindExactRegistry)
	if err != nil {
		return common.Address{}, err
	}
	return helperBinding.UnpackFindExactRegistry(out)
}

// ParentRegistry returns the registry that holds this name's label, or zero
// when the parent is not in the v2 tree.
func (h *UniversalHelper) ParentRegistry(ctx context.Context, name string) (common.Address, error) {
	out, err := h.dnsView(ctx, name, helperBinding.PackFindParentRegistry)
	if err != nil {
		return common.Address{}, err
	}
	return helperBinding.UnpackFindParentRegistry(out)
}

// NearestRegistry returns the deepest registry on the path to the name and
// the DNS offset it serves.
func (h *UniversalHelper) NearestRegistry(ctx context.Context, name string) (common.Address, uint64, error) {
	out, err := h.dnsView(ctx, name, helperBinding.PackFindNearestRegistry)
	if err != nil {
		return common.Address{}, 0, err
	}
	res, err := helperBinding.UnpackFindNearestRegistry(out)
	if err != nil {
		return common.Address{}, 0, err
	}
	return res.Registry, res.Offset.Uint64(), nil
}

// Registries returns the registry chain of a name from the name itself
// (index 0, zero if it has no registry) up to the root.
func (h *UniversalHelper) Registries(ctx context.Context, name string) ([]common.Address, error) {
	out, err := h.dnsView(ctx, name, helperBinding.PackFindRegistries)
	if err != nil {
		return nil, err
	}
	return helperBinding.UnpackFindRegistries(out)
}

// CanonicalName returns the name a registry canonically represents.
func (h *UniversalHelper) CanonicalName(ctx context.Context, registry common.Address) (string, error) {
	out, err := h.client.view(ctx, h.addr, helperBinding.PackFindCanonicalName(registry))
	if err != nil {
		return "", err
	}
	dnsName, err := helperBinding.UnpackFindCanonicalName(out)
	if err != nil {
		return "", err
	}
	return ens.DNSDecode(dnsName)
}

// CanonicalRegistry returns the registry reached by canonical parents only.
func (h *UniversalHelper) CanonicalRegistry(ctx context.Context, name string) (common.Address, error) {
	out, err := h.dnsView(ctx, name, helperBinding.PackFindCanonicalRegistry)
	if err != nil {
		return common.Address{}, err
	}
	return helperBinding.UnpackFindCanonicalRegistry(out)
}
