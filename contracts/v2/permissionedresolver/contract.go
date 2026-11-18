// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package permissionedresolver

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

// PermissionedResolverMetaData contains all meta data concerning the PermissionedResolver contract.
var PermissionedResolverMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"namer\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"dns\",\"type\":\"bytes\"}],\"name\":\"DNSDecodingFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"resource\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"roleBitmap\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"EACCannotGrantRoles\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"resource\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"roleBitmap\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"EACCannotRevokeRoles\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EACInvalidAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roleBitmap\",\"type\":\"uint256\"}],\"name\":\"EACInvalidRoleBitmap\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"resource\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"role\",\"type\":\"uint256\"}],\"name\":\"EACMaxAssignees\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"resource\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"role\",\"type\":\"uint256\"}],\"name\":\"EACMinAssignees\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EACRootResourceNotAllowed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"resource\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"roleBitmap\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"EACUnauthorizedAccountRoles\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"contentType\",\"type\":\"uint256\"}],\"name\":\"InvalidContentType\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"addressBytes\",\"type\":\"bytes\"}],\"name\":\"InvalidEVMAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRecord\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"UnsupportedResolverProfile\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"recordId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"contentType\",\"type\":\"uint256\"}],\"name\":\"ABIUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"recordId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"coinType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"addressBytes\",\"type\":\"bytes\"}],\"name\":\"AddressUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"recordId\",\"type\":\"uint256\"}],\"name\":\"Cleared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"recordId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"hash\",\"type\":\"bytes\"}],\"name\":\"ContenthashUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"recordId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"keyHash\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"DataUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"resource\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldRoleBitmap\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newRoleBitmap\",\"type\":\"uint256\"}],\"name\":\"EACRolesChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"recordId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"implementer\",\"type\":\"address\"}],\"name\":\"InterfaceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"recordId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"}],\"name\":\"Linked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"recordId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"primaryName\",\"type\":\"string\"}],\"name\":\"NameUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ResolverCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"resource\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"arg\",\"type\":\"bytes\"}],\"name\":\"ResourceArgument\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"recordId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"keyHash\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"TextUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ROOT_RESOURCE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"canUpgradeFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"setter\",\"type\":\"bytes\"}],\"name\":\"decodeSetter\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"arg\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"resource\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"roleBitmap\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"resource\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"roleBitmap\",\"type\":\"uint256\"}],\"name\":\"getAssigneeCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"counts\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mask\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRecordCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"getRecordId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"resource\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"roleBitmap\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRoles\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roleBitmap\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRootRoles\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"setter\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantSetterRoles\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"resource\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"roleBitmap\",\"type\":\"uint256\"}],\"name\":\"hasAssignees\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"resource\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"roleBitmap\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRoles\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roleBitmap\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRootRoles\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"roleBitmap\",\"type\":\"uint256\"}],\"internalType\":\"structGrant[]\",\"name\":\"grants\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calls\",\"type\":\"bytes[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"namer\",\"type\":\"address\"}],\"name\":\"isContractNamer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"resource\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"roleBitmap\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isOnlyAssignee\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"sourceName\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"targetNode\",\"type\":\"bytes32\"}],\"name\":\"linkToNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"sourceName\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"recordId\",\"type\":\"uint256\"}],\"name\":\"linkToRecord\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"calls\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes[]\",\"name\":\"calls\",\"type\":\"bytes[]\"}],\"name\":\"multicallWithNodeCheck\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"resolve\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"resource\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"roleBitmap\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRoles\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roleBitmap\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRootRoles\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"resource\",\"type\":\"uint256\"}],\"name\":\"roleCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"resource\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"roles\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"contentType\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"setABI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"coinType\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"addressBytes\",\"type\":\"bytes\"}],\"name\":\"setAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"hash\",\"type\":\"bytes\"}],\"name\":\"setContenthash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"setData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"implementer\",\"type\":\"address\"}],\"name\":\"setInterface\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"primaryName\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"setText\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"feature\",\"type\":\"bytes4\"}],\"name\":\"supportsFeature\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	ID:  "PermissionedResolver",
}

// PermissionedResolver is an auto generated Go binding around an Ethereum contract.
type PermissionedResolver struct {
	abi abi.ABI
}

// GetABI returns the ABI associated with this contract binding.
func (c *PermissionedResolver) GetABI() abi.ABI {
	return c.abi
}

// NewPermissionedResolver creates a new instance of PermissionedResolver.
func NewPermissionedResolver() *PermissionedResolver {
	parsed, err := PermissionedResolverMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &PermissionedResolver{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *PermissionedResolver) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address namer) returns()
