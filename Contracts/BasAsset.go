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

// BasAssetABI is the input ABI used to generate the binding from.
const BasAssetABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"getRootRechargeInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"}],\"name\":\"clearRecord\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expire\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isOpen\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isCustomed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"cusPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"takeoverRoot\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expire\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isRoot\",\"type\":\"bool\"}],\"name\":\"recharge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalDomainSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"}],\"name\":\"DnsDetailsByHash\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"expire\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"ipv4\",\"type\":\"bytes4\"},{\"internalType\":\"bytes16\",\"name\":\"ipv6\",\"type\":\"bytes16\"},{\"internalType\":\"string\",\"name\":\"bcAddr\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"opData\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"aName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"rootHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"openCustomedPrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"getMayAssetIndexOf\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"aName\",\"type\":\"string\"}],\"name\":\"setAlias\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"EXTEND_LEN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes4\",\"name\":\"ipv4\",\"type\":\"bytes4\"},{\"internalType\":\"bytes16\",\"name\":\"ipv6\",\"type\":\"bytes16\"}],\"name\":\"setIP\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ContractOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BADDRESS_LEN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"DnsDetailsByIndex\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"expire\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"ipv4\",\"type\":\"bytes4\"},{\"internalType\":\"bytes16\",\"name\":\"ipv6\",\"type\":\"bytes16\"},{\"internalType\":\"string\",\"name\":\"bcAddr\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"opData\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"aName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"bcAddress\",\"type\":\"string\"}],\"name\":\"setBCAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"isRoot\",\"type\":\"bool\"}],\"name\":\"isExpired\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"rootHash\",\"type\":\"bytes32\"}],\"name\":\"getRootSetting\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"}],\"name\":\"GetExpire\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"}],\"name\":\"revoke\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes4\",\"name\":\"ipv4\",\"type\":\"bytes4\"},{\"internalType\":\"bytes16\",\"name\":\"ipv6\",\"type\":\"bytes16\"},{\"internalType\":\"string\",\"name\":\"bcAddress\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"opData\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"aName\",\"type\":\"string\"}],\"name\":\"setRecord\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"myAssetCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"opData\",\"type\":\"bytes\"}],\"name\":\"setOpData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferContractOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"getSubRechargeInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOANN\",\"type\":\"address\"}],\"name\":\"setOANN\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"Hash\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"rootHash\",\"type\":\"bytes32\"}],\"name\":\"closeCustomedPrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"getDomainOfIndex\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"sname\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"rHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expire\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"mintSubAsset\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"rootHash\",\"type\":\"bytes32\"}],\"name\":\"closeToPublic\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expire\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"takeoverSubName\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"Alias_LEN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"rootHash\",\"type\":\"bytes32\"}],\"name\":\"openToPublic\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"expire\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isOpen\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isCustomed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"cusPrice\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isAType\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"mintRootAsset\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"}],\"name\":\"MintAsset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"TakeoverAsset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"RechargeAsset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"}],\"name\":\"DNSRecordChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"}],\"name\":\"DNSRecordRemove\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"}],\"name\":\"RootAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"}],\"name\":\"RootChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"}],\"name\":\"SubAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"}],\"name\":\"SubChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"AssertTransfer\",\"type\":\"event\"}]"

// BasAsset is an auto generated Go binding around an Ethereum contract.
type BasAsset struct {
	BasAssetCaller     // Read-only binding to the contract
	BasAssetTransactor // Write-only binding to the contract
	BasAssetFilterer   // Log filterer for contract events
}

