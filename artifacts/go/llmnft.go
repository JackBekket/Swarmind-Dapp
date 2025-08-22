// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package LLMNFT

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

// LLMNFTLLM is an auto generated low-level Go binding around an user-defined struct.
type LLMNFTLLM struct {
	Author           string
	Hfid             string
	RoyaltyPrice     *big.Int
	AuthorWallet     common.Address
	ModelType        uint8
	MaxContextWindow *big.Int
}

// LLMNFTMetaData contains all meta data concerning the LLMNFT contract.
var LLMNFTMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"smbl_\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_toTokenId\",\"type\":\"uint256\"}],\"name\":\"BatchMetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"author\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"hfid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"royaltyPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"authorWallet\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumLLMNFT.Llm_type\",\"name\":\"modelType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxContextWindow\",\"type\":\"uint256\"}],\"name\":\"LLMNFTCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"MetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"hf_id_\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"internalType\":\"enumLLMNFT.Llm_type\",\"name\":\"model_type_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"max_context\",\"type\":\"uint256\"}],\"name\":\"CreateLLM_NFT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"token_id\",\"type\":\"uint256\"}],\"name\":\"GetLLM\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"author\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"hfid\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"royalty_price\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"author_wallet\",\"type\":\"address\"},{\"internalType\":\"enumLLMNFT.Llm_type\",\"name\":\"model_type\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"max_context_window\",\"type\":\"uint256\"}],\"internalType\":\"structLLMNFT.LLM\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"Hf_models\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Models_metadata\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"author\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"hfid\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"royalty_price\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"author_wallet\",\"type\":\"address\"},{\"internalType\":\"enumLLMNFT.Llm_type\",\"name\":\"model_type\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"max_context_window\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// LLMNFTABI is the input ABI used to generate the binding from.
// Deprecated: Use LLMNFTMetaData.ABI instead.
var LLMNFTABI = LLMNFTMetaData.ABI

// LLMNFT is an auto generated Go binding around an Ethereum contract.
type LLMNFT struct {
	LLMNFTCaller     // Read-only binding to the contract
	LLMNFTTransactor // Write-only binding to the contract
	LLMNFTFilterer   // Log filterer for contract events
}

// LLMNFTCaller is an auto generated read-only Go binding around an Ethereum contract.
type LLMNFTCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LLMNFTTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LLMNFTTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LLMNFTFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LLMNFTFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LLMNFTSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LLMNFTSession struct {
	Contract     *LLMNFT           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LLMNFTCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LLMNFTCallerSession struct {
	Contract *LLMNFTCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// LLMNFTTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LLMNFTTransactorSession struct {
	Contract     *LLMNFTTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LLMNFTRaw is an auto generated low-level Go binding around an Ethereum contract.
type LLMNFTRaw struct {
	Contract *LLMNFT // Generic contract binding to access the raw methods on
}

// LLMNFTCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LLMNFTCallerRaw struct {
	Contract *LLMNFTCaller // Generic read-only contract binding to access the raw methods on
}

// LLMNFTTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LLMNFTTransactorRaw struct {
	Contract *LLMNFTTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLLMNFT creates a new instance of LLMNFT, bound to a specific deployed contract.
func NewLLMNFT(address common.Address, backend bind.ContractBackend) (*LLMNFT, error) {
	contract, err := bindLLMNFT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LLMNFT{LLMNFTCaller: LLMNFTCaller{contract: contract}, LLMNFTTransactor: LLMNFTTransactor{contract: contract}, LLMNFTFilterer: LLMNFTFilterer{contract: contract}}, nil
}

// NewLLMNFTCaller creates a new read-only instance of LLMNFT, bound to a specific deployed contract.
func NewLLMNFTCaller(address common.Address, caller bind.ContractCaller) (*LLMNFTCaller, error) {
	contract, err := bindLLMNFT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LLMNFTCaller{contract: contract}, nil
}

// NewLLMNFTTransactor creates a new write-only instance of LLMNFT, bound to a specific deployed contract.
func NewLLMNFTTransactor(address common.Address, transactor bind.ContractTransactor) (*LLMNFTTransactor, error) {
	contract, err := bindLLMNFT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LLMNFTTransactor{contract: contract}, nil
}

// NewLLMNFTFilterer creates a new log filterer instance of LLMNFT, bound to a specific deployed contract.
func NewLLMNFTFilterer(address common.Address, filterer bind.ContractFilterer) (*LLMNFTFilterer, error) {
	contract, err := bindLLMNFT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LLMNFTFilterer{contract: contract}, nil
}

// bindLLMNFT binds a generic wrapper to an already deployed contract.
func bindLLMNFT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LLMNFTMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LLMNFT *LLMNFTRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LLMNFT.Contract.LLMNFTCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LLMNFT *LLMNFTRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LLMNFT.Contract.LLMNFTTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LLMNFT *LLMNFTRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LLMNFT.Contract.LLMNFTTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LLMNFT *LLMNFTCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LLMNFT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LLMNFT *LLMNFTTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LLMNFT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LLMNFT *LLMNFTTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LLMNFT.Contract.contract.Transact(opts, method, params...)
}

