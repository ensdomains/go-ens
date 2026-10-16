// Copyright 2026 Weald Technology Trading.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0

// Package chainid probes a contract backend for its chain ID.
package chainid

import (
	"context"
	"math/big"
)

// Prober is implemented by ethclient.Client and most production backends.
type Prober interface {
	ChainID(context.Context) (*big.Int, error)
}

// Probe returns the backend's chain ID. ok is false when the backend does not
// expose one (for example a bare mock), in which case id is nil and err is
// nil. A failed probe returns the backend's error; a cancelled or expired ctx
// is surfaced as ctx.Err().
func Probe(ctx context.Context, backend any) (*big.Int, bool, error) {
	p, isProber := backend.(Prober)
	if !isProber {
		return nil, false, nil
	}
	id, err := p.ChainID(ctx)
	if err != nil {
		if ctx.Err() != nil {
			return nil, true, ctx.Err()
		}
		return nil, true, err
	}
	return id, true, nil
}
