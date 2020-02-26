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

// BasOANNABI is the input ABI used to generate the binding from.
const BasOANNABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newA\",\"type\":\"address\"}],\"name\":\"changeAsset\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"d\",\"type\":\"uint8\"}],\"name\":\"validDuration\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"asset\",\"outputs\":[{\"internalType\":\"contractBasAsset\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"rootHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"openCustomedPrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"rName\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sName\",\"type\":\"bytes\"},{\"internalType\":\"uint8\",\"name\":\"durationInYear\",\"type\":\"uint8\"}],\"name\":\"evalueSubPrice\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"rootOwner\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isCustomed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_DAYS_REG\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"accountant\",\"outputs\":[{\"internalType\":\"contractBasMiner\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_YEAR_REG\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CUSTOMED_PRICE_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ContractOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newT\",\"type\":\"address\"}],\"name\":\"changeToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"durationInYear\",\"type\":\"uint8\"}],\"name\":\"rechargeSub\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isOpen\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isCustomed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"cusPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"durationInYear\",\"type\":\"uint8\"}],\"name\":\"registerRoot\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"AROOT_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"durationInYear\",\"type\":\"uint8\"}],\"name\":\"rechargeRoot\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newGas\",\"type\":\"uint256\"}],\"name\":\"setARootGas\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newGas\",\"type\":\"uint256\"}],\"name\":\"setBRootGas\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferContractOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isCustomed\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"durationInYear\",\"type\":\"uint8\"}],\"name\":\"evalueRootPrice\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isARoot\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"rName\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sName\",\"type\":\"bytes\"},{\"internalType\":\"uint8\",\"name\":\"durationInYear\",\"type\":\"uint8\"}],\"name\":\"registerSub\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"SUB_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newM\",\"type\":\"address\"}],\"name\":\"changeAccountant\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractERC20\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BROOT_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"m\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// BasOANN is an auto generated Go binding around an Ethereum contract.
type BasOANN struct {
	BasOANNCaller     // Read-only binding to the contract
	BasOANNTransactor // Write-only binding to the contract
	BasOANNFilterer   // Log filterer for contract events
}

