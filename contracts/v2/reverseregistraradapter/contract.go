// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package reverseregistraradapter

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

// ReverseRegistrarAdapterMetaData contains all meta data concerning the ReverseRegistrarAdapter contract.
var ReverseRegistrarAdapterMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIReverseRegistrar\",\"name\":\"reverseRegistrar\",\"type\":\"address\"},{\"internalType\":\"contractIStandaloneHCAFactory\",\"name\":\"standaloneHCAFactory\",\"type\":\"address\"},{\"internalType\":\"contractIContractNamer\",\"name\":\"contractNamer\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"hca\",\"type\":\"address\"}],\"name\":\"HCADeploymentNotTrusted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"hcaOwner\",\"type\":\"address\"}],\"name\":\"HCAOwnerMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"namer\",\"type\":\"address\"}],\"name\":\"UnauthorizedNamer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CONTRACT_NAMER\",\"outputs\":[{\"internalType\":\"contractIContractNamer\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REVERSE_REGISTRAR\",\"outputs\":[{\"internalType\":\"contractIReverseRegistrar\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STANDALONE_HCA_FACTORY\",\"outputs\":[{\"internalType\":\"contractIStandaloneHCAFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"}],\"name\":\"claim\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"}],\"name\":\"claimWithHCA\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"namer\",\"type\":\"address\"}],\"name\":\"isContractNamer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	ID:  "ReverseRegistrarAdapter",
}

// ReverseRegistrarAdapter is an auto generated Go binding around an Ethereum contract.
type ReverseRegistrarAdapter struct {
	abi abi.ABI
}

// GetABI returns the ABI associated with this contract binding.
func (c *ReverseRegistrarAdapter) GetABI() abi.ABI {
	return c.abi
}

// NewReverseRegistrarAdapter creates a new instance of ReverseRegistrarAdapter.
func NewReverseRegistrarAdapter() *ReverseRegistrarAdapter {
	parsed, err := ReverseRegistrarAdapterMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &ReverseRegistrarAdapter{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *ReverseRegistrarAdapter) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address reverseRegistrar, address standaloneHCAFactory, address contractNamer) returns()
