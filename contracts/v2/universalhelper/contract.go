// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package universalhelper

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

// UniversalHelperMetaData contains all meta data concerning the UniversalHelper contract.
var UniversalHelperMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPermissionedRegistry\",\"name\":\"rootRegistry\",\"type\":\"address\"},{\"internalType\":\"contractIContractNamer\",\"name\":\"contractNamer\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"dns\",\"type\":\"bytes\"}],\"name\":\"DNSDecodingFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LabelIsEmpty\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"}],\"name\":\"LabelIsTooLong\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIRegistry\",\"name\":\"registry\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"errorData\",\"type\":\"bytes\"}],\"name\":\"ResolutionError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CONTRACT_NAMER\",\"outputs\":[{\"internalType\":\"contractIContractNamer\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ROOT_REGISTRY\",\"outputs\":[{\"internalType\":\"contractIPermissionedRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIRegistry\",\"name\":\"registry\",\"type\":\"address\"}],\"name\":\"findCanonicalName\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"}],\"name\":\"findCanonicalRegistry\",\"outputs\":[{\"internalType\":\"contractIRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"}],\"name\":\"findExactOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"}],\"name\":\"findExactRegistry\",\"outputs\":[{\"internalType\":\"contractIRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"}],\"name\":\"findNearestOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"}],\"name\":\"findNearestRegistry\",\"outputs\":[{\"internalType\":\"contractIRegistry\",\"name\":\"registry\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"}],\"name\":\"findParentRegistry\",\"outputs\":[{\"internalType\":\"contractIRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"}],\"name\":\"findRegistries\",\"outputs\":[{\"internalType\":\"contractIRegistry[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"namer\",\"type\":\"address\"}],\"name\":\"isContractNamer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	ID:  "UniversalHelper",
}

// UniversalHelper is an auto generated Go binding around an Ethereum contract.
type UniversalHelper struct {
	abi abi.ABI
}

// GetABI returns the ABI associated with this contract binding.
func (c *UniversalHelper) GetABI() abi.ABI {
	return c.abi
}

// NewUniversalHelper creates a new instance of UniversalHelper.
func NewUniversalHelper() *UniversalHelper {
	parsed, err := UniversalHelperMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &UniversalHelper{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *UniversalHelper) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address rootRegistry, address contractNamer) returns()
