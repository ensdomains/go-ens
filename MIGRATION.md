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

