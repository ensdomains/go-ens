// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package publicresolver

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

// PublicResolverV2MetaData contains all meta data concerning the PublicResolverV2 contract.
var PublicResolverV2MetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractINameWrapper\",\"name\":\"nameWrapper\",\"type\":\"address\"},{\"internalType\":\"contractIPermissionedRegistry\",\"name\":\"rootRegistry\",\"type\":\"address\"},{\"internalType\":\"contractIContractNamer\",\"name\":\"contractNamer\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"dns\",\"type\":\"bytes\"}],\"name\":\"DNSDecodingFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"addressBytes\",\"type\":\"bytes\"}],\"name\":\"InvalidEVMAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"OffsetOutOfBoundsError\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"contentType\",\"type\":\"uint256\"}],\"name\":\"ABIChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"AddrChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"coinType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"newAddress\",\"type\":\"bytes\"}],\"name\":\"AddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"Approved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"hash\",\"type\":\"bytes\"}],\"name\":\"ContenthashChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"resource\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"record\",\"type\":\"bytes\"}],\"name\":\"DNSRecordChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"resource\",\"type\":\"uint16\"}],\"name\":\"DNSRecordDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"lastzonehash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"zonehash\",\"type\":\"bytes\"}],\"name\":\"DNSZonehashChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"indexedKey\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"indexedData\",\"type\":\"bytes\"}],\"name\":\"DataChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes4\",\"name\":\"interfaceID\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"implementer\",\"type\":\"address\"}],\"name\":\"InterfaceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"NameChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"x\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"y\",\"type\":\"bytes32\"}],\"name\":\"PubkeyChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"indexedKey\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"TextChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newVersion\",\"type\":\"uint64\"}],\"name\":\"VersionChanged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"contentTypes\",\"type\":\"uint256\"}],\"name\":\"ABI\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CONTRACT_NAMER\",\"outputs\":[{\"internalType\":\"contractIContractNamer\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NAME_WRAPPER\",\"outputs\":[{\"internalType\":\"contractINameWrapper\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ROOT_REGISTRY\",\"outputs\":[{\"internalType\":\"contractIPermissionedRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"addr\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"coinType\",\"type\":\"uint256\"}],\"name\":\"addr\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"addressBytes\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"canModifyName\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"clearRecords\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"contenthash\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"data\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"resource\",\"type\":\"uint16\"}],\"name\":\"dnsRecord\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"coinType\",\"type\":\"uint256\"}],\"name\":\"hasAddr\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"}],\"name\":\"hasDNSRecords\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"bytes4\",\"name\":\"interfaceID\",\"type\":\"bytes4\"}],\"name\":\"interfaceImplementer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"isApprovedFor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"namer\",\"type\":\"address\"}],\"name\":\"isContractNamer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nodehash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicallWithNodeCheck\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"pubkey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"x\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"y\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"recordVersions\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"contentType\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"setABI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"coinType\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"addressBytes\",\"type\":\"bytes\"}],\"name\":\"setAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"hash\",\"type\":\"bytes\"}],\"name\":\"setContenthash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"setDNSRecords\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"setData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"bytes4\",\"name\":\"interfaceID\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"implementer\",\"type\":\"address\"}],\"name\":\"setInterface\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"newName\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"x\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"y\",\"type\":\"bytes32\"}],\"name\":\"setPubkey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"setText\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"hash\",\"type\":\"bytes\"}],\"name\":\"setZonehash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"text\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"zonehash\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	ID:  "PublicResolverV2",
}

// PublicResolverV2 is an auto generated Go binding around an Ethereum contract.
type PublicResolverV2 struct {
	abi abi.ABI
}

// GetABI returns the ABI associated with this contract binding.
func (c *PublicResolverV2) GetABI() abi.ABI {
	return c.abi
}

// NewPublicResolverV2 creates a new instance of PublicResolverV2.
func NewPublicResolverV2() *PublicResolverV2 {
	parsed, err := PublicResolverV2MetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &PublicResolverV2{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *PublicResolverV2) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address nameWrapper, address rootRegistry, address contractNamer) returns()
