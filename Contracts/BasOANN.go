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
const BasOANNABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"len\",\"type\":\"uint8\"}],\"name\":\"setRareTypeLength\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"_a_changeAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"y\",\"type\":\"uint8\"}],\"name\":\"validDuration\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newGas\",\"type\":\"uint256\"}],\"name\":\"setCustomedPriceGas\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes4\",\"name\":\"ipv4\",\"type\":\"bytes4\"},{\"internalType\":\"bytes16\",\"name\":\"ipv6\",\"type\":\"bytes16\"},{\"internalType\":\"bytes\",\"name\":\"bca\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"opData\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"aliasName\",\"type\":\"string\"}],\"name\":\"setRecord\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"openCustomedPrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"rName\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sName\",\"type\":\"bytes\"},{\"internalType\":\"uint8\",\"name\":\"durationInYear\",\"type\":\"uint8\"}],\"name\":\"evalueSubPrice\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"totalName\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"rootHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"rootOwner\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isCustomed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"exist\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CUSTOMED_PRICE_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contractCaller\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"year\",\"type\":\"uint256\"}],\"name\":\"setMaxYear\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_YEAR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isOpen\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isCustomed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"cusPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"durationInYear\",\"type\":\"uint8\"}],\"name\":\"registerRoot\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"conAddr\",\"type\":\"address\"}],\"name\":\"_a_changeContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"AROOT_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newGas\",\"type\":\"uint256\"}],\"name\":\"setARootGas\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newGas\",\"type\":\"uint256\"}],\"name\":\"setBRootGas\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"checkAssociatedContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"BasToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"BasOwnership\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"BasAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"BasDNS\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"BasMiner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"BasRule\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"RARE_TYPE_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newGas\",\"type\":\"uint256\"}],\"name\":\"setSubGas\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"internalType\":\"uint8\",\"name\":\"durationInYear\",\"type\":\"uint8\"}],\"name\":\"recharge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"}],\"name\":\"closeCustomedPrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"}],\"name\":\"closeToPublic\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isCustomed\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"durationInYear\",\"type\":\"uint8\"}],\"name\":\"evalueRootPrice\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"exist\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"}],\"name\":\"openToPublic\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"rName\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sName\",\"type\":\"bytes\"},{\"internalType\":\"uint8\",\"name\":\"durationInYear\",\"type\":\"uint8\"}],\"name\":\"registerSub\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"SUB_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BROOT_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_t\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_o\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_d\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_m\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_r\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"option\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Paid\",\"type\":\"event\"}]"

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

// MAXYEAR is a free data retrieval call binding the contract method 0x61c9511e.
//
// Solidity: function MAX_YEAR() constant returns(uint256)
func (_BasOANN *BasOANNCaller) MAXYEAR(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BasOANN.contract.Call(opts, out, "MAX_YEAR")
	return *ret0, err
}

// MAXYEAR is a free data retrieval call binding the contract method 0x61c9511e.
//
// Solidity: function MAX_YEAR() constant returns(uint256)
func (_BasOANN *BasOANNSession) MAXYEAR() (*big.Int, error) {
	return _BasOANN.Contract.MAXYEAR(&_BasOANN.CallOpts)
}

// MAXYEAR is a free data retrieval call binding the contract method 0x61c9511e.
//
// Solidity: function MAX_YEAR() constant returns(uint256)
func (_BasOANN *BasOANNCallerSession) MAXYEAR() (*big.Int, error) {
	return _BasOANN.Contract.MAXYEAR(&_BasOANN.CallOpts)
}

// RARETYPELENGTH is a free data retrieval call binding the contract method 0xa59ea829.
//
// Solidity: function RARE_TYPE_LENGTH() constant returns(uint256)
func (_BasOANN *BasOANNCaller) RARETYPELENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BasOANN.contract.Call(opts, out, "RARE_TYPE_LENGTH")
	return *ret0, err
}

