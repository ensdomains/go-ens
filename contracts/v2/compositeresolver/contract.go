// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package compositeresolver

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

// CompositeResolverMetaData contains all meta data concerning the CompositeResolver contract.
var CompositeResolverMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIGatewayProvider\",\"name\":\"batchGatewayProvider\",\"type\":\"address\"},{\"internalType\":\"contractIContractNamer\",\"name\":\"contractNamer\",\"type\":\"address\"},{\"internalType\":\"contractENS\",\"name\":\"registryV1\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"dns\",\"type\":\"bytes\"}],\"name\":\"DNSDecodingFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBatchGatewayResponse\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"urls\",\"type\":\"string[]\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes4\",\"name\":\"callbackFunction\",\"type\":\"bytes4\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"OffchainLookup\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"OffsetOutOfBoundsError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"}],\"name\":\"UnreachableName\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"UnsafeBatchGatewayResponse\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BATCH_GATEWAY_PROVIDER\",\"outputs\":[{\"internalType\":\"contractIGatewayProvider\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CONTRACT_NAMER\",\"outputs\":[{\"internalType\":\"contractIContractNamer\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REGISTRY_V1\",\"outputs\":[{\"internalType\":\"contractENS\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"hasContext\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"context\",\"type\":\"bytes\"},{\"internalType\":\"string[]\",\"name\":\"batchGateways\",\"type\":\"string[]\"}],\"name\":\"callResolver\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"call\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"}],\"internalType\":\"structCCIPBatcher.Lookup[]\",\"name\":\"lookups\",\"type\":\"tuple[]\"},{\"internalType\":\"string[]\",\"name\":\"gateways\",\"type\":\"string[]\"}],\"internalType\":\"structCCIPBatcher.Batch\",\"name\":\"batch\",\"type\":\"tuple\"}],\"name\":\"ccipBatch\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"call\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"}],\"internalType\":\"structCCIPBatcher.Lookup[]\",\"name\":\"lookups\",\"type\":\"tuple[]\"},{\"internalType\":\"string[]\",\"name\":\"gateways\",\"type\":\"string[]\"}],\"internalType\":\"structCCIPBatcher.Batch\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"response\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"ccipBatchCallback\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"call\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"}],\"internalType\":\"structCCIPBatcher.Lookup[]\",\"name\":\"lookups\",\"type\":\"tuple[]\"},{\"internalType\":\"string[]\",\"name\":\"gateways\",\"type\":\"string[]\"}],\"internalType\":\"structCCIPBatcher.Batch\",\"name\":\"batch\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"response\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"ccipReadCallback\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"}],\"name\":\"getResolver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"namer\",\"type\":\"address\"}],\"name\":\"isContractNamer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"resolve\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"response\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"resolveBatchCallback\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"response\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"resolveDirectImmediateCallback\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"feature\",\"type\":\"bytes4\"}],\"name\":\"supportsFeature\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	ID:  "CompositeResolver",
}

// CompositeResolver is an auto generated Go binding around an Ethereum contract.
type CompositeResolver struct {
	abi abi.ABI
}

// GetABI returns the ABI associated with this contract binding.
func (c *CompositeResolver) GetABI() abi.ABI {
	return c.abi
}

// NewCompositeResolver creates a new instance of CompositeResolver.
func NewCompositeResolver() *CompositeResolver {
	parsed, err := CompositeResolverMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &CompositeResolver{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *CompositeResolver) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address batchGatewayProvider, address contractNamer, address registryV1) returns()
