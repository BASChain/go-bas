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

// BasDNSABI is the input ABI used to generate the binding from.
const BasDNSABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"opData\",\"type\":\"bytes\"}],\"name\":\"_c_setOpData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes4\",\"name\":\"ipv4\",\"type\":\"bytes4\"},{\"internalType\":\"bytes16\",\"name\":\"ipv6\",\"type\":\"bytes16\"},{\"internalType\":\"bytes\",\"name\":\"bca\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"opData\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"aliasName\",\"type\":\"string\"}],\"name\":\"_c_update\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"aName\",\"type\":\"string\"}],\"name\":\"_c_setAlias\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes4\",\"name\":\"ipv4\",\"type\":\"bytes4\"},{\"internalType\":\"bytes16\",\"name\":\"ipv6\",\"type\":\"bytes16\"},{\"internalType\":\"bytes\",\"name\":\"bca\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"opData\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"aliasName\",\"type\":\"string\"}],\"name\":\"_a_update\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"o\",\"outputs\":[{\"internalType\":\"contractBasOwnership\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bcAddress\",\"type\":\"bytes\"}],\"name\":\"_c_setBCAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"}],\"name\":\"_c_clearRecord\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"}],\"name\":\"query\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"ipv4\",\"type\":\"bytes4\"},{\"internalType\":\"bytes16\",\"name\":\"ipv6\",\"type\":\"bytes16\"},{\"internalType\":\"bytes\",\"name\":\"bca\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"opData\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"aliasName\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes4\",\"name\":\"ipv4\",\"type\":\"bytes4\"},{\"internalType\":\"bytes16\",\"name\":\"ipv6\",\"type\":\"bytes16\"}],\"name\":\"_c_setIP\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"DNS\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"ipv4\",\"type\":\"bytes4\"},{\"internalType\":\"bytes16\",\"name\":\"ipv6\",\"type\":\"bytes16\"},{\"internalType\":\"bytes\",\"name\":\"bca\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"opData\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"aliasName\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_o\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"}],\"name\":\"DNSChanged\",\"type\":\"event\"}]"

// BasDNS is an auto generated Go binding around an Ethereum contract.
type BasDNS struct {
	BasDNSCaller     // Read-only binding to the contract
	BasDNSTransactor // Write-only binding to the contract
	BasDNSFilterer   // Log filterer for contract events
}

