# go-ens

[![Tag](https://img.shields.io/github/tag/ensdomains/go-ens.svg)](https://github.com/ensdomains/go-ens/releases/)
[![License](https://img.shields.io/github/license/ensdomains/go-ens.svg)](LICENSE)
[![GoDoc](https://pkg.go.dev/badge/github.com/ensdomains/go-ens/v5.svg)](https://pkg.go.dev/github.com/ensdomains/go-ens/v5)
[![CI](https://github.com/ensdomains/go-ens/actions/workflows/ci.yml/badge.svg)](https://github.com/ensdomains/go-ens/actions/workflows/ci.yml)

Go module for the [Ethereum Name Service](https://ens.domains/), with full support for [ENS v2](docs/ensv2.md).

This is the ENS Go library, the continuation of [`wealdtech/go-ens`](https://github.com/wealdtech/go-ens). Import `github.com/ensdomains/go-ens/v5`. Upgrading: [MIGRATION.md](MIGRATION.md). History: [CHANGELOG.md](CHANGELOG.md).

## Layout

| Package | Purpose |
| --- | --- |
| `github.com/ensdomains/go-ens/v5` | Version-agnostic: `Client`, chain configuration, resolution and reverse resolution through the Universal Resolver, record reads, resolver and owner discovery, migration status, normalisation and hashing. Works for every name, v1 or v2. |
| `github.com/ensdomains/go-ens/v5/v2` | ENS v2: hierarchical registries, the ERC-20 priced `.eth` registrar and renewers, `PermissionedResolver` and `PublicResolverV2`, the proxy factory, subnames, roles, reverse adapters. |
| `github.com/ensdomains/go-ens/v5/v1` | ENS v1: the legacy registry, `BaseRegistrar`, controller, `PublicResolver`, reverse registrars and NameWrapper reads, for names that have not migrated. |
| `github.com/ensdomains/go-ens/v5/ccipread` | ERC-3668 CCIP-Read client. |

## Install

```sh
go get github.com/ensdomains/go-ens/v5
```

Go 1.25 or later.

## Resolution

```go
client, err := ethclient.Dial("https://mainnet.infura.io/v3/SECRET")
c, err := ens.NewClient(ctx, client) // picks the mainnet or Sepolia configuration from the chain ID

address, err := c.ResolveAddress(ctx, "vitalik.eth")
name, err := c.ReverseResolve(ctx, address)   // ErrNoPrimaryName when none is set
fmt.Println(c.Format(ctx, address))            // name, or the hex address

text, err := c.Text(ctx, "vitalik.eth", "url")
recs, err := c.Records(ctx, "vitalik.eth", ens.RecordsRequest{ // one round trip
    Texts: []string{"url", "com.twitter"}, Coins: []uint64{ens.CoinTypeETH, ens.EVMCoinType(10)}, Contenthash: true,
})
```

Every read goes through the ENS Universal Resolver, so ENSIP-10 wildcards and ERC-3668 CCIP-Read names resolve like on-chain ones, and names held in the v1 registry resolve through the v2 mirror after the cutover. `ens.Resolve`, `ens.ReverseResolve` and `ens.Format` remain as package-level shortcuts.

Reverse lookups for other EVM chains follow ENSIP-19:

```go
name, err := c.ReverseResolveCoin(ctx, address.Bytes(), ens.EVMCoinType(8453))
```

## Where a name stands

```go
st, err := c.MigrationStatus(ctx, "vitalik.eth")
switch st.State {
case ens.MigrationStateV1Reserved:   // still in the v1 registrar; renews through ETHRenewerV1
case ens.MigrationStateMigratedFromV1, ens.MigrationStateNativeV2: // held by the v2 .eth registry
case ens.MigrationStateAvailable:    // registrable once v2 registration is enabled
}

info, err := c.WritableResolver(ctx, "vitalik.eth") // the contract that holds the records, and its family
owner, err := c.Owner(ctx, "vitalik.eth")            // v2 owner through the UniversalHelper
v2yet, err := c.IsENSv2(ctx)                          // has the canonical resolver been switched to v2?
```

## Registering and managing names (v2)

```go
v2c := v2.NewClient(c)

open, err := v2c.RegistrationOpen(ctx) // true on mainnet since the 2026-10-26 go-live; false on forks without phase 6

// A resolver for the name, with its first records set during deployment.
pr := v2c.PermissionedResolver(predicted) // predicted := Factory().PredictProxyAddress(ctx, opts.From, salt)
setAddr, _ := pr.BuildSetAddress("example.eth", ens.CoinTypeETH, opts.From.Bytes())
tx, resolver, err := v2c.DeployPermissionedResolver(ctx, opts, v2.DeployResolverParams{Salt: salt, InitCalls: []ens.Call{setAddr}})

// Commit, wait, pay in USDC, register.
secret, _ := v2.NewSecret()
params := v2.RegistrationParams{Commitment: v2.Commitment{
    Label: "example", Owner: opts.From, Secret: secret, Resolver: resolver, Duration: 365 * 24 * time.Hour,
}}
tx, hash, err := v2c.ETHRegistrar().Commit(opts, params.Commitment)
err = v2c.ETHRegistrar().WaitForCommitment(ctx, hash)
price, err := v2c.ETHRegistrar().RegisterPrice(ctx, "example", params.Duration, common.Address{})
tx, err = v2c.EnsureAllowance(ctx, opts, price.Token, v2c.Config().V2.ETHRegistrar, price.Total())
tx, err = v2c.ETHRegistrar().Register(ctx, opts, params)
res, err := v2c.ETHRegistrar().WaitRegistered(ctx, tx)

// Renew (routes to ETHRegistrar or, for unmigrated v1 names, ETHRenewerV1).
tx, err = v2c.Renew(ctx, opts, v2.RenewData{Label: "example", Duration: 365 * 24 * time.Hour}, common.Address{})

// Records, roles, subnames, transfer.
tx, err = pr.SetRecords(opts, "example.eth", v2.RecordSet{Texts: map[string]string{"url": "https://example.org"}})
n, err := v2c.Name(ctx, "example.eth")
tx, err = n.GrantRoles(opts, v2.NewRoles(v2.RoleSetResolver), delegate)
receipts, err := v2c.CreateSubname(ctx, opts, v2.SubnameParams{Name: "sub.example.eth", Owner: friend, Resolver: resolver}, nil)
tx, err = n.Transfer(ctx, opts, newOwner)
```

Registration and renewal are paid in ERC-20 (USDC or DAI on mainnet). `Register` checks the commitment window, availability, allowance and balance before sending and returns typed errors (`errors.Is(err, v2.ErrInsufficientAllowance)`, `errors.As(err, &v2.CommitmentTooNewError{})`, …). Every write has a `Build*` twin returning an `ens.Call` for multisigs or offline signing.

Primary names are still stored by the v1 reverse registrars:

```go
rr, err := v1.NewClient(c).ReverseRegistrar()
tx, err := rr.SetName(opts, "example.eth")
```

## Other chains

Mainnet and Sepolia tables are built in. For a fork or devnet pass a full configuration:

```go
c, err := ens.NewClient(ctx, client, ens.WithChainConfig(cfg))
```

To bypass the canonical proxy and call `UniversalResolverV2` directly:

```go
cfg := ens.MainnetConfig()
c, err := ens.NewClient(ctx, client, ens.WithUniversalResolver(cfg.V2.UniversalResolverV2))
```

## Development

```sh
make test        # unit tests plus the in-process simulated chain
make test-live   # read tests against mainnet (GO_ENS_TEST_RPC, public fallback)
make lint        # golangci-lint under the Go 1.25 toolchain
make generate    # re-vendor ABIs from contracts-v2 and regenerate bindings
```

See [docs/contracts.md](docs/contracts.md) for the binding pipeline and [docs/ensv2.md](docs/ensv2.md) for the ENS v2 model as this library sees it.

## Maintainers

[ENS](https://github.com/ensdomains).

Originally written by Jim McDonald ([@mcdee](https://github.com/mcdee)) at Weald Technology.

## Contribute

Contributions welcome. Please check out [the issues](https://github.com/ensdomains/go-ens/issues).

## License

[Apache-2.0](LICENSE) © 2019 Weald Technology Trading Ltd