func (reverseRegistrarAdapter *ReverseRegistrarAdapter) PackConstructor(reverseRegistrar common.Address, standaloneHCAFactory common.Address, contractNamer common.Address) []byte {
	enc, err := reverseRegistrarAdapter.abi.Pack("", reverseRegistrar, standaloneHCAFactory, contractNamer)
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
func (reverseRegistrarAdapter *ReverseRegistrarAdapter) PackCONTRACTNAMER() []byte {
	enc, err := reverseRegistrarAdapter.abi.Pack("CONTRACT_NAMER")
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
func (reverseRegistrarAdapter *ReverseRegistrarAdapter) TryPackCONTRACTNAMER() ([]byte, error) {
	return reverseRegistrarAdapter.abi.Pack("CONTRACT_NAMER")
}

// UnpackCONTRACTNAMER is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x48ee1bcc.
//
// Solidity: function CONTRACT_NAMER() view returns(address)
func (reverseRegistrarAdapter *ReverseRegistrarAdapter) UnpackCONTRACTNAMER(data []byte) (common.Address, error) {
	out, err := reverseRegistrarAdapter.abi.Unpack("CONTRACT_NAMER", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackREVERSEREGISTRAR is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x952899fc.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function REVERSE_REGISTRAR() view returns(address)
func (reverseRegistrarAdapter *ReverseRegistrarAdapter) PackREVERSEREGISTRAR() []byte {
	enc, err := reverseRegistrarAdapter.abi.Pack("REVERSE_REGISTRAR")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackREVERSEREGISTRAR is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x952899fc.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function REVERSE_REGISTRAR() view returns(address)
func (reverseRegistrarAdapter *ReverseRegistrarAdapter) TryPackREVERSEREGISTRAR() ([]byte, error) {
	return reverseRegistrarAdapter.abi.Pack("REVERSE_REGISTRAR")
}

// UnpackREVERSEREGISTRAR is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x952899fc.
//
// Solidity: function REVERSE_REGISTRAR() view returns(address)
func (reverseRegistrarAdapter *ReverseRegistrarAdapter) UnpackREVERSEREGISTRAR(data []byte) (common.Address, error) {
	out, err := reverseRegistrarAdapter.abi.Unpack("REVERSE_REGISTRAR", data)
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
func (reverseRegistrarAdapter *ReverseRegistrarAdapter) PackSTANDALONEHCAFACTORY() []byte {
	enc, err := reverseRegistrarAdapter.abi.Pack("STANDALONE_HCA_FACTORY")
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
func (reverseRegistrarAdapter *ReverseRegistrarAdapter) TryPackSTANDALONEHCAFACTORY() ([]byte, error) {
	return reverseRegistrarAdapter.abi.Pack("STANDALONE_HCA_FACTORY")
}

// UnpackSTANDALONEHCAFACTORY is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x976d13cf.
//
// Solidity: function STANDALONE_HCA_FACTORY() view returns(address)
func (reverseRegistrarAdapter *ReverseRegistrarAdapter) UnpackSTANDALONEHCAFACTORY(data []byte) (common.Address, error) {
	out, err := reverseRegistrarAdapter.abi.Unpack("STANDALONE_HCA_FACTORY", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackClaim is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x21c0b342.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function claim(address account, address resolver) returns(bytes32)
func (reverseRegistrarAdapter *ReverseRegistrarAdapter) PackClaim(account common.Address, resolver common.Address) []byte {
	enc, err := reverseRegistrarAdapter.abi.Pack("claim", account, resolver)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackClaim is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x21c0b342.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function claim(address account, address resolver) returns(bytes32)
func (reverseRegistrarAdapter *ReverseRegistrarAdapter) TryPackClaim(account common.Address, resolver common.Address) ([]byte, error) {
	return reverseRegistrarAdapter.abi.Pack("claim", account, resolver)
}

// UnpackClaim is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x21c0b342.
//
// Solidity: function claim(address account, address resolver) returns(bytes32)
func (reverseRegistrarAdapter *ReverseRegistrarAdapter) UnpackClaim(data []byte) ([32]byte, error) {
	out, err := reverseRegistrarAdapter.abi.Unpack("claim", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackClaimWithHCA is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc90695df.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function claimWithHCA(address account, address resolver) returns(bytes32)
func (reverseRegistrarAdapter *ReverseRegistrarAdapter) PackClaimWithHCA(account common.Address, resolver common.Address) []byte {
	enc, err := reverseRegistrarAdapter.abi.Pack("claimWithHCA", account, resolver)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackClaimWithHCA is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc90695df.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function claimWithHCA(address account, address resolver) returns(bytes32)
func (reverseRegistrarAdapter *ReverseRegistrarAdapter) TryPackClaimWithHCA(account common.Address, resolver common.Address) ([]byte, error) {
	return reverseRegistrarAdapter.abi.Pack("claimWithHCA", account, resolver)
}

// UnpackClaimWithHCA is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc90695df.
//
// Solidity: function claimWithHCA(address account, address resolver) returns(bytes32)
func (reverseRegistrarAdapter *ReverseRegistrarAdapter) UnpackClaimWithHCA(data []byte) ([32]byte, error) {
	out, err := reverseRegistrarAdapter.abi.Unpack("claimWithHCA", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackIsContractNamer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6f3ff726.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isContractNamer(address namer) view returns(bool)
func (reverseRegistrarAdapter *ReverseRegistrarAdapter) PackIsContractNamer(namer common.Address) []byte {
	enc, err := reverseRegistrarAdapter.abi.Pack("isContractNamer", namer)
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
func (reverseRegistrarAdapter *ReverseRegistrarAdapter) TryPackIsContractNamer(namer common.Address) ([]byte, error) {
	return reverseRegistrarAdapter.abi.Pack("isContractNamer", namer)
}

// UnpackIsContractNamer is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6f3ff726.
//
// Solidity: function isContractNamer(address namer) view returns(bool)
func (reverseRegistrarAdapter *ReverseRegistrarAdapter) UnpackIsContractNamer(data []byte) (bool, error) {
	out, err := reverseRegistrarAdapter.abi.Unpack("isContractNamer", data)
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
func (reverseRegistrarAdapter *ReverseRegistrarAdapter) PackSupportsInterface(interfaceId [4]byte) []byte {
	enc, err := reverseRegistrarAdapter.abi.Pack("supportsInterface", interfaceId)
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
func (reverseRegistrarAdapter *ReverseRegistrarAdapter) TryPackSupportsInterface(interfaceId [4]byte) ([]byte, error) {
	return reverseRegistrarAdapter.abi.Pack("supportsInterface", interfaceId)
}

// UnpackSupportsInterface is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (reverseRegistrarAdapter *ReverseRegistrarAdapter) UnpackSupportsInterface(data []byte) (bool, error) {
	out, err := reverseRegistrarAdapter.abi.Unpack("supportsInterface", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// UnpackError attempts to decode the provided error data using user-defined
// error definitions.
func (reverseRegistrarAdapter *ReverseRegistrarAdapter) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], reverseRegistrarAdapter.abi.Errors["HCADeploymentNotTrusted"].ID.Bytes()[:4]) {
		return reverseRegistrarAdapter.UnpackHCADeploymentNotTrustedError(raw[4:])
	}
	if bytes.Equal(raw[:4], reverseRegistrarAdapter.abi.Errors["HCAOwnerMismatch"].ID.Bytes()[:4]) {
		return reverseRegistrarAdapter.UnpackHCAOwnerMismatchError(raw[4:])
	}
	if bytes.Equal(raw[:4], reverseRegistrarAdapter.abi.Errors["UnauthorizedNamer"].ID.Bytes()[:4]) {
		return reverseRegistrarAdapter.UnpackUnauthorizedNamerError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// ReverseRegistrarAdapterHCADeploymentNotTrusted represents a HCADeploymentNotTrusted error raised by the ReverseRegistrarAdapter contract.
type ReverseRegistrarAdapterHCADeploymentNotTrusted struct {
	Hca common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error HCADeploymentNotTrusted(address hca)
func ReverseRegistrarAdapterHCADeploymentNotTrustedErrorID() common.Hash {
	return common.HexToHash("0x8d859048b6974ccb893e3552b02a721575409f6ee5f6e1a7fb625023605daa84")
}

// UnpackHCADeploymentNotTrustedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error HCADeploymentNotTrusted(address hca)
func (reverseRegistrarAdapter *ReverseRegistrarAdapter) UnpackHCADeploymentNotTrustedError(raw []byte) (*ReverseRegistrarAdapterHCADeploymentNotTrusted, error) {
	out := new(ReverseRegistrarAdapterHCADeploymentNotTrusted)
	if err := reverseRegistrarAdapter.abi.UnpackIntoInterface(out, "HCADeploymentNotTrusted", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ReverseRegistrarAdapterHCAOwnerMismatch represents a HCAOwnerMismatch error raised by the ReverseRegistrarAdapter contract.
type ReverseRegistrarAdapterHCAOwnerMismatch struct {
	Account  common.Address
	HcaOwner common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error HCAOwnerMismatch(address account, address hcaOwner)
func ReverseRegistrarAdapterHCAOwnerMismatchErrorID() common.Hash {
	return common.HexToHash("0xa8f2205f5ffac362d6c601309b57ff6751aaeb24ea70a3e47dddcdbabf63d8e7")
}

// UnpackHCAOwnerMismatchError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error HCAOwnerMismatch(address account, address hcaOwner)
func (reverseRegistrarAdapter *ReverseRegistrarAdapter) UnpackHCAOwnerMismatchError(raw []byte) (*ReverseRegistrarAdapterHCAOwnerMismatch, error) {
	out := new(ReverseRegistrarAdapterHCAOwnerMismatch)
	if err := reverseRegistrarAdapter.abi.UnpackIntoInterface(out, "HCAOwnerMismatch", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ReverseRegistrarAdapterUnauthorizedNamer represents a UnauthorizedNamer error raised by the ReverseRegistrarAdapter contract.
type ReverseRegistrarAdapterUnauthorizedNamer struct {
	Namer common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UnauthorizedNamer(address namer)
func ReverseRegistrarAdapterUnauthorizedNamerErrorID() common.Hash {
	return common.HexToHash("0x0d1b7e4e95d337c8403b427b1e83640e20d09f88944b7647f5d011301f680698")
}

// UnpackUnauthorizedNamerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UnauthorizedNamer(address namer)
func (reverseRegistrarAdapter *ReverseRegistrarAdapter) UnpackUnauthorizedNamerError(raw []byte) (*ReverseRegistrarAdapterUnauthorizedNamer, error) {
	out := new(ReverseRegistrarAdapterUnauthorizedNamer)
	if err := reverseRegistrarAdapter.abi.UnpackIntoInterface(out, "UnauthorizedNamer", raw); err != nil {
		return nil, err
	}
	return out, nil
}
