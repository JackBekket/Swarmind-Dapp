// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package llmnft

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

// LLMNFTLLM is an auto generated low-level Go binding around an user-defined struct.
type LLMNFTLLM struct {
	Author           string
	Hfid             string
	RoyaltyPrice     *big.Int
	AuthorWallet     common.Address
	ModelType        uint8
	MaxContextWindow *big.Int
}

// LlmnftMetaData contains all meta data concerning the Llmnft contract.
var LlmnftMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"smbl_\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_toTokenId\",\"type\":\"uint256\"}],\"name\":\"BatchMetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"MetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"hf_id_\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"internalType\":\"enumLLMNFT.Llm_type\",\"name\":\"model_type_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"max_context\",\"type\":\"uint256\"}],\"name\":\"CreateLLM_NFT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"token_id\",\"type\":\"uint256\"}],\"name\":\"GetLLM\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"author\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"hfid\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"royalty_price\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"author_wallet\",\"type\":\"address\"},{\"internalType\":\"enumLLMNFT.Llm_type\",\"name\":\"model_type\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"max_context_window\",\"type\":\"uint256\"}],\"internalType\":\"structLLMNFT.LLM\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"Hf_models\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Models_metadata\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"author\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"hfid\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"royalty_price\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"author_wallet\",\"type\":\"address\"},{\"internalType\":\"enumLLMNFT.Llm_type\",\"name\":\"model_type\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"max_context_window\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// LlmnftABI is the input ABI used to generate the binding from.
// Deprecated: Use LlmnftMetaData.ABI instead.
var LlmnftABI = LlmnftMetaData.ABI

// Llmnft is an auto generated Go binding around an Ethereum contract.
type Llmnft struct {
	LlmnftCaller     // Read-only binding to the contract
	LlmnftTransactor // Write-only binding to the contract
	LlmnftFilterer   // Log filterer for contract events
}

