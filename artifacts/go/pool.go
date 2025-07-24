// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pool

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

// PoolLLMMeta is an auto generated low-level Go binding around an user-defined struct.
type PoolLLMMeta struct {
	TokenId               *big.Int
	AuthorRoyalty         *big.Int
	AuthorWallet          common.Address
	HwPricePerInputToken  *big.Int
	HwPricePerOutputToken *big.Int
	PayType               uint8
}

// PoolMetaData contains all meta data concerning the Pool contract.
var PoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"llm_nft\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"credit_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"hfswm_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"worker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isApproved\",\"type\":\"bool\"}],\"name\":\"NewWorker\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"llm_id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Payout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"llm_id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"worker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"llmInputTokens\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"llmOutputTokens\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"processingTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Response\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"token_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"hw_price_in\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"hw_price_out\",\"type\":\"uint256\"}],\"name\":\"AddModel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"AddToHFWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"lai_public_key\",\"type\":\"string\"}],\"name\":\"ApproveWorker\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"lai_pub_key\",\"type\":\"string\"}],\"name\":\"Ban\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"hw_cost\",\"type\":\"uint256\"}],\"name\":\"CalculateServiceFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"llm_id\",\"type\":\"uint256\"}],\"name\":\"DeleteModel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DepositCredit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"llm_id\",\"type\":\"uint256\"}],\"name\":\"GetMetaLLM\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"token_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"author_royalty\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"author_wallet\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"hw_price_per_input_token\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"hw_price_per_output_token\",\"type\":\"uint256\"},{\"internalType\":\"enumPool.Pay_type\",\"name\":\"pay_type_\",\"type\":\"uint8\"}],\"internalType\":\"structPool.LLM_meta\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"llm_id\",\"type\":\"uint256\"}],\"name\":\"GetModel\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"token_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"author_royalty\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"author_wallet\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"hw_price_per_input_token\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"hw_price_per_output_token\",\"type\":\"uint256\"},{\"internalType\":\"enumPool.Pay_type\",\"name\":\"pay_type_\",\"type\":\"uint8\"}],\"internalType\":\"structPool.LLM_meta\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"llm_id\",\"type\":\"uint256\"}],\"name\":\"GetTotalPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"lai_id\",\"type\":\"string\"}],\"name\":\"GetWorkerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"HFwhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"llm_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"Pre_request\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"request_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"worker_id\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"llm_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"llmInputTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"llmOutputTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"processingTime\",\"type\":\"uint256\"}],\"name\":\"ProcessResponse\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"lai_public_key\",\"type\":\"string\"}],\"name\":\"RegisterWorker\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"RemoveFromHFWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"lai_pub_key\",\"type\":\"string\"}],\"name\":\"Unban\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"lai_pub_key\",\"type\":\"string\"}],\"name\":\"Unregister\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"llm_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"new_royalty\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"new_address\",\"type\":\"address\"}],\"name\":\"UpdateModel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"old_lai_pub_key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"new_lai_pub_key\",\"type\":\"string\"}],\"name\":\"UpdateWorker\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"blacklist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hfswm\",\"outputs\":[{\"internalType\":\"contractCred\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"isApproved\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PoolABI is the input ABI used to generate the binding from.
// Deprecated: Use PoolMetaData.ABI instead.
var PoolABI = PoolMetaData.ABI

// Pool is an auto generated Go binding around an Ethereum contract.
type Pool struct {
	PoolCaller     // Read-only binding to the contract
	PoolTransactor // Write-only binding to the contract
	PoolFilterer   // Log filterer for contract events
}

// PoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type PoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PoolSession struct {
	Contract     *Pool             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PoolCallerSession struct {
	Contract *PoolCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PoolTransactorSession struct {
	Contract     *PoolTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type PoolRaw struct {
	Contract *Pool // Generic contract binding to access the raw methods on
}

// PoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PoolCallerRaw struct {
	Contract *PoolCaller // Generic read-only contract binding to access the raw methods on
}

// PoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PoolTransactorRaw struct {
	Contract *PoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPool creates a new instance of Pool, bound to a specific deployed contract.
func NewPool(address common.Address, backend bind.ContractBackend) (*Pool, error) {
	contract, err := bindPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pool{PoolCaller: PoolCaller{contract: contract}, PoolTransactor: PoolTransactor{contract: contract}, PoolFilterer: PoolFilterer{contract: contract}}, nil
}

// NewPoolCaller creates a new read-only instance of Pool, bound to a specific deployed contract.
func NewPoolCaller(address common.Address, caller bind.ContractCaller) (*PoolCaller, error) {
	contract, err := bindPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PoolCaller{contract: contract}, nil
}

// NewPoolTransactor creates a new write-only instance of Pool, bound to a specific deployed contract.
func NewPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*PoolTransactor, error) {
	contract, err := bindPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PoolTransactor{contract: contract}, nil
}

// NewPoolFilterer creates a new log filterer instance of Pool, bound to a specific deployed contract.
func NewPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*PoolFilterer, error) {
	contract, err := bindPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PoolFilterer{contract: contract}, nil
}