// BasOANNCaller is an auto generated read-only Go binding around an Ethereum contract.
type BasOANNCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasOANNTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BasOANNTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasOANNFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BasOANNFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasOANNSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BasOANNSession struct {
	Contract     *BasOANN          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BasOANNCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BasOANNCallerSession struct {
	Contract *BasOANNCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// BasOANNTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BasOANNTransactorSession struct {
	Contract     *BasOANNTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// BasOANNRaw is an auto generated low-level Go binding around an Ethereum contract.
type BasOANNRaw struct {
	Contract *BasOANN // Generic contract binding to access the raw methods on
}

// BasOANNCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BasOANNCallerRaw struct {
	Contract *BasOANNCaller // Generic read-only contract binding to access the raw methods on
}

// BasOANNTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BasOANNTransactorRaw struct {
	Contract *BasOANNTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBasOANN creates a new instance of BasOANN, bound to a specific deployed contract.
func NewBasOANN(address common.Address, backend bind.ContractBackend) (*BasOANN, error) {
	contract, err := bindBasOANN(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BasOANN{BasOANNCaller: BasOANNCaller{contract: contract}, BasOANNTransactor: BasOANNTransactor{contract: contract}, BasOANNFilterer: BasOANNFilterer{contract: contract}}, nil
}

// NewBasOANNCaller creates a new read-only instance of BasOANN, bound to a specific deployed contract.
func NewBasOANNCaller(address common.Address, caller bind.ContractCaller) (*BasOANNCaller, error) {
	contract, err := bindBasOANN(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BasOANNCaller{contract: contract}, nil
}

// NewBasOANNTransactor creates a new write-only instance of BasOANN, bound to a specific deployed contract.
func NewBasOANNTransactor(address common.Address, transactor bind.ContractTransactor) (*BasOANNTransactor, error) {
	contract, err := bindBasOANN(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BasOANNTransactor{contract: contract}, nil
}

// NewBasOANNFilterer creates a new log filterer instance of BasOANN, bound to a specific deployed contract.
func NewBasOANNFilterer(address common.Address, filterer bind.ContractFilterer) (*BasOANNFilterer, error) {
	contract, err := bindBasOANN(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BasOANNFilterer{contract: contract}, nil
}

// bindBasOANN binds a generic wrapper to an already deployed contract.
func bindBasOANN(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BasOANNABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasOANN *BasOANNRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BasOANN.Contract.BasOANNCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasOANN *BasOANNRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasOANN.Contract.BasOANNTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasOANN *BasOANNRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasOANN.Contract.BasOANNTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasOANN *BasOANNCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BasOANN.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasOANN *BasOANNTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasOANN.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasOANN *BasOANNTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasOANN.Contract.contract.Transact(opts, method, params...)
}

// AROOTGAS is a free data retrieval call binding the contract method 0x789dba4f.
//
// Solidity: function AROOT_GAS() constant returns(uint256)
func (_BasOANN *BasOANNCaller) AROOTGAS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BasOANN.contract.Call(opts, out, "AROOT_GAS")
	return *ret0, err
}

// AROOTGAS is a free data retrieval call binding the contract method 0x789dba4f.
//
// Solidity: function AROOT_GAS() constant returns(uint256)
func (_BasOANN *BasOANNSession) AROOTGAS() (*big.Int, error) {
	return _BasOANN.Contract.AROOTGAS(&_BasOANN.CallOpts)
}

// AROOTGAS is a free data retrieval call binding the contract method 0x789dba4f.
//
// Solidity: function AROOT_GAS() constant returns(uint256)
func (_BasOANN *BasOANNCallerSession) AROOTGAS() (*big.Int, error) {
	return _BasOANN.Contract.AROOTGAS(&_BasOANN.CallOpts)
}

// BROOTGAS is a free data retrieval call binding the contract method 0xfcb26830.
//
// Solidity: function BROOT_GAS() constant returns(uint256)
func (_BasOANN *BasOANNCaller) BROOTGAS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BasOANN.contract.Call(opts, out, "BROOT_GAS")
	return *ret0, err
}

// BROOTGAS is a free data retrieval call binding the contract method 0xfcb26830.
//
// Solidity: function BROOT_GAS() constant returns(uint256)
func (_BasOANN *BasOANNSession) BROOTGAS() (*big.Int, error) {
	return _BasOANN.Contract.BROOTGAS(&_BasOANN.CallOpts)
}

// BROOTGAS is a free data retrieval call binding the contract method 0xfcb26830.
//
// Solidity: function BROOT_GAS() constant returns(uint256)
func (_BasOANN *BasOANNCallerSession) BROOTGAS() (*big.Int, error) {
	return _BasOANN.Contract.BROOTGAS(&_BasOANN.CallOpts)
}

// CUSTOMEDPRICEGAS is a free data retrieval call binding the contract method 0x536acc1c.
//
// Solidity: function CUSTOMED_PRICE_GAS() constant returns(uint256)
func (_BasOANN *BasOANNCaller) CUSTOMEDPRICEGAS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BasOANN.contract.Call(opts, out, "CUSTOMED_PRICE_GAS")
	return *ret0, err
}

// CUSTOMEDPRICEGAS is a free data retrieval call binding the contract method 0x536acc1c.
//
// Solidity: function CUSTOMED_PRICE_GAS() constant returns(uint256)
func (_BasOANN *BasOANNSession) CUSTOMEDPRICEGAS() (*big.Int, error) {
	return _BasOANN.Contract.CUSTOMEDPRICEGAS(&_BasOANN.CallOpts)
}

// CUSTOMEDPRICEGAS is a free data retrieval call binding the contract method 0x536acc1c.
//
// Solidity: function CUSTOMED_PRICE_GAS() constant returns(uint256)
func (_BasOANN *BasOANNCallerSession) CUSTOMEDPRICEGAS() (*big.Int, error) {
	return _BasOANN.Contract.CUSTOMEDPRICEGAS(&_BasOANN.CallOpts)
}

// ContractOwner is a free data retrieval call binding the contract method 0x5a63fbc9.
//
// Solidity: function ContractOwner() constant returns(address)
func (_BasOANN *BasOANNCaller) ContractOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BasOANN.contract.Call(opts, out, "ContractOwner")
	return *ret0, err
}

// ContractOwner is a free data retrieval call binding the contract method 0x5a63fbc9.
//
// Solidity: function ContractOwner() constant returns(address)
func (_BasOANN *BasOANNSession) ContractOwner() (common.Address, error) {
	return _BasOANN.Contract.ContractOwner(&_BasOANN.CallOpts)
}

// ContractOwner is a free data retrieval call binding the contract method 0x5a63fbc9.
//
// Solidity: function ContractOwner() constant returns(address)
func (_BasOANN *BasOANNCallerSession) ContractOwner() (common.Address, error) {
	return _BasOANN.Contract.ContractOwner(&_BasOANN.CallOpts)
}

// MAXDAYSREG is a free data retrieval call binding the contract method 0x46a01a1a.
//
// Solidity: function MAX_DAYS_REG() constant returns(uint256)
func (_BasOANN *BasOANNCaller) MAXDAYSREG(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BasOANN.contract.Call(opts, out, "MAX_DAYS_REG")
	return *ret0, err
}

// MAXDAYSREG is a free data retrieval call binding the contract method 0x46a01a1a.
//
// Solidity: function MAX_DAYS_REG() constant returns(uint256)
func (_BasOANN *BasOANNSession) MAXDAYSREG() (*big.Int, error) {
	return _BasOANN.Contract.MAXDAYSREG(&_BasOANN.CallOpts)
}

// MAXDAYSREG is a free data retrieval call binding the contract method 0x46a01a1a.
//
// Solidity: function MAX_DAYS_REG() constant returns(uint256)
func (_BasOANN *BasOANNCallerSession) MAXDAYSREG() (*big.Int, error) {
	return _BasOANN.Contract.MAXDAYSREG(&_BasOANN.CallOpts)
}

// MAXYEARREG is a free data retrieval call binding the contract method 0x5286fed7.
//
// Solidity: function MAX_YEAR_REG() constant returns(uint256)
func (_BasOANN *BasOANNCaller) MAXYEARREG(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BasOANN.contract.Call(opts, out, "MAX_YEAR_REG")
	return *ret0, err
}

// MAXYEARREG is a free data retrieval call binding the contract method 0x5286fed7.
//
// Solidity: function MAX_YEAR_REG() constant returns(uint256)
func (_BasOANN *BasOANNSession) MAXYEARREG() (*big.Int, error) {
	return _BasOANN.Contract.MAXYEARREG(&_BasOANN.CallOpts)
}

// MAXYEARREG is a free data retrieval call binding the contract method 0x5286fed7.
//
// Solidity: function MAX_YEAR_REG() constant returns(uint256)
func (_BasOANN *BasOANNCallerSession) MAXYEARREG() (*big.Int, error) {
	return _BasOANN.Contract.MAXYEARREG(&_BasOANN.CallOpts)
}

// SUBGAS is a free data retrieval call binding the contract method 0xfa1826bb.
//
// Solidity: function SUB_GAS() constant returns(uint256)
func (_BasOANN *BasOANNCaller) SUBGAS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BasOANN.contract.Call(opts, out, "SUB_GAS")
	return *ret0, err
}

