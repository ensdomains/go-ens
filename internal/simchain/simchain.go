// Copyright 2026 Weald Technology Trading.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0

// Package simchain deploys the ENS v2 contracts from their vendored creation
// bytecode onto an in-process simulated chain, so the write paths can be
// tested without a network. It mirrors the devnet setup of contracts-v2:
// a root registry holding "eth", the .eth registry, the registrar with the
// mainnet oracle parameters priced in a mock USDC, UniversalResolverV2, the
// UniversalHelper, the VerifiableFactory and the user-registry and
// permissioned-resolver implementations.
package simchain

import (
	"context"
	"crypto/ecdsa"
	"embed"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
	"testing"
	"time"

	ens "github.com/ensdomains/go-ens/v5"
	"github.com/ensdomains/go-ens/v5/contracts/v2/ethregistrar"
	"github.com/ensdomains/go-ens/v5/contracts/v2/labelstore"
	"github.com/ensdomains/go-ens/v5/contracts/v2/permissionedregistry"
	"github.com/ensdomains/go-ens/v5/contracts/v2/permissionedresolver"
	"github.com/ensdomains/go-ens/v5/contracts/v2/rentpriceoracle"
	"github.com/ensdomains/go-ens/v5/contracts/v2/universalhelper"
	"github.com/ensdomains/go-ens/v5/contracts/v2/universalresolver"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	bindv2 "github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/require"
)

//go:embed testdata/*.bin
var bytecodes embed.FS

// ChainID of the simulated chain.
const ChainID = 1337

// Registrar parameters, matching mainnet.
const (
	GracePeriod         = 28 * 24 * time.Hour
	MinCommitmentAge    = 60 * time.Second
	MaxCommitmentAge    = 24 * time.Hour
	MinRegisterDuration = 28 * 24 * time.Hour
)

// Account is a funded key on the simulated chain.
type Account struct {
	Key     *ecdsa.PrivateKey
	Address common.Address
}

// Opts returns transact options for the account.
func (a Account) Opts(t *testing.T) *bind.TransactOpts {
	t.Helper()
	opts, err := bind.NewKeyedTransactorWithChainID(a.Key, big.NewInt(ChainID))
	require.NoError(t, err)
	return opts
}

// Chain is a running simulated deployment.
type Chain struct {
	Backend *simulated.Backend
	Client  simulated.Client
	ENS     *ens.Client
	Config  ens.ChainConfig
	// Deployer holds every root role on both registries and owns the
	// registrar; Alice and Bob are funded users.
	Deployer, Alice, Bob Account
	USDC                 common.Address
	ContractNamer        common.Address
}

// New deploys the v2 set and returns the chain. The backend is closed when
// the test ends.
func New(t *testing.T) *Chain {
	t.Helper()
	deployer := newAccount(t)
	alice := newAccount(t)
	bob := newAccount(t)
	alloc := types.GenesisAlloc{}
	for _, a := range []Account{deployer, alice, bob} {
		alloc[a.Address] = types.Account{Balance: new(big.Int).Mul(big.NewInt(1000), big.NewInt(params.Ether))}
	}
	backend := simulated.NewBackend(alloc, simulated.WithBlockGasLimit(60_000_000))
	t.Cleanup(func() { _ = backend.Close() })
	c := &Chain{Backend: backend, Client: backend.Client(), Deployer: deployer, Alice: alice, Bob: bob}
	c.deploy(t)
	return c
}

func newAccount(t *testing.T) Account {
	t.Helper()
	key, err := crypto.GenerateKey()
	require.NoError(t, err)
	return Account{Key: key, Address: crypto.PubkeyToAddress(key.PublicKey)}
}

func (c *Chain) code(t *testing.T, name string) []byte {
	t.Helper()
	raw, err := bytecodes.ReadFile("testdata/" + name + ".bin")
	require.NoError(t, err, "bytecode %s", name)
	code, err := hex.DecodeString(strings.TrimSpace(strings.TrimPrefix(string(raw), "0x")))
	require.NoError(t, err)
	return code
}

// Mine commits a block.
func (c *Chain) Mine() { c.Backend.Commit() }

// Wait mines a block and returns the receipt of tx, failing on a reverted
// transaction.
func (c *Chain) Wait(t *testing.T, tx *types.Transaction) *types.Receipt {
	t.Helper()
	c.Mine()
	receipt, err := c.Client.TransactionReceipt(context.Background(), tx.Hash())
	require.NoError(t, err)
	require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status, "transaction %s reverted", tx.Hash().Hex())
	return receipt
}

// WaitFunc adapts Wait to the v2 package's WaitFunc.
func (c *Chain) WaitFunc(t *testing.T) func(context.Context, *types.Transaction) (*types.Receipt, error) {
	t.Helper()
	return func(_ context.Context, tx *types.Transaction) (*types.Receipt, error) {
		c.Mine()
		receipt, err := c.Client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			return nil, err
		}
		if receipt.Status != types.ReceiptStatusSuccessful {
			return receipt, fmt.Errorf("transaction %s reverted", tx.Hash().Hex())
		}
		return receipt, nil
	}
}