// bindPool binds a generic wrapper to an already deployed contract.
func bindPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pool *PoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pool.Contract.PoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pool *PoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.Contract.PoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pool *PoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pool.Contract.PoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pool *PoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pool *PoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pool *PoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pool.Contract.contract.Transact(opts, method, params...)
}

// CalculateServiceFee is a free data retrieval call binding the contract method 0x0e5fd6d7.
//
// Solidity: function CalculateServiceFee(uint256 hw_cost) view returns(uint256)
func (_Pool *PoolCaller) CalculateServiceFee(opts *bind.CallOpts, hw_cost *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "CalculateServiceFee", hw_cost)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateServiceFee is a free data retrieval call binding the contract method 0x0e5fd6d7.
//
// Solidity: function CalculateServiceFee(uint256 hw_cost) view returns(uint256)
func (_Pool *PoolSession) CalculateServiceFee(hw_cost *big.Int) (*big.Int, error) {
	return _Pool.Contract.CalculateServiceFee(&_Pool.CallOpts, hw_cost)
}

// CalculateServiceFee is a free data retrieval call binding the contract method 0x0e5fd6d7.
//
// Solidity: function CalculateServiceFee(uint256 hw_cost) view returns(uint256)
func (_Pool *PoolCallerSession) CalculateServiceFee(hw_cost *big.Int) (*big.Int, error) {
	return _Pool.Contract.CalculateServiceFee(&_Pool.CallOpts, hw_cost)
}

// GetMetaLLM is a free data retrieval call binding the contract method 0xad7922d5.
//
// Solidity: function GetMetaLLM(uint256 llm_id) view returns((uint256,uint256,address,uint256,uint256,uint8))
func (_Pool *PoolCaller) GetMetaLLM(opts *bind.CallOpts, llm_id *big.Int) (PoolLLMMeta, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "GetMetaLLM", llm_id)

	if err != nil {
		return *new(PoolLLMMeta), err
	}

	out0 := *abi.ConvertType(out[0], new(PoolLLMMeta)).(*PoolLLMMeta)

	return out0, err

}

// GetMetaLLM is a free data retrieval call binding the contract method 0xad7922d5.
//
// Solidity: function GetMetaLLM(uint256 llm_id) view returns((uint256,uint256,address,uint256,uint256,uint8))
func (_Pool *PoolSession) GetMetaLLM(llm_id *big.Int) (PoolLLMMeta, error) {
	return _Pool.Contract.GetMetaLLM(&_Pool.CallOpts, llm_id)
}

// GetMetaLLM is a free data retrieval call binding the contract method 0xad7922d5.
//
// Solidity: function GetMetaLLM(uint256 llm_id) view returns((uint256,uint256,address,uint256,uint256,uint8))
func (_Pool *PoolCallerSession) GetMetaLLM(llm_id *big.Int) (PoolLLMMeta, error) {
	return _Pool.Contract.GetMetaLLM(&_Pool.CallOpts, llm_id)
}

// GetModel is a free data retrieval call binding the contract method 0xbbfb8caa.
//
// Solidity: function GetModel(uint256 llm_id) view returns((uint256,uint256,address,uint256,uint256,uint8))
func (_Pool *PoolCaller) GetModel(opts *bind.CallOpts, llm_id *big.Int) (PoolLLMMeta, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "GetModel", llm_id)

	if err != nil {
		return *new(PoolLLMMeta), err
	}

	out0 := *abi.ConvertType(out[0], new(PoolLLMMeta)).(*PoolLLMMeta)

	return out0, err

}

// GetModel is a free data retrieval call binding the contract method 0xbbfb8caa.
//
// Solidity: function GetModel(uint256 llm_id) view returns((uint256,uint256,address,uint256,uint256,uint8))
func (_Pool *PoolSession) GetModel(llm_id *big.Int) (PoolLLMMeta, error) {
	return _Pool.Contract.GetModel(&_Pool.CallOpts, llm_id)
}

// GetModel is a free data retrieval call binding the contract method 0xbbfb8caa.
//
// Solidity: function GetModel(uint256 llm_id) view returns((uint256,uint256,address,uint256,uint256,uint8))
func (_Pool *PoolCallerSession) GetModel(llm_id *big.Int) (PoolLLMMeta, error) {
	return _Pool.Contract.GetModel(&_Pool.CallOpts, llm_id)
}

