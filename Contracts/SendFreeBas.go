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

// SendFreeBasABI is the input ABI used to generate the binding from.
const SendFreeBasABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SendTokenByContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"applyRecord\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractERC20\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// SendFreeBas is an auto generated Go binding around an Ethereum contract.
type SendFreeBas struct {
	SendFreeBasCaller     // Read-only binding to the contract
	SendFreeBasTransactor // Write-only binding to the contract
	SendFreeBasFilterer   // Log filterer for contract events
}

// SendFreeBasCaller is an auto generated read-only Go binding around an Ethereum contract.
type SendFreeBasCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SendFreeBasTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SendFreeBasTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SendFreeBasFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SendFreeBasFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SendFreeBasSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SendFreeBasSession struct {
	Contract     *SendFreeBas      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SendFreeBasCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SendFreeBasCallerSession struct {
	Contract *SendFreeBasCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// SendFreeBasTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SendFreeBasTransactorSession struct {
	Contract     *SendFreeBasTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SendFreeBasRaw is an auto generated low-level Go binding around an Ethereum contract.
type SendFreeBasRaw struct {
	Contract *SendFreeBas // Generic contract binding to access the raw methods on
}

// SendFreeBasCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SendFreeBasCallerRaw struct {
	Contract *SendFreeBasCaller // Generic read-only contract binding to access the raw methods on
}

// SendFreeBasTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SendFreeBasTransactorRaw struct {
	Contract *SendFreeBasTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSendFreeBas creates a new instance of SendFreeBas, bound to a specific deployed contract.
func NewSendFreeBas(address common.Address, backend bind.ContractBackend) (*SendFreeBas, error) {
	contract, err := bindSendFreeBas(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SendFreeBas{SendFreeBasCaller: SendFreeBasCaller{contract: contract}, SendFreeBasTransactor: SendFreeBasTransactor{contract: contract}, SendFreeBasFilterer: SendFreeBasFilterer{contract: contract}}, nil
}

// NewSendFreeBasCaller creates a new read-only instance of SendFreeBas, bound to a specific deployed contract.
func NewSendFreeBasCaller(address common.Address, caller bind.ContractCaller) (*SendFreeBasCaller, error) {
	contract, err := bindSendFreeBas(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SendFreeBasCaller{contract: contract}, nil
}

// NewSendFreeBasTransactor creates a new write-only instance of SendFreeBas, bound to a specific deployed contract.
func NewSendFreeBasTransactor(address common.Address, transactor bind.ContractTransactor) (*SendFreeBasTransactor, error) {
	contract, err := bindSendFreeBas(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SendFreeBasTransactor{contract: contract}, nil
}

// NewSendFreeBasFilterer creates a new log filterer instance of SendFreeBas, bound to a specific deployed contract.
func NewSendFreeBasFilterer(address common.Address, filterer bind.ContractFilterer) (*SendFreeBasFilterer, error) {
	contract, err := bindSendFreeBas(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SendFreeBasFilterer{contract: contract}, nil
}

// bindSendFreeBas binds a generic wrapper to an already deployed contract.
func bindSendFreeBas(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SendFreeBasABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SendFreeBas *SendFreeBasRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SendFreeBas.Contract.SendFreeBasCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SendFreeBas *SendFreeBasRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SendFreeBas.Contract.SendFreeBasTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SendFreeBas *SendFreeBasRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SendFreeBas.Contract.SendFreeBasTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SendFreeBas *SendFreeBasCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SendFreeBas.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SendFreeBas *SendFreeBasTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SendFreeBas.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SendFreeBas *SendFreeBasTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SendFreeBas.Contract.contract.Transact(opts, method, params...)
}

// ApplyRecord is a free data retrieval call binding the contract method 0x2b18efb0.
//
// Solidity: function applyRecord(address ) constant returns(bool)
func (_SendFreeBas *SendFreeBasCaller) ApplyRecord(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SendFreeBas.contract.Call(opts, out, "applyRecord", arg0)
	return *ret0, err
}

// ApplyRecord is a free data retrieval call binding the contract method 0x2b18efb0.
//
// Solidity: function applyRecord(address ) constant returns(bool)
func (_SendFreeBas *SendFreeBasSession) ApplyRecord(arg0 common.Address) (bool, error) {
	return _SendFreeBas.Contract.ApplyRecord(&_SendFreeBas.CallOpts, arg0)
}

// ApplyRecord is a free data retrieval call binding the contract method 0x2b18efb0.
//
// Solidity: function applyRecord(address ) constant returns(bool)
func (_SendFreeBas *SendFreeBasCallerSession) ApplyRecord(arg0 common.Address) (bool, error) {
	return _SendFreeBas.Contract.ApplyRecord(&_SendFreeBas.CallOpts, arg0)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_SendFreeBas *SendFreeBasCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SendFreeBas.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_SendFreeBas *SendFreeBasSession) Token() (common.Address, error) {
	return _SendFreeBas.Contract.Token(&_SendFreeBas.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_SendFreeBas *SendFreeBasCallerSession) Token() (common.Address, error) {
	return _SendFreeBas.Contract.Token(&_SendFreeBas.CallOpts)
}

// SendTokenByContract is a paid mutator transaction binding the contract method 0x74076b72.
//
// Solidity: function SendTokenByContract(address to, uint256 amount) returns()
func (_SendFreeBas *SendFreeBasTransactor) SendTokenByContract(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SendFreeBas.contract.Transact(opts, "SendTokenByContract", to, amount)
}

// SendTokenByContract is a paid mutator transaction binding the contract method 0x74076b72.
//
// Solidity: function SendTokenByContract(address to, uint256 amount) returns()
func (_SendFreeBas *SendFreeBasSession) SendTokenByContract(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SendFreeBas.Contract.SendTokenByContract(&_SendFreeBas.TransactOpts, to, amount)
}

// SendTokenByContract is a paid mutator transaction binding the contract method 0x74076b72.
//
// Solidity: function SendTokenByContract(address to, uint256 amount) returns()
func (_SendFreeBas *SendFreeBasTransactorSession) SendTokenByContract(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SendFreeBas.Contract.SendTokenByContract(&_SendFreeBas.TransactOpts, to, amount)
}
