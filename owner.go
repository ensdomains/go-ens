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

	"github.com/ensdomains/go-ens/v5/contracts/v2/universalhelper"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

func (c *Client) helperCall(ctx context.Context, data []byte) ([]byte, error) {
	if c.cfg.V2.UniversalHelper == (common.Address{}) {
		return nil, fmt.Errorf("universal helper: %w", ErrNotConfigured)
	}
	to := c.cfg.V2.UniversalHelper
	out, err := c.backend.CallContract(ctx, ethereum.CallMsg{To: &to, Data: data}, nil)
	if err != nil {
		return nil, urDecoder.Translate(err)
	}
	return out, nil
}

// Owner returns the owner of the exact name in the v2 registry tree via
// UniversalHelper.findExactOwner. Names still held only in the v1 registry
// are reached through the v1 mirror; the zero address with a nil error means
// nobody owns the exact name.
func (c *Client) Owner(ctx context.Context, name string) (common.Address, error) {
	dnsName, err := DNSEncode(name)
	if err != nil {
		return common.Address{}, err
	}
	h := universalhelper.NewUniversalHelper()
	out, err := c.helperCall(ctx, h.PackFindExactOwner(dnsName))
	if err != nil {
		return common.Address{}, err
	}
	return h.UnpackFindExactOwner(out)
}

// NearestOwner returns the owner of the nearest owned ancestor of the name
// and the byte offset into the DNS-encoded name at which it was found.
func (c *Client) NearestOwner(ctx context.Context, name string) (common.Address, uint64, error) {
	dnsName, err := DNSEncode(name)
	if err != nil {
		return common.Address{}, 0, err
	}
	h := universalhelper.NewUniversalHelper()
	out, err := c.helperCall(ctx, h.PackFindNearestOwner(dnsName))
	if err != nil {
		return common.Address{}, 0, err
	}
	res, err := h.UnpackFindNearestOwner(out)
	if err != nil {
		return common.Address{}, 0, err
	}
	return res.Owner, res.Offset.Uint64(), nil
}
