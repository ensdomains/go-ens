# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

This is the continuation of [`wealdtech/go-ens`](https://github.com/wealdtech/go-ens). The last v3 release in this repository is [v3.6.0](https://github.com/ensdomains/go-ens/releases/tag/v3.6.0) (`github.com/wealdtech/go-ens/v3`). v4 lives at `github.com/ensdomains/go-ens/v4`. A step-by-step upgrade is in [MIGRATION.md](MIGRATION.md).

## [4.0.0] - 2026-08-17

First release of the ENS-maintained Go library. Package-level **reads** go through the [Universal Resolver](https://docs.ens.domains/web/ensv2-readiness/), so ENSIP-10 wildcards and ERC-3668 CCIP-Read work. Write and management APIs are unchanged.

This is a breaking major version relative to [v3.6.0](https://github.com/ensdomains/go-ens/releases/tag/v3.6.0).

### Added

- Universal Resolver client (`NewUniversalResolver`) bound to the canonical proxy `0xeEeEEEeE14D718C2B47D9923Deab1335E144EeEe` on Ethereum mainnet and Sepolia.
- Functional options: `WithAddress`, `WithHTTPClient`, `WithMaxHops`.
- `UniversalResolver` read helpers: `Resolve`, `ResolveAddress`, `Reverse`, `Text`, `Contenthash`, `MultiAddress`.
- `ccipread` package: ERC-3668 client that follows `OffchainLookup` reverts (GET `{data}`/`{sender}` or POST JSON), with a 4-hop cap aligned with viem/ethers.
- Typed Universal Resolver errors: `ResolverNotFoundError`, `ResolverNotContractError`, `UnsupportedResolverProfileError`, `ResolverRevertError`, `EmptyAddressError`, `HTTPGatewayError`, `ReverseAddressMismatchError`, plus sentinels `ErrUnregistered` and `ErrNoAddress`.
- `UnknownChainError` when the default constructor is used on a chain without a known UR deployment.
- `DNSEncode` for the DNS wire format the Universal Resolver expects (ENSIP-10).
- ENSIP-15 normalisation via `github.com/adraffy/go-ens-normalize` (Unicode 17.0.0).
- `CoinTypeETH` (`60`) constant.

### Changed

- **Module path:** `github.com/wealdtech/go-ens/v3` → `github.com/ensdomains/go-ens/v4`.
- **`context.Context` is required** on `Resolve`, `ReverseResolve`, `Format`, `PublicResolverAddress`, `Name.Address`, and `DeriveTokenID`. Cancellation and deadlines flow through CCIP-Read hops.
- `Resolve` / `ReverseResolve` / `Format` / `Name.Address` resolve through the Universal Resolver instead of walking the registry.
- `Normalize`, `NormaliseDomain`, `NameHash`, `LabelHash`, and `DNSEncode` use ENSIP-15. Empty labels, the `XX--` label-extension prefix, mid-label underscores, and whole-script confusables now error. Arabic-Indic digits are no longer folded to ASCII, so some namehashes differ from v3.6.0.
- `DNSEncode` allows labels up to 255 bytes; longer labels are replaced with `[hex(keccak256(label))]`. There is no 255-octet cap on the total encoded name (ENSIP-10). A single trailing dot is trimmed before normalisation.
- Reverse resolution now goes through `UniversalResolver.reverse`, which checks that the primary name forward-resolves back to the queried address. A reverse record that fails that check returns `ReverseAddressMismatchError` instead of the name.
- Minimum Go version is 1.25.0 (required by patched `golang.org/x/crypto` / `golang.org/x/net`).
- `github.com/ethereum/go-ethereum` is 1.17.5 (was 1.14.12 in v3.6.0).

### Fixed

- `Resolver.SetABI` reports an error if closing the zlib writer fails.

### Security

- CCIP-Read gateway URLs (untrusted contract input) are restricted to `http`/`https`.
- Gateway redirects are not followed by default, to block SSRF via `Location`.
- Gateway response bodies are capped at 16 MiB and decoded only when `Content-Type` is JSON.
- Bump `golang.org/x/crypto` to 0.55.0, `golang.org/x/net` to 0.58.0, and `github.com/consensys/gnark-crypto` to 0.18.1.

### Removed

- Travis CI (`.travis.yml`). Tests and lint run on GitHub Actions (`main`).

[4.0.0]: https://github.com/ensdomains/go-ens/compare/v3.6.0...v4.0.0
