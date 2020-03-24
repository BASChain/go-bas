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
const BasAssetABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"}],\"name\":\"subExist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"Sub\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"totalName\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"rootHash\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"o\",\"outputs\":[{\"internalType\":\"contractBasOwnership\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"openToPublic\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isCustomed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"customedPrice\",\"type\":\"uint256\"}],\"name\":\"_c_updateRoot\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"totalName\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"rootHash\",\"type\":\"bytes32\"}],\"name\":\"_a_updateSub\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"totalName\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"rootHash\",\"type\":\"bytes32\"}],\"name\":\"_c_addSub\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"rootName\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"openToPublic\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isCustomed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"customedPrice\",\"type\":\"uint256\"}],\"name\":\"_c_addRoot\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"}],\"name\":\"rootExist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"Root\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"rootName\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"openToPublic\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isCustomed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"customedPrice\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"rootName\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"openToPublic\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isCustomed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"customedPrice\",\"type\":\"uint256\"}],\"name\":\"_a_updateRoot\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"}],\"name\":\"Hash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"sName\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"rName\",\"type\":\"bytes\"}],\"name\":\"TotalNameHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_o\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"}],\"name\":\"RootChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"}],\"name\":\"SubChanged\",\"type\":\"event\"}]"

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

// Hash is a free data retrieval call binding the contract method 0xbd1bd35a.
//
// Solidity: function Hash(bytes name) constant returns(bytes32)
func (_BasAsset *BasAssetCaller) Hash(opts *bind.CallOpts, name []byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _BasAsset.contract.Call(opts, out, "Hash", name)
	return *ret0, err
}

// Hash is a free data retrieval call binding the contract method 0xbd1bd35a.
//
// Solidity: function Hash(bytes name) constant returns(bytes32)
func (_BasAsset *BasAssetSession) Hash(name []byte) ([32]byte, error) {
	return _BasAsset.Contract.Hash(&_BasAsset.CallOpts, name)
}

// Hash is a free data retrieval call binding the contract method 0xbd1bd35a.
//
// Solidity: function Hash(bytes name) constant returns(bytes32)
func (_BasAsset *BasAssetCallerSession) Hash(name []byte) ([32]byte, error) {
	return _BasAsset.Contract.Hash(&_BasAsset.CallOpts, name)
}

// Root is a free data retrieval call binding the contract method 0x83b8202f.
//
// Solidity: function Root(bytes32 ) constant returns(bytes rootName, bool openToPublic, bool isCustomed, uint256 customedPrice)
func (_BasAsset *BasAssetCaller) Root(opts *bind.CallOpts, arg0 [32]byte) (struct {
	RootName      []byte
	OpenToPublic  bool
	IsCustomed    bool
	CustomedPrice *big.Int
}, error) {
	ret := new(struct {
		RootName      []byte
		OpenToPublic  bool
		IsCustomed    bool
		CustomedPrice *big.Int
	})
	out := ret
	err := _BasAsset.contract.Call(opts, out, "Root", arg0)
	return *ret, err
}

// Root is a free data retrieval call binding the contract method 0x83b8202f.
//
// Solidity: function Root(bytes32 ) constant returns(bytes rootName, bool openToPublic, bool isCustomed, uint256 customedPrice)
func (_BasAsset *BasAssetSession) Root(arg0 [32]byte) (struct {
	RootName      []byte
	OpenToPublic  bool
	IsCustomed    bool
	CustomedPrice *big.Int
}, error) {
	return _BasAsset.Contract.Root(&_BasAsset.CallOpts, arg0)
}

// Root is a free data retrieval call binding the contract method 0x83b8202f.
//
// Solidity: function Root(bytes32 ) constant returns(bytes rootName, bool openToPublic, bool isCustomed, uint256 customedPrice)
func (_BasAsset *BasAssetCallerSession) Root(arg0 [32]byte) (struct {
	RootName      []byte
	OpenToPublic  bool
	IsCustomed    bool
	CustomedPrice *big.Int
}, error) {
	return _BasAsset.Contract.Root(&_BasAsset.CallOpts, arg0)
}

// Sub is a free data retrieval call binding the contract method 0x333cf44c.
//
// Solidity: function Sub(bytes32 ) constant returns(bytes totalName, bytes32 rootHash)
func (_BasAsset *BasAssetCaller) Sub(opts *bind.CallOpts, arg0 [32]byte) (struct {
	TotalName []byte
	RootHash  [32]byte
}, error) {
	ret := new(struct {
		TotalName []byte
		RootHash  [32]byte
	})
	out := ret
	err := _BasAsset.contract.Call(opts, out, "Sub", arg0)
	return *ret, err
}

