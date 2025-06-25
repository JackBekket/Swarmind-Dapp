// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package feesc

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

// FeescMetaData contains all meta data concerning the Feesc contract.
var FeescMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"scale\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"promille_fee_\",\"type\":\"uint256\"}],\"name\":\"CalculateAbstractFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// FeescABI is the input ABI used to generate the binding from.
// Deprecated: Use FeescMetaData.ABI instead.
var FeescABI = FeescMetaData.ABI

// Feesc is an auto generated Go binding around an Ethereum contract.
type Feesc struct {
	FeescCaller     // Read-only binding to the contract
	FeescTransactor // Write-only binding to the contract
	FeescFilterer   // Log filterer for contract events
}

// FeescCaller is an auto generated read-only Go binding around an Ethereum contract.
type FeescCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeescTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FeescTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeescFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FeescFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeescSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FeescSession struct {
	Contract     *Feesc            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FeescCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FeescCallerSession struct {
	Contract *FeescCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// FeescTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FeescTransactorSession struct {
	Contract     *FeescTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FeescRaw is an auto generated low-level Go binding around an Ethereum contract.
type FeescRaw struct {
	Contract *Feesc // Generic contract binding to access the raw methods on
}

// FeescCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FeescCallerRaw struct {
	Contract *FeescCaller // Generic read-only contract binding to access the raw methods on
}

// FeescTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FeescTransactorRaw struct {
	Contract *FeescTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFeesc creates a new instance of Feesc, bound to a specific deployed contract.
func NewFeesc(address common.Address, backend bind.ContractBackend) (*Feesc, error) {
	contract, err := bindFeesc(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Feesc{FeescCaller: FeescCaller{contract: contract}, FeescTransactor: FeescTransactor{contract: contract}, FeescFilterer: FeescFilterer{contract: contract}}, nil
}

// NewFeescCaller creates a new read-only instance of Feesc, bound to a specific deployed contract.
func NewFeescCaller(address common.Address, caller bind.ContractCaller) (*FeescCaller, error) {
	contract, err := bindFeesc(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FeescCaller{contract: contract}, nil
}

// NewFeescTransactor creates a new write-only instance of Feesc, bound to a specific deployed contract.
func NewFeescTransactor(address common.Address, transactor bind.ContractTransactor) (*FeescTransactor, error) {
	contract, err := bindFeesc(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FeescTransactor{contract: contract}, nil
}

// NewFeescFilterer creates a new log filterer instance of Feesc, bound to a specific deployed contract.
func NewFeescFilterer(address common.Address, filterer bind.ContractFilterer) (*FeescFilterer, error) {
	contract, err := bindFeesc(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FeescFilterer{contract: contract}, nil
}

// bindFeesc binds a generic wrapper to an already deployed contract.
func bindFeesc(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FeescABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Feesc *FeescRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Feesc.Contract.FeescCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Feesc *FeescRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Feesc.Contract.FeescTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Feesc *FeescRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Feesc.Contract.FeescTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Feesc *FeescCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Feesc.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Feesc *FeescTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Feesc.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Feesc *FeescTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Feesc.Contract.contract.Transact(opts, method, params...)
}

// CalculateAbstractFee is a free data retrieval call binding the contract method 0xa285b894.
//
// Solidity: function CalculateAbstractFee(uint256 amount, uint256 scale, uint256 promille_fee_) pure returns(uint256)
func (_Feesc *FeescCaller) CalculateAbstractFee(opts *bind.CallOpts, amount *big.Int, scale *big.Int, promille_fee_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Feesc.contract.Call(opts, &out, "CalculateAbstractFee", amount, scale, promille_fee_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateAbstractFee is a free data retrieval call binding the contract method 0xa285b894.
//
// Solidity: function CalculateAbstractFee(uint256 amount, uint256 scale, uint256 promille_fee_) pure returns(uint256)
func (_Feesc *FeescSession) CalculateAbstractFee(amount *big.Int, scale *big.Int, promille_fee_ *big.Int) (*big.Int, error) {
	return _Feesc.Contract.CalculateAbstractFee(&_Feesc.CallOpts, amount, scale, promille_fee_)
}

// CalculateAbstractFee is a free data retrieval call binding the contract method 0xa285b894.
//
// Solidity: function CalculateAbstractFee(uint256 amount, uint256 scale, uint256 promille_fee_) pure returns(uint256)
func (_Feesc *FeescCallerSession) CalculateAbstractFee(amount *big.Int, scale *big.Int, promille_fee_ *big.Int) (*big.Int, error) {
	return _Feesc.Contract.CalculateAbstractFee(&_Feesc.CallOpts, amount, scale, promille_fee_)
}
