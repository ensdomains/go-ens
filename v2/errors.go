// Copyright 2026 Weald Technology Trading.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0

package v2

import (
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"time"

	ens "github.com/ensdomains/go-ens/v5"
	"github.com/ensdomains/go-ens/v5/internal/revert"
	"github.com/ethereum/go-ethereum/common"
)

// Sentinels for errors.Is. Registrar sentinels shared with v1 live in the
// ens package (ens.ErrNameNotAvailable and friends).
var (
	ErrUnauthorized             = errors.New("account lacks the required roles")
	ErrInvalidRoles             = errors.New("invalid role bitmap")
	ErrTransferDisallowed       = errors.New("transfer disallowed")
	ErrTransferUnsafe           = errors.New("transfer unsafe; use UnsafeTransfer")
	ErrInsufficientAllowance    = errors.New("insufficient ERC-20 allowance")
	ErrInsufficientBalance      = errors.New("insufficient ERC-20 balance")
	ErrPaymentFailed            = errors.New("ERC-20 payment failed")
	ErrPaymentTokenNotSupported = errors.New("payment token not supported")
	ErrInvalidLabel             = errors.New("label is not valid")
	ErrInvalidOwner             = errors.New("invalid owner")
	ErrLabelAlreadyRegistered   = errors.New("label already registered")
	ErrLabelAlreadyReserved     = errors.New("label already reserved")
	ErrLabelExpired             = errors.New("label expired")
	ErrInvalidExpiry            = errors.New("invalid expiry")
	ErrInvalidRecord            = errors.New("invalid resolver record")
	ErrInvalidContentType       = errors.New("ABI content type must be a single bit")
	ErrInvalidEVMAddress        = errors.New("EVM address must be empty or 20 bytes")
	ErrVerificationFailed       = errors.New("proxy verification failed")
	ErrSaltInUse                = errors.New("proxy salt already used by this deployer")
	ErrUnauthorizedNamer        = errors.New("caller may not name this account")
	ErrNoSubregistry            = errors.New("name has no subregistry")
	ErrNotV2Name                = errors.New("name is not held in a v2 registry")
	ErrResolverFamily           = errors.New("resolver is not of the expected family")
	ErrRegistrationClosed       = errors.New("v2 registration is not enabled")
	ErrMixedRenewers            = errors.New("labels route to different renewers")
	ErrSubnameExists            = errors.New("subname already registered")
	ErrCommitmentExpired        = errors.New("commitment expired")
	ErrNoCommitment             = errors.New("no commitment found")
	ErrValueNotAllowed          = errors.New("v2 registration is paid in ERC-20; do not send ETH")
)

// Registrar errors.

// NameNotAvailableError reports that the label is registered, reserved or in grace.
type NameNotAvailableError struct{ Label string }

func (e *NameNotAvailableError) Error() string { return fmt.Sprintf("name %q not available", e.Label) }
func (e *NameNotAvailableError) Unwrap() error { return ens.ErrNameNotAvailable }

// CommitmentTooNewError reports a register call before MIN_COMMITMENT_AGE elapsed.
type CommitmentTooNewError struct {
	Commitment [32]byte
	ValidFrom  time.Time
	BlockTime  time.Time
}

func (e *CommitmentTooNewError) Error() string {
	return fmt.Sprintf("commitment too new: valid from %s (block time %s)", e.ValidFrom.UTC().Format(time.RFC3339), e.BlockTime.UTC().Format(time.RFC3339))
}
func (e *CommitmentTooNewError) Unwrap() error { return ens.ErrCommitmentTooNew }

// CommitmentTooOldError reports a register call after MAX_COMMITMENT_AGE elapsed.
type CommitmentTooOldError struct {
	Commitment [32]byte
	ValidTo    time.Time
	BlockTime  time.Time
}

func (e *CommitmentTooOldError) Error() string {
	return fmt.Sprintf("commitment too old: valid until %s (block time %s)", e.ValidTo.UTC().Format(time.RFC3339), e.BlockTime.UTC().Format(time.RFC3339))
}
func (e *CommitmentTooOldError) Unwrap() error { return ens.ErrCommitmentTooOld }

