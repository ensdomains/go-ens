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
	"errors"
	"fmt"
	"math/big"

	"github.com/ensdomains/go-ens/v5/internal/revert"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// abiEmpty is the ABI handed to bound contracts that only send raw calldata.
var abiEmpty abi.ABI

// Call is an unsigned contract call: the output of every Build* function in
// the v1 and v2 packages. It can be sent with Client.Send, simulated with
// Client.Simulate, or handed to a multisig, batcher or offline signer.
type Call struct {
	To    common.Address
	Data  []byte
	Value *big.Int // nil means zero

	decoder *revert.Decoder
}

// NewCall builds a Call. decoder may be nil; the v1 and v2 packages attach
// their own so reverts decode to their typed errors.
func NewCall(to common.Address, data []byte, value *big.Int, decoder *revert.Decoder) Call {
	return Call{To: to, Data: data, Value: value, decoder: decoder}
}

// Decode translates a revert error from this call into a typed error,
// returning err unchanged when it carries no recognised payload.
func (c Call) Decode(err error) error {
	if c.decoder != nil {
		return c.decoder.Translate(err)
	}
	return urDecoder.Translate(err)
}

// CallMsg converts the call into an ethereum.CallMsg from the given sender.
func (c Call) CallMsg(from common.Address) ethereum.CallMsg {
	to := c.To
	return ethereum.CallMsg{From: from, To: &to, Data: c.Data, Value: c.Value}
}

// Send signs and submits the call. Gas estimation reverts surface as typed
// errors. opts.Value, if set, overrides call.Value.
func (c *Client) Send(opts *bind.TransactOpts, call Call) (*types.Transaction, error) {
	if opts == nil {
		return nil, errors.New("ens: nil transact opts")
	}
	o := *opts
	if o.Value == nil {
		o.Value = call.Value
	}
	contract := bind.NewBoundContract(call.To, abiEmpty, c.backend, c.backend, c.backend)
	tx, err := contract.RawTransact(&o, call.Data)
	if err != nil {
		return nil, call.Decode(err)
	}
	return tx, nil
}

// Simulate executes the call with eth_call from the given sender and
// returns the raw result, decoding reverts into typed errors.
func (c *Client) Simulate(ctx context.Context, from common.Address, call Call) ([]byte, error) {
	out, err := c.backend.CallContract(ctx, call.CallMsg(from), nil)
	if err != nil {
		return nil, call.Decode(err)
	}
	return out, nil
}

// ExplainFailedTx re-executes a mined, failed transaction as an eth_call at
// the block before it was mined and returns the decoded revert reason, or
// nil if the replay succeeds. decoder may be nil.
func (c *Client) ExplainFailedTx(ctx context.Context, tx *types.Transaction, receipt *types.Receipt, decoder *revert.Decoder) error {
	if receipt == nil || tx == nil {
		return errors.New("ens: nil transaction or receipt")
	}
	if receipt.Status == types.ReceiptStatusSuccessful {
		return nil
	}
	from, err := types.Sender(types.LatestSignerForChainID(tx.ChainId()), tx)
	if err != nil {
		return fmt.Errorf("recover sender: %w", err)
	}
	block := new(big.Int).Sub(receipt.BlockNumber, big.NewInt(1))
	_, err = c.backend.CallContract(ctx, ethereum.CallMsg{From: from, To: tx.To(), Data: tx.Data(), Value: tx.Value(), Gas: tx.Gas()}, block)
	if err == nil {
		return nil
	}
	if decoder != nil {
		return decoder.Translate(err)
	}
	return urDecoder.Translate(err)
}

// WaitMinedOK waits for tx to be mined and returns an error if it failed.
// The backend must expose TransactionReceipt (ethclient.Client does).
func WaitMinedOK(ctx context.Context, backend any, tx *types.Transaction) (*types.Receipt, error) {
	db, ok := backend.(bind.DeployBackend)
	if !ok {
		return nil, errors.New("ens: backend cannot fetch transaction receipts")
	}
	receipt, err := bind.WaitMined(ctx, db, tx)
	if err != nil {
		return nil, err
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		return receipt, fmt.Errorf("transaction %s failed", tx.Hash().Hex())
	}
	return receipt, nil
}
