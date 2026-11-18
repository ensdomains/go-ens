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
	"math/big"
	"time"

	ens "github.com/ensdomains/go-ens/v5"
	"github.com/ensdomains/go-ens/v5/contracts/v2/ethrenewerv1"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

var renewerV1Binding = ethrenewerv1.NewETHRenewerV1()

// RenewData describes one renewal: the wire tuple (string,uint64,bytes32).
type RenewData struct {
	Label    string
	Duration time.Duration
	Referrer [32]byte
}

type renewWire struct {
	Label    string
	Duration uint64
	Referrer [32]byte
}

func (rd RenewData) wire() (renewWire, error) {
	label, err := ens.Normalize(rd.Label)
	if err != nil {
		return renewWire{}, err
	}
	return renewWire{Label: label, Duration: durationSeconds(rd.Duration), Referrer: rd.Referrer}, nil
}

// RenewalResult is decoded from a NameRenewed event.
type RenewalResult struct {
	TokenID      *big.Int
	Label        string
	Duration     time.Duration
	NewExpiry    time.Time
	PaymentToken common.Address
	Referrer     [32]byte
	Amount       *big.Int
}

// Renewer is the IETHRenewer interface, implemented by the ETHRegistrar for
// v2-registered names and by ETHRenewerV1 for unmigrated v1 names.
type Renewer interface {
	Address() common.Address
	IsRenewable(ctx context.Context, label string) (bool, error)
	RenewPrice(ctx context.Context, label string, duration time.Duration, token common.Address) (*big.Int, error)
	RemainingGracePeriod(ctx context.Context, label string) (time.Duration, error)
	GracePeriod(ctx context.Context) (time.Duration, error)
	BuildRenew(rd RenewData, token common.Address) (ens.Call, error)
	Renew(opts *bind.TransactOpts, rd RenewData, token common.Address) (*types.Transaction, error)
	BuildRenewBatch(rds []RenewData, token common.Address) (ens.Call, error)
	RenewBatch(opts *bind.TransactOpts, rds []RenewData, token common.Address) (*types.Transaction, error)
	ParseRenewed(receipt *types.Receipt) ([]RenewalResult, error)
}

var (
	_ Renewer = (*Registrar)(nil)
	_ Renewer = (*RenewerV1)(nil)
)

// RenewerV1 is ETHRenewerV1: it renews names that are still held in the v1
// registrar (reserved on v2), keeping the v1 expiry and NameWrapper in sync.
type RenewerV1 struct {
	client *Client
	addr   common.Address
}

func newRenewerV1(c *Client, addr common.Address) *RenewerV1 {
	return &RenewerV1{client: c, addr: addr}
}

// Address returns the renewer contract address.
func (r *RenewerV1) Address() common.Address { return r.addr }

func (r *RenewerV1) view(ctx context.Context, data []byte) ([]byte, error) {
	return r.client.view(ctx, r.addr, data)
}

func (r *RenewerV1) call(data []byte) ens.Call { return r.client.call(r.addr, data) }

// IsRenewable reports whether the label is reserved for a v1 name or in the
// reserved grace period.
func (r *RenewerV1) IsRenewable(ctx context.Context, label string) (bool, error) {
	label, err := ens.Normalize(label)
	if err != nil {
		return false, err
	}
	out, err := r.view(ctx, renewerV1Binding.PackIsRenewable(label))
	if err != nil {
		return false, err
	}
	return renewerV1Binding.UnpackIsRenewable(out)
}

// RenewPrice quotes a renewal in the given token.
func (r *RenewerV1) RenewPrice(ctx context.Context, label string, duration time.Duration, token common.Address) (*big.Int, error) {
	label, err := ens.Normalize(label)
	if err != nil {
		return nil, err
	}
	out, err := r.view(ctx, renewerV1Binding.PackGetRenewPrice(label, durationSeconds(duration), r.client.paymentToken(token)))
	if err != nil {
		return nil, err
	}
	return renewerV1Binding.UnpackGetRenewPrice(out)
}