// Sub is a free data retrieval call binding the contract method 0x333cf44c.
//
// Solidity: function Sub(bytes32 ) constant returns(bytes totalName, bytes32 rootHash)
func (_BasAsset *BasAssetSession) Sub(arg0 [32]byte) (struct {
	TotalName []byte
	RootHash  [32]byte
}, error) {
	return _BasAsset.Contract.Sub(&_BasAsset.CallOpts, arg0)
}

// Sub is a free data retrieval call binding the contract method 0x333cf44c.
//
// Solidity: function Sub(bytes32 ) constant returns(bytes totalName, bytes32 rootHash)
func (_BasAsset *BasAssetCallerSession) Sub(arg0 [32]byte) (struct {
	TotalName []byte
	RootHash  [32]byte
}, error) {
	return _BasAsset.Contract.Sub(&_BasAsset.CallOpts, arg0)
}

// TotalNameHash is a free data retrieval call binding the contract method 0xc3821972.
//
// Solidity: function TotalNameHash(bytes sName, bytes rName) constant returns(bytes32)
func (_BasAsset *BasAssetCaller) TotalNameHash(opts *bind.CallOpts, sName []byte, rName []byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _BasAsset.contract.Call(opts, out, "TotalNameHash", sName, rName)
	return *ret0, err
}

// TotalNameHash is a free data retrieval call binding the contract method 0xc3821972.
//
// Solidity: function TotalNameHash(bytes sName, bytes rName) constant returns(bytes32)
func (_BasAsset *BasAssetSession) TotalNameHash(sName []byte, rName []byte) ([32]byte, error) {
	return _BasAsset.Contract.TotalNameHash(&_BasAsset.CallOpts, sName, rName)
}

// TotalNameHash is a free data retrieval call binding the contract method 0xc3821972.
//
// Solidity: function TotalNameHash(bytes sName, bytes rName) constant returns(bytes32)
func (_BasAsset *BasAssetCallerSession) TotalNameHash(sName []byte, rName []byte) ([32]byte, error) {
	return _BasAsset.Contract.TotalNameHash(&_BasAsset.CallOpts, sName, rName)
}

// O is a free data retrieval call binding the contract method 0x50cd4df2.
//
// Solidity: function o() constant returns(address)
func (_BasAsset *BasAssetCaller) O(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BasAsset.contract.Call(opts, out, "o")
	return *ret0, err
}

// O is a free data retrieval call binding the contract method 0x50cd4df2.
//
// Solidity: function o() constant returns(address)
func (_BasAsset *BasAssetSession) O() (common.Address, error) {
	return _BasAsset.Contract.O(&_BasAsset.CallOpts)
}

// O is a free data retrieval call binding the contract method 0x50cd4df2.
//
// Solidity: function o() constant returns(address)
func (_BasAsset *BasAssetCallerSession) O() (common.Address, error) {
	return _BasAsset.Contract.O(&_BasAsset.CallOpts)
}