// UnexpiredCommitmentExistsError reports that the same commitment was made recently.
type UnexpiredCommitmentExistsError struct{ Commitment [32]byte }

func (e *UnexpiredCommitmentExistsError) Error() string {
	return "unexpired commitment exists; use a new secret"
}
func (e *UnexpiredCommitmentExistsError) Unwrap() error { return ens.ErrUnexpiredCommitment }

// DurationTooShortError reports a duration below the registrar minimum.
type DurationTooShortError struct{ Duration, Min time.Duration }

func (e *DurationTooShortError) Error() string {
	return fmt.Sprintf("duration %s is shorter than the minimum %s", e.Duration, e.Min)
}
func (e *DurationTooShortError) Unwrap() error { return ens.ErrDurationTooShort }

// NameNotRenewableError reports that the renewer cannot renew the label in its state.
type NameNotRenewableError struct{ Label string }

func (e *NameNotRenewableError) Error() string { return fmt.Sprintf("name %q not renewable", e.Label) }
func (e *NameNotRenewableError) Unwrap() error { return ens.ErrNameNotRenewable }

// NotValidError reports that the oracle rejects the label.
type NotValidError struct{ Label string }

func (e *NotValidError) Error() string { return fmt.Sprintf("label %q is not valid", e.Label) }
func (e *NotValidError) Unwrap() error { return ErrInvalidLabel }

// PaymentTokenNotSupportedError reports a token the oracle does not price in.
type PaymentTokenNotSupportedError struct{ Token common.Address }

func (e *PaymentTokenNotSupportedError) Error() string {
	return fmt.Sprintf("payment token %s not supported", e.Token.Hex())
}
func (e *PaymentTokenNotSupportedError) Unwrap() error { return ErrPaymentTokenNotSupported }

// SafeERC20FailedOperationError reports a token that returned false on transfer.
type SafeERC20FailedOperationError struct{ Token common.Address }

func (e *SafeERC20FailedOperationError) Error() string {
	return fmt.Sprintf("ERC-20 operation on %s failed", e.Token.Hex())
}
func (e *SafeERC20FailedOperationError) Unwrap() error { return ErrPaymentFailed }

// InsufficientAllowanceError is raised client-side before a payment that
// would fail, or decoded from an OpenZeppelin v5 token revert.
type InsufficientAllowanceError struct {
	Token, Spender common.Address
	Required, Have *big.Int
}

func (e *InsufficientAllowanceError) Error() string {
	return fmt.Sprintf("allowance of %s for %s is %s, need %s", e.Token.Hex(), e.Spender.Hex(), e.Have, e.Required)
}
func (e *InsufficientAllowanceError) Unwrap() error { return ErrInsufficientAllowance }

// InsufficientBalanceError is raised client-side before a payment that
// would fail, or decoded from an OpenZeppelin v5 token revert.
type InsufficientBalanceError struct {
	Token          common.Address
	Required, Have *big.Int
}

func (e *InsufficientBalanceError) Error() string {
	return fmt.Sprintf("balance of %s is %s, need %s", e.Token.Hex(), e.Have, e.Required)
}
func (e *InsufficientBalanceError) Unwrap() error { return ErrInsufficientBalance }

// Registry errors.

type LabelAlreadyRegisteredError struct{ Label string }

func (e *LabelAlreadyRegisteredError) Error() string {
	return fmt.Sprintf("label %q already registered", e.Label)
}
func (e *LabelAlreadyRegisteredError) Unwrap() error { return ErrLabelAlreadyRegistered }

type LabelAlreadyReservedError struct{ Label string }

func (e *LabelAlreadyReservedError) Error() string {
	return fmt.Sprintf("label %q already reserved", e.Label)
}
func (e *LabelAlreadyReservedError) Unwrap() error { return ErrLabelAlreadyReserved }

type LabelExpiredError struct{ TokenID *big.Int }

func (e *LabelExpiredError) Error() string { return "label expired" }
func (e *LabelExpiredError) Unwrap() error { return ErrLabelExpired }

type CannotReduceExpiryError struct{ Old, New time.Time }