// RARETYPELENGTH is a free data retrieval call binding the contract method 0xa59ea829.
//
// Solidity: function RARE_TYPE_LENGTH() constant returns(uint256)
func (_BasOANN *BasOANNSession) RARETYPELENGTH() (*big.Int, error) {
	return _BasOANN.Contract.RARETYPELENGTH(&_BasOANN.CallOpts)
}

// RARETYPELENGTH is a free data retrieval call binding the contract method 0xa59ea829.
//
// Solidity: function RARE_TYPE_LENGTH() constant returns(uint256)
func (_BasOANN *BasOANNCallerSession) RARETYPELENGTH() (*big.Int, error) {
	return _BasOANN.Contract.RARETYPELENGTH(&_BasOANN.CallOpts)
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

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_BasOANN *BasOANNCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BasOANN.contract.Call(opts, out, "admin")
	return *ret0, err
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_BasOANN *BasOANNSession) Admin() (common.Address, error) {
	return _BasOANN.Contract.Admin(&_BasOANN.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_BasOANN *BasOANNCallerSession) Admin() (common.Address, error) {
	return _BasOANN.Contract.Admin(&_BasOANN.CallOpts)
}

// CheckAssociatedContractAddress is a free data retrieval call binding the contract method 0xa3e0c839.
//
// Solidity: function checkAssociatedContractAddress() constant returns(address BasToken, address BasOwnership, address BasAsset, address BasDNS, address BasMiner, address BasRule)
func (_BasOANN *BasOANNCaller) CheckAssociatedContractAddress(opts *bind.CallOpts) (struct {
	BasToken     common.Address
	BasOwnership common.Address
	BasAsset     common.Address
	BasDNS       common.Address
	BasMiner     common.Address
	BasRule      common.Address
}, error) {
	ret := new(struct {
		BasToken     common.Address
		BasOwnership common.Address
		BasAsset     common.Address
		BasDNS       common.Address
		BasMiner     common.Address
		BasRule      common.Address
	})
	out := ret
	err := _BasOANN.contract.Call(opts, out, "checkAssociatedContractAddress")
	return *ret, err
}

// CheckAssociatedContractAddress is a free data retrieval call binding the contract method 0xa3e0c839.
//
// Solidity: function checkAssociatedContractAddress() constant returns(address BasToken, address BasOwnership, address BasAsset, address BasDNS, address BasMiner, address BasRule)
func (_BasOANN *BasOANNSession) CheckAssociatedContractAddress() (struct {
	BasToken     common.Address
	BasOwnership common.Address
	BasAsset     common.Address
	BasDNS       common.Address
	BasMiner     common.Address
	BasRule      common.Address
}, error) {
	return _BasOANN.Contract.CheckAssociatedContractAddress(&_BasOANN.CallOpts)
}

// CheckAssociatedContractAddress is a free data retrieval call binding the contract method 0xa3e0c839.
//
// Solidity: function checkAssociatedContractAddress() constant returns(address BasToken, address BasOwnership, address BasAsset, address BasDNS, address BasMiner, address BasRule)
func (_BasOANN *BasOANNCallerSession) CheckAssociatedContractAddress() (struct {
	BasToken     common.Address
	BasOwnership common.Address
	BasAsset     common.Address
	BasDNS       common.Address
	BasMiner     common.Address
	BasRule      common.Address
}, error) {
	return _BasOANN.Contract.CheckAssociatedContractAddress(&_BasOANN.CallOpts)
}

// ContractCaller is a free data retrieval call binding the contract method 0x54cf428a.
//
// Solidity: function contractCaller() constant returns(address)
func (_BasOANN *BasOANNCaller) ContractCaller(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BasOANN.contract.Call(opts, out, "contractCaller")
	return *ret0, err
}

// ContractCaller is a free data retrieval call binding the contract method 0x54cf428a.
//
// Solidity: function contractCaller() constant returns(address)
func (_BasOANN *BasOANNSession) ContractCaller() (common.Address, error) {
	return _BasOANN.Contract.ContractCaller(&_BasOANN.CallOpts)
}

// ContractCaller is a free data retrieval call binding the contract method 0x54cf428a.
//
// Solidity: function contractCaller() constant returns(address)
func (_BasOANN *BasOANNCallerSession) ContractCaller() (common.Address, error) {
	return _BasOANN.Contract.ContractCaller(&_BasOANN.CallOpts)
}

// EvalueRootPrice is a free data retrieval call binding the contract method 0xe4ede0c5.
//
// Solidity: function evalueRootPrice(bytes name, bool isCustomed, uint8 durationInYear) constant returns(bytes32 nameHash, bool isValid, uint256 cost, bool exist)
func (_BasOANN *BasOANNCaller) EvalueRootPrice(opts *bind.CallOpts, name []byte, isCustomed bool, durationInYear uint8) (struct {
	NameHash [32]byte
	IsValid  bool
	Cost     *big.Int
	Exist    bool
}, error) {
	ret := new(struct {
		NameHash [32]byte
		IsValid  bool
		Cost     *big.Int
		Exist    bool
	})
	out := ret
	err := _BasOANN.contract.Call(opts, out, "evalueRootPrice", name, isCustomed, durationInYear)
	return *ret, err
}

// EvalueRootPrice is a free data retrieval call binding the contract method 0xe4ede0c5.
//
// Solidity: function evalueRootPrice(bytes name, bool isCustomed, uint8 durationInYear) constant returns(bytes32 nameHash, bool isValid, uint256 cost, bool exist)
func (_BasOANN *BasOANNSession) EvalueRootPrice(name []byte, isCustomed bool, durationInYear uint8) (struct {
	NameHash [32]byte
	IsValid  bool
	Cost     *big.Int
	Exist    bool
}, error) {
	return _BasOANN.Contract.EvalueRootPrice(&_BasOANN.CallOpts, name, isCustomed, durationInYear)
}

// EvalueRootPrice is a free data retrieval call binding the contract method 0xe4ede0c5.
//
// Solidity: function evalueRootPrice(bytes name, bool isCustomed, uint8 durationInYear) constant returns(bytes32 nameHash, bool isValid, uint256 cost, bool exist)
func (_BasOANN *BasOANNCallerSession) EvalueRootPrice(name []byte, isCustomed bool, durationInYear uint8) (struct {
	NameHash [32]byte
	IsValid  bool
	Cost     *big.Int
	Exist    bool
}, error) {
	return _BasOANN.Contract.EvalueRootPrice(&_BasOANN.CallOpts, name, isCustomed, durationInYear)
}

// EvalueSubPrice is a free data retrieval call binding the contract method 0x413e637d.
//
// Solidity: function evalueSubPrice(bytes rName, bytes sName, uint8 durationInYear) constant returns(bytes32 nameHash, bytes totalName, bytes32 rootHash, bool isValid, address rootOwner, bool isCustomed, uint256 cost, bool exist)
func (_BasOANN *BasOANNCaller) EvalueSubPrice(opts *bind.CallOpts, rName []byte, sName []byte, durationInYear uint8) (struct {
	NameHash   [32]byte
	TotalName  []byte
	RootHash   [32]byte
	IsValid    bool
	RootOwner  common.Address
	IsCustomed bool
	Cost       *big.Int
	Exist      bool
}, error) {
	ret := new(struct {
		NameHash   [32]byte
		TotalName  []byte
		RootHash   [32]byte
		IsValid    bool
		RootOwner  common.Address
		IsCustomed bool
		Cost       *big.Int
		Exist      bool
	})
	out := ret
	err := _BasOANN.contract.Call(opts, out, "evalueSubPrice", rName, sName, durationInYear)
	return *ret, err
}

// EvalueSubPrice is a free data retrieval call binding the contract method 0x413e637d.
//
// Solidity: function evalueSubPrice(bytes rName, bytes sName, uint8 durationInYear) constant returns(bytes32 nameHash, bytes totalName, bytes32 rootHash, bool isValid, address rootOwner, bool isCustomed, uint256 cost, bool exist)
func (_BasOANN *BasOANNSession) EvalueSubPrice(rName []byte, sName []byte, durationInYear uint8) (struct {
	NameHash   [32]byte
	TotalName  []byte
	RootHash   [32]byte
	IsValid    bool
	RootOwner  common.Address
	IsCustomed bool
	Cost       *big.Int
	Exist      bool
}, error) {
	return _BasOANN.Contract.EvalueSubPrice(&_BasOANN.CallOpts, rName, sName, durationInYear)
}

// EvalueSubPrice is a free data retrieval call binding the contract method 0x413e637d.
//
// Solidity: function evalueSubPrice(bytes rName, bytes sName, uint8 durationInYear) constant returns(bytes32 nameHash, bytes totalName, bytes32 rootHash, bool isValid, address rootOwner, bool isCustomed, uint256 cost, bool exist)
func (_BasOANN *BasOANNCallerSession) EvalueSubPrice(rName []byte, sName []byte, durationInYear uint8) (struct {
	NameHash   [32]byte
	TotalName  []byte
	RootHash   [32]byte
	IsValid    bool
	RootOwner  common.Address
	IsCustomed bool
	Cost       *big.Int
	Exist      bool
}, error) {
	return _BasOANN.Contract.EvalueSubPrice(&_BasOANN.CallOpts, rName, sName, durationInYear)
}

// ValidDuration is a free data retrieval call binding the contract method 0x2b1e49e4.
//
// Solidity: function validDuration(uint8 y) constant returns(bool)
func (_BasOANN *BasOANNCaller) ValidDuration(opts *bind.CallOpts, y uint8) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BasOANN.contract.Call(opts, out, "validDuration", y)
	return *ret0, err
}

