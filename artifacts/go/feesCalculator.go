// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package feesCalculator

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// FeesCalculatorMetaData contains all meta data concerning the FeesCalculator contract.
var FeesCalculatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"scale\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"promille_fee_\",\"type\":\"uint256\"}],\"name\":\"CalculateAbstractFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// FeesCalculatorABI is the input ABI used to generate the binding from.
// Deprecated: Use FeesCalculatorMetaData.ABI instead.
var FeesCalculatorABI = FeesCalculatorMetaData.ABI

// FeesCalculator is an auto generated Go binding around an Ethereum contract.
type FeesCalculator struct {
	FeesCalculatorCaller     // Read-only binding to the contract
	FeesCalculatorTransactor // Write-only binding to the contract
	FeesCalculatorFilterer   // Log filterer for contract events
}

// FeesCalculatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type FeesCalculatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeesCalculatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FeesCalculatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeesCalculatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FeesCalculatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeesCalculatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FeesCalculatorSession struct {
	Contract     *FeesCalculator   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FeesCalculatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FeesCalculatorCallerSession struct {
	Contract *FeesCalculatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// FeesCalculatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FeesCalculatorTransactorSession struct {
	Contract     *FeesCalculatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// FeesCalculatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type FeesCalculatorRaw struct {
	Contract *FeesCalculator // Generic contract binding to access the raw methods on
}

// FeesCalculatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FeesCalculatorCallerRaw struct {
	Contract *FeesCalculatorCaller // Generic read-only contract binding to access the raw methods on
}

// FeesCalculatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FeesCalculatorTransactorRaw struct {
	Contract *FeesCalculatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFeesCalculator creates a new instance of FeesCalculator, bound to a specific deployed contract.
func NewFeesCalculator(address common.Address, backend bind.ContractBackend) (*FeesCalculator, error) {
	contract, err := bindFeesCalculator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FeesCalculator{FeesCalculatorCaller: FeesCalculatorCaller{contract: contract}, FeesCalculatorTransactor: FeesCalculatorTransactor{contract: contract}, FeesCalculatorFilterer: FeesCalculatorFilterer{contract: contract}}, nil
}

// NewFeesCalculatorCaller creates a new read-only instance of FeesCalculator, bound to a specific deployed contract.
func NewFeesCalculatorCaller(address common.Address, caller bind.ContractCaller) (*FeesCalculatorCaller, error) {
	contract, err := bindFeesCalculator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FeesCalculatorCaller{contract: contract}, nil
}

// NewFeesCalculatorTransactor creates a new write-only instance of FeesCalculator, bound to a specific deployed contract.
func NewFeesCalculatorTransactor(address common.Address, transactor bind.ContractTransactor) (*FeesCalculatorTransactor, error) {
	contract, err := bindFeesCalculator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FeesCalculatorTransactor{contract: contract}, nil
}

// NewFeesCalculatorFilterer creates a new log filterer instance of FeesCalculator, bound to a specific deployed contract.
func NewFeesCalculatorFilterer(address common.Address, filterer bind.ContractFilterer) (*FeesCalculatorFilterer, error) {
	contract, err := bindFeesCalculator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FeesCalculatorFilterer{contract: contract}, nil
}

// bindFeesCalculator binds a generic wrapper to an already deployed contract.
func bindFeesCalculator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FeesCalculatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeesCalculator *FeesCalculatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeesCalculator.Contract.FeesCalculatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeesCalculator *FeesCalculatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeesCalculator.Contract.FeesCalculatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeesCalculator *FeesCalculatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeesCalculator.Contract.FeesCalculatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeesCalculator *FeesCalculatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeesCalculator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeesCalculator *FeesCalculatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeesCalculator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeesCalculator *FeesCalculatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeesCalculator.Contract.contract.Transact(opts, method, params...)
}

// CalculateAbstractFee is a free data retrieval call binding the contract method 0xa285b894.
//
// Solidity: function CalculateAbstractFee(uint256 amount, uint256 scale, uint256 promille_fee_) pure returns(uint256)
func (_FeesCalculator *FeesCalculatorCaller) CalculateAbstractFee(opts *bind.CallOpts, amount *big.Int, scale *big.Int, promille_fee_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FeesCalculator.contract.Call(opts, &out, "CalculateAbstractFee", amount, scale, promille_fee_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateAbstractFee is a free data retrieval call binding the contract method 0xa285b894.
//
// Solidity: function CalculateAbstractFee(uint256 amount, uint256 scale, uint256 promille_fee_) pure returns(uint256)
func (_FeesCalculator *FeesCalculatorSession) CalculateAbstractFee(amount *big.Int, scale *big.Int, promille_fee_ *big.Int) (*big.Int, error) {
	return _FeesCalculator.Contract.CalculateAbstractFee(&_FeesCalculator.CallOpts, amount, scale, promille_fee_)
}

// CalculateAbstractFee is a free data retrieval call binding the contract method 0xa285b894.
//
// Solidity: function CalculateAbstractFee(uint256 amount, uint256 scale, uint256 promille_fee_) pure returns(uint256)
func (_FeesCalculator *FeesCalculatorCallerSession) CalculateAbstractFee(amount *big.Int, scale *big.Int, promille_fee_ *big.Int) (*big.Int, error) {
	return _FeesCalculator.Contract.CalculateAbstractFee(&_FeesCalculator.CallOpts, amount, scale, promille_fee_)
}