// GetTotalPrice is a free data retrieval call binding the contract method 0x97c37d6b.
//
// Solidity: function GetTotalPrice(uint256 llm_id) view returns(uint256)
func (_Pool *PoolCaller) GetTotalPrice(opts *bind.CallOpts, llm_id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "GetTotalPrice", llm_id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalPrice is a free data retrieval call binding the contract method 0x97c37d6b.
//
// Solidity: function GetTotalPrice(uint256 llm_id) view returns(uint256)
func (_Pool *PoolSession) GetTotalPrice(llm_id *big.Int) (*big.Int, error) {
	return _Pool.Contract.GetTotalPrice(&_Pool.CallOpts, llm_id)
}

// GetTotalPrice is a free data retrieval call binding the contract method 0x97c37d6b.
//
// Solidity: function GetTotalPrice(uint256 llm_id) view returns(uint256)
func (_Pool *PoolCallerSession) GetTotalPrice(llm_id *big.Int) (*big.Int, error) {
	return _Pool.Contract.GetTotalPrice(&_Pool.CallOpts, llm_id)
}

// GetWorkerAddress is a free data retrieval call binding the contract method 0x421d2777.
//
// Solidity: function GetWorkerAddress(string lai_id) view returns(address)
func (_Pool *PoolCaller) GetWorkerAddress(opts *bind.CallOpts, lai_id string) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "GetWorkerAddress", lai_id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetWorkerAddress is a free data retrieval call binding the contract method 0x421d2777.
//
// Solidity: function GetWorkerAddress(string lai_id) view returns(address)
func (_Pool *PoolSession) GetWorkerAddress(lai_id string) (common.Address, error) {
	return _Pool.Contract.GetWorkerAddress(&_Pool.CallOpts, lai_id)
}

// GetWorkerAddress is a free data retrieval call binding the contract method 0x421d2777.
//
// Solidity: function GetWorkerAddress(string lai_id) view returns(address)
func (_Pool *PoolCallerSession) GetWorkerAddress(lai_id string) (common.Address, error) {
	return _Pool.Contract.GetWorkerAddress(&_Pool.CallOpts, lai_id)
}

// HFwhitelist is a free data retrieval call binding the contract method 0xb083de54.
//
// Solidity: function HFwhitelist(address ) view returns(bool)
func (_Pool *PoolCaller) HFwhitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "HFwhitelist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HFwhitelist is a free data retrieval call binding the contract method 0xb083de54.
//
// Solidity: function HFwhitelist(address ) view returns(bool)
func (_Pool *PoolSession) HFwhitelist(arg0 common.Address) (bool, error) {
	return _Pool.Contract.HFwhitelist(&_Pool.CallOpts, arg0)
}

// HFwhitelist is a free data retrieval call binding the contract method 0xb083de54.
//
// Solidity: function HFwhitelist(address ) view returns(bool)
func (_Pool *PoolCallerSession) HFwhitelist(arg0 common.Address) (bool, error) {
	return _Pool.Contract.HFwhitelist(&_Pool.CallOpts, arg0)
}

// PreRequest is a free data retrieval call binding the contract method 0xb9359ea2.
//
// Solidity: function Pre_request(uint256 llm_id, address user) view returns()
func (_Pool *PoolCaller) PreRequest(opts *bind.CallOpts, llm_id *big.Int, user common.Address) error {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "Pre_request", llm_id, user)

	if err != nil {
		return err
	}

	return err

}

// PreRequest is a free data retrieval call binding the contract method 0xb9359ea2.
//
// Solidity: function Pre_request(uint256 llm_id, address user) view returns()
func (_Pool *PoolSession) PreRequest(llm_id *big.Int, user common.Address) error {
	return _Pool.Contract.PreRequest(&_Pool.CallOpts, llm_id, user)
}

// PreRequest is a free data retrieval call binding the contract method 0xb9359ea2.
//
// Solidity: function Pre_request(uint256 llm_id, address user) view returns()
func (_Pool *PoolCallerSession) PreRequest(llm_id *big.Int, user common.Address) error {
	return _Pool.Contract.PreRequest(&_Pool.CallOpts, llm_id, user)
}

