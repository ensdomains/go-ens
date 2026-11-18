// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package labelstore

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

// LabelStoreMetaData contains all meta data concerning the LabelStore contract.
var LabelStoreMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIContractNamer\",\"name\":\"contractNamer\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"LabelIsEmpty\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"}],\"name\":\"LabelIsTooLong\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"labelHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"}],\"name\":\"Label\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CONTRACT_NAMER\",\"outputs\":[{\"internalType\":\"contractIContractNamer\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"anyId\",\"type\":\"uint256\"}],\"name\":\"getLabel\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"namer\",\"type\":\"address\"}],\"name\":\"isContractNamer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"}],\"name\":\"setLabel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	ID:  "LabelStore",
}

// LabelStore is an auto generated Go binding around an Ethereum contract.
type LabelStore struct {
	abi abi.ABI
}

// GetABI returns the ABI associated with this contract binding.
func (c *LabelStore) GetABI() abi.ABI {
	return c.abi
}

// NewLabelStore creates a new instance of LabelStore.
func NewLabelStore() *LabelStore {
	parsed, err := LabelStoreMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &LabelStore{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *LabelStore) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address contractNamer) returns()
func (labelStore *LabelStore) PackConstructor(contractNamer common.Address) []byte {
	enc, err := labelStore.abi.Pack("", contractNamer)
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
func (labelStore *LabelStore) PackCONTRACTNAMER() []byte {
	enc, err := labelStore.abi.Pack("CONTRACT_NAMER")
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
func (labelStore *LabelStore) TryPackCONTRACTNAMER() ([]byte, error) {
	return labelStore.abi.Pack("CONTRACT_NAMER")
}

// UnpackCONTRACTNAMER is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x48ee1bcc.
//
// Solidity: function CONTRACT_NAMER() view returns(address)
func (labelStore *LabelStore) UnpackCONTRACTNAMER(data []byte) (common.Address, error) {
	out, err := labelStore.abi.Unpack("CONTRACT_NAMER", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetLabel is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb21bf7fa.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getLabel(uint256 anyId) view returns(string)
func (labelStore *LabelStore) PackGetLabel(anyId *big.Int) []byte {
	enc, err := labelStore.abi.Pack("getLabel", anyId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetLabel is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb21bf7fa.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getLabel(uint256 anyId) view returns(string)
func (labelStore *LabelStore) TryPackGetLabel(anyId *big.Int) ([]byte, error) {
	return labelStore.abi.Pack("getLabel", anyId)
}

// UnpackGetLabel is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb21bf7fa.
//
// Solidity: function getLabel(uint256 anyId) view returns(string)
func (labelStore *LabelStore) UnpackGetLabel(data []byte) (string, error) {
	out, err := labelStore.abi.Unpack("getLabel", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, nil
}

// PackIsContractNamer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6f3ff726.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isContractNamer(address namer) view returns(bool)
func (labelStore *LabelStore) PackIsContractNamer(namer common.Address) []byte {
	enc, err := labelStore.abi.Pack("isContractNamer", namer)
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
func (labelStore *LabelStore) TryPackIsContractNamer(namer common.Address) ([]byte, error) {
	return labelStore.abi.Pack("isContractNamer", namer)
}

// UnpackIsContractNamer is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6f3ff726.
//
// Solidity: function isContractNamer(address namer) view returns(bool)
func (labelStore *LabelStore) UnpackIsContractNamer(data []byte) (bool, error) {
	out, err := labelStore.abi.Unpack("isContractNamer", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackSetLabel is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbf530969.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setLabel(string label) returns()
func (labelStore *LabelStore) PackSetLabel(label string) []byte {
	enc, err := labelStore.abi.Pack("setLabel", label)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetLabel is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbf530969.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setLabel(string label) returns()
func (labelStore *LabelStore) TryPackSetLabel(label string) ([]byte, error) {
	return labelStore.abi.Pack("setLabel", label)
}

// PackSupportsInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01ffc9a7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (labelStore *LabelStore) PackSupportsInterface(interfaceId [4]byte) []byte {
	enc, err := labelStore.abi.Pack("supportsInterface", interfaceId)
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
func (labelStore *LabelStore) TryPackSupportsInterface(interfaceId [4]byte) ([]byte, error) {
	return labelStore.abi.Pack("supportsInterface", interfaceId)
}

// UnpackSupportsInterface is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (labelStore *LabelStore) UnpackSupportsInterface(data []byte) (bool, error) {
	out, err := labelStore.abi.Unpack("supportsInterface", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// LabelStoreLabel represents a Label event raised by the LabelStore contract.
type LabelStoreLabel struct {
	LabelHash [32]byte
	Label     string
	Raw       *types.Log // Blockchain specific contextual infos
}

const LabelStoreLabelEventName = "Label"

// ContractEventName returns the user-defined event name.
func (LabelStoreLabel) ContractEventName() string {
	return LabelStoreLabelEventName
}

// UnpackLabelEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Label(bytes32 indexed labelHash, string label)
func (labelStore *LabelStore) UnpackLabelEvent(log *types.Log) (*LabelStoreLabel, error) {
	event := "Label"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != labelStore.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(LabelStoreLabel)
	if len(log.Data) > 0 {
		if err := labelStore.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range labelStore.abi.Events[event].Inputs {
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
func (labelStore *LabelStore) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], labelStore.abi.Errors["LabelIsEmpty"].ID.Bytes()[:4]) {
		return labelStore.UnpackLabelIsEmptyError(raw[4:])
	}
	if bytes.Equal(raw[:4], labelStore.abi.Errors["LabelIsTooLong"].ID.Bytes()[:4]) {
		return labelStore.UnpackLabelIsTooLongError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// LabelStoreLabelIsEmpty represents a LabelIsEmpty error raised by the LabelStore contract.
type LabelStoreLabelIsEmpty struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error LabelIsEmpty()
func LabelStoreLabelIsEmptyErrorID() common.Hash {
	return common.HexToHash("0xbf9a27402cc67e0a1553ebafddb9a8f09e2a0e17d33881522b49406d3206b53d")
}

// UnpackLabelIsEmptyError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error LabelIsEmpty()
func (labelStore *LabelStore) UnpackLabelIsEmptyError(raw []byte) (*LabelStoreLabelIsEmpty, error) {
	out := new(LabelStoreLabelIsEmpty)
	if err := labelStore.abi.UnpackIntoInterface(out, "LabelIsEmpty", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// LabelStoreLabelIsTooLong represents a LabelIsTooLong error raised by the LabelStore contract.
type LabelStoreLabelIsTooLong struct {
	Label string
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error LabelIsTooLong(string label)
func LabelStoreLabelIsTooLongErrorID() common.Hash {
	return common.HexToHash("0xdab6c73ccd5b3a9a45fe88ea48179efbe644dbc9416a195d7779ed8ccd248e63")
}

// UnpackLabelIsTooLongError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error LabelIsTooLong(string label)
func (labelStore *LabelStore) UnpackLabelIsTooLongError(raw []byte) (*LabelStoreLabelIsTooLong, error) {
	out := new(LabelStoreLabelIsTooLong)
	if err := labelStore.abi.UnpackIntoInterface(out, "LabelIsTooLong", raw); err != nil {
		return nil, err
	}
	return out, nil
}