// AdvanceTime moves the chain clock forward and mines a block.
func (c *Chain) AdvanceTime(t *testing.T, d time.Duration) {
	t.Helper()
	require.NoError(t, c.Backend.AdjustTime(d))
	c.Mine()
}

// Now returns the latest block timestamp.
func (c *Chain) Now(t *testing.T) time.Time {
	t.Helper()
	h, err := c.Client.HeaderByNumber(context.Background(), nil)
	require.NoError(t, err)
	return time.Unix(int64(h.Time), 0).UTC()
}

func (c *Chain) deployRaw(t *testing.T, opts *bind.TransactOpts, name string, ctorArgs []byte) common.Address {
	t.Helper()
	addr, tx, err := bindv2.DeployContract(opts, c.code(t, name), c.Client, ctorArgs)
	require.NoError(t, err, "deploy %s", name)
	c.Wait(t, tx)
	code, err := c.Client.CodeAt(context.Background(), addr, nil)
	require.NoError(t, err)
	require.NotEmpty(t, code, "%s has no code after deployment", name)
	return addr
}

// emptyStringArrayInitCode deploys 15 bytes of runtime that return
// abi.encode(string[]{}) (a 0x20 offset word and a zero length word) for any
// calldata: CODECOPY the runtime to memory and RETURN it.
var emptyStringArrayInitCode = common.FromHex("600f600c600039600f6000f3" + "6020600052600060205260406000f3")

func (c *Chain) deployRawCode(t *testing.T, opts *bind.TransactOpts, initCode []byte) common.Address {
	t.Helper()
	addr, tx, err := bindv2.DeployContract(opts, initCode, c.Client, nil)
	require.NoError(t, err)
	c.Wait(t, tx)
	return addr
}

func (c *Chain) send(t *testing.T, opts *bind.TransactOpts, to common.Address, data []byte) {
	t.Helper()
	contract := bind.NewBoundContract(to, abi.ABI{}, c.Client, c.Client, c.Client)
	tx, err := contract.RawTransact(opts, data)
	require.NoError(t, err)
	c.Wait(t, tx)
}

func mustArgs(t *testing.T, types []string, values ...any) []byte {
	t.Helper()
	args := make(abi.Arguments, 0, len(types))
	for _, ty := range types {
		typ, err := abi.NewType(ty, "", nil)
		require.NoError(t, err)
		args = append(args, abi.Argument{Type: typ})
	}
	out, err := args.Pack(values...)
	require.NoError(t, err)
	return out
}

