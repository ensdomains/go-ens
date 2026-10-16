// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package wrapperregistry

import (
	"bytes"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = bytes.Equal
	_ = errors.New
	_ = big.NewInt
	_ = common.Big1
	_ = types.BloomLookup
	_ = abi.ConvertType
)

// IPermissionedRegistryState is an auto generated low-level Go binding around an user-defined struct.
type IPermissionedRegistryState struct {
	Status      uint8
	Expiry      uint64
	LatestOwner common.Address
	TokenId     *big.Int
	Resource    *big.Int
}

// LibMigrationData is an auto generated low-level Go binding around an user-defined struct.
type LibMigrationData struct {
	Label       string
	Owner       common.Address
	Subregistry common.Address
	Resolver    common.Address
}

// WrapperRegistryMetaData contains all meta data concerning the WrapperRegistry contract.
var WrapperRegistryMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractINameWrapper\",\"name\":\"nameWrapper\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"graveyard\",\"type\":\"address\"},{\"internalType\":\"contractIVerifiableFactory\",\"name\":\"verifiableFactory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"ensV1Resolver\",\"type\":\"address\"},{\"internalType\":\"contractIAddressSet\",\"name\":\"upgradeSet\",\"type\":\"address\"},{\"internalType\":\"contractILabelStore\",\"name\":\"labelStore\",\"type\":\"address\"},{\"internalType\":\"contractIRegistryURIRenderer\",\"name\":\"uriRenderer\",\"type\":\"address\"},{\"internalType\":\"contractIAddressSet\",\"name\":\"publicResolverSet\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"publicResolver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"namer\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"oldExpiry\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"newExpiry\",\"type\":\"uint64\"}],\"name\":\"CannotReduceExpiry\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"expiry\",\"type\":\"uint64\"}],\"name\":\"CannotSetPastExpiry\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"resource\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"roleBitmap\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"EACCannotGrantRoles\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"resource\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"roleBitmap\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"EACCannotRevokeRoles\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EACInvalidAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roleBitmap\",\"type\":\"uint256\"}],\"name\":\"EACInvalidRoleBitmap\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"resource\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"role\",\"type\":\"uint256\"}],\"name\":\"EACMaxAssignees\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"resource\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"role\",\"type\":\"uint256\"}],\"name\":\"EACMinAssignees\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EACRootResourceNotAllowed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"resource\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"roleBitmap\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"EACUnauthorizedAccountRoles\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC1155InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC1155InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idsLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"valuesLength\",\"type\":\"uint256\"}],\"name\":\"ERC1155InvalidArrayLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC1155InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC1155InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC1155InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC1155MissingApprovalForAll\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"FrozenTokenApproval\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"}],\"name\":\"LabelAlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"}],\"name\":\"LabelAlreadyReserved\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"LabelExpired\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"NameDataMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"NameNotLocked\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NameRequiresMigration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"TransferDisallowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferUnsafeUntilRegistryIsEmancipated\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"TransferUnsafeWithMultipleAssignees\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"UnauthorizedCaller\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"UpgradeTargetNotApproved\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"resource\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldRoleBitmap\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newRoleBitmap\",\"type\":\"uint256\"}],\"name\":\"EACRolesChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"newExpiry\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ExpiryUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"labelHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"expiry\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"LabelRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"labelHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"expiry\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"LabelReserved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"LabelUnregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIRegistry\",\"name\":\"parent\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ParentUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"RegistryCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ResolverUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIRegistry\",\"name\":\"subregistry\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SubregistryUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"oldTokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newTokenId\",\"type\":\"uint256\"}],\"name\":\"TokenRegenerated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"resource\",\"type\":\"uint256\"}],\"name\":\"TokenResource\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"renderer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"URIUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"GRAVEYARD\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LABEL_STORE\",\"outputs\":[{\"internalType\":\"contractILabelStore\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NAME_WRAPPER\",\"outputs\":[{\"internalType\":\"contractINameWrapper\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PUBLIC_RESOLVER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PUBLIC_RESOLVER_SET\",\"outputs\":[{\"internalType\":\"contractIAddressSet\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ROOT_RESOURCE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_SET\",\"outputs\":[{\"internalType\":\"contractIAddressSet\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"URI_RENDERER\",\"outputs\":[{\"internalType\":\"contractIRegistryURIRenderer\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"V1_RESOLVER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERIFIABLE_FACTORY\",\"outputs\":[{\"internalType\":\"contractIVerifiableFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WRAPPER_REGISTRY_IMPL\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"canUpgradeFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"}],\"name\":\"findExpiry\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"}],\"name\":\"findOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"}],\"name\":\"findTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"contractIRegistry\",\"name\":\"subregistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"}],\"internalType\":\"structLibMigration.Data[]\",\"name\":\"mds\",\"type\":\"tuple[]\"}],\"name\":\"finishERC1155Migration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"anyId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"roleBitmap\",\"type\":\"uint256\"}],\"name\":\"getAssigneeCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"counts\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mask\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"anyId\",\"type\":\"uint256\"}],\"name\":\"getExpiry\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"anyId\",\"type\":\"uint256\"}],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getParent\",\"outputs\":[{\"internalType\":\"contractIRegistry\",\"name\":\"parent\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"}],\"name\":\"getResolver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"anyId\",\"type\":\"uint256\"}],\"name\":\"getResource\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"anyId\",\"type\":\"uint256\"}],\"name\":\"getState\",\"outputs\":[{\"components\":[{\"internalType\":\"enumIPermissionedRegistry.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"expiry\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"latestOwner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resource\",\"type\":\"uint256\"}],\"internalType\":\"structIPermissionedRegistry.State\",\"name\":\"state\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"anyId\",\"type\":\"uint256\"}],\"name\":\"getStatus\",\"outputs\":[{\"internalType\":\"enumIPermissionedRegistry.Status\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"}],\"name\":\"getSubregistry\",\"outputs\":[{\"internalType\":\"contractIRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"anyId\",\"type\":\"uint256\"}],\"name\":\"getTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"uri_\",\"type\":\"string\"},{\"internalType\":\"contractIRegistryURIRenderer\",\"name\":\"renderer\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWrappedName\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWrappedNode\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"anyId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"roleBitmap\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRoles\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roleBitmap\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRootRoles\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"resource\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"roleBitmap\",\"type\":\"uint256\"}],\"name\":\"hasAssignees\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"anyId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"roleBitmap\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRoles\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roleBitmap\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRootRoles\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"contractIRegistry\",\"name\":\"parentRegistry\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"childLabel\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"roleBitmap\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"namer\",\"type\":\"address\"}],\"name\":\"isContractNamer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isEmancipated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"resource\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"roleBitmap\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isOnlyAssignee\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"latestOwnerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"contractIRegistry\",\"name\":\"registry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"roleBitmap\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"expiry\",\"type\":\"uint64\"}],\"name\":\"register\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"anyId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"newExpiry\",\"type\":\"uint64\"}],\"name\":\"renew\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"anyId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"roleBitmap\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRoles\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roleBitmap\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRootRoles\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"anyId\",\"type\":\"uint256\"}],\"name\":\"roleCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"anyId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"roles\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIRegistry\",\"name\":\"parent\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"}],\"name\":\"setParent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"anyId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"}],\"name\":\"setResolver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"anyId\",\"type\":\"uint256\"},{\"internalType\":\"contractIRegistry\",\"name\":\"registry\",\"type\":\"address\"}],\"name\":\"setSubregistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uri_\",\"type\":\"string\"},{\"internalType\":\"contractIRegistryURIRenderer\",\"name\":\"renderer\",\"type\":\"address\"}],\"name\":\"setURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"anyId\",\"type\":\"uint256\"}],\"name\":\"unregister\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"unsafeBatchTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"unsafeTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	ID:  "WrapperRegistry",
}

// WrapperRegistry is an auto generated Go binding around an Ethereum contract.
type WrapperRegistry struct {
	abi abi.ABI
}

// GetABI returns the ABI associated with this contract binding.
func (c *WrapperRegistry) GetABI() abi.ABI {
	return c.abi
}

// NewWrapperRegistry creates a new instance of WrapperRegistry.
func NewWrapperRegistry() *WrapperRegistry {
	parsed, err := WrapperRegistryMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &WrapperRegistry{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *WrapperRegistry) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address nameWrapper, address graveyard, address verifiableFactory, address ensV1Resolver, address upgradeSet, address labelStore, address uriRenderer, address publicResolverSet, address publicResolver, address namer) returns()