// ValidDuration is a free data retrieval call binding the contract method 0x2b1e49e4.
//
// Solidity: function validDuration(uint8 y) constant returns(bool)
func (_BasOANN *BasOANNSession) ValidDuration(y uint8) (bool, error) {
	return _BasOANN.Contract.ValidDuration(&_BasOANN.CallOpts, y)
}

// ValidDuration is a free data retrieval call binding the contract method 0x2b1e49e4.
//
// Solidity: function validDuration(uint8 y) constant returns(bool)
func (_BasOANN *BasOANNCallerSession) ValidDuration(y uint8) (bool, error) {
	return _BasOANN.Contract.ValidDuration(&_BasOANN.CallOpts, y)
}

// AChangeAdmin is a paid mutator transaction binding the contract method 0x1062f6cc.
//
// Solidity: function _a_changeAdmin(address newOwner) returns()
func (_BasOANN *BasOANNTransactor) AChangeAdmin(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BasOANN.contract.Transact(opts, "_a_changeAdmin", newOwner)
}

// AChangeAdmin is a paid mutator transaction binding the contract method 0x1062f6cc.
//
// Solidity: function _a_changeAdmin(address newOwner) returns()
func (_BasOANN *BasOANNSession) AChangeAdmin(newOwner common.Address) (*types.Transaction, error) {
	return _BasOANN.Contract.AChangeAdmin(&_BasOANN.TransactOpts, newOwner)
}