func (publicResolverV2 *PublicResolverV2) PackConstructor(nameWrapper common.Address, rootRegistry common.Address, contractNamer common.Address) []byte {
	enc, err := publicResolverV2.abi.Pack("", nameWrapper, rootRegistry, contractNamer)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackABI is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2203ab56.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function ABI(bytes32 node, uint256 contentTypes) view returns(uint256, bytes)
func (publicResolverV2 *PublicResolverV2) PackABI(node [32]byte, contentTypes *big.Int) []byte {
	enc, err := publicResolverV2.abi.Pack("ABI", node, contentTypes)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackABI is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2203ab56.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function ABI(bytes32 node, uint256 contentTypes) view returns(uint256, bytes)
func (publicResolverV2 *PublicResolverV2) TryPackABI(node [32]byte, contentTypes *big.Int) ([]byte, error) {
	return publicResolverV2.abi.Pack("ABI", node, contentTypes)
}

// ABIOutput serves as a container for the return parameters of contract
// method ABI.
type ABIOutput struct {
	Arg0 *big.Int
	Arg1 []byte
}

// UnpackABI is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x2203ab56.
//
// Solidity: function ABI(bytes32 node, uint256 contentTypes) view returns(uint256, bytes)
func (publicResolverV2 *PublicResolverV2) UnpackABI(data []byte) (ABIOutput, error) {
	out, err := publicResolverV2.abi.Unpack("ABI", data)
	outstruct := new(ABIOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Arg0 = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.Arg1 = *abi.ConvertType(out[1], new([]byte)).(*[]byte)
	return *outstruct, nil
}

// PackCONTRACTNAMER is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x48ee1bcc.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function CONTRACT_NAMER() view returns(address)
func (publicResolverV2 *PublicResolverV2) PackCONTRACTNAMER() []byte {
	enc, err := publicResolverV2.abi.Pack("CONTRACT_NAMER")
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
func (publicResolverV2 *PublicResolverV2) TryPackCONTRACTNAMER() ([]byte, error) {
	return publicResolverV2.abi.Pack("CONTRACT_NAMER")
}

// UnpackCONTRACTNAMER is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x48ee1bcc.
//
// Solidity: function CONTRACT_NAMER() view returns(address)
func (publicResolverV2 *PublicResolverV2) UnpackCONTRACTNAMER(data []byte) (common.Address, error) {
	out, err := publicResolverV2.abi.Unpack("CONTRACT_NAMER", data)
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
func (publicResolverV2 *PublicResolverV2) PackNAMEWRAPPER() []byte {
	enc, err := publicResolverV2.abi.Pack("NAME_WRAPPER")
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
func (publicResolverV2 *PublicResolverV2) TryPackNAMEWRAPPER() ([]byte, error) {
	return publicResolverV2.abi.Pack("NAME_WRAPPER")
}

// UnpackNAMEWRAPPER is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x192cf07d.
//
// Solidity: function NAME_WRAPPER() view returns(address)
func (publicResolverV2 *PublicResolverV2) UnpackNAMEWRAPPER(data []byte) (common.Address, error) {
	out, err := publicResolverV2.abi.Unpack("NAME_WRAPPER", data)
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
func (publicResolverV2 *PublicResolverV2) PackROOTREGISTRY() []byte {
	enc, err := publicResolverV2.abi.Pack("ROOT_REGISTRY")
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
func (publicResolverV2 *PublicResolverV2) TryPackROOTREGISTRY() ([]byte, error) {
	return publicResolverV2.abi.Pack("ROOT_REGISTRY")
}

// UnpackROOTREGISTRY is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc92cc49a.
//
// Solidity: function ROOT_REGISTRY() view returns(address)
func (publicResolverV2 *PublicResolverV2) UnpackROOTREGISTRY(data []byte) (common.Address, error) {
	out, err := publicResolverV2.abi.Unpack("ROOT_REGISTRY", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackAddr is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3b3b57de.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function addr(bytes32 node) view returns(address)
func (publicResolverV2 *PublicResolverV2) PackAddr(node [32]byte) []byte {
	enc, err := publicResolverV2.abi.Pack("addr", node)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAddr is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3b3b57de.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function addr(bytes32 node) view returns(address)
func (publicResolverV2 *PublicResolverV2) TryPackAddr(node [32]byte) ([]byte, error) {
	return publicResolverV2.abi.Pack("addr", node)
}

// UnpackAddr is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3b3b57de.
//
// Solidity: function addr(bytes32 node) view returns(address)
func (publicResolverV2 *PublicResolverV2) UnpackAddr(data []byte) (common.Address, error) {
	out, err := publicResolverV2.abi.Unpack("addr", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackAddr0 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf1cb7e06.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function addr(bytes32 node, uint256 coinType) view returns(bytes addressBytes)
func (publicResolverV2 *PublicResolverV2) PackAddr0(node [32]byte, coinType *big.Int) []byte {
	enc, err := publicResolverV2.abi.Pack("addr0", node, coinType)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAddr0 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf1cb7e06.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function addr(bytes32 node, uint256 coinType) view returns(bytes addressBytes)
func (publicResolverV2 *PublicResolverV2) TryPackAddr0(node [32]byte, coinType *big.Int) ([]byte, error) {
	return publicResolverV2.abi.Pack("addr0", node, coinType)
}

// UnpackAddr0 is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf1cb7e06.
//
// Solidity: function addr(bytes32 node, uint256 coinType) view returns(bytes addressBytes)
func (publicResolverV2 *PublicResolverV2) UnpackAddr0(data []byte) ([]byte, error) {
	out, err := publicResolverV2.abi.Unpack("addr0", data)
	if err != nil {
		return *new([]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	return out0, nil
}

// PackApprove is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa4b91a01.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function approve(bytes32 node, address delegate, bool approved) returns()
func (publicResolverV2 *PublicResolverV2) PackApprove(node [32]byte, delegate common.Address, approved bool) []byte {
	enc, err := publicResolverV2.abi.Pack("approve", node, delegate, approved)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackApprove is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa4b91a01.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function approve(bytes32 node, address delegate, bool approved) returns()
func (publicResolverV2 *PublicResolverV2) TryPackApprove(node [32]byte, delegate common.Address, approved bool) ([]byte, error) {
	return publicResolverV2.abi.Pack("approve", node, delegate, approved)
}

// PackCanModifyName is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x41415eab.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function canModifyName(bytes32 node, address operator) view returns(bool)
func (publicResolverV2 *PublicResolverV2) PackCanModifyName(node [32]byte, operator common.Address) []byte {
	enc, err := publicResolverV2.abi.Pack("canModifyName", node, operator)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCanModifyName is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x41415eab.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function canModifyName(bytes32 node, address operator) view returns(bool)
func (publicResolverV2 *PublicResolverV2) TryPackCanModifyName(node [32]byte, operator common.Address) ([]byte, error) {
	return publicResolverV2.abi.Pack("canModifyName", node, operator)
}

// UnpackCanModifyName is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x41415eab.
//
// Solidity: function canModifyName(bytes32 node, address operator) view returns(bool)
func (publicResolverV2 *PublicResolverV2) UnpackCanModifyName(data []byte) (bool, error) {
	out, err := publicResolverV2.abi.Unpack("canModifyName", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackClearRecords is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3603d758.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function clearRecords(bytes32 node) returns()
func (publicResolverV2 *PublicResolverV2) PackClearRecords(node [32]byte) []byte {
	enc, err := publicResolverV2.abi.Pack("clearRecords", node)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackClearRecords is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3603d758.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function clearRecords(bytes32 node) returns()
func (publicResolverV2 *PublicResolverV2) TryPackClearRecords(node [32]byte) ([]byte, error) {
	return publicResolverV2.abi.Pack("clearRecords", node)
}

// PackContenthash is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbc1c58d1.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function contenthash(bytes32 node) view returns(bytes)
func (publicResolverV2 *PublicResolverV2) PackContenthash(node [32]byte) []byte {
	enc, err := publicResolverV2.abi.Pack("contenthash", node)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackContenthash is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbc1c58d1.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function contenthash(bytes32 node) view returns(bytes)
func (publicResolverV2 *PublicResolverV2) TryPackContenthash(node [32]byte) ([]byte, error) {
	return publicResolverV2.abi.Pack("contenthash", node)
}

// UnpackContenthash is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbc1c58d1.
//
// Solidity: function contenthash(bytes32 node) view returns(bytes)
func (publicResolverV2 *PublicResolverV2) UnpackContenthash(data []byte) ([]byte, error) {
	out, err := publicResolverV2.abi.Unpack("contenthash", data)
	if err != nil {
		return *new([]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	return out0, nil
}

// PackData is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xecbfada3.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function data(bytes32 node, string key) view returns(bytes)
func (publicResolverV2 *PublicResolverV2) PackData(node [32]byte, key string) []byte {
	enc, err := publicResolverV2.abi.Pack("data", node, key)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackData is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xecbfada3.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function data(bytes32 node, string key) view returns(bytes)
func (publicResolverV2 *PublicResolverV2) TryPackData(node [32]byte, key string) ([]byte, error) {
	return publicResolverV2.abi.Pack("data", node, key)
}

// UnpackData is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xecbfada3.
//
// Solidity: function data(bytes32 node, string key) view returns(bytes)
func (publicResolverV2 *PublicResolverV2) UnpackData(data []byte) ([]byte, error) {
	out, err := publicResolverV2.abi.Unpack("data", data)
	if err != nil {
		return *new([]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	return out0, nil
}

// PackDnsRecord is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa8fa5682.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function dnsRecord(bytes32 node, bytes32 name, uint16 resource) view returns(bytes)
func (publicResolverV2 *PublicResolverV2) PackDnsRecord(node [32]byte, name [32]byte, resource uint16) []byte {
	enc, err := publicResolverV2.abi.Pack("dnsRecord", node, name, resource)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDnsRecord is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa8fa5682.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function dnsRecord(bytes32 node, bytes32 name, uint16 resource) view returns(bytes)
func (publicResolverV2 *PublicResolverV2) TryPackDnsRecord(node [32]byte, name [32]byte, resource uint16) ([]byte, error) {
	return publicResolverV2.abi.Pack("dnsRecord", node, name, resource)
}

// UnpackDnsRecord is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa8fa5682.
//
// Solidity: function dnsRecord(bytes32 node, bytes32 name, uint16 resource) view returns(bytes)
func (publicResolverV2 *PublicResolverV2) UnpackDnsRecord(data []byte) ([]byte, error) {
	out, err := publicResolverV2.abi.Unpack("dnsRecord", data)
	if err != nil {
		return *new([]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	return out0, nil
}

// PackHasAddr is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x32f111d7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function hasAddr(bytes32 node, uint256 coinType) view returns(bool)
func (publicResolverV2 *PublicResolverV2) PackHasAddr(node [32]byte, coinType *big.Int) []byte {
	enc, err := publicResolverV2.abi.Pack("hasAddr", node, coinType)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackHasAddr is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x32f111d7.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function hasAddr(bytes32 node, uint256 coinType) view returns(bool)
func (publicResolverV2 *PublicResolverV2) TryPackHasAddr(node [32]byte, coinType *big.Int) ([]byte, error) {
	return publicResolverV2.abi.Pack("hasAddr", node, coinType)
}

// UnpackHasAddr is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x32f111d7.
//
// Solidity: function hasAddr(bytes32 node, uint256 coinType) view returns(bool)
func (publicResolverV2 *PublicResolverV2) UnpackHasAddr(data []byte) (bool, error) {
	out, err := publicResolverV2.abi.Unpack("hasAddr", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackHasDNSRecords is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4cbf6ba4.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function hasDNSRecords(bytes32 node, bytes32 name) view returns(bool)
func (publicResolverV2 *PublicResolverV2) PackHasDNSRecords(node [32]byte, name [32]byte) []byte {
	enc, err := publicResolverV2.abi.Pack("hasDNSRecords", node, name)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackHasDNSRecords is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4cbf6ba4.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function hasDNSRecords(bytes32 node, bytes32 name) view returns(bool)
func (publicResolverV2 *PublicResolverV2) TryPackHasDNSRecords(node [32]byte, name [32]byte) ([]byte, error) {
	return publicResolverV2.abi.Pack("hasDNSRecords", node, name)
}

// UnpackHasDNSRecords is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x4cbf6ba4.
//
// Solidity: function hasDNSRecords(bytes32 node, bytes32 name) view returns(bool)
func (publicResolverV2 *PublicResolverV2) UnpackHasDNSRecords(data []byte) (bool, error) {
	out, err := publicResolverV2.abi.Unpack("hasDNSRecords", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackInterfaceImplementer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x124a319c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function interfaceImplementer(bytes32 node, bytes4 interfaceID) view returns(address)
func (publicResolverV2 *PublicResolverV2) PackInterfaceImplementer(node [32]byte, interfaceID [4]byte) []byte {
	enc, err := publicResolverV2.abi.Pack("interfaceImplementer", node, interfaceID)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackInterfaceImplementer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x124a319c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function interfaceImplementer(bytes32 node, bytes4 interfaceID) view returns(address)
func (publicResolverV2 *PublicResolverV2) TryPackInterfaceImplementer(node [32]byte, interfaceID [4]byte) ([]byte, error) {
	return publicResolverV2.abi.Pack("interfaceImplementer", node, interfaceID)
}

// UnpackInterfaceImplementer is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x124a319c.
//
// Solidity: function interfaceImplementer(bytes32 node, bytes4 interfaceID) view returns(address)
func (publicResolverV2 *PublicResolverV2) UnpackInterfaceImplementer(data []byte) (common.Address, error) {
	out, err := publicResolverV2.abi.Unpack("interfaceImplementer", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackIsApprovedFor is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa9784b3e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isApprovedFor(address owner, bytes32 node, address delegate) view returns(bool)
func (publicResolverV2 *PublicResolverV2) PackIsApprovedFor(owner common.Address, node [32]byte, delegate common.Address) []byte {
	enc, err := publicResolverV2.abi.Pack("isApprovedFor", owner, node, delegate)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsApprovedFor is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa9784b3e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isApprovedFor(address owner, bytes32 node, address delegate) view returns(bool)
func (publicResolverV2 *PublicResolverV2) TryPackIsApprovedFor(owner common.Address, node [32]byte, delegate common.Address) ([]byte, error) {
	return publicResolverV2.abi.Pack("isApprovedFor", owner, node, delegate)
}

// UnpackIsApprovedFor is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa9784b3e.
//
// Solidity: function isApprovedFor(address owner, bytes32 node, address delegate) view returns(bool)
func (publicResolverV2 *PublicResolverV2) UnpackIsApprovedFor(data []byte) (bool, error) {
	out, err := publicResolverV2.abi.Unpack("isApprovedFor", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackIsApprovedForAll is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe985e9c5.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (publicResolverV2 *PublicResolverV2) PackIsApprovedForAll(owner common.Address, operator common.Address) []byte {
	enc, err := publicResolverV2.abi.Pack("isApprovedForAll", owner, operator)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsApprovedForAll is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe985e9c5.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (publicResolverV2 *PublicResolverV2) TryPackIsApprovedForAll(owner common.Address, operator common.Address) ([]byte, error) {
	return publicResolverV2.abi.Pack("isApprovedForAll", owner, operator)
}

// UnpackIsApprovedForAll is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (publicResolverV2 *PublicResolverV2) UnpackIsApprovedForAll(data []byte) (bool, error) {
	out, err := publicResolverV2.abi.Unpack("isApprovedForAll", data)
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
func (publicResolverV2 *PublicResolverV2) PackIsContractNamer(namer common.Address) []byte {
	enc, err := publicResolverV2.abi.Pack("isContractNamer", namer)
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
func (publicResolverV2 *PublicResolverV2) TryPackIsContractNamer(namer common.Address) ([]byte, error) {
	return publicResolverV2.abi.Pack("isContractNamer", namer)
}

// UnpackIsContractNamer is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6f3ff726.
//
// Solidity: function isContractNamer(address namer) view returns(bool)
func (publicResolverV2 *PublicResolverV2) UnpackIsContractNamer(data []byte) (bool, error) {
	out, err := publicResolverV2.abi.Unpack("isContractNamer", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackMulticall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xac9650d8.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (publicResolverV2 *PublicResolverV2) PackMulticall(data [][]byte) []byte {
	enc, err := publicResolverV2.abi.Pack("multicall", data)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackMulticall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xac9650d8.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (publicResolverV2 *PublicResolverV2) TryPackMulticall(data [][]byte) ([]byte, error) {
	return publicResolverV2.abi.Pack("multicall", data)
}

// UnpackMulticall is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (publicResolverV2 *PublicResolverV2) UnpackMulticall(data []byte) ([][]byte, error) {
	out, err := publicResolverV2.abi.Unpack("multicall", data)
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
// Solidity: function multicallWithNodeCheck(bytes32 nodehash, bytes[] data) returns(bytes[] results)
func (publicResolverV2 *PublicResolverV2) PackMulticallWithNodeCheck(nodehash [32]byte, data [][]byte) []byte {
	enc, err := publicResolverV2.abi.Pack("multicallWithNodeCheck", nodehash, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackMulticallWithNodeCheck is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe32954eb.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function multicallWithNodeCheck(bytes32 nodehash, bytes[] data) returns(bytes[] results)
func (publicResolverV2 *PublicResolverV2) TryPackMulticallWithNodeCheck(nodehash [32]byte, data [][]byte) ([]byte, error) {
	return publicResolverV2.abi.Pack("multicallWithNodeCheck", nodehash, data)
}

// UnpackMulticallWithNodeCheck is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe32954eb.
//
// Solidity: function multicallWithNodeCheck(bytes32 nodehash, bytes[] data) returns(bytes[] results)
func (publicResolverV2 *PublicResolverV2) UnpackMulticallWithNodeCheck(data []byte) ([][]byte, error) {
	out, err := publicResolverV2.abi.Unpack("multicallWithNodeCheck", data)
	if err != nil {
		return *new([][]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)
	return out0, nil
}

// PackName is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x691f3431.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function name(bytes32 node) view returns(string)
func (publicResolverV2 *PublicResolverV2) PackName(node [32]byte) []byte {
	enc, err := publicResolverV2.abi.Pack("name", node)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackName is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x691f3431.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function name(bytes32 node) view returns(string)
func (publicResolverV2 *PublicResolverV2) TryPackName(node [32]byte) ([]byte, error) {
	return publicResolverV2.abi.Pack("name", node)
}

// UnpackName is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x691f3431.
//
// Solidity: function name(bytes32 node) view returns(string)
func (publicResolverV2 *PublicResolverV2) UnpackName(data []byte) (string, error) {
	out, err := publicResolverV2.abi.Unpack("name", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, nil
}

// PackPubkey is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc8690233.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function pubkey(bytes32 node) view returns(bytes32 x, bytes32 y)
func (publicResolverV2 *PublicResolverV2) PackPubkey(node [32]byte) []byte {
	enc, err := publicResolverV2.abi.Pack("pubkey", node)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackPubkey is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc8690233.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function pubkey(bytes32 node) view returns(bytes32 x, bytes32 y)
func (publicResolverV2 *PublicResolverV2) TryPackPubkey(node [32]byte) ([]byte, error) {
	return publicResolverV2.abi.Pack("pubkey", node)
}

// PubkeyOutput serves as a container for the return parameters of contract
// method Pubkey.
type PubkeyOutput struct {
	X [32]byte
	Y [32]byte
}

// UnpackPubkey is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc8690233.
//
// Solidity: function pubkey(bytes32 node) view returns(bytes32 x, bytes32 y)
func (publicResolverV2 *PublicResolverV2) UnpackPubkey(data []byte) (PubkeyOutput, error) {
	out, err := publicResolverV2.abi.Unpack("pubkey", data)
	outstruct := new(PubkeyOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.X = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Y = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	return *outstruct, nil
}

// PackRecordVersions is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd700ff33.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function recordVersions(bytes32 ) view returns(uint64)
func (publicResolverV2 *PublicResolverV2) PackRecordVersions(arg0 [32]byte) []byte {
	enc, err := publicResolverV2.abi.Pack("recordVersions", arg0)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRecordVersions is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd700ff33.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function recordVersions(bytes32 ) view returns(uint64)
func (publicResolverV2 *PublicResolverV2) TryPackRecordVersions(arg0 [32]byte) ([]byte, error) {
	return publicResolverV2.abi.Pack("recordVersions", arg0)
}

// UnpackRecordVersions is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd700ff33.
//
// Solidity: function recordVersions(bytes32 ) view returns(uint64)
func (publicResolverV2 *PublicResolverV2) UnpackRecordVersions(data []byte) (uint64, error) {
	out, err := publicResolverV2.abi.Unpack("recordVersions", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, nil
}

// PackSetABI is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x623195b0.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setABI(bytes32 node, uint256 contentType, bytes data) returns()
func (publicResolverV2 *PublicResolverV2) PackSetABI(node [32]byte, contentType *big.Int, data []byte) []byte {
	enc, err := publicResolverV2.abi.Pack("setABI", node, contentType, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetABI is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x623195b0.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setABI(bytes32 node, uint256 contentType, bytes data) returns()
func (publicResolverV2 *PublicResolverV2) TryPackSetABI(node [32]byte, contentType *big.Int, data []byte) ([]byte, error) {
	return publicResolverV2.abi.Pack("setABI", node, contentType, data)
}

// PackSetAddr is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8b95dd71.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setAddr(bytes32 node, uint256 coinType, bytes addressBytes) returns()
func (publicResolverV2 *PublicResolverV2) PackSetAddr(node [32]byte, coinType *big.Int, addressBytes []byte) []byte {
	enc, err := publicResolverV2.abi.Pack("setAddr", node, coinType, addressBytes)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetAddr is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8b95dd71.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setAddr(bytes32 node, uint256 coinType, bytes addressBytes) returns()
func (publicResolverV2 *PublicResolverV2) TryPackSetAddr(node [32]byte, coinType *big.Int, addressBytes []byte) ([]byte, error) {
	return publicResolverV2.abi.Pack("setAddr", node, coinType, addressBytes)
}

// PackSetAddr0 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd5fa2b00.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setAddr(bytes32 node, address _addr) returns()
func (publicResolverV2 *PublicResolverV2) PackSetAddr0(node [32]byte, addr common.Address) []byte {
	enc, err := publicResolverV2.abi.Pack("setAddr0", node, addr)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetAddr0 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd5fa2b00.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setAddr(bytes32 node, address _addr) returns()
func (publicResolverV2 *PublicResolverV2) TryPackSetAddr0(node [32]byte, addr common.Address) ([]byte, error) {
	return publicResolverV2.abi.Pack("setAddr0", node, addr)
}

// PackSetApprovalForAll is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa22cb465.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (publicResolverV2 *PublicResolverV2) PackSetApprovalForAll(operator common.Address, approved bool) []byte {
	enc, err := publicResolverV2.abi.Pack("setApprovalForAll", operator, approved)
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
func (publicResolverV2 *PublicResolverV2) TryPackSetApprovalForAll(operator common.Address, approved bool) ([]byte, error) {
	return publicResolverV2.abi.Pack("setApprovalForAll", operator, approved)
}

// PackSetContenthash is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x304e6ade.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setContenthash(bytes32 node, bytes hash) returns()
func (publicResolverV2 *PublicResolverV2) PackSetContenthash(node [32]byte, hash []byte) []byte {
	enc, err := publicResolverV2.abi.Pack("setContenthash", node, hash)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetContenthash is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x304e6ade.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setContenthash(bytes32 node, bytes hash) returns()
func (publicResolverV2 *PublicResolverV2) TryPackSetContenthash(node [32]byte, hash []byte) ([]byte, error) {
	return publicResolverV2.abi.Pack("setContenthash", node, hash)
}

// PackSetDNSRecords is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0af179d7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setDNSRecords(bytes32 node, bytes data) returns()
func (publicResolverV2 *PublicResolverV2) PackSetDNSRecords(node [32]byte, data []byte) []byte {
	enc, err := publicResolverV2.abi.Pack("setDNSRecords", node, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetDNSRecords is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0af179d7.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setDNSRecords(bytes32 node, bytes data) returns()
func (publicResolverV2 *PublicResolverV2) TryPackSetDNSRecords(node [32]byte, data []byte) ([]byte, error) {
	return publicResolverV2.abi.Pack("setDNSRecords", node, data)
}

// PackSetData is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4eb9c45e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setData(bytes32 node, string key, bytes value) returns()
func (publicResolverV2 *PublicResolverV2) PackSetData(node [32]byte, key string, value []byte) []byte {
	enc, err := publicResolverV2.abi.Pack("setData", node, key, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetData is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4eb9c45e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setData(bytes32 node, string key, bytes value) returns()
func (publicResolverV2 *PublicResolverV2) TryPackSetData(node [32]byte, key string, value []byte) ([]byte, error) {
	return publicResolverV2.abi.Pack("setData", node, key, value)
}

// PackSetInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe59d895d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setInterface(bytes32 node, bytes4 interfaceID, address implementer) returns()
func (publicResolverV2 *PublicResolverV2) PackSetInterface(node [32]byte, interfaceID [4]byte, implementer common.Address) []byte {
	enc, err := publicResolverV2.abi.Pack("setInterface", node, interfaceID, implementer)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe59d895d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setInterface(bytes32 node, bytes4 interfaceID, address implementer) returns()
func (publicResolverV2 *PublicResolverV2) TryPackSetInterface(node [32]byte, interfaceID [4]byte, implementer common.Address) ([]byte, error) {
	return publicResolverV2.abi.Pack("setInterface", node, interfaceID, implementer)
}

// PackSetName is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x77372213.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setName(bytes32 node, string newName) returns()
func (publicResolverV2 *PublicResolverV2) PackSetName(node [32]byte, newName string) []byte {
	enc, err := publicResolverV2.abi.Pack("setName", node, newName)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetName is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x77372213.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setName(bytes32 node, string newName) returns()
func (publicResolverV2 *PublicResolverV2) TryPackSetName(node [32]byte, newName string) ([]byte, error) {
	return publicResolverV2.abi.Pack("setName", node, newName)
}

// PackSetPubkey is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x29cd62ea.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setPubkey(bytes32 node, bytes32 x, bytes32 y) returns()
func (publicResolverV2 *PublicResolverV2) PackSetPubkey(node [32]byte, x [32]byte, y [32]byte) []byte {
	enc, err := publicResolverV2.abi.Pack("setPubkey", node, x, y)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetPubkey is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x29cd62ea.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setPubkey(bytes32 node, bytes32 x, bytes32 y) returns()
func (publicResolverV2 *PublicResolverV2) TryPackSetPubkey(node [32]byte, x [32]byte, y [32]byte) ([]byte, error) {
	return publicResolverV2.abi.Pack("setPubkey", node, x, y)
}

// PackSetText is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x10f13a8c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setText(bytes32 node, string key, string value) returns()
func (publicResolverV2 *PublicResolverV2) PackSetText(node [32]byte, key string, value string) []byte {
	enc, err := publicResolverV2.abi.Pack("setText", node, key, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetText is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x10f13a8c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setText(bytes32 node, string key, string value) returns()
func (publicResolverV2 *PublicResolverV2) TryPackSetText(node [32]byte, key string, value string) ([]byte, error) {
	return publicResolverV2.abi.Pack("setText", node, key, value)
}

// PackSetZonehash is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xce3decdc.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setZonehash(bytes32 node, bytes hash) returns()
func (publicResolverV2 *PublicResolverV2) PackSetZonehash(node [32]byte, hash []byte) []byte {
	enc, err := publicResolverV2.abi.Pack("setZonehash", node, hash)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetZonehash is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xce3decdc.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setZonehash(bytes32 node, bytes hash) returns()
func (publicResolverV2 *PublicResolverV2) TryPackSetZonehash(node [32]byte, hash []byte) ([]byte, error) {
	return publicResolverV2.abi.Pack("setZonehash", node, hash)
}

// PackSupportsInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01ffc9a7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (publicResolverV2 *PublicResolverV2) PackSupportsInterface(interfaceId [4]byte) []byte {
	enc, err := publicResolverV2.abi.Pack("supportsInterface", interfaceId)
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
func (publicResolverV2 *PublicResolverV2) TryPackSupportsInterface(interfaceId [4]byte) ([]byte, error) {
	return publicResolverV2.abi.Pack("supportsInterface", interfaceId)
}

// UnpackSupportsInterface is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (publicResolverV2 *PublicResolverV2) UnpackSupportsInterface(data []byte) (bool, error) {
	out, err := publicResolverV2.abi.Unpack("supportsInterface", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackText is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x59d1d43c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function text(bytes32 node, string key) view returns(string)
func (publicResolverV2 *PublicResolverV2) PackText(node [32]byte, key string) []byte {
	enc, err := publicResolverV2.abi.Pack("text", node, key)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackText is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x59d1d43c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function text(bytes32 node, string key) view returns(string)
func (publicResolverV2 *PublicResolverV2) TryPackText(node [32]byte, key string) ([]byte, error) {
	return publicResolverV2.abi.Pack("text", node, key)
}

// UnpackText is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x59d1d43c.
//
// Solidity: function text(bytes32 node, string key) view returns(string)
func (publicResolverV2 *PublicResolverV2) UnpackText(data []byte) (string, error) {
	out, err := publicResolverV2.abi.Unpack("text", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, nil
}

// PackZonehash is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5c98042b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function zonehash(bytes32 node) view returns(bytes)
func (publicResolverV2 *PublicResolverV2) PackZonehash(node [32]byte) []byte {
	enc, err := publicResolverV2.abi.Pack("zonehash", node)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackZonehash is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5c98042b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function zonehash(bytes32 node) view returns(bytes)
func (publicResolverV2 *PublicResolverV2) TryPackZonehash(node [32]byte) ([]byte, error) {
	return publicResolverV2.abi.Pack("zonehash", node)
}

// UnpackZonehash is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5c98042b.
//
// Solidity: function zonehash(bytes32 node) view returns(bytes)
func (publicResolverV2 *PublicResolverV2) UnpackZonehash(data []byte) ([]byte, error) {
	out, err := publicResolverV2.abi.Unpack("zonehash", data)
	if err != nil {
		return *new([]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	return out0, nil
}

// PublicResolverV2ABIChanged represents a ABIChanged event raised by the PublicResolverV2 contract.
type PublicResolverV2ABIChanged struct {
	Node        [32]byte
	ContentType *big.Int
	Raw         *types.Log // Blockchain specific contextual infos
}

const PublicResolverV2ABIChangedEventName = "ABIChanged"

// ContractEventName returns the user-defined event name.
func (PublicResolverV2ABIChanged) ContractEventName() string {
	return PublicResolverV2ABIChangedEventName
}

// UnpackABIChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ABIChanged(bytes32 indexed node, uint256 indexed contentType)
func (publicResolverV2 *PublicResolverV2) UnpackABIChangedEvent(log *types.Log) (*PublicResolverV2ABIChanged, error) {
	event := "ABIChanged"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != publicResolverV2.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(PublicResolverV2ABIChanged)
	if len(log.Data) > 0 {
		if err := publicResolverV2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range publicResolverV2.abi.Events[event].Inputs {
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

// PublicResolverV2AddrChanged represents a AddrChanged event raised by the PublicResolverV2 contract.
type PublicResolverV2AddrChanged struct {
	Node [32]byte
	A    common.Address
	Raw  *types.Log // Blockchain specific contextual infos
}

const PublicResolverV2AddrChangedEventName = "AddrChanged"

// ContractEventName returns the user-defined event name.
func (PublicResolverV2AddrChanged) ContractEventName() string {
	return PublicResolverV2AddrChangedEventName
}

// UnpackAddrChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event AddrChanged(bytes32 indexed node, address a)
func (publicResolverV2 *PublicResolverV2) UnpackAddrChangedEvent(log *types.Log) (*PublicResolverV2AddrChanged, error) {
	event := "AddrChanged"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != publicResolverV2.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(PublicResolverV2AddrChanged)
	if len(log.Data) > 0 {
		if err := publicResolverV2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range publicResolverV2.abi.Events[event].Inputs {
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

// PublicResolverV2AddressChanged represents a AddressChanged event raised by the PublicResolverV2 contract.
type PublicResolverV2AddressChanged struct {
	Node       [32]byte
	CoinType   *big.Int
	NewAddress []byte
	Raw        *types.Log // Blockchain specific contextual infos
}

const PublicResolverV2AddressChangedEventName = "AddressChanged"

// ContractEventName returns the user-defined event name.
func (PublicResolverV2AddressChanged) ContractEventName() string {
	return PublicResolverV2AddressChangedEventName
}

// UnpackAddressChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event AddressChanged(bytes32 indexed node, uint256 coinType, bytes newAddress)
func (publicResolverV2 *PublicResolverV2) UnpackAddressChangedEvent(log *types.Log) (*PublicResolverV2AddressChanged, error) {
	event := "AddressChanged"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != publicResolverV2.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(PublicResolverV2AddressChanged)
	if len(log.Data) > 0 {
		if err := publicResolverV2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range publicResolverV2.abi.Events[event].Inputs {
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

// PublicResolverV2ApprovalForAll represents a ApprovalForAll event raised by the PublicResolverV2 contract.
type PublicResolverV2ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      *types.Log // Blockchain specific contextual infos
}

const PublicResolverV2ApprovalForAllEventName = "ApprovalForAll"

// ContractEventName returns the user-defined event name.
func (PublicResolverV2ApprovalForAll) ContractEventName() string {
	return PublicResolverV2ApprovalForAllEventName
}

// UnpackApprovalForAllEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (publicResolverV2 *PublicResolverV2) UnpackApprovalForAllEvent(log *types.Log) (*PublicResolverV2ApprovalForAll, error) {
	event := "ApprovalForAll"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != publicResolverV2.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(PublicResolverV2ApprovalForAll)
	if len(log.Data) > 0 {
		if err := publicResolverV2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range publicResolverV2.abi.Events[event].Inputs {
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

// PublicResolverV2Approved represents a Approved event raised by the PublicResolverV2 contract.
type PublicResolverV2Approved struct {
	Owner    common.Address
	Node     [32]byte
	Delegate common.Address
	Approved bool
	Raw      *types.Log // Blockchain specific contextual infos
}

const PublicResolverV2ApprovedEventName = "Approved"

// ContractEventName returns the user-defined event name.
func (PublicResolverV2Approved) ContractEventName() string {
	return PublicResolverV2ApprovedEventName
}

// UnpackApprovedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Approved(address owner, bytes32 indexed node, address indexed delegate, bool indexed approved)
func (publicResolverV2 *PublicResolverV2) UnpackApprovedEvent(log *types.Log) (*PublicResolverV2Approved, error) {
	event := "Approved"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != publicResolverV2.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(PublicResolverV2Approved)
	if len(log.Data) > 0 {
		if err := publicResolverV2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range publicResolverV2.abi.Events[event].Inputs {
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

// PublicResolverV2ContenthashChanged represents a ContenthashChanged event raised by the PublicResolverV2 contract.
type PublicResolverV2ContenthashChanged struct {
	Node [32]byte
	Hash []byte
	Raw  *types.Log // Blockchain specific contextual infos
}

const PublicResolverV2ContenthashChangedEventName = "ContenthashChanged"

// ContractEventName returns the user-defined event name.
func (PublicResolverV2ContenthashChanged) ContractEventName() string {
	return PublicResolverV2ContenthashChangedEventName
}

// UnpackContenthashChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ContenthashChanged(bytes32 indexed node, bytes hash)
func (publicResolverV2 *PublicResolverV2) UnpackContenthashChangedEvent(log *types.Log) (*PublicResolverV2ContenthashChanged, error) {
	event := "ContenthashChanged"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != publicResolverV2.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(PublicResolverV2ContenthashChanged)
	if len(log.Data) > 0 {
		if err := publicResolverV2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range publicResolverV2.abi.Events[event].Inputs {
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

// PublicResolverV2DNSRecordChanged represents a DNSRecordChanged event raised by the PublicResolverV2 contract.
type PublicResolverV2DNSRecordChanged struct {
	Node     [32]byte
	Name     []byte
	Resource uint16
	Record   []byte
	Raw      *types.Log // Blockchain specific contextual infos
}

const PublicResolverV2DNSRecordChangedEventName = "DNSRecordChanged"

// ContractEventName returns the user-defined event name.
func (PublicResolverV2DNSRecordChanged) ContractEventName() string {
	return PublicResolverV2DNSRecordChangedEventName
}

// UnpackDNSRecordChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DNSRecordChanged(bytes32 indexed node, bytes name, uint16 resource, bytes record)
func (publicResolverV2 *PublicResolverV2) UnpackDNSRecordChangedEvent(log *types.Log) (*PublicResolverV2DNSRecordChanged, error) {
	event := "DNSRecordChanged"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != publicResolverV2.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(PublicResolverV2DNSRecordChanged)
	if len(log.Data) > 0 {
		if err := publicResolverV2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range publicResolverV2.abi.Events[event].Inputs {
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

// PublicResolverV2DNSRecordDeleted represents a DNSRecordDeleted event raised by the PublicResolverV2 contract.
type PublicResolverV2DNSRecordDeleted struct {
	Node     [32]byte
	Name     []byte
	Resource uint16
	Raw      *types.Log // Blockchain specific contextual infos
}

const PublicResolverV2DNSRecordDeletedEventName = "DNSRecordDeleted"

// ContractEventName returns the user-defined event name.
func (PublicResolverV2DNSRecordDeleted) ContractEventName() string {
	return PublicResolverV2DNSRecordDeletedEventName
}

// UnpackDNSRecordDeletedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DNSRecordDeleted(bytes32 indexed node, bytes name, uint16 resource)
func (publicResolverV2 *PublicResolverV2) UnpackDNSRecordDeletedEvent(log *types.Log) (*PublicResolverV2DNSRecordDeleted, error) {
	event := "DNSRecordDeleted"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != publicResolverV2.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(PublicResolverV2DNSRecordDeleted)
	if len(log.Data) > 0 {
		if err := publicResolverV2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range publicResolverV2.abi.Events[event].Inputs {
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

// PublicResolverV2DNSZonehashChanged represents a DNSZonehashChanged event raised by the PublicResolverV2 contract.
type PublicResolverV2DNSZonehashChanged struct {
	Node         [32]byte
	Lastzonehash []byte
	Zonehash     []byte
	Raw          *types.Log // Blockchain specific contextual infos
}

const PublicResolverV2DNSZonehashChangedEventName = "DNSZonehashChanged"

// ContractEventName returns the user-defined event name.
func (PublicResolverV2DNSZonehashChanged) ContractEventName() string {
	return PublicResolverV2DNSZonehashChangedEventName
}

// UnpackDNSZonehashChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DNSZonehashChanged(bytes32 indexed node, bytes lastzonehash, bytes zonehash)
func (publicResolverV2 *PublicResolverV2) UnpackDNSZonehashChangedEvent(log *types.Log) (*PublicResolverV2DNSZonehashChanged, error) {
	event := "DNSZonehashChanged"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != publicResolverV2.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(PublicResolverV2DNSZonehashChanged)
	if len(log.Data) > 0 {
		if err := publicResolverV2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range publicResolverV2.abi.Events[event].Inputs {
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

// PublicResolverV2DataChanged represents a DataChanged event raised by the PublicResolverV2 contract.
type PublicResolverV2DataChanged struct {
	Node        [32]byte
	IndexedKey  common.Hash
	Key         string
	IndexedData common.Hash
	Raw         *types.Log // Blockchain specific contextual infos
}

const PublicResolverV2DataChangedEventName = "DataChanged"

// ContractEventName returns the user-defined event name.
func (PublicResolverV2DataChanged) ContractEventName() string {
	return PublicResolverV2DataChangedEventName
}

// UnpackDataChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DataChanged(bytes32 indexed node, string indexed indexedKey, string key, bytes indexed indexedData)
func (publicResolverV2 *PublicResolverV2) UnpackDataChangedEvent(log *types.Log) (*PublicResolverV2DataChanged, error) {
	event := "DataChanged"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != publicResolverV2.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(PublicResolverV2DataChanged)
	if len(log.Data) > 0 {
		if err := publicResolverV2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range publicResolverV2.abi.Events[event].Inputs {
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

// PublicResolverV2InterfaceChanged represents a InterfaceChanged event raised by the PublicResolverV2 contract.
type PublicResolverV2InterfaceChanged struct {
	Node        [32]byte
	InterfaceID [4]byte
	Implementer common.Address
	Raw         *types.Log // Blockchain specific contextual infos
}

const PublicResolverV2InterfaceChangedEventName = "InterfaceChanged"

// ContractEventName returns the user-defined event name.
func (PublicResolverV2InterfaceChanged) ContractEventName() string {
	return PublicResolverV2InterfaceChangedEventName
}

// UnpackInterfaceChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event InterfaceChanged(bytes32 indexed node, bytes4 indexed interfaceID, address implementer)
func (publicResolverV2 *PublicResolverV2) UnpackInterfaceChangedEvent(log *types.Log) (*PublicResolverV2InterfaceChanged, error) {
	event := "InterfaceChanged"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != publicResolverV2.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(PublicResolverV2InterfaceChanged)
	if len(log.Data) > 0 {
		if err := publicResolverV2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range publicResolverV2.abi.Events[event].Inputs {
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

// PublicResolverV2NameChanged represents a NameChanged event raised by the PublicResolverV2 contract.
type PublicResolverV2NameChanged struct {
	Node [32]byte
	Name string
	Raw  *types.Log // Blockchain specific contextual infos
}

const PublicResolverV2NameChangedEventName = "NameChanged"

// ContractEventName returns the user-defined event name.
func (PublicResolverV2NameChanged) ContractEventName() string {
	return PublicResolverV2NameChangedEventName
}

// UnpackNameChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event NameChanged(bytes32 indexed node, string name)
func (publicResolverV2 *PublicResolverV2) UnpackNameChangedEvent(log *types.Log) (*PublicResolverV2NameChanged, error) {
	event := "NameChanged"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != publicResolverV2.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(PublicResolverV2NameChanged)
	if len(log.Data) > 0 {
		if err := publicResolverV2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range publicResolverV2.abi.Events[event].Inputs {
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

// PublicResolverV2PubkeyChanged represents a PubkeyChanged event raised by the PublicResolverV2 contract.
type PublicResolverV2PubkeyChanged struct {
	Node [32]byte
	X    [32]byte
	Y    [32]byte
	Raw  *types.Log // Blockchain specific contextual infos
}

const PublicResolverV2PubkeyChangedEventName = "PubkeyChanged"

// ContractEventName returns the user-defined event name.
func (PublicResolverV2PubkeyChanged) ContractEventName() string {
	return PublicResolverV2PubkeyChangedEventName
}

// UnpackPubkeyChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event PubkeyChanged(bytes32 indexed node, bytes32 x, bytes32 y)
func (publicResolverV2 *PublicResolverV2) UnpackPubkeyChangedEvent(log *types.Log) (*PublicResolverV2PubkeyChanged, error) {
	event := "PubkeyChanged"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != publicResolverV2.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(PublicResolverV2PubkeyChanged)
	if len(log.Data) > 0 {
		if err := publicResolverV2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range publicResolverV2.abi.Events[event].Inputs {
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

// PublicResolverV2TextChanged represents a TextChanged event raised by the PublicResolverV2 contract.
type PublicResolverV2TextChanged struct {
	Node       [32]byte
	IndexedKey common.Hash
	Key        string
	Value      string
	Raw        *types.Log // Blockchain specific contextual infos
}

const PublicResolverV2TextChangedEventName = "TextChanged"

// ContractEventName returns the user-defined event name.
func (PublicResolverV2TextChanged) ContractEventName() string {
	return PublicResolverV2TextChangedEventName
}

// UnpackTextChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TextChanged(bytes32 indexed node, string indexed indexedKey, string key, string value)
func (publicResolverV2 *PublicResolverV2) UnpackTextChangedEvent(log *types.Log) (*PublicResolverV2TextChanged, error) {
	event := "TextChanged"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != publicResolverV2.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(PublicResolverV2TextChanged)
	if len(log.Data) > 0 {
		if err := publicResolverV2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range publicResolverV2.abi.Events[event].Inputs {
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

// PublicResolverV2VersionChanged represents a VersionChanged event raised by the PublicResolverV2 contract.
type PublicResolverV2VersionChanged struct {
	Node       [32]byte
	NewVersion uint64
	Raw        *types.Log // Blockchain specific contextual infos
}

const PublicResolverV2VersionChangedEventName = "VersionChanged"

// ContractEventName returns the user-defined event name.
func (PublicResolverV2VersionChanged) ContractEventName() string {
	return PublicResolverV2VersionChangedEventName
}

// UnpackVersionChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event VersionChanged(bytes32 indexed node, uint64 newVersion)
func (publicResolverV2 *PublicResolverV2) UnpackVersionChangedEvent(log *types.Log) (*PublicResolverV2VersionChanged, error) {
	event := "VersionChanged"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != publicResolverV2.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(PublicResolverV2VersionChanged)
	if len(log.Data) > 0 {
		if err := publicResolverV2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range publicResolverV2.abi.Events[event].Inputs {
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
func (publicResolverV2 *PublicResolverV2) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], publicResolverV2.abi.Errors["DNSDecodingFailed"].ID.Bytes()[:4]) {
		return publicResolverV2.UnpackDNSDecodingFailedError(raw[4:])
	}
	if bytes.Equal(raw[:4], publicResolverV2.abi.Errors["InvalidEVMAddress"].ID.Bytes()[:4]) {
		return publicResolverV2.UnpackInvalidEVMAddressError(raw[4:])
	}
	if bytes.Equal(raw[:4], publicResolverV2.abi.Errors["OffsetOutOfBoundsError"].ID.Bytes()[:4]) {
		return publicResolverV2.UnpackOffsetOutOfBoundsErrorError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// PublicResolverV2DNSDecodingFailed represents a DNSDecodingFailed error raised by the PublicResolverV2 contract.
type PublicResolverV2DNSDecodingFailed struct {
	Dns []byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error DNSDecodingFailed(bytes dns)
func PublicResolverV2DNSDecodingFailedErrorID() common.Hash {
	return common.HexToHash("0xba4adc2398f6571d4d3236eee815a96fcbbf84bbc4320e35b771331f71afdc72")
}

// UnpackDNSDecodingFailedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error DNSDecodingFailed(bytes dns)
func (publicResolverV2 *PublicResolverV2) UnpackDNSDecodingFailedError(raw []byte) (*PublicResolverV2DNSDecodingFailed, error) {
	out := new(PublicResolverV2DNSDecodingFailed)
	if err := publicResolverV2.abi.UnpackIntoInterface(out, "DNSDecodingFailed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PublicResolverV2InvalidEVMAddress represents a InvalidEVMAddress error raised by the PublicResolverV2 contract.
type PublicResolverV2InvalidEVMAddress struct {
	AddressBytes []byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidEVMAddress(bytes addressBytes)
func PublicResolverV2InvalidEVMAddressErrorID() common.Hash {
	return common.HexToHash("0x8d666f60151e99b4b75dcde8fac04e7b1bfe46b40ae9842f449dbad1817d17b8")
}

// UnpackInvalidEVMAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidEVMAddress(bytes addressBytes)
func (publicResolverV2 *PublicResolverV2) UnpackInvalidEVMAddressError(raw []byte) (*PublicResolverV2InvalidEVMAddress, error) {
	out := new(PublicResolverV2InvalidEVMAddress)
	if err := publicResolverV2.abi.UnpackIntoInterface(out, "InvalidEVMAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PublicResolverV2OffsetOutOfBoundsError represents a OffsetOutOfBoundsError error raised by the PublicResolverV2 contract.
type PublicResolverV2OffsetOutOfBoundsError struct {
	Offset *big.Int
	Length *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OffsetOutOfBoundsError(uint256 offset, uint256 length)
func PublicResolverV2OffsetOutOfBoundsErrorErrorID() common.Hash {
	return common.HexToHash("0x8a3c1cfb98bd88fadc1ef4bc86d0b221827194fb788772725d6bd414678ba0eb")
}

// UnpackOffsetOutOfBoundsErrorError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OffsetOutOfBoundsError(uint256 offset, uint256 length)
func (publicResolverV2 *PublicResolverV2) UnpackOffsetOutOfBoundsErrorError(raw []byte) (*PublicResolverV2OffsetOutOfBoundsError, error) {
	out := new(PublicResolverV2OffsetOutOfBoundsError)
	if err := publicResolverV2.abi.UnpackIntoInterface(out, "OffsetOutOfBoundsError", raw); err != nil {
		return nil, err
	}
	return out, nil
}