func (e *CannotReduceExpiryError) Error() string {
	return fmt.Sprintf("cannot reduce expiry from %s to %s", e.Old.UTC().Format(time.RFC3339), e.New.UTC().Format(time.RFC3339))
}
func (e *CannotReduceExpiryError) Unwrap() error { return ErrInvalidExpiry }

type CannotSetPastExpiryError struct{ Expiry time.Time }

func (e *CannotSetPastExpiryError) Error() string {
	return fmt.Sprintf("cannot set expiry in the past (%s)", e.Expiry.UTC().Format(time.RFC3339))
}
func (e *CannotSetPastExpiryError) Unwrap() error { return ErrInvalidExpiry }

type InvalidOwnerError struct{}

func (e *InvalidOwnerError) Error() string { return "invalid owner" }
func (e *InvalidOwnerError) Unwrap() error { return ErrInvalidOwner }

// TransferDisallowedError reports that the sender lacks ROLE_CAN_TRANSFER_ADMIN.
type TransferDisallowedError struct {
	TokenID *big.Int
	From    common.Address
}

func (e *TransferDisallowedError) Error() string {
	return "transfer disallowed: sender cannot transfer this name"
}
func (e *TransferDisallowedError) Unwrap() error { return ErrTransferDisallowed }

// TransferUnsafeReason explains why safeTransferFrom refused.
type TransferUnsafeReason int

const (
	TransferUnsafeNotEmancipated TransferUnsafeReason = iota + 1
	TransferUnsafeMultipleAssignees
)

// TransferUnsafeError reports that safeTransferFrom refused because the registry is
// not emancipated or other accounts hold roles on the name; UnsafeTransfer
// skips these checks.
type TransferUnsafeError struct {
	Reason  TransferUnsafeReason
	TokenID *big.Int
	From    common.Address
}

func (e *TransferUnsafeError) Error() string {
	switch e.Reason {
	case TransferUnsafeNotEmancipated:
		return "transfer unsafe until the registry is emancipated"
	case TransferUnsafeMultipleAssignees:
		return "transfer unsafe: other accounts hold roles on this name"
	default:
		return "transfer unsafe: other accounts hold roles on this name"
	}
}
func (e *TransferUnsafeError) Unwrap() error { return ErrTransferUnsafe }

// ERC1155Error is any OpenZeppelin ERC-1155 custom error.
type ERC1155Error struct {
	Name string
	Args []any
}

func (e *ERC1155Error) Error() string { return "erc1155: " + e.Name }

// EAC errors.

// UnauthorizedRolesError is EACUnauthorizedAccountRoles.
type UnauthorizedRolesError struct {
	Resource *big.Int
	Roles    Roles
	Account  common.Address
}

func (e *UnauthorizedRolesError) Error() string {
	return fmt.Sprintf("account %s lacks roles %s on resource 0x%s", e.Account.Hex(), e.Roles, e.Resource.Text(16))
}
func (e *UnauthorizedRolesError) Unwrap() error { return ErrUnauthorized }

type CannotGrantRolesError struct {
	Resource *big.Int
	Roles    Roles
	Account  common.Address
}

func (e *CannotGrantRolesError) Error() string {
	return fmt.Sprintf("cannot grant roles %s to %s", e.Roles, e.Account.Hex())
}
func (e *CannotGrantRolesError) Unwrap() error { return ErrUnauthorized }

type CannotRevokeRolesError struct {
	Resource *big.Int
	Roles    Roles
	Account  common.Address
}

func (e *CannotRevokeRolesError) Error() string {
	return fmt.Sprintf("cannot revoke roles %s from %s", e.Roles, e.Account.Hex())
}
func (e *CannotRevokeRolesError) Unwrap() error { return ErrUnauthorized }

type RootResourceNotAllowedError struct{}

func (e *RootResourceNotAllowedError) Error() string {
	return "root resource not allowed here; use the root-role functions"
}

type MaxAssigneesError struct{ Resource, Role *big.Int }

func (e *MaxAssigneesError) Error() string { return "a role already has the maximum of 15 assignees" }

type MinAssigneesError struct{ Resource, Role *big.Int }

func (e *MinAssigneesError) Error() string { return "a role cannot lose its last assignee" }