// AChangeAdmin is a paid mutator transaction binding the contract method 0x1062f6cc.
//
// Solidity: function _a_changeAdmin(address newOwner) returns()
func (_BasOANN *BasOANNTransactorSession) AChangeAdmin(newOwner common.Address) (*types.Transaction, error) {
	return _BasOANN.Contract.AChangeAdmin(&_BasOANN.TransactOpts, newOwner)
}

// AChangeContract is a paid mutator transaction binding the contract method 0x733c1950.
//
// Solidity: function _a_changeContract(address conAddr) returns()
func (_BasOANN *BasOANNTransactor) AChangeContract(opts *bind.TransactOpts, conAddr common.Address) (*types.Transaction, error) {
	return _BasOANN.contract.Transact(opts, "_a_changeContract", conAddr)
}

// AChangeContract is a paid mutator transaction binding the contract method 0x733c1950.
//
// Solidity: function _a_changeContract(address conAddr) returns()
func (_BasOANN *BasOANNSession) AChangeContract(conAddr common.Address) (*types.Transaction, error) {
	return _BasOANN.Contract.AChangeContract(&_BasOANN.TransactOpts, conAddr)
}

// AChangeContract is a paid mutator transaction binding the contract method 0x733c1950.
//
// Solidity: function _a_changeContract(address conAddr) returns()
func (_BasOANN *BasOANNTransactorSession) AChangeContract(conAddr common.Address) (*types.Transaction, error) {
	return _BasOANN.Contract.AChangeContract(&_BasOANN.TransactOpts, conAddr)
}

