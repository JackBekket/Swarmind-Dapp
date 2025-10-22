// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package WalletHub

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
	_ = abi.ConvertType
)

// WalletHubMetaData contains all meta data concerning the WalletHub contract.
var WalletHubMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"credit_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"treasury_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"Credit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"Debit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"addPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"credit\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"debit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"pay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pools\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"removePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTreasury\",\"type\":\"address\"}],\"name\":\"setTreasury\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"treasuryTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"userBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// WalletHubABI is the input ABI used to generate the binding from.
// Deprecated: Use WalletHubMetaData.ABI instead.
var WalletHubABI = WalletHubMetaData.ABI

// WalletHub is an auto generated Go binding around an Ethereum contract.
type WalletHub struct {
	WalletHubCaller     // Read-only binding to the contract
	WalletHubTransactor // Write-only binding to the contract
	WalletHubFilterer   // Log filterer for contract events
}

// WalletHubCaller is an auto generated read-only Go binding around an Ethereum contract.
type WalletHubCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletHubTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WalletHubTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletHubFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WalletHubFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletHubSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WalletHubSession struct {
	Contract     *WalletHub        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WalletHubCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WalletHubCallerSession struct {
	Contract *WalletHubCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// WalletHubTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WalletHubTransactorSession struct {
	Contract     *WalletHubTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// WalletHubRaw is an auto generated low-level Go binding around an Ethereum contract.
type WalletHubRaw struct {
	Contract *WalletHub // Generic contract binding to access the raw methods on
}

// WalletHubCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WalletHubCallerRaw struct {
	Contract *WalletHubCaller // Generic read-only contract binding to access the raw methods on
}

// WalletHubTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WalletHubTransactorRaw struct {
	Contract *WalletHubTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWalletHub creates a new instance of WalletHub, bound to a specific deployed contract.
func NewWalletHub(address common.Address, backend bind.ContractBackend) (*WalletHub, error) {
	contract, err := bindWalletHub(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WalletHub{WalletHubCaller: WalletHubCaller{contract: contract}, WalletHubTransactor: WalletHubTransactor{contract: contract}, WalletHubFilterer: WalletHubFilterer{contract: contract}}, nil
}

// NewWalletHubCaller creates a new read-only instance of WalletHub, bound to a specific deployed contract.
func NewWalletHubCaller(address common.Address, caller bind.ContractCaller) (*WalletHubCaller, error) {
	contract, err := bindWalletHub(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WalletHubCaller{contract: contract}, nil
}

// NewWalletHubTransactor creates a new write-only instance of WalletHub, bound to a specific deployed contract.
func NewWalletHubTransactor(address common.Address, transactor bind.ContractTransactor) (*WalletHubTransactor, error) {
	contract, err := bindWalletHub(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WalletHubTransactor{contract: contract}, nil
}

// NewWalletHubFilterer creates a new log filterer instance of WalletHub, bound to a specific deployed contract.
func NewWalletHubFilterer(address common.Address, filterer bind.ContractFilterer) (*WalletHubFilterer, error) {
	contract, err := bindWalletHub(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WalletHubFilterer{contract: contract}, nil
}

// bindWalletHub binds a generic wrapper to an already deployed contract.
func bindWalletHub(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WalletHubMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WalletHub *WalletHubRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WalletHub.Contract.WalletHubCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WalletHub *WalletHubRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletHub.Contract.WalletHubTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WalletHub *WalletHubRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WalletHub.Contract.WalletHubTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WalletHub *WalletHubCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WalletHub.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WalletHub *WalletHubTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletHub.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WalletHub *WalletHubTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WalletHub.Contract.contract.Transact(opts, method, params...)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_WalletHub *WalletHubCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WalletHub.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_WalletHub *WalletHubSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _WalletHub.Contract.Balances(&_WalletHub.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_WalletHub *WalletHubCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _WalletHub.Contract.Balances(&_WalletHub.CallOpts, arg0)
}

// Credit is a free data retrieval call binding the contract method 0xa06d083c.
//
// Solidity: function credit() view returns(address)
func (_WalletHub *WalletHubCaller) Credit(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WalletHub.contract.Call(opts, &out, "credit")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Credit is a free data retrieval call binding the contract method 0xa06d083c.
//
// Solidity: function credit() view returns(address)
func (_WalletHub *WalletHubSession) Credit() (common.Address, error) {
	return _WalletHub.Contract.Credit(&_WalletHub.CallOpts)
}

// Credit is a free data retrieval call binding the contract method 0xa06d083c.
//
// Solidity: function credit() view returns(address)
func (_WalletHub *WalletHubCallerSession) Credit() (common.Address, error) {
	return _WalletHub.Contract.Credit(&_WalletHub.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WalletHub *WalletHubCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WalletHub.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WalletHub *WalletHubSession) Owner() (common.Address, error) {
	return _WalletHub.Contract.Owner(&_WalletHub.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WalletHub *WalletHubCallerSession) Owner() (common.Address, error) {
	return _WalletHub.Contract.Owner(&_WalletHub.CallOpts)
}

// Pools is a free data retrieval call binding the contract method 0xa4063dbc.
//
// Solidity: function pools(address ) view returns(bool)
func (_WalletHub *WalletHubCaller) Pools(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _WalletHub.contract.Call(opts, &out, "pools", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Pools is a free data retrieval call binding the contract method 0xa4063dbc.
//
// Solidity: function pools(address ) view returns(bool)
func (_WalletHub *WalletHubSession) Pools(arg0 common.Address) (bool, error) {
	return _WalletHub.Contract.Pools(&_WalletHub.CallOpts, arg0)
}

// Pools is a free data retrieval call binding the contract method 0xa4063dbc.
//
// Solidity: function pools(address ) view returns(bool)
func (_WalletHub *WalletHubCallerSession) Pools(arg0 common.Address) (bool, error) {
	return _WalletHub.Contract.Pools(&_WalletHub.CallOpts, arg0)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_WalletHub *WalletHubCaller) Treasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WalletHub.contract.Call(opts, &out, "treasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_WalletHub *WalletHubSession) Treasury() (common.Address, error) {
	return _WalletHub.Contract.Treasury(&_WalletHub.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_WalletHub *WalletHubCallerSession) Treasury() (common.Address, error) {
	return _WalletHub.Contract.Treasury(&_WalletHub.CallOpts)
}

// UserBalance is a free data retrieval call binding the contract method 0x0103c92b.
//
// Solidity: function userBalance(address user) view returns(uint256)
func (_WalletHub *WalletHubCaller) UserBalance(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WalletHub.contract.Call(opts, &out, "userBalance", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserBalance is a free data retrieval call binding the contract method 0x0103c92b.
//
// Solidity: function userBalance(address user) view returns(uint256)
func (_WalletHub *WalletHubSession) UserBalance(user common.Address) (*big.Int, error) {
	return _WalletHub.Contract.UserBalance(&_WalletHub.CallOpts, user)
}

// UserBalance is a free data retrieval call binding the contract method 0x0103c92b.
//
// Solidity: function userBalance(address user) view returns(uint256)
func (_WalletHub *WalletHubCallerSession) UserBalance(user common.Address) (*big.Int, error) {
	return _WalletHub.Contract.UserBalance(&_WalletHub.CallOpts, user)
}

// AddPool is a paid mutator transaction binding the contract method 0xd914cd4b.
//
// Solidity: function addPool(address pool) returns()
func (_WalletHub *WalletHubTransactor) AddPool(opts *bind.TransactOpts, pool common.Address) (*types.Transaction, error) {
	return _WalletHub.contract.Transact(opts, "addPool", pool)
}

// AddPool is a paid mutator transaction binding the contract method 0xd914cd4b.
//
// Solidity: function addPool(address pool) returns()
func (_WalletHub *WalletHubSession) AddPool(pool common.Address) (*types.Transaction, error) {
	return _WalletHub.Contract.AddPool(&_WalletHub.TransactOpts, pool)
}

// AddPool is a paid mutator transaction binding the contract method 0xd914cd4b.
//
// Solidity: function addPool(address pool) returns()
func (_WalletHub *WalletHubTransactorSession) AddPool(pool common.Address) (*types.Transaction, error) {
	return _WalletHub.Contract.AddPool(&_WalletHub.TransactOpts, pool)
}

// Debit is a paid mutator transaction binding the contract method 0x636be27a.
//
// Solidity: function debit(address user, uint256 amount) returns()
func (_WalletHub *WalletHubTransactor) Debit(opts *bind.TransactOpts, user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WalletHub.contract.Transact(opts, "debit", user, amount)
}

// Debit is a paid mutator transaction binding the contract method 0x636be27a.
//
// Solidity: function debit(address user, uint256 amount) returns()
func (_WalletHub *WalletHubSession) Debit(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WalletHub.Contract.Debit(&_WalletHub.TransactOpts, user, amount)
}

// Debit is a paid mutator transaction binding the contract method 0x636be27a.
//
// Solidity: function debit(address user, uint256 amount) returns()
func (_WalletHub *WalletHubTransactorSession) Debit(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WalletHub.Contract.Debit(&_WalletHub.TransactOpts, user, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_WalletHub *WalletHubTransactor) Deposit(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _WalletHub.contract.Transact(opts, "deposit", amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_WalletHub *WalletHubSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _WalletHub.Contract.Deposit(&_WalletHub.TransactOpts, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_WalletHub *WalletHubTransactorSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _WalletHub.Contract.Deposit(&_WalletHub.TransactOpts, amount)
}

// Pay is a paid mutator transaction binding the contract method 0xc4076876.
//
// Solidity: function pay(address to, uint256 amount) returns()
func (_WalletHub *WalletHubTransactor) Pay(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WalletHub.contract.Transact(opts, "pay", to, amount)
}

// Pay is a paid mutator transaction binding the contract method 0xc4076876.
//
// Solidity: function pay(address to, uint256 amount) returns()
func (_WalletHub *WalletHubSession) Pay(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WalletHub.Contract.Pay(&_WalletHub.TransactOpts, to, amount)
}

// Pay is a paid mutator transaction binding the contract method 0xc4076876.
//
// Solidity: function pay(address to, uint256 amount) returns()
func (_WalletHub *WalletHubTransactorSession) Pay(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WalletHub.Contract.Pay(&_WalletHub.TransactOpts, to, amount)
}

// RemovePool is a paid mutator transaction binding the contract method 0x3b7d0946.
//
// Solidity: function removePool(address pool) returns()
func (_WalletHub *WalletHubTransactor) RemovePool(opts *bind.TransactOpts, pool common.Address) (*types.Transaction, error) {
	return _WalletHub.contract.Transact(opts, "removePool", pool)
}

// RemovePool is a paid mutator transaction binding the contract method 0x3b7d0946.
//
// Solidity: function removePool(address pool) returns()
func (_WalletHub *WalletHubSession) RemovePool(pool common.Address) (*types.Transaction, error) {
	return _WalletHub.Contract.RemovePool(&_WalletHub.TransactOpts, pool)
}

// RemovePool is a paid mutator transaction binding the contract method 0x3b7d0946.
//
// Solidity: function removePool(address pool) returns()
func (_WalletHub *WalletHubTransactorSession) RemovePool(pool common.Address) (*types.Transaction, error) {
	return _WalletHub.Contract.RemovePool(&_WalletHub.TransactOpts, pool)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WalletHub *WalletHubTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletHub.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WalletHub *WalletHubSession) RenounceOwnership() (*types.Transaction, error) {
	return _WalletHub.Contract.RenounceOwnership(&_WalletHub.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WalletHub *WalletHubTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _WalletHub.Contract.RenounceOwnership(&_WalletHub.TransactOpts)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address newTreasury) returns()
func (_WalletHub *WalletHubTransactor) SetTreasury(opts *bind.TransactOpts, newTreasury common.Address) (*types.Transaction, error) {
	return _WalletHub.contract.Transact(opts, "setTreasury", newTreasury)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address newTreasury) returns()
func (_WalletHub *WalletHubSession) SetTreasury(newTreasury common.Address) (*types.Transaction, error) {
	return _WalletHub.Contract.SetTreasury(&_WalletHub.TransactOpts, newTreasury)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address newTreasury) returns()
func (_WalletHub *WalletHubTransactorSession) SetTreasury(newTreasury common.Address) (*types.Transaction, error) {
	return _WalletHub.Contract.SetTreasury(&_WalletHub.TransactOpts, newTreasury)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WalletHub *WalletHubTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _WalletHub.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WalletHub *WalletHubSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WalletHub.Contract.TransferOwnership(&_WalletHub.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WalletHub *WalletHubTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WalletHub.Contract.TransferOwnership(&_WalletHub.TransactOpts, newOwner)
}

// TreasuryTransfer is a paid mutator transaction binding the contract method 0xeffbe39a.
//
// Solidity: function treasuryTransfer(uint256 amount) returns()
func (_WalletHub *WalletHubTransactor) TreasuryTransfer(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _WalletHub.contract.Transact(opts, "treasuryTransfer", amount)
}

// TreasuryTransfer is a paid mutator transaction binding the contract method 0xeffbe39a.
//
// Solidity: function treasuryTransfer(uint256 amount) returns()
func (_WalletHub *WalletHubSession) TreasuryTransfer(amount *big.Int) (*types.Transaction, error) {
	return _WalletHub.Contract.TreasuryTransfer(&_WalletHub.TransactOpts, amount)
}

// TreasuryTransfer is a paid mutator transaction binding the contract method 0xeffbe39a.
//
// Solidity: function treasuryTransfer(uint256 amount) returns()
func (_WalletHub *WalletHubTransactorSession) TreasuryTransfer(amount *big.Int) (*types.Transaction, error) {
	return _WalletHub.Contract.TreasuryTransfer(&_WalletHub.TransactOpts, amount)
}

// WalletHubCreditIterator is returned from FilterCredit and is used to iterate over the raw logs and unpacked data for Credit events raised by the WalletHub contract.
type WalletHubCreditIterator struct {
	Event *WalletHubCredit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WalletHubCreditIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletHubCredit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WalletHubCredit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WalletHubCreditIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletHubCreditIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletHubCredit represents a Credit event raised by the WalletHub contract.
type WalletHubCredit struct {
	Recipient common.Address
	Amount    *big.Int
	Pool      common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCredit is a free log retrieval operation binding the contract event 0xd8850388b259dc88b9d78382c6dc6475dc041b530972549e97d4dc969fbfcc51.
//
// Solidity: event Credit(address indexed recipient, uint256 amount, address indexed pool)
func (_WalletHub *WalletHubFilterer) FilterCredit(opts *bind.FilterOpts, recipient []common.Address, pool []common.Address) (*WalletHubCreditIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _WalletHub.contract.FilterLogs(opts, "Credit", recipientRule, poolRule)
	if err != nil {
		return nil, err
	}
	return &WalletHubCreditIterator{contract: _WalletHub.contract, event: "Credit", logs: logs, sub: sub}, nil
}

// WatchCredit is a free log subscription operation binding the contract event 0xd8850388b259dc88b9d78382c6dc6475dc041b530972549e97d4dc969fbfcc51.
//
// Solidity: event Credit(address indexed recipient, uint256 amount, address indexed pool)
func (_WalletHub *WalletHubFilterer) WatchCredit(opts *bind.WatchOpts, sink chan<- *WalletHubCredit, recipient []common.Address, pool []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _WalletHub.contract.WatchLogs(opts, "Credit", recipientRule, poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletHubCredit)
				if err := _WalletHub.contract.UnpackLog(event, "Credit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCredit is a log parse operation binding the contract event 0xd8850388b259dc88b9d78382c6dc6475dc041b530972549e97d4dc969fbfcc51.
//
// Solidity: event Credit(address indexed recipient, uint256 amount, address indexed pool)
func (_WalletHub *WalletHubFilterer) ParseCredit(log types.Log) (*WalletHubCredit, error) {
	event := new(WalletHubCredit)
	if err := _WalletHub.contract.UnpackLog(event, "Credit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletHubDebitIterator is returned from FilterDebit and is used to iterate over the raw logs and unpacked data for Debit events raised by the WalletHub contract.
type WalletHubDebitIterator struct {
	Event *WalletHubDebit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WalletHubDebitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletHubDebit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WalletHubDebit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WalletHubDebitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletHubDebitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletHubDebit represents a Debit event raised by the WalletHub contract.
type WalletHubDebit struct {
	User   common.Address
	Amount *big.Int
	Pool   common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDebit is a free log retrieval operation binding the contract event 0x966154e1a039c40d0ab727e2d4a0b051386bf0584466eab1254c1374e99e48ed.
//
// Solidity: event Debit(address indexed user, uint256 amount, address indexed pool)
func (_WalletHub *WalletHubFilterer) FilterDebit(opts *bind.FilterOpts, user []common.Address, pool []common.Address) (*WalletHubDebitIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _WalletHub.contract.FilterLogs(opts, "Debit", userRule, poolRule)
	if err != nil {
		return nil, err
	}
	return &WalletHubDebitIterator{contract: _WalletHub.contract, event: "Debit", logs: logs, sub: sub}, nil
}

// WatchDebit is a free log subscription operation binding the contract event 0x966154e1a039c40d0ab727e2d4a0b051386bf0584466eab1254c1374e99e48ed.
//
// Solidity: event Debit(address indexed user, uint256 amount, address indexed pool)
func (_WalletHub *WalletHubFilterer) WatchDebit(opts *bind.WatchOpts, sink chan<- *WalletHubDebit, user []common.Address, pool []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _WalletHub.contract.WatchLogs(opts, "Debit", userRule, poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletHubDebit)
				if err := _WalletHub.contract.UnpackLog(event, "Debit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDebit is a log parse operation binding the contract event 0x966154e1a039c40d0ab727e2d4a0b051386bf0584466eab1254c1374e99e48ed.
//
// Solidity: event Debit(address indexed user, uint256 amount, address indexed pool)
func (_WalletHub *WalletHubFilterer) ParseDebit(log types.Log) (*WalletHubDebit, error) {
	event := new(WalletHubDebit)
	if err := _WalletHub.contract.UnpackLog(event, "Debit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletHubDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the WalletHub contract.
type WalletHubDepositIterator struct {
	Event *WalletHubDeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WalletHubDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletHubDeposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WalletHubDeposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WalletHubDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletHubDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletHubDeposit represents a Deposit event raised by the WalletHub contract.
type WalletHubDeposit struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed user, uint256 amount)
func (_WalletHub *WalletHubFilterer) FilterDeposit(opts *bind.FilterOpts, user []common.Address) (*WalletHubDepositIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _WalletHub.contract.FilterLogs(opts, "Deposit", userRule)
	if err != nil {
		return nil, err
	}
	return &WalletHubDepositIterator{contract: _WalletHub.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed user, uint256 amount)
func (_WalletHub *WalletHubFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *WalletHubDeposit, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _WalletHub.contract.WatchLogs(opts, "Deposit", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletHubDeposit)
				if err := _WalletHub.contract.UnpackLog(event, "Deposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed user, uint256 amount)
func (_WalletHub *WalletHubFilterer) ParseDeposit(log types.Log) (*WalletHubDeposit, error) {
	event := new(WalletHubDeposit)
	if err := _WalletHub.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletHubOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the WalletHub contract.
type WalletHubOwnershipTransferredIterator struct {
	Event *WalletHubOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WalletHubOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletHubOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WalletHubOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WalletHubOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletHubOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletHubOwnershipTransferred represents a OwnershipTransferred event raised by the WalletHub contract.
type WalletHubOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WalletHub *WalletHubFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*WalletHubOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WalletHub.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &WalletHubOwnershipTransferredIterator{contract: _WalletHub.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WalletHub *WalletHubFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *WalletHubOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WalletHub.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletHubOwnershipTransferred)
				if err := _WalletHub.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WalletHub *WalletHubFilterer) ParseOwnershipTransferred(log types.Log) (*WalletHubOwnershipTransferred, error) {
	event := new(WalletHubOwnershipTransferred)
	if err := _WalletHub.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
