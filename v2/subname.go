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

// Step names of a SubnamePlan.
const (
	StepDeployRegistry = "deployUserRegistry"
	StepSetSubregistry = "setSubregistry"
	StepSetParent      = "setParent"
	StepRegister       = "register"
	StepSetResolver    = "setResolver"
)

// Step is one transaction of a plan.
type Step struct {
	Name string
	Call ens.Call
	// Skipped marks a step the chain already satisfies.
	Skipped bool
}

// SubnamePlan is the ordered transaction list that creates a subname. Steps
// must be mined in order: later ones call the registry the first deploys.
type SubnamePlan struct {
	// Name is the subname being created.
	Name string
	// Label is its first label; Parent the rest.
	Label  string
	Parent string
	// ParentRegistry is the registry that will hold Label: existing, or the
	// predicted address of the one the plan deploys.
	ParentRegistry common.Address
	// GrandparentRegistry is the registry holding Parent's label.
	GrandparentRegistry common.Address
	ParentLabel         string
	Steps               []Step
}

// Pending returns the steps that are not skipped.
func (p *SubnamePlan) Pending() []Step {
	var out []Step
	for _, s := range p.Steps {
		if !s.Skipped {
			out = append(out, s)
		}
	}
	return out
}

// SubnameParams configures PlanCreateSubname.
type SubnameParams struct {
	// Name is the subname, e.g. "sub.example.eth".
	Name  string
	Owner common.Address
	// Resolver and Subregistry for the new name; zero for none.
	Resolver    common.Address
	Subregistry common.Address
	// Roles granted to Owner; empty means AllRoles.
	Roles Roles
	// Expiry zero means MaxExpiry.
	Expiry time.Time
	// RegistrySalt is used if the parent needs a registry deployed; nil
	// means UserRegistrySalt(parent, 0), which keeps the plan idempotent.
	RegistrySalt *big.Int
	// RegistryGrants are the root-role holders of a deployed registry;
	// empty means [{from, AllRoles}].
	RegistryGrants []Grant
	// Overwrite lets the plan update resolver and subregistry of an existing
	// subname instead of failing with ErrSubnameExists.
	Overwrite bool
}

