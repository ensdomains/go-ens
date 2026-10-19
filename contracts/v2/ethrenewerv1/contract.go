// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethrenewerv1

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

// ETHRenewerV1MetaData contains all meta data concerning the ETHRenewerV1 contract.
var ETHRenewerV1MetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"},{\"internalType\":\"contractIPermissionedRegistry\",\"name\":\"ethRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"contractIRentPriceOracle\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"gracePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"bonusPeriod\",\"type\":\"uint64\"},{\"internalType\":\"contractINameWrapper\",\"name\":\"nameWrapper\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"wrappedController\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"duration\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minDuration\",\"type\":\"uint64\"}],\"name\":\"DurationTooShort\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"}],\"name\":\"NameNotRenewable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"duration\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newExpiry\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"referrer\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"NameRenewed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIRentPriceOracle\",\"name\":\"oracle\",\"type\":\"address\"}],\"name\":\"RentPriceOracleUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BASE_REGISTRAR\",\"outputs\":[{\"internalType\":\"contractBaseRegistrarImplementation\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BENEFICIARY\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETH_REGISTRY\",\"outputs\":[{\"internalType\":\"contractIPermissionedRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GRACE_PERIOD\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_RENEW_DURATION\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NAME_WRAPPER\",\"outputs\":[{\"internalType\":\"contractINameWrapper\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WRAPPED_CONTROLLER\",\"outputs\":[{\"internalType\":\"contractIWrappedETHRegistrarController\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"}],\"name\":\"getRemainingGracePeriod\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"duration\",\"type\":\"uint64\"},{\"internalType\":\"contractIERC20\",\"name\":\"paymentToken\",\"type\":\"address\"}],\"name\":\"getRenewPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"}],\"name\":\"isRenewable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"duration\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"referrer\",\"type\":\"bytes32\"}],\"internalType\":\"structRenewData\",\"name\":\"rd\",\"type\":\"tuple\"},{\"internalType\":\"contractIERC20\",\"name\":\"paymentToken\",\"type\":\"address\"}],\"name\":\"renew\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"duration\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"referrer\",\"type\":\"bytes32\"}],\"internalType\":\"structRenewData[]\",\"name\":\"rds\",\"type\":\"tuple[]\"},{\"internalType\":\"contractIERC20\",\"name\":\"paymentToken\",\"type\":\"address\"}],\"name\":\"renewBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rentPriceOracle\",\"outputs\":[{\"internalType\":\"contractIRentPriceOracle\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"}],\"name\":\"setRegistrarResolver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIRentPriceOracle\",\"name\":\"oracle\",\"type\":\"address\"}],\"name\":\"setRentPriceOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"labels\",\"type\":\"string[]\"}],\"name\":\"syncWrapper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferRegistrarOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "ETHRenewerV1",
}

// ETHRenewerV1 is an auto generated Go binding around an Ethereum contract.
type ETHRenewerV1 struct {
	abi abi.ABI
}

// GetABI returns the ABI associated with this contract binding.
func (c *ETHRenewerV1) GetABI() abi.ABI {
	return c.abi
}

// NewETHRenewerV1 creates a new instance of ETHRenewerV1.
func NewETHRenewerV1() *ETHRenewerV1 {
	parsed, err := ETHRenewerV1MetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &ETHRenewerV1{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *ETHRenewerV1) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address owner_, address ethRegistry, address beneficiary, address oracle, uint64 gracePeriod, uint64 bonusPeriod, address nameWrapper, address wrappedController) returns()