type InvalidRoleBitmapError struct{ Bitmap *big.Int }

func (e *InvalidRoleBitmapError) Error() string { return "invalid role bitmap 0x" + e.Bitmap.Text(16) }
func (e *InvalidRoleBitmapError) Unwrap() error { return ErrInvalidRoles }

type InvalidAccountError struct{}

func (e *InvalidAccountError) Error() string { return "invalid account" }

// Resolver errors.

type InvalidRecordError struct{}

func (e *InvalidRecordError) Error() string { return "invalid resolver record" }
func (e *InvalidRecordError) Unwrap() error { return ErrInvalidRecord }

type InvalidContentTypeError struct{ ContentType *big.Int }

func (e *InvalidContentTypeError) Error() string {
	return "invalid ABI content type " + e.ContentType.String()
}
func (e *InvalidContentTypeError) Unwrap() error { return ErrInvalidContentType }

type InvalidEVMAddressError struct{ Data []byte }

func (e *InvalidEVMAddressError) Error() string {
	return "invalid EVM address 0x" + hex.EncodeToString(e.Data)
}
func (e *InvalidEVMAddressError) Unwrap() error { return ErrInvalidEVMAddress }

// Factory and proxy errors.

type VerificationFailedError struct{ Proxy common.Address }

func (e *VerificationFailedError) Error() string {
	return "proxy " + e.Proxy.Hex() + " failed verification"
}
func (e *VerificationFailedError) Unwrap() error { return ErrVerificationFailed }

// ProxyError is any OpenZeppelin proxy/initializer/UUPS custom error.
type ProxyError struct {
	Name string
	Args []any
}

func (e *ProxyError) Error() string { return "proxy: " + e.Name }

// Reverse errors.

type UnauthorizedNamerError struct{ Namer common.Address }

func (e *UnauthorizedNamerError) Error() string {
	return e.Namer.Hex() + " is not authorised to name this account"
}
func (e *UnauthorizedNamerError) Unwrap() error { return ErrUnauthorizedNamer }

// OwnableError is an OpenZeppelin Ownable custom error.
type OwnableError struct {
	Name    string
	Account common.Address
}

func (e *OwnableError) Error() string { return "ownable: " + e.Name + " " + e.Account.Hex() }

// decoder translates every v2 custom error, plus those of the Universal
// Resolver.
var decoder = buildDecoder()

// DecodeRevert translates a revert error from any v2 contract into a typed
// error. Errors without a recognised payload are returned unchanged;
// Error(string) reverts from mainnet USDC and DAI about allowance or balance
// map to ErrInsufficientAllowance and ErrInsufficientBalance.
func DecodeRevert(err error) error {
	out := decoder.Translate(err)
	var s *ens.StringRevertError
	if errors.As(out, &s) {
		switch {
		case strings.Contains(s.Reason, "allowance"):
			return fmt.Errorf("%w: %s", ErrInsufficientAllowance, s.Reason)
		case strings.Contains(s.Reason, "balance"):
			return fmt.Errorf("%w: %s", ErrInsufficientBalance, s.Reason)
		}
	}
	return out
}

func unixTime(v any) time.Time {
	switch x := v.(type) {
	case uint64:
		return time.Unix(int64(x), 0).UTC()
	case *big.Int:
		return time.Unix(x.Int64(), 0).UTC()
	}
	return time.Time{}
}

func bigArg(a []any, i int) *big.Int {
	if i < len(a) {
		if x, ok := a[i].(*big.Int); ok {
			return x
		}
	}
	return new(big.Int)
}

func addrArg(a []any, i int) common.Address {
	if i < len(a) {
		if x, ok := a[i].(common.Address); ok {
			return x
		}
	}
	return common.Address{}
}

func strArg(a []any, i int) string { //nolint:unparam // symmetry with the other arg helpers
	if i < len(a) {
		if x, ok := a[i].(string); ok {
			return x
		}
	}
	return ""
}

func bytesArg(a []any, i int) []byte {
	if i < len(a) {
		if x, ok := a[i].([]byte); ok {
			return x
		}
	}
	return nil
}