// BasAssetCaller is an auto generated read-only Go binding around an Ethereum contract.
type BasAssetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasAssetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BasAssetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasAssetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BasAssetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasAssetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BasAssetSession struct {
	Contract     *BasAsset         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BasAssetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BasAssetCallerSession struct {
	Contract *BasAssetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BasAssetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BasAssetTransactorSession struct {
	Contract     *BasAssetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BasAssetRaw is an auto generated low-level Go binding around an Ethereum contract.
type BasAssetRaw struct {
	Contract *BasAsset // Generic contract binding to access the raw methods on
}

// BasAssetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BasAssetCallerRaw struct {
	Contract *BasAssetCaller // Generic read-only contract binding to access the raw methods on
}

// BasAssetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BasAssetTransactorRaw struct {
	Contract *BasAssetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBasAsset creates a new instance of BasAsset, bound to a specific deployed contract.
func NewBasAsset(address common.Address, backend bind.ContractBackend) (*BasAsset, error) {
	contract, err := bindBasAsset(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BasAsset{BasAssetCaller: BasAssetCaller{contract: contract}, BasAssetTransactor: BasAssetTransactor{contract: contract}, BasAssetFilterer: BasAssetFilterer{contract: contract}}, nil
}

// NewBasAssetCaller creates a new read-only instance of BasAsset, bound to a specific deployed contract.
func NewBasAssetCaller(address common.Address, caller bind.ContractCaller) (*BasAssetCaller, error) {
	contract, err := bindBasAsset(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BasAssetCaller{contract: contract}, nil
}

// NewBasAssetTransactor creates a new write-only instance of BasAsset, bound to a specific deployed contract.
func NewBasAssetTransactor(address common.Address, transactor bind.ContractTransactor) (*BasAssetTransactor, error) {
	contract, err := bindBasAsset(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BasAssetTransactor{contract: contract}, nil
}

// NewBasAssetFilterer creates a new log filterer instance of BasAsset, bound to a specific deployed contract.
func NewBasAssetFilterer(address common.Address, filterer bind.ContractFilterer) (*BasAssetFilterer, error) {
	contract, err := bindBasAsset(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BasAssetFilterer{contract: contract}, nil
}

// bindBasAsset binds a generic wrapper to an already deployed contract.
func bindBasAsset(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BasAssetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasAsset *BasAssetRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BasAsset.Contract.BasAssetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasAsset *BasAssetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasAsset.Contract.BasAssetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasAsset *BasAssetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasAsset.Contract.BasAssetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasAsset *BasAssetCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BasAsset.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasAsset *BasAssetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasAsset.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasAsset *BasAssetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasAsset.Contract.contract.Transact(opts, method, params...)
}

// AliasLEN is a free data retrieval call binding the contract method 0xda203e69.
//
// Solidity: function Alias_LEN() constant returns(uint256)
func (_BasAsset *BasAssetCaller) AliasLEN(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BasAsset.contract.Call(opts, out, "Alias_LEN")
	return *ret0, err
}

// AliasLEN is a free data retrieval call binding the contract method 0xda203e69.
//
// Solidity: function Alias_LEN() constant returns(uint256)
func (_BasAsset *BasAssetSession) AliasLEN() (*big.Int, error) {
	return _BasAsset.Contract.AliasLEN(&_BasAsset.CallOpts)
}

// AliasLEN is a free data retrieval call binding the contract method 0xda203e69.
//
// Solidity: function Alias_LEN() constant returns(uint256)
func (_BasAsset *BasAssetCallerSession) AliasLEN() (*big.Int, error) {
	return _BasAsset.Contract.AliasLEN(&_BasAsset.CallOpts)
}

// BADDRESSLEN is a free data retrieval call binding the contract method 0x5ce6f0a4.
//
// Solidity: function BADDRESS_LEN() constant returns(uint256)
func (_BasAsset *BasAssetCaller) BADDRESSLEN(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BasAsset.contract.Call(opts, out, "BADDRESS_LEN")
	return *ret0, err
}

// BADDRESSLEN is a free data retrieval call binding the contract method 0x5ce6f0a4.
//
// Solidity: function BADDRESS_LEN() constant returns(uint256)
func (_BasAsset *BasAssetSession) BADDRESSLEN() (*big.Int, error) {
	return _BasAsset.Contract.BADDRESSLEN(&_BasAsset.CallOpts)
}

// BADDRESSLEN is a free data retrieval call binding the contract method 0x5ce6f0a4.
//
// Solidity: function BADDRESS_LEN() constant returns(uint256)
func (_BasAsset *BasAssetCallerSession) BADDRESSLEN() (*big.Int, error) {
	return _BasAsset.Contract.BADDRESSLEN(&_BasAsset.CallOpts)
}

// ContractOwner is a free data retrieval call binding the contract method 0x5a63fbc9.
//
// Solidity: function ContractOwner() constant returns(address)
func (_BasAsset *BasAssetCaller) ContractOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BasAsset.contract.Call(opts, out, "ContractOwner")
	return *ret0, err
}

// ContractOwner is a free data retrieval call binding the contract method 0x5a63fbc9.
//
// Solidity: function ContractOwner() constant returns(address)
func (_BasAsset *BasAssetSession) ContractOwner() (common.Address, error) {
	return _BasAsset.Contract.ContractOwner(&_BasAsset.CallOpts)
}

// ContractOwner is a free data retrieval call binding the contract method 0x5a63fbc9.
//
// Solidity: function ContractOwner() constant returns(address)
func (_BasAsset *BasAssetCallerSession) ContractOwner() (common.Address, error) {
	return _BasAsset.Contract.ContractOwner(&_BasAsset.CallOpts)
}

// DnsDetailsByHash is a free data retrieval call binding the contract method 0x39a9a72e.
//
// Solidity: function DnsDetailsByHash(bytes32 nameHash) constant returns(bytes name, uint256 expire, bytes4 ipv4, bytes16 ipv6, string bcAddr, bytes opData, string aName, address owner)
func (_BasAsset *BasAssetCaller) DnsDetailsByHash(opts *bind.CallOpts, nameHash [32]byte) (struct {
	Name   []byte
	Expire *big.Int
	Ipv4   [4]byte
	Ipv6   [16]byte
	BcAddr string
	OpData []byte
	AName  string
	Owner  common.Address
}, error) {
	ret := new(struct {
		Name   []byte
		Expire *big.Int
		Ipv4   [4]byte
		Ipv6   [16]byte
		BcAddr string
		OpData []byte
		AName  string
		Owner  common.Address
	})
	out := ret
	err := _BasAsset.contract.Call(opts, out, "DnsDetailsByHash", nameHash)
	return *ret, err
}

// DnsDetailsByHash is a free data retrieval call binding the contract method 0x39a9a72e.
//
// Solidity: function DnsDetailsByHash(bytes32 nameHash) constant returns(bytes name, uint256 expire, bytes4 ipv4, bytes16 ipv6, string bcAddr, bytes opData, string aName, address owner)
func (_BasAsset *BasAssetSession) DnsDetailsByHash(nameHash [32]byte) (struct {
	Name   []byte
	Expire *big.Int
	Ipv4   [4]byte
	Ipv6   [16]byte
	BcAddr string
	OpData []byte
	AName  string
	Owner  common.Address
}, error) {
	return _BasAsset.Contract.DnsDetailsByHash(&_BasAsset.CallOpts, nameHash)
}

// DnsDetailsByHash is a free data retrieval call binding the contract method 0x39a9a72e.
//
// Solidity: function DnsDetailsByHash(bytes32 nameHash) constant returns(bytes name, uint256 expire, bytes4 ipv4, bytes16 ipv6, string bcAddr, bytes opData, string aName, address owner)
func (_BasAsset *BasAssetCallerSession) DnsDetailsByHash(nameHash [32]byte) (struct {
	Name   []byte
	Expire *big.Int
	Ipv4   [4]byte
	Ipv6   [16]byte
	BcAddr string
	OpData []byte
	AName  string
	Owner  common.Address
}, error) {
	return _BasAsset.Contract.DnsDetailsByHash(&_BasAsset.CallOpts, nameHash)
}

// DnsDetailsByIndex is a free data retrieval call binding the contract method 0x6ad1fb6a.
//
// Solidity: function DnsDetailsByIndex(uint256 index) constant returns(bytes name, uint256 expire, bytes4 ipv4, bytes16 ipv6, string bcAddr, bytes opData, string aName, address owner)
func (_BasAsset *BasAssetCaller) DnsDetailsByIndex(opts *bind.CallOpts, index *big.Int) (struct {
	Name   []byte
	Expire *big.Int
	Ipv4   [4]byte
	Ipv6   [16]byte
	BcAddr string
	OpData []byte
	AName  string
	Owner  common.Address
}, error) {
	ret := new(struct {
		Name   []byte
		Expire *big.Int
		Ipv4   [4]byte
		Ipv6   [16]byte
		BcAddr string
		OpData []byte
		AName  string
		Owner  common.Address
	})
	out := ret
	err := _BasAsset.contract.Call(opts, out, "DnsDetailsByIndex", index)
	return *ret, err
}

// DnsDetailsByIndex is a free data retrieval call binding the contract method 0x6ad1fb6a.
//
// Solidity: function DnsDetailsByIndex(uint256 index) constant returns(bytes name, uint256 expire, bytes4 ipv4, bytes16 ipv6, string bcAddr, bytes opData, string aName, address owner)
func (_BasAsset *BasAssetSession) DnsDetailsByIndex(index *big.Int) (struct {
	Name   []byte
	Expire *big.Int
	Ipv4   [4]byte
	Ipv6   [16]byte
	BcAddr string
	OpData []byte
	AName  string
	Owner  common.Address
}, error) {
	return _BasAsset.Contract.DnsDetailsByIndex(&_BasAsset.CallOpts, index)
}

// DnsDetailsByIndex is a free data retrieval call binding the contract method 0x6ad1fb6a.
//
// Solidity: function DnsDetailsByIndex(uint256 index) constant returns(bytes name, uint256 expire, bytes4 ipv4, bytes16 ipv6, string bcAddr, bytes opData, string aName, address owner)
func (_BasAsset *BasAssetCallerSession) DnsDetailsByIndex(index *big.Int) (struct {
	Name   []byte
	Expire *big.Int
	Ipv4   [4]byte
	Ipv6   [16]byte
	BcAddr string
	OpData []byte
	AName  string
	Owner  common.Address
}, error) {
	return _BasAsset.Contract.DnsDetailsByIndex(&_BasAsset.CallOpts, index)
}

// EXTENDLEN is a free data retrieval call binding the contract method 0x4b50acc8.
//
// Solidity: function EXTEND_LEN() constant returns(uint256)
func (_BasAsset *BasAssetCaller) EXTENDLEN(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BasAsset.contract.Call(opts, out, "EXTEND_LEN")
	return *ret0, err
}

// EXTENDLEN is a free data retrieval call binding the contract method 0x4b50acc8.
//
// Solidity: function EXTEND_LEN() constant returns(uint256)
func (_BasAsset *BasAssetSession) EXTENDLEN() (*big.Int, error) {
	return _BasAsset.Contract.EXTENDLEN(&_BasAsset.CallOpts)
}

// EXTENDLEN is a free data retrieval call binding the contract method 0x4b50acc8.
//
// Solidity: function EXTEND_LEN() constant returns(uint256)
func (_BasAsset *BasAssetCallerSession) EXTENDLEN() (*big.Int, error) {
	return _BasAsset.Contract.EXTENDLEN(&_BasAsset.CallOpts)
}

// GetExpire is a free data retrieval call binding the contract method 0x8303ed38.
//
// Solidity: function GetExpire(bytes32 nameHash) constant returns(uint256)
func (_BasAsset *BasAssetCaller) GetExpire(opts *bind.CallOpts, nameHash [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BasAsset.contract.Call(opts, out, "GetExpire", nameHash)
	return *ret0, err
}

// GetExpire is a free data retrieval call binding the contract method 0x8303ed38.
//
// Solidity: function GetExpire(bytes32 nameHash) constant returns(uint256)
func (_BasAsset *BasAssetSession) GetExpire(nameHash [32]byte) (*big.Int, error) {
	return _BasAsset.Contract.GetExpire(&_BasAsset.CallOpts, nameHash)
}

// GetExpire is a free data retrieval call binding the contract method 0x8303ed38.
//
// Solidity: function GetExpire(bytes32 nameHash) constant returns(uint256)
func (_BasAsset *BasAssetCallerSession) GetExpire(nameHash [32]byte) (*big.Int, error) {
	return _BasAsset.Contract.GetExpire(&_BasAsset.CallOpts, nameHash)
}

// Hash is a free data retrieval call binding the contract method 0xbc780a0c.
//
// Solidity: function Hash(string name) constant returns(bytes, bytes32)
func (_BasAsset *BasAssetCaller) Hash(opts *bind.CallOpts, name string) ([]byte, [32]byte, error) {
	var (
		ret0 = new([]byte)
		ret1 = new([32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _BasAsset.contract.Call(opts, out, "Hash", name)
	return *ret0, *ret1, err
}

// Hash is a free data retrieval call binding the contract method 0xbc780a0c.
//
// Solidity: function Hash(string name) constant returns(bytes, bytes32)
func (_BasAsset *BasAssetSession) Hash(name string) ([]byte, [32]byte, error) {
	return _BasAsset.Contract.Hash(&_BasAsset.CallOpts, name)
}

// Hash is a free data retrieval call binding the contract method 0xbc780a0c.
//
// Solidity: function Hash(string name) constant returns(bytes, bytes32)
func (_BasAsset *BasAssetCallerSession) Hash(name string) ([]byte, [32]byte, error) {
	return _BasAsset.Contract.Hash(&_BasAsset.CallOpts, name)
}

// Allowance is a free data retrieval call binding the contract method 0xbcf9fcad.
//
// Solidity: function allowance(address owner, bytes32 nameHash) constant returns(address)
func (_BasAsset *BasAssetCaller) Allowance(opts *bind.CallOpts, owner common.Address, nameHash [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BasAsset.contract.Call(opts, out, "allowance", owner, nameHash)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xbcf9fcad.
//
// Solidity: function allowance(address owner, bytes32 nameHash) constant returns(address)
func (_BasAsset *BasAssetSession) Allowance(owner common.Address, nameHash [32]byte) (common.Address, error) {
	return _BasAsset.Contract.Allowance(&_BasAsset.CallOpts, owner, nameHash)
}

// Allowance is a free data retrieval call binding the contract method 0xbcf9fcad.
//
// Solidity: function allowance(address owner, bytes32 nameHash) constant returns(address)
func (_BasAsset *BasAssetCallerSession) Allowance(owner common.Address, nameHash [32]byte) (common.Address, error) {
	return _BasAsset.Contract.Allowance(&_BasAsset.CallOpts, owner, nameHash)
}

// GetDomainOfIndex is a free data retrieval call binding the contract method 0xc5efdc3d.
//
// Solidity: function getDomainOfIndex(uint256 idx) constant returns(bytes32)
func (_BasAsset *BasAssetCaller) GetDomainOfIndex(opts *bind.CallOpts, idx *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _BasAsset.contract.Call(opts, out, "getDomainOfIndex", idx)
	return *ret0, err
}

// GetDomainOfIndex is a free data retrieval call binding the contract method 0xc5efdc3d.
//
// Solidity: function getDomainOfIndex(uint256 idx) constant returns(bytes32)
func (_BasAsset *BasAssetSession) GetDomainOfIndex(idx *big.Int) ([32]byte, error) {
	return _BasAsset.Contract.GetDomainOfIndex(&_BasAsset.CallOpts, idx)
}

// GetDomainOfIndex is a free data retrieval call binding the contract method 0xc5efdc3d.
//
// Solidity: function getDomainOfIndex(uint256 idx) constant returns(bytes32)
func (_BasAsset *BasAssetCallerSession) GetDomainOfIndex(idx *big.Int) ([32]byte, error) {
	return _BasAsset.Contract.GetDomainOfIndex(&_BasAsset.CallOpts, idx)
}

// GetMayAssetIndexOf is a free data retrieval call binding the contract method 0x3db8b32c.
//
// Solidity: function getMayAssetIndexOf(uint256 idx) constant returns(bytes32)
func (_BasAsset *BasAssetCaller) GetMayAssetIndexOf(opts *bind.CallOpts, idx *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _BasAsset.contract.Call(opts, out, "getMayAssetIndexOf", idx)
	return *ret0, err
}

// GetMayAssetIndexOf is a free data retrieval call binding the contract method 0x3db8b32c.
//
// Solidity: function getMayAssetIndexOf(uint256 idx) constant returns(bytes32)
func (_BasAsset *BasAssetSession) GetMayAssetIndexOf(idx *big.Int) ([32]byte, error) {
	return _BasAsset.Contract.GetMayAssetIndexOf(&_BasAsset.CallOpts, idx)
}

// GetMayAssetIndexOf is a free data retrieval call binding the contract method 0x3db8b32c.
//
// Solidity: function getMayAssetIndexOf(uint256 idx) constant returns(bytes32)
func (_BasAsset *BasAssetCallerSession) GetMayAssetIndexOf(idx *big.Int) ([32]byte, error) {
	return _BasAsset.Contract.GetMayAssetIndexOf(&_BasAsset.CallOpts, idx)
}

// GetRootRechargeInfo is a free data retrieval call binding the contract method 0x04607bb4.
//
// Solidity: function getRootRechargeInfo(bytes32 hash) constant returns(uint256, bool)
func (_BasAsset *BasAssetCaller) GetRootRechargeInfo(opts *bind.CallOpts, hash [32]byte) (*big.Int, bool, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _BasAsset.contract.Call(opts, out, "getRootRechargeInfo", hash)
	return *ret0, *ret1, err
}

// GetRootRechargeInfo is a free data retrieval call binding the contract method 0x04607bb4.
//
// Solidity: function getRootRechargeInfo(bytes32 hash) constant returns(uint256, bool)
func (_BasAsset *BasAssetSession) GetRootRechargeInfo(hash [32]byte) (*big.Int, bool, error) {
	return _BasAsset.Contract.GetRootRechargeInfo(&_BasAsset.CallOpts, hash)
}

// GetRootRechargeInfo is a free data retrieval call binding the contract method 0x04607bb4.
//
// Solidity: function getRootRechargeInfo(bytes32 hash) constant returns(uint256, bool)
func (_BasAsset *BasAssetCallerSession) GetRootRechargeInfo(hash [32]byte) (*big.Int, bool, error) {
	return _BasAsset.Contract.GetRootRechargeInfo(&_BasAsset.CallOpts, hash)
}

// GetRootSetting is a free data retrieval call binding the contract method 0x7a6bbd99.
//
// Solidity: function getRootSetting(bytes32 rootHash) constant returns(uint256, bool, bool)
func (_BasAsset *BasAssetCaller) GetRootSetting(opts *bind.CallOpts, rootHash [32]byte) (*big.Int, bool, bool, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(bool)
		ret2 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _BasAsset.contract.Call(opts, out, "getRootSetting", rootHash)
	return *ret0, *ret1, *ret2, err
}

// GetRootSetting is a free data retrieval call binding the contract method 0x7a6bbd99.
//
// Solidity: function getRootSetting(bytes32 rootHash) constant returns(uint256, bool, bool)
func (_BasAsset *BasAssetSession) GetRootSetting(rootHash [32]byte) (*big.Int, bool, bool, error) {
	return _BasAsset.Contract.GetRootSetting(&_BasAsset.CallOpts, rootHash)
}

// GetRootSetting is a free data retrieval call binding the contract method 0x7a6bbd99.
//
// Solidity: function getRootSetting(bytes32 rootHash) constant returns(uint256, bool, bool)
func (_BasAsset *BasAssetCallerSession) GetRootSetting(rootHash [32]byte) (*big.Int, bool, bool, error) {
	return _BasAsset.Contract.GetRootSetting(&_BasAsset.CallOpts, rootHash)
}

// GetSubRechargeInfo is a free data retrieval call binding the contract method 0xbbe472b8.
//
// Solidity: function getSubRechargeInfo(bytes32 hash) constant returns(uint256, bytes32)
func (_BasAsset *BasAssetCaller) GetSubRechargeInfo(opts *bind.CallOpts, hash [32]byte) (*big.Int, [32]byte, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new([32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _BasAsset.contract.Call(opts, out, "getSubRechargeInfo", hash)
	return *ret0, *ret1, err
}

// GetSubRechargeInfo is a free data retrieval call binding the contract method 0xbbe472b8.
//
// Solidity: function getSubRechargeInfo(bytes32 hash) constant returns(uint256, bytes32)
func (_BasAsset *BasAssetSession) GetSubRechargeInfo(hash [32]byte) (*big.Int, [32]byte, error) {
	return _BasAsset.Contract.GetSubRechargeInfo(&_BasAsset.CallOpts, hash)
}

// GetSubRechargeInfo is a free data retrieval call binding the contract method 0xbbe472b8.
//
// Solidity: function getSubRechargeInfo(bytes32 hash) constant returns(uint256, bytes32)
func (_BasAsset *BasAssetCallerSession) GetSubRechargeInfo(hash [32]byte) (*big.Int, [32]byte, error) {
	return _BasAsset.Contract.GetSubRechargeInfo(&_BasAsset.CallOpts, hash)
}

// IsExpired is a free data retrieval call binding the contract method 0x73d5b5cd.
//
// Solidity: function isExpired(bytes32 nameHash, bool isRoot) constant returns(bool)
func (_BasAsset *BasAssetCaller) IsExpired(opts *bind.CallOpts, nameHash [32]byte, isRoot bool) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BasAsset.contract.Call(opts, out, "isExpired", nameHash, isRoot)
	return *ret0, err
}

// IsExpired is a free data retrieval call binding the contract method 0x73d5b5cd.
//
// Solidity: function isExpired(bytes32 nameHash, bool isRoot) constant returns(bool)
func (_BasAsset *BasAssetSession) IsExpired(nameHash [32]byte, isRoot bool) (bool, error) {
	return _BasAsset.Contract.IsExpired(&_BasAsset.CallOpts, nameHash, isRoot)
}

// IsExpired is a free data retrieval call binding the contract method 0x73d5b5cd.
//
// Solidity: function isExpired(bytes32 nameHash, bool isRoot) constant returns(bool)
func (_BasAsset *BasAssetCallerSession) IsExpired(nameHash [32]byte, isRoot bool) (bool, error) {
	return _BasAsset.Contract.IsExpired(&_BasAsset.CallOpts, nameHash, isRoot)
}

// MyAssetCount is a free data retrieval call binding the contract method 0x96c9de20.
//
// Solidity: function myAssetCount(address owner) constant returns(uint256)
func (_BasAsset *BasAssetCaller) MyAssetCount(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BasAsset.contract.Call(opts, out, "myAssetCount", owner)
	return *ret0, err
}

// MyAssetCount is a free data retrieval call binding the contract method 0x96c9de20.
//
// Solidity: function myAssetCount(address owner) constant returns(uint256)
func (_BasAsset *BasAssetSession) MyAssetCount(owner common.Address) (*big.Int, error) {
	return _BasAsset.Contract.MyAssetCount(&_BasAsset.CallOpts, owner)
}

// MyAssetCount is a free data retrieval call binding the contract method 0x96c9de20.
//
// Solidity: function myAssetCount(address owner) constant returns(uint256)
func (_BasAsset *BasAssetCallerSession) MyAssetCount(owner common.Address) (*big.Int, error) {
	return _BasAsset.Contract.MyAssetCount(&_BasAsset.CallOpts, owner)
}

// OwnerOf is a free data retrieval call binding the contract method 0x7dd56411.
//
// Solidity: function ownerOf(bytes32 nameHash) constant returns(address)
func (_BasAsset *BasAssetCaller) OwnerOf(opts *bind.CallOpts, nameHash [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BasAsset.contract.Call(opts, out, "ownerOf", nameHash)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x7dd56411.
//
// Solidity: function ownerOf(bytes32 nameHash) constant returns(address)
func (_BasAsset *BasAssetSession) OwnerOf(nameHash [32]byte) (common.Address, error) {
	return _BasAsset.Contract.OwnerOf(&_BasAsset.CallOpts, nameHash)
}

// OwnerOf is a free data retrieval call binding the contract method 0x7dd56411.
//
// Solidity: function ownerOf(bytes32 nameHash) constant returns(address)
func (_BasAsset *BasAssetCallerSession) OwnerOf(nameHash [32]byte) (common.Address, error) {
	return _BasAsset.Contract.OwnerOf(&_BasAsset.CallOpts, nameHash)
}

// TotalDomainSize is a free data retrieval call binding the contract method 0x24ea2e6b.
//
// Solidity: function totalDomainSize() constant returns(uint256)
func (_BasAsset *BasAssetCaller) TotalDomainSize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BasAsset.contract.Call(opts, out, "totalDomainSize")
	return *ret0, err
}

// TotalDomainSize is a free data retrieval call binding the contract method 0x24ea2e6b.
//
// Solidity: function totalDomainSize() constant returns(uint256)
func (_BasAsset *BasAssetSession) TotalDomainSize() (*big.Int, error) {
	return _BasAsset.Contract.TotalDomainSize(&_BasAsset.CallOpts)
}

// TotalDomainSize is a free data retrieval call binding the contract method 0x24ea2e6b.
//
// Solidity: function totalDomainSize() constant returns(uint256)
func (_BasAsset *BasAssetCallerSession) TotalDomainSize() (*big.Int, error) {
	return _BasAsset.Contract.TotalDomainSize(&_BasAsset.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x5cd2f4d3.
//
// Solidity: function approve(address spender, bytes32 nameHash) returns()
func (_BasAsset *BasAssetTransactor) Approve(opts *bind.TransactOpts, spender common.Address, nameHash [32]byte) (*types.Transaction, error) {
	return _BasAsset.contract.Transact(opts, "approve", spender, nameHash)
}

// Approve is a paid mutator transaction binding the contract method 0x5cd2f4d3.
//
// Solidity: function approve(address spender, bytes32 nameHash) returns()
func (_BasAsset *BasAssetSession) Approve(spender common.Address, nameHash [32]byte) (*types.Transaction, error) {
	return _BasAsset.Contract.Approve(&_BasAsset.TransactOpts, spender, nameHash)
}

// Approve is a paid mutator transaction binding the contract method 0x5cd2f4d3.
//
// Solidity: function approve(address spender, bytes32 nameHash) returns()
func (_BasAsset *BasAssetTransactorSession) Approve(spender common.Address, nameHash [32]byte) (*types.Transaction, error) {
	return _BasAsset.Contract.Approve(&_BasAsset.TransactOpts, spender, nameHash)
}

// ClearRecord is a paid mutator transaction binding the contract method 0x081780f4.
//
// Solidity: function clearRecord(bytes32 nameHash) returns()
func (_BasAsset *BasAssetTransactor) ClearRecord(opts *bind.TransactOpts, nameHash [32]byte) (*types.Transaction, error) {
	return _BasAsset.contract.Transact(opts, "clearRecord", nameHash)
}

// ClearRecord is a paid mutator transaction binding the contract method 0x081780f4.
//
// Solidity: function clearRecord(bytes32 nameHash) returns()
func (_BasAsset *BasAssetSession) ClearRecord(nameHash [32]byte) (*types.Transaction, error) {
	return _BasAsset.Contract.ClearRecord(&_BasAsset.TransactOpts, nameHash)
}

// ClearRecord is a paid mutator transaction binding the contract method 0x081780f4.
//
// Solidity: function clearRecord(bytes32 nameHash) returns()
func (_BasAsset *BasAssetTransactorSession) ClearRecord(nameHash [32]byte) (*types.Transaction, error) {
	return _BasAsset.Contract.ClearRecord(&_BasAsset.TransactOpts, nameHash)
}

// CloseCustomedPrice is a paid mutator transaction binding the contract method 0xc0bd516d.
//
// Solidity: function closeCustomedPrice(bytes32 rootHash) returns()
func (_BasAsset *BasAssetTransactor) CloseCustomedPrice(opts *bind.TransactOpts, rootHash [32]byte) (*types.Transaction, error) {
	return _BasAsset.contract.Transact(opts, "closeCustomedPrice", rootHash)
}

// CloseCustomedPrice is a paid mutator transaction binding the contract method 0xc0bd516d.
//
// Solidity: function closeCustomedPrice(bytes32 rootHash) returns()
func (_BasAsset *BasAssetSession) CloseCustomedPrice(rootHash [32]byte) (*types.Transaction, error) {
	return _BasAsset.Contract.CloseCustomedPrice(&_BasAsset.TransactOpts, rootHash)
}

// CloseCustomedPrice is a paid mutator transaction binding the contract method 0xc0bd516d.
//
// Solidity: function closeCustomedPrice(bytes32 rootHash) returns()
func (_BasAsset *BasAssetTransactorSession) CloseCustomedPrice(rootHash [32]byte) (*types.Transaction, error) {
	return _BasAsset.Contract.CloseCustomedPrice(&_BasAsset.TransactOpts, rootHash)
}

// CloseToPublic is a paid mutator transaction binding the contract method 0xd0f0b43d.
//
// Solidity: function closeToPublic(bytes32 rootHash) returns()
func (_BasAsset *BasAssetTransactor) CloseToPublic(opts *bind.TransactOpts, rootHash [32]byte) (*types.Transaction, error) {
	return _BasAsset.contract.Transact(opts, "closeToPublic", rootHash)
}

// CloseToPublic is a paid mutator transaction binding the contract method 0xd0f0b43d.
//
// Solidity: function closeToPublic(bytes32 rootHash) returns()
func (_BasAsset *BasAssetSession) CloseToPublic(rootHash [32]byte) (*types.Transaction, error) {
	return _BasAsset.Contract.CloseToPublic(&_BasAsset.TransactOpts, rootHash)
}

// CloseToPublic is a paid mutator transaction binding the contract method 0xd0f0b43d.
//
// Solidity: function closeToPublic(bytes32 rootHash) returns()
func (_BasAsset *BasAssetTransactorSession) CloseToPublic(rootHash [32]byte) (*types.Transaction, error) {
	return _BasAsset.Contract.CloseToPublic(&_BasAsset.TransactOpts, rootHash)
}

// MintRootAsset is a paid mutator transaction binding the contract method 0xfbc94c0b.
//
// Solidity: function mintRootAsset(bytes32 nameHash, bytes name, uint256 expire, bool isOpen, bool isCustomed, uint256 cusPrice, bool isAType, address owner) returns()
func (_BasAsset *BasAssetTransactor) MintRootAsset(opts *bind.TransactOpts, nameHash [32]byte, name []byte, expire *big.Int, isOpen bool, isCustomed bool, cusPrice *big.Int, isAType bool, owner common.Address) (*types.Transaction, error) {
	return _BasAsset.contract.Transact(opts, "mintRootAsset", nameHash, name, expire, isOpen, isCustomed, cusPrice, isAType, owner)
}

// MintRootAsset is a paid mutator transaction binding the contract method 0xfbc94c0b.
//
// Solidity: function mintRootAsset(bytes32 nameHash, bytes name, uint256 expire, bool isOpen, bool isCustomed, uint256 cusPrice, bool isAType, address owner) returns()
func (_BasAsset *BasAssetSession) MintRootAsset(nameHash [32]byte, name []byte, expire *big.Int, isOpen bool, isCustomed bool, cusPrice *big.Int, isAType bool, owner common.Address) (*types.Transaction, error) {
	return _BasAsset.Contract.MintRootAsset(&_BasAsset.TransactOpts, nameHash, name, expire, isOpen, isCustomed, cusPrice, isAType, owner)
}

// MintRootAsset is a paid mutator transaction binding the contract method 0xfbc94c0b.
//
// Solidity: function mintRootAsset(bytes32 nameHash, bytes name, uint256 expire, bool isOpen, bool isCustomed, uint256 cusPrice, bool isAType, address owner) returns()
func (_BasAsset *BasAssetTransactorSession) MintRootAsset(nameHash [32]byte, name []byte, expire *big.Int, isOpen bool, isCustomed bool, cusPrice *big.Int, isAType bool, owner common.Address) (*types.Transaction, error) {
	return _BasAsset.Contract.MintRootAsset(&_BasAsset.TransactOpts, nameHash, name, expire, isOpen, isCustomed, cusPrice, isAType, owner)
}

// MintSubAsset is a paid mutator transaction binding the contract method 0xcc2312f1.
//
// Solidity: function mintSubAsset(bytes32 sHash, bytes sname, bytes32 rHash, uint256 expire, address owner) returns()
func (_BasAsset *BasAssetTransactor) MintSubAsset(opts *bind.TransactOpts, sHash [32]byte, sname []byte, rHash [32]byte, expire *big.Int, owner common.Address) (*types.Transaction, error) {
	return _BasAsset.contract.Transact(opts, "mintSubAsset", sHash, sname, rHash, expire, owner)
}

// MintSubAsset is a paid mutator transaction binding the contract method 0xcc2312f1.
//
// Solidity: function mintSubAsset(bytes32 sHash, bytes sname, bytes32 rHash, uint256 expire, address owner) returns()
func (_BasAsset *BasAssetSession) MintSubAsset(sHash [32]byte, sname []byte, rHash [32]byte, expire *big.Int, owner common.Address) (*types.Transaction, error) {
	return _BasAsset.Contract.MintSubAsset(&_BasAsset.TransactOpts, sHash, sname, rHash, expire, owner)
}

// MintSubAsset is a paid mutator transaction binding the contract method 0xcc2312f1.
//
// Solidity: function mintSubAsset(bytes32 sHash, bytes sname, bytes32 rHash, uint256 expire, address owner) returns()
func (_BasAsset *BasAssetTransactorSession) MintSubAsset(sHash [32]byte, sname []byte, rHash [32]byte, expire *big.Int, owner common.Address) (*types.Transaction, error) {
	return _BasAsset.Contract.MintSubAsset(&_BasAsset.TransactOpts, sHash, sname, rHash, expire, owner)
}

// OpenCustomedPrice is a paid mutator transaction binding the contract method 0x3d04f2fb.
//
// Solidity: function openCustomedPrice(bytes32 rootHash, uint256 price) returns()
func (_BasAsset *BasAssetTransactor) OpenCustomedPrice(opts *bind.TransactOpts, rootHash [32]byte, price *big.Int) (*types.Transaction, error) {
	return _BasAsset.contract.Transact(opts, "openCustomedPrice", rootHash, price)
}

// OpenCustomedPrice is a paid mutator transaction binding the contract method 0x3d04f2fb.
//
// Solidity: function openCustomedPrice(bytes32 rootHash, uint256 price) returns()
func (_BasAsset *BasAssetSession) OpenCustomedPrice(rootHash [32]byte, price *big.Int) (*types.Transaction, error) {
	return _BasAsset.Contract.OpenCustomedPrice(&_BasAsset.TransactOpts, rootHash, price)
}

// OpenCustomedPrice is a paid mutator transaction binding the contract method 0x3d04f2fb.
//
// Solidity: function openCustomedPrice(bytes32 rootHash, uint256 price) returns()
func (_BasAsset *BasAssetTransactorSession) OpenCustomedPrice(rootHash [32]byte, price *big.Int) (*types.Transaction, error) {
	return _BasAsset.Contract.OpenCustomedPrice(&_BasAsset.TransactOpts, rootHash, price)
}

// OpenToPublic is a paid mutator transaction binding the contract method 0xed26e5ae.
//
// Solidity: function openToPublic(bytes32 rootHash) returns()
func (_BasAsset *BasAssetTransactor) OpenToPublic(opts *bind.TransactOpts, rootHash [32]byte) (*types.Transaction, error) {
	return _BasAsset.contract.Transact(opts, "openToPublic", rootHash)
}

// OpenToPublic is a paid mutator transaction binding the contract method 0xed26e5ae.
//
// Solidity: function openToPublic(bytes32 rootHash) returns()
func (_BasAsset *BasAssetSession) OpenToPublic(rootHash [32]byte) (*types.Transaction, error) {
	return _BasAsset.Contract.OpenToPublic(&_BasAsset.TransactOpts, rootHash)
}

// OpenToPublic is a paid mutator transaction binding the contract method 0xed26e5ae.
//
// Solidity: function openToPublic(bytes32 rootHash) returns()
func (_BasAsset *BasAssetTransactorSession) OpenToPublic(rootHash [32]byte) (*types.Transaction, error) {
	return _BasAsset.Contract.OpenToPublic(&_BasAsset.TransactOpts, rootHash)
}

// Recharge is a paid mutator transaction binding the contract method 0x1c31e2da.
//
// Solidity: function recharge(bytes32 nameHash, uint256 expire, bool isRoot) returns()
func (_BasAsset *BasAssetTransactor) Recharge(opts *bind.TransactOpts, nameHash [32]byte, expire *big.Int, isRoot bool) (*types.Transaction, error) {
	return _BasAsset.contract.Transact(opts, "recharge", nameHash, expire, isRoot)
}

// Recharge is a paid mutator transaction binding the contract method 0x1c31e2da.
//
// Solidity: function recharge(bytes32 nameHash, uint256 expire, bool isRoot) returns()
func (_BasAsset *BasAssetSession) Recharge(nameHash [32]byte, expire *big.Int, isRoot bool) (*types.Transaction, error) {
	return _BasAsset.Contract.Recharge(&_BasAsset.TransactOpts, nameHash, expire, isRoot)
}

// Recharge is a paid mutator transaction binding the contract method 0x1c31e2da.
//
// Solidity: function recharge(bytes32 nameHash, uint256 expire, bool isRoot) returns()
func (_BasAsset *BasAssetTransactorSession) Recharge(nameHash [32]byte, expire *big.Int, isRoot bool) (*types.Transaction, error) {
	return _BasAsset.Contract.Recharge(&_BasAsset.TransactOpts, nameHash, expire, isRoot)
}

// Revoke is a paid mutator transaction binding the contract method 0x88e62721.
//
// Solidity: function revoke(address from, bytes32 nameHash) returns()
func (_BasAsset *BasAssetTransactor) Revoke(opts *bind.TransactOpts, from common.Address, nameHash [32]byte) (*types.Transaction, error) {
	return _BasAsset.contract.Transact(opts, "revoke", from, nameHash)
}

// Revoke is a paid mutator transaction binding the contract method 0x88e62721.
//
// Solidity: function revoke(address from, bytes32 nameHash) returns()
func (_BasAsset *BasAssetSession) Revoke(from common.Address, nameHash [32]byte) (*types.Transaction, error) {
	return _BasAsset.Contract.Revoke(&_BasAsset.TransactOpts, from, nameHash)
}

// Revoke is a paid mutator transaction binding the contract method 0x88e62721.
//
// Solidity: function revoke(address from, bytes32 nameHash) returns()
func (_BasAsset *BasAssetTransactorSession) Revoke(from common.Address, nameHash [32]byte) (*types.Transaction, error) {
	return _BasAsset.Contract.Revoke(&_BasAsset.TransactOpts, from, nameHash)
}

// SetAlias is a paid mutator transaction binding the contract method 0x468c024c.
//
// Solidity: function setAlias(bytes32 nameHash, string aName) returns()
func (_BasAsset *BasAssetTransactor) SetAlias(opts *bind.TransactOpts, nameHash [32]byte, aName string) (*types.Transaction, error) {
	return _BasAsset.contract.Transact(opts, "setAlias", nameHash, aName)
}

// SetAlias is a paid mutator transaction binding the contract method 0x468c024c.
//
// Solidity: function setAlias(bytes32 nameHash, string aName) returns()
func (_BasAsset *BasAssetSession) SetAlias(nameHash [32]byte, aName string) (*types.Transaction, error) {
	return _BasAsset.Contract.SetAlias(&_BasAsset.TransactOpts, nameHash, aName)
}

// SetAlias is a paid mutator transaction binding the contract method 0x468c024c.
//
// Solidity: function setAlias(bytes32 nameHash, string aName) returns()
func (_BasAsset *BasAssetTransactorSession) SetAlias(nameHash [32]byte, aName string) (*types.Transaction, error) {
	return _BasAsset.Contract.SetAlias(&_BasAsset.TransactOpts, nameHash, aName)
}

// SetBCAddress is a paid mutator transaction binding the contract method 0x706ef095.
//
// Solidity: function setBCAddress(bytes32 nameHash, string bcAddress) returns()
func (_BasAsset *BasAssetTransactor) SetBCAddress(opts *bind.TransactOpts, nameHash [32]byte, bcAddress string) (*types.Transaction, error) {
	return _BasAsset.contract.Transact(opts, "setBCAddress", nameHash, bcAddress)
}

// SetBCAddress is a paid mutator transaction binding the contract method 0x706ef095.
//
// Solidity: function setBCAddress(bytes32 nameHash, string bcAddress) returns()
func (_BasAsset *BasAssetSession) SetBCAddress(nameHash [32]byte, bcAddress string) (*types.Transaction, error) {
	return _BasAsset.Contract.SetBCAddress(&_BasAsset.TransactOpts, nameHash, bcAddress)
}

// SetBCAddress is a paid mutator transaction binding the contract method 0x706ef095.
//
// Solidity: function setBCAddress(bytes32 nameHash, string bcAddress) returns()
func (_BasAsset *BasAssetTransactorSession) SetBCAddress(nameHash [32]byte, bcAddress string) (*types.Transaction, error) {
	return _BasAsset.Contract.SetBCAddress(&_BasAsset.TransactOpts, nameHash, bcAddress)
}

// SetIP is a paid mutator transaction binding the contract method 0x4fd2d18a.
//
// Solidity: function setIP(bytes32 nameHash, bytes4 ipv4, bytes16 ipv6) returns()
func (_BasAsset *BasAssetTransactor) SetIP(opts *bind.TransactOpts, nameHash [32]byte, ipv4 [4]byte, ipv6 [16]byte) (*types.Transaction, error) {
	return _BasAsset.contract.Transact(opts, "setIP", nameHash, ipv4, ipv6)
}

// SetIP is a paid mutator transaction binding the contract method 0x4fd2d18a.
//
// Solidity: function setIP(bytes32 nameHash, bytes4 ipv4, bytes16 ipv6) returns()
func (_BasAsset *BasAssetSession) SetIP(nameHash [32]byte, ipv4 [4]byte, ipv6 [16]byte) (*types.Transaction, error) {
	return _BasAsset.Contract.SetIP(&_BasAsset.TransactOpts, nameHash, ipv4, ipv6)
}

// SetIP is a paid mutator transaction binding the contract method 0x4fd2d18a.
//
// Solidity: function setIP(bytes32 nameHash, bytes4 ipv4, bytes16 ipv6) returns()
func (_BasAsset *BasAssetTransactorSession) SetIP(nameHash [32]byte, ipv4 [4]byte, ipv6 [16]byte) (*types.Transaction, error) {
	return _BasAsset.Contract.SetIP(&_BasAsset.TransactOpts, nameHash, ipv4, ipv6)
}

// SetOANN is a paid mutator transaction binding the contract method 0xbc63d8b9.
//
// Solidity: function setOANN(address newOANN) returns()
func (_BasAsset *BasAssetTransactor) SetOANN(opts *bind.TransactOpts, newOANN common.Address) (*types.Transaction, error) {
	return _BasAsset.contract.Transact(opts, "setOANN", newOANN)
}

// SetOANN is a paid mutator transaction binding the contract method 0xbc63d8b9.
//
// Solidity: function setOANN(address newOANN) returns()
func (_BasAsset *BasAssetSession) SetOANN(newOANN common.Address) (*types.Transaction, error) {
	return _BasAsset.Contract.SetOANN(&_BasAsset.TransactOpts, newOANN)
}

// SetOANN is a paid mutator transaction binding the contract method 0xbc63d8b9.
//
// Solidity: function setOANN(address newOANN) returns()
func (_BasAsset *BasAssetTransactorSession) SetOANN(newOANN common.Address) (*types.Transaction, error) {
	return _BasAsset.Contract.SetOANN(&_BasAsset.TransactOpts, newOANN)
}

// SetOpData is a paid mutator transaction binding the contract method 0x9e19a994.
//
// Solidity: function setOpData(bytes32 nameHash, bytes opData) returns()
func (_BasAsset *BasAssetTransactor) SetOpData(opts *bind.TransactOpts, nameHash [32]byte, opData []byte) (*types.Transaction, error) {
	return _BasAsset.contract.Transact(opts, "setOpData", nameHash, opData)
}

// SetOpData is a paid mutator transaction binding the contract method 0x9e19a994.
//
// Solidity: function setOpData(bytes32 nameHash, bytes opData) returns()
func (_BasAsset *BasAssetSession) SetOpData(nameHash [32]byte, opData []byte) (*types.Transaction, error) {
	return _BasAsset.Contract.SetOpData(&_BasAsset.TransactOpts, nameHash, opData)
}

// SetOpData is a paid mutator transaction binding the contract method 0x9e19a994.
//
// Solidity: function setOpData(bytes32 nameHash, bytes opData) returns()
func (_BasAsset *BasAssetTransactorSession) SetOpData(nameHash [32]byte, opData []byte) (*types.Transaction, error) {
	return _BasAsset.Contract.SetOpData(&_BasAsset.TransactOpts, nameHash, opData)
}

// SetRecord is a paid mutator transaction binding the contract method 0x9385b5e1.
//
// Solidity: function setRecord(bytes32 nameHash, bytes4 ipv4, bytes16 ipv6, string bcAddress, bytes opData, string aName) returns()
func (_BasAsset *BasAssetTransactor) SetRecord(opts *bind.TransactOpts, nameHash [32]byte, ipv4 [4]byte, ipv6 [16]byte, bcAddress string, opData []byte, aName string) (*types.Transaction, error) {
	return _BasAsset.contract.Transact(opts, "setRecord", nameHash, ipv4, ipv6, bcAddress, opData, aName)
}

// SetRecord is a paid mutator transaction binding the contract method 0x9385b5e1.
//
// Solidity: function setRecord(bytes32 nameHash, bytes4 ipv4, bytes16 ipv6, string bcAddress, bytes opData, string aName) returns()
func (_BasAsset *BasAssetSession) SetRecord(nameHash [32]byte, ipv4 [4]byte, ipv6 [16]byte, bcAddress string, opData []byte, aName string) (*types.Transaction, error) {
	return _BasAsset.Contract.SetRecord(&_BasAsset.TransactOpts, nameHash, ipv4, ipv6, bcAddress, opData, aName)
}

// SetRecord is a paid mutator transaction binding the contract method 0x9385b5e1.
//
// Solidity: function setRecord(bytes32 nameHash, bytes4 ipv4, bytes16 ipv6, string bcAddress, bytes opData, string aName) returns()
func (_BasAsset *BasAssetTransactorSession) SetRecord(nameHash [32]byte, ipv4 [4]byte, ipv6 [16]byte, bcAddress string, opData []byte, aName string) (*types.Transaction, error) {
	return _BasAsset.Contract.SetRecord(&_BasAsset.TransactOpts, nameHash, ipv4, ipv6, bcAddress, opData, aName)
}

// TakeoverRoot is a paid mutator transaction binding the contract method 0x1bdf10ff.
//
// Solidity: function takeoverRoot(bytes32 nameHash, uint256 expire, bool isOpen, bool isCustomed, uint256 cusPrice, address oldOwner, address newOwner) returns()
func (_BasAsset *BasAssetTransactor) TakeoverRoot(opts *bind.TransactOpts, nameHash [32]byte, expire *big.Int, isOpen bool, isCustomed bool, cusPrice *big.Int, oldOwner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _BasAsset.contract.Transact(opts, "takeoverRoot", nameHash, expire, isOpen, isCustomed, cusPrice, oldOwner, newOwner)
}

// TakeoverRoot is a paid mutator transaction binding the contract method 0x1bdf10ff.
//
// Solidity: function takeoverRoot(bytes32 nameHash, uint256 expire, bool isOpen, bool isCustomed, uint256 cusPrice, address oldOwner, address newOwner) returns()
func (_BasAsset *BasAssetSession) TakeoverRoot(nameHash [32]byte, expire *big.Int, isOpen bool, isCustomed bool, cusPrice *big.Int, oldOwner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _BasAsset.Contract.TakeoverRoot(&_BasAsset.TransactOpts, nameHash, expire, isOpen, isCustomed, cusPrice, oldOwner, newOwner)
}

// TakeoverRoot is a paid mutator transaction binding the contract method 0x1bdf10ff.
//
// Solidity: function takeoverRoot(bytes32 nameHash, uint256 expire, bool isOpen, bool isCustomed, uint256 cusPrice, address oldOwner, address newOwner) returns()
func (_BasAsset *BasAssetTransactorSession) TakeoverRoot(nameHash [32]byte, expire *big.Int, isOpen bool, isCustomed bool, cusPrice *big.Int, oldOwner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _BasAsset.Contract.TakeoverRoot(&_BasAsset.TransactOpts, nameHash, expire, isOpen, isCustomed, cusPrice, oldOwner, newOwner)
}

// TakeoverSubName is a paid mutator transaction binding the contract method 0xd59a0e68.
//
// Solidity: function takeoverSubName(bytes32 sHash, uint256 expire, address oldOwner, address newOwner) returns()
func (_BasAsset *BasAssetTransactor) TakeoverSubName(opts *bind.TransactOpts, sHash [32]byte, expire *big.Int, oldOwner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _BasAsset.contract.Transact(opts, "takeoverSubName", sHash, expire, oldOwner, newOwner)
}

// TakeoverSubName is a paid mutator transaction binding the contract method 0xd59a0e68.
//
// Solidity: function takeoverSubName(bytes32 sHash, uint256 expire, address oldOwner, address newOwner) returns()
func (_BasAsset *BasAssetSession) TakeoverSubName(sHash [32]byte, expire *big.Int, oldOwner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _BasAsset.Contract.TakeoverSubName(&_BasAsset.TransactOpts, sHash, expire, oldOwner, newOwner)
}

// TakeoverSubName is a paid mutator transaction binding the contract method 0xd59a0e68.
//
// Solidity: function takeoverSubName(bytes32 sHash, uint256 expire, address oldOwner, address newOwner) returns()
func (_BasAsset *BasAssetTransactorSession) TakeoverSubName(sHash [32]byte, expire *big.Int, oldOwner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _BasAsset.Contract.TakeoverSubName(&_BasAsset.TransactOpts, sHash, expire, oldOwner, newOwner)
}

// Transfer is a paid mutator transaction binding the contract method 0x7d32e7bd.
//
// Solidity: function transfer(address to, bytes32 nameHash) returns()
func (_BasAsset *BasAssetTransactor) Transfer(opts *bind.TransactOpts, to common.Address, nameHash [32]byte) (*types.Transaction, error) {
	return _BasAsset.contract.Transact(opts, "transfer", to, nameHash)
}

// Transfer is a paid mutator transaction binding the contract method 0x7d32e7bd.
//
// Solidity: function transfer(address to, bytes32 nameHash) returns()
func (_BasAsset *BasAssetSession) Transfer(to common.Address, nameHash [32]byte) (*types.Transaction, error) {
	return _BasAsset.Contract.Transfer(&_BasAsset.TransactOpts, to, nameHash)
}

// Transfer is a paid mutator transaction binding the contract method 0x7d32e7bd.
//
// Solidity: function transfer(address to, bytes32 nameHash) returns()
func (_BasAsset *BasAssetTransactorSession) Transfer(to common.Address, nameHash [32]byte) (*types.Transaction, error) {
	return _BasAsset.Contract.Transfer(&_BasAsset.TransactOpts, to, nameHash)
}

// TransferContractOwnership is a paid mutator transaction binding the contract method 0xa843c51f.
//
// Solidity: function transferContractOwnership(address newOwner) returns()
func (_BasAsset *BasAssetTransactor) TransferContractOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BasAsset.contract.Transact(opts, "transferContractOwnership", newOwner)
}

// TransferContractOwnership is a paid mutator transaction binding the contract method 0xa843c51f.
//
// Solidity: function transferContractOwnership(address newOwner) returns()
func (_BasAsset *BasAssetSession) TransferContractOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BasAsset.Contract.TransferContractOwnership(&_BasAsset.TransactOpts, newOwner)
}

// TransferContractOwnership is a paid mutator transaction binding the contract method 0xa843c51f.
//
// Solidity: function transferContractOwnership(address newOwner) returns()
func (_BasAsset *BasAssetTransactorSession) TransferContractOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BasAsset.Contract.TransferContractOwnership(&_BasAsset.TransactOpts, newOwner)
}

// TransferFrom is a paid mutator transaction binding the contract method 0xb3c06f50.
//
// Solidity: function transferFrom(address from, address to, bytes32 nameHash) returns()
func (_BasAsset *BasAssetTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, nameHash [32]byte) (*types.Transaction, error) {
	return _BasAsset.contract.Transact(opts, "transferFrom", from, to, nameHash)
}

// TransferFrom is a paid mutator transaction binding the contract method 0xb3c06f50.
//
// Solidity: function transferFrom(address from, address to, bytes32 nameHash) returns()
func (_BasAsset *BasAssetSession) TransferFrom(from common.Address, to common.Address, nameHash [32]byte) (*types.Transaction, error) {
	return _BasAsset.Contract.TransferFrom(&_BasAsset.TransactOpts, from, to, nameHash)
}

// TransferFrom is a paid mutator transaction binding the contract method 0xb3c06f50.
//
// Solidity: function transferFrom(address from, address to, bytes32 nameHash) returns()
func (_BasAsset *BasAssetTransactorSession) TransferFrom(from common.Address, to common.Address, nameHash [32]byte) (*types.Transaction, error) {
	return _BasAsset.Contract.TransferFrom(&_BasAsset.TransactOpts, from, to, nameHash)
}

// BasAssetAssertTransferIterator is returned from FilterAssertTransfer and is used to iterate over the raw logs and unpacked data for AssertTransfer events raised by the BasAsset contract.
type BasAssetAssertTransferIterator struct {
	Event *BasAssetAssertTransfer // Event containing the contract specifics and raw log

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
func (it *BasAssetAssertTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasAssetAssertTransfer)
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
		it.Event = new(BasAssetAssertTransfer)
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
func (it *BasAssetAssertTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasAssetAssertTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasAssetAssertTransfer represents a AssertTransfer event raised by the BasAsset contract.
type BasAssetAssertTransfer struct {
	NameHash [32]byte
	From     common.Address
	To       common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAssertTransfer is a free log retrieval operation binding the contract event 0xcd4efa8d78d4de87992137b0ae50f3f3db87dc11a7ffeabc0994bf58b0e65a6a.
//
// Solidity: event AssertTransfer(bytes32 nameHash, address from, address to)
func (_BasAsset *BasAssetFilterer) FilterAssertTransfer(opts *bind.FilterOpts) (*BasAssetAssertTransferIterator, error) {

	logs, sub, err := _BasAsset.contract.FilterLogs(opts, "AssertTransfer")
	if err != nil {
		return nil, err
	}
	return &BasAssetAssertTransferIterator{contract: _BasAsset.contract, event: "AssertTransfer", logs: logs, sub: sub}, nil
}

// WatchAssertTransfer is a free log subscription operation binding the contract event 0xcd4efa8d78d4de87992137b0ae50f3f3db87dc11a7ffeabc0994bf58b0e65a6a.
//
// Solidity: event AssertTransfer(bytes32 nameHash, address from, address to)
func (_BasAsset *BasAssetFilterer) WatchAssertTransfer(opts *bind.WatchOpts, sink chan<- *BasAssetAssertTransfer) (event.Subscription, error) {

	logs, sub, err := _BasAsset.contract.WatchLogs(opts, "AssertTransfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasAssetAssertTransfer)
				if err := _BasAsset.contract.UnpackLog(event, "AssertTransfer", log); err != nil {
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

// ParseAssertTransfer is a log parse operation binding the contract event 0xcd4efa8d78d4de87992137b0ae50f3f3db87dc11a7ffeabc0994bf58b0e65a6a.
//
// Solidity: event AssertTransfer(bytes32 nameHash, address from, address to)
func (_BasAsset *BasAssetFilterer) ParseAssertTransfer(log types.Log) (*BasAssetAssertTransfer, error) {
	event := new(BasAssetAssertTransfer)
	if err := _BasAsset.contract.UnpackLog(event, "AssertTransfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BasAssetDNSRecordChangeIterator is returned from FilterDNSRecordChange and is used to iterate over the raw logs and unpacked data for DNSRecordChange events raised by the BasAsset contract.
type BasAssetDNSRecordChangeIterator struct {
	Event *BasAssetDNSRecordChange // Event containing the contract specifics and raw log

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
func (it *BasAssetDNSRecordChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasAssetDNSRecordChange)
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
		it.Event = new(BasAssetDNSRecordChange)
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
func (it *BasAssetDNSRecordChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasAssetDNSRecordChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasAssetDNSRecordChange represents a DNSRecordChange event raised by the BasAsset contract.
type BasAssetDNSRecordChange struct {
	NameHash [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDNSRecordChange is a free log retrieval operation binding the contract event 0xc4aaff00d9c0e5be36460c841ba5920da4c842307539cdbe500387afa617e621.
//
// Solidity: event DNSRecordChange(bytes32 nameHash)
func (_BasAsset *BasAssetFilterer) FilterDNSRecordChange(opts *bind.FilterOpts) (*BasAssetDNSRecordChangeIterator, error) {

	logs, sub, err := _BasAsset.contract.FilterLogs(opts, "DNSRecordChange")
	if err != nil {
		return nil, err
	}
	return &BasAssetDNSRecordChangeIterator{contract: _BasAsset.contract, event: "DNSRecordChange", logs: logs, sub: sub}, nil
}

// WatchDNSRecordChange is a free log subscription operation binding the contract event 0xc4aaff00d9c0e5be36460c841ba5920da4c842307539cdbe500387afa617e621.
//
// Solidity: event DNSRecordChange(bytes32 nameHash)
func (_BasAsset *BasAssetFilterer) WatchDNSRecordChange(opts *bind.WatchOpts, sink chan<- *BasAssetDNSRecordChange) (event.Subscription, error) {

	logs, sub, err := _BasAsset.contract.WatchLogs(opts, "DNSRecordChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasAssetDNSRecordChange)
				if err := _BasAsset.contract.UnpackLog(event, "DNSRecordChange", log); err != nil {
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

// ParseDNSRecordChange is a log parse operation binding the contract event 0xc4aaff00d9c0e5be36460c841ba5920da4c842307539cdbe500387afa617e621.
//
// Solidity: event DNSRecordChange(bytes32 nameHash)
func (_BasAsset *BasAssetFilterer) ParseDNSRecordChange(log types.Log) (*BasAssetDNSRecordChange, error) {
	event := new(BasAssetDNSRecordChange)
	if err := _BasAsset.contract.UnpackLog(event, "DNSRecordChange", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BasAssetDNSRecordRemoveIterator is returned from FilterDNSRecordRemove and is used to iterate over the raw logs and unpacked data for DNSRecordRemove events raised by the BasAsset contract.
type BasAssetDNSRecordRemoveIterator struct {
	Event *BasAssetDNSRecordRemove // Event containing the contract specifics and raw log

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
func (it *BasAssetDNSRecordRemoveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasAssetDNSRecordRemove)
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
		it.Event = new(BasAssetDNSRecordRemove)
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
func (it *BasAssetDNSRecordRemoveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasAssetDNSRecordRemoveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasAssetDNSRecordRemove represents a DNSRecordRemove event raised by the BasAsset contract.
type BasAssetDNSRecordRemove struct {
	NameHash [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDNSRecordRemove is a free log retrieval operation binding the contract event 0xcf86f295942fcb9fa5bae52547a73e10a7a197b753af74ec610b8a9283a21439.
//
// Solidity: event DNSRecordRemove(bytes32 nameHash)
func (_BasAsset *BasAssetFilterer) FilterDNSRecordRemove(opts *bind.FilterOpts) (*BasAssetDNSRecordRemoveIterator, error) {

	logs, sub, err := _BasAsset.contract.FilterLogs(opts, "DNSRecordRemove")
	if err != nil {
		return nil, err
	}
	return &BasAssetDNSRecordRemoveIterator{contract: _BasAsset.contract, event: "DNSRecordRemove", logs: logs, sub: sub}, nil
}

// WatchDNSRecordRemove is a free log subscription operation binding the contract event 0xcf86f295942fcb9fa5bae52547a73e10a7a197b753af74ec610b8a9283a21439.
//
// Solidity: event DNSRecordRemove(bytes32 nameHash)
func (_BasAsset *BasAssetFilterer) WatchDNSRecordRemove(opts *bind.WatchOpts, sink chan<- *BasAssetDNSRecordRemove) (event.Subscription, error) {

	logs, sub, err := _BasAsset.contract.WatchLogs(opts, "DNSRecordRemove")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasAssetDNSRecordRemove)
				if err := _BasAsset.contract.UnpackLog(event, "DNSRecordRemove", log); err != nil {
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

// ParseDNSRecordRemove is a log parse operation binding the contract event 0xcf86f295942fcb9fa5bae52547a73e10a7a197b753af74ec610b8a9283a21439.
//
// Solidity: event DNSRecordRemove(bytes32 nameHash)
func (_BasAsset *BasAssetFilterer) ParseDNSRecordRemove(log types.Log) (*BasAssetDNSRecordRemove, error) {
	event := new(BasAssetDNSRecordRemove)
	if err := _BasAsset.contract.UnpackLog(event, "DNSRecordRemove", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BasAssetMintAssetIterator is returned from FilterMintAsset and is used to iterate over the raw logs and unpacked data for MintAsset events raised by the BasAsset contract.
type BasAssetMintAssetIterator struct {
	Event *BasAssetMintAsset // Event containing the contract specifics and raw log

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
func (it *BasAssetMintAssetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasAssetMintAsset)
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
		it.Event = new(BasAssetMintAsset)
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
func (it *BasAssetMintAssetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasAssetMintAssetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasAssetMintAsset represents a MintAsset event raised by the BasAsset contract.
type BasAssetMintAsset struct {
	Owner common.Address
	Hash  [32]byte
	Name  []byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMintAsset is a free log retrieval operation binding the contract event 0x52d8e8dc2982b49519977efd1fdb015121fc21583eafae05e568c207b3109059.
//
// Solidity: event MintAsset(address owner, bytes32 hash, bytes name)
func (_BasAsset *BasAssetFilterer) FilterMintAsset(opts *bind.FilterOpts) (*BasAssetMintAssetIterator, error) {

	logs, sub, err := _BasAsset.contract.FilterLogs(opts, "MintAsset")
	if err != nil {
		return nil, err
	}
	return &BasAssetMintAssetIterator{contract: _BasAsset.contract, event: "MintAsset", logs: logs, sub: sub}, nil
}

// WatchMintAsset is a free log subscription operation binding the contract event 0x52d8e8dc2982b49519977efd1fdb015121fc21583eafae05e568c207b3109059.
//
// Solidity: event MintAsset(address owner, bytes32 hash, bytes name)
func (_BasAsset *BasAssetFilterer) WatchMintAsset(opts *bind.WatchOpts, sink chan<- *BasAssetMintAsset) (event.Subscription, error) {

	logs, sub, err := _BasAsset.contract.WatchLogs(opts, "MintAsset")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasAssetMintAsset)
				if err := _BasAsset.contract.UnpackLog(event, "MintAsset", log); err != nil {
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

// ParseMintAsset is a log parse operation binding the contract event 0x52d8e8dc2982b49519977efd1fdb015121fc21583eafae05e568c207b3109059.
//
// Solidity: event MintAsset(address owner, bytes32 hash, bytes name)
func (_BasAsset *BasAssetFilterer) ParseMintAsset(log types.Log) (*BasAssetMintAsset, error) {
	event := new(BasAssetMintAsset)
	if err := _BasAsset.contract.UnpackLog(event, "MintAsset", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BasAssetRechargeAssetIterator is returned from FilterRechargeAsset and is used to iterate over the raw logs and unpacked data for RechargeAsset events raised by the BasAsset contract.
type BasAssetRechargeAssetIterator struct {
	Event *BasAssetRechargeAsset // Event containing the contract specifics and raw log

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
func (it *BasAssetRechargeAssetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasAssetRechargeAsset)
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
		it.Event = new(BasAssetRechargeAsset)
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
func (it *BasAssetRechargeAssetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasAssetRechargeAssetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasAssetRechargeAsset represents a RechargeAsset event raised by the BasAsset contract.
type BasAssetRechargeAsset struct {
	Hash     [32]byte
	Duration *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRechargeAsset is a free log retrieval operation binding the contract event 0x0429f51dcf78228d5341fb0e92fdbd3330a07a9aede215114103ac1ae0d8395a.
//
// Solidity: event RechargeAsset(bytes32 hash, uint256 duration)
func (_BasAsset *BasAssetFilterer) FilterRechargeAsset(opts *bind.FilterOpts) (*BasAssetRechargeAssetIterator, error) {

	logs, sub, err := _BasAsset.contract.FilterLogs(opts, "RechargeAsset")
	if err != nil {
		return nil, err
	}
	return &BasAssetRechargeAssetIterator{contract: _BasAsset.contract, event: "RechargeAsset", logs: logs, sub: sub}, nil
}

// WatchRechargeAsset is a free log subscription operation binding the contract event 0x0429f51dcf78228d5341fb0e92fdbd3330a07a9aede215114103ac1ae0d8395a.
//
// Solidity: event RechargeAsset(bytes32 hash, uint256 duration)
func (_BasAsset *BasAssetFilterer) WatchRechargeAsset(opts *bind.WatchOpts, sink chan<- *BasAssetRechargeAsset) (event.Subscription, error) {

	logs, sub, err := _BasAsset.contract.WatchLogs(opts, "RechargeAsset")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasAssetRechargeAsset)
				if err := _BasAsset.contract.UnpackLog(event, "RechargeAsset", log); err != nil {
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

// ParseRechargeAsset is a log parse operation binding the contract event 0x0429f51dcf78228d5341fb0e92fdbd3330a07a9aede215114103ac1ae0d8395a.
//
// Solidity: event RechargeAsset(bytes32 hash, uint256 duration)
func (_BasAsset *BasAssetFilterer) ParseRechargeAsset(log types.Log) (*BasAssetRechargeAsset, error) {
	event := new(BasAssetRechargeAsset)
	if err := _BasAsset.contract.UnpackLog(event, "RechargeAsset", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BasAssetRootAddedIterator is returned from FilterRootAdded and is used to iterate over the raw logs and unpacked data for RootAdded events raised by the BasAsset contract.
type BasAssetRootAddedIterator struct {
	Event *BasAssetRootAdded // Event containing the contract specifics and raw log

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
func (it *BasAssetRootAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasAssetRootAdded)
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
		it.Event = new(BasAssetRootAdded)
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
func (it *BasAssetRootAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasAssetRootAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasAssetRootAdded represents a RootAdded event raised by the BasAsset contract.
type BasAssetRootAdded struct {
	NameHash [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRootAdded is a free log retrieval operation binding the contract event 0x388a5f4af784e29fc791e14e90d372b6057e8379d9d5556e843545420c675dea.
//
// Solidity: event RootAdded(bytes32 nameHash)
func (_BasAsset *BasAssetFilterer) FilterRootAdded(opts *bind.FilterOpts) (*BasAssetRootAddedIterator, error) {

	logs, sub, err := _BasAsset.contract.FilterLogs(opts, "RootAdded")
	if err != nil {
		return nil, err
	}
	return &BasAssetRootAddedIterator{contract: _BasAsset.contract, event: "RootAdded", logs: logs, sub: sub}, nil
}

// WatchRootAdded is a free log subscription operation binding the contract event 0x388a5f4af784e29fc791e14e90d372b6057e8379d9d5556e843545420c675dea.
//
// Solidity: event RootAdded(bytes32 nameHash)
func (_BasAsset *BasAssetFilterer) WatchRootAdded(opts *bind.WatchOpts, sink chan<- *BasAssetRootAdded) (event.Subscription, error) {

	logs, sub, err := _BasAsset.contract.WatchLogs(opts, "RootAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasAssetRootAdded)
				if err := _BasAsset.contract.UnpackLog(event, "RootAdded", log); err != nil {
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

// ParseRootAdded is a log parse operation binding the contract event 0x388a5f4af784e29fc791e14e90d372b6057e8379d9d5556e843545420c675dea.
//
// Solidity: event RootAdded(bytes32 nameHash)
func (_BasAsset *BasAssetFilterer) ParseRootAdded(log types.Log) (*BasAssetRootAdded, error) {
	event := new(BasAssetRootAdded)
	if err := _BasAsset.contract.UnpackLog(event, "RootAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BasAssetRootChangedIterator is returned from FilterRootChanged and is used to iterate over the raw logs and unpacked data for RootChanged events raised by the BasAsset contract.
type BasAssetRootChangedIterator struct {
	Event *BasAssetRootChanged // Event containing the contract specifics and raw log

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
func (it *BasAssetRootChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasAssetRootChanged)
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
		it.Event = new(BasAssetRootChanged)
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
func (it *BasAssetRootChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasAssetRootChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasAssetRootChanged represents a RootChanged event raised by the BasAsset contract.
type BasAssetRootChanged struct {
	NameHash [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRootChanged is a free log retrieval operation binding the contract event 0x545a99af2f74d472d3ceb11889ff68b31a1a02f48a9431f04cda814892ee57e2.
//
// Solidity: event RootChanged(bytes32 nameHash)
func (_BasAsset *BasAssetFilterer) FilterRootChanged(opts *bind.FilterOpts) (*BasAssetRootChangedIterator, error) {

	logs, sub, err := _BasAsset.contract.FilterLogs(opts, "RootChanged")
	if err != nil {
		return nil, err
	}
	return &BasAssetRootChangedIterator{contract: _BasAsset.contract, event: "RootChanged", logs: logs, sub: sub}, nil
}

// WatchRootChanged is a free log subscription operation binding the contract event 0x545a99af2f74d472d3ceb11889ff68b31a1a02f48a9431f04cda814892ee57e2.
//
// Solidity: event RootChanged(bytes32 nameHash)
func (_BasAsset *BasAssetFilterer) WatchRootChanged(opts *bind.WatchOpts, sink chan<- *BasAssetRootChanged) (event.Subscription, error) {

	logs, sub, err := _BasAsset.contract.WatchLogs(opts, "RootChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasAssetRootChanged)
				if err := _BasAsset.contract.UnpackLog(event, "RootChanged", log); err != nil {
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

// ParseRootChanged is a log parse operation binding the contract event 0x545a99af2f74d472d3ceb11889ff68b31a1a02f48a9431f04cda814892ee57e2.
//
// Solidity: event RootChanged(bytes32 nameHash)
func (_BasAsset *BasAssetFilterer) ParseRootChanged(log types.Log) (*BasAssetRootChanged, error) {
	event := new(BasAssetRootChanged)
	if err := _BasAsset.contract.UnpackLog(event, "RootChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BasAssetSubAddedIterator is returned from FilterSubAdded and is used to iterate over the raw logs and unpacked data for SubAdded events raised by the BasAsset contract.
type BasAssetSubAddedIterator struct {
	Event *BasAssetSubAdded // Event containing the contract specifics and raw log

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
func (it *BasAssetSubAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasAssetSubAdded)
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
		it.Event = new(BasAssetSubAdded)
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
func (it *BasAssetSubAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasAssetSubAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasAssetSubAdded represents a SubAdded event raised by the BasAsset contract.
type BasAssetSubAdded struct {
	NameHash [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSubAdded is a free log retrieval operation binding the contract event 0x89553af7cea334c9dc75d0da93f13782badf6eb8ff1c31cf1792297cf3a8a914.
//
// Solidity: event SubAdded(bytes32 nameHash)
func (_BasAsset *BasAssetFilterer) FilterSubAdded(opts *bind.FilterOpts) (*BasAssetSubAddedIterator, error) {

	logs, sub, err := _BasAsset.contract.FilterLogs(opts, "SubAdded")
	if err != nil {
		return nil, err
	}
	return &BasAssetSubAddedIterator{contract: _BasAsset.contract, event: "SubAdded", logs: logs, sub: sub}, nil
}

// WatchSubAdded is a free log subscription operation binding the contract event 0x89553af7cea334c9dc75d0da93f13782badf6eb8ff1c31cf1792297cf3a8a914.
//
// Solidity: event SubAdded(bytes32 nameHash)
func (_BasAsset *BasAssetFilterer) WatchSubAdded(opts *bind.WatchOpts, sink chan<- *BasAssetSubAdded) (event.Subscription, error) {

	logs, sub, err := _BasAsset.contract.WatchLogs(opts, "SubAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasAssetSubAdded)
				if err := _BasAsset.contract.UnpackLog(event, "SubAdded", log); err != nil {
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

// ParseSubAdded is a log parse operation binding the contract event 0x89553af7cea334c9dc75d0da93f13782badf6eb8ff1c31cf1792297cf3a8a914.
//
// Solidity: event SubAdded(bytes32 nameHash)
func (_BasAsset *BasAssetFilterer) ParseSubAdded(log types.Log) (*BasAssetSubAdded, error) {
	event := new(BasAssetSubAdded)
	if err := _BasAsset.contract.UnpackLog(event, "SubAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BasAssetSubChangedIterator is returned from FilterSubChanged and is used to iterate over the raw logs and unpacked data for SubChanged events raised by the BasAsset contract.
type BasAssetSubChangedIterator struct {
	Event *BasAssetSubChanged // Event containing the contract specifics and raw log

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
func (it *BasAssetSubChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasAssetSubChanged)
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
		it.Event = new(BasAssetSubChanged)
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
func (it *BasAssetSubChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasAssetSubChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasAssetSubChanged represents a SubChanged event raised by the BasAsset contract.
type BasAssetSubChanged struct {
	NameHash [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSubChanged is a free log retrieval operation binding the contract event 0x5bbf7d10adbd21c487c35c286faeeea72d10232dbc9c7e55d05f10f40f5ac3c3.
//
// Solidity: event SubChanged(bytes32 nameHash)
func (_BasAsset *BasAssetFilterer) FilterSubChanged(opts *bind.FilterOpts) (*BasAssetSubChangedIterator, error) {

	logs, sub, err := _BasAsset.contract.FilterLogs(opts, "SubChanged")
	if err != nil {
		return nil, err
	}
	return &BasAssetSubChangedIterator{contract: _BasAsset.contract, event: "SubChanged", logs: logs, sub: sub}, nil
}

// WatchSubChanged is a free log subscription operation binding the contract event 0x5bbf7d10adbd21c487c35c286faeeea72d10232dbc9c7e55d05f10f40f5ac3c3.
//
// Solidity: event SubChanged(bytes32 nameHash)
func (_BasAsset *BasAssetFilterer) WatchSubChanged(opts *bind.WatchOpts, sink chan<- *BasAssetSubChanged) (event.Subscription, error) {

	logs, sub, err := _BasAsset.contract.WatchLogs(opts, "SubChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasAssetSubChanged)
				if err := _BasAsset.contract.UnpackLog(event, "SubChanged", log); err != nil {
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

// ParseSubChanged is a log parse operation binding the contract event 0x5bbf7d10adbd21c487c35c286faeeea72d10232dbc9c7e55d05f10f40f5ac3c3.
//
// Solidity: event SubChanged(bytes32 nameHash)
func (_BasAsset *BasAssetFilterer) ParseSubChanged(log types.Log) (*BasAssetSubChanged, error) {
	event := new(BasAssetSubChanged)
	if err := _BasAsset.contract.UnpackLog(event, "SubChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BasAssetTakeoverAssetIterator is returned from FilterTakeoverAsset and is used to iterate over the raw logs and unpacked data for TakeoverAsset events raised by the BasAsset contract.
type BasAssetTakeoverAssetIterator struct {
	Event *BasAssetTakeoverAsset // Event containing the contract specifics and raw log

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
func (it *BasAssetTakeoverAssetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasAssetTakeoverAsset)
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
		it.Event = new(BasAssetTakeoverAsset)
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
func (it *BasAssetTakeoverAssetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasAssetTakeoverAssetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasAssetTakeoverAsset represents a TakeoverAsset event raised by the BasAsset contract.
type BasAssetTakeoverAsset struct {
	OldOwner common.Address
	NewOwner common.Address
	Hash     [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTakeoverAsset is a free log retrieval operation binding the contract event 0x9224ee6f06391fbcf8b3f9ab9c1f0c339a04d38a72fdaa9b4abec371afbcfd00.
//
// Solidity: event TakeoverAsset(address oldOwner, address newOwner, bytes32 hash)
func (_BasAsset *BasAssetFilterer) FilterTakeoverAsset(opts *bind.FilterOpts) (*BasAssetTakeoverAssetIterator, error) {

	logs, sub, err := _BasAsset.contract.FilterLogs(opts, "TakeoverAsset")
	if err != nil {
		return nil, err
	}
	return &BasAssetTakeoverAssetIterator{contract: _BasAsset.contract, event: "TakeoverAsset", logs: logs, sub: sub}, nil
}

// WatchTakeoverAsset is a free log subscription operation binding the contract event 0x9224ee6f06391fbcf8b3f9ab9c1f0c339a04d38a72fdaa9b4abec371afbcfd00.
//
// Solidity: event TakeoverAsset(address oldOwner, address newOwner, bytes32 hash)
func (_BasAsset *BasAssetFilterer) WatchTakeoverAsset(opts *bind.WatchOpts, sink chan<- *BasAssetTakeoverAsset) (event.Subscription, error) {

	logs, sub, err := _BasAsset.contract.WatchLogs(opts, "TakeoverAsset")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasAssetTakeoverAsset)
				if err := _BasAsset.contract.UnpackLog(event, "TakeoverAsset", log); err != nil {
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

// ParseTakeoverAsset is a log parse operation binding the contract event 0x9224ee6f06391fbcf8b3f9ab9c1f0c339a04d38a72fdaa9b4abec371afbcfd00.
//
// Solidity: event TakeoverAsset(address oldOwner, address newOwner, bytes32 hash)
func (_BasAsset *BasAssetFilterer) ParseTakeoverAsset(log types.Log) (*BasAssetTakeoverAsset, error) {
	event := new(BasAssetTakeoverAsset)
	if err := _BasAsset.contract.UnpackLog(event, "TakeoverAsset", log); err != nil {
		return nil, err
	}
	return event, nil
}