func (permissionedResolver *PermissionedResolver) PackConstructor(namer common.Address) []byte {
	enc, err := permissionedResolver.abi.Pack("", namer)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackROOTRESOURCE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1c3fc3eb.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function ROOT_RESOURCE() view returns(uint256)
func (permissionedResolver *PermissionedResolver) PackROOTRESOURCE() []byte {
	enc, err := permissionedResolver.abi.Pack("ROOT_RESOURCE")
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
func (permissionedResolver *PermissionedResolver) TryPackROOTRESOURCE() ([]byte, error) {
	return permissionedResolver.abi.Pack("ROOT_RESOURCE")
}

// UnpackROOTRESOURCE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x1c3fc3eb.
//
// Solidity: function ROOT_RESOURCE() view returns(uint256)
func (permissionedResolver *PermissionedResolver) UnpackROOTRESOURCE(data []byte) (*big.Int, error) {
	out, err := permissionedResolver.abi.Unpack("ROOT_RESOURCE", data)
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
func (permissionedResolver *PermissionedResolver) PackUPGRADEINTERFACEVERSION() []byte {
	enc, err := permissionedResolver.abi.Pack("UPGRADE_INTERFACE_VERSION")
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
func (permissionedResolver *PermissionedResolver) TryPackUPGRADEINTERFACEVERSION() ([]byte, error) {
	return permissionedResolver.abi.Pack("UPGRADE_INTERFACE_VERSION")
}

// UnpackUPGRADEINTERFACEVERSION is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (permissionedResolver *PermissionedResolver) UnpackUPGRADEINTERFACEVERSION(data []byte) (string, error) {
	out, err := permissionedResolver.abi.Unpack("UPGRADE_INTERFACE_VERSION", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, nil
}

// PackCanUpgradeFrom is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf41a143d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function canUpgradeFrom(address ) pure returns(bool allowed)
func (permissionedResolver *PermissionedResolver) PackCanUpgradeFrom(arg0 common.Address) []byte {
	enc, err := permissionedResolver.abi.Pack("canUpgradeFrom", arg0)
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
func (permissionedResolver *PermissionedResolver) TryPackCanUpgradeFrom(arg0 common.Address) ([]byte, error) {
	return permissionedResolver.abi.Pack("canUpgradeFrom", arg0)
}

// UnpackCanUpgradeFrom is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf41a143d.
//
// Solidity: function canUpgradeFrom(address ) pure returns(bool allowed)
func (permissionedResolver *PermissionedResolver) UnpackCanUpgradeFrom(data []byte) (bool, error) {
	out, err := permissionedResolver.abi.Unpack("canUpgradeFrom", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackDecodeSetter is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x18308aef.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function decodeSetter(bytes setter) pure returns(bytes arg, uint256 resource, uint256 roleBitmap)
func (permissionedResolver *PermissionedResolver) PackDecodeSetter(setter []byte) []byte {
	enc, err := permissionedResolver.abi.Pack("decodeSetter", setter)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDecodeSetter is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x18308aef.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function decodeSetter(bytes setter) pure returns(bytes arg, uint256 resource, uint256 roleBitmap)
func (permissionedResolver *PermissionedResolver) TryPackDecodeSetter(setter []byte) ([]byte, error) {
	return permissionedResolver.abi.Pack("decodeSetter", setter)
}

// DecodeSetterOutput serves as a container for the return parameters of contract
// method DecodeSetter.
type DecodeSetterOutput struct {
	Arg        []byte
	Resource   *big.Int
	RoleBitmap *big.Int
}

// UnpackDecodeSetter is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x18308aef.
//
// Solidity: function decodeSetter(bytes setter) pure returns(bytes arg, uint256 resource, uint256 roleBitmap)
func (permissionedResolver *PermissionedResolver) UnpackDecodeSetter(data []byte) (DecodeSetterOutput, error) {
	out, err := permissionedResolver.abi.Unpack("decodeSetter", data)
	outstruct := new(DecodeSetterOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Arg = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.Resource = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	outstruct.RoleBitmap = abi.ConvertType(out[2], new(big.Int)).(*big.Int)
	return *outstruct, nil
}

// PackGetAssigneeCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3634f911.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getAssigneeCount(uint256 resource, uint256 roleBitmap) view returns(uint256 counts, uint256 mask)
func (permissionedResolver *PermissionedResolver) PackGetAssigneeCount(resource *big.Int, roleBitmap *big.Int) []byte {
	enc, err := permissionedResolver.abi.Pack("getAssigneeCount", resource, roleBitmap)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetAssigneeCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3634f911.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getAssigneeCount(uint256 resource, uint256 roleBitmap) view returns(uint256 counts, uint256 mask)
func (permissionedResolver *PermissionedResolver) TryPackGetAssigneeCount(resource *big.Int, roleBitmap *big.Int) ([]byte, error) {
	return permissionedResolver.abi.Pack("getAssigneeCount", resource, roleBitmap)
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
// Solidity: function getAssigneeCount(uint256 resource, uint256 roleBitmap) view returns(uint256 counts, uint256 mask)
func (permissionedResolver *PermissionedResolver) UnpackGetAssigneeCount(data []byte) (GetAssigneeCountOutput, error) {
	out, err := permissionedResolver.abi.Unpack("getAssigneeCount", data)
	outstruct := new(GetAssigneeCountOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Counts = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.Mask = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	return *outstruct, nil
}

// PackGetRecordCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xca267f28.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getRecordCount() view returns(uint256)
func (permissionedResolver *PermissionedResolver) PackGetRecordCount() []byte {
	enc, err := permissionedResolver.abi.Pack("getRecordCount")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetRecordCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xca267f28.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getRecordCount() view returns(uint256)
func (permissionedResolver *PermissionedResolver) TryPackGetRecordCount() ([]byte, error) {
	return permissionedResolver.abi.Pack("getRecordCount")
}

// UnpackGetRecordCount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xca267f28.
//
// Solidity: function getRecordCount() view returns(uint256)
func (permissionedResolver *PermissionedResolver) UnpackGetRecordCount(data []byte) (*big.Int, error) {
	out, err := permissionedResolver.abi.Unpack("getRecordCount", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetRecordId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfaf10086.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getRecordId(bytes32 node) view returns(uint256)
func (permissionedResolver *PermissionedResolver) PackGetRecordId(node [32]byte) []byte {
	enc, err := permissionedResolver.abi.Pack("getRecordId", node)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetRecordId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfaf10086.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getRecordId(bytes32 node) view returns(uint256)
func (permissionedResolver *PermissionedResolver) TryPackGetRecordId(node [32]byte) ([]byte, error) {
	return permissionedResolver.abi.Pack("getRecordId", node)
}

// UnpackGetRecordId is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xfaf10086.
//
// Solidity: function getRecordId(bytes32 node) view returns(uint256)
func (permissionedResolver *PermissionedResolver) UnpackGetRecordId(data []byte) (*big.Int, error) {
	out, err := permissionedResolver.abi.Unpack("getRecordId", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGrantRoles is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7c300586.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function grantRoles(uint256 resource, uint256 roleBitmap, address account) pure returns(bool)
func (permissionedResolver *PermissionedResolver) PackGrantRoles(resource *big.Int, roleBitmap *big.Int, account common.Address) []byte {
	enc, err := permissionedResolver.abi.Pack("grantRoles", resource, roleBitmap, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGrantRoles is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7c300586.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function grantRoles(uint256 resource, uint256 roleBitmap, address account) pure returns(bool)
func (permissionedResolver *PermissionedResolver) TryPackGrantRoles(resource *big.Int, roleBitmap *big.Int, account common.Address) ([]byte, error) {
	return permissionedResolver.abi.Pack("grantRoles", resource, roleBitmap, account)
}

// UnpackGrantRoles is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7c300586.
//
// Solidity: function grantRoles(uint256 resource, uint256 roleBitmap, address account) pure returns(bool)
func (permissionedResolver *PermissionedResolver) UnpackGrantRoles(data []byte) (bool, error) {
	out, err := permissionedResolver.abi.Unpack("grantRoles", data)
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
func (permissionedResolver *PermissionedResolver) PackGrantRootRoles(roleBitmap *big.Int, account common.Address) []byte {
	enc, err := permissionedResolver.abi.Pack("grantRootRoles", roleBitmap, account)
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
func (permissionedResolver *PermissionedResolver) TryPackGrantRootRoles(roleBitmap *big.Int, account common.Address) ([]byte, error) {
	return permissionedResolver.abi.Pack("grantRootRoles", roleBitmap, account)
}

// UnpackGrantRootRoles is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x072d5d77.
//
// Solidity: function grantRootRoles(uint256 roleBitmap, address account) returns(bool)
func (permissionedResolver *PermissionedResolver) UnpackGrantRootRoles(data []byte) (bool, error) {
	out, err := permissionedResolver.abi.Unpack("grantRootRoles", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackGrantSetterRoles is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xccd3eaff.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function grantSetterRoles(bytes setter, address account) returns(bool)
func (permissionedResolver *PermissionedResolver) PackGrantSetterRoles(setter []byte, account common.Address) []byte {
	enc, err := permissionedResolver.abi.Pack("grantSetterRoles", setter, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGrantSetterRoles is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xccd3eaff.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function grantSetterRoles(bytes setter, address account) returns(bool)
func (permissionedResolver *PermissionedResolver) TryPackGrantSetterRoles(setter []byte, account common.Address) ([]byte, error) {
	return permissionedResolver.abi.Pack("grantSetterRoles", setter, account)
}

// UnpackGrantSetterRoles is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xccd3eaff.
//
// Solidity: function grantSetterRoles(bytes setter, address account) returns(bool)
func (permissionedResolver *PermissionedResolver) UnpackGrantSetterRoles(data []byte) (bool, error) {
	out, err := permissionedResolver.abi.Unpack("grantSetterRoles", data)
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
func (permissionedResolver *PermissionedResolver) PackHasAssignees(resource *big.Int, roleBitmap *big.Int) []byte {
	enc, err := permissionedResolver.abi.Pack("hasAssignees", resource, roleBitmap)
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
func (permissionedResolver *PermissionedResolver) TryPackHasAssignees(resource *big.Int, roleBitmap *big.Int) ([]byte, error) {
	return permissionedResolver.abi.Pack("hasAssignees", resource, roleBitmap)
}

// UnpackHasAssignees is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x11b8e00a.
//
// Solidity: function hasAssignees(uint256 resource, uint256 roleBitmap) view returns(bool)
func (permissionedResolver *PermissionedResolver) UnpackHasAssignees(data []byte) (bool, error) {
	out, err := permissionedResolver.abi.Unpack("hasAssignees", data)
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
// Solidity: function hasRoles(uint256 resource, uint256 roleBitmap, address account) view returns(bool)
func (permissionedResolver *PermissionedResolver) PackHasRoles(resource *big.Int, roleBitmap *big.Int, account common.Address) []byte {
	enc, err := permissionedResolver.abi.Pack("hasRoles", resource, roleBitmap, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackHasRoles is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd3bf89b1.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function hasRoles(uint256 resource, uint256 roleBitmap, address account) view returns(bool)
func (permissionedResolver *PermissionedResolver) TryPackHasRoles(resource *big.Int, roleBitmap *big.Int, account common.Address) ([]byte, error) {
	return permissionedResolver.abi.Pack("hasRoles", resource, roleBitmap, account)
}

// UnpackHasRoles is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd3bf89b1.
//
// Solidity: function hasRoles(uint256 resource, uint256 roleBitmap, address account) view returns(bool)
func (permissionedResolver *PermissionedResolver) UnpackHasRoles(data []byte) (bool, error) {
	out, err := permissionedResolver.abi.Unpack("hasRoles", data)
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
func (permissionedResolver *PermissionedResolver) PackHasRootRoles(roleBitmap *big.Int, account common.Address) []byte {
	enc, err := permissionedResolver.abi.Pack("hasRootRoles", roleBitmap, account)
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
func (permissionedResolver *PermissionedResolver) TryPackHasRootRoles(roleBitmap *big.Int, account common.Address) ([]byte, error) {
	return permissionedResolver.abi.Pack("hasRootRoles", roleBitmap, account)
}

// UnpackHasRootRoles is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x781ef8db.
//
// Solidity: function hasRootRoles(uint256 roleBitmap, address account) view returns(bool)
func (permissionedResolver *PermissionedResolver) UnpackHasRootRoles(data []byte) (bool, error) {
	out, err := permissionedResolver.abi.Unpack("hasRootRoles", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x33cc44a0.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function initialize((address,uint256)[] grants, bytes[] calls) returns()
func (permissionedResolver *PermissionedResolver) PackInitialize(grants []Grant, calls [][]byte) []byte {
	enc, err := permissionedResolver.abi.Pack("initialize", grants, calls)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x33cc44a0.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function initialize((address,uint256)[] grants, bytes[] calls) returns()
func (permissionedResolver *PermissionedResolver) TryPackInitialize(grants []Grant, calls [][]byte) ([]byte, error) {
	return permissionedResolver.abi.Pack("initialize", grants, calls)
}

// PackIsContractNamer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6f3ff726.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isContractNamer(address namer) view returns(bool)
func (permissionedResolver *PermissionedResolver) PackIsContractNamer(namer common.Address) []byte {
	enc, err := permissionedResolver.abi.Pack("isContractNamer", namer)
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
func (permissionedResolver *PermissionedResolver) TryPackIsContractNamer(namer common.Address) ([]byte, error) {
	return permissionedResolver.abi.Pack("isContractNamer", namer)
}

// UnpackIsContractNamer is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6f3ff726.
//
// Solidity: function isContractNamer(address namer) view returns(bool)
func (permissionedResolver *PermissionedResolver) UnpackIsContractNamer(data []byte) (bool, error) {
	out, err := permissionedResolver.abi.Unpack("isContractNamer", data)
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
func (permissionedResolver *PermissionedResolver) PackIsOnlyAssignee(resource *big.Int, roleBitmap *big.Int, account common.Address) []byte {
	enc, err := permissionedResolver.abi.Pack("isOnlyAssignee", resource, roleBitmap, account)
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
func (permissionedResolver *PermissionedResolver) TryPackIsOnlyAssignee(resource *big.Int, roleBitmap *big.Int, account common.Address) ([]byte, error) {
	return permissionedResolver.abi.Pack("isOnlyAssignee", resource, roleBitmap, account)
}

// UnpackIsOnlyAssignee is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8e77c95f.
//
// Solidity: function isOnlyAssignee(uint256 resource, uint256 roleBitmap, address account) view returns(bool)
func (permissionedResolver *PermissionedResolver) UnpackIsOnlyAssignee(data []byte) (bool, error) {
	out, err := permissionedResolver.abi.Unpack("isOnlyAssignee", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackLinkToNode is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5d27b8e5.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function linkToNode(bytes sourceName, bytes32 targetNode) returns()
func (permissionedResolver *PermissionedResolver) PackLinkToNode(sourceName []byte, targetNode [32]byte) []byte {
	enc, err := permissionedResolver.abi.Pack("linkToNode", sourceName, targetNode)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackLinkToNode is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5d27b8e5.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function linkToNode(bytes sourceName, bytes32 targetNode) returns()
func (permissionedResolver *PermissionedResolver) TryPackLinkToNode(sourceName []byte, targetNode [32]byte) ([]byte, error) {
	return permissionedResolver.abi.Pack("linkToNode", sourceName, targetNode)
}

// PackLinkToRecord is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x35378097.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function linkToRecord(bytes sourceName, uint256 recordId) returns()
func (permissionedResolver *PermissionedResolver) PackLinkToRecord(sourceName []byte, recordId *big.Int) []byte {
	enc, err := permissionedResolver.abi.Pack("linkToRecord", sourceName, recordId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackLinkToRecord is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x35378097.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function linkToRecord(bytes sourceName, uint256 recordId) returns()
func (permissionedResolver *PermissionedResolver) TryPackLinkToRecord(sourceName []byte, recordId *big.Int) ([]byte, error) {
	return permissionedResolver.abi.Pack("linkToRecord", sourceName, recordId)
}

// PackMulticall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xac9650d8.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function multicall(bytes[] calls) returns(bytes[] results)
func (permissionedResolver *PermissionedResolver) PackMulticall(calls [][]byte) []byte {
	enc, err := permissionedResolver.abi.Pack("multicall", calls)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackMulticall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xac9650d8.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function multicall(bytes[] calls) returns(bytes[] results)
func (permissionedResolver *PermissionedResolver) TryPackMulticall(calls [][]byte) ([]byte, error) {
	return permissionedResolver.abi.Pack("multicall", calls)
}

// UnpackMulticall is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xac9650d8.
//
// Solidity: function multicall(bytes[] calls) returns(bytes[] results)
func (permissionedResolver *PermissionedResolver) UnpackMulticall(data []byte) ([][]byte, error) {
	out, err := permissionedResolver.abi.Unpack("multicall", data)
	if err != nil {
		return *new([][]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)
	return out0, nil
}

// PackMulticallWithNodeCheck is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe32954eb.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function multicallWithNodeCheck(bytes32 , bytes[] calls) returns(bytes[])
func (permissionedResolver *PermissionedResolver) PackMulticallWithNodeCheck(arg0 [32]byte, calls [][]byte) []byte {
	enc, err := permissionedResolver.abi.Pack("multicallWithNodeCheck", arg0, calls)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackMulticallWithNodeCheck is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe32954eb.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function multicallWithNodeCheck(bytes32 , bytes[] calls) returns(bytes[])
func (permissionedResolver *PermissionedResolver) TryPackMulticallWithNodeCheck(arg0 [32]byte, calls [][]byte) ([]byte, error) {
	return permissionedResolver.abi.Pack("multicallWithNodeCheck", arg0, calls)
}

// UnpackMulticallWithNodeCheck is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe32954eb.
//
// Solidity: function multicallWithNodeCheck(bytes32 , bytes[] calls) returns(bytes[])
func (permissionedResolver *PermissionedResolver) UnpackMulticallWithNodeCheck(data []byte) ([][]byte, error) {
	out, err := permissionedResolver.abi.Unpack("multicallWithNodeCheck", data)
	if err != nil {
		return *new([][]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)
	return out0, nil
}

// PackProxiableUUID is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x52d1902d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (permissionedResolver *PermissionedResolver) PackProxiableUUID() []byte {
	enc, err := permissionedResolver.abi.Pack("proxiableUUID")
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
func (permissionedResolver *PermissionedResolver) TryPackProxiableUUID() ([]byte, error) {
	return permissionedResolver.abi.Pack("proxiableUUID")
}

// UnpackProxiableUUID is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (permissionedResolver *PermissionedResolver) UnpackProxiableUUID(data []byte) ([32]byte, error) {
	out, err := permissionedResolver.abi.Unpack("proxiableUUID", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackResolve is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9061b923.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function resolve(bytes name, bytes data) view returns(bytes)
func (permissionedResolver *PermissionedResolver) PackResolve(name []byte, data []byte) []byte {
	enc, err := permissionedResolver.abi.Pack("resolve", name, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackResolve is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9061b923.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function resolve(bytes name, bytes data) view returns(bytes)
func (permissionedResolver *PermissionedResolver) TryPackResolve(name []byte, data []byte) ([]byte, error) {
	return permissionedResolver.abi.Pack("resolve", name, data)
}

// UnpackResolve is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9061b923.
//
// Solidity: function resolve(bytes name, bytes data) view returns(bytes)
func (permissionedResolver *PermissionedResolver) UnpackResolve(data []byte) ([]byte, error) {
	out, err := permissionedResolver.abi.Unpack("resolve", data)
	if err != nil {
		return *new([]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	return out0, nil
}

// PackRevokeRoles is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdfa70d8b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function revokeRoles(uint256 resource, uint256 roleBitmap, address account) returns(bool)
func (permissionedResolver *PermissionedResolver) PackRevokeRoles(resource *big.Int, roleBitmap *big.Int, account common.Address) []byte {
	enc, err := permissionedResolver.abi.Pack("revokeRoles", resource, roleBitmap, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRevokeRoles is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdfa70d8b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function revokeRoles(uint256 resource, uint256 roleBitmap, address account) returns(bool)
func (permissionedResolver *PermissionedResolver) TryPackRevokeRoles(resource *big.Int, roleBitmap *big.Int, account common.Address) ([]byte, error) {
	return permissionedResolver.abi.Pack("revokeRoles", resource, roleBitmap, account)
}

// UnpackRevokeRoles is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xdfa70d8b.
//
// Solidity: function revokeRoles(uint256 resource, uint256 roleBitmap, address account) returns(bool)
func (permissionedResolver *PermissionedResolver) UnpackRevokeRoles(data []byte) (bool, error) {
	out, err := permissionedResolver.abi.Unpack("revokeRoles", data)
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
func (permissionedResolver *PermissionedResolver) PackRevokeRootRoles(roleBitmap *big.Int, account common.Address) []byte {
	enc, err := permissionedResolver.abi.Pack("revokeRootRoles", roleBitmap, account)
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
func (permissionedResolver *PermissionedResolver) TryPackRevokeRootRoles(roleBitmap *big.Int, account common.Address) ([]byte, error) {
	return permissionedResolver.abi.Pack("revokeRootRoles", roleBitmap, account)
}

// UnpackRevokeRootRoles is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xce156e82.
//
// Solidity: function revokeRootRoles(uint256 roleBitmap, address account) returns(bool)
func (permissionedResolver *PermissionedResolver) UnpackRevokeRootRoles(data []byte) (bool, error) {
	out, err := permissionedResolver.abi.Unpack("revokeRootRoles", data)
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
// Solidity: function roleCount(uint256 resource) view returns(uint256)
func (permissionedResolver *PermissionedResolver) PackRoleCount(resource *big.Int) []byte {
	enc, err := permissionedResolver.abi.Pack("roleCount", resource)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRoleCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2f27fa24.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function roleCount(uint256 resource) view returns(uint256)
func (permissionedResolver *PermissionedResolver) TryPackRoleCount(resource *big.Int) ([]byte, error) {
	return permissionedResolver.abi.Pack("roleCount", resource)
}

// UnpackRoleCount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x2f27fa24.
//
// Solidity: function roleCount(uint256 resource) view returns(uint256)
func (permissionedResolver *PermissionedResolver) UnpackRoleCount(data []byte) (*big.Int, error) {
	out, err := permissionedResolver.abi.Unpack("roleCount", data)
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
// Solidity: function roles(uint256 resource, address account) view returns(uint256)
func (permissionedResolver *PermissionedResolver) PackRoles(resource *big.Int, account common.Address) []byte {
	enc, err := permissionedResolver.abi.Pack("roles", resource, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRoles is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5adf4724.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function roles(uint256 resource, address account) view returns(uint256)
func (permissionedResolver *PermissionedResolver) TryPackRoles(resource *big.Int, account common.Address) ([]byte, error) {
	return permissionedResolver.abi.Pack("roles", resource, account)
}

// UnpackRoles is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5adf4724.
//
// Solidity: function roles(uint256 resource, address account) view returns(uint256)
func (permissionedResolver *PermissionedResolver) UnpackRoles(data []byte) (*big.Int, error) {
	out, err := permissionedResolver.abi.Unpack("roles", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackSetABI is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd26f550e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setABI(bytes name, uint256 contentType, bytes data) returns()
func (permissionedResolver *PermissionedResolver) PackSetABI(name []byte, contentType *big.Int, data []byte) []byte {
	enc, err := permissionedResolver.abi.Pack("setABI", name, contentType, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetABI is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd26f550e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setABI(bytes name, uint256 contentType, bytes data) returns()
func (permissionedResolver *PermissionedResolver) TryPackSetABI(name []byte, contentType *big.Int, data []byte) ([]byte, error) {
	return permissionedResolver.abi.Pack("setABI", name, contentType, data)
}

// PackSetAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb4436dde.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setAddress(bytes name, uint256 coinType, bytes addressBytes) returns()
func (permissionedResolver *PermissionedResolver) PackSetAddress(name []byte, coinType *big.Int, addressBytes []byte) []byte {
	enc, err := permissionedResolver.abi.Pack("setAddress", name, coinType, addressBytes)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb4436dde.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setAddress(bytes name, uint256 coinType, bytes addressBytes) returns()
func (permissionedResolver *PermissionedResolver) TryPackSetAddress(name []byte, coinType *big.Int, addressBytes []byte) ([]byte, error) {
	return permissionedResolver.abi.Pack("setAddress", name, coinType, addressBytes)
}

// PackSetContenthash is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc5d7badd.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setContenthash(bytes name, bytes hash) returns()
func (permissionedResolver *PermissionedResolver) PackSetContenthash(name []byte, hash []byte) []byte {
	enc, err := permissionedResolver.abi.Pack("setContenthash", name, hash)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetContenthash is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc5d7badd.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setContenthash(bytes name, bytes hash) returns()
func (permissionedResolver *PermissionedResolver) TryPackSetContenthash(name []byte, hash []byte) ([]byte, error) {
	return permissionedResolver.abi.Pack("setContenthash", name, hash)
}

// PackSetData is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xeb4b73bb.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setData(bytes name, string key, bytes value) returns()
func (permissionedResolver *PermissionedResolver) PackSetData(name []byte, key string, value []byte) []byte {
	enc, err := permissionedResolver.abi.Pack("setData", name, key, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetData is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xeb4b73bb.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setData(bytes name, string key, bytes value) returns()
func (permissionedResolver *PermissionedResolver) TryPackSetData(name []byte, key string, value []byte) ([]byte, error) {
	return permissionedResolver.abi.Pack("setData", name, key, value)
}

// PackSetInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9ce8c375.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setInterface(bytes name, bytes4 interfaceId, address implementer) returns()
func (permissionedResolver *PermissionedResolver) PackSetInterface(name []byte, interfaceId [4]byte, implementer common.Address) []byte {
	enc, err := permissionedResolver.abi.Pack("setInterface", name, interfaceId, implementer)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9ce8c375.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setInterface(bytes name, bytes4 interfaceId, address implementer) returns()
func (permissionedResolver *PermissionedResolver) TryPackSetInterface(name []byte, interfaceId [4]byte, implementer common.Address) ([]byte, error) {
	return permissionedResolver.abi.Pack("setInterface", name, interfaceId, implementer)
}

// PackSetName is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0ce0112a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setName(bytes name, string primaryName) returns()
func (permissionedResolver *PermissionedResolver) PackSetName(name []byte, primaryName string) []byte {
	enc, err := permissionedResolver.abi.Pack("setName", name, primaryName)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetName is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0ce0112a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setName(bytes name, string primaryName) returns()
func (permissionedResolver *PermissionedResolver) TryPackSetName(name []byte, primaryName string) ([]byte, error) {
	return permissionedResolver.abi.Pack("setName", name, primaryName)
}

// PackSetText is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc7279f88.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setText(bytes name, string key, string value) returns()
func (permissionedResolver *PermissionedResolver) PackSetText(name []byte, key string, value string) []byte {
	enc, err := permissionedResolver.abi.Pack("setText", name, key, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetText is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc7279f88.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setText(bytes name, string key, string value) returns()
func (permissionedResolver *PermissionedResolver) TryPackSetText(name []byte, key string, value string) ([]byte, error) {
	return permissionedResolver.abi.Pack("setText", name, key, value)
}

// PackSupportsFeature is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x582de3e7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function supportsFeature(bytes4 feature) pure returns(bool)
func (permissionedResolver *PermissionedResolver) PackSupportsFeature(feature [4]byte) []byte {
	enc, err := permissionedResolver.abi.Pack("supportsFeature", feature)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSupportsFeature is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x582de3e7.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function supportsFeature(bytes4 feature) pure returns(bool)
func (permissionedResolver *PermissionedResolver) TryPackSupportsFeature(feature [4]byte) ([]byte, error) {
	return permissionedResolver.abi.Pack("supportsFeature", feature)
}

// UnpackSupportsFeature is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x582de3e7.
//
// Solidity: function supportsFeature(bytes4 feature) pure returns(bool)
func (permissionedResolver *PermissionedResolver) UnpackSupportsFeature(data []byte) (bool, error) {
	out, err := permissionedResolver.abi.Unpack("supportsFeature", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackSupportsInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01ffc9a7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (permissionedResolver *PermissionedResolver) PackSupportsInterface(interfaceId [4]byte) []byte {
	enc, err := permissionedResolver.abi.Pack("supportsInterface", interfaceId)
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
func (permissionedResolver *PermissionedResolver) TryPackSupportsInterface(interfaceId [4]byte) ([]byte, error) {
	return permissionedResolver.abi.Pack("supportsInterface", interfaceId)
}

// UnpackSupportsInterface is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (permissionedResolver *PermissionedResolver) UnpackSupportsInterface(data []byte) (bool, error) {
	out, err := permissionedResolver.abi.Unpack("supportsInterface", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackUpgradeToAndCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4f1ef286.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (permissionedResolver *PermissionedResolver) PackUpgradeToAndCall(newImplementation common.Address, data []byte) []byte {
	enc, err := permissionedResolver.abi.Pack("upgradeToAndCall", newImplementation, data)
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
func (permissionedResolver *PermissionedResolver) TryPackUpgradeToAndCall(newImplementation common.Address, data []byte) ([]byte, error) {
	return permissionedResolver.abi.Pack("upgradeToAndCall", newImplementation, data)
}

// PermissionedResolverABIUpdated represents a ABIUpdated event raised by the PermissionedResolver contract.
type PermissionedResolverABIUpdated struct {
	RecordId    *big.Int
	ContentType *big.Int
	Raw         *types.Log // Blockchain specific contextual infos
}

const PermissionedResolverABIUpdatedEventName = "ABIUpdated"

// ContractEventName returns the user-defined event name.
func (PermissionedResolverABIUpdated) ContractEventName() string {
	return PermissionedResolverABIUpdatedEventName
}

// UnpackABIUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ABIUpdated(uint256 indexed recordId, uint256 indexed contentType)
func (permissionedResolver *PermissionedResolver) UnpackABIUpdatedEvent(log *types.Log) (*PermissionedResolverABIUpdated, error) {
	event := "ABIUpdated"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != permissionedResolver.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(PermissionedResolverABIUpdated)
	if len(log.Data) > 0 {
		if err := permissionedResolver.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range permissionedResolver.abi.Events[event].Inputs {
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

// PermissionedResolverAddressUpdated represents a AddressUpdated event raised by the PermissionedResolver contract.
type PermissionedResolverAddressUpdated struct {
	RecordId     *big.Int
	CoinType     *big.Int
	AddressBytes []byte
	Raw          *types.Log // Blockchain specific contextual infos
}

const PermissionedResolverAddressUpdatedEventName = "AddressUpdated"

// ContractEventName returns the user-defined event name.
func (PermissionedResolverAddressUpdated) ContractEventName() string {
	return PermissionedResolverAddressUpdatedEventName
}

// UnpackAddressUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event AddressUpdated(uint256 indexed recordId, uint256 coinType, bytes addressBytes)
func (permissionedResolver *PermissionedResolver) UnpackAddressUpdatedEvent(log *types.Log) (*PermissionedResolverAddressUpdated, error) {
	event := "AddressUpdated"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != permissionedResolver.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(PermissionedResolverAddressUpdated)
	if len(log.Data) > 0 {
		if err := permissionedResolver.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range permissionedResolver.abi.Events[event].Inputs {
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

// PermissionedResolverCleared represents a Cleared event raised by the PermissionedResolver contract.
type PermissionedResolverCleared struct {
	RecordId *big.Int
	Raw      *types.Log // Blockchain specific contextual infos
}

const PermissionedResolverClearedEventName = "Cleared"

// ContractEventName returns the user-defined event name.
func (PermissionedResolverCleared) ContractEventName() string {
	return PermissionedResolverClearedEventName
}

// UnpackClearedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Cleared(uint256 indexed recordId)
func (permissionedResolver *PermissionedResolver) UnpackClearedEvent(log *types.Log) (*PermissionedResolverCleared, error) {
	event := "Cleared"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != permissionedResolver.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(PermissionedResolverCleared)
	if len(log.Data) > 0 {
		if err := permissionedResolver.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range permissionedResolver.abi.Events[event].Inputs {
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

// PermissionedResolverContenthashUpdated represents a ContenthashUpdated event raised by the PermissionedResolver contract.
type PermissionedResolverContenthashUpdated struct {
	RecordId *big.Int
	Hash     []byte
	Raw      *types.Log // Blockchain specific contextual infos
}

const PermissionedResolverContenthashUpdatedEventName = "ContenthashUpdated"

// ContractEventName returns the user-defined event name.
func (PermissionedResolverContenthashUpdated) ContractEventName() string {
	return PermissionedResolverContenthashUpdatedEventName
}

// UnpackContenthashUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ContenthashUpdated(uint256 indexed recordId, bytes hash)
func (permissionedResolver *PermissionedResolver) UnpackContenthashUpdatedEvent(log *types.Log) (*PermissionedResolverContenthashUpdated, error) {
	event := "ContenthashUpdated"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != permissionedResolver.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(PermissionedResolverContenthashUpdated)
	if len(log.Data) > 0 {
		if err := permissionedResolver.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range permissionedResolver.abi.Events[event].Inputs {
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

// PermissionedResolverDataUpdated represents a DataUpdated event raised by the PermissionedResolver contract.
type PermissionedResolverDataUpdated struct {
	RecordId *big.Int
	KeyHash  common.Hash
	Key      string
	Value    []byte
	Raw      *types.Log // Blockchain specific contextual infos
}

const PermissionedResolverDataUpdatedEventName = "DataUpdated"

// ContractEventName returns the user-defined event name.
func (PermissionedResolverDataUpdated) ContractEventName() string {
	return PermissionedResolverDataUpdatedEventName
}

// UnpackDataUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DataUpdated(uint256 indexed recordId, string indexed keyHash, string key, bytes value)
func (permissionedResolver *PermissionedResolver) UnpackDataUpdatedEvent(log *types.Log) (*PermissionedResolverDataUpdated, error) {
	event := "DataUpdated"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != permissionedResolver.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(PermissionedResolverDataUpdated)
	if len(log.Data) > 0 {
		if err := permissionedResolver.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range permissionedResolver.abi.Events[event].Inputs {
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

// PermissionedResolverEACRolesChanged represents a EACRolesChanged event raised by the PermissionedResolver contract.
type PermissionedResolverEACRolesChanged struct {
	Resource      *big.Int
	Account       common.Address
	OldRoleBitmap *big.Int
	NewRoleBitmap *big.Int
	Raw           *types.Log // Blockchain specific contextual infos
}

const PermissionedResolverEACRolesChangedEventName = "EACRolesChanged"

// ContractEventName returns the user-defined event name.
func (PermissionedResolverEACRolesChanged) ContractEventName() string {
	return PermissionedResolverEACRolesChangedEventName
}

// UnpackEACRolesChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event EACRolesChanged(uint256 indexed resource, address indexed account, uint256 oldRoleBitmap, uint256 newRoleBitmap)
func (permissionedResolver *PermissionedResolver) UnpackEACRolesChangedEvent(log *types.Log) (*PermissionedResolverEACRolesChanged, error) {
	event := "EACRolesChanged"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != permissionedResolver.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(PermissionedResolverEACRolesChanged)
	if len(log.Data) > 0 {
		if err := permissionedResolver.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range permissionedResolver.abi.Events[event].Inputs {
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

// PermissionedResolverInitialized represents a Initialized event raised by the PermissionedResolver contract.
type PermissionedResolverInitialized struct {
	Version uint64
	Raw     *types.Log // Blockchain specific contextual infos
}

const PermissionedResolverInitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (PermissionedResolverInitialized) ContractEventName() string {
	return PermissionedResolverInitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint64 version)
func (permissionedResolver *PermissionedResolver) UnpackInitializedEvent(log *types.Log) (*PermissionedResolverInitialized, error) {
	event := "Initialized"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != permissionedResolver.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(PermissionedResolverInitialized)
	if len(log.Data) > 0 {
		if err := permissionedResolver.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range permissionedResolver.abi.Events[event].Inputs {
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

// PermissionedResolverInterfaceUpdated represents a InterfaceUpdated event raised by the PermissionedResolver contract.
type PermissionedResolverInterfaceUpdated struct {
	RecordId    *big.Int
	InterfaceId [4]byte
	Implementer common.Address
	Raw         *types.Log // Blockchain specific contextual infos
}

const PermissionedResolverInterfaceUpdatedEventName = "InterfaceUpdated"

// ContractEventName returns the user-defined event name.
func (PermissionedResolverInterfaceUpdated) ContractEventName() string {
	return PermissionedResolverInterfaceUpdatedEventName
}

// UnpackInterfaceUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event InterfaceUpdated(uint256 indexed recordId, bytes4 indexed interfaceId, address implementer)
func (permissionedResolver *PermissionedResolver) UnpackInterfaceUpdatedEvent(log *types.Log) (*PermissionedResolverInterfaceUpdated, error) {
	event := "InterfaceUpdated"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != permissionedResolver.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(PermissionedResolverInterfaceUpdated)
	if len(log.Data) > 0 {
		if err := permissionedResolver.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range permissionedResolver.abi.Events[event].Inputs {
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

// PermissionedResolverLinked represents a Linked event raised by the PermissionedResolver contract.
type PermissionedResolverLinked struct {
	RecordId *big.Int
	Node     [32]byte
	Name     []byte
	Raw      *types.Log // Blockchain specific contextual infos
}

const PermissionedResolverLinkedEventName = "Linked"

// ContractEventName returns the user-defined event name.
func (PermissionedResolverLinked) ContractEventName() string {
	return PermissionedResolverLinkedEventName
}

// UnpackLinkedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Linked(uint256 indexed recordId, bytes32 indexed node, bytes name)
func (permissionedResolver *PermissionedResolver) UnpackLinkedEvent(log *types.Log) (*PermissionedResolverLinked, error) {
	event := "Linked"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != permissionedResolver.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(PermissionedResolverLinked)
	if len(log.Data) > 0 {
		if err := permissionedResolver.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range permissionedResolver.abi.Events[event].Inputs {
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

// PermissionedResolverNameUpdated represents a NameUpdated event raised by the PermissionedResolver contract.
type PermissionedResolverNameUpdated struct {
	RecordId    *big.Int
	PrimaryName string
	Raw         *types.Log // Blockchain specific contextual infos
}

const PermissionedResolverNameUpdatedEventName = "NameUpdated"

// ContractEventName returns the user-defined event name.
func (PermissionedResolverNameUpdated) ContractEventName() string {
	return PermissionedResolverNameUpdatedEventName
}

// UnpackNameUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event NameUpdated(uint256 indexed recordId, string primaryName)
func (permissionedResolver *PermissionedResolver) UnpackNameUpdatedEvent(log *types.Log) (*PermissionedResolverNameUpdated, error) {
	event := "NameUpdated"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != permissionedResolver.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(PermissionedResolverNameUpdated)
	if len(log.Data) > 0 {
		if err := permissionedResolver.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range permissionedResolver.abi.Events[event].Inputs {
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

// PermissionedResolverResolverCreated represents a ResolverCreated event raised by the PermissionedResolver contract.
type PermissionedResolverResolverCreated struct {
	Raw *types.Log // Blockchain specific contextual infos
}

const PermissionedResolverResolverCreatedEventName = "ResolverCreated"

// ContractEventName returns the user-defined event name.
func (PermissionedResolverResolverCreated) ContractEventName() string {
	return PermissionedResolverResolverCreatedEventName
}

// UnpackResolverCreatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ResolverCreated()
func (permissionedResolver *PermissionedResolver) UnpackResolverCreatedEvent(log *types.Log) (*PermissionedResolverResolverCreated, error) {
	event := "ResolverCreated"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != permissionedResolver.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(PermissionedResolverResolverCreated)
	if len(log.Data) > 0 {
		if err := permissionedResolver.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range permissionedResolver.abi.Events[event].Inputs {
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

// PermissionedResolverResourceArgument represents a ResourceArgument event raised by the PermissionedResolver contract.
type PermissionedResolverResourceArgument struct {
	Resource *big.Int
	Arg      []byte
	Raw      *types.Log // Blockchain specific contextual infos
}

const PermissionedResolverResourceArgumentEventName = "ResourceArgument"

// ContractEventName returns the user-defined event name.
func (PermissionedResolverResourceArgument) ContractEventName() string {
	return PermissionedResolverResourceArgumentEventName
}

// UnpackResourceArgumentEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ResourceArgument(uint256 indexed resource, bytes arg)
func (permissionedResolver *PermissionedResolver) UnpackResourceArgumentEvent(log *types.Log) (*PermissionedResolverResourceArgument, error) {
	event := "ResourceArgument"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != permissionedResolver.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(PermissionedResolverResourceArgument)
	if len(log.Data) > 0 {
		if err := permissionedResolver.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range permissionedResolver.abi.Events[event].Inputs {
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

// PermissionedResolverTextUpdated represents a TextUpdated event raised by the PermissionedResolver contract.
type PermissionedResolverTextUpdated struct {
	RecordId *big.Int
	KeyHash  common.Hash
	Key      string
	Value    string
	Raw      *types.Log // Blockchain specific contextual infos
}

const PermissionedResolverTextUpdatedEventName = "TextUpdated"

// ContractEventName returns the user-defined event name.
func (PermissionedResolverTextUpdated) ContractEventName() string {
	return PermissionedResolverTextUpdatedEventName
}

// UnpackTextUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TextUpdated(uint256 indexed recordId, string indexed keyHash, string key, string value)
func (permissionedResolver *PermissionedResolver) UnpackTextUpdatedEvent(log *types.Log) (*PermissionedResolverTextUpdated, error) {
	event := "TextUpdated"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != permissionedResolver.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(PermissionedResolverTextUpdated)
	if len(log.Data) > 0 {
		if err := permissionedResolver.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range permissionedResolver.abi.Events[event].Inputs {
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

// PermissionedResolverUpgraded represents a Upgraded event raised by the PermissionedResolver contract.
type PermissionedResolverUpgraded struct {
	Implementation common.Address
	Raw            *types.Log // Blockchain specific contextual infos
}

const PermissionedResolverUpgradedEventName = "Upgraded"

// ContractEventName returns the user-defined event name.
func (PermissionedResolverUpgraded) ContractEventName() string {
	return PermissionedResolverUpgradedEventName
}

// UnpackUpgradedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Upgraded(address indexed implementation)
func (permissionedResolver *PermissionedResolver) UnpackUpgradedEvent(log *types.Log) (*PermissionedResolverUpgraded, error) {
	event := "Upgraded"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != permissionedResolver.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(PermissionedResolverUpgraded)
	if len(log.Data) > 0 {
		if err := permissionedResolver.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range permissionedResolver.abi.Events[event].Inputs {
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
func (permissionedResolver *PermissionedResolver) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], permissionedResolver.abi.Errors["AddressEmptyCode"].ID.Bytes()[:4]) {
		return permissionedResolver.UnpackAddressEmptyCodeError(raw[4:])
	}
	if bytes.Equal(raw[:4], permissionedResolver.abi.Errors["DNSDecodingFailed"].ID.Bytes()[:4]) {
		return permissionedResolver.UnpackDNSDecodingFailedError(raw[4:])
	}
	if bytes.Equal(raw[:4], permissionedResolver.abi.Errors["EACCannotGrantRoles"].ID.Bytes()[:4]) {
		return permissionedResolver.UnpackEACCannotGrantRolesError(raw[4:])
	}
	if bytes.Equal(raw[:4], permissionedResolver.abi.Errors["EACCannotRevokeRoles"].ID.Bytes()[:4]) {
		return permissionedResolver.UnpackEACCannotRevokeRolesError(raw[4:])
	}
	if bytes.Equal(raw[:4], permissionedResolver.abi.Errors["EACInvalidAccount"].ID.Bytes()[:4]) {
		return permissionedResolver.UnpackEACInvalidAccountError(raw[4:])
	}
	if bytes.Equal(raw[:4], permissionedResolver.abi.Errors["EACInvalidRoleBitmap"].ID.Bytes()[:4]) {
		return permissionedResolver.UnpackEACInvalidRoleBitmapError(raw[4:])
	}
	if bytes.Equal(raw[:4], permissionedResolver.abi.Errors["EACMaxAssignees"].ID.Bytes()[:4]) {
		return permissionedResolver.UnpackEACMaxAssigneesError(raw[4:])
	}
	if bytes.Equal(raw[:4], permissionedResolver.abi.Errors["EACMinAssignees"].ID.Bytes()[:4]) {
		return permissionedResolver.UnpackEACMinAssigneesError(raw[4:])
	}
	if bytes.Equal(raw[:4], permissionedResolver.abi.Errors["EACRootResourceNotAllowed"].ID.Bytes()[:4]) {
		return permissionedResolver.UnpackEACRootResourceNotAllowedError(raw[4:])
	}
	if bytes.Equal(raw[:4], permissionedResolver.abi.Errors["EACUnauthorizedAccountRoles"].ID.Bytes()[:4]) {
		return permissionedResolver.UnpackEACUnauthorizedAccountRolesError(raw[4:])
	}
	if bytes.Equal(raw[:4], permissionedResolver.abi.Errors["ERC1967InvalidImplementation"].ID.Bytes()[:4]) {
		return permissionedResolver.UnpackERC1967InvalidImplementationError(raw[4:])
	}
	if bytes.Equal(raw[:4], permissionedResolver.abi.Errors["ERC1967NonPayable"].ID.Bytes()[:4]) {
		return permissionedResolver.UnpackERC1967NonPayableError(raw[4:])
	}
	if bytes.Equal(raw[:4], permissionedResolver.abi.Errors["FailedCall"].ID.Bytes()[:4]) {
		return permissionedResolver.UnpackFailedCallError(raw[4:])
	}
	if bytes.Equal(raw[:4], permissionedResolver.abi.Errors["InvalidContentType"].ID.Bytes()[:4]) {
		return permissionedResolver.UnpackInvalidContentTypeError(raw[4:])
	}
	if bytes.Equal(raw[:4], permissionedResolver.abi.Errors["InvalidEVMAddress"].ID.Bytes()[:4]) {
		return permissionedResolver.UnpackInvalidEVMAddressError(raw[4:])
	}
	if bytes.Equal(raw[:4], permissionedResolver.abi.Errors["InvalidInitialization"].ID.Bytes()[:4]) {
		return permissionedResolver.UnpackInvalidInitializationError(raw[4:])
	}
	if bytes.Equal(raw[:4], permissionedResolver.abi.Errors["InvalidRecord"].ID.Bytes()[:4]) {
		return permissionedResolver.UnpackInvalidRecordError(raw[4:])
	}
	if bytes.Equal(raw[:4], permissionedResolver.abi.Errors["NotInitializing"].ID.Bytes()[:4]) {
		return permissionedResolver.UnpackNotInitializingError(raw[4:])
	}
	if bytes.Equal(raw[:4], permissionedResolver.abi.Errors["UUPSUnauthorizedCallContext"].ID.Bytes()[:4]) {
		return permissionedResolver.UnpackUUPSUnauthorizedCallContextError(raw[4:])
	}
	if bytes.Equal(raw[:4], permissionedResolver.abi.Errors["UUPSUnsupportedProxiableUUID"].ID.Bytes()[:4]) {
		return permissionedResolver.UnpackUUPSUnsupportedProxiableUUIDError(raw[4:])
	}
	if bytes.Equal(raw[:4], permissionedResolver.abi.Errors["UnsupportedResolverProfile"].ID.Bytes()[:4]) {
		return permissionedResolver.UnpackUnsupportedResolverProfileError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// PermissionedResolverAddressEmptyCode represents a AddressEmptyCode error raised by the PermissionedResolver contract.
type PermissionedResolverAddressEmptyCode struct {
	Target common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AddressEmptyCode(address target)
func PermissionedResolverAddressEmptyCodeErrorID() common.Hash {
	return common.HexToHash("0x9996b315c842ff135b8fc4a08ad5df1c344efbc03d2687aecc0678050d2aac89")
}

// UnpackAddressEmptyCodeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AddressEmptyCode(address target)
func (permissionedResolver *PermissionedResolver) UnpackAddressEmptyCodeError(raw []byte) (*PermissionedResolverAddressEmptyCode, error) {
	out := new(PermissionedResolverAddressEmptyCode)
	if err := permissionedResolver.abi.UnpackIntoInterface(out, "AddressEmptyCode", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionedResolverDNSDecodingFailed represents a DNSDecodingFailed error raised by the PermissionedResolver contract.
type PermissionedResolverDNSDecodingFailed struct {
	Dns []byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error DNSDecodingFailed(bytes dns)
func PermissionedResolverDNSDecodingFailedErrorID() common.Hash {
	return common.HexToHash("0xba4adc2398f6571d4d3236eee815a96fcbbf84bbc4320e35b771331f71afdc72")
}

// UnpackDNSDecodingFailedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error DNSDecodingFailed(bytes dns)
func (permissionedResolver *PermissionedResolver) UnpackDNSDecodingFailedError(raw []byte) (*PermissionedResolverDNSDecodingFailed, error) {
	out := new(PermissionedResolverDNSDecodingFailed)
	if err := permissionedResolver.abi.UnpackIntoInterface(out, "DNSDecodingFailed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionedResolverEACCannotGrantRoles represents a EACCannotGrantRoles error raised by the PermissionedResolver contract.
type PermissionedResolverEACCannotGrantRoles struct {
	Resource   *big.Int
	RoleBitmap *big.Int
	Account    common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error EACCannotGrantRoles(uint256 resource, uint256 roleBitmap, address account)
func PermissionedResolverEACCannotGrantRolesErrorID() common.Hash {
	return common.HexToHash("0xd1a3b355154fa30324f856ca6b25179bcc6c9d2936e28429b1b4293509e97f4a")
}

// UnpackEACCannotGrantRolesError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error EACCannotGrantRoles(uint256 resource, uint256 roleBitmap, address account)
func (permissionedResolver *PermissionedResolver) UnpackEACCannotGrantRolesError(raw []byte) (*PermissionedResolverEACCannotGrantRoles, error) {
	out := new(PermissionedResolverEACCannotGrantRoles)
	if err := permissionedResolver.abi.UnpackIntoInterface(out, "EACCannotGrantRoles", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionedResolverEACCannotRevokeRoles represents a EACCannotRevokeRoles error raised by the PermissionedResolver contract.
type PermissionedResolverEACCannotRevokeRoles struct {
	Resource   *big.Int
	RoleBitmap *big.Int
	Account    common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error EACCannotRevokeRoles(uint256 resource, uint256 roleBitmap, address account)
func PermissionedResolverEACCannotRevokeRolesErrorID() common.Hash {
	return common.HexToHash("0xa604e318421f3889d41e189ce5a1bbb63773eb6844a6d371417f9eed8f6cb153")
}

// UnpackEACCannotRevokeRolesError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error EACCannotRevokeRoles(uint256 resource, uint256 roleBitmap, address account)
func (permissionedResolver *PermissionedResolver) UnpackEACCannotRevokeRolesError(raw []byte) (*PermissionedResolverEACCannotRevokeRoles, error) {
	out := new(PermissionedResolverEACCannotRevokeRoles)
	if err := permissionedResolver.abi.UnpackIntoInterface(out, "EACCannotRevokeRoles", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionedResolverEACInvalidAccount represents a EACInvalidAccount error raised by the PermissionedResolver contract.
type PermissionedResolverEACInvalidAccount struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error EACInvalidAccount()
func PermissionedResolverEACInvalidAccountErrorID() common.Hash {
	return common.HexToHash("0xec3fc5924c67c1c1507473878ec47d207e41f4b88da4d7a835c6e5043b63da51")
}

// UnpackEACInvalidAccountError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error EACInvalidAccount()
func (permissionedResolver *PermissionedResolver) UnpackEACInvalidAccountError(raw []byte) (*PermissionedResolverEACInvalidAccount, error) {
	out := new(PermissionedResolverEACInvalidAccount)
	if err := permissionedResolver.abi.UnpackIntoInterface(out, "EACInvalidAccount", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionedResolverEACInvalidRoleBitmap represents a EACInvalidRoleBitmap error raised by the PermissionedResolver contract.
type PermissionedResolverEACInvalidRoleBitmap struct {
	RoleBitmap *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error EACInvalidRoleBitmap(uint256 roleBitmap)
func PermissionedResolverEACInvalidRoleBitmapErrorID() common.Hash {
	return common.HexToHash("0x2a7b2d20191e44fd4e4227272caf529612678184b58adb291afe5231bf5b2734")
}

// UnpackEACInvalidRoleBitmapError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error EACInvalidRoleBitmap(uint256 roleBitmap)
func (permissionedResolver *PermissionedResolver) UnpackEACInvalidRoleBitmapError(raw []byte) (*PermissionedResolverEACInvalidRoleBitmap, error) {
	out := new(PermissionedResolverEACInvalidRoleBitmap)
	if err := permissionedResolver.abi.UnpackIntoInterface(out, "EACInvalidRoleBitmap", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionedResolverEACMaxAssignees represents a EACMaxAssignees error raised by the PermissionedResolver contract.
type PermissionedResolverEACMaxAssignees struct {
	Resource *big.Int
	Role     *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error EACMaxAssignees(uint256 resource, uint256 role)
func PermissionedResolverEACMaxAssigneesErrorID() common.Hash {
	return common.HexToHash("0xf91653484f73ddf3683475f443ccba4f564362c5896702ddf28889a64960b109")
}

// UnpackEACMaxAssigneesError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error EACMaxAssignees(uint256 resource, uint256 role)
func (permissionedResolver *PermissionedResolver) UnpackEACMaxAssigneesError(raw []byte) (*PermissionedResolverEACMaxAssignees, error) {
	out := new(PermissionedResolverEACMaxAssignees)
	if err := permissionedResolver.abi.UnpackIntoInterface(out, "EACMaxAssignees", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionedResolverEACMinAssignees represents a EACMinAssignees error raised by the PermissionedResolver contract.
type PermissionedResolverEACMinAssignees struct {
	Resource *big.Int
	Role     *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error EACMinAssignees(uint256 resource, uint256 role)
func PermissionedResolverEACMinAssigneesErrorID() common.Hash {
	return common.HexToHash("0x1f80c19ba45f0369eb8ace81d631dde7c47ae7c144205c7e77e1555ff7bb091b")
}

// UnpackEACMinAssigneesError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error EACMinAssignees(uint256 resource, uint256 role)
func (permissionedResolver *PermissionedResolver) UnpackEACMinAssigneesError(raw []byte) (*PermissionedResolverEACMinAssignees, error) {
	out := new(PermissionedResolverEACMinAssignees)
	if err := permissionedResolver.abi.UnpackIntoInterface(out, "EACMinAssignees", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionedResolverEACRootResourceNotAllowed represents a EACRootResourceNotAllowed error raised by the PermissionedResolver contract.
type PermissionedResolverEACRootResourceNotAllowed struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error EACRootResourceNotAllowed()
func PermissionedResolverEACRootResourceNotAllowedErrorID() common.Hash {
	return common.HexToHash("0xc2842458b56deb79dcc9694a98069c21c928dfa5456787ba7ae1728a0d863238")
}

// UnpackEACRootResourceNotAllowedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error EACRootResourceNotAllowed()
func (permissionedResolver *PermissionedResolver) UnpackEACRootResourceNotAllowedError(raw []byte) (*PermissionedResolverEACRootResourceNotAllowed, error) {
	out := new(PermissionedResolverEACRootResourceNotAllowed)
	if err := permissionedResolver.abi.UnpackIntoInterface(out, "EACRootResourceNotAllowed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionedResolverEACUnauthorizedAccountRoles represents a EACUnauthorizedAccountRoles error raised by the PermissionedResolver contract.
type PermissionedResolverEACUnauthorizedAccountRoles struct {
	Resource   *big.Int
	RoleBitmap *big.Int
	Account    common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error EACUnauthorizedAccountRoles(uint256 resource, uint256 roleBitmap, address account)
func PermissionedResolverEACUnauthorizedAccountRolesErrorID() common.Hash {
	return common.HexToHash("0x4b27a13377de460ddbf323b4163e94f59f4092e2fe9befdabaf9bd06c2dd939b")
}

// UnpackEACUnauthorizedAccountRolesError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error EACUnauthorizedAccountRoles(uint256 resource, uint256 roleBitmap, address account)
func (permissionedResolver *PermissionedResolver) UnpackEACUnauthorizedAccountRolesError(raw []byte) (*PermissionedResolverEACUnauthorizedAccountRoles, error) {
	out := new(PermissionedResolverEACUnauthorizedAccountRoles)
	if err := permissionedResolver.abi.UnpackIntoInterface(out, "EACUnauthorizedAccountRoles", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionedResolverERC1967InvalidImplementation represents a ERC1967InvalidImplementation error raised by the PermissionedResolver contract.
type PermissionedResolverERC1967InvalidImplementation struct {
	Implementation common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1967InvalidImplementation(address implementation)
func PermissionedResolverERC1967InvalidImplementationErrorID() common.Hash {
	return common.HexToHash("0x4c9c8ce3ceb3130f17f7cdba48d89b5b0129f266a8bac114e6e315a41879b617")
}

// UnpackERC1967InvalidImplementationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1967InvalidImplementation(address implementation)
func (permissionedResolver *PermissionedResolver) UnpackERC1967InvalidImplementationError(raw []byte) (*PermissionedResolverERC1967InvalidImplementation, error) {
	out := new(PermissionedResolverERC1967InvalidImplementation)
	if err := permissionedResolver.abi.UnpackIntoInterface(out, "ERC1967InvalidImplementation", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionedResolverERC1967NonPayable represents a ERC1967NonPayable error raised by the PermissionedResolver contract.
type PermissionedResolverERC1967NonPayable struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1967NonPayable()
func PermissionedResolverERC1967NonPayableErrorID() common.Hash {
	return common.HexToHash("0xb398979fa84f543c8e222f17890372c487baf85e062276c127fef521eea7224b")
}

// UnpackERC1967NonPayableError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1967NonPayable()
func (permissionedResolver *PermissionedResolver) UnpackERC1967NonPayableError(raw []byte) (*PermissionedResolverERC1967NonPayable, error) {
	out := new(PermissionedResolverERC1967NonPayable)
	if err := permissionedResolver.abi.UnpackIntoInterface(out, "ERC1967NonPayable", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionedResolverFailedCall represents a FailedCall error raised by the PermissionedResolver contract.
type PermissionedResolverFailedCall struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FailedCall()
func PermissionedResolverFailedCallErrorID() common.Hash {
	return common.HexToHash("0xd6bda27508c0fb6d8a39b4b122878dab26f731a7d4e4abe711dd3731899052a4")
}

// UnpackFailedCallError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FailedCall()
func (permissionedResolver *PermissionedResolver) UnpackFailedCallError(raw []byte) (*PermissionedResolverFailedCall, error) {
	out := new(PermissionedResolverFailedCall)
	if err := permissionedResolver.abi.UnpackIntoInterface(out, "FailedCall", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionedResolverInvalidContentType represents a InvalidContentType error raised by the PermissionedResolver contract.
type PermissionedResolverInvalidContentType struct {
	ContentType *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidContentType(uint256 contentType)
func PermissionedResolverInvalidContentTypeErrorID() common.Hash {
	return common.HexToHash("0x5742bb26bee89fc719a606bbf927887fd81c90c2e3bb884a32ca7ffc6c432078")
}

// UnpackInvalidContentTypeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidContentType(uint256 contentType)
func (permissionedResolver *PermissionedResolver) UnpackInvalidContentTypeError(raw []byte) (*PermissionedResolverInvalidContentType, error) {
	out := new(PermissionedResolverInvalidContentType)
	if err := permissionedResolver.abi.UnpackIntoInterface(out, "InvalidContentType", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionedResolverInvalidEVMAddress represents a InvalidEVMAddress error raised by the PermissionedResolver contract.
type PermissionedResolverInvalidEVMAddress struct {
	AddressBytes []byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidEVMAddress(bytes addressBytes)
func PermissionedResolverInvalidEVMAddressErrorID() common.Hash {
	return common.HexToHash("0x8d666f60151e99b4b75dcde8fac04e7b1bfe46b40ae9842f449dbad1817d17b8")
}

// UnpackInvalidEVMAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidEVMAddress(bytes addressBytes)
func (permissionedResolver *PermissionedResolver) UnpackInvalidEVMAddressError(raw []byte) (*PermissionedResolverInvalidEVMAddress, error) {
	out := new(PermissionedResolverInvalidEVMAddress)
	if err := permissionedResolver.abi.UnpackIntoInterface(out, "InvalidEVMAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionedResolverInvalidInitialization represents a InvalidInitialization error raised by the PermissionedResolver contract.
type PermissionedResolverInvalidInitialization struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidInitialization()
func PermissionedResolverInvalidInitializationErrorID() common.Hash {
	return common.HexToHash("0xf92ee8a957075833165f68c320933b1a1294aafc84ee6e0dd3fb178008f9aaf5")
}

// UnpackInvalidInitializationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidInitialization()
func (permissionedResolver *PermissionedResolver) UnpackInvalidInitializationError(raw []byte) (*PermissionedResolverInvalidInitialization, error) {
	out := new(PermissionedResolverInvalidInitialization)
	if err := permissionedResolver.abi.UnpackIntoInterface(out, "InvalidInitialization", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionedResolverInvalidRecord represents a InvalidRecord error raised by the PermissionedResolver contract.
type PermissionedResolverInvalidRecord struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidRecord()
func PermissionedResolverInvalidRecordErrorID() common.Hash {
	return common.HexToHash("0xf2a3e8db448861a247b65de5494e309151be044fd9cd088ac91171257e4ddfef")
}

// UnpackInvalidRecordError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidRecord()
func (permissionedResolver *PermissionedResolver) UnpackInvalidRecordError(raw []byte) (*PermissionedResolverInvalidRecord, error) {
	out := new(PermissionedResolverInvalidRecord)
	if err := permissionedResolver.abi.UnpackIntoInterface(out, "InvalidRecord", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionedResolverNotInitializing represents a NotInitializing error raised by the PermissionedResolver contract.
type PermissionedResolverNotInitializing struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotInitializing()
func PermissionedResolverNotInitializingErrorID() common.Hash {
	return common.HexToHash("0xd7e6bcf8597daa127dc9f0048d2f08d5ef140a2cb659feabd700beff1f7a8302")
}

// UnpackNotInitializingError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotInitializing()
func (permissionedResolver *PermissionedResolver) UnpackNotInitializingError(raw []byte) (*PermissionedResolverNotInitializing, error) {
	out := new(PermissionedResolverNotInitializing)
	if err := permissionedResolver.abi.UnpackIntoInterface(out, "NotInitializing", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionedResolverUUPSUnauthorizedCallContext represents a UUPSUnauthorizedCallContext error raised by the PermissionedResolver contract.
type PermissionedResolverUUPSUnauthorizedCallContext struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UUPSUnauthorizedCallContext()
func PermissionedResolverUUPSUnauthorizedCallContextErrorID() common.Hash {
	return common.HexToHash("0xe07c8dba242a06571ac65fe4bbe20522c9fb111cb33599b799ff8039c1ed18f4")
}

// UnpackUUPSUnauthorizedCallContextError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UUPSUnauthorizedCallContext()
func (permissionedResolver *PermissionedResolver) UnpackUUPSUnauthorizedCallContextError(raw []byte) (*PermissionedResolverUUPSUnauthorizedCallContext, error) {
	out := new(PermissionedResolverUUPSUnauthorizedCallContext)
	if err := permissionedResolver.abi.UnpackIntoInterface(out, "UUPSUnauthorizedCallContext", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionedResolverUUPSUnsupportedProxiableUUID represents a UUPSUnsupportedProxiableUUID error raised by the PermissionedResolver contract.
type PermissionedResolverUUPSUnsupportedProxiableUUID struct {
	Slot [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UUPSUnsupportedProxiableUUID(bytes32 slot)
func PermissionedResolverUUPSUnsupportedProxiableUUIDErrorID() common.Hash {
	return common.HexToHash("0xaa1d49a4c084bfa9aeeee2a0be65267a7f19ba7e1476b114dac513d2c14cb563")
}

// UnpackUUPSUnsupportedProxiableUUIDError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UUPSUnsupportedProxiableUUID(bytes32 slot)
func (permissionedResolver *PermissionedResolver) UnpackUUPSUnsupportedProxiableUUIDError(raw []byte) (*PermissionedResolverUUPSUnsupportedProxiableUUID, error) {
	out := new(PermissionedResolverUUPSUnsupportedProxiableUUID)
	if err := permissionedResolver.abi.UnpackIntoInterface(out, "UUPSUnsupportedProxiableUUID", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionedResolverUnsupportedResolverProfile represents a UnsupportedResolverProfile error raised by the PermissionedResolver contract.
type PermissionedResolverUnsupportedResolverProfile struct {
	Selector [4]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UnsupportedResolverProfile(bytes4 selector)
func PermissionedResolverUnsupportedResolverProfileErrorID() common.Hash {
	return common.HexToHash("0x7b1c461b9d31387598812449409a83d2fcd215e9f43e0d58ccc6333187d37d79")
}

// UnpackUnsupportedResolverProfileError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UnsupportedResolverProfile(bytes4 selector)
func (permissionedResolver *PermissionedResolver) UnpackUnsupportedResolverProfileError(raw []byte) (*PermissionedResolverUnsupportedResolverProfile, error) {
	out := new(PermissionedResolverUnsupportedResolverProfile)
	if err := permissionedResolver.abi.UnpackIntoInterface(out, "UnsupportedResolverProfile", raw); err != nil {
		return nil, err
	}
	return out, nil
}