// RootExist is a free data retrieval call binding the contract method 0x83026106.
//
// Solidity: function rootExist(bytes32 nameHash) constant returns(bool)
func (_BasAsset *BasAssetCaller) RootExist(opts *bind.CallOpts, nameHash [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BasAsset.contract.Call(opts, out, "rootExist", nameHash)
	return *ret0, err
}

// RootExist is a free data retrieval call binding the contract method 0x83026106.
//
// Solidity: function rootExist(bytes32 nameHash) constant returns(bool)
func (_BasAsset *BasAssetSession) RootExist(nameHash [32]byte) (bool, error) {
	return _BasAsset.Contract.RootExist(&_BasAsset.CallOpts, nameHash)
}

// RootExist is a free data retrieval call binding the contract method 0x83026106.
//
// Solidity: function rootExist(bytes32 nameHash) constant returns(bool)
func (_BasAsset *BasAssetCallerSession) RootExist(nameHash [32]byte) (bool, error) {
	return _BasAsset.Contract.RootExist(&_BasAsset.CallOpts, nameHash)
}

// SubExist is a free data retrieval call binding the contract method 0x0119a9e6.
//
// Solidity: function subExist(bytes32 nameHash) constant returns(bool)
func (_BasAsset *BasAssetCaller) SubExist(opts *bind.CallOpts, nameHash [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BasAsset.contract.Call(opts, out, "subExist", nameHash)
	return *ret0, err
}

// SubExist is a free data retrieval call binding the contract method 0x0119a9e6.
//
// Solidity: function subExist(bytes32 nameHash) constant returns(bool)
func (_BasAsset *BasAssetSession) SubExist(nameHash [32]byte) (bool, error) {
	return _BasAsset.Contract.SubExist(&_BasAsset.CallOpts, nameHash)
}

// SubExist is a free data retrieval call binding the contract method 0x0119a9e6.
//
// Solidity: function subExist(bytes32 nameHash) constant returns(bool)
func (_BasAsset *BasAssetCallerSession) SubExist(nameHash [32]byte) (bool, error) {
	return _BasAsset.Contract.SubExist(&_BasAsset.CallOpts, nameHash)
}

// AUpdateRoot is a paid mutator transaction binding the contract method 0xaa57635e.
//
// Solidity: function _a_updateRoot(bytes rootName, bool openToPublic, bool isCustomed, uint256 customedPrice) returns()
func (_BasAsset *BasAssetTransactor) AUpdateRoot(opts *bind.TransactOpts, rootName []byte, openToPublic bool, isCustomed bool, customedPrice *big.Int) (*types.Transaction, error) {
	return _BasAsset.contract.Transact(opts, "_a_updateRoot", rootName, openToPublic, isCustomed, customedPrice)
}

// AUpdateRoot is a paid mutator transaction binding the contract method 0xaa57635e.
//
// Solidity: function _a_updateRoot(bytes rootName, bool openToPublic, bool isCustomed, uint256 customedPrice) returns()
func (_BasAsset *BasAssetSession) AUpdateRoot(rootName []byte, openToPublic bool, isCustomed bool, customedPrice *big.Int) (*types.Transaction, error) {
	return _BasAsset.Contract.AUpdateRoot(&_BasAsset.TransactOpts, rootName, openToPublic, isCustomed, customedPrice)
}

// AUpdateRoot is a paid mutator transaction binding the contract method 0xaa57635e.
//
// Solidity: function _a_updateRoot(bytes rootName, bool openToPublic, bool isCustomed, uint256 customedPrice) returns()
func (_BasAsset *BasAssetTransactorSession) AUpdateRoot(rootName []byte, openToPublic bool, isCustomed bool, customedPrice *big.Int) (*types.Transaction, error) {
	return _BasAsset.Contract.AUpdateRoot(&_BasAsset.TransactOpts, rootName, openToPublic, isCustomed, customedPrice)
}

// AUpdateSub is a paid mutator transaction binding the contract method 0x603f29cd.
//
// Solidity: function _a_updateSub(bytes totalName, bytes32 rootHash) returns()
func (_BasAsset *BasAssetTransactor) AUpdateSub(opts *bind.TransactOpts, totalName []byte, rootHash [32]byte) (*types.Transaction, error) {
	return _BasAsset.contract.Transact(opts, "_a_updateSub", totalName, rootHash)
}

// AUpdateSub is a paid mutator transaction binding the contract method 0x603f29cd.
//
// Solidity: function _a_updateSub(bytes totalName, bytes32 rootHash) returns()
func (_BasAsset *BasAssetSession) AUpdateSub(totalName []byte, rootHash [32]byte) (*types.Transaction, error) {
	return _BasAsset.Contract.AUpdateSub(&_BasAsset.TransactOpts, totalName, rootHash)
}

// AUpdateSub is a paid mutator transaction binding the contract method 0x603f29cd.
//
// Solidity: function _a_updateSub(bytes totalName, bytes32 rootHash) returns()
func (_BasAsset *BasAssetTransactorSession) AUpdateSub(totalName []byte, rootHash [32]byte) (*types.Transaction, error) {
	return _BasAsset.Contract.AUpdateSub(&_BasAsset.TransactOpts, totalName, rootHash)
}

// CAddRoot is a paid mutator transaction binding the contract method 0x6f84208d.
//
// Solidity: function _c_addRoot(bytes rootName, bool openToPublic, bool isCustomed, uint256 customedPrice) returns()
func (_BasAsset *BasAssetTransactor) CAddRoot(opts *bind.TransactOpts, rootName []byte, openToPublic bool, isCustomed bool, customedPrice *big.Int) (*types.Transaction, error) {
	return _BasAsset.contract.Transact(opts, "_c_addRoot", rootName, openToPublic, isCustomed, customedPrice)
}

// CAddRoot is a paid mutator transaction binding the contract method 0x6f84208d.
//
// Solidity: function _c_addRoot(bytes rootName, bool openToPublic, bool isCustomed, uint256 customedPrice) returns()
func (_BasAsset *BasAssetSession) CAddRoot(rootName []byte, openToPublic bool, isCustomed bool, customedPrice *big.Int) (*types.Transaction, error) {
	return _BasAsset.Contract.CAddRoot(&_BasAsset.TransactOpts, rootName, openToPublic, isCustomed, customedPrice)
}

// CAddRoot is a paid mutator transaction binding the contract method 0x6f84208d.
//
// Solidity: function _c_addRoot(bytes rootName, bool openToPublic, bool isCustomed, uint256 customedPrice) returns()
func (_BasAsset *BasAssetTransactorSession) CAddRoot(rootName []byte, openToPublic bool, isCustomed bool, customedPrice *big.Int) (*types.Transaction, error) {
	return _BasAsset.Contract.CAddRoot(&_BasAsset.TransactOpts, rootName, openToPublic, isCustomed, customedPrice)
}

// CAddSub is a paid mutator transaction binding the contract method 0x66967cfb.
//
// Solidity: function _c_addSub(bytes totalName, bytes32 rootHash) returns()
func (_BasAsset *BasAssetTransactor) CAddSub(opts *bind.TransactOpts, totalName []byte, rootHash [32]byte) (*types.Transaction, error) {
	return _BasAsset.contract.Transact(opts, "_c_addSub", totalName, rootHash)
}

// CAddSub is a paid mutator transaction binding the contract method 0x66967cfb.
//
// Solidity: function _c_addSub(bytes totalName, bytes32 rootHash) returns()
func (_BasAsset *BasAssetSession) CAddSub(totalName []byte, rootHash [32]byte) (*types.Transaction, error) {
	return _BasAsset.Contract.CAddSub(&_BasAsset.TransactOpts, totalName, rootHash)
}

// CAddSub is a paid mutator transaction binding the contract method 0x66967cfb.
//
// Solidity: function _c_addSub(bytes totalName, bytes32 rootHash) returns()
func (_BasAsset *BasAssetTransactorSession) CAddSub(totalName []byte, rootHash [32]byte) (*types.Transaction, error) {
	return _BasAsset.Contract.CAddSub(&_BasAsset.TransactOpts, totalName, rootHash)
}

// CUpdateRoot is a paid mutator transaction binding the contract method 0x514bfdca.
//
// Solidity: function _c_updateRoot(bytes32 nameHash, bool openToPublic, bool isCustomed, uint256 customedPrice) returns()
func (_BasAsset *BasAssetTransactor) CUpdateRoot(opts *bind.TransactOpts, nameHash [32]byte, openToPublic bool, isCustomed bool, customedPrice *big.Int) (*types.Transaction, error) {
	return _BasAsset.contract.Transact(opts, "_c_updateRoot", nameHash, openToPublic, isCustomed, customedPrice)
}

// CUpdateRoot is a paid mutator transaction binding the contract method 0x514bfdca.
//
// Solidity: function _c_updateRoot(bytes32 nameHash, bool openToPublic, bool isCustomed, uint256 customedPrice) returns()
func (_BasAsset *BasAssetSession) CUpdateRoot(nameHash [32]byte, openToPublic bool, isCustomed bool, customedPrice *big.Int) (*types.Transaction, error) {
	return _BasAsset.Contract.CUpdateRoot(&_BasAsset.TransactOpts, nameHash, openToPublic, isCustomed, customedPrice)
}

// CUpdateRoot is a paid mutator transaction binding the contract method 0x514bfdca.
//
// Solidity: function _c_updateRoot(bytes32 nameHash, bool openToPublic, bool isCustomed, uint256 customedPrice) returns()
func (_BasAsset *BasAssetTransactorSession) CUpdateRoot(nameHash [32]byte, openToPublic bool, isCustomed bool, customedPrice *big.Int) (*types.Transaction, error) {
	return _BasAsset.Contract.CUpdateRoot(&_BasAsset.TransactOpts, nameHash, openToPublic, isCustomed, customedPrice)
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