func b32Arg(a []any, i int) [32]byte {
	if i < len(a) {
		if x, ok := a[i].([32]byte); ok {
			return x
		}
	}
	return [32]byte{}
}

func secondsArg(a []any, i int) time.Duration {
	if i < len(a) {
		if x, ok := a[i].(uint64); ok {
			return time.Duration(x) * time.Second
		}
	}
	return 0
}

func buildDecoder() *revert.Decoder {
	d := ens.URDecoder()
	// Registrar / renewer / oracle.
	d.Register("UnexpiredCommitmentExists(bytes32)", func(a []any) error {
		return &UnexpiredCommitmentExistsError{Commitment: b32Arg(a, 0)}
	})
	d.Register("CommitmentTooNew(bytes32,uint64,uint64)", func(a []any) error {
		return &CommitmentTooNewError{Commitment: b32Arg(a, 0), ValidFrom: unixTime(a[1]), BlockTime: unixTime(a[2])}
	})
	d.Register("CommitmentTooOld(bytes32,uint64,uint64)", func(a []any) error {
		return &CommitmentTooOldError{Commitment: b32Arg(a, 0), ValidTo: unixTime(a[1]), BlockTime: unixTime(a[2])}
	})
	d.Register("NameNotAvailable(string)", func(a []any) error { return &NameNotAvailableError{Label: strArg(a, 0)} })
	d.Register("DurationTooShort(uint64,uint64)", func(a []any) error {
		return &DurationTooShortError{Duration: secondsArg(a, 0), Min: secondsArg(a, 1)}
	})
	d.Register("NameNotRenewable(string)", func(a []any) error { return &NameNotRenewableError{Label: strArg(a, 0)} })
	d.Register("NotValid(string)", func(a []any) error { return &NotValidError{Label: strArg(a, 0)} })
	d.Register("PaymentTokenNotSupported(address)", func(a []any) error {
		return &PaymentTokenNotSupportedError{Token: addrArg(a, 0)}
	})
	d.Register("SafeERC20FailedOperation(address)", func(a []any) error {
		return &SafeERC20FailedOperationError{Token: addrArg(a, 0)}
	})
	d.Register("ERC20InsufficientAllowance(address,uint256,uint256)", func(a []any) error {
		return &InsufficientAllowanceError{Spender: addrArg(a, 0), Have: bigArg(a, 1), Required: bigArg(a, 2)}
	})
	d.Register("ERC20InsufficientBalance(address,uint256,uint256)", func(a []any) error {
		return &InsufficientBalanceError{Have: bigArg(a, 1), Required: bigArg(a, 2)}
	})
	d.Register("InvalidOwner()", func([]any) error { return &InvalidOwnerError{} })
	d.Register("MaxCommitmentAgeTooLow()", func([]any) error { return errors.New("max commitment age too low") })
	d.Register("UnauthorizedCaller(address)", func(a []any) error {
		return fmt.Errorf("%w: caller %s", ErrUnauthorized, addrArg(a, 0).Hex())
	})
	// Registry.
	d.Register("LabelAlreadyRegistered(string)", func(a []any) error { return &LabelAlreadyRegisteredError{Label: strArg(a, 0)} })
	d.Register("LabelAlreadyReserved(string)", func(a []any) error { return &LabelAlreadyReservedError{Label: strArg(a, 0)} })
	d.Register("LabelExpired(uint256)", func(a []any) error { return &LabelExpiredError{TokenID: bigArg(a, 0)} })
	d.Register("CannotReduceExpiry(uint64,uint64)", func(a []any) error {
		return &CannotReduceExpiryError{Old: unixTime(a[0]), New: unixTime(a[1])}
	})
	d.Register("CannotSetPastExpiry(uint64)", func(a []any) error { return &CannotSetPastExpiryError{Expiry: unixTime(a[0])} })
	d.Register("TransferDisallowed(uint256,address)", func(a []any) error {
		return &TransferDisallowedError{TokenID: bigArg(a, 0), From: addrArg(a, 1)}
	})
	d.Register("TransferUnsafeUntilRegistryIsEmancipated()", func([]any) error {
		return &TransferUnsafeError{Reason: TransferUnsafeNotEmancipated}
	})
	d.Register("TransferUnsafeWithMultipleAssignees(uint256,address)", func(a []any) error {
		return &TransferUnsafeError{Reason: TransferUnsafeMultipleAssignees, TokenID: bigArg(a, 0), From: addrArg(a, 1)}
	})
	for _, sig := range []string{
		"ERC1155InsufficientBalance(address,uint256,uint256,uint256)", "ERC1155InvalidSender(address)",
		"ERC1155InvalidReceiver(address)", "ERC1155MissingApprovalForAll(address,address)",
		"ERC1155InvalidApprover(address)", "ERC1155InvalidOperator(address)", "ERC1155InvalidArrayLength(uint256,uint256)",
	} {
		name := sig[:strings.IndexByte(sig, '(')]
		d.Register(sig, func(a []any) error { return &ERC1155Error{Name: name, Args: a} })
	}
	// EnhancedAccessControl.
	d.Register("EACUnauthorizedAccountRoles(uint256,uint256,address)", func(a []any) error {
		return &UnauthorizedRolesError{Resource: bigArg(a, 0), Roles: mustRoles(bigArg(a, 1)), Account: addrArg(a, 2)}
	})
	d.Register("EACCannotGrantRoles(uint256,uint256,address)", func(a []any) error {
		return &CannotGrantRolesError{Resource: bigArg(a, 0), Roles: mustRoles(bigArg(a, 1)), Account: addrArg(a, 2)}
	})
	d.Register("EACCannotRevokeRoles(uint256,uint256,address)", func(a []any) error {
		return &CannotRevokeRolesError{Resource: bigArg(a, 0), Roles: mustRoles(bigArg(a, 1)), Account: addrArg(a, 2)}
	})
	d.Register("EACRootResourceNotAllowed()", func([]any) error { return &RootResourceNotAllowedError{} })
	d.Register("EACMaxAssignees(uint256,uint256)", func(a []any) error { return &MaxAssigneesError{Resource: bigArg(a, 0), Role: bigArg(a, 1)} })
	d.Register("EACMinAssignees(uint256,uint256)", func(a []any) error { return &MinAssigneesError{Resource: bigArg(a, 0), Role: bigArg(a, 1)} })
	d.Register("EACInvalidRoleBitmap(uint256)", func(a []any) error { return &InvalidRoleBitmapError{Bitmap: bigArg(a, 0)} })
	d.Register("EACInvalidAccount()", func([]any) error { return &InvalidAccountError{} })
	// Resolver.
	d.Register("InvalidRecord()", func([]any) error { return &InvalidRecordError{} })
	d.Register("InvalidContentType(uint256)", func(a []any) error { return &InvalidContentTypeError{ContentType: bigArg(a, 0)} })
	d.Register("InvalidEVMAddress(bytes)", func(a []any) error { return &InvalidEVMAddressError{Data: bytesArg(a, 0)} })
	// Factory and OpenZeppelin proxies.
	d.Register("VerificationFailed(address)", func(a []any) error { return &VerificationFailedError{Proxy: addrArg(a, 0)} })
	for _, sig := range []string{
		"InvalidInitialization()", "NotInitializing()", "FailedCall()", "FailedDeployment()", "AddressEmptyCode(address)",
		"ERC1967InvalidImplementation(address)", "ERC1967NonPayable()", "UUPSUnauthorizedCallContext()",
		"UUPSUnsupportedProxiableUUID(bytes32)",
	} {
		name := sig[:strings.IndexByte(sig, '(')]
		d.Register(sig, func(a []any) error { return &ProxyError{Name: name, Args: a} })
	}
	// Reverse adapters and Ownable.
	d.Register("UnauthorizedNamer(address)", func(a []any) error { return &UnauthorizedNamerError{Namer: addrArg(a, 0)} })
	d.Register("OwnableUnauthorizedAccount(address)", func(a []any) error {
		return &OwnableError{Name: "unauthorized account", Account: addrArg(a, 0)}
	})
	d.Register("OwnableInvalidOwner(address)", func(a []any) error {
		return &OwnableError{Name: "invalid owner", Account: addrArg(a, 0)}
	})
	return d
}
