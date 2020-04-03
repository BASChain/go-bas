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

// BasMinerABI is the input ABI used to generate the binding from.
const BasMinerABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"_a_changeAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"admin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"burn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"miner\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"}],\"name\":\"_a_changeDefaultSubSetting\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"GetAllMainNodeAddress\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"customedSubSetting\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"toAdmin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toBurn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toMiner\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toRootOwner\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"admin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"burn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"miner\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"}],\"name\":\"_a_changeSelfSubSetting\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"AllocationRoot\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"admin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"burn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"miner\",\"type\":\"uint256\"}],\"name\":\"_a_changeRootSetting\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"AllocationCustomedSub\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"OANNAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contractCaller\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"oldM\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newM\",\"type\":\"address\"}],\"name\":\"_a_replaceMiner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"MainNode\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"AllocationSelfSub\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"allocateType\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"receiptNumber\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"rootOwner\",\"type\":\"address\"}],\"name\":\"_c_allocateProfit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"conAddr\",\"type\":\"address\"}],\"name\":\"_a_changeContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rootSetting\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"toAdmin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toBurn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toMiner\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toRootOwner\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"no\",\"type\":\"uint256\"}],\"name\":\"_a_emergencyWithdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"m\",\"type\":\"address\"}],\"name\":\"_a_addMiner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"}],\"name\":\"_a_removeMiner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"Satoshi_Nakamoto\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"admin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"burn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"miner\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"}],\"name\":\"_a_changeCustomedSubSetting\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"defaultSubSetting\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"toAdmin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toBurn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toMiner\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toRootOwner\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"selfSubSetting\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"toAdmin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toBurn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toMiner\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toRootOwner\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MainNodeSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"AllocationSub\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractBasToken\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"team\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"receiptNumber\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amout\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"allocation\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"Receipt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"drawer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amout\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"allocateType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toAdmin\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toBurn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toMiner\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toRoot\",\"type\":\"uint256\"}],\"name\":\"AllocationChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"}],\"name\":\"MinerAdd\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"}],\"name\":\"MinerRemove\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldMiner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newMiner\",\"type\":\"address\"}],\"name\":\"MinerReplace\",\"type\":\"event\"}]"

// BasMiner is an auto generated Go binding around an Ethereum contract.
type BasMiner struct {
	BasMinerCaller     // Read-only binding to the contract
	BasMinerTransactor // Write-only binding to the contract
	BasMinerFilterer   // Log filterer for contract events
}