// GetLLM is a free data retrieval call binding the contract method 0x40a24c67.
//
// Solidity: function GetLLM(uint256 token_id) view returns((string,string,uint256,address,uint8,uint256))
func (_LLMNFT *LLMNFTCaller) GetLLM(opts *bind.CallOpts, token_id *big.Int) (LLMNFTLLM, error) {
	var out []interface{}
	err := _LLMNFT.contract.Call(opts, &out, "GetLLM", token_id)

	if err != nil {
		return *new(LLMNFTLLM), err
	}

	out0 := *abi.ConvertType(out[0], new(LLMNFTLLM)).(*LLMNFTLLM)

	return out0, err

}

// GetLLM is a free data retrieval call binding the contract method 0x40a24c67.
//
// Solidity: function GetLLM(uint256 token_id) view returns((string,string,uint256,address,uint8,uint256))
func (_LLMNFT *LLMNFTSession) GetLLM(token_id *big.Int) (LLMNFTLLM, error) {
	return _LLMNFT.Contract.GetLLM(&_LLMNFT.CallOpts, token_id)
}

// GetLLM is a free data retrieval call binding the contract method 0x40a24c67.
//
// Solidity: function GetLLM(uint256 token_id) view returns((string,string,uint256,address,uint8,uint256))
func (_LLMNFT *LLMNFTCallerSession) GetLLM(token_id *big.Int) (LLMNFTLLM, error) {
	return _LLMNFT.Contract.GetLLM(&_LLMNFT.CallOpts, token_id)
}