// CloseCustomedPrice is a paid mutator transaction binding the contract method 0xc0bd516d.
//
// Solidity: function closeCustomedPrice(bytes32 nameHash) returns()
func (_BasOANN *BasOANNTransactor) CloseCustomedPrice(opts *bind.TransactOpts, nameHash [32]byte) (*types.Transaction, error) {
	return _BasOANN.contract.Transact(opts, "closeCustomedPrice", nameHash)
}

// CloseCustomedPrice is a paid mutator transaction binding the contract method 0xc0bd516d.
//
// Solidity: function closeCustomedPrice(bytes32 nameHash) returns()
func (_BasOANN *BasOANNSession) CloseCustomedPrice(nameHash [32]byte) (*types.Transaction, error) {
	return _BasOANN.Contract.CloseCustomedPrice(&_BasOANN.TransactOpts, nameHash)
}

// CloseCustomedPrice is a paid mutator transaction binding the contract method 0xc0bd516d.
//
// Solidity: function closeCustomedPrice(bytes32 nameHash) returns()
func (_BasOANN *BasOANNTransactorSession) CloseCustomedPrice(nameHash [32]byte) (*types.Transaction, error) {
	return _BasOANN.Contract.CloseCustomedPrice(&_BasOANN.TransactOpts, nameHash)
}

// CloseToPublic is a paid mutator transaction binding the contract method 0xd0f0b43d.
//
// Solidity: function closeToPublic(bytes32 nameHash) returns()
func (_BasOANN *BasOANNTransactor) CloseToPublic(opts *bind.TransactOpts, nameHash [32]byte) (*types.Transaction, error) {
	return _BasOANN.contract.Transact(opts, "closeToPublic", nameHash)
}

// CloseToPublic is a paid mutator transaction binding the contract method 0xd0f0b43d.
//
// Solidity: function closeToPublic(bytes32 nameHash) returns()
func (_BasOANN *BasOANNSession) CloseToPublic(nameHash [32]byte) (*types.Transaction, error) {
	return _BasOANN.Contract.CloseToPublic(&_BasOANN.TransactOpts, nameHash)
}

// CloseToPublic is a paid mutator transaction binding the contract method 0xd0f0b43d.
//
// Solidity: function closeToPublic(bytes32 nameHash) returns()
func (_BasOANN *BasOANNTransactorSession) CloseToPublic(nameHash [32]byte) (*types.Transaction, error) {
	return _BasOANN.Contract.CloseToPublic(&_BasOANN.TransactOpts, nameHash)
}

// OpenCustomedPrice is a paid mutator transaction binding the contract method 0x3d04f2fb.
//
// Solidity: function openCustomedPrice(bytes32 nameHash, uint256 price) returns()
func (_BasOANN *BasOANNTransactor) OpenCustomedPrice(opts *bind.TransactOpts, nameHash [32]byte, price *big.Int) (*types.Transaction, error) {
	return _BasOANN.contract.Transact(opts, "openCustomedPrice", nameHash, price)
}

// OpenCustomedPrice is a paid mutator transaction binding the contract method 0x3d04f2fb.
//
// Solidity: function openCustomedPrice(bytes32 nameHash, uint256 price) returns()
func (_BasOANN *BasOANNSession) OpenCustomedPrice(nameHash [32]byte, price *big.Int) (*types.Transaction, error) {
	return _BasOANN.Contract.OpenCustomedPrice(&_BasOANN.TransactOpts, nameHash, price)
}