// PlanCreateSubname inspects the chain and returns the transactions that
// from must send to create the subname. It checks that the parent is
// registered in a v2 registry and that from may set its subregistry.
func (c *Client) PlanCreateSubname(ctx context.Context, from common.Address, p SubnameParams) (*SubnamePlan, error) {
	name, err := ens.Normalize(p.Name)
	if err != nil {
		return nil, err
	}
	i := strings.IndexByte(name, '.')
	if i <= 0 || i == len(name)-1 {
		return nil, fmt.Errorf("%q is not a subname", p.Name)
	}
	label, parent := name[:i], name[i+1:]
	j := strings.IndexByte(parent, '.')
	parentLabel := parent
	if j > 0 {
		parentLabel = parent[:j]
	}
	regs, err := c.Helper().Registries(ctx, parent)
	if err != nil {
		return nil, err
	}
	if len(regs) < 2 || regs[1] == (common.Address{}) {
		return nil, fmt.Errorf("%q: %w", parent, ErrNotV2Name)
	}
	grandparent := c.Registry(regs[1])
	st, err := grandparent.State(ctx, parentLabel)
	if err != nil {
		return nil, err
	}
	if st.Status != StatusRegistered {
		return nil, fmt.Errorf("%q is %s in its registry: %w", parent, st.Status, ErrNotV2Name)
	}
	plan := &SubnamePlan{Name: name, Label: label, Parent: parent, GrandparentRegistry: regs[1], ParentLabel: parentLabel}
	parentReg := regs[0]
	if parentReg == (common.Address{}) {
		ok, err := grandparent.HasRoles(ctx, parentLabel, NewRoles(RoleSetSubregistry), from)
		if err != nil {
			return nil, err
		}
		if !ok {
			return nil, fmt.Errorf("%s cannot set the subregistry of %q: %w", from.Hex(), parent, ErrUnauthorized)
		}
		salt := p.RegistrySalt
		if salt == nil {
			if salt, err = UserRegistrySalt(parent, 0); err != nil {
				return nil, err
			}
		}
		predicted, err := c.Factory().PredictProxyAddress(ctx, from, salt)
		if err != nil {
			return nil, err
		}
		grants := p.RegistryGrants
		if len(grants) == 0 {
			grants = []Grant{{Account: from, Roles: AllRoles}}
		}
		init, err := EncodeUserRegistryInit(grants)
		if err != nil {
			return nil, err
		}
		deployed, err := c.hasCode(ctx, predicted)
		if err != nil {
			return nil, err
		}
		plan.Steps = append(plan.Steps, Step{Name: StepDeployRegistry, Call: c.Factory().BuildDeployProxy(c.cfg.V2.UserRegistryImpl, salt, init), Skipped: deployed})
		attach, err := grandparent.BuildSetSubregistry(parentLabel, predicted)
		if err != nil {
			return nil, err
		}
		plan.Steps = append(plan.Steps, Step{Name: StepSetSubregistry, Call: attach})
		setParent, err := c.Registry(predicted).BuildSetParent(regs[1], parentLabel)
		if err != nil {
			return nil, err
		}
		plan.Steps = append(plan.Steps, Step{Name: StepSetParent, Call: setParent})
		parentReg = predicted
	} else {
		curParent, curLabel, err := c.Registry(parentReg).Parent(ctx)
		if err == nil && (curParent != regs[1] || curLabel != parentLabel) {
			setParent, err := c.Registry(parentReg).BuildSetParent(regs[1], parentLabel)
			if err != nil {
				return nil, err
			}
			plan.Steps = append(plan.Steps, Step{Name: StepSetParent, Call: setParent})
		}
	}
	plan.ParentRegistry = parentReg
	roles := p.Roles
	if roles.IsZero() {
		roles = AllRoles
	}
	reg := c.Registry(parentReg)
	existing := StatusAvailable
	if regs[0] != (common.Address{}) {
		est, err := reg.State(ctx, label)
		if err != nil {
			return nil, err
		}
		existing = est.Status
	}
	if existing == StatusRegistered {
		if !p.Overwrite {
			return nil, fmt.Errorf("%q: %w", name, ErrSubnameExists)
		}
		setRes, err := reg.BuildSetResolver(label, p.Resolver)
		if err != nil {
			return nil, err
		}
		plan.Steps = append(plan.Steps, Step{Name: StepSetResolver, Call: setRes})
		setSub, err := reg.BuildSetSubregistry(label, p.Subregistry)
		if err != nil {
			return nil, err
		}
		plan.Steps = append(plan.Steps, Step{Name: StepSetSubregistry, Call: setSub})
		return plan, nil
	}
	register, err := reg.BuildRegister(RegisterParams{Label: label, Owner: p.Owner, Subregistry: p.Subregistry, Resolver: p.Resolver, Roles: roles, Expiry: p.Expiry})
	if err != nil {
		return nil, err
	}
	plan.Steps = append(plan.Steps, Step{Name: StepRegister, Call: register})
	return plan, nil
}

// ExecutePlan sends the pending steps in order, waiting for each to be
// mined with wait (nil means the client's default). It returns the receipts
// of the steps that were sent; on failure the error names the step.
func (c *Client) ExecutePlan(ctx context.Context, opts *bind.TransactOpts, plan *SubnamePlan, wait WaitFunc) ([]*types.Receipt, error) {
	if wait == nil {
		wait = c.DefaultWait()
	}
	var receipts []*types.Receipt
	for _, step := range plan.Pending() {
		tx, err := c.send(opts, step.Call)
		if err != nil {
			return receipts, fmt.Errorf("%s: %w", step.Name, err)
		}
		receipt, err := wait(ctx, tx)
		if err != nil {
			return receipts, fmt.Errorf("%s: %w", step.Name, err)
		}
		receipts = append(receipts, receipt)
	}
	return receipts, nil
}

// CreateSubname plans and executes the creation of a subname.
func (c *Client) CreateSubname(ctx context.Context, opts *bind.TransactOpts, p SubnameParams, wait WaitFunc) ([]*types.Receipt, error) {
	plan, err := c.PlanCreateSubname(ctx, opts.From, p)
	if err != nil {
		return nil, err
	}
	return c.ExecutePlan(ctx, opts, plan, wait)
}

// DeleteSubname unregisters a label from a registry. Needs ROLE_UNREGISTER.
func (c *Client) DeleteSubname(opts *bind.TransactOpts, registry common.Address, label string) (*types.Transaction, error) {
	return c.Registry(registry).Unregister(opts, label)
}
