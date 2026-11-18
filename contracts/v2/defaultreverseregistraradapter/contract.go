// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package defaultreverseregistraradapter

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

// DefaultReverseRegistrarAdapterMetaData contains all meta data concerning the DefaultReverseRegistrarAdapter contract.
var DefaultReverseRegistrarAdapterMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIDefaultReverseRegistrar\",\"name\":\"defaultReverseRegistrar\",\"type\":\"address\"},{\"internalType\":\"contractIStandaloneHCAFactory\",\"name\":\"standaloneHCAFactory\",\"type\":\"address\"},{\"internalType\":\"contractIContractNamer\",\"name\":\"contractNamer\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"hca\",\"type\":\"address\"}],\"name\":\"HCADeploymentNotTrusted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"hcaOwner\",\"type\":\"address\"}],\"name\":\"HCAOwnerMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"namer\",\"type\":\"address\"}],\"name\":\"UnauthorizedNamer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CONTRACT_NAMER\",\"outputs\":[{\"internalType\":\"contractIContractNamer\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_REVERSE_REGISTRAR\",\"outputs\":[{\"internalType\":\"contractIDefaultReverseRegistrar\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STANDALONE_HCA_FACTORY\",\"outputs\":[{\"internalType\":\"contractIStandaloneHCAFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"namer\",\"type\":\"address\"}],\"name\":\"isContractNamer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"setNameWithHCA\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	ID:  "DefaultReverseRegistrarAdapter",
}

// DefaultReverseRegistrarAdapter is an auto generated Go binding around an Ethereum contract.
type DefaultReverseRegistrarAdapter struct {
	abi abi.ABI
}

// GetABI returns the ABI associated with this contract binding.
func (c *DefaultReverseRegistrarAdapter) GetABI() abi.ABI {
	return c.abi
}

// NewDefaultReverseRegistrarAdapter creates a new instance of DefaultReverseRegistrarAdapter.
func NewDefaultReverseRegistrarAdapter() *DefaultReverseRegistrarAdapter {
	parsed, err := DefaultReverseRegistrarAdapterMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &DefaultReverseRegistrarAdapter{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *DefaultReverseRegistrarAdapter) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address defaultReverseRegistrar, address standaloneHCAFactory, address contractNamer) returns()
