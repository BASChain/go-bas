// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Contract

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// BasOwnershipABI is the input ABI used to generate the binding from.
const BasOwnershipABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"check\",\"type\":\"address\"}],\"name\":\"isValid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"_a_changeAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"}],\"name\":\"expired\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expire\",\"type\":\"uint256\"}],\"name\":\"_c_add\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contractCaller\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expire\",\"type\":\"uint256\"}],\"name\":\"_c_takeover\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"conAddr\",\"type\":\"address\"}],\"name\":\"_a_changeContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"}],\"name\":\"exist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expire\",\"type\":\"uint256\"}],\"name\":\"_a_update\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"extend\",\"type\":\"uint256\"}],\"name\":\"_c_extend\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"}],\"name\":\"removeExpired\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"}],\"name\":\"remove\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"}],\"name\":\"query\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"}],\"name\":\"isWild\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"}],\"name\":\"revoke\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"expirationOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"Add\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"Update\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"Extend\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Takeover\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"by\",\"type\":\"address\"}],\"name\":\"TransferFrom\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"}],\"name\":\"Remove\",\"type\":\"event\"}]"

// BasOwnership is an auto generated Go binding around an Ethereum contract.
type BasOwnership struct {
	BasOwnershipCaller     // Read-only binding to the contract
	BasOwnershipTransactor // Write-only binding to the contract
	BasOwnershipFilterer   // Log filterer for contract events
}