// OpenCustomedPrice is a paid mutator transaction binding the contract method 0x3d04f2fb.
//
// Solidity: function openCustomedPrice(bytes32 nameHash, uint256 price) returns()
func (_BasOANN *BasOANNTransactorSession) OpenCustomedPrice(nameHash [32]byte, price *big.Int) (*types.Transaction, error) {
	return _BasOANN.Contract.OpenCustomedPrice(&_BasOANN.TransactOpts, nameHash, price)
}

// OpenToPublic is a paid mutator transaction binding the contract method 0xed26e5ae.
//
// Solidity: function openToPublic(bytes32 nameHash) returns()
func (_BasOANN *BasOANNTransactor) OpenToPublic(opts *bind.TransactOpts, nameHash [32]byte) (*types.Transaction, error) {
	return _BasOANN.contract.Transact(opts, "openToPublic", nameHash)
}

// OpenToPublic is a paid mutator transaction binding the contract method 0xed26e5ae.
//
// Solidity: function openToPublic(bytes32 nameHash) returns()
func (_BasOANN *BasOANNSession) OpenToPublic(nameHash [32]byte) (*types.Transaction, error) {
	return _BasOANN.Contract.OpenToPublic(&_BasOANN.TransactOpts, nameHash)
}

// OpenToPublic is a paid mutator transaction binding the contract method 0xed26e5ae.
//
// Solidity: function openToPublic(bytes32 nameHash) returns()
func (_BasOANN *BasOANNTransactorSession) OpenToPublic(nameHash [32]byte) (*types.Transaction, error) {
	return _BasOANN.Contract.OpenToPublic(&_BasOANN.TransactOpts, nameHash)
}

// Recharge is a paid mutator transaction binding the contract method 0xb84f5f20.
//
// Solidity: function recharge(bytes name, uint8 durationInYear) returns()
func (_BasOANN *BasOANNTransactor) Recharge(opts *bind.TransactOpts, name []byte, durationInYear uint8) (*types.Transaction, error) {
	return _BasOANN.contract.Transact(opts, "recharge", name, durationInYear)
}

// Recharge is a paid mutator transaction binding the contract method 0xb84f5f20.
//
// Solidity: function recharge(bytes name, uint8 durationInYear) returns()
func (_BasOANN *BasOANNSession) Recharge(name []byte, durationInYear uint8) (*types.Transaction, error) {
	return _BasOANN.Contract.Recharge(&_BasOANN.TransactOpts, name, durationInYear)
}

