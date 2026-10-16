// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package universalresolver

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

// CCIPBatcherBatch is an auto generated low-level Go binding around an user-defined struct.
type CCIPBatcherBatch struct {
	Lookups  []CCIPBatcherLookup
	Gateways []string
}

// CCIPBatcherLookup is an auto generated low-level Go binding around an user-defined struct.
type CCIPBatcherLookup struct {
	Target common.Address
	Call   []byte
	Data   []byte
	Flags  *big.Int
}

// IUniversalResolverExtendedResolverInfo is an auto generated low-level Go binding around an user-defined struct.
type IUniversalResolverExtendedResolverInfo struct {
	Name     []byte
	Offset   *big.Int
	Node     [32]byte
	Resolver common.Address
	Extended bool
}

// UniversalResolverV2MetaData contains all meta data concerning the UniversalResolverV2 contract.
var UniversalResolverV2MetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPermissionedRegistry\",\"name\":\"rootRegistry\",\"type\":\"address\"},{\"internalType\":\"contractIGatewayProvider\",\"name\":\"batchGatewayProvider\",\"type\":\"address\"},{\"internalType\":\"contractIContractNamer\",\"name\":\"contractNamer\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"dns\",\"type\":\"bytes\"}],\"name\":\"DNSDecodingFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"ens\",\"type\":\"string\"}],\"name\":\"DNSEncodingFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"status\",\"type\":\"uint16\"},{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"HttpError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBatchGatewayResponse\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LabelIsEmpty\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"}],\"name\":\"LabelIsTooLong\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"normalizedName\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"}],\"name\":\"NormalizationChangedName\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"urls\",\"type\":\"string[]\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes4\",\"name\":\"callbackFunction\",\"type\":\"bytes4\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"OffchainLookup\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"OffsetOutOfBoundsError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"primary\",\"type\":\"string\"}],\"name\":\"PrimaryNameNotNormalized\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"errorData\",\"type\":\"bytes\"}],\"name\":\"ResolverError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"}],\"name\":\"ResolverNotContract\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"}],\"name\":\"ResolverNotFound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"primary\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"primaryAddress\",\"type\":\"bytes\"}],\"name\":\"ReverseAddressMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"UnsafeBatchGatewayResponse\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"UnsupportedResolverProfile\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BATCH_GATEWAY_PROVIDER\",\"outputs\":[{\"internalType\":\"contractIGatewayProvider\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CONTRACT_NAMER\",\"outputs\":[{\"internalType\":\"contractIContractNamer\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ROOT_REGISTRY\",\"outputs\":[{\"internalType\":\"contractIPermissionedRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"batchGatewayProvider\",\"outputs\":[{\"internalType\":\"contractIGatewayProvider\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"call\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"}],\"internalType\":\"structCCIPBatcher.Lookup[]\",\"name\":\"lookups\",\"type\":\"tuple[]\"},{\"internalType\":\"string[]\",\"name\":\"gateways\",\"type\":\"string[]\"}],\"internalType\":\"structCCIPBatcher.Batch\",\"name\":\"batch\",\"type\":\"tuple\"}],\"name\":\"ccipBatch\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"call\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"}],\"internalType\":\"structCCIPBatcher.Lookup[]\",\"name\":\"lookups\",\"type\":\"tuple[]\"},{\"internalType\":\"string[]\",\"name\":\"gateways\",\"type\":\"string[]\"}],\"internalType\":\"structCCIPBatcher.Batch\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"response\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"ccipBatchCallback\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"call\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"}],\"internalType\":\"structCCIPBatcher.Lookup[]\",\"name\":\"lookups\",\"type\":\"tuple[]\"},{\"internalType\":\"string[]\",\"name\":\"gateways\",\"type\":\"string[]\"}],\"internalType\":\"structCCIPBatcher.Batch\",\"name\":\"batch\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"response\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"ccipReadCallback\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"}],\"name\":\"findResolver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"namer\",\"type\":\"address\"}],\"name\":\"isContractNamer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isENSv2\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"internalType\":\"contractIENSIP15\",\"name\":\"ensip15\",\"type\":\"address\"}],\"name\":\"normalize\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"}],\"name\":\"requireResolver\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"extended\",\"type\":\"bool\"}],\"internalType\":\"structIUniversalResolverExtended.ResolverInfo\",\"name\":\"info\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"resolve\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"response\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"resolveBatchCallback\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"response\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"resolveCallback\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"response\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"resolveDirectCallback\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"response\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"resolveDirectCallbackError\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"string[]\",\"name\":\"gateways\",\"type\":\"string[]\"}],\"name\":\"resolveWithGateways\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"string[]\",\"name\":\"gateways\",\"type\":\"string[]\"},{\"internalType\":\"contractIENSIP15\",\"name\":\"ensip15\",\"type\":\"address\"}],\"name\":\"resolveWithGatewaysAndNormalization\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"contractIENSIP15\",\"name\":\"ensip15\",\"type\":\"address\"}],\"name\":\"resolveWithNormalization\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"string[]\",\"name\":\"gateways\",\"type\":\"string[]\"}],\"name\":\"resolveWithResolver\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"lookupAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"coinType\",\"type\":\"uint256\"}],\"name\":\"reverse\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"response\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"reverseAddressCallback\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"primary\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"reverseResolver\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"response\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"reverseNameCallback\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"primary\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"lookupAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"coinType\",\"type\":\"uint256\"},{\"internalType\":\"string[]\",\"name\":\"gateways\",\"type\":\"string[]\"}],\"name\":\"reverseWithGateways\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"lookupAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"coinType\",\"type\":\"uint256\"},{\"internalType\":\"string[]\",\"name\":\"gateways\",\"type\":\"string[]\"},{\"internalType\":\"contractIENSIP15\",\"name\":\"ensip15\",\"type\":\"address\"}],\"name\":\"reverseWithGatewaysAndNormalization\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"lookupAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"coinType\",\"type\":\"uint256\"},{\"internalType\":\"contractIENSIP15\",\"name\":\"ensip15\",\"type\":\"address\"}],\"name\":\"reverseWithNormalization\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	ID:  "UniversalResolverV2",
}

// UniversalResolverV2 is an auto generated Go binding around an Ethereum contract.
type UniversalResolverV2 struct {
	abi abi.ABI
}

// GetABI returns the ABI associated with this contract binding.
func (c *UniversalResolverV2) GetABI() abi.ABI {
	return c.abi
}

// NewUniversalResolverV2 creates a new instance of UniversalResolverV2.
func NewUniversalResolverV2() *UniversalResolverV2 {
	parsed, err := UniversalResolverV2MetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &UniversalResolverV2{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *UniversalResolverV2) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address rootRegistry, address batchGatewayProvider, address contractNamer) returns()