// Blacklist is a free data retrieval call binding the contract method 0x8c0233e9.
//
// Solidity: function blacklist(string ) view returns(bool)
func (_Pool *PoolCaller) Blacklist(opts *bind.CallOpts, arg0 string) (bool, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "blacklist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Blacklist is a free data retrieval call binding the contract method 0x8c0233e9.
//
// Solidity: function blacklist(string ) view returns(bool)
func (_Pool *PoolSession) Blacklist(arg0 string) (bool, error) {
	return _Pool.Contract.Blacklist(&_Pool.CallOpts, arg0)
}

// Blacklist is a free data retrieval call binding the contract method 0x8c0233e9.
//
// Solidity: function blacklist(string ) view returns(bool)
func (_Pool *PoolCallerSession) Blacklist(arg0 string) (bool, error) {
	return _Pool.Contract.Blacklist(&_Pool.CallOpts, arg0)
}

// Hfswm is a free data retrieval call binding the contract method 0xe03ec511.
//
// Solidity: function hfswm() view returns(address)
func (_Pool *PoolCaller) Hfswm(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "hfswm")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Hfswm is a free data retrieval call binding the contract method 0xe03ec511.
//
// Solidity: function hfswm() view returns(address)
func (_Pool *PoolSession) Hfswm() (common.Address, error) {
	return _Pool.Contract.Hfswm(&_Pool.CallOpts)
}

// Hfswm is a free data retrieval call binding the contract method 0xe03ec511.
//
// Solidity: function hfswm() view returns(address)
func (_Pool *PoolCallerSession) Hfswm() (common.Address, error) {
	return _Pool.Contract.Hfswm(&_Pool.CallOpts)
}

// IsApproved is a free data retrieval call binding the contract method 0xf31fbc9f.
//
// Solidity: function isApproved(string ) view returns(bool)
func (_Pool *PoolCaller) IsApproved(opts *bind.CallOpts, arg0 string) (bool, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "isApproved", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApproved is a free data retrieval call binding the contract method 0xf31fbc9f.
//
// Solidity: function isApproved(string ) view returns(bool)
func (_Pool *PoolSession) IsApproved(arg0 string) (bool, error) {
	return _Pool.Contract.IsApproved(&_Pool.CallOpts, arg0)
}

// IsApproved is a free data retrieval call binding the contract method 0xf31fbc9f.
//
// Solidity: function isApproved(string ) view returns(bool)
func (_Pool *PoolCallerSession) IsApproved(arg0 string) (bool, error) {
	return _Pool.Contract.IsApproved(&_Pool.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pool *PoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pool *PoolSession) Owner() (common.Address, error) {
	return _Pool.Contract.Owner(&_Pool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pool *PoolCallerSession) Owner() (common.Address, error) {
	return _Pool.Contract.Owner(&_Pool.CallOpts)
}

// AddModel is a paid mutator transaction binding the contract method 0xaa389692.
//
// Solidity: function AddModel(uint256 token_id, uint256 hw_price_in, uint256 hw_price_out) returns()
func (_Pool *PoolTransactor) AddModel(opts *bind.TransactOpts, token_id *big.Int, hw_price_in *big.Int, hw_price_out *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "AddModel", token_id, hw_price_in, hw_price_out)
}

// AddModel is a paid mutator transaction binding the contract method 0xaa389692.
//
// Solidity: function AddModel(uint256 token_id, uint256 hw_price_in, uint256 hw_price_out) returns()
func (_Pool *PoolSession) AddModel(token_id *big.Int, hw_price_in *big.Int, hw_price_out *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.AddModel(&_Pool.TransactOpts, token_id, hw_price_in, hw_price_out)
}

// AddModel is a paid mutator transaction binding the contract method 0xaa389692.
//
// Solidity: function AddModel(uint256 token_id, uint256 hw_price_in, uint256 hw_price_out) returns()
func (_Pool *PoolTransactorSession) AddModel(token_id *big.Int, hw_price_in *big.Int, hw_price_out *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.AddModel(&_Pool.TransactOpts, token_id, hw_price_in, hw_price_out)
}

// AddToHFWhitelist is a paid mutator transaction binding the contract method 0xba68ce9c.
//
// Solidity: function AddToHFWhitelist(address addr) returns()
func (_Pool *PoolTransactor) AddToHFWhitelist(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "AddToHFWhitelist", addr)
}

// AddToHFWhitelist is a paid mutator transaction binding the contract method 0xba68ce9c.
//
// Solidity: function AddToHFWhitelist(address addr) returns()
func (_Pool *PoolSession) AddToHFWhitelist(addr common.Address) (*types.Transaction, error) {
	return _Pool.Contract.AddToHFWhitelist(&_Pool.TransactOpts, addr)
}

// AddToHFWhitelist is a paid mutator transaction binding the contract method 0xba68ce9c.
//
// Solidity: function AddToHFWhitelist(address addr) returns()
func (_Pool *PoolTransactorSession) AddToHFWhitelist(addr common.Address) (*types.Transaction, error) {
	return _Pool.Contract.AddToHFWhitelist(&_Pool.TransactOpts, addr)
}

// ApproveWorker is a paid mutator transaction binding the contract method 0x431f90f3.
//
// Solidity: function ApproveWorker(string lai_public_key) returns()
func (_Pool *PoolTransactor) ApproveWorker(opts *bind.TransactOpts, lai_public_key string) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "ApproveWorker", lai_public_key)
}

// ApproveWorker is a paid mutator transaction binding the contract method 0x431f90f3.
//
// Solidity: function ApproveWorker(string lai_public_key) returns()
func (_Pool *PoolSession) ApproveWorker(lai_public_key string) (*types.Transaction, error) {
	return _Pool.Contract.ApproveWorker(&_Pool.TransactOpts, lai_public_key)
}

// ApproveWorker is a paid mutator transaction binding the contract method 0x431f90f3.
//
// Solidity: function ApproveWorker(string lai_public_key) returns()
func (_Pool *PoolTransactorSession) ApproveWorker(lai_public_key string) (*types.Transaction, error) {
	return _Pool.Contract.ApproveWorker(&_Pool.TransactOpts, lai_public_key)
}

// Ban is a paid mutator transaction binding the contract method 0xf6bd2ac8.
//
// Solidity: function Ban(string lai_pub_key) returns()
func (_Pool *PoolTransactor) Ban(opts *bind.TransactOpts, lai_pub_key string) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "Ban", lai_pub_key)
}

