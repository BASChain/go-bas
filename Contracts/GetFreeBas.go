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

// GetFreeBasABI is the input ABI used to generate the binding from.
const GetFreeBasABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"v\",\"type\":\"uint256\"}],\"name\":\"setAmount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"applyRecord\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"ApplyToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ContractOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferContractOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"amount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"GetBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractERC20\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// GetFreeBas is an auto generated Go binding around an Ethereum contract.
type GetFreeBas struct {
	GetFreeBasCaller     // Read-only binding to the contract
	GetFreeBasTransactor // Write-only binding to the contract
	GetFreeBasFilterer   // Log filterer for contract events
}

// GetFreeBasCaller is an auto generated read-only Go binding around an Ethereum contract.
type GetFreeBasCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GetFreeBasTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GetFreeBasTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GetFreeBasFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GetFreeBasFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GetFreeBasSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GetFreeBasSession struct {
	Contract     *GetFreeBas       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GetFreeBasCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GetFreeBasCallerSession struct {
	Contract *GetFreeBasCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// GetFreeBasTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GetFreeBasTransactorSession struct {
	Contract     *GetFreeBasTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// GetFreeBasRaw is an auto generated low-level Go binding around an Ethereum contract.
type GetFreeBasRaw struct {
	Contract *GetFreeBas // Generic contract binding to access the raw methods on
}

// GetFreeBasCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GetFreeBasCallerRaw struct {
	Contract *GetFreeBasCaller // Generic read-only contract binding to access the raw methods on
}

// GetFreeBasTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GetFreeBasTransactorRaw struct {
	Contract *GetFreeBasTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGetFreeBas creates a new instance of GetFreeBas, bound to a specific deployed contract.
func NewGetFreeBas(address common.Address, backend bind.ContractBackend) (*GetFreeBas, error) {
	contract, err := bindGetFreeBas(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GetFreeBas{GetFreeBasCaller: GetFreeBasCaller{contract: contract}, GetFreeBasTransactor: GetFreeBasTransactor{contract: contract}, GetFreeBasFilterer: GetFreeBasFilterer{contract: contract}}, nil
}

// NewGetFreeBasCaller creates a new read-only instance of GetFreeBas, bound to a specific deployed contract.
func NewGetFreeBasCaller(address common.Address, caller bind.ContractCaller) (*GetFreeBasCaller, error) {
	contract, err := bindGetFreeBas(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GetFreeBasCaller{contract: contract}, nil
}

// NewGetFreeBasTransactor creates a new write-only instance of GetFreeBas, bound to a specific deployed contract.
func NewGetFreeBasTransactor(address common.Address, transactor bind.ContractTransactor) (*GetFreeBasTransactor, error) {
	contract, err := bindGetFreeBas(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GetFreeBasTransactor{contract: contract}, nil
}

// NewGetFreeBasFilterer creates a new log filterer instance of GetFreeBas, bound to a specific deployed contract.
func NewGetFreeBasFilterer(address common.Address, filterer bind.ContractFilterer) (*GetFreeBasFilterer, error) {
	contract, err := bindGetFreeBas(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GetFreeBasFilterer{contract: contract}, nil
}

// bindGetFreeBas binds a generic wrapper to an already deployed contract.
func bindGetFreeBas(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GetFreeBasABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GetFreeBas *GetFreeBasRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GetFreeBas.Contract.GetFreeBasCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GetFreeBas *GetFreeBasRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GetFreeBas.Contract.GetFreeBasTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GetFreeBas *GetFreeBasRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GetFreeBas.Contract.GetFreeBasTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GetFreeBas *GetFreeBasCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GetFreeBas.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GetFreeBas *GetFreeBasTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GetFreeBas.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GetFreeBas *GetFreeBasTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GetFreeBas.Contract.contract.Transact(opts, method, params...)
}

// ContractOwner is a free data retrieval call binding the contract method 0x5a63fbc9.
//
// Solidity: function ContractOwner() constant returns(address)
func (_GetFreeBas *GetFreeBasCaller) ContractOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _GetFreeBas.contract.Call(opts, out, "ContractOwner")
	return *ret0, err
}

// ContractOwner is a free data retrieval call binding the contract method 0x5a63fbc9.
//
// Solidity: function ContractOwner() constant returns(address)
func (_GetFreeBas *GetFreeBasSession) ContractOwner() (common.Address, error) {
	return _GetFreeBas.Contract.ContractOwner(&_GetFreeBas.CallOpts)
}

// ContractOwner is a free data retrieval call binding the contract method 0x5a63fbc9.
//
// Solidity: function ContractOwner() constant returns(address)
func (_GetFreeBas *GetFreeBasCallerSession) ContractOwner() (common.Address, error) {
	return _GetFreeBas.Contract.ContractOwner(&_GetFreeBas.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8f8a912.
//
// Solidity: function GetBalance() constant returns(uint256)
func (_GetFreeBas *GetFreeBasCaller) GetBalance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _GetFreeBas.contract.Call(opts, out, "GetBalance")
	return *ret0, err
}

// GetBalance is a free data retrieval call binding the contract method 0xf8f8a912.
//
// Solidity: function GetBalance() constant returns(uint256)
func (_GetFreeBas *GetFreeBasSession) GetBalance() (*big.Int, error) {
	return _GetFreeBas.Contract.GetBalance(&_GetFreeBas.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8f8a912.
//
// Solidity: function GetBalance() constant returns(uint256)
func (_GetFreeBas *GetFreeBasCallerSession) GetBalance() (*big.Int, error) {
	return _GetFreeBas.Contract.GetBalance(&_GetFreeBas.CallOpts)
}

// Amount is a free data retrieval call binding the contract method 0xaa8c217c.
//
// Solidity: function amount() constant returns(uint256)
func (_GetFreeBas *GetFreeBasCaller) Amount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _GetFreeBas.contract.Call(opts, out, "amount")
	return *ret0, err
}

// Amount is a free data retrieval call binding the contract method 0xaa8c217c.
//
// Solidity: function amount() constant returns(uint256)
func (_GetFreeBas *GetFreeBasSession) Amount() (*big.Int, error) {
	return _GetFreeBas.Contract.Amount(&_GetFreeBas.CallOpts)
}

// Amount is a free data retrieval call binding the contract method 0xaa8c217c.
//
// Solidity: function amount() constant returns(uint256)
func (_GetFreeBas *GetFreeBasCallerSession) Amount() (*big.Int, error) {
	return _GetFreeBas.Contract.Amount(&_GetFreeBas.CallOpts)
}

// ApplyRecord is a free data retrieval call binding the contract method 0x2b18efb0.
//
// Solidity: function applyRecord(address ) constant returns(bool)
func (_GetFreeBas *GetFreeBasCaller) ApplyRecord(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _GetFreeBas.contract.Call(opts, out, "applyRecord", arg0)
	return *ret0, err
}

// ApplyRecord is a free data retrieval call binding the contract method 0x2b18efb0.
//
// Solidity: function applyRecord(address ) constant returns(bool)
func (_GetFreeBas *GetFreeBasSession) ApplyRecord(arg0 common.Address) (bool, error) {
	return _GetFreeBas.Contract.ApplyRecord(&_GetFreeBas.CallOpts, arg0)
}

// ApplyRecord is a free data retrieval call binding the contract method 0x2b18efb0.
//
// Solidity: function applyRecord(address ) constant returns(bool)
func (_GetFreeBas *GetFreeBasCallerSession) ApplyRecord(arg0 common.Address) (bool, error) {
	return _GetFreeBas.Contract.ApplyRecord(&_GetFreeBas.CallOpts, arg0)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_GetFreeBas *GetFreeBasCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _GetFreeBas.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_GetFreeBas *GetFreeBasSession) Token() (common.Address, error) {
	return _GetFreeBas.Contract.Token(&_GetFreeBas.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_GetFreeBas *GetFreeBasCallerSession) Token() (common.Address, error) {
	return _GetFreeBas.Contract.Token(&_GetFreeBas.CallOpts)
}

// ApplyToken is a paid mutator transaction binding the contract method 0x402f5c12.
//
// Solidity: function ApplyToken() returns()
func (_GetFreeBas *GetFreeBasTransactor) ApplyToken(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GetFreeBas.contract.Transact(opts, "ApplyToken")
}

// ApplyToken is a paid mutator transaction binding the contract method 0x402f5c12.
//
// Solidity: function ApplyToken() returns()
func (_GetFreeBas *GetFreeBasSession) ApplyToken() (*types.Transaction, error) {
	return _GetFreeBas.Contract.ApplyToken(&_GetFreeBas.TransactOpts)
}

// ApplyToken is a paid mutator transaction binding the contract method 0x402f5c12.
//
// Solidity: function ApplyToken() returns()
func (_GetFreeBas *GetFreeBasTransactorSession) ApplyToken() (*types.Transaction, error) {
	return _GetFreeBas.Contract.ApplyToken(&_GetFreeBas.TransactOpts)
}

// SetAmount is a paid mutator transaction binding the contract method 0x271f88b4.
//
// Solidity: function setAmount(uint256 v) returns()
func (_GetFreeBas *GetFreeBasTransactor) SetAmount(opts *bind.TransactOpts, v *big.Int) (*types.Transaction, error) {
	return _GetFreeBas.contract.Transact(opts, "setAmount", v)
}

// SetAmount is a paid mutator transaction binding the contract method 0x271f88b4.
//
// Solidity: function setAmount(uint256 v) returns()
func (_GetFreeBas *GetFreeBasSession) SetAmount(v *big.Int) (*types.Transaction, error) {
	return _GetFreeBas.Contract.SetAmount(&_GetFreeBas.TransactOpts, v)
}

// SetAmount is a paid mutator transaction binding the contract method 0x271f88b4.
//
// Solidity: function setAmount(uint256 v) returns()
func (_GetFreeBas *GetFreeBasTransactorSession) SetAmount(v *big.Int) (*types.Transaction, error) {
	return _GetFreeBas.Contract.SetAmount(&_GetFreeBas.TransactOpts, v)
}

// TransferContractOwnership is a paid mutator transaction binding the contract method 0xa843c51f.
//
// Solidity: function transferContractOwnership(address newOwner) returns()
func (_GetFreeBas *GetFreeBasTransactor) TransferContractOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _GetFreeBas.contract.Transact(opts, "transferContractOwnership", newOwner)
}

// TransferContractOwnership is a paid mutator transaction binding the contract method 0xa843c51f.
//
// Solidity: function transferContractOwnership(address newOwner) returns()
func (_GetFreeBas *GetFreeBasSession) TransferContractOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GetFreeBas.Contract.TransferContractOwnership(&_GetFreeBas.TransactOpts, newOwner)
}

// TransferContractOwnership is a paid mutator transaction binding the contract method 0xa843c51f.
//
// Solidity: function transferContractOwnership(address newOwner) returns()
func (_GetFreeBas *GetFreeBasTransactorSession) TransferContractOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GetFreeBas.Contract.TransferContractOwnership(&_GetFreeBas.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_GetFreeBas *GetFreeBasTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GetFreeBas.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_GetFreeBas *GetFreeBasSession) Withdraw() (*types.Transaction, error) {
	return _GetFreeBas.Contract.Withdraw(&_GetFreeBas.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_GetFreeBas *GetFreeBasTransactorSession) Withdraw() (*types.Transaction, error) {
	return _GetFreeBas.Contract.Withdraw(&_GetFreeBas.TransactOpts)
}
