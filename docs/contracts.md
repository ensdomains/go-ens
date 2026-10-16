# Contract bindings

## Sources

| Tree | Source | Generator |
| --- | --- | --- |
| `contracts/v1/*` | ENS v1 ABIs carried over from earlier releases, plus `namewrapper` and `defaultreverseregistrar` from `ensdomains/ens-contracts` at the submodule revision contracts-v2 pins | abigen v1 bindings (`-type Contract`) |
| `contracts/v2/*` | `ensdomains/contracts-v2` deployment artifacts (`contracts/deployments/mainnet/*.json`), pinned by commit in `contracts/v2/manifest.json` | abigen `--v2` bindings |
| `contracts/v2/erc20` | hand-written minimal ERC-20 ABI (plus the mock tokens' `mint` and the OpenZeppelin v5 insufficiency errors) | abigen `--v2` |

Mainnet and Sepolia ship byte-identical ABIs; the vendoring tool fails if they ever differ.

`contracts/v2/deployments/{mainnet,sepolia}.json` record every artifact's address and decoded constructor arguments. `chains_test.go` pins the `ens.ChainConfig` tables to them, and cross-checks the v1 addresses against the constructor arguments the v2 contracts were deployed with, so the tables cannot drift from the deployment.

| Package | Artifact | Note |
| --- | --- | --- |
| `universalresolver` | `UniversalResolverV2.json` | 28 functions, 17 errors |
| `universalhelper` | `UniversalHelper.json` | |
| `permissionedregistry` | `UserRegistryImpl.json` | strict superset of the `ETHRegistry` and `RootRegistry` ABIs; one binding serves every registry |
| `wrapperregistry` | `WrapperRegistryImpl.json` | vendored for the v5.1 migration flows |
| `ethregistrar`, `ethrenewerv1`, `rentpriceoracle` | respective artifacts | |
| `permissionedresolver` | `PermissionedResolverImpl.json` | |
| `publicresolver` | `PublicResolverV2.json` | the only overloaded ABI: `addr`/`addr0`, `setAddr`/`setAddr0`; entry order is significant |
| `verifiablefactory`, `reverseregistraradapter`, `defaultreverseregistraradapter`, `labelstore` | respective artifacts | |
| `compositeresolver` | `ENSV1Resolver.json` | used only for `getResolver(bytes)` unwrapping |

## Regenerating

```sh
make generate            # fetches from GitHub at the pinned revision
make generate VENDOR_LOCAL=/path/to/contracts-v2/contracts/deployments VENDOR_LOCAL_V1=/path/with/NameWrapper.json
make check-generated     # what CI runs: re-vendor into a temp dir and diff
```

`internal/cmd/vendorabi` is the vendoring tool. To move to a new contracts-v2 revision, update `ContractsV2Ref` (and `ENSContractsRef` if the ens-contracts submodule moved) in `internal/cmd/vendorabi/main.go`, run `make generate`, then fix whatever the drift test and the compiler report.

abigen is run as `go run github.com/ethereum/go-ethereum/cmd/abigen@v1.17.5`, matching the go-ethereum version in `go.mod`, so nothing has to be installed. It does not build on Go 1.27 yet, hence `GOTOOLCHAIN=go1.25.0` in the Makefile.

## Simulated chain

`internal/simchain` deploys the creation bytecode vendored into `internal/simchain/testdata/*.bin` onto an `ethclient/simulated` backend: the contract namer, label store, root and `.eth` registries, the mainnet oracle parameters priced in a mock USDC, the registrar, `UniversalResolverV2` with a stub gateway provider, the helper, the factory and the two proxy implementations. The `.eth` registry is emancipated as on mainnet. `v2/sim_test.go` runs the full write path against it under `go test -short`.

## Sepolia

`deployments/sepolia` is the current Sepolia namespace in contracts-v2; the dated directories beside it are immutable archives of earlier deployments. Sepolia already serves `UniversalResolverV2` through the canonical proxy and has v2 registration enabled, so the Sepolia tests exercise the post-cutover path. Its oracle accepts the mock USDC and DAI (which expose `mint`) and Circle's Sepolia USDC. The Sepolia `ETHRegistrarController` address in `ens.SepoliaConfig` is the wrapped controller `ETHRenewerV1` was deployed against, which differs from the one ensjs lists.

Live Sepolia fixtures (from contracts-v2 `deployments/sepolia/fixtures.md`): `bnmig-0601-fe-address-001-r04.eth` is an unwrapped v1 name reserved on v2, `bnmig-0665-fe-perm-001-r04.eth` a wrapped, locked one. `TestSepoliaWritePath` registers a throwaway name when `GO_ENS_TEST_SEPOLIA_KEY` holds a funded key.
