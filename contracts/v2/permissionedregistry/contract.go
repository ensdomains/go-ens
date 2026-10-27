// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package permissionedregistry

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

// Grant is an auto generated low-level Go binding around an user-defined struct.
type Grant struct {
	Account    common.Address
	RoleBitmap *big.Int
}

// IPermissionedRegistryState is an auto generated low-level Go binding around an user-defined struct.
type IPermissionedRegistryState struct {
	Status      uint8
	Expiry      uint64
	LatestOwner common.Address
	TokenId     *big.Int
	Resource    *big.Int
}

// PermissionedRegistryMetaData contains all meta data concerning the PermissionedRegistry contract.
var PermissionedRegistryMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractILabelStore\",\"name\":\"labelStore\",\"type\":\"address\"},{\"internalType\":\"contractIRegistryURIRenderer\",\"name\":\"uriRenderer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"namer\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"oldExpiry\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"newExpiry\",\"type\":\"uint64\"}],\"name\":\"CannotReduceExpiry\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"expiry\",\"type\":\"uint64\"}],\"name\":\"CannotSetPastExpiry\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"resource\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"roleBitmap\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"EACCannotGrantRoles\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"resource\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"roleBitmap\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"EACCannotRevokeRoles\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EACInvalidAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roleBitmap\",\"type\":\"uint256\"}],\"name\":\"EACInvalidRoleBitmap\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"resource\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"role\",\"type\":\"uint256\"}],\"name\":\"EACMaxAssignees\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"resource\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"role\",\"type\":\"uint256\"}],\"name\":\"EACMinAssignees\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EACRootResourceNotAllowed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"resource\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"roleBitmap\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"EACUnauthorizedAccountRoles\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC1155InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC1155InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idsLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"valuesLength\",\"type\":\"uint256\"}],\"name\":\"ERC1155InvalidArrayLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC1155InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC1155InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC1155InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC1155MissingApprovalForAll\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"}],\"name\":\"LabelAlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"}],\"name\":\"LabelAlreadyReserved\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"LabelExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"TransferDisallowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferUnsafeUntilRegistryIsEmancipated\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"TransferUnsafeWithMultipleAssignees\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"resource\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldRoleBitmap\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newRoleBitmap\",\"type\":\"uint256\"}],\"name\":\"EACRolesChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"newExpiry\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ExpiryUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"labelHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"expiry\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"LabelRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"labelHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"expiry\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"LabelReserved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"LabelUnregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIRegistry\",\"name\":\"parent\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ParentUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"RegistryCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ResolverUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIRegistry\",\"name\":\"subregistry\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SubregistryUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"oldTokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newTokenId\",\"type\":\"uint256\"}],\"name\":\"TokenRegenerated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"resource\",\"type\":\"uint256\"}],\"name\":\"TokenResource\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"renderer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"URIUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"LABEL_STORE\",\"outputs\":[{\"internalType\":\"contractILabelStore\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ROOT_RESOURCE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"URI_RENDERER\",\"outputs\":[{\"internalType\":\"contractIRegistryURIRenderer\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"canUpgradeFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"}],\"name\":\"findExpiry\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"}],\"name\":\"findOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"}],\"name\":\"findTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"anyId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"roleBitmap\",\"type\":\"uint256\"}],\"name\":\"getAssigneeCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"counts\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mask\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"anyId\",\"type\":\"uint256\"}],\"name\":\"getExpiry\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"anyId\",\"type\":\"uint256\"}],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getParent\",\"outputs\":[{\"internalType\":\"contractIRegistry\",\"name\":\"parent\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"}],\"name\":\"getResolver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"anyId\",\"type\":\"uint256\"}],\"name\":\"getResource\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"anyId\",\"type\":\"uint256\"}],\"name\":\"getState\",\"outputs\":[{\"components\":[{\"internalType\":\"enumIPermissionedRegistry.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"expiry\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"latestOwner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resource\",\"type\":\"uint256\"}],\"internalType\":\"structIPermissionedRegistry.State\",\"name\":\"state\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"anyId\",\"type\":\"uint256\"}],\"name\":\"getStatus\",\"outputs\":[{\"internalType\":\"enumIPermissionedRegistry.Status\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"}],\"name\":\"getSubregistry\",\"outputs\":[{\"internalType\":\"contractIRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"anyId\",\"type\":\"uint256\"}],\"name\":\"getTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"uri_\",\"type\":\"string\"},{\"internalType\":\"contractIRegistryURIRenderer\",\"name\":\"renderer\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"anyId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"roleBitmap\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRoles\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roleBitmap\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRootRoles\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"resource\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"roleBitmap\",\"type\":\"uint256\"}],\"name\":\"hasAssignees\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"anyId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"roleBitmap\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRoles\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roleBitmap\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRootRoles\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"roleBitmap\",\"type\":\"uint256\"}],\"internalType\":\"structGrant[]\",\"name\":\"grants\",\"type\":\"tuple[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"namer\",\"type\":\"address\"}],\"name\":\"isContractNamer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isEmancipated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"resource\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"roleBitmap\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isOnlyAssignee\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"latestOwnerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"contractIRegistry\",\"name\":\"registry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"roleBitmap\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"expiry\",\"type\":\"uint64\"}],\"name\":\"register\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"anyId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"newExpiry\",\"type\":\"uint64\"}],\"name\":\"renew\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"anyId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"roleBitmap\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRoles\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roleBitmap\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRootRoles\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"anyId\",\"type\":\"uint256\"}],\"name\":\"roleCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"anyId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"roles\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIRegistry\",\"name\":\"parent\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"}],\"name\":\"setParent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"anyId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"}],\"name\":\"setResolver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"anyId\",\"type\":\"uint256\"},{\"internalType\":\"contractIRegistry\",\"name\":\"registry\",\"type\":\"address\"}],\"name\":\"setSubregistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uri_\",\"type\":\"string\"},{\"internalType\":\"contractIRegistryURIRenderer\",\"name\":\"renderer\",\"type\":\"address\"}],\"name\":\"setURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"anyId\",\"type\":\"uint256\"}],\"name\":\"unregister\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"unsafeBatchTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"unsafeTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	ID:  "PermissionedRegistry",
}

// PermissionedRegistry is an auto generated Go binding around an Ethereum contract.
type PermissionedRegistry struct {
	abi abi.ABI
}

// GetABI returns the ABI associated with this contract binding.
func (c *PermissionedRegistry) GetABI() abi.ABI {
	return c.abi
}

// NewPermissionedRegistry creates a new instance of PermissionedRegistry.
func NewPermissionedRegistry() *PermissionedRegistry {
	parsed, err := PermissionedRegistryMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &PermissionedRegistry{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *PermissionedRegistry) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address labelStore, address uriRenderer, address namer) returns()
