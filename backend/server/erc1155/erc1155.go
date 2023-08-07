// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc1155

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

// Token1155MetaData contains all meta data concerning the Token1155 contract.
var Token1155MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_contractURI\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"_transferSend\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_mintAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_number\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"endUser\",\"type\":\"address\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"endUser\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amount\",\"type\":\"uint256[]\"}],\"name\":\"claimBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBaseURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_state\",\"type\":\"bool\"}],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_newBaseURI\",\"type\":\"string\"}],\"name\":\"setBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"contractURI_\",\"type\":\"string\"}],\"name\":\"setContractURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupplyMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelistClaimed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"whitelistedAddresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// Token1155ABI is the input ABI used to generate the binding from.
// Deprecated: Use Token1155MetaData.ABI instead.
var Token1155ABI = Token1155MetaData.ABI

// Token1155 is an auto generated Go binding around an Ethereum contract.
type Token1155 struct {
	Token1155Caller     // Read-only binding to the contract
	Token1155Transactor // Write-only binding to the contract
	Token1155Filterer   // Log filterer for contract events
}

// Token1155Caller is an auto generated read-only Go binding around an Ethereum contract.
type Token1155Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Token1155Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Token1155Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Token1155Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Token1155Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Token1155Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Token1155Session struct {
	Contract     *Token1155        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Token1155CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Token1155CallerSession struct {
	Contract *Token1155Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// Token1155TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Token1155TransactorSession struct {
	Contract     *Token1155Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// Token1155Raw is an auto generated low-level Go binding around an Ethereum contract.
type Token1155Raw struct {
	Contract *Token1155 // Generic contract binding to access the raw methods on
}

// Token1155CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Token1155CallerRaw struct {
	Contract *Token1155Caller // Generic read-only contract binding to access the raw methods on
}

// Token1155TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Token1155TransactorRaw struct {
	Contract *Token1155Transactor // Generic write-only contract binding to access the raw methods on
}