// Recharge is a paid mutator transaction binding the contract method 0xb84f5f20.
//
// Solidity: function recharge(bytes name, uint8 durationInYear) returns()
func (_BasOANN *BasOANNTransactorSession) Recharge(name []byte, durationInYear uint8) (*types.Transaction, error) {
	return _BasOANN.Contract.Recharge(&_BasOANN.TransactOpts, name, durationInYear)
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

// SetCustomedPriceGas is a paid mutator transaction binding the contract method 0x30ffb316.
//
// Solidity: function setCustomedPriceGas(uint256 newGas) returns()
func (_BasOANN *BasOANNTransactor) SetCustomedPriceGas(opts *bind.TransactOpts, newGas *big.Int) (*types.Transaction, error) {
	return _BasOANN.contract.Transact(opts, "setCustomedPriceGas", newGas)
}

// SetCustomedPriceGas is a paid mutator transaction binding the contract method 0x30ffb316.
//
// Solidity: function setCustomedPriceGas(uint256 newGas) returns()
func (_BasOANN *BasOANNSession) SetCustomedPriceGas(newGas *big.Int) (*types.Transaction, error) {
	return _BasOANN.Contract.SetCustomedPriceGas(&_BasOANN.TransactOpts, newGas)
}

// SetCustomedPriceGas is a paid mutator transaction binding the contract method 0x30ffb316.
//
// Solidity: function setCustomedPriceGas(uint256 newGas) returns()
func (_BasOANN *BasOANNTransactorSession) SetCustomedPriceGas(newGas *big.Int) (*types.Transaction, error) {
	return _BasOANN.Contract.SetCustomedPriceGas(&_BasOANN.TransactOpts, newGas)
}

// SetMaxYear is a paid mutator transaction binding the contract method 0x6000e099.
//
// Solidity: function setMaxYear(uint256 year) returns()
func (_BasOANN *BasOANNTransactor) SetMaxYear(opts *bind.TransactOpts, year *big.Int) (*types.Transaction, error) {
	return _BasOANN.contract.Transact(opts, "setMaxYear", year)
}

// SetMaxYear is a paid mutator transaction binding the contract method 0x6000e099.
//
// Solidity: function setMaxYear(uint256 year) returns()
func (_BasOANN *BasOANNSession) SetMaxYear(year *big.Int) (*types.Transaction, error) {
	return _BasOANN.Contract.SetMaxYear(&_BasOANN.TransactOpts, year)
}

// SetMaxYear is a paid mutator transaction binding the contract method 0x6000e099.
//
// Solidity: function setMaxYear(uint256 year) returns()
func (_BasOANN *BasOANNTransactorSession) SetMaxYear(year *big.Int) (*types.Transaction, error) {
	return _BasOANN.Contract.SetMaxYear(&_BasOANN.TransactOpts, year)
}

// SetRareTypeLength is a paid mutator transaction binding the contract method 0x05790fd1.
//
// Solidity: function setRareTypeLength(uint8 len) returns()
func (_BasOANN *BasOANNTransactor) SetRareTypeLength(opts *bind.TransactOpts, len uint8) (*types.Transaction, error) {
	return _BasOANN.contract.Transact(opts, "setRareTypeLength", len)
}

// SetRareTypeLength is a paid mutator transaction binding the contract method 0x05790fd1.
//
// Solidity: function setRareTypeLength(uint8 len) returns()
func (_BasOANN *BasOANNSession) SetRareTypeLength(len uint8) (*types.Transaction, error) {
	return _BasOANN.Contract.SetRareTypeLength(&_BasOANN.TransactOpts, len)
}

// SetRareTypeLength is a paid mutator transaction binding the contract method 0x05790fd1.
//
// Solidity: function setRareTypeLength(uint8 len) returns()
func (_BasOANN *BasOANNTransactorSession) SetRareTypeLength(len uint8) (*types.Transaction, error) {
	return _BasOANN.Contract.SetRareTypeLength(&_BasOANN.TransactOpts, len)
}

// SetRecord is a paid mutator transaction binding the contract method 0x3834ae56.
//
// Solidity: function setRecord(bytes32 nameHash, bytes4 ipv4, bytes16 ipv6, bytes bca, bytes opData, string aliasName) returns()
func (_BasOANN *BasOANNTransactor) SetRecord(opts *bind.TransactOpts, nameHash [32]byte, ipv4 [4]byte, ipv6 [16]byte, bca []byte, opData []byte, aliasName string) (*types.Transaction, error) {
	return _BasOANN.contract.Transact(opts, "setRecord", nameHash, ipv4, ipv6, bca, opData, aliasName)
}

// SetRecord is a paid mutator transaction binding the contract method 0x3834ae56.
//
// Solidity: function setRecord(bytes32 nameHash, bytes4 ipv4, bytes16 ipv6, bytes bca, bytes opData, string aliasName) returns()
func (_BasOANN *BasOANNSession) SetRecord(nameHash [32]byte, ipv4 [4]byte, ipv6 [16]byte, bca []byte, opData []byte, aliasName string) (*types.Transaction, error) {
	return _BasOANN.Contract.SetRecord(&_BasOANN.TransactOpts, nameHash, ipv4, ipv6, bca, opData, aliasName)
}

// SetRecord is a paid mutator transaction binding the contract method 0x3834ae56.
//
// Solidity: function setRecord(bytes32 nameHash, bytes4 ipv4, bytes16 ipv6, bytes bca, bytes opData, string aliasName) returns()
func (_BasOANN *BasOANNTransactorSession) SetRecord(nameHash [32]byte, ipv4 [4]byte, ipv6 [16]byte, bca []byte, opData []byte, aliasName string) (*types.Transaction, error) {
	return _BasOANN.Contract.SetRecord(&_BasOANN.TransactOpts, nameHash, ipv4, ipv6, bca, opData, aliasName)
}

// SetSubGas is a paid mutator transaction binding the contract method 0xa7a6907b.
//
// Solidity: function setSubGas(uint256 newGas) returns()
func (_BasOANN *BasOANNTransactor) SetSubGas(opts *bind.TransactOpts, newGas *big.Int) (*types.Transaction, error) {
	return _BasOANN.contract.Transact(opts, "setSubGas", newGas)
}

// SetSubGas is a paid mutator transaction binding the contract method 0xa7a6907b.
//
// Solidity: function setSubGas(uint256 newGas) returns()
func (_BasOANN *BasOANNSession) SetSubGas(newGas *big.Int) (*types.Transaction, error) {
	return _BasOANN.Contract.SetSubGas(&_BasOANN.TransactOpts, newGas)
}

// SetSubGas is a paid mutator transaction binding the contract method 0xa7a6907b.
//
// Solidity: function setSubGas(uint256 newGas) returns()
func (_BasOANN *BasOANNTransactorSession) SetSubGas(newGas *big.Int) (*types.Transaction, error) {
	return _BasOANN.Contract.SetSubGas(&_BasOANN.TransactOpts, newGas)
}

// BasOANNPaidIterator is returned from FilterPaid and is used to iterate over the raw logs and unpacked data for Paid events raised by the BasOANN contract.
type BasOANNPaidIterator struct {
	Event *BasOANNPaid // Event containing the contract specifics and raw log

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
func (it *BasOANNPaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasOANNPaid)
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
		it.Event = new(BasOANNPaid)
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
func (it *BasOANNPaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasOANNPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasOANNPaid represents a Paid event raised by the BasOANN contract.
type BasOANNPaid struct {
	Payer  common.Address
	Name   []byte
	Option string
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPaid is a free log retrieval operation binding the contract event 0x6a6b5d53d5a52b3e062c7acfa65003e51072e45029558311e0cff6219afa2351.
//
// Solidity: event Paid(address payer, bytes name, string option, uint256 amount)
func (_BasOANN *BasOANNFilterer) FilterPaid(opts *bind.FilterOpts) (*BasOANNPaidIterator, error) {

	logs, sub, err := _BasOANN.contract.FilterLogs(opts, "Paid")
	if err != nil {
		return nil, err
	}
	return &BasOANNPaidIterator{contract: _BasOANN.contract, event: "Paid", logs: logs, sub: sub}, nil
}

// WatchPaid is a free log subscription operation binding the contract event 0x6a6b5d53d5a52b3e062c7acfa65003e51072e45029558311e0cff6219afa2351.
//
// Solidity: event Paid(address payer, bytes name, string option, uint256 amount)
func (_BasOANN *BasOANNFilterer) WatchPaid(opts *bind.WatchOpts, sink chan<- *BasOANNPaid) (event.Subscription, error) {

	logs, sub, err := _BasOANN.contract.WatchLogs(opts, "Paid")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasOANNPaid)
				if err := _BasOANN.contract.UnpackLog(event, "Paid", log); err != nil {
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

// ParsePaid is a log parse operation binding the contract event 0x6a6b5d53d5a52b3e062c7acfa65003e51072e45029558311e0cff6219afa2351.
//
// Solidity: event Paid(address payer, bytes name, string option, uint256 amount)
func (_BasOANN *BasOANNFilterer) ParsePaid(log types.Log) (*BasOANNPaid, error) {
	event := new(BasOANNPaid)
	if err := _BasOANN.contract.UnpackLog(event, "Paid", log); err != nil {
		return nil, err
	}
	return event, nil
}
