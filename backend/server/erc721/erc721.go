// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc721

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

// Token721MetaData contains all meta data concerning the Token721 contract.
var Token721MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_initBaseURI\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"_contractURI\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ApprovalCallerNotOwnerNorApproved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ApprovalQueryForNonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ApprovalToCurrentOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ApproveToCaller\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BalanceQueryForZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintToZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintZeroQuantity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerQueryForNonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferCallerNotOwnerNorApproved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFromIncorrectOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToNonERC721ReceiverImplementer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"URIQueryForNonexistentToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"_transferSend\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"_checkContractOnERC721Received\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"_transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractIsFreeMint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"getAddressLastMintedTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBaseURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxMintAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getNFTURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"getNftEtherValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getQtdAvailableTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"isAffliliated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOnlyWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"}],\"name\":\"isValid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"isWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"merkleProof\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_mintAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"_recommendedBy\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_indicationType\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"endUser\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_state\",\"type\":\"bool\"}],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"root\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"state\",\"type\":\"bool\"}],\"name\":\"setAffiliate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"percentage\",\"type\":\"uint256\"}],\"name\":\"setAffiliateProgramPercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_state\",\"type\":\"bool\"}],\"name\":\"setAllowAffiliateProgram\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_newBaseURI\",\"type\":\"string\"}],\"name\":\"setBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"contractURI_\",\"type\":\"string\"}],\"name\":\"setContractURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_limit\",\"type\":\"uint256\"}],\"name\":\"setFreeMintPerAddressLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"state\",\"type\":\"bool\"}],\"name\":\"setIsFreeMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxMintAmount\",\"type\":\"uint256\"}],\"name\":\"setMaxMintAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxSupply\",\"type\":\"uint256\"}],\"name\":\"setMaxSupply\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nftEtherValue\",\"type\":\"uint256\"}],\"name\":\"setNftEtherValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_state\",\"type\":\"bool\"}],\"name\":\"setOnlyWhitelisted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_mintAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"_recommendedBy\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_indicationType\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"split\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"baseExtension\",\"type\":\"string\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"walletOfOwner\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelistClaimed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"whitelistedAddresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// Token721ABI is the input ABI used to generate the binding from.
// Deprecated: Use Token721MetaData.ABI instead.
var Token721ABI = Token721MetaData.ABI

// Token721 is an auto generated Go binding around an Ethereum contract.
type Token721 struct {
	Token721Caller     // Read-only binding to the contract
	Token721Transactor // Write-only binding to the contract
	Token721Filterer   // Log filterer for contract events
}

// Token721Caller is an auto generated read-only Go binding around an Ethereum contract.
type Token721Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Token721Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Token721Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Token721Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Token721Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Token721Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Token721Session struct {
	Contract     *Token721         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Token721CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Token721CallerSession struct {
	Contract *Token721Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// Token721TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Token721TransactorSession struct {
	Contract     *Token721Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// Token721Raw is an auto generated low-level Go binding around an Ethereum contract.
type Token721Raw struct {
	Contract *Token721 // Generic contract binding to access the raw methods on
}

// Token721CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Token721CallerRaw struct {
	Contract *Token721Caller // Generic read-only contract binding to access the raw methods on
}

// Token721TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Token721TransactorRaw struct {
	Contract *Token721Transactor // Generic write-only contract binding to access the raw methods on
}

