// Copyright 2026 Weald Technology Trading.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0

package v2_test

import (
	"context"
	"errors"
	"math/big"
	"testing"
	"time"

	ens "github.com/ensdomains/go-ens/v5"
	"github.com/ensdomains/go-ens/v5/internal/simchain"
	v2 "github.com/ensdomains/go-ens/v5/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

// TestSimulatedEndToEnd runs the whole v2 write path on an in-process chain
// built from the vendored mainnet bytecode: register example.eth with USDC
// after deploying a PermissionedResolver with initial records, read the
// records back through UniversalResolverV2, renew, create a subname with a
// user registry, change roles and observe the token id regenerate, and
// transfer the name.
func TestSimulatedEndToEnd(t *testing.T) {
	chain := simchain.New(t)
	ctx := context.Background()
	c := v2.NewClient(chain.ENS)
	alice := chain.Alice
	opts := alice.Opts(t)
	wait := chain.WaitFunc(t)

	ok, err := chain.ENS.IsENSv2(ctx)
	require.NoError(t, err)
	require.True(t, ok)
	open, err := c.RegistrationOpen(ctx)
	require.NoError(t, err)
	require.True(t, open)

	reg := c.ETHRegistrar()
	consts, err := reg.Constants(ctx)
	require.NoError(t, err)
	require.Equal(t, simchain.MinCommitmentAge, consts.MinCommitmentAge)

	// Resolver first, with initial records set inside initialize().
	salt := v2.OwnedResolverSalt(alice.Address, 0)
	resolverAddr, err := c.Factory().PredictProxyAddress(ctx, alice.Address, salt)
	require.NoError(t, err)
	pr := c.PermissionedResolver(resolverAddr)
	setAddr, err := pr.BuildSetAddress("example.eth", ens.CoinTypeETH, alice.Address.Bytes())
	require.NoError(t, err)
	setText, err := pr.BuildSetText("example.eth", "description", "hello from go-ens")
	require.NoError(t, err)
	tx, predicted, err := c.DeployPermissionedResolver(ctx, opts, v2.DeployResolverParams{Salt: salt, InitCalls: []ens.Call{setAddr, setText}})
	require.NoError(t, err)
	require.Equal(t, resolverAddr, predicted)
	ev, err := c.Factory().ParseProxyDeployed(chain.Wait(t, tx))
	require.NoError(t, err)
	require.Equal(t, resolverAddr, ev.Proxy)
	require.Equal(t, alice.Address, ev.Sender)
	impl, err := c.Factory().VerifyContract(ctx, resolverAddr)
	require.NoError(t, err)
	require.Equal(t, chain.Config.V2.PermissionedResolverImpl, impl)
	hasRoot, err := pr.HasRootRoles(ctx, v2.ResolverAllRoles, alice.Address)
	require.NoError(t, err)
	require.True(t, hasRoot)

	// Price and payment.
	available, err := reg.IsAvailable(ctx, "example")
	require.NoError(t, err)
	require.True(t, available)
	price, err := reg.RegisterPrice(ctx, "example", 365*24*time.Hour, common.Address{})
	require.NoError(t, err)
	require.Equal(t, chain.USDC, price.Token)
	require.True(t, price.Base.Sign() > 0)
	require.Equal(t, 0, price.Premium.Sign())
	t.Logf("example.eth for a year: %s USDC base units", price.Total())

	secret, err := v2.NewSecret()
	require.NoError(t, err)
	params := v2.RegistrationParams{Commitment: v2.Commitment{
		Label: "example", Owner: alice.Address, Secret: secret, Resolver: resolverAddr, Duration: 365 * 24 * time.Hour,
	}}

	// Registration before any commitment is refused client-side.
	_, err = reg.Register(ctx, opts, params)
	require.ErrorIs(t, err, v2.ErrNoCommitment)

	tx, h, err := reg.Commit(opts, params.Commitment)
	require.NoError(t, err)
	chain.Wait(t, tx)
	onchain, err := reg.MakeCommitmentOnChain(ctx, params.Commitment)
	require.NoError(t, err)
	require.Equal(t, onchain, h)
	at, err := reg.CommitmentAt(ctx, h)
	require.NoError(t, err)
	require.False(t, at.IsZero())

	// Too new: the client-side pre-flight reports it without sending.
	_, err = reg.Register(ctx, opts, params)
	var tooNew *v2.CommitmentTooNewError
	require.ErrorAs(t, err, &tooNew)
	chain.AdvanceTime(t, simchain.MinCommitmentAge+time.Second)

	// No allowance yet: typed error before gas is spent.
	_, err = reg.Register(ctx, opts, params)
	require.ErrorIs(t, err, v2.ErrInsufficientBalance)
	chain.MintUSDC(t, alice.Address, new(big.Int).Mul(price.Total(), big.NewInt(10)))
	_, err = reg.Register(ctx, opts, params)
	var allowance *v2.InsufficientAllowanceError
	require.ErrorAs(t, err, &allowance)
	require.Equal(t, chain.Config.V2.ETHRegistrar, allowance.Spender)
	tx, err = c.EnsureAllowance(ctx, opts, common.Address{}, chain.Config.V2.ETHRegistrar, price.Total())
	require.NoError(t, err)
	require.NotNil(t, tx)
	chain.Wait(t, tx)
	tx, err = c.EnsureAllowance(ctx, opts, common.Address{}, chain.Config.V2.ETHRegistrar, price.Total())
	require.NoError(t, err)
	require.Nil(t, tx, "allowance already sufficient")

	// Register.
	tx, err = reg.Register(ctx, opts, params)
	require.NoError(t, err)
	receipt := chain.Wait(t, tx)
	res, err := reg.ParseRegistered(receipt)
	require.NoError(t, err)
	require.Equal(t, "example", res.Label)
	require.Equal(t, alice.Address, res.Owner)
	require.Equal(t, resolverAddr, res.Resolver)
	require.Equal(t, 0, res.Base.Cmp(price.Base))

	ethReg := c.ETHRegistry()
	emancipated, err := ethReg.IsEmancipated(ctx)
	require.NoError(t, err)
	require.True(t, emancipated, "harness must emancipate the .eth registry like mainnet")
	st, err := ethReg.State(ctx, "example")
	require.NoError(t, err)
	require.Equal(t, v2.StatusRegistered, st.Status)
	require.Equal(t, alice.Address, st.LatestOwner)
	require.Equal(t, 0, st.TokenID.Cmp(res.TokenID))
	owner, err := ethReg.Owner(ctx, "example")
	require.NoError(t, err)
	require.Equal(t, alice.Address, owner)
	roles, err := ethReg.Roles(ctx, "example", alice.Address)
	require.NoError(t, err)
	require.True(t, roles.HasAll(v2.RegistrationRoles))
	require.False(t, roles.Has(v2.RoleRenew))

	// Reads through UniversalResolverV2 land on the PermissionedResolver.
	addr, err := chain.ENS.ResolveAddress(ctx, "example.eth")
	require.NoError(t, err)
	require.Equal(t, alice.Address, addr)
	text, err := chain.ENS.Text(ctx, "example.eth", "description")
	require.NoError(t, err)
	require.Equal(t, "hello from go-ens", text)
	recs, err := chain.ENS.Records(ctx, "example.eth", ens.RecordsRequest{Texts: []string{"description", "missing"}, Coins: []uint64{60}, Contenthash: true})
	require.NoError(t, err)
	require.Equal(t, "hello from go-ens", recs.Texts["description"])
	require.NotContains(t, recs.Texts, "missing")
	require.Equal(t, alice.Address.Bytes(), recs.Coins[60])
	require.Empty(t, recs.Errors)
	// pubkey is not a PermissionedResolver profile.
	_, _, err = chain.ENS.PubKey(ctx, "example.eth")
	require.ErrorIs(t, err, ens.ErrUnsupportedProfile)
	info, err := chain.ENS.WritableResolver(ctx, "example.eth")
	require.NoError(t, err)
	require.Equal(t, resolverAddr, info.Underlying)
	require.Equal(t, ens.ResolverFamilyPermissioned, info.Family)
	require.True(t, info.Writable)
	v2owner, err := chain.ENS.Owner(ctx, "example.eth")
	require.NoError(t, err)
	require.Equal(t, alice.Address, v2owner)

	// Wildcard: a subname without its own records resolves through the
	// parent's resolver, which has no record for it.
	_, err = chain.ENS.ResolveAddress(ctx, "sub.example.eth")
	require.ErrorIs(t, err, ens.ErrNoAddress)

	// More records through SetRecords, then a multicoin address.
	tx, err = pr.SetRecords(opts, "example.eth", v2.RecordSet{Texts: map[string]string{"url": "https://example.org"}, Coins: map[uint64][]byte{ens.EVMCoinType(10): alice.Address.Bytes()}})
	require.NoError(t, err)
	chain.Wait(t, tx)
	op, err := chain.ENS.MultiAddress(ctx, "example.eth", ens.EVMCoinType(10))
	require.NoError(t, err)
	require.Equal(t, alice.Address.Bytes(), op)
	_, err = pr.BuildSetAddress("example.eth", ens.EVMCoinType(10), []byte{1, 2, 3})
	require.ErrorIs(t, err, v2.ErrInvalidEVMAddress)

	// Migration status of a native v2 registration.
	ms, err := chain.ENS.MigrationStatus(ctx, "example.eth")
	require.NoError(t, err)
	require.Equal(t, ens.MigrationStateNativeV2, ms.State)
	require.Equal(t, alice.Address, ms.Owner)

	// Renewal routes to the registrar.
	renewer, _, err := c.ChooseRenewer(ctx, "example")
	require.NoError(t, err)
	require.Equal(t, chain.Config.V2.ETHRegistrar, renewer.Address())
	renewPrice, err := renewer.RenewPrice(ctx, "example", 365*24*time.Hour, common.Address{})
	require.NoError(t, err)
	tx, err = c.ApprovePayment(opts, common.Address{}, renewer.Address(), renewPrice)
	require.NoError(t, err)
	chain.Wait(t, tx)
	before := st.Expiry
	tx, err = c.Renew(ctx, opts, v2.RenewData{Label: "example", Duration: 365 * 24 * time.Hour}, common.Address{})
	require.NoError(t, err)
	receipt = chain.Wait(t, tx)
	renewals, err := renewer.ParseRenewed(receipt)
	require.NoError(t, err)
	require.Len(t, renewals, 1)
	require.Equal(t, before.Add(365*24*time.Hour), renewals[0].NewExpiry)

	// Subname with a user registry: plan, execute, read back.
	n, err := c.Name(ctx, "example.eth")
	require.NoError(t, err)
	require.True(t, n.IsETH2LD())
	_, err = n.Subregistry(ctx)
	require.ErrorIs(t, err, v2.ErrNoSubregistry)
	plan, err := c.PlanCreateSubname(ctx, alice.Address, v2.SubnameParams{Name: "sub.example.eth", Owner: chain.Bob.Address, Resolver: resolverAddr})
	require.NoError(t, err)
	require.Len(t, plan.Pending(), 4)
	require.Equal(t, []string{v2.StepDeployRegistry, v2.StepSetSubregistry, v2.StepSetParent, v2.StepRegister}, stepNames(plan))
	receipts, err := c.ExecutePlan(ctx, opts, plan, wait)
	require.NoError(t, err)
	require.Len(t, receipts, 4)
	subReg, err := n.Subregistry(ctx)
	require.NoError(t, err)
	require.Equal(t, plan.ParentRegistry, subReg.Address())
	parent, label, err := subReg.Parent(ctx)
	require.NoError(t, err)
	require.Equal(t, chain.Config.V2.ETHRegistry, parent)
	require.Equal(t, "example", label)
	subOwner, err := subReg.Owner(ctx, "sub")
	require.NoError(t, err)
	require.Equal(t, chain.Bob.Address, subOwner)
	subName, err := c.Name(ctx, "sub.example.eth")
	require.NoError(t, err)
	require.Equal(t, subReg.Address(), subName.Registry.Address())
	regs, err := c.Helper().Registries(ctx, "sub.example.eth")
	require.NoError(t, err)
	require.Equal(t, []common.Address{{}, subReg.Address(), chain.Config.V2.ETHRegistry, chain.Config.V2.RootRegistry}, regs)
	// Re-planning is idempotent: the registry exists, only the register
	// step remains and it fails because the subname exists.
	_, err = c.PlanCreateSubname(ctx, alice.Address, v2.SubnameParams{Name: "sub.example.eth", Owner: chain.Bob.Address})
	require.ErrorIs(t, err, v2.ErrSubnameExists)
	plan2, err := c.PlanCreateSubname(ctx, alice.Address, v2.SubnameParams{Name: "other.example.eth", Owner: chain.Bob.Address})
	require.NoError(t, err)
	require.Equal(t, []string{v2.StepRegister}, stepNames(plan2))
	// Bob owns the subname and can set its records on a resolver he
	// controls; here he cannot write to Alice's resolver.
	_, err = pr.SetText(chain.Bob.Opts(t), "sub.example.eth", "x", "y")
	require.ErrorIs(t, err, v2.ErrUnauthorized)

	// Roles: granting regenerates the token id; the transfer helper
	// refuses while another account holds roles, UnsafeTransfer works.
	tokenBefore, err := ethReg.TokenID(ctx, "example")
	require.NoError(t, err)
	tx, err = n.GrantRoles(opts, v2.NewRoles(v2.RoleSetResolver), chain.Bob.Address)
	require.NoError(t, err)
	receipt = chain.Wait(t, tx)
	regen := ethReg.ParseTokenRegenerated(receipt)
	require.Len(t, regen, 1)
	require.Equal(t, 0, regen[0].OldTokenID.Cmp(tokenBefore))
	tokenAfter, err := ethReg.TokenID(ctx, "example")
	require.NoError(t, err)
	require.NotEqual(t, 0, tokenAfter.Cmp(tokenBefore))
	require.True(t, v2.SameName(tokenBefore, tokenAfter))
	stale, err := ethReg.OwnerOf(ctx, tokenBefore)
	require.NoError(t, err)
	require.Equal(t, common.Address{}, stale, "stale token id has no owner")
	bobHas, err := ethReg.HasRoles(ctx, "example", v2.NewRoles(v2.RoleSetResolver), chain.Bob.Address)
	require.NoError(t, err)
	require.True(t, bobHas)

	_, err = n.Transfer(ctx, opts, chain.Bob.Address)
	var unsafe *v2.TransferUnsafeError
	require.ErrorAs(t, err, &unsafe)
	require.Equal(t, v2.TransferUnsafeMultipleAssignees, unsafe.Reason)
	tx, err = n.RevokeRoles(opts, v2.NewRoles(v2.RoleSetResolver), chain.Bob.Address)
	require.NoError(t, err)
	chain.Wait(t, tx)
	tx, err = n.Transfer(ctx, opts, chain.Bob.Address)
	require.NoError(t, err)
	chain.Wait(t, tx)
	owner, err = n.Owner(ctx)
	require.NoError(t, err)
	require.Equal(t, chain.Bob.Address, owner)

	// Alice can no longer act; the revert decodes to a typed error.
	_, err = n.SetResolver(opts, common.Address{})
	require.True(t, errors.Is(err, v2.ErrUnauthorized), "%v", err)
}

func stepNames(p *v2.SubnamePlan) []string {
	var out []string
	for _, s := range p.Pending() {
		out = append(out, s.Name)
	}
	return out
}