// Ban is a paid mutator transaction binding the contract method 0xf6bd2ac8.
//
// Solidity: function Ban(string lai_pub_key) returns()
func (_Pool *PoolSession) Ban(lai_pub_key string) (*types.Transaction, error) {
	return _Pool.Contract.Ban(&_Pool.TransactOpts, lai_pub_key)
}

// Ban is a paid mutator transaction binding the contract method 0xf6bd2ac8.
//
// Solidity: function Ban(string lai_pub_key) returns()
func (_Pool *PoolTransactorSession) Ban(lai_pub_key string) (*types.Transaction, error) {
	return _Pool.Contract.Ban(&_Pool.TransactOpts, lai_pub_key)
}

// DeleteModel is a paid mutator transaction binding the contract method 0x27801534.
//
// Solidity: function DeleteModel(uint256 llm_id) returns()
func (_Pool *PoolTransactor) DeleteModel(opts *bind.TransactOpts, llm_id *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "DeleteModel", llm_id)
}

// DeleteModel is a paid mutator transaction binding the contract method 0x27801534.
//
// Solidity: function DeleteModel(uint256 llm_id) returns()
func (_Pool *PoolSession) DeleteModel(llm_id *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.DeleteModel(&_Pool.TransactOpts, llm_id)
}

// DeleteModel is a paid mutator transaction binding the contract method 0x27801534.
//
// Solidity: function DeleteModel(uint256 llm_id) returns()
func (_Pool *PoolTransactorSession) DeleteModel(llm_id *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.DeleteModel(&_Pool.TransactOpts, llm_id)
}

// DepositCredit is a paid mutator transaction binding the contract method 0x1b988ce4.
//
// Solidity: function DepositCredit(uint256 amount) returns()
func (_Pool *PoolTransactor) DepositCredit(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "DepositCredit", amount)
}

// DepositCredit is a paid mutator transaction binding the contract method 0x1b988ce4.
//
// Solidity: function DepositCredit(uint256 amount) returns()
func (_Pool *PoolSession) DepositCredit(amount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.DepositCredit(&_Pool.TransactOpts, amount)
}

// DepositCredit is a paid mutator transaction binding the contract method 0x1b988ce4.
//
// Solidity: function DepositCredit(uint256 amount) returns()
func (_Pool *PoolTransactorSession) DepositCredit(amount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.DepositCredit(&_Pool.TransactOpts, amount)
}

// ProcessResponse is a paid mutator transaction binding the contract method 0xc7567cd7.
//
// Solidity: function ProcessResponse(uint256 request_id, string worker_id, uint256 llm_id, uint256 llmInputTokens, uint256 llmOutputTokens, uint256 processingTime) returns()
func (_Pool *PoolTransactor) ProcessResponse(opts *bind.TransactOpts, request_id *big.Int, worker_id string, llm_id *big.Int, llmInputTokens *big.Int, llmOutputTokens *big.Int, processingTime *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "ProcessResponse", request_id, worker_id, llm_id, llmInputTokens, llmOutputTokens, processingTime)
}

// ProcessResponse is a paid mutator transaction binding the contract method 0xc7567cd7.
//
// Solidity: function ProcessResponse(uint256 request_id, string worker_id, uint256 llm_id, uint256 llmInputTokens, uint256 llmOutputTokens, uint256 processingTime) returns()
func (_Pool *PoolSession) ProcessResponse(request_id *big.Int, worker_id string, llm_id *big.Int, llmInputTokens *big.Int, llmOutputTokens *big.Int, processingTime *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.ProcessResponse(&_Pool.TransactOpts, request_id, worker_id, llm_id, llmInputTokens, llmOutputTokens, processingTime)
}

// ProcessResponse is a paid mutator transaction binding the contract method 0xc7567cd7.
//
// Solidity: function ProcessResponse(uint256 request_id, string worker_id, uint256 llm_id, uint256 llmInputTokens, uint256 llmOutputTokens, uint256 processingTime) returns()
func (_Pool *PoolTransactorSession) ProcessResponse(request_id *big.Int, worker_id string, llm_id *big.Int, llmInputTokens *big.Int, llmOutputTokens *big.Int, processingTime *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.ProcessResponse(&_Pool.TransactOpts, request_id, worker_id, llm_id, llmInputTokens, llmOutputTokens, processingTime)
}

// RegisterWorker is a paid mutator transaction binding the contract method 0x9fcd8c9d.
//
// Solidity: function RegisterWorker(string lai_public_key) returns()
func (_Pool *PoolTransactor) RegisterWorker(opts *bind.TransactOpts, lai_public_key string) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "RegisterWorker", lai_public_key)
}

// RegisterWorker is a paid mutator transaction binding the contract method 0x9fcd8c9d.
//
// Solidity: function RegisterWorker(string lai_public_key) returns()
func (_Pool *PoolSession) RegisterWorker(lai_public_key string) (*types.Transaction, error) {
	return _Pool.Contract.RegisterWorker(&_Pool.TransactOpts, lai_public_key)
}