func (compositeResolver *CompositeResolver) PackConstructor(batchGatewayProvider common.Address, contractNamer common.Address, registryV1 common.Address) []byte {
	enc, err := compositeResolver.abi.Pack("", batchGatewayProvider, contractNamer, registryV1)
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
func (compositeResolver *CompositeResolver) PackBATCHGATEWAYPROVIDER() []byte {
	enc, err := compositeResolver.abi.Pack("BATCH_GATEWAY_PROVIDER")
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
func (compositeResolver *CompositeResolver) TryPackBATCHGATEWAYPROVIDER() ([]byte, error) {
	return compositeResolver.abi.Pack("BATCH_GATEWAY_PROVIDER")
}

// UnpackBATCHGATEWAYPROVIDER is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6ccb8660.
//
// Solidity: function BATCH_GATEWAY_PROVIDER() view returns(address)
func (compositeResolver *CompositeResolver) UnpackBATCHGATEWAYPROVIDER(data []byte) (common.Address, error) {
	out, err := compositeResolver.abi.Unpack("BATCH_GATEWAY_PROVIDER", data)
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
func (compositeResolver *CompositeResolver) PackCONTRACTNAMER() []byte {
	enc, err := compositeResolver.abi.Pack("CONTRACT_NAMER")
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
func (compositeResolver *CompositeResolver) TryPackCONTRACTNAMER() ([]byte, error) {
	return compositeResolver.abi.Pack("CONTRACT_NAMER")
}

// UnpackCONTRACTNAMER is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x48ee1bcc.
//
// Solidity: function CONTRACT_NAMER() view returns(address)
func (compositeResolver *CompositeResolver) UnpackCONTRACTNAMER(data []byte) (common.Address, error) {
	out, err := compositeResolver.abi.Unpack("CONTRACT_NAMER", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackREGISTRYV1 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe370ecbe.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function REGISTRY_V1() view returns(address)
func (compositeResolver *CompositeResolver) PackREGISTRYV1() []byte {
	enc, err := compositeResolver.abi.Pack("REGISTRY_V1")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackREGISTRYV1 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe370ecbe.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function REGISTRY_V1() view returns(address)
func (compositeResolver *CompositeResolver) TryPackREGISTRYV1() ([]byte, error) {
	return compositeResolver.abi.Pack("REGISTRY_V1")
}

// UnpackREGISTRYV1 is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe370ecbe.
//
// Solidity: function REGISTRY_V1() view returns(address)
func (compositeResolver *CompositeResolver) UnpackREGISTRYV1(data []byte) (common.Address, error) {
	out, err := compositeResolver.abi.Unpack("REGISTRY_V1", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackCallResolver is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf394443a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function callResolver(address resolver, bytes name, bytes data, bool hasContext, bytes context, string[] batchGateways) view returns(bytes)
func (compositeResolver *CompositeResolver) PackCallResolver(resolver common.Address, name []byte, data []byte, hasContext bool, context []byte, batchGateways []string) []byte {
	enc, err := compositeResolver.abi.Pack("callResolver", resolver, name, data, hasContext, context, batchGateways)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCallResolver is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf394443a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function callResolver(address resolver, bytes name, bytes data, bool hasContext, bytes context, string[] batchGateways) view returns(bytes)
func (compositeResolver *CompositeResolver) TryPackCallResolver(resolver common.Address, name []byte, data []byte, hasContext bool, context []byte, batchGateways []string) ([]byte, error) {
	return compositeResolver.abi.Pack("callResolver", resolver, name, data, hasContext, context, batchGateways)
}

// UnpackCallResolver is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf394443a.
//
// Solidity: function callResolver(address resolver, bytes name, bytes data, bool hasContext, bytes context, string[] batchGateways) view returns(bytes)
func (compositeResolver *CompositeResolver) UnpackCallResolver(data []byte) ([]byte, error) {
	out, err := compositeResolver.abi.Unpack("callResolver", data)
	if err != nil {
		return *new([]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	return out0, nil
}

// PackCcipBatch is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9f28e99d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function ccipBatch(((address,bytes,bytes,uint256)[],string[]) batch) view returns(((address,bytes,bytes,uint256)[],string[]))
func (compositeResolver *CompositeResolver) PackCcipBatch(batch CCIPBatcherBatch) []byte {
	enc, err := compositeResolver.abi.Pack("ccipBatch", batch)
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
func (compositeResolver *CompositeResolver) TryPackCcipBatch(batch CCIPBatcherBatch) ([]byte, error) {
	return compositeResolver.abi.Pack("ccipBatch", batch)
}

// UnpackCcipBatch is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9f28e99d.
//
// Solidity: function ccipBatch(((address,bytes,bytes,uint256)[],string[]) batch) view returns(((address,bytes,bytes,uint256)[],string[]))
func (compositeResolver *CompositeResolver) UnpackCcipBatch(data []byte) (CCIPBatcherBatch, error) {
	out, err := compositeResolver.abi.Unpack("ccipBatch", data)
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
func (compositeResolver *CompositeResolver) PackCcipBatchCallback(response []byte, extraData []byte) []byte {
	enc, err := compositeResolver.abi.Pack("ccipBatchCallback", response, extraData)
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
func (compositeResolver *CompositeResolver) TryPackCcipBatchCallback(response []byte, extraData []byte) ([]byte, error) {
	return compositeResolver.abi.Pack("ccipBatchCallback", response, extraData)
}

// UnpackCcipBatchCallback is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb536af76.
//
// Solidity: function ccipBatchCallback(bytes response, bytes extraData) view returns(((address,bytes,bytes,uint256)[],string[]) batch)
func (compositeResolver *CompositeResolver) UnpackCcipBatchCallback(data []byte) (CCIPBatcherBatch, error) {
	out, err := compositeResolver.abi.Unpack("ccipBatchCallback", data)
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
func (compositeResolver *CompositeResolver) PackCcipReadCallback(response []byte, extraData []byte) []byte {
	enc, err := compositeResolver.abi.Pack("ccipReadCallback", response, extraData)
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
func (compositeResolver *CompositeResolver) TryPackCcipReadCallback(response []byte, extraData []byte) ([]byte, error) {
	return compositeResolver.abi.Pack("ccipReadCallback", response, extraData)
}

// PackGetResolver is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xeea330f9.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getResolver(bytes name) view returns(address, bool)
func (compositeResolver *CompositeResolver) PackGetResolver(name []byte) []byte {
	enc, err := compositeResolver.abi.Pack("getResolver", name)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetResolver is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xeea330f9.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getResolver(bytes name) view returns(address, bool)
func (compositeResolver *CompositeResolver) TryPackGetResolver(name []byte) ([]byte, error) {
	return compositeResolver.abi.Pack("getResolver", name)
}

// GetResolverOutput serves as a container for the return parameters of contract
// method GetResolver.
type GetResolverOutput struct {
	Arg0 common.Address
	Arg1 bool
}

// UnpackGetResolver is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xeea330f9.
//
// Solidity: function getResolver(bytes name) view returns(address, bool)
func (compositeResolver *CompositeResolver) UnpackGetResolver(data []byte) (GetResolverOutput, error) {
	out, err := compositeResolver.abi.Unpack("getResolver", data)
	outstruct := new(GetResolverOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Arg0 = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Arg1 = *abi.ConvertType(out[1], new(bool)).(*bool)
	return *outstruct, nil
}

// PackIsContractNamer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6f3ff726.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isContractNamer(address namer) view returns(bool)
func (compositeResolver *CompositeResolver) PackIsContractNamer(namer common.Address) []byte {
	enc, err := compositeResolver.abi.Pack("isContractNamer", namer)
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
func (compositeResolver *CompositeResolver) TryPackIsContractNamer(namer common.Address) ([]byte, error) {
	return compositeResolver.abi.Pack("isContractNamer", namer)
}

// UnpackIsContractNamer is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6f3ff726.
//
// Solidity: function isContractNamer(address namer) view returns(bool)
func (compositeResolver *CompositeResolver) UnpackIsContractNamer(data []byte) (bool, error) {
	out, err := compositeResolver.abi.Unpack("isContractNamer", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackResolve is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9061b923.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function resolve(bytes name, bytes data) view returns(bytes)
func (compositeResolver *CompositeResolver) PackResolve(name []byte, data []byte) []byte {
	enc, err := compositeResolver.abi.Pack("resolve", name, data)
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
func (compositeResolver *CompositeResolver) TryPackResolve(name []byte, data []byte) ([]byte, error) {
	return compositeResolver.abi.Pack("resolve", name, data)
}

// UnpackResolve is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9061b923.
//
// Solidity: function resolve(bytes name, bytes data) view returns(bytes)
func (compositeResolver *CompositeResolver) UnpackResolve(data []byte) ([]byte, error) {
	out, err := compositeResolver.abi.Unpack("resolve", data)
	if err != nil {
		return *new([]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	return out0, nil
}

// PackResolveBatchCallback is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x491fc4f9.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function resolveBatchCallback(bytes response, bytes extraData) pure returns(bytes)
func (compositeResolver *CompositeResolver) PackResolveBatchCallback(response []byte, extraData []byte) []byte {
	enc, err := compositeResolver.abi.Pack("resolveBatchCallback", response, extraData)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackResolveBatchCallback is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x491fc4f9.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function resolveBatchCallback(bytes response, bytes extraData) pure returns(bytes)
func (compositeResolver *CompositeResolver) TryPackResolveBatchCallback(response []byte, extraData []byte) ([]byte, error) {
	return compositeResolver.abi.Pack("resolveBatchCallback", response, extraData)
}

// UnpackResolveBatchCallback is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x491fc4f9.
//
// Solidity: function resolveBatchCallback(bytes response, bytes extraData) pure returns(bytes)
func (compositeResolver *CompositeResolver) UnpackResolveBatchCallback(data []byte) ([]byte, error) {
	out, err := compositeResolver.abi.Unpack("resolveBatchCallback", data)
	if err != nil {
		return *new([]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	return out0, nil
}

// PackResolveDirectImmediateCallback is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6d6dd540.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function resolveDirectImmediateCallback(bytes response, bytes ) pure returns(bytes)
func (compositeResolver *CompositeResolver) PackResolveDirectImmediateCallback(response []byte, arg1 []byte) []byte {
	enc, err := compositeResolver.abi.Pack("resolveDirectImmediateCallback", response, arg1)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackResolveDirectImmediateCallback is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6d6dd540.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function resolveDirectImmediateCallback(bytes response, bytes ) pure returns(bytes)
func (compositeResolver *CompositeResolver) TryPackResolveDirectImmediateCallback(response []byte, arg1 []byte) ([]byte, error) {
	return compositeResolver.abi.Pack("resolveDirectImmediateCallback", response, arg1)
}

// UnpackResolveDirectImmediateCallback is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6d6dd540.
//
// Solidity: function resolveDirectImmediateCallback(bytes response, bytes ) pure returns(bytes)
func (compositeResolver *CompositeResolver) UnpackResolveDirectImmediateCallback(data []byte) ([]byte, error) {
	out, err := compositeResolver.abi.Unpack("resolveDirectImmediateCallback", data)
	if err != nil {
		return *new([]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	return out0, nil
}

// PackSupportsFeature is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x582de3e7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function supportsFeature(bytes4 feature) pure returns(bool)
func (compositeResolver *CompositeResolver) PackSupportsFeature(feature [4]byte) []byte {
	enc, err := compositeResolver.abi.Pack("supportsFeature", feature)
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
func (compositeResolver *CompositeResolver) TryPackSupportsFeature(feature [4]byte) ([]byte, error) {
	return compositeResolver.abi.Pack("supportsFeature", feature)
}

// UnpackSupportsFeature is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x582de3e7.
//
// Solidity: function supportsFeature(bytes4 feature) pure returns(bool)
func (compositeResolver *CompositeResolver) UnpackSupportsFeature(data []byte) (bool, error) {
	out, err := compositeResolver.abi.Unpack("supportsFeature", data)
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
func (compositeResolver *CompositeResolver) PackSupportsInterface(interfaceId [4]byte) []byte {
	enc, err := compositeResolver.abi.Pack("supportsInterface", interfaceId)
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
func (compositeResolver *CompositeResolver) TryPackSupportsInterface(interfaceId [4]byte) ([]byte, error) {
	return compositeResolver.abi.Pack("supportsInterface", interfaceId)
}

// UnpackSupportsInterface is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (compositeResolver *CompositeResolver) UnpackSupportsInterface(data []byte) (bool, error) {
	out, err := compositeResolver.abi.Unpack("supportsInterface", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// UnpackError attempts to decode the provided error data using user-defined
// error definitions.
func (compositeResolver *CompositeResolver) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], compositeResolver.abi.Errors["DNSDecodingFailed"].ID.Bytes()[:4]) {
		return compositeResolver.UnpackDNSDecodingFailedError(raw[4:])
	}
	if bytes.Equal(raw[:4], compositeResolver.abi.Errors["InvalidBatchGatewayResponse"].ID.Bytes()[:4]) {
		return compositeResolver.UnpackInvalidBatchGatewayResponseError(raw[4:])
	}
	if bytes.Equal(raw[:4], compositeResolver.abi.Errors["OffchainLookup"].ID.Bytes()[:4]) {
		return compositeResolver.UnpackOffchainLookupError(raw[4:])
	}
	if bytes.Equal(raw[:4], compositeResolver.abi.Errors["OffsetOutOfBoundsError"].ID.Bytes()[:4]) {
		return compositeResolver.UnpackOffsetOutOfBoundsErrorError(raw[4:])
	}
	if bytes.Equal(raw[:4], compositeResolver.abi.Errors["UnreachableName"].ID.Bytes()[:4]) {
		return compositeResolver.UnpackUnreachableNameError(raw[4:])
	}
	if bytes.Equal(raw[:4], compositeResolver.abi.Errors["UnsafeBatchGatewayResponse"].ID.Bytes()[:4]) {
		return compositeResolver.UnpackUnsafeBatchGatewayResponseError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// CompositeResolverDNSDecodingFailed represents a DNSDecodingFailed error raised by the CompositeResolver contract.
type CompositeResolverDNSDecodingFailed struct {
	Dns []byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error DNSDecodingFailed(bytes dns)
func CompositeResolverDNSDecodingFailedErrorID() common.Hash {
	return common.HexToHash("0xba4adc2398f6571d4d3236eee815a96fcbbf84bbc4320e35b771331f71afdc72")
}

// UnpackDNSDecodingFailedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error DNSDecodingFailed(bytes dns)
func (compositeResolver *CompositeResolver) UnpackDNSDecodingFailedError(raw []byte) (*CompositeResolverDNSDecodingFailed, error) {
	out := new(CompositeResolverDNSDecodingFailed)
	if err := compositeResolver.abi.UnpackIntoInterface(out, "DNSDecodingFailed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CompositeResolverInvalidBatchGatewayResponse represents a InvalidBatchGatewayResponse error raised by the CompositeResolver contract.
type CompositeResolverInvalidBatchGatewayResponse struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidBatchGatewayResponse()
func CompositeResolverInvalidBatchGatewayResponseErrorID() common.Hash {
	return common.HexToHash("0x4a5c31ea8f03bbbdc9e81c61816d9410834a9fb0deb4eda201146ddb972cdb7a")
}

// UnpackInvalidBatchGatewayResponseError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidBatchGatewayResponse()
func (compositeResolver *CompositeResolver) UnpackInvalidBatchGatewayResponseError(raw []byte) (*CompositeResolverInvalidBatchGatewayResponse, error) {
	out := new(CompositeResolverInvalidBatchGatewayResponse)
	if err := compositeResolver.abi.UnpackIntoInterface(out, "InvalidBatchGatewayResponse", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CompositeResolverOffchainLookup represents a OffchainLookup error raised by the CompositeResolver contract.
type CompositeResolverOffchainLookup struct {
	Sender           common.Address
	Urls             []string
	CallData         []byte
	CallbackFunction [4]byte
	ExtraData        []byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OffchainLookup(address sender, string[] urls, bytes callData, bytes4 callbackFunction, bytes extraData)
func CompositeResolverOffchainLookupErrorID() common.Hash {
	return common.HexToHash("0x556f18302ad20acabcb7abd78e46143543eff9de9c87236fb097a1398d41115e")
}

// UnpackOffchainLookupError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OffchainLookup(address sender, string[] urls, bytes callData, bytes4 callbackFunction, bytes extraData)
func (compositeResolver *CompositeResolver) UnpackOffchainLookupError(raw []byte) (*CompositeResolverOffchainLookup, error) {
	out := new(CompositeResolverOffchainLookup)
	if err := compositeResolver.abi.UnpackIntoInterface(out, "OffchainLookup", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CompositeResolverOffsetOutOfBoundsError represents a OffsetOutOfBoundsError error raised by the CompositeResolver contract.
type CompositeResolverOffsetOutOfBoundsError struct {
	Offset *big.Int
	Length *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OffsetOutOfBoundsError(uint256 offset, uint256 length)
func CompositeResolverOffsetOutOfBoundsErrorErrorID() common.Hash {
	return common.HexToHash("0x8a3c1cfb98bd88fadc1ef4bc86d0b221827194fb788772725d6bd414678ba0eb")
}

// UnpackOffsetOutOfBoundsErrorError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OffsetOutOfBoundsError(uint256 offset, uint256 length)
func (compositeResolver *CompositeResolver) UnpackOffsetOutOfBoundsErrorError(raw []byte) (*CompositeResolverOffsetOutOfBoundsError, error) {
	out := new(CompositeResolverOffsetOutOfBoundsError)
	if err := compositeResolver.abi.UnpackIntoInterface(out, "OffsetOutOfBoundsError", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CompositeResolverUnreachableName represents a UnreachableName error raised by the CompositeResolver contract.
type CompositeResolverUnreachableName struct {
	Name []byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UnreachableName(bytes name)
func CompositeResolverUnreachableNameErrorID() common.Hash {
	return common.HexToHash("0x5fe9a5df6f815b4549d6ea64689d1bd343919bbac2703a97aaf0038e91db0b67")
}

// UnpackUnreachableNameError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UnreachableName(bytes name)
func (compositeResolver *CompositeResolver) UnpackUnreachableNameError(raw []byte) (*CompositeResolverUnreachableName, error) {
	out := new(CompositeResolverUnreachableName)
	if err := compositeResolver.abi.UnpackIntoInterface(out, "UnreachableName", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CompositeResolverUnsafeBatchGatewayResponse represents a UnsafeBatchGatewayResponse error raised by the CompositeResolver contract.
type CompositeResolverUnsafeBatchGatewayResponse struct {
	Arg0 []byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UnsafeBatchGatewayResponse(bytes arg0)
func CompositeResolverUnsafeBatchGatewayResponseErrorID() common.Hash {
	return common.HexToHash("0x2fa87250bd9d28eeb5d50d876277daead92ec8f4759ce95a149361cc36fe344a")
}

// UnpackUnsafeBatchGatewayResponseError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UnsafeBatchGatewayResponse(bytes arg0)
func (compositeResolver *CompositeResolver) UnpackUnsafeBatchGatewayResponseError(raw []byte) (*CompositeResolverUnsafeBatchGatewayResponse, error) {
	out := new(CompositeResolverUnsafeBatchGatewayResponse)
	if err := compositeResolver.abi.UnpackIntoInterface(out, "UnsafeBatchGatewayResponse", raw); err != nil {
		return nil, err
	}
	return out, nil
}