// BasDNSCaller is an auto generated read-only Go binding around an Ethereum contract.
type BasDNSCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasDNSTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BasDNSTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasDNSFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BasDNSFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasDNSSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BasDNSSession struct {
	Contract     *BasDNS           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BasDNSCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BasDNSCallerSession struct {
	Contract *BasDNSCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BasDNSTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BasDNSTransactorSession struct {
	Contract     *BasDNSTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BasDNSRaw is an auto generated low-level Go binding around an Ethereum contract.
type BasDNSRaw struct {
	Contract *BasDNS // Generic contract binding to access the raw methods on
}

// BasDNSCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BasDNSCallerRaw struct {
	Contract *BasDNSCaller // Generic read-only contract binding to access the raw methods on
}

// BasDNSTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BasDNSTransactorRaw struct {
	Contract *BasDNSTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBasDNS creates a new instance of BasDNS, bound to a specific deployed contract.
func NewBasDNS(address common.Address, backend bind.ContractBackend) (*BasDNS, error) {
	contract, err := bindBasDNS(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BasDNS{BasDNSCaller: BasDNSCaller{contract: contract}, BasDNSTransactor: BasDNSTransactor{contract: contract}, BasDNSFilterer: BasDNSFilterer{contract: contract}}, nil
}

// NewBasDNSCaller creates a new read-only instance of BasDNS, bound to a specific deployed contract.
func NewBasDNSCaller(address common.Address, caller bind.ContractCaller) (*BasDNSCaller, error) {
	contract, err := bindBasDNS(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BasDNSCaller{contract: contract}, nil
}

// NewBasDNSTransactor creates a new write-only instance of BasDNS, bound to a specific deployed contract.
func NewBasDNSTransactor(address common.Address, transactor bind.ContractTransactor) (*BasDNSTransactor, error) {
	contract, err := bindBasDNS(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BasDNSTransactor{contract: contract}, nil
}

// NewBasDNSFilterer creates a new log filterer instance of BasDNS, bound to a specific deployed contract.
func NewBasDNSFilterer(address common.Address, filterer bind.ContractFilterer) (*BasDNSFilterer, error) {
	contract, err := bindBasDNS(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BasDNSFilterer{contract: contract}, nil
}

// bindBasDNS binds a generic wrapper to an already deployed contract.
func bindBasDNS(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BasDNSABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasDNS *BasDNSRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BasDNS.Contract.BasDNSCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasDNS *BasDNSRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasDNS.Contract.BasDNSTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasDNS *BasDNSRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasDNS.Contract.BasDNSTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasDNS *BasDNSCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BasDNS.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasDNS *BasDNSTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasDNS.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasDNS *BasDNSTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasDNS.Contract.contract.Transact(opts, method, params...)
}

// DNS is a free data retrieval call binding the contract method 0xd241cf8e.
//
// Solidity: function DNS(bytes32 ) constant returns(bytes4 ipv4, bytes16 ipv6, bytes bca, bytes opData, string aliasName)
func (_BasDNS *BasDNSCaller) DNS(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Ipv4      [4]byte
	Ipv6      [16]byte
	Bca       []byte
	OpData    []byte
	AliasName string
}, error) {
	ret := new(struct {
		Ipv4      [4]byte
		Ipv6      [16]byte
		Bca       []byte
		OpData    []byte
		AliasName string
	})
	out := ret
	err := _BasDNS.contract.Call(opts, out, "DNS", arg0)
	return *ret, err
}

// DNS is a free data retrieval call binding the contract method 0xd241cf8e.
//
// Solidity: function DNS(bytes32 ) constant returns(bytes4 ipv4, bytes16 ipv6, bytes bca, bytes opData, string aliasName)
func (_BasDNS *BasDNSSession) DNS(arg0 [32]byte) (struct {
	Ipv4      [4]byte
	Ipv6      [16]byte
	Bca       []byte
	OpData    []byte
	AliasName string
}, error) {
	return _BasDNS.Contract.DNS(&_BasDNS.CallOpts, arg0)
}

// DNS is a free data retrieval call binding the contract method 0xd241cf8e.
//
// Solidity: function DNS(bytes32 ) constant returns(bytes4 ipv4, bytes16 ipv6, bytes bca, bytes opData, string aliasName)
func (_BasDNS *BasDNSCallerSession) DNS(arg0 [32]byte) (struct {
	Ipv4      [4]byte
	Ipv6      [16]byte
	Bca       []byte
	OpData    []byte
	AliasName string
}, error) {
	return _BasDNS.Contract.DNS(&_BasDNS.CallOpts, arg0)
}

// O is a free data retrieval call binding the contract method 0x50cd4df2.
//
// Solidity: function o() constant returns(address)
func (_BasDNS *BasDNSCaller) O(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BasDNS.contract.Call(opts, out, "o")
	return *ret0, err
}

// O is a free data retrieval call binding the contract method 0x50cd4df2.
//
// Solidity: function o() constant returns(address)
func (_BasDNS *BasDNSSession) O() (common.Address, error) {
	return _BasDNS.Contract.O(&_BasDNS.CallOpts)
}

// O is a free data retrieval call binding the contract method 0x50cd4df2.
//
// Solidity: function o() constant returns(address)
func (_BasDNS *BasDNSCallerSession) O() (common.Address, error) {
	return _BasDNS.Contract.O(&_BasDNS.CallOpts)
}

// Query is a free data retrieval call binding the contract method 0xa3b61a9c.
//
// Solidity: function query(bytes32 nameHash) constant returns(bytes4 ipv4, bytes16 ipv6, bytes bca, bytes opData, string aliasName)
func (_BasDNS *BasDNSCaller) Query(opts *bind.CallOpts, nameHash [32]byte) (struct {
	Ipv4      [4]byte
	Ipv6      [16]byte
	Bca       []byte
	OpData    []byte
	AliasName string
}, error) {
	ret := new(struct {
		Ipv4      [4]byte
		Ipv6      [16]byte
		Bca       []byte
		OpData    []byte
		AliasName string
	})
	out := ret
	err := _BasDNS.contract.Call(opts, out, "query", nameHash)
	return *ret, err
}

// Query is a free data retrieval call binding the contract method 0xa3b61a9c.
//
// Solidity: function query(bytes32 nameHash) constant returns(bytes4 ipv4, bytes16 ipv6, bytes bca, bytes opData, string aliasName)
func (_BasDNS *BasDNSSession) Query(nameHash [32]byte) (struct {
	Ipv4      [4]byte
	Ipv6      [16]byte
	Bca       []byte
	OpData    []byte
	AliasName string
}, error) {
	return _BasDNS.Contract.Query(&_BasDNS.CallOpts, nameHash)
}

// Query is a free data retrieval call binding the contract method 0xa3b61a9c.
//
// Solidity: function query(bytes32 nameHash) constant returns(bytes4 ipv4, bytes16 ipv6, bytes bca, bytes opData, string aliasName)
func (_BasDNS *BasDNSCallerSession) Query(nameHash [32]byte) (struct {
	Ipv4      [4]byte
	Ipv6      [16]byte
	Bca       []byte
	OpData    []byte
	AliasName string
}, error) {
	return _BasDNS.Contract.Query(&_BasDNS.CallOpts, nameHash)
}

// AUpdate is a paid mutator transaction binding the contract method 0x47beb2fd.
//
// Solidity: function _a_update(bytes32 nameHash, bytes4 ipv4, bytes16 ipv6, bytes bca, bytes opData, string aliasName) returns()
func (_BasDNS *BasDNSTransactor) AUpdate(opts *bind.TransactOpts, nameHash [32]byte, ipv4 [4]byte, ipv6 [16]byte, bca []byte, opData []byte, aliasName string) (*types.Transaction, error) {
	return _BasDNS.contract.Transact(opts, "_a_update", nameHash, ipv4, ipv6, bca, opData, aliasName)
}

// AUpdate is a paid mutator transaction binding the contract method 0x47beb2fd.
//
// Solidity: function _a_update(bytes32 nameHash, bytes4 ipv4, bytes16 ipv6, bytes bca, bytes opData, string aliasName) returns()
func (_BasDNS *BasDNSSession) AUpdate(nameHash [32]byte, ipv4 [4]byte, ipv6 [16]byte, bca []byte, opData []byte, aliasName string) (*types.Transaction, error) {
	return _BasDNS.Contract.AUpdate(&_BasDNS.TransactOpts, nameHash, ipv4, ipv6, bca, opData, aliasName)
}

// AUpdate is a paid mutator transaction binding the contract method 0x47beb2fd.
//
// Solidity: function _a_update(bytes32 nameHash, bytes4 ipv4, bytes16 ipv6, bytes bca, bytes opData, string aliasName) returns()
func (_BasDNS *BasDNSTransactorSession) AUpdate(nameHash [32]byte, ipv4 [4]byte, ipv6 [16]byte, bca []byte, opData []byte, aliasName string) (*types.Transaction, error) {
	return _BasDNS.Contract.AUpdate(&_BasDNS.TransactOpts, nameHash, ipv4, ipv6, bca, opData, aliasName)
}

// CClearRecord is a paid mutator transaction binding the contract method 0x89507a23.
//
// Solidity: function _c_clearRecord(bytes32 nameHash) returns()
func (_BasDNS *BasDNSTransactor) CClearRecord(opts *bind.TransactOpts, nameHash [32]byte) (*types.Transaction, error) {
	return _BasDNS.contract.Transact(opts, "_c_clearRecord", nameHash)
}

// CClearRecord is a paid mutator transaction binding the contract method 0x89507a23.
//
// Solidity: function _c_clearRecord(bytes32 nameHash) returns()
func (_BasDNS *BasDNSSession) CClearRecord(nameHash [32]byte) (*types.Transaction, error) {
	return _BasDNS.Contract.CClearRecord(&_BasDNS.TransactOpts, nameHash)
}

// CClearRecord is a paid mutator transaction binding the contract method 0x89507a23.
//
// Solidity: function _c_clearRecord(bytes32 nameHash) returns()
func (_BasDNS *BasDNSTransactorSession) CClearRecord(nameHash [32]byte) (*types.Transaction, error) {
	return _BasDNS.Contract.CClearRecord(&_BasDNS.TransactOpts, nameHash)
}

// CSetAlias is a paid mutator transaction binding the contract method 0x309ba2c3.
//
// Solidity: function _c_setAlias(bytes32 nameHash, string aName) returns()
func (_BasDNS *BasDNSTransactor) CSetAlias(opts *bind.TransactOpts, nameHash [32]byte, aName string) (*types.Transaction, error) {
	return _BasDNS.contract.Transact(opts, "_c_setAlias", nameHash, aName)
}

// CSetAlias is a paid mutator transaction binding the contract method 0x309ba2c3.
//
// Solidity: function _c_setAlias(bytes32 nameHash, string aName) returns()
func (_BasDNS *BasDNSSession) CSetAlias(nameHash [32]byte, aName string) (*types.Transaction, error) {
	return _BasDNS.Contract.CSetAlias(&_BasDNS.TransactOpts, nameHash, aName)
}

// CSetAlias is a paid mutator transaction binding the contract method 0x309ba2c3.
//
// Solidity: function _c_setAlias(bytes32 nameHash, string aName) returns()
func (_BasDNS *BasDNSTransactorSession) CSetAlias(nameHash [32]byte, aName string) (*types.Transaction, error) {
	return _BasDNS.Contract.CSetAlias(&_BasDNS.TransactOpts, nameHash, aName)
}

// CSetBCAddress is a paid mutator transaction binding the contract method 0x59547fec.
//
// Solidity: function _c_setBCAddress(bytes32 nameHash, bytes bcAddress) returns()
func (_BasDNS *BasDNSTransactor) CSetBCAddress(opts *bind.TransactOpts, nameHash [32]byte, bcAddress []byte) (*types.Transaction, error) {
	return _BasDNS.contract.Transact(opts, "_c_setBCAddress", nameHash, bcAddress)
}

// CSetBCAddress is a paid mutator transaction binding the contract method 0x59547fec.
//
// Solidity: function _c_setBCAddress(bytes32 nameHash, bytes bcAddress) returns()
func (_BasDNS *BasDNSSession) CSetBCAddress(nameHash [32]byte, bcAddress []byte) (*types.Transaction, error) {
	return _BasDNS.Contract.CSetBCAddress(&_BasDNS.TransactOpts, nameHash, bcAddress)
}

// CSetBCAddress is a paid mutator transaction binding the contract method 0x59547fec.
//
// Solidity: function _c_setBCAddress(bytes32 nameHash, bytes bcAddress) returns()
func (_BasDNS *BasDNSTransactorSession) CSetBCAddress(nameHash [32]byte, bcAddress []byte) (*types.Transaction, error) {
	return _BasDNS.Contract.CSetBCAddress(&_BasDNS.TransactOpts, nameHash, bcAddress)
}

// CSetIP is a paid mutator transaction binding the contract method 0xd187e4f6.
//
// Solidity: function _c_setIP(bytes32 nameHash, bytes4 ipv4, bytes16 ipv6) returns()
func (_BasDNS *BasDNSTransactor) CSetIP(opts *bind.TransactOpts, nameHash [32]byte, ipv4 [4]byte, ipv6 [16]byte) (*types.Transaction, error) {
	return _BasDNS.contract.Transact(opts, "_c_setIP", nameHash, ipv4, ipv6)
}

// CSetIP is a paid mutator transaction binding the contract method 0xd187e4f6.
//
// Solidity: function _c_setIP(bytes32 nameHash, bytes4 ipv4, bytes16 ipv6) returns()
func (_BasDNS *BasDNSSession) CSetIP(nameHash [32]byte, ipv4 [4]byte, ipv6 [16]byte) (*types.Transaction, error) {
	return _BasDNS.Contract.CSetIP(&_BasDNS.TransactOpts, nameHash, ipv4, ipv6)
}

// CSetIP is a paid mutator transaction binding the contract method 0xd187e4f6.
//
// Solidity: function _c_setIP(bytes32 nameHash, bytes4 ipv4, bytes16 ipv6) returns()
func (_BasDNS *BasDNSTransactorSession) CSetIP(nameHash [32]byte, ipv4 [4]byte, ipv6 [16]byte) (*types.Transaction, error) {
	return _BasDNS.Contract.CSetIP(&_BasDNS.TransactOpts, nameHash, ipv4, ipv6)
}

// CSetOpData is a paid mutator transaction binding the contract method 0x0883fb43.
//
// Solidity: function _c_setOpData(bytes32 nameHash, bytes opData) returns()
func (_BasDNS *BasDNSTransactor) CSetOpData(opts *bind.TransactOpts, nameHash [32]byte, opData []byte) (*types.Transaction, error) {
	return _BasDNS.contract.Transact(opts, "_c_setOpData", nameHash, opData)
}

// CSetOpData is a paid mutator transaction binding the contract method 0x0883fb43.
//
// Solidity: function _c_setOpData(bytes32 nameHash, bytes opData) returns()
func (_BasDNS *BasDNSSession) CSetOpData(nameHash [32]byte, opData []byte) (*types.Transaction, error) {
	return _BasDNS.Contract.CSetOpData(&_BasDNS.TransactOpts, nameHash, opData)
}

// CSetOpData is a paid mutator transaction binding the contract method 0x0883fb43.
//
// Solidity: function _c_setOpData(bytes32 nameHash, bytes opData) returns()
func (_BasDNS *BasDNSTransactorSession) CSetOpData(nameHash [32]byte, opData []byte) (*types.Transaction, error) {
	return _BasDNS.Contract.CSetOpData(&_BasDNS.TransactOpts, nameHash, opData)
}

// CUpdate is a paid mutator transaction binding the contract method 0x1a004ea2.
//
// Solidity: function _c_update(bytes32 nameHash, bytes4 ipv4, bytes16 ipv6, bytes bca, bytes opData, string aliasName) returns()
func (_BasDNS *BasDNSTransactor) CUpdate(opts *bind.TransactOpts, nameHash [32]byte, ipv4 [4]byte, ipv6 [16]byte, bca []byte, opData []byte, aliasName string) (*types.Transaction, error) {
	return _BasDNS.contract.Transact(opts, "_c_update", nameHash, ipv4, ipv6, bca, opData, aliasName)
}

// CUpdate is a paid mutator transaction binding the contract method 0x1a004ea2.
//
// Solidity: function _c_update(bytes32 nameHash, bytes4 ipv4, bytes16 ipv6, bytes bca, bytes opData, string aliasName) returns()
func (_BasDNS *BasDNSSession) CUpdate(nameHash [32]byte, ipv4 [4]byte, ipv6 [16]byte, bca []byte, opData []byte, aliasName string) (*types.Transaction, error) {
	return _BasDNS.Contract.CUpdate(&_BasDNS.TransactOpts, nameHash, ipv4, ipv6, bca, opData, aliasName)
}

// CUpdate is a paid mutator transaction binding the contract method 0x1a004ea2.
//
// Solidity: function _c_update(bytes32 nameHash, bytes4 ipv4, bytes16 ipv6, bytes bca, bytes opData, string aliasName) returns()
func (_BasDNS *BasDNSTransactorSession) CUpdate(nameHash [32]byte, ipv4 [4]byte, ipv6 [16]byte, bca []byte, opData []byte, aliasName string) (*types.Transaction, error) {
	return _BasDNS.Contract.CUpdate(&_BasDNS.TransactOpts, nameHash, ipv4, ipv6, bca, opData, aliasName)
}

// BasDNSDNSChangedIterator is returned from FilterDNSChanged and is used to iterate over the raw logs and unpacked data for DNSChanged events raised by the BasDNS contract.
type BasDNSDNSChangedIterator struct {
	Event *BasDNSDNSChanged // Event containing the contract specifics and raw log

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
func (it *BasDNSDNSChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasDNSDNSChanged)
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
		it.Event = new(BasDNSDNSChanged)
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
func (it *BasDNSDNSChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasDNSDNSChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasDNSDNSChanged represents a DNSChanged event raised by the BasDNS contract.
type BasDNSDNSChanged struct {
	NameHash [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDNSChanged is a free log retrieval operation binding the contract event 0x03168db5f4e31983b2b5ac149fba9bf46e9ff27416789c624ccf91f4be4057a7.
//
// Solidity: event DNSChanged(bytes32 nameHash)
func (_BasDNS *BasDNSFilterer) FilterDNSChanged(opts *bind.FilterOpts) (*BasDNSDNSChangedIterator, error) {

	logs, sub, err := _BasDNS.contract.FilterLogs(opts, "DNSChanged")
	if err != nil {
		return nil, err
	}
	return &BasDNSDNSChangedIterator{contract: _BasDNS.contract, event: "DNSChanged", logs: logs, sub: sub}, nil
}

// WatchDNSChanged is a free log subscription operation binding the contract event 0x03168db5f4e31983b2b5ac149fba9bf46e9ff27416789c624ccf91f4be4057a7.
//
// Solidity: event DNSChanged(bytes32 nameHash)
func (_BasDNS *BasDNSFilterer) WatchDNSChanged(opts *bind.WatchOpts, sink chan<- *BasDNSDNSChanged) (event.Subscription, error) {

	logs, sub, err := _BasDNS.contract.WatchLogs(opts, "DNSChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasDNSDNSChanged)
				if err := _BasDNS.contract.UnpackLog(event, "DNSChanged", log); err != nil {
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

// ParseDNSChanged is a log parse operation binding the contract event 0x03168db5f4e31983b2b5ac149fba9bf46e9ff27416789c624ccf91f4be4057a7.
//
// Solidity: event DNSChanged(bytes32 nameHash)
func (_BasDNS *BasDNSFilterer) ParseDNSChanged(log types.Log) (*BasDNSDNSChanged, error) {
	event := new(BasDNSDNSChanged)
	if err := _BasDNS.contract.UnpackLog(event, "DNSChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}
