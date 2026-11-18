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
	"math/big"

	ens "github.com/ensdomains/go-ens/v5"
	"github.com/ensdomains/go-ens/v5/contracts/v2/permissionedregistry"
	"github.com/ensdomains/go-ens/v5/contracts/v2/verifiablefactory"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

var factoryBinding = verifiablefactory.NewVerifiableFactory()

// Factory is the VerifiableFactory, which deploys verifiable UUPS proxies
// of the UserRegistry, WrapperRegistry and PermissionedResolver
// implementations at CREATE2 addresses derived from (deployer, salt).
type Factory struct {
	client *Client
	addr   common.Address
}

func newFactory(c *Client, addr common.Address) *Factory { return &Factory{client: c, addr: addr} }

// Address returns the factory contract address.
func (f *Factory) Address() common.Address { return f.addr }

// RandomSalt returns a random 256-bit salt. Proxy addresses depend on the
// deployer and salt, so a salt is single-use per deployer.
func RandomSalt() (*big.Int, error) {
	var b [32]byte
	if _, err := rand.Read(b[:]); err != nil {
		return nil, err
	}
	return new(big.Int).SetBytes(b[:]), nil
}

var saltArgs abi.Arguments

func init() {
	b32, _ := abi.NewType("bytes32", "", nil)
	addr, _ := abi.NewType("address", "", nil)
	u, _ := abi.NewType("uint256", "", nil)
	saltArgs = abi.Arguments{{Type: b32}, {Type: addr}, {Type: u}}
	_ = b32
	registrySaltArgs = abi.Arguments{{Type: b32}, {Type: b32}, {Type: u}}
}

var registrySaltArgs abi.Arguments

// UserRegistrySalt is the deterministic salt convention of the ENS
// deployment scripts for the registry of a name:
// keccak256(abi.encode(keccak256("UserRegistry"), namehash(name), version)).
// Deterministic salts make subname creation idempotent for one deployer.
func UserRegistrySalt(name string, version uint64) (*big.Int, error) {
	node, err := ens.NameHash(name)
	if err != nil {
		return nil, err
	}
	packed, err := registrySaltArgs.Pack(crypto.Keccak256Hash([]byte("UserRegistry")), node, new(big.Int).SetUint64(version))
	if err != nil {
		return nil, err
	}
	return new(big.Int).SetBytes(crypto.Keccak256(packed)), nil
}

// OwnedResolverSalt is a deterministic salt for an account's resolver:
// keccak256(abi.encode(keccak256("OwnedResolver"), owner, version)).
func OwnedResolverSalt(owner common.Address, version uint64) *big.Int {
	packed, _ := saltArgs.Pack(crypto.Keccak256Hash([]byte("OwnedResolver")), owner, new(big.Int).SetUint64(version))
	return new(big.Int).SetBytes(crypto.Keccak256(packed))
}

// PredictProxyAddress asks the factory for the address a deployer's salt
// yields.
func (f *Factory) PredictProxyAddress(ctx context.Context, deployer common.Address, salt *big.Int) (common.Address, error) {
	out, err := f.client.view(ctx, f.addr, factoryBinding.PackPredictProxyAddress(deployer, salt))
	if err != nil {
		return common.Address{}, err
	}
	return factoryBinding.UnpackPredictProxyAddress(out)
}

// VerifyContract returns the implementation behind a proxy the factory
// deployed, or *VerificationFailedError.
func (f *Factory) VerifyContract(ctx context.Context, proxy common.Address) (common.Address, error) {
	out, err := f.client.view(ctx, f.addr, factoryBinding.PackVerifyContract(proxy))
	if err != nil {
		return common.Address{}, err
	}
	return factoryBinding.UnpackVerifyContract(out)
}

// BuildDeployProxy packs deployProxy(implementation, salt, initData).
func (f *Factory) BuildDeployProxy(impl common.Address, salt *big.Int, initData []byte) ens.Call {
	return f.client.call(f.addr, factoryBinding.PackDeployProxy(impl, salt, initData))
}

// DeployProxy deploys a proxy of impl initialised with initData. The
// resulting address is PredictProxyAddress(opts.From, salt); reusing a salt
// reverts without data, which Decode maps to ErrSaltInUse when code already
// exists at the predicted address.
func (f *Factory) DeployProxy(ctx context.Context, opts *bind.TransactOpts, impl common.Address, salt *big.Int, initData []byte) (*types.Transaction, common.Address, error) {
	predicted, err := f.PredictProxyAddress(ctx, opts.From, salt)
	if err != nil {
		return nil, common.Address{}, err
	}
	tx, err := f.client.send(opts, f.BuildDeployProxy(impl, salt, initData))
	if err != nil {
		var empty *ens.EmptyRevertError
		if errors.As(err, &empty) {
			if has, cerr := f.client.hasCode(ctx, predicted); cerr == nil && has {
				return nil, predicted, ErrSaltInUse
			}
		}
		return nil, predicted, err
	}
	return tx, predicted, nil
}

