# Migration guides

- [v4 to v5](#migrating-from-v4-to-v5): ENS v2 support, `Client`, `v1`/`v2` packages.
- [v3 to v4](#migrating-from-v3-to-v4): Universal Resolver reads, `context.Context`, ENSIP-15.

# Migrating from v4 to v5

This guide covers upgrading from [v4.0.0](https://github.com/ensdomains/go-ens/releases/tag/v4.0.0) (`github.com/ensdomains/go-ens/v4`) to `github.com/ensdomains/go-ens/v5`.

v5 is a breaking major version. It adds full [ENS v2](docs/ensv2.md) support and moves the v1-era name management into a `v1` package. Nothing is deleted: every v4 symbol is either kept at the root, moved to `v1`, or has a v2 replacement listed below.

## 1. Change the module path

```sh
go get github.com/ensdomains/go-ens/v5
```

```go
// v4
import ens "github.com/ensdomains/go-ens/v4"

// v5
import (
    ens "github.com/ensdomains/go-ens/v5"
    v1 "github.com/ensdomains/go-ens/v5/v1" // only if you use v1 name management
    v2 "github.com/ensdomains/go-ens/v5/v2" // v2 registries, registrar, resolvers
)
```

Go 1.25.0 or later is required. `golangci-lint` and `abigen` do not build on Go 1.27 yet; run them with `GOTOOLCHAIN=go1.25.0` (the `Makefile` does).

## 2. Hold a `Client`

The package-level shortcuts `ens.Resolve`, `ens.ReverseResolve` and `ens.Format` still work and still take `(ctx, backend, ...)`. Each one now builds a throwaway client, which costs one `eth_chainId` call, so servers should hold a `Client`:

```go
c, err := ens.NewClient(ctx, backend)          // probes the chain ID, picks mainnet or Sepolia
addr, err := c.ResolveAddress(ctx, "vitalik.eth")
name, err := c.ReverseResolve(ctx, addr)
recs, err := c.Records(ctx, "vitalik.eth", ens.RecordsRequest{Texts: []string{"url"}, Coins: []uint64{60}})
```

Options: `WithChainConfig` and `WithChainID` skip the probe, `WithUniversalResolver` overrides the resolver address, `WithGatewayURLs`, `WithHTTPClient` and `WithMaxHops` tune CCIP-Read.

## 3. Where every v4 symbol went

| v4 (root) | v5 | Notes |
| --- | --- | --- |
| `Resolve`, `ReverseResolve`, `Format` | root, unchanged signatures | `ReverseResolve` returns `ErrNoPrimaryName` (same message) |
| `NewUniversalResolver`, `UniversalResolver` and its methods | root | binding regenerated from `UniversalResolverV2`; new `FindResolver`, `IsENSv2`, `ResolveDNS`, `Records`, `PubKey`, `ABI`, `InterfaceImplementer`, `NameRecord`, `Data`, `HasAddr` |
| `Parameter`, `WithAddress`, `WithHTTPClient`, `WithMaxHops` | root | `Parameter` is an alias of `Option`; `WithAddress` also skips the chain-ID probe |
| `UniversalResolverContractAddress`, `CoinTypeETH`, `UnknownAddress` | root | plus `CoinTypeDefault`, `ChainIDMainnet`, `ChainIDSepolia` |
| `UnknownChainError` | root | `ChainID` is now `uint64` |
| `ErrUnregistered`, `ErrNoAddress`, the seven typed UR errors | root | nine new typed errors; `UnsupportedResolverProfileError` unwraps to `ErrUnsupportedProfile` |
| `Normalize`, `NameHash`, `LabelHash`, `DNSEncode`, `NormaliseDomain*`, `DomainLevel`, `Tld`, `Domain`, `DomainPart`, `UnqualifiedName` | root | `DNSEncode` errors on labels over 255 bytes (`ErrLabelTooLong`) instead of substituting `[hash]` |
| `StringToContenthash`, `ContenthashToString` | root | |
| `ccipread.*`, `util.*` | unchanged | |
| `Name`, `NewName` and methods | `v1.Name` (deprecated) | v2: `v2.Name`, `v2.Registrar`, `v2.Client.Renew`, `v2.Registry.GrantRoles`, `v2.Client.CreateSubname` |
| `Registry`, `NewRegistry`, `NewRegistryAt`, `RegistryContractAddress`, `RegistryContractFromRegistrar`, `SetResolver`, `SetSubdomainOwner`, `CreateRegistrySession` | `v1.*` | v2: `v2.Registry` |
| `BaseRegistrar`, `NewBaseRegistrar` and methods | `v1.*` (+ `NewBaseRegistrarAt`) | reads keep working for unmigrated names |
| `ETHController`, `NewETHController`, `NewETHControllerAt` and methods | `v1.*` | controller address now comes from `ChainConfig`; v2: `v2.Registrar` |
| `Resolver`, `NewResolver`, `NewResolverAt`, `PublicResolverAddress` and methods | `v1.*` | reads: `ens.Client`; writes for v2 names: `v2.PermissionedResolver` or `v2.PublicResolverV2` |
| `ReverseRegistrar`, `NewReverseRegistrar`, `NewReverseRegistrarAt`, `SetName`, `DefaultResolverAddress` | `v1.*` | still the write path for primary names; new `v1.DefaultReverseRegistrar`, `v2.ReverseAdapters` |
| `ReverseResolver`, `NewReverseResolver*`, `Name` | `v1.*` (deprecated) | use `ens.Client.ReverseResolve` |
| `AuctionRegistrar`, `Deed`, `DNSRegistrar`, `DNSResolver`, `DNSSECOracle`, `DNSWireFormat*`, `RegistrarContractAddress` | `v1.*` | v1-only; DNS names resolve through `ens.Client` |
| `DeriveTokenID` | `v1.DeriveTokenID` | v1 ERC-721 id; v2: `v2.LabelID`, `v2.Registry.TokenID` |
| `contracts/<pkg>` | `contracts/v1/<pkg>` | `contracts/universalresolver` is replaced by `contracts/v2/universalresolver` |

## 4. Behaviour changes to check

- The resolver address returned by `Resolve` and `Reverse` is the contract the Universal Resolver called. Since the v2 cutover that is a read-only mirror (`ENSV1Resolver`) for names still held in v1. Use `Client.WritableResolver` to find the contract that holds the records and its family (`node-keyed`, `permissioned`, `mirror`).
- `resolve("eth")` has no resolver on v2 (`ErrUnregistered`).
- `PubKey`, DNS records and `zonehash` are not implemented by the v2 `PermissionedResolver`; reads surface `ErrUnsupportedProfile`.
- v2 registration and renewal are paid in ERC-20 (USDC or DAI on mainnet), never ETH. `v2.Registrar.Register` refuses an `opts.Value`.
- Token ids regenerate on every role change. Key everything by label; fetch `Registry.TokenID` right before a transfer.

## 5. Migrating writes

```go
v2c := v2.NewClient(c)

// True on mainnet since the 2026-10-26 go-live; false on forks that skipped phase 6.
open, err := v2c.RegistrationOpen(ctx)

// Deploy a resolver with initial records, then commit and register.
tx, resolver, err := v2c.DeployPermissionedResolver(ctx, opts, v2.DeployResolverParams{})
params := v2.RegistrationParams{Commitment: v2.Commitment{Label: "example", Owner: opts.From, Secret: secret, Resolver: resolver, Duration: 365 * 24 * time.Hour}}
tx, hash, err := v2c.ETHRegistrar().Commit(opts, params.Commitment)
err = v2c.ETHRegistrar().WaitForCommitment(ctx, hash)
price, err := v2c.ETHRegistrar().RegisterPrice(ctx, "example", params.Duration, common.Address{})
tx, err = v2c.EnsureAllowance(ctx, opts, price.Token, v2c.Config().V2.ETHRegistrar, price.Total())
tx, err = v2c.ETHRegistrar().Register(ctx, opts, params)

// Renewal routes by state: ETHRegistrar for v2 names, ETHRenewerV1 for unmigrated v1 names.
tx, err = v2c.Renew(ctx, opts, v2.RenewData{Label: "vitalik", Duration: 365 * 24 * time.Hour}, common.Address{})

// Records.
pr := v2c.PermissionedResolver(resolver)
tx, err = pr.SetRecords(opts, "example.eth", v2.RecordSet{Texts: map[string]string{"url": "https://example.org"}})

// Primary name (an EOA can keep using v1).
rr, err := v1.NewClient(c).ReverseRegistrar()
tx, err = rr.SetName(opts, "example.eth")
```

Every write has a `Build*` twin returning an `ens.Call` for multisigs, batching or offline signing; `Client.Send` submits one.

## 6. Checklist

1. `go get github.com/ensdomains/go-ens/v5` and rewrite imports; add `v1` or `v2` imports where name management is used.
2. Replace per-call `ens.Resolve` in servers with a long-lived `ens.Client`.
3. Switch v1 registration and renewal to `v2.Registrar` and `v2.Client.Renew`; budget ERC-20 approvals.
4. Find writable resolvers with `Client.WritableResolver` and write through `v2.PermissionedResolver` for v2 names.
5. Check `Client.MigrationStatus` before acting on a `.eth` name you do not control.
6. Re-run anything that `DNSEncode`s very long labels; they now error.

# Migrating from v3 to v4

This guide covers upgrading from [v3.6.0](https://github.com/ensdomains/go-ens/releases/tag/v3.6.0) (`github.com/wealdtech/go-ens/v3`) to `github.com/ensdomains/go-ens/v4`.

v4 is a breaking major version. The write/management surface is the same; **name reads** now go through the ENS [Universal Resolver](https://docs.ens.domains/web/ensv2-readiness/), take a `context.Context`, and normalise names with [ENSIP-15](https://docs.ens.domains/ensip/15).

Stay on v3 if you need the old registry-walking resolver lookup and do not want wildcards, CCIP-Read, or ENSIP-15.

## 1. Change the module path

```sh
go get github.com/ensdomains/go-ens/v4
```

```go
// v3
import ens "github.com/wealdtech/go-ens/v3"

// v4
import ens "github.com/ensdomains/go-ens/v4"
```

Go 1.25.0 or later is required.

This module depends on `github.com/ethereum/go-ethereum` v1.17.5. If your project still pins geth 1.14.x, `go get` will pull 1.17; `accounts/abi/bind` v1 still works but is deprecated upstream.

## 2. Pass a context into every public read

These signatures now take `ctx` as the first argument:


| Function                | v3                   | v4                        |
| ----------------------- | -------------------- | ------------------------- |
| `Resolve`               | `(backend, input)`   | `(ctx, backend, input)`   |
| `ReverseResolve`        | `(backend, address)` | `(ctx, backend, address)` |
| `Format`                | `(backend, address)` | `(ctx, backend, address)` |
| `PublicResolverAddress` | `(backend)`          | `(ctx, backend)`          |
| `Name.Address`          | `(coinType)`         | `(ctx, coinType)`         |
| `DeriveTokenID`         | `(backend, domain)`  | `(ctx, backend, domain)`  |


```go
ctx := context.Background()

address, err := ens.Resolve(ctx, client, "ethereum.eth")
name, err := ens.ReverseResolve(ctx, client, address)
fmt.Println(ens.Format(ctx, client, address))

n, err := ens.NewName(client, "mydomain.eth")
addr, err := n.Address(ctx, ens.CoinTypeETH)
```

`ctx` is honoured for cancellation and deadlines, including HTTP hops during CCIP-Read. Prefer a request-scoped deadline over `context.Background()` in servers.

Constructors such as `NewName`, `NewResolver`, and `NewRegistry`, and all write methods (`SetAddress`, `SetText`, registration, …), are unchanged and still do not take a context.

## 3. Reads go through the Universal Resolver

`Resolve`, `ReverseResolve`, `Format`, and `Name.Address` no longer walk the registry to find a resolver and call it. They call the Universal Resolver at:

```
0xeEeEEEeE14D718C2B47D9923Deab1335E144EeEe
```

That proxy is deployed by the ENS DAO on **Ethereum mainnet (1)** and **Sepolia (11155111)**. The result:

- ENSIP-10 wildcard names resolve.
- ERC-3668 CCIP-Read (offchain / L2) names resolve.
- Reverse resolution requires a **primary name**: the reverse record must forward-resolve back to the same address.

### `Resolver` is still the v3 object

`Resolver.Address`, `Resolver.Text`, `Resolver.Contenthash`, and `Resolver.MultiAddress` call the bound resolver contract directly. They do **not** follow wildcards or CCIP-Read.


| Want…                                                          | Use                                                       |
| -------------------------------------------------------------- | --------------------------------------------------------- |
| Address of any name, including wildcard / CCIP-Read            | `ens.Resolve` or `Name.Address`                           |
| Text / contenthash / multicoin addr with wildcards / CCIP-Read | `UniversalResolver.Text` / `Contenthash` / `MultiAddress` |
| Set records, or read a resolver you already hold               | `Resolver`                                                |


```go
// Wildcard- and CCIP-aware (v4 default).
addr, err := ens.Resolve(ctx, client, "foo.example.eth")

// Legacy on-chain read of one resolver. Offchain names will not resolve here.
r, err := ens.NewResolver(client, "foo.example.eth")
onchain, err := r.Address()
```

### Other chains, forks, and custom deployments

`NewUniversalResolver` checks the backend chain ID when you bind to the canonical address. Unknown chains return `*UnknownChainError`. Pass `WithAddress` to skip the guardrail:

```go
ur, err := ens.NewUniversalResolver(ctx, client,
    ens.WithAddress(customUR),
    ens.WithHTTPClient(httpClient), // optional
    ens.WithMaxHops(4),             // optional; default 4
)
addr, err := ur.ResolveAddress(ctx, "foo.eth")
```

### RPC requirements

CCIP-Read needs the JSON-RPC `data` field on revert responses. Some providers strip it; if offchain names fail with a generic revert and no `OffchainLookup`, switch endpoint. The test default (`https://ethereum.publicnode.com`) surfaces revert data. Override in tests with `GO_ENS_TEST_RPC`.

Live-RPC tests skip under `go test -short`.

## 4. Handle the new errors

v3 returned plain `errors.New("unregistered name")` / `"no resolution"` strings. v4 still produces those messages in the common cases, but the concrete types changed. Prefer `errors.Is` / `errors.As`:

```go
addr, err := ens.Resolve(ctx, client, name)
switch {
case errors.Is(err, ens.ErrUnregistered):
    // no resolver anywhere up the name
case errors.Is(err, ens.ErrNoAddress):
    // resolver exists, address record missing or zero
default:
    var mismatch *ens.ReverseAddressMismatchError
    var httpErr *ens.HTTPGatewayError
    var unknown *ens.UnknownChainError
    _ = mismatch
    _ = httpErr
    _ = unknown
}
```


| Situation                                            | v3                                    | v4                                                  |
| ---------------------------------------------------- | ------------------------------------- | --------------------------------------------------- |
| No resolver                                          | `"unregistered name"`                 | `*ResolverNotFoundError` wrapping `ErrUnregistered` |
| Empty / zero address                                 | often zero address or a generic error | `ErrNoAddress` / `*EmptyAddressError`               |
| Reverse name does not forward-resolve to the address | returned the name                     | `*ReverseAddressMismatchError`                      |
| CCIP-Read gateway HTTP error                         | n/a                                   | `*HTTPGatewayError` or `ccipread.GatewayHTTPError`  |
| Resolver reverted                                    | raw RPC error                         | `*ResolverRevertError`                              |
| Resolver has no code                                 | `"no resolver"`                       | `*ResolverNotContractError`                         |
| Profile not implemented                              | raw revert                            | `*UnsupportedResolverProfileError`                  |


`ReverseResolve` still returns `"no resolution"` when the Universal Resolver succeeds with an empty name.

## 5. ENSIP-15 normalisation

`Normalize`, `NormaliseDomain`, `NameHash`, `LabelHash`, and `DNSEncode` use [ENSIP-15](https://docs.ens.domains/ensip/15) (`github.com/adraffy/go-ens-normalize`). v3 used `golang.org/x/net/idna` with `MapForLookup`, which let through names the on-chain Universal Resolver rejects.

These now **error** (they did not in v3):

- Empty labels: `.eth`, `foo..eth`, `foo.eth.` (via `Normalize` / `NameHash`)
- `XX--` label-extension prefix: `te--st.eth`
- Underscores **inside** a label: `a_b.eth` (`_foo.eth` is still valid)
- Whole-script confusables (for example Cyrillic mimicking Latin)

Arabic-Indic digits are **not** folded to ASCII. `۰۱۲۳۷۸۹.eth` is a different namehash than in v3.

`NormaliseDomainStrict` is unchanged: it is still IDNA with `StrictDomainName=true` (DNS-strict, no underscores). It is **not** ENSIP-15. Use it only when publishing to a DNS-backed registrar.

`DNSEncode` (used by the Universal Resolver) trims a **single** trailing dot before ENSIP-15, so DNS-style `"foo.eth."` still encodes. `Normalize("foo.eth.")` itself still errors.

`NormaliseDomain` still preserves a leading `*.` wildcard prefix.

If you persist namehashes, re-hash any names that contain Arabic-Indic digits or that v3 accepted and v4 rejects.

## 6. What did not change

- Registration, renewals, transfers, controller, subdomains.
- `Resolver.SetAddress` / `SetText` / `SetABI` / `SetContenthash` / `SetMultiAddress` / `SetPubKey`.
- `Registry`, `BaseRegistrar`, `ETHController`, `ReverseRegistrar`, auction/DNS registrar helpers.
- `NewName` (still no context).
- Contenthash encode/decode, token ID derivation logic (aside from the new `ctx` and the fact that `DeriveTokenID` now resolves through the UR).

## 7. Checklist

1. `go get github.com/ensdomains/go-ens/v4` and rewrite imports.
2. Thread `context.Context` into the six read functions above.
3. Use `ens.Resolve` / `Name.Address` / `UniversalResolver` for wildcard and CCIP-Read names; keep `Resolver` for writes.
4. Switch string-matched errors to `errors.Is` / `errors.As`.
5. Re-run anything that hashes or validates names; expect ENSIP-15 rejections and a few hash changes.
6. Confirm the RPC surfaces revert `data`, or CCIP-Read names will fail.
