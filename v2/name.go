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
	"fmt"
	"math/big"
	"strings"
	"time"

	ens "github.com/ensdomains/go-ens/v5"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Name is a thin handle on a v2 name: its label in the registry that holds
// it. Every method is a one-liner over Registry or ens.Client, so there is
// no hidden routing.
type Name struct {
	// FQDN is the full, normalised name.
	FQDN string
	// Label is the first label; Registry holds it.
	Label    string
	Registry *Registry
	client   *Client
}

// Name binds a name by locating the registry that holds its first label
// through the UniversalHelper. It returns ErrNotV2Name when the parent is
// not in the v2 tree.
func (c *Client) Name(ctx context.Context, fqdn string) (*Name, error) {
	n, err := ens.Normalize(fqdn)
	if err != nil {
		return nil, err
	}
	parent, err := c.Helper().ParentRegistry(ctx, n)
	if err != nil {
		return nil, err
	}
	if parent == (common.Address{}) {
		return nil, fmt.Errorf("%q: %w", fqdn, ErrNotV2Name)
	}
	return NewName(c, parent, n)
}

// NewName binds a name to a known parent registry without a network call.
func NewName(c *Client, parentRegistry common.Address, fqdn string) (*Name, error) {
	n, err := ens.Normalize(fqdn)
	if err != nil {
		return nil, err
	}
	label := n
	if i := strings.IndexByte(n, '.'); i > 0 {
		label = n[:i]
	}
	return &Name{FQDN: n, Label: label, Registry: c.Registry(parentRegistry), client: c}, nil
}

// ETHName binds a second-level .eth name to the .eth registry.
func (c *Client) ETHName(label string) (*Name, error) {
	return NewName(c, c.cfg.V2.ETHRegistry, strings.TrimSuffix(label, ".eth")+".eth")
}

// State returns the registry's record of the name.
func (n *Name) State(ctx context.Context) (*State, error) { return n.Registry.State(ctx, n.Label) }

// Owner returns the name's owner, or zero once expired.
func (n *Name) Owner(ctx context.Context) (common.Address, error) {
	return n.Registry.Owner(ctx, n.Label)
}

// Expiry returns the name's expiry.
func (n *Name) Expiry(ctx context.Context) (time.Time, error) { return n.Registry.Expiry(ctx, n.Label) }

// TokenID returns the current token id.
func (n *Name) TokenID(ctx context.Context) (*big.Int, error) {
	return n.Registry.TokenID(ctx, n.Label)
}

// ResolverAddress returns the resolver set in the parent registry.
func (n *Name) ResolverAddress(ctx context.Context) (common.Address, error) {
	return n.Registry.Resolver(ctx, n.Label)
}

// Resolver finds the writable resolver that serves the name.
func (n *Name) Resolver(ctx context.Context) (*ens.ResolverInfo, error) {
	return n.client.ens.WritableResolver(ctx, n.FQDN)
}

// PermissionedResolver returns a write handle on the name's resolver,
// which must be a PermissionedResolver.
func (n *Name) PermissionedResolver(ctx context.Context) (*PermissionedResolver, error) {
	info, err := n.Resolver(ctx)
	if err != nil {
		return nil, err
	}
	if info.Family != ens.ResolverFamilyPermissioned {
		return nil, fmt.Errorf("resolver %s is %s: %w", info.Underlying.Hex(), info.Family, ErrResolverFamily)
	}
	return n.client.PermissionedResolver(info.Underlying), nil
}

// Subregistry returns the registry of the name's children, or ErrNoSubregistry.
func (n *Name) Subregistry(ctx context.Context) (*Registry, error) {
	addr, err := n.Registry.Subregistry(ctx, n.Label)
	if err != nil {
		return nil, err
	}
	if addr == (common.Address{}) {
		return nil, fmt.Errorf("%q: %w", n.FQDN, ErrNoSubregistry)
	}
	return n.client.Registry(addr), nil
}

// Subname binds a child of this name.
func (n *Name) Subname(ctx context.Context, label string) (*Name, error) {
	sub, err := n.Subregistry(ctx)
	if err != nil {
		return nil, err
	}
	return NewName(n.client, sub.Address(), label+"."+n.FQDN)
}

// Roles returns the roles an account holds on the name.
func (n *Name) Roles(ctx context.Context, account common.Address) (Roles, error) {
	return n.Registry.Roles(ctx, n.Label, account)
}

// HasRoles reports whether an account holds every given role on the name.
func (n *Name) HasRoles(ctx context.Context, roles Roles, account common.Address) (bool, error) {
	return n.Registry.HasRoles(ctx, n.Label, roles, account)
}

// SetResolver sets the name's resolver.
func (n *Name) SetResolver(opts *bind.TransactOpts, resolver common.Address) (*types.Transaction, error) {
	return n.Registry.SetResolver(opts, n.Label, resolver)
}

// SetSubregistry sets the name's child registry.
func (n *Name) SetSubregistry(opts *bind.TransactOpts, sub common.Address) (*types.Transaction, error) {
	return n.Registry.SetSubregistry(opts, n.Label, sub)
}

// GrantRoles grants roles on the name.
func (n *Name) GrantRoles(opts *bind.TransactOpts, roles Roles, account common.Address) (*types.Transaction, error) {
	return n.Registry.GrantRoles(opts, n.Label, roles, account)
}

// RevokeRoles revokes roles on the name.
func (n *Name) RevokeRoles(opts *bind.TransactOpts, roles Roles, account common.Address) (*types.Transaction, error) {
	return n.Registry.RevokeRoles(opts, n.Label, roles, account)
}

// Transfer moves the name to another account.
func (n *Name) Transfer(ctx context.Context, opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return n.Registry.Transfer(ctx, opts, n.Label, to, nil)
}

// Unregister burns the name.
func (n *Name) Unregister(opts *bind.TransactOpts) (*types.Transaction, error) {
	return n.Registry.Unregister(opts, n.Label)
}

// IsETH2LD reports whether the name is held by the .eth registry.
func (n *Name) IsETH2LD() bool { return n.Registry.Address() == n.client.cfg.V2.ETHRegistry }

// Renew renews a .eth name through the routed renewer.
func (n *Name) Renew(ctx context.Context, opts *bind.TransactOpts, duration time.Duration, token common.Address) (*types.Transaction, error) {
	if !n.IsETH2LD() {
		return nil, fmt.Errorf("%q is not a .eth name; renew through its registry", n.FQDN)
	}
	return n.client.Renew(ctx, opts, RenewData{Label: n.Label, Duration: duration}, token)
}

// MigrationStatus reports a .eth name's migration state.
func (n *Name) MigrationStatus(ctx context.Context) (*ens.MigrationStatus, error) {
	return n.client.ens.MigrationStatus(ctx, n.FQDN)
}

// Address resolves the name's Ethereum address through the Universal Resolver.
func (n *Name) Address(ctx context.Context) (common.Address, error) {
	return n.client.ens.ResolveAddress(ctx, n.FQDN)
}

// Text resolves a text record.
func (n *Name) Text(ctx context.Context, key string) (string, error) {
	return n.client.ens.Text(ctx, n.FQDN, key)
}

// Records fetches several records in one call.
func (n *Name) Records(ctx context.Context, req ens.RecordsRequest) (*ens.Records, error) {
	return n.client.ens.Records(ctx, n.FQDN, req)
}