// ProxyDeployed is the factory's deployment event.
type ProxyDeployed struct {
	Sender         common.Address
	Proxy          common.Address
	Salt           *big.Int
	Implementation common.Address
}

// ParseProxyDeployed finds the ProxyDeployed event in a receipt.
func (f *Factory) ParseProxyDeployed(receipt *types.Receipt) (*ProxyDeployed, error) {
	for _, log := range receipt.Logs {
		if log.Address != f.addr {
			continue
		}
		ev, err := factoryBinding.UnpackProxyDeployedEvent(log)
		if err != nil {
			continue
		}
		return &ProxyDeployed{Sender: ev.Sender, Proxy: ev.ProxyAddress, Salt: ev.Salt, Implementation: ev.Implementation}, nil
	}
	return nil, errors.New("no ProxyDeployed event in receipt")
}

// WaitDeployed waits for a deployment and returns the proxy address.
func (f *Factory) WaitDeployed(ctx context.Context, tx *types.Transaction) (common.Address, error) {
	receipt, err := ens.WaitMinedOK(ctx, f.client.Backend(), tx)
	if err != nil {
		return common.Address{}, err
	}
	ev, err := f.ParseProxyDeployed(receipt)
	if err != nil {
		return common.Address{}, err
	}
	return ev.Proxy, nil
}

// EncodeUserRegistryInit packs initialize(grants) for a UserRegistry proxy.
func EncodeUserRegistryInit(grants []Grant) ([]byte, error) {
	gs := make([]permissionedregistry.Grant, 0, len(grants))
	for _, g := range grants {
		gs = append(gs, permissionedregistry.Grant{Account: g.Account, RoleBitmap: g.Roles.Big()})
	}
	return registryBinding.TryPackInitialize(gs)
}

// DeployRegistryParams configures DeployUserRegistry.
type DeployRegistryParams struct {
	// Grants are the initial root-role holders; empty means
	// [{opts.From, AllRoles}]. The registry rejects an empty list.
	Grants []Grant
	// Salt is the CREATE2 salt; nil means RandomSalt().
	Salt *big.Int
}

// DeployUserRegistry deploys a UserRegistry proxy and returns the
// transaction and the predicted address.
func (c *Client) DeployUserRegistry(ctx context.Context, opts *bind.TransactOpts, p DeployRegistryParams) (*types.Transaction, common.Address, error) {
	if c.cfg.V2.UserRegistryImpl == (common.Address{}) {
		return nil, common.Address{}, ens.ErrNotConfigured
	}
	grants := p.Grants
	if len(grants) == 0 {
		grants = []Grant{{Account: opts.From, Roles: AllRoles}}
	}
	salt := p.Salt
	if salt == nil {
		var err error
		if salt, err = RandomSalt(); err != nil {
			return nil, common.Address{}, err
		}
	}
	init, err := EncodeUserRegistryInit(grants)
	if err != nil {
		return nil, common.Address{}, err
	}
	return c.Factory().DeployProxy(ctx, opts, c.cfg.V2.UserRegistryImpl, salt, init)
}

// DeployResolverParams configures DeployPermissionedResolver.
type DeployResolverParams struct {
	// Grants are the initial root-role holders; empty means
	// [{opts.From, AllRoles}].
	Grants []Grant
	// Salt is the CREATE2 salt; nil means RandomSalt().
	Salt *big.Int
	// InitCalls are setter calls (built with a PermissionedResolver bound to
	// the predicted address) that run inside initialize without role checks.
	InitCalls []ens.Call
}

// DeployPermissionedResolver deploys a PermissionedResolver proxy and
// returns the transaction and the predicted address.
func (c *Client) DeployPermissionedResolver(ctx context.Context, opts *bind.TransactOpts, p DeployResolverParams) (*types.Transaction, common.Address, error) {
	if c.cfg.V2.PermissionedResolverImpl == (common.Address{}) {
		return nil, common.Address{}, ens.ErrNotConfigured
	}
	grants := p.Grants
	if len(grants) == 0 {
		grants = []Grant{{Account: opts.From, Roles: AllRoles}}
	}
	salt := p.Salt
	if salt == nil {
		var err error
		if salt, err = RandomSalt(); err != nil {
			return nil, common.Address{}, err
		}
	}
	calls := make([][]byte, 0, len(p.InitCalls))
	for _, call := range p.InitCalls {
		calls = append(calls, call.Data)
	}
	init, err := EncodePermissionedResolverInit(grants, calls)
	if err != nil {
		return nil, common.Address{}, err
	}
	return c.Factory().DeployProxy(ctx, opts, c.cfg.V2.PermissionedResolverImpl, salt, init)
}
