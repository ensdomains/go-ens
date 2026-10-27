// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethregistrar

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

// RenewData is an auto generated low-level Go binding around an user-defined struct.
type RenewData struct {
	Label    string
	Duration uint64
	Referrer [32]byte
}

// ETHRegistrarMetaData contains all meta data concerning the ETHRegistrar contract.
var ETHRegistrarMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"},{\"internalType\":\"contractIPermissionedRegistry\",\"name\":\"ethRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"contractIRentPriceOracle\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"gracePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minCommitmentAge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxCommitmentAge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minRegisterDuration\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"validFrom\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"blockTimestamp\",\"type\":\"uint64\"}],\"name\":\"CommitmentTooNew\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"validTo\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"blockTimestamp\",\"type\":\"uint64\"}],\"name\":\"CommitmentTooOld\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"duration\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minDuration\",\"type\":\"uint64\"}],\"name\":\"DurationTooShort\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxCommitmentAgeTooLow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"}],\"name\":\"NameNotAvailable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"}],\"name\":\"NameNotRenewable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"name\":\"UnexpiredCommitmentExists\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"name\":\"CommitmentMade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIRegistry\",\"name\":\"subregistry\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"duration\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"referrer\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"base\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"premium\",\"type\":\"uint256\"}],\"name\":\"NameRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"duration\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newExpiry\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"referrer\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"NameRenewed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIRentPriceOracle\",\"name\":\"oracle\",\"type\":\"address\"}],\"name\":\"RentPriceOracleUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BENEFICIARY\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETH_REGISTRY\",\"outputs\":[{\"internalType\":\"contractIPermissionedRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GRACE_PERIOD\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_COMMITMENT_AGE\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_COMMITMENT_AGE\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_REGISTER_DURATION\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_RENEW_DURATION\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"name\":\"commit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"name\":\"commitmentAt\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"commitTime\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"duration\",\"type\":\"uint64\"},{\"internalType\":\"contractIERC20\",\"name\":\"paymentToken\",\"type\":\"address\"}],\"name\":\"getRegisterPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"base\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"premium\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"}],\"name\":\"getRemainingGracePeriod\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"duration\",\"type\":\"uint64\"},{\"internalType\":\"contractIERC20\",\"name\":\"paymentToken\",\"type\":\"address\"}],\"name\":\"getRenewPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"}],\"name\":\"isAvailable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"}],\"name\":\"isRenewable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"secret\",\"type\":\"bytes32\"},{\"internalType\":\"contractIRegistry\",\"name\":\"subregistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"duration\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"referrer\",\"type\":\"bytes32\"}],\"name\":\"makeCommitment\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"secret\",\"type\":\"bytes32\"},{\"internalType\":\"contractIRegistry\",\"name\":\"subregistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"duration\",\"type\":\"uint64\"},{\"internalType\":\"contractIERC20\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"referrer\",\"type\":\"bytes32\"}],\"name\":\"register\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"duration\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"referrer\",\"type\":\"bytes32\"}],\"internalType\":\"structRenewData\",\"name\":\"rd\",\"type\":\"tuple\"},{\"internalType\":\"contractIERC20\",\"name\":\"paymentToken\",\"type\":\"address\"}],\"name\":\"renew\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"duration\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"referrer\",\"type\":\"bytes32\"}],\"internalType\":\"structRenewData[]\",\"name\":\"rds\",\"type\":\"tuple[]\"},{\"internalType\":\"contractIERC20\",\"name\":\"paymentToken\",\"type\":\"address\"}],\"name\":\"renewBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rentPriceOracle\",\"outputs\":[{\"internalType\":\"contractIRentPriceOracle\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIRentPriceOracle\",\"name\":\"oracle\",\"type\":\"address\"}],\"name\":\"setRentPriceOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "ETHRegistrar",
}

// ETHRegistrar is an auto generated Go binding around an Ethereum contract.
type ETHRegistrar struct {
	abi abi.ABI
}

// GetABI returns the ABI associated with this contract binding.
func (c *ETHRegistrar) GetABI() abi.ABI {
	return c.abi
}

// NewETHRegistrar creates a new instance of ETHRegistrar.
func NewETHRegistrar() *ETHRegistrar {
	parsed, err := ETHRegistrarMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &ETHRegistrar{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *ETHRegistrar) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address owner_, address ethRegistry, address beneficiary, address oracle, uint64 gracePeriod, uint64 minCommitmentAge, uint64 maxCommitmentAge, uint64 minRegisterDuration) returns()
