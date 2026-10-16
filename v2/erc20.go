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

	ens "github.com/ensdomains/go-ens/v5"
	"github.com/ensdomains/go-ens/v5/contracts/v2/erc20"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

var erc20Binding = erc20.NewERC20()

// ERC20 is a minimal token handle for registration payments.
type ERC20 struct {
	client *Client
	addr   common.Address
}

func newERC20(c *Client, addr common.Address) *ERC20 { return &ERC20{client: c, addr: addr} }

// Address returns the token contract address.
func (t *ERC20) Address() common.Address { return t.addr }

// Allowance returns how much spender may pull from owner.
func (t *ERC20) Allowance(ctx context.Context, owner, spender common.Address) (*big.Int, error) {
	out, err := t.client.view(ctx, t.addr, erc20Binding.PackAllowance(owner, spender))
	if err != nil {
		return nil, err
	}
	return erc20Binding.UnpackAllowance(out)
}

// BalanceOf returns the balance of an account.
func (t *ERC20) BalanceOf(ctx context.Context, account common.Address) (*big.Int, error) {
	out, err := t.client.view(ctx, t.addr, erc20Binding.PackBalanceOf(account))
	if err != nil {
		return nil, err
	}
	return erc20Binding.UnpackBalanceOf(out)
}

// Decimals returns the token's decimals.
func (t *ERC20) Decimals(ctx context.Context) (uint8, error) {
	out, err := t.client.view(ctx, t.addr, erc20Binding.PackDecimals())
	if err != nil {
		return 0, err
	}
	return erc20Binding.UnpackDecimals(out)
}

// Symbol returns the token's symbol.
func (t *ERC20) Symbol(ctx context.Context) (string, error) {
	out, err := t.client.view(ctx, t.addr, erc20Binding.PackSymbol())
	if err != nil {
		return "", err
	}
	return erc20Binding.UnpackSymbol(out)
}

// BuildApprove packs approve(spender, amount).
func (t *ERC20) BuildApprove(spender common.Address, amount *big.Int) ens.Call {
	return t.client.call(t.addr, erc20Binding.PackApprove(spender, amount))
}

// Approve lets spender pull up to amount from the sender.
func (t *ERC20) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return t.client.send(opts, t.BuildApprove(spender, amount))
}

// Mint calls mint(to, amount), which only the Sepolia mock tokens expose.
func (t *ERC20) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return t.client.send(opts, t.client.call(t.addr, erc20Binding.PackMint(to, amount)))
}

// PaymentCheck is the outcome of a payment pre-flight.
type PaymentCheck struct {
	Token               common.Address
	Required            *big.Int
	Allowance           *big.Int
	Balance             *big.Int
	NeedsApproval       bool
	InsufficientBalance bool
}

// CheckPayment reads the allowance payer granted spender and payer's
// balance for a required amount.
func (c *Client) CheckPayment(ctx context.Context, token, payer, spender common.Address, required *big.Int) (*PaymentCheck, error) {
	token = c.paymentToken(token)
	t := c.Token(token)
	allowance, err := t.Allowance(ctx, payer, spender)
	if err != nil {
		return nil, err
	}
	balance, err := t.BalanceOf(ctx, payer)
	if err != nil {
		return nil, err
	}
	return &PaymentCheck{
		Token: token, Required: required, Allowance: allowance, Balance: balance,
		NeedsApproval:       allowance.Cmp(required) < 0,
		InsufficientBalance: balance.Cmp(required) < 0,
	}, nil
}

// checkPayment is CheckPayment as a pre-flight that returns typed errors.
func (c *Client) checkPayment(ctx context.Context, token, payer, spender common.Address, required *big.Int) error {
	if payer == (common.Address{}) {
		return nil // no sender to check against (for example building offline)
	}
	chk, err := c.CheckPayment(ctx, token, payer, spender, required)
	if err != nil {
		return err
	}
	if chk.InsufficientBalance {
		return &InsufficientBalanceError{Token: chk.Token, Required: required, Have: chk.Balance}
	}
	if chk.NeedsApproval {
		return &InsufficientAllowanceError{Token: chk.Token, Spender: spender, Required: required, Have: chk.Allowance}
	}
	return nil
}

// ApprovePayment approves spender (the ETHRegistrar or ETHRenewerV1) to
// pull amount of token from the sender.
func (c *Client) ApprovePayment(opts *bind.TransactOpts, token, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return c.Token(c.paymentToken(token)).Approve(opts, spender, amount)
}

// EnsureAllowance approves spender for amount unless the sender's allowance
// already covers it, in which case it returns a nil transaction.
func (c *Client) EnsureAllowance(ctx context.Context, opts *bind.TransactOpts, token, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	chk, err := c.CheckPayment(ctx, token, opts.From, spender, amount)
	if err != nil {
		return nil, err
	}
	if !chk.NeedsApproval {
		return nil, nil
	}
	return c.ApprovePayment(opts, token, spender, amount)
}

// paymentToken substitutes the chain default for the zero address.
func (c *Client) paymentToken(token common.Address) common.Address {
	if token == (common.Address{}) {
		return c.cfg.DefaultPaymentToken
	}
	return token
}

// FormatPrice renders a price as a decimal amount with the token symbol,
// using the configured decimals when known and reading them otherwise.
func (c *Client) FormatPrice(ctx context.Context, p Price) (string, error) {
	var symbol string
	var decimals uint8
	if pt, ok := c.cfg.PaymentTokenByAddress(p.Token); ok {
		symbol, decimals = pt.Symbol, pt.Decimals
	} else {
		t := c.Token(p.Token)
		var err error
		if decimals, err = t.Decimals(ctx); err != nil {
			return "", err
		}
		if symbol, err = t.Symbol(ctx); err != nil {
			return "", err
		}
	}
	return FormatAmount(p.Total(), decimals) + " " + symbol, nil
}

// FormatAmount renders a token amount as a decimal string.
func FormatAmount(amount *big.Int, decimals uint8) string {
	if amount == nil {
		return "0"
	}
	s := amount.String()
	if decimals == 0 {
		return s
	}
	neg := strings.HasPrefix(s, "-")
	s = strings.TrimPrefix(s, "-")
	for len(s) <= int(decimals) {
		s = "0" + s
	}
	cut := len(s) - int(decimals)
	frac := strings.TrimRight(s[cut:], "0")
	out := s[:cut]
	if frac != "" {
		out += "." + frac
	}
	if neg {
		out = "-" + out
	}
	return out
}

// String renders the quote for logs.
func (p Price) String() string {
	return fmt.Sprintf("base=%s premium=%s token=%s", p.Base, p.Premium, p.Token.Hex())
}
