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
	"github.com/ensdomains/go-ens/v5/contracts/v2/rentpriceoracle"
	"github.com/ethereum/go-ethereum/common"
)

var oracleBinding = rentpriceoracle.NewStandardRentPriceOracle()

// PriceDecimals is the number of decimals of the oracle's standard unit:
// prices are quoted in USD times 1e12 before conversion to a token.
const PriceDecimals = 12

// Oracle is the StandardRentPriceOracle: label validity, base rates by
// label length, multi-year discounts, the post-expiry premium and the
// conversion of standard units to payment tokens.
type Oracle struct {
	client *Client
	addr   common.Address
}

func newOracle(c *Client, addr common.Address) *Oracle { return &Oracle{client: c, addr: addr} }

// Address returns the oracle contract address.
func (o *Oracle) Address() common.Address { return o.addr }

func (o *Oracle) view(ctx context.Context, data []byte) ([]byte, error) {
	return o.client.view(ctx, o.addr, data)
}

// IsValid reports whether the oracle prices the label at all (it rejects
// labels of fewer than three codepoints on mainnet).
func (o *Oracle) IsValid(ctx context.Context, label string) (bool, error) {
	label, err := ens.Normalize(label)
	if err != nil {
		return false, nil //nolint:nilerr // an unnormalisable label is not valid
	}
	out, err := o.view(ctx, oracleBinding.PackIsValid(label))
	if err != nil {
		return false, err
	}
	return oracleBinding.UnpackIsValid(out)
}

// IsPaymentToken reports whether the oracle accepts a token.
func (o *Oracle) IsPaymentToken(ctx context.Context, token common.Address) (bool, error) {
	out, err := o.view(ctx, oracleBinding.PackIsPaymentToken(token))
	if err != nil {
		return false, err
	}
	return oracleBinding.UnpackIsPaymentToken(out)
}

// PaymentTokenRatio returns the numerator and denominator that convert
// standard units to the token: amount = ceil(value * numer / denom).
func (o *Oracle) PaymentTokenRatio(ctx context.Context, token common.Address) (*big.Int, *big.Int, error) {
	out, err := o.view(ctx, oracleBinding.PackGetPaymentTokenRatio(token))
	if err != nil {
		return nil, nil, err
	}
	r, err := oracleBinding.UnpackGetPaymentTokenRatio(out)
	if err != nil {
		return nil, nil, err
	}
	return r.Numer, r.Denom, nil
}

// BaseRates returns the per-second price in standard units indexed by
// label length in codepoints (index 0 is unused; the last entry applies to
// every longer label).
func (o *Oracle) BaseRates(ctx context.Context) ([]*big.Int, error) {
	out, err := o.view(ctx, oracleBinding.PackGetBaseRates())
	if err != nil {
		return nil, err
	}
	return oracleBinding.UnpackGetBaseRates(out)
}

// DiscountPoint is a multi-year discount: durations of at least Duration
// pay Numer / DiscountDenominator of the base price.
type DiscountPoint struct {
	Duration time.Duration
	Numer    *big.Int
}

// DiscountPoints returns the discount schedule.
func (o *Oracle) DiscountPoints(ctx context.Context) ([]DiscountPoint, error) {
	out, err := o.view(ctx, oracleBinding.PackGetDiscountPoints())
	if err != nil {
		return nil, err
	}
	pts, err := oracleBinding.UnpackGetDiscountPoints(out)
	if err != nil {
		return nil, err
	}
	res := make([]DiscountPoint, len(pts))
	for i, p := range pts {
		res[i] = DiscountPoint{Duration: time.Duration(p.Duration) * time.Second, Numer: p.Numer}
	}
	return res, nil
}

// DiscountDenominator returns the denominator of the discount numerators.
func (o *Oracle) DiscountDenominator(ctx context.Context) (*big.Int, error) {
	out, err := o.view(ctx, oracleBinding.PackDISCOUNTDENOMINATOR())
	if err != nil {
		return nil, err
	}
	return oracleBinding.UnpackDISCOUNTDENOMINATOR(out)
}

// ApplyDiscount applies the duration discount to a standard-unit amount.
func (o *Oracle) ApplyDiscount(ctx context.Context, amount *big.Int, duration time.Duration) (*big.Int, error) {
	out, err := o.view(ctx, oracleBinding.PackApplyDiscount(amount, durationSeconds(duration)))
	if err != nil {
		return nil, err
	}
	return oracleBinding.UnpackApplyDiscount(out)
}

// BasePrice returns the discounted base price of a label for a duration,
// in standard units.
func (o *Oracle) BasePrice(ctx context.Context, label string, duration time.Duration) (*big.Int, error) {
	label, err := ens.Normalize(label)
	if err != nil {
		return nil, err
	}
	out, err := o.view(ctx, oracleBinding.PackGetBasePrice(label, durationSeconds(duration)))
	if err != nil {
		return nil, err
	}
	return oracleBinding.UnpackGetBasePrice(out)
}

// PremiumPriceAfter returns the premium, in standard units, at a given
// time after a name became available.
func (o *Oracle) PremiumPriceAfter(ctx context.Context, since time.Duration) (*big.Int, error) {
	out, err := o.view(ctx, oracleBinding.PackGetPremiumPriceAfter(durationSeconds(since)))
	if err != nil {
		return nil, err
	}
	return oracleBinding.UnpackGetPremiumPriceAfter(out)
}

// PremiumParameters returns the initial premium (standard units), its
// halving period and the period after which it reaches zero.
func (o *Oracle) PremiumParameters(ctx context.Context) (*big.Int, time.Duration, time.Duration, error) {
	out, err := o.view(ctx, oracleBinding.PackPREMIUMPRICEINITIAL())
	if err != nil {
		return nil, 0, 0, err
	}
	initial, err := oracleBinding.UnpackPREMIUMPRICEINITIAL(out)
	if err != nil {
		return nil, 0, 0, err
	}
	if out, err = o.view(ctx, oracleBinding.PackPREMIUMHALVINGPERIOD()); err != nil {
		return nil, 0, 0, err
	}
	h, err := oracleBinding.UnpackPREMIUMHALVINGPERIOD(out)
	if err != nil {
		return nil, 0, 0, err
	}
	if out, err = o.view(ctx, oracleBinding.PackPREMIUMPERIOD()); err != nil {
		return nil, 0, 0, err
	}
	p, err := oracleBinding.UnpackPREMIUMPERIOD(out)
	if err != nil {
		return nil, 0, 0, err
	}
	return initial, time.Duration(h) * time.Second, time.Duration(p) * time.Second, nil
}

// ConvertUnits converts a standard-unit amount to token units.
func (o *Oracle) ConvertUnits(ctx context.Context, amount *big.Int, token common.Address) (*big.Int, error) {
	out, err := o.view(ctx, oracleBinding.PackConvertUnits(amount, token))
	if err != nil {
		return nil, err
	}
	return oracleBinding.UnpackConvertUnits(out)
}

// Length returns the label length the oracle prices on (codepoints).
func (o *Oracle) Length(ctx context.Context, label string) (uint64, error) {
	label, err := ens.Normalize(label)
	if err != nil {
		return 0, err
	}
	out, err := o.view(ctx, oracleBinding.PackGetLength(label))
	if err != nil {
		return 0, err
	}
	n, err := oracleBinding.UnpackGetLength(out)
	if err != nil {
		return 0, err
	}
	return n.Uint64(), nil
}