// RegisterWorker is a paid mutator transaction binding the contract method 0x9fcd8c9d.
//
// Solidity: function RegisterWorker(string lai_public_key) returns()
func (_Pool *PoolTransactorSession) RegisterWorker(lai_public_key string) (*types.Transaction, error) {
	return _Pool.Contract.RegisterWorker(&_Pool.TransactOpts, lai_public_key)
}

// RemoveFromHFWhitelist is a paid mutator transaction binding the contract method 0x00345d3d.
//
// Solidity: function RemoveFromHFWhitelist(address addr) returns()
func (_Pool *PoolTransactor) RemoveFromHFWhitelist(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "RemoveFromHFWhitelist", addr)
}

// RemoveFromHFWhitelist is a paid mutator transaction binding the contract method 0x00345d3d.
//
// Solidity: function RemoveFromHFWhitelist(address addr) returns()
func (_Pool *PoolSession) RemoveFromHFWhitelist(addr common.Address) (*types.Transaction, error) {
	return _Pool.Contract.RemoveFromHFWhitelist(&_Pool.TransactOpts, addr)
}

// RemoveFromHFWhitelist is a paid mutator transaction binding the contract method 0x00345d3d.
//
// Solidity: function RemoveFromHFWhitelist(address addr) returns()
func (_Pool *PoolTransactorSession) RemoveFromHFWhitelist(addr common.Address) (*types.Transaction, error) {
	return _Pool.Contract.RemoveFromHFWhitelist(&_Pool.TransactOpts, addr)
}

// Unban is a paid mutator transaction binding the contract method 0xbad3191a.
//
// Solidity: function Unban(string lai_pub_key) returns()
func (_Pool *PoolTransactor) Unban(opts *bind.TransactOpts, lai_pub_key string) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "Unban", lai_pub_key)
}

// Unban is a paid mutator transaction binding the contract method 0xbad3191a.
//
// Solidity: function Unban(string lai_pub_key) returns()
func (_Pool *PoolSession) Unban(lai_pub_key string) (*types.Transaction, error) {
	return _Pool.Contract.Unban(&_Pool.TransactOpts, lai_pub_key)
}

// Unban is a paid mutator transaction binding the contract method 0xbad3191a.
//
// Solidity: function Unban(string lai_pub_key) returns()
func (_Pool *PoolTransactorSession) Unban(lai_pub_key string) (*types.Transaction, error) {
	return _Pool.Contract.Unban(&_Pool.TransactOpts, lai_pub_key)
}

// Unregister is a paid mutator transaction binding the contract method 0x1e99f349.
//
// Solidity: function Unregister(string lai_pub_key) returns()
func (_Pool *PoolTransactor) Unregister(opts *bind.TransactOpts, lai_pub_key string) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "Unregister", lai_pub_key)
}

// Unregister is a paid mutator transaction binding the contract method 0x1e99f349.
//
// Solidity: function Unregister(string lai_pub_key) returns()
func (_Pool *PoolSession) Unregister(lai_pub_key string) (*types.Transaction, error) {
	return _Pool.Contract.Unregister(&_Pool.TransactOpts, lai_pub_key)
}

// Unregister is a paid mutator transaction binding the contract method 0x1e99f349.
//
// Solidity: function Unregister(string lai_pub_key) returns()
func (_Pool *PoolTransactorSession) Unregister(lai_pub_key string) (*types.Transaction, error) {
	return _Pool.Contract.Unregister(&_Pool.TransactOpts, lai_pub_key)
}

// UpdateModel is a paid mutator transaction binding the contract method 0x3266e290.
//
// Solidity: function UpdateModel(uint256 llm_id, uint256 new_royalty, address new_address) returns()
func (_Pool *PoolTransactor) UpdateModel(opts *bind.TransactOpts, llm_id *big.Int, new_royalty *big.Int, new_address common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "UpdateModel", llm_id, new_royalty, new_address)
}

// UpdateModel is a paid mutator transaction binding the contract method 0x3266e290.
//
// Solidity: function UpdateModel(uint256 llm_id, uint256 new_royalty, address new_address) returns()
func (_Pool *PoolSession) UpdateModel(llm_id *big.Int, new_royalty *big.Int, new_address common.Address) (*types.Transaction, error) {
	return _Pool.Contract.UpdateModel(&_Pool.TransactOpts, llm_id, new_royalty, new_address)
}

// UpdateModel is a paid mutator transaction binding the contract method 0x3266e290.
//
// Solidity: function UpdateModel(uint256 llm_id, uint256 new_royalty, address new_address) returns()
func (_Pool *PoolTransactorSession) UpdateModel(llm_id *big.Int, new_royalty *big.Int, new_address common.Address) (*types.Transaction, error) {
	return _Pool.Contract.UpdateModel(&_Pool.TransactOpts, llm_id, new_royalty, new_address)
}

