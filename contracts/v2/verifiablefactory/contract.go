// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package verifiablefactory

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

// VerifiableFactoryMetaData contains all meta data concerning the VerifiableFactory contract.
var VerifiableFactoryMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proxy\",\"type\":\"address\"}],\"name\":\"VerificationFailed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"proxyAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ProxyDeployed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"deployProxy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"proxy\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"deployer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"}],\"name\":\"predictProxyAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"proxy\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxyLogic\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proxy\",\"type\":\"address\"}],\"name\":\"verifyContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	ID:  "VerifiableFactory",
}

// VerifiableFactory is an auto generated Go binding around an Ethereum contract.
type VerifiableFactory struct {
	abi abi.ABI
}

// GetABI returns the ABI associated with this contract binding.
func (c *VerifiableFactory) GetABI() abi.ABI {
	return c.abi
}

// NewVerifiableFactory creates a new instance of VerifiableFactory.
func NewVerifiableFactory() *VerifiableFactory {
	parsed, err := VerifiableFactoryMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &VerifiableFactory{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *VerifiableFactory) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackDeployProxy is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5d84121a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function deployProxy(address implementation, uint256 salt, bytes data) returns(address proxy)
func (verifiableFactory *VerifiableFactory) PackDeployProxy(implementation common.Address, salt *big.Int, data []byte) []byte {
	enc, err := verifiableFactory.abi.Pack("deployProxy", implementation, salt, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDeployProxy is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5d84121a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function deployProxy(address implementation, uint256 salt, bytes data) returns(address proxy)
func (verifiableFactory *VerifiableFactory) TryPackDeployProxy(implementation common.Address, salt *big.Int, data []byte) ([]byte, error) {
	return verifiableFactory.abi.Pack("deployProxy", implementation, salt, data)
}

// UnpackDeployProxy is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5d84121a.
//
// Solidity: function deployProxy(address implementation, uint256 salt, bytes data) returns(address proxy)
func (verifiableFactory *VerifiableFactory) UnpackDeployProxy(data []byte) (common.Address, error) {
	out, err := verifiableFactory.abi.Unpack("deployProxy", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackPredictProxyAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x17d35d7b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function predictProxyAddress(address deployer, uint256 salt) view returns(address proxy)
func (verifiableFactory *VerifiableFactory) PackPredictProxyAddress(deployer common.Address, salt *big.Int) []byte {
	enc, err := verifiableFactory.abi.Pack("predictProxyAddress", deployer, salt)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackPredictProxyAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x17d35d7b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function predictProxyAddress(address deployer, uint256 salt) view returns(address proxy)
func (verifiableFactory *VerifiableFactory) TryPackPredictProxyAddress(deployer common.Address, salt *big.Int) ([]byte, error) {
	return verifiableFactory.abi.Pack("predictProxyAddress", deployer, salt)
}

// UnpackPredictProxyAddress is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x17d35d7b.
//
// Solidity: function predictProxyAddress(address deployer, uint256 salt) view returns(address proxy)
func (verifiableFactory *VerifiableFactory) UnpackPredictProxyAddress(data []byte) (common.Address, error) {
	out, err := verifiableFactory.abi.Unpack("predictProxyAddress", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackProxyLogic is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x813845bc.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function proxyLogic() view returns(address)
func (verifiableFactory *VerifiableFactory) PackProxyLogic() []byte {
	enc, err := verifiableFactory.abi.Pack("proxyLogic")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackProxyLogic is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x813845bc.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function proxyLogic() view returns(address)
func (verifiableFactory *VerifiableFactory) TryPackProxyLogic() ([]byte, error) {
	return verifiableFactory.abi.Pack("proxyLogic")
}

// UnpackProxyLogic is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x813845bc.
//
// Solidity: function proxyLogic() view returns(address)
func (verifiableFactory *VerifiableFactory) UnpackProxyLogic(data []byte) (common.Address, error) {
	out, err := verifiableFactory.abi.Unpack("proxyLogic", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackVerifyContract is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3d200b45.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function verifyContract(address proxy) view returns(address implementation)
func (verifiableFactory *VerifiableFactory) PackVerifyContract(proxy common.Address) []byte {
	enc, err := verifiableFactory.abi.Pack("verifyContract", proxy)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackVerifyContract is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3d200b45.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function verifyContract(address proxy) view returns(address implementation)
func (verifiableFactory *VerifiableFactory) TryPackVerifyContract(proxy common.Address) ([]byte, error) {
	return verifiableFactory.abi.Pack("verifyContract", proxy)
}

// UnpackVerifyContract is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3d200b45.
//
// Solidity: function verifyContract(address proxy) view returns(address implementation)
func (verifiableFactory *VerifiableFactory) UnpackVerifyContract(data []byte) (common.Address, error) {
	out, err := verifiableFactory.abi.Unpack("verifyContract", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// VerifiableFactoryProxyDeployed represents a ProxyDeployed event raised by the VerifiableFactory contract.
type VerifiableFactoryProxyDeployed struct {
	Sender         common.Address
	ProxyAddress   common.Address
	Salt           *big.Int
	Implementation common.Address
	Raw            *types.Log // Blockchain specific contextual infos
}

const VerifiableFactoryProxyDeployedEventName = "ProxyDeployed"

// ContractEventName returns the user-defined event name.
func (VerifiableFactoryProxyDeployed) ContractEventName() string {
	return VerifiableFactoryProxyDeployedEventName
}

// UnpackProxyDeployedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ProxyDeployed(address indexed sender, address indexed proxyAddress, uint256 salt, address implementation)
func (verifiableFactory *VerifiableFactory) UnpackProxyDeployedEvent(log *types.Log) (*VerifiableFactoryProxyDeployed, error) {
	event := "ProxyDeployed"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != verifiableFactory.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(VerifiableFactoryProxyDeployed)
	if len(log.Data) > 0 {
		if err := verifiableFactory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range verifiableFactory.abi.Events[event].Inputs {
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
func (verifiableFactory *VerifiableFactory) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], verifiableFactory.abi.Errors["VerificationFailed"].ID.Bytes()[:4]) {
		return verifiableFactory.UnpackVerificationFailedError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// VerifiableFactoryVerificationFailed represents a VerificationFailed error raised by the VerifiableFactory contract.
type VerifiableFactoryVerificationFailed struct {
	Proxy common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error VerificationFailed(address proxy)
func VerifiableFactoryVerificationFailedErrorID() common.Hash {
	return common.HexToHash("0x4c87e2b695a6641e4d63f88e9d449494080344c0c2700faaf708e320b2c2f8bf")
}

// UnpackVerificationFailedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error VerificationFailed(address proxy)
func (verifiableFactory *VerifiableFactory) UnpackVerificationFailedError(raw []byte) (*VerifiableFactoryVerificationFailed, error) {
	out := new(VerifiableFactoryVerificationFailed)
	if err := verifiableFactory.abi.UnpackIntoInterface(out, "VerificationFailed", raw); err != nil {
		return nil, err
	}
	return out, nil
}