// NewToken1155 creates a new instance of Token1155, bound to a specific deployed contract.
func NewToken1155(address common.Address, backend bind.ContractBackend) (*Token1155, error) {
	contract, err := bindToken1155(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Token1155{Token1155Caller: Token1155Caller{contract: contract}, Token1155Transactor: Token1155Transactor{contract: contract}, Token1155Filterer: Token1155Filterer{contract: contract}}, nil
}

// NewToken1155Caller creates a new read-only instance of Token1155, bound to a specific deployed contract.
func NewToken1155Caller(address common.Address, caller bind.ContractCaller) (*Token1155Caller, error) {
	contract, err := bindToken1155(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Token1155Caller{contract: contract}, nil
}

// NewToken1155Transactor creates a new write-only instance of Token1155, bound to a specific deployed contract.
func NewToken1155Transactor(address common.Address, transactor bind.ContractTransactor) (*Token1155Transactor, error) {
	contract, err := bindToken1155(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Token1155Transactor{contract: contract}, nil
}

// NewToken1155Filterer creates a new log filterer instance of Token1155, bound to a specific deployed contract.
func NewToken1155Filterer(address common.Address, filterer bind.ContractFilterer) (*Token1155Filterer, error) {
	contract, err := bindToken1155(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Token1155Filterer{contract: contract}, nil
}

// bindToken1155 binds a generic wrapper to an already deployed contract.
func bindToken1155(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Token1155ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token1155 *Token1155Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Token1155.Contract.Token1155Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token1155 *Token1155Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token1155.Contract.Token1155Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token1155 *Token1155Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token1155.Contract.Token1155Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token1155 *Token1155CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Token1155.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token1155 *Token1155TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token1155.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token1155 *Token1155TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token1155.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_Token1155 *Token1155Caller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Token1155.contract.Call(opts, &out, "balanceOf", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_Token1155 *Token1155Session) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _Token1155.Contract.BalanceOf(&_Token1155.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_Token1155 *Token1155CallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _Token1155.Contract.BalanceOf(&_Token1155.CallOpts, account, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_Token1155 *Token1155Caller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Token1155.contract.Call(opts, &out, "balanceOfBatch", accounts, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_Token1155 *Token1155Session) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _Token1155.Contract.BalanceOfBatch(&_Token1155.CallOpts, accounts, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_Token1155 *Token1155CallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _Token1155.Contract.BalanceOfBatch(&_Token1155.CallOpts, accounts, ids)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Token1155 *Token1155Caller) ContractURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Token1155.contract.Call(opts, &out, "contractURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Token1155 *Token1155Session) ContractURI() (string, error) {
	return _Token1155.Contract.ContractURI(&_Token1155.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Token1155 *Token1155CallerSession) ContractURI() (string, error) {
	return _Token1155.Contract.ContractURI(&_Token1155.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Token1155 *Token1155Caller) GetBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token1155.contract.Call(opts, &out, "getBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Token1155 *Token1155Session) GetBalance() (*big.Int, error) {
	return _Token1155.Contract.GetBalance(&_Token1155.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Token1155 *Token1155CallerSession) GetBalance() (*big.Int, error) {
	return _Token1155.Contract.GetBalance(&_Token1155.CallOpts)
}

// GetBaseURI is a free data retrieval call binding the contract method 0x714c5398.
//
// Solidity: function getBaseURI() view returns(string)
func (_Token1155 *Token1155Caller) GetBaseURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Token1155.contract.Call(opts, &out, "getBaseURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetBaseURI is a free data retrieval call binding the contract method 0x714c5398.
//
// Solidity: function getBaseURI() view returns(string)
func (_Token1155 *Token1155Session) GetBaseURI() (string, error) {
	return _Token1155.Contract.GetBaseURI(&_Token1155.CallOpts)
}

// GetBaseURI is a free data retrieval call binding the contract method 0x714c5398.
//
// Solidity: function getBaseURI() view returns(string)
func (_Token1155 *Token1155CallerSession) GetBaseURI() (string, error) {
	return _Token1155.Contract.GetBaseURI(&_Token1155.CallOpts)
}

// GetMaxSupply is a free data retrieval call binding the contract method 0x4c0f38c2.
//
// Solidity: function getMaxSupply() view returns(uint256)
func (_Token1155 *Token1155Caller) GetMaxSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token1155.contract.Call(opts, &out, "getMaxSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxSupply is a free data retrieval call binding the contract method 0x4c0f38c2.
//
// Solidity: function getMaxSupply() view returns(uint256)
func (_Token1155 *Token1155Session) GetMaxSupply() (*big.Int, error) {
	return _Token1155.Contract.GetMaxSupply(&_Token1155.CallOpts)
}

// GetMaxSupply is a free data retrieval call binding the contract method 0x4c0f38c2.
//
// Solidity: function getMaxSupply() view returns(uint256)
func (_Token1155 *Token1155CallerSession) GetMaxSupply() (*big.Int, error) {
	return _Token1155.Contract.GetMaxSupply(&_Token1155.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_Token1155 *Token1155Caller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Token1155.contract.Call(opts, &out, "isApprovedForAll", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_Token1155 *Token1155Session) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _Token1155.Contract.IsApprovedForAll(&_Token1155.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_Token1155 *Token1155CallerSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _Token1155.Contract.IsApprovedForAll(&_Token1155.CallOpts, account, operator)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Token1155 *Token1155Caller) IsPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Token1155.contract.Call(opts, &out, "isPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Token1155 *Token1155Session) IsPaused() (bool, error) {
	return _Token1155.Contract.IsPaused(&_Token1155.CallOpts)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Token1155 *Token1155CallerSession) IsPaused() (bool, error) {
	return _Token1155.Contract.IsPaused(&_Token1155.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Token1155 *Token1155Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Token1155.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Token1155 *Token1155Session) Name() (string, error) {
	return _Token1155.Contract.Name(&_Token1155.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Token1155 *Token1155CallerSession) Name() (string, error) {
	return _Token1155.Contract.Name(&_Token1155.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Token1155 *Token1155Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Token1155.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Token1155 *Token1155Session) Owner() (common.Address, error) {
	return _Token1155.Contract.Owner(&_Token1155.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Token1155 *Token1155CallerSession) Owner() (common.Address, error) {
	return _Token1155.Contract.Owner(&_Token1155.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Token1155 *Token1155Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Token1155.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Token1155 *Token1155Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Token1155.Contract.SupportsInterface(&_Token1155.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Token1155 *Token1155CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Token1155.Contract.SupportsInterface(&_Token1155.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Token1155 *Token1155Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Token1155.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Token1155 *Token1155Session) Symbol() (string, error) {
	return _Token1155.Contract.Symbol(&_Token1155.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Token1155 *Token1155CallerSession) Symbol() (string, error) {
	return _Token1155.Contract.Symbol(&_Token1155.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Token1155 *Token1155Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token1155.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Token1155 *Token1155Session) TotalSupply() (*big.Int, error) {
	return _Token1155.Contract.TotalSupply(&_Token1155.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Token1155 *Token1155CallerSession) TotalSupply() (*big.Int, error) {
	return _Token1155.Contract.TotalSupply(&_Token1155.CallOpts)
}

// TotalSupplyMint is a free data retrieval call binding the contract method 0xf5b0778b.
//
// Solidity: function totalSupplyMint() view returns(uint256)
func (_Token1155 *Token1155Caller) TotalSupplyMint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token1155.contract.Call(opts, &out, "totalSupplyMint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupplyMint is a free data retrieval call binding the contract method 0xf5b0778b.
//
// Solidity: function totalSupplyMint() view returns(uint256)
func (_Token1155 *Token1155Session) TotalSupplyMint() (*big.Int, error) {
	return _Token1155.Contract.TotalSupplyMint(&_Token1155.CallOpts)
}

// TotalSupplyMint is a free data retrieval call binding the contract method 0xf5b0778b.
//
// Solidity: function totalSupplyMint() view returns(uint256)
func (_Token1155 *Token1155CallerSession) TotalSupplyMint() (*big.Int, error) {
	return _Token1155.Contract.TotalSupplyMint(&_Token1155.CallOpts)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_Token1155 *Token1155Caller) Uri(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Token1155.contract.Call(opts, &out, "uri", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_Token1155 *Token1155Session) Uri(arg0 *big.Int) (string, error) {
	return _Token1155.Contract.Uri(&_Token1155.CallOpts, arg0)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_Token1155 *Token1155CallerSession) Uri(arg0 *big.Int) (string, error) {
	return _Token1155.Contract.Uri(&_Token1155.CallOpts, arg0)
}

// WhitelistClaimed is a free data retrieval call binding the contract method 0xdb4bec44.
//
// Solidity: function whitelistClaimed(address ) view returns(bool)
func (_Token1155 *Token1155Caller) WhitelistClaimed(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Token1155.contract.Call(opts, &out, "whitelistClaimed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WhitelistClaimed is a free data retrieval call binding the contract method 0xdb4bec44.
//
// Solidity: function whitelistClaimed(address ) view returns(bool)
func (_Token1155 *Token1155Session) WhitelistClaimed(arg0 common.Address) (bool, error) {
	return _Token1155.Contract.WhitelistClaimed(&_Token1155.CallOpts, arg0)
}

// WhitelistClaimed is a free data retrieval call binding the contract method 0xdb4bec44.
//
// Solidity: function whitelistClaimed(address ) view returns(bool)
func (_Token1155 *Token1155CallerSession) WhitelistClaimed(arg0 common.Address) (bool, error) {
	return _Token1155.Contract.WhitelistClaimed(&_Token1155.CallOpts, arg0)
}

// WhitelistedAddresses is a free data retrieval call binding the contract method 0xba4e5c49.
//
// Solidity: function whitelistedAddresses(uint256 ) view returns(address)
func (_Token1155 *Token1155Caller) WhitelistedAddresses(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Token1155.contract.Call(opts, &out, "whitelistedAddresses", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WhitelistedAddresses is a free data retrieval call binding the contract method 0xba4e5c49.
//
// Solidity: function whitelistedAddresses(uint256 ) view returns(address)
func (_Token1155 *Token1155Session) WhitelistedAddresses(arg0 *big.Int) (common.Address, error) {
	return _Token1155.Contract.WhitelistedAddresses(&_Token1155.CallOpts, arg0)
}

// WhitelistedAddresses is a free data retrieval call binding the contract method 0xba4e5c49.
//
// Solidity: function whitelistedAddresses(uint256 ) view returns(address)
func (_Token1155 *Token1155CallerSession) WhitelistedAddresses(arg0 *big.Int) (common.Address, error) {
	return _Token1155.Contract.WhitelistedAddresses(&_Token1155.CallOpts, arg0)
}

// Claim is a paid mutator transaction binding the contract method 0xac44ff31.
//
// Solidity: function claim(uint256 _mintAmount, uint256 _number, address endUser) returns()
func (_Token1155 *Token1155Transactor) Claim(opts *bind.TransactOpts, _mintAmount *big.Int, _number *big.Int, endUser common.Address) (*types.Transaction, error) {
	return _Token1155.contract.Transact(opts, "claim", _mintAmount, _number, endUser)
}

// Claim is a paid mutator transaction binding the contract method 0xac44ff31.
//
// Solidity: function claim(uint256 _mintAmount, uint256 _number, address endUser) returns()
func (_Token1155 *Token1155Session) Claim(_mintAmount *big.Int, _number *big.Int, endUser common.Address) (*types.Transaction, error) {
	return _Token1155.Contract.Claim(&_Token1155.TransactOpts, _mintAmount, _number, endUser)
}

// Claim is a paid mutator transaction binding the contract method 0xac44ff31.
//
// Solidity: function claim(uint256 _mintAmount, uint256 _number, address endUser) returns()
func (_Token1155 *Token1155TransactorSession) Claim(_mintAmount *big.Int, _number *big.Int, endUser common.Address) (*types.Transaction, error) {
	return _Token1155.Contract.Claim(&_Token1155.TransactOpts, _mintAmount, _number, endUser)
}

// ClaimBatch is a paid mutator transaction binding the contract method 0xd7c2dfc8.
//
// Solidity: function claimBatch(address[] endUser, uint256[] amount) returns()
func (_Token1155 *Token1155Transactor) ClaimBatch(opts *bind.TransactOpts, endUser []common.Address, amount []*big.Int) (*types.Transaction, error) {
	return _Token1155.contract.Transact(opts, "claimBatch", endUser, amount)
}

// ClaimBatch is a paid mutator transaction binding the contract method 0xd7c2dfc8.
//
// Solidity: function claimBatch(address[] endUser, uint256[] amount) returns()
func (_Token1155 *Token1155Session) ClaimBatch(endUser []common.Address, amount []*big.Int) (*types.Transaction, error) {
	return _Token1155.Contract.ClaimBatch(&_Token1155.TransactOpts, endUser, amount)
}

// ClaimBatch is a paid mutator transaction binding the contract method 0xd7c2dfc8.
//
// Solidity: function claimBatch(address[] endUser, uint256[] amount) returns()
func (_Token1155 *Token1155TransactorSession) ClaimBatch(endUser []common.Address, amount []*big.Int) (*types.Transaction, error) {
	return _Token1155.Contract.ClaimBatch(&_Token1155.TransactOpts, endUser, amount)
}

// Pause is a paid mutator transaction binding the contract method 0x02329a29.
//
// Solidity: function pause(bool _state) returns()
func (_Token1155 *Token1155Transactor) Pause(opts *bind.TransactOpts, _state bool) (*types.Transaction, error) {
	return _Token1155.contract.Transact(opts, "pause", _state)
}

// Pause is a paid mutator transaction binding the contract method 0x02329a29.
//
// Solidity: function pause(bool _state) returns()
func (_Token1155 *Token1155Session) Pause(_state bool) (*types.Transaction, error) {
	return _Token1155.Contract.Pause(&_Token1155.TransactOpts, _state)
}

// Pause is a paid mutator transaction binding the contract method 0x02329a29.
//
// Solidity: function pause(bool _state) returns()
func (_Token1155 *Token1155TransactorSession) Pause(_state bool) (*types.Transaction, error) {
	return _Token1155.Contract.Pause(&_Token1155.TransactOpts, _state)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Token1155 *Token1155Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token1155.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Token1155 *Token1155Session) RenounceOwnership() (*types.Transaction, error) {
	return _Token1155.Contract.RenounceOwnership(&_Token1155.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Token1155 *Token1155TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Token1155.Contract.RenounceOwnership(&_Token1155.TransactOpts)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Token1155 *Token1155Transactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Token1155.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Token1155 *Token1155Session) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Token1155.Contract.SafeBatchTransferFrom(&_Token1155.TransactOpts, from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Token1155 *Token1155TransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Token1155.Contract.SafeBatchTransferFrom(&_Token1155.TransactOpts, from, to, ids, amounts, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Token1155 *Token1155Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Token1155.contract.Transact(opts, "safeTransferFrom", from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Token1155 *Token1155Session) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Token1155.Contract.SafeTransferFrom(&_Token1155.TransactOpts, from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Token1155 *Token1155TransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Token1155.Contract.SafeTransferFrom(&_Token1155.TransactOpts, from, to, id, amount, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Token1155 *Token1155Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Token1155.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Token1155 *Token1155Session) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Token1155.Contract.SetApprovalForAll(&_Token1155.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Token1155 *Token1155TransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Token1155.Contract.SetApprovalForAll(&_Token1155.TransactOpts, operator, approved)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string _newBaseURI) returns()
func (_Token1155 *Token1155Transactor) SetBaseURI(opts *bind.TransactOpts, _newBaseURI string) (*types.Transaction, error) {
	return _Token1155.contract.Transact(opts, "setBaseURI", _newBaseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string _newBaseURI) returns()
func (_Token1155 *Token1155Session) SetBaseURI(_newBaseURI string) (*types.Transaction, error) {
	return _Token1155.Contract.SetBaseURI(&_Token1155.TransactOpts, _newBaseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string _newBaseURI) returns()
func (_Token1155 *Token1155TransactorSession) SetBaseURI(_newBaseURI string) (*types.Transaction, error) {
	return _Token1155.Contract.SetBaseURI(&_Token1155.TransactOpts, _newBaseURI)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string contractURI_) returns()
func (_Token1155 *Token1155Transactor) SetContractURI(opts *bind.TransactOpts, contractURI_ string) (*types.Transaction, error) {
	return _Token1155.contract.Transact(opts, "setContractURI", contractURI_)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string contractURI_) returns()
func (_Token1155 *Token1155Session) SetContractURI(contractURI_ string) (*types.Transaction, error) {
	return _Token1155.Contract.SetContractURI(&_Token1155.TransactOpts, contractURI_)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string contractURI_) returns()
func (_Token1155 *Token1155TransactorSession) SetContractURI(contractURI_ string) (*types.Transaction, error) {
	return _Token1155.Contract.SetContractURI(&_Token1155.TransactOpts, contractURI_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Token1155 *Token1155Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Token1155.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Token1155 *Token1155Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Token1155.Contract.TransferOwnership(&_Token1155.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Token1155 *Token1155TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Token1155.Contract.TransferOwnership(&_Token1155.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() payable returns()
func (_Token1155 *Token1155Transactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token1155.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() payable returns()
func (_Token1155 *Token1155Session) Withdraw() (*types.Transaction, error) {
	return _Token1155.Contract.Withdraw(&_Token1155.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() payable returns()
func (_Token1155 *Token1155TransactorSession) Withdraw() (*types.Transaction, error) {
	return _Token1155.Contract.Withdraw(&_Token1155.TransactOpts)
}

// Token1155ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Token1155 contract.
type Token1155ApprovalForAllIterator struct {
	Event *Token1155ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *Token1155ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Token1155ApprovalForAll)
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
		it.Event = new(Token1155ApprovalForAll)
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
func (it *Token1155ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Token1155ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Token1155ApprovalForAll represents a ApprovalForAll event raised by the Token1155 contract.
type Token1155ApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_Token1155 *Token1155Filterer) FilterApprovalForAll(opts *bind.FilterOpts, account []common.Address, operator []common.Address) (*Token1155ApprovalForAllIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Token1155.contract.FilterLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &Token1155ApprovalForAllIterator{contract: _Token1155.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_Token1155 *Token1155Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *Token1155ApprovalForAll, account []common.Address, operator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Token1155.contract.WatchLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Token1155ApprovalForAll)
				if err := _Token1155.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_Token1155 *Token1155Filterer) ParseApprovalForAll(log types.Log) (*Token1155ApprovalForAll, error) {
	event := new(Token1155ApprovalForAll)
	if err := _Token1155.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Token1155OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Token1155 contract.
type Token1155OwnershipTransferredIterator struct {
	Event *Token1155OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *Token1155OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Token1155OwnershipTransferred)
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
		it.Event = new(Token1155OwnershipTransferred)
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
func (it *Token1155OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Token1155OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Token1155OwnershipTransferred represents a OwnershipTransferred event raised by the Token1155 contract.
type Token1155OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Token1155 *Token1155Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*Token1155OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Token1155.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &Token1155OwnershipTransferredIterator{contract: _Token1155.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Token1155 *Token1155Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *Token1155OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Token1155.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Token1155OwnershipTransferred)
				if err := _Token1155.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Token1155 *Token1155Filterer) ParseOwnershipTransferred(log types.Log) (*Token1155OwnershipTransferred, error) {
	event := new(Token1155OwnershipTransferred)
	if err := _Token1155.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Token1155TransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the Token1155 contract.
type Token1155TransferBatchIterator struct {
	Event *Token1155TransferBatch // Event containing the contract specifics and raw log

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
func (it *Token1155TransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Token1155TransferBatch)
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
		it.Event = new(Token1155TransferBatch)
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
func (it *Token1155TransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Token1155TransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Token1155TransferBatch represents a TransferBatch event raised by the Token1155 contract.
type Token1155TransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_Token1155 *Token1155Filterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*Token1155TransferBatchIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Token1155.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Token1155TransferBatchIterator{contract: _Token1155.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_Token1155 *Token1155Filterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *Token1155TransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Token1155.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Token1155TransferBatch)
				if err := _Token1155.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_Token1155 *Token1155Filterer) ParseTransferBatch(log types.Log) (*Token1155TransferBatch, error) {
	event := new(Token1155TransferBatch)
	if err := _Token1155.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Token1155TransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the Token1155 contract.
type Token1155TransferSingleIterator struct {
	Event *Token1155TransferSingle // Event containing the contract specifics and raw log

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
func (it *Token1155TransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Token1155TransferSingle)
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
		it.Event = new(Token1155TransferSingle)
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
func (it *Token1155TransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Token1155TransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Token1155TransferSingle represents a TransferSingle event raised by the Token1155 contract.
type Token1155TransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_Token1155 *Token1155Filterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*Token1155TransferSingleIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Token1155.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Token1155TransferSingleIterator{contract: _Token1155.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_Token1155 *Token1155Filterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *Token1155TransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Token1155.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Token1155TransferSingle)
				if err := _Token1155.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_Token1155 *Token1155Filterer) ParseTransferSingle(log types.Log) (*Token1155TransferSingle, error) {
	event := new(Token1155TransferSingle)
	if err := _Token1155.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Token1155URIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the Token1155 contract.
type Token1155URIIterator struct {
	Event *Token1155URI // Event containing the contract specifics and raw log

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
func (it *Token1155URIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Token1155URI)
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
		it.Event = new(Token1155URI)
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
func (it *Token1155URIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Token1155URIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Token1155URI represents a URI event raised by the Token1155 contract.
type Token1155URI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Token1155 *Token1155Filterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*Token1155URIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Token1155.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &Token1155URIIterator{contract: _Token1155.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Token1155 *Token1155Filterer) WatchURI(opts *bind.WatchOpts, sink chan<- *Token1155URI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Token1155.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Token1155URI)
				if err := _Token1155.contract.UnpackLog(event, "URI", log); err != nil {
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

// ParseURI is a log parse operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Token1155 *Token1155Filterer) ParseURI(log types.Log) (*Token1155URI, error) {
	event := new(Token1155URI)
	if err := _Token1155.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Token1155TransferSendIterator is returned from FilterTransferSend and is used to iterate over the raw logs and unpacked data for TransferSend events raised by the Token1155 contract.
type Token1155TransferSendIterator struct {
	Event *Token1155TransferSend // Event containing the contract specifics and raw log

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
func (it *Token1155TransferSendIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Token1155TransferSend)
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
		it.Event = new(Token1155TransferSend)
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
func (it *Token1155TransferSendIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Token1155TransferSendIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Token1155TransferSend represents a TransferSend event raised by the Token1155 contract.
type Token1155TransferSend struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransferSend is a free log retrieval operation binding the contract event 0xf70a742b6ec5de639ce8e8d0df10a44a88c323ac660a09f1fc0c190e830aea74.
//
// Solidity: event _transferSend(address _from, address _to, uint256 _amount)
func (_Token1155 *Token1155Filterer) FilterTransferSend(opts *bind.FilterOpts) (*Token1155TransferSendIterator, error) {

	logs, sub, err := _Token1155.contract.FilterLogs(opts, "_transferSend")
	if err != nil {
		return nil, err
	}
	return &Token1155TransferSendIterator{contract: _Token1155.contract, event: "_transferSend", logs: logs, sub: sub}, nil
}

// WatchTransferSend is a free log subscription operation binding the contract event 0xf70a742b6ec5de639ce8e8d0df10a44a88c323ac660a09f1fc0c190e830aea74.
//
// Solidity: event _transferSend(address _from, address _to, uint256 _amount)
func (_Token1155 *Token1155Filterer) WatchTransferSend(opts *bind.WatchOpts, sink chan<- *Token1155TransferSend) (event.Subscription, error) {

	logs, sub, err := _Token1155.contract.WatchLogs(opts, "_transferSend")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Token1155TransferSend)
				if err := _Token1155.contract.UnpackLog(event, "_transferSend", log); err != nil {
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

// ParseTransferSend is a log parse operation binding the contract event 0xf70a742b6ec5de639ce8e8d0df10a44a88c323ac660a09f1fc0c190e830aea74.
//
// Solidity: event _transferSend(address _from, address _to, uint256 _amount)
func (_Token1155 *Token1155Filterer) ParseTransferSend(log types.Log) (*Token1155TransferSend, error) {
	event := new(Token1155TransferSend)
	if err := _Token1155.contract.UnpackLog(event, "_transferSend", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
