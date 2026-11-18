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
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"sync"
	"time"

	ens "github.com/ensdomains/go-ens/v5"
	"github.com/ensdomains/go-ens/v5/contracts/v2/ethregistrar"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

var registrarBinding = ethregistrar.NewETHRegistrar()

// Secret is the commit-reveal secret of a registration.
type Secret [32]byte

// NewSecret returns a cryptographically random secret.
func NewSecret() (Secret, error) {
	var s Secret
	if _, err := rand.Read(s[:]); err != nil {
		return Secret{}, err
	}
	return s, nil
}

// Commitment holds every input of the registrar's commitment hash. The
// payment token is deliberately absent: it is chosen at registration time.
type Commitment struct {
	Label       string         // unqualified label, normalised by Hash
	Owner       common.Address // must be non-zero
	Secret      Secret
	Subregistry common.Address // zero for none
	Resolver    common.Address // zero for none
	Duration    time.Duration  // at least the registrar minimum (28 days on mainnet)
	Referrer    [32]byte
}

var commitmentArgs abi.Arguments

func init() {
	mk := func(t string) abi.Type {
		typ, err := abi.NewType(t, "", nil)
		if err != nil {
			panic(err)
		}
		return typ
	}
	commitmentArgs = abi.Arguments{
		{Type: mk("string")},
		{Type: mk("address")},
		{Type: mk("bytes32")},
		{Type: mk("address")},
		{Type: mk("address")},
		{Type: mk("uint64")},
		{Type: mk("bytes32")},
	}
}

func durationSeconds(d time.Duration) uint64 {
	if d <= 0 {
		return 0
	}
	return uint64(d / time.Second)
}

// Hash computes the commitment exactly as ETHRegistrar.makeCommitment does:
// keccak256(abi.encode(label, owner, secret, subregistry, resolver,
// duration, referrer)).
func (c Commitment) Hash() (common.Hash, error) {
	label, err := ens.Normalize(c.Label)
	if err != nil {
		return common.Hash{}, err
	}
	packed, err := commitmentArgs.Pack(label, c.Owner, [32]byte(c.Secret), c.Subregistry, c.Resolver, durationSeconds(c.Duration), c.Referrer)
	if err != nil {
		return common.Hash{}, err
	}
	return crypto.Keccak256Hash(packed), nil
}

// MakeCommitment is Commitment.Hash as a function.
func MakeCommitment(c Commitment) (common.Hash, error) { return c.Hash() }

// RegistrationParams are the inputs of a registration.
type RegistrationParams struct {
	Commitment

	// PaymentToken is the ERC-20 to pay with; zero selects the chain's
	// default (USDC).
	PaymentToken common.Address
}

// Price is a registration or renewal quote in token base units.
type Price struct {
	Token   common.Address
	Base    *big.Int
	Premium *big.Int
}

// Total returns Base plus Premium.
func (p Price) Total() *big.Int {
	t := new(big.Int)
	if p.Base != nil {
		t.Add(t, p.Base)
	}
	if p.Premium != nil {
		t.Add(t, p.Premium)
	}
	return t
}

// Constants are the registrar's immutable parameters.
type Constants struct {
	MinCommitmentAge    time.Duration
	MaxCommitmentAge    time.Duration
	MinRegisterDuration time.Duration
	MinRenewDuration    time.Duration
	GracePeriod         time.Duration
	Beneficiary         common.Address
	Oracle              common.Address
	ETHRegistry         common.Address
}

// Registrar is the ETHRegistrar: v2 registration and renewal of .eth names,
// paid in ERC-20.
type Registrar struct {
	client *Client
	addr   common.Address

	constOnce sync.Once
	consts    Constants
	constErr  error
}

func newRegistrar(c *Client, addr common.Address) *Registrar {
	return &Registrar{client: c, addr: addr}
}

// Address returns the registrar contract address.
func (r *Registrar) Address() common.Address { return r.addr }

func (r *Registrar) view(ctx context.Context, data []byte) ([]byte, error) {
	return r.client.view(ctx, r.addr, data)
}

func (r *Registrar) call(data []byte) ens.Call { return r.client.call(r.addr, data) }