// HfModels is a free data retrieval call binding the contract method 0x65783bbb.
//
// Solidity: function Hf_models(string ) view returns(uint256)
func (_LLMNFT *LLMNFTCaller) HfModels(opts *bind.CallOpts, arg0 string) (*big.Int, error) {
	var out []interface{}
	err := _LLMNFT.contract.Call(opts, &out, "Hf_models", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HfModels is a free data retrieval call binding the contract method 0x65783bbb.
//
// Solidity: function Hf_models(string ) view returns(uint256)
func (_LLMNFT *LLMNFTSession) HfModels(arg0 string) (*big.Int, error) {
	return _LLMNFT.Contract.HfModels(&_LLMNFT.CallOpts, arg0)
}

// HfModels is a free data retrieval call binding the contract method 0x65783bbb.
//
// Solidity: function Hf_models(string ) view returns(uint256)
func (_LLMNFT *LLMNFTCallerSession) HfModels(arg0 string) (*big.Int, error) {
	return _LLMNFT.Contract.HfModels(&_LLMNFT.CallOpts, arg0)
}

// ModelsMetadata is a free data retrieval call binding the contract method 0x8a87e958.
//
// Solidity: function Models_metadata(uint256 ) view returns(string author, string hfid, uint256 royalty_price, address author_wallet, uint8 model_type, uint256 max_context_window)
func (_LLMNFT *LLMNFTCaller) ModelsMetadata(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Author           string
	Hfid             string
	RoyaltyPrice     *big.Int
	AuthorWallet     common.Address
	ModelType        uint8
	MaxContextWindow *big.Int
}, error) {
	var out []interface{}
	err := _LLMNFT.contract.Call(opts, &out, "Models_metadata", arg0)

	outstruct := new(struct {
		Author           string
		Hfid             string
		RoyaltyPrice     *big.Int
		AuthorWallet     common.Address
		ModelType        uint8
		MaxContextWindow *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Author = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Hfid = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.RoyaltyPrice = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.AuthorWallet = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.ModelType = *abi.ConvertType(out[4], new(uint8)).(*uint8)
	outstruct.MaxContextWindow = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ModelsMetadata is a free data retrieval call binding the contract method 0x8a87e958.
//
// Solidity: function Models_metadata(uint256 ) view returns(string author, string hfid, uint256 royalty_price, address author_wallet, uint8 model_type, uint256 max_context_window)
func (_LLMNFT *LLMNFTSession) ModelsMetadata(arg0 *big.Int) (struct {
	Author           string
	Hfid             string
	RoyaltyPrice     *big.Int
	AuthorWallet     common.Address
	ModelType        uint8
	MaxContextWindow *big.Int
}, error) {
	return _LLMNFT.Contract.ModelsMetadata(&_LLMNFT.CallOpts, arg0)
}

// ModelsMetadata is a free data retrieval call binding the contract method 0x8a87e958.
//
// Solidity: function Models_metadata(uint256 ) view returns(string author, string hfid, uint256 royalty_price, address author_wallet, uint8 model_type, uint256 max_context_window)
func (_LLMNFT *LLMNFTCallerSession) ModelsMetadata(arg0 *big.Int) (struct {
	Author           string
	Hfid             string
	RoyaltyPrice     *big.Int
	AuthorWallet     common.Address
	ModelType        uint8
	MaxContextWindow *big.Int
}, error) {
	return _LLMNFT.Contract.ModelsMetadata(&_LLMNFT.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_LLMNFT *LLMNFTCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LLMNFT.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_LLMNFT *LLMNFTSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _LLMNFT.Contract.BalanceOf(&_LLMNFT.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_LLMNFT *LLMNFTCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _LLMNFT.Contract.BalanceOf(&_LLMNFT.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_LLMNFT *LLMNFTCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _LLMNFT.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_LLMNFT *LLMNFTSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _LLMNFT.Contract.GetApproved(&_LLMNFT.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_LLMNFT *LLMNFTCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _LLMNFT.Contract.GetApproved(&_LLMNFT.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_LLMNFT *LLMNFTCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _LLMNFT.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_LLMNFT *LLMNFTSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _LLMNFT.Contract.IsApprovedForAll(&_LLMNFT.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_LLMNFT *LLMNFTCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _LLMNFT.Contract.IsApprovedForAll(&_LLMNFT.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LLMNFT *LLMNFTCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LLMNFT.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LLMNFT *LLMNFTSession) Name() (string, error) {
	return _LLMNFT.Contract.Name(&_LLMNFT.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LLMNFT *LLMNFTCallerSession) Name() (string, error) {
	return _LLMNFT.Contract.Name(&_LLMNFT.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_LLMNFT *LLMNFTCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _LLMNFT.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_LLMNFT *LLMNFTSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _LLMNFT.Contract.OwnerOf(&_LLMNFT.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_LLMNFT *LLMNFTCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _LLMNFT.Contract.OwnerOf(&_LLMNFT.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LLMNFT *LLMNFTCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _LLMNFT.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LLMNFT *LLMNFTSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LLMNFT.Contract.SupportsInterface(&_LLMNFT.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LLMNFT *LLMNFTCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LLMNFT.Contract.SupportsInterface(&_LLMNFT.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LLMNFT *LLMNFTCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LLMNFT.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LLMNFT *LLMNFTSession) Symbol() (string, error) {
	return _LLMNFT.Contract.Symbol(&_LLMNFT.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LLMNFT *LLMNFTCallerSession) Symbol() (string, error) {
	return _LLMNFT.Contract.Symbol(&_LLMNFT.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_LLMNFT *LLMNFTCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _LLMNFT.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_LLMNFT *LLMNFTSession) TokenURI(tokenId *big.Int) (string, error) {
	return _LLMNFT.Contract.TokenURI(&_LLMNFT.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_LLMNFT *LLMNFTCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _LLMNFT.Contract.TokenURI(&_LLMNFT.CallOpts, tokenId)
}

// CreateLLMNFT is a paid mutator transaction binding the contract method 0x31d99e7a.
//
// Solidity: function CreateLLM_NFT(string hf_id_, uint256 price, address wallet, uint8 model_type_, uint256 max_context) returns(uint256)
func (_LLMNFT *LLMNFTTransactor) CreateLLMNFT(opts *bind.TransactOpts, hf_id_ string, price *big.Int, wallet common.Address, model_type_ uint8, max_context *big.Int) (*types.Transaction, error) {
	return _LLMNFT.contract.Transact(opts, "CreateLLM_NFT", hf_id_, price, wallet, model_type_, max_context)
}

// CreateLLMNFT is a paid mutator transaction binding the contract method 0x31d99e7a.
//
// Solidity: function CreateLLM_NFT(string hf_id_, uint256 price, address wallet, uint8 model_type_, uint256 max_context) returns(uint256)
func (_LLMNFT *LLMNFTSession) CreateLLMNFT(hf_id_ string, price *big.Int, wallet common.Address, model_type_ uint8, max_context *big.Int) (*types.Transaction, error) {
	return _LLMNFT.Contract.CreateLLMNFT(&_LLMNFT.TransactOpts, hf_id_, price, wallet, model_type_, max_context)
}

// CreateLLMNFT is a paid mutator transaction binding the contract method 0x31d99e7a.
//
// Solidity: function CreateLLM_NFT(string hf_id_, uint256 price, address wallet, uint8 model_type_, uint256 max_context) returns(uint256)
func (_LLMNFT *LLMNFTTransactorSession) CreateLLMNFT(hf_id_ string, price *big.Int, wallet common.Address, model_type_ uint8, max_context *big.Int) (*types.Transaction, error) {
	return _LLMNFT.Contract.CreateLLMNFT(&_LLMNFT.TransactOpts, hf_id_, price, wallet, model_type_, max_context)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_LLMNFT *LLMNFTTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LLMNFT.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_LLMNFT *LLMNFTSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LLMNFT.Contract.Approve(&_LLMNFT.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_LLMNFT *LLMNFTTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LLMNFT.Contract.Approve(&_LLMNFT.TransactOpts, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_LLMNFT *LLMNFTTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LLMNFT.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_LLMNFT *LLMNFTSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LLMNFT.Contract.SafeTransferFrom(&_LLMNFT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_LLMNFT *LLMNFTTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LLMNFT.Contract.SafeTransferFrom(&_LLMNFT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_LLMNFT *LLMNFTTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _LLMNFT.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_LLMNFT *LLMNFTSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _LLMNFT.Contract.SafeTransferFrom0(&_LLMNFT.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_LLMNFT *LLMNFTTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _LLMNFT.Contract.SafeTransferFrom0(&_LLMNFT.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_LLMNFT *LLMNFTTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _LLMNFT.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_LLMNFT *LLMNFTSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _LLMNFT.Contract.SetApprovalForAll(&_LLMNFT.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_LLMNFT *LLMNFTTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _LLMNFT.Contract.SetApprovalForAll(&_LLMNFT.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_LLMNFT *LLMNFTTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LLMNFT.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_LLMNFT *LLMNFTSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LLMNFT.Contract.TransferFrom(&_LLMNFT.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_LLMNFT *LLMNFTTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LLMNFT.Contract.TransferFrom(&_LLMNFT.TransactOpts, from, to, tokenId)
}

// LLMNFTApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the LLMNFT contract.
type LLMNFTApprovalIterator struct {
	Event *LLMNFTApproval // Event containing the contract specifics and raw log

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
func (it *LLMNFTApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LLMNFTApproval)
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
		it.Event = new(LLMNFTApproval)
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
func (it *LLMNFTApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LLMNFTApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LLMNFTApproval represents a Approval event raised by the LLMNFT contract.
type LLMNFTApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_LLMNFT *LLMNFTFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*LLMNFTApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _LLMNFT.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &LLMNFTApprovalIterator{contract: _LLMNFT.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_LLMNFT *LLMNFTFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *LLMNFTApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _LLMNFT.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LLMNFTApproval)
				if err := _LLMNFT.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_LLMNFT *LLMNFTFilterer) ParseApproval(log types.Log) (*LLMNFTApproval, error) {
	event := new(LLMNFTApproval)
	if err := _LLMNFT.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LLMNFTApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the LLMNFT contract.
type LLMNFTApprovalForAllIterator struct {
	Event *LLMNFTApprovalForAll // Event containing the contract specifics and raw log

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
func (it *LLMNFTApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LLMNFTApprovalForAll)
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
		it.Event = new(LLMNFTApprovalForAll)
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
func (it *LLMNFTApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LLMNFTApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LLMNFTApprovalForAll represents a ApprovalForAll event raised by the LLMNFT contract.
type LLMNFTApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_LLMNFT *LLMNFTFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*LLMNFTApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _LLMNFT.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &LLMNFTApprovalForAllIterator{contract: _LLMNFT.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_LLMNFT *LLMNFTFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *LLMNFTApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _LLMNFT.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LLMNFTApprovalForAll)
				if err := _LLMNFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_LLMNFT *LLMNFTFilterer) ParseApprovalForAll(log types.Log) (*LLMNFTApprovalForAll, error) {
	event := new(LLMNFTApprovalForAll)
	if err := _LLMNFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LLMNFTBatchMetadataUpdateIterator is returned from FilterBatchMetadataUpdate and is used to iterate over the raw logs and unpacked data for BatchMetadataUpdate events raised by the LLMNFT contract.
type LLMNFTBatchMetadataUpdateIterator struct {
	Event *LLMNFTBatchMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *LLMNFTBatchMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LLMNFTBatchMetadataUpdate)
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
		it.Event = new(LLMNFTBatchMetadataUpdate)
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
func (it *LLMNFTBatchMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LLMNFTBatchMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LLMNFTBatchMetadataUpdate represents a BatchMetadataUpdate event raised by the LLMNFT contract.
type LLMNFTBatchMetadataUpdate struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBatchMetadataUpdate is a free log retrieval operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_LLMNFT *LLMNFTFilterer) FilterBatchMetadataUpdate(opts *bind.FilterOpts) (*LLMNFTBatchMetadataUpdateIterator, error) {

	logs, sub, err := _LLMNFT.contract.FilterLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &LLMNFTBatchMetadataUpdateIterator{contract: _LLMNFT.contract, event: "BatchMetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchBatchMetadataUpdate is a free log subscription operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_LLMNFT *LLMNFTFilterer) WatchBatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *LLMNFTBatchMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _LLMNFT.contract.WatchLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LLMNFTBatchMetadataUpdate)
				if err := _LLMNFT.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
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

// ParseBatchMetadataUpdate is a log parse operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_LLMNFT *LLMNFTFilterer) ParseBatchMetadataUpdate(log types.Log) (*LLMNFTBatchMetadataUpdate, error) {
	event := new(LLMNFTBatchMetadataUpdate)
	if err := _LLMNFT.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LLMNFTLLMNFTCreatedIterator is returned from FilterLLMNFTCreated and is used to iterate over the raw logs and unpacked data for LLMNFTCreated events raised by the LLMNFT contract.
type LLMNFTLLMNFTCreatedIterator struct {
	Event *LLMNFTLLMNFTCreated // Event containing the contract specifics and raw log

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
func (it *LLMNFTLLMNFTCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LLMNFTLLMNFTCreated)
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
		it.Event = new(LLMNFTLLMNFTCreated)
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
func (it *LLMNFTLLMNFTCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LLMNFTLLMNFTCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LLMNFTLLMNFTCreated represents a LLMNFTCreated event raised by the LLMNFT contract.
type LLMNFTLLMNFTCreated struct {
	TokenId          *big.Int
	Author           common.Address
	Hfid             string
	RoyaltyPrice     *big.Int
	AuthorWallet     common.Address
	ModelType        uint8
	MaxContextWindow *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLLMNFTCreated is a free log retrieval operation binding the contract event 0xdb52b24549ce275e6027cbfb3bad43901bbe705f5c7a0a539ea5b0b8a0be39ad.
//
// Solidity: event LLMNFTCreated(uint256 indexed tokenId, address author, string hfid, uint256 royaltyPrice, address authorWallet, uint8 modelType, uint256 maxContextWindow)
func (_LLMNFT *LLMNFTFilterer) FilterLLMNFTCreated(opts *bind.FilterOpts, tokenId []*big.Int) (*LLMNFTLLMNFTCreatedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _LLMNFT.contract.FilterLogs(opts, "LLMNFTCreated", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &LLMNFTLLMNFTCreatedIterator{contract: _LLMNFT.contract, event: "LLMNFTCreated", logs: logs, sub: sub}, nil
}

// WatchLLMNFTCreated is a free log subscription operation binding the contract event 0xdb52b24549ce275e6027cbfb3bad43901bbe705f5c7a0a539ea5b0b8a0be39ad.
//
// Solidity: event LLMNFTCreated(uint256 indexed tokenId, address author, string hfid, uint256 royaltyPrice, address authorWallet, uint8 modelType, uint256 maxContextWindow)
func (_LLMNFT *LLMNFTFilterer) WatchLLMNFTCreated(opts *bind.WatchOpts, sink chan<- *LLMNFTLLMNFTCreated, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _LLMNFT.contract.WatchLogs(opts, "LLMNFTCreated", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LLMNFTLLMNFTCreated)
				if err := _LLMNFT.contract.UnpackLog(event, "LLMNFTCreated", log); err != nil {
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

// ParseLLMNFTCreated is a log parse operation binding the contract event 0xdb52b24549ce275e6027cbfb3bad43901bbe705f5c7a0a539ea5b0b8a0be39ad.
//
// Solidity: event LLMNFTCreated(uint256 indexed tokenId, address author, string hfid, uint256 royaltyPrice, address authorWallet, uint8 modelType, uint256 maxContextWindow)
func (_LLMNFT *LLMNFTFilterer) ParseLLMNFTCreated(log types.Log) (*LLMNFTLLMNFTCreated, error) {
	event := new(LLMNFTLLMNFTCreated)
	if err := _LLMNFT.contract.UnpackLog(event, "LLMNFTCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LLMNFTMetadataUpdateIterator is returned from FilterMetadataUpdate and is used to iterate over the raw logs and unpacked data for MetadataUpdate events raised by the LLMNFT contract.
type LLMNFTMetadataUpdateIterator struct {
	Event *LLMNFTMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *LLMNFTMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LLMNFTMetadataUpdate)
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
		it.Event = new(LLMNFTMetadataUpdate)
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
func (it *LLMNFTMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LLMNFTMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LLMNFTMetadataUpdate represents a MetadataUpdate event raised by the LLMNFT contract.
type LLMNFTMetadataUpdate struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMetadataUpdate is a free log retrieval operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_LLMNFT *LLMNFTFilterer) FilterMetadataUpdate(opts *bind.FilterOpts) (*LLMNFTMetadataUpdateIterator, error) {

	logs, sub, err := _LLMNFT.contract.FilterLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &LLMNFTMetadataUpdateIterator{contract: _LLMNFT.contract, event: "MetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchMetadataUpdate is a free log subscription operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_LLMNFT *LLMNFTFilterer) WatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *LLMNFTMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _LLMNFT.contract.WatchLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LLMNFTMetadataUpdate)
				if err := _LLMNFT.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
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

// ParseMetadataUpdate is a log parse operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_LLMNFT *LLMNFTFilterer) ParseMetadataUpdate(log types.Log) (*LLMNFTMetadataUpdate, error) {
	event := new(LLMNFTMetadataUpdate)
	if err := _LLMNFT.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LLMNFTTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the LLMNFT contract.
type LLMNFTTransferIterator struct {
	Event *LLMNFTTransfer // Event containing the contract specifics and raw log

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
func (it *LLMNFTTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LLMNFTTransfer)
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
		it.Event = new(LLMNFTTransfer)
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
func (it *LLMNFTTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LLMNFTTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LLMNFTTransfer represents a Transfer event raised by the LLMNFT contract.
type LLMNFTTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_LLMNFT *LLMNFTFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*LLMNFTTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _LLMNFT.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &LLMNFTTransferIterator{contract: _LLMNFT.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_LLMNFT *LLMNFTFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *LLMNFTTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _LLMNFT.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LLMNFTTransfer)
				if err := _LLMNFT.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_LLMNFT *LLMNFTFilterer) ParseTransfer(log types.Log) (*LLMNFTTransfer, error) {
	event := new(LLMNFTTransfer)
	if err := _LLMNFT.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
