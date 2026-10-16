# ENS v2 as go-ens sees it

This page explains the on-chain model the `v2` package wraps and what it means for callers. The authoritative design is the [ENSv2 design doc](http://go.ens.xyz/ensv2) and the [contracts-v2](https://github.com/ensdomains/contracts-v2) repository.

## Registries

ENS v2 replaces the flat v1 registry with one registry contract per name, each responsible for its direct children:

```
RootRegistry  --"eth"-->  ETHRegistry  --"example"-->  example.eth's UserRegistry  --"sub"-->  ...
```

Every registry is an ERC-1155 contract (`PermissionedRegistry`) with one token per child name. For each label it stores a status (`available`, `reserved`, `registered`), an expiry, the latest owner, the child registry ("subregistry") and the resolver. `v2.Registry` wraps all of them; `v2.Client.RootRegistry()` and `ETHRegistry()` are the two well-known ones and `UniversalHelper` walks the tree (`Registries`, `ParentRegistry`, `ExactOwner`).

**Ids.** A name's canonical id is its labelhash with the low 32 bits cleared. The ERC-1155 token id and the access-control resource are that canonical id plus a version in the low 32 bits. The version changes whenever roles on the name change, on unregistration and on re-registration, so a token id you cached can stop resolving to an owner at any time. Key everything by label and fetch `Registry.TokenID` right before a transfer. `v2.LabelID`, `CanonicalID`, `WithVersion`, `VersionOf` and `SameName` implement the arithmetic.

**Roles.** There is no registrant/controller split. The token owner holds roles, one nybble each in a 256-bit bitmap, and delegates them with `GrantRoles`. Each role has an admin role (shifted left by 128 bits) that only the owner can hold. `v2.Roles` is the set type; the constants mirror `RegistryRolesLib`. The `.eth` registrar grants new owners `v2.RegistrationRoles` (set resolver, set subregistry, their admins, and the transfer admin), deliberately not `RoleRenew`: `.eth` names renew only through the registrar.

**Transfers.** `safeTransferFrom` refuses unless the registry is emancipated (no root role can reach registered names), the sender holds the transfer admin role and nobody else holds roles on the name. `Registry.Transfer` runs those checks first and returns `*TransferUnsafeError` or `*TransferDisallowedError`; `UnsafeTransfer` skips the first and third.

## Registration

`ETHRegistrar` keeps commit and reveal. The commitment is `keccak256(abi.encode(label, owner, secret, subregistry, resolver, duration, referrer))`; `v2.Commitment.Hash` computes it locally and the live tests check it against the contract's pure `makeCommitment`. Registration is paid in ERC-20 through `safeTransferFrom`, so the payer must approve the registrar first. On mainnet the `StandardRentPriceOracle` accepts USDC and DAI; prices are quoted in USD × 1e12 and converted per token (`v2.Oracle`). A year of a five-plus character name is about 8 USD; three and four character names cost more; expired names carry a premium that starts at 100 million USD and halves daily for 21 days after the 28-day grace period.

`v2.Registrar.Register` runs the pre-flight the contract would otherwise revert on: non-zero owner, minimum duration, commitment window, availability, allowance and balance. `Client.RegistrationOpen` reports whether the ENS DAO has granted the registrar its role on the `.eth` registry yet.

## Resolvers

Two resolver families exist under v2:

- **`PermissionedResolver`** (one proxy per owner, deployed through `VerifiableFactory`): setters take the DNS-encoded name, so one resolver serves any number of names, and writes are authorised by roles on the resolver, not by name ownership. It has no direct getters; it answers only through `resolve(bytes,bytes)`, which the Universal Resolver calls. It does not implement `pubkey`, DNS records or `zonehash`, and `clearRecords` is replaced by linking a name to record 0. `v2.PermissionedResolver` wraps it; `Client.DeployPermissionedResolver` deploys one, optionally with initial records set inside `initialize`.
- **`PublicResolverV2`** (shared, node-keyed, v1-compatible): serves wrapped v1 names migrated with `CANNOT_SET_RESOLVER`. Authorisation goes through the v1 NameWrapper's knowledge of the name plus the v2 owner.

Two read-only mirrors also appear: `ENSV1Resolver` (the resolver the v2 tree points at for names still held in v1) and `ENSV2Resolver` (installed in the v1 registry for `.eth`, so v1 clients reach v2 names). `ens.Client.WritableResolver` unwraps them with `ICompositeResolver.getResolver` and reports the family of the contract that holds the records.

## Resolution

`UniversalResolverV2` keeps the `resolve(bytes,bytes)` and `reverse(bytes,uint256)` entry points of the v1 resolver, walks the v2 tree, enforces ENSIP-10 and follows CCIP-Read through a batch gateway. It does not normalise names; go-ens normalises (ENSIP-15) and DNS-encodes on the client. Labels over 255 bytes cannot be encoded for v2, so `DNSEncode` errors on them.

The canonical proxy `0xeEeEEEeE14D718C2B47D9923Deab1335E144EeEe` is the DAO-designated entry point before and after the cutover. `Client.IsENSv2` reports whether it has been switched; `WithUniversalResolver(cfg.V2.UniversalResolverV2)` targets the v2 implementation directly.

## Reverse resolution

Primary names stay in the v1 reverse registrars (`addr.reverse`, and `default.reverse` for ENSIP-19 defaults); the v2 root registers `reverse` with `ENSV1Resolver` so lookups mirror into v1. The `reverse` call checks that the name resolves forward to the queried address and reverts `ReverseAddressMismatch` otherwise. `ens.EVMCoinType` maps a chain id to its ENSIP-19 coin type and `ReverseName` builds the reverse name.

Writes: an externally owned account keeps using `v1.ReverseRegistrar.SetName`. Contracts use the v2 adapters (`v2.ReverseAdapters`), which the ENS DAO authorises as controllers of the v1 registrars; `SetPrimaryNameWithPermissionedResolver` claims the reverse node with a `PermissionedResolver` and stores the name there.

## Migration of v1 names

Every active v1 `.eth` name was reserved in the `.eth` registry before the cutover with the v1 expiry plus a bonus period and `ENSV1Resolver` as resolver. `ens.Client.MigrationStatus` reads the state:

| State | Meaning |
| --- | --- |
| `V1Reserved` | Unmigrated. Owner in `BaseRegistrar` or `NameWrapper`; renews through `ETHRenewerV1` (`v2.Client.Renew` routes there). |
| `V1ReservedGrace` | Reservation lapsed less than the grace period ago; only its v1 owner can still act. |
| `MigratedFromV1` | Registered on v2 with the `WAS_RESERVED` role. |
| `NativeV2` | Registered directly on v2. |
| `V2Expired` | A v2 registration lapsed; its owner may renew within the grace period. |
| `Available` | Registrable. |
| `Unknown` | The v2 registry never saw the label. |

Transfer-based migration (sending the v1 token to the migration controllers, or batching through `MigrationHelper`) is planned for v5.1. `v1.NameWrapper` and `Fuses` already report whether a wrapped name is locked, which decides the controller it will go through.

## What is v1-only now

`v1.ETHController` registration and renewal (disabled by the ENS DAO at migration phase 3), `v1.Registry` writes, the auction registrar and deed, and DNS claiming through `v1.DNSRegistrar`. DNS names resolve under v2 through `DNSTLDResolver` with no claim step.