// NewToken721 creates a new instance of Token721, bound to a specific deployed contract.
func NewToken721(address common.Address, backend bind.ContractBackend) (*Token721, error) {
	contract, err := bindToken721(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Token721{Token721Caller: Token721Caller{contract: contract}, Token721Transactor: Token721Transactor{contract: contract}, Token721Filterer: Token721Filterer{contract: contract}}, nil
}

// NewToken721Caller creates a new read-only instance of Token721, bound to a specific deployed contract.
func NewToken721Caller(address common.Address, caller bind.ContractCaller) (*Token721Caller, error) {
	contract, err := bindToken721(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Token721Caller{contract: contract}, nil
}

// NewToken721Transactor creates a new write-only instance of Token721, bound to a specific deployed contract.
func NewToken721Transactor(address common.Address, transactor bind.ContractTransactor) (*Token721Transactor, error) {
	contract, err := bindToken721(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Token721Transactor{contract: contract}, nil
}

// NewToken721Filterer creates a new log filterer instance of Token721, bound to a specific deployed contract.
func NewToken721Filterer(address common.Address, filterer bind.ContractFilterer) (*Token721Filterer, error) {
	contract, err := bindToken721(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Token721Filterer{contract: contract}, nil
}

// bindToken721 binds a generic wrapper to an already deployed contract.
func bindToken721(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Token721ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token721 *Token721Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Token721.Contract.Token721Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token721 *Token721Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token721.Contract.Token721Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token721 *Token721Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token721.Contract.Token721Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token721 *Token721CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Token721.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token721 *Token721TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token721.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token721 *Token721TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token721.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Token721 *Token721Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Token721.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Token721 *Token721Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Token721.Contract.BalanceOf(&_Token721.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Token721 *Token721CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Token721.Contract.BalanceOf(&_Token721.CallOpts, owner)
}

// ContractIsFreeMint is a free data retrieval call binding the contract method 0x4fc9f84a.
//
// Solidity: function contractIsFreeMint() view returns(bool)
func (_Token721 *Token721Caller) ContractIsFreeMint(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Token721.contract.Call(opts, &out, "contractIsFreeMint")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ContractIsFreeMint is a free data retrieval call binding the contract method 0x4fc9f84a.
//
// Solidity: function contractIsFreeMint() view returns(bool)
func (_Token721 *Token721Session) ContractIsFreeMint() (bool, error) {
	return _Token721.Contract.ContractIsFreeMint(&_Token721.CallOpts)
}

// ContractIsFreeMint is a free data retrieval call binding the contract method 0x4fc9f84a.
//
// Solidity: function contractIsFreeMint() view returns(bool)
func (_Token721 *Token721CallerSession) ContractIsFreeMint() (bool, error) {
	return _Token721.Contract.ContractIsFreeMint(&_Token721.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Token721 *Token721Caller) ContractURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Token721.contract.Call(opts, &out, "contractURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Token721 *Token721Session) ContractURI() (string, error) {
	return _Token721.Contract.ContractURI(&_Token721.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Token721 *Token721CallerSession) ContractURI() (string, error) {
	return _Token721.Contract.ContractURI(&_Token721.CallOpts)
}

// GetAddressLastMintedTokenId is a free data retrieval call binding the contract method 0x6c21fb00.
//
// Solidity: function getAddressLastMintedTokenId(address wallet) view returns(uint256)
func (_Token721 *Token721Caller) GetAddressLastMintedTokenId(opts *bind.CallOpts, wallet common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Token721.contract.Call(opts, &out, "getAddressLastMintedTokenId", wallet)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAddressLastMintedTokenId is a free data retrieval call binding the contract method 0x6c21fb00.
//
// Solidity: function getAddressLastMintedTokenId(address wallet) view returns(uint256)
func (_Token721 *Token721Session) GetAddressLastMintedTokenId(wallet common.Address) (*big.Int, error) {
	return _Token721.Contract.GetAddressLastMintedTokenId(&_Token721.CallOpts, wallet)
}

// GetAddressLastMintedTokenId is a free data retrieval call binding the contract method 0x6c21fb00.
//
// Solidity: function getAddressLastMintedTokenId(address wallet) view returns(uint256)
func (_Token721 *Token721CallerSession) GetAddressLastMintedTokenId(wallet common.Address) (*big.Int, error) {
	return _Token721.Contract.GetAddressLastMintedTokenId(&_Token721.CallOpts, wallet)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Token721 *Token721Caller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Token721.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Token721 *Token721Session) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Token721.Contract.GetApproved(&_Token721.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Token721 *Token721CallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Token721.Contract.GetApproved(&_Token721.CallOpts, tokenId)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Token721 *Token721Caller) GetBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token721.contract.Call(opts, &out, "getBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Token721 *Token721Session) GetBalance() (*big.Int, error) {
	return _Token721.Contract.GetBalance(&_Token721.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Token721 *Token721CallerSession) GetBalance() (*big.Int, error) {
	return _Token721.Contract.GetBalance(&_Token721.CallOpts)
}

// GetBaseURI is a free data retrieval call binding the contract method 0x714c5398.
//
// Solidity: function getBaseURI() view returns(string)
func (_Token721 *Token721Caller) GetBaseURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Token721.contract.Call(opts, &out, "getBaseURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetBaseURI is a free data retrieval call binding the contract method 0x714c5398.
//
// Solidity: function getBaseURI() view returns(string)
func (_Token721 *Token721Session) GetBaseURI() (string, error) {
	return _Token721.Contract.GetBaseURI(&_Token721.CallOpts)
}

// GetBaseURI is a free data retrieval call binding the contract method 0x714c5398.
//
// Solidity: function getBaseURI() view returns(string)
func (_Token721 *Token721CallerSession) GetBaseURI() (string, error) {
	return _Token721.Contract.GetBaseURI(&_Token721.CallOpts)
}

// GetMaxMintAmount is a free data retrieval call binding the contract method 0xd6e7b8e4.
//
// Solidity: function getMaxMintAmount() view returns(uint256)
func (_Token721 *Token721Caller) GetMaxMintAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token721.contract.Call(opts, &out, "getMaxMintAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxMintAmount is a free data retrieval call binding the contract method 0xd6e7b8e4.
//
// Solidity: function getMaxMintAmount() view returns(uint256)
func (_Token721 *Token721Session) GetMaxMintAmount() (*big.Int, error) {
	return _Token721.Contract.GetMaxMintAmount(&_Token721.CallOpts)
}

// GetMaxMintAmount is a free data retrieval call binding the contract method 0xd6e7b8e4.
//
// Solidity: function getMaxMintAmount() view returns(uint256)
func (_Token721 *Token721CallerSession) GetMaxMintAmount() (*big.Int, error) {
	return _Token721.Contract.GetMaxMintAmount(&_Token721.CallOpts)
}

// GetMaxSupply is a free data retrieval call binding the contract method 0x4c0f38c2.
//
// Solidity: function getMaxSupply() view returns(uint256)
func (_Token721 *Token721Caller) GetMaxSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token721.contract.Call(opts, &out, "getMaxSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxSupply is a free data retrieval call binding the contract method 0x4c0f38c2.
//
// Solidity: function getMaxSupply() view returns(uint256)
func (_Token721 *Token721Session) GetMaxSupply() (*big.Int, error) {
	return _Token721.Contract.GetMaxSupply(&_Token721.CallOpts)
}

// GetMaxSupply is a free data retrieval call binding the contract method 0x4c0f38c2.
//
// Solidity: function getMaxSupply() view returns(uint256)
func (_Token721 *Token721CallerSession) GetMaxSupply() (*big.Int, error) {
	return _Token721.Contract.GetMaxSupply(&_Token721.CallOpts)
}

// GetNFTURI is a free data retrieval call binding the contract method 0x00621fda.
//
// Solidity: function getNFTURI(uint256 tokenId) view returns(string)
func (_Token721 *Token721Caller) GetNFTURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Token721.contract.Call(opts, &out, "getNFTURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetNFTURI is a free data retrieval call binding the contract method 0x00621fda.
//
// Solidity: function getNFTURI(uint256 tokenId) view returns(string)
func (_Token721 *Token721Session) GetNFTURI(tokenId *big.Int) (string, error) {
	return _Token721.Contract.GetNFTURI(&_Token721.CallOpts, tokenId)
}

// GetNFTURI is a free data retrieval call binding the contract method 0x00621fda.
//
// Solidity: function getNFTURI(uint256 tokenId) view returns(string)
func (_Token721 *Token721CallerSession) GetNFTURI(tokenId *big.Int) (string, error) {
	return _Token721.Contract.GetNFTURI(&_Token721.CallOpts, tokenId)
}

// GetNftEtherValue is a free data retrieval call binding the contract method 0x77dbd17f.
//
// Solidity: function getNftEtherValue(address wallet, bytes32[] proof) view returns(uint256)
func (_Token721 *Token721Caller) GetNftEtherValue(opts *bind.CallOpts, wallet common.Address, proof [][32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Token721.contract.Call(opts, &out, "getNftEtherValue", wallet, proof)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNftEtherValue is a free data retrieval call binding the contract method 0x77dbd17f.
//
// Solidity: function getNftEtherValue(address wallet, bytes32[] proof) view returns(uint256)
func (_Token721 *Token721Session) GetNftEtherValue(wallet common.Address, proof [][32]byte) (*big.Int, error) {
	return _Token721.Contract.GetNftEtherValue(&_Token721.CallOpts, wallet, proof)
}

// GetNftEtherValue is a free data retrieval call binding the contract method 0x77dbd17f.
//
// Solidity: function getNftEtherValue(address wallet, bytes32[] proof) view returns(uint256)
func (_Token721 *Token721CallerSession) GetNftEtherValue(wallet common.Address, proof [][32]byte) (*big.Int, error) {
	return _Token721.Contract.GetNftEtherValue(&_Token721.CallOpts, wallet, proof)
}

// GetQtdAvailableTokens is a free data retrieval call binding the contract method 0xb4ed16b3.
//
// Solidity: function getQtdAvailableTokens() view returns(uint256)
func (_Token721 *Token721Caller) GetQtdAvailableTokens(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token721.contract.Call(opts, &out, "getQtdAvailableTokens")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQtdAvailableTokens is a free data retrieval call binding the contract method 0xb4ed16b3.
//
// Solidity: function getQtdAvailableTokens() view returns(uint256)
func (_Token721 *Token721Session) GetQtdAvailableTokens() (*big.Int, error) {
	return _Token721.Contract.GetQtdAvailableTokens(&_Token721.CallOpts)
}

// GetQtdAvailableTokens is a free data retrieval call binding the contract method 0xb4ed16b3.
//
// Solidity: function getQtdAvailableTokens() view returns(uint256)
func (_Token721 *Token721CallerSession) GetQtdAvailableTokens() (*big.Int, error) {
	return _Token721.Contract.GetQtdAvailableTokens(&_Token721.CallOpts)
}

// IsAffliliated is a free data retrieval call binding the contract method 0x9b0e665d.
//
// Solidity: function isAffliliated(address wallet) view returns(bool)
func (_Token721 *Token721Caller) IsAffliliated(opts *bind.CallOpts, wallet common.Address) (bool, error) {
	var out []interface{}
	err := _Token721.contract.Call(opts, &out, "isAffliliated", wallet)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAffliliated is a free data retrieval call binding the contract method 0x9b0e665d.
//
// Solidity: function isAffliliated(address wallet) view returns(bool)
func (_Token721 *Token721Session) IsAffliliated(wallet common.Address) (bool, error) {
	return _Token721.Contract.IsAffliliated(&_Token721.CallOpts, wallet)
}

// IsAffliliated is a free data retrieval call binding the contract method 0x9b0e665d.
//
// Solidity: function isAffliliated(address wallet) view returns(bool)
func (_Token721 *Token721CallerSession) IsAffliliated(wallet common.Address) (bool, error) {
	return _Token721.Contract.IsAffliliated(&_Token721.CallOpts, wallet)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Token721 *Token721Caller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Token721.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Token721 *Token721Session) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Token721.Contract.IsApprovedForAll(&_Token721.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Token721 *Token721CallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Token721.Contract.IsApprovedForAll(&_Token721.CallOpts, owner, operator)
}

// IsOnlyWhitelist is a free data retrieval call binding the contract method 0x6f50c93c.
//
// Solidity: function isOnlyWhitelist() view returns(bool)
func (_Token721 *Token721Caller) IsOnlyWhitelist(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Token721.contract.Call(opts, &out, "isOnlyWhitelist")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOnlyWhitelist is a free data retrieval call binding the contract method 0x6f50c93c.
//
// Solidity: function isOnlyWhitelist() view returns(bool)
func (_Token721 *Token721Session) IsOnlyWhitelist() (bool, error) {
	return _Token721.Contract.IsOnlyWhitelist(&_Token721.CallOpts)
}

// IsOnlyWhitelist is a free data retrieval call binding the contract method 0x6f50c93c.
//
// Solidity: function isOnlyWhitelist() view returns(bool)
func (_Token721 *Token721CallerSession) IsOnlyWhitelist() (bool, error) {
	return _Token721.Contract.IsOnlyWhitelist(&_Token721.CallOpts)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Token721 *Token721Caller) IsPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Token721.contract.Call(opts, &out, "isPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Token721 *Token721Session) IsPaused() (bool, error) {
	return _Token721.Contract.IsPaused(&_Token721.CallOpts)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Token721 *Token721CallerSession) IsPaused() (bool, error) {
	return _Token721.Contract.IsPaused(&_Token721.CallOpts)
}

// IsValid is a free data retrieval call binding the contract method 0xb8a20ed0.
//
// Solidity: function isValid(bytes32[] proof, bytes32 leaf) view returns(bool)
func (_Token721 *Token721Caller) IsValid(opts *bind.CallOpts, proof [][32]byte, leaf [32]byte) (bool, error) {
	var out []interface{}
	err := _Token721.contract.Call(opts, &out, "isValid", proof, leaf)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValid is a free data retrieval call binding the contract method 0xb8a20ed0.
//
// Solidity: function isValid(bytes32[] proof, bytes32 leaf) view returns(bool)
func (_Token721 *Token721Session) IsValid(proof [][32]byte, leaf [32]byte) (bool, error) {
	return _Token721.Contract.IsValid(&_Token721.CallOpts, proof, leaf)
}

// IsValid is a free data retrieval call binding the contract method 0xb8a20ed0.
//
// Solidity: function isValid(bytes32[] proof, bytes32 leaf) view returns(bool)
func (_Token721 *Token721CallerSession) IsValid(proof [][32]byte, leaf [32]byte) (bool, error) {
	return _Token721.Contract.IsValid(&_Token721.CallOpts, proof, leaf)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x5a23dd99.
//
// Solidity: function isWhitelisted(address _user, bytes32[] proof) view returns(bool)
func (_Token721 *Token721Caller) IsWhitelisted(opts *bind.CallOpts, _user common.Address, proof [][32]byte) (bool, error) {
	var out []interface{}
	err := _Token721.contract.Call(opts, &out, "isWhitelisted", _user, proof)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWhitelisted is a free data retrieval call binding the contract method 0x5a23dd99.
//
// Solidity: function isWhitelisted(address _user, bytes32[] proof) view returns(bool)
func (_Token721 *Token721Session) IsWhitelisted(_user common.Address, proof [][32]byte) (bool, error) {
	return _Token721.Contract.IsWhitelisted(&_Token721.CallOpts, _user, proof)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x5a23dd99.
//
// Solidity: function isWhitelisted(address _user, bytes32[] proof) view returns(bool)
func (_Token721 *Token721CallerSession) IsWhitelisted(_user common.Address, proof [][32]byte) (bool, error) {
	return _Token721.Contract.IsWhitelisted(&_Token721.CallOpts, _user, proof)
}

// MerkleProof is a free data retrieval call binding the contract method 0xfa9a309d.
//
// Solidity: function merkleProof(uint256 ) view returns(bytes32)
func (_Token721 *Token721Caller) MerkleProof(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Token721.contract.Call(opts, &out, "merkleProof", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MerkleProof is a free data retrieval call binding the contract method 0xfa9a309d.
//
// Solidity: function merkleProof(uint256 ) view returns(bytes32)
func (_Token721 *Token721Session) MerkleProof(arg0 *big.Int) ([32]byte, error) {
	return _Token721.Contract.MerkleProof(&_Token721.CallOpts, arg0)
}

// MerkleProof is a free data retrieval call binding the contract method 0xfa9a309d.
//
// Solidity: function merkleProof(uint256 ) view returns(bytes32)
func (_Token721 *Token721CallerSession) MerkleProof(arg0 *big.Int) ([32]byte, error) {
	return _Token721.Contract.MerkleProof(&_Token721.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Token721 *Token721Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Token721.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Token721 *Token721Session) Name() (string, error) {
	return _Token721.Contract.Name(&_Token721.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Token721 *Token721CallerSession) Name() (string, error) {
	return _Token721.Contract.Name(&_Token721.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Token721 *Token721Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Token721.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Token721 *Token721Session) Owner() (common.Address, error) {
	return _Token721.Contract.Owner(&_Token721.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Token721 *Token721CallerSession) Owner() (common.Address, error) {
	return _Token721.Contract.Owner(&_Token721.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Token721 *Token721Caller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Token721.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Token721 *Token721Session) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Token721.Contract.OwnerOf(&_Token721.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Token721 *Token721CallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Token721.Contract.OwnerOf(&_Token721.CallOpts, tokenId)
}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() view returns(bytes32)
func (_Token721 *Token721Caller) Root(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Token721.contract.Call(opts, &out, "root")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() view returns(bytes32)
func (_Token721 *Token721Session) Root() ([32]byte, error) {
	return _Token721.Contract.Root(&_Token721.CallOpts)
}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() view returns(bytes32)
func (_Token721 *Token721CallerSession) Root() ([32]byte, error) {
	return _Token721.Contract.Root(&_Token721.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Token721 *Token721Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Token721.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Token721 *Token721Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Token721.Contract.SupportsInterface(&_Token721.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Token721 *Token721CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Token721.Contract.SupportsInterface(&_Token721.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Token721 *Token721Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Token721.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Token721 *Token721Session) Symbol() (string, error) {
	return _Token721.Contract.Symbol(&_Token721.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Token721 *Token721CallerSession) Symbol() (string, error) {
	return _Token721.Contract.Symbol(&_Token721.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0x41dcf454.
//
// Solidity: function tokenURI(uint256 tokenId, string baseExtension) view returns(string)
func (_Token721 *Token721Caller) TokenURI(opts *bind.CallOpts, tokenId *big.Int, baseExtension string) (string, error) {
	var out []interface{}
	err := _Token721.contract.Call(opts, &out, "tokenURI", tokenId, baseExtension)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0x41dcf454.
//
// Solidity: function tokenURI(uint256 tokenId, string baseExtension) view returns(string)
func (_Token721 *Token721Session) TokenURI(tokenId *big.Int, baseExtension string) (string, error) {
	return _Token721.Contract.TokenURI(&_Token721.CallOpts, tokenId, baseExtension)
}

// TokenURI is a free data retrieval call binding the contract method 0x41dcf454.
//
// Solidity: function tokenURI(uint256 tokenId, string baseExtension) view returns(string)
func (_Token721 *Token721CallerSession) TokenURI(tokenId *big.Int, baseExtension string) (string, error) {
	return _Token721.Contract.TokenURI(&_Token721.CallOpts, tokenId, baseExtension)
}

// TokenURI0 is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Token721 *Token721Caller) TokenURI0(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Token721.contract.Call(opts, &out, "tokenURI0", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI0 is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Token721 *Token721Session) TokenURI0(tokenId *big.Int) (string, error) {
	return _Token721.Contract.TokenURI0(&_Token721.CallOpts, tokenId)
}

// TokenURI0 is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Token721 *Token721CallerSession) TokenURI0(tokenId *big.Int) (string, error) {
	return _Token721.Contract.TokenURI0(&_Token721.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Token721 *Token721Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token721.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Token721 *Token721Session) TotalSupply() (*big.Int, error) {
	return _Token721.Contract.TotalSupply(&_Token721.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Token721 *Token721CallerSession) TotalSupply() (*big.Int, error) {
	return _Token721.Contract.TotalSupply(&_Token721.CallOpts)
}

// WalletOfOwner is a free data retrieval call binding the contract method 0x438b6300.
//
// Solidity: function walletOfOwner(address _owner) view returns(uint256)
func (_Token721 *Token721Caller) WalletOfOwner(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Token721.contract.Call(opts, &out, "walletOfOwner", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WalletOfOwner is a free data retrieval call binding the contract method 0x438b6300.
//
// Solidity: function walletOfOwner(address _owner) view returns(uint256)
func (_Token721 *Token721Session) WalletOfOwner(_owner common.Address) (*big.Int, error) {
	return _Token721.Contract.WalletOfOwner(&_Token721.CallOpts, _owner)
}

// WalletOfOwner is a free data retrieval call binding the contract method 0x438b6300.
//
// Solidity: function walletOfOwner(address _owner) view returns(uint256)
func (_Token721 *Token721CallerSession) WalletOfOwner(_owner common.Address) (*big.Int, error) {
	return _Token721.Contract.WalletOfOwner(&_Token721.CallOpts, _owner)
}

// WhitelistClaimed is a free data retrieval call binding the contract method 0xdb4bec44.
//
// Solidity: function whitelistClaimed(address ) view returns(bool)
func (_Token721 *Token721Caller) WhitelistClaimed(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Token721.contract.Call(opts, &out, "whitelistClaimed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WhitelistClaimed is a free data retrieval call binding the contract method 0xdb4bec44.
//
// Solidity: function whitelistClaimed(address ) view returns(bool)
func (_Token721 *Token721Session) WhitelistClaimed(arg0 common.Address) (bool, error) {
	return _Token721.Contract.WhitelistClaimed(&_Token721.CallOpts, arg0)
}

// WhitelistClaimed is a free data retrieval call binding the contract method 0xdb4bec44.
//
// Solidity: function whitelistClaimed(address ) view returns(bool)
func (_Token721 *Token721CallerSession) WhitelistClaimed(arg0 common.Address) (bool, error) {
	return _Token721.Contract.WhitelistClaimed(&_Token721.CallOpts, arg0)
}

// WhitelistedAddresses is a free data retrieval call binding the contract method 0xba4e5c49.
//
// Solidity: function whitelistedAddresses(uint256 ) view returns(address)
func (_Token721 *Token721Caller) WhitelistedAddresses(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Token721.contract.Call(opts, &out, "whitelistedAddresses", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WhitelistedAddresses is a free data retrieval call binding the contract method 0xba4e5c49.
//
// Solidity: function whitelistedAddresses(uint256 ) view returns(address)
func (_Token721 *Token721Session) WhitelistedAddresses(arg0 *big.Int) (common.Address, error) {
	return _Token721.Contract.WhitelistedAddresses(&_Token721.CallOpts, arg0)
}

// WhitelistedAddresses is a free data retrieval call binding the contract method 0xba4e5c49.
//
// Solidity: function whitelistedAddresses(uint256 ) view returns(address)
func (_Token721 *Token721CallerSession) WhitelistedAddresses(arg0 *big.Int) (common.Address, error) {
	return _Token721.Contract.WhitelistedAddresses(&_Token721.CallOpts, arg0)
}

// CheckContractOnERC721Received is a paid mutator transaction binding the contract method 0xd88343e2.
//
// Solidity: function _checkContractOnERC721Received(address from, address to, uint256 tokenId, bytes _data) returns(bool)
func (_Token721 *Token721Transactor) CheckContractOnERC721Received(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Token721.contract.Transact(opts, "_checkContractOnERC721Received", from, to, tokenId, _data)
}

// CheckContractOnERC721Received is a paid mutator transaction binding the contract method 0xd88343e2.
//
// Solidity: function _checkContractOnERC721Received(address from, address to, uint256 tokenId, bytes _data) returns(bool)
func (_Token721 *Token721Session) CheckContractOnERC721Received(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Token721.Contract.CheckContractOnERC721Received(&_Token721.TransactOpts, from, to, tokenId, _data)
}

// CheckContractOnERC721Received is a paid mutator transaction binding the contract method 0xd88343e2.
//
// Solidity: function _checkContractOnERC721Received(address from, address to, uint256 tokenId, bytes _data) returns(bool)
func (_Token721 *Token721TransactorSession) CheckContractOnERC721Received(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Token721.Contract.CheckContractOnERC721Received(&_Token721.TransactOpts, from, to, tokenId, _data)
}

// Transfer is a paid mutator transaction binding the contract method 0x30e0789e.
//
// Solidity: function _transfer(address from, address to, uint256 tokenId) returns()
func (_Token721 *Token721Transactor) Transfer(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Token721.contract.Transact(opts, "_transfer", from, to, tokenId)
}

// Transfer is a paid mutator transaction binding the contract method 0x30e0789e.
//
// Solidity: function _transfer(address from, address to, uint256 tokenId) returns()
func (_Token721 *Token721Session) Transfer(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Token721.Contract.Transfer(&_Token721.TransactOpts, from, to, tokenId)
}

// Transfer is a paid mutator transaction binding the contract method 0x30e0789e.
//
// Solidity: function _transfer(address from, address to, uint256 tokenId) returns()
func (_Token721 *Token721TransactorSession) Transfer(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Token721.Contract.Transfer(&_Token721.TransactOpts, from, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Token721 *Token721Transactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Token721.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Token721 *Token721Session) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Token721.Contract.Approve(&_Token721.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Token721 *Token721TransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Token721.Contract.Approve(&_Token721.TransactOpts, to, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x9293237d.
//
// Solidity: function mint(uint256 _mintAmount, address _recommendedBy, uint256 _indicationType, address endUser, bytes32[] proof) payable returns()
func (_Token721 *Token721Transactor) Mint(opts *bind.TransactOpts, _mintAmount *big.Int, _recommendedBy common.Address, _indicationType *big.Int, endUser common.Address, proof [][32]byte) (*types.Transaction, error) {
	return _Token721.contract.Transact(opts, "mint", _mintAmount, _recommendedBy, _indicationType, endUser, proof)
}

// Mint is a paid mutator transaction binding the contract method 0x9293237d.
//
// Solidity: function mint(uint256 _mintAmount, address _recommendedBy, uint256 _indicationType, address endUser, bytes32[] proof) payable returns()
func (_Token721 *Token721Session) Mint(_mintAmount *big.Int, _recommendedBy common.Address, _indicationType *big.Int, endUser common.Address, proof [][32]byte) (*types.Transaction, error) {
	return _Token721.Contract.Mint(&_Token721.TransactOpts, _mintAmount, _recommendedBy, _indicationType, endUser, proof)
}

// Mint is a paid mutator transaction binding the contract method 0x9293237d.
//
// Solidity: function mint(uint256 _mintAmount, address _recommendedBy, uint256 _indicationType, address endUser, bytes32[] proof) payable returns()
func (_Token721 *Token721TransactorSession) Mint(_mintAmount *big.Int, _recommendedBy common.Address, _indicationType *big.Int, endUser common.Address, proof [][32]byte) (*types.Transaction, error) {
	return _Token721.Contract.Mint(&_Token721.TransactOpts, _mintAmount, _recommendedBy, _indicationType, endUser, proof)
}

// Pause is a paid mutator transaction binding the contract method 0x02329a29.
//
// Solidity: function pause(bool _state) returns()
func (_Token721 *Token721Transactor) Pause(opts *bind.TransactOpts, _state bool) (*types.Transaction, error) {
	return _Token721.contract.Transact(opts, "pause", _state)
}

// Pause is a paid mutator transaction binding the contract method 0x02329a29.
//
// Solidity: function pause(bool _state) returns()
func (_Token721 *Token721Session) Pause(_state bool) (*types.Transaction, error) {
	return _Token721.Contract.Pause(&_Token721.TransactOpts, _state)
}

// Pause is a paid mutator transaction binding the contract method 0x02329a29.
//
// Solidity: function pause(bool _state) returns()
func (_Token721 *Token721TransactorSession) Pause(_state bool) (*types.Transaction, error) {
	return _Token721.Contract.Pause(&_Token721.TransactOpts, _state)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Token721 *Token721Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token721.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Token721 *Token721Session) RenounceOwnership() (*types.Transaction, error) {
	return _Token721.Contract.RenounceOwnership(&_Token721.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Token721 *Token721TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Token721.Contract.RenounceOwnership(&_Token721.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Token721 *Token721Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Token721.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Token721 *Token721Session) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Token721.Contract.SafeTransferFrom(&_Token721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Token721 *Token721TransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Token721.Contract.SafeTransferFrom(&_Token721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Token721 *Token721Transactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Token721.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Token721 *Token721Session) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Token721.Contract.SafeTransferFrom0(&_Token721.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Token721 *Token721TransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Token721.Contract.SafeTransferFrom0(&_Token721.TransactOpts, from, to, tokenId, _data)
}

// SetAffiliate is a paid mutator transaction binding the contract method 0x8863ebbb.
//
// Solidity: function setAffiliate(address manager, bool state) returns()
func (_Token721 *Token721Transactor) SetAffiliate(opts *bind.TransactOpts, manager common.Address, state bool) (*types.Transaction, error) {
	return _Token721.contract.Transact(opts, "setAffiliate", manager, state)
}

// SetAffiliate is a paid mutator transaction binding the contract method 0x8863ebbb.
//
// Solidity: function setAffiliate(address manager, bool state) returns()
func (_Token721 *Token721Session) SetAffiliate(manager common.Address, state bool) (*types.Transaction, error) {
	return _Token721.Contract.SetAffiliate(&_Token721.TransactOpts, manager, state)
}

// SetAffiliate is a paid mutator transaction binding the contract method 0x8863ebbb.
//
// Solidity: function setAffiliate(address manager, bool state) returns()
func (_Token721 *Token721TransactorSession) SetAffiliate(manager common.Address, state bool) (*types.Transaction, error) {
	return _Token721.Contract.SetAffiliate(&_Token721.TransactOpts, manager, state)
}

// SetAffiliateProgramPercentage is a paid mutator transaction binding the contract method 0x474c68ec.
//
// Solidity: function setAffiliateProgramPercentage(uint256 percentage) returns()
func (_Token721 *Token721Transactor) SetAffiliateProgramPercentage(opts *bind.TransactOpts, percentage *big.Int) (*types.Transaction, error) {
	return _Token721.contract.Transact(opts, "setAffiliateProgramPercentage", percentage)
}

// SetAffiliateProgramPercentage is a paid mutator transaction binding the contract method 0x474c68ec.
//
// Solidity: function setAffiliateProgramPercentage(uint256 percentage) returns()
func (_Token721 *Token721Session) SetAffiliateProgramPercentage(percentage *big.Int) (*types.Transaction, error) {
	return _Token721.Contract.SetAffiliateProgramPercentage(&_Token721.TransactOpts, percentage)
}

// SetAffiliateProgramPercentage is a paid mutator transaction binding the contract method 0x474c68ec.
//
// Solidity: function setAffiliateProgramPercentage(uint256 percentage) returns()
func (_Token721 *Token721TransactorSession) SetAffiliateProgramPercentage(percentage *big.Int) (*types.Transaction, error) {
	return _Token721.Contract.SetAffiliateProgramPercentage(&_Token721.TransactOpts, percentage)
}

// SetAllowAffiliateProgram is a paid mutator transaction binding the contract method 0x94cf4b6a.
//
// Solidity: function setAllowAffiliateProgram(bool _state) returns()
func (_Token721 *Token721Transactor) SetAllowAffiliateProgram(opts *bind.TransactOpts, _state bool) (*types.Transaction, error) {
	return _Token721.contract.Transact(opts, "setAllowAffiliateProgram", _state)
}

// SetAllowAffiliateProgram is a paid mutator transaction binding the contract method 0x94cf4b6a.
//
// Solidity: function setAllowAffiliateProgram(bool _state) returns()
func (_Token721 *Token721Session) SetAllowAffiliateProgram(_state bool) (*types.Transaction, error) {
	return _Token721.Contract.SetAllowAffiliateProgram(&_Token721.TransactOpts, _state)
}

// SetAllowAffiliateProgram is a paid mutator transaction binding the contract method 0x94cf4b6a.
//
// Solidity: function setAllowAffiliateProgram(bool _state) returns()
func (_Token721 *Token721TransactorSession) SetAllowAffiliateProgram(_state bool) (*types.Transaction, error) {
	return _Token721.Contract.SetAllowAffiliateProgram(&_Token721.TransactOpts, _state)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Token721 *Token721Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Token721.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Token721 *Token721Session) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Token721.Contract.SetApprovalForAll(&_Token721.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Token721 *Token721TransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Token721.Contract.SetApprovalForAll(&_Token721.TransactOpts, operator, approved)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string _newBaseURI) returns()
func (_Token721 *Token721Transactor) SetBaseURI(opts *bind.TransactOpts, _newBaseURI string) (*types.Transaction, error) {
	return _Token721.contract.Transact(opts, "setBaseURI", _newBaseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string _newBaseURI) returns()
func (_Token721 *Token721Session) SetBaseURI(_newBaseURI string) (*types.Transaction, error) {
	return _Token721.Contract.SetBaseURI(&_Token721.TransactOpts, _newBaseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string _newBaseURI) returns()
func (_Token721 *Token721TransactorSession) SetBaseURI(_newBaseURI string) (*types.Transaction, error) {
	return _Token721.Contract.SetBaseURI(&_Token721.TransactOpts, _newBaseURI)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string contractURI_) returns()
func (_Token721 *Token721Transactor) SetContractURI(opts *bind.TransactOpts, contractURI_ string) (*types.Transaction, error) {
	return _Token721.contract.Transact(opts, "setContractURI", contractURI_)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string contractURI_) returns()
func (_Token721 *Token721Session) SetContractURI(contractURI_ string) (*types.Transaction, error) {
	return _Token721.Contract.SetContractURI(&_Token721.TransactOpts, contractURI_)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string contractURI_) returns()
func (_Token721 *Token721TransactorSession) SetContractURI(contractURI_ string) (*types.Transaction, error) {
	return _Token721.Contract.SetContractURI(&_Token721.TransactOpts, contractURI_)
}

// SetFreeMintPerAddressLimit is a paid mutator transaction binding the contract method 0x08181816.
//
// Solidity: function setFreeMintPerAddressLimit(uint256 _limit) returns()
func (_Token721 *Token721Transactor) SetFreeMintPerAddressLimit(opts *bind.TransactOpts, _limit *big.Int) (*types.Transaction, error) {
	return _Token721.contract.Transact(opts, "setFreeMintPerAddressLimit", _limit)
}

// SetFreeMintPerAddressLimit is a paid mutator transaction binding the contract method 0x08181816.
//
// Solidity: function setFreeMintPerAddressLimit(uint256 _limit) returns()
func (_Token721 *Token721Session) SetFreeMintPerAddressLimit(_limit *big.Int) (*types.Transaction, error) {
	return _Token721.Contract.SetFreeMintPerAddressLimit(&_Token721.TransactOpts, _limit)
}

// SetFreeMintPerAddressLimit is a paid mutator transaction binding the contract method 0x08181816.
//
// Solidity: function setFreeMintPerAddressLimit(uint256 _limit) returns()
func (_Token721 *Token721TransactorSession) SetFreeMintPerAddressLimit(_limit *big.Int) (*types.Transaction, error) {
	return _Token721.Contract.SetFreeMintPerAddressLimit(&_Token721.TransactOpts, _limit)
}

// SetIsFreeMint is a paid mutator transaction binding the contract method 0xef5d87d7.
//
// Solidity: function setIsFreeMint(bool state) returns()
func (_Token721 *Token721Transactor) SetIsFreeMint(opts *bind.TransactOpts, state bool) (*types.Transaction, error) {
	return _Token721.contract.Transact(opts, "setIsFreeMint", state)
}

// SetIsFreeMint is a paid mutator transaction binding the contract method 0xef5d87d7.
//
// Solidity: function setIsFreeMint(bool state) returns()
func (_Token721 *Token721Session) SetIsFreeMint(state bool) (*types.Transaction, error) {
	return _Token721.Contract.SetIsFreeMint(&_Token721.TransactOpts, state)
}

// SetIsFreeMint is a paid mutator transaction binding the contract method 0xef5d87d7.
//
// Solidity: function setIsFreeMint(bool state) returns()
func (_Token721 *Token721TransactorSession) SetIsFreeMint(state bool) (*types.Transaction, error) {
	return _Token721.Contract.SetIsFreeMint(&_Token721.TransactOpts, state)
}

// SetMaxMintAmount is a paid mutator transaction binding the contract method 0x088a4ed0.
//
// Solidity: function setMaxMintAmount(uint256 _maxMintAmount) returns()
func (_Token721 *Token721Transactor) SetMaxMintAmount(opts *bind.TransactOpts, _maxMintAmount *big.Int) (*types.Transaction, error) {
	return _Token721.contract.Transact(opts, "setMaxMintAmount", _maxMintAmount)
}

// SetMaxMintAmount is a paid mutator transaction binding the contract method 0x088a4ed0.
//
// Solidity: function setMaxMintAmount(uint256 _maxMintAmount) returns()
func (_Token721 *Token721Session) SetMaxMintAmount(_maxMintAmount *big.Int) (*types.Transaction, error) {
	return _Token721.Contract.SetMaxMintAmount(&_Token721.TransactOpts, _maxMintAmount)
}

// SetMaxMintAmount is a paid mutator transaction binding the contract method 0x088a4ed0.
//
// Solidity: function setMaxMintAmount(uint256 _maxMintAmount) returns()
func (_Token721 *Token721TransactorSession) SetMaxMintAmount(_maxMintAmount *big.Int) (*types.Transaction, error) {
	return _Token721.Contract.SetMaxMintAmount(&_Token721.TransactOpts, _maxMintAmount)
}

// SetMaxSupply is a paid mutator transaction binding the contract method 0x6f8b44b0.
//
// Solidity: function setMaxSupply(uint256 _maxSupply) returns()
func (_Token721 *Token721Transactor) SetMaxSupply(opts *bind.TransactOpts, _maxSupply *big.Int) (*types.Transaction, error) {
	return _Token721.contract.Transact(opts, "setMaxSupply", _maxSupply)
}

// SetMaxSupply is a paid mutator transaction binding the contract method 0x6f8b44b0.
//
// Solidity: function setMaxSupply(uint256 _maxSupply) returns()
func (_Token721 *Token721Session) SetMaxSupply(_maxSupply *big.Int) (*types.Transaction, error) {
	return _Token721.Contract.SetMaxSupply(&_Token721.TransactOpts, _maxSupply)
}

// SetMaxSupply is a paid mutator transaction binding the contract method 0x6f8b44b0.
//
// Solidity: function setMaxSupply(uint256 _maxSupply) returns()
func (_Token721 *Token721TransactorSession) SetMaxSupply(_maxSupply *big.Int) (*types.Transaction, error) {
	return _Token721.Contract.SetMaxSupply(&_Token721.TransactOpts, _maxSupply)
}

// SetNftEtherValue is a paid mutator transaction binding the contract method 0x6e17f797.
//
// Solidity: function setNftEtherValue(uint256 nftEtherValue) returns()
func (_Token721 *Token721Transactor) SetNftEtherValue(opts *bind.TransactOpts, nftEtherValue *big.Int) (*types.Transaction, error) {
	return _Token721.contract.Transact(opts, "setNftEtherValue", nftEtherValue)
}

// SetNftEtherValue is a paid mutator transaction binding the contract method 0x6e17f797.
//
// Solidity: function setNftEtherValue(uint256 nftEtherValue) returns()
func (_Token721 *Token721Session) SetNftEtherValue(nftEtherValue *big.Int) (*types.Transaction, error) {
	return _Token721.Contract.SetNftEtherValue(&_Token721.TransactOpts, nftEtherValue)
}

// SetNftEtherValue is a paid mutator transaction binding the contract method 0x6e17f797.
//
// Solidity: function setNftEtherValue(uint256 nftEtherValue) returns()
func (_Token721 *Token721TransactorSession) SetNftEtherValue(nftEtherValue *big.Int) (*types.Transaction, error) {
	return _Token721.Contract.SetNftEtherValue(&_Token721.TransactOpts, nftEtherValue)
}

// SetOnlyWhitelisted is a paid mutator transaction binding the contract method 0x3c952764.
//
// Solidity: function setOnlyWhitelisted(bool _state) returns()
func (_Token721 *Token721Transactor) SetOnlyWhitelisted(opts *bind.TransactOpts, _state bool) (*types.Transaction, error) {
	return _Token721.contract.Transact(opts, "setOnlyWhitelisted", _state)
}

// SetOnlyWhitelisted is a paid mutator transaction binding the contract method 0x3c952764.
//
// Solidity: function setOnlyWhitelisted(bool _state) returns()
func (_Token721 *Token721Session) SetOnlyWhitelisted(_state bool) (*types.Transaction, error) {
	return _Token721.Contract.SetOnlyWhitelisted(&_Token721.TransactOpts, _state)
}

// SetOnlyWhitelisted is a paid mutator transaction binding the contract method 0x3c952764.
//
// Solidity: function setOnlyWhitelisted(bool _state) returns()
func (_Token721 *Token721TransactorSession) SetOnlyWhitelisted(_state bool) (*types.Transaction, error) {
	return _Token721.Contract.SetOnlyWhitelisted(&_Token721.TransactOpts, _state)
}

// Split is a paid mutator transaction binding the contract method 0x8c584664.
//
// Solidity: function split(uint256 _mintAmount, address _recommendedBy, uint256 _indicationType, bytes32[] proof) payable returns()
func (_Token721 *Token721Transactor) Split(opts *bind.TransactOpts, _mintAmount *big.Int, _recommendedBy common.Address, _indicationType *big.Int, proof [][32]byte) (*types.Transaction, error) {
	return _Token721.contract.Transact(opts, "split", _mintAmount, _recommendedBy, _indicationType, proof)
}

// Split is a paid mutator transaction binding the contract method 0x8c584664.
//
// Solidity: function split(uint256 _mintAmount, address _recommendedBy, uint256 _indicationType, bytes32[] proof) payable returns()
func (_Token721 *Token721Session) Split(_mintAmount *big.Int, _recommendedBy common.Address, _indicationType *big.Int, proof [][32]byte) (*types.Transaction, error) {
	return _Token721.Contract.Split(&_Token721.TransactOpts, _mintAmount, _recommendedBy, _indicationType, proof)
}

// Split is a paid mutator transaction binding the contract method 0x8c584664.
//
// Solidity: function split(uint256 _mintAmount, address _recommendedBy, uint256 _indicationType, bytes32[] proof) payable returns()
func (_Token721 *Token721TransactorSession) Split(_mintAmount *big.Int, _recommendedBy common.Address, _indicationType *big.Int, proof [][32]byte) (*types.Transaction, error) {
	return _Token721.Contract.Split(&_Token721.TransactOpts, _mintAmount, _recommendedBy, _indicationType, proof)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Token721 *Token721Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Token721.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Token721 *Token721Session) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Token721.Contract.TransferFrom(&_Token721.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Token721 *Token721TransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Token721.Contract.TransferFrom(&_Token721.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Token721 *Token721Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Token721.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Token721 *Token721Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Token721.Contract.TransferOwnership(&_Token721.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Token721 *Token721TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Token721.Contract.TransferOwnership(&_Token721.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() payable returns()
func (_Token721 *Token721Transactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token721.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() payable returns()
func (_Token721 *Token721Session) Withdraw() (*types.Transaction, error) {
	return _Token721.Contract.Withdraw(&_Token721.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() payable returns()
func (_Token721 *Token721TransactorSession) Withdraw() (*types.Transaction, error) {
	return _Token721.Contract.Withdraw(&_Token721.TransactOpts)
}

// Token721ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Token721 contract.
type Token721ApprovalIterator struct {
	Event *Token721Approval // Event containing the contract specifics and raw log

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
func (it *Token721ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Token721Approval)
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
		it.Event = new(Token721Approval)
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
func (it *Token721ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Token721ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Token721Approval represents a Approval event raised by the Token721 contract.
type Token721Approval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Token721 *Token721Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*Token721ApprovalIterator, error) {

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

	logs, sub, err := _Token721.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &Token721ApprovalIterator{contract: _Token721.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Token721 *Token721Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Token721Approval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Token721.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Token721Approval)
				if err := _Token721.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Token721 *Token721Filterer) ParseApproval(log types.Log) (*Token721Approval, error) {
	event := new(Token721Approval)
	if err := _Token721.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Token721ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Token721 contract.
type Token721ApprovalForAllIterator struct {
	Event *Token721ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *Token721ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Token721ApprovalForAll)
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
		it.Event = new(Token721ApprovalForAll)
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
func (it *Token721ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Token721ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Token721ApprovalForAll represents a ApprovalForAll event raised by the Token721 contract.
type Token721ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Token721 *Token721Filterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*Token721ApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Token721.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &Token721ApprovalForAllIterator{contract: _Token721.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Token721 *Token721Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *Token721ApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Token721.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Token721ApprovalForAll)
				if err := _Token721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_Token721 *Token721Filterer) ParseApprovalForAll(log types.Log) (*Token721ApprovalForAll, error) {
	event := new(Token721ApprovalForAll)
	if err := _Token721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Token721OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Token721 contract.
type Token721OwnershipTransferredIterator struct {
	Event *Token721OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *Token721OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Token721OwnershipTransferred)
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
		it.Event = new(Token721OwnershipTransferred)
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
func (it *Token721OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Token721OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Token721OwnershipTransferred represents a OwnershipTransferred event raised by the Token721 contract.
type Token721OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Token721 *Token721Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*Token721OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Token721.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &Token721OwnershipTransferredIterator{contract: _Token721.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Token721 *Token721Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *Token721OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Token721.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Token721OwnershipTransferred)
				if err := _Token721.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Token721 *Token721Filterer) ParseOwnershipTransferred(log types.Log) (*Token721OwnershipTransferred, error) {
	event := new(Token721OwnershipTransferred)
	if err := _Token721.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Token721TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Token721 contract.
type Token721TransferIterator struct {
	Event *Token721Transfer // Event containing the contract specifics and raw log

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
func (it *Token721TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Token721Transfer)
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
		it.Event = new(Token721Transfer)
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
func (it *Token721TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Token721TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Token721Transfer represents a Transfer event raised by the Token721 contract.
type Token721Transfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Token721 *Token721Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*Token721TransferIterator, error) {

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

	logs, sub, err := _Token721.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &Token721TransferIterator{contract: _Token721.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Token721 *Token721Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Token721Transfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Token721.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Token721Transfer)
				if err := _Token721.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Token721 *Token721Filterer) ParseTransfer(log types.Log) (*Token721Transfer, error) {
	event := new(Token721Transfer)
	if err := _Token721.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Token721TransferSendIterator is returned from FilterTransferSend and is used to iterate over the raw logs and unpacked data for TransferSend events raised by the Token721 contract.
type Token721TransferSendIterator struct {
	Event *Token721TransferSend // Event containing the contract specifics and raw log

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
func (it *Token721TransferSendIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Token721TransferSend)
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
		it.Event = new(Token721TransferSend)
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
func (it *Token721TransferSendIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Token721TransferSendIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Token721TransferSend represents a TransferSend event raised by the Token721 contract.
type Token721TransferSend struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransferSend is a free log retrieval operation binding the contract event 0xf70a742b6ec5de639ce8e8d0df10a44a88c323ac660a09f1fc0c190e830aea74.
//
// Solidity: event _transferSend(address _from, address _to, uint256 _amount)
func (_Token721 *Token721Filterer) FilterTransferSend(opts *bind.FilterOpts) (*Token721TransferSendIterator, error) {

	logs, sub, err := _Token721.contract.FilterLogs(opts, "_transferSend")
	if err != nil {
		return nil, err
	}
	return &Token721TransferSendIterator{contract: _Token721.contract, event: "_transferSend", logs: logs, sub: sub}, nil
}

// WatchTransferSend is a free log subscription operation binding the contract event 0xf70a742b6ec5de639ce8e8d0df10a44a88c323ac660a09f1fc0c190e830aea74.
//
// Solidity: event _transferSend(address _from, address _to, uint256 _amount)
func (_Token721 *Token721Filterer) WatchTransferSend(opts *bind.WatchOpts, sink chan<- *Token721TransferSend) (event.Subscription, error) {

	logs, sub, err := _Token721.contract.WatchLogs(opts, "_transferSend")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Token721TransferSend)
				if err := _Token721.contract.UnpackLog(event, "_transferSend", log); err != nil {
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
func (_Token721 *Token721Filterer) ParseTransferSend(log types.Log) (*Token721TransferSend, error) {
	event := new(Token721TransferSend)
	if err := _Token721.contract.UnpackLog(event, "_transferSend", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