func (eTHRegistrar *ETHRegistrar) PackConstructor(owner_ common.Address, ethRegistry common.Address, beneficiary common.Address, oracle common.Address, gracePeriod uint64, minCommitmentAge uint64, maxCommitmentAge uint64, minRegisterDuration uint64) []byte {
	enc, err := eTHRegistrar.abi.Pack("", owner_, ethRegistry, beneficiary, oracle, gracePeriod, minCommitmentAge, maxCommitmentAge, minRegisterDuration)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackBENEFICIARY is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2f99c6cc.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function BENEFICIARY() view returns(address)
func (eTHRegistrar *ETHRegistrar) PackBENEFICIARY() []byte {
	enc, err := eTHRegistrar.abi.Pack("BENEFICIARY")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackBENEFICIARY is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2f99c6cc.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function BENEFICIARY() view returns(address)
func (eTHRegistrar *ETHRegistrar) TryPackBENEFICIARY() ([]byte, error) {
	return eTHRegistrar.abi.Pack("BENEFICIARY")
}

// UnpackBENEFICIARY is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x2f99c6cc.
//
// Solidity: function BENEFICIARY() view returns(address)
func (eTHRegistrar *ETHRegistrar) UnpackBENEFICIARY(data []byte) (common.Address, error) {
	out, err := eTHRegistrar.abi.Unpack("BENEFICIARY", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackETHREGISTRY is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x47500708.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function ETH_REGISTRY() view returns(address)
func (eTHRegistrar *ETHRegistrar) PackETHREGISTRY() []byte {
	enc, err := eTHRegistrar.abi.Pack("ETH_REGISTRY")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackETHREGISTRY is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x47500708.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function ETH_REGISTRY() view returns(address)
func (eTHRegistrar *ETHRegistrar) TryPackETHREGISTRY() ([]byte, error) {
	return eTHRegistrar.abi.Pack("ETH_REGISTRY")
}

// UnpackETHREGISTRY is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x47500708.
//
// Solidity: function ETH_REGISTRY() view returns(address)
func (eTHRegistrar *ETHRegistrar) UnpackETHREGISTRY(data []byte) (common.Address, error) {
	out, err := eTHRegistrar.abi.Unpack("ETH_REGISTRY", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGRACEPERIOD is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc1a287e2.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function GRACE_PERIOD() view returns(uint64)
func (eTHRegistrar *ETHRegistrar) PackGRACEPERIOD() []byte {
	enc, err := eTHRegistrar.abi.Pack("GRACE_PERIOD")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGRACEPERIOD is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc1a287e2.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function GRACE_PERIOD() view returns(uint64)
func (eTHRegistrar *ETHRegistrar) TryPackGRACEPERIOD() ([]byte, error) {
	return eTHRegistrar.abi.Pack("GRACE_PERIOD")
}

// UnpackGRACEPERIOD is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc1a287e2.
//
// Solidity: function GRACE_PERIOD() view returns(uint64)
func (eTHRegistrar *ETHRegistrar) UnpackGRACEPERIOD(data []byte) (uint64, error) {
	out, err := eTHRegistrar.abi.Unpack("GRACE_PERIOD", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, nil
}

// PackMAXCOMMITMENTAGE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8ccb9ea6.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function MAX_COMMITMENT_AGE() view returns(uint64)
func (eTHRegistrar *ETHRegistrar) PackMAXCOMMITMENTAGE() []byte {
	enc, err := eTHRegistrar.abi.Pack("MAX_COMMITMENT_AGE")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackMAXCOMMITMENTAGE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8ccb9ea6.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function MAX_COMMITMENT_AGE() view returns(uint64)
func (eTHRegistrar *ETHRegistrar) TryPackMAXCOMMITMENTAGE() ([]byte, error) {
	return eTHRegistrar.abi.Pack("MAX_COMMITMENT_AGE")
}

// UnpackMAXCOMMITMENTAGE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8ccb9ea6.
//
// Solidity: function MAX_COMMITMENT_AGE() view returns(uint64)
func (eTHRegistrar *ETHRegistrar) UnpackMAXCOMMITMENTAGE(data []byte) (uint64, error) {
	out, err := eTHRegistrar.abi.Unpack("MAX_COMMITMENT_AGE", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, nil
}

// PackMINCOMMITMENTAGE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2e4f692a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function MIN_COMMITMENT_AGE() view returns(uint64)
func (eTHRegistrar *ETHRegistrar) PackMINCOMMITMENTAGE() []byte {
	enc, err := eTHRegistrar.abi.Pack("MIN_COMMITMENT_AGE")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackMINCOMMITMENTAGE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2e4f692a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function MIN_COMMITMENT_AGE() view returns(uint64)
func (eTHRegistrar *ETHRegistrar) TryPackMINCOMMITMENTAGE() ([]byte, error) {
	return eTHRegistrar.abi.Pack("MIN_COMMITMENT_AGE")
}

// UnpackMINCOMMITMENTAGE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x2e4f692a.
//
// Solidity: function MIN_COMMITMENT_AGE() view returns(uint64)
func (eTHRegistrar *ETHRegistrar) UnpackMINCOMMITMENTAGE(data []byte) (uint64, error) {
	out, err := eTHRegistrar.abi.Unpack("MIN_COMMITMENT_AGE", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, nil
}

// PackMINREGISTERDURATION is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa40938bd.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function MIN_REGISTER_DURATION() view returns(uint64)
func (eTHRegistrar *ETHRegistrar) PackMINREGISTERDURATION() []byte {
	enc, err := eTHRegistrar.abi.Pack("MIN_REGISTER_DURATION")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackMINREGISTERDURATION is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa40938bd.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function MIN_REGISTER_DURATION() view returns(uint64)
func (eTHRegistrar *ETHRegistrar) TryPackMINREGISTERDURATION() ([]byte, error) {
	return eTHRegistrar.abi.Pack("MIN_REGISTER_DURATION")
}

// UnpackMINREGISTERDURATION is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa40938bd.
//
// Solidity: function MIN_REGISTER_DURATION() view returns(uint64)
func (eTHRegistrar *ETHRegistrar) UnpackMINREGISTERDURATION(data []byte) (uint64, error) {
	out, err := eTHRegistrar.abi.Unpack("MIN_REGISTER_DURATION", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, nil
}

// PackMINRENEWDURATION is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x307a64a5.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function MIN_RENEW_DURATION() view returns(uint64)
func (eTHRegistrar *ETHRegistrar) PackMINRENEWDURATION() []byte {
	enc, err := eTHRegistrar.abi.Pack("MIN_RENEW_DURATION")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackMINRENEWDURATION is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x307a64a5.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function MIN_RENEW_DURATION() view returns(uint64)
func (eTHRegistrar *ETHRegistrar) TryPackMINRENEWDURATION() ([]byte, error) {
	return eTHRegistrar.abi.Pack("MIN_RENEW_DURATION")
}

// UnpackMINRENEWDURATION is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x307a64a5.
//
// Solidity: function MIN_RENEW_DURATION() view returns(uint64)
func (eTHRegistrar *ETHRegistrar) UnpackMINRENEWDURATION(data []byte) (uint64, error) {
	out, err := eTHRegistrar.abi.Unpack("MIN_RENEW_DURATION", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, nil
}

// PackCommit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf14fcbc8.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function commit(bytes32 commitment) returns()
func (eTHRegistrar *ETHRegistrar) PackCommit(commitment [32]byte) []byte {
	enc, err := eTHRegistrar.abi.Pack("commit", commitment)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCommit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf14fcbc8.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function commit(bytes32 commitment) returns()
func (eTHRegistrar *ETHRegistrar) TryPackCommit(commitment [32]byte) ([]byte, error) {
	return eTHRegistrar.abi.Pack("commit", commitment)
}

// PackCommitmentAt is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x16a92535.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function commitmentAt(bytes32 commitment) view returns(uint64 commitTime)
func (eTHRegistrar *ETHRegistrar) PackCommitmentAt(commitment [32]byte) []byte {
	enc, err := eTHRegistrar.abi.Pack("commitmentAt", commitment)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCommitmentAt is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x16a92535.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function commitmentAt(bytes32 commitment) view returns(uint64 commitTime)
func (eTHRegistrar *ETHRegistrar) TryPackCommitmentAt(commitment [32]byte) ([]byte, error) {
	return eTHRegistrar.abi.Pack("commitmentAt", commitment)
}

// UnpackCommitmentAt is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x16a92535.
//
// Solidity: function commitmentAt(bytes32 commitment) view returns(uint64 commitTime)
func (eTHRegistrar *ETHRegistrar) UnpackCommitmentAt(data []byte) (uint64, error) {
	out, err := eTHRegistrar.abi.Unpack("commitmentAt", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, nil
}

// PackGetRegisterPrice is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x61907b12.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getRegisterPrice(string label, uint64 duration, address paymentToken) view returns(uint256 base, uint256 premium)
func (eTHRegistrar *ETHRegistrar) PackGetRegisterPrice(label string, duration uint64, paymentToken common.Address) []byte {
	enc, err := eTHRegistrar.abi.Pack("getRegisterPrice", label, duration, paymentToken)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetRegisterPrice is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x61907b12.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getRegisterPrice(string label, uint64 duration, address paymentToken) view returns(uint256 base, uint256 premium)
func (eTHRegistrar *ETHRegistrar) TryPackGetRegisterPrice(label string, duration uint64, paymentToken common.Address) ([]byte, error) {
	return eTHRegistrar.abi.Pack("getRegisterPrice", label, duration, paymentToken)
}

// GetRegisterPriceOutput serves as a container for the return parameters of contract
// method GetRegisterPrice.
type GetRegisterPriceOutput struct {
	Base    *big.Int
	Premium *big.Int
}

// UnpackGetRegisterPrice is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x61907b12.
//
// Solidity: function getRegisterPrice(string label, uint64 duration, address paymentToken) view returns(uint256 base, uint256 premium)
func (eTHRegistrar *ETHRegistrar) UnpackGetRegisterPrice(data []byte) (GetRegisterPriceOutput, error) {
	out, err := eTHRegistrar.abi.Unpack("getRegisterPrice", data)
	outstruct := new(GetRegisterPriceOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Base = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.Premium = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	return *outstruct, nil
}

// PackGetRemainingGracePeriod is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x802295ef.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getRemainingGracePeriod(string label) view returns(uint64)
func (eTHRegistrar *ETHRegistrar) PackGetRemainingGracePeriod(label string) []byte {
	enc, err := eTHRegistrar.abi.Pack("getRemainingGracePeriod", label)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetRemainingGracePeriod is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x802295ef.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getRemainingGracePeriod(string label) view returns(uint64)
func (eTHRegistrar *ETHRegistrar) TryPackGetRemainingGracePeriod(label string) ([]byte, error) {
	return eTHRegistrar.abi.Pack("getRemainingGracePeriod", label)
}

// UnpackGetRemainingGracePeriod is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x802295ef.
//
// Solidity: function getRemainingGracePeriod(string label) view returns(uint64)
func (eTHRegistrar *ETHRegistrar) UnpackGetRemainingGracePeriod(data []byte) (uint64, error) {
	out, err := eTHRegistrar.abi.Unpack("getRemainingGracePeriod", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, nil
}

// PackGetRenewPrice is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xddf0effc.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getRenewPrice(string label, uint64 duration, address paymentToken) view returns(uint256)
func (eTHRegistrar *ETHRegistrar) PackGetRenewPrice(label string, duration uint64, paymentToken common.Address) []byte {
	enc, err := eTHRegistrar.abi.Pack("getRenewPrice", label, duration, paymentToken)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetRenewPrice is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xddf0effc.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getRenewPrice(string label, uint64 duration, address paymentToken) view returns(uint256)
func (eTHRegistrar *ETHRegistrar) TryPackGetRenewPrice(label string, duration uint64, paymentToken common.Address) ([]byte, error) {
	return eTHRegistrar.abi.Pack("getRenewPrice", label, duration, paymentToken)
}

// UnpackGetRenewPrice is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xddf0effc.
//
// Solidity: function getRenewPrice(string label, uint64 duration, address paymentToken) view returns(uint256)
func (eTHRegistrar *ETHRegistrar) UnpackGetRenewPrice(data []byte) (*big.Int, error) {
	out, err := eTHRegistrar.abi.Unpack("getRenewPrice", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackIsAvailable is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x965306aa.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isAvailable(string label) view returns(bool)
func (eTHRegistrar *ETHRegistrar) PackIsAvailable(label string) []byte {
	enc, err := eTHRegistrar.abi.Pack("isAvailable", label)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsAvailable is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x965306aa.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isAvailable(string label) view returns(bool)
func (eTHRegistrar *ETHRegistrar) TryPackIsAvailable(label string) ([]byte, error) {
	return eTHRegistrar.abi.Pack("isAvailable", label)
}

// UnpackIsAvailable is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x965306aa.
//
// Solidity: function isAvailable(string label) view returns(bool)
func (eTHRegistrar *ETHRegistrar) UnpackIsAvailable(data []byte) (bool, error) {
	out, err := eTHRegistrar.abi.Unpack("isAvailable", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackIsRenewable is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x130d6f00.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isRenewable(string label) view returns(bool)
func (eTHRegistrar *ETHRegistrar) PackIsRenewable(label string) []byte {
	enc, err := eTHRegistrar.abi.Pack("isRenewable", label)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsRenewable is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x130d6f00.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isRenewable(string label) view returns(bool)
func (eTHRegistrar *ETHRegistrar) TryPackIsRenewable(label string) ([]byte, error) {
	return eTHRegistrar.abi.Pack("isRenewable", label)
}

// UnpackIsRenewable is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x130d6f00.
//
// Solidity: function isRenewable(string label) view returns(bool)
func (eTHRegistrar *ETHRegistrar) UnpackIsRenewable(data []byte) (bool, error) {
	out, err := eTHRegistrar.abi.Unpack("isRenewable", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackMakeCommitment is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1e966f07.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function makeCommitment(string label, address owner, bytes32 secret, address subregistry, address resolver, uint64 duration, bytes32 referrer) pure returns(bytes32)
func (eTHRegistrar *ETHRegistrar) PackMakeCommitment(label string, owner common.Address, secret [32]byte, subregistry common.Address, resolver common.Address, duration uint64, referrer [32]byte) []byte {
	enc, err := eTHRegistrar.abi.Pack("makeCommitment", label, owner, secret, subregistry, resolver, duration, referrer)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackMakeCommitment is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1e966f07.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function makeCommitment(string label, address owner, bytes32 secret, address subregistry, address resolver, uint64 duration, bytes32 referrer) pure returns(bytes32)
func (eTHRegistrar *ETHRegistrar) TryPackMakeCommitment(label string, owner common.Address, secret [32]byte, subregistry common.Address, resolver common.Address, duration uint64, referrer [32]byte) ([]byte, error) {
	return eTHRegistrar.abi.Pack("makeCommitment", label, owner, secret, subregistry, resolver, duration, referrer)
}

// UnpackMakeCommitment is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x1e966f07.
//
// Solidity: function makeCommitment(string label, address owner, bytes32 secret, address subregistry, address resolver, uint64 duration, bytes32 referrer) pure returns(bytes32)
func (eTHRegistrar *ETHRegistrar) UnpackMakeCommitment(data []byte) ([32]byte, error) {
	out, err := eTHRegistrar.abi.Unpack("makeCommitment", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8da5cb5b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function owner() view returns(address)
func (eTHRegistrar *ETHRegistrar) PackOwner() []byte {
	enc, err := eTHRegistrar.abi.Pack("owner")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8da5cb5b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function owner() view returns(address)
func (eTHRegistrar *ETHRegistrar) TryPackOwner() ([]byte, error) {
	return eTHRegistrar.abi.Pack("owner")
}

// UnpackOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (eTHRegistrar *ETHRegistrar) UnpackOwner(data []byte) (common.Address, error) {
	out, err := eTHRegistrar.abi.Unpack("owner", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackRegister is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xcff3e7c2.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function register(string label, address owner, bytes32 secret, address subregistry, address resolver, uint64 duration, address paymentToken, bytes32 referrer) returns(uint256 tokenId)
func (eTHRegistrar *ETHRegistrar) PackRegister(label string, owner common.Address, secret [32]byte, subregistry common.Address, resolver common.Address, duration uint64, paymentToken common.Address, referrer [32]byte) []byte {
	enc, err := eTHRegistrar.abi.Pack("register", label, owner, secret, subregistry, resolver, duration, paymentToken, referrer)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRegister is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xcff3e7c2.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function register(string label, address owner, bytes32 secret, address subregistry, address resolver, uint64 duration, address paymentToken, bytes32 referrer) returns(uint256 tokenId)
func (eTHRegistrar *ETHRegistrar) TryPackRegister(label string, owner common.Address, secret [32]byte, subregistry common.Address, resolver common.Address, duration uint64, paymentToken common.Address, referrer [32]byte) ([]byte, error) {
	return eTHRegistrar.abi.Pack("register", label, owner, secret, subregistry, resolver, duration, paymentToken, referrer)
}

// UnpackRegister is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xcff3e7c2.
//
// Solidity: function register(string label, address owner, bytes32 secret, address subregistry, address resolver, uint64 duration, address paymentToken, bytes32 referrer) returns(uint256 tokenId)
func (eTHRegistrar *ETHRegistrar) UnpackRegister(data []byte) (*big.Int, error) {
	out, err := eTHRegistrar.abi.Unpack("register", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackRenew is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2d06649b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function renew((string,uint64,bytes32) rd, address paymentToken) returns()
func (eTHRegistrar *ETHRegistrar) PackRenew(rd RenewData, paymentToken common.Address) []byte {
	enc, err := eTHRegistrar.abi.Pack("renew", rd, paymentToken)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRenew is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2d06649b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function renew((string,uint64,bytes32) rd, address paymentToken) returns()
func (eTHRegistrar *ETHRegistrar) TryPackRenew(rd RenewData, paymentToken common.Address) ([]byte, error) {
	return eTHRegistrar.abi.Pack("renew", rd, paymentToken)
}

// PackRenewBatch is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x959d530d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function renewBatch((string,uint64,bytes32)[] rds, address paymentToken) returns()
func (eTHRegistrar *ETHRegistrar) PackRenewBatch(rds []RenewData, paymentToken common.Address) []byte {
	enc, err := eTHRegistrar.abi.Pack("renewBatch", rds, paymentToken)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRenewBatch is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x959d530d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function renewBatch((string,uint64,bytes32)[] rds, address paymentToken) returns()
func (eTHRegistrar *ETHRegistrar) TryPackRenewBatch(rds []RenewData, paymentToken common.Address) ([]byte, error) {
	return eTHRegistrar.abi.Pack("renewBatch", rds, paymentToken)
}

// PackRenounceOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x715018a6.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function renounceOwnership() returns()
func (eTHRegistrar *ETHRegistrar) PackRenounceOwnership() []byte {
	enc, err := eTHRegistrar.abi.Pack("renounceOwnership")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRenounceOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x715018a6.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function renounceOwnership() returns()
func (eTHRegistrar *ETHRegistrar) TryPackRenounceOwnership() ([]byte, error) {
	return eTHRegistrar.abi.Pack("renounceOwnership")
}

// PackRentPriceOracle is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7b39ba16.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function rentPriceOracle() view returns(address)
func (eTHRegistrar *ETHRegistrar) PackRentPriceOracle() []byte {
	enc, err := eTHRegistrar.abi.Pack("rentPriceOracle")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRentPriceOracle is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7b39ba16.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function rentPriceOracle() view returns(address)
func (eTHRegistrar *ETHRegistrar) TryPackRentPriceOracle() ([]byte, error) {
	return eTHRegistrar.abi.Pack("rentPriceOracle")
}

// UnpackRentPriceOracle is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7b39ba16.
//
// Solidity: function rentPriceOracle() view returns(address)
func (eTHRegistrar *ETHRegistrar) UnpackRentPriceOracle(data []byte) (common.Address, error) {
	out, err := eTHRegistrar.abi.Unpack("rentPriceOracle", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackSetRentPriceOracle is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa2a11fbe.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setRentPriceOracle(address oracle) returns()
func (eTHRegistrar *ETHRegistrar) PackSetRentPriceOracle(oracle common.Address) []byte {
	enc, err := eTHRegistrar.abi.Pack("setRentPriceOracle", oracle)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetRentPriceOracle is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa2a11fbe.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setRentPriceOracle(address oracle) returns()
func (eTHRegistrar *ETHRegistrar) TryPackSetRentPriceOracle(oracle common.Address) ([]byte, error) {
	return eTHRegistrar.abi.Pack("setRentPriceOracle", oracle)
}

// PackSupportsInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01ffc9a7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (eTHRegistrar *ETHRegistrar) PackSupportsInterface(interfaceId [4]byte) []byte {
	enc, err := eTHRegistrar.abi.Pack("supportsInterface", interfaceId)
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
func (eTHRegistrar *ETHRegistrar) TryPackSupportsInterface(interfaceId [4]byte) ([]byte, error) {
	return eTHRegistrar.abi.Pack("supportsInterface", interfaceId)
}

// UnpackSupportsInterface is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (eTHRegistrar *ETHRegistrar) UnpackSupportsInterface(data []byte) (bool, error) {
	out, err := eTHRegistrar.abi.Unpack("supportsInterface", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackTransferOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf2fde38b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (eTHRegistrar *ETHRegistrar) PackTransferOwnership(newOwner common.Address) []byte {
	enc, err := eTHRegistrar.abi.Pack("transferOwnership", newOwner)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTransferOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf2fde38b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (eTHRegistrar *ETHRegistrar) TryPackTransferOwnership(newOwner common.Address) ([]byte, error) {
	return eTHRegistrar.abi.Pack("transferOwnership", newOwner)
}

// ETHRegistrarCommitmentMade represents a CommitmentMade event raised by the ETHRegistrar contract.
type ETHRegistrarCommitmentMade struct {
	Commitment [32]byte
	Raw        *types.Log // Blockchain specific contextual infos
}

const ETHRegistrarCommitmentMadeEventName = "CommitmentMade"

// ContractEventName returns the user-defined event name.
func (ETHRegistrarCommitmentMade) ContractEventName() string {
	return ETHRegistrarCommitmentMadeEventName
}

// UnpackCommitmentMadeEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event CommitmentMade(bytes32 commitment)
func (eTHRegistrar *ETHRegistrar) UnpackCommitmentMadeEvent(log *types.Log) (*ETHRegistrarCommitmentMade, error) {
	event := "CommitmentMade"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != eTHRegistrar.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(ETHRegistrarCommitmentMade)
	if len(log.Data) > 0 {
		if err := eTHRegistrar.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eTHRegistrar.abi.Events[event].Inputs {
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

// ETHRegistrarNameRegistered represents a NameRegistered event raised by the ETHRegistrar contract.
type ETHRegistrarNameRegistered struct {
	TokenId      *big.Int
	Label        string
	Owner        common.Address
	Subregistry  common.Address
	Resolver     common.Address
	Duration     uint64
	PaymentToken common.Address
	Referrer     [32]byte
	Base         *big.Int
	Premium      *big.Int
	Raw          *types.Log // Blockchain specific contextual infos
}

const ETHRegistrarNameRegisteredEventName = "NameRegistered"

// ContractEventName returns the user-defined event name.
func (ETHRegistrarNameRegistered) ContractEventName() string {
	return ETHRegistrarNameRegisteredEventName
}

// UnpackNameRegisteredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event NameRegistered(uint256 indexed tokenId, string label, address owner, address subregistry, address resolver, uint64 duration, address paymentToken, bytes32 indexed referrer, uint256 base, uint256 premium)
func (eTHRegistrar *ETHRegistrar) UnpackNameRegisteredEvent(log *types.Log) (*ETHRegistrarNameRegistered, error) {
	event := "NameRegistered"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != eTHRegistrar.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(ETHRegistrarNameRegistered)
	if len(log.Data) > 0 {
		if err := eTHRegistrar.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eTHRegistrar.abi.Events[event].Inputs {
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

// ETHRegistrarNameRenewed represents a NameRenewed event raised by the ETHRegistrar contract.
type ETHRegistrarNameRenewed struct {
	TokenId      *big.Int
	Label        string
	Duration     uint64
	NewExpiry    uint64
	PaymentToken common.Address
	Referrer     [32]byte
	Amount       *big.Int
	Raw          *types.Log // Blockchain specific contextual infos
}

const ETHRegistrarNameRenewedEventName = "NameRenewed"

// ContractEventName returns the user-defined event name.
func (ETHRegistrarNameRenewed) ContractEventName() string {
	return ETHRegistrarNameRenewedEventName
}

// UnpackNameRenewedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event NameRenewed(uint256 indexed tokenId, string label, uint64 duration, uint64 newExpiry, address paymentToken, bytes32 indexed referrer, uint256 amount)
func (eTHRegistrar *ETHRegistrar) UnpackNameRenewedEvent(log *types.Log) (*ETHRegistrarNameRenewed, error) {
	event := "NameRenewed"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != eTHRegistrar.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(ETHRegistrarNameRenewed)
	if len(log.Data) > 0 {
		if err := eTHRegistrar.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eTHRegistrar.abi.Events[event].Inputs {
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

// ETHRegistrarOwnershipTransferred represents a OwnershipTransferred event raised by the ETHRegistrar contract.
type ETHRegistrarOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           *types.Log // Blockchain specific contextual infos
}

const ETHRegistrarOwnershipTransferredEventName = "OwnershipTransferred"

// ContractEventName returns the user-defined event name.
func (ETHRegistrarOwnershipTransferred) ContractEventName() string {
	return ETHRegistrarOwnershipTransferredEventName
}

// UnpackOwnershipTransferredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (eTHRegistrar *ETHRegistrar) UnpackOwnershipTransferredEvent(log *types.Log) (*ETHRegistrarOwnershipTransferred, error) {
	event := "OwnershipTransferred"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != eTHRegistrar.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(ETHRegistrarOwnershipTransferred)
	if len(log.Data) > 0 {
		if err := eTHRegistrar.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eTHRegistrar.abi.Events[event].Inputs {
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

// ETHRegistrarRentPriceOracleUpdated represents a RentPriceOracleUpdated event raised by the ETHRegistrar contract.
type ETHRegistrarRentPriceOracleUpdated struct {
	Oracle common.Address
	Raw    *types.Log // Blockchain specific contextual infos
}

const ETHRegistrarRentPriceOracleUpdatedEventName = "RentPriceOracleUpdated"

// ContractEventName returns the user-defined event name.
func (ETHRegistrarRentPriceOracleUpdated) ContractEventName() string {
	return ETHRegistrarRentPriceOracleUpdatedEventName
}

// UnpackRentPriceOracleUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RentPriceOracleUpdated(address oracle)
func (eTHRegistrar *ETHRegistrar) UnpackRentPriceOracleUpdatedEvent(log *types.Log) (*ETHRegistrarRentPriceOracleUpdated, error) {
	event := "RentPriceOracleUpdated"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != eTHRegistrar.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(ETHRegistrarRentPriceOracleUpdated)
	if len(log.Data) > 0 {
		if err := eTHRegistrar.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eTHRegistrar.abi.Events[event].Inputs {
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
func (eTHRegistrar *ETHRegistrar) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], eTHRegistrar.abi.Errors["CommitmentTooNew"].ID.Bytes()[:4]) {
		return eTHRegistrar.UnpackCommitmentTooNewError(raw[4:])
	}
	if bytes.Equal(raw[:4], eTHRegistrar.abi.Errors["CommitmentTooOld"].ID.Bytes()[:4]) {
		return eTHRegistrar.UnpackCommitmentTooOldError(raw[4:])
	}
	if bytes.Equal(raw[:4], eTHRegistrar.abi.Errors["DurationTooShort"].ID.Bytes()[:4]) {
		return eTHRegistrar.UnpackDurationTooShortError(raw[4:])
	}
	if bytes.Equal(raw[:4], eTHRegistrar.abi.Errors["InvalidOwner"].ID.Bytes()[:4]) {
		return eTHRegistrar.UnpackInvalidOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], eTHRegistrar.abi.Errors["MaxCommitmentAgeTooLow"].ID.Bytes()[:4]) {
		return eTHRegistrar.UnpackMaxCommitmentAgeTooLowError(raw[4:])
	}
	if bytes.Equal(raw[:4], eTHRegistrar.abi.Errors["NameNotAvailable"].ID.Bytes()[:4]) {
		return eTHRegistrar.UnpackNameNotAvailableError(raw[4:])
	}
	if bytes.Equal(raw[:4], eTHRegistrar.abi.Errors["NameNotRenewable"].ID.Bytes()[:4]) {
		return eTHRegistrar.UnpackNameNotRenewableError(raw[4:])
	}
	if bytes.Equal(raw[:4], eTHRegistrar.abi.Errors["OwnableInvalidOwner"].ID.Bytes()[:4]) {
		return eTHRegistrar.UnpackOwnableInvalidOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], eTHRegistrar.abi.Errors["OwnableUnauthorizedAccount"].ID.Bytes()[:4]) {
		return eTHRegistrar.UnpackOwnableUnauthorizedAccountError(raw[4:])
	}
	if bytes.Equal(raw[:4], eTHRegistrar.abi.Errors["SafeERC20FailedOperation"].ID.Bytes()[:4]) {
		return eTHRegistrar.UnpackSafeERC20FailedOperationError(raw[4:])
	}
	if bytes.Equal(raw[:4], eTHRegistrar.abi.Errors["UnexpiredCommitmentExists"].ID.Bytes()[:4]) {
		return eTHRegistrar.UnpackUnexpiredCommitmentExistsError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// ETHRegistrarCommitmentTooNew represents a CommitmentTooNew error raised by the ETHRegistrar contract.
type ETHRegistrarCommitmentTooNew struct {
	Commitment     [32]byte
	ValidFrom      uint64
	BlockTimestamp uint64
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error CommitmentTooNew(bytes32 commitment, uint64 validFrom, uint64 blockTimestamp)
func ETHRegistrarCommitmentTooNewErrorID() common.Hash {
	return common.HexToHash("0x6be614e365c2739141d32186ed9721e572e63acab0d27ca700bc28900ac94f1a")
}

// UnpackCommitmentTooNewError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error CommitmentTooNew(bytes32 commitment, uint64 validFrom, uint64 blockTimestamp)
func (eTHRegistrar *ETHRegistrar) UnpackCommitmentTooNewError(raw []byte) (*ETHRegistrarCommitmentTooNew, error) {
	out := new(ETHRegistrarCommitmentTooNew)
	if err := eTHRegistrar.abi.UnpackIntoInterface(out, "CommitmentTooNew", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ETHRegistrarCommitmentTooOld represents a CommitmentTooOld error raised by the ETHRegistrar contract.
type ETHRegistrarCommitmentTooOld struct {
	Commitment     [32]byte
	ValidTo        uint64
	BlockTimestamp uint64
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error CommitmentTooOld(bytes32 commitment, uint64 validTo, uint64 blockTimestamp)
func ETHRegistrarCommitmentTooOldErrorID() common.Hash {
	return common.HexToHash("0x0cb9df3fe11b7e41b0d3a44b31725ab337d14d1d4a6a4c74895687b6b47a9c3f")
}

// UnpackCommitmentTooOldError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error CommitmentTooOld(bytes32 commitment, uint64 validTo, uint64 blockTimestamp)
func (eTHRegistrar *ETHRegistrar) UnpackCommitmentTooOldError(raw []byte) (*ETHRegistrarCommitmentTooOld, error) {
	out := new(ETHRegistrarCommitmentTooOld)
	if err := eTHRegistrar.abi.UnpackIntoInterface(out, "CommitmentTooOld", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ETHRegistrarDurationTooShort represents a DurationTooShort error raised by the ETHRegistrar contract.
type ETHRegistrarDurationTooShort struct {
	Duration    uint64
	MinDuration uint64
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error DurationTooShort(uint64 duration, uint64 minDuration)
func ETHRegistrarDurationTooShortErrorID() common.Hash {
	return common.HexToHash("0xa096b844523daa5d1bb83afd97baa99f239f9b57ef2a99fb5e67d6abf00e360b")
}

// UnpackDurationTooShortError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error DurationTooShort(uint64 duration, uint64 minDuration)
func (eTHRegistrar *ETHRegistrar) UnpackDurationTooShortError(raw []byte) (*ETHRegistrarDurationTooShort, error) {
	out := new(ETHRegistrarDurationTooShort)
	if err := eTHRegistrar.abi.UnpackIntoInterface(out, "DurationTooShort", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ETHRegistrarInvalidOwner represents a InvalidOwner error raised by the ETHRegistrar contract.
type ETHRegistrarInvalidOwner struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidOwner()
func ETHRegistrarInvalidOwnerErrorID() common.Hash {
	return common.HexToHash("0x49e27cffb37b1ca4a9bf5318243b2014d13f940af232b8552c208bdea15739da")
}

// UnpackInvalidOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidOwner()
func (eTHRegistrar *ETHRegistrar) UnpackInvalidOwnerError(raw []byte) (*ETHRegistrarInvalidOwner, error) {
	out := new(ETHRegistrarInvalidOwner)
	if err := eTHRegistrar.abi.UnpackIntoInterface(out, "InvalidOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ETHRegistrarMaxCommitmentAgeTooLow represents a MaxCommitmentAgeTooLow error raised by the ETHRegistrar contract.
type ETHRegistrarMaxCommitmentAgeTooLow struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error MaxCommitmentAgeTooLow()
func ETHRegistrarMaxCommitmentAgeTooLowErrorID() common.Hash {
	return common.HexToHash("0x3e5aa838b636c9cff9cb272fdf56fe3bdca93a67dcfd6adc1f1f04e2a41a9531")
}

// UnpackMaxCommitmentAgeTooLowError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error MaxCommitmentAgeTooLow()
func (eTHRegistrar *ETHRegistrar) UnpackMaxCommitmentAgeTooLowError(raw []byte) (*ETHRegistrarMaxCommitmentAgeTooLow, error) {
	out := new(ETHRegistrarMaxCommitmentAgeTooLow)
	if err := eTHRegistrar.abi.UnpackIntoInterface(out, "MaxCommitmentAgeTooLow", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ETHRegistrarNameNotAvailable represents a NameNotAvailable error raised by the ETHRegistrar contract.
type ETHRegistrarNameNotAvailable struct {
	Label string
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NameNotAvailable(string label)
func ETHRegistrarNameNotAvailableErrorID() common.Hash {
	return common.HexToHash("0x477707e82010494a7688132ab8503d9aa6b224a0fa45c45c7ca9f122c67e3287")
}

// UnpackNameNotAvailableError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NameNotAvailable(string label)
func (eTHRegistrar *ETHRegistrar) UnpackNameNotAvailableError(raw []byte) (*ETHRegistrarNameNotAvailable, error) {
	out := new(ETHRegistrarNameNotAvailable)
	if err := eTHRegistrar.abi.UnpackIntoInterface(out, "NameNotAvailable", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ETHRegistrarNameNotRenewable represents a NameNotRenewable error raised by the ETHRegistrar contract.
type ETHRegistrarNameNotRenewable struct {
	Label string
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NameNotRenewable(string label)
func ETHRegistrarNameNotRenewableErrorID() common.Hash {
	return common.HexToHash("0x1caefaa0709e9c51b7dc0085c2af73ef0ee280cad1968e0631809301c96a86e7")
}

// UnpackNameNotRenewableError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NameNotRenewable(string label)
func (eTHRegistrar *ETHRegistrar) UnpackNameNotRenewableError(raw []byte) (*ETHRegistrarNameNotRenewable, error) {
	out := new(ETHRegistrarNameNotRenewable)
	if err := eTHRegistrar.abi.UnpackIntoInterface(out, "NameNotRenewable", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ETHRegistrarOwnableInvalidOwner represents a OwnableInvalidOwner error raised by the ETHRegistrar contract.
type ETHRegistrarOwnableInvalidOwner struct {
	Owner common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OwnableInvalidOwner(address owner)
func ETHRegistrarOwnableInvalidOwnerErrorID() common.Hash {
	return common.HexToHash("0x1e4fbdf7f3ef8bcaa855599e3abf48b232380f183f08f6f813d9ffa5bd585188")
}

// UnpackOwnableInvalidOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OwnableInvalidOwner(address owner)
func (eTHRegistrar *ETHRegistrar) UnpackOwnableInvalidOwnerError(raw []byte) (*ETHRegistrarOwnableInvalidOwner, error) {
	out := new(ETHRegistrarOwnableInvalidOwner)
	if err := eTHRegistrar.abi.UnpackIntoInterface(out, "OwnableInvalidOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ETHRegistrarOwnableUnauthorizedAccount represents a OwnableUnauthorizedAccount error raised by the ETHRegistrar contract.
type ETHRegistrarOwnableUnauthorizedAccount struct {
	Account common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OwnableUnauthorizedAccount(address account)
func ETHRegistrarOwnableUnauthorizedAccountErrorID() common.Hash {
	return common.HexToHash("0x118cdaa7a341953d1887a2245fd6665d741c67c8c50581daa59e1d03373fa188")
}

// UnpackOwnableUnauthorizedAccountError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OwnableUnauthorizedAccount(address account)
func (eTHRegistrar *ETHRegistrar) UnpackOwnableUnauthorizedAccountError(raw []byte) (*ETHRegistrarOwnableUnauthorizedAccount, error) {
	out := new(ETHRegistrarOwnableUnauthorizedAccount)
	if err := eTHRegistrar.abi.UnpackIntoInterface(out, "OwnableUnauthorizedAccount", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ETHRegistrarSafeERC20FailedOperation represents a SafeERC20FailedOperation error raised by the ETHRegistrar contract.
type ETHRegistrarSafeERC20FailedOperation struct {
	Token common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SafeERC20FailedOperation(address token)
func ETHRegistrarSafeERC20FailedOperationErrorID() common.Hash {
	return common.HexToHash("0x5274afe73c98b4749fc91ffae6b7b574e7842cb2144a159e9377a5f20b32edf9")
}

// UnpackSafeERC20FailedOperationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SafeERC20FailedOperation(address token)
func (eTHRegistrar *ETHRegistrar) UnpackSafeERC20FailedOperationError(raw []byte) (*ETHRegistrarSafeERC20FailedOperation, error) {
	out := new(ETHRegistrarSafeERC20FailedOperation)
	if err := eTHRegistrar.abi.UnpackIntoInterface(out, "SafeERC20FailedOperation", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ETHRegistrarUnexpiredCommitmentExists represents a UnexpiredCommitmentExists error raised by the ETHRegistrar contract.
type ETHRegistrarUnexpiredCommitmentExists struct {
	Commitment [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UnexpiredCommitmentExists(bytes32 commitment)
func ETHRegistrarUnexpiredCommitmentExistsErrorID() common.Hash {
	return common.HexToHash("0x0a059d7159cc5b59fb271d3def536a153d37926e0d20a02ace2fa1e6ca4a8738")
}

// UnpackUnexpiredCommitmentExistsError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UnexpiredCommitmentExists(bytes32 commitment)
func (eTHRegistrar *ETHRegistrar) UnpackUnexpiredCommitmentExistsError(raw []byte) (*ETHRegistrarUnexpiredCommitmentExists, error) {
	out := new(ETHRegistrarUnexpiredCommitmentExists)
	if err := eTHRegistrar.abi.UnpackIntoInterface(out, "UnexpiredCommitmentExists", raw); err != nil {
		return nil, err
	}
	return out, nil
}