func (universalResolverV2 *UniversalResolverV2) PackConstructor(rootRegistry common.Address, batchGatewayProvider common.Address, contractNamer common.Address) []byte {
	enc, err := universalResolverV2.abi.Pack("", rootRegistry, batchGatewayProvider, contractNamer)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackBATCHGATEWAYPROVIDER is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6ccb8660.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function BATCH_GATEWAY_PROVIDER() view returns(address)
func (universalResolverV2 *UniversalResolverV2) PackBATCHGATEWAYPROVIDER() []byte {
	enc, err := universalResolverV2.abi.Pack("BATCH_GATEWAY_PROVIDER")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackBATCHGATEWAYPROVIDER is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6ccb8660.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function BATCH_GATEWAY_PROVIDER() view returns(address)
func (universalResolverV2 *UniversalResolverV2) TryPackBATCHGATEWAYPROVIDER() ([]byte, error) {
	return universalResolverV2.abi.Pack("BATCH_GATEWAY_PROVIDER")
}

// UnpackBATCHGATEWAYPROVIDER is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6ccb8660.
//
// Solidity: function BATCH_GATEWAY_PROVIDER() view returns(address)
func (universalResolverV2 *UniversalResolverV2) UnpackBATCHGATEWAYPROVIDER(data []byte) (common.Address, error) {
	out, err := universalResolverV2.abi.Unpack("BATCH_GATEWAY_PROVIDER", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackCONTRACTNAMER is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x48ee1bcc.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function CONTRACT_NAMER() view returns(address)
func (universalResolverV2 *UniversalResolverV2) PackCONTRACTNAMER() []byte {
	enc, err := universalResolverV2.abi.Pack("CONTRACT_NAMER")
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
func (universalResolverV2 *UniversalResolverV2) TryPackCONTRACTNAMER() ([]byte, error) {
	return universalResolverV2.abi.Pack("CONTRACT_NAMER")
}

// UnpackCONTRACTNAMER is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x48ee1bcc.
//
// Solidity: function CONTRACT_NAMER() view returns(address)
func (universalResolverV2 *UniversalResolverV2) UnpackCONTRACTNAMER(data []byte) (common.Address, error) {
	out, err := universalResolverV2.abi.Unpack("CONTRACT_NAMER", data)
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
func (universalResolverV2 *UniversalResolverV2) PackROOTREGISTRY() []byte {
	enc, err := universalResolverV2.abi.Pack("ROOT_REGISTRY")
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
func (universalResolverV2 *UniversalResolverV2) TryPackROOTREGISTRY() ([]byte, error) {
	return universalResolverV2.abi.Pack("ROOT_REGISTRY")
}

// UnpackROOTREGISTRY is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc92cc49a.
//
// Solidity: function ROOT_REGISTRY() view returns(address)
func (universalResolverV2 *UniversalResolverV2) UnpackROOTREGISTRY(data []byte) (common.Address, error) {
	out, err := universalResolverV2.abi.Unpack("ROOT_REGISTRY", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackBatchGatewayProvider is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x02cf2578.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function batchGatewayProvider() view returns(address)
func (universalResolverV2 *UniversalResolverV2) PackBatchGatewayProvider() []byte {
	enc, err := universalResolverV2.abi.Pack("batchGatewayProvider")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackBatchGatewayProvider is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x02cf2578.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function batchGatewayProvider() view returns(address)
func (universalResolverV2 *UniversalResolverV2) TryPackBatchGatewayProvider() ([]byte, error) {
	return universalResolverV2.abi.Pack("batchGatewayProvider")
}

// UnpackBatchGatewayProvider is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x02cf2578.
//
// Solidity: function batchGatewayProvider() view returns(address)
func (universalResolverV2 *UniversalResolverV2) UnpackBatchGatewayProvider(data []byte) (common.Address, error) {
	out, err := universalResolverV2.abi.Unpack("batchGatewayProvider", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackCcipBatch is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9f28e99d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function ccipBatch(((address,bytes,bytes,uint256)[],string[]) batch) view returns(((address,bytes,bytes,uint256)[],string[]))
func (universalResolverV2 *UniversalResolverV2) PackCcipBatch(batch CCIPBatcherBatch) []byte {
	enc, err := universalResolverV2.abi.Pack("ccipBatch", batch)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCcipBatch is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9f28e99d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function ccipBatch(((address,bytes,bytes,uint256)[],string[]) batch) view returns(((address,bytes,bytes,uint256)[],string[]))
func (universalResolverV2 *UniversalResolverV2) TryPackCcipBatch(batch CCIPBatcherBatch) ([]byte, error) {
	return universalResolverV2.abi.Pack("ccipBatch", batch)
}

// UnpackCcipBatch is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9f28e99d.
//
// Solidity: function ccipBatch(((address,bytes,bytes,uint256)[],string[]) batch) view returns(((address,bytes,bytes,uint256)[],string[]))
func (universalResolverV2 *UniversalResolverV2) UnpackCcipBatch(data []byte) (CCIPBatcherBatch, error) {
	out, err := universalResolverV2.abi.Unpack("ccipBatch", data)
	if err != nil {
		return *new(CCIPBatcherBatch), err
	}
	out0 := *abi.ConvertType(out[0], new(CCIPBatcherBatch)).(*CCIPBatcherBatch)
	return out0, nil
}

// PackCcipBatchCallback is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb536af76.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function ccipBatchCallback(bytes response, bytes extraData) view returns(((address,bytes,bytes,uint256)[],string[]) batch)
func (universalResolverV2 *UniversalResolverV2) PackCcipBatchCallback(response []byte, extraData []byte) []byte {
	enc, err := universalResolverV2.abi.Pack("ccipBatchCallback", response, extraData)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCcipBatchCallback is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb536af76.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function ccipBatchCallback(bytes response, bytes extraData) view returns(((address,bytes,bytes,uint256)[],string[]) batch)
func (universalResolverV2 *UniversalResolverV2) TryPackCcipBatchCallback(response []byte, extraData []byte) ([]byte, error) {
	return universalResolverV2.abi.Pack("ccipBatchCallback", response, extraData)
}

// UnpackCcipBatchCallback is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb536af76.
//
// Solidity: function ccipBatchCallback(bytes response, bytes extraData) view returns(((address,bytes,bytes,uint256)[],string[]) batch)
func (universalResolverV2 *UniversalResolverV2) UnpackCcipBatchCallback(data []byte) (CCIPBatcherBatch, error) {
	out, err := universalResolverV2.abi.Unpack("ccipBatchCallback", data)
	if err != nil {
		return *new(CCIPBatcherBatch), err
	}
	out0 := *abi.ConvertType(out[0], new(CCIPBatcherBatch)).(*CCIPBatcherBatch)
	return out0, nil
}

// PackCcipReadCallback is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xef46c0b8.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function ccipReadCallback(bytes response, bytes extraData) view returns()
func (universalResolverV2 *UniversalResolverV2) PackCcipReadCallback(response []byte, extraData []byte) []byte {
	enc, err := universalResolverV2.abi.Pack("ccipReadCallback", response, extraData)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCcipReadCallback is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xef46c0b8.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function ccipReadCallback(bytes response, bytes extraData) view returns()
func (universalResolverV2 *UniversalResolverV2) TryPackCcipReadCallback(response []byte, extraData []byte) ([]byte, error) {
	return universalResolverV2.abi.Pack("ccipReadCallback", response, extraData)
}

// PackFindResolver is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa1cbcbaf.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function findResolver(bytes name) view returns(address resolver, bytes32 node, uint256 offset)
func (universalResolverV2 *UniversalResolverV2) PackFindResolver(name []byte) []byte {
	enc, err := universalResolverV2.abi.Pack("findResolver", name)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFindResolver is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa1cbcbaf.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function findResolver(bytes name) view returns(address resolver, bytes32 node, uint256 offset)
func (universalResolverV2 *UniversalResolverV2) TryPackFindResolver(name []byte) ([]byte, error) {
	return universalResolverV2.abi.Pack("findResolver", name)
}

// FindResolverOutput serves as a container for the return parameters of contract
// method FindResolver.
type FindResolverOutput struct {
	Resolver common.Address
	Node     [32]byte
	Offset   *big.Int
}

// UnpackFindResolver is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa1cbcbaf.
//
// Solidity: function findResolver(bytes name) view returns(address resolver, bytes32 node, uint256 offset)
func (universalResolverV2 *UniversalResolverV2) UnpackFindResolver(data []byte) (FindResolverOutput, error) {
	out, err := universalResolverV2.abi.Unpack("findResolver", data)
	outstruct := new(FindResolverOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Resolver = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Node = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.Offset = abi.ConvertType(out[2], new(big.Int)).(*big.Int)
	return *outstruct, nil
}

// PackIsContractNamer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6f3ff726.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isContractNamer(address namer) view returns(bool)
func (universalResolverV2 *UniversalResolverV2) PackIsContractNamer(namer common.Address) []byte {
	enc, err := universalResolverV2.abi.Pack("isContractNamer", namer)
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
func (universalResolverV2 *UniversalResolverV2) TryPackIsContractNamer(namer common.Address) ([]byte, error) {
	return universalResolverV2.abi.Pack("isContractNamer", namer)
}

// UnpackIsContractNamer is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6f3ff726.
//
// Solidity: function isContractNamer(address namer) view returns(bool)
func (universalResolverV2 *UniversalResolverV2) UnpackIsContractNamer(data []byte) (bool, error) {
	out, err := universalResolverV2.abi.Unpack("isContractNamer", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackIsENSv2 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1a6cc9f0.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isENSv2() pure returns(bool)
func (universalResolverV2 *UniversalResolverV2) PackIsENSv2() []byte {
	enc, err := universalResolverV2.abi.Pack("isENSv2")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsENSv2 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1a6cc9f0.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isENSv2() pure returns(bool)
func (universalResolverV2 *UniversalResolverV2) TryPackIsENSv2() ([]byte, error) {
	return universalResolverV2.abi.Pack("isENSv2")
}

// UnpackIsENSv2 is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x1a6cc9f0.
//
// Solidity: function isENSv2() pure returns(bool)
func (universalResolverV2 *UniversalResolverV2) UnpackIsENSv2(data []byte) (bool, error) {
	out, err := universalResolverV2.abi.Unpack("isENSv2", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackNormalize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x649a7ed8.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function normalize(bytes name, address ensip15) view returns(bool, bytes)
func (universalResolverV2 *UniversalResolverV2) PackNormalize(name []byte, ensip15 common.Address) []byte {
	enc, err := universalResolverV2.abi.Pack("normalize", name, ensip15)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackNormalize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x649a7ed8.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function normalize(bytes name, address ensip15) view returns(bool, bytes)
func (universalResolverV2 *UniversalResolverV2) TryPackNormalize(name []byte, ensip15 common.Address) ([]byte, error) {
	return universalResolverV2.abi.Pack("normalize", name, ensip15)
}

// NormalizeOutput serves as a container for the return parameters of contract
// method Normalize.
type NormalizeOutput struct {
	Arg0 bool
	Arg1 []byte
}

// UnpackNormalize is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x649a7ed8.
//
// Solidity: function normalize(bytes name, address ensip15) view returns(bool, bytes)
func (universalResolverV2 *UniversalResolverV2) UnpackNormalize(data []byte) (NormalizeOutput, error) {
	out, err := universalResolverV2.abi.Unpack("normalize", data)
	outstruct := new(NormalizeOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Arg0 = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Arg1 = *abi.ConvertType(out[1], new([]byte)).(*[]byte)
	return *outstruct, nil
}

// PackRequireResolver is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc285238a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function requireResolver(bytes name) view returns((bytes,uint256,bytes32,address,bool) info)
func (universalResolverV2 *UniversalResolverV2) PackRequireResolver(name []byte) []byte {
	enc, err := universalResolverV2.abi.Pack("requireResolver", name)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRequireResolver is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc285238a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function requireResolver(bytes name) view returns((bytes,uint256,bytes32,address,bool) info)
func (universalResolverV2 *UniversalResolverV2) TryPackRequireResolver(name []byte) ([]byte, error) {
	return universalResolverV2.abi.Pack("requireResolver", name)
}

// UnpackRequireResolver is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc285238a.
//
// Solidity: function requireResolver(bytes name) view returns((bytes,uint256,bytes32,address,bool) info)
func (universalResolverV2 *UniversalResolverV2) UnpackRequireResolver(data []byte) (IUniversalResolverExtendedResolverInfo, error) {
	out, err := universalResolverV2.abi.Unpack("requireResolver", data)
	if err != nil {
		return *new(IUniversalResolverExtendedResolverInfo), err
	}
	out0 := *abi.ConvertType(out[0], new(IUniversalResolverExtendedResolverInfo)).(*IUniversalResolverExtendedResolverInfo)
	return out0, nil
}

// PackResolve is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9061b923.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function resolve(bytes name, bytes data) view returns(bytes, address)
func (universalResolverV2 *UniversalResolverV2) PackResolve(name []byte, data []byte) []byte {
	enc, err := universalResolverV2.abi.Pack("resolve", name, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackResolve is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9061b923.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function resolve(bytes name, bytes data) view returns(bytes, address)
func (universalResolverV2 *UniversalResolverV2) TryPackResolve(name []byte, data []byte) ([]byte, error) {
	return universalResolverV2.abi.Pack("resolve", name, data)
}

// ResolveOutput serves as a container for the return parameters of contract
// method Resolve.
type ResolveOutput struct {
	Arg0 []byte
	Arg1 common.Address
}

// UnpackResolve is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9061b923.
//
// Solidity: function resolve(bytes name, bytes data) view returns(bytes, address)
func (universalResolverV2 *UniversalResolverV2) UnpackResolve(data []byte) (ResolveOutput, error) {
	out, err := universalResolverV2.abi.Unpack("resolve", data)
	outstruct := new(ResolveOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Arg0 = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.Arg1 = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	return *outstruct, nil
}

// PackResolveBatchCallback is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x491fc4f9.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function resolveBatchCallback(bytes response, bytes extraData) view returns()
func (universalResolverV2 *UniversalResolverV2) PackResolveBatchCallback(response []byte, extraData []byte) []byte {
	enc, err := universalResolverV2.abi.Pack("resolveBatchCallback", response, extraData)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackResolveBatchCallback is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x491fc4f9.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function resolveBatchCallback(bytes response, bytes extraData) view returns()
func (universalResolverV2 *UniversalResolverV2) TryPackResolveBatchCallback(response []byte, extraData []byte) ([]byte, error) {
	return universalResolverV2.abi.Pack("resolveBatchCallback", response, extraData)
}

// PackResolveCallback is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb4a85801.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function resolveCallback(bytes response, bytes extraData) pure returns(bytes, address)
func (universalResolverV2 *UniversalResolverV2) PackResolveCallback(response []byte, extraData []byte) []byte {
	enc, err := universalResolverV2.abi.Pack("resolveCallback", response, extraData)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackResolveCallback is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb4a85801.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function resolveCallback(bytes response, bytes extraData) pure returns(bytes, address)
func (universalResolverV2 *UniversalResolverV2) TryPackResolveCallback(response []byte, extraData []byte) ([]byte, error) {
	return universalResolverV2.abi.Pack("resolveCallback", response, extraData)
}

// ResolveCallbackOutput serves as a container for the return parameters of contract
// method ResolveCallback.
type ResolveCallbackOutput struct {
	Arg0 []byte
	Arg1 common.Address
}

// UnpackResolveCallback is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb4a85801.
//
// Solidity: function resolveCallback(bytes response, bytes extraData) pure returns(bytes, address)
func (universalResolverV2 *UniversalResolverV2) UnpackResolveCallback(data []byte) (ResolveCallbackOutput, error) {
	out, err := universalResolverV2.abi.Unpack("resolveCallback", data)
	outstruct := new(ResolveCallbackOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Arg0 = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.Arg1 = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	return *outstruct, nil
}

// PackResolveDirectCallback is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x55391bb8.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function resolveDirectCallback(bytes response, bytes extraData) view returns()
func (universalResolverV2 *UniversalResolverV2) PackResolveDirectCallback(response []byte, extraData []byte) []byte {
	enc, err := universalResolverV2.abi.Pack("resolveDirectCallback", response, extraData)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackResolveDirectCallback is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x55391bb8.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function resolveDirectCallback(bytes response, bytes extraData) view returns()
func (universalResolverV2 *UniversalResolverV2) TryPackResolveDirectCallback(response []byte, extraData []byte) ([]byte, error) {
	return universalResolverV2.abi.Pack("resolveDirectCallback", response, extraData)
}

// PackResolveDirectCallbackError is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3c6cbda8.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function resolveDirectCallbackError(bytes response, bytes ) pure returns()
func (universalResolverV2 *UniversalResolverV2) PackResolveDirectCallbackError(response []byte, arg1 []byte) []byte {
	enc, err := universalResolverV2.abi.Pack("resolveDirectCallbackError", response, arg1)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackResolveDirectCallbackError is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3c6cbda8.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function resolveDirectCallbackError(bytes response, bytes ) pure returns()
func (universalResolverV2 *UniversalResolverV2) TryPackResolveDirectCallbackError(response []byte, arg1 []byte) ([]byte, error) {
	return universalResolverV2.abi.Pack("resolveDirectCallbackError", response, arg1)
}

// PackResolveWithGateways is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa1472844.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function resolveWithGateways(bytes name, bytes data, string[] gateways) view returns(bytes, address)
func (universalResolverV2 *UniversalResolverV2) PackResolveWithGateways(name []byte, data []byte, gateways []string) []byte {
	enc, err := universalResolverV2.abi.Pack("resolveWithGateways", name, data, gateways)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackResolveWithGateways is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa1472844.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function resolveWithGateways(bytes name, bytes data, string[] gateways) view returns(bytes, address)
func (universalResolverV2 *UniversalResolverV2) TryPackResolveWithGateways(name []byte, data []byte, gateways []string) ([]byte, error) {
	return universalResolverV2.abi.Pack("resolveWithGateways", name, data, gateways)
}

// ResolveWithGatewaysOutput serves as a container for the return parameters of contract
// method ResolveWithGateways.
type ResolveWithGatewaysOutput struct {
	Arg0 []byte
	Arg1 common.Address
}

// UnpackResolveWithGateways is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa1472844.
//
// Solidity: function resolveWithGateways(bytes name, bytes data, string[] gateways) view returns(bytes, address)
func (universalResolverV2 *UniversalResolverV2) UnpackResolveWithGateways(data []byte) (ResolveWithGatewaysOutput, error) {
	out, err := universalResolverV2.abi.Unpack("resolveWithGateways", data)
	outstruct := new(ResolveWithGatewaysOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Arg0 = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.Arg1 = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	return *outstruct, nil
}

// PackResolveWithGatewaysAndNormalization is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6166ddf2.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function resolveWithGatewaysAndNormalization(bytes name, bytes data, string[] gateways, address ensip15) view returns(bytes, address)
func (universalResolverV2 *UniversalResolverV2) PackResolveWithGatewaysAndNormalization(name []byte, data []byte, gateways []string, ensip15 common.Address) []byte {
	enc, err := universalResolverV2.abi.Pack("resolveWithGatewaysAndNormalization", name, data, gateways, ensip15)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackResolveWithGatewaysAndNormalization is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6166ddf2.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function resolveWithGatewaysAndNormalization(bytes name, bytes data, string[] gateways, address ensip15) view returns(bytes, address)
func (universalResolverV2 *UniversalResolverV2) TryPackResolveWithGatewaysAndNormalization(name []byte, data []byte, gateways []string, ensip15 common.Address) ([]byte, error) {
	return universalResolverV2.abi.Pack("resolveWithGatewaysAndNormalization", name, data, gateways, ensip15)
}

// ResolveWithGatewaysAndNormalizationOutput serves as a container for the return parameters of contract
// method ResolveWithGatewaysAndNormalization.
type ResolveWithGatewaysAndNormalizationOutput struct {
	Arg0 []byte
	Arg1 common.Address
}

// UnpackResolveWithGatewaysAndNormalization is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6166ddf2.
//
// Solidity: function resolveWithGatewaysAndNormalization(bytes name, bytes data, string[] gateways, address ensip15) view returns(bytes, address)
func (universalResolverV2 *UniversalResolverV2) UnpackResolveWithGatewaysAndNormalization(data []byte) (ResolveWithGatewaysAndNormalizationOutput, error) {
	out, err := universalResolverV2.abi.Unpack("resolveWithGatewaysAndNormalization", data)
	outstruct := new(ResolveWithGatewaysAndNormalizationOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Arg0 = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.Arg1 = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	return *outstruct, nil
}

// PackResolveWithNormalization is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc5d8f5dd.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function resolveWithNormalization(bytes name, bytes data, address ensip15) view returns(bytes, address)
func (universalResolverV2 *UniversalResolverV2) PackResolveWithNormalization(name []byte, data []byte, ensip15 common.Address) []byte {
	enc, err := universalResolverV2.abi.Pack("resolveWithNormalization", name, data, ensip15)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackResolveWithNormalization is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc5d8f5dd.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function resolveWithNormalization(bytes name, bytes data, address ensip15) view returns(bytes, address)
func (universalResolverV2 *UniversalResolverV2) TryPackResolveWithNormalization(name []byte, data []byte, ensip15 common.Address) ([]byte, error) {
	return universalResolverV2.abi.Pack("resolveWithNormalization", name, data, ensip15)
}

// ResolveWithNormalizationOutput serves as a container for the return parameters of contract
// method ResolveWithNormalization.
type ResolveWithNormalizationOutput struct {
	Arg0 []byte
	Arg1 common.Address
}

// UnpackResolveWithNormalization is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc5d8f5dd.
//
// Solidity: function resolveWithNormalization(bytes name, bytes data, address ensip15) view returns(bytes, address)
func (universalResolverV2 *UniversalResolverV2) UnpackResolveWithNormalization(data []byte) (ResolveWithNormalizationOutput, error) {
	out, err := universalResolverV2.abi.Unpack("resolveWithNormalization", data)
	outstruct := new(ResolveWithNormalizationOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Arg0 = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.Arg1 = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	return *outstruct, nil
}

// PackResolveWithResolver is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4a3e3994.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function resolveWithResolver(address resolver, bytes name, bytes data, string[] gateways) view returns(bytes)
func (universalResolverV2 *UniversalResolverV2) PackResolveWithResolver(resolver common.Address, name []byte, data []byte, gateways []string) []byte {
	enc, err := universalResolverV2.abi.Pack("resolveWithResolver", resolver, name, data, gateways)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackResolveWithResolver is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4a3e3994.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function resolveWithResolver(address resolver, bytes name, bytes data, string[] gateways) view returns(bytes)
func (universalResolverV2 *UniversalResolverV2) TryPackResolveWithResolver(resolver common.Address, name []byte, data []byte, gateways []string) ([]byte, error) {
	return universalResolverV2.abi.Pack("resolveWithResolver", resolver, name, data, gateways)
}

// UnpackResolveWithResolver is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x4a3e3994.
//
// Solidity: function resolveWithResolver(address resolver, bytes name, bytes data, string[] gateways) view returns(bytes)
func (universalResolverV2 *UniversalResolverV2) UnpackResolveWithResolver(data []byte) ([]byte, error) {
	out, err := universalResolverV2.abi.Unpack("resolveWithResolver", data)
	if err != nil {
		return *new([]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	return out0, nil
}

// PackReverse is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5d78a217.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function reverse(bytes lookupAddress, uint256 coinType) view returns(string, address, address)
func (universalResolverV2 *UniversalResolverV2) PackReverse(lookupAddress []byte, coinType *big.Int) []byte {
	enc, err := universalResolverV2.abi.Pack("reverse", lookupAddress, coinType)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackReverse is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5d78a217.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function reverse(bytes lookupAddress, uint256 coinType) view returns(string, address, address)
func (universalResolverV2 *UniversalResolverV2) TryPackReverse(lookupAddress []byte, coinType *big.Int) ([]byte, error) {
	return universalResolverV2.abi.Pack("reverse", lookupAddress, coinType)
}

// ReverseOutput serves as a container for the return parameters of contract
// method Reverse.
type ReverseOutput struct {
	Arg0 string
	Arg1 common.Address
	Arg2 common.Address
}

// UnpackReverse is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5d78a217.
//
// Solidity: function reverse(bytes lookupAddress, uint256 coinType) view returns(string, address, address)
func (universalResolverV2 *UniversalResolverV2) UnpackReverse(data []byte) (ReverseOutput, error) {
	out, err := universalResolverV2.abi.Unpack("reverse", data)
	outstruct := new(ReverseOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Arg0 = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Arg1 = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Arg2 = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	return *outstruct, nil
}

// PackReverseAddressCallback is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x94fbfa87.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function reverseAddressCallback(bytes response, bytes extraData) pure returns(string primary, address resolver, address reverseResolver)
func (universalResolverV2 *UniversalResolverV2) PackReverseAddressCallback(response []byte, extraData []byte) []byte {
	enc, err := universalResolverV2.abi.Pack("reverseAddressCallback", response, extraData)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackReverseAddressCallback is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x94fbfa87.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function reverseAddressCallback(bytes response, bytes extraData) pure returns(string primary, address resolver, address reverseResolver)
func (universalResolverV2 *UniversalResolverV2) TryPackReverseAddressCallback(response []byte, extraData []byte) ([]byte, error) {
	return universalResolverV2.abi.Pack("reverseAddressCallback", response, extraData)
}

// ReverseAddressCallbackOutput serves as a container for the return parameters of contract
// method ReverseAddressCallback.
type ReverseAddressCallbackOutput struct {
	Primary         string
	Resolver        common.Address
	ReverseResolver common.Address
}

// UnpackReverseAddressCallback is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x94fbfa87.
//
// Solidity: function reverseAddressCallback(bytes response, bytes extraData) pure returns(string primary, address resolver, address reverseResolver)
func (universalResolverV2 *UniversalResolverV2) UnpackReverseAddressCallback(data []byte) (ReverseAddressCallbackOutput, error) {
	out, err := universalResolverV2.abi.Unpack("reverseAddressCallback", data)
	outstruct := new(ReverseAddressCallbackOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Primary = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Resolver = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.ReverseResolver = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	return *outstruct, nil
}

// PackReverseNameCallback is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x575de750.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function reverseNameCallback(bytes response, bytes extraData) view returns(string primary, address, address)
func (universalResolverV2 *UniversalResolverV2) PackReverseNameCallback(response []byte, extraData []byte) []byte {
	enc, err := universalResolverV2.abi.Pack("reverseNameCallback", response, extraData)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackReverseNameCallback is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x575de750.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function reverseNameCallback(bytes response, bytes extraData) view returns(string primary, address, address)
func (universalResolverV2 *UniversalResolverV2) TryPackReverseNameCallback(response []byte, extraData []byte) ([]byte, error) {
	return universalResolverV2.abi.Pack("reverseNameCallback", response, extraData)
}

// ReverseNameCallbackOutput serves as a container for the return parameters of contract
// method ReverseNameCallback.
type ReverseNameCallbackOutput struct {
	Primary string
	Arg1    common.Address
	Arg2    common.Address
}

// UnpackReverseNameCallback is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x575de750.
//
// Solidity: function reverseNameCallback(bytes response, bytes extraData) view returns(string primary, address, address)
func (universalResolverV2 *UniversalResolverV2) UnpackReverseNameCallback(data []byte) (ReverseNameCallbackOutput, error) {
	out, err := universalResolverV2.abi.Unpack("reverseNameCallback", data)
	outstruct := new(ReverseNameCallbackOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Primary = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Arg1 = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Arg2 = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	return *outstruct, nil
}

// PackReverseWithGateways is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb7d6ca64.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function reverseWithGateways(bytes lookupAddress, uint256 coinType, string[] gateways) view returns(string, address, address)
func (universalResolverV2 *UniversalResolverV2) PackReverseWithGateways(lookupAddress []byte, coinType *big.Int, gateways []string) []byte {
	enc, err := universalResolverV2.abi.Pack("reverseWithGateways", lookupAddress, coinType, gateways)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackReverseWithGateways is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb7d6ca64.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function reverseWithGateways(bytes lookupAddress, uint256 coinType, string[] gateways) view returns(string, address, address)
func (universalResolverV2 *UniversalResolverV2) TryPackReverseWithGateways(lookupAddress []byte, coinType *big.Int, gateways []string) ([]byte, error) {
	return universalResolverV2.abi.Pack("reverseWithGateways", lookupAddress, coinType, gateways)
}

// ReverseWithGatewaysOutput serves as a container for the return parameters of contract
// method ReverseWithGateways.
type ReverseWithGatewaysOutput struct {
	Arg0 string
	Arg1 common.Address
	Arg2 common.Address
}

// UnpackReverseWithGateways is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb7d6ca64.
//
// Solidity: function reverseWithGateways(bytes lookupAddress, uint256 coinType, string[] gateways) view returns(string, address, address)
func (universalResolverV2 *UniversalResolverV2) UnpackReverseWithGateways(data []byte) (ReverseWithGatewaysOutput, error) {
	out, err := universalResolverV2.abi.Unpack("reverseWithGateways", data)
	outstruct := new(ReverseWithGatewaysOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Arg0 = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Arg1 = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Arg2 = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	return *outstruct, nil
}

// PackReverseWithGatewaysAndNormalization is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x76abcf2b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function reverseWithGatewaysAndNormalization(bytes lookupAddress, uint256 coinType, string[] gateways, address ensip15) view returns(string, address, address)
func (universalResolverV2 *UniversalResolverV2) PackReverseWithGatewaysAndNormalization(lookupAddress []byte, coinType *big.Int, gateways []string, ensip15 common.Address) []byte {
	enc, err := universalResolverV2.abi.Pack("reverseWithGatewaysAndNormalization", lookupAddress, coinType, gateways, ensip15)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackReverseWithGatewaysAndNormalization is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x76abcf2b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function reverseWithGatewaysAndNormalization(bytes lookupAddress, uint256 coinType, string[] gateways, address ensip15) view returns(string, address, address)
func (universalResolverV2 *UniversalResolverV2) TryPackReverseWithGatewaysAndNormalization(lookupAddress []byte, coinType *big.Int, gateways []string, ensip15 common.Address) ([]byte, error) {
	return universalResolverV2.abi.Pack("reverseWithGatewaysAndNormalization", lookupAddress, coinType, gateways, ensip15)
}

// ReverseWithGatewaysAndNormalizationOutput serves as a container for the return parameters of contract
// method ReverseWithGatewaysAndNormalization.
type ReverseWithGatewaysAndNormalizationOutput struct {
	Arg0 string
	Arg1 common.Address
	Arg2 common.Address
}

// UnpackReverseWithGatewaysAndNormalization is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x76abcf2b.
//
// Solidity: function reverseWithGatewaysAndNormalization(bytes lookupAddress, uint256 coinType, string[] gateways, address ensip15) view returns(string, address, address)
func (universalResolverV2 *UniversalResolverV2) UnpackReverseWithGatewaysAndNormalization(data []byte) (ReverseWithGatewaysAndNormalizationOutput, error) {
	out, err := universalResolverV2.abi.Unpack("reverseWithGatewaysAndNormalization", data)
	outstruct := new(ReverseWithGatewaysAndNormalizationOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Arg0 = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Arg1 = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Arg2 = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	return *outstruct, nil
}

// PackReverseWithNormalization is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4525ef23.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function reverseWithNormalization(bytes lookupAddress, uint256 coinType, address ensip15) view returns(string, address, address)
func (universalResolverV2 *UniversalResolverV2) PackReverseWithNormalization(lookupAddress []byte, coinType *big.Int, ensip15 common.Address) []byte {
	enc, err := universalResolverV2.abi.Pack("reverseWithNormalization", lookupAddress, coinType, ensip15)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackReverseWithNormalization is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4525ef23.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function reverseWithNormalization(bytes lookupAddress, uint256 coinType, address ensip15) view returns(string, address, address)
func (universalResolverV2 *UniversalResolverV2) TryPackReverseWithNormalization(lookupAddress []byte, coinType *big.Int, ensip15 common.Address) ([]byte, error) {
	return universalResolverV2.abi.Pack("reverseWithNormalization", lookupAddress, coinType, ensip15)
}

// ReverseWithNormalizationOutput serves as a container for the return parameters of contract
// method ReverseWithNormalization.
type ReverseWithNormalizationOutput struct {
	Arg0 string
	Arg1 common.Address
	Arg2 common.Address
}

// UnpackReverseWithNormalization is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x4525ef23.
//
// Solidity: function reverseWithNormalization(bytes lookupAddress, uint256 coinType, address ensip15) view returns(string, address, address)
func (universalResolverV2 *UniversalResolverV2) UnpackReverseWithNormalization(data []byte) (ReverseWithNormalizationOutput, error) {
	out, err := universalResolverV2.abi.Unpack("reverseWithNormalization", data)
	outstruct := new(ReverseWithNormalizationOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Arg0 = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Arg1 = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Arg2 = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	return *outstruct, nil
}

// PackSupportsInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01ffc9a7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (universalResolverV2 *UniversalResolverV2) PackSupportsInterface(interfaceId [4]byte) []byte {
	enc, err := universalResolverV2.abi.Pack("supportsInterface", interfaceId)
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
func (universalResolverV2 *UniversalResolverV2) TryPackSupportsInterface(interfaceId [4]byte) ([]byte, error) {
	return universalResolverV2.abi.Pack("supportsInterface", interfaceId)
}

// UnpackSupportsInterface is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (universalResolverV2 *UniversalResolverV2) UnpackSupportsInterface(data []byte) (bool, error) {
	out, err := universalResolverV2.abi.Unpack("supportsInterface", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// UnpackError attempts to decode the provided error data using user-defined
// error definitions.
func (universalResolverV2 *UniversalResolverV2) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], universalResolverV2.abi.Errors["DNSDecodingFailed"].ID.Bytes()[:4]) {
		return universalResolverV2.UnpackDNSDecodingFailedError(raw[4:])
	}
	if bytes.Equal(raw[:4], universalResolverV2.abi.Errors["DNSEncodingFailed"].ID.Bytes()[:4]) {
		return universalResolverV2.UnpackDNSEncodingFailedError(raw[4:])
	}
	if bytes.Equal(raw[:4], universalResolverV2.abi.Errors["EmptyAddress"].ID.Bytes()[:4]) {
		return universalResolverV2.UnpackEmptyAddressError(raw[4:])
	}
	if bytes.Equal(raw[:4], universalResolverV2.abi.Errors["HttpError"].ID.Bytes()[:4]) {
		return universalResolverV2.UnpackHttpErrorError(raw[4:])
	}
	if bytes.Equal(raw[:4], universalResolverV2.abi.Errors["InvalidBatchGatewayResponse"].ID.Bytes()[:4]) {
		return universalResolverV2.UnpackInvalidBatchGatewayResponseError(raw[4:])
	}
	if bytes.Equal(raw[:4], universalResolverV2.abi.Errors["LabelIsEmpty"].ID.Bytes()[:4]) {
		return universalResolverV2.UnpackLabelIsEmptyError(raw[4:])
	}
	if bytes.Equal(raw[:4], universalResolverV2.abi.Errors["LabelIsTooLong"].ID.Bytes()[:4]) {
		return universalResolverV2.UnpackLabelIsTooLongError(raw[4:])
	}
	if bytes.Equal(raw[:4], universalResolverV2.abi.Errors["NormalizationChangedName"].ID.Bytes()[:4]) {
		return universalResolverV2.UnpackNormalizationChangedNameError(raw[4:])
	}
	if bytes.Equal(raw[:4], universalResolverV2.abi.Errors["OffchainLookup"].ID.Bytes()[:4]) {
		return universalResolverV2.UnpackOffchainLookupError(raw[4:])
	}
	if bytes.Equal(raw[:4], universalResolverV2.abi.Errors["OffsetOutOfBoundsError"].ID.Bytes()[:4]) {
		return universalResolverV2.UnpackOffsetOutOfBoundsErrorError(raw[4:])
	}
	if bytes.Equal(raw[:4], universalResolverV2.abi.Errors["PrimaryNameNotNormalized"].ID.Bytes()[:4]) {
		return universalResolverV2.UnpackPrimaryNameNotNormalizedError(raw[4:])
	}
	if bytes.Equal(raw[:4], universalResolverV2.abi.Errors["ResolverError"].ID.Bytes()[:4]) {
		return universalResolverV2.UnpackResolverErrorError(raw[4:])
	}
	if bytes.Equal(raw[:4], universalResolverV2.abi.Errors["ResolverNotContract"].ID.Bytes()[:4]) {
		return universalResolverV2.UnpackResolverNotContractError(raw[4:])
	}
	if bytes.Equal(raw[:4], universalResolverV2.abi.Errors["ResolverNotFound"].ID.Bytes()[:4]) {
		return universalResolverV2.UnpackResolverNotFoundError(raw[4:])
	}
	if bytes.Equal(raw[:4], universalResolverV2.abi.Errors["ReverseAddressMismatch"].ID.Bytes()[:4]) {
		return universalResolverV2.UnpackReverseAddressMismatchError(raw[4:])
	}
	if bytes.Equal(raw[:4], universalResolverV2.abi.Errors["UnsafeBatchGatewayResponse"].ID.Bytes()[:4]) {
		return universalResolverV2.UnpackUnsafeBatchGatewayResponseError(raw[4:])
	}
	if bytes.Equal(raw[:4], universalResolverV2.abi.Errors["UnsupportedResolverProfile"].ID.Bytes()[:4]) {
		return universalResolverV2.UnpackUnsupportedResolverProfileError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// UniversalResolverV2DNSDecodingFailed represents a DNSDecodingFailed error raised by the UniversalResolverV2 contract.
type UniversalResolverV2DNSDecodingFailed struct {
	Dns []byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error DNSDecodingFailed(bytes dns)
func UniversalResolverV2DNSDecodingFailedErrorID() common.Hash {
	return common.HexToHash("0xba4adc2398f6571d4d3236eee815a96fcbbf84bbc4320e35b771331f71afdc72")
}

// UnpackDNSDecodingFailedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error DNSDecodingFailed(bytes dns)
func (universalResolverV2 *UniversalResolverV2) UnpackDNSDecodingFailedError(raw []byte) (*UniversalResolverV2DNSDecodingFailed, error) {
	out := new(UniversalResolverV2DNSDecodingFailed)
	if err := universalResolverV2.abi.UnpackIntoInterface(out, "DNSDecodingFailed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// UniversalResolverV2DNSEncodingFailed represents a DNSEncodingFailed error raised by the UniversalResolverV2 contract.
type UniversalResolverV2DNSEncodingFailed struct {
	Ens string
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error DNSEncodingFailed(string ens)
func UniversalResolverV2DNSEncodingFailedErrorID() common.Hash {
	return common.HexToHash("0x9a4c3e3b0f6c823624c6920224b1e57879d6eb956162632cc9b2fd0ae3def923")
}

// UnpackDNSEncodingFailedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error DNSEncodingFailed(string ens)
func (universalResolverV2 *UniversalResolverV2) UnpackDNSEncodingFailedError(raw []byte) (*UniversalResolverV2DNSEncodingFailed, error) {
	out := new(UniversalResolverV2DNSEncodingFailed)
	if err := universalResolverV2.abi.UnpackIntoInterface(out, "DNSEncodingFailed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// UniversalResolverV2EmptyAddress represents a EmptyAddress error raised by the UniversalResolverV2 contract.
type UniversalResolverV2EmptyAddress struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error EmptyAddress()
func UniversalResolverV2EmptyAddressErrorID() common.Hash {
	return common.HexToHash("0x7138356f89d71559551c4ff5ae4279ad57f02e8d7fad66641fbbfdba57c12a42")
}

// UnpackEmptyAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error EmptyAddress()
func (universalResolverV2 *UniversalResolverV2) UnpackEmptyAddressError(raw []byte) (*UniversalResolverV2EmptyAddress, error) {
	out := new(UniversalResolverV2EmptyAddress)
	if err := universalResolverV2.abi.UnpackIntoInterface(out, "EmptyAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// UniversalResolverV2HttpError represents a HttpError error raised by the UniversalResolverV2 contract.
type UniversalResolverV2HttpError struct {
	Status  uint16
	Message string
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error HttpError(uint16 status, string message)
func UniversalResolverV2HttpErrorErrorID() common.Hash {
	return common.HexToHash("0x0180015209d29d4a351f5e42b26b2968cfc25d5caabbbc5e3816b762c64e0581")
}

// UnpackHttpErrorError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error HttpError(uint16 status, string message)
func (universalResolverV2 *UniversalResolverV2) UnpackHttpErrorError(raw []byte) (*UniversalResolverV2HttpError, error) {
	out := new(UniversalResolverV2HttpError)
	if err := universalResolverV2.abi.UnpackIntoInterface(out, "HttpError", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// UniversalResolverV2InvalidBatchGatewayResponse represents a InvalidBatchGatewayResponse error raised by the UniversalResolverV2 contract.
type UniversalResolverV2InvalidBatchGatewayResponse struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidBatchGatewayResponse()
func UniversalResolverV2InvalidBatchGatewayResponseErrorID() common.Hash {
	return common.HexToHash("0x4a5c31ea8f03bbbdc9e81c61816d9410834a9fb0deb4eda201146ddb972cdb7a")
}

// UnpackInvalidBatchGatewayResponseError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidBatchGatewayResponse()
func (universalResolverV2 *UniversalResolverV2) UnpackInvalidBatchGatewayResponseError(raw []byte) (*UniversalResolverV2InvalidBatchGatewayResponse, error) {
	out := new(UniversalResolverV2InvalidBatchGatewayResponse)
	if err := universalResolverV2.abi.UnpackIntoInterface(out, "InvalidBatchGatewayResponse", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// UniversalResolverV2LabelIsEmpty represents a LabelIsEmpty error raised by the UniversalResolverV2 contract.
type UniversalResolverV2LabelIsEmpty struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error LabelIsEmpty()
func UniversalResolverV2LabelIsEmptyErrorID() common.Hash {
	return common.HexToHash("0xbf9a27402cc67e0a1553ebafddb9a8f09e2a0e17d33881522b49406d3206b53d")
}

// UnpackLabelIsEmptyError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error LabelIsEmpty()
func (universalResolverV2 *UniversalResolverV2) UnpackLabelIsEmptyError(raw []byte) (*UniversalResolverV2LabelIsEmpty, error) {
	out := new(UniversalResolverV2LabelIsEmpty)
	if err := universalResolverV2.abi.UnpackIntoInterface(out, "LabelIsEmpty", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// UniversalResolverV2LabelIsTooLong represents a LabelIsTooLong error raised by the UniversalResolverV2 contract.
type UniversalResolverV2LabelIsTooLong struct {
	Label string
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error LabelIsTooLong(string label)
func UniversalResolverV2LabelIsTooLongErrorID() common.Hash {
	return common.HexToHash("0xdab6c73ccd5b3a9a45fe88ea48179efbe644dbc9416a195d7779ed8ccd248e63")
}

// UnpackLabelIsTooLongError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error LabelIsTooLong(string label)
func (universalResolverV2 *UniversalResolverV2) UnpackLabelIsTooLongError(raw []byte) (*UniversalResolverV2LabelIsTooLong, error) {
	out := new(UniversalResolverV2LabelIsTooLong)
	if err := universalResolverV2.abi.UnpackIntoInterface(out, "LabelIsTooLong", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// UniversalResolverV2NormalizationChangedName represents a NormalizationChangedName error raised by the UniversalResolverV2 contract.
type UniversalResolverV2NormalizationChangedName struct {
	NormalizedName []byte
	Result         []byte
	Resolver       common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NormalizationChangedName(bytes normalizedName, bytes result, address resolver)
func UniversalResolverV2NormalizationChangedNameErrorID() common.Hash {
	return common.HexToHash("0x2dba135371e1afb266a6a49b4df807b2ff61b237b3141eb5fd82353172532747")
}

// UnpackNormalizationChangedNameError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NormalizationChangedName(bytes normalizedName, bytes result, address resolver)
func (universalResolverV2 *UniversalResolverV2) UnpackNormalizationChangedNameError(raw []byte) (*UniversalResolverV2NormalizationChangedName, error) {
	out := new(UniversalResolverV2NormalizationChangedName)
	if err := universalResolverV2.abi.UnpackIntoInterface(out, "NormalizationChangedName", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// UniversalResolverV2OffchainLookup represents a OffchainLookup error raised by the UniversalResolverV2 contract.
type UniversalResolverV2OffchainLookup struct {
	Sender           common.Address
	Urls             []string
	CallData         []byte
	CallbackFunction [4]byte
	ExtraData        []byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OffchainLookup(address sender, string[] urls, bytes callData, bytes4 callbackFunction, bytes extraData)
func UniversalResolverV2OffchainLookupErrorID() common.Hash {
	return common.HexToHash("0x556f18302ad20acabcb7abd78e46143543eff9de9c87236fb097a1398d41115e")
}

// UnpackOffchainLookupError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OffchainLookup(address sender, string[] urls, bytes callData, bytes4 callbackFunction, bytes extraData)
func (universalResolverV2 *UniversalResolverV2) UnpackOffchainLookupError(raw []byte) (*UniversalResolverV2OffchainLookup, error) {
	out := new(UniversalResolverV2OffchainLookup)
	if err := universalResolverV2.abi.UnpackIntoInterface(out, "OffchainLookup", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// UniversalResolverV2OffsetOutOfBoundsError represents a OffsetOutOfBoundsError error raised by the UniversalResolverV2 contract.
type UniversalResolverV2OffsetOutOfBoundsError struct {
	Offset *big.Int
	Length *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OffsetOutOfBoundsError(uint256 offset, uint256 length)
func UniversalResolverV2OffsetOutOfBoundsErrorErrorID() common.Hash {
	return common.HexToHash("0x8a3c1cfb98bd88fadc1ef4bc86d0b221827194fb788772725d6bd414678ba0eb")
}

// UnpackOffsetOutOfBoundsErrorError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OffsetOutOfBoundsError(uint256 offset, uint256 length)
func (universalResolverV2 *UniversalResolverV2) UnpackOffsetOutOfBoundsErrorError(raw []byte) (*UniversalResolverV2OffsetOutOfBoundsError, error) {
	out := new(UniversalResolverV2OffsetOutOfBoundsError)
	if err := universalResolverV2.abi.UnpackIntoInterface(out, "OffsetOutOfBoundsError", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// UniversalResolverV2PrimaryNameNotNormalized represents a PrimaryNameNotNormalized error raised by the UniversalResolverV2 contract.
type UniversalResolverV2PrimaryNameNotNormalized struct {
	Primary string
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error PrimaryNameNotNormalized(string primary)
func UniversalResolverV2PrimaryNameNotNormalizedErrorID() common.Hash {
	return common.HexToHash("0x28c6aa0cca3df69075cecd38112e99d5092ec910e1c923f527d1616a2d473b19")
}

// UnpackPrimaryNameNotNormalizedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error PrimaryNameNotNormalized(string primary)
func (universalResolverV2 *UniversalResolverV2) UnpackPrimaryNameNotNormalizedError(raw []byte) (*UniversalResolverV2PrimaryNameNotNormalized, error) {
	out := new(UniversalResolverV2PrimaryNameNotNormalized)
	if err := universalResolverV2.abi.UnpackIntoInterface(out, "PrimaryNameNotNormalized", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// UniversalResolverV2ResolverError represents a ResolverError error raised by the UniversalResolverV2 contract.
type UniversalResolverV2ResolverError struct {
	ErrorData []byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ResolverError(bytes errorData)
func UniversalResolverV2ResolverErrorErrorID() common.Hash {
	return common.HexToHash("0x95c0c7522237b23bbb1e132d04090a8ff995b438a91575150cb3fc7a7741bfd4")
}

// UnpackResolverErrorError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ResolverError(bytes errorData)
func (universalResolverV2 *UniversalResolverV2) UnpackResolverErrorError(raw []byte) (*UniversalResolverV2ResolverError, error) {
	out := new(UniversalResolverV2ResolverError)
	if err := universalResolverV2.abi.UnpackIntoInterface(out, "ResolverError", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// UniversalResolverV2ResolverNotContract represents a ResolverNotContract error raised by the UniversalResolverV2 contract.
type UniversalResolverV2ResolverNotContract struct {
	Name     []byte
	Resolver common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ResolverNotContract(bytes name, address resolver)
func UniversalResolverV2ResolverNotContractErrorID() common.Hash {
	return common.HexToHash("0x1e9535f20018e88ec379670696a41c8379b2e032c8ae329aa11058a21994d8ef")
}

// UnpackResolverNotContractError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ResolverNotContract(bytes name, address resolver)
func (universalResolverV2 *UniversalResolverV2) UnpackResolverNotContractError(raw []byte) (*UniversalResolverV2ResolverNotContract, error) {
	out := new(UniversalResolverV2ResolverNotContract)
	if err := universalResolverV2.abi.UnpackIntoInterface(out, "ResolverNotContract", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// UniversalResolverV2ResolverNotFound represents a ResolverNotFound error raised by the UniversalResolverV2 contract.
type UniversalResolverV2ResolverNotFound struct {
	Name []byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ResolverNotFound(bytes name)
func UniversalResolverV2ResolverNotFoundErrorID() common.Hash {
	return common.HexToHash("0x77209fe8d5fcf80d9dd1a5bf50067976a316191331973ed6d4876c758ab3475b")
}

// UnpackResolverNotFoundError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ResolverNotFound(bytes name)
func (universalResolverV2 *UniversalResolverV2) UnpackResolverNotFoundError(raw []byte) (*UniversalResolverV2ResolverNotFound, error) {
	out := new(UniversalResolverV2ResolverNotFound)
	if err := universalResolverV2.abi.UnpackIntoInterface(out, "ResolverNotFound", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// UniversalResolverV2ReverseAddressMismatch represents a ReverseAddressMismatch error raised by the UniversalResolverV2 contract.
type UniversalResolverV2ReverseAddressMismatch struct {
	Primary        string
	PrimaryAddress []byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ReverseAddressMismatch(string primary, bytes primaryAddress)
func UniversalResolverV2ReverseAddressMismatchErrorID() common.Hash {
	return common.HexToHash("0xef9c03ce87b28a9b7d889641b3468f4d0fbafa4e8001ea211cc120f82577ae23")
}

// UnpackReverseAddressMismatchError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ReverseAddressMismatch(string primary, bytes primaryAddress)
func (universalResolverV2 *UniversalResolverV2) UnpackReverseAddressMismatchError(raw []byte) (*UniversalResolverV2ReverseAddressMismatch, error) {
	out := new(UniversalResolverV2ReverseAddressMismatch)
	if err := universalResolverV2.abi.UnpackIntoInterface(out, "ReverseAddressMismatch", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// UniversalResolverV2UnsafeBatchGatewayResponse represents a UnsafeBatchGatewayResponse error raised by the UniversalResolverV2 contract.
type UniversalResolverV2UnsafeBatchGatewayResponse struct {
	Arg0 []byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UnsafeBatchGatewayResponse(bytes arg0)
func UniversalResolverV2UnsafeBatchGatewayResponseErrorID() common.Hash {
	return common.HexToHash("0x2fa87250bd9d28eeb5d50d876277daead92ec8f4759ce95a149361cc36fe344a")
}

// UnpackUnsafeBatchGatewayResponseError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UnsafeBatchGatewayResponse(bytes arg0)
func (universalResolverV2 *UniversalResolverV2) UnpackUnsafeBatchGatewayResponseError(raw []byte) (*UniversalResolverV2UnsafeBatchGatewayResponse, error) {
	out := new(UniversalResolverV2UnsafeBatchGatewayResponse)
	if err := universalResolverV2.abi.UnpackIntoInterface(out, "UnsafeBatchGatewayResponse", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// UniversalResolverV2UnsupportedResolverProfile represents a UnsupportedResolverProfile error raised by the UniversalResolverV2 contract.
type UniversalResolverV2UnsupportedResolverProfile struct {
	Selector [4]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UnsupportedResolverProfile(bytes4 selector)
func UniversalResolverV2UnsupportedResolverProfileErrorID() common.Hash {
	return common.HexToHash("0x7b1c461b9d31387598812449409a83d2fcd215e9f43e0d58ccc6333187d37d79")
}

// UnpackUnsupportedResolverProfileError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UnsupportedResolverProfile(bytes4 selector)
func (universalResolverV2 *UniversalResolverV2) UnpackUnsupportedResolverProfileError(raw []byte) (*UniversalResolverV2UnsupportedResolverProfile, error) {
	out := new(UniversalResolverV2UnsupportedResolverProfile)
	if err := universalResolverV2.abi.UnpackIntoInterface(out, "UnsupportedResolverProfile", raw); err != nil {
		return nil, err
	}
	return out, nil
}
