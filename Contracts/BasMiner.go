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
const BasMinerABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"admin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"burn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"miner\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"}],\"name\":\"setDefaultSubSetting\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"oldM\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newM\",\"type\":\"address\"}],\"name\":\"replaceMiner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"}],\"name\":\"removeMiner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"_a_changeAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"GetAllMainNodeAddress\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"admin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"burn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"miner\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"}],\"name\":\"setCustomedSubSetting\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"customedSubSetting\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"toAdmin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toBurn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toMiner\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toRootOwner\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"OANNAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contractCaller\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"MainNode\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ViceNodeSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"admin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"burn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"miner\",\"type\":\"uint256\"}],\"name\":\"setRootSetting\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"receiptNumber\",\"type\":\"bytes32\"}],\"name\":\"rootAllocateProfit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"conAddr\",\"type\":\"address\"}],\"name\":\"_a_changeContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rootSetting\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"toAdmin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toBurn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toMiner\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"no\",\"type\":\"uint256\"}],\"name\":\"emergencyWithdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"internalType\":\"enumBasMiner.ProfitType\",\"name\":\"typ\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"rOwner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"receiptNumber\",\"type\":\"bytes32\"}],\"name\":\"subNameProfit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"oann\",\"type\":\"address\"}],\"name\":\"setOANN\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"Satoshi_Nakamoto\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ViceNode\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"defaultSubSetting\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"toAdmin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toBurn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toMiner\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toRootOwner\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"selfSubSetting\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"toAdmin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toBurn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toMiner\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toRootOwner\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MainNodeSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"m\",\"type\":\"address\"}],\"name\":\"addMiner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractBasToken\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"team\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"receiptNumber\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"Receipt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"drawer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amout\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"AllocationChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"MinerChanged\",\"type\":\"event\"}]"

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

// ViceNode is a free data retrieval call binding the contract method 0xcccccfcf.
//
// Solidity: function ViceNode(uint256 ) constant returns(address)
func (_BasMiner *BasMinerCaller) ViceNode(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BasMiner.contract.Call(opts, out, "ViceNode", arg0)
	return *ret0, err
}

// ViceNode is a free data retrieval call binding the contract method 0xcccccfcf.
//
// Solidity: function ViceNode(uint256 ) constant returns(address)
func (_BasMiner *BasMinerSession) ViceNode(arg0 *big.Int) (common.Address, error) {
	return _BasMiner.Contract.ViceNode(&_BasMiner.CallOpts, arg0)
}

// ViceNode is a free data retrieval call binding the contract method 0xcccccfcf.
//
// Solidity: function ViceNode(uint256 ) constant returns(address)
func (_BasMiner *BasMinerCallerSession) ViceNode(arg0 *big.Int) (common.Address, error) {
	return _BasMiner.Contract.ViceNode(&_BasMiner.CallOpts, arg0)
}

// ViceNodeSize is a free data retrieval call binding the contract method 0x5911606f.
//
// Solidity: function ViceNodeSize() constant returns(uint256)
func (_BasMiner *BasMinerCaller) ViceNodeSize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BasMiner.contract.Call(opts, out, "ViceNodeSize")
	return *ret0, err
}

// ViceNodeSize is a free data retrieval call binding the contract method 0x5911606f.
//
// Solidity: function ViceNodeSize() constant returns(uint256)
func (_BasMiner *BasMinerSession) ViceNodeSize() (*big.Int, error) {
	return _BasMiner.Contract.ViceNodeSize(&_BasMiner.CallOpts)
}