// RemainingGracePeriod returns how long the v1 name stays renewable.
func (r *RenewerV1) RemainingGracePeriod(ctx context.Context, label string) (time.Duration, error) {
	label, err := ens.Normalize(label)
	if err != nil {
		return 0, err
	}
	out, err := r.view(ctx, renewerV1Binding.PackGetRemainingGracePeriod(label))
	if err != nil {
		return 0, err
	}
	v, err := renewerV1Binding.UnpackGetRemainingGracePeriod(out)
	return time.Duration(v) * time.Second, err
}

// GracePeriod returns the renewer's grace period (the v2 grace plus the
// migration bonus).
func (r *RenewerV1) GracePeriod(ctx context.Context) (time.Duration, error) {
	out, err := r.view(ctx, renewerV1Binding.PackGRACEPERIOD())
	if err != nil {
		return 0, err
	}
	v, err := renewerV1Binding.UnpackGRACEPERIOD(out)
	return time.Duration(v) * time.Second, err
}

// BuildRenew packs renew((label,duration,referrer), token).
func (r *RenewerV1) BuildRenew(rd RenewData, token common.Address) (ens.Call, error) {
	w, err := rd.wire()
	if err != nil {
		return ens.Call{}, err
	}
	data, err := renewerV1Binding.TryPackRenew(ethrenewerv1.RenewData(w), r.client.paymentToken(token))
	if err != nil {
		return ens.Call{}, err
	}
	return r.call(data), nil
}

// Renew extends an unmigrated v1 name.
func (r *RenewerV1) Renew(opts *bind.TransactOpts, rd RenewData, token common.Address) (*types.Transaction, error) {
	c, err := r.BuildRenew(rd, token)
	if err != nil {
		return nil, err
	}
	return r.client.send(opts, c)
}

// BuildRenewBatch packs renewBatch(rds, token).
func (r *RenewerV1) BuildRenewBatch(rds []RenewData, token common.Address) (ens.Call, error) {
	wires := make([]ethrenewerv1.RenewData, 0, len(rds))
	for _, rd := range rds {
		w, err := rd.wire()
		if err != nil {
			return ens.Call{}, err
		}
		wires = append(wires, ethrenewerv1.RenewData(w))
	}
	data, err := renewerV1Binding.TryPackRenewBatch(wires, r.client.paymentToken(token))
	if err != nil {
		return ens.Call{}, err
	}
	return r.call(data), nil
}

// RenewBatch renews several unmigrated v1 names in one transaction.
func (r *RenewerV1) RenewBatch(opts *bind.TransactOpts, rds []RenewData, token common.Address) (*types.Transaction, error) {
	c, err := r.BuildRenewBatch(rds, token)
	if err != nil {
		return nil, err
	}
	return r.client.send(opts, c)
}

// BuildSyncWrapper packs syncWrapper(labels), which refreshes NameWrapper
// expiries of wrapped v1 names after a renewal.
func (r *RenewerV1) BuildSyncWrapper(labels []string) (ens.Call, error) {
	norm := make([]string, 0, len(labels))
	for _, l := range labels {
		n, err := ens.Normalize(l)
		if err != nil {
			return ens.Call{}, err
		}
		norm = append(norm, n)
	}
	data, err := renewerV1Binding.TryPackSyncWrapper(norm)
	if err != nil {
		return ens.Call{}, err
	}
	return r.call(data), nil
}

// SyncWrapper sends syncWrapper(labels).
func (r *RenewerV1) SyncWrapper(opts *bind.TransactOpts, labels []string) (*types.Transaction, error) {
	c, err := r.BuildSyncWrapper(labels)
	if err != nil {
		return nil, err
	}
	return r.client.send(opts, c)
}