func (r *Registrar) uint64View(ctx context.Context, data []byte, unpack func([]byte) (uint64, error)) (uint64, error) {
	out, err := r.view(ctx, data)
	if err != nil {
		return 0, err
	}
	return unpack(out)
}

// Constants reads the registrar parameters once and caches them.
func (r *Registrar) Constants(ctx context.Context) (Constants, error) {
	r.constOnce.Do(func() {
		var c Constants
		secs := func(data []byte, unpack func([]byte) (uint64, error)) time.Duration {
			if r.constErr != nil {
				return 0
			}
			v, err := r.uint64View(ctx, data, unpack)
			if err != nil {
				r.constErr = err
			}
			return time.Duration(v) * time.Second
		}
		c.MinCommitmentAge = secs(registrarBinding.PackMINCOMMITMENTAGE(), registrarBinding.UnpackMINCOMMITMENTAGE)
		c.MaxCommitmentAge = secs(registrarBinding.PackMAXCOMMITMENTAGE(), registrarBinding.UnpackMAXCOMMITMENTAGE)
		c.MinRegisterDuration = secs(registrarBinding.PackMINREGISTERDURATION(), registrarBinding.UnpackMINREGISTERDURATION)
		c.MinRenewDuration = secs(registrarBinding.PackMINRENEWDURATION(), registrarBinding.UnpackMINRENEWDURATION)
		c.GracePeriod = secs(registrarBinding.PackGRACEPERIOD(), registrarBinding.UnpackGRACEPERIOD)
		addr := func(data []byte, unpack func([]byte) (common.Address, error)) common.Address {
			if r.constErr != nil {
				return common.Address{}
			}
			out, err := r.view(ctx, data)
			if err != nil {
				r.constErr = err
				return common.Address{}
			}
			a, err := unpack(out)
			if err != nil {
				r.constErr = err
			}
			return a
		}
		c.Beneficiary = addr(registrarBinding.PackBENEFICIARY(), registrarBinding.UnpackBENEFICIARY)
		c.Oracle = addr(registrarBinding.PackRentPriceOracle(), registrarBinding.UnpackRentPriceOracle)
		c.ETHRegistry = addr(registrarBinding.PackETHREGISTRY(), registrarBinding.UnpackETHREGISTRY)
		r.consts = c
	})
	return r.consts, r.constErr
}

// OracleAddress returns the oracle the registrar currently prices with.
func (r *Registrar) OracleAddress(ctx context.Context) (common.Address, error) {
	out, err := r.view(ctx, registrarBinding.PackRentPriceOracle())
	if err != nil {
		return common.Address{}, err
	}
	return registrarBinding.UnpackRentPriceOracle(out)
}

// IsAvailable reports whether a label can be registered: it is neither
// registered, reserved for a v1 name, nor within the grace period.
func (r *Registrar) IsAvailable(ctx context.Context, label string) (bool, error) {
	label, err := ens.Normalize(label)
	if err != nil {
		return false, err
	}
	out, err := r.view(ctx, registrarBinding.PackIsAvailable(label))
	if err != nil {
		return false, err
	}
	return registrarBinding.UnpackIsAvailable(out)
}

// IsValid reports whether the oracle accepts the label (length and
// characters) and the label is ENSIP-15 normalised.
func (r *Registrar) IsValid(ctx context.Context, label string) (bool, error) {
	normalised, err := ens.Normalize(label)
	if err != nil || normalised != label {
		return false, nil //nolint:nilerr // an unnormalisable label is not valid
	}
	o, err := r.client.Oracle(ctx)
	if err != nil {
		return false, err
	}
	return o.IsValid(ctx, label)
}

// RegisterPrice quotes a registration in the given token. It reverts (as a
// typed error) when the name is not available or the duration too short.
func (r *Registrar) RegisterPrice(ctx context.Context, label string, duration time.Duration, token common.Address) (Price, error) {
	label, err := ens.Normalize(label)
	if err != nil {
		return Price{}, err
	}
	token = r.client.paymentToken(token)
	out, err := r.view(ctx, registrarBinding.PackGetRegisterPrice(label, durationSeconds(duration), token))
	if err != nil {
		return Price{}, err
	}
	p, err := registrarBinding.UnpackGetRegisterPrice(out)
	if err != nil {
		return Price{}, err
	}
	return Price{Token: token, Base: p.Base, Premium: p.Premium}, nil
}