// LlmnftCaller is an auto generated read-only Go binding around an Ethereum contract.
type LlmnftCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LlmnftTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LlmnftTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LlmnftFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LlmnftFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LlmnftSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LlmnftSession struct {
	Contract     *Llmnft           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LlmnftCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LlmnftCallerSession struct {
	Contract *LlmnftCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// LlmnftTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LlmnftTransactorSession struct {
	Contract     *LlmnftTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LlmnftRaw is an auto generated low-level Go binding around an Ethereum contract.
type LlmnftRaw struct {
	Contract *Llmnft // Generic contract binding to access the raw methods on
}

// LlmnftCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LlmnftCallerRaw struct {
	Contract *LlmnftCaller // Generic read-only contract binding to access the raw methods on
}

// LlmnftTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LlmnftTransactorRaw struct {
	Contract *LlmnftTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLlmnft creates a new instance of Llmnft, bound to a specific deployed contract.
func NewLlmnft(address common.Address, backend bind.ContractBackend) (*Llmnft, error) {
	contract, err := bindLlmnft(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Llmnft{LlmnftCaller: LlmnftCaller{contract: contract}, LlmnftTransactor: LlmnftTransactor{contract: contract}, LlmnftFilterer: LlmnftFilterer{contract: contract}}, nil
}

// NewLlmnftCaller creates a new read-only instance of Llmnft, bound to a specific deployed contract.
func NewLlmnftCaller(address common.Address, caller bind.ContractCaller) (*LlmnftCaller, error) {
	contract, err := bindLlmnft(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LlmnftCaller{contract: contract}, nil
}

// NewLlmnftTransactor creates a new write-only instance of Llmnft, bound to a specific deployed contract.
func NewLlmnftTransactor(address common.Address, transactor bind.ContractTransactor) (*LlmnftTransactor, error) {
	contract, err := bindLlmnft(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LlmnftTransactor{contract: contract}, nil
}

// NewLlmnftFilterer creates a new log filterer instance of Llmnft, bound to a specific deployed contract.
func NewLlmnftFilterer(address common.Address, filterer bind.ContractFilterer) (*LlmnftFilterer, error) {
	contract, err := bindLlmnft(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LlmnftFilterer{contract: contract}, nil
}

// bindLlmnft binds a generic wrapper to an already deployed contract.
func bindLlmnft(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LlmnftABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Llmnft *LlmnftRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Llmnft.Contract.LlmnftCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Llmnft *LlmnftRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Llmnft.Contract.LlmnftTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Llmnft *LlmnftRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Llmnft.Contract.LlmnftTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Llmnft *LlmnftCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Llmnft.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Llmnft *LlmnftTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Llmnft.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Llmnft *LlmnftTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Llmnft.Contract.contract.Transact(opts, method, params...)
}

// GetLLM is a free data retrieval call binding the contract method 0x40a24c67.
//
// Solidity: function GetLLM(uint256 token_id) view returns((string,string,uint256,address,uint8,uint256))
func (_Llmnft *LlmnftCaller) GetLLM(opts *bind.CallOpts, token_id *big.Int) (LLMNFTLLM, error) {
	var out []interface{}
	err := _Llmnft.contract.Call(opts, &out, "GetLLM", token_id)

	if err != nil {
		return *new(LLMNFTLLM), err
	}

	out0 := *abi.ConvertType(out[0], new(LLMNFTLLM)).(*LLMNFTLLM)

	return out0, err

}

// GetLLM is a free data retrieval call binding the contract method 0x40a24c67.
//
// Solidity: function GetLLM(uint256 token_id) view returns((string,string,uint256,address,uint8,uint256))
func (_Llmnft *LlmnftSession) GetLLM(token_id *big.Int) (LLMNFTLLM, error) {
	return _Llmnft.Contract.GetLLM(&_Llmnft.CallOpts, token_id)
}

// GetLLM is a free data retrieval call binding the contract method 0x40a24c67.
//
// Solidity: function GetLLM(uint256 token_id) view returns((string,string,uint256,address,uint8,uint256))
func (_Llmnft *LlmnftCallerSession) GetLLM(token_id *big.Int) (LLMNFTLLM, error) {
	return _Llmnft.Contract.GetLLM(&_Llmnft.CallOpts, token_id)
}

// HfModels is a free data retrieval call binding the contract method 0x65783bbb.
//
// Solidity: function Hf_models(string ) view returns(uint256)
func (_Llmnft *LlmnftCaller) HfModels(opts *bind.CallOpts, arg0 string) (*big.Int, error) {
	var out []interface{}
	err := _Llmnft.contract.Call(opts, &out, "Hf_models", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HfModels is a free data retrieval call binding the contract method 0x65783bbb.
//
// Solidity: function Hf_models(string ) view returns(uint256)
func (_Llmnft *LlmnftSession) HfModels(arg0 string) (*big.Int, error) {
	return _Llmnft.Contract.HfModels(&_Llmnft.CallOpts, arg0)
}

// HfModels is a free data retrieval call binding the contract method 0x65783bbb.
//
// Solidity: function Hf_models(string ) view returns(uint256)
func (_Llmnft *LlmnftCallerSession) HfModels(arg0 string) (*big.Int, error) {
	return _Llmnft.Contract.HfModels(&_Llmnft.CallOpts, arg0)
}

// ModelsMetadata is a free data retrieval call binding the contract method 0x8a87e958.
//
// Solidity: function Models_metadata(uint256 ) view returns(string author, string hfid, uint256 royalty_price, address author_wallet, uint8 model_type, uint256 max_context_window)
func (_Llmnft *LlmnftCaller) ModelsMetadata(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Author           string
	Hfid             string
	RoyaltyPrice     *big.Int
	AuthorWallet     common.Address
	ModelType        uint8
	MaxContextWindow *big.Int
}, error) {
	var out []interface{}
	err := _Llmnft.contract.Call(opts, &out, "Models_metadata", arg0)

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
func (_Llmnft *LlmnftSession) ModelsMetadata(arg0 *big.Int) (struct {
	Author           string
	Hfid             string
	RoyaltyPrice     *big.Int
	AuthorWallet     common.Address
	ModelType        uint8
	MaxContextWindow *big.Int
}, error) {
	return _Llmnft.Contract.ModelsMetadata(&_Llmnft.CallOpts, arg0)
}

// ModelsMetadata is a free data retrieval call binding the contract method 0x8a87e958.
//
// Solidity: function Models_metadata(uint256 ) view returns(string author, string hfid, uint256 royalty_price, address author_wallet, uint8 model_type, uint256 max_context_window)
func (_Llmnft *LlmnftCallerSession) ModelsMetadata(arg0 *big.Int) (struct {
	Author           string
	Hfid             string
	RoyaltyPrice     *big.Int
	AuthorWallet     common.Address
	ModelType        uint8
	MaxContextWindow *big.Int
}, error) {
	return _Llmnft.Contract.ModelsMetadata(&_Llmnft.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Llmnft *LlmnftCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Llmnft.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Llmnft *LlmnftSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Llmnft.Contract.BalanceOf(&_Llmnft.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Llmnft *LlmnftCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Llmnft.Contract.BalanceOf(&_Llmnft.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Llmnft *LlmnftCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Llmnft.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Llmnft *LlmnftSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Llmnft.Contract.GetApproved(&_Llmnft.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Llmnft *LlmnftCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Llmnft.Contract.GetApproved(&_Llmnft.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Llmnft *LlmnftCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Llmnft.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Llmnft *LlmnftSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Llmnft.Contract.IsApprovedForAll(&_Llmnft.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Llmnft *LlmnftCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Llmnft.Contract.IsApprovedForAll(&_Llmnft.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Llmnft *LlmnftCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Llmnft.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Llmnft *LlmnftSession) Name() (string, error) {
	return _Llmnft.Contract.Name(&_Llmnft.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Llmnft *LlmnftCallerSession) Name() (string, error) {
	return _Llmnft.Contract.Name(&_Llmnft.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Llmnft *LlmnftCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Llmnft.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Llmnft *LlmnftSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Llmnft.Contract.OwnerOf(&_Llmnft.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Llmnft *LlmnftCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Llmnft.Contract.OwnerOf(&_Llmnft.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Llmnft *LlmnftCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Llmnft.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Llmnft *LlmnftSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Llmnft.Contract.SupportsInterface(&_Llmnft.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Llmnft *LlmnftCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Llmnft.Contract.SupportsInterface(&_Llmnft.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Llmnft *LlmnftCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Llmnft.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Llmnft *LlmnftSession) Symbol() (string, error) {
	return _Llmnft.Contract.Symbol(&_Llmnft.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Llmnft *LlmnftCallerSession) Symbol() (string, error) {
	return _Llmnft.Contract.Symbol(&_Llmnft.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Llmnft *LlmnftCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Llmnft.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Llmnft *LlmnftSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Llmnft.Contract.TokenURI(&_Llmnft.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Llmnft *LlmnftCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Llmnft.Contract.TokenURI(&_Llmnft.CallOpts, tokenId)
}

// CreateLLMNFT is a paid mutator transaction binding the contract method 0x31d99e7a.
//
// Solidity: function CreateLLM_NFT(string hf_id_, uint256 price, address wallet, uint8 model_type_, uint256 max_context) returns(uint256)
func (_Llmnft *LlmnftTransactor) CreateLLMNFT(opts *bind.TransactOpts, hf_id_ string, price *big.Int, wallet common.Address, model_type_ uint8, max_context *big.Int) (*types.Transaction, error) {
	return _Llmnft.contract.Transact(opts, "CreateLLM_NFT", hf_id_, price, wallet, model_type_, max_context)
}

// CreateLLMNFT is a paid mutator transaction binding the contract method 0x31d99e7a.
//
// Solidity: function CreateLLM_NFT(string hf_id_, uint256 price, address wallet, uint8 model_type_, uint256 max_context) returns(uint256)
func (_Llmnft *LlmnftSession) CreateLLMNFT(hf_id_ string, price *big.Int, wallet common.Address, model_type_ uint8, max_context *big.Int) (*types.Transaction, error) {
	return _Llmnft.Contract.CreateLLMNFT(&_Llmnft.TransactOpts, hf_id_, price, wallet, model_type_, max_context)
}

// CreateLLMNFT is a paid mutator transaction binding the contract method 0x31d99e7a.
//
// Solidity: function CreateLLM_NFT(string hf_id_, uint256 price, address wallet, uint8 model_type_, uint256 max_context) returns(uint256)
func (_Llmnft *LlmnftTransactorSession) CreateLLMNFT(hf_id_ string, price *big.Int, wallet common.Address, model_type_ uint8, max_context *big.Int) (*types.Transaction, error) {
	return _Llmnft.Contract.CreateLLMNFT(&_Llmnft.TransactOpts, hf_id_, price, wallet, model_type_, max_context)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Llmnft *LlmnftTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Llmnft.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Llmnft *LlmnftSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Llmnft.Contract.Approve(&_Llmnft.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Llmnft *LlmnftTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Llmnft.Contract.Approve(&_Llmnft.TransactOpts, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Llmnft *LlmnftTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Llmnft.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Llmnft *LlmnftSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Llmnft.Contract.SafeTransferFrom(&_Llmnft.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Llmnft *LlmnftTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Llmnft.Contract.SafeTransferFrom(&_Llmnft.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Llmnft *LlmnftTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Llmnft.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Llmnft *LlmnftSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Llmnft.Contract.SafeTransferFrom0(&_Llmnft.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Llmnft *LlmnftTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Llmnft.Contract.SafeTransferFrom0(&_Llmnft.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Llmnft *LlmnftTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Llmnft.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Llmnft *LlmnftSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Llmnft.Contract.SetApprovalForAll(&_Llmnft.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Llmnft *LlmnftTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Llmnft.Contract.SetApprovalForAll(&_Llmnft.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Llmnft *LlmnftTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Llmnft.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Llmnft *LlmnftSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Llmnft.Contract.TransferFrom(&_Llmnft.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Llmnft *LlmnftTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Llmnft.Contract.TransferFrom(&_Llmnft.TransactOpts, from, to, tokenId)
}

// LlmnftApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Llmnft contract.
type LlmnftApprovalIterator struct {
	Event *LlmnftApproval // Event containing the contract specifics and raw log

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
func (it *LlmnftApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LlmnftApproval)
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
		it.Event = new(LlmnftApproval)
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
func (it *LlmnftApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LlmnftApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LlmnftApproval represents a Approval event raised by the Llmnft contract.
type LlmnftApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Llmnft *LlmnftFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*LlmnftApprovalIterator, error) {

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

	logs, sub, err := _Llmnft.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &LlmnftApprovalIterator{contract: _Llmnft.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Llmnft *LlmnftFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *LlmnftApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Llmnft.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LlmnftApproval)
				if err := _Llmnft.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Llmnft *LlmnftFilterer) ParseApproval(log types.Log) (*LlmnftApproval, error) {
	event := new(LlmnftApproval)
	if err := _Llmnft.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LlmnftApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Llmnft contract.
type LlmnftApprovalForAllIterator struct {
	Event *LlmnftApprovalForAll // Event containing the contract specifics and raw log

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
func (it *LlmnftApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LlmnftApprovalForAll)
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
		it.Event = new(LlmnftApprovalForAll)
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
func (it *LlmnftApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LlmnftApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LlmnftApprovalForAll represents a ApprovalForAll event raised by the Llmnft contract.
type LlmnftApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Llmnft *LlmnftFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*LlmnftApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Llmnft.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &LlmnftApprovalForAllIterator{contract: _Llmnft.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Llmnft *LlmnftFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *LlmnftApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Llmnft.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LlmnftApprovalForAll)
				if err := _Llmnft.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_Llmnft *LlmnftFilterer) ParseApprovalForAll(log types.Log) (*LlmnftApprovalForAll, error) {
	event := new(LlmnftApprovalForAll)
	if err := _Llmnft.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LlmnftBatchMetadataUpdateIterator is returned from FilterBatchMetadataUpdate and is used to iterate over the raw logs and unpacked data for BatchMetadataUpdate events raised by the Llmnft contract.
type LlmnftBatchMetadataUpdateIterator struct {
	Event *LlmnftBatchMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *LlmnftBatchMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LlmnftBatchMetadataUpdate)
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
		it.Event = new(LlmnftBatchMetadataUpdate)
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
func (it *LlmnftBatchMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LlmnftBatchMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LlmnftBatchMetadataUpdate represents a BatchMetadataUpdate event raised by the Llmnft contract.
type LlmnftBatchMetadataUpdate struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBatchMetadataUpdate is a free log retrieval operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_Llmnft *LlmnftFilterer) FilterBatchMetadataUpdate(opts *bind.FilterOpts) (*LlmnftBatchMetadataUpdateIterator, error) {

	logs, sub, err := _Llmnft.contract.FilterLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &LlmnftBatchMetadataUpdateIterator{contract: _Llmnft.contract, event: "BatchMetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchBatchMetadataUpdate is a free log subscription operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_Llmnft *LlmnftFilterer) WatchBatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *LlmnftBatchMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _Llmnft.contract.WatchLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LlmnftBatchMetadataUpdate)
				if err := _Llmnft.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
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
func (_Llmnft *LlmnftFilterer) ParseBatchMetadataUpdate(log types.Log) (*LlmnftBatchMetadataUpdate, error) {
	event := new(LlmnftBatchMetadataUpdate)
	if err := _Llmnft.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LlmnftMetadataUpdateIterator is returned from FilterMetadataUpdate and is used to iterate over the raw logs and unpacked data for MetadataUpdate events raised by the Llmnft contract.
type LlmnftMetadataUpdateIterator struct {
	Event *LlmnftMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *LlmnftMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LlmnftMetadataUpdate)
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
		it.Event = new(LlmnftMetadataUpdate)
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
func (it *LlmnftMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LlmnftMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LlmnftMetadataUpdate represents a MetadataUpdate event raised by the Llmnft contract.
type LlmnftMetadataUpdate struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMetadataUpdate is a free log retrieval operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_Llmnft *LlmnftFilterer) FilterMetadataUpdate(opts *bind.FilterOpts) (*LlmnftMetadataUpdateIterator, error) {

	logs, sub, err := _Llmnft.contract.FilterLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &LlmnftMetadataUpdateIterator{contract: _Llmnft.contract, event: "MetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchMetadataUpdate is a free log subscription operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_Llmnft *LlmnftFilterer) WatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *LlmnftMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _Llmnft.contract.WatchLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LlmnftMetadataUpdate)
				if err := _Llmnft.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
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
func (_Llmnft *LlmnftFilterer) ParseMetadataUpdate(log types.Log) (*LlmnftMetadataUpdate, error) {
	event := new(LlmnftMetadataUpdate)
	if err := _Llmnft.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LlmnftTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Llmnft contract.
type LlmnftTransferIterator struct {
	Event *LlmnftTransfer // Event containing the contract specifics and raw log

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
func (it *LlmnftTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LlmnftTransfer)
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
		it.Event = new(LlmnftTransfer)
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
func (it *LlmnftTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LlmnftTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LlmnftTransfer represents a Transfer event raised by the Llmnft contract.
type LlmnftTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Llmnft *LlmnftFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*LlmnftTransferIterator, error) {

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

	logs, sub, err := _Llmnft.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &LlmnftTransferIterator{contract: _Llmnft.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Llmnft *LlmnftFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *LlmnftTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Llmnft.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LlmnftTransfer)
				if err := _Llmnft.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Llmnft *LlmnftFilterer) ParseTransfer(log types.Log) (*LlmnftTransfer, error) {
	event := new(LlmnftTransfer)
	if err := _Llmnft.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