// UpdateWorker is a paid mutator transaction binding the contract method 0x8f977f36.
//
// Solidity: function UpdateWorker(string old_lai_pub_key, string new_lai_pub_key) returns()
func (_Pool *PoolTransactor) UpdateWorker(opts *bind.TransactOpts, old_lai_pub_key string, new_lai_pub_key string) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "UpdateWorker", old_lai_pub_key, new_lai_pub_key)
}

// UpdateWorker is a paid mutator transaction binding the contract method 0x8f977f36.
//
// Solidity: function UpdateWorker(string old_lai_pub_key, string new_lai_pub_key) returns()
func (_Pool *PoolSession) UpdateWorker(old_lai_pub_key string, new_lai_pub_key string) (*types.Transaction, error) {
	return _Pool.Contract.UpdateWorker(&_Pool.TransactOpts, old_lai_pub_key, new_lai_pub_key)
}

// UpdateWorker is a paid mutator transaction binding the contract method 0x8f977f36.
//
// Solidity: function UpdateWorker(string old_lai_pub_key, string new_lai_pub_key) returns()
func (_Pool *PoolTransactorSession) UpdateWorker(old_lai_pub_key string, new_lai_pub_key string) (*types.Transaction, error) {
	return _Pool.Contract.UpdateWorker(&_Pool.TransactOpts, old_lai_pub_key, new_lai_pub_key)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Pool *PoolTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Pool *PoolSession) RenounceOwnership() (*types.Transaction, error) {
	return _Pool.Contract.RenounceOwnership(&_Pool.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Pool *PoolTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Pool.Contract.RenounceOwnership(&_Pool.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Pool *PoolTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Pool *PoolSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Pool.Contract.TransferOwnership(&_Pool.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Pool *PoolTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Pool.Contract.TransferOwnership(&_Pool.TransactOpts, newOwner)
}

// PoolDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Pool contract.
type PoolDepositIterator struct {
	Event *PoolDeposit // Event containing the contract specifics and raw log

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
func (it *PoolDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolDeposit)
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
		it.Event = new(PoolDeposit)
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
func (it *PoolDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolDeposit represents a Deposit event raised by the Pool contract.
type PoolDeposit struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address user, uint256 amount)
func (_Pool *PoolFilterer) FilterDeposit(opts *bind.FilterOpts) (*PoolDepositIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &PoolDepositIterator{contract: _Pool.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address user, uint256 amount)
func (_Pool *PoolFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *PoolDeposit) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolDeposit)
				if err := _Pool.contract.UnpackLog(event, "Deposit", log); err != nil {
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
// Solidity: event Deposit(address user, uint256 amount)
func (_Pool *PoolFilterer) ParseDeposit(log types.Log) (*PoolDeposit, error) {
	event := new(PoolDeposit)
	if err := _Pool.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolNewWorkerIterator is returned from FilterNewWorker and is used to iterate over the raw logs and unpacked data for NewWorker events raised by the Pool contract.
type PoolNewWorkerIterator struct {
	Event *PoolNewWorker // Event containing the contract specifics and raw log

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
func (it *PoolNewWorkerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolNewWorker)
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
		it.Event = new(PoolNewWorker)
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
func (it *PoolNewWorkerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolNewWorkerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolNewWorker represents a NewWorker event raised by the Pool contract.
type PoolNewWorker struct {
	Worker     common.Address
	IsApproved bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewWorker is a free log retrieval operation binding the contract event 0xee06bd32f8b2583fd50c0dd94d872913c9fe17bb486e0f9d4b4865752d79efa0.
//
// Solidity: event NewWorker(address worker, bool isApproved)
func (_Pool *PoolFilterer) FilterNewWorker(opts *bind.FilterOpts) (*PoolNewWorkerIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "NewWorker")
	if err != nil {
		return nil, err
	}
	return &PoolNewWorkerIterator{contract: _Pool.contract, event: "NewWorker", logs: logs, sub: sub}, nil
}

// WatchNewWorker is a free log subscription operation binding the contract event 0xee06bd32f8b2583fd50c0dd94d872913c9fe17bb486e0f9d4b4865752d79efa0.
//
// Solidity: event NewWorker(address worker, bool isApproved)
func (_Pool *PoolFilterer) WatchNewWorker(opts *bind.WatchOpts, sink chan<- *PoolNewWorker) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "NewWorker")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolNewWorker)
				if err := _Pool.contract.UnpackLog(event, "NewWorker", log); err != nil {
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

// ParseNewWorker is a log parse operation binding the contract event 0xee06bd32f8b2583fd50c0dd94d872913c9fe17bb486e0f9d4b4865752d79efa0.
//
// Solidity: event NewWorker(address worker, bool isApproved)
func (_Pool *PoolFilterer) ParseNewWorker(log types.Log) (*PoolNewWorker, error) {
	event := new(PoolNewWorker)
	if err := _Pool.contract.UnpackLog(event, "NewWorker", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Pool contract.
type PoolOwnershipTransferredIterator struct {
	Event *PoolOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PoolOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolOwnershipTransferred)
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
		it.Event = new(PoolOwnershipTransferred)
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
func (it *PoolOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolOwnershipTransferred represents a OwnershipTransferred event raised by the Pool contract.
type PoolOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Pool *PoolFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PoolOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PoolOwnershipTransferredIterator{contract: _Pool.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Pool *PoolFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PoolOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolOwnershipTransferred)
				if err := _Pool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Pool *PoolFilterer) ParseOwnershipTransferred(log types.Log) (*PoolOwnershipTransferred, error) {
	event := new(PoolOwnershipTransferred)
	if err := _Pool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolPayoutIterator is returned from FilterPayout and is used to iterate over the raw logs and unpacked data for Payout events raised by the Pool contract.
type PoolPayoutIterator struct {
	Event *PoolPayout // Event containing the contract specifics and raw log

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
func (it *PoolPayoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolPayout)
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
		it.Event = new(PoolPayout)
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
func (it *PoolPayoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolPayoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolPayout represents a Payout event raised by the Pool contract.
type PoolPayout struct {
	Provider common.Address
	LlmId    *big.Int
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPayout is a free log retrieval operation binding the contract event 0x634235fcf5af0adbca1a405ec65f6f6c08f55e1f379c2c45cd10f23cb29e0e31.
//
// Solidity: event Payout(address provider, uint256 llm_id, uint256 amount)
func (_Pool *PoolFilterer) FilterPayout(opts *bind.FilterOpts) (*PoolPayoutIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "Payout")
	if err != nil {
		return nil, err
	}
	return &PoolPayoutIterator{contract: _Pool.contract, event: "Payout", logs: logs, sub: sub}, nil
}

// WatchPayout is a free log subscription operation binding the contract event 0x634235fcf5af0adbca1a405ec65f6f6c08f55e1f379c2c45cd10f23cb29e0e31.
//
// Solidity: event Payout(address provider, uint256 llm_id, uint256 amount)
func (_Pool *PoolFilterer) WatchPayout(opts *bind.WatchOpts, sink chan<- *PoolPayout) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "Payout")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolPayout)
				if err := _Pool.contract.UnpackLog(event, "Payout", log); err != nil {
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

// ParsePayout is a log parse operation binding the contract event 0x634235fcf5af0adbca1a405ec65f6f6c08f55e1f379c2c45cd10f23cb29e0e31.
//
// Solidity: event Payout(address provider, uint256 llm_id, uint256 amount)
func (_Pool *PoolFilterer) ParsePayout(log types.Log) (*PoolPayout, error) {
	event := new(PoolPayout)
	if err := _Pool.contract.UnpackLog(event, "Payout", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolResponseIterator is returned from FilterResponse and is used to iterate over the raw logs and unpacked data for Response events raised by the Pool contract.
type PoolResponseIterator struct {
	Event *PoolResponse // Event containing the contract specifics and raw log

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
func (it *PoolResponseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolResponse)
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
		it.Event = new(PoolResponse)
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
func (it *PoolResponseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolResponseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolResponse represents a Response event raised by the Pool contract.
type PoolResponse struct {
	Id              *big.Int
	LlmId           *big.Int
	Worker          common.Address
	LlmInputTokens  *big.Int
	LlmOutputTokens *big.Int
	Cost            *big.Int
	ProcessingTime  *big.Int
	Timestamp       *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterResponse is a free log retrieval operation binding the contract event 0xe682774c095c98275661d05ff664616a641de1592fa52e2412789919e84b01ad.
//
// Solidity: event Response(uint256 id, uint256 llm_id, address worker, uint256 llmInputTokens, uint256 llmOutputTokens, uint256 cost, uint256 processingTime, uint256 timestamp)
func (_Pool *PoolFilterer) FilterResponse(opts *bind.FilterOpts) (*PoolResponseIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "Response")
	if err != nil {
		return nil, err
	}
	return &PoolResponseIterator{contract: _Pool.contract, event: "Response", logs: logs, sub: sub}, nil
}

// WatchResponse is a free log subscription operation binding the contract event 0xe682774c095c98275661d05ff664616a641de1592fa52e2412789919e84b01ad.
//
// Solidity: event Response(uint256 id, uint256 llm_id, address worker, uint256 llmInputTokens, uint256 llmOutputTokens, uint256 cost, uint256 processingTime, uint256 timestamp)
func (_Pool *PoolFilterer) WatchResponse(opts *bind.WatchOpts, sink chan<- *PoolResponse) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "Response")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolResponse)
				if err := _Pool.contract.UnpackLog(event, "Response", log); err != nil {
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

// ParseResponse is a log parse operation binding the contract event 0xe682774c095c98275661d05ff664616a641de1592fa52e2412789919e84b01ad.
//
// Solidity: event Response(uint256 id, uint256 llm_id, address worker, uint256 llmInputTokens, uint256 llmOutputTokens, uint256 cost, uint256 processingTime, uint256 timestamp)
func (_Pool *PoolFilterer) ParseResponse(log types.Log) (*PoolResponse, error) {
	event := new(PoolResponse)
	if err := _Pool.contract.UnpackLog(event, "Response", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