// SUBGAS is a free data retrieval call binding the contract method 0xfa1826bb.
//
// Solidity: function SUB_GAS() constant returns(uint256)
func (_BasOANN *BasOANNSession) SUBGAS() (*big.Int, error) {
	return _BasOANN.Contract.SUBGAS(&_BasOANN.CallOpts)
}

// SUBGAS is a free data retrieval call binding the contract method 0xfa1826bb.
//
// Solidity: function SUB_GAS() constant returns(uint256)
func (_BasOANN *BasOANNCallerSession) SUBGAS() (*big.Int, error) {
	return _BasOANN.Contract.SUBGAS(&_BasOANN.CallOpts)
}

// Accountant is a free data retrieval call binding the contract method 0x4fb3ccc5.
//
// Solidity: function accountant() constant returns(address)
func (_BasOANN *BasOANNCaller) Accountant(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BasOANN.contract.Call(opts, out, "accountant")
	return *ret0, err
}

// Accountant is a free data retrieval call binding the contract method 0x4fb3ccc5.
//
// Solidity: function accountant() constant returns(address)
func (_BasOANN *BasOANNSession) Accountant() (common.Address, error) {
	return _BasOANN.Contract.Accountant(&_BasOANN.CallOpts)
}

// Accountant is a free data retrieval call binding the contract method 0x4fb3ccc5.
//
// Solidity: function accountant() constant returns(address)
func (_BasOANN *BasOANNCallerSession) Accountant() (common.Address, error) {
	return _BasOANN.Contract.Accountant(&_BasOANN.CallOpts)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() constant returns(address)
func (_BasOANN *BasOANNCaller) Asset(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BasOANN.contract.Call(opts, out, "asset")
	return *ret0, err
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() constant returns(address)
func (_BasOANN *BasOANNSession) Asset() (common.Address, error) {
	return _BasOANN.Contract.Asset(&_BasOANN.CallOpts)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() constant returns(address)
func (_BasOANN *BasOANNCallerSession) Asset() (common.Address, error) {
	return _BasOANN.Contract.Asset(&_BasOANN.CallOpts)
}

// EvalueRootPrice is a free data retrieval call binding the contract method 0xe4ede0c5.
//
// Solidity: function evalueRootPrice(bytes name, bool isCustomed, uint8 durationInYear) constant returns(bool isValid, bool isARoot, uint256 cost)
func (_BasOANN *BasOANNCaller) EvalueRootPrice(opts *bind.CallOpts, name []byte, isCustomed bool, durationInYear uint8) (struct {
	IsValid bool
	IsARoot bool
	Cost    *big.Int
}, error) {
	ret := new(struct {
		IsValid bool
		IsARoot bool
		Cost    *big.Int
	})
	out := ret
	err := _BasOANN.contract.Call(opts, out, "evalueRootPrice", name, isCustomed, durationInYear)
	return *ret, err
}

// EvalueRootPrice is a free data retrieval call binding the contract method 0xe4ede0c5.
//
// Solidity: function evalueRootPrice(bytes name, bool isCustomed, uint8 durationInYear) constant returns(bool isValid, bool isARoot, uint256 cost)
func (_BasOANN *BasOANNSession) EvalueRootPrice(name []byte, isCustomed bool, durationInYear uint8) (struct {
	IsValid bool
	IsARoot bool
	Cost    *big.Int
}, error) {
	return _BasOANN.Contract.EvalueRootPrice(&_BasOANN.CallOpts, name, isCustomed, durationInYear)
}

// EvalueRootPrice is a free data retrieval call binding the contract method 0xe4ede0c5.
//
// Solidity: function evalueRootPrice(bytes name, bool isCustomed, uint8 durationInYear) constant returns(bool isValid, bool isARoot, uint256 cost)
func (_BasOANN *BasOANNCallerSession) EvalueRootPrice(name []byte, isCustomed bool, durationInYear uint8) (struct {
	IsValid bool
	IsARoot bool
	Cost    *big.Int
}, error) {
	return _BasOANN.Contract.EvalueRootPrice(&_BasOANN.CallOpts, name, isCustomed, durationInYear)
}

// EvalueSubPrice is a free data retrieval call binding the contract method 0x413e637d.
//
// Solidity: function evalueSubPrice(bytes rName, bytes sName, uint8 durationInYear) constant returns(bool isValid, address rootOwner, bool isCustomed, uint256 cost)
func (_BasOANN *BasOANNCaller) EvalueSubPrice(opts *bind.CallOpts, rName []byte, sName []byte, durationInYear uint8) (struct {
	IsValid    bool
	RootOwner  common.Address
	IsCustomed bool
	Cost       *big.Int
}, error) {
	ret := new(struct {
		IsValid    bool
		RootOwner  common.Address
		IsCustomed bool
		Cost       *big.Int
	})
	out := ret
	err := _BasOANN.contract.Call(opts, out, "evalueSubPrice", rName, sName, durationInYear)
	return *ret, err
}

// EvalueSubPrice is a free data retrieval call binding the contract method 0x413e637d.
//
// Solidity: function evalueSubPrice(bytes rName, bytes sName, uint8 durationInYear) constant returns(bool isValid, address rootOwner, bool isCustomed, uint256 cost)
func (_BasOANN *BasOANNSession) EvalueSubPrice(rName []byte, sName []byte, durationInYear uint8) (struct {
	IsValid    bool
	RootOwner  common.Address
	IsCustomed bool
	Cost       *big.Int
}, error) {
	return _BasOANN.Contract.EvalueSubPrice(&_BasOANN.CallOpts, rName, sName, durationInYear)
}

// EvalueSubPrice is a free data retrieval call binding the contract method 0x413e637d.
//
// Solidity: function evalueSubPrice(bytes rName, bytes sName, uint8 durationInYear) constant returns(bool isValid, address rootOwner, bool isCustomed, uint256 cost)
func (_BasOANN *BasOANNCallerSession) EvalueSubPrice(rName []byte, sName []byte, durationInYear uint8) (struct {
	IsValid    bool
	RootOwner  common.Address
	IsCustomed bool
	Cost       *big.Int
}, error) {
	return _BasOANN.Contract.EvalueSubPrice(&_BasOANN.CallOpts, rName, sName, durationInYear)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_BasOANN *BasOANNCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BasOANN.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_BasOANN *BasOANNSession) Token() (common.Address, error) {
	return _BasOANN.Contract.Token(&_BasOANN.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_BasOANN *BasOANNCallerSession) Token() (common.Address, error) {
	return _BasOANN.Contract.Token(&_BasOANN.CallOpts)
}

// ValidDuration is a free data retrieval call binding the contract method 0x2b1e49e4.
//
// Solidity: function validDuration(uint8 d) constant returns(bool)
func (_BasOANN *BasOANNCaller) ValidDuration(opts *bind.CallOpts, d uint8) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BasOANN.contract.Call(opts, out, "validDuration", d)
	return *ret0, err
}

// ValidDuration is a free data retrieval call binding the contract method 0x2b1e49e4.
//
// Solidity: function validDuration(uint8 d) constant returns(bool)
func (_BasOANN *BasOANNSession) ValidDuration(d uint8) (bool, error) {
	return _BasOANN.Contract.ValidDuration(&_BasOANN.CallOpts, d)
}

// ValidDuration is a free data retrieval call binding the contract method 0x2b1e49e4.
//
// Solidity: function validDuration(uint8 d) constant returns(bool)
func (_BasOANN *BasOANNCallerSession) ValidDuration(d uint8) (bool, error) {
	return _BasOANN.Contract.ValidDuration(&_BasOANN.CallOpts, d)
}

// ChangeAccountant is a paid mutator transaction binding the contract method 0xfad68b30.
//
// Solidity: function changeAccountant(address newM) returns()
func (_BasOANN *BasOANNTransactor) ChangeAccountant(opts *bind.TransactOpts, newM common.Address) (*types.Transaction, error) {
	return _BasOANN.contract.Transact(opts, "changeAccountant", newM)
}

// ChangeAccountant is a paid mutator transaction binding the contract method 0xfad68b30.
//
// Solidity: function changeAccountant(address newM) returns()
func (_BasOANN *BasOANNSession) ChangeAccountant(newM common.Address) (*types.Transaction, error) {
	return _BasOANN.Contract.ChangeAccountant(&_BasOANN.TransactOpts, newM)
}

// ChangeAccountant is a paid mutator transaction binding the contract method 0xfad68b30.
//
// Solidity: function changeAccountant(address newM) returns()
func (_BasOANN *BasOANNTransactorSession) ChangeAccountant(newM common.Address) (*types.Transaction, error) {
	return _BasOANN.Contract.ChangeAccountant(&_BasOANN.TransactOpts, newM)
}

// ChangeAsset is a paid mutator transaction binding the contract method 0x10bbb707.
//
// Solidity: function changeAsset(address newA) returns()
func (_BasOANN *BasOANNTransactor) ChangeAsset(opts *bind.TransactOpts, newA common.Address) (*types.Transaction, error) {
	return _BasOANN.contract.Transact(opts, "changeAsset", newA)
}

// ChangeAsset is a paid mutator transaction binding the contract method 0x10bbb707.
//
// Solidity: function changeAsset(address newA) returns()
func (_BasOANN *BasOANNSession) ChangeAsset(newA common.Address) (*types.Transaction, error) {
	return _BasOANN.Contract.ChangeAsset(&_BasOANN.TransactOpts, newA)
}

// ChangeAsset is a paid mutator transaction binding the contract method 0x10bbb707.
//
// Solidity: function changeAsset(address newA) returns()
func (_BasOANN *BasOANNTransactorSession) ChangeAsset(newA common.Address) (*types.Transaction, error) {
	return _BasOANN.Contract.ChangeAsset(&_BasOANN.TransactOpts, newA)
}

// ChangeToken is a paid mutator transaction binding the contract method 0x66829b16.
//
// Solidity: function changeToken(address newT) returns()
func (_BasOANN *BasOANNTransactor) ChangeToken(opts *bind.TransactOpts, newT common.Address) (*types.Transaction, error) {
	return _BasOANN.contract.Transact(opts, "changeToken", newT)
}

// ChangeToken is a paid mutator transaction binding the contract method 0x66829b16.
//
// Solidity: function changeToken(address newT) returns()
func (_BasOANN *BasOANNSession) ChangeToken(newT common.Address) (*types.Transaction, error) {
	return _BasOANN.Contract.ChangeToken(&_BasOANN.TransactOpts, newT)
}

// ChangeToken is a paid mutator transaction binding the contract method 0x66829b16.
//
// Solidity: function changeToken(address newT) returns()
func (_BasOANN *BasOANNTransactorSession) ChangeToken(newT common.Address) (*types.Transaction, error) {
	return _BasOANN.Contract.ChangeToken(&_BasOANN.TransactOpts, newT)
}

// OpenCustomedPrice is a paid mutator transaction binding the contract method 0x3d04f2fb.
//
// Solidity: function openCustomedPrice(bytes32 rootHash, uint256 price) returns()
func (_BasOANN *BasOANNTransactor) OpenCustomedPrice(opts *bind.TransactOpts, rootHash [32]byte, price *big.Int) (*types.Transaction, error) {
	return _BasOANN.contract.Transact(opts, "openCustomedPrice", rootHash, price)
}

// OpenCustomedPrice is a paid mutator transaction binding the contract method 0x3d04f2fb.
//
// Solidity: function openCustomedPrice(bytes32 rootHash, uint256 price) returns()
func (_BasOANN *BasOANNSession) OpenCustomedPrice(rootHash [32]byte, price *big.Int) (*types.Transaction, error) {
	return _BasOANN.Contract.OpenCustomedPrice(&_BasOANN.TransactOpts, rootHash, price)
}

// OpenCustomedPrice is a paid mutator transaction binding the contract method 0x3d04f2fb.
//
// Solidity: function openCustomedPrice(bytes32 rootHash, uint256 price) returns()
func (_BasOANN *BasOANNTransactorSession) OpenCustomedPrice(rootHash [32]byte, price *big.Int) (*types.Transaction, error) {
	return _BasOANN.Contract.OpenCustomedPrice(&_BasOANN.TransactOpts, rootHash, price)
}

// RechargeRoot is a paid mutator transaction binding the contract method 0x79607e1f.
//
// Solidity: function rechargeRoot(bytes32 nameHash, uint8 durationInYear) returns()
func (_BasOANN *BasOANNTransactor) RechargeRoot(opts *bind.TransactOpts, nameHash [32]byte, durationInYear uint8) (*types.Transaction, error) {
	return _BasOANN.contract.Transact(opts, "rechargeRoot", nameHash, durationInYear)
}

// RechargeRoot is a paid mutator transaction binding the contract method 0x79607e1f.
//
// Solidity: function rechargeRoot(bytes32 nameHash, uint8 durationInYear) returns()
func (_BasOANN *BasOANNSession) RechargeRoot(nameHash [32]byte, durationInYear uint8) (*types.Transaction, error) {
	return _BasOANN.Contract.RechargeRoot(&_BasOANN.TransactOpts, nameHash, durationInYear)
}

// RechargeRoot is a paid mutator transaction binding the contract method 0x79607e1f.
//
// Solidity: function rechargeRoot(bytes32 nameHash, uint8 durationInYear) returns()
func (_BasOANN *BasOANNTransactorSession) RechargeRoot(nameHash [32]byte, durationInYear uint8) (*types.Transaction, error) {
	return _BasOANN.Contract.RechargeRoot(&_BasOANN.TransactOpts, nameHash, durationInYear)
}

// RechargeSub is a paid mutator transaction binding the contract method 0x6f6c774a.
//
// Solidity: function rechargeSub(bytes32 nameHash, uint8 durationInYear) returns()
func (_BasOANN *BasOANNTransactor) RechargeSub(opts *bind.TransactOpts, nameHash [32]byte, durationInYear uint8) (*types.Transaction, error) {
	return _BasOANN.contract.Transact(opts, "rechargeSub", nameHash, durationInYear)
}

// RechargeSub is a paid mutator transaction binding the contract method 0x6f6c774a.
//
// Solidity: function rechargeSub(bytes32 nameHash, uint8 durationInYear) returns()
func (_BasOANN *BasOANNSession) RechargeSub(nameHash [32]byte, durationInYear uint8) (*types.Transaction, error) {
	return _BasOANN.Contract.RechargeSub(&_BasOANN.TransactOpts, nameHash, durationInYear)
}

// RechargeSub is a paid mutator transaction binding the contract method 0x6f6c774a.
//
// Solidity: function rechargeSub(bytes32 nameHash, uint8 durationInYear) returns()
func (_BasOANN *BasOANNTransactorSession) RechargeSub(nameHash [32]byte, durationInYear uint8) (*types.Transaction, error) {
	return _BasOANN.Contract.RechargeSub(&_BasOANN.TransactOpts, nameHash, durationInYear)
}

// RegisterRoot is a paid mutator transaction binding the contract method 0x70a1495c.
//
// Solidity: function registerRoot(bytes name, bool isOpen, bool isCustomed, uint256 cusPrice, uint8 durationInYear) returns()
func (_BasOANN *BasOANNTransactor) RegisterRoot(opts *bind.TransactOpts, name []byte, isOpen bool, isCustomed bool, cusPrice *big.Int, durationInYear uint8) (*types.Transaction, error) {
	return _BasOANN.contract.Transact(opts, "registerRoot", name, isOpen, isCustomed, cusPrice, durationInYear)
}

// RegisterRoot is a paid mutator transaction binding the contract method 0x70a1495c.
//
// Solidity: function registerRoot(bytes name, bool isOpen, bool isCustomed, uint256 cusPrice, uint8 durationInYear) returns()
func (_BasOANN *BasOANNSession) RegisterRoot(name []byte, isOpen bool, isCustomed bool, cusPrice *big.Int, durationInYear uint8) (*types.Transaction, error) {
	return _BasOANN.Contract.RegisterRoot(&_BasOANN.TransactOpts, name, isOpen, isCustomed, cusPrice, durationInYear)
}

// RegisterRoot is a paid mutator transaction binding the contract method 0x70a1495c.
//
// Solidity: function registerRoot(bytes name, bool isOpen, bool isCustomed, uint256 cusPrice, uint8 durationInYear) returns()
func (_BasOANN *BasOANNTransactorSession) RegisterRoot(name []byte, isOpen bool, isCustomed bool, cusPrice *big.Int, durationInYear uint8) (*types.Transaction, error) {
	return _BasOANN.Contract.RegisterRoot(&_BasOANN.TransactOpts, name, isOpen, isCustomed, cusPrice, durationInYear)
}

// RegisterSub is a paid mutator transaction binding the contract method 0xee9a512a.
//
// Solidity: function registerSub(bytes rName, bytes sName, uint8 durationInYear) returns()
func (_BasOANN *BasOANNTransactor) RegisterSub(opts *bind.TransactOpts, rName []byte, sName []byte, durationInYear uint8) (*types.Transaction, error) {
	return _BasOANN.contract.Transact(opts, "registerSub", rName, sName, durationInYear)
}

// RegisterSub is a paid mutator transaction binding the contract method 0xee9a512a.
//
// Solidity: function registerSub(bytes rName, bytes sName, uint8 durationInYear) returns()
func (_BasOANN *BasOANNSession) RegisterSub(rName []byte, sName []byte, durationInYear uint8) (*types.Transaction, error) {
	return _BasOANN.Contract.RegisterSub(&_BasOANN.TransactOpts, rName, sName, durationInYear)
}

// RegisterSub is a paid mutator transaction binding the contract method 0xee9a512a.
//
// Solidity: function registerSub(bytes rName, bytes sName, uint8 durationInYear) returns()
func (_BasOANN *BasOANNTransactorSession) RegisterSub(rName []byte, sName []byte, durationInYear uint8) (*types.Transaction, error) {
	return _BasOANN.Contract.RegisterSub(&_BasOANN.TransactOpts, rName, sName, durationInYear)
}

// SetARootGas is a paid mutator transaction binding the contract method 0x9c626de4.
//
// Solidity: function setARootGas(uint256 newGas) returns()
func (_BasOANN *BasOANNTransactor) SetARootGas(opts *bind.TransactOpts, newGas *big.Int) (*types.Transaction, error) {
	return _BasOANN.contract.Transact(opts, "setARootGas", newGas)
}

// SetARootGas is a paid mutator transaction binding the contract method 0x9c626de4.
//
// Solidity: function setARootGas(uint256 newGas) returns()
func (_BasOANN *BasOANNSession) SetARootGas(newGas *big.Int) (*types.Transaction, error) {
	return _BasOANN.Contract.SetARootGas(&_BasOANN.TransactOpts, newGas)
}

// SetARootGas is a paid mutator transaction binding the contract method 0x9c626de4.
//
// Solidity: function setARootGas(uint256 newGas) returns()
func (_BasOANN *BasOANNTransactorSession) SetARootGas(newGas *big.Int) (*types.Transaction, error) {
	return _BasOANN.Contract.SetARootGas(&_BasOANN.TransactOpts, newGas)
}

// SetBRootGas is a paid mutator transaction binding the contract method 0xa23a7245.
//
// Solidity: function setBRootGas(uint256 newGas) returns()
func (_BasOANN *BasOANNTransactor) SetBRootGas(opts *bind.TransactOpts, newGas *big.Int) (*types.Transaction, error) {
	return _BasOANN.contract.Transact(opts, "setBRootGas", newGas)
}

// SetBRootGas is a paid mutator transaction binding the contract method 0xa23a7245.
//
// Solidity: function setBRootGas(uint256 newGas) returns()
func (_BasOANN *BasOANNSession) SetBRootGas(newGas *big.Int) (*types.Transaction, error) {
	return _BasOANN.Contract.SetBRootGas(&_BasOANN.TransactOpts, newGas)
}

// SetBRootGas is a paid mutator transaction binding the contract method 0xa23a7245.
//
// Solidity: function setBRootGas(uint256 newGas) returns()
func (_BasOANN *BasOANNTransactorSession) SetBRootGas(newGas *big.Int) (*types.Transaction, error) {
	return _BasOANN.Contract.SetBRootGas(&_BasOANN.TransactOpts, newGas)
}

// TransferContractOwnership is a paid mutator transaction binding the contract method 0xa843c51f.
//
// Solidity: function transferContractOwnership(address newOwner) returns()
func (_BasOANN *BasOANNTransactor) TransferContractOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BasOANN.contract.Transact(opts, "transferContractOwnership", newOwner)
}

// TransferContractOwnership is a paid mutator transaction binding the contract method 0xa843c51f.
//
// Solidity: function transferContractOwnership(address newOwner) returns()
func (_BasOANN *BasOANNSession) TransferContractOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BasOANN.Contract.TransferContractOwnership(&_BasOANN.TransactOpts, newOwner)
}

// TransferContractOwnership is a paid mutator transaction binding the contract method 0xa843c51f.
//
// Solidity: function transferContractOwnership(address newOwner) returns()
func (_BasOANN *BasOANNTransactorSession) TransferContractOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BasOANN.Contract.TransferContractOwnership(&_BasOANN.TransactOpts, newOwner)
}
