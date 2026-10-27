# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

This is the continuation of [`wealdtech/go-ens`](https://github.com/wealdtech/go-ens). The last v3 release in this repository is [v3.6.0](https://github.com/ensdomains/go-ens/releases/tag/v3.6.0) (`github.com/wealdtech/go-ens/v3`). v4 lives at `github.com/ensdomains/go-ens/v4`, v5 at `github.com/ensdomains/go-ens/v5`. Step-by-step upgrades are in [MIGRATION.md](MIGRATION.md).

## [5.0.0] - Unreleased

ENS v2 support. The library now speaks the hierarchical v2 registries, the ERC-20 priced `ETHRegistrar`, `PermissionedResolver` and `UniversalResolverV2`, while keeping the v1 surface for names that have not migrated. See [docs/ensv2.md](docs/ensv2.md) for the model and [MIGRATION.md](MIGRATION.md) for the upgrade.

This is a breaking major version relative to [v4.0.0](https://github.com/ensdomains/go-ens/releases/tag/v4.0.0).

### Added

- `ens.Client`: `NewClient(ctx, backend, opts...)` resolves the chain configuration once (`WithChainConfig`, `WithChainID`, `WithUniversalResolver`, `WithGatewayURLs`, `WithHTTPClient`, `WithMaxHops`) and exposes every version-agnostic read: `ResolveAddress`, `MultiAddress`, `HasAddr`, `Text`, `Contenthash`, `ContenthashString`, `PubKey`, `ABI`, `ABIJSON`, `InterfaceImplementer`, `NameRecord`, `Data`, batched `Records` (one `multicall` round trip), `Reverse`, `ReverseResolve`, `ReverseResolveCoin`, `Format`, `FormatCoin`, `FindResolver`, `WritableResolver`, `Owner`, `NearestOwner`, `MigrationStatus`, `IsENSv2`, `Send`, `Simulate`, `ExplainFailedTx`.
- `ens.ChainConfig` with baked-in tables for Ethereum mainnet and Sepolia (`MainnetConfig`, `SepoliaConfig`, `ConfigForChain`), covering the v2 deployment, the v1 contracts still in use and the accepted payment tokens. A test pins every address to the vendored deployment artifacts.
- `ens.Call`: every write in `v1` and `v2` has a `Build*` twin that returns the unsigned call for batching, multisigs or offline signing.
- ENSIP-19 helpers: `CoinTypeDefault`, `EVMCoinType`, `IsEVMCoinType`, `ChainIDFromCoinType`, `ReverseName`; `DNSDecode`.
- `ens.MigrationStatus` reports where a `.eth` name stands in the v1 to v2 migration (`V1Reserved`, `V1ReservedGrace`, `MigratedFromV1`, `NativeV2`, `V2Expired`, `Available`, `Unknown`) including the v1 owner, NameWrapper fuses and whether the name is locked.
- Typed Universal Resolver errors for the v2 custom errors: `DNSDecodingFailedError`, `DNSEncodingFailedError`, `LabelIsEmptyError`, `LabelTooLongError`, `InvalidBatchGatewayResponseError`, `UnsafeBatchGatewayResponseError`, `OffsetOutOfBoundsError`, `NormalizationChangedNameError`, `PrimaryNameNotNormalizedError`, `RegistryRevertError` (`ResolutionError`, a registry reverting during a tree walk), `InvalidOffchainLookupError`; `ResolverRevertError` and `RegistryRevertError` unwrap to the decoded inner error. Sentinels `ErrUnsupportedProfile`, `ErrNoPrimaryName`, `ErrLabelTooLong`, `ErrEmptyLabel`, `ErrNotConfigured`, `ErrNot2LD`, `ErrNotENSv2`, and the registrar sentinels shared by `v1` and `v2`.
- `v2` package: `Client`; `Registry` (reads, writes, roles, safe and unsafe transfers, events) for root, `.eth`, user and wrapper registries; `Registrar` (commit and reveal with a locally computed, contract-verified commitment hash, ERC-20 pricing, client-side pre-flight for the commitment window, availability, allowance and balance); `RenewerV1`, the `Renewer` interface and state-routed `Client.Renew` / `RenewBatch`; `Oracle`; `ERC20` with `CheckPayment`, `ApprovePayment`, `EnsureAllowance`, `FormatPrice`; `PermissionedResolver` (DNS-name keyed setters, `SetRecords`, `Multicall`, record linking, root and scoped roles); `PublicResolverV2`; `Factory` with `DeployPermissionedResolver` and `DeployUserRegistry`; `UniversalHelper`; `ReverseAdapters` and `SetPrimaryNameWithPermissionedResolver`; `PlanCreateSubname` / `ExecutePlan` / `CreateSubname`; a thin `Name` handle; `Role`, `Roles`, `SetterScope` and the registry and resolver role constants; `LabelID`, `CanonicalID`, `WithVersion`, `VersionOf`, `SameName`; `DecodeRevert` with typed errors for every registrar, registry, access-control, resolver, factory and adapter custom error.
- `v1` package: `Client`, `NameWrapper` reads with a `Fuses` type, `DefaultReverseRegistrar`, `NewBaseRegistrarAt`.
- Bindings generated with abigen `--v2` under `contracts/v2`, vendored from `ensdomains/contracts-v2` at a pinned revision by `internal/cmd/vendorabi`, with `make check-generated` in CI.
- An in-process simulated chain (`internal/simchain`) that deploys the real v2 bytecode so the whole write path is tested offline by `go test -short`.
- Live tests against mainnet and Sepolia, a nightly workflow and a cutover watch that records whether the canonical Universal Resolver serves v2. Mainnet went live on v2 on 2026-10-26; the tables track the 2026-10-18 and 2026-10-26 contract replacements.
- `v2.UnemancipatedRoles`, the roles that reach a name; `Registry.CheckTransfer` applies the relaxed safe-transfer rule of registries deployed after 2026-10-26 and also reports the strict one the mainnet root and `.eth` registries keep.

### Changed

- **Module path:** `github.com/ensdomains/go-ens/v4` → `github.com/ensdomains/go-ens/v5`.
- **Package layout mirrors ensjs:** the root package is version-agnostic; the v1-era name management (`Name`, `Registry`, `BaseRegistrar`, `ETHController`, `Resolver`, `ReverseRegistrar`, `ReverseResolver`, auction, deed and DNS helpers, `DeriveTokenID`) moved unchanged to the `v1` package, and bindings moved from `contracts/<pkg>` to `contracts/v1/<pkg>`.
- `Resolve`, `ReverseResolve` and `Format` remain as package-level shortcuts and construct a throwaway `Client`.
- `ReverseResolve` returns `ErrNoPrimaryName` (message unchanged, `"no resolution"`) instead of an ad-hoc error.
- `DNSEncode` returns `*LabelTooLongError` for labels over 255 bytes instead of substituting a bracketed labelhash, which `UniversalResolverV2` cannot decode.
- `NewUniversalResolver` keeps its signature; `Parameter` is now an alias of `Option`, and `WithAddress` also skips the chain-ID probe. The binding is regenerated from `UniversalResolverV2`, adding `FindResolver`, `IsENSv2`, `ResolveDNS` and gateway overrides.
- `UnknownChainError.ChainID` is a `uint64`.
- `v1.NewETHController` takes the controller address from the chain configuration instead of the `.eth` resolver's interface-implementer record, which stops working once the ENS DAO repoints that resolver at the v2 mirror. `v1.NewName` no longer calls `IsValid`.
- `v1.Name` and `v1.ReverseResolver` are deprecated: v1 registration and renewal stop working at migration phase 3, and direct `name(bytes32)` reads fail against a `PermissionedResolver`.
- `golangci-lint` and `abigen` run under the Go 1.25 toolchain (`GOTOOLCHAIN=go1.25.0`) because neither builds on Go 1.27 yet; `make lint` and `make generate` set it.

### Removed

- `github.com/pkg/errors` dependency.
- `contracts/universalresolver` (the two-function v1 ABI); use `contracts/v2/universalresolver`.
- The `ur.integration-tests.eth` wildcard fixtures that depended on the legacy resolver's caller address.

### Security

- Reverts from every v2 write are decoded into typed errors before gas estimation, and `Registrar.Register` refuses to send when the commitment window, availability, allowance or balance would make the transaction fail.

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

[5.0.0]: https://github.com/ensdomains/go-ens/compare/v4.0.0...v5.0.0
[4.0.0]: https://github.com/ensdomains/go-ens/compare/v3.6.0...v4.0.0
