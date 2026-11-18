// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rentpriceoracle

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

// DiscountPoint is an auto generated low-level Go binding around an user-defined struct.
type DiscountPoint struct {
	Duration uint64
	Numer    *big.Int
}

// PaymentRatio is an auto generated low-level Go binding around an user-defined struct.
type PaymentRatio struct {
	PaymentToken common.Address
	Numer        *big.Int
	Denom        *big.Int
}

// StandardRentPriceOracleMetaData contains all meta data concerning the StandardRentPriceOracle contract.
var StandardRentPriceOracleMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rootAccount\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"baseRatePerCp\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"duration\",\"type\":\"uint64\"},{\"internalType\":\"uint128\",\"name\":\"numer\",\"type\":\"uint128\"}],\"internalType\":\"structDiscountPoint[]\",\"name\":\"discountPoints\",\"type\":\"tuple[]\"},{\"internalType\":\"uint128\",\"name\":\"discountDenominator\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"premiumPriceInitial\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"premiumHalvingPeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"premiumPeriod\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"contractIERC20\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"numer\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"denom\",\"type\":\"uint128\"}],\"internalType\":\"structPaymentRatio[]\",\"name\":\"paymentRatios\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"resource\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"roleBitmap\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"EACCannotGrantRoles\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"resource\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"roleBitmap\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"EACCannotRevokeRoles\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EACInvalidAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roleBitmap\",\"type\":\"uint256\"}],\"name\":\"EACInvalidRoleBitmap\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"resource\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"role\",\"type\":\"uint256\"}],\"name\":\"EACMaxAssignees\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"resource\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"role\",\"type\":\"uint256\"}],\"name\":\"EACMinAssignees\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EACRootResourceNotAllowed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"resource\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"roleBitmap\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"EACUnauthorizedAccountRoles\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBaseRates\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDiscount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRatio\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"}],\"name\":\"NotValid\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"paymentToken\",\"type\":\"address\"}],\"name\":\"PaymentTokenNotSupported\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"resource\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldRoleBitmap\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newRoleBitmap\",\"type\":\"uint256\"}],\"name\":\"EACRolesChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"numer\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"denom\",\"type\":\"uint128\"}],\"name\":\"PaymentTokenUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DISCOUNT_DENOMINATOR\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PREMIUM_HALVING_PERIOD\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PREMIUM_PERIOD\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PREMIUM_PRICE_INITIAL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PREMIUM_PRICE_OFFSET\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ROOT_RESOURCE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"duration\",\"type\":\"uint64\"}],\"name\":\"applyDiscount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"paymentToken\",\"type\":\"address\"}],\"name\":\"convertUnits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"paymentToken\",\"type\":\"address\"}],\"name\":\"disablePaymentToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"resource\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"roleBitmap\",\"type\":\"uint256\"}],\"name\":\"getAssigneeCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"counts\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mask\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"duration\",\"type\":\"uint64\"}],\"name\":\"getBasePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBaseRates\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDiscountPoints\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"duration\",\"type\":\"uint64\"},{\"internalType\":\"uint128\",\"name\":\"numer\",\"type\":\"uint128\"}],\"internalType\":\"structDiscountPoint[]\",\"name\":\"v\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"}],\"name\":\"getLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"paymentToken\",\"type\":\"address\"}],\"name\":\"getPaymentTokenRatio\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"numer\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"denom\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"duration\",\"type\":\"uint64\"}],\"name\":\"getPremiumPriceAfter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"available\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"duration\",\"type\":\"uint64\"},{\"internalType\":\"contractIERC20\",\"name\":\"paymentToken\",\"type\":\"address\"}],\"name\":\"getRegisterPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"base\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"premium\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"duration\",\"type\":\"uint64\"},{\"internalType\":\"contractIERC20\",\"name\":\"paymentToken\",\"type\":\"address\"}],\"name\":\"getRenewPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"resource\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"roleBitmap\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRoles\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roleBitmap\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRootRoles\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"resource\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"roleBitmap\",\"type\":\"uint256\"}],\"name\":\"hasAssignees\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"resource\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"roleBitmap\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRoles\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roleBitmap\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRootRoles\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"namer\",\"type\":\"address\"}],\"name\":\"isContractNamer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"resource\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"roleBitmap\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isOnlyAssignee\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"paymentToken\",\"type\":\"address\"}],\"name\":\"isPaymentToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"}],\"name\":\"isValid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"resource\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"roleBitmap\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRoles\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roleBitmap\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRootRoles\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"resource\",\"type\":\"uint256\"}],\"name\":\"roleCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"resource\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"roles\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"numer\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"denom\",\"type\":\"uint128\"}],\"name\":\"updatePaymentToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "StandardRentPriceOracle",
}

// StandardRentPriceOracle is an auto generated Go binding around an Ethereum contract.
type StandardRentPriceOracle struct {
	abi abi.ABI
}

// GetABI returns the ABI associated with this contract binding.
func (c *StandardRentPriceOracle) GetABI() abi.ABI {
	return c.abi
}

// NewStandardRentPriceOracle creates a new instance of StandardRentPriceOracle.
func NewStandardRentPriceOracle() *StandardRentPriceOracle {
	parsed, err := StandardRentPriceOracleMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &StandardRentPriceOracle{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *StandardRentPriceOracle) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address rootAccount, uint256[] baseRatePerCp, (uint64,uint128)[] discountPoints, uint128 discountDenominator, uint256 premiumPriceInitial, uint64 premiumHalvingPeriod, uint64 premiumPeriod, (address,uint128,uint128)[] paymentRatios) returns()