// RenewPrice quotes a renewal in the given token. Renewals carry no premium.
func (r *Registrar) RenewPrice(ctx context.Context, label string, duration time.Duration, token common.Address) (*big.Int, error) {
	label, err := ens.Normalize(label)
	if err != nil {
		return nil, err
	}
	out, err := r.view(ctx, registrarBinding.PackGetRenewPrice(label, durationSeconds(duration), r.client.paymentToken(token)))
	if err != nil {
		return nil, err
	}
	return registrarBinding.UnpackGetRenewPrice(out)
}

// IsRenewable reports whether this registrar can renew the label: it is
// registered on v2 or in the post-registration grace period.
func (r *Registrar) IsRenewable(ctx context.Context, label string) (bool, error) {
	label, err := ens.Normalize(label)
	if err != nil {
		return false, err
	}
	out, err := r.view(ctx, registrarBinding.PackIsRenewable(label))
	if err != nil {
		return false, err
	}
	return registrarBinding.UnpackIsRenewable(out)
}

// RemainingGracePeriod returns how long an expired name stays renewable.
func (r *Registrar) RemainingGracePeriod(ctx context.Context, label string) (time.Duration, error) {
	label, err := ens.Normalize(label)
	if err != nil {
		return 0, err
	}
	v, err := r.uint64View(ctx, registrarBinding.PackGetRemainingGracePeriod(label), registrarBinding.UnpackGetRemainingGracePeriod)
	return time.Duration(v) * time.Second, err
}

// GracePeriod returns the registrar grace period.
func (r *Registrar) GracePeriod(ctx context.Context) (time.Duration, error) {
	c, err := r.Constants(ctx)
	return c.GracePeriod, err
}

// MakeCommitmentOnChain asks the registrar for the commitment hash; it is a
// pure function and exists mainly to verify Commitment.Hash.
func (r *Registrar) MakeCommitmentOnChain(ctx context.Context, c Commitment) (common.Hash, error) {
	label, err := ens.Normalize(c.Label)
	if err != nil {
		return common.Hash{}, err
	}
	data, err := registrarBinding.TryPackMakeCommitment(label, c.Owner, c.Secret, c.Subregistry, c.Resolver, durationSeconds(c.Duration), c.Referrer)
	if err != nil {
		return common.Hash{}, err
	}
	out, err := r.view(ctx, data)
	if err != nil {
		return common.Hash{}, err
	}
	h, err := registrarBinding.UnpackMakeCommitment(out)
	return common.Hash(h), err
}

// CommitmentAt returns when a commitment was made, or the zero time.
func (r *Registrar) CommitmentAt(ctx context.Context, commitment common.Hash) (time.Time, error) {
	v, err := r.uint64View(ctx, registrarBinding.PackCommitmentAt(commitment), registrarBinding.UnpackCommitmentAt)
	if err != nil || v == 0 {
		return time.Time{}, err
	}
	return time.Unix(int64(v), 0).UTC(), nil
}

// CommitmentWindow returns the interval during which register() accepts
// the commitment. Both times are zero when there is no commitment.
func (r *Registrar) CommitmentWindow(ctx context.Context, commitment common.Hash) (time.Time, time.Time, error) {
	at, err := r.CommitmentAt(ctx, commitment)
	if err != nil || at.IsZero() {
		return time.Time{}, time.Time{}, err
	}
	c, err := r.Constants(ctx)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}
	return at.Add(c.MinCommitmentAge), at.Add(c.MaxCommitmentAge), nil
}

// WaitForCommitment blocks until the commitment is old enough to register
// against, polling the chain's head timestamp. It returns ErrNoCommitment
// if none exists and ErrCommitmentExpired once the window has passed.
func (r *Registrar) WaitForCommitment(ctx context.Context, commitment common.Hash) error {
	from, to, err := r.CommitmentWindow(ctx, commitment)
	if err != nil {
		return err
	}
	if from.IsZero() {
		return ErrNoCommitment
	}
	for {
		now, err := r.chainTime(ctx)
		if err != nil {
			return err
		}
		if !now.Before(to) {
			return ErrCommitmentExpired
		}
		if !now.Before(from) {
			return nil
		}
		wait := from.Sub(now)
		if wait > 5*time.Second {
			wait = 5 * time.Second
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(wait):
		}
	}
}