func (permissionedRegistry *PermissionedRegistry) PackConstructor(labelStore common.Address, uriRenderer common.Address, namer common.Address) []byte {
	enc, err := permissionedRegistry.abi.Pack("", labelStore, uriRenderer, namer)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackLABELSTORE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9dbba19d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function LABEL_STORE() view returns(address)
func (permissionedRegistry *PermissionedRegistry) PackLABELSTORE() []byte {
	enc, err := permissionedRegistry.abi.Pack("LABEL_STORE")
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
func (permissionedRegistry *PermissionedRegistry) TryPackLABELSTORE() ([]byte, error) {
	return permissionedRegistry.abi.Pack("LABEL_STORE")
}

// UnpackLABELSTORE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9dbba19d.
//
// Solidity: function LABEL_STORE() view returns(address)
func (permissionedRegistry *PermissionedRegistry) UnpackLABELSTORE(data []byte) (common.Address, error) {
	out, err := permissionedRegistry.abi.Unpack("LABEL_STORE", data)
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
func (permissionedRegistry *PermissionedRegistry) PackROOTRESOURCE() []byte {
	enc, err := permissionedRegistry.abi.Pack("ROOT_RESOURCE")
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
func (permissionedRegistry *PermissionedRegistry) TryPackROOTRESOURCE() ([]byte, error) {
	return permissionedRegistry.abi.Pack("ROOT_RESOURCE")
}

// UnpackROOTRESOURCE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x1c3fc3eb.
//
// Solidity: function ROOT_RESOURCE() view returns(uint256)
func (permissionedRegistry *PermissionedRegistry) UnpackROOTRESOURCE(data []byte) (*big.Int, error) {
	out, err := permissionedRegistry.abi.Unpack("ROOT_RESOURCE", data)
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
func (permissionedRegistry *PermissionedRegistry) PackUPGRADEINTERFACEVERSION() []byte {
	enc, err := permissionedRegistry.abi.Pack("UPGRADE_INTERFACE_VERSION")
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
func (permissionedRegistry *PermissionedRegistry) TryPackUPGRADEINTERFACEVERSION() ([]byte, error) {
	return permissionedRegistry.abi.Pack("UPGRADE_INTERFACE_VERSION")
}

// UnpackUPGRADEINTERFACEVERSION is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (permissionedRegistry *PermissionedRegistry) UnpackUPGRADEINTERFACEVERSION(data []byte) (string, error) {
	out, err := permissionedRegistry.abi.Unpack("UPGRADE_INTERFACE_VERSION", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, nil
}

// PackURIRENDERER is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9c42d7e8.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function URI_RENDERER() view returns(address)
func (permissionedRegistry *PermissionedRegistry) PackURIRENDERER() []byte {
	enc, err := permissionedRegistry.abi.Pack("URI_RENDERER")
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
func (permissionedRegistry *PermissionedRegistry) TryPackURIRENDERER() ([]byte, error) {
	return permissionedRegistry.abi.Pack("URI_RENDERER")
}

// UnpackURIRENDERER is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9c42d7e8.
//
// Solidity: function URI_RENDERER() view returns(address)
func (permissionedRegistry *PermissionedRegistry) UnpackURIRENDERER(data []byte) (common.Address, error) {
	out, err := permissionedRegistry.abi.Unpack("URI_RENDERER", data)
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
func (permissionedRegistry *PermissionedRegistry) PackBalanceOf(account common.Address, id *big.Int) []byte {
	enc, err := permissionedRegistry.abi.Pack("balanceOf", account, id)
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
func (permissionedRegistry *PermissionedRegistry) TryPackBalanceOf(account common.Address, id *big.Int) ([]byte, error) {
	return permissionedRegistry.abi.Pack("balanceOf", account, id)
}

// UnpackBalanceOf is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (permissionedRegistry *PermissionedRegistry) UnpackBalanceOf(data []byte) (*big.Int, error) {
	out, err := permissionedRegistry.abi.Unpack("balanceOf", data)
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
func (permissionedRegistry *PermissionedRegistry) PackBalanceOfBatch(accounts []common.Address, ids []*big.Int) []byte {
	enc, err := permissionedRegistry.abi.Pack("balanceOfBatch", accounts, ids)
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
func (permissionedRegistry *PermissionedRegistry) TryPackBalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]byte, error) {
	return permissionedRegistry.abi.Pack("balanceOfBatch", accounts, ids)
}

// UnpackBalanceOfBatch is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (permissionedRegistry *PermissionedRegistry) UnpackBalanceOfBatch(data []byte) ([]*big.Int, error) {
	out, err := permissionedRegistry.abi.Unpack("balanceOfBatch", data)
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
func (permissionedRegistry *PermissionedRegistry) PackCanUpgradeFrom(arg0 common.Address) []byte {
	enc, err := permissionedRegistry.abi.Pack("canUpgradeFrom", arg0)
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
func (permissionedRegistry *PermissionedRegistry) TryPackCanUpgradeFrom(arg0 common.Address) ([]byte, error) {
	return permissionedRegistry.abi.Pack("canUpgradeFrom", arg0)
}

// UnpackCanUpgradeFrom is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf41a143d.
//
// Solidity: function canUpgradeFrom(address ) pure returns(bool allowed)
func (permissionedRegistry *PermissionedRegistry) UnpackCanUpgradeFrom(data []byte) (bool, error) {
	out, err := permissionedRegistry.abi.Unpack("canUpgradeFrom", data)
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
func (permissionedRegistry *PermissionedRegistry) PackFindExpiry(label string) []byte {
	enc, err := permissionedRegistry.abi.Pack("findExpiry", label)
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
func (permissionedRegistry *PermissionedRegistry) TryPackFindExpiry(label string) ([]byte, error) {
	return permissionedRegistry.abi.Pack("findExpiry", label)
}

// UnpackFindExpiry is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6f537c72.
//
// Solidity: function findExpiry(string label) view returns(uint64)
func (permissionedRegistry *PermissionedRegistry) UnpackFindExpiry(data []byte) (uint64, error) {
	out, err := permissionedRegistry.abi.Unpack("findExpiry", data)
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
func (permissionedRegistry *PermissionedRegistry) PackFindOwner(label string) []byte {
	enc, err := permissionedRegistry.abi.Pack("findOwner", label)
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
func (permissionedRegistry *PermissionedRegistry) TryPackFindOwner(label string) ([]byte, error) {
	return permissionedRegistry.abi.Pack("findOwner", label)
}

// UnpackFindOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x63560a8e.
//
// Solidity: function findOwner(string label) view returns(address)
func (permissionedRegistry *PermissionedRegistry) UnpackFindOwner(data []byte) (common.Address, error) {
	out, err := permissionedRegistry.abi.Unpack("findOwner", data)
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
func (permissionedRegistry *PermissionedRegistry) PackFindTokenId(label string) []byte {
	enc, err := permissionedRegistry.abi.Pack("findTokenId", label)
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
func (permissionedRegistry *PermissionedRegistry) TryPackFindTokenId(label string) ([]byte, error) {
	return permissionedRegistry.abi.Pack("findTokenId", label)
}

// UnpackFindTokenId is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x91b3c037.
//
// Solidity: function findTokenId(string label) view returns(uint256)
func (permissionedRegistry *PermissionedRegistry) UnpackFindTokenId(data []byte) (*big.Int, error) {
	out, err := permissionedRegistry.abi.Unpack("findTokenId", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetAssigneeCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3634f911.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getAssigneeCount(uint256 anyId, uint256 roleBitmap) view returns(uint256 counts, uint256 mask)
func (permissionedRegistry *PermissionedRegistry) PackGetAssigneeCount(anyId *big.Int, roleBitmap *big.Int) []byte {
	enc, err := permissionedRegistry.abi.Pack("getAssigneeCount", anyId, roleBitmap)
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
func (permissionedRegistry *PermissionedRegistry) TryPackGetAssigneeCount(anyId *big.Int, roleBitmap *big.Int) ([]byte, error) {
	return permissionedRegistry.abi.Pack("getAssigneeCount", anyId, roleBitmap)
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
func (permissionedRegistry *PermissionedRegistry) UnpackGetAssigneeCount(data []byte) (GetAssigneeCountOutput, error) {
	out, err := permissionedRegistry.abi.Unpack("getAssigneeCount", data)
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
func (permissionedRegistry *PermissionedRegistry) PackGetExpiry(anyId *big.Int) []byte {
	enc, err := permissionedRegistry.abi.Pack("getExpiry", anyId)
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
func (permissionedRegistry *PermissionedRegistry) TryPackGetExpiry(anyId *big.Int) ([]byte, error) {
	return permissionedRegistry.abi.Pack("getExpiry", anyId)
}

// UnpackGetExpiry is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x13c72608.
//
// Solidity: function getExpiry(uint256 anyId) view returns(uint64)
func (permissionedRegistry *PermissionedRegistry) UnpackGetExpiry(data []byte) (uint64, error) {
	out, err := permissionedRegistry.abi.Unpack("getExpiry", data)
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
func (permissionedRegistry *PermissionedRegistry) PackGetOwner(anyId *big.Int) []byte {
	enc, err := permissionedRegistry.abi.Pack("getOwner", anyId)
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
func (permissionedRegistry *PermissionedRegistry) TryPackGetOwner(anyId *big.Int) ([]byte, error) {
	return permissionedRegistry.abi.Pack("getOwner", anyId)
}

// UnpackGetOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc41a360a.
//
// Solidity: function getOwner(uint256 anyId) view returns(address)
func (permissionedRegistry *PermissionedRegistry) UnpackGetOwner(data []byte) (common.Address, error) {
	out, err := permissionedRegistry.abi.Unpack("getOwner", data)
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
func (permissionedRegistry *PermissionedRegistry) PackGetParent() []byte {
	enc, err := permissionedRegistry.abi.Pack("getParent")
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
func (permissionedRegistry *PermissionedRegistry) TryPackGetParent() ([]byte, error) {
	return permissionedRegistry.abi.Pack("getParent")
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
func (permissionedRegistry *PermissionedRegistry) UnpackGetParent(data []byte) (GetParentOutput, error) {
	out, err := permissionedRegistry.abi.Unpack("getParent", data)
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
func (permissionedRegistry *PermissionedRegistry) PackGetResolver(label string) []byte {
	enc, err := permissionedRegistry.abi.Pack("getResolver", label)
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
func (permissionedRegistry *PermissionedRegistry) TryPackGetResolver(label string) ([]byte, error) {
	return permissionedRegistry.abi.Pack("getResolver", label)
}

// UnpackGetResolver is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe4ae7d77.
//
// Solidity: function getResolver(string label) view returns(address)
func (permissionedRegistry *PermissionedRegistry) UnpackGetResolver(data []byte) (common.Address, error) {
	out, err := permissionedRegistry.abi.Unpack("getResolver", data)
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
func (permissionedRegistry *PermissionedRegistry) PackGetResource(anyId *big.Int) []byte {
	enc, err := permissionedRegistry.abi.Pack("getResource", anyId)
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
func (permissionedRegistry *PermissionedRegistry) TryPackGetResource(anyId *big.Int) ([]byte, error) {
	return permissionedRegistry.abi.Pack("getResource", anyId)
}

// UnpackGetResource is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x1e8fca2d.
//
// Solidity: function getResource(uint256 anyId) view returns(uint256)
func (permissionedRegistry *PermissionedRegistry) UnpackGetResource(data []byte) (*big.Int, error) {
	out, err := permissionedRegistry.abi.Unpack("getResource", data)
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
func (permissionedRegistry *PermissionedRegistry) PackGetState(anyId *big.Int) []byte {
	enc, err := permissionedRegistry.abi.Pack("getState", anyId)
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
func (permissionedRegistry *PermissionedRegistry) TryPackGetState(anyId *big.Int) ([]byte, error) {
	return permissionedRegistry.abi.Pack("getState", anyId)
}

// UnpackGetState is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x44c9af28.
//
// Solidity: function getState(uint256 anyId) view returns((uint8,uint64,address,uint256,uint256) state)
func (permissionedRegistry *PermissionedRegistry) UnpackGetState(data []byte) (IPermissionedRegistryState, error) {
	out, err := permissionedRegistry.abi.Unpack("getState", data)
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
func (permissionedRegistry *PermissionedRegistry) PackGetStatus(anyId *big.Int) []byte {
	enc, err := permissionedRegistry.abi.Pack("getStatus", anyId)
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
func (permissionedRegistry *PermissionedRegistry) TryPackGetStatus(anyId *big.Int) ([]byte, error) {
	return permissionedRegistry.abi.Pack("getStatus", anyId)
}

// UnpackGetStatus is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5c622a0e.
//
// Solidity: function getStatus(uint256 anyId) view returns(uint8)
func (permissionedRegistry *PermissionedRegistry) UnpackGetStatus(data []byte) (uint8, error) {
	out, err := permissionedRegistry.abi.Unpack("getStatus", data)
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
func (permissionedRegistry *PermissionedRegistry) PackGetSubregistry(label string) []byte {
	enc, err := permissionedRegistry.abi.Pack("getSubregistry", label)
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
func (permissionedRegistry *PermissionedRegistry) TryPackGetSubregistry(label string) ([]byte, error) {
	return permissionedRegistry.abi.Pack("getSubregistry", label)
}

// UnpackGetSubregistry is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x35af6216.
//
// Solidity: function getSubregistry(string label) view returns(address)
func (permissionedRegistry *PermissionedRegistry) UnpackGetSubregistry(data []byte) (common.Address, error) {
	out, err := permissionedRegistry.abi.Unpack("getSubregistry", data)
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
func (permissionedRegistry *PermissionedRegistry) PackGetTokenId(anyId *big.Int) []byte {
	enc, err := permissionedRegistry.abi.Pack("getTokenId", anyId)
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
func (permissionedRegistry *PermissionedRegistry) TryPackGetTokenId(anyId *big.Int) ([]byte, error) {
	return permissionedRegistry.abi.Pack("getTokenId", anyId)
}

// UnpackGetTokenId is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x14ff5ea3.
//
// Solidity: function getTokenId(uint256 anyId) view returns(uint256)
func (permissionedRegistry *PermissionedRegistry) UnpackGetTokenId(data []byte) (*big.Int, error) {
	out, err := permissionedRegistry.abi.Unpack("getTokenId", data)
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
func (permissionedRegistry *PermissionedRegistry) PackGetURI() []byte {
	enc, err := permissionedRegistry.abi.Pack("getURI")
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
func (permissionedRegistry *PermissionedRegistry) TryPackGetURI() ([]byte, error) {
	return permissionedRegistry.abi.Pack("getURI")
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
func (permissionedRegistry *PermissionedRegistry) UnpackGetURI(data []byte) (GetURIOutput, error) {
	out, err := permissionedRegistry.abi.Unpack("getURI", data)
	outstruct := new(GetURIOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Uri = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Renderer = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	return *outstruct, nil
}

// PackGrantRoles is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7c300586.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function grantRoles(uint256 anyId, uint256 roleBitmap, address account) returns(bool)
func (permissionedRegistry *PermissionedRegistry) PackGrantRoles(anyId *big.Int, roleBitmap *big.Int, account common.Address) []byte {
	enc, err := permissionedRegistry.abi.Pack("grantRoles", anyId, roleBitmap, account)
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
func (permissionedRegistry *PermissionedRegistry) TryPackGrantRoles(anyId *big.Int, roleBitmap *big.Int, account common.Address) ([]byte, error) {
	return permissionedRegistry.abi.Pack("grantRoles", anyId, roleBitmap, account)
}

// UnpackGrantRoles is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7c300586.
//
// Solidity: function grantRoles(uint256 anyId, uint256 roleBitmap, address account) returns(bool)
func (permissionedRegistry *PermissionedRegistry) UnpackGrantRoles(data []byte) (bool, error) {
	out, err := permissionedRegistry.abi.Unpack("grantRoles", data)
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
func (permissionedRegistry *PermissionedRegistry) PackGrantRootRoles(roleBitmap *big.Int, account common.Address) []byte {
	enc, err := permissionedRegistry.abi.Pack("grantRootRoles", roleBitmap, account)
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
func (permissionedRegistry *PermissionedRegistry) TryPackGrantRootRoles(roleBitmap *big.Int, account common.Address) ([]byte, error) {
	return permissionedRegistry.abi.Pack("grantRootRoles", roleBitmap, account)
}

// UnpackGrantRootRoles is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x072d5d77.
//
// Solidity: function grantRootRoles(uint256 roleBitmap, address account) returns(bool)
func (permissionedRegistry *PermissionedRegistry) UnpackGrantRootRoles(data []byte) (bool, error) {
	out, err := permissionedRegistry.abi.Unpack("grantRootRoles", data)
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
func (permissionedRegistry *PermissionedRegistry) PackHasAssignees(resource *big.Int, roleBitmap *big.Int) []byte {
	enc, err := permissionedRegistry.abi.Pack("hasAssignees", resource, roleBitmap)
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
func (permissionedRegistry *PermissionedRegistry) TryPackHasAssignees(resource *big.Int, roleBitmap *big.Int) ([]byte, error) {
	return permissionedRegistry.abi.Pack("hasAssignees", resource, roleBitmap)
}

// UnpackHasAssignees is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x11b8e00a.
//
// Solidity: function hasAssignees(uint256 resource, uint256 roleBitmap) view returns(bool)
func (permissionedRegistry *PermissionedRegistry) UnpackHasAssignees(data []byte) (bool, error) {
	out, err := permissionedRegistry.abi.Unpack("hasAssignees", data)
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
func (permissionedRegistry *PermissionedRegistry) PackHasRoles(anyId *big.Int, roleBitmap *big.Int, account common.Address) []byte {
	enc, err := permissionedRegistry.abi.Pack("hasRoles", anyId, roleBitmap, account)
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
func (permissionedRegistry *PermissionedRegistry) TryPackHasRoles(anyId *big.Int, roleBitmap *big.Int, account common.Address) ([]byte, error) {
	return permissionedRegistry.abi.Pack("hasRoles", anyId, roleBitmap, account)
}

// UnpackHasRoles is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd3bf89b1.
//
// Solidity: function hasRoles(uint256 anyId, uint256 roleBitmap, address account) view returns(bool)
func (permissionedRegistry *PermissionedRegistry) UnpackHasRoles(data []byte) (bool, error) {
	out, err := permissionedRegistry.abi.Unpack("hasRoles", data)
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
func (permissionedRegistry *PermissionedRegistry) PackHasRootRoles(roleBitmap *big.Int, account common.Address) []byte {
	enc, err := permissionedRegistry.abi.Pack("hasRootRoles", roleBitmap, account)
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
func (permissionedRegistry *PermissionedRegistry) TryPackHasRootRoles(roleBitmap *big.Int, account common.Address) ([]byte, error) {
	return permissionedRegistry.abi.Pack("hasRootRoles", roleBitmap, account)
}

// UnpackHasRootRoles is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x781ef8db.
//
// Solidity: function hasRootRoles(uint256 roleBitmap, address account) view returns(bool)
func (permissionedRegistry *PermissionedRegistry) UnpackHasRootRoles(data []byte) (bool, error) {
	out, err := permissionedRegistry.abi.Unpack("hasRootRoles", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x37cb53a8.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function initialize((address,uint256)[] grants) returns()
func (permissionedRegistry *PermissionedRegistry) PackInitialize(grants []Grant) []byte {
	enc, err := permissionedRegistry.abi.Pack("initialize", grants)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x37cb53a8.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function initialize((address,uint256)[] grants) returns()
func (permissionedRegistry *PermissionedRegistry) TryPackInitialize(grants []Grant) ([]byte, error) {
	return permissionedRegistry.abi.Pack("initialize", grants)
}

// PackIsApprovedForAll is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe985e9c5.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (permissionedRegistry *PermissionedRegistry) PackIsApprovedForAll(account common.Address, operator common.Address) []byte {
	enc, err := permissionedRegistry.abi.Pack("isApprovedForAll", account, operator)
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
func (permissionedRegistry *PermissionedRegistry) TryPackIsApprovedForAll(account common.Address, operator common.Address) ([]byte, error) {
	return permissionedRegistry.abi.Pack("isApprovedForAll", account, operator)
}

// UnpackIsApprovedForAll is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (permissionedRegistry *PermissionedRegistry) UnpackIsApprovedForAll(data []byte) (bool, error) {
	out, err := permissionedRegistry.abi.Unpack("isApprovedForAll", data)
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
func (permissionedRegistry *PermissionedRegistry) PackIsContractNamer(namer common.Address) []byte {
	enc, err := permissionedRegistry.abi.Pack("isContractNamer", namer)
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
func (permissionedRegistry *PermissionedRegistry) TryPackIsContractNamer(namer common.Address) ([]byte, error) {
	return permissionedRegistry.abi.Pack("isContractNamer", namer)
}

// UnpackIsContractNamer is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6f3ff726.
//
// Solidity: function isContractNamer(address namer) view returns(bool)
func (permissionedRegistry *PermissionedRegistry) UnpackIsContractNamer(data []byte) (bool, error) {
	out, err := permissionedRegistry.abi.Unpack("isContractNamer", data)
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
// Solidity: function isEmancipated() view returns(bool)
func (permissionedRegistry *PermissionedRegistry) PackIsEmancipated() []byte {
	enc, err := permissionedRegistry.abi.Pack("isEmancipated")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsEmancipated is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x955266f5.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isEmancipated() view returns(bool)
func (permissionedRegistry *PermissionedRegistry) TryPackIsEmancipated() ([]byte, error) {
	return permissionedRegistry.abi.Pack("isEmancipated")
}

// UnpackIsEmancipated is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x955266f5.
//
// Solidity: function isEmancipated() view returns(bool)
func (permissionedRegistry *PermissionedRegistry) UnpackIsEmancipated(data []byte) (bool, error) {
	out, err := permissionedRegistry.abi.Unpack("isEmancipated", data)
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
func (permissionedRegistry *PermissionedRegistry) PackIsOnlyAssignee(resource *big.Int, roleBitmap *big.Int, account common.Address) []byte {
	enc, err := permissionedRegistry.abi.Pack("isOnlyAssignee", resource, roleBitmap, account)
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
func (permissionedRegistry *PermissionedRegistry) TryPackIsOnlyAssignee(resource *big.Int, roleBitmap *big.Int, account common.Address) ([]byte, error) {
	return permissionedRegistry.abi.Pack("isOnlyAssignee", resource, roleBitmap, account)
}

// UnpackIsOnlyAssignee is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8e77c95f.
//
// Solidity: function isOnlyAssignee(uint256 resource, uint256 roleBitmap, address account) view returns(bool)
func (permissionedRegistry *PermissionedRegistry) UnpackIsOnlyAssignee(data []byte) (bool, error) {
	out, err := permissionedRegistry.abi.Unpack("isOnlyAssignee", data)
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
func (permissionedRegistry *PermissionedRegistry) PackLatestOwnerOf(tokenId *big.Int) []byte {
	enc, err := permissionedRegistry.abi.Pack("latestOwnerOf", tokenId)
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
func (permissionedRegistry *PermissionedRegistry) TryPackLatestOwnerOf(tokenId *big.Int) ([]byte, error) {
	return permissionedRegistry.abi.Pack("latestOwnerOf", tokenId)
}

// UnpackLatestOwnerOf is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbd242bcb.
//
// Solidity: function latestOwnerOf(uint256 tokenId) view returns(address)
func (permissionedRegistry *PermissionedRegistry) UnpackLatestOwnerOf(data []byte) (common.Address, error) {
	out, err := permissionedRegistry.abi.Unpack("latestOwnerOf", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackOwnerOf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6352211e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (permissionedRegistry *PermissionedRegistry) PackOwnerOf(tokenId *big.Int) []byte {
	enc, err := permissionedRegistry.abi.Pack("ownerOf", tokenId)
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
func (permissionedRegistry *PermissionedRegistry) TryPackOwnerOf(tokenId *big.Int) ([]byte, error) {
	return permissionedRegistry.abi.Pack("ownerOf", tokenId)
}

// UnpackOwnerOf is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (permissionedRegistry *PermissionedRegistry) UnpackOwnerOf(data []byte) (common.Address, error) {
	out, err := permissionedRegistry.abi.Unpack("ownerOf", data)
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
func (permissionedRegistry *PermissionedRegistry) PackProxiableUUID() []byte {
	enc, err := permissionedRegistry.abi.Pack("proxiableUUID")
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
func (permissionedRegistry *PermissionedRegistry) TryPackProxiableUUID() ([]byte, error) {
	return permissionedRegistry.abi.Pack("proxiableUUID")
}

// UnpackProxiableUUID is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (permissionedRegistry *PermissionedRegistry) UnpackProxiableUUID(data []byte) ([32]byte, error) {
	out, err := permissionedRegistry.abi.Unpack("proxiableUUID", data)
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
// Solidity: function register(string label, address owner, address registry, address resolver, uint256 roleBitmap, uint64 expiry) returns(uint256)
func (permissionedRegistry *PermissionedRegistry) PackRegister(label string, owner common.Address, registry common.Address, resolver common.Address, roleBitmap *big.Int, expiry uint64) []byte {
	enc, err := permissionedRegistry.abi.Pack("register", label, owner, registry, resolver, roleBitmap, expiry)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRegister is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x85f3e643.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function register(string label, address owner, address registry, address resolver, uint256 roleBitmap, uint64 expiry) returns(uint256)
func (permissionedRegistry *PermissionedRegistry) TryPackRegister(label string, owner common.Address, registry common.Address, resolver common.Address, roleBitmap *big.Int, expiry uint64) ([]byte, error) {
	return permissionedRegistry.abi.Pack("register", label, owner, registry, resolver, roleBitmap, expiry)
}

// UnpackRegister is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x85f3e643.
//
// Solidity: function register(string label, address owner, address registry, address resolver, uint256 roleBitmap, uint64 expiry) returns(uint256)
func (permissionedRegistry *PermissionedRegistry) UnpackRegister(data []byte) (*big.Int, error) {
	out, err := permissionedRegistry.abi.Unpack("register", data)
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
func (permissionedRegistry *PermissionedRegistry) PackRenew(anyId *big.Int, newExpiry uint64) []byte {
	enc, err := permissionedRegistry.abi.Pack("renew", anyId, newExpiry)
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
func (permissionedRegistry *PermissionedRegistry) TryPackRenew(anyId *big.Int, newExpiry uint64) ([]byte, error) {
	return permissionedRegistry.abi.Pack("renew", anyId, newExpiry)
}

// PackRevokeRoles is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdfa70d8b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function revokeRoles(uint256 anyId, uint256 roleBitmap, address account) returns(bool)
func (permissionedRegistry *PermissionedRegistry) PackRevokeRoles(anyId *big.Int, roleBitmap *big.Int, account common.Address) []byte {
	enc, err := permissionedRegistry.abi.Pack("revokeRoles", anyId, roleBitmap, account)
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
func (permissionedRegistry *PermissionedRegistry) TryPackRevokeRoles(anyId *big.Int, roleBitmap *big.Int, account common.Address) ([]byte, error) {
	return permissionedRegistry.abi.Pack("revokeRoles", anyId, roleBitmap, account)
}

// UnpackRevokeRoles is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xdfa70d8b.
//
// Solidity: function revokeRoles(uint256 anyId, uint256 roleBitmap, address account) returns(bool)
func (permissionedRegistry *PermissionedRegistry) UnpackRevokeRoles(data []byte) (bool, error) {
	out, err := permissionedRegistry.abi.Unpack("revokeRoles", data)
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
func (permissionedRegistry *PermissionedRegistry) PackRevokeRootRoles(roleBitmap *big.Int, account common.Address) []byte {
	enc, err := permissionedRegistry.abi.Pack("revokeRootRoles", roleBitmap, account)
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
func (permissionedRegistry *PermissionedRegistry) TryPackRevokeRootRoles(roleBitmap *big.Int, account common.Address) ([]byte, error) {
	return permissionedRegistry.abi.Pack("revokeRootRoles", roleBitmap, account)
}

// UnpackRevokeRootRoles is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xce156e82.
//
// Solidity: function revokeRootRoles(uint256 roleBitmap, address account) returns(bool)
func (permissionedRegistry *PermissionedRegistry) UnpackRevokeRootRoles(data []byte) (bool, error) {
	out, err := permissionedRegistry.abi.Unpack("revokeRootRoles", data)
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
func (permissionedRegistry *PermissionedRegistry) PackRoleCount(anyId *big.Int) []byte {
	enc, err := permissionedRegistry.abi.Pack("roleCount", anyId)
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
func (permissionedRegistry *PermissionedRegistry) TryPackRoleCount(anyId *big.Int) ([]byte, error) {
	return permissionedRegistry.abi.Pack("roleCount", anyId)
}

// UnpackRoleCount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x2f27fa24.
//
// Solidity: function roleCount(uint256 anyId) view returns(uint256)
func (permissionedRegistry *PermissionedRegistry) UnpackRoleCount(data []byte) (*big.Int, error) {
	out, err := permissionedRegistry.abi.Unpack("roleCount", data)
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
func (permissionedRegistry *PermissionedRegistry) PackRoles(anyId *big.Int, account common.Address) []byte {
	enc, err := permissionedRegistry.abi.Pack("roles", anyId, account)
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
func (permissionedRegistry *PermissionedRegistry) TryPackRoles(anyId *big.Int, account common.Address) ([]byte, error) {
	return permissionedRegistry.abi.Pack("roles", anyId, account)
}

// UnpackRoles is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5adf4724.
//
// Solidity: function roles(uint256 anyId, address account) view returns(uint256)
func (permissionedRegistry *PermissionedRegistry) UnpackRoles(data []byte) (*big.Int, error) {
	out, err := permissionedRegistry.abi.Unpack("roles", data)
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
func (permissionedRegistry *PermissionedRegistry) PackSafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, values []*big.Int, data []byte) []byte {
	enc, err := permissionedRegistry.abi.Pack("safeBatchTransferFrom", from, to, ids, values, data)
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
func (permissionedRegistry *PermissionedRegistry) TryPackSafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, values []*big.Int, data []byte) ([]byte, error) {
	return permissionedRegistry.abi.Pack("safeBatchTransferFrom", from, to, ids, values, data)
}

// PackSafeTransferFrom is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf242432a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 value, bytes data) returns()
func (permissionedRegistry *PermissionedRegistry) PackSafeTransferFrom(from common.Address, to common.Address, id *big.Int, value *big.Int, data []byte) []byte {
	enc, err := permissionedRegistry.abi.Pack("safeTransferFrom", from, to, id, value, data)
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
func (permissionedRegistry *PermissionedRegistry) TryPackSafeTransferFrom(from common.Address, to common.Address, id *big.Int, value *big.Int, data []byte) ([]byte, error) {
	return permissionedRegistry.abi.Pack("safeTransferFrom", from, to, id, value, data)
}

// PackSetApprovalForAll is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa22cb465.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (permissionedRegistry *PermissionedRegistry) PackSetApprovalForAll(operator common.Address, approved bool) []byte {
	enc, err := permissionedRegistry.abi.Pack("setApprovalForAll", operator, approved)
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
func (permissionedRegistry *PermissionedRegistry) TryPackSetApprovalForAll(operator common.Address, approved bool) ([]byte, error) {
	return permissionedRegistry.abi.Pack("setApprovalForAll", operator, approved)
}

// PackSetParent is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5357263f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setParent(address parent, string label) returns()
func (permissionedRegistry *PermissionedRegistry) PackSetParent(parent common.Address, label string) []byte {
	enc, err := permissionedRegistry.abi.Pack("setParent", parent, label)
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
func (permissionedRegistry *PermissionedRegistry) TryPackSetParent(parent common.Address, label string) ([]byte, error) {
	return permissionedRegistry.abi.Pack("setParent", parent, label)
}

// PackSetResolver is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbc7b6d62.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setResolver(uint256 anyId, address resolver) returns()
func (permissionedRegistry *PermissionedRegistry) PackSetResolver(anyId *big.Int, resolver common.Address) []byte {
	enc, err := permissionedRegistry.abi.Pack("setResolver", anyId, resolver)
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
func (permissionedRegistry *PermissionedRegistry) TryPackSetResolver(anyId *big.Int, resolver common.Address) ([]byte, error) {
	return permissionedRegistry.abi.Pack("setResolver", anyId, resolver)
}

// PackSetSubregistry is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x341ec559.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setSubregistry(uint256 anyId, address registry) returns()
func (permissionedRegistry *PermissionedRegistry) PackSetSubregistry(anyId *big.Int, registry common.Address) []byte {
	enc, err := permissionedRegistry.abi.Pack("setSubregistry", anyId, registry)
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
func (permissionedRegistry *PermissionedRegistry) TryPackSetSubregistry(anyId *big.Int, registry common.Address) ([]byte, error) {
	return permissionedRegistry.abi.Pack("setSubregistry", anyId, registry)
}

// PackSetURI is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x48688f95.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setURI(string uri_, address renderer) returns()
func (permissionedRegistry *PermissionedRegistry) PackSetURI(uri string, renderer common.Address) []byte {
	enc, err := permissionedRegistry.abi.Pack("setURI", uri, renderer)
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
func (permissionedRegistry *PermissionedRegistry) TryPackSetURI(uri string, renderer common.Address) ([]byte, error) {
	return permissionedRegistry.abi.Pack("setURI", uri, renderer)
}

// PackSupportsInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01ffc9a7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (permissionedRegistry *PermissionedRegistry) PackSupportsInterface(interfaceId [4]byte) []byte {
	enc, err := permissionedRegistry.abi.Pack("supportsInterface", interfaceId)
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
func (permissionedRegistry *PermissionedRegistry) TryPackSupportsInterface(interfaceId [4]byte) ([]byte, error) {
	return permissionedRegistry.abi.Pack("supportsInterface", interfaceId)
}

// UnpackSupportsInterface is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (permissionedRegistry *PermissionedRegistry) UnpackSupportsInterface(data []byte) (bool, error) {
	out, err := permissionedRegistry.abi.Unpack("supportsInterface", data)
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
func (permissionedRegistry *PermissionedRegistry) PackUnregister(anyId *big.Int) []byte {
	enc, err := permissionedRegistry.abi.Pack("unregister", anyId)
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
func (permissionedRegistry *PermissionedRegistry) TryPackUnregister(anyId *big.Int) ([]byte, error) {
	return permissionedRegistry.abi.Pack("unregister", anyId)
}

// PackUnsafeBatchTransfer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3dd5ee86.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function unsafeBatchTransfer(address to, uint256[] tokenIds, bytes data) returns()
func (permissionedRegistry *PermissionedRegistry) PackUnsafeBatchTransfer(to common.Address, tokenIds []*big.Int, data []byte) []byte {
	enc, err := permissionedRegistry.abi.Pack("unsafeBatchTransfer", to, tokenIds, data)
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
func (permissionedRegistry *PermissionedRegistry) TryPackUnsafeBatchTransfer(to common.Address, tokenIds []*big.Int, data []byte) ([]byte, error) {
	return permissionedRegistry.abi.Pack("unsafeBatchTransfer", to, tokenIds, data)
}

// PackUnsafeTransfer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x087b0790.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function unsafeTransfer(address to, uint256 tokenId, bytes data) returns()
func (permissionedRegistry *PermissionedRegistry) PackUnsafeTransfer(to common.Address, tokenId *big.Int, data []byte) []byte {
	enc, err := permissionedRegistry.abi.Pack("unsafeTransfer", to, tokenId, data)
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
func (permissionedRegistry *PermissionedRegistry) TryPackUnsafeTransfer(to common.Address, tokenId *big.Int, data []byte) ([]byte, error) {
	return permissionedRegistry.abi.Pack("unsafeTransfer", to, tokenId, data)
}

// PackUpgradeToAndCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4f1ef286.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (permissionedRegistry *PermissionedRegistry) PackUpgradeToAndCall(newImplementation common.Address, data []byte) []byte {
	enc, err := permissionedRegistry.abi.Pack("upgradeToAndCall", newImplementation, data)
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
func (permissionedRegistry *PermissionedRegistry) TryPackUpgradeToAndCall(newImplementation common.Address, data []byte) ([]byte, error) {
	return permissionedRegistry.abi.Pack("upgradeToAndCall", newImplementation, data)
}

// PackUri is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0e89341c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function uri(uint256 tokenId) view returns(string)
func (permissionedRegistry *PermissionedRegistry) PackUri(tokenId *big.Int) []byte {
	enc, err := permissionedRegistry.abi.Pack("uri", tokenId)
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
func (permissionedRegistry *PermissionedRegistry) TryPackUri(tokenId *big.Int) ([]byte, error) {
	return permissionedRegistry.abi.Pack("uri", tokenId)
}

// UnpackUri is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x0e89341c.
//
// Solidity: function uri(uint256 tokenId) view returns(string)
func (permissionedRegistry *PermissionedRegistry) UnpackUri(data []byte) (string, error) {
	out, err := permissionedRegistry.abi.Unpack("uri", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, nil
}

// PermissionedRegistryApprovalForAll represents a ApprovalForAll event raised by the PermissionedRegistry contract.
type PermissionedRegistryApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      *types.Log // Blockchain specific contextual infos
}

const PermissionedRegistryApprovalForAllEventName = "ApprovalForAll"

// ContractEventName returns the user-defined event name.
func (PermissionedRegistryApprovalForAll) ContractEventName() string {
	return PermissionedRegistryApprovalForAllEventName
}

// UnpackApprovalForAllEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (permissionedRegistry *PermissionedRegistry) UnpackApprovalForAllEvent(log *types.Log) (*PermissionedRegistryApprovalForAll, error) {
	event := "ApprovalForAll"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != permissionedRegistry.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(PermissionedRegistryApprovalForAll)
	if len(log.Data) > 0 {
		if err := permissionedRegistry.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range permissionedRegistry.abi.Events[event].Inputs {
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

// PermissionedRegistryEACRolesChanged represents a EACRolesChanged event raised by the PermissionedRegistry contract.
type PermissionedRegistryEACRolesChanged struct {
	Resource      *big.Int
	Account       common.Address
	OldRoleBitmap *big.Int
	NewRoleBitmap *big.Int
	Raw           *types.Log // Blockchain specific contextual infos
}

const PermissionedRegistryEACRolesChangedEventName = "EACRolesChanged"

// ContractEventName returns the user-defined event name.
func (PermissionedRegistryEACRolesChanged) ContractEventName() string {
	return PermissionedRegistryEACRolesChangedEventName
}

// UnpackEACRolesChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event EACRolesChanged(uint256 indexed resource, address indexed account, uint256 oldRoleBitmap, uint256 newRoleBitmap)
func (permissionedRegistry *PermissionedRegistry) UnpackEACRolesChangedEvent(log *types.Log) (*PermissionedRegistryEACRolesChanged, error) {
	event := "EACRolesChanged"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != permissionedRegistry.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(PermissionedRegistryEACRolesChanged)
	if len(log.Data) > 0 {
		if err := permissionedRegistry.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range permissionedRegistry.abi.Events[event].Inputs {
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

// PermissionedRegistryExpiryUpdated represents a ExpiryUpdated event raised by the PermissionedRegistry contract.
type PermissionedRegistryExpiryUpdated struct {
	TokenId   *big.Int
	NewExpiry uint64
	Sender    common.Address
	Raw       *types.Log // Blockchain specific contextual infos
}

const PermissionedRegistryExpiryUpdatedEventName = "ExpiryUpdated"

// ContractEventName returns the user-defined event name.
func (PermissionedRegistryExpiryUpdated) ContractEventName() string {
	return PermissionedRegistryExpiryUpdatedEventName
}

// UnpackExpiryUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ExpiryUpdated(uint256 indexed tokenId, uint64 indexed newExpiry, address indexed sender)
func (permissionedRegistry *PermissionedRegistry) UnpackExpiryUpdatedEvent(log *types.Log) (*PermissionedRegistryExpiryUpdated, error) {
	event := "ExpiryUpdated"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != permissionedRegistry.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(PermissionedRegistryExpiryUpdated)
	if len(log.Data) > 0 {
		if err := permissionedRegistry.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range permissionedRegistry.abi.Events[event].Inputs {
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

// PermissionedRegistryInitialized represents a Initialized event raised by the PermissionedRegistry contract.
type PermissionedRegistryInitialized struct {
	Version uint64
	Raw     *types.Log // Blockchain specific contextual infos
}

const PermissionedRegistryInitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (PermissionedRegistryInitialized) ContractEventName() string {
	return PermissionedRegistryInitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint64 version)
func (permissionedRegistry *PermissionedRegistry) UnpackInitializedEvent(log *types.Log) (*PermissionedRegistryInitialized, error) {
	event := "Initialized"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != permissionedRegistry.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(PermissionedRegistryInitialized)
	if len(log.Data) > 0 {
		if err := permissionedRegistry.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range permissionedRegistry.abi.Events[event].Inputs {
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

// PermissionedRegistryLabelRegistered represents a LabelRegistered event raised by the PermissionedRegistry contract.
type PermissionedRegistryLabelRegistered struct {
	TokenId   *big.Int
	LabelHash [32]byte
	Label     string
	Owner     common.Address
	Expiry    uint64
	Sender    common.Address
	Raw       *types.Log // Blockchain specific contextual infos
}

const PermissionedRegistryLabelRegisteredEventName = "LabelRegistered"

// ContractEventName returns the user-defined event name.
func (PermissionedRegistryLabelRegistered) ContractEventName() string {
	return PermissionedRegistryLabelRegisteredEventName
}

// UnpackLabelRegisteredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event LabelRegistered(uint256 indexed tokenId, bytes32 indexed labelHash, string label, address owner, uint64 expiry, address indexed sender)
func (permissionedRegistry *PermissionedRegistry) UnpackLabelRegisteredEvent(log *types.Log) (*PermissionedRegistryLabelRegistered, error) {
	event := "LabelRegistered"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != permissionedRegistry.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(PermissionedRegistryLabelRegistered)
	if len(log.Data) > 0 {
		if err := permissionedRegistry.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range permissionedRegistry.abi.Events[event].Inputs {
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

// PermissionedRegistryLabelReserved represents a LabelReserved event raised by the PermissionedRegistry contract.
type PermissionedRegistryLabelReserved struct {
	TokenId   *big.Int
	LabelHash [32]byte
	Label     string
	Expiry    uint64
	Sender    common.Address
	Raw       *types.Log // Blockchain specific contextual infos
}

const PermissionedRegistryLabelReservedEventName = "LabelReserved"

// ContractEventName returns the user-defined event name.
func (PermissionedRegistryLabelReserved) ContractEventName() string {
	return PermissionedRegistryLabelReservedEventName
}

// UnpackLabelReservedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event LabelReserved(uint256 indexed tokenId, bytes32 indexed labelHash, string label, uint64 expiry, address indexed sender)
func (permissionedRegistry *PermissionedRegistry) UnpackLabelReservedEvent(log *types.Log) (*PermissionedRegistryLabelReserved, error) {
	event := "LabelReserved"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != permissionedRegistry.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(PermissionedRegistryLabelReserved)
	if len(log.Data) > 0 {
		if err := permissionedRegistry.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range permissionedRegistry.abi.Events[event].Inputs {
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

// PermissionedRegistryLabelUnregistered represents a LabelUnregistered event raised by the PermissionedRegistry contract.
type PermissionedRegistryLabelUnregistered struct {
	TokenId *big.Int
	Sender  common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const PermissionedRegistryLabelUnregisteredEventName = "LabelUnregistered"

// ContractEventName returns the user-defined event name.
func (PermissionedRegistryLabelUnregistered) ContractEventName() string {
	return PermissionedRegistryLabelUnregisteredEventName
}

// UnpackLabelUnregisteredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event LabelUnregistered(uint256 indexed tokenId, address indexed sender)
func (permissionedRegistry *PermissionedRegistry) UnpackLabelUnregisteredEvent(log *types.Log) (*PermissionedRegistryLabelUnregistered, error) {
	event := "LabelUnregistered"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != permissionedRegistry.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(PermissionedRegistryLabelUnregistered)
	if len(log.Data) > 0 {
		if err := permissionedRegistry.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range permissionedRegistry.abi.Events[event].Inputs {
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

// PermissionedRegistryParentUpdated represents a ParentUpdated event raised by the PermissionedRegistry contract.
type PermissionedRegistryParentUpdated struct {
	Parent common.Address
	Label  string
	Sender common.Address
	Raw    *types.Log // Blockchain specific contextual infos
}

const PermissionedRegistryParentUpdatedEventName = "ParentUpdated"

// ContractEventName returns the user-defined event name.
func (PermissionedRegistryParentUpdated) ContractEventName() string {
	return PermissionedRegistryParentUpdatedEventName
}

// UnpackParentUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ParentUpdated(address indexed parent, string label, address indexed sender)
func (permissionedRegistry *PermissionedRegistry) UnpackParentUpdatedEvent(log *types.Log) (*PermissionedRegistryParentUpdated, error) {
	event := "ParentUpdated"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != permissionedRegistry.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(PermissionedRegistryParentUpdated)
	if len(log.Data) > 0 {
		if err := permissionedRegistry.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range permissionedRegistry.abi.Events[event].Inputs {
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

// PermissionedRegistryRegistryCreated represents a RegistryCreated event raised by the PermissionedRegistry contract.
type PermissionedRegistryRegistryCreated struct {
	Raw *types.Log // Blockchain specific contextual infos
}

const PermissionedRegistryRegistryCreatedEventName = "RegistryCreated"

// ContractEventName returns the user-defined event name.
func (PermissionedRegistryRegistryCreated) ContractEventName() string {
	return PermissionedRegistryRegistryCreatedEventName
}

// UnpackRegistryCreatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RegistryCreated()
func (permissionedRegistry *PermissionedRegistry) UnpackRegistryCreatedEvent(log *types.Log) (*PermissionedRegistryRegistryCreated, error) {
	event := "RegistryCreated"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != permissionedRegistry.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(PermissionedRegistryRegistryCreated)
	if len(log.Data) > 0 {
		if err := permissionedRegistry.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range permissionedRegistry.abi.Events[event].Inputs {
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

// PermissionedRegistryResolverUpdated represents a ResolverUpdated event raised by the PermissionedRegistry contract.
type PermissionedRegistryResolverUpdated struct {
	TokenId  *big.Int
	Resolver common.Address
	Sender   common.Address
	Raw      *types.Log // Blockchain specific contextual infos
}

const PermissionedRegistryResolverUpdatedEventName = "ResolverUpdated"

// ContractEventName returns the user-defined event name.
func (PermissionedRegistryResolverUpdated) ContractEventName() string {
	return PermissionedRegistryResolverUpdatedEventName
}

// UnpackResolverUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ResolverUpdated(uint256 indexed tokenId, address indexed resolver, address indexed sender)
func (permissionedRegistry *PermissionedRegistry) UnpackResolverUpdatedEvent(log *types.Log) (*PermissionedRegistryResolverUpdated, error) {
	event := "ResolverUpdated"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != permissionedRegistry.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(PermissionedRegistryResolverUpdated)
	if len(log.Data) > 0 {
		if err := permissionedRegistry.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range permissionedRegistry.abi.Events[event].Inputs {
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

// PermissionedRegistrySubregistryUpdated represents a SubregistryUpdated event raised by the PermissionedRegistry contract.
type PermissionedRegistrySubregistryUpdated struct {
	TokenId     *big.Int
	Subregistry common.Address
	Sender      common.Address
	Raw         *types.Log // Blockchain specific contextual infos
}

const PermissionedRegistrySubregistryUpdatedEventName = "SubregistryUpdated"

// ContractEventName returns the user-defined event name.
func (PermissionedRegistrySubregistryUpdated) ContractEventName() string {
	return PermissionedRegistrySubregistryUpdatedEventName
}

// UnpackSubregistryUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event SubregistryUpdated(uint256 indexed tokenId, address indexed subregistry, address indexed sender)
func (permissionedRegistry *PermissionedRegistry) UnpackSubregistryUpdatedEvent(log *types.Log) (*PermissionedRegistrySubregistryUpdated, error) {
	event := "SubregistryUpdated"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != permissionedRegistry.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(PermissionedRegistrySubregistryUpdated)
	if len(log.Data) > 0 {
		if err := permissionedRegistry.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range permissionedRegistry.abi.Events[event].Inputs {
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

// PermissionedRegistryTokenRegenerated represents a TokenRegenerated event raised by the PermissionedRegistry contract.
type PermissionedRegistryTokenRegenerated struct {
	OldTokenId *big.Int
	NewTokenId *big.Int
	Raw        *types.Log // Blockchain specific contextual infos
}

const PermissionedRegistryTokenRegeneratedEventName = "TokenRegenerated"

// ContractEventName returns the user-defined event name.
func (PermissionedRegistryTokenRegenerated) ContractEventName() string {
	return PermissionedRegistryTokenRegeneratedEventName
}

// UnpackTokenRegeneratedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TokenRegenerated(uint256 indexed oldTokenId, uint256 indexed newTokenId)
func (permissionedRegistry *PermissionedRegistry) UnpackTokenRegeneratedEvent(log *types.Log) (*PermissionedRegistryTokenRegenerated, error) {
	event := "TokenRegenerated"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != permissionedRegistry.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(PermissionedRegistryTokenRegenerated)
	if len(log.Data) > 0 {
		if err := permissionedRegistry.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range permissionedRegistry.abi.Events[event].Inputs {
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

// PermissionedRegistryTokenResource represents a TokenResource event raised by the PermissionedRegistry contract.
type PermissionedRegistryTokenResource struct {
	TokenId  *big.Int
	Resource *big.Int
	Raw      *types.Log // Blockchain specific contextual infos
}

const PermissionedRegistryTokenResourceEventName = "TokenResource"

// ContractEventName returns the user-defined event name.
func (PermissionedRegistryTokenResource) ContractEventName() string {
	return PermissionedRegistryTokenResourceEventName
}

// UnpackTokenResourceEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TokenResource(uint256 indexed tokenId, uint256 indexed resource)
func (permissionedRegistry *PermissionedRegistry) UnpackTokenResourceEvent(log *types.Log) (*PermissionedRegistryTokenResource, error) {
	event := "TokenResource"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != permissionedRegistry.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(PermissionedRegistryTokenResource)
	if len(log.Data) > 0 {
		if err := permissionedRegistry.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range permissionedRegistry.abi.Events[event].Inputs {
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

// PermissionedRegistryTransferBatch represents a TransferBatch event raised by the PermissionedRegistry contract.
type PermissionedRegistryTransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      *types.Log // Blockchain specific contextual infos
}

const PermissionedRegistryTransferBatchEventName = "TransferBatch"

// ContractEventName returns the user-defined event name.
func (PermissionedRegistryTransferBatch) ContractEventName() string {
	return PermissionedRegistryTransferBatchEventName
}

// UnpackTransferBatchEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (permissionedRegistry *PermissionedRegistry) UnpackTransferBatchEvent(log *types.Log) (*PermissionedRegistryTransferBatch, error) {
	event := "TransferBatch"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != permissionedRegistry.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(PermissionedRegistryTransferBatch)
	if len(log.Data) > 0 {
		if err := permissionedRegistry.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range permissionedRegistry.abi.Events[event].Inputs {
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

// PermissionedRegistryTransferSingle represents a TransferSingle event raised by the PermissionedRegistry contract.
type PermissionedRegistryTransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      *types.Log // Blockchain specific contextual infos
}

const PermissionedRegistryTransferSingleEventName = "TransferSingle"

// ContractEventName returns the user-defined event name.
func (PermissionedRegistryTransferSingle) ContractEventName() string {
	return PermissionedRegistryTransferSingleEventName
}

// UnpackTransferSingleEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (permissionedRegistry *PermissionedRegistry) UnpackTransferSingleEvent(log *types.Log) (*PermissionedRegistryTransferSingle, error) {
	event := "TransferSingle"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != permissionedRegistry.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(PermissionedRegistryTransferSingle)
	if len(log.Data) > 0 {
		if err := permissionedRegistry.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range permissionedRegistry.abi.Events[event].Inputs {
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

// PermissionedRegistryURI represents a URI event raised by the PermissionedRegistry contract.
type PermissionedRegistryURI struct {
	Value string
	Id    *big.Int
	Raw   *types.Log // Blockchain specific contextual infos
}

const PermissionedRegistryURIEventName = "URI"

// ContractEventName returns the user-defined event name.
func (PermissionedRegistryURI) ContractEventName() string {
	return PermissionedRegistryURIEventName
}

// UnpackURIEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event URI(string value, uint256 indexed id)
func (permissionedRegistry *PermissionedRegistry) UnpackURIEvent(log *types.Log) (*PermissionedRegistryURI, error) {
	event := "URI"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != permissionedRegistry.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(PermissionedRegistryURI)
	if len(log.Data) > 0 {
		if err := permissionedRegistry.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range permissionedRegistry.abi.Events[event].Inputs {
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

// PermissionedRegistryURIUpdated represents a URIUpdated event raised by the PermissionedRegistry contract.
type PermissionedRegistryURIUpdated struct {
	Uri      string
	Renderer common.Address
	Sender   common.Address
	Raw      *types.Log // Blockchain specific contextual infos
}

const PermissionedRegistryURIUpdatedEventName = "URIUpdated"

// ContractEventName returns the user-defined event name.
func (PermissionedRegistryURIUpdated) ContractEventName() string {
	return PermissionedRegistryURIUpdatedEventName
}

// UnpackURIUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event URIUpdated(string uri, address renderer, address indexed sender)
func (permissionedRegistry *PermissionedRegistry) UnpackURIUpdatedEvent(log *types.Log) (*PermissionedRegistryURIUpdated, error) {
	event := "URIUpdated"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != permissionedRegistry.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(PermissionedRegistryURIUpdated)
	if len(log.Data) > 0 {
		if err := permissionedRegistry.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range permissionedRegistry.abi.Events[event].Inputs {
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

// PermissionedRegistryUpgraded represents a Upgraded event raised by the PermissionedRegistry contract.
type PermissionedRegistryUpgraded struct {
	Implementation common.Address
	Raw            *types.Log // Blockchain specific contextual infos
}

const PermissionedRegistryUpgradedEventName = "Upgraded"

// ContractEventName returns the user-defined event name.
func (PermissionedRegistryUpgraded) ContractEventName() string {
	return PermissionedRegistryUpgradedEventName
}

// UnpackUpgradedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Upgraded(address indexed implementation)
func (permissionedRegistry *PermissionedRegistry) UnpackUpgradedEvent(log *types.Log) (*PermissionedRegistryUpgraded, error) {
	event := "Upgraded"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != permissionedRegistry.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(PermissionedRegistryUpgraded)
	if len(log.Data) > 0 {
		if err := permissionedRegistry.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range permissionedRegistry.abi.Events[event].Inputs {
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
func (permissionedRegistry *PermissionedRegistry) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], permissionedRegistry.abi.Errors["AddressEmptyCode"].ID.Bytes()[:4]) {
		return permissionedRegistry.UnpackAddressEmptyCodeError(raw[4:])
	}
	if bytes.Equal(raw[:4], permissionedRegistry.abi.Errors["CannotReduceExpiry"].ID.Bytes()[:4]) {
		return permissionedRegistry.UnpackCannotReduceExpiryError(raw[4:])
	}
	if bytes.Equal(raw[:4], permissionedRegistry.abi.Errors["CannotSetPastExpiry"].ID.Bytes()[:4]) {
		return permissionedRegistry.UnpackCannotSetPastExpiryError(raw[4:])
	}
	if bytes.Equal(raw[:4], permissionedRegistry.abi.Errors["EACCannotGrantRoles"].ID.Bytes()[:4]) {
		return permissionedRegistry.UnpackEACCannotGrantRolesError(raw[4:])
	}
	if bytes.Equal(raw[:4], permissionedRegistry.abi.Errors["EACCannotRevokeRoles"].ID.Bytes()[:4]) {
		return permissionedRegistry.UnpackEACCannotRevokeRolesError(raw[4:])
	}
	if bytes.Equal(raw[:4], permissionedRegistry.abi.Errors["EACInvalidAccount"].ID.Bytes()[:4]) {
		return permissionedRegistry.UnpackEACInvalidAccountError(raw[4:])
	}
	if bytes.Equal(raw[:4], permissionedRegistry.abi.Errors["EACInvalidRoleBitmap"].ID.Bytes()[:4]) {
		return permissionedRegistry.UnpackEACInvalidRoleBitmapError(raw[4:])
	}
	if bytes.Equal(raw[:4], permissionedRegistry.abi.Errors["EACMaxAssignees"].ID.Bytes()[:4]) {
		return permissionedRegistry.UnpackEACMaxAssigneesError(raw[4:])
	}
	if bytes.Equal(raw[:4], permissionedRegistry.abi.Errors["EACMinAssignees"].ID.Bytes()[:4]) {
		return permissionedRegistry.UnpackEACMinAssigneesError(raw[4:])
	}
	if bytes.Equal(raw[:4], permissionedRegistry.abi.Errors["EACRootResourceNotAllowed"].ID.Bytes()[:4]) {
		return permissionedRegistry.UnpackEACRootResourceNotAllowedError(raw[4:])
	}
	if bytes.Equal(raw[:4], permissionedRegistry.abi.Errors["EACUnauthorizedAccountRoles"].ID.Bytes()[:4]) {
		return permissionedRegistry.UnpackEACUnauthorizedAccountRolesError(raw[4:])
	}
	if bytes.Equal(raw[:4], permissionedRegistry.abi.Errors["ERC1155InsufficientBalance"].ID.Bytes()[:4]) {
		return permissionedRegistry.UnpackERC1155InsufficientBalanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], permissionedRegistry.abi.Errors["ERC1155InvalidApprover"].ID.Bytes()[:4]) {
		return permissionedRegistry.UnpackERC1155InvalidApproverError(raw[4:])
	}
	if bytes.Equal(raw[:4], permissionedRegistry.abi.Errors["ERC1155InvalidArrayLength"].ID.Bytes()[:4]) {
		return permissionedRegistry.UnpackERC1155InvalidArrayLengthError(raw[4:])
	}
	if bytes.Equal(raw[:4], permissionedRegistry.abi.Errors["ERC1155InvalidOperator"].ID.Bytes()[:4]) {
		return permissionedRegistry.UnpackERC1155InvalidOperatorError(raw[4:])
	}
	if bytes.Equal(raw[:4], permissionedRegistry.abi.Errors["ERC1155InvalidReceiver"].ID.Bytes()[:4]) {
		return permissionedRegistry.UnpackERC1155InvalidReceiverError(raw[4:])
	}
	if bytes.Equal(raw[:4], permissionedRegistry.abi.Errors["ERC1155InvalidSender"].ID.Bytes()[:4]) {
		return permissionedRegistry.UnpackERC1155InvalidSenderError(raw[4:])
	}
	if bytes.Equal(raw[:4], permissionedRegistry.abi.Errors["ERC1155MissingApprovalForAll"].ID.Bytes()[:4]) {
		return permissionedRegistry.UnpackERC1155MissingApprovalForAllError(raw[4:])
	}
	if bytes.Equal(raw[:4], permissionedRegistry.abi.Errors["ERC1967InvalidImplementation"].ID.Bytes()[:4]) {
		return permissionedRegistry.UnpackERC1967InvalidImplementationError(raw[4:])
	}
	if bytes.Equal(raw[:4], permissionedRegistry.abi.Errors["ERC1967NonPayable"].ID.Bytes()[:4]) {
		return permissionedRegistry.UnpackERC1967NonPayableError(raw[4:])
	}
	if bytes.Equal(raw[:4], permissionedRegistry.abi.Errors["FailedCall"].ID.Bytes()[:4]) {
		return permissionedRegistry.UnpackFailedCallError(raw[4:])
	}
	if bytes.Equal(raw[:4], permissionedRegistry.abi.Errors["InvalidInitialization"].ID.Bytes()[:4]) {
		return permissionedRegistry.UnpackInvalidInitializationError(raw[4:])
	}
	if bytes.Equal(raw[:4], permissionedRegistry.abi.Errors["InvalidOwner"].ID.Bytes()[:4]) {
		return permissionedRegistry.UnpackInvalidOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], permissionedRegistry.abi.Errors["LabelAlreadyRegistered"].ID.Bytes()[:4]) {
		return permissionedRegistry.UnpackLabelAlreadyRegisteredError(raw[4:])
	}
	if bytes.Equal(raw[:4], permissionedRegistry.abi.Errors["LabelAlreadyReserved"].ID.Bytes()[:4]) {
		return permissionedRegistry.UnpackLabelAlreadyReservedError(raw[4:])
	}
	if bytes.Equal(raw[:4], permissionedRegistry.abi.Errors["LabelExpired"].ID.Bytes()[:4]) {
		return permissionedRegistry.UnpackLabelExpiredError(raw[4:])
	}
	if bytes.Equal(raw[:4], permissionedRegistry.abi.Errors["NotInitializing"].ID.Bytes()[:4]) {
		return permissionedRegistry.UnpackNotInitializingError(raw[4:])
	}
	if bytes.Equal(raw[:4], permissionedRegistry.abi.Errors["TransferDisallowed"].ID.Bytes()[:4]) {
		return permissionedRegistry.UnpackTransferDisallowedError(raw[4:])
	}
	if bytes.Equal(raw[:4], permissionedRegistry.abi.Errors["TransferUnsafeUntilRegistryIsEmancipated"].ID.Bytes()[:4]) {
		return permissionedRegistry.UnpackTransferUnsafeUntilRegistryIsEmancipatedError(raw[4:])
	}
	if bytes.Equal(raw[:4], permissionedRegistry.abi.Errors["TransferUnsafeWithMultipleAssignees"].ID.Bytes()[:4]) {
		return permissionedRegistry.UnpackTransferUnsafeWithMultipleAssigneesError(raw[4:])
	}
	if bytes.Equal(raw[:4], permissionedRegistry.abi.Errors["UUPSUnauthorizedCallContext"].ID.Bytes()[:4]) {
		return permissionedRegistry.UnpackUUPSUnauthorizedCallContextError(raw[4:])
	}
	if bytes.Equal(raw[:4], permissionedRegistry.abi.Errors["UUPSUnsupportedProxiableUUID"].ID.Bytes()[:4]) {
		return permissionedRegistry.UnpackUUPSUnsupportedProxiableUUIDError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// PermissionedRegistryAddressEmptyCode represents a AddressEmptyCode error raised by the PermissionedRegistry contract.
type PermissionedRegistryAddressEmptyCode struct {
	Target common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AddressEmptyCode(address target)
func PermissionedRegistryAddressEmptyCodeErrorID() common.Hash {
	return common.HexToHash("0x9996b315c842ff135b8fc4a08ad5df1c344efbc03d2687aecc0678050d2aac89")
}

// UnpackAddressEmptyCodeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AddressEmptyCode(address target)
func (permissionedRegistry *PermissionedRegistry) UnpackAddressEmptyCodeError(raw []byte) (*PermissionedRegistryAddressEmptyCode, error) {
	out := new(PermissionedRegistryAddressEmptyCode)
	if err := permissionedRegistry.abi.UnpackIntoInterface(out, "AddressEmptyCode", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionedRegistryCannotReduceExpiry represents a CannotReduceExpiry error raised by the PermissionedRegistry contract.
type PermissionedRegistryCannotReduceExpiry struct {
	OldExpiry uint64
	NewExpiry uint64
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error CannotReduceExpiry(uint64 oldExpiry, uint64 newExpiry)
func PermissionedRegistryCannotReduceExpiryErrorID() common.Hash {
	return common.HexToHash("0x68c1425a85303a923d9963b6b77886de6c657b9cc761a422fa2090222a51b5ca")
}

// UnpackCannotReduceExpiryError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error CannotReduceExpiry(uint64 oldExpiry, uint64 newExpiry)
func (permissionedRegistry *PermissionedRegistry) UnpackCannotReduceExpiryError(raw []byte) (*PermissionedRegistryCannotReduceExpiry, error) {
	out := new(PermissionedRegistryCannotReduceExpiry)
	if err := permissionedRegistry.abi.UnpackIntoInterface(out, "CannotReduceExpiry", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionedRegistryCannotSetPastExpiry represents a CannotSetPastExpiry error raised by the PermissionedRegistry contract.
type PermissionedRegistryCannotSetPastExpiry struct {
	Expiry uint64
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error CannotSetPastExpiry(uint64 expiry)
func PermissionedRegistryCannotSetPastExpiryErrorID() common.Hash {
	return common.HexToHash("0xf1d446c394c7afd3028f17f9fc3b3d460c3e997e3b2a7fc80b87d86295b1a896")
}

// UnpackCannotSetPastExpiryError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error CannotSetPastExpiry(uint64 expiry)
func (permissionedRegistry *PermissionedRegistry) UnpackCannotSetPastExpiryError(raw []byte) (*PermissionedRegistryCannotSetPastExpiry, error) {
	out := new(PermissionedRegistryCannotSetPastExpiry)
	if err := permissionedRegistry.abi.UnpackIntoInterface(out, "CannotSetPastExpiry", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionedRegistryEACCannotGrantRoles represents a EACCannotGrantRoles error raised by the PermissionedRegistry contract.
type PermissionedRegistryEACCannotGrantRoles struct {
	Resource   *big.Int
	RoleBitmap *big.Int
	Account    common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error EACCannotGrantRoles(uint256 resource, uint256 roleBitmap, address account)
func PermissionedRegistryEACCannotGrantRolesErrorID() common.Hash {
	return common.HexToHash("0xd1a3b355154fa30324f856ca6b25179bcc6c9d2936e28429b1b4293509e97f4a")
}

// UnpackEACCannotGrantRolesError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error EACCannotGrantRoles(uint256 resource, uint256 roleBitmap, address account)
func (permissionedRegistry *PermissionedRegistry) UnpackEACCannotGrantRolesError(raw []byte) (*PermissionedRegistryEACCannotGrantRoles, error) {
	out := new(PermissionedRegistryEACCannotGrantRoles)
	if err := permissionedRegistry.abi.UnpackIntoInterface(out, "EACCannotGrantRoles", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionedRegistryEACCannotRevokeRoles represents a EACCannotRevokeRoles error raised by the PermissionedRegistry contract.
type PermissionedRegistryEACCannotRevokeRoles struct {
	Resource   *big.Int
	RoleBitmap *big.Int
	Account    common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error EACCannotRevokeRoles(uint256 resource, uint256 roleBitmap, address account)
func PermissionedRegistryEACCannotRevokeRolesErrorID() common.Hash {
	return common.HexToHash("0xa604e318421f3889d41e189ce5a1bbb63773eb6844a6d371417f9eed8f6cb153")
}

// UnpackEACCannotRevokeRolesError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error EACCannotRevokeRoles(uint256 resource, uint256 roleBitmap, address account)
func (permissionedRegistry *PermissionedRegistry) UnpackEACCannotRevokeRolesError(raw []byte) (*PermissionedRegistryEACCannotRevokeRoles, error) {
	out := new(PermissionedRegistryEACCannotRevokeRoles)
	if err := permissionedRegistry.abi.UnpackIntoInterface(out, "EACCannotRevokeRoles", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionedRegistryEACInvalidAccount represents a EACInvalidAccount error raised by the PermissionedRegistry contract.
type PermissionedRegistryEACInvalidAccount struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error EACInvalidAccount()
func PermissionedRegistryEACInvalidAccountErrorID() common.Hash {
	return common.HexToHash("0xec3fc5924c67c1c1507473878ec47d207e41f4b88da4d7a835c6e5043b63da51")
}

// UnpackEACInvalidAccountError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error EACInvalidAccount()
func (permissionedRegistry *PermissionedRegistry) UnpackEACInvalidAccountError(raw []byte) (*PermissionedRegistryEACInvalidAccount, error) {
	out := new(PermissionedRegistryEACInvalidAccount)
	if err := permissionedRegistry.abi.UnpackIntoInterface(out, "EACInvalidAccount", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionedRegistryEACInvalidRoleBitmap represents a EACInvalidRoleBitmap error raised by the PermissionedRegistry contract.
type PermissionedRegistryEACInvalidRoleBitmap struct {
	RoleBitmap *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error EACInvalidRoleBitmap(uint256 roleBitmap)
func PermissionedRegistryEACInvalidRoleBitmapErrorID() common.Hash {
	return common.HexToHash("0x2a7b2d20191e44fd4e4227272caf529612678184b58adb291afe5231bf5b2734")
}

// UnpackEACInvalidRoleBitmapError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error EACInvalidRoleBitmap(uint256 roleBitmap)
func (permissionedRegistry *PermissionedRegistry) UnpackEACInvalidRoleBitmapError(raw []byte) (*PermissionedRegistryEACInvalidRoleBitmap, error) {
	out := new(PermissionedRegistryEACInvalidRoleBitmap)
	if err := permissionedRegistry.abi.UnpackIntoInterface(out, "EACInvalidRoleBitmap", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionedRegistryEACMaxAssignees represents a EACMaxAssignees error raised by the PermissionedRegistry contract.
type PermissionedRegistryEACMaxAssignees struct {
	Resource *big.Int
	Role     *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error EACMaxAssignees(uint256 resource, uint256 role)
func PermissionedRegistryEACMaxAssigneesErrorID() common.Hash {
	return common.HexToHash("0xf91653484f73ddf3683475f443ccba4f564362c5896702ddf28889a64960b109")
}

// UnpackEACMaxAssigneesError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error EACMaxAssignees(uint256 resource, uint256 role)
func (permissionedRegistry *PermissionedRegistry) UnpackEACMaxAssigneesError(raw []byte) (*PermissionedRegistryEACMaxAssignees, error) {
	out := new(PermissionedRegistryEACMaxAssignees)
	if err := permissionedRegistry.abi.UnpackIntoInterface(out, "EACMaxAssignees", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionedRegistryEACMinAssignees represents a EACMinAssignees error raised by the PermissionedRegistry contract.
type PermissionedRegistryEACMinAssignees struct {
	Resource *big.Int
	Role     *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error EACMinAssignees(uint256 resource, uint256 role)
func PermissionedRegistryEACMinAssigneesErrorID() common.Hash {
	return common.HexToHash("0x1f80c19ba45f0369eb8ace81d631dde7c47ae7c144205c7e77e1555ff7bb091b")
}

// UnpackEACMinAssigneesError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error EACMinAssignees(uint256 resource, uint256 role)
func (permissionedRegistry *PermissionedRegistry) UnpackEACMinAssigneesError(raw []byte) (*PermissionedRegistryEACMinAssignees, error) {
	out := new(PermissionedRegistryEACMinAssignees)
	if err := permissionedRegistry.abi.UnpackIntoInterface(out, "EACMinAssignees", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionedRegistryEACRootResourceNotAllowed represents a EACRootResourceNotAllowed error raised by the PermissionedRegistry contract.
type PermissionedRegistryEACRootResourceNotAllowed struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error EACRootResourceNotAllowed()
func PermissionedRegistryEACRootResourceNotAllowedErrorID() common.Hash {
	return common.HexToHash("0xc2842458b56deb79dcc9694a98069c21c928dfa5456787ba7ae1728a0d863238")
}

// UnpackEACRootResourceNotAllowedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error EACRootResourceNotAllowed()
func (permissionedRegistry *PermissionedRegistry) UnpackEACRootResourceNotAllowedError(raw []byte) (*PermissionedRegistryEACRootResourceNotAllowed, error) {
	out := new(PermissionedRegistryEACRootResourceNotAllowed)
	if err := permissionedRegistry.abi.UnpackIntoInterface(out, "EACRootResourceNotAllowed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionedRegistryEACUnauthorizedAccountRoles represents a EACUnauthorizedAccountRoles error raised by the PermissionedRegistry contract.
type PermissionedRegistryEACUnauthorizedAccountRoles struct {
	Resource   *big.Int
	RoleBitmap *big.Int
	Account    common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error EACUnauthorizedAccountRoles(uint256 resource, uint256 roleBitmap, address account)
func PermissionedRegistryEACUnauthorizedAccountRolesErrorID() common.Hash {
	return common.HexToHash("0x4b27a13377de460ddbf323b4163e94f59f4092e2fe9befdabaf9bd06c2dd939b")
}

// UnpackEACUnauthorizedAccountRolesError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error EACUnauthorizedAccountRoles(uint256 resource, uint256 roleBitmap, address account)
func (permissionedRegistry *PermissionedRegistry) UnpackEACUnauthorizedAccountRolesError(raw []byte) (*PermissionedRegistryEACUnauthorizedAccountRoles, error) {
	out := new(PermissionedRegistryEACUnauthorizedAccountRoles)
	if err := permissionedRegistry.abi.UnpackIntoInterface(out, "EACUnauthorizedAccountRoles", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionedRegistryERC1155InsufficientBalance represents a ERC1155InsufficientBalance error raised by the PermissionedRegistry contract.
type PermissionedRegistryERC1155InsufficientBalance struct {
	Sender  common.Address
	Balance *big.Int
	Needed  *big.Int
	TokenId *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1155InsufficientBalance(address sender, uint256 balance, uint256 needed, uint256 tokenId)
func PermissionedRegistryERC1155InsufficientBalanceErrorID() common.Hash {
	return common.HexToHash("0x03dee4c573c982787b5f3537d6323ffaca9d864448aa6bd828ada9e5d0837036")
}

// UnpackERC1155InsufficientBalanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1155InsufficientBalance(address sender, uint256 balance, uint256 needed, uint256 tokenId)
func (permissionedRegistry *PermissionedRegistry) UnpackERC1155InsufficientBalanceError(raw []byte) (*PermissionedRegistryERC1155InsufficientBalance, error) {
	out := new(PermissionedRegistryERC1155InsufficientBalance)
	if err := permissionedRegistry.abi.UnpackIntoInterface(out, "ERC1155InsufficientBalance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionedRegistryERC1155InvalidApprover represents a ERC1155InvalidApprover error raised by the PermissionedRegistry contract.
type PermissionedRegistryERC1155InvalidApprover struct {
	Approver common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1155InvalidApprover(address approver)
func PermissionedRegistryERC1155InvalidApproverErrorID() common.Hash {
	return common.HexToHash("0x3e31884e33c33ce0039d1905e3c252950ae3b95240f36d4fff81f5ff6752ef99")
}

// UnpackERC1155InvalidApproverError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1155InvalidApprover(address approver)
func (permissionedRegistry *PermissionedRegistry) UnpackERC1155InvalidApproverError(raw []byte) (*PermissionedRegistryERC1155InvalidApprover, error) {
	out := new(PermissionedRegistryERC1155InvalidApprover)
	if err := permissionedRegistry.abi.UnpackIntoInterface(out, "ERC1155InvalidApprover", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionedRegistryERC1155InvalidArrayLength represents a ERC1155InvalidArrayLength error raised by the PermissionedRegistry contract.
type PermissionedRegistryERC1155InvalidArrayLength struct {
	IdsLength    *big.Int
	ValuesLength *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1155InvalidArrayLength(uint256 idsLength, uint256 valuesLength)
func PermissionedRegistryERC1155InvalidArrayLengthErrorID() common.Hash {
	return common.HexToHash("0x5b0599913619cfa5633692652638ed25cafcd079c9beae8c251b12c23dcc83f2")
}

// UnpackERC1155InvalidArrayLengthError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1155InvalidArrayLength(uint256 idsLength, uint256 valuesLength)
func (permissionedRegistry *PermissionedRegistry) UnpackERC1155InvalidArrayLengthError(raw []byte) (*PermissionedRegistryERC1155InvalidArrayLength, error) {
	out := new(PermissionedRegistryERC1155InvalidArrayLength)
	if err := permissionedRegistry.abi.UnpackIntoInterface(out, "ERC1155InvalidArrayLength", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionedRegistryERC1155InvalidOperator represents a ERC1155InvalidOperator error raised by the PermissionedRegistry contract.
type PermissionedRegistryERC1155InvalidOperator struct {
	Operator common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1155InvalidOperator(address operator)
func PermissionedRegistryERC1155InvalidOperatorErrorID() common.Hash {
	return common.HexToHash("0xced3e10010b9d2aa24827119d0db4a8feec73aea48b4b3e470d8a9f3ff723569")
}

// UnpackERC1155InvalidOperatorError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1155InvalidOperator(address operator)
func (permissionedRegistry *PermissionedRegistry) UnpackERC1155InvalidOperatorError(raw []byte) (*PermissionedRegistryERC1155InvalidOperator, error) {
	out := new(PermissionedRegistryERC1155InvalidOperator)
	if err := permissionedRegistry.abi.UnpackIntoInterface(out, "ERC1155InvalidOperator", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionedRegistryERC1155InvalidReceiver represents a ERC1155InvalidReceiver error raised by the PermissionedRegistry contract.
type PermissionedRegistryERC1155InvalidReceiver struct {
	Receiver common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1155InvalidReceiver(address receiver)
func PermissionedRegistryERC1155InvalidReceiverErrorID() common.Hash {
	return common.HexToHash("0x57f447ceed621d9e134e26de5772c88799abb7322ce2a87f95dce247d47105c6")
}

// UnpackERC1155InvalidReceiverError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1155InvalidReceiver(address receiver)
func (permissionedRegistry *PermissionedRegistry) UnpackERC1155InvalidReceiverError(raw []byte) (*PermissionedRegistryERC1155InvalidReceiver, error) {
	out := new(PermissionedRegistryERC1155InvalidReceiver)
	if err := permissionedRegistry.abi.UnpackIntoInterface(out, "ERC1155InvalidReceiver", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionedRegistryERC1155InvalidSender represents a ERC1155InvalidSender error raised by the PermissionedRegistry contract.
type PermissionedRegistryERC1155InvalidSender struct {
	Sender common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1155InvalidSender(address sender)
func PermissionedRegistryERC1155InvalidSenderErrorID() common.Hash {
	return common.HexToHash("0x01a83514e94b34009110b75cac6742ba33bd7c62f18a3616bafea52855d3b175")
}

// UnpackERC1155InvalidSenderError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1155InvalidSender(address sender)
func (permissionedRegistry *PermissionedRegistry) UnpackERC1155InvalidSenderError(raw []byte) (*PermissionedRegistryERC1155InvalidSender, error) {
	out := new(PermissionedRegistryERC1155InvalidSender)
	if err := permissionedRegistry.abi.UnpackIntoInterface(out, "ERC1155InvalidSender", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionedRegistryERC1155MissingApprovalForAll represents a ERC1155MissingApprovalForAll error raised by the PermissionedRegistry contract.
type PermissionedRegistryERC1155MissingApprovalForAll struct {
	Operator common.Address
	Owner    common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1155MissingApprovalForAll(address operator, address owner)
func PermissionedRegistryERC1155MissingApprovalForAllErrorID() common.Hash {
	return common.HexToHash("0xe237d922be9fac42efeaaaffb42cc6b57e0ff95d94a1b74daeff69adc7657754")
}

// UnpackERC1155MissingApprovalForAllError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1155MissingApprovalForAll(address operator, address owner)
func (permissionedRegistry *PermissionedRegistry) UnpackERC1155MissingApprovalForAllError(raw []byte) (*PermissionedRegistryERC1155MissingApprovalForAll, error) {
	out := new(PermissionedRegistryERC1155MissingApprovalForAll)
	if err := permissionedRegistry.abi.UnpackIntoInterface(out, "ERC1155MissingApprovalForAll", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionedRegistryERC1967InvalidImplementation represents a ERC1967InvalidImplementation error raised by the PermissionedRegistry contract.
type PermissionedRegistryERC1967InvalidImplementation struct {
	Implementation common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1967InvalidImplementation(address implementation)
func PermissionedRegistryERC1967InvalidImplementationErrorID() common.Hash {
	return common.HexToHash("0x4c9c8ce3ceb3130f17f7cdba48d89b5b0129f266a8bac114e6e315a41879b617")
}

// UnpackERC1967InvalidImplementationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1967InvalidImplementation(address implementation)
func (permissionedRegistry *PermissionedRegistry) UnpackERC1967InvalidImplementationError(raw []byte) (*PermissionedRegistryERC1967InvalidImplementation, error) {
	out := new(PermissionedRegistryERC1967InvalidImplementation)
	if err := permissionedRegistry.abi.UnpackIntoInterface(out, "ERC1967InvalidImplementation", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionedRegistryERC1967NonPayable represents a ERC1967NonPayable error raised by the PermissionedRegistry contract.
type PermissionedRegistryERC1967NonPayable struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1967NonPayable()
func PermissionedRegistryERC1967NonPayableErrorID() common.Hash {
	return common.HexToHash("0xb398979fa84f543c8e222f17890372c487baf85e062276c127fef521eea7224b")
}

// UnpackERC1967NonPayableError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1967NonPayable()
func (permissionedRegistry *PermissionedRegistry) UnpackERC1967NonPayableError(raw []byte) (*PermissionedRegistryERC1967NonPayable, error) {
	out := new(PermissionedRegistryERC1967NonPayable)
	if err := permissionedRegistry.abi.UnpackIntoInterface(out, "ERC1967NonPayable", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionedRegistryFailedCall represents a FailedCall error raised by the PermissionedRegistry contract.
type PermissionedRegistryFailedCall struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FailedCall()
func PermissionedRegistryFailedCallErrorID() common.Hash {
	return common.HexToHash("0xd6bda27508c0fb6d8a39b4b122878dab26f731a7d4e4abe711dd3731899052a4")
}

// UnpackFailedCallError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FailedCall()
func (permissionedRegistry *PermissionedRegistry) UnpackFailedCallError(raw []byte) (*PermissionedRegistryFailedCall, error) {
	out := new(PermissionedRegistryFailedCall)
	if err := permissionedRegistry.abi.UnpackIntoInterface(out, "FailedCall", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionedRegistryInvalidInitialization represents a InvalidInitialization error raised by the PermissionedRegistry contract.
type PermissionedRegistryInvalidInitialization struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidInitialization()
func PermissionedRegistryInvalidInitializationErrorID() common.Hash {
	return common.HexToHash("0xf92ee8a957075833165f68c320933b1a1294aafc84ee6e0dd3fb178008f9aaf5")
}

// UnpackInvalidInitializationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidInitialization()
func (permissionedRegistry *PermissionedRegistry) UnpackInvalidInitializationError(raw []byte) (*PermissionedRegistryInvalidInitialization, error) {
	out := new(PermissionedRegistryInvalidInitialization)
	if err := permissionedRegistry.abi.UnpackIntoInterface(out, "InvalidInitialization", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionedRegistryInvalidOwner represents a InvalidOwner error raised by the PermissionedRegistry contract.
type PermissionedRegistryInvalidOwner struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidOwner()
func PermissionedRegistryInvalidOwnerErrorID() common.Hash {
	return common.HexToHash("0x49e27cffb37b1ca4a9bf5318243b2014d13f940af232b8552c208bdea15739da")
}

// UnpackInvalidOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidOwner()
func (permissionedRegistry *PermissionedRegistry) UnpackInvalidOwnerError(raw []byte) (*PermissionedRegistryInvalidOwner, error) {
	out := new(PermissionedRegistryInvalidOwner)
	if err := permissionedRegistry.abi.UnpackIntoInterface(out, "InvalidOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionedRegistryLabelAlreadyRegistered represents a LabelAlreadyRegistered error raised by the PermissionedRegistry contract.
type PermissionedRegistryLabelAlreadyRegistered struct {
	Label string
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error LabelAlreadyRegistered(string label)
func PermissionedRegistryLabelAlreadyRegisteredErrorID() common.Hash {
	return common.HexToHash("0xdef545a48127e18e04d05455886a7023f5b6c21f969846ca4ef9c9aa7cc2267b")
}

// UnpackLabelAlreadyRegisteredError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error LabelAlreadyRegistered(string label)
func (permissionedRegistry *PermissionedRegistry) UnpackLabelAlreadyRegisteredError(raw []byte) (*PermissionedRegistryLabelAlreadyRegistered, error) {
	out := new(PermissionedRegistryLabelAlreadyRegistered)
	if err := permissionedRegistry.abi.UnpackIntoInterface(out, "LabelAlreadyRegistered", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionedRegistryLabelAlreadyReserved represents a LabelAlreadyReserved error raised by the PermissionedRegistry contract.
type PermissionedRegistryLabelAlreadyReserved struct {
	Label string
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error LabelAlreadyReserved(string label)
func PermissionedRegistryLabelAlreadyReservedErrorID() common.Hash {
	return common.HexToHash("0xf60759e04b9d51b7561cf8055ffc1d1b1c222340d6daaf37c7a4c5819b3b9207")
}

// UnpackLabelAlreadyReservedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error LabelAlreadyReserved(string label)
func (permissionedRegistry *PermissionedRegistry) UnpackLabelAlreadyReservedError(raw []byte) (*PermissionedRegistryLabelAlreadyReserved, error) {
	out := new(PermissionedRegistryLabelAlreadyReserved)
	if err := permissionedRegistry.abi.UnpackIntoInterface(out, "LabelAlreadyReserved", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionedRegistryLabelExpired represents a LabelExpired error raised by the PermissionedRegistry contract.
type PermissionedRegistryLabelExpired struct {
	TokenId *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error LabelExpired(uint256 tokenId)
func PermissionedRegistryLabelExpiredErrorID() common.Hash {
	return common.HexToHash("0xc44e2374535590feeaeb825026475cb21e2c4eac2832b552c8811b602a69f016")
}

// UnpackLabelExpiredError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error LabelExpired(uint256 tokenId)
func (permissionedRegistry *PermissionedRegistry) UnpackLabelExpiredError(raw []byte) (*PermissionedRegistryLabelExpired, error) {
	out := new(PermissionedRegistryLabelExpired)
	if err := permissionedRegistry.abi.UnpackIntoInterface(out, "LabelExpired", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionedRegistryNotInitializing represents a NotInitializing error raised by the PermissionedRegistry contract.
type PermissionedRegistryNotInitializing struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotInitializing()
func PermissionedRegistryNotInitializingErrorID() common.Hash {
	return common.HexToHash("0xd7e6bcf8597daa127dc9f0048d2f08d5ef140a2cb659feabd700beff1f7a8302")
}

// UnpackNotInitializingError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotInitializing()
func (permissionedRegistry *PermissionedRegistry) UnpackNotInitializingError(raw []byte) (*PermissionedRegistryNotInitializing, error) {
	out := new(PermissionedRegistryNotInitializing)
	if err := permissionedRegistry.abi.UnpackIntoInterface(out, "NotInitializing", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionedRegistryTransferDisallowed represents a TransferDisallowed error raised by the PermissionedRegistry contract.
type PermissionedRegistryTransferDisallowed struct {
	TokenId *big.Int
	From    common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TransferDisallowed(uint256 tokenId, address from)
func PermissionedRegistryTransferDisallowedErrorID() common.Hash {
	return common.HexToHash("0xe58f6d5a5faf9cca464192e088054deae4ec88a7ee4aafacd834f996253c270d")
}

// UnpackTransferDisallowedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TransferDisallowed(uint256 tokenId, address from)
func (permissionedRegistry *PermissionedRegistry) UnpackTransferDisallowedError(raw []byte) (*PermissionedRegistryTransferDisallowed, error) {
	out := new(PermissionedRegistryTransferDisallowed)
	if err := permissionedRegistry.abi.UnpackIntoInterface(out, "TransferDisallowed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionedRegistryTransferUnsafeUntilRegistryIsEmancipated represents a TransferUnsafeUntilRegistryIsEmancipated error raised by the PermissionedRegistry contract.
type PermissionedRegistryTransferUnsafeUntilRegistryIsEmancipated struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TransferUnsafeUntilRegistryIsEmancipated()
func PermissionedRegistryTransferUnsafeUntilRegistryIsEmancipatedErrorID() common.Hash {
	return common.HexToHash("0x54838f98c10e7309e5345cfa4bd39b2f43a7576ced98070099678ab413097390")
}

// UnpackTransferUnsafeUntilRegistryIsEmancipatedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TransferUnsafeUntilRegistryIsEmancipated()
func (permissionedRegistry *PermissionedRegistry) UnpackTransferUnsafeUntilRegistryIsEmancipatedError(raw []byte) (*PermissionedRegistryTransferUnsafeUntilRegistryIsEmancipated, error) {
	out := new(PermissionedRegistryTransferUnsafeUntilRegistryIsEmancipated)
	if err := permissionedRegistry.abi.UnpackIntoInterface(out, "TransferUnsafeUntilRegistryIsEmancipated", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionedRegistryTransferUnsafeWithMultipleAssignees represents a TransferUnsafeWithMultipleAssignees error raised by the PermissionedRegistry contract.
type PermissionedRegistryTransferUnsafeWithMultipleAssignees struct {
	TokenId *big.Int
	From    common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TransferUnsafeWithMultipleAssignees(uint256 tokenId, address from)
func PermissionedRegistryTransferUnsafeWithMultipleAssigneesErrorID() common.Hash {
	return common.HexToHash("0x677f1c18057fe4713efbb723b401a9d6b67633796a8433232c79a8e6fa834e27")
}

// UnpackTransferUnsafeWithMultipleAssigneesError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TransferUnsafeWithMultipleAssignees(uint256 tokenId, address from)
func (permissionedRegistry *PermissionedRegistry) UnpackTransferUnsafeWithMultipleAssigneesError(raw []byte) (*PermissionedRegistryTransferUnsafeWithMultipleAssignees, error) {
	out := new(PermissionedRegistryTransferUnsafeWithMultipleAssignees)
	if err := permissionedRegistry.abi.UnpackIntoInterface(out, "TransferUnsafeWithMultipleAssignees", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionedRegistryUUPSUnauthorizedCallContext represents a UUPSUnauthorizedCallContext error raised by the PermissionedRegistry contract.
type PermissionedRegistryUUPSUnauthorizedCallContext struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UUPSUnauthorizedCallContext()
func PermissionedRegistryUUPSUnauthorizedCallContextErrorID() common.Hash {
	return common.HexToHash("0xe07c8dba242a06571ac65fe4bbe20522c9fb111cb33599b799ff8039c1ed18f4")
}

// UnpackUUPSUnauthorizedCallContextError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UUPSUnauthorizedCallContext()
func (permissionedRegistry *PermissionedRegistry) UnpackUUPSUnauthorizedCallContextError(raw []byte) (*PermissionedRegistryUUPSUnauthorizedCallContext, error) {
	out := new(PermissionedRegistryUUPSUnauthorizedCallContext)
	if err := permissionedRegistry.abi.UnpackIntoInterface(out, "UUPSUnauthorizedCallContext", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionedRegistryUUPSUnsupportedProxiableUUID represents a UUPSUnsupportedProxiableUUID error raised by the PermissionedRegistry contract.
type PermissionedRegistryUUPSUnsupportedProxiableUUID struct {
	Slot [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UUPSUnsupportedProxiableUUID(bytes32 slot)
func PermissionedRegistryUUPSUnsupportedProxiableUUIDErrorID() common.Hash {
	return common.HexToHash("0xaa1d49a4c084bfa9aeeee2a0be65267a7f19ba7e1476b114dac513d2c14cb563")
}

// UnpackUUPSUnsupportedProxiableUUIDError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UUPSUnsupportedProxiableUUID(bytes32 slot)
func (permissionedRegistry *PermissionedRegistry) UnpackUUPSUnsupportedProxiableUUIDError(raw []byte) (*PermissionedRegistryUUPSUnsupportedProxiableUUID, error) {
	out := new(PermissionedRegistryUUPSUnsupportedProxiableUUID)
	if err := permissionedRegistry.abi.UnpackIntoInterface(out, "UUPSUnsupportedProxiableUUID", raw); err != nil {
		return nil, err
	}
	return out, nil
}