// BasOwnershipCaller is an auto generated read-only Go binding around an Ethereum contract.
type BasOwnershipCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasOwnershipTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BasOwnershipTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasOwnershipFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BasOwnershipFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasOwnershipSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BasOwnershipSession struct {
	Contract     *BasOwnership     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BasOwnershipCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BasOwnershipCallerSession struct {
	Contract *BasOwnershipCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// BasOwnershipTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BasOwnershipTransactorSession struct {
	Contract     *BasOwnershipTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// BasOwnershipRaw is an auto generated low-level Go binding around an Ethereum contract.
type BasOwnershipRaw struct {
	Contract *BasOwnership // Generic contract binding to access the raw methods on
}

// BasOwnershipCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BasOwnershipCallerRaw struct {
	Contract *BasOwnershipCaller // Generic read-only contract binding to access the raw methods on
}

// BasOwnershipTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BasOwnershipTransactorRaw struct {
	Contract *BasOwnershipTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBasOwnership creates a new instance of BasOwnership, bound to a specific deployed contract.
func NewBasOwnership(address common.Address, backend bind.ContractBackend) (*BasOwnership, error) {
	contract, err := bindBasOwnership(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BasOwnership{BasOwnershipCaller: BasOwnershipCaller{contract: contract}, BasOwnershipTransactor: BasOwnershipTransactor{contract: contract}, BasOwnershipFilterer: BasOwnershipFilterer{contract: contract}}, nil
}

// NewBasOwnershipCaller creates a new read-only instance of BasOwnership, bound to a specific deployed contract.
func NewBasOwnershipCaller(address common.Address, caller bind.ContractCaller) (*BasOwnershipCaller, error) {
	contract, err := bindBasOwnership(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BasOwnershipCaller{contract: contract}, nil
}

// NewBasOwnershipTransactor creates a new write-only instance of BasOwnership, bound to a specific deployed contract.
func NewBasOwnershipTransactor(address common.Address, transactor bind.ContractTransactor) (*BasOwnershipTransactor, error) {
	contract, err := bindBasOwnership(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BasOwnershipTransactor{contract: contract}, nil
}

// NewBasOwnershipFilterer creates a new log filterer instance of BasOwnership, bound to a specific deployed contract.
func NewBasOwnershipFilterer(address common.Address, filterer bind.ContractFilterer) (*BasOwnershipFilterer, error) {
	contract, err := bindBasOwnership(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BasOwnershipFilterer{contract: contract}, nil
}

// bindBasOwnership binds a generic wrapper to an already deployed contract.
func bindBasOwnership(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BasOwnershipABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasOwnership *BasOwnershipRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BasOwnership.Contract.BasOwnershipCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasOwnership *BasOwnershipRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasOwnership.Contract.BasOwnershipTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasOwnership *BasOwnershipRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasOwnership.Contract.BasOwnershipTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasOwnership *BasOwnershipCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BasOwnership.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasOwnership *BasOwnershipTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasOwnership.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasOwnership *BasOwnershipTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasOwnership.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_BasOwnership *BasOwnershipCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BasOwnership.contract.Call(opts, out, "admin")
	return *ret0, err
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_BasOwnership *BasOwnershipSession) Admin() (common.Address, error) {
	return _BasOwnership.Contract.Admin(&_BasOwnership.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_BasOwnership *BasOwnershipCallerSession) Admin() (common.Address, error) {
	return _BasOwnership.Contract.Admin(&_BasOwnership.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xbcf9fcad.
//
// Solidity: function allowance(address owner, bytes32 nameHash) constant returns(address)
func (_BasOwnership *BasOwnershipCaller) Allowance(opts *bind.CallOpts, owner common.Address, nameHash [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BasOwnership.contract.Call(opts, out, "allowance", owner, nameHash)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xbcf9fcad.
//
// Solidity: function allowance(address owner, bytes32 nameHash) constant returns(address)
func (_BasOwnership *BasOwnershipSession) Allowance(owner common.Address, nameHash [32]byte) (common.Address, error) {
	return _BasOwnership.Contract.Allowance(&_BasOwnership.CallOpts, owner, nameHash)
}

// Allowance is a free data retrieval call binding the contract method 0xbcf9fcad.
//
// Solidity: function allowance(address owner, bytes32 nameHash) constant returns(address)
func (_BasOwnership *BasOwnershipCallerSession) Allowance(owner common.Address, nameHash [32]byte) (common.Address, error) {
	return _BasOwnership.Contract.Allowance(&_BasOwnership.CallOpts, owner, nameHash)
}

// ContractCaller is a free data retrieval call binding the contract method 0x54cf428a.
//
// Solidity: function contractCaller() constant returns(address)
func (_BasOwnership *BasOwnershipCaller) ContractCaller(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BasOwnership.contract.Call(opts, out, "contractCaller")
	return *ret0, err
}

// ContractCaller is a free data retrieval call binding the contract method 0x54cf428a.
//
// Solidity: function contractCaller() constant returns(address)
func (_BasOwnership *BasOwnershipSession) ContractCaller() (common.Address, error) {
	return _BasOwnership.Contract.ContractCaller(&_BasOwnership.CallOpts)
}

// ContractCaller is a free data retrieval call binding the contract method 0x54cf428a.
//
// Solidity: function contractCaller() constant returns(address)
func (_BasOwnership *BasOwnershipCallerSession) ContractCaller() (common.Address, error) {
	return _BasOwnership.Contract.ContractCaller(&_BasOwnership.CallOpts)
}

// Exist is a free data retrieval call binding the contract method 0x73e8b3d4.
//
// Solidity: function exist(bytes32 nameHash) constant returns(bool)
func (_BasOwnership *BasOwnershipCaller) Exist(opts *bind.CallOpts, nameHash [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BasOwnership.contract.Call(opts, out, "exist", nameHash)
	return *ret0, err
}

// Exist is a free data retrieval call binding the contract method 0x73e8b3d4.
//
// Solidity: function exist(bytes32 nameHash) constant returns(bool)
func (_BasOwnership *BasOwnershipSession) Exist(nameHash [32]byte) (bool, error) {
	return _BasOwnership.Contract.Exist(&_BasOwnership.CallOpts, nameHash)
}

// Exist is a free data retrieval call binding the contract method 0x73e8b3d4.
//
// Solidity: function exist(bytes32 nameHash) constant returns(bool)
func (_BasOwnership *BasOwnershipCallerSession) Exist(nameHash [32]byte) (bool, error) {
	return _BasOwnership.Contract.Exist(&_BasOwnership.CallOpts, nameHash)
}

// ExpirationOf is a free data retrieval call binding the contract method 0xf7f67ef9.
//
// Solidity: function expirationOf(bytes32 ) constant returns(uint256)
func (_BasOwnership *BasOwnershipCaller) ExpirationOf(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BasOwnership.contract.Call(opts, out, "expirationOf", arg0)
	return *ret0, err
}

// ExpirationOf is a free data retrieval call binding the contract method 0xf7f67ef9.
//
// Solidity: function expirationOf(bytes32 ) constant returns(uint256)
func (_BasOwnership *BasOwnershipSession) ExpirationOf(arg0 [32]byte) (*big.Int, error) {
	return _BasOwnership.Contract.ExpirationOf(&_BasOwnership.CallOpts, arg0)
}

// ExpirationOf is a free data retrieval call binding the contract method 0xf7f67ef9.
//
// Solidity: function expirationOf(bytes32 ) constant returns(uint256)
func (_BasOwnership *BasOwnershipCallerSession) ExpirationOf(arg0 [32]byte) (*big.Int, error) {
	return _BasOwnership.Contract.ExpirationOf(&_BasOwnership.CallOpts, arg0)
}

// Expired is a free data retrieval call binding the contract method 0x30cc3597.
//
// Solidity: function expired(bytes32 nameHash) constant returns(bool)
func (_BasOwnership *BasOwnershipCaller) Expired(opts *bind.CallOpts, nameHash [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BasOwnership.contract.Call(opts, out, "expired", nameHash)
	return *ret0, err
}

// Expired is a free data retrieval call binding the contract method 0x30cc3597.
//
// Solidity: function expired(bytes32 nameHash) constant returns(bool)
func (_BasOwnership *BasOwnershipSession) Expired(nameHash [32]byte) (bool, error) {
	return _BasOwnership.Contract.Expired(&_BasOwnership.CallOpts, nameHash)
}

// Expired is a free data retrieval call binding the contract method 0x30cc3597.
//
// Solidity: function expired(bytes32 nameHash) constant returns(bool)
func (_BasOwnership *BasOwnershipCallerSession) Expired(nameHash [32]byte) (bool, error) {
	return _BasOwnership.Contract.Expired(&_BasOwnership.CallOpts, nameHash)
}

// IsValid is a free data retrieval call binding the contract method 0x0407abec.
//
// Solidity: function isValid(bytes32 nameHash, address check) constant returns(bool)
func (_BasOwnership *BasOwnershipCaller) IsValid(opts *bind.CallOpts, nameHash [32]byte, check common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BasOwnership.contract.Call(opts, out, "isValid", nameHash, check)
	return *ret0, err
}

// IsValid is a free data retrieval call binding the contract method 0x0407abec.
//
// Solidity: function isValid(bytes32 nameHash, address check) constant returns(bool)
func (_BasOwnership *BasOwnershipSession) IsValid(nameHash [32]byte, check common.Address) (bool, error) {
	return _BasOwnership.Contract.IsValid(&_BasOwnership.CallOpts, nameHash, check)
}

// IsValid is a free data retrieval call binding the contract method 0x0407abec.
//
// Solidity: function isValid(bytes32 nameHash, address check) constant returns(bool)
func (_BasOwnership *BasOwnershipCallerSession) IsValid(nameHash [32]byte, check common.Address) (bool, error) {
	return _BasOwnership.Contract.IsValid(&_BasOwnership.CallOpts, nameHash, check)
}

// IsWild is a free data retrieval call binding the contract method 0xa87587f6.
//
// Solidity: function isWild(bytes32 nameHash) constant returns(bool)
func (_BasOwnership *BasOwnershipCaller) IsWild(opts *bind.CallOpts, nameHash [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BasOwnership.contract.Call(opts, out, "isWild", nameHash)
	return *ret0, err
}

// IsWild is a free data retrieval call binding the contract method 0xa87587f6.
//
// Solidity: function isWild(bytes32 nameHash) constant returns(bool)
func (_BasOwnership *BasOwnershipSession) IsWild(nameHash [32]byte) (bool, error) {
	return _BasOwnership.Contract.IsWild(&_BasOwnership.CallOpts, nameHash)
}

// IsWild is a free data retrieval call binding the contract method 0xa87587f6.
//
// Solidity: function isWild(bytes32 nameHash) constant returns(bool)
func (_BasOwnership *BasOwnershipCallerSession) IsWild(nameHash [32]byte) (bool, error) {
	return _BasOwnership.Contract.IsWild(&_BasOwnership.CallOpts, nameHash)
}

// OwnerOf is a free data retrieval call binding the contract method 0x7dd56411.
//
// Solidity: function ownerOf(bytes32 ) constant returns(address)
func (_BasOwnership *BasOwnershipCaller) OwnerOf(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BasOwnership.contract.Call(opts, out, "ownerOf", arg0)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x7dd56411.
//
// Solidity: function ownerOf(bytes32 ) constant returns(address)
func (_BasOwnership *BasOwnershipSession) OwnerOf(arg0 [32]byte) (common.Address, error) {
	return _BasOwnership.Contract.OwnerOf(&_BasOwnership.CallOpts, arg0)
}

// OwnerOf is a free data retrieval call binding the contract method 0x7dd56411.
//
// Solidity: function ownerOf(bytes32 ) constant returns(address)
func (_BasOwnership *BasOwnershipCallerSession) OwnerOf(arg0 [32]byte) (common.Address, error) {
	return _BasOwnership.Contract.OwnerOf(&_BasOwnership.CallOpts, arg0)
}

// Query is a free data retrieval call binding the contract method 0xa3b61a9c.
//
// Solidity: function query(bytes32 nameHash) constant returns(address, uint256)
func (_BasOwnership *BasOwnershipCaller) Query(opts *bind.CallOpts, nameHash [32]byte) (common.Address, *big.Int, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _BasOwnership.contract.Call(opts, out, "query", nameHash)
	return *ret0, *ret1, err
}

// Query is a free data retrieval call binding the contract method 0xa3b61a9c.
//
// Solidity: function query(bytes32 nameHash) constant returns(address, uint256)
func (_BasOwnership *BasOwnershipSession) Query(nameHash [32]byte) (common.Address, *big.Int, error) {
	return _BasOwnership.Contract.Query(&_BasOwnership.CallOpts, nameHash)
}

// Query is a free data retrieval call binding the contract method 0xa3b61a9c.
//
// Solidity: function query(bytes32 nameHash) constant returns(address, uint256)
func (_BasOwnership *BasOwnershipCallerSession) Query(nameHash [32]byte) (common.Address, *big.Int, error) {
	return _BasOwnership.Contract.Query(&_BasOwnership.CallOpts, nameHash)
}

// AChangeAdmin is a paid mutator transaction binding the contract method 0x1062f6cc.
//
// Solidity: function _a_changeAdmin(address newOwner) returns()
func (_BasOwnership *BasOwnershipTransactor) AChangeAdmin(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BasOwnership.contract.Transact(opts, "_a_changeAdmin", newOwner)
}

// AChangeAdmin is a paid mutator transaction binding the contract method 0x1062f6cc.
//
// Solidity: function _a_changeAdmin(address newOwner) returns()
func (_BasOwnership *BasOwnershipSession) AChangeAdmin(newOwner common.Address) (*types.Transaction, error) {
	return _BasOwnership.Contract.AChangeAdmin(&_BasOwnership.TransactOpts, newOwner)
}

// AChangeAdmin is a paid mutator transaction binding the contract method 0x1062f6cc.
//
// Solidity: function _a_changeAdmin(address newOwner) returns()
func (_BasOwnership *BasOwnershipTransactorSession) AChangeAdmin(newOwner common.Address) (*types.Transaction, error) {
	return _BasOwnership.Contract.AChangeAdmin(&_BasOwnership.TransactOpts, newOwner)
}

// AChangeContract is a paid mutator transaction binding the contract method 0x733c1950.
//
// Solidity: function _a_changeContract(address conAddr) returns()
func (_BasOwnership *BasOwnershipTransactor) AChangeContract(opts *bind.TransactOpts, conAddr common.Address) (*types.Transaction, error) {
	return _BasOwnership.contract.Transact(opts, "_a_changeContract", conAddr)
}

// AChangeContract is a paid mutator transaction binding the contract method 0x733c1950.
//
// Solidity: function _a_changeContract(address conAddr) returns()
func (_BasOwnership *BasOwnershipSession) AChangeContract(conAddr common.Address) (*types.Transaction, error) {
	return _BasOwnership.Contract.AChangeContract(&_BasOwnership.TransactOpts, conAddr)
}

// AChangeContract is a paid mutator transaction binding the contract method 0x733c1950.
//
// Solidity: function _a_changeContract(address conAddr) returns()
func (_BasOwnership *BasOwnershipTransactorSession) AChangeContract(conAddr common.Address) (*types.Transaction, error) {
	return _BasOwnership.Contract.AChangeContract(&_BasOwnership.TransactOpts, conAddr)
}

// AUpdate is a paid mutator transaction binding the contract method 0x788ccdea.
//
// Solidity: function _a_update(bytes32 nameHash, address owner, uint256 expire) returns()
func (_BasOwnership *BasOwnershipTransactor) AUpdate(opts *bind.TransactOpts, nameHash [32]byte, owner common.Address, expire *big.Int) (*types.Transaction, error) {
	return _BasOwnership.contract.Transact(opts, "_a_update", nameHash, owner, expire)
}

// AUpdate is a paid mutator transaction binding the contract method 0x788ccdea.
//
// Solidity: function _a_update(bytes32 nameHash, address owner, uint256 expire) returns()
func (_BasOwnership *BasOwnershipSession) AUpdate(nameHash [32]byte, owner common.Address, expire *big.Int) (*types.Transaction, error) {
	return _BasOwnership.Contract.AUpdate(&_BasOwnership.TransactOpts, nameHash, owner, expire)
}

// AUpdate is a paid mutator transaction binding the contract method 0x788ccdea.
//
// Solidity: function _a_update(bytes32 nameHash, address owner, uint256 expire) returns()
func (_BasOwnership *BasOwnershipTransactorSession) AUpdate(nameHash [32]byte, owner common.Address, expire *big.Int) (*types.Transaction, error) {
	return _BasOwnership.Contract.AUpdate(&_BasOwnership.TransactOpts, nameHash, owner, expire)
}

// CAdd is a paid mutator transaction binding the contract method 0x33c7fb4c.
//
// Solidity: function _c_add(bytes32 nameHash, address owner, uint256 expire) returns()
func (_BasOwnership *BasOwnershipTransactor) CAdd(opts *bind.TransactOpts, nameHash [32]byte, owner common.Address, expire *big.Int) (*types.Transaction, error) {
	return _BasOwnership.contract.Transact(opts, "_c_add", nameHash, owner, expire)
}

// CAdd is a paid mutator transaction binding the contract method 0x33c7fb4c.
//
// Solidity: function _c_add(bytes32 nameHash, address owner, uint256 expire) returns()
func (_BasOwnership *BasOwnershipSession) CAdd(nameHash [32]byte, owner common.Address, expire *big.Int) (*types.Transaction, error) {
	return _BasOwnership.Contract.CAdd(&_BasOwnership.TransactOpts, nameHash, owner, expire)
}

// CAdd is a paid mutator transaction binding the contract method 0x33c7fb4c.
//
// Solidity: function _c_add(bytes32 nameHash, address owner, uint256 expire) returns()
func (_BasOwnership *BasOwnershipTransactorSession) CAdd(nameHash [32]byte, owner common.Address, expire *big.Int) (*types.Transaction, error) {
	return _BasOwnership.Contract.CAdd(&_BasOwnership.TransactOpts, nameHash, owner, expire)
}

// CExtend is a paid mutator transaction binding the contract method 0x7f77bbcd.
//
// Solidity: function _c_extend(bytes32 nameHash, uint256 extend) returns()
func (_BasOwnership *BasOwnershipTransactor) CExtend(opts *bind.TransactOpts, nameHash [32]byte, extend *big.Int) (*types.Transaction, error) {
	return _BasOwnership.contract.Transact(opts, "_c_extend", nameHash, extend)
}

// CExtend is a paid mutator transaction binding the contract method 0x7f77bbcd.
//
// Solidity: function _c_extend(bytes32 nameHash, uint256 extend) returns()
func (_BasOwnership *BasOwnershipSession) CExtend(nameHash [32]byte, extend *big.Int) (*types.Transaction, error) {
	return _BasOwnership.Contract.CExtend(&_BasOwnership.TransactOpts, nameHash, extend)
}

// CExtend is a paid mutator transaction binding the contract method 0x7f77bbcd.
//
// Solidity: function _c_extend(bytes32 nameHash, uint256 extend) returns()
func (_BasOwnership *BasOwnershipTransactorSession) CExtend(nameHash [32]byte, extend *big.Int) (*types.Transaction, error) {
	return _BasOwnership.Contract.CExtend(&_BasOwnership.TransactOpts, nameHash, extend)
}

// CTakeover is a paid mutator transaction binding the contract method 0x56b7f1a0.
//
// Solidity: function _c_takeover(bytes32 nameHash, address owner, uint256 expire) returns()
func (_BasOwnership *BasOwnershipTransactor) CTakeover(opts *bind.TransactOpts, nameHash [32]byte, owner common.Address, expire *big.Int) (*types.Transaction, error) {
	return _BasOwnership.contract.Transact(opts, "_c_takeover", nameHash, owner, expire)
}

// CTakeover is a paid mutator transaction binding the contract method 0x56b7f1a0.
//
// Solidity: function _c_takeover(bytes32 nameHash, address owner, uint256 expire) returns()
func (_BasOwnership *BasOwnershipSession) CTakeover(nameHash [32]byte, owner common.Address, expire *big.Int) (*types.Transaction, error) {
	return _BasOwnership.Contract.CTakeover(&_BasOwnership.TransactOpts, nameHash, owner, expire)
}

// CTakeover is a paid mutator transaction binding the contract method 0x56b7f1a0.
//
// Solidity: function _c_takeover(bytes32 nameHash, address owner, uint256 expire) returns()
func (_BasOwnership *BasOwnershipTransactorSession) CTakeover(nameHash [32]byte, owner common.Address, expire *big.Int) (*types.Transaction, error) {
	return _BasOwnership.Contract.CTakeover(&_BasOwnership.TransactOpts, nameHash, owner, expire)
}

// Approve is a paid mutator transaction binding the contract method 0xb2e2c99b.
//
// Solidity: function approve(bytes32 nameHash, address spender) returns()
func (_BasOwnership *BasOwnershipTransactor) Approve(opts *bind.TransactOpts, nameHash [32]byte, spender common.Address) (*types.Transaction, error) {
	return _BasOwnership.contract.Transact(opts, "approve", nameHash, spender)
}

// Approve is a paid mutator transaction binding the contract method 0xb2e2c99b.
//
// Solidity: function approve(bytes32 nameHash, address spender) returns()
func (_BasOwnership *BasOwnershipSession) Approve(nameHash [32]byte, spender common.Address) (*types.Transaction, error) {
	return _BasOwnership.Contract.Approve(&_BasOwnership.TransactOpts, nameHash, spender)
}

// Approve is a paid mutator transaction binding the contract method 0xb2e2c99b.
//
// Solidity: function approve(bytes32 nameHash, address spender) returns()
func (_BasOwnership *BasOwnershipTransactorSession) Approve(nameHash [32]byte, spender common.Address) (*types.Transaction, error) {
	return _BasOwnership.Contract.Approve(&_BasOwnership.TransactOpts, nameHash, spender)
}

// Remove is a paid mutator transaction binding the contract method 0x95bc2673.
//
// Solidity: function remove(bytes32 nameHash) returns()
func (_BasOwnership *BasOwnershipTransactor) Remove(opts *bind.TransactOpts, nameHash [32]byte) (*types.Transaction, error) {
	return _BasOwnership.contract.Transact(opts, "remove", nameHash)
}

// Remove is a paid mutator transaction binding the contract method 0x95bc2673.
//
// Solidity: function remove(bytes32 nameHash) returns()
func (_BasOwnership *BasOwnershipSession) Remove(nameHash [32]byte) (*types.Transaction, error) {
	return _BasOwnership.Contract.Remove(&_BasOwnership.TransactOpts, nameHash)
}

// Remove is a paid mutator transaction binding the contract method 0x95bc2673.
//
// Solidity: function remove(bytes32 nameHash) returns()
func (_BasOwnership *BasOwnershipTransactorSession) Remove(nameHash [32]byte) (*types.Transaction, error) {
	return _BasOwnership.Contract.Remove(&_BasOwnership.TransactOpts, nameHash)
}

// RemoveExpired is a paid mutator transaction binding the contract method 0x8380d30c.
//
// Solidity: function removeExpired(bytes32 nameHash) returns()
func (_BasOwnership *BasOwnershipTransactor) RemoveExpired(opts *bind.TransactOpts, nameHash [32]byte) (*types.Transaction, error) {
	return _BasOwnership.contract.Transact(opts, "removeExpired", nameHash)
}

// RemoveExpired is a paid mutator transaction binding the contract method 0x8380d30c.
//
// Solidity: function removeExpired(bytes32 nameHash) returns()
func (_BasOwnership *BasOwnershipSession) RemoveExpired(nameHash [32]byte) (*types.Transaction, error) {
	return _BasOwnership.Contract.RemoveExpired(&_BasOwnership.TransactOpts, nameHash)
}

// RemoveExpired is a paid mutator transaction binding the contract method 0x8380d30c.
//
// Solidity: function removeExpired(bytes32 nameHash) returns()
func (_BasOwnership *BasOwnershipTransactorSession) RemoveExpired(nameHash [32]byte) (*types.Transaction, error) {
	return _BasOwnership.Contract.RemoveExpired(&_BasOwnership.TransactOpts, nameHash)
}

// Revoke is a paid mutator transaction binding the contract method 0xb75c7dc6.
//
// Solidity: function revoke(bytes32 nameHash) returns()
func (_BasOwnership *BasOwnershipTransactor) Revoke(opts *bind.TransactOpts, nameHash [32]byte) (*types.Transaction, error) {
	return _BasOwnership.contract.Transact(opts, "revoke", nameHash)
}

// Revoke is a paid mutator transaction binding the contract method 0xb75c7dc6.
//
// Solidity: function revoke(bytes32 nameHash) returns()
func (_BasOwnership *BasOwnershipSession) Revoke(nameHash [32]byte) (*types.Transaction, error) {
	return _BasOwnership.Contract.Revoke(&_BasOwnership.TransactOpts, nameHash)
}

// Revoke is a paid mutator transaction binding the contract method 0xb75c7dc6.
//
// Solidity: function revoke(bytes32 nameHash) returns()
func (_BasOwnership *BasOwnershipTransactorSession) Revoke(nameHash [32]byte) (*types.Transaction, error) {
	return _BasOwnership.Contract.Revoke(&_BasOwnership.TransactOpts, nameHash)
}

// Transfer is a paid mutator transaction binding the contract method 0x79ce9fac.
//
// Solidity: function transfer(bytes32 nameHash, address to) returns()
func (_BasOwnership *BasOwnershipTransactor) Transfer(opts *bind.TransactOpts, nameHash [32]byte, to common.Address) (*types.Transaction, error) {
	return _BasOwnership.contract.Transact(opts, "transfer", nameHash, to)
}

// Transfer is a paid mutator transaction binding the contract method 0x79ce9fac.
//
// Solidity: function transfer(bytes32 nameHash, address to) returns()
func (_BasOwnership *BasOwnershipSession) Transfer(nameHash [32]byte, to common.Address) (*types.Transaction, error) {
	return _BasOwnership.Contract.Transfer(&_BasOwnership.TransactOpts, nameHash, to)
}

// Transfer is a paid mutator transaction binding the contract method 0x79ce9fac.
//
// Solidity: function transfer(bytes32 nameHash, address to) returns()
func (_BasOwnership *BasOwnershipTransactorSession) Transfer(nameHash [32]byte, to common.Address) (*types.Transaction, error) {
	return _BasOwnership.Contract.Transfer(&_BasOwnership.TransactOpts, nameHash, to)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x82d7bd27.
//
// Solidity: function transferFrom(bytes32 nameHash, address from, address to) returns()
func (_BasOwnership *BasOwnershipTransactor) TransferFrom(opts *bind.TransactOpts, nameHash [32]byte, from common.Address, to common.Address) (*types.Transaction, error) {
	return _BasOwnership.contract.Transact(opts, "transferFrom", nameHash, from, to)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x82d7bd27.
//
// Solidity: function transferFrom(bytes32 nameHash, address from, address to) returns()
func (_BasOwnership *BasOwnershipSession) TransferFrom(nameHash [32]byte, from common.Address, to common.Address) (*types.Transaction, error) {
	return _BasOwnership.Contract.TransferFrom(&_BasOwnership.TransactOpts, nameHash, from, to)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x82d7bd27.
//
// Solidity: function transferFrom(bytes32 nameHash, address from, address to) returns()
func (_BasOwnership *BasOwnershipTransactorSession) TransferFrom(nameHash [32]byte, from common.Address, to common.Address) (*types.Transaction, error) {
	return _BasOwnership.Contract.TransferFrom(&_BasOwnership.TransactOpts, nameHash, from, to)
}

// BasOwnershipAddIterator is returned from FilterAdd and is used to iterate over the raw logs and unpacked data for Add events raised by the BasOwnership contract.
type BasOwnershipAddIterator struct {
	Event *BasOwnershipAdd // Event containing the contract specifics and raw log

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
func (it *BasOwnershipAddIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasOwnershipAdd)
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
		it.Event = new(BasOwnershipAdd)
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
func (it *BasOwnershipAddIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasOwnershipAddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasOwnershipAdd represents a Add event raised by the BasOwnership contract.
type BasOwnershipAdd struct {
	NameHash [32]byte
	Owner    common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAdd is a free log retrieval operation binding the contract event 0x5f17f6377b8024fbe396282bbccde46c3c704613385bd2f4b4b3ff955fd56a35.
//
// Solidity: event Add(bytes32 nameHash, address owner)
func (_BasOwnership *BasOwnershipFilterer) FilterAdd(opts *bind.FilterOpts) (*BasOwnershipAddIterator, error) {

	logs, sub, err := _BasOwnership.contract.FilterLogs(opts, "Add")
	if err != nil {
		return nil, err
	}
	return &BasOwnershipAddIterator{contract: _BasOwnership.contract, event: "Add", logs: logs, sub: sub}, nil
}

// WatchAdd is a free log subscription operation binding the contract event 0x5f17f6377b8024fbe396282bbccde46c3c704613385bd2f4b4b3ff955fd56a35.
//
// Solidity: event Add(bytes32 nameHash, address owner)
func (_BasOwnership *BasOwnershipFilterer) WatchAdd(opts *bind.WatchOpts, sink chan<- *BasOwnershipAdd) (event.Subscription, error) {

	logs, sub, err := _BasOwnership.contract.WatchLogs(opts, "Add")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasOwnershipAdd)
				if err := _BasOwnership.contract.UnpackLog(event, "Add", log); err != nil {
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

// ParseAdd is a log parse operation binding the contract event 0x5f17f6377b8024fbe396282bbccde46c3c704613385bd2f4b4b3ff955fd56a35.
//
// Solidity: event Add(bytes32 nameHash, address owner)
func (_BasOwnership *BasOwnershipFilterer) ParseAdd(log types.Log) (*BasOwnershipAdd, error) {
	event := new(BasOwnershipAdd)
	if err := _BasOwnership.contract.UnpackLog(event, "Add", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BasOwnershipExtendIterator is returned from FilterExtend and is used to iterate over the raw logs and unpacked data for Extend events raised by the BasOwnership contract.
type BasOwnershipExtendIterator struct {
	Event *BasOwnershipExtend // Event containing the contract specifics and raw log

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
func (it *BasOwnershipExtendIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasOwnershipExtend)
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
		it.Event = new(BasOwnershipExtend)
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
func (it *BasOwnershipExtendIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasOwnershipExtendIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasOwnershipExtend represents a Extend event raised by the BasOwnership contract.
type BasOwnershipExtend struct {
	NameHash [32]byte
	Time     *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterExtend is a free log retrieval operation binding the contract event 0xe898863b26adf3af54e82709ac1a76136eaaa5f3a7d9b790ce6d539d508e75ff.
//
// Solidity: event Extend(bytes32 nameHash, uint256 time)
func (_BasOwnership *BasOwnershipFilterer) FilterExtend(opts *bind.FilterOpts) (*BasOwnershipExtendIterator, error) {

	logs, sub, err := _BasOwnership.contract.FilterLogs(opts, "Extend")
	if err != nil {
		return nil, err
	}
	return &BasOwnershipExtendIterator{contract: _BasOwnership.contract, event: "Extend", logs: logs, sub: sub}, nil
}

// WatchExtend is a free log subscription operation binding the contract event 0xe898863b26adf3af54e82709ac1a76136eaaa5f3a7d9b790ce6d539d508e75ff.
//
// Solidity: event Extend(bytes32 nameHash, uint256 time)
func (_BasOwnership *BasOwnershipFilterer) WatchExtend(opts *bind.WatchOpts, sink chan<- *BasOwnershipExtend) (event.Subscription, error) {

	logs, sub, err := _BasOwnership.contract.WatchLogs(opts, "Extend")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasOwnershipExtend)
				if err := _BasOwnership.contract.UnpackLog(event, "Extend", log); err != nil {
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

// ParseExtend is a log parse operation binding the contract event 0xe898863b26adf3af54e82709ac1a76136eaaa5f3a7d9b790ce6d539d508e75ff.
//
// Solidity: event Extend(bytes32 nameHash, uint256 time)
func (_BasOwnership *BasOwnershipFilterer) ParseExtend(log types.Log) (*BasOwnershipExtend, error) {
	event := new(BasOwnershipExtend)
	if err := _BasOwnership.contract.UnpackLog(event, "Extend", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BasOwnershipRemoveIterator is returned from FilterRemove and is used to iterate over the raw logs and unpacked data for Remove events raised by the BasOwnership contract.
type BasOwnershipRemoveIterator struct {
	Event *BasOwnershipRemove // Event containing the contract specifics and raw log

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
func (it *BasOwnershipRemoveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasOwnershipRemove)
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
		it.Event = new(BasOwnershipRemove)
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
func (it *BasOwnershipRemoveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasOwnershipRemoveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasOwnershipRemove represents a Remove event raised by the BasOwnership contract.
type BasOwnershipRemove struct {
	NameHash [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRemove is a free log retrieval operation binding the contract event 0xa56fb2a6d4126f324526f0668c53927c0cd8e08f41ba0fe0f2d6090a84bc75c8.
//
// Solidity: event Remove(bytes32 nameHash)
func (_BasOwnership *BasOwnershipFilterer) FilterRemove(opts *bind.FilterOpts) (*BasOwnershipRemoveIterator, error) {

	logs, sub, err := _BasOwnership.contract.FilterLogs(opts, "Remove")
	if err != nil {
		return nil, err
	}
	return &BasOwnershipRemoveIterator{contract: _BasOwnership.contract, event: "Remove", logs: logs, sub: sub}, nil
}

// WatchRemove is a free log subscription operation binding the contract event 0xa56fb2a6d4126f324526f0668c53927c0cd8e08f41ba0fe0f2d6090a84bc75c8.
//
// Solidity: event Remove(bytes32 nameHash)
func (_BasOwnership *BasOwnershipFilterer) WatchRemove(opts *bind.WatchOpts, sink chan<- *BasOwnershipRemove) (event.Subscription, error) {

	logs, sub, err := _BasOwnership.contract.WatchLogs(opts, "Remove")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasOwnershipRemove)
				if err := _BasOwnership.contract.UnpackLog(event, "Remove", log); err != nil {
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

// ParseRemove is a log parse operation binding the contract event 0xa56fb2a6d4126f324526f0668c53927c0cd8e08f41ba0fe0f2d6090a84bc75c8.
//
// Solidity: event Remove(bytes32 nameHash)
func (_BasOwnership *BasOwnershipFilterer) ParseRemove(log types.Log) (*BasOwnershipRemove, error) {
	event := new(BasOwnershipRemove)
	if err := _BasOwnership.contract.UnpackLog(event, "Remove", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BasOwnershipTakeoverIterator is returned from FilterTakeover and is used to iterate over the raw logs and unpacked data for Takeover events raised by the BasOwnership contract.
type BasOwnershipTakeoverIterator struct {
	Event *BasOwnershipTakeover // Event containing the contract specifics and raw log

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
func (it *BasOwnershipTakeoverIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasOwnershipTakeover)
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
		it.Event = new(BasOwnershipTakeover)
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
func (it *BasOwnershipTakeoverIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasOwnershipTakeoverIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasOwnershipTakeover represents a Takeover event raised by the BasOwnership contract.
type BasOwnershipTakeover struct {
	NameHash [32]byte
	From     common.Address
	To       common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTakeover is a free log retrieval operation binding the contract event 0xae53ff61a227b196db64e2dea26e80d455bb2d90fd10372db78b44ace9cb0f62.
//
// Solidity: event Takeover(bytes32 nameHash, address from, address to)
func (_BasOwnership *BasOwnershipFilterer) FilterTakeover(opts *bind.FilterOpts) (*BasOwnershipTakeoverIterator, error) {

	logs, sub, err := _BasOwnership.contract.FilterLogs(opts, "Takeover")
	if err != nil {
		return nil, err
	}
	return &BasOwnershipTakeoverIterator{contract: _BasOwnership.contract, event: "Takeover", logs: logs, sub: sub}, nil
}

// WatchTakeover is a free log subscription operation binding the contract event 0xae53ff61a227b196db64e2dea26e80d455bb2d90fd10372db78b44ace9cb0f62.
//
// Solidity: event Takeover(bytes32 nameHash, address from, address to)
func (_BasOwnership *BasOwnershipFilterer) WatchTakeover(opts *bind.WatchOpts, sink chan<- *BasOwnershipTakeover) (event.Subscription, error) {

	logs, sub, err := _BasOwnership.contract.WatchLogs(opts, "Takeover")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasOwnershipTakeover)
				if err := _BasOwnership.contract.UnpackLog(event, "Takeover", log); err != nil {
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

// ParseTakeover is a log parse operation binding the contract event 0xae53ff61a227b196db64e2dea26e80d455bb2d90fd10372db78b44ace9cb0f62.
//
// Solidity: event Takeover(bytes32 nameHash, address from, address to)
func (_BasOwnership *BasOwnershipFilterer) ParseTakeover(log types.Log) (*BasOwnershipTakeover, error) {
	event := new(BasOwnershipTakeover)
	if err := _BasOwnership.contract.UnpackLog(event, "Takeover", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BasOwnershipTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BasOwnership contract.
type BasOwnershipTransferIterator struct {
	Event *BasOwnershipTransfer // Event containing the contract specifics and raw log

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
func (it *BasOwnershipTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasOwnershipTransfer)
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
		it.Event = new(BasOwnershipTransfer)
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
func (it *BasOwnershipTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasOwnershipTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasOwnershipTransfer represents a Transfer event raised by the BasOwnership contract.
type BasOwnershipTransfer struct {
	NameHash [32]byte
	From     common.Address
	To       common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0x57972c0bcb114a9f52d3501767c95a65e93901ff48da6677c7399b5593b4e999.
//
// Solidity: event Transfer(bytes32 nameHash, address from, address to)
func (_BasOwnership *BasOwnershipFilterer) FilterTransfer(opts *bind.FilterOpts) (*BasOwnershipTransferIterator, error) {

	logs, sub, err := _BasOwnership.contract.FilterLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return &BasOwnershipTransferIterator{contract: _BasOwnership.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0x57972c0bcb114a9f52d3501767c95a65e93901ff48da6677c7399b5593b4e999.
//
// Solidity: event Transfer(bytes32 nameHash, address from, address to)
func (_BasOwnership *BasOwnershipFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BasOwnershipTransfer) (event.Subscription, error) {

	logs, sub, err := _BasOwnership.contract.WatchLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasOwnershipTransfer)
				if err := _BasOwnership.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0x57972c0bcb114a9f52d3501767c95a65e93901ff48da6677c7399b5593b4e999.
//
// Solidity: event Transfer(bytes32 nameHash, address from, address to)
func (_BasOwnership *BasOwnershipFilterer) ParseTransfer(log types.Log) (*BasOwnershipTransfer, error) {
	event := new(BasOwnershipTransfer)
	if err := _BasOwnership.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BasOwnershipTransferFromIterator is returned from FilterTransferFrom and is used to iterate over the raw logs and unpacked data for TransferFrom events raised by the BasOwnership contract.
type BasOwnershipTransferFromIterator struct {
	Event *BasOwnershipTransferFrom // Event containing the contract specifics and raw log

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
func (it *BasOwnershipTransferFromIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasOwnershipTransferFrom)
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
		it.Event = new(BasOwnershipTransferFrom)
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
func (it *BasOwnershipTransferFromIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasOwnershipTransferFromIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasOwnershipTransferFrom represents a TransferFrom event raised by the BasOwnership contract.
type BasOwnershipTransferFrom struct {
	NameHash [32]byte
	From     common.Address
	To       common.Address
	By       common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferFrom is a free log retrieval operation binding the contract event 0x3f8c833a27e9521421e46d68c5a31e603cc465fc4bbff2226b98d964252ebb3f.
//
// Solidity: event TransferFrom(bytes32 nameHash, address from, address to, address by)
func (_BasOwnership *BasOwnershipFilterer) FilterTransferFrom(opts *bind.FilterOpts) (*BasOwnershipTransferFromIterator, error) {

	logs, sub, err := _BasOwnership.contract.FilterLogs(opts, "TransferFrom")
	if err != nil {
		return nil, err
	}
	return &BasOwnershipTransferFromIterator{contract: _BasOwnership.contract, event: "TransferFrom", logs: logs, sub: sub}, nil
}

// WatchTransferFrom is a free log subscription operation binding the contract event 0x3f8c833a27e9521421e46d68c5a31e603cc465fc4bbff2226b98d964252ebb3f.
//
// Solidity: event TransferFrom(bytes32 nameHash, address from, address to, address by)
func (_BasOwnership *BasOwnershipFilterer) WatchTransferFrom(opts *bind.WatchOpts, sink chan<- *BasOwnershipTransferFrom) (event.Subscription, error) {

	logs, sub, err := _BasOwnership.contract.WatchLogs(opts, "TransferFrom")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasOwnershipTransferFrom)
				if err := _BasOwnership.contract.UnpackLog(event, "TransferFrom", log); err != nil {
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

// ParseTransferFrom is a log parse operation binding the contract event 0x3f8c833a27e9521421e46d68c5a31e603cc465fc4bbff2226b98d964252ebb3f.
//
// Solidity: event TransferFrom(bytes32 nameHash, address from, address to, address by)
func (_BasOwnership *BasOwnershipFilterer) ParseTransferFrom(log types.Log) (*BasOwnershipTransferFrom, error) {
	event := new(BasOwnershipTransferFrom)
	if err := _BasOwnership.contract.UnpackLog(event, "TransferFrom", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BasOwnershipUpdateIterator is returned from FilterUpdate and is used to iterate over the raw logs and unpacked data for Update events raised by the BasOwnership contract.
type BasOwnershipUpdateIterator struct {
	Event *BasOwnershipUpdate // Event containing the contract specifics and raw log

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
func (it *BasOwnershipUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasOwnershipUpdate)
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
		it.Event = new(BasOwnershipUpdate)
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
func (it *BasOwnershipUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasOwnershipUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasOwnershipUpdate represents a Update event raised by the BasOwnership contract.
type BasOwnershipUpdate struct {
	NameHash [32]byte
	Owner    common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdate is a free log retrieval operation binding the contract event 0xb2b848c23b147f2d3c11308c74cb0ff81a30b54628e5c73c4d70bf623eb98f72.
//
// Solidity: event Update(bytes32 nameHash, address owner)
func (_BasOwnership *BasOwnershipFilterer) FilterUpdate(opts *bind.FilterOpts) (*BasOwnershipUpdateIterator, error) {

	logs, sub, err := _BasOwnership.contract.FilterLogs(opts, "Update")
	if err != nil {
		return nil, err
	}
	return &BasOwnershipUpdateIterator{contract: _BasOwnership.contract, event: "Update", logs: logs, sub: sub}, nil
}

// WatchUpdate is a free log subscription operation binding the contract event 0xb2b848c23b147f2d3c11308c74cb0ff81a30b54628e5c73c4d70bf623eb98f72.
//
// Solidity: event Update(bytes32 nameHash, address owner)
func (_BasOwnership *BasOwnershipFilterer) WatchUpdate(opts *bind.WatchOpts, sink chan<- *BasOwnershipUpdate) (event.Subscription, error) {

	logs, sub, err := _BasOwnership.contract.WatchLogs(opts, "Update")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasOwnershipUpdate)
				if err := _BasOwnership.contract.UnpackLog(event, "Update", log); err != nil {
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

// ParseUpdate is a log parse operation binding the contract event 0xb2b848c23b147f2d3c11308c74cb0ff81a30b54628e5c73c4d70bf623eb98f72.
//
// Solidity: event Update(bytes32 nameHash, address owner)
func (_BasOwnership *BasOwnershipFilterer) ParseUpdate(log types.Log) (*BasOwnershipUpdate, error) {
	event := new(BasOwnershipUpdate)
	if err := _BasOwnership.contract.UnpackLog(event, "Update", log); err != nil {
		return nil, err
	}
	return event, nil
}