// ParseRenewed finds every NameRenewed event of the renewer in a receipt.
func (r *RenewerV1) ParseRenewed(receipt *types.Receipt) ([]RenewalResult, error) {
	var out []RenewalResult
	for _, log := range receipt.Logs {
		if log.Address != r.addr {
			continue
		}
		ev, err := renewerV1Binding.UnpackNameRenewedEvent(log)
		if err != nil {
			continue
		}
		out = append(out, RenewalResult{
			TokenID: ev.TokenId, Label: ev.Label, Duration: time.Duration(ev.Duration) * time.Second,
			NewExpiry: time.Unix(int64(ev.NewExpiry), 0).UTC(), PaymentToken: ev.PaymentToken, Referrer: ev.Referrer, Amount: ev.Amount,
		})
	}
	if len(out) == 0 {
		return nil, &NameNotRenewableError{}
	}
	return out, nil
}

// Routing.

// ChooseRenewer picks the renewer for a .eth label from the .eth registry's
// state and confirms the choice with the renewer's own isRenewable:
//
//	REGISTERED, or AVAILABLE with a previous owner inside the grace period: ETHRegistrar
//	RESERVED, or AVAILABLE with no owner inside the grace period:           ETHRenewerV1
//	anything else:                                                           ens.ErrNameNotRenewable
func (c *Client) ChooseRenewer(ctx context.Context, label string) (Renewer, *State, error) {
	st, err := c.ETHRegistry().State(ctx, label)
	if err != nil {
		return nil, nil, err
	}
	var candidate Renewer
	switch st.Status {
	case StatusRegistered:
		candidate = c.ETHRegistrar()
	case StatusReserved:
		candidate = c.RenewerV1()
	case StatusAvailable:
		fallthrough
	default:
		if st.Expiry.IsZero() {
			return nil, st, &NameNotRenewableError{Label: label}
		}
		grace, err := c.ETHRegistrar().GracePeriod(ctx)
		if err != nil {
			return nil, st, err
		}
		if time.Now().After(st.Expiry.Add(grace)) {
			return nil, st, &NameNotRenewableError{Label: label}
		}
		if st.LatestOwner != (common.Address{}) {
			candidate = c.ETHRegistrar()
		} else {
			candidate = c.RenewerV1()
		}
	}
	ok, err := candidate.IsRenewable(ctx, label)
	if err != nil {
		return nil, st, err
	}
	if !ok {
		return nil, st, &NameNotRenewableError{Label: label}
	}
	return candidate, st, nil
}

// Renew routes a renewal to the right renewer, checks the ERC-20 allowance
// and balance of opts.From, and sends it.
func (c *Client) Renew(ctx context.Context, opts *bind.TransactOpts, rd RenewData, token common.Address) (*types.Transaction, error) {
	renewer, _, err := c.ChooseRenewer(ctx, rd.Label)
	if err != nil {
		return nil, err
	}
	token = c.paymentToken(token)
	amount, err := renewer.RenewPrice(ctx, rd.Label, rd.Duration, token)
	if err != nil {
		return nil, err
	}
	if err := c.checkPayment(ctx, token, opts.From, renewer.Address(), amount); err != nil {
		return nil, err
	}
	return renewer.Renew(opts, rd, token)
}

// RenewBatch renews several labels that all route to the same renewer. It
// returns ErrMixedRenewers when they do not.
func (c *Client) RenewBatch(ctx context.Context, opts *bind.TransactOpts, rds []RenewData, token common.Address) (*types.Transaction, error) {
	if len(rds) == 0 {
		return nil, &NameNotRenewableError{}
	}
	var renewer Renewer
	token = c.paymentToken(token)
	total := new(big.Int)
	for _, rd := range rds {
		r, _, err := c.ChooseRenewer(ctx, rd.Label)
		if err != nil {
			return nil, err
		}
		if renewer != nil && r.Address() != renewer.Address() {
			return nil, ErrMixedRenewers
		}
		renewer = r
		amount, err := r.RenewPrice(ctx, rd.Label, rd.Duration, token)
		if err != nil {
			return nil, err
		}
		total.Add(total, amount)
	}
	if err := c.checkPayment(ctx, token, opts.From, renewer.Address(), total); err != nil {
		return nil, err
	}
	return renewer.RenewBatch(opts, rds, token)
}