func (eTHRenewerV1 *ETHRenewerV1) PackConstructor(owner_ common.Address, ethRegistry common.Address, beneficiary common.Address, oracle common.Address, gracePeriod uint64, bonusPeriod uint64, nameWrapper common.Address, wrappedController common.Address) []byte {
	enc, err := eTHRenewerV1.abi.Pack("", owner_, ethRegistry, beneficiary, oracle, gracePeriod, bonusPeriod, nameWrapper, wrappedController)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackBASEREGISTRAR is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xcd93adf5.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function BASE_REGISTRAR() view returns(address)
func (eTHRenewerV1 *ETHRenewerV1) PackBASEREGISTRAR() []byte {
	enc, err := eTHRenewerV1.abi.Pack("BASE_REGISTRAR")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackBASEREGISTRAR is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xcd93adf5.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function BASE_REGISTRAR() view returns(address)
func (eTHRenewerV1 *ETHRenewerV1) TryPackBASEREGISTRAR() ([]byte, error) {
	return eTHRenewerV1.abi.Pack("BASE_REGISTRAR")
}

// UnpackBASEREGISTRAR is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xcd93adf5.
//
// Solidity: function BASE_REGISTRAR() view returns(address)
func (eTHRenewerV1 *ETHRenewerV1) UnpackBASEREGISTRAR(data []byte) (common.Address, error) {
	out, err := eTHRenewerV1.abi.Unpack("BASE_REGISTRAR", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackBENEFICIARY is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2f99c6cc.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function BENEFICIARY() view returns(address)
func (eTHRenewerV1 *ETHRenewerV1) PackBENEFICIARY() []byte {
	enc, err := eTHRenewerV1.abi.Pack("BENEFICIARY")
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
func (eTHRenewerV1 *ETHRenewerV1) TryPackBENEFICIARY() ([]byte, error) {
	return eTHRenewerV1.abi.Pack("BENEFICIARY")
}

// UnpackBENEFICIARY is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x2f99c6cc.
//
// Solidity: function BENEFICIARY() view returns(address)
func (eTHRenewerV1 *ETHRenewerV1) UnpackBENEFICIARY(data []byte) (common.Address, error) {
	out, err := eTHRenewerV1.abi.Unpack("BENEFICIARY", data)
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
func (eTHRenewerV1 *ETHRenewerV1) PackETHREGISTRY() []byte {
	enc, err := eTHRenewerV1.abi.Pack("ETH_REGISTRY")
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
func (eTHRenewerV1 *ETHRenewerV1) TryPackETHREGISTRY() ([]byte, error) {
	return eTHRenewerV1.abi.Pack("ETH_REGISTRY")
}

// UnpackETHREGISTRY is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x47500708.
//
// Solidity: function ETH_REGISTRY() view returns(address)
func (eTHRenewerV1 *ETHRenewerV1) UnpackETHREGISTRY(data []byte) (common.Address, error) {
	out, err := eTHRenewerV1.abi.Unpack("ETH_REGISTRY", data)
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
func (eTHRenewerV1 *ETHRenewerV1) PackGRACEPERIOD() []byte {
	enc, err := eTHRenewerV1.abi.Pack("GRACE_PERIOD")
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
func (eTHRenewerV1 *ETHRenewerV1) TryPackGRACEPERIOD() ([]byte, error) {
	return eTHRenewerV1.abi.Pack("GRACE_PERIOD")
}

// UnpackGRACEPERIOD is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc1a287e2.
//
// Solidity: function GRACE_PERIOD() view returns(uint64)
func (eTHRenewerV1 *ETHRenewerV1) UnpackGRACEPERIOD(data []byte) (uint64, error) {
	out, err := eTHRenewerV1.abi.Unpack("GRACE_PERIOD", data)
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
func (eTHRenewerV1 *ETHRenewerV1) PackMINRENEWDURATION() []byte {
	enc, err := eTHRenewerV1.abi.Pack("MIN_RENEW_DURATION")
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
func (eTHRenewerV1 *ETHRenewerV1) TryPackMINRENEWDURATION() ([]byte, error) {
	return eTHRenewerV1.abi.Pack("MIN_RENEW_DURATION")
}

// UnpackMINRENEWDURATION is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x307a64a5.
//
// Solidity: function MIN_RENEW_DURATION() view returns(uint64)
func (eTHRenewerV1 *ETHRenewerV1) UnpackMINRENEWDURATION(data []byte) (uint64, error) {
	out, err := eTHRenewerV1.abi.Unpack("MIN_RENEW_DURATION", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, nil
}

// PackNAMEWRAPPER is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x192cf07d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function NAME_WRAPPER() view returns(address)
func (eTHRenewerV1 *ETHRenewerV1) PackNAMEWRAPPER() []byte {
	enc, err := eTHRenewerV1.abi.Pack("NAME_WRAPPER")
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
func (eTHRenewerV1 *ETHRenewerV1) TryPackNAMEWRAPPER() ([]byte, error) {
	return eTHRenewerV1.abi.Pack("NAME_WRAPPER")
}

// UnpackNAMEWRAPPER is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x192cf07d.
//
// Solidity: function NAME_WRAPPER() view returns(address)
func (eTHRenewerV1 *ETHRenewerV1) UnpackNAMEWRAPPER(data []byte) (common.Address, error) {
	out, err := eTHRenewerV1.abi.Unpack("NAME_WRAPPER", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackWRAPPEDCONTROLLER is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdba10021.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function WRAPPED_CONTROLLER() view returns(address)
func (eTHRenewerV1 *ETHRenewerV1) PackWRAPPEDCONTROLLER() []byte {
	enc, err := eTHRenewerV1.abi.Pack("WRAPPED_CONTROLLER")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackWRAPPEDCONTROLLER is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdba10021.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function WRAPPED_CONTROLLER() view returns(address)
func (eTHRenewerV1 *ETHRenewerV1) TryPackWRAPPEDCONTROLLER() ([]byte, error) {
	return eTHRenewerV1.abi.Pack("WRAPPED_CONTROLLER")
}

// UnpackWRAPPEDCONTROLLER is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xdba10021.
//
// Solidity: function WRAPPED_CONTROLLER() view returns(address)
func (eTHRenewerV1 *ETHRenewerV1) UnpackWRAPPEDCONTROLLER(data []byte) (common.Address, error) {
	out, err := eTHRenewerV1.abi.Unpack("WRAPPED_CONTROLLER", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetRemainingGracePeriod is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x802295ef.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getRemainingGracePeriod(string label) view returns(uint64)
func (eTHRenewerV1 *ETHRenewerV1) PackGetRemainingGracePeriod(label string) []byte {
	enc, err := eTHRenewerV1.abi.Pack("getRemainingGracePeriod", label)
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
func (eTHRenewerV1 *ETHRenewerV1) TryPackGetRemainingGracePeriod(label string) ([]byte, error) {
	return eTHRenewerV1.abi.Pack("getRemainingGracePeriod", label)
}

// UnpackGetRemainingGracePeriod is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x802295ef.
//
// Solidity: function getRemainingGracePeriod(string label) view returns(uint64)
func (eTHRenewerV1 *ETHRenewerV1) UnpackGetRemainingGracePeriod(data []byte) (uint64, error) {
	out, err := eTHRenewerV1.abi.Unpack("getRemainingGracePeriod", data)
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
func (eTHRenewerV1 *ETHRenewerV1) PackGetRenewPrice(label string, duration uint64, paymentToken common.Address) []byte {
	enc, err := eTHRenewerV1.abi.Pack("getRenewPrice", label, duration, paymentToken)
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
func (eTHRenewerV1 *ETHRenewerV1) TryPackGetRenewPrice(label string, duration uint64, paymentToken common.Address) ([]byte, error) {
	return eTHRenewerV1.abi.Pack("getRenewPrice", label, duration, paymentToken)
}

// UnpackGetRenewPrice is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xddf0effc.
//
// Solidity: function getRenewPrice(string label, uint64 duration, address paymentToken) view returns(uint256)
func (eTHRenewerV1 *ETHRenewerV1) UnpackGetRenewPrice(data []byte) (*big.Int, error) {
	out, err := eTHRenewerV1.abi.Unpack("getRenewPrice", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackIsRenewable is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x130d6f00.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isRenewable(string label) view returns(bool)
func (eTHRenewerV1 *ETHRenewerV1) PackIsRenewable(label string) []byte {
	enc, err := eTHRenewerV1.abi.Pack("isRenewable", label)
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
func (eTHRenewerV1 *ETHRenewerV1) TryPackIsRenewable(label string) ([]byte, error) {
	return eTHRenewerV1.abi.Pack("isRenewable", label)
}

// UnpackIsRenewable is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x130d6f00.
//
// Solidity: function isRenewable(string label) view returns(bool)
func (eTHRenewerV1 *ETHRenewerV1) UnpackIsRenewable(data []byte) (bool, error) {
	out, err := eTHRenewerV1.abi.Unpack("isRenewable", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8da5cb5b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function owner() view returns(address)
func (eTHRenewerV1 *ETHRenewerV1) PackOwner() []byte {
	enc, err := eTHRenewerV1.abi.Pack("owner")
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
func (eTHRenewerV1 *ETHRenewerV1) TryPackOwner() ([]byte, error) {
	return eTHRenewerV1.abi.Pack("owner")
}

// UnpackOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (eTHRenewerV1 *ETHRenewerV1) UnpackOwner(data []byte) (common.Address, error) {
	out, err := eTHRenewerV1.abi.Unpack("owner", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackRenew is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2d06649b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function renew((string,uint64,bytes32) rd, address paymentToken) returns()
func (eTHRenewerV1 *ETHRenewerV1) PackRenew(rd RenewData, paymentToken common.Address) []byte {
	enc, err := eTHRenewerV1.abi.Pack("renew", rd, paymentToken)
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
func (eTHRenewerV1 *ETHRenewerV1) TryPackRenew(rd RenewData, paymentToken common.Address) ([]byte, error) {
	return eTHRenewerV1.abi.Pack("renew", rd, paymentToken)
}

// PackRenewBatch is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x959d530d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function renewBatch((string,uint64,bytes32)[] rds, address paymentToken) returns()
func (eTHRenewerV1 *ETHRenewerV1) PackRenewBatch(rds []RenewData, paymentToken common.Address) []byte {
	enc, err := eTHRenewerV1.abi.Pack("renewBatch", rds, paymentToken)
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
func (eTHRenewerV1 *ETHRenewerV1) TryPackRenewBatch(rds []RenewData, paymentToken common.Address) ([]byte, error) {
	return eTHRenewerV1.abi.Pack("renewBatch", rds, paymentToken)
}

// PackRenounceOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x715018a6.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function renounceOwnership() returns()
func (eTHRenewerV1 *ETHRenewerV1) PackRenounceOwnership() []byte {
	enc, err := eTHRenewerV1.abi.Pack("renounceOwnership")
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
func (eTHRenewerV1 *ETHRenewerV1) TryPackRenounceOwnership() ([]byte, error) {
	return eTHRenewerV1.abi.Pack("renounceOwnership")
}

// PackRentPriceOracle is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7b39ba16.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function rentPriceOracle() view returns(address)
func (eTHRenewerV1 *ETHRenewerV1) PackRentPriceOracle() []byte {
	enc, err := eTHRenewerV1.abi.Pack("rentPriceOracle")
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
func (eTHRenewerV1 *ETHRenewerV1) TryPackRentPriceOracle() ([]byte, error) {
	return eTHRenewerV1.abi.Pack("rentPriceOracle")
}

// UnpackRentPriceOracle is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7b39ba16.
//
// Solidity: function rentPriceOracle() view returns(address)
func (eTHRenewerV1 *ETHRenewerV1) UnpackRentPriceOracle(data []byte) (common.Address, error) {
	out, err := eTHRenewerV1.abi.Unpack("rentPriceOracle", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackSetRegistrarResolver is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x198e4236.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setRegistrarResolver(address resolver) returns()
func (eTHRenewerV1 *ETHRenewerV1) PackSetRegistrarResolver(resolver common.Address) []byte {
	enc, err := eTHRenewerV1.abi.Pack("setRegistrarResolver", resolver)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetRegistrarResolver is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x198e4236.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setRegistrarResolver(address resolver) returns()
func (eTHRenewerV1 *ETHRenewerV1) TryPackSetRegistrarResolver(resolver common.Address) ([]byte, error) {
	return eTHRenewerV1.abi.Pack("setRegistrarResolver", resolver)
}

// PackSetRentPriceOracle is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa2a11fbe.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setRentPriceOracle(address oracle) returns()
func (eTHRenewerV1 *ETHRenewerV1) PackSetRentPriceOracle(oracle common.Address) []byte {
	enc, err := eTHRenewerV1.abi.Pack("setRentPriceOracle", oracle)
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
func (eTHRenewerV1 *ETHRenewerV1) TryPackSetRentPriceOracle(oracle common.Address) ([]byte, error) {
	return eTHRenewerV1.abi.Pack("setRentPriceOracle", oracle)
}

// PackSupportsInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01ffc9a7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (eTHRenewerV1 *ETHRenewerV1) PackSupportsInterface(interfaceId [4]byte) []byte {
	enc, err := eTHRenewerV1.abi.Pack("supportsInterface", interfaceId)
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
func (eTHRenewerV1 *ETHRenewerV1) TryPackSupportsInterface(interfaceId [4]byte) ([]byte, error) {
	return eTHRenewerV1.abi.Pack("supportsInterface", interfaceId)
}

// UnpackSupportsInterface is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (eTHRenewerV1 *ETHRenewerV1) UnpackSupportsInterface(data []byte) (bool, error) {
	out, err := eTHRenewerV1.abi.Unpack("supportsInterface", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackSyncWrapper is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa2596702.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function syncWrapper(string[] labels) returns()
func (eTHRenewerV1 *ETHRenewerV1) PackSyncWrapper(labels []string) []byte {
	enc, err := eTHRenewerV1.abi.Pack("syncWrapper", labels)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSyncWrapper is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa2596702.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function syncWrapper(string[] labels) returns()
func (eTHRenewerV1 *ETHRenewerV1) TryPackSyncWrapper(labels []string) ([]byte, error) {
	return eTHRenewerV1.abi.Pack("syncWrapper", labels)
}

// PackTransferOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf2fde38b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (eTHRenewerV1 *ETHRenewerV1) PackTransferOwnership(newOwner common.Address) []byte {
	enc, err := eTHRenewerV1.abi.Pack("transferOwnership", newOwner)
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
func (eTHRenewerV1 *ETHRenewerV1) TryPackTransferOwnership(newOwner common.Address) ([]byte, error) {
	return eTHRenewerV1.abi.Pack("transferOwnership", newOwner)
}

// PackTransferRegistrarOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x154fbf15.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function transferRegistrarOwnership(address newOwner) returns()
func (eTHRenewerV1 *ETHRenewerV1) PackTransferRegistrarOwnership(newOwner common.Address) []byte {
	enc, err := eTHRenewerV1.abi.Pack("transferRegistrarOwnership", newOwner)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTransferRegistrarOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x154fbf15.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function transferRegistrarOwnership(address newOwner) returns()
func (eTHRenewerV1 *ETHRenewerV1) TryPackTransferRegistrarOwnership(newOwner common.Address) ([]byte, error) {
	return eTHRenewerV1.abi.Pack("transferRegistrarOwnership", newOwner)
}

// ETHRenewerV1NameRenewed represents a NameRenewed event raised by the ETHRenewerV1 contract.
type ETHRenewerV1NameRenewed struct {
	TokenId      *big.Int
	Label        string
	Duration     uint64
	NewExpiry    uint64
	PaymentToken common.Address
	Referrer     [32]byte
	Amount       *big.Int
	Raw          *types.Log // Blockchain specific contextual infos
}

const ETHRenewerV1NameRenewedEventName = "NameRenewed"

// ContractEventName returns the user-defined event name.
func (ETHRenewerV1NameRenewed) ContractEventName() string {
	return ETHRenewerV1NameRenewedEventName
}

// UnpackNameRenewedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event NameRenewed(uint256 indexed tokenId, string label, uint64 duration, uint64 newExpiry, address paymentToken, bytes32 indexed referrer, uint256 amount)
func (eTHRenewerV1 *ETHRenewerV1) UnpackNameRenewedEvent(log *types.Log) (*ETHRenewerV1NameRenewed, error) {
	event := "NameRenewed"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != eTHRenewerV1.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(ETHRenewerV1NameRenewed)
	if len(log.Data) > 0 {
		if err := eTHRenewerV1.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eTHRenewerV1.abi.Events[event].Inputs {
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

// ETHRenewerV1OwnershipTransferred represents a OwnershipTransferred event raised by the ETHRenewerV1 contract.
type ETHRenewerV1OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           *types.Log // Blockchain specific contextual infos
}

const ETHRenewerV1OwnershipTransferredEventName = "OwnershipTransferred"

// ContractEventName returns the user-defined event name.
func (ETHRenewerV1OwnershipTransferred) ContractEventName() string {
	return ETHRenewerV1OwnershipTransferredEventName
}

// UnpackOwnershipTransferredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (eTHRenewerV1 *ETHRenewerV1) UnpackOwnershipTransferredEvent(log *types.Log) (*ETHRenewerV1OwnershipTransferred, error) {
	event := "OwnershipTransferred"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != eTHRenewerV1.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(ETHRenewerV1OwnershipTransferred)
	if len(log.Data) > 0 {
		if err := eTHRenewerV1.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eTHRenewerV1.abi.Events[event].Inputs {
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

// ETHRenewerV1RentPriceOracleUpdated represents a RentPriceOracleUpdated event raised by the ETHRenewerV1 contract.
type ETHRenewerV1RentPriceOracleUpdated struct {
	Oracle common.Address
	Raw    *types.Log // Blockchain specific contextual infos
}

const ETHRenewerV1RentPriceOracleUpdatedEventName = "RentPriceOracleUpdated"

// ContractEventName returns the user-defined event name.
func (ETHRenewerV1RentPriceOracleUpdated) ContractEventName() string {
	return ETHRenewerV1RentPriceOracleUpdatedEventName
}

// UnpackRentPriceOracleUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RentPriceOracleUpdated(address oracle)
func (eTHRenewerV1 *ETHRenewerV1) UnpackRentPriceOracleUpdatedEvent(log *types.Log) (*ETHRenewerV1RentPriceOracleUpdated, error) {
	event := "RentPriceOracleUpdated"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != eTHRenewerV1.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(ETHRenewerV1RentPriceOracleUpdated)
	if len(log.Data) > 0 {
		if err := eTHRenewerV1.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eTHRenewerV1.abi.Events[event].Inputs {
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
func (eTHRenewerV1 *ETHRenewerV1) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], eTHRenewerV1.abi.Errors["DurationTooShort"].ID.Bytes()[:4]) {
		return eTHRenewerV1.UnpackDurationTooShortError(raw[4:])
	}
	if bytes.Equal(raw[:4], eTHRenewerV1.abi.Errors["NameNotRenewable"].ID.Bytes()[:4]) {
		return eTHRenewerV1.UnpackNameNotRenewableError(raw[4:])
	}
	if bytes.Equal(raw[:4], eTHRenewerV1.abi.Errors["OwnableInvalidOwner"].ID.Bytes()[:4]) {
		return eTHRenewerV1.UnpackOwnableInvalidOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], eTHRenewerV1.abi.Errors["OwnableUnauthorizedAccount"].ID.Bytes()[:4]) {
		return eTHRenewerV1.UnpackOwnableUnauthorizedAccountError(raw[4:])
	}
	if bytes.Equal(raw[:4], eTHRenewerV1.abi.Errors["SafeERC20FailedOperation"].ID.Bytes()[:4]) {
		return eTHRenewerV1.UnpackSafeERC20FailedOperationError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// ETHRenewerV1DurationTooShort represents a DurationTooShort error raised by the ETHRenewerV1 contract.
type ETHRenewerV1DurationTooShort struct {
	Duration    uint64
	MinDuration uint64
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error DurationTooShort(uint64 duration, uint64 minDuration)
func ETHRenewerV1DurationTooShortErrorID() common.Hash {
	return common.HexToHash("0xa096b844523daa5d1bb83afd97baa99f239f9b57ef2a99fb5e67d6abf00e360b")
}

// UnpackDurationTooShortError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error DurationTooShort(uint64 duration, uint64 minDuration)
func (eTHRenewerV1 *ETHRenewerV1) UnpackDurationTooShortError(raw []byte) (*ETHRenewerV1DurationTooShort, error) {
	out := new(ETHRenewerV1DurationTooShort)
	if err := eTHRenewerV1.abi.UnpackIntoInterface(out, "DurationTooShort", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ETHRenewerV1NameNotRenewable represents a NameNotRenewable error raised by the ETHRenewerV1 contract.
type ETHRenewerV1NameNotRenewable struct {
	Label string
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NameNotRenewable(string label)
func ETHRenewerV1NameNotRenewableErrorID() common.Hash {
	return common.HexToHash("0x1caefaa0709e9c51b7dc0085c2af73ef0ee280cad1968e0631809301c96a86e7")
}

// UnpackNameNotRenewableError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NameNotRenewable(string label)
func (eTHRenewerV1 *ETHRenewerV1) UnpackNameNotRenewableError(raw []byte) (*ETHRenewerV1NameNotRenewable, error) {
	out := new(ETHRenewerV1NameNotRenewable)
	if err := eTHRenewerV1.abi.UnpackIntoInterface(out, "NameNotRenewable", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ETHRenewerV1OwnableInvalidOwner represents a OwnableInvalidOwner error raised by the ETHRenewerV1 contract.
type ETHRenewerV1OwnableInvalidOwner struct {
	Owner common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OwnableInvalidOwner(address owner)
func ETHRenewerV1OwnableInvalidOwnerErrorID() common.Hash {
	return common.HexToHash("0x1e4fbdf7f3ef8bcaa855599e3abf48b232380f183f08f6f813d9ffa5bd585188")
}

// UnpackOwnableInvalidOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OwnableInvalidOwner(address owner)
func (eTHRenewerV1 *ETHRenewerV1) UnpackOwnableInvalidOwnerError(raw []byte) (*ETHRenewerV1OwnableInvalidOwner, error) {
	out := new(ETHRenewerV1OwnableInvalidOwner)
	if err := eTHRenewerV1.abi.UnpackIntoInterface(out, "OwnableInvalidOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ETHRenewerV1OwnableUnauthorizedAccount represents a OwnableUnauthorizedAccount error raised by the ETHRenewerV1 contract.
type ETHRenewerV1OwnableUnauthorizedAccount struct {
	Account common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OwnableUnauthorizedAccount(address account)
func ETHRenewerV1OwnableUnauthorizedAccountErrorID() common.Hash {
	return common.HexToHash("0x118cdaa7a341953d1887a2245fd6665d741c67c8c50581daa59e1d03373fa188")
}

// UnpackOwnableUnauthorizedAccountError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OwnableUnauthorizedAccount(address account)
func (eTHRenewerV1 *ETHRenewerV1) UnpackOwnableUnauthorizedAccountError(raw []byte) (*ETHRenewerV1OwnableUnauthorizedAccount, error) {
	out := new(ETHRenewerV1OwnableUnauthorizedAccount)
	if err := eTHRenewerV1.abi.UnpackIntoInterface(out, "OwnableUnauthorizedAccount", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ETHRenewerV1SafeERC20FailedOperation represents a SafeERC20FailedOperation error raised by the ETHRenewerV1 contract.
type ETHRenewerV1SafeERC20FailedOperation struct {
	Token common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SafeERC20FailedOperation(address token)
func ETHRenewerV1SafeERC20FailedOperationErrorID() common.Hash {
	return common.HexToHash("0x5274afe73c98b4749fc91ffae6b7b574e7842cb2144a159e9377a5f20b32edf9")
}

// UnpackSafeERC20FailedOperationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SafeERC20FailedOperation(address token)
func (eTHRenewerV1 *ETHRenewerV1) UnpackSafeERC20FailedOperationError(raw []byte) (*ETHRenewerV1SafeERC20FailedOperation, error) {
	out := new(ETHRenewerV1SafeERC20FailedOperation)
	if err := eTHRenewerV1.abi.UnpackIntoInterface(out, "SafeERC20FailedOperation", raw); err != nil {
		return nil, err
	}
	return out, nil
}