func (defaultReverseRegistrarAdapter *DefaultReverseRegistrarAdapter) PackConstructor(defaultReverseRegistrar common.Address, standaloneHCAFactory common.Address, contractNamer common.Address) []byte {
	enc, err := defaultReverseRegistrarAdapter.abi.Pack("", defaultReverseRegistrar, standaloneHCAFactory, contractNamer)
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
func (defaultReverseRegistrarAdapter *DefaultReverseRegistrarAdapter) PackCONTRACTNAMER() []byte {
	enc, err := defaultReverseRegistrarAdapter.abi.Pack("CONTRACT_NAMER")
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
func (defaultReverseRegistrarAdapter *DefaultReverseRegistrarAdapter) TryPackCONTRACTNAMER() ([]byte, error) {
	return defaultReverseRegistrarAdapter.abi.Pack("CONTRACT_NAMER")
}

// UnpackCONTRACTNAMER is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x48ee1bcc.
//
// Solidity: function CONTRACT_NAMER() view returns(address)
func (defaultReverseRegistrarAdapter *DefaultReverseRegistrarAdapter) UnpackCONTRACTNAMER(data []byte) (common.Address, error) {
	out, err := defaultReverseRegistrarAdapter.abi.Unpack("CONTRACT_NAMER", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackDEFAULTREVERSEREGISTRAR is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8633886d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function DEFAULT_REVERSE_REGISTRAR() view returns(address)
func (defaultReverseRegistrarAdapter *DefaultReverseRegistrarAdapter) PackDEFAULTREVERSEREGISTRAR() []byte {
	enc, err := defaultReverseRegistrarAdapter.abi.Pack("DEFAULT_REVERSE_REGISTRAR")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDEFAULTREVERSEREGISTRAR is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8633886d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function DEFAULT_REVERSE_REGISTRAR() view returns(address)
func (defaultReverseRegistrarAdapter *DefaultReverseRegistrarAdapter) TryPackDEFAULTREVERSEREGISTRAR() ([]byte, error) {
	return defaultReverseRegistrarAdapter.abi.Pack("DEFAULT_REVERSE_REGISTRAR")
}

// UnpackDEFAULTREVERSEREGISTRAR is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8633886d.
//
// Solidity: function DEFAULT_REVERSE_REGISTRAR() view returns(address)
func (defaultReverseRegistrarAdapter *DefaultReverseRegistrarAdapter) UnpackDEFAULTREVERSEREGISTRAR(data []byte) (common.Address, error) {
	out, err := defaultReverseRegistrarAdapter.abi.Unpack("DEFAULT_REVERSE_REGISTRAR", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackSTANDALONEHCAFACTORY is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x976d13cf.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function STANDALONE_HCA_FACTORY() view returns(address)
func (defaultReverseRegistrarAdapter *DefaultReverseRegistrarAdapter) PackSTANDALONEHCAFACTORY() []byte {
	enc, err := defaultReverseRegistrarAdapter.abi.Pack("STANDALONE_HCA_FACTORY")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSTANDALONEHCAFACTORY is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x976d13cf.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function STANDALONE_HCA_FACTORY() view returns(address)
func (defaultReverseRegistrarAdapter *DefaultReverseRegistrarAdapter) TryPackSTANDALONEHCAFACTORY() ([]byte, error) {
	return defaultReverseRegistrarAdapter.abi.Pack("STANDALONE_HCA_FACTORY")
}

// UnpackSTANDALONEHCAFACTORY is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x976d13cf.
//
// Solidity: function STANDALONE_HCA_FACTORY() view returns(address)
func (defaultReverseRegistrarAdapter *DefaultReverseRegistrarAdapter) UnpackSTANDALONEHCAFACTORY(data []byte) (common.Address, error) {
	out, err := defaultReverseRegistrarAdapter.abi.Unpack("STANDALONE_HCA_FACTORY", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackIsContractNamer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6f3ff726.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isContractNamer(address namer) view returns(bool)
func (defaultReverseRegistrarAdapter *DefaultReverseRegistrarAdapter) PackIsContractNamer(namer common.Address) []byte {
	enc, err := defaultReverseRegistrarAdapter.abi.Pack("isContractNamer", namer)
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
func (defaultReverseRegistrarAdapter *DefaultReverseRegistrarAdapter) TryPackIsContractNamer(namer common.Address) ([]byte, error) {
	return defaultReverseRegistrarAdapter.abi.Pack("isContractNamer", namer)
}

// UnpackIsContractNamer is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6f3ff726.
//
// Solidity: function isContractNamer(address namer) view returns(bool)
func (defaultReverseRegistrarAdapter *DefaultReverseRegistrarAdapter) UnpackIsContractNamer(data []byte) (bool, error) {
	out, err := defaultReverseRegistrarAdapter.abi.Unpack("isContractNamer", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackSetName is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3121db1c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setName(address account, string name) returns()
func (defaultReverseRegistrarAdapter *DefaultReverseRegistrarAdapter) PackSetName(account common.Address, name string) []byte {
	enc, err := defaultReverseRegistrarAdapter.abi.Pack("setName", account, name)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetName is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3121db1c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setName(address account, string name) returns()
func (defaultReverseRegistrarAdapter *DefaultReverseRegistrarAdapter) TryPackSetName(account common.Address, name string) ([]byte, error) {
	return defaultReverseRegistrarAdapter.abi.Pack("setName", account, name)
}

// PackSetNameWithHCA is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xab863445.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setNameWithHCA(address account, string name) returns()
func (defaultReverseRegistrarAdapter *DefaultReverseRegistrarAdapter) PackSetNameWithHCA(account common.Address, name string) []byte {
	enc, err := defaultReverseRegistrarAdapter.abi.Pack("setNameWithHCA", account, name)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetNameWithHCA is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xab863445.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setNameWithHCA(address account, string name) returns()
func (defaultReverseRegistrarAdapter *DefaultReverseRegistrarAdapter) TryPackSetNameWithHCA(account common.Address, name string) ([]byte, error) {
	return defaultReverseRegistrarAdapter.abi.Pack("setNameWithHCA", account, name)
}

// PackSupportsInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01ffc9a7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (defaultReverseRegistrarAdapter *DefaultReverseRegistrarAdapter) PackSupportsInterface(interfaceId [4]byte) []byte {
	enc, err := defaultReverseRegistrarAdapter.abi.Pack("supportsInterface", interfaceId)
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
func (defaultReverseRegistrarAdapter *DefaultReverseRegistrarAdapter) TryPackSupportsInterface(interfaceId [4]byte) ([]byte, error) {
	return defaultReverseRegistrarAdapter.abi.Pack("supportsInterface", interfaceId)
}

// UnpackSupportsInterface is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (defaultReverseRegistrarAdapter *DefaultReverseRegistrarAdapter) UnpackSupportsInterface(data []byte) (bool, error) {
	out, err := defaultReverseRegistrarAdapter.abi.Unpack("supportsInterface", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// UnpackError attempts to decode the provided error data using user-defined
// error definitions.
func (defaultReverseRegistrarAdapter *DefaultReverseRegistrarAdapter) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], defaultReverseRegistrarAdapter.abi.Errors["HCADeploymentNotTrusted"].ID.Bytes()[:4]) {
		return defaultReverseRegistrarAdapter.UnpackHCADeploymentNotTrustedError(raw[4:])
	}
	if bytes.Equal(raw[:4], defaultReverseRegistrarAdapter.abi.Errors["HCAOwnerMismatch"].ID.Bytes()[:4]) {
		return defaultReverseRegistrarAdapter.UnpackHCAOwnerMismatchError(raw[4:])
	}
	if bytes.Equal(raw[:4], defaultReverseRegistrarAdapter.abi.Errors["UnauthorizedNamer"].ID.Bytes()[:4]) {
		return defaultReverseRegistrarAdapter.UnpackUnauthorizedNamerError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// DefaultReverseRegistrarAdapterHCADeploymentNotTrusted represents a HCADeploymentNotTrusted error raised by the DefaultReverseRegistrarAdapter contract.
type DefaultReverseRegistrarAdapterHCADeploymentNotTrusted struct {
	Hca common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error HCADeploymentNotTrusted(address hca)
func DefaultReverseRegistrarAdapterHCADeploymentNotTrustedErrorID() common.Hash {
	return common.HexToHash("0x8d859048b6974ccb893e3552b02a721575409f6ee5f6e1a7fb625023605daa84")
}

// UnpackHCADeploymentNotTrustedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error HCADeploymentNotTrusted(address hca)
func (defaultReverseRegistrarAdapter *DefaultReverseRegistrarAdapter) UnpackHCADeploymentNotTrustedError(raw []byte) (*DefaultReverseRegistrarAdapterHCADeploymentNotTrusted, error) {
	out := new(DefaultReverseRegistrarAdapterHCADeploymentNotTrusted)
	if err := defaultReverseRegistrarAdapter.abi.UnpackIntoInterface(out, "HCADeploymentNotTrusted", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// DefaultReverseRegistrarAdapterHCAOwnerMismatch represents a HCAOwnerMismatch error raised by the DefaultReverseRegistrarAdapter contract.
type DefaultReverseRegistrarAdapterHCAOwnerMismatch struct {
	Account  common.Address
	HcaOwner common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error HCAOwnerMismatch(address account, address hcaOwner)
func DefaultReverseRegistrarAdapterHCAOwnerMismatchErrorID() common.Hash {
	return common.HexToHash("0xa8f2205f5ffac362d6c601309b57ff6751aaeb24ea70a3e47dddcdbabf63d8e7")
}

// UnpackHCAOwnerMismatchError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error HCAOwnerMismatch(address account, address hcaOwner)
func (defaultReverseRegistrarAdapter *DefaultReverseRegistrarAdapter) UnpackHCAOwnerMismatchError(raw []byte) (*DefaultReverseRegistrarAdapterHCAOwnerMismatch, error) {
	out := new(DefaultReverseRegistrarAdapterHCAOwnerMismatch)
	if err := defaultReverseRegistrarAdapter.abi.UnpackIntoInterface(out, "HCAOwnerMismatch", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// DefaultReverseRegistrarAdapterUnauthorizedNamer represents a UnauthorizedNamer error raised by the DefaultReverseRegistrarAdapter contract.
type DefaultReverseRegistrarAdapterUnauthorizedNamer struct {
	Namer common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UnauthorizedNamer(address namer)
func DefaultReverseRegistrarAdapterUnauthorizedNamerErrorID() common.Hash {
	return common.HexToHash("0x0d1b7e4e95d337c8403b427b1e83640e20d09f88944b7647f5d011301f680698")
}

// UnpackUnauthorizedNamerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UnauthorizedNamer(address namer)
func (defaultReverseRegistrarAdapter *DefaultReverseRegistrarAdapter) UnpackUnauthorizedNamerError(raw []byte) (*DefaultReverseRegistrarAdapterUnauthorizedNamer, error) {
	out := new(DefaultReverseRegistrarAdapterUnauthorizedNamer)
	if err := defaultReverseRegistrarAdapter.abi.UnpackIntoInterface(out, "UnauthorizedNamer", raw); err != nil {
		return nil, err
	}
	return out, nil
}
