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

// BasMarketABI is the input ABI used to generate the binding from.
const BasMarketABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protectiveRemainTime\",\"type\":\"uint256\"}],\"name\":\"ChangeAsksData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"SellOrders\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"ChangeSellPrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"AT_LEAST_REMAIN_TIME\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protectiveRemainTime\",\"type\":\"uint256\"}],\"name\":\"AddToAsks\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"BuyFromSells\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"AskOrders\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protectiveRemainTime\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"o\",\"outputs\":[{\"internalType\":\"contractBasOwnership\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"AddToSells\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"t\",\"outputs\":[{\"internalType\":\"contractERC20\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"}],\"name\":\"RemoveAskOrder\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"}],\"name\":\"RemoveSellOrder\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"SellToAsks\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_t\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_o\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"SoldBySell\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"SoldByAsk\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"SellAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"SellChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"SellRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"AskAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"AskChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"AskRemoved\",\"type\":\"event\"}]"

// BasMarket is an auto generated Go binding around an Ethereum contract.
type BasMarket struct {
	BasMarketCaller     // Read-only binding to the contract
	BasMarketTransactor // Write-only binding to the contract
	BasMarketFilterer   // Log filterer for contract events
}

// BasMarketCaller is an auto generated read-only Go binding around an Ethereum contract.
type BasMarketCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasMarketTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BasMarketTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasMarketFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BasMarketFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasMarketSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BasMarketSession struct {
	Contract     *BasMarket        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BasMarketCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BasMarketCallerSession struct {
	Contract *BasMarketCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// BasMarketTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BasMarketTransactorSession struct {
	Contract     *BasMarketTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// BasMarketRaw is an auto generated low-level Go binding around an Ethereum contract.
type BasMarketRaw struct {
	Contract *BasMarket // Generic contract binding to access the raw methods on
}

// BasMarketCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BasMarketCallerRaw struct {
	Contract *BasMarketCaller // Generic read-only contract binding to access the raw methods on
}

// BasMarketTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BasMarketTransactorRaw struct {
	Contract *BasMarketTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBasMarket creates a new instance of BasMarket, bound to a specific deployed contract.
func NewBasMarket(address common.Address, backend bind.ContractBackend) (*BasMarket, error) {
	contract, err := bindBasMarket(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BasMarket{BasMarketCaller: BasMarketCaller{contract: contract}, BasMarketTransactor: BasMarketTransactor{contract: contract}, BasMarketFilterer: BasMarketFilterer{contract: contract}}, nil
}

// NewBasMarketCaller creates a new read-only instance of BasMarket, bound to a specific deployed contract.
func NewBasMarketCaller(address common.Address, caller bind.ContractCaller) (*BasMarketCaller, error) {
	contract, err := bindBasMarket(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BasMarketCaller{contract: contract}, nil
}

// NewBasMarketTransactor creates a new write-only instance of BasMarket, bound to a specific deployed contract.
func NewBasMarketTransactor(address common.Address, transactor bind.ContractTransactor) (*BasMarketTransactor, error) {
	contract, err := bindBasMarket(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BasMarketTransactor{contract: contract}, nil
}

// NewBasMarketFilterer creates a new log filterer instance of BasMarket, bound to a specific deployed contract.
func NewBasMarketFilterer(address common.Address, filterer bind.ContractFilterer) (*BasMarketFilterer, error) {
	contract, err := bindBasMarket(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BasMarketFilterer{contract: contract}, nil
}

// bindBasMarket binds a generic wrapper to an already deployed contract.
func bindBasMarket(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BasMarketABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasMarket *BasMarketRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BasMarket.Contract.BasMarketCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasMarket *BasMarketRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasMarket.Contract.BasMarketTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasMarket *BasMarketRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasMarket.Contract.BasMarketTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasMarket *BasMarketCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BasMarket.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasMarket *BasMarketTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasMarket.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasMarket *BasMarketTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasMarket.Contract.contract.Transact(opts, method, params...)
}

// ATLEASTREMAINTIME is a free data retrieval call binding the contract method 0x1d6a1ed0.
//
// Solidity: function AT_LEAST_REMAIN_TIME() constant returns(uint256)
func (_BasMarket *BasMarketCaller) ATLEASTREMAINTIME(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BasMarket.contract.Call(opts, out, "AT_LEAST_REMAIN_TIME")
	return *ret0, err
}

// ATLEASTREMAINTIME is a free data retrieval call binding the contract method 0x1d6a1ed0.
//
// Solidity: function AT_LEAST_REMAIN_TIME() constant returns(uint256)
func (_BasMarket *BasMarketSession) ATLEASTREMAINTIME() (*big.Int, error) {
	return _BasMarket.Contract.ATLEASTREMAINTIME(&_BasMarket.CallOpts)
}

// ATLEASTREMAINTIME is a free data retrieval call binding the contract method 0x1d6a1ed0.
//
// Solidity: function AT_LEAST_REMAIN_TIME() constant returns(uint256)
func (_BasMarket *BasMarketCallerSession) ATLEASTREMAINTIME() (*big.Int, error) {
	return _BasMarket.Contract.ATLEASTREMAINTIME(&_BasMarket.CallOpts)
}

// AskOrders is a free data retrieval call binding the contract method 0x43d4a000.
//
// Solidity: function AskOrders(address , bytes32 ) constant returns(uint256 price, uint256 protectiveRemainTime)
func (_BasMarket *BasMarketCaller) AskOrders(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (struct {
	Price                *big.Int
	ProtectiveRemainTime *big.Int
}, error) {
	ret := new(struct {
		Price                *big.Int
		ProtectiveRemainTime *big.Int
	})
	out := ret
	err := _BasMarket.contract.Call(opts, out, "AskOrders", arg0, arg1)
	return *ret, err
}

// AskOrders is a free data retrieval call binding the contract method 0x43d4a000.
//
// Solidity: function AskOrders(address , bytes32 ) constant returns(uint256 price, uint256 protectiveRemainTime)
func (_BasMarket *BasMarketSession) AskOrders(arg0 common.Address, arg1 [32]byte) (struct {
	Price                *big.Int
	ProtectiveRemainTime *big.Int
}, error) {
	return _BasMarket.Contract.AskOrders(&_BasMarket.CallOpts, arg0, arg1)
}

// AskOrders is a free data retrieval call binding the contract method 0x43d4a000.
//
// Solidity: function AskOrders(address , bytes32 ) constant returns(uint256 price, uint256 protectiveRemainTime)
func (_BasMarket *BasMarketCallerSession) AskOrders(arg0 common.Address, arg1 [32]byte) (struct {
	Price                *big.Int
	ProtectiveRemainTime *big.Int
}, error) {
	return _BasMarket.Contract.AskOrders(&_BasMarket.CallOpts, arg0, arg1)
}

// SellOrders is a free data retrieval call binding the contract method 0x0db8ae09.
//
// Solidity: function SellOrders(address , bytes32 ) constant returns(uint256)
func (_BasMarket *BasMarketCaller) SellOrders(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BasMarket.contract.Call(opts, out, "SellOrders", arg0, arg1)
	return *ret0, err
}

// SellOrders is a free data retrieval call binding the contract method 0x0db8ae09.
//
// Solidity: function SellOrders(address , bytes32 ) constant returns(uint256)
func (_BasMarket *BasMarketSession) SellOrders(arg0 common.Address, arg1 [32]byte) (*big.Int, error) {
	return _BasMarket.Contract.SellOrders(&_BasMarket.CallOpts, arg0, arg1)
}

// SellOrders is a free data retrieval call binding the contract method 0x0db8ae09.
//
// Solidity: function SellOrders(address , bytes32 ) constant returns(uint256)
func (_BasMarket *BasMarketCallerSession) SellOrders(arg0 common.Address, arg1 [32]byte) (*big.Int, error) {
	return _BasMarket.Contract.SellOrders(&_BasMarket.CallOpts, arg0, arg1)
}

// O is a free data retrieval call binding the contract method 0x50cd4df2.
//
// Solidity: function o() constant returns(address)
func (_BasMarket *BasMarketCaller) O(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BasMarket.contract.Call(opts, out, "o")
	return *ret0, err
}

// O is a free data retrieval call binding the contract method 0x50cd4df2.
//
// Solidity: function o() constant returns(address)
func (_BasMarket *BasMarketSession) O() (common.Address, error) {
	return _BasMarket.Contract.O(&_BasMarket.CallOpts)
}

// O is a free data retrieval call binding the contract method 0x50cd4df2.
//
// Solidity: function o() constant returns(address)
func (_BasMarket *BasMarketCallerSession) O() (common.Address, error) {
	return _BasMarket.Contract.O(&_BasMarket.CallOpts)
}

// T is a free data retrieval call binding the contract method 0x92d0d153.
//
// Solidity: function t() constant returns(address)
func (_BasMarket *BasMarketCaller) T(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BasMarket.contract.Call(opts, out, "t")
	return *ret0, err
}

// T is a free data retrieval call binding the contract method 0x92d0d153.
//
// Solidity: function t() constant returns(address)
func (_BasMarket *BasMarketSession) T() (common.Address, error) {
	return _BasMarket.Contract.T(&_BasMarket.CallOpts)
}

// T is a free data retrieval call binding the contract method 0x92d0d153.
//
// Solidity: function t() constant returns(address)
func (_BasMarket *BasMarketCallerSession) T() (common.Address, error) {
	return _BasMarket.Contract.T(&_BasMarket.CallOpts)
}

// AddToAsks is a paid mutator transaction binding the contract method 0x32e99397.
//
// Solidity: function AddToAsks(bytes32 nameHash, uint256 price, uint256 protectiveRemainTime) returns()
func (_BasMarket *BasMarketTransactor) AddToAsks(opts *bind.TransactOpts, nameHash [32]byte, price *big.Int, protectiveRemainTime *big.Int) (*types.Transaction, error) {
	return _BasMarket.contract.Transact(opts, "AddToAsks", nameHash, price, protectiveRemainTime)
}

// AddToAsks is a paid mutator transaction binding the contract method 0x32e99397.
//
// Solidity: function AddToAsks(bytes32 nameHash, uint256 price, uint256 protectiveRemainTime) returns()
func (_BasMarket *BasMarketSession) AddToAsks(nameHash [32]byte, price *big.Int, protectiveRemainTime *big.Int) (*types.Transaction, error) {
	return _BasMarket.Contract.AddToAsks(&_BasMarket.TransactOpts, nameHash, price, protectiveRemainTime)
}

// AddToAsks is a paid mutator transaction binding the contract method 0x32e99397.
//
// Solidity: function AddToAsks(bytes32 nameHash, uint256 price, uint256 protectiveRemainTime) returns()
func (_BasMarket *BasMarketTransactorSession) AddToAsks(nameHash [32]byte, price *big.Int, protectiveRemainTime *big.Int) (*types.Transaction, error) {
	return _BasMarket.Contract.AddToAsks(&_BasMarket.TransactOpts, nameHash, price, protectiveRemainTime)
}

// AddToSells is a paid mutator transaction binding the contract method 0x75a69f19.
//
// Solidity: function AddToSells(bytes32 nameHash, uint256 price) returns()
func (_BasMarket *BasMarketTransactor) AddToSells(opts *bind.TransactOpts, nameHash [32]byte, price *big.Int) (*types.Transaction, error) {
	return _BasMarket.contract.Transact(opts, "AddToSells", nameHash, price)
}

// AddToSells is a paid mutator transaction binding the contract method 0x75a69f19.
//
// Solidity: function AddToSells(bytes32 nameHash, uint256 price) returns()
func (_BasMarket *BasMarketSession) AddToSells(nameHash [32]byte, price *big.Int) (*types.Transaction, error) {
	return _BasMarket.Contract.AddToSells(&_BasMarket.TransactOpts, nameHash, price)
}

// AddToSells is a paid mutator transaction binding the contract method 0x75a69f19.
//
// Solidity: function AddToSells(bytes32 nameHash, uint256 price) returns()
func (_BasMarket *BasMarketTransactorSession) AddToSells(nameHash [32]byte, price *big.Int) (*types.Transaction, error) {
	return _BasMarket.Contract.AddToSells(&_BasMarket.TransactOpts, nameHash, price)
}

// BuyFromSells is a paid mutator transaction binding the contract method 0x3a2ef923.
//
// Solidity: function BuyFromSells(bytes32 nameHash, address owner) returns()
func (_BasMarket *BasMarketTransactor) BuyFromSells(opts *bind.TransactOpts, nameHash [32]byte, owner common.Address) (*types.Transaction, error) {
	return _BasMarket.contract.Transact(opts, "BuyFromSells", nameHash, owner)
}

// BuyFromSells is a paid mutator transaction binding the contract method 0x3a2ef923.
//
// Solidity: function BuyFromSells(bytes32 nameHash, address owner) returns()
func (_BasMarket *BasMarketSession) BuyFromSells(nameHash [32]byte, owner common.Address) (*types.Transaction, error) {
	return _BasMarket.Contract.BuyFromSells(&_BasMarket.TransactOpts, nameHash, owner)
}

// BuyFromSells is a paid mutator transaction binding the contract method 0x3a2ef923.
//
// Solidity: function BuyFromSells(bytes32 nameHash, address owner) returns()
func (_BasMarket *BasMarketTransactorSession) BuyFromSells(nameHash [32]byte, owner common.Address) (*types.Transaction, error) {
	return _BasMarket.Contract.BuyFromSells(&_BasMarket.TransactOpts, nameHash, owner)
}

// ChangeAsksData is a paid mutator transaction binding the contract method 0x01f922a4.
//
// Solidity: function ChangeAsksData(bytes32 nameHash, uint256 price, uint256 protectiveRemainTime) returns()
func (_BasMarket *BasMarketTransactor) ChangeAsksData(opts *bind.TransactOpts, nameHash [32]byte, price *big.Int, protectiveRemainTime *big.Int) (*types.Transaction, error) {
	return _BasMarket.contract.Transact(opts, "ChangeAsksData", nameHash, price, protectiveRemainTime)
}

// ChangeAsksData is a paid mutator transaction binding the contract method 0x01f922a4.
//
// Solidity: function ChangeAsksData(bytes32 nameHash, uint256 price, uint256 protectiveRemainTime) returns()
func (_BasMarket *BasMarketSession) ChangeAsksData(nameHash [32]byte, price *big.Int, protectiveRemainTime *big.Int) (*types.Transaction, error) {
	return _BasMarket.Contract.ChangeAsksData(&_BasMarket.TransactOpts, nameHash, price, protectiveRemainTime)
}

// ChangeAsksData is a paid mutator transaction binding the contract method 0x01f922a4.
//
// Solidity: function ChangeAsksData(bytes32 nameHash, uint256 price, uint256 protectiveRemainTime) returns()
func (_BasMarket *BasMarketTransactorSession) ChangeAsksData(nameHash [32]byte, price *big.Int, protectiveRemainTime *big.Int) (*types.Transaction, error) {
	return _BasMarket.Contract.ChangeAsksData(&_BasMarket.TransactOpts, nameHash, price, protectiveRemainTime)
}

// ChangeSellPrice is a paid mutator transaction binding the contract method 0x1d4edd20.
//
// Solidity: function ChangeSellPrice(bytes32 nameHash, uint256 price) returns()
func (_BasMarket *BasMarketTransactor) ChangeSellPrice(opts *bind.TransactOpts, nameHash [32]byte, price *big.Int) (*types.Transaction, error) {
	return _BasMarket.contract.Transact(opts, "ChangeSellPrice", nameHash, price)
}

// ChangeSellPrice is a paid mutator transaction binding the contract method 0x1d4edd20.
//
// Solidity: function ChangeSellPrice(bytes32 nameHash, uint256 price) returns()
func (_BasMarket *BasMarketSession) ChangeSellPrice(nameHash [32]byte, price *big.Int) (*types.Transaction, error) {
	return _BasMarket.Contract.ChangeSellPrice(&_BasMarket.TransactOpts, nameHash, price)
}

// ChangeSellPrice is a paid mutator transaction binding the contract method 0x1d4edd20.
//
// Solidity: function ChangeSellPrice(bytes32 nameHash, uint256 price) returns()
func (_BasMarket *BasMarketTransactorSession) ChangeSellPrice(nameHash [32]byte, price *big.Int) (*types.Transaction, error) {
	return _BasMarket.Contract.ChangeSellPrice(&_BasMarket.TransactOpts, nameHash, price)
}

// RemoveAskOrder is a paid mutator transaction binding the contract method 0xa86bec4c.
//
// Solidity: function RemoveAskOrder(bytes32 nameHash) returns()
func (_BasMarket *BasMarketTransactor) RemoveAskOrder(opts *bind.TransactOpts, nameHash [32]byte) (*types.Transaction, error) {
	return _BasMarket.contract.Transact(opts, "RemoveAskOrder", nameHash)
}

// RemoveAskOrder is a paid mutator transaction binding the contract method 0xa86bec4c.
//
// Solidity: function RemoveAskOrder(bytes32 nameHash) returns()
func (_BasMarket *BasMarketSession) RemoveAskOrder(nameHash [32]byte) (*types.Transaction, error) {
	return _BasMarket.Contract.RemoveAskOrder(&_BasMarket.TransactOpts, nameHash)
}

// RemoveAskOrder is a paid mutator transaction binding the contract method 0xa86bec4c.
//
// Solidity: function RemoveAskOrder(bytes32 nameHash) returns()
func (_BasMarket *BasMarketTransactorSession) RemoveAskOrder(nameHash [32]byte) (*types.Transaction, error) {
	return _BasMarket.Contract.RemoveAskOrder(&_BasMarket.TransactOpts, nameHash)
}

// RemoveSellOrder is a paid mutator transaction binding the contract method 0xe53427ca.
//
// Solidity: function RemoveSellOrder(bytes32 nameHash) returns()
func (_BasMarket *BasMarketTransactor) RemoveSellOrder(opts *bind.TransactOpts, nameHash [32]byte) (*types.Transaction, error) {
	return _BasMarket.contract.Transact(opts, "RemoveSellOrder", nameHash)
}

// RemoveSellOrder is a paid mutator transaction binding the contract method 0xe53427ca.
//
// Solidity: function RemoveSellOrder(bytes32 nameHash) returns()
func (_BasMarket *BasMarketSession) RemoveSellOrder(nameHash [32]byte) (*types.Transaction, error) {
	return _BasMarket.Contract.RemoveSellOrder(&_BasMarket.TransactOpts, nameHash)
}

// RemoveSellOrder is a paid mutator transaction binding the contract method 0xe53427ca.
//
// Solidity: function RemoveSellOrder(bytes32 nameHash) returns()
func (_BasMarket *BasMarketTransactorSession) RemoveSellOrder(nameHash [32]byte) (*types.Transaction, error) {
	return _BasMarket.Contract.RemoveSellOrder(&_BasMarket.TransactOpts, nameHash)
}

// SellToAsks is a paid mutator transaction binding the contract method 0xebfaafda.
//
// Solidity: function SellToAsks(bytes32 nameHash, address to) returns()
func (_BasMarket *BasMarketTransactor) SellToAsks(opts *bind.TransactOpts, nameHash [32]byte, to common.Address) (*types.Transaction, error) {
	return _BasMarket.contract.Transact(opts, "SellToAsks", nameHash, to)
}

// SellToAsks is a paid mutator transaction binding the contract method 0xebfaafda.
//
// Solidity: function SellToAsks(bytes32 nameHash, address to) returns()
func (_BasMarket *BasMarketSession) SellToAsks(nameHash [32]byte, to common.Address) (*types.Transaction, error) {
	return _BasMarket.Contract.SellToAsks(&_BasMarket.TransactOpts, nameHash, to)
}

// SellToAsks is a paid mutator transaction binding the contract method 0xebfaafda.
//
// Solidity: function SellToAsks(bytes32 nameHash, address to) returns()
func (_BasMarket *BasMarketTransactorSession) SellToAsks(nameHash [32]byte, to common.Address) (*types.Transaction, error) {
	return _BasMarket.Contract.SellToAsks(&_BasMarket.TransactOpts, nameHash, to)
}

// BasMarketAskAddedIterator is returned from FilterAskAdded and is used to iterate over the raw logs and unpacked data for AskAdded events raised by the BasMarket contract.
type BasMarketAskAddedIterator struct {
	Event *BasMarketAskAdded // Event containing the contract specifics and raw log

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
func (it *BasMarketAskAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasMarketAskAdded)
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
		it.Event = new(BasMarketAskAdded)
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
func (it *BasMarketAskAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasMarketAskAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasMarketAskAdded represents a AskAdded event raised by the BasMarket contract.
type BasMarketAskAdded struct {
	NameHash [32]byte
	Operator common.Address
	Price    *big.Int
	Time     *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAskAdded is a free log retrieval operation binding the contract event 0xeabc52f62acd1d372f720cd90e895e31d674f029df61e3cb880f5641ae32916f.
//
// Solidity: event AskAdded(bytes32 nameHash, address operator, uint256 price, uint256 time)
func (_BasMarket *BasMarketFilterer) FilterAskAdded(opts *bind.FilterOpts) (*BasMarketAskAddedIterator, error) {

	logs, sub, err := _BasMarket.contract.FilterLogs(opts, "AskAdded")
	if err != nil {
		return nil, err
	}
	return &BasMarketAskAddedIterator{contract: _BasMarket.contract, event: "AskAdded", logs: logs, sub: sub}, nil
}

// WatchAskAdded is a free log subscription operation binding the contract event 0xeabc52f62acd1d372f720cd90e895e31d674f029df61e3cb880f5641ae32916f.
//
// Solidity: event AskAdded(bytes32 nameHash, address operator, uint256 price, uint256 time)
func (_BasMarket *BasMarketFilterer) WatchAskAdded(opts *bind.WatchOpts, sink chan<- *BasMarketAskAdded) (event.Subscription, error) {

	logs, sub, err := _BasMarket.contract.WatchLogs(opts, "AskAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasMarketAskAdded)
				if err := _BasMarket.contract.UnpackLog(event, "AskAdded", log); err != nil {
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

// ParseAskAdded is a log parse operation binding the contract event 0xeabc52f62acd1d372f720cd90e895e31d674f029df61e3cb880f5641ae32916f.
//
// Solidity: event AskAdded(bytes32 nameHash, address operator, uint256 price, uint256 time)
func (_BasMarket *BasMarketFilterer) ParseAskAdded(log types.Log) (*BasMarketAskAdded, error) {
	event := new(BasMarketAskAdded)
	if err := _BasMarket.contract.UnpackLog(event, "AskAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BasMarketAskChangedIterator is returned from FilterAskChanged and is used to iterate over the raw logs and unpacked data for AskChanged events raised by the BasMarket contract.
type BasMarketAskChangedIterator struct {
	Event *BasMarketAskChanged // Event containing the contract specifics and raw log

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
func (it *BasMarketAskChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasMarketAskChanged)
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
		it.Event = new(BasMarketAskChanged)
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
func (it *BasMarketAskChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasMarketAskChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasMarketAskChanged represents a AskChanged event raised by the BasMarket contract.
type BasMarketAskChanged struct {
	NameHash [32]byte
	Operator common.Address
	Price    *big.Int
	Time     *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAskChanged is a free log retrieval operation binding the contract event 0x745fa0d1805a08074546ad265a3e7a3e1d8799287dcb35d3947370ba469eb4e0.
//
// Solidity: event AskChanged(bytes32 nameHash, address operator, uint256 price, uint256 time)
func (_BasMarket *BasMarketFilterer) FilterAskChanged(opts *bind.FilterOpts) (*BasMarketAskChangedIterator, error) {

	logs, sub, err := _BasMarket.contract.FilterLogs(opts, "AskChanged")
	if err != nil {
		return nil, err
	}
	return &BasMarketAskChangedIterator{contract: _BasMarket.contract, event: "AskChanged", logs: logs, sub: sub}, nil
}

// WatchAskChanged is a free log subscription operation binding the contract event 0x745fa0d1805a08074546ad265a3e7a3e1d8799287dcb35d3947370ba469eb4e0.
//
// Solidity: event AskChanged(bytes32 nameHash, address operator, uint256 price, uint256 time)
func (_BasMarket *BasMarketFilterer) WatchAskChanged(opts *bind.WatchOpts, sink chan<- *BasMarketAskChanged) (event.Subscription, error) {

	logs, sub, err := _BasMarket.contract.WatchLogs(opts, "AskChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasMarketAskChanged)
				if err := _BasMarket.contract.UnpackLog(event, "AskChanged", log); err != nil {
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

// ParseAskChanged is a log parse operation binding the contract event 0x745fa0d1805a08074546ad265a3e7a3e1d8799287dcb35d3947370ba469eb4e0.
//
// Solidity: event AskChanged(bytes32 nameHash, address operator, uint256 price, uint256 time)
func (_BasMarket *BasMarketFilterer) ParseAskChanged(log types.Log) (*BasMarketAskChanged, error) {
	event := new(BasMarketAskChanged)
	if err := _BasMarket.contract.UnpackLog(event, "AskChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BasMarketAskRemovedIterator is returned from FilterAskRemoved and is used to iterate over the raw logs and unpacked data for AskRemoved events raised by the BasMarket contract.
type BasMarketAskRemovedIterator struct {
	Event *BasMarketAskRemoved // Event containing the contract specifics and raw log

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
func (it *BasMarketAskRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasMarketAskRemoved)
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
		it.Event = new(BasMarketAskRemoved)
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
func (it *BasMarketAskRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasMarketAskRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasMarketAskRemoved represents a AskRemoved event raised by the BasMarket contract.
type BasMarketAskRemoved struct {
	NameHash [32]byte
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAskRemoved is a free log retrieval operation binding the contract event 0xb715f983de08cb4ac1b5dd69df079431cfcf98fa642f38a0b9b266de6db4a7cc.
//
// Solidity: event AskRemoved(bytes32 nameHash, address operator)
func (_BasMarket *BasMarketFilterer) FilterAskRemoved(opts *bind.FilterOpts) (*BasMarketAskRemovedIterator, error) {

	logs, sub, err := _BasMarket.contract.FilterLogs(opts, "AskRemoved")
	if err != nil {
		return nil, err
	}
	return &BasMarketAskRemovedIterator{contract: _BasMarket.contract, event: "AskRemoved", logs: logs, sub: sub}, nil
}

// WatchAskRemoved is a free log subscription operation binding the contract event 0xb715f983de08cb4ac1b5dd69df079431cfcf98fa642f38a0b9b266de6db4a7cc.
//
// Solidity: event AskRemoved(bytes32 nameHash, address operator)
func (_BasMarket *BasMarketFilterer) WatchAskRemoved(opts *bind.WatchOpts, sink chan<- *BasMarketAskRemoved) (event.Subscription, error) {

	logs, sub, err := _BasMarket.contract.WatchLogs(opts, "AskRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasMarketAskRemoved)
				if err := _BasMarket.contract.UnpackLog(event, "AskRemoved", log); err != nil {
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

// ParseAskRemoved is a log parse operation binding the contract event 0xb715f983de08cb4ac1b5dd69df079431cfcf98fa642f38a0b9b266de6db4a7cc.
//
// Solidity: event AskRemoved(bytes32 nameHash, address operator)
func (_BasMarket *BasMarketFilterer) ParseAskRemoved(log types.Log) (*BasMarketAskRemoved, error) {
	event := new(BasMarketAskRemoved)
	if err := _BasMarket.contract.UnpackLog(event, "AskRemoved", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BasMarketSellAddedIterator is returned from FilterSellAdded and is used to iterate over the raw logs and unpacked data for SellAdded events raised by the BasMarket contract.
type BasMarketSellAddedIterator struct {
	Event *BasMarketSellAdded // Event containing the contract specifics and raw log

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
func (it *BasMarketSellAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasMarketSellAdded)
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
		it.Event = new(BasMarketSellAdded)
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
func (it *BasMarketSellAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasMarketSellAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasMarketSellAdded represents a SellAdded event raised by the BasMarket contract.
type BasMarketSellAdded struct {
	NameHash [32]byte
	Operator common.Address
	Price    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSellAdded is a free log retrieval operation binding the contract event 0x10f7d9b00b30f2c94496f80f416e54d140fcd3c6040b5440ae55c39a3294086c.
//
// Solidity: event SellAdded(bytes32 nameHash, address operator, uint256 price)
func (_BasMarket *BasMarketFilterer) FilterSellAdded(opts *bind.FilterOpts) (*BasMarketSellAddedIterator, error) {

	logs, sub, err := _BasMarket.contract.FilterLogs(opts, "SellAdded")
	if err != nil {
		return nil, err
	}
	return &BasMarketSellAddedIterator{contract: _BasMarket.contract, event: "SellAdded", logs: logs, sub: sub}, nil
}

// WatchSellAdded is a free log subscription operation binding the contract event 0x10f7d9b00b30f2c94496f80f416e54d140fcd3c6040b5440ae55c39a3294086c.
//
// Solidity: event SellAdded(bytes32 nameHash, address operator, uint256 price)
func (_BasMarket *BasMarketFilterer) WatchSellAdded(opts *bind.WatchOpts, sink chan<- *BasMarketSellAdded) (event.Subscription, error) {

	logs, sub, err := _BasMarket.contract.WatchLogs(opts, "SellAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasMarketSellAdded)
				if err := _BasMarket.contract.UnpackLog(event, "SellAdded", log); err != nil {
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

// ParseSellAdded is a log parse operation binding the contract event 0x10f7d9b00b30f2c94496f80f416e54d140fcd3c6040b5440ae55c39a3294086c.
//
// Solidity: event SellAdded(bytes32 nameHash, address operator, uint256 price)
func (_BasMarket *BasMarketFilterer) ParseSellAdded(log types.Log) (*BasMarketSellAdded, error) {
	event := new(BasMarketSellAdded)
	if err := _BasMarket.contract.UnpackLog(event, "SellAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BasMarketSellChangedIterator is returned from FilterSellChanged and is used to iterate over the raw logs and unpacked data for SellChanged events raised by the BasMarket contract.
type BasMarketSellChangedIterator struct {
	Event *BasMarketSellChanged // Event containing the contract specifics and raw log

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
func (it *BasMarketSellChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasMarketSellChanged)
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
		it.Event = new(BasMarketSellChanged)
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
func (it *BasMarketSellChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasMarketSellChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasMarketSellChanged represents a SellChanged event raised by the BasMarket contract.
type BasMarketSellChanged struct {
	NameHash [32]byte
	Operator common.Address
	Price    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSellChanged is a free log retrieval operation binding the contract event 0xe2fbe11a69dd3397016a8ac5cc0d982dd57cc1586c3400f8e3a3167d0dcd3634.
//
// Solidity: event SellChanged(bytes32 nameHash, address operator, uint256 price)
func (_BasMarket *BasMarketFilterer) FilterSellChanged(opts *bind.FilterOpts) (*BasMarketSellChangedIterator, error) {

	logs, sub, err := _BasMarket.contract.FilterLogs(opts, "SellChanged")
	if err != nil {
		return nil, err
	}
	return &BasMarketSellChangedIterator{contract: _BasMarket.contract, event: "SellChanged", logs: logs, sub: sub}, nil
}

// WatchSellChanged is a free log subscription operation binding the contract event 0xe2fbe11a69dd3397016a8ac5cc0d982dd57cc1586c3400f8e3a3167d0dcd3634.
//
// Solidity: event SellChanged(bytes32 nameHash, address operator, uint256 price)
func (_BasMarket *BasMarketFilterer) WatchSellChanged(opts *bind.WatchOpts, sink chan<- *BasMarketSellChanged) (event.Subscription, error) {

	logs, sub, err := _BasMarket.contract.WatchLogs(opts, "SellChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasMarketSellChanged)
				if err := _BasMarket.contract.UnpackLog(event, "SellChanged", log); err != nil {
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

// ParseSellChanged is a log parse operation binding the contract event 0xe2fbe11a69dd3397016a8ac5cc0d982dd57cc1586c3400f8e3a3167d0dcd3634.
//
// Solidity: event SellChanged(bytes32 nameHash, address operator, uint256 price)
func (_BasMarket *BasMarketFilterer) ParseSellChanged(log types.Log) (*BasMarketSellChanged, error) {
	event := new(BasMarketSellChanged)
	if err := _BasMarket.contract.UnpackLog(event, "SellChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BasMarketSellRemovedIterator is returned from FilterSellRemoved and is used to iterate over the raw logs and unpacked data for SellRemoved events raised by the BasMarket contract.
type BasMarketSellRemovedIterator struct {
	Event *BasMarketSellRemoved // Event containing the contract specifics and raw log

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
func (it *BasMarketSellRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasMarketSellRemoved)
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
		it.Event = new(BasMarketSellRemoved)
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
func (it *BasMarketSellRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasMarketSellRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasMarketSellRemoved represents a SellRemoved event raised by the BasMarket contract.
type BasMarketSellRemoved struct {
	NameHash [32]byte
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSellRemoved is a free log retrieval operation binding the contract event 0xd369082d239e5af317337ed91a37711a5809487a596e485189541c3219a35a0a.
//
// Solidity: event SellRemoved(bytes32 nameHash, address operator)
func (_BasMarket *BasMarketFilterer) FilterSellRemoved(opts *bind.FilterOpts) (*BasMarketSellRemovedIterator, error) {

	logs, sub, err := _BasMarket.contract.FilterLogs(opts, "SellRemoved")
	if err != nil {
		return nil, err
	}
	return &BasMarketSellRemovedIterator{contract: _BasMarket.contract, event: "SellRemoved", logs: logs, sub: sub}, nil
}

// WatchSellRemoved is a free log subscription operation binding the contract event 0xd369082d239e5af317337ed91a37711a5809487a596e485189541c3219a35a0a.
//
// Solidity: event SellRemoved(bytes32 nameHash, address operator)
func (_BasMarket *BasMarketFilterer) WatchSellRemoved(opts *bind.WatchOpts, sink chan<- *BasMarketSellRemoved) (event.Subscription, error) {

	logs, sub, err := _BasMarket.contract.WatchLogs(opts, "SellRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasMarketSellRemoved)
				if err := _BasMarket.contract.UnpackLog(event, "SellRemoved", log); err != nil {
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

// ParseSellRemoved is a log parse operation binding the contract event 0xd369082d239e5af317337ed91a37711a5809487a596e485189541c3219a35a0a.
//
// Solidity: event SellRemoved(bytes32 nameHash, address operator)
func (_BasMarket *BasMarketFilterer) ParseSellRemoved(log types.Log) (*BasMarketSellRemoved, error) {
	event := new(BasMarketSellRemoved)
	if err := _BasMarket.contract.UnpackLog(event, "SellRemoved", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BasMarketSoldByAskIterator is returned from FilterSoldByAsk and is used to iterate over the raw logs and unpacked data for SoldByAsk events raised by the BasMarket contract.
type BasMarketSoldByAskIterator struct {
	Event *BasMarketSoldByAsk // Event containing the contract specifics and raw log

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
func (it *BasMarketSoldByAskIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasMarketSoldByAsk)
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
		it.Event = new(BasMarketSoldByAsk)
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
func (it *BasMarketSoldByAskIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasMarketSoldByAskIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasMarketSoldByAsk represents a SoldByAsk event raised by the BasMarket contract.
type BasMarketSoldByAsk struct {
	NameHash [32]byte
	From     common.Address
	To       common.Address
	Price    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSoldByAsk is a free log retrieval operation binding the contract event 0xf4a9277c75325361938cc6d0e4f286721bcd4b3f14e95104242d23a7e953172e.
//
// Solidity: event SoldByAsk(bytes32 nameHash, address from, address to, uint256 price)
func (_BasMarket *BasMarketFilterer) FilterSoldByAsk(opts *bind.FilterOpts) (*BasMarketSoldByAskIterator, error) {

	logs, sub, err := _BasMarket.contract.FilterLogs(opts, "SoldByAsk")
	if err != nil {
		return nil, err
	}
	return &BasMarketSoldByAskIterator{contract: _BasMarket.contract, event: "SoldByAsk", logs: logs, sub: sub}, nil
}

// WatchSoldByAsk is a free log subscription operation binding the contract event 0xf4a9277c75325361938cc6d0e4f286721bcd4b3f14e95104242d23a7e953172e.
//
// Solidity: event SoldByAsk(bytes32 nameHash, address from, address to, uint256 price)
func (_BasMarket *BasMarketFilterer) WatchSoldByAsk(opts *bind.WatchOpts, sink chan<- *BasMarketSoldByAsk) (event.Subscription, error) {

	logs, sub, err := _BasMarket.contract.WatchLogs(opts, "SoldByAsk")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasMarketSoldByAsk)
				if err := _BasMarket.contract.UnpackLog(event, "SoldByAsk", log); err != nil {
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

// ParseSoldByAsk is a log parse operation binding the contract event 0xf4a9277c75325361938cc6d0e4f286721bcd4b3f14e95104242d23a7e953172e.
//
// Solidity: event SoldByAsk(bytes32 nameHash, address from, address to, uint256 price)
func (_BasMarket *BasMarketFilterer) ParseSoldByAsk(log types.Log) (*BasMarketSoldByAsk, error) {
	event := new(BasMarketSoldByAsk)
	if err := _BasMarket.contract.UnpackLog(event, "SoldByAsk", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BasMarketSoldBySellIterator is returned from FilterSoldBySell and is used to iterate over the raw logs and unpacked data for SoldBySell events raised by the BasMarket contract.
type BasMarketSoldBySellIterator struct {
	Event *BasMarketSoldBySell // Event containing the contract specifics and raw log

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
func (it *BasMarketSoldBySellIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasMarketSoldBySell)
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
		it.Event = new(BasMarketSoldBySell)
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
func (it *BasMarketSoldBySellIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasMarketSoldBySellIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasMarketSoldBySell represents a SoldBySell event raised by the BasMarket contract.
type BasMarketSoldBySell struct {
	NameHash [32]byte
	From     common.Address
	To       common.Address
	Price    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSoldBySell is a free log retrieval operation binding the contract event 0x9c0f845844791139cd02a1448f071f31346e7d1fee8cafc98925d507dc4d6244.
//
// Solidity: event SoldBySell(bytes32 nameHash, address from, address to, uint256 price)
func (_BasMarket *BasMarketFilterer) FilterSoldBySell(opts *bind.FilterOpts) (*BasMarketSoldBySellIterator, error) {

	logs, sub, err := _BasMarket.contract.FilterLogs(opts, "SoldBySell")
	if err != nil {
		return nil, err
	}
	return &BasMarketSoldBySellIterator{contract: _BasMarket.contract, event: "SoldBySell", logs: logs, sub: sub}, nil
}

// WatchSoldBySell is a free log subscription operation binding the contract event 0x9c0f845844791139cd02a1448f071f31346e7d1fee8cafc98925d507dc4d6244.
//
// Solidity: event SoldBySell(bytes32 nameHash, address from, address to, uint256 price)
func (_BasMarket *BasMarketFilterer) WatchSoldBySell(opts *bind.WatchOpts, sink chan<- *BasMarketSoldBySell) (event.Subscription, error) {

	logs, sub, err := _BasMarket.contract.WatchLogs(opts, "SoldBySell")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasMarketSoldBySell)
				if err := _BasMarket.contract.UnpackLog(event, "SoldBySell", log); err != nil {
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

// ParseSoldBySell is a log parse operation binding the contract event 0x9c0f845844791139cd02a1448f071f31346e7d1fee8cafc98925d507dc4d6244.
//
// Solidity: event SoldBySell(bytes32 nameHash, address from, address to, uint256 price)
func (_BasMarket *BasMarketFilterer) ParseSoldBySell(log types.Log) (*BasMarketSoldBySell, error) {
	event := new(BasMarketSoldBySell)
	if err := _BasMarket.contract.UnpackLog(event, "SoldBySell", log); err != nil {
		return nil, err
	}
	return event, nil
}
