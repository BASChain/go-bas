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

// BasRuleABI is the input ABI used to generate the binding from.
const BasRuleABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"rareLength\",\"type\":\"uint256\"}],\"name\":\"classifyRoot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"rareLength\",\"type\":\"uint256\"}],\"name\":\"classifyRootS\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes1\",\"name\":\"c\",\"type\":\"bytes1\"}],\"name\":\"validChar\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"rootLength\",\"type\":\"uint256\"}],\"name\":\"verifySub\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"rootLength\",\"type\":\"uint256\"}],\"name\":\"verifySubS\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// BasRule is an auto generated Go binding around an Ethereum contract.
type BasRule struct {
	BasRuleCaller     // Read-only binding to the contract
	BasRuleTransactor // Write-only binding to the contract
	BasRuleFilterer   // Log filterer for contract events
}

// BasRuleCaller is an auto generated read-only Go binding around an Ethereum contract.
type BasRuleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasRuleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BasRuleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasRuleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BasRuleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasRuleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BasRuleSession struct {
	Contract     *BasRule          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BasRuleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BasRuleCallerSession struct {
	Contract *BasRuleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// BasRuleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BasRuleTransactorSession struct {
	Contract     *BasRuleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// BasRuleRaw is an auto generated low-level Go binding around an Ethereum contract.
type BasRuleRaw struct {
	Contract *BasRule // Generic contract binding to access the raw methods on
}

// BasRuleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BasRuleCallerRaw struct {
	Contract *BasRuleCaller // Generic read-only contract binding to access the raw methods on
}

// BasRuleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BasRuleTransactorRaw struct {
	Contract *BasRuleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBasRule creates a new instance of BasRule, bound to a specific deployed contract.
func NewBasRule(address common.Address, backend bind.ContractBackend) (*BasRule, error) {
	contract, err := bindBasRule(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BasRule{BasRuleCaller: BasRuleCaller{contract: contract}, BasRuleTransactor: BasRuleTransactor{contract: contract}, BasRuleFilterer: BasRuleFilterer{contract: contract}}, nil
}

// NewBasRuleCaller creates a new read-only instance of BasRule, bound to a specific deployed contract.
func NewBasRuleCaller(address common.Address, caller bind.ContractCaller) (*BasRuleCaller, error) {
	contract, err := bindBasRule(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BasRuleCaller{contract: contract}, nil
}

// NewBasRuleTransactor creates a new write-only instance of BasRule, bound to a specific deployed contract.
func NewBasRuleTransactor(address common.Address, transactor bind.ContractTransactor) (*BasRuleTransactor, error) {
	contract, err := bindBasRule(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BasRuleTransactor{contract: contract}, nil
}

// NewBasRuleFilterer creates a new log filterer instance of BasRule, bound to a specific deployed contract.
func NewBasRuleFilterer(address common.Address, filterer bind.ContractFilterer) (*BasRuleFilterer, error) {
	contract, err := bindBasRule(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BasRuleFilterer{contract: contract}, nil
}

// bindBasRule binds a generic wrapper to an already deployed contract.
func bindBasRule(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BasRuleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasRule *BasRuleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BasRule.Contract.BasRuleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasRule *BasRuleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasRule.Contract.BasRuleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasRule *BasRuleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasRule.Contract.BasRuleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasRule *BasRuleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BasRule.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasRule *BasRuleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasRule.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasRule *BasRuleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasRule.Contract.contract.Transact(opts, method, params...)
}

// ClassifyRoot is a free data retrieval call binding the contract method 0x20d912fa.
//
// Solidity: function classifyRoot(bytes name, uint256 rareLength) constant returns(bool, bool)
func (_BasRule *BasRuleCaller) ClassifyRoot(opts *bind.CallOpts, name []byte, rareLength *big.Int) (bool, bool, error) {
	var (
		ret0 = new(bool)
		ret1 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _BasRule.contract.Call(opts, out, "classifyRoot", name, rareLength)
	return *ret0, *ret1, err
}

// ClassifyRoot is a free data retrieval call binding the contract method 0x20d912fa.
//
// Solidity: function classifyRoot(bytes name, uint256 rareLength) constant returns(bool, bool)
func (_BasRule *BasRuleSession) ClassifyRoot(name []byte, rareLength *big.Int) (bool, bool, error) {
	return _BasRule.Contract.ClassifyRoot(&_BasRule.CallOpts, name, rareLength)
}

// ClassifyRoot is a free data retrieval call binding the contract method 0x20d912fa.
//
// Solidity: function classifyRoot(bytes name, uint256 rareLength) constant returns(bool, bool)
func (_BasRule *BasRuleCallerSession) ClassifyRoot(name []byte, rareLength *big.Int) (bool, bool, error) {
	return _BasRule.Contract.ClassifyRoot(&_BasRule.CallOpts, name, rareLength)
}

// ClassifyRootS is a free data retrieval call binding the contract method 0x804e05dd.
//
// Solidity: function classifyRootS(string name, uint256 rareLength) constant returns(bool, bool)
func (_BasRule *BasRuleCaller) ClassifyRootS(opts *bind.CallOpts, name string, rareLength *big.Int) (bool, bool, error) {
	var (
		ret0 = new(bool)
		ret1 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _BasRule.contract.Call(opts, out, "classifyRootS", name, rareLength)
	return *ret0, *ret1, err
}

// ClassifyRootS is a free data retrieval call binding the contract method 0x804e05dd.
//
// Solidity: function classifyRootS(string name, uint256 rareLength) constant returns(bool, bool)
func (_BasRule *BasRuleSession) ClassifyRootS(name string, rareLength *big.Int) (bool, bool, error) {
	return _BasRule.Contract.ClassifyRootS(&_BasRule.CallOpts, name, rareLength)
}

// ClassifyRootS is a free data retrieval call binding the contract method 0x804e05dd.
//
// Solidity: function classifyRootS(string name, uint256 rareLength) constant returns(bool, bool)
func (_BasRule *BasRuleCallerSession) ClassifyRootS(name string, rareLength *big.Int) (bool, bool, error) {
	return _BasRule.Contract.ClassifyRootS(&_BasRule.CallOpts, name, rareLength)
}

// ValidChar is a free data retrieval call binding the contract method 0xb508db39.
//
// Solidity: function validChar(bytes1 c) constant returns(bool)
func (_BasRule *BasRuleCaller) ValidChar(opts *bind.CallOpts, c [1]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BasRule.contract.Call(opts, out, "validChar", c)
	return *ret0, err
}

// ValidChar is a free data retrieval call binding the contract method 0xb508db39.
//
// Solidity: function validChar(bytes1 c) constant returns(bool)
func (_BasRule *BasRuleSession) ValidChar(c [1]byte) (bool, error) {
	return _BasRule.Contract.ValidChar(&_BasRule.CallOpts, c)
}

// ValidChar is a free data retrieval call binding the contract method 0xb508db39.
//
// Solidity: function validChar(bytes1 c) constant returns(bool)
func (_BasRule *BasRuleCallerSession) ValidChar(c [1]byte) (bool, error) {
	return _BasRule.Contract.ValidChar(&_BasRule.CallOpts, c)
}

// VerifySub is a free data retrieval call binding the contract method 0x1cd0c24d.
//
// Solidity: function verifySub(bytes name, uint256 rootLength) constant returns(bool)
func (_BasRule *BasRuleCaller) VerifySub(opts *bind.CallOpts, name []byte, rootLength *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BasRule.contract.Call(opts, out, "verifySub", name, rootLength)
	return *ret0, err
}

// VerifySub is a free data retrieval call binding the contract method 0x1cd0c24d.
//
// Solidity: function verifySub(bytes name, uint256 rootLength) constant returns(bool)
func (_BasRule *BasRuleSession) VerifySub(name []byte, rootLength *big.Int) (bool, error) {
	return _BasRule.Contract.VerifySub(&_BasRule.CallOpts, name, rootLength)
}

// VerifySub is a free data retrieval call binding the contract method 0x1cd0c24d.
//
// Solidity: function verifySub(bytes name, uint256 rootLength) constant returns(bool)
func (_BasRule *BasRuleCallerSession) VerifySub(name []byte, rootLength *big.Int) (bool, error) {
	return _BasRule.Contract.VerifySub(&_BasRule.CallOpts, name, rootLength)
}

// VerifySubS is a free data retrieval call binding the contract method 0x5ced286a.
//
// Solidity: function verifySubS(string name, uint256 rootLength) constant returns(bool)
func (_BasRule *BasRuleCaller) VerifySubS(opts *bind.CallOpts, name string, rootLength *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BasRule.contract.Call(opts, out, "verifySubS", name, rootLength)
	return *ret0, err
}

// VerifySubS is a free data retrieval call binding the contract method 0x5ced286a.
//
// Solidity: function verifySubS(string name, uint256 rootLength) constant returns(bool)
func (_BasRule *BasRuleSession) VerifySubS(name string, rootLength *big.Int) (bool, error) {
	return _BasRule.Contract.VerifySubS(&_BasRule.CallOpts, name, rootLength)
}

// VerifySubS is a free data retrieval call binding the contract method 0x5ced286a.
//
// Solidity: function verifySubS(string name, uint256 rootLength) constant returns(bool)
func (_BasRule *BasRuleCallerSession) VerifySubS(name string, rootLength *big.Int) (bool, error) {
	return _BasRule.Contract.VerifySubS(&_BasRule.CallOpts, name, rootLength)
}