func (c *Chain) deploy(t *testing.T) {
	t.Helper()
	opts := c.Deployer.Opts(t)
	deployer := c.Deployer.Address
	ctx := context.Background()

	// ContractNamer: implementation + ERC-1967 proxy initialised to the deployer.
	namerImpl := c.deployRaw(t, opts, "ContractNamer_Implementation", nil)
	initSel := ens.Selector("initialize(address)")
	initData := append(initSel[:], common.LeftPadBytes(deployer.Bytes(), 32)...)
	c.ContractNamer = c.deployRaw(t, opts, "ContractNamer_Proxy", mustArgs(t, []string{"address", "bytes"}, namerImpl, initData))

	// LabelStore.
	labelStore := c.deployRaw(t, opts, "LabelStore", labelstore.NewLabelStore().PackConstructor(c.ContractNamer))

	// Registries: ETHRegistry bytecode is a PermissionedRegistry whose
	// constructor is (labelStore, rootAccount, roleBitmap).
	allRoles, _ := new(big.Int).SetString(strings.Repeat("1", 64), 16)
	regArgs := mustArgs(t, []string{"address", "address", "uint256"}, labelStore, deployer, allRoles)
	root := c.deployRaw(t, opts, "ETHRegistry", regArgs)
	eth := c.deployRaw(t, opts, "ETHRegistry", regArgs)

	// Mock USDC with mint().
	c.USDC = c.deployRaw(t, opts, "MockUSDC", mustArgs(t, []string{"string", "uint8"}, "USDC", uint8(6)))

	// Oracle with the mainnet parameters, priced in the mock USDC.
	oracleB := rentpriceoracle.NewStandardRentPriceOracle()
	denom, _ := new(big.Int).SetString("100000000000000000000000000000000000000", 10)
	num := func(s string) *big.Int { v, _ := new(big.Int).SetString(s, 10); return v }
	oracleCtor := oracleB.PackConstructor(
		deployer,
		[]*big.Int{big.NewInt(0), big.NewInt(0), big.NewInt(20294267), big.NewInt(5073567), big.NewInt(253679)},
		[]rentpriceoracle.DiscountPoint{
			{Duration: 63072000, Numer: num("87500000000000000000000000000000000000")},
			{Duration: 94608000, Numer: num("68750000000000000000000000000000000000")},
			{Duration: 189216000, Numer: num("56250000000000000000000000000000000000")},
		},
		denom,
		num("100000000000000000000"), 86400, 1814400,
		[]rentpriceoracle.PaymentRatio{{PaymentToken: c.USDC, Numer: big.NewInt(1), Denom: big.NewInt(1000000)}},
	)
	oracle := c.deployRaw(t, opts, "StandardRentPriceOracle", oracleCtor)

	// Registrar.
	regB := ethregistrar.NewETHRegistrar()
	registrarCtor := regB.PackConstructor(deployer, eth, deployer, oracle,
		uint64(GracePeriod/time.Second), uint64(MinCommitmentAge/time.Second), uint64(MaxCommitmentAge/time.Second), uint64(MinRegisterDuration/time.Second))
	registrar := c.deployRaw(t, opts, "ETHRegistrar", registrarCtor)

	// Resolution and helpers. The batch gateway provider is a stub whose
	// every call returns an empty string[], so resolvers that support
	// direct calls (ERC-7996) are used in-process and nothing goes offchain.
	gatewayProvider := c.deployRawCode(t, opts, emptyStringArrayInitCode)
	urCtor := universalresolver.NewUniversalResolverV2().PackConstructor(root, gatewayProvider, c.ContractNamer)
	ur := c.deployRaw(t, opts, "UniversalResolverV2", urCtor)
	helperCtor := universalhelper.NewUniversalHelper().PackConstructor(root, c.ContractNamer)
	helper := c.deployRaw(t, opts, "UniversalHelper", helperCtor)
	factory := c.deployRaw(t, opts, "VerifiableFactory", nil)
	userImplCtor := permissionedregistry.NewPermissionedRegistry().PackConstructor(labelStore, common.Address{}, c.ContractNamer)
	userImpl := c.deployRaw(t, opts, "UserRegistryImpl", userImplCtor)
	resolverImplCtor := permissionedresolver.NewPermissionedResolver().PackConstructor(c.ContractNamer)
	resolverImpl := c.deployRaw(t, opts, "PermissionedResolverImpl", resolverImplCtor)

	// Wire the tree: root holds "eth" -> ETHRegistry; the registrar may
	// register and renew on the .eth registry.
	regBind := permissionedregistry.NewPermissionedRegistry()
	c.send(t, opts, root, regBind.PackRegister("eth", deployer, eth, common.Address{}, allRoles, ^uint64(0)))
	c.send(t, opts, eth, regBind.PackSetParent(root, "eth"))
	registrarRoles := new(big.Int).Or(new(big.Int).Lsh(big.NewInt(1), 0), new(big.Int).Lsh(big.NewInt(1), 16))
	c.send(t, opts, eth, regBind.PackGrantRootRoles(registrarRoles, registrar))
	// Emancipate the .eth registry as mainnet does: the deployer drops the
	// root roles that could reach a registered name (unregister, set
	// subregistry, set resolver, transfer, upgrade) and their admin roles.
	dangerous := new(big.Int)
	for _, nybble := range []uint{12, 20, 24, 124} {
		dangerous.Or(dangerous, new(big.Int).Lsh(big.NewInt(1), nybble))
		dangerous.Or(dangerous, new(big.Int).Lsh(big.NewInt(1), nybble+128))
	}
	dangerous.Or(dangerous, new(big.Int).Lsh(big.NewInt(1), 28+128))
	c.send(t, opts, eth, regBind.PackRevokeRootRoles(dangerous, deployer))

	c.Config = ens.ChainConfig{
		ChainID: ChainID,
		Name:    "simulated",
		V2: ens.V2Contracts{
			UniversalResolver:        ur,
			UniversalResolverV2:      ur,
			UniversalHelper:          helper,
			RootRegistry:             root,
			ETHRegistry:              eth,
			ETHRegistrar:             registrar,
			RentPriceOracle:          oracle,
			VerifiableFactory:        factory,
			UserRegistryImpl:         userImpl,
			PermissionedResolverImpl: resolverImpl,
			LabelStore:               labelStore,
		},
		PaymentTokens:       []ens.PaymentToken{{Symbol: "USDC", Address: c.USDC, Decimals: 6}},
		DefaultPaymentToken: c.USDC,
	}
	client, err := ens.NewClient(ctx, c.Client, ens.WithChainConfig(c.Config))
	require.NoError(t, err)
	c.ENS = client
}

// MintUSDC gives an account mock USDC.
func (c *Chain) MintUSDC(t *testing.T, to common.Address, amount *big.Int) {
	t.Helper()
	sel := ens.Selector("mint(address,uint256)")
	data := append(sel[:], mustArgs(t, []string{"address", "uint256"}, to, amount)...)
	c.send(t, c.Deployer.Opts(t), c.USDC, data)
}