// chainTime returns the latest block timestamp when the backend exposes
// headers, else the local clock.
func (r *Registrar) chainTime(ctx context.Context) (time.Time, error) {
	if hb, ok := r.client.Backend().(interface {
		HeaderByNumber(context.Context, *big.Int) (*types.Header, error)
	}); ok {
		h, err := hb.HeaderByNumber(ctx, nil)
		if err != nil {
			return time.Time{}, err
		}
		return time.Unix(int64(h.Time), 0).UTC(), nil
	}
	return time.Now().UTC(), nil
}

// BuildCommit packs commit(hash) and returns the hash.
func (r *Registrar) BuildCommit(c Commitment) (ens.Call, common.Hash, error) {
	h, err := c.Hash()
	if err != nil {
		return ens.Call{}, common.Hash{}, err
	}
	return r.call(registrarBinding.PackCommit(h)), h, nil
}

// Commit sends the commitment. An unexpired identical commitment makes the
// registrar revert with *UnexpiredCommitmentExistsError; use a new secret.
func (r *Registrar) Commit(opts *bind.TransactOpts, c Commitment) (*types.Transaction, common.Hash, error) {
	call, h, err := r.BuildCommit(c)
	if err != nil {
		return nil, common.Hash{}, err
	}
	tx, err := r.client.send(opts, call)
	return tx, h, err
}

// BuildRegister packs register(...) without any pre-flight.
func (r *Registrar) BuildRegister(p RegistrationParams) (ens.Call, error) {
	label, err := ens.Normalize(p.Label)
	if err != nil {
		return ens.Call{}, err
	}
	data, err := registrarBinding.TryPackRegister(label, p.Owner, p.Secret, p.Subregistry, p.Resolver, durationSeconds(p.Duration), r.client.paymentToken(p.PaymentToken), p.Referrer)
	if err != nil {
		return ens.Call{}, err
	}
	return r.call(data), nil
}

// Register completes a registration. Before sending it checks, and returns
// typed errors for, a zero owner, a short duration, the commitment window,
// availability, an ETH value on opts, and the ERC-20 allowance and balance
// of opts.From. It never approves on the caller's behalf: see
// Client.ApprovePayment.
func (r *Registrar) Register(ctx context.Context, opts *bind.TransactOpts, p RegistrationParams) (*types.Transaction, error) {
	if opts == nil {
		return nil, errors.New("v2: nil transact opts")
	}
	if opts.Value != nil && opts.Value.Sign() != 0 {
		return nil, ErrValueNotAllowed
	}
	if p.Owner == (common.Address{}) {
		return nil, &InvalidOwnerError{}
	}
	consts, err := r.Constants(ctx)
	if err != nil {
		return nil, err
	}
	if p.Duration < consts.MinRegisterDuration {
		return nil, &DurationTooShortError{Duration: p.Duration, Min: consts.MinRegisterDuration}
	}
	h, err := p.Hash()
	if err != nil {
		return nil, err
	}
	from, to, err := r.CommitmentWindow(ctx, h)
	if err != nil {
		return nil, err
	}
	if from.IsZero() {
		return nil, ErrNoCommitment
	}
	now, err := r.chainTime(ctx)
	if err != nil {
		return nil, err
	}
	if now.Before(from) {
		return nil, &CommitmentTooNewError{Commitment: h, ValidFrom: from, BlockTime: now}
	}
	if !now.Before(to) {
		return nil, &CommitmentTooOldError{Commitment: h, ValidTo: to, BlockTime: now}
	}
	available, err := r.IsAvailable(ctx, p.Label)
	if err != nil {
		return nil, err
	}
	if !available {
		return nil, &NameNotAvailableError{Label: p.Label}
	}
	price, err := r.RegisterPrice(ctx, p.Label, p.Duration, p.PaymentToken)
	if err != nil {
		return nil, err
	}
	if err := r.client.checkPayment(ctx, price.Token, opts.From, r.addr, price.Total()); err != nil {
		return nil, err
	}
	call, err := r.BuildRegister(p)
	if err != nil {
		return nil, err
	}
	return r.client.send(opts, call)
}