func (universalHelper *UniversalHelper) PackConstructor(rootRegistry common.Address, contractNamer common.Address) []byte {
	enc, err := universalHelper.abi.Pack("", rootRegistry, contractNamer)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackCONTRACTNAMER is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x48ee1bcc.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function CONTRACT_NAMER() view returns(address)
func (universalHelper *UniversalHelper) PackCONTRACTNAMER() []byte {
	enc, err := universalHelper.abi.Pack("CONTRACT_NAMER")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCONTRACTNAMER is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x48ee1bcc.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function CONTRACT_NAMER() view returns(address)
func (universalHelper *UniversalHelper) TryPackCONTRACTNAMER() ([]byte, error) {
	return universalHelper.abi.Pack("CONTRACT_NAMER")
}

// UnpackCONTRACTNAMER is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x48ee1bcc.
//
// Solidity: function CONTRACT_NAMER() view returns(address)
func (universalHelper *UniversalHelper) UnpackCONTRACTNAMER(data []byte) (common.Address, error) {
	out, err := universalHelper.abi.Unpack("CONTRACT_NAMER", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackROOTREGISTRY is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc92cc49a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function ROOT_REGISTRY() view returns(address)
func (universalHelper *UniversalHelper) PackROOTREGISTRY() []byte {
	enc, err := universalHelper.abi.Pack("ROOT_REGISTRY")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackROOTREGISTRY is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc92cc49a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function ROOT_REGISTRY() view returns(address)
func (universalHelper *UniversalHelper) TryPackROOTREGISTRY() ([]byte, error) {
	return universalHelper.abi.Pack("ROOT_REGISTRY")
}

// UnpackROOTREGISTRY is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc92cc49a.
//
// Solidity: function ROOT_REGISTRY() view returns(address)
func (universalHelper *UniversalHelper) UnpackROOTREGISTRY(data []byte) (common.Address, error) {
	out, err := universalHelper.abi.Unpack("ROOT_REGISTRY", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackFindCanonicalName is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x83a64339.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function findCanonicalName(address registry) view returns(bytes)
func (universalHelper *UniversalHelper) PackFindCanonicalName(registry common.Address) []byte {
	enc, err := universalHelper.abi.Pack("findCanonicalName", registry)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFindCanonicalName is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x83a64339.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function findCanonicalName(address registry) view returns(bytes)
func (universalHelper *UniversalHelper) TryPackFindCanonicalName(registry common.Address) ([]byte, error) {
	return universalHelper.abi.Pack("findCanonicalName", registry)
}

// UnpackFindCanonicalName is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x83a64339.
//
// Solidity: function findCanonicalName(address registry) view returns(bytes)
func (universalHelper *UniversalHelper) UnpackFindCanonicalName(data []byte) ([]byte, error) {
	out, err := universalHelper.abi.Unpack("findCanonicalName", data)
	if err != nil {
		return *new([]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	return out0, nil
}

// PackFindCanonicalRegistry is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb363cc73.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function findCanonicalRegistry(bytes name) view returns(address)
func (universalHelper *UniversalHelper) PackFindCanonicalRegistry(name []byte) []byte {
	enc, err := universalHelper.abi.Pack("findCanonicalRegistry", name)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFindCanonicalRegistry is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb363cc73.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function findCanonicalRegistry(bytes name) view returns(address)
func (universalHelper *UniversalHelper) TryPackFindCanonicalRegistry(name []byte) ([]byte, error) {
	return universalHelper.abi.Pack("findCanonicalRegistry", name)
}

// UnpackFindCanonicalRegistry is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb363cc73.
//
// Solidity: function findCanonicalRegistry(bytes name) view returns(address)
func (universalHelper *UniversalHelper) UnpackFindCanonicalRegistry(data []byte) (common.Address, error) {
	out, err := universalHelper.abi.Unpack("findCanonicalRegistry", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackFindExactOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x78b8187f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function findExactOwner(bytes name) view returns(address)
func (universalHelper *UniversalHelper) PackFindExactOwner(name []byte) []byte {
	enc, err := universalHelper.abi.Pack("findExactOwner", name)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFindExactOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x78b8187f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function findExactOwner(bytes name) view returns(address)
func (universalHelper *UniversalHelper) TryPackFindExactOwner(name []byte) ([]byte, error) {
	return universalHelper.abi.Pack("findExactOwner", name)
}

// UnpackFindExactOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x78b8187f.
//
// Solidity: function findExactOwner(bytes name) view returns(address)
func (universalHelper *UniversalHelper) UnpackFindExactOwner(data []byte) (common.Address, error) {
	out, err := universalHelper.abi.Unpack("findExactOwner", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackFindExactRegistry is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4878c6dd.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function findExactRegistry(bytes name) view returns(address)
func (universalHelper *UniversalHelper) PackFindExactRegistry(name []byte) []byte {
	enc, err := universalHelper.abi.Pack("findExactRegistry", name)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFindExactRegistry is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4878c6dd.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function findExactRegistry(bytes name) view returns(address)
func (universalHelper *UniversalHelper) TryPackFindExactRegistry(name []byte) ([]byte, error) {
	return universalHelper.abi.Pack("findExactRegistry", name)
}

// UnpackFindExactRegistry is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x4878c6dd.
//
// Solidity: function findExactRegistry(bytes name) view returns(address)
func (universalHelper *UniversalHelper) UnpackFindExactRegistry(data []byte) (common.Address, error) {
	out, err := universalHelper.abi.Unpack("findExactRegistry", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackFindNearestOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4f50b731.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function findNearestOwner(bytes name) view returns(address owner, uint256 offset)
func (universalHelper *UniversalHelper) PackFindNearestOwner(name []byte) []byte {
	enc, err := universalHelper.abi.Pack("findNearestOwner", name)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFindNearestOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4f50b731.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function findNearestOwner(bytes name) view returns(address owner, uint256 offset)
func (universalHelper *UniversalHelper) TryPackFindNearestOwner(name []byte) ([]byte, error) {
	return universalHelper.abi.Pack("findNearestOwner", name)
}

// FindNearestOwnerOutput serves as a container for the return parameters of contract
// method FindNearestOwner.
type FindNearestOwnerOutput struct {
	Owner  common.Address
	Offset *big.Int
}

// UnpackFindNearestOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x4f50b731.
//
// Solidity: function findNearestOwner(bytes name) view returns(address owner, uint256 offset)
func (universalHelper *UniversalHelper) UnpackFindNearestOwner(data []byte) (FindNearestOwnerOutput, error) {
	out, err := universalHelper.abi.Unpack("findNearestOwner", data)
	outstruct := new(FindNearestOwnerOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Offset = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	return *outstruct, nil
}

// PackFindNearestRegistry is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5acaafb9.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function findNearestRegistry(bytes name) view returns(address registry, uint256 offset)
func (universalHelper *UniversalHelper) PackFindNearestRegistry(name []byte) []byte {
	enc, err := universalHelper.abi.Pack("findNearestRegistry", name)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFindNearestRegistry is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5acaafb9.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function findNearestRegistry(bytes name) view returns(address registry, uint256 offset)
func (universalHelper *UniversalHelper) TryPackFindNearestRegistry(name []byte) ([]byte, error) {
	return universalHelper.abi.Pack("findNearestRegistry", name)
}

// FindNearestRegistryOutput serves as a container for the return parameters of contract
// method FindNearestRegistry.
type FindNearestRegistryOutput struct {
	Registry common.Address
	Offset   *big.Int
}

// UnpackFindNearestRegistry is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5acaafb9.
//
// Solidity: function findNearestRegistry(bytes name) view returns(address registry, uint256 offset)
func (universalHelper *UniversalHelper) UnpackFindNearestRegistry(data []byte) (FindNearestRegistryOutput, error) {
	out, err := universalHelper.abi.Unpack("findNearestRegistry", data)
	outstruct := new(FindNearestRegistryOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Registry = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Offset = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	return *outstruct, nil
}

// PackFindParentRegistry is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe4f8ce05.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function findParentRegistry(bytes name) view returns(address)
func (universalHelper *UniversalHelper) PackFindParentRegistry(name []byte) []byte {
	enc, err := universalHelper.abi.Pack("findParentRegistry", name)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFindParentRegistry is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe4f8ce05.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function findParentRegistry(bytes name) view returns(address)
func (universalHelper *UniversalHelper) TryPackFindParentRegistry(name []byte) ([]byte, error) {
	return universalHelper.abi.Pack("findParentRegistry", name)
}

// UnpackFindParentRegistry is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe4f8ce05.
//
// Solidity: function findParentRegistry(bytes name) view returns(address)
func (universalHelper *UniversalHelper) UnpackFindParentRegistry(data []byte) (common.Address, error) {
	out, err := universalHelper.abi.Unpack("findParentRegistry", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackFindRegistries is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf272e2af.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function findRegistries(bytes name) view returns(address[])
func (universalHelper *UniversalHelper) PackFindRegistries(name []byte) []byte {
	enc, err := universalHelper.abi.Pack("findRegistries", name)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFindRegistries is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf272e2af.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function findRegistries(bytes name) view returns(address[])
func (universalHelper *UniversalHelper) TryPackFindRegistries(name []byte) ([]byte, error) {
	return universalHelper.abi.Pack("findRegistries", name)
}

// UnpackFindRegistries is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf272e2af.
//
// Solidity: function findRegistries(bytes name) view returns(address[])
func (universalHelper *UniversalHelper) UnpackFindRegistries(data []byte) ([]common.Address, error) {
	out, err := universalHelper.abi.Unpack("findRegistries", data)
	if err != nil {
		return *new([]common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	return out0, nil
}

// PackIsContractNamer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6f3ff726.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isContractNamer(address namer) view returns(bool)
func (universalHelper *UniversalHelper) PackIsContractNamer(namer common.Address) []byte {
	enc, err := universalHelper.abi.Pack("isContractNamer", namer)
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
func (universalHelper *UniversalHelper) TryPackIsContractNamer(namer common.Address) ([]byte, error) {
	return universalHelper.abi.Pack("isContractNamer", namer)
}

// UnpackIsContractNamer is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6f3ff726.
//
// Solidity: function isContractNamer(address namer) view returns(bool)
func (universalHelper *UniversalHelper) UnpackIsContractNamer(data []byte) (bool, error) {
	out, err := universalHelper.abi.Unpack("isContractNamer", data)
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
func (universalHelper *UniversalHelper) PackSupportsInterface(interfaceId [4]byte) []byte {
	enc, err := universalHelper.abi.Pack("supportsInterface", interfaceId)
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
func (universalHelper *UniversalHelper) TryPackSupportsInterface(interfaceId [4]byte) ([]byte, error) {
	return universalHelper.abi.Pack("supportsInterface", interfaceId)
}

// UnpackSupportsInterface is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (universalHelper *UniversalHelper) UnpackSupportsInterface(data []byte) (bool, error) {
	out, err := universalHelper.abi.Unpack("supportsInterface", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// UnpackError attempts to decode the provided error data using user-defined
// error definitions.
func (universalHelper *UniversalHelper) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], universalHelper.abi.Errors["DNSDecodingFailed"].ID.Bytes()[:4]) {
		return universalHelper.UnpackDNSDecodingFailedError(raw[4:])
	}
	if bytes.Equal(raw[:4], universalHelper.abi.Errors["LabelIsEmpty"].ID.Bytes()[:4]) {
		return universalHelper.UnpackLabelIsEmptyError(raw[4:])
	}
	if bytes.Equal(raw[:4], universalHelper.abi.Errors["LabelIsTooLong"].ID.Bytes()[:4]) {
		return universalHelper.UnpackLabelIsTooLongError(raw[4:])
	}
	if bytes.Equal(raw[:4], universalHelper.abi.Errors["ResolutionError"].ID.Bytes()[:4]) {
		return universalHelper.UnpackResolutionErrorError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// UniversalHelperDNSDecodingFailed represents a DNSDecodingFailed error raised by the UniversalHelper contract.
type UniversalHelperDNSDecodingFailed struct {
	Dns []byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error DNSDecodingFailed(bytes dns)
func UniversalHelperDNSDecodingFailedErrorID() common.Hash {
	return common.HexToHash("0xba4adc2398f6571d4d3236eee815a96fcbbf84bbc4320e35b771331f71afdc72")
}

// UnpackDNSDecodingFailedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error DNSDecodingFailed(bytes dns)
func (universalHelper *UniversalHelper) UnpackDNSDecodingFailedError(raw []byte) (*UniversalHelperDNSDecodingFailed, error) {
	out := new(UniversalHelperDNSDecodingFailed)
	if err := universalHelper.abi.UnpackIntoInterface(out, "DNSDecodingFailed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// UniversalHelperLabelIsEmpty represents a LabelIsEmpty error raised by the UniversalHelper contract.
type UniversalHelperLabelIsEmpty struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error LabelIsEmpty()
func UniversalHelperLabelIsEmptyErrorID() common.Hash {
	return common.HexToHash("0xbf9a27402cc67e0a1553ebafddb9a8f09e2a0e17d33881522b49406d3206b53d")
}

// UnpackLabelIsEmptyError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error LabelIsEmpty()
func (universalHelper *UniversalHelper) UnpackLabelIsEmptyError(raw []byte) (*UniversalHelperLabelIsEmpty, error) {
	out := new(UniversalHelperLabelIsEmpty)
	if err := universalHelper.abi.UnpackIntoInterface(out, "LabelIsEmpty", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// UniversalHelperLabelIsTooLong represents a LabelIsTooLong error raised by the UniversalHelper contract.
type UniversalHelperLabelIsTooLong struct {
	Label string
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error LabelIsTooLong(string label)
func UniversalHelperLabelIsTooLongErrorID() common.Hash {
	return common.HexToHash("0xdab6c73ccd5b3a9a45fe88ea48179efbe644dbc9416a195d7779ed8ccd248e63")
}

// UnpackLabelIsTooLongError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error LabelIsTooLong(string label)
func (universalHelper *UniversalHelper) UnpackLabelIsTooLongError(raw []byte) (*UniversalHelperLabelIsTooLong, error) {
	out := new(UniversalHelperLabelIsTooLong)
	if err := universalHelper.abi.UnpackIntoInterface(out, "LabelIsTooLong", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// UniversalHelperResolutionError represents a ResolutionError error raised by the UniversalHelper contract.
type UniversalHelperResolutionError struct {
	Registry  common.Address
	ErrorData []byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ResolutionError(address registry, bytes errorData)
func UniversalHelperResolutionErrorErrorID() common.Hash {
	return common.HexToHash("0x4e8547e9ab8fa91db4a07393325f95c7d6bd9e823f4d7816b99b5eb3582b6d9f")
}

// UnpackResolutionErrorError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ResolutionError(address registry, bytes errorData)
func (universalHelper *UniversalHelper) UnpackResolutionErrorError(raw []byte) (*UniversalHelperResolutionError, error) {
	out := new(UniversalHelperResolutionError)
	if err := universalHelper.abi.UnpackIntoInterface(out, "ResolutionError", raw); err != nil {
		return nil, err
	}
	return out, nil
}