func (standardRentPriceOracle *StandardRentPriceOracle) PackConstructor(rootAccount common.Address, baseRatePerCp []*big.Int, discountPoints []DiscountPoint, discountDenominator *big.Int, premiumPriceInitial *big.Int, premiumHalvingPeriod uint64, premiumPeriod uint64, paymentRatios []PaymentRatio) []byte {
	enc, err := standardRentPriceOracle.abi.Pack("", rootAccount, baseRatePerCp, discountPoints, discountDenominator, premiumPriceInitial, premiumHalvingPeriod, premiumPeriod, paymentRatios)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackDISCOUNTDENOMINATOR is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x415175bb.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function DISCOUNT_DENOMINATOR() view returns(uint128)
func (standardRentPriceOracle *StandardRentPriceOracle) PackDISCOUNTDENOMINATOR() []byte {
	enc, err := standardRentPriceOracle.abi.Pack("DISCOUNT_DENOMINATOR")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDISCOUNTDENOMINATOR is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x415175bb.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function DISCOUNT_DENOMINATOR() view returns(uint128)
func (standardRentPriceOracle *StandardRentPriceOracle) TryPackDISCOUNTDENOMINATOR() ([]byte, error) {
	return standardRentPriceOracle.abi.Pack("DISCOUNT_DENOMINATOR")
}

// UnpackDISCOUNTDENOMINATOR is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x415175bb.
//
// Solidity: function DISCOUNT_DENOMINATOR() view returns(uint128)
func (standardRentPriceOracle *StandardRentPriceOracle) UnpackDISCOUNTDENOMINATOR(data []byte) (*big.Int, error) {
	out, err := standardRentPriceOracle.abi.Unpack("DISCOUNT_DENOMINATOR", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackPREMIUMHALVINGPERIOD is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb553630e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function PREMIUM_HALVING_PERIOD() view returns(uint64)
func (standardRentPriceOracle *StandardRentPriceOracle) PackPREMIUMHALVINGPERIOD() []byte {
	enc, err := standardRentPriceOracle.abi.Pack("PREMIUM_HALVING_PERIOD")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackPREMIUMHALVINGPERIOD is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb553630e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function PREMIUM_HALVING_PERIOD() view returns(uint64)
func (standardRentPriceOracle *StandardRentPriceOracle) TryPackPREMIUMHALVINGPERIOD() ([]byte, error) {
	return standardRentPriceOracle.abi.Pack("PREMIUM_HALVING_PERIOD")
}

// UnpackPREMIUMHALVINGPERIOD is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb553630e.
//
// Solidity: function PREMIUM_HALVING_PERIOD() view returns(uint64)
func (standardRentPriceOracle *StandardRentPriceOracle) UnpackPREMIUMHALVINGPERIOD(data []byte) (uint64, error) {
	out, err := standardRentPriceOracle.abi.Unpack("PREMIUM_HALVING_PERIOD", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, nil
}

// PackPREMIUMPERIOD is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6bab30b7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function PREMIUM_PERIOD() view returns(uint64)
func (standardRentPriceOracle *StandardRentPriceOracle) PackPREMIUMPERIOD() []byte {
	enc, err := standardRentPriceOracle.abi.Pack("PREMIUM_PERIOD")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackPREMIUMPERIOD is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6bab30b7.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function PREMIUM_PERIOD() view returns(uint64)
func (standardRentPriceOracle *StandardRentPriceOracle) TryPackPREMIUMPERIOD() ([]byte, error) {
	return standardRentPriceOracle.abi.Pack("PREMIUM_PERIOD")
}

// UnpackPREMIUMPERIOD is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6bab30b7.
//
// Solidity: function PREMIUM_PERIOD() view returns(uint64)
func (standardRentPriceOracle *StandardRentPriceOracle) UnpackPREMIUMPERIOD(data []byte) (uint64, error) {
	out, err := standardRentPriceOracle.abi.Unpack("PREMIUM_PERIOD", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, nil
}

// PackPREMIUMPRICEINITIAL is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xac0d63d0.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function PREMIUM_PRICE_INITIAL() view returns(uint256)
func (standardRentPriceOracle *StandardRentPriceOracle) PackPREMIUMPRICEINITIAL() []byte {
	enc, err := standardRentPriceOracle.abi.Pack("PREMIUM_PRICE_INITIAL")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackPREMIUMPRICEINITIAL is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xac0d63d0.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function PREMIUM_PRICE_INITIAL() view returns(uint256)
func (standardRentPriceOracle *StandardRentPriceOracle) TryPackPREMIUMPRICEINITIAL() ([]byte, error) {
	return standardRentPriceOracle.abi.Pack("PREMIUM_PRICE_INITIAL")
}

// UnpackPREMIUMPRICEINITIAL is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xac0d63d0.
//
// Solidity: function PREMIUM_PRICE_INITIAL() view returns(uint256)
func (standardRentPriceOracle *StandardRentPriceOracle) UnpackPREMIUMPRICEINITIAL(data []byte) (*big.Int, error) {
	out, err := standardRentPriceOracle.abi.Unpack("PREMIUM_PRICE_INITIAL", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackPREMIUMPRICEOFFSET is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x48b6781f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function PREMIUM_PRICE_OFFSET() view returns(uint256)
func (standardRentPriceOracle *StandardRentPriceOracle) PackPREMIUMPRICEOFFSET() []byte {
	enc, err := standardRentPriceOracle.abi.Pack("PREMIUM_PRICE_OFFSET")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackPREMIUMPRICEOFFSET is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x48b6781f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function PREMIUM_PRICE_OFFSET() view returns(uint256)
func (standardRentPriceOracle *StandardRentPriceOracle) TryPackPREMIUMPRICEOFFSET() ([]byte, error) {
	return standardRentPriceOracle.abi.Pack("PREMIUM_PRICE_OFFSET")
}

// UnpackPREMIUMPRICEOFFSET is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x48b6781f.
//
// Solidity: function PREMIUM_PRICE_OFFSET() view returns(uint256)
func (standardRentPriceOracle *StandardRentPriceOracle) UnpackPREMIUMPRICEOFFSET(data []byte) (*big.Int, error) {
	out, err := standardRentPriceOracle.abi.Unpack("PREMIUM_PRICE_OFFSET", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackROOTRESOURCE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1c3fc3eb.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function ROOT_RESOURCE() view returns(uint256)
func (standardRentPriceOracle *StandardRentPriceOracle) PackROOTRESOURCE() []byte {
	enc, err := standardRentPriceOracle.abi.Pack("ROOT_RESOURCE")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackROOTRESOURCE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1c3fc3eb.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function ROOT_RESOURCE() view returns(uint256)
func (standardRentPriceOracle *StandardRentPriceOracle) TryPackROOTRESOURCE() ([]byte, error) {
	return standardRentPriceOracle.abi.Pack("ROOT_RESOURCE")
}

// UnpackROOTRESOURCE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x1c3fc3eb.
//
// Solidity: function ROOT_RESOURCE() view returns(uint256)
func (standardRentPriceOracle *StandardRentPriceOracle) UnpackROOTRESOURCE(data []byte) (*big.Int, error) {
	out, err := standardRentPriceOracle.abi.Unpack("ROOT_RESOURCE", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackApplyDiscount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x34d82622.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function applyDiscount(uint256 value, uint64 duration) view returns(uint256)
func (standardRentPriceOracle *StandardRentPriceOracle) PackApplyDiscount(value *big.Int, duration uint64) []byte {
	enc, err := standardRentPriceOracle.abi.Pack("applyDiscount", value, duration)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackApplyDiscount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x34d82622.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function applyDiscount(uint256 value, uint64 duration) view returns(uint256)
func (standardRentPriceOracle *StandardRentPriceOracle) TryPackApplyDiscount(value *big.Int, duration uint64) ([]byte, error) {
	return standardRentPriceOracle.abi.Pack("applyDiscount", value, duration)
}

// UnpackApplyDiscount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x34d82622.
//
// Solidity: function applyDiscount(uint256 value, uint64 duration) view returns(uint256)
func (standardRentPriceOracle *StandardRentPriceOracle) UnpackApplyDiscount(data []byte) (*big.Int, error) {
	out, err := standardRentPriceOracle.abi.Unpack("applyDiscount", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackConvertUnits is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8b59de43.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function convertUnits(uint256 value, address paymentToken) view returns(uint256)
func (standardRentPriceOracle *StandardRentPriceOracle) PackConvertUnits(value *big.Int, paymentToken common.Address) []byte {
	enc, err := standardRentPriceOracle.abi.Pack("convertUnits", value, paymentToken)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackConvertUnits is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8b59de43.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function convertUnits(uint256 value, address paymentToken) view returns(uint256)
func (standardRentPriceOracle *StandardRentPriceOracle) TryPackConvertUnits(value *big.Int, paymentToken common.Address) ([]byte, error) {
	return standardRentPriceOracle.abi.Pack("convertUnits", value, paymentToken)
}

// UnpackConvertUnits is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8b59de43.
//
// Solidity: function convertUnits(uint256 value, address paymentToken) view returns(uint256)
func (standardRentPriceOracle *StandardRentPriceOracle) UnpackConvertUnits(data []byte) (*big.Int, error) {
	out, err := standardRentPriceOracle.abi.Unpack("convertUnits", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackDisablePaymentToken is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9c7aa7f8.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function disablePaymentToken(address paymentToken) returns()
func (standardRentPriceOracle *StandardRentPriceOracle) PackDisablePaymentToken(paymentToken common.Address) []byte {
	enc, err := standardRentPriceOracle.abi.Pack("disablePaymentToken", paymentToken)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDisablePaymentToken is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9c7aa7f8.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function disablePaymentToken(address paymentToken) returns()
func (standardRentPriceOracle *StandardRentPriceOracle) TryPackDisablePaymentToken(paymentToken common.Address) ([]byte, error) {
	return standardRentPriceOracle.abi.Pack("disablePaymentToken", paymentToken)
}

// PackGetAssigneeCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3634f911.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getAssigneeCount(uint256 resource, uint256 roleBitmap) view returns(uint256 counts, uint256 mask)
func (standardRentPriceOracle *StandardRentPriceOracle) PackGetAssigneeCount(resource *big.Int, roleBitmap *big.Int) []byte {
	enc, err := standardRentPriceOracle.abi.Pack("getAssigneeCount", resource, roleBitmap)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetAssigneeCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3634f911.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getAssigneeCount(uint256 resource, uint256 roleBitmap) view returns(uint256 counts, uint256 mask)
func (standardRentPriceOracle *StandardRentPriceOracle) TryPackGetAssigneeCount(resource *big.Int, roleBitmap *big.Int) ([]byte, error) {
	return standardRentPriceOracle.abi.Pack("getAssigneeCount", resource, roleBitmap)
}

// GetAssigneeCountOutput serves as a container for the return parameters of contract
// method GetAssigneeCount.
type GetAssigneeCountOutput struct {
	Counts *big.Int
	Mask   *big.Int
}

// UnpackGetAssigneeCount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3634f911.
//
// Solidity: function getAssigneeCount(uint256 resource, uint256 roleBitmap) view returns(uint256 counts, uint256 mask)
func (standardRentPriceOracle *StandardRentPriceOracle) UnpackGetAssigneeCount(data []byte) (GetAssigneeCountOutput, error) {
	out, err := standardRentPriceOracle.abi.Unpack("getAssigneeCount", data)
	outstruct := new(GetAssigneeCountOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Counts = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.Mask = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	return *outstruct, nil
}

// PackGetBasePrice is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6e455111.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getBasePrice(string label, uint64 duration) view returns(uint256)
func (standardRentPriceOracle *StandardRentPriceOracle) PackGetBasePrice(label string, duration uint64) []byte {
	enc, err := standardRentPriceOracle.abi.Pack("getBasePrice", label, duration)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetBasePrice is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6e455111.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getBasePrice(string label, uint64 duration) view returns(uint256)
func (standardRentPriceOracle *StandardRentPriceOracle) TryPackGetBasePrice(label string, duration uint64) ([]byte, error) {
	return standardRentPriceOracle.abi.Pack("getBasePrice", label, duration)
}

// UnpackGetBasePrice is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6e455111.
//
// Solidity: function getBasePrice(string label, uint64 duration) view returns(uint256)
func (standardRentPriceOracle *StandardRentPriceOracle) UnpackGetBasePrice(data []byte) (*big.Int, error) {
	out, err := standardRentPriceOracle.abi.Unpack("getBasePrice", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetBaseRates is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc50f093b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getBaseRates() view returns(uint256[])
func (standardRentPriceOracle *StandardRentPriceOracle) PackGetBaseRates() []byte {
	enc, err := standardRentPriceOracle.abi.Pack("getBaseRates")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetBaseRates is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc50f093b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getBaseRates() view returns(uint256[])
func (standardRentPriceOracle *StandardRentPriceOracle) TryPackGetBaseRates() ([]byte, error) {
	return standardRentPriceOracle.abi.Pack("getBaseRates")
}

// UnpackGetBaseRates is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc50f093b.
//
// Solidity: function getBaseRates() view returns(uint256[])
func (standardRentPriceOracle *StandardRentPriceOracle) UnpackGetBaseRates(data []byte) ([]*big.Int, error) {
	out, err := standardRentPriceOracle.abi.Unpack("getBaseRates", data)
	if err != nil {
		return *new([]*big.Int), err
	}
	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	return out0, nil
}

// PackGetDiscountPoints is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x39ac7a2a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getDiscountPoints() view returns((uint64,uint128)[] v)
func (standardRentPriceOracle *StandardRentPriceOracle) PackGetDiscountPoints() []byte {
	enc, err := standardRentPriceOracle.abi.Pack("getDiscountPoints")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetDiscountPoints is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x39ac7a2a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getDiscountPoints() view returns((uint64,uint128)[] v)
func (standardRentPriceOracle *StandardRentPriceOracle) TryPackGetDiscountPoints() ([]byte, error) {
	return standardRentPriceOracle.abi.Pack("getDiscountPoints")
}

// UnpackGetDiscountPoints is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x39ac7a2a.
//
// Solidity: function getDiscountPoints() view returns((uint64,uint128)[] v)
func (standardRentPriceOracle *StandardRentPriceOracle) UnpackGetDiscountPoints(data []byte) ([]DiscountPoint, error) {
	out, err := standardRentPriceOracle.abi.Unpack("getDiscountPoints", data)
	if err != nil {
		return *new([]DiscountPoint), err
	}
	out0 := *abi.ConvertType(out[0], new([]DiscountPoint)).(*[]DiscountPoint)
	return out0, nil
}

// PackGetLength is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6981b5f4.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getLength(string label) pure returns(uint256)
func (standardRentPriceOracle *StandardRentPriceOracle) PackGetLength(label string) []byte {
	enc, err := standardRentPriceOracle.abi.Pack("getLength", label)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetLength is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6981b5f4.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getLength(string label) pure returns(uint256)
func (standardRentPriceOracle *StandardRentPriceOracle) TryPackGetLength(label string) ([]byte, error) {
	return standardRentPriceOracle.abi.Pack("getLength", label)
}

// UnpackGetLength is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6981b5f4.
//
// Solidity: function getLength(string label) pure returns(uint256)
func (standardRentPriceOracle *StandardRentPriceOracle) UnpackGetLength(data []byte) (*big.Int, error) {
	out, err := standardRentPriceOracle.abi.Unpack("getLength", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetPaymentTokenRatio is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x30897dba.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getPaymentTokenRatio(address paymentToken) view returns(uint128 numer, uint128 denom)
func (standardRentPriceOracle *StandardRentPriceOracle) PackGetPaymentTokenRatio(paymentToken common.Address) []byte {
	enc, err := standardRentPriceOracle.abi.Pack("getPaymentTokenRatio", paymentToken)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetPaymentTokenRatio is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x30897dba.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getPaymentTokenRatio(address paymentToken) view returns(uint128 numer, uint128 denom)
func (standardRentPriceOracle *StandardRentPriceOracle) TryPackGetPaymentTokenRatio(paymentToken common.Address) ([]byte, error) {
	return standardRentPriceOracle.abi.Pack("getPaymentTokenRatio", paymentToken)
}

// GetPaymentTokenRatioOutput serves as a container for the return parameters of contract
// method GetPaymentTokenRatio.
type GetPaymentTokenRatioOutput struct {
	Numer *big.Int
	Denom *big.Int
}

// UnpackGetPaymentTokenRatio is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x30897dba.
//
// Solidity: function getPaymentTokenRatio(address paymentToken) view returns(uint128 numer, uint128 denom)
func (standardRentPriceOracle *StandardRentPriceOracle) UnpackGetPaymentTokenRatio(data []byte) (GetPaymentTokenRatioOutput, error) {
	out, err := standardRentPriceOracle.abi.Unpack("getPaymentTokenRatio", data)
	outstruct := new(GetPaymentTokenRatioOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Numer = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.Denom = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	return *outstruct, nil
}

// PackGetPremiumPriceAfter is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9e6d1a93.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getPremiumPriceAfter(uint64 duration) view returns(uint256)
func (standardRentPriceOracle *StandardRentPriceOracle) PackGetPremiumPriceAfter(duration uint64) []byte {
	enc, err := standardRentPriceOracle.abi.Pack("getPremiumPriceAfter", duration)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetPremiumPriceAfter is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9e6d1a93.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getPremiumPriceAfter(uint64 duration) view returns(uint256)
func (standardRentPriceOracle *StandardRentPriceOracle) TryPackGetPremiumPriceAfter(duration uint64) ([]byte, error) {
	return standardRentPriceOracle.abi.Pack("getPremiumPriceAfter", duration)
}

// UnpackGetPremiumPriceAfter is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9e6d1a93.
//
// Solidity: function getPremiumPriceAfter(uint64 duration) view returns(uint256)
func (standardRentPriceOracle *StandardRentPriceOracle) UnpackGetPremiumPriceAfter(data []byte) (*big.Int, error) {
	out, err := standardRentPriceOracle.abi.Unpack("getPremiumPriceAfter", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetRegisterPrice is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe1de9c83.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getRegisterPrice(string label, uint64 available, uint64 duration, address paymentToken) view returns(uint256 base, uint256 premium)
func (standardRentPriceOracle *StandardRentPriceOracle) PackGetRegisterPrice(label string, available uint64, duration uint64, paymentToken common.Address) []byte {
	enc, err := standardRentPriceOracle.abi.Pack("getRegisterPrice", label, available, duration, paymentToken)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetRegisterPrice is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe1de9c83.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getRegisterPrice(string label, uint64 available, uint64 duration, address paymentToken) view returns(uint256 base, uint256 premium)
func (standardRentPriceOracle *StandardRentPriceOracle) TryPackGetRegisterPrice(label string, available uint64, duration uint64, paymentToken common.Address) ([]byte, error) {
	return standardRentPriceOracle.abi.Pack("getRegisterPrice", label, available, duration, paymentToken)
}

// GetRegisterPriceOutput serves as a container for the return parameters of contract
// method GetRegisterPrice.
type GetRegisterPriceOutput struct {
	Base    *big.Int
	Premium *big.Int
}

// UnpackGetRegisterPrice is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe1de9c83.
//
// Solidity: function getRegisterPrice(string label, uint64 available, uint64 duration, address paymentToken) view returns(uint256 base, uint256 premium)
func (standardRentPriceOracle *StandardRentPriceOracle) UnpackGetRegisterPrice(data []byte) (GetRegisterPriceOutput, error) {
	out, err := standardRentPriceOracle.abi.Unpack("getRegisterPrice", data)
	outstruct := new(GetRegisterPriceOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Base = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.Premium = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	return *outstruct, nil
}

// PackGetRenewPrice is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3ad86083.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getRenewPrice(string label, uint64 , uint64 duration, address paymentToken) view returns(uint256)
func (standardRentPriceOracle *StandardRentPriceOracle) PackGetRenewPrice(label string, arg1 uint64, duration uint64, paymentToken common.Address) []byte {
	enc, err := standardRentPriceOracle.abi.Pack("getRenewPrice", label, arg1, duration, paymentToken)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetRenewPrice is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3ad86083.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getRenewPrice(string label, uint64 , uint64 duration, address paymentToken) view returns(uint256)
func (standardRentPriceOracle *StandardRentPriceOracle) TryPackGetRenewPrice(label string, arg1 uint64, duration uint64, paymentToken common.Address) ([]byte, error) {
	return standardRentPriceOracle.abi.Pack("getRenewPrice", label, arg1, duration, paymentToken)
}

// UnpackGetRenewPrice is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3ad86083.
//
// Solidity: function getRenewPrice(string label, uint64 , uint64 duration, address paymentToken) view returns(uint256)
func (standardRentPriceOracle *StandardRentPriceOracle) UnpackGetRenewPrice(data []byte) (*big.Int, error) {
	out, err := standardRentPriceOracle.abi.Unpack("getRenewPrice", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGrantRoles is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7c300586.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function grantRoles(uint256 resource, uint256 roleBitmap, address account) returns(bool)
func (standardRentPriceOracle *StandardRentPriceOracle) PackGrantRoles(resource *big.Int, roleBitmap *big.Int, account common.Address) []byte {
	enc, err := standardRentPriceOracle.abi.Pack("grantRoles", resource, roleBitmap, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGrantRoles is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7c300586.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function grantRoles(uint256 resource, uint256 roleBitmap, address account) returns(bool)
func (standardRentPriceOracle *StandardRentPriceOracle) TryPackGrantRoles(resource *big.Int, roleBitmap *big.Int, account common.Address) ([]byte, error) {
	return standardRentPriceOracle.abi.Pack("grantRoles", resource, roleBitmap, account)
}

// UnpackGrantRoles is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7c300586.
//
// Solidity: function grantRoles(uint256 resource, uint256 roleBitmap, address account) returns(bool)
func (standardRentPriceOracle *StandardRentPriceOracle) UnpackGrantRoles(data []byte) (bool, error) {
	out, err := standardRentPriceOracle.abi.Unpack("grantRoles", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackGrantRootRoles is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x072d5d77.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function grantRootRoles(uint256 roleBitmap, address account) returns(bool)
func (standardRentPriceOracle *StandardRentPriceOracle) PackGrantRootRoles(roleBitmap *big.Int, account common.Address) []byte {
	enc, err := standardRentPriceOracle.abi.Pack("grantRootRoles", roleBitmap, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGrantRootRoles is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x072d5d77.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function grantRootRoles(uint256 roleBitmap, address account) returns(bool)
func (standardRentPriceOracle *StandardRentPriceOracle) TryPackGrantRootRoles(roleBitmap *big.Int, account common.Address) ([]byte, error) {
	return standardRentPriceOracle.abi.Pack("grantRootRoles", roleBitmap, account)
}

// UnpackGrantRootRoles is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x072d5d77.
//
// Solidity: function grantRootRoles(uint256 roleBitmap, address account) returns(bool)
func (standardRentPriceOracle *StandardRentPriceOracle) UnpackGrantRootRoles(data []byte) (bool, error) {
	out, err := standardRentPriceOracle.abi.Unpack("grantRootRoles", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackHasAssignees is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x11b8e00a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function hasAssignees(uint256 resource, uint256 roleBitmap) view returns(bool)
func (standardRentPriceOracle *StandardRentPriceOracle) PackHasAssignees(resource *big.Int, roleBitmap *big.Int) []byte {
	enc, err := standardRentPriceOracle.abi.Pack("hasAssignees", resource, roleBitmap)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackHasAssignees is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x11b8e00a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function hasAssignees(uint256 resource, uint256 roleBitmap) view returns(bool)
func (standardRentPriceOracle *StandardRentPriceOracle) TryPackHasAssignees(resource *big.Int, roleBitmap *big.Int) ([]byte, error) {
	return standardRentPriceOracle.abi.Pack("hasAssignees", resource, roleBitmap)
}

// UnpackHasAssignees is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x11b8e00a.
//
// Solidity: function hasAssignees(uint256 resource, uint256 roleBitmap) view returns(bool)
func (standardRentPriceOracle *StandardRentPriceOracle) UnpackHasAssignees(data []byte) (bool, error) {
	out, err := standardRentPriceOracle.abi.Unpack("hasAssignees", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackHasRoles is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd3bf89b1.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function hasRoles(uint256 resource, uint256 roleBitmap, address account) view returns(bool)
func (standardRentPriceOracle *StandardRentPriceOracle) PackHasRoles(resource *big.Int, roleBitmap *big.Int, account common.Address) []byte {
	enc, err := standardRentPriceOracle.abi.Pack("hasRoles", resource, roleBitmap, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackHasRoles is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd3bf89b1.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function hasRoles(uint256 resource, uint256 roleBitmap, address account) view returns(bool)
func (standardRentPriceOracle *StandardRentPriceOracle) TryPackHasRoles(resource *big.Int, roleBitmap *big.Int, account common.Address) ([]byte, error) {
	return standardRentPriceOracle.abi.Pack("hasRoles", resource, roleBitmap, account)
}

// UnpackHasRoles is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd3bf89b1.
//
// Solidity: function hasRoles(uint256 resource, uint256 roleBitmap, address account) view returns(bool)
func (standardRentPriceOracle *StandardRentPriceOracle) UnpackHasRoles(data []byte) (bool, error) {
	out, err := standardRentPriceOracle.abi.Unpack("hasRoles", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackHasRootRoles is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x781ef8db.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function hasRootRoles(uint256 roleBitmap, address account) view returns(bool)
func (standardRentPriceOracle *StandardRentPriceOracle) PackHasRootRoles(roleBitmap *big.Int, account common.Address) []byte {
	enc, err := standardRentPriceOracle.abi.Pack("hasRootRoles", roleBitmap, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackHasRootRoles is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x781ef8db.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function hasRootRoles(uint256 roleBitmap, address account) view returns(bool)
func (standardRentPriceOracle *StandardRentPriceOracle) TryPackHasRootRoles(roleBitmap *big.Int, account common.Address) ([]byte, error) {
	return standardRentPriceOracle.abi.Pack("hasRootRoles", roleBitmap, account)
}

// UnpackHasRootRoles is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x781ef8db.
//
// Solidity: function hasRootRoles(uint256 roleBitmap, address account) view returns(bool)
func (standardRentPriceOracle *StandardRentPriceOracle) UnpackHasRootRoles(data []byte) (bool, error) {
	out, err := standardRentPriceOracle.abi.Unpack("hasRootRoles", data)
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
func (standardRentPriceOracle *StandardRentPriceOracle) PackIsContractNamer(namer common.Address) []byte {
	enc, err := standardRentPriceOracle.abi.Pack("isContractNamer", namer)
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
func (standardRentPriceOracle *StandardRentPriceOracle) TryPackIsContractNamer(namer common.Address) ([]byte, error) {
	return standardRentPriceOracle.abi.Pack("isContractNamer", namer)
}

// UnpackIsContractNamer is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6f3ff726.
//
// Solidity: function isContractNamer(address namer) view returns(bool)
func (standardRentPriceOracle *StandardRentPriceOracle) UnpackIsContractNamer(data []byte) (bool, error) {
	out, err := standardRentPriceOracle.abi.Unpack("isContractNamer", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackIsOnlyAssignee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8e77c95f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isOnlyAssignee(uint256 resource, uint256 roleBitmap, address account) view returns(bool)
func (standardRentPriceOracle *StandardRentPriceOracle) PackIsOnlyAssignee(resource *big.Int, roleBitmap *big.Int, account common.Address) []byte {
	enc, err := standardRentPriceOracle.abi.Pack("isOnlyAssignee", resource, roleBitmap, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsOnlyAssignee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8e77c95f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isOnlyAssignee(uint256 resource, uint256 roleBitmap, address account) view returns(bool)
func (standardRentPriceOracle *StandardRentPriceOracle) TryPackIsOnlyAssignee(resource *big.Int, roleBitmap *big.Int, account common.Address) ([]byte, error) {
	return standardRentPriceOracle.abi.Pack("isOnlyAssignee", resource, roleBitmap, account)
}

// UnpackIsOnlyAssignee is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8e77c95f.
//
// Solidity: function isOnlyAssignee(uint256 resource, uint256 roleBitmap, address account) view returns(bool)
func (standardRentPriceOracle *StandardRentPriceOracle) UnpackIsOnlyAssignee(data []byte) (bool, error) {
	out, err := standardRentPriceOracle.abi.Unpack("isOnlyAssignee", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackIsPaymentToken is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x930eaddc.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isPaymentToken(address paymentToken) view returns(bool)
func (standardRentPriceOracle *StandardRentPriceOracle) PackIsPaymentToken(paymentToken common.Address) []byte {
	enc, err := standardRentPriceOracle.abi.Pack("isPaymentToken", paymentToken)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsPaymentToken is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x930eaddc.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isPaymentToken(address paymentToken) view returns(bool)
func (standardRentPriceOracle *StandardRentPriceOracle) TryPackIsPaymentToken(paymentToken common.Address) ([]byte, error) {
	return standardRentPriceOracle.abi.Pack("isPaymentToken", paymentToken)
}

// UnpackIsPaymentToken is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x930eaddc.
//
// Solidity: function isPaymentToken(address paymentToken) view returns(bool)
func (standardRentPriceOracle *StandardRentPriceOracle) UnpackIsPaymentToken(data []byte) (bool, error) {
	out, err := standardRentPriceOracle.abi.Unpack("isPaymentToken", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackIsValid is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf1dfaefb.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isValid(string label) view returns(bool)
func (standardRentPriceOracle *StandardRentPriceOracle) PackIsValid(label string) []byte {
	enc, err := standardRentPriceOracle.abi.Pack("isValid", label)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsValid is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf1dfaefb.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isValid(string label) view returns(bool)
func (standardRentPriceOracle *StandardRentPriceOracle) TryPackIsValid(label string) ([]byte, error) {
	return standardRentPriceOracle.abi.Pack("isValid", label)
}

// UnpackIsValid is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf1dfaefb.
//
// Solidity: function isValid(string label) view returns(bool)
func (standardRentPriceOracle *StandardRentPriceOracle) UnpackIsValid(data []byte) (bool, error) {
	out, err := standardRentPriceOracle.abi.Unpack("isValid", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackRevokeRoles is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdfa70d8b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function revokeRoles(uint256 resource, uint256 roleBitmap, address account) returns(bool)
func (standardRentPriceOracle *StandardRentPriceOracle) PackRevokeRoles(resource *big.Int, roleBitmap *big.Int, account common.Address) []byte {
	enc, err := standardRentPriceOracle.abi.Pack("revokeRoles", resource, roleBitmap, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRevokeRoles is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdfa70d8b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function revokeRoles(uint256 resource, uint256 roleBitmap, address account) returns(bool)
func (standardRentPriceOracle *StandardRentPriceOracle) TryPackRevokeRoles(resource *big.Int, roleBitmap *big.Int, account common.Address) ([]byte, error) {
	return standardRentPriceOracle.abi.Pack("revokeRoles", resource, roleBitmap, account)
}

// UnpackRevokeRoles is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xdfa70d8b.
//
// Solidity: function revokeRoles(uint256 resource, uint256 roleBitmap, address account) returns(bool)
func (standardRentPriceOracle *StandardRentPriceOracle) UnpackRevokeRoles(data []byte) (bool, error) {
	out, err := standardRentPriceOracle.abi.Unpack("revokeRoles", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackRevokeRootRoles is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xce156e82.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function revokeRootRoles(uint256 roleBitmap, address account) returns(bool)
func (standardRentPriceOracle *StandardRentPriceOracle) PackRevokeRootRoles(roleBitmap *big.Int, account common.Address) []byte {
	enc, err := standardRentPriceOracle.abi.Pack("revokeRootRoles", roleBitmap, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRevokeRootRoles is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xce156e82.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function revokeRootRoles(uint256 roleBitmap, address account) returns(bool)
func (standardRentPriceOracle *StandardRentPriceOracle) TryPackRevokeRootRoles(roleBitmap *big.Int, account common.Address) ([]byte, error) {
	return standardRentPriceOracle.abi.Pack("revokeRootRoles", roleBitmap, account)
}

// UnpackRevokeRootRoles is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xce156e82.
//
// Solidity: function revokeRootRoles(uint256 roleBitmap, address account) returns(bool)
func (standardRentPriceOracle *StandardRentPriceOracle) UnpackRevokeRootRoles(data []byte) (bool, error) {
	out, err := standardRentPriceOracle.abi.Unpack("revokeRootRoles", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackRoleCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2f27fa24.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function roleCount(uint256 resource) view returns(uint256)
func (standardRentPriceOracle *StandardRentPriceOracle) PackRoleCount(resource *big.Int) []byte {
	enc, err := standardRentPriceOracle.abi.Pack("roleCount", resource)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRoleCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2f27fa24.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function roleCount(uint256 resource) view returns(uint256)
func (standardRentPriceOracle *StandardRentPriceOracle) TryPackRoleCount(resource *big.Int) ([]byte, error) {
	return standardRentPriceOracle.abi.Pack("roleCount", resource)
}

// UnpackRoleCount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x2f27fa24.
//
// Solidity: function roleCount(uint256 resource) view returns(uint256)
func (standardRentPriceOracle *StandardRentPriceOracle) UnpackRoleCount(data []byte) (*big.Int, error) {
	out, err := standardRentPriceOracle.abi.Unpack("roleCount", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackRoles is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5adf4724.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function roles(uint256 resource, address account) view returns(uint256)
func (standardRentPriceOracle *StandardRentPriceOracle) PackRoles(resource *big.Int, account common.Address) []byte {
	enc, err := standardRentPriceOracle.abi.Pack("roles", resource, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRoles is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5adf4724.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function roles(uint256 resource, address account) view returns(uint256)
func (standardRentPriceOracle *StandardRentPriceOracle) TryPackRoles(resource *big.Int, account common.Address) ([]byte, error) {
	return standardRentPriceOracle.abi.Pack("roles", resource, account)
}

// UnpackRoles is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5adf4724.
//
// Solidity: function roles(uint256 resource, address account) view returns(uint256)
func (standardRentPriceOracle *StandardRentPriceOracle) UnpackRoles(data []byte) (*big.Int, error) {
	out, err := standardRentPriceOracle.abi.Unpack("roles", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackSupportsInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01ffc9a7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (standardRentPriceOracle *StandardRentPriceOracle) PackSupportsInterface(interfaceId [4]byte) []byte {
	enc, err := standardRentPriceOracle.abi.Pack("supportsInterface", interfaceId)
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
func (standardRentPriceOracle *StandardRentPriceOracle) TryPackSupportsInterface(interfaceId [4]byte) ([]byte, error) {
	return standardRentPriceOracle.abi.Pack("supportsInterface", interfaceId)
}

// UnpackSupportsInterface is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (standardRentPriceOracle *StandardRentPriceOracle) UnpackSupportsInterface(data []byte) (bool, error) {
	out, err := standardRentPriceOracle.abi.Unpack("supportsInterface", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackUpdatePaymentToken is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdff521a8.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function updatePaymentToken(address paymentToken, uint128 numer, uint128 denom) returns()
func (standardRentPriceOracle *StandardRentPriceOracle) PackUpdatePaymentToken(paymentToken common.Address, numer *big.Int, denom *big.Int) []byte {
	enc, err := standardRentPriceOracle.abi.Pack("updatePaymentToken", paymentToken, numer, denom)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUpdatePaymentToken is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdff521a8.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function updatePaymentToken(address paymentToken, uint128 numer, uint128 denom) returns()
func (standardRentPriceOracle *StandardRentPriceOracle) TryPackUpdatePaymentToken(paymentToken common.Address, numer *big.Int, denom *big.Int) ([]byte, error) {
	return standardRentPriceOracle.abi.Pack("updatePaymentToken", paymentToken, numer, denom)
}

// StandardRentPriceOracleEACRolesChanged represents a EACRolesChanged event raised by the StandardRentPriceOracle contract.
type StandardRentPriceOracleEACRolesChanged struct {
	Resource      *big.Int
	Account       common.Address
	OldRoleBitmap *big.Int
	NewRoleBitmap *big.Int
	Raw           *types.Log // Blockchain specific contextual infos
}

const StandardRentPriceOracleEACRolesChangedEventName = "EACRolesChanged"

// ContractEventName returns the user-defined event name.
func (StandardRentPriceOracleEACRolesChanged) ContractEventName() string {
	return StandardRentPriceOracleEACRolesChangedEventName
}

// UnpackEACRolesChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event EACRolesChanged(uint256 indexed resource, address indexed account, uint256 oldRoleBitmap, uint256 newRoleBitmap)
func (standardRentPriceOracle *StandardRentPriceOracle) UnpackEACRolesChangedEvent(log *types.Log) (*StandardRentPriceOracleEACRolesChanged, error) {
	event := "EACRolesChanged"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != standardRentPriceOracle.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(StandardRentPriceOracleEACRolesChanged)
	if len(log.Data) > 0 {
		if err := standardRentPriceOracle.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range standardRentPriceOracle.abi.Events[event].Inputs {
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

// StandardRentPriceOraclePaymentTokenUpdated represents a PaymentTokenUpdated event raised by the StandardRentPriceOracle contract.
type StandardRentPriceOraclePaymentTokenUpdated struct {
	PaymentToken common.Address
	Numer        *big.Int
	Denom        *big.Int
	Raw          *types.Log // Blockchain specific contextual infos
}

const StandardRentPriceOraclePaymentTokenUpdatedEventName = "PaymentTokenUpdated"

// ContractEventName returns the user-defined event name.
func (StandardRentPriceOraclePaymentTokenUpdated) ContractEventName() string {
	return StandardRentPriceOraclePaymentTokenUpdatedEventName
}

// UnpackPaymentTokenUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event PaymentTokenUpdated(address indexed paymentToken, uint128 numer, uint128 denom)
func (standardRentPriceOracle *StandardRentPriceOracle) UnpackPaymentTokenUpdatedEvent(log *types.Log) (*StandardRentPriceOraclePaymentTokenUpdated, error) {
	event := "PaymentTokenUpdated"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != standardRentPriceOracle.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(StandardRentPriceOraclePaymentTokenUpdated)
	if len(log.Data) > 0 {
		if err := standardRentPriceOracle.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range standardRentPriceOracle.abi.Events[event].Inputs {
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
func (standardRentPriceOracle *StandardRentPriceOracle) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], standardRentPriceOracle.abi.Errors["EACCannotGrantRoles"].ID.Bytes()[:4]) {
		return standardRentPriceOracle.UnpackEACCannotGrantRolesError(raw[4:])
	}
	if bytes.Equal(raw[:4], standardRentPriceOracle.abi.Errors["EACCannotRevokeRoles"].ID.Bytes()[:4]) {
		return standardRentPriceOracle.UnpackEACCannotRevokeRolesError(raw[4:])
	}
	if bytes.Equal(raw[:4], standardRentPriceOracle.abi.Errors["EACInvalidAccount"].ID.Bytes()[:4]) {
		return standardRentPriceOracle.UnpackEACInvalidAccountError(raw[4:])
	}
	if bytes.Equal(raw[:4], standardRentPriceOracle.abi.Errors["EACInvalidRoleBitmap"].ID.Bytes()[:4]) {
		return standardRentPriceOracle.UnpackEACInvalidRoleBitmapError(raw[4:])
	}
	if bytes.Equal(raw[:4], standardRentPriceOracle.abi.Errors["EACMaxAssignees"].ID.Bytes()[:4]) {
		return standardRentPriceOracle.UnpackEACMaxAssigneesError(raw[4:])
	}
	if bytes.Equal(raw[:4], standardRentPriceOracle.abi.Errors["EACMinAssignees"].ID.Bytes()[:4]) {
		return standardRentPriceOracle.UnpackEACMinAssigneesError(raw[4:])
	}
	if bytes.Equal(raw[:4], standardRentPriceOracle.abi.Errors["EACRootResourceNotAllowed"].ID.Bytes()[:4]) {
		return standardRentPriceOracle.UnpackEACRootResourceNotAllowedError(raw[4:])
	}
	if bytes.Equal(raw[:4], standardRentPriceOracle.abi.Errors["EACUnauthorizedAccountRoles"].ID.Bytes()[:4]) {
		return standardRentPriceOracle.UnpackEACUnauthorizedAccountRolesError(raw[4:])
	}
	if bytes.Equal(raw[:4], standardRentPriceOracle.abi.Errors["InvalidBaseRates"].ID.Bytes()[:4]) {
		return standardRentPriceOracle.UnpackInvalidBaseRatesError(raw[4:])
	}
	if bytes.Equal(raw[:4], standardRentPriceOracle.abi.Errors["InvalidDiscount"].ID.Bytes()[:4]) {
		return standardRentPriceOracle.UnpackInvalidDiscountError(raw[4:])
	}
	if bytes.Equal(raw[:4], standardRentPriceOracle.abi.Errors["InvalidRatio"].ID.Bytes()[:4]) {
		return standardRentPriceOracle.UnpackInvalidRatioError(raw[4:])
	}
	if bytes.Equal(raw[:4], standardRentPriceOracle.abi.Errors["NotValid"].ID.Bytes()[:4]) {
		return standardRentPriceOracle.UnpackNotValidError(raw[4:])
	}
	if bytes.Equal(raw[:4], standardRentPriceOracle.abi.Errors["PaymentTokenNotSupported"].ID.Bytes()[:4]) {
		return standardRentPriceOracle.UnpackPaymentTokenNotSupportedError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// StandardRentPriceOracleEACCannotGrantRoles represents a EACCannotGrantRoles error raised by the StandardRentPriceOracle contract.
type StandardRentPriceOracleEACCannotGrantRoles struct {
	Resource   *big.Int
	RoleBitmap *big.Int
	Account    common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error EACCannotGrantRoles(uint256 resource, uint256 roleBitmap, address account)
func StandardRentPriceOracleEACCannotGrantRolesErrorID() common.Hash {
	return common.HexToHash("0xd1a3b355154fa30324f856ca6b25179bcc6c9d2936e28429b1b4293509e97f4a")
}

// UnpackEACCannotGrantRolesError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error EACCannotGrantRoles(uint256 resource, uint256 roleBitmap, address account)
func (standardRentPriceOracle *StandardRentPriceOracle) UnpackEACCannotGrantRolesError(raw []byte) (*StandardRentPriceOracleEACCannotGrantRoles, error) {
	out := new(StandardRentPriceOracleEACCannotGrantRoles)
	if err := standardRentPriceOracle.abi.UnpackIntoInterface(out, "EACCannotGrantRoles", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// StandardRentPriceOracleEACCannotRevokeRoles represents a EACCannotRevokeRoles error raised by the StandardRentPriceOracle contract.
type StandardRentPriceOracleEACCannotRevokeRoles struct {
	Resource   *big.Int
	RoleBitmap *big.Int
	Account    common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error EACCannotRevokeRoles(uint256 resource, uint256 roleBitmap, address account)
func StandardRentPriceOracleEACCannotRevokeRolesErrorID() common.Hash {
	return common.HexToHash("0xa604e318421f3889d41e189ce5a1bbb63773eb6844a6d371417f9eed8f6cb153")
}

// UnpackEACCannotRevokeRolesError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error EACCannotRevokeRoles(uint256 resource, uint256 roleBitmap, address account)
func (standardRentPriceOracle *StandardRentPriceOracle) UnpackEACCannotRevokeRolesError(raw []byte) (*StandardRentPriceOracleEACCannotRevokeRoles, error) {
	out := new(StandardRentPriceOracleEACCannotRevokeRoles)
	if err := standardRentPriceOracle.abi.UnpackIntoInterface(out, "EACCannotRevokeRoles", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// StandardRentPriceOracleEACInvalidAccount represents a EACInvalidAccount error raised by the StandardRentPriceOracle contract.
type StandardRentPriceOracleEACInvalidAccount struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error EACInvalidAccount()
func StandardRentPriceOracleEACInvalidAccountErrorID() common.Hash {
	return common.HexToHash("0xec3fc5924c67c1c1507473878ec47d207e41f4b88da4d7a835c6e5043b63da51")
}

// UnpackEACInvalidAccountError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error EACInvalidAccount()
func (standardRentPriceOracle *StandardRentPriceOracle) UnpackEACInvalidAccountError(raw []byte) (*StandardRentPriceOracleEACInvalidAccount, error) {
	out := new(StandardRentPriceOracleEACInvalidAccount)
	if err := standardRentPriceOracle.abi.UnpackIntoInterface(out, "EACInvalidAccount", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// StandardRentPriceOracleEACInvalidRoleBitmap represents a EACInvalidRoleBitmap error raised by the StandardRentPriceOracle contract.
type StandardRentPriceOracleEACInvalidRoleBitmap struct {
	RoleBitmap *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error EACInvalidRoleBitmap(uint256 roleBitmap)
func StandardRentPriceOracleEACInvalidRoleBitmapErrorID() common.Hash {
	return common.HexToHash("0x2a7b2d20191e44fd4e4227272caf529612678184b58adb291afe5231bf5b2734")
}

// UnpackEACInvalidRoleBitmapError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error EACInvalidRoleBitmap(uint256 roleBitmap)
func (standardRentPriceOracle *StandardRentPriceOracle) UnpackEACInvalidRoleBitmapError(raw []byte) (*StandardRentPriceOracleEACInvalidRoleBitmap, error) {
	out := new(StandardRentPriceOracleEACInvalidRoleBitmap)
	if err := standardRentPriceOracle.abi.UnpackIntoInterface(out, "EACInvalidRoleBitmap", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// StandardRentPriceOracleEACMaxAssignees represents a EACMaxAssignees error raised by the StandardRentPriceOracle contract.
type StandardRentPriceOracleEACMaxAssignees struct {
	Resource *big.Int
	Role     *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error EACMaxAssignees(uint256 resource, uint256 role)
func StandardRentPriceOracleEACMaxAssigneesErrorID() common.Hash {
	return common.HexToHash("0xf91653484f73ddf3683475f443ccba4f564362c5896702ddf28889a64960b109")
}

// UnpackEACMaxAssigneesError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error EACMaxAssignees(uint256 resource, uint256 role)
func (standardRentPriceOracle *StandardRentPriceOracle) UnpackEACMaxAssigneesError(raw []byte) (*StandardRentPriceOracleEACMaxAssignees, error) {
	out := new(StandardRentPriceOracleEACMaxAssignees)
	if err := standardRentPriceOracle.abi.UnpackIntoInterface(out, "EACMaxAssignees", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// StandardRentPriceOracleEACMinAssignees represents a EACMinAssignees error raised by the StandardRentPriceOracle contract.
type StandardRentPriceOracleEACMinAssignees struct {
	Resource *big.Int
	Role     *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error EACMinAssignees(uint256 resource, uint256 role)
func StandardRentPriceOracleEACMinAssigneesErrorID() common.Hash {
	return common.HexToHash("0x1f80c19ba45f0369eb8ace81d631dde7c47ae7c144205c7e77e1555ff7bb091b")
}

// UnpackEACMinAssigneesError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error EACMinAssignees(uint256 resource, uint256 role)
func (standardRentPriceOracle *StandardRentPriceOracle) UnpackEACMinAssigneesError(raw []byte) (*StandardRentPriceOracleEACMinAssignees, error) {
	out := new(StandardRentPriceOracleEACMinAssignees)
	if err := standardRentPriceOracle.abi.UnpackIntoInterface(out, "EACMinAssignees", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// StandardRentPriceOracleEACRootResourceNotAllowed represents a EACRootResourceNotAllowed error raised by the StandardRentPriceOracle contract.
type StandardRentPriceOracleEACRootResourceNotAllowed struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error EACRootResourceNotAllowed()
func StandardRentPriceOracleEACRootResourceNotAllowedErrorID() common.Hash {
	return common.HexToHash("0xc2842458b56deb79dcc9694a98069c21c928dfa5456787ba7ae1728a0d863238")
}

// UnpackEACRootResourceNotAllowedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error EACRootResourceNotAllowed()
func (standardRentPriceOracle *StandardRentPriceOracle) UnpackEACRootResourceNotAllowedError(raw []byte) (*StandardRentPriceOracleEACRootResourceNotAllowed, error) {
	out := new(StandardRentPriceOracleEACRootResourceNotAllowed)
	if err := standardRentPriceOracle.abi.UnpackIntoInterface(out, "EACRootResourceNotAllowed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// StandardRentPriceOracleEACUnauthorizedAccountRoles represents a EACUnauthorizedAccountRoles error raised by the StandardRentPriceOracle contract.
type StandardRentPriceOracleEACUnauthorizedAccountRoles struct {
	Resource   *big.Int
	RoleBitmap *big.Int
	Account    common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error EACUnauthorizedAccountRoles(uint256 resource, uint256 roleBitmap, address account)
func StandardRentPriceOracleEACUnauthorizedAccountRolesErrorID() common.Hash {
	return common.HexToHash("0x4b27a13377de460ddbf323b4163e94f59f4092e2fe9befdabaf9bd06c2dd939b")
}

// UnpackEACUnauthorizedAccountRolesError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error EACUnauthorizedAccountRoles(uint256 resource, uint256 roleBitmap, address account)
func (standardRentPriceOracle *StandardRentPriceOracle) UnpackEACUnauthorizedAccountRolesError(raw []byte) (*StandardRentPriceOracleEACUnauthorizedAccountRoles, error) {
	out := new(StandardRentPriceOracleEACUnauthorizedAccountRoles)
	if err := standardRentPriceOracle.abi.UnpackIntoInterface(out, "EACUnauthorizedAccountRoles", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// StandardRentPriceOracleInvalidBaseRates represents a InvalidBaseRates error raised by the StandardRentPriceOracle contract.
type StandardRentPriceOracleInvalidBaseRates struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidBaseRates()
func StandardRentPriceOracleInvalidBaseRatesErrorID() common.Hash {
	return common.HexToHash("0xde276447951957ad91a7574738e436e196ede3e15019fb01eba310fe5a26dea6")
}

// UnpackInvalidBaseRatesError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidBaseRates()
func (standardRentPriceOracle *StandardRentPriceOracle) UnpackInvalidBaseRatesError(raw []byte) (*StandardRentPriceOracleInvalidBaseRates, error) {
	out := new(StandardRentPriceOracleInvalidBaseRates)
	if err := standardRentPriceOracle.abi.UnpackIntoInterface(out, "InvalidBaseRates", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// StandardRentPriceOracleInvalidDiscount represents a InvalidDiscount error raised by the StandardRentPriceOracle contract.
type StandardRentPriceOracleInvalidDiscount struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidDiscount()
func StandardRentPriceOracleInvalidDiscountErrorID() common.Hash {
	return common.HexToHash("0x997ea3601b175042a225bcd305c8e0434449bb795e78e4cb3f93a718c1a88caf")
}

// UnpackInvalidDiscountError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidDiscount()
func (standardRentPriceOracle *StandardRentPriceOracle) UnpackInvalidDiscountError(raw []byte) (*StandardRentPriceOracleInvalidDiscount, error) {
	out := new(StandardRentPriceOracleInvalidDiscount)
	if err := standardRentPriceOracle.abi.UnpackIntoInterface(out, "InvalidDiscount", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// StandardRentPriceOracleInvalidRatio represents a InvalidRatio error raised by the StandardRentPriceOracle contract.
type StandardRentPriceOracleInvalidRatio struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidRatio()
func StandardRentPriceOracleInvalidRatioErrorID() common.Hash {
	return common.HexToHash("0x648564d30353f59fffa3a83ca2af1eb61f1bc5278aee48327e5101c81cc1b9f3")
}

// UnpackInvalidRatioError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidRatio()
func (standardRentPriceOracle *StandardRentPriceOracle) UnpackInvalidRatioError(raw []byte) (*StandardRentPriceOracleInvalidRatio, error) {
	out := new(StandardRentPriceOracleInvalidRatio)
	if err := standardRentPriceOracle.abi.UnpackIntoInterface(out, "InvalidRatio", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// StandardRentPriceOracleNotValid represents a NotValid error raised by the StandardRentPriceOracle contract.
type StandardRentPriceOracleNotValid struct {
	Label string
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotValid(string label)
func StandardRentPriceOracleNotValidErrorID() common.Hash {
	return common.HexToHash("0xdbfa28867b89f5fb7cc22f5a333f6ab7c5a65daecdcecdb7af6cb31f0ecf1149")
}

// UnpackNotValidError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotValid(string label)
func (standardRentPriceOracle *StandardRentPriceOracle) UnpackNotValidError(raw []byte) (*StandardRentPriceOracleNotValid, error) {
	out := new(StandardRentPriceOracleNotValid)
	if err := standardRentPriceOracle.abi.UnpackIntoInterface(out, "NotValid", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// StandardRentPriceOraclePaymentTokenNotSupported represents a PaymentTokenNotSupported error raised by the StandardRentPriceOracle contract.
type StandardRentPriceOraclePaymentTokenNotSupported struct {
	PaymentToken common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error PaymentTokenNotSupported(address paymentToken)
func StandardRentPriceOraclePaymentTokenNotSupportedErrorID() common.Hash {
	return common.HexToHash("0x02e2ae9e114104839024ac8d95818f45f9c85ac6717c08caab7252ffbb93537b")
}

// UnpackPaymentTokenNotSupportedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error PaymentTokenNotSupported(address paymentToken)
func (standardRentPriceOracle *StandardRentPriceOracle) UnpackPaymentTokenNotSupportedError(raw []byte) (*StandardRentPriceOraclePaymentTokenNotSupported, error) {
	out := new(StandardRentPriceOraclePaymentTokenNotSupported)
	if err := standardRentPriceOracle.abi.UnpackIntoInterface(out, "PaymentTokenNotSupported", raw); err != nil {
		return nil, err
	}
	return out, nil
}