// BasMinerCaller is an auto generated read-only Go binding around an Ethereum contract.
type BasMinerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasMinerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BasMinerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasMinerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BasMinerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasMinerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BasMinerSession struct {
	Contract     *BasMiner         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BasMinerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BasMinerCallerSession struct {
	Contract *BasMinerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BasMinerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BasMinerTransactorSession struct {
	Contract     *BasMinerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BasMinerRaw is an auto generated low-level Go binding around an Ethereum contract.
type BasMinerRaw struct {
	Contract *BasMiner // Generic contract binding to access the raw methods on
}

// BasMinerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BasMinerCallerRaw struct {
	Contract *BasMinerCaller // Generic read-only contract binding to access the raw methods on
}

// BasMinerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BasMinerTransactorRaw struct {
	Contract *BasMinerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBasMiner creates a new instance of BasMiner, bound to a specific deployed contract.
func NewBasMiner(address common.Address, backend bind.ContractBackend) (*BasMiner, error) {
	contract, err := bindBasMiner(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BasMiner{BasMinerCaller: BasMinerCaller{contract: contract}, BasMinerTransactor: BasMinerTransactor{contract: contract}, BasMinerFilterer: BasMinerFilterer{contract: contract}}, nil
}

// NewBasMinerCaller creates a new read-only instance of BasMiner, bound to a specific deployed contract.
func NewBasMinerCaller(address common.Address, caller bind.ContractCaller) (*BasMinerCaller, error) {
	contract, err := bindBasMiner(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BasMinerCaller{contract: contract}, nil
}

// NewBasMinerTransactor creates a new write-only instance of BasMiner, bound to a specific deployed contract.
func NewBasMinerTransactor(address common.Address, transactor bind.ContractTransactor) (*BasMinerTransactor, error) {
	contract, err := bindBasMiner(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BasMinerTransactor{contract: contract}, nil
}

// NewBasMinerFilterer creates a new log filterer instance of BasMiner, bound to a specific deployed contract.
func NewBasMinerFilterer(address common.Address, filterer bind.ContractFilterer) (*BasMinerFilterer, error) {
	contract, err := bindBasMiner(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BasMinerFilterer{contract: contract}, nil
}

// bindBasMiner binds a generic wrapper to an already deployed contract.
func bindBasMiner(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BasMinerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasMiner *BasMinerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BasMiner.Contract.BasMinerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasMiner *BasMinerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasMiner.Contract.BasMinerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasMiner *BasMinerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasMiner.Contract.BasMinerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasMiner *BasMinerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BasMiner.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasMiner *BasMinerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasMiner.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasMiner *BasMinerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasMiner.Contract.contract.Transact(opts, method, params...)
}

// AllocationCustomedSub is a free data retrieval call binding the contract method 0x4cd9e136.
//
// Solidity: function AllocationCustomedSub() constant returns(uint8)
func (_BasMiner *BasMinerCaller) AllocationCustomedSub(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _BasMiner.contract.Call(opts, out, "AllocationCustomedSub")
	return *ret0, err
}

// AllocationCustomedSub is a free data retrieval call binding the contract method 0x4cd9e136.
//
// Solidity: function AllocationCustomedSub() constant returns(uint8)
func (_BasMiner *BasMinerSession) AllocationCustomedSub() (uint8, error) {
	return _BasMiner.Contract.AllocationCustomedSub(&_BasMiner.CallOpts)
}

// AllocationCustomedSub is a free data retrieval call binding the contract method 0x4cd9e136.
//
// Solidity: function AllocationCustomedSub() constant returns(uint8)
func (_BasMiner *BasMinerCallerSession) AllocationCustomedSub() (uint8, error) {
	return _BasMiner.Contract.AllocationCustomedSub(&_BasMiner.CallOpts)
}

// AllocationRoot is a free data retrieval call binding the contract method 0x41e9a9fe.
//
// Solidity: function AllocationRoot() constant returns(uint8)
func (_BasMiner *BasMinerCaller) AllocationRoot(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _BasMiner.contract.Call(opts, out, "AllocationRoot")
	return *ret0, err
}

// AllocationRoot is a free data retrieval call binding the contract method 0x41e9a9fe.
//
// Solidity: function AllocationRoot() constant returns(uint8)
func (_BasMiner *BasMinerSession) AllocationRoot() (uint8, error) {
	return _BasMiner.Contract.AllocationRoot(&_BasMiner.CallOpts)
}

// AllocationRoot is a free data retrieval call binding the contract method 0x41e9a9fe.
//
// Solidity: function AllocationRoot() constant returns(uint8)
func (_BasMiner *BasMinerCallerSession) AllocationRoot() (uint8, error) {
	return _BasMiner.Contract.AllocationRoot(&_BasMiner.CallOpts)
}

// AllocationSelfSub is a free data retrieval call binding the contract method 0x63831e0e.
//
// Solidity: function AllocationSelfSub() constant returns(uint8)
func (_BasMiner *BasMinerCaller) AllocationSelfSub(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _BasMiner.contract.Call(opts, out, "AllocationSelfSub")
	return *ret0, err
}

// AllocationSelfSub is a free data retrieval call binding the contract method 0x63831e0e.
//
// Solidity: function AllocationSelfSub() constant returns(uint8)
func (_BasMiner *BasMinerSession) AllocationSelfSub() (uint8, error) {
	return _BasMiner.Contract.AllocationSelfSub(&_BasMiner.CallOpts)
}

// AllocationSelfSub is a free data retrieval call binding the contract method 0x63831e0e.
//
// Solidity: function AllocationSelfSub() constant returns(uint8)
func (_BasMiner *BasMinerCallerSession) AllocationSelfSub() (uint8, error) {
	return _BasMiner.Contract.AllocationSelfSub(&_BasMiner.CallOpts)
}

// AllocationSub is a free data retrieval call binding the contract method 0xf22c9b00.
//
// Solidity: function AllocationSub() constant returns(uint8)
func (_BasMiner *BasMinerCaller) AllocationSub(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _BasMiner.contract.Call(opts, out, "AllocationSub")
	return *ret0, err
}

// AllocationSub is a free data retrieval call binding the contract method 0xf22c9b00.
//
// Solidity: function AllocationSub() constant returns(uint8)
func (_BasMiner *BasMinerSession) AllocationSub() (uint8, error) {
	return _BasMiner.Contract.AllocationSub(&_BasMiner.CallOpts)
}

// AllocationSub is a free data retrieval call binding the contract method 0xf22c9b00.
//
// Solidity: function AllocationSub() constant returns(uint8)
func (_BasMiner *BasMinerCallerSession) AllocationSub() (uint8, error) {
	return _BasMiner.Contract.AllocationSub(&_BasMiner.CallOpts)
}

// GetAllMainNodeAddress is a free data retrieval call binding the contract method 0x179e0f00.
//
// Solidity: function GetAllMainNodeAddress() constant returns(address[])
func (_BasMiner *BasMinerCaller) GetAllMainNodeAddress(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _BasMiner.contract.Call(opts, out, "GetAllMainNodeAddress")
	return *ret0, err
}

// GetAllMainNodeAddress is a free data retrieval call binding the contract method 0x179e0f00.
//
// Solidity: function GetAllMainNodeAddress() constant returns(address[])
func (_BasMiner *BasMinerSession) GetAllMainNodeAddress() ([]common.Address, error) {
	return _BasMiner.Contract.GetAllMainNodeAddress(&_BasMiner.CallOpts)
}

// GetAllMainNodeAddress is a free data retrieval call binding the contract method 0x179e0f00.
//
// Solidity: function GetAllMainNodeAddress() constant returns(address[])
func (_BasMiner *BasMinerCallerSession) GetAllMainNodeAddress() ([]common.Address, error) {
	return _BasMiner.Contract.GetAllMainNodeAddress(&_BasMiner.CallOpts)
}

// MainNode is a free data retrieval call binding the contract method 0x57d3a2ba.
//
// Solidity: function MainNode(uint256 ) constant returns(address)
func (_BasMiner *BasMinerCaller) MainNode(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BasMiner.contract.Call(opts, out, "MainNode", arg0)
	return *ret0, err
}

// MainNode is a free data retrieval call binding the contract method 0x57d3a2ba.
//
// Solidity: function MainNode(uint256 ) constant returns(address)
func (_BasMiner *BasMinerSession) MainNode(arg0 *big.Int) (common.Address, error) {
	return _BasMiner.Contract.MainNode(&_BasMiner.CallOpts, arg0)
}

// MainNode is a free data retrieval call binding the contract method 0x57d3a2ba.
//
// Solidity: function MainNode(uint256 ) constant returns(address)
func (_BasMiner *BasMinerCallerSession) MainNode(arg0 *big.Int) (common.Address, error) {
	return _BasMiner.Contract.MainNode(&_BasMiner.CallOpts, arg0)
}

// MainNodeSize is a free data retrieval call binding the contract method 0xdea21241.
//
// Solidity: function MainNodeSize() constant returns(uint256)
func (_BasMiner *BasMinerCaller) MainNodeSize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BasMiner.contract.Call(opts, out, "MainNodeSize")
	return *ret0, err
}

// MainNodeSize is a free data retrieval call binding the contract method 0xdea21241.
//
// Solidity: function MainNodeSize() constant returns(uint256)
func (_BasMiner *BasMinerSession) MainNodeSize() (*big.Int, error) {
	return _BasMiner.Contract.MainNodeSize(&_BasMiner.CallOpts)
}

// MainNodeSize is a free data retrieval call binding the contract method 0xdea21241.
//
// Solidity: function MainNodeSize() constant returns(uint256)
func (_BasMiner *BasMinerCallerSession) MainNodeSize() (*big.Int, error) {
	return _BasMiner.Contract.MainNodeSize(&_BasMiner.CallOpts)
}

// OANNAddress is a free data retrieval call binding the contract method 0x54cc5d54.
//
// Solidity: function OANNAddress() constant returns(address)
func (_BasMiner *BasMinerCaller) OANNAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BasMiner.contract.Call(opts, out, "OANNAddress")
	return *ret0, err
}

// OANNAddress is a free data retrieval call binding the contract method 0x54cc5d54.
//
// Solidity: function OANNAddress() constant returns(address)
func (_BasMiner *BasMinerSession) OANNAddress() (common.Address, error) {
	return _BasMiner.Contract.OANNAddress(&_BasMiner.CallOpts)
}

// OANNAddress is a free data retrieval call binding the contract method 0x54cc5d54.
//
// Solidity: function OANNAddress() constant returns(address)
func (_BasMiner *BasMinerCallerSession) OANNAddress() (common.Address, error) {
	return _BasMiner.Contract.OANNAddress(&_BasMiner.CallOpts)
}

// SatoshiNakamoto is a free data retrieval call binding the contract method 0xbf6f1121.
//
// Solidity: function Satoshi_Nakamoto() constant returns(address)
func (_BasMiner *BasMinerCaller) SatoshiNakamoto(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BasMiner.contract.Call(opts, out, "Satoshi_Nakamoto")
	return *ret0, err
}

// SatoshiNakamoto is a free data retrieval call binding the contract method 0xbf6f1121.
//
// Solidity: function Satoshi_Nakamoto() constant returns(address)
func (_BasMiner *BasMinerSession) SatoshiNakamoto() (common.Address, error) {
	return _BasMiner.Contract.SatoshiNakamoto(&_BasMiner.CallOpts)
}

// SatoshiNakamoto is a free data retrieval call binding the contract method 0xbf6f1121.
//
// Solidity: function Satoshi_Nakamoto() constant returns(address)
func (_BasMiner *BasMinerCallerSession) SatoshiNakamoto() (common.Address, error) {
	return _BasMiner.Contract.SatoshiNakamoto(&_BasMiner.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_BasMiner *BasMinerCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BasMiner.contract.Call(opts, out, "admin")
	return *ret0, err
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_BasMiner *BasMinerSession) Admin() (common.Address, error) {
	return _BasMiner.Contract.Admin(&_BasMiner.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_BasMiner *BasMinerCallerSession) Admin() (common.Address, error) {
	return _BasMiner.Contract.Admin(&_BasMiner.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) constant returns(uint256)
func (_BasMiner *BasMinerCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BasMiner.contract.Call(opts, out, "balanceOf", arg0)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) constant returns(uint256)
func (_BasMiner *BasMinerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _BasMiner.Contract.BalanceOf(&_BasMiner.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) constant returns(uint256)
func (_BasMiner *BasMinerCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _BasMiner.Contract.BalanceOf(&_BasMiner.CallOpts, arg0)
}

// ContractCaller is a free data retrieval call binding the contract method 0x54cf428a.
//
// Solidity: function contractCaller() constant returns(address)
func (_BasMiner *BasMinerCaller) ContractCaller(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BasMiner.contract.Call(opts, out, "contractCaller")
	return *ret0, err
}

// ContractCaller is a free data retrieval call binding the contract method 0x54cf428a.
//
// Solidity: function contractCaller() constant returns(address)
func (_BasMiner *BasMinerSession) ContractCaller() (common.Address, error) {
	return _BasMiner.Contract.ContractCaller(&_BasMiner.CallOpts)
}

// ContractCaller is a free data retrieval call binding the contract method 0x54cf428a.
//
// Solidity: function contractCaller() constant returns(address)
func (_BasMiner *BasMinerCallerSession) ContractCaller() (common.Address, error) {
	return _BasMiner.Contract.ContractCaller(&_BasMiner.CallOpts)
}

// CustomedSubSetting is a free data retrieval call binding the contract method 0x289acc25.
//
// Solidity: function customedSubSetting() constant returns(uint256 toAdmin, uint256 toBurn, uint256 toMiner, uint256 toRootOwner)
func (_BasMiner *BasMinerCaller) CustomedSubSetting(opts *bind.CallOpts) (struct {
	ToAdmin     *big.Int
	ToBurn      *big.Int
	ToMiner     *big.Int
	ToRootOwner *big.Int
}, error) {
	ret := new(struct {
		ToAdmin     *big.Int
		ToBurn      *big.Int
		ToMiner     *big.Int
		ToRootOwner *big.Int
	})
	out := ret
	err := _BasMiner.contract.Call(opts, out, "customedSubSetting")
	return *ret, err
}

// CustomedSubSetting is a free data retrieval call binding the contract method 0x289acc25.
//
// Solidity: function customedSubSetting() constant returns(uint256 toAdmin, uint256 toBurn, uint256 toMiner, uint256 toRootOwner)
func (_BasMiner *BasMinerSession) CustomedSubSetting() (struct {
	ToAdmin     *big.Int
	ToBurn      *big.Int
	ToMiner     *big.Int
	ToRootOwner *big.Int
}, error) {
	return _BasMiner.Contract.CustomedSubSetting(&_BasMiner.CallOpts)
}

// CustomedSubSetting is a free data retrieval call binding the contract method 0x289acc25.
//
// Solidity: function customedSubSetting() constant returns(uint256 toAdmin, uint256 toBurn, uint256 toMiner, uint256 toRootOwner)
func (_BasMiner *BasMinerCallerSession) CustomedSubSetting() (struct {
	ToAdmin     *big.Int
	ToBurn      *big.Int
	ToMiner     *big.Int
	ToRootOwner *big.Int
}, error) {
	return _BasMiner.Contract.CustomedSubSetting(&_BasMiner.CallOpts)
}

// DefaultSubSetting is a free data retrieval call binding the contract method 0xcf9a400e.
//
// Solidity: function defaultSubSetting() constant returns(uint256 toAdmin, uint256 toBurn, uint256 toMiner, uint256 toRootOwner)
func (_BasMiner *BasMinerCaller) DefaultSubSetting(opts *bind.CallOpts) (struct {
	ToAdmin     *big.Int
	ToBurn      *big.Int
	ToMiner     *big.Int
	ToRootOwner *big.Int
}, error) {
	ret := new(struct {
		ToAdmin     *big.Int
		ToBurn      *big.Int
		ToMiner     *big.Int
		ToRootOwner *big.Int
	})
	out := ret
	err := _BasMiner.contract.Call(opts, out, "defaultSubSetting")
	return *ret, err
}

// DefaultSubSetting is a free data retrieval call binding the contract method 0xcf9a400e.
//
// Solidity: function defaultSubSetting() constant returns(uint256 toAdmin, uint256 toBurn, uint256 toMiner, uint256 toRootOwner)
func (_BasMiner *BasMinerSession) DefaultSubSetting() (struct {
	ToAdmin     *big.Int
	ToBurn      *big.Int
	ToMiner     *big.Int
	ToRootOwner *big.Int
}, error) {
	return _BasMiner.Contract.DefaultSubSetting(&_BasMiner.CallOpts)
}

// DefaultSubSetting is a free data retrieval call binding the contract method 0xcf9a400e.
//
// Solidity: function defaultSubSetting() constant returns(uint256 toAdmin, uint256 toBurn, uint256 toMiner, uint256 toRootOwner)
func (_BasMiner *BasMinerCallerSession) DefaultSubSetting() (struct {
	ToAdmin     *big.Int
	ToBurn      *big.Int
	ToMiner     *big.Int
	ToRootOwner *big.Int
}, error) {
	return _BasMiner.Contract.DefaultSubSetting(&_BasMiner.CallOpts)
}

// RootSetting is a free data retrieval call binding the contract method 0x75d4b7f7.
//
// Solidity: function rootSetting() constant returns(uint256 toAdmin, uint256 toBurn, uint256 toMiner, uint256 toRootOwner)
func (_BasMiner *BasMinerCaller) RootSetting(opts *bind.CallOpts) (struct {
	ToAdmin     *big.Int
	ToBurn      *big.Int
	ToMiner     *big.Int
	ToRootOwner *big.Int
}, error) {
	ret := new(struct {
		ToAdmin     *big.Int
		ToBurn      *big.Int
		ToMiner     *big.Int
		ToRootOwner *big.Int
	})
	out := ret
	err := _BasMiner.contract.Call(opts, out, "rootSetting")
	return *ret, err
}

// RootSetting is a free data retrieval call binding the contract method 0x75d4b7f7.
//
// Solidity: function rootSetting() constant returns(uint256 toAdmin, uint256 toBurn, uint256 toMiner, uint256 toRootOwner)
func (_BasMiner *BasMinerSession) RootSetting() (struct {
	ToAdmin     *big.Int
	ToBurn      *big.Int
	ToMiner     *big.Int
	ToRootOwner *big.Int
}, error) {
	return _BasMiner.Contract.RootSetting(&_BasMiner.CallOpts)
}

// RootSetting is a free data retrieval call binding the contract method 0x75d4b7f7.
//
// Solidity: function rootSetting() constant returns(uint256 toAdmin, uint256 toBurn, uint256 toMiner, uint256 toRootOwner)
func (_BasMiner *BasMinerCallerSession) RootSetting() (struct {
	ToAdmin     *big.Int
	ToBurn      *big.Int
	ToMiner     *big.Int
	ToRootOwner *big.Int
}, error) {
	return _BasMiner.Contract.RootSetting(&_BasMiner.CallOpts)
}

// SelfSubSetting is a free data retrieval call binding the contract method 0xda3926c6.
//
// Solidity: function selfSubSetting() constant returns(uint256 toAdmin, uint256 toBurn, uint256 toMiner, uint256 toRootOwner)
func (_BasMiner *BasMinerCaller) SelfSubSetting(opts *bind.CallOpts) (struct {
	ToAdmin     *big.Int
	ToBurn      *big.Int
	ToMiner     *big.Int
	ToRootOwner *big.Int
}, error) {
	ret := new(struct {
		ToAdmin     *big.Int
		ToBurn      *big.Int
		ToMiner     *big.Int
		ToRootOwner *big.Int
	})
	out := ret
	err := _BasMiner.contract.Call(opts, out, "selfSubSetting")
	return *ret, err
}

// SelfSubSetting is a free data retrieval call binding the contract method 0xda3926c6.
//
// Solidity: function selfSubSetting() constant returns(uint256 toAdmin, uint256 toBurn, uint256 toMiner, uint256 toRootOwner)
func (_BasMiner *BasMinerSession) SelfSubSetting() (struct {
	ToAdmin     *big.Int
	ToBurn      *big.Int
	ToMiner     *big.Int
	ToRootOwner *big.Int
}, error) {
	return _BasMiner.Contract.SelfSubSetting(&_BasMiner.CallOpts)
}

// SelfSubSetting is a free data retrieval call binding the contract method 0xda3926c6.
//
// Solidity: function selfSubSetting() constant returns(uint256 toAdmin, uint256 toBurn, uint256 toMiner, uint256 toRootOwner)
func (_BasMiner *BasMinerCallerSession) SelfSubSetting() (struct {
	ToAdmin     *big.Int
	ToBurn      *big.Int
	ToMiner     *big.Int
	ToRootOwner *big.Int
}, error) {
	return _BasMiner.Contract.SelfSubSetting(&_BasMiner.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_BasMiner *BasMinerCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BasMiner.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_BasMiner *BasMinerSession) Token() (common.Address, error) {
	return _BasMiner.Contract.Token(&_BasMiner.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_BasMiner *BasMinerCallerSession) Token() (common.Address, error) {
	return _BasMiner.Contract.Token(&_BasMiner.CallOpts)
}

// AAddMiner is a paid mutator transaction binding the contract method 0xb963b0d2.
//
// Solidity: function _a_addMiner(address m) returns()
func (_BasMiner *BasMinerTransactor) AAddMiner(opts *bind.TransactOpts, m common.Address) (*types.Transaction, error) {
	return _BasMiner.contract.Transact(opts, "_a_addMiner", m)
}

// AAddMiner is a paid mutator transaction binding the contract method 0xb963b0d2.
//
// Solidity: function _a_addMiner(address m) returns()
func (_BasMiner *BasMinerSession) AAddMiner(m common.Address) (*types.Transaction, error) {
	return _BasMiner.Contract.AAddMiner(&_BasMiner.TransactOpts, m)
}

// AAddMiner is a paid mutator transaction binding the contract method 0xb963b0d2.
//
// Solidity: function _a_addMiner(address m) returns()
func (_BasMiner *BasMinerTransactorSession) AAddMiner(m common.Address) (*types.Transaction, error) {
	return _BasMiner.Contract.AAddMiner(&_BasMiner.TransactOpts, m)
}

// AChangeAdmin is a paid mutator transaction binding the contract method 0x1062f6cc.
//
// Solidity: function _a_changeAdmin(address newOwner) returns()
func (_BasMiner *BasMinerTransactor) AChangeAdmin(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BasMiner.contract.Transact(opts, "_a_changeAdmin", newOwner)
}

// AChangeAdmin is a paid mutator transaction binding the contract method 0x1062f6cc.
//
// Solidity: function _a_changeAdmin(address newOwner) returns()
func (_BasMiner *BasMinerSession) AChangeAdmin(newOwner common.Address) (*types.Transaction, error) {
	return _BasMiner.Contract.AChangeAdmin(&_BasMiner.TransactOpts, newOwner)
}

// AChangeAdmin is a paid mutator transaction binding the contract method 0x1062f6cc.
//
// Solidity: function _a_changeAdmin(address newOwner) returns()
func (_BasMiner *BasMinerTransactorSession) AChangeAdmin(newOwner common.Address) (*types.Transaction, error) {
	return _BasMiner.Contract.AChangeAdmin(&_BasMiner.TransactOpts, newOwner)
}

// AChangeContract is a paid mutator transaction binding the contract method 0x733c1950.
//
// Solidity: function _a_changeContract(address conAddr) returns()
func (_BasMiner *BasMinerTransactor) AChangeContract(opts *bind.TransactOpts, conAddr common.Address) (*types.Transaction, error) {
	return _BasMiner.contract.Transact(opts, "_a_changeContract", conAddr)
}

// AChangeContract is a paid mutator transaction binding the contract method 0x733c1950.
//
// Solidity: function _a_changeContract(address conAddr) returns()
func (_BasMiner *BasMinerSession) AChangeContract(conAddr common.Address) (*types.Transaction, error) {
	return _BasMiner.Contract.AChangeContract(&_BasMiner.TransactOpts, conAddr)
}

// AChangeContract is a paid mutator transaction binding the contract method 0x733c1950.
//
// Solidity: function _a_changeContract(address conAddr) returns()
func (_BasMiner *BasMinerTransactorSession) AChangeContract(conAddr common.Address) (*types.Transaction, error) {
	return _BasMiner.Contract.AChangeContract(&_BasMiner.TransactOpts, conAddr)
}

// AChangeCustomedSubSetting is a paid mutator transaction binding the contract method 0xce803e2b.
//
// Solidity: function _a_changeCustomedSubSetting(uint256 admin, uint256 burn, uint256 miner, uint256 root) returns()
func (_BasMiner *BasMinerTransactor) AChangeCustomedSubSetting(opts *bind.TransactOpts, admin *big.Int, burn *big.Int, miner *big.Int, root *big.Int) (*types.Transaction, error) {
	return _BasMiner.contract.Transact(opts, "_a_changeCustomedSubSetting", admin, burn, miner, root)
}

// AChangeCustomedSubSetting is a paid mutator transaction binding the contract method 0xce803e2b.
//
// Solidity: function _a_changeCustomedSubSetting(uint256 admin, uint256 burn, uint256 miner, uint256 root) returns()
func (_BasMiner *BasMinerSession) AChangeCustomedSubSetting(admin *big.Int, burn *big.Int, miner *big.Int, root *big.Int) (*types.Transaction, error) {
	return _BasMiner.Contract.AChangeCustomedSubSetting(&_BasMiner.TransactOpts, admin, burn, miner, root)
}

// AChangeCustomedSubSetting is a paid mutator transaction binding the contract method 0xce803e2b.
//
// Solidity: function _a_changeCustomedSubSetting(uint256 admin, uint256 burn, uint256 miner, uint256 root) returns()
func (_BasMiner *BasMinerTransactorSession) AChangeCustomedSubSetting(admin *big.Int, burn *big.Int, miner *big.Int, root *big.Int) (*types.Transaction, error) {
	return _BasMiner.Contract.AChangeCustomedSubSetting(&_BasMiner.TransactOpts, admin, burn, miner, root)
}

// AChangeDefaultSubSetting is a paid mutator transaction binding the contract method 0x13612f56.
//
// Solidity: function _a_changeDefaultSubSetting(uint256 admin, uint256 burn, uint256 miner, uint256 root) returns()
func (_BasMiner *BasMinerTransactor) AChangeDefaultSubSetting(opts *bind.TransactOpts, admin *big.Int, burn *big.Int, miner *big.Int, root *big.Int) (*types.Transaction, error) {
	return _BasMiner.contract.Transact(opts, "_a_changeDefaultSubSetting", admin, burn, miner, root)
}

// AChangeDefaultSubSetting is a paid mutator transaction binding the contract method 0x13612f56.
//
// Solidity: function _a_changeDefaultSubSetting(uint256 admin, uint256 burn, uint256 miner, uint256 root) returns()
func (_BasMiner *BasMinerSession) AChangeDefaultSubSetting(admin *big.Int, burn *big.Int, miner *big.Int, root *big.Int) (*types.Transaction, error) {
	return _BasMiner.Contract.AChangeDefaultSubSetting(&_BasMiner.TransactOpts, admin, burn, miner, root)
}

// AChangeDefaultSubSetting is a paid mutator transaction binding the contract method 0x13612f56.
//
// Solidity: function _a_changeDefaultSubSetting(uint256 admin, uint256 burn, uint256 miner, uint256 root) returns()
func (_BasMiner *BasMinerTransactorSession) AChangeDefaultSubSetting(admin *big.Int, burn *big.Int, miner *big.Int, root *big.Int) (*types.Transaction, error) {
	return _BasMiner.Contract.AChangeDefaultSubSetting(&_BasMiner.TransactOpts, admin, burn, miner, root)
}

// AChangeRootSetting is a paid mutator transaction binding the contract method 0x43b7919a.
//
// Solidity: function _a_changeRootSetting(uint256 admin, uint256 burn, uint256 miner) returns()
func (_BasMiner *BasMinerTransactor) AChangeRootSetting(opts *bind.TransactOpts, admin *big.Int, burn *big.Int, miner *big.Int) (*types.Transaction, error) {
	return _BasMiner.contract.Transact(opts, "_a_changeRootSetting", admin, burn, miner)
}

// AChangeRootSetting is a paid mutator transaction binding the contract method 0x43b7919a.
//
// Solidity: function _a_changeRootSetting(uint256 admin, uint256 burn, uint256 miner) returns()
func (_BasMiner *BasMinerSession) AChangeRootSetting(admin *big.Int, burn *big.Int, miner *big.Int) (*types.Transaction, error) {
	return _BasMiner.Contract.AChangeRootSetting(&_BasMiner.TransactOpts, admin, burn, miner)
}

// AChangeRootSetting is a paid mutator transaction binding the contract method 0x43b7919a.
//
// Solidity: function _a_changeRootSetting(uint256 admin, uint256 burn, uint256 miner) returns()
func (_BasMiner *BasMinerTransactorSession) AChangeRootSetting(admin *big.Int, burn *big.Int, miner *big.Int) (*types.Transaction, error) {
	return _BasMiner.Contract.AChangeRootSetting(&_BasMiner.TransactOpts, admin, burn, miner)
}

// AChangeSelfSubSetting is a paid mutator transaction binding the contract method 0x3f2ed91e.
//
// Solidity: function _a_changeSelfSubSetting(uint256 admin, uint256 burn, uint256 miner, uint256 root) returns()
func (_BasMiner *BasMinerTransactor) AChangeSelfSubSetting(opts *bind.TransactOpts, admin *big.Int, burn *big.Int, miner *big.Int, root *big.Int) (*types.Transaction, error) {
	return _BasMiner.contract.Transact(opts, "_a_changeSelfSubSetting", admin, burn, miner, root)
}

// AChangeSelfSubSetting is a paid mutator transaction binding the contract method 0x3f2ed91e.
//
// Solidity: function _a_changeSelfSubSetting(uint256 admin, uint256 burn, uint256 miner, uint256 root) returns()
func (_BasMiner *BasMinerSession) AChangeSelfSubSetting(admin *big.Int, burn *big.Int, miner *big.Int, root *big.Int) (*types.Transaction, error) {
	return _BasMiner.Contract.AChangeSelfSubSetting(&_BasMiner.TransactOpts, admin, burn, miner, root)
}

// AChangeSelfSubSetting is a paid mutator transaction binding the contract method 0x3f2ed91e.
//
// Solidity: function _a_changeSelfSubSetting(uint256 admin, uint256 burn, uint256 miner, uint256 root) returns()
func (_BasMiner *BasMinerTransactorSession) AChangeSelfSubSetting(admin *big.Int, burn *big.Int, miner *big.Int, root *big.Int) (*types.Transaction, error) {
	return _BasMiner.Contract.AChangeSelfSubSetting(&_BasMiner.TransactOpts, admin, burn, miner, root)
}

// AEmergencyWithdraw is a paid mutator transaction binding the contract method 0xac6e389d.
//
// Solidity: function _a_emergencyWithdraw(address to, uint256 no) returns()
func (_BasMiner *BasMinerTransactor) AEmergencyWithdraw(opts *bind.TransactOpts, to common.Address, no *big.Int) (*types.Transaction, error) {
	return _BasMiner.contract.Transact(opts, "_a_emergencyWithdraw", to, no)
}

// AEmergencyWithdraw is a paid mutator transaction binding the contract method 0xac6e389d.
//
// Solidity: function _a_emergencyWithdraw(address to, uint256 no) returns()
func (_BasMiner *BasMinerSession) AEmergencyWithdraw(to common.Address, no *big.Int) (*types.Transaction, error) {
	return _BasMiner.Contract.AEmergencyWithdraw(&_BasMiner.TransactOpts, to, no)
}

// AEmergencyWithdraw is a paid mutator transaction binding the contract method 0xac6e389d.
//
// Solidity: function _a_emergencyWithdraw(address to, uint256 no) returns()
func (_BasMiner *BasMinerTransactorSession) AEmergencyWithdraw(to common.Address, no *big.Int) (*types.Transaction, error) {
	return _BasMiner.Contract.AEmergencyWithdraw(&_BasMiner.TransactOpts, to, no)
}

// ARemoveMiner is a paid mutator transaction binding the contract method 0xbe7738d8.
//
// Solidity: function _a_removeMiner(address miner) returns()
func (_BasMiner *BasMinerTransactor) ARemoveMiner(opts *bind.TransactOpts, miner common.Address) (*types.Transaction, error) {
	return _BasMiner.contract.Transact(opts, "_a_removeMiner", miner)
}

// ARemoveMiner is a paid mutator transaction binding the contract method 0xbe7738d8.
//
// Solidity: function _a_removeMiner(address miner) returns()
func (_BasMiner *BasMinerSession) ARemoveMiner(miner common.Address) (*types.Transaction, error) {
	return _BasMiner.Contract.ARemoveMiner(&_BasMiner.TransactOpts, miner)
}

// ARemoveMiner is a paid mutator transaction binding the contract method 0xbe7738d8.
//
// Solidity: function _a_removeMiner(address miner) returns()
func (_BasMiner *BasMinerTransactorSession) ARemoveMiner(miner common.Address) (*types.Transaction, error) {
	return _BasMiner.Contract.ARemoveMiner(&_BasMiner.TransactOpts, miner)
}

// AReplaceMiner is a paid mutator transaction binding the contract method 0x55c120cc.
//
// Solidity: function _a_replaceMiner(address oldM, address newM) returns()
func (_BasMiner *BasMinerTransactor) AReplaceMiner(opts *bind.TransactOpts, oldM common.Address, newM common.Address) (*types.Transaction, error) {
	return _BasMiner.contract.Transact(opts, "_a_replaceMiner", oldM, newM)
}

// AReplaceMiner is a paid mutator transaction binding the contract method 0x55c120cc.
//
// Solidity: function _a_replaceMiner(address oldM, address newM) returns()
func (_BasMiner *BasMinerSession) AReplaceMiner(oldM common.Address, newM common.Address) (*types.Transaction, error) {
	return _BasMiner.Contract.AReplaceMiner(&_BasMiner.TransactOpts, oldM, newM)
}

// AReplaceMiner is a paid mutator transaction binding the contract method 0x55c120cc.
//
// Solidity: function _a_replaceMiner(address oldM, address newM) returns()
func (_BasMiner *BasMinerTransactorSession) AReplaceMiner(oldM common.Address, newM common.Address) (*types.Transaction, error) {
	return _BasMiner.Contract.AReplaceMiner(&_BasMiner.TransactOpts, oldM, newM)
}

// CAllocateProfit is a paid mutator transaction binding the contract method 0x6f3cb28f.
//
// Solidity: function _c_allocateProfit(uint256 cost, uint8 allocateType, bytes32 receiptNumber, address rootOwner) returns()
func (_BasMiner *BasMinerTransactor) CAllocateProfit(opts *bind.TransactOpts, cost *big.Int, allocateType uint8, receiptNumber [32]byte, rootOwner common.Address) (*types.Transaction, error) {
	return _BasMiner.contract.Transact(opts, "_c_allocateProfit", cost, allocateType, receiptNumber, rootOwner)
}

// CAllocateProfit is a paid mutator transaction binding the contract method 0x6f3cb28f.
//
// Solidity: function _c_allocateProfit(uint256 cost, uint8 allocateType, bytes32 receiptNumber, address rootOwner) returns()
func (_BasMiner *BasMinerSession) CAllocateProfit(cost *big.Int, allocateType uint8, receiptNumber [32]byte, rootOwner common.Address) (*types.Transaction, error) {
	return _BasMiner.Contract.CAllocateProfit(&_BasMiner.TransactOpts, cost, allocateType, receiptNumber, rootOwner)
}

// CAllocateProfit is a paid mutator transaction binding the contract method 0x6f3cb28f.
//
// Solidity: function _c_allocateProfit(uint256 cost, uint8 allocateType, bytes32 receiptNumber, address rootOwner) returns()
func (_BasMiner *BasMinerTransactorSession) CAllocateProfit(cost *big.Int, allocateType uint8, receiptNumber [32]byte, rootOwner common.Address) (*types.Transaction, error) {
	return _BasMiner.Contract.CAllocateProfit(&_BasMiner.TransactOpts, cost, allocateType, receiptNumber, rootOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_BasMiner *BasMinerTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasMiner.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_BasMiner *BasMinerSession) Withdraw() (*types.Transaction, error) {
	return _BasMiner.Contract.Withdraw(&_BasMiner.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_BasMiner *BasMinerTransactorSession) Withdraw() (*types.Transaction, error) {
	return _BasMiner.Contract.Withdraw(&_BasMiner.TransactOpts)
}

// BasMinerAllocationChangedIterator is returned from FilterAllocationChanged and is used to iterate over the raw logs and unpacked data for AllocationChanged events raised by the BasMiner contract.
type BasMinerAllocationChangedIterator struct {
	Event *BasMinerAllocationChanged // Event containing the contract specifics and raw log

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
func (it *BasMinerAllocationChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasMinerAllocationChanged)
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
		it.Event = new(BasMinerAllocationChanged)
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
func (it *BasMinerAllocationChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasMinerAllocationChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasMinerAllocationChanged represents a AllocationChanged event raised by the BasMiner contract.
type BasMinerAllocationChanged struct {
	AllocateType uint8
	ToAdmin      *big.Int
	ToBurn       *big.Int
	ToMiner      *big.Int
	ToRoot       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAllocationChanged is a free log retrieval operation binding the contract event 0x0244e311880cf8969cc8b01ad7c19317d1b2f9b197754df3ae5088d8dfa58a37.
//
// Solidity: event AllocationChanged(uint8 allocateType, uint256 toAdmin, uint256 toBurn, uint256 toMiner, uint256 toRoot)
func (_BasMiner *BasMinerFilterer) FilterAllocationChanged(opts *bind.FilterOpts) (*BasMinerAllocationChangedIterator, error) {

	logs, sub, err := _BasMiner.contract.FilterLogs(opts, "AllocationChanged")
	if err != nil {
		return nil, err
	}
	return &BasMinerAllocationChangedIterator{contract: _BasMiner.contract, event: "AllocationChanged", logs: logs, sub: sub}, nil
}

// WatchAllocationChanged is a free log subscription operation binding the contract event 0x0244e311880cf8969cc8b01ad7c19317d1b2f9b197754df3ae5088d8dfa58a37.
//
// Solidity: event AllocationChanged(uint8 allocateType, uint256 toAdmin, uint256 toBurn, uint256 toMiner, uint256 toRoot)
func (_BasMiner *BasMinerFilterer) WatchAllocationChanged(opts *bind.WatchOpts, sink chan<- *BasMinerAllocationChanged) (event.Subscription, error) {

	logs, sub, err := _BasMiner.contract.WatchLogs(opts, "AllocationChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasMinerAllocationChanged)
				if err := _BasMiner.contract.UnpackLog(event, "AllocationChanged", log); err != nil {
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

// ParseAllocationChanged is a log parse operation binding the contract event 0x0244e311880cf8969cc8b01ad7c19317d1b2f9b197754df3ae5088d8dfa58a37.
//
// Solidity: event AllocationChanged(uint8 allocateType, uint256 toAdmin, uint256 toBurn, uint256 toMiner, uint256 toRoot)
func (_BasMiner *BasMinerFilterer) ParseAllocationChanged(log types.Log) (*BasMinerAllocationChanged, error) {
	event := new(BasMinerAllocationChanged)
	if err := _BasMiner.contract.UnpackLog(event, "AllocationChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BasMinerMinerAddIterator is returned from FilterMinerAdd and is used to iterate over the raw logs and unpacked data for MinerAdd events raised by the BasMiner contract.
type BasMinerMinerAddIterator struct {
	Event *BasMinerMinerAdd // Event containing the contract specifics and raw log

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
func (it *BasMinerMinerAddIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasMinerMinerAdd)
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
		it.Event = new(BasMinerMinerAdd)
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
func (it *BasMinerMinerAddIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasMinerMinerAddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasMinerMinerAdd represents a MinerAdd event raised by the BasMiner contract.
type BasMinerMinerAdd struct {
	Miner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMinerAdd is a free log retrieval operation binding the contract event 0x424346829819678c1f897cd0d980f28004ae4eb06bb95f034a791e0e4e9cebf5.
//
// Solidity: event MinerAdd(address miner)
func (_BasMiner *BasMinerFilterer) FilterMinerAdd(opts *bind.FilterOpts) (*BasMinerMinerAddIterator, error) {

	logs, sub, err := _BasMiner.contract.FilterLogs(opts, "MinerAdd")
	if err != nil {
		return nil, err
	}
	return &BasMinerMinerAddIterator{contract: _BasMiner.contract, event: "MinerAdd", logs: logs, sub: sub}, nil
}

// WatchMinerAdd is a free log subscription operation binding the contract event 0x424346829819678c1f897cd0d980f28004ae4eb06bb95f034a791e0e4e9cebf5.
//
// Solidity: event MinerAdd(address miner)
func (_BasMiner *BasMinerFilterer) WatchMinerAdd(opts *bind.WatchOpts, sink chan<- *BasMinerMinerAdd) (event.Subscription, error) {

	logs, sub, err := _BasMiner.contract.WatchLogs(opts, "MinerAdd")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasMinerMinerAdd)
				if err := _BasMiner.contract.UnpackLog(event, "MinerAdd", log); err != nil {
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

// ParseMinerAdd is a log parse operation binding the contract event 0x424346829819678c1f897cd0d980f28004ae4eb06bb95f034a791e0e4e9cebf5.
//
// Solidity: event MinerAdd(address miner)
func (_BasMiner *BasMinerFilterer) ParseMinerAdd(log types.Log) (*BasMinerMinerAdd, error) {
	event := new(BasMinerMinerAdd)
	if err := _BasMiner.contract.UnpackLog(event, "MinerAdd", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BasMinerMinerRemoveIterator is returned from FilterMinerRemove and is used to iterate over the raw logs and unpacked data for MinerRemove events raised by the BasMiner contract.
type BasMinerMinerRemoveIterator struct {
	Event *BasMinerMinerRemove // Event containing the contract specifics and raw log

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
func (it *BasMinerMinerRemoveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasMinerMinerRemove)
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
		it.Event = new(BasMinerMinerRemove)
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
func (it *BasMinerMinerRemoveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasMinerMinerRemoveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasMinerMinerRemove represents a MinerRemove event raised by the BasMiner contract.
type BasMinerMinerRemove struct {
	Miner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMinerRemove is a free log retrieval operation binding the contract event 0x872f9c238db5a645adcaea90821e7c7884c6e063b9e047fd2658592f7cea397d.
//
// Solidity: event MinerRemove(address miner)
func (_BasMiner *BasMinerFilterer) FilterMinerRemove(opts *bind.FilterOpts) (*BasMinerMinerRemoveIterator, error) {

	logs, sub, err := _BasMiner.contract.FilterLogs(opts, "MinerRemove")
	if err != nil {
		return nil, err
	}
	return &BasMinerMinerRemoveIterator{contract: _BasMiner.contract, event: "MinerRemove", logs: logs, sub: sub}, nil
}

// WatchMinerRemove is a free log subscription operation binding the contract event 0x872f9c238db5a645adcaea90821e7c7884c6e063b9e047fd2658592f7cea397d.
//
// Solidity: event MinerRemove(address miner)
func (_BasMiner *BasMinerFilterer) WatchMinerRemove(opts *bind.WatchOpts, sink chan<- *BasMinerMinerRemove) (event.Subscription, error) {

	logs, sub, err := _BasMiner.contract.WatchLogs(opts, "MinerRemove")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasMinerMinerRemove)
				if err := _BasMiner.contract.UnpackLog(event, "MinerRemove", log); err != nil {
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

// ParseMinerRemove is a log parse operation binding the contract event 0x872f9c238db5a645adcaea90821e7c7884c6e063b9e047fd2658592f7cea397d.
//
// Solidity: event MinerRemove(address miner)
func (_BasMiner *BasMinerFilterer) ParseMinerRemove(log types.Log) (*BasMinerMinerRemove, error) {
	event := new(BasMinerMinerRemove)
	if err := _BasMiner.contract.UnpackLog(event, "MinerRemove", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BasMinerMinerReplaceIterator is returned from FilterMinerReplace and is used to iterate over the raw logs and unpacked data for MinerReplace events raised by the BasMiner contract.
type BasMinerMinerReplaceIterator struct {
	Event *BasMinerMinerReplace // Event containing the contract specifics and raw log

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
func (it *BasMinerMinerReplaceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasMinerMinerReplace)
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
		it.Event = new(BasMinerMinerReplace)
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
func (it *BasMinerMinerReplaceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasMinerMinerReplaceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasMinerMinerReplace represents a MinerReplace event raised by the BasMiner contract.
type BasMinerMinerReplace struct {
	OldMiner common.Address
	NewMiner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMinerReplace is a free log retrieval operation binding the contract event 0xc205c6ca767906b347bde3b64a6a00763decbd4fb0ea515cb8afb411647591ee.
//
// Solidity: event MinerReplace(address oldMiner, address newMiner)
func (_BasMiner *BasMinerFilterer) FilterMinerReplace(opts *bind.FilterOpts) (*BasMinerMinerReplaceIterator, error) {

	logs, sub, err := _BasMiner.contract.FilterLogs(opts, "MinerReplace")
	if err != nil {
		return nil, err
	}
	return &BasMinerMinerReplaceIterator{contract: _BasMiner.contract, event: "MinerReplace", logs: logs, sub: sub}, nil
}

// WatchMinerReplace is a free log subscription operation binding the contract event 0xc205c6ca767906b347bde3b64a6a00763decbd4fb0ea515cb8afb411647591ee.
//
// Solidity: event MinerReplace(address oldMiner, address newMiner)
func (_BasMiner *BasMinerFilterer) WatchMinerReplace(opts *bind.WatchOpts, sink chan<- *BasMinerMinerReplace) (event.Subscription, error) {

	logs, sub, err := _BasMiner.contract.WatchLogs(opts, "MinerReplace")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasMinerMinerReplace)
				if err := _BasMiner.contract.UnpackLog(event, "MinerReplace", log); err != nil {
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

// ParseMinerReplace is a log parse operation binding the contract event 0xc205c6ca767906b347bde3b64a6a00763decbd4fb0ea515cb8afb411647591ee.
//
// Solidity: event MinerReplace(address oldMiner, address newMiner)
func (_BasMiner *BasMinerFilterer) ParseMinerReplace(log types.Log) (*BasMinerMinerReplace, error) {
	event := new(BasMinerMinerReplace)
	if err := _BasMiner.contract.UnpackLog(event, "MinerReplace", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BasMinerReceiptIterator is returned from FilterReceipt and is used to iterate over the raw logs and unpacked data for Receipt events raised by the BasMiner contract.
type BasMinerReceiptIterator struct {
	Event *BasMinerReceipt // Event containing the contract specifics and raw log

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
func (it *BasMinerReceiptIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasMinerReceipt)
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
		it.Event = new(BasMinerReceipt)
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
func (it *BasMinerReceiptIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasMinerReceiptIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasMinerReceipt represents a Receipt event raised by the BasMiner contract.
type BasMinerReceipt struct {
	ReceiptNumber [32]byte
	Amout         *big.Int
	Allocation    uint8
	From          common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterReceipt is a free log retrieval operation binding the contract event 0x3758759a794787144866c77ee84ba833a2225c12a6d58982691b2e88363544ff.
//
// Solidity: event Receipt(bytes32 receiptNumber, uint256 amout, uint8 allocation, address from)
func (_BasMiner *BasMinerFilterer) FilterReceipt(opts *bind.FilterOpts) (*BasMinerReceiptIterator, error) {

	logs, sub, err := _BasMiner.contract.FilterLogs(opts, "Receipt")
	if err != nil {
		return nil, err
	}
	return &BasMinerReceiptIterator{contract: _BasMiner.contract, event: "Receipt", logs: logs, sub: sub}, nil
}

// WatchReceipt is a free log subscription operation binding the contract event 0x3758759a794787144866c77ee84ba833a2225c12a6d58982691b2e88363544ff.
//
// Solidity: event Receipt(bytes32 receiptNumber, uint256 amout, uint8 allocation, address from)
func (_BasMiner *BasMinerFilterer) WatchReceipt(opts *bind.WatchOpts, sink chan<- *BasMinerReceipt) (event.Subscription, error) {

	logs, sub, err := _BasMiner.contract.WatchLogs(opts, "Receipt")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasMinerReceipt)
				if err := _BasMiner.contract.UnpackLog(event, "Receipt", log); err != nil {
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

// ParseReceipt is a log parse operation binding the contract event 0x3758759a794787144866c77ee84ba833a2225c12a6d58982691b2e88363544ff.
//
// Solidity: event Receipt(bytes32 receiptNumber, uint256 amout, uint8 allocation, address from)
func (_BasMiner *BasMinerFilterer) ParseReceipt(log types.Log) (*BasMinerReceipt, error) {
	event := new(BasMinerReceipt)
	if err := _BasMiner.contract.UnpackLog(event, "Receipt", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BasMinerWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the BasMiner contract.
type BasMinerWithdrawIterator struct {
	Event *BasMinerWithdraw // Event containing the contract specifics and raw log

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
func (it *BasMinerWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasMinerWithdraw)
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
		it.Event = new(BasMinerWithdraw)
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
func (it *BasMinerWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasMinerWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasMinerWithdraw represents a Withdraw event raised by the BasMiner contract.
type BasMinerWithdraw struct {
	Drawer common.Address
	Amout  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address drawer, uint256 amout)
func (_BasMiner *BasMinerFilterer) FilterWithdraw(opts *bind.FilterOpts) (*BasMinerWithdrawIterator, error) {

	logs, sub, err := _BasMiner.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &BasMinerWithdrawIterator{contract: _BasMiner.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address drawer, uint256 amout)
func (_BasMiner *BasMinerFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *BasMinerWithdraw) (event.Subscription, error) {

	logs, sub, err := _BasMiner.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasMinerWithdraw)
				if err := _BasMiner.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address drawer, uint256 amout)
func (_BasMiner *BasMinerFilterer) ParseWithdraw(log types.Log) (*BasMinerWithdraw, error) {
	event := new(BasMinerWithdraw)
	if err := _BasMiner.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	return event, nil
}