// RegistrationResult is decoded from the NameRegistered event.
type RegistrationResult struct {
	TokenID      *big.Int
	Label        string
	Owner        common.Address
	Subregistry  common.Address
	Resolver     common.Address
	Duration     time.Duration
	PaymentToken common.Address
	Referrer     [32]byte
	Base         *big.Int
	Premium      *big.Int
}

// ParseRegistered finds the registrar's NameRegistered event in a receipt.
func (r *Registrar) ParseRegistered(receipt *types.Receipt) (*RegistrationResult, error) {
	for _, log := range receipt.Logs {
		if log.Address != r.addr {
			continue
		}
		ev, err := registrarBinding.UnpackNameRegisteredEvent(log)
		if err != nil {
			continue
		}
		return &RegistrationResult{
			TokenID: ev.TokenId, Label: ev.Label, Owner: ev.Owner, Subregistry: ev.Subregistry, Resolver: ev.Resolver,
			Duration: time.Duration(ev.Duration) * time.Second, PaymentToken: ev.PaymentToken, Referrer: ev.Referrer,
			Base: ev.Base, Premium: ev.Premium,
		}, nil
	}
	return nil, errors.New("no NameRegistered event in receipt")
}

// WaitRegistered waits for a registration transaction and decodes its result.
func (r *Registrar) WaitRegistered(ctx context.Context, tx *types.Transaction) (*RegistrationResult, error) {
	receipt, err := ens.WaitMinedOK(ctx, r.client.Backend(), tx)
	if err != nil {
		return nil, err
	}
	return r.ParseRegistered(receipt)
}

// Renewer implementation.

// BuildRenew packs renew((label,duration,referrer), token).
func (r *Registrar) BuildRenew(rd RenewData, token common.Address) (ens.Call, error) {
	wire, err := rd.wire()
	if err != nil {
		return ens.Call{}, err
	}
	data, err := registrarBinding.TryPackRenew(ethregistrar.RenewData(wire), r.client.paymentToken(token))
	if err != nil {
		return ens.Call{}, err
	}
	return r.call(data), nil
}

// Renew extends a v2-registered name. See Client.Renew for state-based
// routing between this registrar and the v1 renewer.
func (r *Registrar) Renew(opts *bind.TransactOpts, rd RenewData, token common.Address) (*types.Transaction, error) {
	c, err := r.BuildRenew(rd, token)
	if err != nil {
		return nil, err
	}
	return r.client.send(opts, c)
}

// BuildRenewBatch packs renewBatch(rds, token).
func (r *Registrar) BuildRenewBatch(rds []RenewData, token common.Address) (ens.Call, error) {
	wires := make([]ethregistrar.RenewData, 0, len(rds))
	for _, rd := range rds {
		w, err := rd.wire()
		if err != nil {
			return ens.Call{}, err
		}
		wires = append(wires, ethregistrar.RenewData(w))
	}
	data, err := registrarBinding.TryPackRenewBatch(wires, r.client.paymentToken(token))
	if err != nil {
		return ens.Call{}, err
	}
	return r.call(data), nil
}

// RenewBatch renews several v2-registered names in one transaction.
func (r *Registrar) RenewBatch(opts *bind.TransactOpts, rds []RenewData, token common.Address) (*types.Transaction, error) {
	c, err := r.BuildRenewBatch(rds, token)
	if err != nil {
		return nil, err
	}
	return r.client.send(opts, c)
}

// ParseRenewed finds every NameRenewed event of the registrar in a receipt.
func (r *Registrar) ParseRenewed(receipt *types.Receipt) ([]RenewalResult, error) {
	var out []RenewalResult
	for _, log := range receipt.Logs {
		if log.Address != r.addr {
			continue
		}
		ev, err := registrarBinding.UnpackNameRenewedEvent(log)
		if err != nil {
			continue
		}
		out = append(out, RenewalResult{
			TokenID: ev.TokenId, Label: ev.Label, Duration: time.Duration(ev.Duration) * time.Second,
			NewExpiry: time.Unix(int64(ev.NewExpiry), 0).UTC(), PaymentToken: ev.PaymentToken, Referrer: ev.Referrer, Amount: ev.Amount,
		})
	}
	if len(out) == 0 {
		return nil, fmt.Errorf("no NameRenewed event in receipt")
	}
	return out, nil
}