// ViceNodeSize is a free data retrieval call binding the contract method 0x5911606f.
//
// Solidity: function ViceNodeSize() constant returns(uint256)
func (_BasMiner *BasMinerCallerSession) ViceNodeSize() (*big.Int, error) {
	return _BasMiner.Contract.ViceNodeSize(&_BasMiner.CallOpts)
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
// Solidity: function rootSetting() constant returns(uint256 toAdmin, uint256 toBurn, uint256 toMiner)
func (_BasMiner *BasMinerCaller) RootSetting(opts *bind.CallOpts) (struct {
	ToAdmin *big.Int
	ToBurn  *big.Int
	ToMiner *big.Int
}, error) {
	ret := new(struct {
		ToAdmin *big.Int
		ToBurn  *big.Int
		ToMiner *big.Int
	})
	out := ret
	err := _BasMiner.contract.Call(opts, out, "rootSetting")
	return *ret, err
}

// RootSetting is a free data retrieval call binding the contract method 0x75d4b7f7.
//
// Solidity: function rootSetting() constant returns(uint256 toAdmin, uint256 toBurn, uint256 toMiner)
func (_BasMiner *BasMinerSession) RootSetting() (struct {
	ToAdmin *big.Int
	ToBurn  *big.Int
	ToMiner *big.Int
}, error) {
	return _BasMiner.Contract.RootSetting(&_BasMiner.CallOpts)
}

// RootSetting is a free data retrieval call binding the contract method 0x75d4b7f7.
//
// Solidity: function rootSetting() constant returns(uint256 toAdmin, uint256 toBurn, uint256 toMiner)
func (_BasMiner *BasMinerCallerSession) RootSetting() (struct {
	ToAdmin *big.Int
	ToBurn  *big.Int
	ToMiner *big.Int
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

// AddMiner is a paid mutator transaction binding the contract method 0xf3982e5e.
//
// Solidity: function addMiner(address m) returns()
func (_BasMiner *BasMinerTransactor) AddMiner(opts *bind.TransactOpts, m common.Address) (*types.Transaction, error) {
	return _BasMiner.contract.Transact(opts, "addMiner", m)
}

// AddMiner is a paid mutator transaction binding the contract method 0xf3982e5e.
//
// Solidity: function addMiner(address m) returns()
func (_BasMiner *BasMinerSession) AddMiner(m common.Address) (*types.Transaction, error) {
	return _BasMiner.Contract.AddMiner(&_BasMiner.TransactOpts, m)
}

// AddMiner is a paid mutator transaction binding the contract method 0xf3982e5e.
//
// Solidity: function addMiner(address m) returns()
func (_BasMiner *BasMinerTransactorSession) AddMiner(m common.Address) (*types.Transaction, error) {
	return _BasMiner.Contract.AddMiner(&_BasMiner.TransactOpts, m)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x95ccea67.
//
// Solidity: function emergencyWithdraw(address to, uint256 no) returns()
func (_BasMiner *BasMinerTransactor) EmergencyWithdraw(opts *bind.TransactOpts, to common.Address, no *big.Int) (*types.Transaction, error) {
	return _BasMiner.contract.Transact(opts, "emergencyWithdraw", to, no)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x95ccea67.
//
// Solidity: function emergencyWithdraw(address to, uint256 no) returns()
func (_BasMiner *BasMinerSession) EmergencyWithdraw(to common.Address, no *big.Int) (*types.Transaction, error) {
	return _BasMiner.Contract.EmergencyWithdraw(&_BasMiner.TransactOpts, to, no)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x95ccea67.
//
// Solidity: function emergencyWithdraw(address to, uint256 no) returns()
func (_BasMiner *BasMinerTransactorSession) EmergencyWithdraw(to common.Address, no *big.Int) (*types.Transaction, error) {
	return _BasMiner.Contract.EmergencyWithdraw(&_BasMiner.TransactOpts, to, no)
}

// RemoveMiner is a paid mutator transaction binding the contract method 0x10242590.
//
// Solidity: function removeMiner(address miner) returns()
func (_BasMiner *BasMinerTransactor) RemoveMiner(opts *bind.TransactOpts, miner common.Address) (*types.Transaction, error) {
	return _BasMiner.contract.Transact(opts, "removeMiner", miner)
}

// RemoveMiner is a paid mutator transaction binding the contract method 0x10242590.
//
// Solidity: function removeMiner(address miner) returns()
func (_BasMiner *BasMinerSession) RemoveMiner(miner common.Address) (*types.Transaction, error) {
	return _BasMiner.Contract.RemoveMiner(&_BasMiner.TransactOpts, miner)
}

// RemoveMiner is a paid mutator transaction binding the contract method 0x10242590.
//
// Solidity: function removeMiner(address miner) returns()
func (_BasMiner *BasMinerTransactorSession) RemoveMiner(miner common.Address) (*types.Transaction, error) {
	return _BasMiner.Contract.RemoveMiner(&_BasMiner.TransactOpts, miner)
}

// ReplaceMiner is a paid mutator transaction binding the contract method 0x0bdf2fd3.
//
// Solidity: function replaceMiner(address oldM, address newM) returns()
func (_BasMiner *BasMinerTransactor) ReplaceMiner(opts *bind.TransactOpts, oldM common.Address, newM common.Address) (*types.Transaction, error) {
	return _BasMiner.contract.Transact(opts, "replaceMiner", oldM, newM)
}

// ReplaceMiner is a paid mutator transaction binding the contract method 0x0bdf2fd3.
//
// Solidity: function replaceMiner(address oldM, address newM) returns()
func (_BasMiner *BasMinerSession) ReplaceMiner(oldM common.Address, newM common.Address) (*types.Transaction, error) {
	return _BasMiner.Contract.ReplaceMiner(&_BasMiner.TransactOpts, oldM, newM)
}

// ReplaceMiner is a paid mutator transaction binding the contract method 0x0bdf2fd3.
//
// Solidity: function replaceMiner(address oldM, address newM) returns()
func (_BasMiner *BasMinerTransactorSession) ReplaceMiner(oldM common.Address, newM common.Address) (*types.Transaction, error) {
	return _BasMiner.Contract.ReplaceMiner(&_BasMiner.TransactOpts, oldM, newM)
}

// RootAllocateProfit is a paid mutator transaction binding the contract method 0x645a3675.
//
// Solidity: function rootAllocateProfit(uint256 cost, bytes32 receiptNumber) returns()
func (_BasMiner *BasMinerTransactor) RootAllocateProfit(opts *bind.TransactOpts, cost *big.Int, receiptNumber [32]byte) (*types.Transaction, error) {
	return _BasMiner.contract.Transact(opts, "rootAllocateProfit", cost, receiptNumber)
}

// RootAllocateProfit is a paid mutator transaction binding the contract method 0x645a3675.
//
// Solidity: function rootAllocateProfit(uint256 cost, bytes32 receiptNumber) returns()
func (_BasMiner *BasMinerSession) RootAllocateProfit(cost *big.Int, receiptNumber [32]byte) (*types.Transaction, error) {
	return _BasMiner.Contract.RootAllocateProfit(&_BasMiner.TransactOpts, cost, receiptNumber)
}

// RootAllocateProfit is a paid mutator transaction binding the contract method 0x645a3675.
//
// Solidity: function rootAllocateProfit(uint256 cost, bytes32 receiptNumber) returns()
func (_BasMiner *BasMinerTransactorSession) RootAllocateProfit(cost *big.Int, receiptNumber [32]byte) (*types.Transaction, error) {
	return _BasMiner.Contract.RootAllocateProfit(&_BasMiner.TransactOpts, cost, receiptNumber)
}

// SetCustomedSubSetting is a paid mutator transaction binding the contract method 0x1918f4ef.
//
// Solidity: function setCustomedSubSetting(uint256 admin, uint256 burn, uint256 miner, uint256 root) returns()
func (_BasMiner *BasMinerTransactor) SetCustomedSubSetting(opts *bind.TransactOpts, admin *big.Int, burn *big.Int, miner *big.Int, root *big.Int) (*types.Transaction, error) {
	return _BasMiner.contract.Transact(opts, "setCustomedSubSetting", admin, burn, miner, root)
}

// SetCustomedSubSetting is a paid mutator transaction binding the contract method 0x1918f4ef.
//
// Solidity: function setCustomedSubSetting(uint256 admin, uint256 burn, uint256 miner, uint256 root) returns()
func (_BasMiner *BasMinerSession) SetCustomedSubSetting(admin *big.Int, burn *big.Int, miner *big.Int, root *big.Int) (*types.Transaction, error) {
	return _BasMiner.Contract.SetCustomedSubSetting(&_BasMiner.TransactOpts, admin, burn, miner, root)
}

// SetCustomedSubSetting is a paid mutator transaction binding the contract method 0x1918f4ef.
//
// Solidity: function setCustomedSubSetting(uint256 admin, uint256 burn, uint256 miner, uint256 root) returns()
func (_BasMiner *BasMinerTransactorSession) SetCustomedSubSetting(admin *big.Int, burn *big.Int, miner *big.Int, root *big.Int) (*types.Transaction, error) {
	return _BasMiner.Contract.SetCustomedSubSetting(&_BasMiner.TransactOpts, admin, burn, miner, root)
}

// SetDefaultSubSetting is a paid mutator transaction binding the contract method 0x0767a424.
//
// Solidity: function setDefaultSubSetting(uint256 admin, uint256 burn, uint256 miner, uint256 root) returns()
func (_BasMiner *BasMinerTransactor) SetDefaultSubSetting(opts *bind.TransactOpts, admin *big.Int, burn *big.Int, miner *big.Int, root *big.Int) (*types.Transaction, error) {
	return _BasMiner.contract.Transact(opts, "setDefaultSubSetting", admin, burn, miner, root)
}

// SetDefaultSubSetting is a paid mutator transaction binding the contract method 0x0767a424.
//
// Solidity: function setDefaultSubSetting(uint256 admin, uint256 burn, uint256 miner, uint256 root) returns()
func (_BasMiner *BasMinerSession) SetDefaultSubSetting(admin *big.Int, burn *big.Int, miner *big.Int, root *big.Int) (*types.Transaction, error) {
	return _BasMiner.Contract.SetDefaultSubSetting(&_BasMiner.TransactOpts, admin, burn, miner, root)
}

// SetDefaultSubSetting is a paid mutator transaction binding the contract method 0x0767a424.
//
// Solidity: function setDefaultSubSetting(uint256 admin, uint256 burn, uint256 miner, uint256 root) returns()
func (_BasMiner *BasMinerTransactorSession) SetDefaultSubSetting(admin *big.Int, burn *big.Int, miner *big.Int, root *big.Int) (*types.Transaction, error) {
	return _BasMiner.Contract.SetDefaultSubSetting(&_BasMiner.TransactOpts, admin, burn, miner, root)
}

// SetOANN is a paid mutator transaction binding the contract method 0xbc63d8b9.
//
// Solidity: function setOANN(address oann) returns()
func (_BasMiner *BasMinerTransactor) SetOANN(opts *bind.TransactOpts, oann common.Address) (*types.Transaction, error) {
	return _BasMiner.contract.Transact(opts, "setOANN", oann)
}

// SetOANN is a paid mutator transaction binding the contract method 0xbc63d8b9.
//
// Solidity: function setOANN(address oann) returns()
func (_BasMiner *BasMinerSession) SetOANN(oann common.Address) (*types.Transaction, error) {
	return _BasMiner.Contract.SetOANN(&_BasMiner.TransactOpts, oann)
}

// SetOANN is a paid mutator transaction binding the contract method 0xbc63d8b9.
//
// Solidity: function setOANN(address oann) returns()
func (_BasMiner *BasMinerTransactorSession) SetOANN(oann common.Address) (*types.Transaction, error) {
	return _BasMiner.Contract.SetOANN(&_BasMiner.TransactOpts, oann)
}

// SetRootSetting is a paid mutator transaction binding the contract method 0x5d89ae36.
//
// Solidity: function setRootSetting(uint256 admin, uint256 burn, uint256 miner) returns()
func (_BasMiner *BasMinerTransactor) SetRootSetting(opts *bind.TransactOpts, admin *big.Int, burn *big.Int, miner *big.Int) (*types.Transaction, error) {
	return _BasMiner.contract.Transact(opts, "setRootSetting", admin, burn, miner)
}

// SetRootSetting is a paid mutator transaction binding the contract method 0x5d89ae36.
//
// Solidity: function setRootSetting(uint256 admin, uint256 burn, uint256 miner) returns()
func (_BasMiner *BasMinerSession) SetRootSetting(admin *big.Int, burn *big.Int, miner *big.Int) (*types.Transaction, error) {
	return _BasMiner.Contract.SetRootSetting(&_BasMiner.TransactOpts, admin, burn, miner)
}

// SetRootSetting is a paid mutator transaction binding the contract method 0x5d89ae36.
//
// Solidity: function setRootSetting(uint256 admin, uint256 burn, uint256 miner) returns()
func (_BasMiner *BasMinerTransactorSession) SetRootSetting(admin *big.Int, burn *big.Int, miner *big.Int) (*types.Transaction, error) {
	return _BasMiner.Contract.SetRootSetting(&_BasMiner.TransactOpts, admin, burn, miner)
}

// SubNameProfit is a paid mutator transaction binding the contract method 0xafd878d8.
//
// Solidity: function subNameProfit(uint256 cost, uint8 typ, address rOwner, bytes32 receiptNumber) returns()
func (_BasMiner *BasMinerTransactor) SubNameProfit(opts *bind.TransactOpts, cost *big.Int, typ uint8, rOwner common.Address, receiptNumber [32]byte) (*types.Transaction, error) {
	return _BasMiner.contract.Transact(opts, "subNameProfit", cost, typ, rOwner, receiptNumber)
}

// SubNameProfit is a paid mutator transaction binding the contract method 0xafd878d8.
//
// Solidity: function subNameProfit(uint256 cost, uint8 typ, address rOwner, bytes32 receiptNumber) returns()
func (_BasMiner *BasMinerSession) SubNameProfit(cost *big.Int, typ uint8, rOwner common.Address, receiptNumber [32]byte) (*types.Transaction, error) {
	return _BasMiner.Contract.SubNameProfit(&_BasMiner.TransactOpts, cost, typ, rOwner, receiptNumber)
}

// SubNameProfit is a paid mutator transaction binding the contract method 0xafd878d8.
//
// Solidity: function subNameProfit(uint256 cost, uint8 typ, address rOwner, bytes32 receiptNumber) returns()
func (_BasMiner *BasMinerTransactorSession) SubNameProfit(cost *big.Int, typ uint8, rOwner common.Address, receiptNumber [32]byte) (*types.Transaction, error) {
	return _BasMiner.Contract.SubNameProfit(&_BasMiner.TransactOpts, cost, typ, rOwner, receiptNumber)
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
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAllocationChanged is a free log retrieval operation binding the contract event 0x315c884dd6c5858ebe0d5c16dc4eed2fc70df851669a1dbd221990202eaa733c.
//
// Solidity: event AllocationChanged()
func (_BasMiner *BasMinerFilterer) FilterAllocationChanged(opts *bind.FilterOpts) (*BasMinerAllocationChangedIterator, error) {

	logs, sub, err := _BasMiner.contract.FilterLogs(opts, "AllocationChanged")
	if err != nil {
		return nil, err
	}
	return &BasMinerAllocationChangedIterator{contract: _BasMiner.contract, event: "AllocationChanged", logs: logs, sub: sub}, nil
}

// WatchAllocationChanged is a free log subscription operation binding the contract event 0x315c884dd6c5858ebe0d5c16dc4eed2fc70df851669a1dbd221990202eaa733c.
//
// Solidity: event AllocationChanged()
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

// ParseAllocationChanged is a log parse operation binding the contract event 0x315c884dd6c5858ebe0d5c16dc4eed2fc70df851669a1dbd221990202eaa733c.
//
// Solidity: event AllocationChanged()
func (_BasMiner *BasMinerFilterer) ParseAllocationChanged(log types.Log) (*BasMinerAllocationChanged, error) {
	event := new(BasMinerAllocationChanged)
	if err := _BasMiner.contract.UnpackLog(event, "AllocationChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BasMinerMinerChangedIterator is returned from FilterMinerChanged and is used to iterate over the raw logs and unpacked data for MinerChanged events raised by the BasMiner contract.
type BasMinerMinerChangedIterator struct {
	Event *BasMinerMinerChanged // Event containing the contract specifics and raw log

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
func (it *BasMinerMinerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasMinerMinerChanged)
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
		it.Event = new(BasMinerMinerChanged)
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
func (it *BasMinerMinerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasMinerMinerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasMinerMinerChanged represents a MinerChanged event raised by the BasMiner contract.
type BasMinerMinerChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterMinerChanged is a free log retrieval operation binding the contract event 0x7e308cc084e00d620063a76a644739d79059ab4b54ac41e7739b3635e569445b.
//
// Solidity: event MinerChanged()
func (_BasMiner *BasMinerFilterer) FilterMinerChanged(opts *bind.FilterOpts) (*BasMinerMinerChangedIterator, error) {

	logs, sub, err := _BasMiner.contract.FilterLogs(opts, "MinerChanged")
	if err != nil {
		return nil, err
	}
	return &BasMinerMinerChangedIterator{contract: _BasMiner.contract, event: "MinerChanged", logs: logs, sub: sub}, nil
}

// WatchMinerChanged is a free log subscription operation binding the contract event 0x7e308cc084e00d620063a76a644739d79059ab4b54ac41e7739b3635e569445b.
//
// Solidity: event MinerChanged()
func (_BasMiner *BasMinerFilterer) WatchMinerChanged(opts *bind.WatchOpts, sink chan<- *BasMinerMinerChanged) (event.Subscription, error) {

	logs, sub, err := _BasMiner.contract.WatchLogs(opts, "MinerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasMinerMinerChanged)
				if err := _BasMiner.contract.UnpackLog(event, "MinerChanged", log); err != nil {
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

// ParseMinerChanged is a log parse operation binding the contract event 0x7e308cc084e00d620063a76a644739d79059ab4b54ac41e7739b3635e569445b.
//
// Solidity: event MinerChanged()
func (_BasMiner *BasMinerFilterer) ParseMinerChanged(log types.Log) (*BasMinerMinerChanged, error) {
	event := new(BasMinerMinerChanged)
	if err := _BasMiner.contract.UnpackLog(event, "MinerChanged", log); err != nil {
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
	From          common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterReceipt is a free log retrieval operation binding the contract event 0xe0279b2669a15826965c3f2418236311d0e48e112d0f8302163e3ddc159c2dba.
//
// Solidity: event Receipt(bytes32 receiptNumber, address from)
func (_BasMiner *BasMinerFilterer) FilterReceipt(opts *bind.FilterOpts) (*BasMinerReceiptIterator, error) {

	logs, sub, err := _BasMiner.contract.FilterLogs(opts, "Receipt")
	if err != nil {
		return nil, err
	}
	return &BasMinerReceiptIterator{contract: _BasMiner.contract, event: "Receipt", logs: logs, sub: sub}, nil
}

// WatchReceipt is a free log subscription operation binding the contract event 0xe0279b2669a15826965c3f2418236311d0e48e112d0f8302163e3ddc159c2dba.
//
// Solidity: event Receipt(bytes32 receiptNumber, address from)
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

// ParseReceipt is a log parse operation binding the contract event 0xe0279b2669a15826965c3f2418236311d0e48e112d0f8302163e3ddc159c2dba.
//
// Solidity: event Receipt(bytes32 receiptNumber, address from)
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