func (wrapperRegistry *WrapperRegistry) PackConstructor(nameWrapper common.Address, graveyard common.Address, verifiableFactory common.Address, ensV1Resolver common.Address, upgradeSet common.Address, labelStore common.Address, uriRenderer common.Address, publicResolverSet common.Address, publicResolver common.Address, namer common.Address) []byte {
	enc, err := wrapperRegistry.abi.Pack("", nameWrapper, graveyard, verifiableFactory, ensV1Resolver, upgradeSet, labelStore, uriRenderer, publicResolverSet, publicResolver, namer)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackGRAVEYARD is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5c1a6b68.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function GRAVEYARD() view returns(address)
func (wrapperRegistry *WrapperRegistry) PackGRAVEYARD() []byte {
	enc, err := wrapperRegistry.abi.Pack("GRAVEYARD")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGRAVEYARD is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5c1a6b68.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function GRAVEYARD() view returns(address)
func (wrapperRegistry *WrapperRegistry) TryPackGRAVEYARD() ([]byte, error) {
	return wrapperRegistry.abi.Pack("GRAVEYARD")
}

// UnpackGRAVEYARD is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5c1a6b68.
//
// Solidity: function GRAVEYARD() view returns(address)
func (wrapperRegistry *WrapperRegistry) UnpackGRAVEYARD(data []byte) (common.Address, error) {
	out, err := wrapperRegistry.abi.Unpack("GRAVEYARD", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackLABELSTORE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9dbba19d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function LABEL_STORE() view returns(address)
func (wrapperRegistry *WrapperRegistry) PackLABELSTORE() []byte {
	enc, err := wrapperRegistry.abi.Pack("LABEL_STORE")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackLABELSTORE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9dbba19d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function LABEL_STORE() view returns(address)
func (wrapperRegistry *WrapperRegistry) TryPackLABELSTORE() ([]byte, error) {
	return wrapperRegistry.abi.Pack("LABEL_STORE")
}

// UnpackLABELSTORE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9dbba19d.
//
// Solidity: function LABEL_STORE() view returns(address)
func (wrapperRegistry *WrapperRegistry) UnpackLABELSTORE(data []byte) (common.Address, error) {
	out, err := wrapperRegistry.abi.Unpack("LABEL_STORE", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackNAMEWRAPPER is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x192cf07d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function NAME_WRAPPER() view returns(address)
func (wrapperRegistry *WrapperRegistry) PackNAMEWRAPPER() []byte {
	enc, err := wrapperRegistry.abi.Pack("NAME_WRAPPER")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackNAMEWRAPPER is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x192cf07d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function NAME_WRAPPER() view returns(address)
func (wrapperRegistry *WrapperRegistry) TryPackNAMEWRAPPER() ([]byte, error) {
	return wrapperRegistry.abi.Pack("NAME_WRAPPER")
}

// UnpackNAMEWRAPPER is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x192cf07d.
//
// Solidity: function NAME_WRAPPER() view returns(address)
func (wrapperRegistry *WrapperRegistry) UnpackNAMEWRAPPER(data []byte) (common.Address, error) {
	out, err := wrapperRegistry.abi.Unpack("NAME_WRAPPER", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackPUBLICRESOLVER is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xffeb4a30.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function PUBLIC_RESOLVER() view returns(address)
func (wrapperRegistry *WrapperRegistry) PackPUBLICRESOLVER() []byte {
	enc, err := wrapperRegistry.abi.Pack("PUBLIC_RESOLVER")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackPUBLICRESOLVER is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xffeb4a30.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function PUBLIC_RESOLVER() view returns(address)
func (wrapperRegistry *WrapperRegistry) TryPackPUBLICRESOLVER() ([]byte, error) {
	return wrapperRegistry.abi.Pack("PUBLIC_RESOLVER")
}

// UnpackPUBLICRESOLVER is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xffeb4a30.
//
// Solidity: function PUBLIC_RESOLVER() view returns(address)
func (wrapperRegistry *WrapperRegistry) UnpackPUBLICRESOLVER(data []byte) (common.Address, error) {
	out, err := wrapperRegistry.abi.Unpack("PUBLIC_RESOLVER", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackPUBLICRESOLVERSET is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf923b685.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function PUBLIC_RESOLVER_SET() view returns(address)
func (wrapperRegistry *WrapperRegistry) PackPUBLICRESOLVERSET() []byte {
	enc, err := wrapperRegistry.abi.Pack("PUBLIC_RESOLVER_SET")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackPUBLICRESOLVERSET is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf923b685.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function PUBLIC_RESOLVER_SET() view returns(address)
func (wrapperRegistry *WrapperRegistry) TryPackPUBLICRESOLVERSET() ([]byte, error) {
	return wrapperRegistry.abi.Pack("PUBLIC_RESOLVER_SET")
}

// UnpackPUBLICRESOLVERSET is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf923b685.
//
// Solidity: function PUBLIC_RESOLVER_SET() view returns(address)
func (wrapperRegistry *WrapperRegistry) UnpackPUBLICRESOLVERSET(data []byte) (common.Address, error) {
	out, err := wrapperRegistry.abi.Unpack("PUBLIC_RESOLVER_SET", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackROOTRESOURCE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1c3fc3eb.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function ROOT_RESOURCE() view returns(uint256)
func (wrapperRegistry *WrapperRegistry) PackROOTRESOURCE() []byte {
	enc, err := wrapperRegistry.abi.Pack("ROOT_RESOURCE")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackROOTRESOURCE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1c3fc3eb.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function ROOT_RESOURCE() view returns(uint256)
func (wrapperRegistry *WrapperRegistry) TryPackROOTRESOURCE() ([]byte, error) {
	return wrapperRegistry.abi.Pack("ROOT_RESOURCE")
}

// UnpackROOTRESOURCE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x1c3fc3eb.
//
// Solidity: function ROOT_RESOURCE() view returns(uint256)
func (wrapperRegistry *WrapperRegistry) UnpackROOTRESOURCE(data []byte) (*big.Int, error) {
	out, err := wrapperRegistry.abi.Unpack("ROOT_RESOURCE", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackUPGRADEINTERFACEVERSION is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xad3cb1cc.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (wrapperRegistry *WrapperRegistry) PackUPGRADEINTERFACEVERSION() []byte {
	enc, err := wrapperRegistry.abi.Pack("UPGRADE_INTERFACE_VERSION")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUPGRADEINTERFACEVERSION is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xad3cb1cc.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (wrapperRegistry *WrapperRegistry) TryPackUPGRADEINTERFACEVERSION() ([]byte, error) {
	return wrapperRegistry.abi.Pack("UPGRADE_INTERFACE_VERSION")
}

// UnpackUPGRADEINTERFACEVERSION is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (wrapperRegistry *WrapperRegistry) UnpackUPGRADEINTERFACEVERSION(data []byte) (string, error) {
	out, err := wrapperRegistry.abi.Unpack("UPGRADE_INTERFACE_VERSION", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, nil
}

// PackUPGRADESET is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbf08765b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function UPGRADE_SET() view returns(address)
func (wrapperRegistry *WrapperRegistry) PackUPGRADESET() []byte {
	enc, err := wrapperRegistry.abi.Pack("UPGRADE_SET")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUPGRADESET is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbf08765b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function UPGRADE_SET() view returns(address)
func (wrapperRegistry *WrapperRegistry) TryPackUPGRADESET() ([]byte, error) {
	return wrapperRegistry.abi.Pack("UPGRADE_SET")
}

// UnpackUPGRADESET is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbf08765b.
//
// Solidity: function UPGRADE_SET() view returns(address)
func (wrapperRegistry *WrapperRegistry) UnpackUPGRADESET(data []byte) (common.Address, error) {
	out, err := wrapperRegistry.abi.Unpack("UPGRADE_SET", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackURIRENDERER is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9c42d7e8.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function URI_RENDERER() view returns(address)
func (wrapperRegistry *WrapperRegistry) PackURIRENDERER() []byte {
	enc, err := wrapperRegistry.abi.Pack("URI_RENDERER")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackURIRENDERER is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9c42d7e8.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function URI_RENDERER() view returns(address)
func (wrapperRegistry *WrapperRegistry) TryPackURIRENDERER() ([]byte, error) {
	return wrapperRegistry.abi.Pack("URI_RENDERER")
}

// UnpackURIRENDERER is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9c42d7e8.
//
// Solidity: function URI_RENDERER() view returns(address)
func (wrapperRegistry *WrapperRegistry) UnpackURIRENDERER(data []byte) (common.Address, error) {
	out, err := wrapperRegistry.abi.Unpack("URI_RENDERER", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackV1RESOLVER is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4b30228c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function V1_RESOLVER() view returns(address)
func (wrapperRegistry *WrapperRegistry) PackV1RESOLVER() []byte {
	enc, err := wrapperRegistry.abi.Pack("V1_RESOLVER")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackV1RESOLVER is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4b30228c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function V1_RESOLVER() view returns(address)
func (wrapperRegistry *WrapperRegistry) TryPackV1RESOLVER() ([]byte, error) {
	return wrapperRegistry.abi.Pack("V1_RESOLVER")
}

// UnpackV1RESOLVER is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x4b30228c.
//
// Solidity: function V1_RESOLVER() view returns(address)
func (wrapperRegistry *WrapperRegistry) UnpackV1RESOLVER(data []byte) (common.Address, error) {
	out, err := wrapperRegistry.abi.Unpack("V1_RESOLVER", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackVERIFIABLEFACTORY is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x18ad9b71.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function VERIFIABLE_FACTORY() view returns(address)
func (wrapperRegistry *WrapperRegistry) PackVERIFIABLEFACTORY() []byte {
	enc, err := wrapperRegistry.abi.Pack("VERIFIABLE_FACTORY")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackVERIFIABLEFACTORY is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x18ad9b71.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function VERIFIABLE_FACTORY() view returns(address)
func (wrapperRegistry *WrapperRegistry) TryPackVERIFIABLEFACTORY() ([]byte, error) {
	return wrapperRegistry.abi.Pack("VERIFIABLE_FACTORY")
}

// UnpackVERIFIABLEFACTORY is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x18ad9b71.
//
// Solidity: function VERIFIABLE_FACTORY() view returns(address)
func (wrapperRegistry *WrapperRegistry) UnpackVERIFIABLEFACTORY(data []byte) (common.Address, error) {
	out, err := wrapperRegistry.abi.Unpack("VERIFIABLE_FACTORY", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackWRAPPERREGISTRYIMPL is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x547c9d2d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function WRAPPER_REGISTRY_IMPL() view returns(address)
func (wrapperRegistry *WrapperRegistry) PackWRAPPERREGISTRYIMPL() []byte {
	enc, err := wrapperRegistry.abi.Pack("WRAPPER_REGISTRY_IMPL")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackWRAPPERREGISTRYIMPL is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x547c9d2d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function WRAPPER_REGISTRY_IMPL() view returns(address)
func (wrapperRegistry *WrapperRegistry) TryPackWRAPPERREGISTRYIMPL() ([]byte, error) {
	return wrapperRegistry.abi.Pack("WRAPPER_REGISTRY_IMPL")
}

// UnpackWRAPPERREGISTRYIMPL is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x547c9d2d.
//
// Solidity: function WRAPPER_REGISTRY_IMPL() view returns(address)
func (wrapperRegistry *WrapperRegistry) UnpackWRAPPERREGISTRYIMPL(data []byte) (common.Address, error) {
	out, err := wrapperRegistry.abi.Unpack("WRAPPER_REGISTRY_IMPL", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackBalanceOf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x00fdd58e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (wrapperRegistry *WrapperRegistry) PackBalanceOf(account common.Address, id *big.Int) []byte {
	enc, err := wrapperRegistry.abi.Pack("balanceOf", account, id)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackBalanceOf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x00fdd58e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (wrapperRegistry *WrapperRegistry) TryPackBalanceOf(account common.Address, id *big.Int) ([]byte, error) {
	return wrapperRegistry.abi.Pack("balanceOf", account, id)
}

// UnpackBalanceOf is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (wrapperRegistry *WrapperRegistry) UnpackBalanceOf(data []byte) (*big.Int, error) {
	out, err := wrapperRegistry.abi.Unpack("balanceOf", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackBalanceOfBatch is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4e1273f4.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (wrapperRegistry *WrapperRegistry) PackBalanceOfBatch(accounts []common.Address, ids []*big.Int) []byte {
	enc, err := wrapperRegistry.abi.Pack("balanceOfBatch", accounts, ids)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackBalanceOfBatch is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4e1273f4.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (wrapperRegistry *WrapperRegistry) TryPackBalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]byte, error) {
	return wrapperRegistry.abi.Pack("balanceOfBatch", accounts, ids)
}

// UnpackBalanceOfBatch is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (wrapperRegistry *WrapperRegistry) UnpackBalanceOfBatch(data []byte) ([]*big.Int, error) {
	out, err := wrapperRegistry.abi.Unpack("balanceOfBatch", data)
	if err != nil {
		return *new([]*big.Int), err
	}
	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	return out0, nil
}

// PackCanUpgradeFrom is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf41a143d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function canUpgradeFrom(address ) pure returns(bool allowed)
func (wrapperRegistry *WrapperRegistry) PackCanUpgradeFrom(arg0 common.Address) []byte {
	enc, err := wrapperRegistry.abi.Pack("canUpgradeFrom", arg0)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCanUpgradeFrom is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf41a143d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function canUpgradeFrom(address ) pure returns(bool allowed)
func (wrapperRegistry *WrapperRegistry) TryPackCanUpgradeFrom(arg0 common.Address) ([]byte, error) {
	return wrapperRegistry.abi.Pack("canUpgradeFrom", arg0)
}

// UnpackCanUpgradeFrom is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf41a143d.
//
// Solidity: function canUpgradeFrom(address ) pure returns(bool allowed)
func (wrapperRegistry *WrapperRegistry) UnpackCanUpgradeFrom(data []byte) (bool, error) {
	out, err := wrapperRegistry.abi.Unpack("canUpgradeFrom", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackFindExpiry is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6f537c72.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function findExpiry(string label) view returns(uint64)
func (wrapperRegistry *WrapperRegistry) PackFindExpiry(label string) []byte {
	enc, err := wrapperRegistry.abi.Pack("findExpiry", label)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFindExpiry is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6f537c72.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function findExpiry(string label) view returns(uint64)
func (wrapperRegistry *WrapperRegistry) TryPackFindExpiry(label string) ([]byte, error) {
	return wrapperRegistry.abi.Pack("findExpiry", label)
}

// UnpackFindExpiry is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6f537c72.
//
// Solidity: function findExpiry(string label) view returns(uint64)
func (wrapperRegistry *WrapperRegistry) UnpackFindExpiry(data []byte) (uint64, error) {
	out, err := wrapperRegistry.abi.Unpack("findExpiry", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, nil
}

// PackFindOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x63560a8e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function findOwner(string label) view returns(address)
func (wrapperRegistry *WrapperRegistry) PackFindOwner(label string) []byte {
	enc, err := wrapperRegistry.abi.Pack("findOwner", label)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFindOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x63560a8e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function findOwner(string label) view returns(address)
func (wrapperRegistry *WrapperRegistry) TryPackFindOwner(label string) ([]byte, error) {
	return wrapperRegistry.abi.Pack("findOwner", label)
}

// UnpackFindOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x63560a8e.
//
// Solidity: function findOwner(string label) view returns(address)
func (wrapperRegistry *WrapperRegistry) UnpackFindOwner(data []byte) (common.Address, error) {
	out, err := wrapperRegistry.abi.Unpack("findOwner", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackFindTokenId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x91b3c037.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function findTokenId(string label) view returns(uint256)
func (wrapperRegistry *WrapperRegistry) PackFindTokenId(label string) []byte {
	enc, err := wrapperRegistry.abi.Pack("findTokenId", label)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFindTokenId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x91b3c037.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function findTokenId(string label) view returns(uint256)
func (wrapperRegistry *WrapperRegistry) TryPackFindTokenId(label string) ([]byte, error) {
	return wrapperRegistry.abi.Pack("findTokenId", label)
}

// UnpackFindTokenId is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x91b3c037.
//
// Solidity: function findTokenId(string label) view returns(uint256)
func (wrapperRegistry *WrapperRegistry) UnpackFindTokenId(data []byte) (*big.Int, error) {
	out, err := wrapperRegistry.abi.Unpack("findTokenId", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackFinishERC1155Migration is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5d05f049.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function finishERC1155Migration(uint256[] ids, (string,address,address,address)[] mds) returns()
func (wrapperRegistry *WrapperRegistry) PackFinishERC1155Migration(ids []*big.Int, mds []LibMigrationData) []byte {
	enc, err := wrapperRegistry.abi.Pack("finishERC1155Migration", ids, mds)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFinishERC1155Migration is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5d05f049.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function finishERC1155Migration(uint256[] ids, (string,address,address,address)[] mds) returns()
func (wrapperRegistry *WrapperRegistry) TryPackFinishERC1155Migration(ids []*big.Int, mds []LibMigrationData) ([]byte, error) {
	return wrapperRegistry.abi.Pack("finishERC1155Migration", ids, mds)
}

// PackGetAssigneeCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3634f911.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getAssigneeCount(uint256 anyId, uint256 roleBitmap) view returns(uint256 counts, uint256 mask)
func (wrapperRegistry *WrapperRegistry) PackGetAssigneeCount(anyId *big.Int, roleBitmap *big.Int) []byte {
	enc, err := wrapperRegistry.abi.Pack("getAssigneeCount", anyId, roleBitmap)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetAssigneeCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3634f911.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getAssigneeCount(uint256 anyId, uint256 roleBitmap) view returns(uint256 counts, uint256 mask)
func (wrapperRegistry *WrapperRegistry) TryPackGetAssigneeCount(anyId *big.Int, roleBitmap *big.Int) ([]byte, error) {
	return wrapperRegistry.abi.Pack("getAssigneeCount", anyId, roleBitmap)
}

// GetAssigneeCountOutput serves as a container for the return parameters of contract
// method GetAssigneeCount.
type GetAssigneeCountOutput struct {
	Counts *big.Int
	Mask   *big.Int
}

// UnpackGetAssigneeCount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3634f911.
//
// Solidity: function getAssigneeCount(uint256 anyId, uint256 roleBitmap) view returns(uint256 counts, uint256 mask)
func (wrapperRegistry *WrapperRegistry) UnpackGetAssigneeCount(data []byte) (GetAssigneeCountOutput, error) {
	out, err := wrapperRegistry.abi.Unpack("getAssigneeCount", data)
	outstruct := new(GetAssigneeCountOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Counts = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.Mask = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	return *outstruct, nil
}

// PackGetExpiry is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x13c72608.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getExpiry(uint256 anyId) view returns(uint64)
func (wrapperRegistry *WrapperRegistry) PackGetExpiry(anyId *big.Int) []byte {
	enc, err := wrapperRegistry.abi.Pack("getExpiry", anyId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetExpiry is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x13c72608.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getExpiry(uint256 anyId) view returns(uint64)
func (wrapperRegistry *WrapperRegistry) TryPackGetExpiry(anyId *big.Int) ([]byte, error) {
	return wrapperRegistry.abi.Pack("getExpiry", anyId)
}

// UnpackGetExpiry is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x13c72608.
//
// Solidity: function getExpiry(uint256 anyId) view returns(uint64)
func (wrapperRegistry *WrapperRegistry) UnpackGetExpiry(data []byte) (uint64, error) {
	out, err := wrapperRegistry.abi.Unpack("getExpiry", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, nil
}

// PackGetOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc41a360a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getOwner(uint256 anyId) view returns(address)
func (wrapperRegistry *WrapperRegistry) PackGetOwner(anyId *big.Int) []byte {
	enc, err := wrapperRegistry.abi.Pack("getOwner", anyId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc41a360a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getOwner(uint256 anyId) view returns(address)
func (wrapperRegistry *WrapperRegistry) TryPackGetOwner(anyId *big.Int) ([]byte, error) {
	return wrapperRegistry.abi.Pack("getOwner", anyId)
}

// UnpackGetOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc41a360a.
//
// Solidity: function getOwner(uint256 anyId) view returns(address)
func (wrapperRegistry *WrapperRegistry) UnpackGetOwner(data []byte) (common.Address, error) {
	out, err := wrapperRegistry.abi.Unpack("getOwner", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetParent is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x80f76021.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getParent() view returns(address parent, string label)
func (wrapperRegistry *WrapperRegistry) PackGetParent() []byte {
	enc, err := wrapperRegistry.abi.Pack("getParent")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetParent is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x80f76021.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getParent() view returns(address parent, string label)
func (wrapperRegistry *WrapperRegistry) TryPackGetParent() ([]byte, error) {
	return wrapperRegistry.abi.Pack("getParent")
}

// GetParentOutput serves as a container for the return parameters of contract
// method GetParent.
type GetParentOutput struct {
	Parent common.Address
	Label  string
}

// UnpackGetParent is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x80f76021.
//
// Solidity: function getParent() view returns(address parent, string label)
func (wrapperRegistry *WrapperRegistry) UnpackGetParent(data []byte) (GetParentOutput, error) {
	out, err := wrapperRegistry.abi.Unpack("getParent", data)
	outstruct := new(GetParentOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Parent = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Label = *abi.ConvertType(out[1], new(string)).(*string)
	return *outstruct, nil
}

// PackGetResolver is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe4ae7d77.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getResolver(string label) view returns(address)
func (wrapperRegistry *WrapperRegistry) PackGetResolver(label string) []byte {
	enc, err := wrapperRegistry.abi.Pack("getResolver", label)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetResolver is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe4ae7d77.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getResolver(string label) view returns(address)
func (wrapperRegistry *WrapperRegistry) TryPackGetResolver(label string) ([]byte, error) {
	return wrapperRegistry.abi.Pack("getResolver", label)
}

// UnpackGetResolver is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe4ae7d77.
//
// Solidity: function getResolver(string label) view returns(address)
func (wrapperRegistry *WrapperRegistry) UnpackGetResolver(data []byte) (common.Address, error) {
	out, err := wrapperRegistry.abi.Unpack("getResolver", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetResource is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1e8fca2d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getResource(uint256 anyId) view returns(uint256)
func (wrapperRegistry *WrapperRegistry) PackGetResource(anyId *big.Int) []byte {
	enc, err := wrapperRegistry.abi.Pack("getResource", anyId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetResource is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1e8fca2d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getResource(uint256 anyId) view returns(uint256)
func (wrapperRegistry *WrapperRegistry) TryPackGetResource(anyId *big.Int) ([]byte, error) {
	return wrapperRegistry.abi.Pack("getResource", anyId)
}

// UnpackGetResource is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x1e8fca2d.
//
// Solidity: function getResource(uint256 anyId) view returns(uint256)
func (wrapperRegistry *WrapperRegistry) UnpackGetResource(data []byte) (*big.Int, error) {
	out, err := wrapperRegistry.abi.Unpack("getResource", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetState is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x44c9af28.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getState(uint256 anyId) view returns((uint8,uint64,address,uint256,uint256) state)
func (wrapperRegistry *WrapperRegistry) PackGetState(anyId *big.Int) []byte {
	enc, err := wrapperRegistry.abi.Pack("getState", anyId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetState is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x44c9af28.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getState(uint256 anyId) view returns((uint8,uint64,address,uint256,uint256) state)
func (wrapperRegistry *WrapperRegistry) TryPackGetState(anyId *big.Int) ([]byte, error) {
	return wrapperRegistry.abi.Pack("getState", anyId)
}

// UnpackGetState is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x44c9af28.
//
// Solidity: function getState(uint256 anyId) view returns((uint8,uint64,address,uint256,uint256) state)
func (wrapperRegistry *WrapperRegistry) UnpackGetState(data []byte) (IPermissionedRegistryState, error) {
	out, err := wrapperRegistry.abi.Unpack("getState", data)
	if err != nil {
		return *new(IPermissionedRegistryState), err
	}
	out0 := *abi.ConvertType(out[0], new(IPermissionedRegistryState)).(*IPermissionedRegistryState)
	return out0, nil
}

// PackGetStatus is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5c622a0e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getStatus(uint256 anyId) view returns(uint8)
func (wrapperRegistry *WrapperRegistry) PackGetStatus(anyId *big.Int) []byte {
	enc, err := wrapperRegistry.abi.Pack("getStatus", anyId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetStatus is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5c622a0e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getStatus(uint256 anyId) view returns(uint8)
func (wrapperRegistry *WrapperRegistry) TryPackGetStatus(anyId *big.Int) ([]byte, error) {
	return wrapperRegistry.abi.Pack("getStatus", anyId)
}

// UnpackGetStatus is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5c622a0e.
//
// Solidity: function getStatus(uint256 anyId) view returns(uint8)
func (wrapperRegistry *WrapperRegistry) UnpackGetStatus(data []byte) (uint8, error) {
	out, err := wrapperRegistry.abi.Unpack("getStatus", data)
	if err != nil {
		return *new(uint8), err
	}
	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	return out0, nil
}

// PackGetSubregistry is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x35af6216.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getSubregistry(string label) view returns(address)
func (wrapperRegistry *WrapperRegistry) PackGetSubregistry(label string) []byte {
	enc, err := wrapperRegistry.abi.Pack("getSubregistry", label)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetSubregistry is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x35af6216.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getSubregistry(string label) view returns(address)
func (wrapperRegistry *WrapperRegistry) TryPackGetSubregistry(label string) ([]byte, error) {
	return wrapperRegistry.abi.Pack("getSubregistry", label)
}

// UnpackGetSubregistry is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x35af6216.
//
// Solidity: function getSubregistry(string label) view returns(address)
func (wrapperRegistry *WrapperRegistry) UnpackGetSubregistry(data []byte) (common.Address, error) {
	out, err := wrapperRegistry.abi.Unpack("getSubregistry", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetTokenId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x14ff5ea3.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getTokenId(uint256 anyId) view returns(uint256)
func (wrapperRegistry *WrapperRegistry) PackGetTokenId(anyId *big.Int) []byte {
	enc, err := wrapperRegistry.abi.Pack("getTokenId", anyId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetTokenId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x14ff5ea3.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getTokenId(uint256 anyId) view returns(uint256)
func (wrapperRegistry *WrapperRegistry) TryPackGetTokenId(anyId *big.Int) ([]byte, error) {
	return wrapperRegistry.abi.Pack("getTokenId", anyId)
}

// UnpackGetTokenId is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x14ff5ea3.
//
// Solidity: function getTokenId(uint256 anyId) view returns(uint256)
func (wrapperRegistry *WrapperRegistry) UnpackGetTokenId(data []byte) (*big.Int, error) {
	out, err := wrapperRegistry.abi.Unpack("getTokenId", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetURI is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7754305c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getURI() view returns(string uri_, address renderer)
func (wrapperRegistry *WrapperRegistry) PackGetURI() []byte {
	enc, err := wrapperRegistry.abi.Pack("getURI")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetURI is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7754305c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getURI() view returns(string uri_, address renderer)
func (wrapperRegistry *WrapperRegistry) TryPackGetURI() ([]byte, error) {
	return wrapperRegistry.abi.Pack("getURI")
}

// GetURIOutput serves as a container for the return parameters of contract
// method GetURI.
type GetURIOutput struct {
	Uri      string
	Renderer common.Address
}

// UnpackGetURI is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7754305c.
//
// Solidity: function getURI() view returns(string uri_, address renderer)
func (wrapperRegistry *WrapperRegistry) UnpackGetURI(data []byte) (GetURIOutput, error) {
	out, err := wrapperRegistry.abi.Unpack("getURI", data)
	outstruct := new(GetURIOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Uri = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Renderer = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	return *outstruct, nil
}

// PackGetWrappedName is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9b224b1d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getWrappedName() view returns(bytes)
func (wrapperRegistry *WrapperRegistry) PackGetWrappedName() []byte {
	enc, err := wrapperRegistry.abi.Pack("getWrappedName")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetWrappedName is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9b224b1d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getWrappedName() view returns(bytes)
func (wrapperRegistry *WrapperRegistry) TryPackGetWrappedName() ([]byte, error) {
	return wrapperRegistry.abi.Pack("getWrappedName")
}

// UnpackGetWrappedName is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9b224b1d.
//
// Solidity: function getWrappedName() view returns(bytes)
func (wrapperRegistry *WrapperRegistry) UnpackGetWrappedName(data []byte) ([]byte, error) {
	out, err := wrapperRegistry.abi.Unpack("getWrappedName", data)
	if err != nil {
		return *new([]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	return out0, nil
}

// PackGetWrappedNode is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6e7a2116.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getWrappedNode() view returns(bytes32)
func (wrapperRegistry *WrapperRegistry) PackGetWrappedNode() []byte {
	enc, err := wrapperRegistry.abi.Pack("getWrappedNode")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetWrappedNode is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6e7a2116.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getWrappedNode() view returns(bytes32)
func (wrapperRegistry *WrapperRegistry) TryPackGetWrappedNode() ([]byte, error) {
	return wrapperRegistry.abi.Pack("getWrappedNode")
}

// UnpackGetWrappedNode is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6e7a2116.
//
// Solidity: function getWrappedNode() view returns(bytes32)
func (wrapperRegistry *WrapperRegistry) UnpackGetWrappedNode(data []byte) ([32]byte, error) {
	out, err := wrapperRegistry.abi.Unpack("getWrappedNode", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackGrantRoles is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7c300586.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function grantRoles(uint256 anyId, uint256 roleBitmap, address account) returns(bool)
func (wrapperRegistry *WrapperRegistry) PackGrantRoles(anyId *big.Int, roleBitmap *big.Int, account common.Address) []byte {
	enc, err := wrapperRegistry.abi.Pack("grantRoles", anyId, roleBitmap, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGrantRoles is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7c300586.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function grantRoles(uint256 anyId, uint256 roleBitmap, address account) returns(bool)
func (wrapperRegistry *WrapperRegistry) TryPackGrantRoles(anyId *big.Int, roleBitmap *big.Int, account common.Address) ([]byte, error) {
	return wrapperRegistry.abi.Pack("grantRoles", anyId, roleBitmap, account)
}

// UnpackGrantRoles is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7c300586.
//
// Solidity: function grantRoles(uint256 anyId, uint256 roleBitmap, address account) returns(bool)
func (wrapperRegistry *WrapperRegistry) UnpackGrantRoles(data []byte) (bool, error) {
	out, err := wrapperRegistry.abi.Unpack("grantRoles", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackGrantRootRoles is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x072d5d77.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function grantRootRoles(uint256 roleBitmap, address account) returns(bool)
func (wrapperRegistry *WrapperRegistry) PackGrantRootRoles(roleBitmap *big.Int, account common.Address) []byte {
	enc, err := wrapperRegistry.abi.Pack("grantRootRoles", roleBitmap, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGrantRootRoles is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x072d5d77.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function grantRootRoles(uint256 roleBitmap, address account) returns(bool)
func (wrapperRegistry *WrapperRegistry) TryPackGrantRootRoles(roleBitmap *big.Int, account common.Address) ([]byte, error) {
	return wrapperRegistry.abi.Pack("grantRootRoles", roleBitmap, account)
}

// UnpackGrantRootRoles is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x072d5d77.
//
// Solidity: function grantRootRoles(uint256 roleBitmap, address account) returns(bool)
func (wrapperRegistry *WrapperRegistry) UnpackGrantRootRoles(data []byte) (bool, error) {
	out, err := wrapperRegistry.abi.Unpack("grantRootRoles", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackHasAssignees is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x11b8e00a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function hasAssignees(uint256 resource, uint256 roleBitmap) view returns(bool)
func (wrapperRegistry *WrapperRegistry) PackHasAssignees(resource *big.Int, roleBitmap *big.Int) []byte {
	enc, err := wrapperRegistry.abi.Pack("hasAssignees", resource, roleBitmap)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackHasAssignees is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x11b8e00a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function hasAssignees(uint256 resource, uint256 roleBitmap) view returns(bool)
func (wrapperRegistry *WrapperRegistry) TryPackHasAssignees(resource *big.Int, roleBitmap *big.Int) ([]byte, error) {
	return wrapperRegistry.abi.Pack("hasAssignees", resource, roleBitmap)
}

// UnpackHasAssignees is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x11b8e00a.
//
// Solidity: function hasAssignees(uint256 resource, uint256 roleBitmap) view returns(bool)
func (wrapperRegistry *WrapperRegistry) UnpackHasAssignees(data []byte) (bool, error) {
	out, err := wrapperRegistry.abi.Unpack("hasAssignees", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackHasRoles is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd3bf89b1.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function hasRoles(uint256 anyId, uint256 roleBitmap, address account) view returns(bool)
func (wrapperRegistry *WrapperRegistry) PackHasRoles(anyId *big.Int, roleBitmap *big.Int, account common.Address) []byte {
	enc, err := wrapperRegistry.abi.Pack("hasRoles", anyId, roleBitmap, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackHasRoles is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd3bf89b1.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function hasRoles(uint256 anyId, uint256 roleBitmap, address account) view returns(bool)
func (wrapperRegistry *WrapperRegistry) TryPackHasRoles(anyId *big.Int, roleBitmap *big.Int, account common.Address) ([]byte, error) {
	return wrapperRegistry.abi.Pack("hasRoles", anyId, roleBitmap, account)
}

// UnpackHasRoles is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd3bf89b1.
//
// Solidity: function hasRoles(uint256 anyId, uint256 roleBitmap, address account) view returns(bool)
func (wrapperRegistry *WrapperRegistry) UnpackHasRoles(data []byte) (bool, error) {
	out, err := wrapperRegistry.abi.Unpack("hasRoles", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackHasRootRoles is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x781ef8db.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function hasRootRoles(uint256 roleBitmap, address account) view returns(bool)
func (wrapperRegistry *WrapperRegistry) PackHasRootRoles(roleBitmap *big.Int, account common.Address) []byte {
	enc, err := wrapperRegistry.abi.Pack("hasRootRoles", roleBitmap, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackHasRootRoles is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x781ef8db.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function hasRootRoles(uint256 roleBitmap, address account) view returns(bool)
func (wrapperRegistry *WrapperRegistry) TryPackHasRootRoles(roleBitmap *big.Int, account common.Address) ([]byte, error) {
	return wrapperRegistry.abi.Pack("hasRootRoles", roleBitmap, account)
}

// UnpackHasRootRoles is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x781ef8db.
//
// Solidity: function hasRootRoles(uint256 roleBitmap, address account) view returns(bool)
func (wrapperRegistry *WrapperRegistry) UnpackHasRootRoles(data []byte) (bool, error) {
	out, err := wrapperRegistry.abi.Unpack("hasRootRoles", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1542c01a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function initialize(bytes32 node, address parentRegistry, string childLabel, uint256 roleBitmap) returns()
func (wrapperRegistry *WrapperRegistry) PackInitialize(node [32]byte, parentRegistry common.Address, childLabel string, roleBitmap *big.Int) []byte {
	enc, err := wrapperRegistry.abi.Pack("initialize", node, parentRegistry, childLabel, roleBitmap)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1542c01a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function initialize(bytes32 node, address parentRegistry, string childLabel, uint256 roleBitmap) returns()
func (wrapperRegistry *WrapperRegistry) TryPackInitialize(node [32]byte, parentRegistry common.Address, childLabel string, roleBitmap *big.Int) ([]byte, error) {
	return wrapperRegistry.abi.Pack("initialize", node, parentRegistry, childLabel, roleBitmap)
}

// PackIsApprovedForAll is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe985e9c5.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (wrapperRegistry *WrapperRegistry) PackIsApprovedForAll(account common.Address, operator common.Address) []byte {
	enc, err := wrapperRegistry.abi.Pack("isApprovedForAll", account, operator)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsApprovedForAll is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe985e9c5.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (wrapperRegistry *WrapperRegistry) TryPackIsApprovedForAll(account common.Address, operator common.Address) ([]byte, error) {
	return wrapperRegistry.abi.Pack("isApprovedForAll", account, operator)
}

// UnpackIsApprovedForAll is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (wrapperRegistry *WrapperRegistry) UnpackIsApprovedForAll(data []byte) (bool, error) {
	out, err := wrapperRegistry.abi.Unpack("isApprovedForAll", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackIsContractNamer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6f3ff726.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isContractNamer(address namer) view returns(bool)
func (wrapperRegistry *WrapperRegistry) PackIsContractNamer(namer common.Address) []byte {
	enc, err := wrapperRegistry.abi.Pack("isContractNamer", namer)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsContractNamer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6f3ff726.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isContractNamer(address namer) view returns(bool)
func (wrapperRegistry *WrapperRegistry) TryPackIsContractNamer(namer common.Address) ([]byte, error) {
	return wrapperRegistry.abi.Pack("isContractNamer", namer)
}

// UnpackIsContractNamer is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6f3ff726.
//
// Solidity: function isContractNamer(address namer) view returns(bool)
func (wrapperRegistry *WrapperRegistry) UnpackIsContractNamer(data []byte) (bool, error) {
	out, err := wrapperRegistry.abi.Unpack("isContractNamer", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackIsEmancipated is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x955266f5.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isEmancipated() pure returns(bool)
func (wrapperRegistry *WrapperRegistry) PackIsEmancipated() []byte {
	enc, err := wrapperRegistry.abi.Pack("isEmancipated")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsEmancipated is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x955266f5.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isEmancipated() pure returns(bool)
func (wrapperRegistry *WrapperRegistry) TryPackIsEmancipated() ([]byte, error) {
	return wrapperRegistry.abi.Pack("isEmancipated")
}

// UnpackIsEmancipated is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x955266f5.
//
// Solidity: function isEmancipated() pure returns(bool)
func (wrapperRegistry *WrapperRegistry) UnpackIsEmancipated(data []byte) (bool, error) {
	out, err := wrapperRegistry.abi.Unpack("isEmancipated", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackIsOnlyAssignee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8e77c95f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isOnlyAssignee(uint256 resource, uint256 roleBitmap, address account) view returns(bool)
func (wrapperRegistry *WrapperRegistry) PackIsOnlyAssignee(resource *big.Int, roleBitmap *big.Int, account common.Address) []byte {
	enc, err := wrapperRegistry.abi.Pack("isOnlyAssignee", resource, roleBitmap, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsOnlyAssignee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8e77c95f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isOnlyAssignee(uint256 resource, uint256 roleBitmap, address account) view returns(bool)
func (wrapperRegistry *WrapperRegistry) TryPackIsOnlyAssignee(resource *big.Int, roleBitmap *big.Int, account common.Address) ([]byte, error) {
	return wrapperRegistry.abi.Pack("isOnlyAssignee", resource, roleBitmap, account)
}

// UnpackIsOnlyAssignee is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8e77c95f.
//
// Solidity: function isOnlyAssignee(uint256 resource, uint256 roleBitmap, address account) view returns(bool)
func (wrapperRegistry *WrapperRegistry) UnpackIsOnlyAssignee(data []byte) (bool, error) {
	out, err := wrapperRegistry.abi.Unpack("isOnlyAssignee", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackLatestOwnerOf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbd242bcb.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function latestOwnerOf(uint256 tokenId) view returns(address)
func (wrapperRegistry *WrapperRegistry) PackLatestOwnerOf(tokenId *big.Int) []byte {
	enc, err := wrapperRegistry.abi.Pack("latestOwnerOf", tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackLatestOwnerOf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbd242bcb.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function latestOwnerOf(uint256 tokenId) view returns(address)
func (wrapperRegistry *WrapperRegistry) TryPackLatestOwnerOf(tokenId *big.Int) ([]byte, error) {
	return wrapperRegistry.abi.Pack("latestOwnerOf", tokenId)
}

// UnpackLatestOwnerOf is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbd242bcb.
//
// Solidity: function latestOwnerOf(uint256 tokenId) view returns(address)
func (wrapperRegistry *WrapperRegistry) UnpackLatestOwnerOf(data []byte) (common.Address, error) {
	out, err := wrapperRegistry.abi.Unpack("latestOwnerOf", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackOnERC1155BatchReceived is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbc197c81.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] ids, uint256[] , bytes data) returns(bytes4)
func (wrapperRegistry *WrapperRegistry) PackOnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, ids []*big.Int, arg3 []*big.Int, data []byte) []byte {
	enc, err := wrapperRegistry.abi.Pack("onERC1155BatchReceived", arg0, arg1, ids, arg3, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackOnERC1155BatchReceived is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbc197c81.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] ids, uint256[] , bytes data) returns(bytes4)
func (wrapperRegistry *WrapperRegistry) TryPackOnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, ids []*big.Int, arg3 []*big.Int, data []byte) ([]byte, error) {
	return wrapperRegistry.abi.Pack("onERC1155BatchReceived", arg0, arg1, ids, arg3, data)
}

// UnpackOnERC1155BatchReceived is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] ids, uint256[] , bytes data) returns(bytes4)
func (wrapperRegistry *WrapperRegistry) UnpackOnERC1155BatchReceived(data []byte) ([4]byte, error) {
	out, err := wrapperRegistry.abi.Unpack("onERC1155BatchReceived", data)
	if err != nil {
		return *new([4]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)
	return out0, nil
}

// PackOnERC1155Received is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf23a6e61.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function onERC1155Received(address , address , uint256 id, uint256 , bytes data) returns(bytes4)
func (wrapperRegistry *WrapperRegistry) PackOnERC1155Received(arg0 common.Address, arg1 common.Address, id *big.Int, arg3 *big.Int, data []byte) []byte {
	enc, err := wrapperRegistry.abi.Pack("onERC1155Received", arg0, arg1, id, arg3, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackOnERC1155Received is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf23a6e61.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function onERC1155Received(address , address , uint256 id, uint256 , bytes data) returns(bytes4)
func (wrapperRegistry *WrapperRegistry) TryPackOnERC1155Received(arg0 common.Address, arg1 common.Address, id *big.Int, arg3 *big.Int, data []byte) ([]byte, error) {
	return wrapperRegistry.abi.Pack("onERC1155Received", arg0, arg1, id, arg3, data)
}

// UnpackOnERC1155Received is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 id, uint256 , bytes data) returns(bytes4)
func (wrapperRegistry *WrapperRegistry) UnpackOnERC1155Received(data []byte) ([4]byte, error) {
	out, err := wrapperRegistry.abi.Unpack("onERC1155Received", data)
	if err != nil {
		return *new([4]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)
	return out0, nil
}

// PackOwnerOf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6352211e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (wrapperRegistry *WrapperRegistry) PackOwnerOf(tokenId *big.Int) []byte {
	enc, err := wrapperRegistry.abi.Pack("ownerOf", tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackOwnerOf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6352211e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (wrapperRegistry *WrapperRegistry) TryPackOwnerOf(tokenId *big.Int) ([]byte, error) {
	return wrapperRegistry.abi.Pack("ownerOf", tokenId)
}

// UnpackOwnerOf is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (wrapperRegistry *WrapperRegistry) UnpackOwnerOf(data []byte) (common.Address, error) {
	out, err := wrapperRegistry.abi.Unpack("ownerOf", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackProxiableUUID is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x52d1902d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (wrapperRegistry *WrapperRegistry) PackProxiableUUID() []byte {
	enc, err := wrapperRegistry.abi.Pack("proxiableUUID")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackProxiableUUID is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x52d1902d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (wrapperRegistry *WrapperRegistry) TryPackProxiableUUID() ([]byte, error) {
	return wrapperRegistry.abi.Pack("proxiableUUID")
}

// UnpackProxiableUUID is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (wrapperRegistry *WrapperRegistry) UnpackProxiableUUID(data []byte) ([32]byte, error) {
	out, err := wrapperRegistry.abi.Unpack("proxiableUUID", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackRegister is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x85f3e643.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function register(string label, address owner, address registry, address resolver, uint256 roleBitmap, uint64 expiry) returns(uint256 tokenId)
func (wrapperRegistry *WrapperRegistry) PackRegister(label string, owner common.Address, registry common.Address, resolver common.Address, roleBitmap *big.Int, expiry uint64) []byte {
	enc, err := wrapperRegistry.abi.Pack("register", label, owner, registry, resolver, roleBitmap, expiry)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRegister is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x85f3e643.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function register(string label, address owner, address registry, address resolver, uint256 roleBitmap, uint64 expiry) returns(uint256 tokenId)
func (wrapperRegistry *WrapperRegistry) TryPackRegister(label string, owner common.Address, registry common.Address, resolver common.Address, roleBitmap *big.Int, expiry uint64) ([]byte, error) {
	return wrapperRegistry.abi.Pack("register", label, owner, registry, resolver, roleBitmap, expiry)
}

// UnpackRegister is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x85f3e643.
//
// Solidity: function register(string label, address owner, address registry, address resolver, uint256 roleBitmap, uint64 expiry) returns(uint256 tokenId)
func (wrapperRegistry *WrapperRegistry) UnpackRegister(data []byte) (*big.Int, error) {
	out, err := wrapperRegistry.abi.Unpack("register", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackRenew is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5569f33d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function renew(uint256 anyId, uint64 newExpiry) returns()
func (wrapperRegistry *WrapperRegistry) PackRenew(anyId *big.Int, newExpiry uint64) []byte {
	enc, err := wrapperRegistry.abi.Pack("renew", anyId, newExpiry)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRenew is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5569f33d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function renew(uint256 anyId, uint64 newExpiry) returns()
func (wrapperRegistry *WrapperRegistry) TryPackRenew(anyId *big.Int, newExpiry uint64) ([]byte, error) {
	return wrapperRegistry.abi.Pack("renew", anyId, newExpiry)
}

// PackRevokeRoles is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdfa70d8b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function revokeRoles(uint256 anyId, uint256 roleBitmap, address account) returns(bool)
func (wrapperRegistry *WrapperRegistry) PackRevokeRoles(anyId *big.Int, roleBitmap *big.Int, account common.Address) []byte {
	enc, err := wrapperRegistry.abi.Pack("revokeRoles", anyId, roleBitmap, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRevokeRoles is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdfa70d8b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function revokeRoles(uint256 anyId, uint256 roleBitmap, address account) returns(bool)
func (wrapperRegistry *WrapperRegistry) TryPackRevokeRoles(anyId *big.Int, roleBitmap *big.Int, account common.Address) ([]byte, error) {
	return wrapperRegistry.abi.Pack("revokeRoles", anyId, roleBitmap, account)
}

// UnpackRevokeRoles is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xdfa70d8b.
//
// Solidity: function revokeRoles(uint256 anyId, uint256 roleBitmap, address account) returns(bool)
func (wrapperRegistry *WrapperRegistry) UnpackRevokeRoles(data []byte) (bool, error) {
	out, err := wrapperRegistry.abi.Unpack("revokeRoles", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackRevokeRootRoles is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xce156e82.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function revokeRootRoles(uint256 roleBitmap, address account) returns(bool)
func (wrapperRegistry *WrapperRegistry) PackRevokeRootRoles(roleBitmap *big.Int, account common.Address) []byte {
	enc, err := wrapperRegistry.abi.Pack("revokeRootRoles", roleBitmap, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRevokeRootRoles is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xce156e82.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function revokeRootRoles(uint256 roleBitmap, address account) returns(bool)
func (wrapperRegistry *WrapperRegistry) TryPackRevokeRootRoles(roleBitmap *big.Int, account common.Address) ([]byte, error) {
	return wrapperRegistry.abi.Pack("revokeRootRoles", roleBitmap, account)
}

// UnpackRevokeRootRoles is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xce156e82.
//
// Solidity: function revokeRootRoles(uint256 roleBitmap, address account) returns(bool)
func (wrapperRegistry *WrapperRegistry) UnpackRevokeRootRoles(data []byte) (bool, error) {
	out, err := wrapperRegistry.abi.Unpack("revokeRootRoles", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackRoleCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2f27fa24.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function roleCount(uint256 anyId) view returns(uint256)
func (wrapperRegistry *WrapperRegistry) PackRoleCount(anyId *big.Int) []byte {
	enc, err := wrapperRegistry.abi.Pack("roleCount", anyId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRoleCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2f27fa24.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function roleCount(uint256 anyId) view returns(uint256)
func (wrapperRegistry *WrapperRegistry) TryPackRoleCount(anyId *big.Int) ([]byte, error) {
	return wrapperRegistry.abi.Pack("roleCount", anyId)
}

// UnpackRoleCount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x2f27fa24.
//
// Solidity: function roleCount(uint256 anyId) view returns(uint256)
func (wrapperRegistry *WrapperRegistry) UnpackRoleCount(data []byte) (*big.Int, error) {
	out, err := wrapperRegistry.abi.Unpack("roleCount", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackRoles is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5adf4724.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function roles(uint256 anyId, address account) view returns(uint256)
func (wrapperRegistry *WrapperRegistry) PackRoles(anyId *big.Int, account common.Address) []byte {
	enc, err := wrapperRegistry.abi.Pack("roles", anyId, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRoles is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5adf4724.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function roles(uint256 anyId, address account) view returns(uint256)
func (wrapperRegistry *WrapperRegistry) TryPackRoles(anyId *big.Int, account common.Address) ([]byte, error) {
	return wrapperRegistry.abi.Pack("roles", anyId, account)
}

// UnpackRoles is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5adf4724.
//
// Solidity: function roles(uint256 anyId, address account) view returns(uint256)
func (wrapperRegistry *WrapperRegistry) UnpackRoles(data []byte) (*big.Int, error) {
	out, err := wrapperRegistry.abi.Unpack("roles", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackSafeBatchTransferFrom is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2eb2c2d6.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] values, bytes data) returns()
func (wrapperRegistry *WrapperRegistry) PackSafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, values []*big.Int, data []byte) []byte {
	enc, err := wrapperRegistry.abi.Pack("safeBatchTransferFrom", from, to, ids, values, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSafeBatchTransferFrom is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2eb2c2d6.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] values, bytes data) returns()
func (wrapperRegistry *WrapperRegistry) TryPackSafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, values []*big.Int, data []byte) ([]byte, error) {
	return wrapperRegistry.abi.Pack("safeBatchTransferFrom", from, to, ids, values, data)
}

// PackSafeTransferFrom is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf242432a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 value, bytes data) returns()
func (wrapperRegistry *WrapperRegistry) PackSafeTransferFrom(from common.Address, to common.Address, id *big.Int, value *big.Int, data []byte) []byte {
	enc, err := wrapperRegistry.abi.Pack("safeTransferFrom", from, to, id, value, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSafeTransferFrom is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf242432a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 value, bytes data) returns()
func (wrapperRegistry *WrapperRegistry) TryPackSafeTransferFrom(from common.Address, to common.Address, id *big.Int, value *big.Int, data []byte) ([]byte, error) {
	return wrapperRegistry.abi.Pack("safeTransferFrom", from, to, id, value, data)
}

// PackSetApprovalForAll is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa22cb465.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (wrapperRegistry *WrapperRegistry) PackSetApprovalForAll(operator common.Address, approved bool) []byte {
	enc, err := wrapperRegistry.abi.Pack("setApprovalForAll", operator, approved)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetApprovalForAll is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa22cb465.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (wrapperRegistry *WrapperRegistry) TryPackSetApprovalForAll(operator common.Address, approved bool) ([]byte, error) {
	return wrapperRegistry.abi.Pack("setApprovalForAll", operator, approved)
}

// PackSetParent is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5357263f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setParent(address parent, string label) returns()
func (wrapperRegistry *WrapperRegistry) PackSetParent(parent common.Address, label string) []byte {
	enc, err := wrapperRegistry.abi.Pack("setParent", parent, label)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetParent is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5357263f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setParent(address parent, string label) returns()
func (wrapperRegistry *WrapperRegistry) TryPackSetParent(parent common.Address, label string) ([]byte, error) {
	return wrapperRegistry.abi.Pack("setParent", parent, label)
}

// PackSetResolver is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbc7b6d62.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setResolver(uint256 anyId, address resolver) returns()
func (wrapperRegistry *WrapperRegistry) PackSetResolver(anyId *big.Int, resolver common.Address) []byte {
	enc, err := wrapperRegistry.abi.Pack("setResolver", anyId, resolver)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetResolver is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbc7b6d62.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setResolver(uint256 anyId, address resolver) returns()
func (wrapperRegistry *WrapperRegistry) TryPackSetResolver(anyId *big.Int, resolver common.Address) ([]byte, error) {
	return wrapperRegistry.abi.Pack("setResolver", anyId, resolver)
}

// PackSetSubregistry is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x341ec559.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setSubregistry(uint256 anyId, address registry) returns()
func (wrapperRegistry *WrapperRegistry) PackSetSubregistry(anyId *big.Int, registry common.Address) []byte {
	enc, err := wrapperRegistry.abi.Pack("setSubregistry", anyId, registry)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetSubregistry is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x341ec559.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setSubregistry(uint256 anyId, address registry) returns()
func (wrapperRegistry *WrapperRegistry) TryPackSetSubregistry(anyId *big.Int, registry common.Address) ([]byte, error) {
	return wrapperRegistry.abi.Pack("setSubregistry", anyId, registry)
}

// PackSetURI is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x48688f95.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setURI(string uri_, address renderer) returns()
func (wrapperRegistry *WrapperRegistry) PackSetURI(uri string, renderer common.Address) []byte {
	enc, err := wrapperRegistry.abi.Pack("setURI", uri, renderer)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetURI is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x48688f95.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setURI(string uri_, address renderer) returns()
func (wrapperRegistry *WrapperRegistry) TryPackSetURI(uri string, renderer common.Address) ([]byte, error) {
	return wrapperRegistry.abi.Pack("setURI", uri, renderer)
}

// PackSupportsInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01ffc9a7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (wrapperRegistry *WrapperRegistry) PackSupportsInterface(interfaceId [4]byte) []byte {
	enc, err := wrapperRegistry.abi.Pack("supportsInterface", interfaceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSupportsInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01ffc9a7.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (wrapperRegistry *WrapperRegistry) TryPackSupportsInterface(interfaceId [4]byte) ([]byte, error) {
	return wrapperRegistry.abi.Pack("supportsInterface", interfaceId)
}

// UnpackSupportsInterface is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (wrapperRegistry *WrapperRegistry) UnpackSupportsInterface(data []byte) (bool, error) {
	out, err := wrapperRegistry.abi.Unpack("supportsInterface", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackUnregister is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa02b161e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function unregister(uint256 anyId) returns()
func (wrapperRegistry *WrapperRegistry) PackUnregister(anyId *big.Int) []byte {
	enc, err := wrapperRegistry.abi.Pack("unregister", anyId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUnregister is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa02b161e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function unregister(uint256 anyId) returns()
func (wrapperRegistry *WrapperRegistry) TryPackUnregister(anyId *big.Int) ([]byte, error) {
	return wrapperRegistry.abi.Pack("unregister", anyId)
}

// PackUnsafeBatchTransfer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3dd5ee86.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function unsafeBatchTransfer(address to, uint256[] tokenIds, bytes data) returns()
func (wrapperRegistry *WrapperRegistry) PackUnsafeBatchTransfer(to common.Address, tokenIds []*big.Int, data []byte) []byte {
	enc, err := wrapperRegistry.abi.Pack("unsafeBatchTransfer", to, tokenIds, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUnsafeBatchTransfer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3dd5ee86.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function unsafeBatchTransfer(address to, uint256[] tokenIds, bytes data) returns()
func (wrapperRegistry *WrapperRegistry) TryPackUnsafeBatchTransfer(to common.Address, tokenIds []*big.Int, data []byte) ([]byte, error) {
	return wrapperRegistry.abi.Pack("unsafeBatchTransfer", to, tokenIds, data)
}

// PackUnsafeTransfer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x087b0790.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function unsafeTransfer(address to, uint256 tokenId, bytes data) returns()
func (wrapperRegistry *WrapperRegistry) PackUnsafeTransfer(to common.Address, tokenId *big.Int, data []byte) []byte {
	enc, err := wrapperRegistry.abi.Pack("unsafeTransfer", to, tokenId, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUnsafeTransfer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x087b0790.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function unsafeTransfer(address to, uint256 tokenId, bytes data) returns()
func (wrapperRegistry *WrapperRegistry) TryPackUnsafeTransfer(to common.Address, tokenId *big.Int, data []byte) ([]byte, error) {
	return wrapperRegistry.abi.Pack("unsafeTransfer", to, tokenId, data)
}

// PackUpgradeToAndCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4f1ef286.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (wrapperRegistry *WrapperRegistry) PackUpgradeToAndCall(newImplementation common.Address, data []byte) []byte {
	enc, err := wrapperRegistry.abi.Pack("upgradeToAndCall", newImplementation, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUpgradeToAndCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4f1ef286.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (wrapperRegistry *WrapperRegistry) TryPackUpgradeToAndCall(newImplementation common.Address, data []byte) ([]byte, error) {
	return wrapperRegistry.abi.Pack("upgradeToAndCall", newImplementation, data)
}

// PackUri is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0e89341c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function uri(uint256 tokenId) view returns(string)
func (wrapperRegistry *WrapperRegistry) PackUri(tokenId *big.Int) []byte {
	enc, err := wrapperRegistry.abi.Pack("uri", tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUri is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0e89341c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function uri(uint256 tokenId) view returns(string)
func (wrapperRegistry *WrapperRegistry) TryPackUri(tokenId *big.Int) ([]byte, error) {
	return wrapperRegistry.abi.Pack("uri", tokenId)
}

// UnpackUri is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x0e89341c.
//
// Solidity: function uri(uint256 tokenId) view returns(string)
func (wrapperRegistry *WrapperRegistry) UnpackUri(data []byte) (string, error) {
	out, err := wrapperRegistry.abi.Unpack("uri", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, nil
}

// WrapperRegistryApprovalForAll represents a ApprovalForAll event raised by the WrapperRegistry contract.
type WrapperRegistryApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      *types.Log // Blockchain specific contextual infos
}

const WrapperRegistryApprovalForAllEventName = "ApprovalForAll"

// ContractEventName returns the user-defined event name.
func (WrapperRegistryApprovalForAll) ContractEventName() string {
	return WrapperRegistryApprovalForAllEventName
}

// UnpackApprovalForAllEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (wrapperRegistry *WrapperRegistry) UnpackApprovalForAllEvent(log *types.Log) (*WrapperRegistryApprovalForAll, error) {
	event := "ApprovalForAll"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != wrapperRegistry.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(WrapperRegistryApprovalForAll)
	if len(log.Data) > 0 {
		if err := wrapperRegistry.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range wrapperRegistry.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// WrapperRegistryEACRolesChanged represents a EACRolesChanged event raised by the WrapperRegistry contract.
type WrapperRegistryEACRolesChanged struct {
	Resource      *big.Int
	Account       common.Address
	OldRoleBitmap *big.Int
	NewRoleBitmap *big.Int
	Raw           *types.Log // Blockchain specific contextual infos
}

const WrapperRegistryEACRolesChangedEventName = "EACRolesChanged"

// ContractEventName returns the user-defined event name.
func (WrapperRegistryEACRolesChanged) ContractEventName() string {
	return WrapperRegistryEACRolesChangedEventName
}

// UnpackEACRolesChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event EACRolesChanged(uint256 indexed resource, address indexed account, uint256 oldRoleBitmap, uint256 newRoleBitmap)
func (wrapperRegistry *WrapperRegistry) UnpackEACRolesChangedEvent(log *types.Log) (*WrapperRegistryEACRolesChanged, error) {
	event := "EACRolesChanged"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != wrapperRegistry.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(WrapperRegistryEACRolesChanged)
	if len(log.Data) > 0 {
		if err := wrapperRegistry.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range wrapperRegistry.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// WrapperRegistryExpiryUpdated represents a ExpiryUpdated event raised by the WrapperRegistry contract.
type WrapperRegistryExpiryUpdated struct {
	TokenId   *big.Int
	NewExpiry uint64
	Sender    common.Address
	Raw       *types.Log // Blockchain specific contextual infos
}

const WrapperRegistryExpiryUpdatedEventName = "ExpiryUpdated"

// ContractEventName returns the user-defined event name.
func (WrapperRegistryExpiryUpdated) ContractEventName() string {
	return WrapperRegistryExpiryUpdatedEventName
}

// UnpackExpiryUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ExpiryUpdated(uint256 indexed tokenId, uint64 indexed newExpiry, address indexed sender)
func (wrapperRegistry *WrapperRegistry) UnpackExpiryUpdatedEvent(log *types.Log) (*WrapperRegistryExpiryUpdated, error) {
	event := "ExpiryUpdated"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != wrapperRegistry.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(WrapperRegistryExpiryUpdated)
	if len(log.Data) > 0 {
		if err := wrapperRegistry.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range wrapperRegistry.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// WrapperRegistryInitialized represents a Initialized event raised by the WrapperRegistry contract.
type WrapperRegistryInitialized struct {
	Version uint64
	Raw     *types.Log // Blockchain specific contextual infos
}

const WrapperRegistryInitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (WrapperRegistryInitialized) ContractEventName() string {
	return WrapperRegistryInitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint64 version)
func (wrapperRegistry *WrapperRegistry) UnpackInitializedEvent(log *types.Log) (*WrapperRegistryInitialized, error) {
	event := "Initialized"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != wrapperRegistry.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(WrapperRegistryInitialized)
	if len(log.Data) > 0 {
		if err := wrapperRegistry.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range wrapperRegistry.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// WrapperRegistryLabelRegistered represents a LabelRegistered event raised by the WrapperRegistry contract.
type WrapperRegistryLabelRegistered struct {
	TokenId   *big.Int
	LabelHash [32]byte
	Label     string
	Owner     common.Address
	Expiry    uint64
	Sender    common.Address
	Raw       *types.Log // Blockchain specific contextual infos
}

const WrapperRegistryLabelRegisteredEventName = "LabelRegistered"

// ContractEventName returns the user-defined event name.
func (WrapperRegistryLabelRegistered) ContractEventName() string {
	return WrapperRegistryLabelRegisteredEventName
}

// UnpackLabelRegisteredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event LabelRegistered(uint256 indexed tokenId, bytes32 indexed labelHash, string label, address owner, uint64 expiry, address indexed sender)
func (wrapperRegistry *WrapperRegistry) UnpackLabelRegisteredEvent(log *types.Log) (*WrapperRegistryLabelRegistered, error) {
	event := "LabelRegistered"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != wrapperRegistry.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(WrapperRegistryLabelRegistered)
	if len(log.Data) > 0 {
		if err := wrapperRegistry.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range wrapperRegistry.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// WrapperRegistryLabelReserved represents a LabelReserved event raised by the WrapperRegistry contract.
type WrapperRegistryLabelReserved struct {
	TokenId   *big.Int
	LabelHash [32]byte
	Label     string
	Expiry    uint64
	Sender    common.Address
	Raw       *types.Log // Blockchain specific contextual infos
}

const WrapperRegistryLabelReservedEventName = "LabelReserved"

// ContractEventName returns the user-defined event name.
func (WrapperRegistryLabelReserved) ContractEventName() string {
	return WrapperRegistryLabelReservedEventName
}

// UnpackLabelReservedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event LabelReserved(uint256 indexed tokenId, bytes32 indexed labelHash, string label, uint64 expiry, address indexed sender)
func (wrapperRegistry *WrapperRegistry) UnpackLabelReservedEvent(log *types.Log) (*WrapperRegistryLabelReserved, error) {
	event := "LabelReserved"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != wrapperRegistry.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(WrapperRegistryLabelReserved)
	if len(log.Data) > 0 {
		if err := wrapperRegistry.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range wrapperRegistry.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// WrapperRegistryLabelUnregistered represents a LabelUnregistered event raised by the WrapperRegistry contract.
type WrapperRegistryLabelUnregistered struct {
	TokenId *big.Int
	Sender  common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const WrapperRegistryLabelUnregisteredEventName = "LabelUnregistered"

// ContractEventName returns the user-defined event name.
func (WrapperRegistryLabelUnregistered) ContractEventName() string {
	return WrapperRegistryLabelUnregisteredEventName
}

// UnpackLabelUnregisteredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event LabelUnregistered(uint256 indexed tokenId, address indexed sender)
func (wrapperRegistry *WrapperRegistry) UnpackLabelUnregisteredEvent(log *types.Log) (*WrapperRegistryLabelUnregistered, error) {
	event := "LabelUnregistered"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != wrapperRegistry.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(WrapperRegistryLabelUnregistered)
	if len(log.Data) > 0 {
		if err := wrapperRegistry.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range wrapperRegistry.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// WrapperRegistryParentUpdated represents a ParentUpdated event raised by the WrapperRegistry contract.
type WrapperRegistryParentUpdated struct {
	Parent common.Address
	Label  string
	Sender common.Address
	Raw    *types.Log // Blockchain specific contextual infos
}

const WrapperRegistryParentUpdatedEventName = "ParentUpdated"

// ContractEventName returns the user-defined event name.
func (WrapperRegistryParentUpdated) ContractEventName() string {
	return WrapperRegistryParentUpdatedEventName
}

// UnpackParentUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ParentUpdated(address indexed parent, string label, address indexed sender)
func (wrapperRegistry *WrapperRegistry) UnpackParentUpdatedEvent(log *types.Log) (*WrapperRegistryParentUpdated, error) {
	event := "ParentUpdated"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != wrapperRegistry.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(WrapperRegistryParentUpdated)
	if len(log.Data) > 0 {
		if err := wrapperRegistry.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range wrapperRegistry.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// WrapperRegistryRegistryCreated represents a RegistryCreated event raised by the WrapperRegistry contract.
type WrapperRegistryRegistryCreated struct {
	Raw *types.Log // Blockchain specific contextual infos
}

const WrapperRegistryRegistryCreatedEventName = "RegistryCreated"

// ContractEventName returns the user-defined event name.
func (WrapperRegistryRegistryCreated) ContractEventName() string {
	return WrapperRegistryRegistryCreatedEventName
}

// UnpackRegistryCreatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RegistryCreated()
func (wrapperRegistry *WrapperRegistry) UnpackRegistryCreatedEvent(log *types.Log) (*WrapperRegistryRegistryCreated, error) {
	event := "RegistryCreated"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != wrapperRegistry.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(WrapperRegistryRegistryCreated)
	if len(log.Data) > 0 {
		if err := wrapperRegistry.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range wrapperRegistry.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// WrapperRegistryResolverUpdated represents a ResolverUpdated event raised by the WrapperRegistry contract.
type WrapperRegistryResolverUpdated struct {
	TokenId  *big.Int
	Resolver common.Address
	Sender   common.Address
	Raw      *types.Log // Blockchain specific contextual infos
}

const WrapperRegistryResolverUpdatedEventName = "ResolverUpdated"

// ContractEventName returns the user-defined event name.
func (WrapperRegistryResolverUpdated) ContractEventName() string {
	return WrapperRegistryResolverUpdatedEventName
}

// UnpackResolverUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ResolverUpdated(uint256 indexed tokenId, address indexed resolver, address indexed sender)
func (wrapperRegistry *WrapperRegistry) UnpackResolverUpdatedEvent(log *types.Log) (*WrapperRegistryResolverUpdated, error) {
	event := "ResolverUpdated"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != wrapperRegistry.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(WrapperRegistryResolverUpdated)
	if len(log.Data) > 0 {
		if err := wrapperRegistry.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range wrapperRegistry.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// WrapperRegistrySubregistryUpdated represents a SubregistryUpdated event raised by the WrapperRegistry contract.
type WrapperRegistrySubregistryUpdated struct {
	TokenId     *big.Int
	Subregistry common.Address
	Sender      common.Address
	Raw         *types.Log // Blockchain specific contextual infos
}

const WrapperRegistrySubregistryUpdatedEventName = "SubregistryUpdated"

// ContractEventName returns the user-defined event name.
func (WrapperRegistrySubregistryUpdated) ContractEventName() string {
	return WrapperRegistrySubregistryUpdatedEventName
}

// UnpackSubregistryUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event SubregistryUpdated(uint256 indexed tokenId, address indexed subregistry, address indexed sender)
func (wrapperRegistry *WrapperRegistry) UnpackSubregistryUpdatedEvent(log *types.Log) (*WrapperRegistrySubregistryUpdated, error) {
	event := "SubregistryUpdated"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != wrapperRegistry.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(WrapperRegistrySubregistryUpdated)
	if len(log.Data) > 0 {
		if err := wrapperRegistry.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range wrapperRegistry.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// WrapperRegistryTokenRegenerated represents a TokenRegenerated event raised by the WrapperRegistry contract.
type WrapperRegistryTokenRegenerated struct {
	OldTokenId *big.Int
	NewTokenId *big.Int
	Raw        *types.Log // Blockchain specific contextual infos
}

const WrapperRegistryTokenRegeneratedEventName = "TokenRegenerated"

// ContractEventName returns the user-defined event name.
func (WrapperRegistryTokenRegenerated) ContractEventName() string {
	return WrapperRegistryTokenRegeneratedEventName
}

// UnpackTokenRegeneratedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TokenRegenerated(uint256 indexed oldTokenId, uint256 indexed newTokenId)
func (wrapperRegistry *WrapperRegistry) UnpackTokenRegeneratedEvent(log *types.Log) (*WrapperRegistryTokenRegenerated, error) {
	event := "TokenRegenerated"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != wrapperRegistry.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(WrapperRegistryTokenRegenerated)
	if len(log.Data) > 0 {
		if err := wrapperRegistry.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range wrapperRegistry.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// WrapperRegistryTokenResource represents a TokenResource event raised by the WrapperRegistry contract.
type WrapperRegistryTokenResource struct {
	TokenId  *big.Int
	Resource *big.Int
	Raw      *types.Log // Blockchain specific contextual infos
}

const WrapperRegistryTokenResourceEventName = "TokenResource"

// ContractEventName returns the user-defined event name.
func (WrapperRegistryTokenResource) ContractEventName() string {
	return WrapperRegistryTokenResourceEventName
}

// UnpackTokenResourceEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TokenResource(uint256 indexed tokenId, uint256 indexed resource)
func (wrapperRegistry *WrapperRegistry) UnpackTokenResourceEvent(log *types.Log) (*WrapperRegistryTokenResource, error) {
	event := "TokenResource"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != wrapperRegistry.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(WrapperRegistryTokenResource)
	if len(log.Data) > 0 {
		if err := wrapperRegistry.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range wrapperRegistry.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// WrapperRegistryTransferBatch represents a TransferBatch event raised by the WrapperRegistry contract.
type WrapperRegistryTransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      *types.Log // Blockchain specific contextual infos
}

const WrapperRegistryTransferBatchEventName = "TransferBatch"

// ContractEventName returns the user-defined event name.
func (WrapperRegistryTransferBatch) ContractEventName() string {
	return WrapperRegistryTransferBatchEventName
}

// UnpackTransferBatchEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (wrapperRegistry *WrapperRegistry) UnpackTransferBatchEvent(log *types.Log) (*WrapperRegistryTransferBatch, error) {
	event := "TransferBatch"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != wrapperRegistry.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(WrapperRegistryTransferBatch)
	if len(log.Data) > 0 {
		if err := wrapperRegistry.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range wrapperRegistry.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// WrapperRegistryTransferSingle represents a TransferSingle event raised by the WrapperRegistry contract.
type WrapperRegistryTransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      *types.Log // Blockchain specific contextual infos
}

const WrapperRegistryTransferSingleEventName = "TransferSingle"

// ContractEventName returns the user-defined event name.
func (WrapperRegistryTransferSingle) ContractEventName() string {
	return WrapperRegistryTransferSingleEventName
}

// UnpackTransferSingleEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (wrapperRegistry *WrapperRegistry) UnpackTransferSingleEvent(log *types.Log) (*WrapperRegistryTransferSingle, error) {
	event := "TransferSingle"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != wrapperRegistry.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(WrapperRegistryTransferSingle)
	if len(log.Data) > 0 {
		if err := wrapperRegistry.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range wrapperRegistry.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// WrapperRegistryURI represents a URI event raised by the WrapperRegistry contract.
type WrapperRegistryURI struct {
	Value string
	Id    *big.Int
	Raw   *types.Log // Blockchain specific contextual infos
}

const WrapperRegistryURIEventName = "URI"

// ContractEventName returns the user-defined event name.
func (WrapperRegistryURI) ContractEventName() string {
	return WrapperRegistryURIEventName
}

// UnpackURIEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event URI(string value, uint256 indexed id)
func (wrapperRegistry *WrapperRegistry) UnpackURIEvent(log *types.Log) (*WrapperRegistryURI, error) {
	event := "URI"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != wrapperRegistry.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(WrapperRegistryURI)
	if len(log.Data) > 0 {
		if err := wrapperRegistry.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range wrapperRegistry.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// WrapperRegistryURIUpdated represents a URIUpdated event raised by the WrapperRegistry contract.
type WrapperRegistryURIUpdated struct {
	Uri      string
	Renderer common.Address
	Sender   common.Address
	Raw      *types.Log // Blockchain specific contextual infos
}

const WrapperRegistryURIUpdatedEventName = "URIUpdated"

// ContractEventName returns the user-defined event name.
func (WrapperRegistryURIUpdated) ContractEventName() string {
	return WrapperRegistryURIUpdatedEventName
}

// UnpackURIUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event URIUpdated(string uri, address renderer, address indexed sender)
func (wrapperRegistry *WrapperRegistry) UnpackURIUpdatedEvent(log *types.Log) (*WrapperRegistryURIUpdated, error) {
	event := "URIUpdated"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != wrapperRegistry.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(WrapperRegistryURIUpdated)
	if len(log.Data) > 0 {
		if err := wrapperRegistry.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range wrapperRegistry.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// WrapperRegistryUpgraded represents a Upgraded event raised by the WrapperRegistry contract.
type WrapperRegistryUpgraded struct {
	Implementation common.Address
	Raw            *types.Log // Blockchain specific contextual infos
}

const WrapperRegistryUpgradedEventName = "Upgraded"

// ContractEventName returns the user-defined event name.
func (WrapperRegistryUpgraded) ContractEventName() string {
	return WrapperRegistryUpgradedEventName
}

// UnpackUpgradedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Upgraded(address indexed implementation)
func (wrapperRegistry *WrapperRegistry) UnpackUpgradedEvent(log *types.Log) (*WrapperRegistryUpgraded, error) {
	event := "Upgraded"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != wrapperRegistry.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(WrapperRegistryUpgraded)
	if len(log.Data) > 0 {
		if err := wrapperRegistry.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range wrapperRegistry.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// UnpackError attempts to decode the provided error data using user-defined
// error definitions.
func (wrapperRegistry *WrapperRegistry) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], wrapperRegistry.abi.Errors["AddressEmptyCode"].ID.Bytes()[:4]) {
		return wrapperRegistry.UnpackAddressEmptyCodeError(raw[4:])
	}
	if bytes.Equal(raw[:4], wrapperRegistry.abi.Errors["CannotReduceExpiry"].ID.Bytes()[:4]) {
		return wrapperRegistry.UnpackCannotReduceExpiryError(raw[4:])
	}
	if bytes.Equal(raw[:4], wrapperRegistry.abi.Errors["CannotSetPastExpiry"].ID.Bytes()[:4]) {
		return wrapperRegistry.UnpackCannotSetPastExpiryError(raw[4:])
	}
	if bytes.Equal(raw[:4], wrapperRegistry.abi.Errors["EACCannotGrantRoles"].ID.Bytes()[:4]) {
		return wrapperRegistry.UnpackEACCannotGrantRolesError(raw[4:])
	}
	if bytes.Equal(raw[:4], wrapperRegistry.abi.Errors["EACCannotRevokeRoles"].ID.Bytes()[:4]) {
		return wrapperRegistry.UnpackEACCannotRevokeRolesError(raw[4:])
	}
	if bytes.Equal(raw[:4], wrapperRegistry.abi.Errors["EACInvalidAccount"].ID.Bytes()[:4]) {
		return wrapperRegistry.UnpackEACInvalidAccountError(raw[4:])
	}
	if bytes.Equal(raw[:4], wrapperRegistry.abi.Errors["EACInvalidRoleBitmap"].ID.Bytes()[:4]) {
		return wrapperRegistry.UnpackEACInvalidRoleBitmapError(raw[4:])
	}
	if bytes.Equal(raw[:4], wrapperRegistry.abi.Errors["EACMaxAssignees"].ID.Bytes()[:4]) {
		return wrapperRegistry.UnpackEACMaxAssigneesError(raw[4:])
	}
	if bytes.Equal(raw[:4], wrapperRegistry.abi.Errors["EACMinAssignees"].ID.Bytes()[:4]) {
		return wrapperRegistry.UnpackEACMinAssigneesError(raw[4:])
	}
	if bytes.Equal(raw[:4], wrapperRegistry.abi.Errors["EACRootResourceNotAllowed"].ID.Bytes()[:4]) {
		return wrapperRegistry.UnpackEACRootResourceNotAllowedError(raw[4:])
	}
	if bytes.Equal(raw[:4], wrapperRegistry.abi.Errors["EACUnauthorizedAccountRoles"].ID.Bytes()[:4]) {
		return wrapperRegistry.UnpackEACUnauthorizedAccountRolesError(raw[4:])
	}
	if bytes.Equal(raw[:4], wrapperRegistry.abi.Errors["ERC1155InsufficientBalance"].ID.Bytes()[:4]) {
		return wrapperRegistry.UnpackERC1155InsufficientBalanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], wrapperRegistry.abi.Errors["ERC1155InvalidApprover"].ID.Bytes()[:4]) {
		return wrapperRegistry.UnpackERC1155InvalidApproverError(raw[4:])
	}
	if bytes.Equal(raw[:4], wrapperRegistry.abi.Errors["ERC1155InvalidArrayLength"].ID.Bytes()[:4]) {
		return wrapperRegistry.UnpackERC1155InvalidArrayLengthError(raw[4:])
	}
	if bytes.Equal(raw[:4], wrapperRegistry.abi.Errors["ERC1155InvalidOperator"].ID.Bytes()[:4]) {
		return wrapperRegistry.UnpackERC1155InvalidOperatorError(raw[4:])
	}
	if bytes.Equal(raw[:4], wrapperRegistry.abi.Errors["ERC1155InvalidReceiver"].ID.Bytes()[:4]) {
		return wrapperRegistry.UnpackERC1155InvalidReceiverError(raw[4:])
	}
	if bytes.Equal(raw[:4], wrapperRegistry.abi.Errors["ERC1155InvalidSender"].ID.Bytes()[:4]) {
		return wrapperRegistry.UnpackERC1155InvalidSenderError(raw[4:])
	}
	if bytes.Equal(raw[:4], wrapperRegistry.abi.Errors["ERC1155MissingApprovalForAll"].ID.Bytes()[:4]) {
		return wrapperRegistry.UnpackERC1155MissingApprovalForAllError(raw[4:])
	}
	if bytes.Equal(raw[:4], wrapperRegistry.abi.Errors["ERC1967InvalidImplementation"].ID.Bytes()[:4]) {
		return wrapperRegistry.UnpackERC1967InvalidImplementationError(raw[4:])
	}
	if bytes.Equal(raw[:4], wrapperRegistry.abi.Errors["ERC1967NonPayable"].ID.Bytes()[:4]) {
		return wrapperRegistry.UnpackERC1967NonPayableError(raw[4:])
	}
	if bytes.Equal(raw[:4], wrapperRegistry.abi.Errors["FailedCall"].ID.Bytes()[:4]) {
		return wrapperRegistry.UnpackFailedCallError(raw[4:])
	}
	if bytes.Equal(raw[:4], wrapperRegistry.abi.Errors["FrozenTokenApproval"].ID.Bytes()[:4]) {
		return wrapperRegistry.UnpackFrozenTokenApprovalError(raw[4:])
	}
	if bytes.Equal(raw[:4], wrapperRegistry.abi.Errors["InvalidInitialization"].ID.Bytes()[:4]) {
		return wrapperRegistry.UnpackInvalidInitializationError(raw[4:])
	}
	if bytes.Equal(raw[:4], wrapperRegistry.abi.Errors["InvalidOwner"].ID.Bytes()[:4]) {
		return wrapperRegistry.UnpackInvalidOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], wrapperRegistry.abi.Errors["LabelAlreadyRegistered"].ID.Bytes()[:4]) {
		return wrapperRegistry.UnpackLabelAlreadyRegisteredError(raw[4:])
	}
	if bytes.Equal(raw[:4], wrapperRegistry.abi.Errors["LabelAlreadyReserved"].ID.Bytes()[:4]) {
		return wrapperRegistry.UnpackLabelAlreadyReservedError(raw[4:])
	}
	if bytes.Equal(raw[:4], wrapperRegistry.abi.Errors["LabelExpired"].ID.Bytes()[:4]) {
		return wrapperRegistry.UnpackLabelExpiredError(raw[4:])
	}
	if bytes.Equal(raw[:4], wrapperRegistry.abi.Errors["NameDataMismatch"].ID.Bytes()[:4]) {
		return wrapperRegistry.UnpackNameDataMismatchError(raw[4:])
	}
	if bytes.Equal(raw[:4], wrapperRegistry.abi.Errors["NameNotLocked"].ID.Bytes()[:4]) {
		return wrapperRegistry.UnpackNameNotLockedError(raw[4:])
	}
	if bytes.Equal(raw[:4], wrapperRegistry.abi.Errors["NameRequiresMigration"].ID.Bytes()[:4]) {
		return wrapperRegistry.UnpackNameRequiresMigrationError(raw[4:])
	}
	if bytes.Equal(raw[:4], wrapperRegistry.abi.Errors["NotInitializing"].ID.Bytes()[:4]) {
		return wrapperRegistry.UnpackNotInitializingError(raw[4:])
	}
	if bytes.Equal(raw[:4], wrapperRegistry.abi.Errors["TransferDisallowed"].ID.Bytes()[:4]) {
		return wrapperRegistry.UnpackTransferDisallowedError(raw[4:])
	}
	if bytes.Equal(raw[:4], wrapperRegistry.abi.Errors["TransferUnsafeUntilRegistryIsEmancipated"].ID.Bytes()[:4]) {
		return wrapperRegistry.UnpackTransferUnsafeUntilRegistryIsEmancipatedError(raw[4:])
	}
	if bytes.Equal(raw[:4], wrapperRegistry.abi.Errors["TransferUnsafeWithMultipleAssignees"].ID.Bytes()[:4]) {
		return wrapperRegistry.UnpackTransferUnsafeWithMultipleAssigneesError(raw[4:])
	}
	if bytes.Equal(raw[:4], wrapperRegistry.abi.Errors["UUPSUnauthorizedCallContext"].ID.Bytes()[:4]) {
		return wrapperRegistry.UnpackUUPSUnauthorizedCallContextError(raw[4:])
	}
	if bytes.Equal(raw[:4], wrapperRegistry.abi.Errors["UUPSUnsupportedProxiableUUID"].ID.Bytes()[:4]) {
		return wrapperRegistry.UnpackUUPSUnsupportedProxiableUUIDError(raw[4:])
	}
	if bytes.Equal(raw[:4], wrapperRegistry.abi.Errors["UnauthorizedCaller"].ID.Bytes()[:4]) {
		return wrapperRegistry.UnpackUnauthorizedCallerError(raw[4:])
	}
	if bytes.Equal(raw[:4], wrapperRegistry.abi.Errors["UpgradeTargetNotApproved"].ID.Bytes()[:4]) {
		return wrapperRegistry.UnpackUpgradeTargetNotApprovedError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// WrapperRegistryAddressEmptyCode represents a AddressEmptyCode error raised by the WrapperRegistry contract.
type WrapperRegistryAddressEmptyCode struct {
	Target common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AddressEmptyCode(address target)
func WrapperRegistryAddressEmptyCodeErrorID() common.Hash {
	return common.HexToHash("0x9996b315c842ff135b8fc4a08ad5df1c344efbc03d2687aecc0678050d2aac89")
}

// UnpackAddressEmptyCodeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AddressEmptyCode(address target)
func (wrapperRegistry *WrapperRegistry) UnpackAddressEmptyCodeError(raw []byte) (*WrapperRegistryAddressEmptyCode, error) {
	out := new(WrapperRegistryAddressEmptyCode)
	if err := wrapperRegistry.abi.UnpackIntoInterface(out, "AddressEmptyCode", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WrapperRegistryCannotReduceExpiry represents a CannotReduceExpiry error raised by the WrapperRegistry contract.
type WrapperRegistryCannotReduceExpiry struct {
	OldExpiry uint64
	NewExpiry uint64
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error CannotReduceExpiry(uint64 oldExpiry, uint64 newExpiry)
func WrapperRegistryCannotReduceExpiryErrorID() common.Hash {
	return common.HexToHash("0x68c1425a85303a923d9963b6b77886de6c657b9cc761a422fa2090222a51b5ca")
}

// UnpackCannotReduceExpiryError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error CannotReduceExpiry(uint64 oldExpiry, uint64 newExpiry)
func (wrapperRegistry *WrapperRegistry) UnpackCannotReduceExpiryError(raw []byte) (*WrapperRegistryCannotReduceExpiry, error) {
	out := new(WrapperRegistryCannotReduceExpiry)
	if err := wrapperRegistry.abi.UnpackIntoInterface(out, "CannotReduceExpiry", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WrapperRegistryCannotSetPastExpiry represents a CannotSetPastExpiry error raised by the WrapperRegistry contract.
type WrapperRegistryCannotSetPastExpiry struct {
	Expiry uint64
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error CannotSetPastExpiry(uint64 expiry)
func WrapperRegistryCannotSetPastExpiryErrorID() common.Hash {
	return common.HexToHash("0xf1d446c394c7afd3028f17f9fc3b3d460c3e997e3b2a7fc80b87d86295b1a896")
}

// UnpackCannotSetPastExpiryError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error CannotSetPastExpiry(uint64 expiry)
func (wrapperRegistry *WrapperRegistry) UnpackCannotSetPastExpiryError(raw []byte) (*WrapperRegistryCannotSetPastExpiry, error) {
	out := new(WrapperRegistryCannotSetPastExpiry)
	if err := wrapperRegistry.abi.UnpackIntoInterface(out, "CannotSetPastExpiry", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WrapperRegistryEACCannotGrantRoles represents a EACCannotGrantRoles error raised by the WrapperRegistry contract.
type WrapperRegistryEACCannotGrantRoles struct {
	Resource   *big.Int
	RoleBitmap *big.Int
	Account    common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error EACCannotGrantRoles(uint256 resource, uint256 roleBitmap, address account)
func WrapperRegistryEACCannotGrantRolesErrorID() common.Hash {
	return common.HexToHash("0xd1a3b355154fa30324f856ca6b25179bcc6c9d2936e28429b1b4293509e97f4a")
}

// UnpackEACCannotGrantRolesError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error EACCannotGrantRoles(uint256 resource, uint256 roleBitmap, address account)
func (wrapperRegistry *WrapperRegistry) UnpackEACCannotGrantRolesError(raw []byte) (*WrapperRegistryEACCannotGrantRoles, error) {
	out := new(WrapperRegistryEACCannotGrantRoles)
	if err := wrapperRegistry.abi.UnpackIntoInterface(out, "EACCannotGrantRoles", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WrapperRegistryEACCannotRevokeRoles represents a EACCannotRevokeRoles error raised by the WrapperRegistry contract.
type WrapperRegistryEACCannotRevokeRoles struct {
	Resource   *big.Int
	RoleBitmap *big.Int
	Account    common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error EACCannotRevokeRoles(uint256 resource, uint256 roleBitmap, address account)
func WrapperRegistryEACCannotRevokeRolesErrorID() common.Hash {
	return common.HexToHash("0xa604e318421f3889d41e189ce5a1bbb63773eb6844a6d371417f9eed8f6cb153")
}

// UnpackEACCannotRevokeRolesError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error EACCannotRevokeRoles(uint256 resource, uint256 roleBitmap, address account)
func (wrapperRegistry *WrapperRegistry) UnpackEACCannotRevokeRolesError(raw []byte) (*WrapperRegistryEACCannotRevokeRoles, error) {
	out := new(WrapperRegistryEACCannotRevokeRoles)
	if err := wrapperRegistry.abi.UnpackIntoInterface(out, "EACCannotRevokeRoles", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WrapperRegistryEACInvalidAccount represents a EACInvalidAccount error raised by the WrapperRegistry contract.
type WrapperRegistryEACInvalidAccount struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error EACInvalidAccount()
func WrapperRegistryEACInvalidAccountErrorID() common.Hash {
	return common.HexToHash("0xec3fc5924c67c1c1507473878ec47d207e41f4b88da4d7a835c6e5043b63da51")
}

// UnpackEACInvalidAccountError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error EACInvalidAccount()
func (wrapperRegistry *WrapperRegistry) UnpackEACInvalidAccountError(raw []byte) (*WrapperRegistryEACInvalidAccount, error) {
	out := new(WrapperRegistryEACInvalidAccount)
	if err := wrapperRegistry.abi.UnpackIntoInterface(out, "EACInvalidAccount", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WrapperRegistryEACInvalidRoleBitmap represents a EACInvalidRoleBitmap error raised by the WrapperRegistry contract.
type WrapperRegistryEACInvalidRoleBitmap struct {
	RoleBitmap *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error EACInvalidRoleBitmap(uint256 roleBitmap)
func WrapperRegistryEACInvalidRoleBitmapErrorID() common.Hash {
	return common.HexToHash("0x2a7b2d20191e44fd4e4227272caf529612678184b58adb291afe5231bf5b2734")
}

// UnpackEACInvalidRoleBitmapError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error EACInvalidRoleBitmap(uint256 roleBitmap)
func (wrapperRegistry *WrapperRegistry) UnpackEACInvalidRoleBitmapError(raw []byte) (*WrapperRegistryEACInvalidRoleBitmap, error) {
	out := new(WrapperRegistryEACInvalidRoleBitmap)
	if err := wrapperRegistry.abi.UnpackIntoInterface(out, "EACInvalidRoleBitmap", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WrapperRegistryEACMaxAssignees represents a EACMaxAssignees error raised by the WrapperRegistry contract.
type WrapperRegistryEACMaxAssignees struct {
	Resource *big.Int
	Role     *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error EACMaxAssignees(uint256 resource, uint256 role)
func WrapperRegistryEACMaxAssigneesErrorID() common.Hash {
	return common.HexToHash("0xf91653484f73ddf3683475f443ccba4f564362c5896702ddf28889a64960b109")
}

// UnpackEACMaxAssigneesError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error EACMaxAssignees(uint256 resource, uint256 role)
func (wrapperRegistry *WrapperRegistry) UnpackEACMaxAssigneesError(raw []byte) (*WrapperRegistryEACMaxAssignees, error) {
	out := new(WrapperRegistryEACMaxAssignees)
	if err := wrapperRegistry.abi.UnpackIntoInterface(out, "EACMaxAssignees", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WrapperRegistryEACMinAssignees represents a EACMinAssignees error raised by the WrapperRegistry contract.
type WrapperRegistryEACMinAssignees struct {
	Resource *big.Int
	Role     *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error EACMinAssignees(uint256 resource, uint256 role)
func WrapperRegistryEACMinAssigneesErrorID() common.Hash {
	return common.HexToHash("0x1f80c19ba45f0369eb8ace81d631dde7c47ae7c144205c7e77e1555ff7bb091b")
}

// UnpackEACMinAssigneesError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error EACMinAssignees(uint256 resource, uint256 role)
func (wrapperRegistry *WrapperRegistry) UnpackEACMinAssigneesError(raw []byte) (*WrapperRegistryEACMinAssignees, error) {
	out := new(WrapperRegistryEACMinAssignees)
	if err := wrapperRegistry.abi.UnpackIntoInterface(out, "EACMinAssignees", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WrapperRegistryEACRootResourceNotAllowed represents a EACRootResourceNotAllowed error raised by the WrapperRegistry contract.
type WrapperRegistryEACRootResourceNotAllowed struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error EACRootResourceNotAllowed()
func WrapperRegistryEACRootResourceNotAllowedErrorID() common.Hash {
	return common.HexToHash("0xc2842458b56deb79dcc9694a98069c21c928dfa5456787ba7ae1728a0d863238")
}

// UnpackEACRootResourceNotAllowedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error EACRootResourceNotAllowed()
func (wrapperRegistry *WrapperRegistry) UnpackEACRootResourceNotAllowedError(raw []byte) (*WrapperRegistryEACRootResourceNotAllowed, error) {
	out := new(WrapperRegistryEACRootResourceNotAllowed)
	if err := wrapperRegistry.abi.UnpackIntoInterface(out, "EACRootResourceNotAllowed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WrapperRegistryEACUnauthorizedAccountRoles represents a EACUnauthorizedAccountRoles error raised by the WrapperRegistry contract.
type WrapperRegistryEACUnauthorizedAccountRoles struct {
	Resource   *big.Int
	RoleBitmap *big.Int
	Account    common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error EACUnauthorizedAccountRoles(uint256 resource, uint256 roleBitmap, address account)
func WrapperRegistryEACUnauthorizedAccountRolesErrorID() common.Hash {
	return common.HexToHash("0x4b27a13377de460ddbf323b4163e94f59f4092e2fe9befdabaf9bd06c2dd939b")
}

// UnpackEACUnauthorizedAccountRolesError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error EACUnauthorizedAccountRoles(uint256 resource, uint256 roleBitmap, address account)
func (wrapperRegistry *WrapperRegistry) UnpackEACUnauthorizedAccountRolesError(raw []byte) (*WrapperRegistryEACUnauthorizedAccountRoles, error) {
	out := new(WrapperRegistryEACUnauthorizedAccountRoles)
	if err := wrapperRegistry.abi.UnpackIntoInterface(out, "EACUnauthorizedAccountRoles", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WrapperRegistryERC1155InsufficientBalance represents a ERC1155InsufficientBalance error raised by the WrapperRegistry contract.
type WrapperRegistryERC1155InsufficientBalance struct {
	Sender  common.Address
	Balance *big.Int
	Needed  *big.Int
	TokenId *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1155InsufficientBalance(address sender, uint256 balance, uint256 needed, uint256 tokenId)
func WrapperRegistryERC1155InsufficientBalanceErrorID() common.Hash {
	return common.HexToHash("0x03dee4c573c982787b5f3537d6323ffaca9d864448aa6bd828ada9e5d0837036")
}

// UnpackERC1155InsufficientBalanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1155InsufficientBalance(address sender, uint256 balance, uint256 needed, uint256 tokenId)
func (wrapperRegistry *WrapperRegistry) UnpackERC1155InsufficientBalanceError(raw []byte) (*WrapperRegistryERC1155InsufficientBalance, error) {
	out := new(WrapperRegistryERC1155InsufficientBalance)
	if err := wrapperRegistry.abi.UnpackIntoInterface(out, "ERC1155InsufficientBalance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WrapperRegistryERC1155InvalidApprover represents a ERC1155InvalidApprover error raised by the WrapperRegistry contract.
type WrapperRegistryERC1155InvalidApprover struct {
	Approver common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1155InvalidApprover(address approver)
func WrapperRegistryERC1155InvalidApproverErrorID() common.Hash {
	return common.HexToHash("0x3e31884e33c33ce0039d1905e3c252950ae3b95240f36d4fff81f5ff6752ef99")
}

// UnpackERC1155InvalidApproverError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1155InvalidApprover(address approver)
func (wrapperRegistry *WrapperRegistry) UnpackERC1155InvalidApproverError(raw []byte) (*WrapperRegistryERC1155InvalidApprover, error) {
	out := new(WrapperRegistryERC1155InvalidApprover)
	if err := wrapperRegistry.abi.UnpackIntoInterface(out, "ERC1155InvalidApprover", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WrapperRegistryERC1155InvalidArrayLength represents a ERC1155InvalidArrayLength error raised by the WrapperRegistry contract.
type WrapperRegistryERC1155InvalidArrayLength struct {
	IdsLength    *big.Int
	ValuesLength *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1155InvalidArrayLength(uint256 idsLength, uint256 valuesLength)
func WrapperRegistryERC1155InvalidArrayLengthErrorID() common.Hash {
	return common.HexToHash("0x5b0599913619cfa5633692652638ed25cafcd079c9beae8c251b12c23dcc83f2")
}

// UnpackERC1155InvalidArrayLengthError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1155InvalidArrayLength(uint256 idsLength, uint256 valuesLength)
func (wrapperRegistry *WrapperRegistry) UnpackERC1155InvalidArrayLengthError(raw []byte) (*WrapperRegistryERC1155InvalidArrayLength, error) {
	out := new(WrapperRegistryERC1155InvalidArrayLength)
	if err := wrapperRegistry.abi.UnpackIntoInterface(out, "ERC1155InvalidArrayLength", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WrapperRegistryERC1155InvalidOperator represents a ERC1155InvalidOperator error raised by the WrapperRegistry contract.
type WrapperRegistryERC1155InvalidOperator struct {
	Operator common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1155InvalidOperator(address operator)
func WrapperRegistryERC1155InvalidOperatorErrorID() common.Hash {
	return common.HexToHash("0xced3e10010b9d2aa24827119d0db4a8feec73aea48b4b3e470d8a9f3ff723569")
}

// UnpackERC1155InvalidOperatorError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1155InvalidOperator(address operator)
func (wrapperRegistry *WrapperRegistry) UnpackERC1155InvalidOperatorError(raw []byte) (*WrapperRegistryERC1155InvalidOperator, error) {
	out := new(WrapperRegistryERC1155InvalidOperator)
	if err := wrapperRegistry.abi.UnpackIntoInterface(out, "ERC1155InvalidOperator", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WrapperRegistryERC1155InvalidReceiver represents a ERC1155InvalidReceiver error raised by the WrapperRegistry contract.
type WrapperRegistryERC1155InvalidReceiver struct {
	Receiver common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1155InvalidReceiver(address receiver)
func WrapperRegistryERC1155InvalidReceiverErrorID() common.Hash {
	return common.HexToHash("0x57f447ceed621d9e134e26de5772c88799abb7322ce2a87f95dce247d47105c6")
}

// UnpackERC1155InvalidReceiverError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1155InvalidReceiver(address receiver)
func (wrapperRegistry *WrapperRegistry) UnpackERC1155InvalidReceiverError(raw []byte) (*WrapperRegistryERC1155InvalidReceiver, error) {
	out := new(WrapperRegistryERC1155InvalidReceiver)
	if err := wrapperRegistry.abi.UnpackIntoInterface(out, "ERC1155InvalidReceiver", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WrapperRegistryERC1155InvalidSender represents a ERC1155InvalidSender error raised by the WrapperRegistry contract.
type WrapperRegistryERC1155InvalidSender struct {
	Sender common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1155InvalidSender(address sender)
func WrapperRegistryERC1155InvalidSenderErrorID() common.Hash {
	return common.HexToHash("0x01a83514e94b34009110b75cac6742ba33bd7c62f18a3616bafea52855d3b175")
}

// UnpackERC1155InvalidSenderError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1155InvalidSender(address sender)
func (wrapperRegistry *WrapperRegistry) UnpackERC1155InvalidSenderError(raw []byte) (*WrapperRegistryERC1155InvalidSender, error) {
	out := new(WrapperRegistryERC1155InvalidSender)
	if err := wrapperRegistry.abi.UnpackIntoInterface(out, "ERC1155InvalidSender", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WrapperRegistryERC1155MissingApprovalForAll represents a ERC1155MissingApprovalForAll error raised by the WrapperRegistry contract.
type WrapperRegistryERC1155MissingApprovalForAll struct {
	Operator common.Address
	Owner    common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1155MissingApprovalForAll(address operator, address owner)
func WrapperRegistryERC1155MissingApprovalForAllErrorID() common.Hash {
	return common.HexToHash("0xe237d922be9fac42efeaaaffb42cc6b57e0ff95d94a1b74daeff69adc7657754")
}

// UnpackERC1155MissingApprovalForAllError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1155MissingApprovalForAll(address operator, address owner)
func (wrapperRegistry *WrapperRegistry) UnpackERC1155MissingApprovalForAllError(raw []byte) (*WrapperRegistryERC1155MissingApprovalForAll, error) {
	out := new(WrapperRegistryERC1155MissingApprovalForAll)
	if err := wrapperRegistry.abi.UnpackIntoInterface(out, "ERC1155MissingApprovalForAll", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WrapperRegistryERC1967InvalidImplementation represents a ERC1967InvalidImplementation error raised by the WrapperRegistry contract.
type WrapperRegistryERC1967InvalidImplementation struct {
	Implementation common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1967InvalidImplementation(address implementation)
func WrapperRegistryERC1967InvalidImplementationErrorID() common.Hash {
	return common.HexToHash("0x4c9c8ce3ceb3130f17f7cdba48d89b5b0129f266a8bac114e6e315a41879b617")
}

// UnpackERC1967InvalidImplementationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1967InvalidImplementation(address implementation)
func (wrapperRegistry *WrapperRegistry) UnpackERC1967InvalidImplementationError(raw []byte) (*WrapperRegistryERC1967InvalidImplementation, error) {
	out := new(WrapperRegistryERC1967InvalidImplementation)
	if err := wrapperRegistry.abi.UnpackIntoInterface(out, "ERC1967InvalidImplementation", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WrapperRegistryERC1967NonPayable represents a ERC1967NonPayable error raised by the WrapperRegistry contract.
type WrapperRegistryERC1967NonPayable struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1967NonPayable()
func WrapperRegistryERC1967NonPayableErrorID() common.Hash {
	return common.HexToHash("0xb398979fa84f543c8e222f17890372c487baf85e062276c127fef521eea7224b")
}

// UnpackERC1967NonPayableError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1967NonPayable()
func (wrapperRegistry *WrapperRegistry) UnpackERC1967NonPayableError(raw []byte) (*WrapperRegistryERC1967NonPayable, error) {
	out := new(WrapperRegistryERC1967NonPayable)
	if err := wrapperRegistry.abi.UnpackIntoInterface(out, "ERC1967NonPayable", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WrapperRegistryFailedCall represents a FailedCall error raised by the WrapperRegistry contract.
type WrapperRegistryFailedCall struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FailedCall()
func WrapperRegistryFailedCallErrorID() common.Hash {
	return common.HexToHash("0xd6bda27508c0fb6d8a39b4b122878dab26f731a7d4e4abe711dd3731899052a4")
}

// UnpackFailedCallError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FailedCall()
func (wrapperRegistry *WrapperRegistry) UnpackFailedCallError(raw []byte) (*WrapperRegistryFailedCall, error) {
	out := new(WrapperRegistryFailedCall)
	if err := wrapperRegistry.abi.UnpackIntoInterface(out, "FailedCall", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WrapperRegistryFrozenTokenApproval represents a FrozenTokenApproval error raised by the WrapperRegistry contract.
type WrapperRegistryFrozenTokenApproval struct {
	TokenId *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FrozenTokenApproval(uint256 tokenId)
func WrapperRegistryFrozenTokenApprovalErrorID() common.Hash {
	return common.HexToHash("0xa4f0771302f5d351751357424809d2b38e1240ff6d38091e0f202aef5ce0f754")
}

// UnpackFrozenTokenApprovalError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FrozenTokenApproval(uint256 tokenId)
func (wrapperRegistry *WrapperRegistry) UnpackFrozenTokenApprovalError(raw []byte) (*WrapperRegistryFrozenTokenApproval, error) {
	out := new(WrapperRegistryFrozenTokenApproval)
	if err := wrapperRegistry.abi.UnpackIntoInterface(out, "FrozenTokenApproval", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WrapperRegistryInvalidInitialization represents a InvalidInitialization error raised by the WrapperRegistry contract.
type WrapperRegistryInvalidInitialization struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidInitialization()
func WrapperRegistryInvalidInitializationErrorID() common.Hash {
	return common.HexToHash("0xf92ee8a957075833165f68c320933b1a1294aafc84ee6e0dd3fb178008f9aaf5")
}

// UnpackInvalidInitializationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidInitialization()
func (wrapperRegistry *WrapperRegistry) UnpackInvalidInitializationError(raw []byte) (*WrapperRegistryInvalidInitialization, error) {
	out := new(WrapperRegistryInvalidInitialization)
	if err := wrapperRegistry.abi.UnpackIntoInterface(out, "InvalidInitialization", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WrapperRegistryInvalidOwner represents a InvalidOwner error raised by the WrapperRegistry contract.
type WrapperRegistryInvalidOwner struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidOwner()
func WrapperRegistryInvalidOwnerErrorID() common.Hash {
	return common.HexToHash("0x49e27cffb37b1ca4a9bf5318243b2014d13f940af232b8552c208bdea15739da")
}

// UnpackInvalidOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidOwner()
func (wrapperRegistry *WrapperRegistry) UnpackInvalidOwnerError(raw []byte) (*WrapperRegistryInvalidOwner, error) {
	out := new(WrapperRegistryInvalidOwner)
	if err := wrapperRegistry.abi.UnpackIntoInterface(out, "InvalidOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WrapperRegistryLabelAlreadyRegistered represents a LabelAlreadyRegistered error raised by the WrapperRegistry contract.
type WrapperRegistryLabelAlreadyRegistered struct {
	Label string
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error LabelAlreadyRegistered(string label)
func WrapperRegistryLabelAlreadyRegisteredErrorID() common.Hash {
	return common.HexToHash("0xdef545a48127e18e04d05455886a7023f5b6c21f969846ca4ef9c9aa7cc2267b")
}

// UnpackLabelAlreadyRegisteredError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error LabelAlreadyRegistered(string label)
func (wrapperRegistry *WrapperRegistry) UnpackLabelAlreadyRegisteredError(raw []byte) (*WrapperRegistryLabelAlreadyRegistered, error) {
	out := new(WrapperRegistryLabelAlreadyRegistered)
	if err := wrapperRegistry.abi.UnpackIntoInterface(out, "LabelAlreadyRegistered", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WrapperRegistryLabelAlreadyReserved represents a LabelAlreadyReserved error raised by the WrapperRegistry contract.
type WrapperRegistryLabelAlreadyReserved struct {
	Label string
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error LabelAlreadyReserved(string label)
func WrapperRegistryLabelAlreadyReservedErrorID() common.Hash {
	return common.HexToHash("0xf60759e04b9d51b7561cf8055ffc1d1b1c222340d6daaf37c7a4c5819b3b9207")
}

// UnpackLabelAlreadyReservedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error LabelAlreadyReserved(string label)
func (wrapperRegistry *WrapperRegistry) UnpackLabelAlreadyReservedError(raw []byte) (*WrapperRegistryLabelAlreadyReserved, error) {
	out := new(WrapperRegistryLabelAlreadyReserved)
	if err := wrapperRegistry.abi.UnpackIntoInterface(out, "LabelAlreadyReserved", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WrapperRegistryLabelExpired represents a LabelExpired error raised by the WrapperRegistry contract.
type WrapperRegistryLabelExpired struct {
	TokenId *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error LabelExpired(uint256 tokenId)
func WrapperRegistryLabelExpiredErrorID() common.Hash {
	return common.HexToHash("0xc44e2374535590feeaeb825026475cb21e2c4eac2832b552c8811b602a69f016")
}

// UnpackLabelExpiredError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error LabelExpired(uint256 tokenId)
func (wrapperRegistry *WrapperRegistry) UnpackLabelExpiredError(raw []byte) (*WrapperRegistryLabelExpired, error) {
	out := new(WrapperRegistryLabelExpired)
	if err := wrapperRegistry.abi.UnpackIntoInterface(out, "LabelExpired", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WrapperRegistryNameDataMismatch represents a NameDataMismatch error raised by the WrapperRegistry contract.
type WrapperRegistryNameDataMismatch struct {
	TokenId *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NameDataMismatch(uint256 tokenId)
func WrapperRegistryNameDataMismatchErrorID() common.Hash {
	return common.HexToHash("0xedec3569de29bfd1a5c3fca776a007dbb804b5d63ec5b1b45f6cdd291ab34782")
}

// UnpackNameDataMismatchError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NameDataMismatch(uint256 tokenId)
func (wrapperRegistry *WrapperRegistry) UnpackNameDataMismatchError(raw []byte) (*WrapperRegistryNameDataMismatch, error) {
	out := new(WrapperRegistryNameDataMismatch)
	if err := wrapperRegistry.abi.UnpackIntoInterface(out, "NameDataMismatch", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WrapperRegistryNameNotLocked represents a NameNotLocked error raised by the WrapperRegistry contract.
type WrapperRegistryNameNotLocked struct {
	TokenId *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NameNotLocked(uint256 tokenId)
func WrapperRegistryNameNotLockedErrorID() common.Hash {
	return common.HexToHash("0x1bfe8f0a6060f4ab19fbc6bac4c5b41676d9b2147c5ce73d61105f9b4e7e90aa")
}

// UnpackNameNotLockedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NameNotLocked(uint256 tokenId)
func (wrapperRegistry *WrapperRegistry) UnpackNameNotLockedError(raw []byte) (*WrapperRegistryNameNotLocked, error) {
	out := new(WrapperRegistryNameNotLocked)
	if err := wrapperRegistry.abi.UnpackIntoInterface(out, "NameNotLocked", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WrapperRegistryNameRequiresMigration represents a NameRequiresMigration error raised by the WrapperRegistry contract.
type WrapperRegistryNameRequiresMigration struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NameRequiresMigration()
func WrapperRegistryNameRequiresMigrationErrorID() common.Hash {
	return common.HexToHash("0x408fa1b8ee4a4cd7bf27a10d8522d3e2949fc5c6ea14d8f739f5b903f2fd4522")
}

// UnpackNameRequiresMigrationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NameRequiresMigration()
func (wrapperRegistry *WrapperRegistry) UnpackNameRequiresMigrationError(raw []byte) (*WrapperRegistryNameRequiresMigration, error) {
	out := new(WrapperRegistryNameRequiresMigration)
	if err := wrapperRegistry.abi.UnpackIntoInterface(out, "NameRequiresMigration", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WrapperRegistryNotInitializing represents a NotInitializing error raised by the WrapperRegistry contract.
type WrapperRegistryNotInitializing struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotInitializing()
func WrapperRegistryNotInitializingErrorID() common.Hash {
	return common.HexToHash("0xd7e6bcf8597daa127dc9f0048d2f08d5ef140a2cb659feabd700beff1f7a8302")
}

// UnpackNotInitializingError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotInitializing()
func (wrapperRegistry *WrapperRegistry) UnpackNotInitializingError(raw []byte) (*WrapperRegistryNotInitializing, error) {
	out := new(WrapperRegistryNotInitializing)
	if err := wrapperRegistry.abi.UnpackIntoInterface(out, "NotInitializing", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WrapperRegistryTransferDisallowed represents a TransferDisallowed error raised by the WrapperRegistry contract.
type WrapperRegistryTransferDisallowed struct {
	TokenId *big.Int
	From    common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TransferDisallowed(uint256 tokenId, address from)
func WrapperRegistryTransferDisallowedErrorID() common.Hash {
	return common.HexToHash("0xe58f6d5a5faf9cca464192e088054deae4ec88a7ee4aafacd834f996253c270d")
}

// UnpackTransferDisallowedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TransferDisallowed(uint256 tokenId, address from)
func (wrapperRegistry *WrapperRegistry) UnpackTransferDisallowedError(raw []byte) (*WrapperRegistryTransferDisallowed, error) {
	out := new(WrapperRegistryTransferDisallowed)
	if err := wrapperRegistry.abi.UnpackIntoInterface(out, "TransferDisallowed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WrapperRegistryTransferUnsafeUntilRegistryIsEmancipated represents a TransferUnsafeUntilRegistryIsEmancipated error raised by the WrapperRegistry contract.
type WrapperRegistryTransferUnsafeUntilRegistryIsEmancipated struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TransferUnsafeUntilRegistryIsEmancipated()
func WrapperRegistryTransferUnsafeUntilRegistryIsEmancipatedErrorID() common.Hash {
	return common.HexToHash("0x54838f98c10e7309e5345cfa4bd39b2f43a7576ced98070099678ab413097390")
}

// UnpackTransferUnsafeUntilRegistryIsEmancipatedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TransferUnsafeUntilRegistryIsEmancipated()
func (wrapperRegistry *WrapperRegistry) UnpackTransferUnsafeUntilRegistryIsEmancipatedError(raw []byte) (*WrapperRegistryTransferUnsafeUntilRegistryIsEmancipated, error) {
	out := new(WrapperRegistryTransferUnsafeUntilRegistryIsEmancipated)
	if err := wrapperRegistry.abi.UnpackIntoInterface(out, "TransferUnsafeUntilRegistryIsEmancipated", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WrapperRegistryTransferUnsafeWithMultipleAssignees represents a TransferUnsafeWithMultipleAssignees error raised by the WrapperRegistry contract.
type WrapperRegistryTransferUnsafeWithMultipleAssignees struct {
	TokenId *big.Int
	From    common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TransferUnsafeWithMultipleAssignees(uint256 tokenId, address from)
func WrapperRegistryTransferUnsafeWithMultipleAssigneesErrorID() common.Hash {
	return common.HexToHash("0x677f1c18057fe4713efbb723b401a9d6b67633796a8433232c79a8e6fa834e27")
}

// UnpackTransferUnsafeWithMultipleAssigneesError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TransferUnsafeWithMultipleAssignees(uint256 tokenId, address from)
func (wrapperRegistry *WrapperRegistry) UnpackTransferUnsafeWithMultipleAssigneesError(raw []byte) (*WrapperRegistryTransferUnsafeWithMultipleAssignees, error) {
	out := new(WrapperRegistryTransferUnsafeWithMultipleAssignees)
	if err := wrapperRegistry.abi.UnpackIntoInterface(out, "TransferUnsafeWithMultipleAssignees", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WrapperRegistryUUPSUnauthorizedCallContext represents a UUPSUnauthorizedCallContext error raised by the WrapperRegistry contract.
type WrapperRegistryUUPSUnauthorizedCallContext struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UUPSUnauthorizedCallContext()
func WrapperRegistryUUPSUnauthorizedCallContextErrorID() common.Hash {
	return common.HexToHash("0xe07c8dba242a06571ac65fe4bbe20522c9fb111cb33599b799ff8039c1ed18f4")
}

// UnpackUUPSUnauthorizedCallContextError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UUPSUnauthorizedCallContext()
func (wrapperRegistry *WrapperRegistry) UnpackUUPSUnauthorizedCallContextError(raw []byte) (*WrapperRegistryUUPSUnauthorizedCallContext, error) {
	out := new(WrapperRegistryUUPSUnauthorizedCallContext)
	if err := wrapperRegistry.abi.UnpackIntoInterface(out, "UUPSUnauthorizedCallContext", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WrapperRegistryUUPSUnsupportedProxiableUUID represents a UUPSUnsupportedProxiableUUID error raised by the WrapperRegistry contract.
type WrapperRegistryUUPSUnsupportedProxiableUUID struct {
	Slot [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UUPSUnsupportedProxiableUUID(bytes32 slot)
func WrapperRegistryUUPSUnsupportedProxiableUUIDErrorID() common.Hash {
	return common.HexToHash("0xaa1d49a4c084bfa9aeeee2a0be65267a7f19ba7e1476b114dac513d2c14cb563")
}

// UnpackUUPSUnsupportedProxiableUUIDError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UUPSUnsupportedProxiableUUID(bytes32 slot)
func (wrapperRegistry *WrapperRegistry) UnpackUUPSUnsupportedProxiableUUIDError(raw []byte) (*WrapperRegistryUUPSUnsupportedProxiableUUID, error) {
	out := new(WrapperRegistryUUPSUnsupportedProxiableUUID)
	if err := wrapperRegistry.abi.UnpackIntoInterface(out, "UUPSUnsupportedProxiableUUID", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WrapperRegistryUnauthorizedCaller represents a UnauthorizedCaller error raised by the WrapperRegistry contract.
type WrapperRegistryUnauthorizedCaller struct {
	Caller common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UnauthorizedCaller(address caller)
func WrapperRegistryUnauthorizedCallerErrorID() common.Hash {
	return common.HexToHash("0xd86ad9cfaa7befeb3f173b3707581b7abd2af82a8982c91af7ea00f9ee034b67")
}

// UnpackUnauthorizedCallerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UnauthorizedCaller(address caller)
func (wrapperRegistry *WrapperRegistry) UnpackUnauthorizedCallerError(raw []byte) (*WrapperRegistryUnauthorizedCaller, error) {
	out := new(WrapperRegistryUnauthorizedCaller)
	if err := wrapperRegistry.abi.UnpackIntoInterface(out, "UnauthorizedCaller", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WrapperRegistryUpgradeTargetNotApproved represents a UpgradeTargetNotApproved error raised by the WrapperRegistry contract.
type WrapperRegistryUpgradeTargetNotApproved struct {
	Implementation common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UpgradeTargetNotApproved(address implementation)
func WrapperRegistryUpgradeTargetNotApprovedErrorID() common.Hash {
	return common.HexToHash("0xf74d7dd08e97da4111e6fce526692a2faac65a0edb5971a381c0d749e273b4b8")
}

// UnpackUpgradeTargetNotApprovedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UpgradeTargetNotApproved(address implementation)
func (wrapperRegistry *WrapperRegistry) UnpackUpgradeTargetNotApprovedError(raw []byte) (*WrapperRegistryUpgradeTargetNotApproved, error) {
	out := new(WrapperRegistryUpgradeTargetNotApproved)
	if err := wrapperRegistry.abi.UnpackIntoInterface(out, "UpgradeTargetNotApproved", raw); err != nil {
		return nil, err
	}
	return out, nil
}
