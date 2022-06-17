// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package kmanage

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
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// IAuthABI is the input ABI used to generate the binding from.
const IAuthABI = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_h\",\"type\":\"bytes32\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"perm\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IAuthFuncSigs maps the 4-byte function signature to its string representation.
var IAuthFuncSigs = map[string]string{
	"a96bba9d": "perm(bytes32,bytes[5])",
}

// IAuth is an auto generated Go binding around an Ethereum contract.
type IAuth struct {
	IAuthCaller     // Read-only binding to the contract
	IAuthTransactor // Write-only binding to the contract
	IAuthFilterer   // Log filterer for contract events
}

// IAuthCaller is an auto generated read-only Go binding around an Ethereum contract.
type IAuthCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAuthTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IAuthTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAuthFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAuthFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAuthSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IAuthSession struct {
	Contract     *IAuth            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAuthCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IAuthCallerSession struct {
	Contract *IAuthCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IAuthTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IAuthTransactorSession struct {
	Contract     *IAuthTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAuthRaw is an auto generated low-level Go binding around an Ethereum contract.
type IAuthRaw struct {
	Contract *IAuth // Generic contract binding to access the raw methods on
}

// IAuthCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IAuthCallerRaw struct {
	Contract *IAuthCaller // Generic read-only contract binding to access the raw methods on
}

// IAuthTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IAuthTransactorRaw struct {
	Contract *IAuthTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIAuth creates a new instance of IAuth, bound to a specific deployed contract.
func NewIAuth(address common.Address, backend bind.ContractBackend) (*IAuth, error) {
	contract, err := bindIAuth(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAuth{IAuthCaller: IAuthCaller{contract: contract}, IAuthTransactor: IAuthTransactor{contract: contract}, IAuthFilterer: IAuthFilterer{contract: contract}}, nil
}

// NewIAuthCaller creates a new read-only instance of IAuth, bound to a specific deployed contract.
func NewIAuthCaller(address common.Address, caller bind.ContractCaller) (*IAuthCaller, error) {
	contract, err := bindIAuth(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAuthCaller{contract: contract}, nil
}

// NewIAuthTransactor creates a new write-only instance of IAuth, bound to a specific deployed contract.
func NewIAuthTransactor(address common.Address, transactor bind.ContractTransactor) (*IAuthTransactor, error) {
	contract, err := bindIAuth(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAuthTransactor{contract: contract}, nil
}

// NewIAuthFilterer creates a new log filterer instance of IAuth, bound to a specific deployed contract.
func NewIAuthFilterer(address common.Address, filterer bind.ContractFilterer) (*IAuthFilterer, error) {
	contract, err := bindIAuth(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAuthFilterer{contract: contract}, nil
}

// bindIAuth binds a generic wrapper to an already deployed contract.
func bindIAuth(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IAuthABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAuth *IAuthRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAuth.Contract.IAuthCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAuth *IAuthRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAuth.Contract.IAuthTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAuth *IAuthRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAuth.Contract.IAuthTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAuth *IAuthCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAuth.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAuth *IAuthTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAuth.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAuth *IAuthTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAuth.Contract.contract.Transact(opts, method, params...)
}

// Perm is a paid mutator transaction binding the contract method 0xa96bba9d.
//
// Solidity: function perm(bytes32 _h, bytes[5] signs) returns()
func (_IAuth *IAuthTransactor) Perm(opts *bind.TransactOpts, _h [32]byte, signs [5][]byte) (*types.Transaction, error) {
	return _IAuth.contract.Transact(opts, "perm", _h, signs)
}

// Perm is a paid mutator transaction binding the contract method 0xa96bba9d.
//
// Solidity: function perm(bytes32 _h, bytes[5] signs) returns()
func (_IAuth *IAuthSession) Perm(_h [32]byte, signs [5][]byte) (*types.Transaction, error) {
	return _IAuth.Contract.Perm(&_IAuth.TransactOpts, _h, signs)
}

// Perm is a paid mutator transaction binding the contract method 0xa96bba9d.
//
// Solidity: function perm(bytes32 _h, bytes[5] signs) returns()
func (_IAuth *IAuthTransactorSession) Perm(_h [32]byte, signs [5][]byte) (*types.Transaction, error) {
	return _IAuth.Contract.Perm(&_IAuth.TransactOpts, _h, signs)
}

// IKmanageABI is the input ABI used to generate the binding from.
const IKmanageABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"ki\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"cnt\",\"type\":\"uint64\"}],\"name\":\"AddCnt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"ti\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"}],\"name\":\"AddProfit\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_ki\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_cnt\",\"type\":\"uint64\"}],\"name\":\"addCnt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_ki\",\"type\":\"uint64\"}],\"name\":\"addKeeper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_money\",\"type\":\"uint256\"}],\"name\":\"addProfit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_ki\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"}],\"name\":\"getK\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getKCnt\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"}],\"name\":\"getPf\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"manageRate\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_ki\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IKmanageFuncSigs maps the 4-byte function signature to its string representation.
var IKmanageFuncSigs = map[string]string{
	"024130e4": "addCnt(uint64,uint64)",
	"50cbb46f": "addKeeper(uint64)",
	"55d3d7ef": "addProfit(uint8,uint256)",
	"fc3ba0ad": "balanceOf(uint64,uint8)",
	"dc0e7f45": "getK(uint64)",
	"39f2db96": "getKCnt()",
	"4f3c2eab": "getPf(uint8)",
	"6d23f6c8": "manageRate()",
	"259c6d5e": "withdraw(uint64,uint8,uint256)",
}

// IKmanage is an auto generated Go binding around an Ethereum contract.
type IKmanage struct {
	IKmanageCaller     // Read-only binding to the contract
	IKmanageTransactor // Write-only binding to the contract
	IKmanageFilterer   // Log filterer for contract events
}

// IKmanageCaller is an auto generated read-only Go binding around an Ethereum contract.
type IKmanageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IKmanageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IKmanageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IKmanageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IKmanageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IKmanageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IKmanageSession struct {
	Contract     *IKmanage         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IKmanageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IKmanageCallerSession struct {
	Contract *IKmanageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IKmanageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IKmanageTransactorSession struct {
	Contract     *IKmanageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IKmanageRaw is an auto generated low-level Go binding around an Ethereum contract.
type IKmanageRaw struct {
	Contract *IKmanage // Generic contract binding to access the raw methods on
}

// IKmanageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IKmanageCallerRaw struct {
	Contract *IKmanageCaller // Generic read-only contract binding to access the raw methods on
}

// IKmanageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IKmanageTransactorRaw struct {
	Contract *IKmanageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIKmanage creates a new instance of IKmanage, bound to a specific deployed contract.
func NewIKmanage(address common.Address, backend bind.ContractBackend) (*IKmanage, error) {
	contract, err := bindIKmanage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IKmanage{IKmanageCaller: IKmanageCaller{contract: contract}, IKmanageTransactor: IKmanageTransactor{contract: contract}, IKmanageFilterer: IKmanageFilterer{contract: contract}}, nil
}

// NewIKmanageCaller creates a new read-only instance of IKmanage, bound to a specific deployed contract.
func NewIKmanageCaller(address common.Address, caller bind.ContractCaller) (*IKmanageCaller, error) {
	contract, err := bindIKmanage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IKmanageCaller{contract: contract}, nil
}

// NewIKmanageTransactor creates a new write-only instance of IKmanage, bound to a specific deployed contract.
func NewIKmanageTransactor(address common.Address, transactor bind.ContractTransactor) (*IKmanageTransactor, error) {
	contract, err := bindIKmanage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IKmanageTransactor{contract: contract}, nil
}

// NewIKmanageFilterer creates a new log filterer instance of IKmanage, bound to a specific deployed contract.
func NewIKmanageFilterer(address common.Address, filterer bind.ContractFilterer) (*IKmanageFilterer, error) {
	contract, err := bindIKmanage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IKmanageFilterer{contract: contract}, nil
}

// bindIKmanage binds a generic wrapper to an already deployed contract.
func bindIKmanage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IKmanageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IKmanage *IKmanageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IKmanage.Contract.IKmanageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IKmanage *IKmanageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IKmanage.Contract.IKmanageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IKmanage *IKmanageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IKmanage.Contract.IKmanageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IKmanage *IKmanageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IKmanage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IKmanage *IKmanageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IKmanage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IKmanage *IKmanageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IKmanage.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0xfc3ba0ad.
//
// Solidity: function balanceOf(uint64 _ki, uint8 _ti) view returns(uint256, uint256)
func (_IKmanage *IKmanageCaller) BalanceOf(opts *bind.CallOpts, _ki uint64, _ti uint8) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _IKmanage.contract.Call(opts, &out, "balanceOf", _ki, _ti)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// BalanceOf is a free data retrieval call binding the contract method 0xfc3ba0ad.
//
// Solidity: function balanceOf(uint64 _ki, uint8 _ti) view returns(uint256, uint256)
func (_IKmanage *IKmanageSession) BalanceOf(_ki uint64, _ti uint8) (*big.Int, *big.Int, error) {
	return _IKmanage.Contract.BalanceOf(&_IKmanage.CallOpts, _ki, _ti)
}

// BalanceOf is a free data retrieval call binding the contract method 0xfc3ba0ad.
//
// Solidity: function balanceOf(uint64 _ki, uint8 _ti) view returns(uint256, uint256)
func (_IKmanage *IKmanageCallerSession) BalanceOf(_ki uint64, _ti uint8) (*big.Int, *big.Int, error) {
	return _IKmanage.Contract.BalanceOf(&_IKmanage.CallOpts, _ki, _ti)
}

// GetK is a free data retrieval call binding the contract method 0xdc0e7f45.
//
// Solidity: function getK(uint64 _i) view returns(uint64)
func (_IKmanage *IKmanageCaller) GetK(opts *bind.CallOpts, _i uint64) (uint64, error) {
	var out []interface{}
	err := _IKmanage.contract.Call(opts, &out, "getK", _i)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetK is a free data retrieval call binding the contract method 0xdc0e7f45.
//
// Solidity: function getK(uint64 _i) view returns(uint64)
func (_IKmanage *IKmanageSession) GetK(_i uint64) (uint64, error) {
	return _IKmanage.Contract.GetK(&_IKmanage.CallOpts, _i)
}

// GetK is a free data retrieval call binding the contract method 0xdc0e7f45.
//
// Solidity: function getK(uint64 _i) view returns(uint64)
func (_IKmanage *IKmanageCallerSession) GetK(_i uint64) (uint64, error) {
	return _IKmanage.Contract.GetK(&_IKmanage.CallOpts, _i)
}

// GetKCnt is a free data retrieval call binding the contract method 0x39f2db96.
//
// Solidity: function getKCnt() view returns(uint8)
func (_IKmanage *IKmanageCaller) GetKCnt(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _IKmanage.contract.Call(opts, &out, "getKCnt")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetKCnt is a free data retrieval call binding the contract method 0x39f2db96.
//
// Solidity: function getKCnt() view returns(uint8)
func (_IKmanage *IKmanageSession) GetKCnt() (uint8, error) {
	return _IKmanage.Contract.GetKCnt(&_IKmanage.CallOpts)
}

// GetKCnt is a free data retrieval call binding the contract method 0x39f2db96.
//
// Solidity: function getKCnt() view returns(uint8)
func (_IKmanage *IKmanageCallerSession) GetKCnt() (uint8, error) {
	return _IKmanage.Contract.GetKCnt(&_IKmanage.CallOpts)
}

// GetPf is a free data retrieval call binding the contract method 0x4f3c2eab.
//
// Solidity: function getPf(uint8 _ti) view returns(uint64, uint256)
func (_IKmanage *IKmanageCaller) GetPf(opts *bind.CallOpts, _ti uint8) (uint64, *big.Int, error) {
	var out []interface{}
	err := _IKmanage.contract.Call(opts, &out, "getPf", _ti)

	if err != nil {
		return *new(uint64), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetPf is a free data retrieval call binding the contract method 0x4f3c2eab.
//
// Solidity: function getPf(uint8 _ti) view returns(uint64, uint256)
func (_IKmanage *IKmanageSession) GetPf(_ti uint8) (uint64, *big.Int, error) {
	return _IKmanage.Contract.GetPf(&_IKmanage.CallOpts, _ti)
}

// GetPf is a free data retrieval call binding the contract method 0x4f3c2eab.
//
// Solidity: function getPf(uint8 _ti) view returns(uint64, uint256)
func (_IKmanage *IKmanageCallerSession) GetPf(_ti uint8) (uint64, *big.Int, error) {
	return _IKmanage.Contract.GetPf(&_IKmanage.CallOpts, _ti)
}

// ManageRate is a free data retrieval call binding the contract method 0x6d23f6c8.
//
// Solidity: function manageRate() view returns(uint8)
func (_IKmanage *IKmanageCaller) ManageRate(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _IKmanage.contract.Call(opts, &out, "manageRate")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ManageRate is a free data retrieval call binding the contract method 0x6d23f6c8.
//
// Solidity: function manageRate() view returns(uint8)
func (_IKmanage *IKmanageSession) ManageRate() (uint8, error) {
	return _IKmanage.Contract.ManageRate(&_IKmanage.CallOpts)
}

// ManageRate is a free data retrieval call binding the contract method 0x6d23f6c8.
//
// Solidity: function manageRate() view returns(uint8)
func (_IKmanage *IKmanageCallerSession) ManageRate() (uint8, error) {
	return _IKmanage.Contract.ManageRate(&_IKmanage.CallOpts)
}

// AddCnt is a paid mutator transaction binding the contract method 0x024130e4.
//
// Solidity: function addCnt(uint64 _ki, uint64 _cnt) returns()
func (_IKmanage *IKmanageTransactor) AddCnt(opts *bind.TransactOpts, _ki uint64, _cnt uint64) (*types.Transaction, error) {
	return _IKmanage.contract.Transact(opts, "addCnt", _ki, _cnt)
}

// AddCnt is a paid mutator transaction binding the contract method 0x024130e4.
//
// Solidity: function addCnt(uint64 _ki, uint64 _cnt) returns()
func (_IKmanage *IKmanageSession) AddCnt(_ki uint64, _cnt uint64) (*types.Transaction, error) {
	return _IKmanage.Contract.AddCnt(&_IKmanage.TransactOpts, _ki, _cnt)
}

// AddCnt is a paid mutator transaction binding the contract method 0x024130e4.
//
// Solidity: function addCnt(uint64 _ki, uint64 _cnt) returns()
func (_IKmanage *IKmanageTransactorSession) AddCnt(_ki uint64, _cnt uint64) (*types.Transaction, error) {
	return _IKmanage.Contract.AddCnt(&_IKmanage.TransactOpts, _ki, _cnt)
}

// AddKeeper is a paid mutator transaction binding the contract method 0x50cbb46f.
//
// Solidity: function addKeeper(uint64 _ki) returns()
func (_IKmanage *IKmanageTransactor) AddKeeper(opts *bind.TransactOpts, _ki uint64) (*types.Transaction, error) {
	return _IKmanage.contract.Transact(opts, "addKeeper", _ki)
}

// AddKeeper is a paid mutator transaction binding the contract method 0x50cbb46f.
//
// Solidity: function addKeeper(uint64 _ki) returns()
func (_IKmanage *IKmanageSession) AddKeeper(_ki uint64) (*types.Transaction, error) {
	return _IKmanage.Contract.AddKeeper(&_IKmanage.TransactOpts, _ki)
}

// AddKeeper is a paid mutator transaction binding the contract method 0x50cbb46f.
//
// Solidity: function addKeeper(uint64 _ki) returns()
func (_IKmanage *IKmanageTransactorSession) AddKeeper(_ki uint64) (*types.Transaction, error) {
	return _IKmanage.Contract.AddKeeper(&_IKmanage.TransactOpts, _ki)
}

// AddProfit is a paid mutator transaction binding the contract method 0x55d3d7ef.
//
// Solidity: function addProfit(uint8 _ti, uint256 _money) returns()
func (_IKmanage *IKmanageTransactor) AddProfit(opts *bind.TransactOpts, _ti uint8, _money *big.Int) (*types.Transaction, error) {
	return _IKmanage.contract.Transact(opts, "addProfit", _ti, _money)
}

// AddProfit is a paid mutator transaction binding the contract method 0x55d3d7ef.
//
// Solidity: function addProfit(uint8 _ti, uint256 _money) returns()
func (_IKmanage *IKmanageSession) AddProfit(_ti uint8, _money *big.Int) (*types.Transaction, error) {
	return _IKmanage.Contract.AddProfit(&_IKmanage.TransactOpts, _ti, _money)
}

// AddProfit is a paid mutator transaction binding the contract method 0x55d3d7ef.
//
// Solidity: function addProfit(uint8 _ti, uint256 _money) returns()
func (_IKmanage *IKmanageTransactorSession) AddProfit(_ti uint8, _money *big.Int) (*types.Transaction, error) {
	return _IKmanage.Contract.AddProfit(&_IKmanage.TransactOpts, _ti, _money)
}

// Withdraw is a paid mutator transaction binding the contract method 0x259c6d5e.
//
// Solidity: function withdraw(uint64 _ki, uint8 _ti, uint256 money) returns(uint256)
func (_IKmanage *IKmanageTransactor) Withdraw(opts *bind.TransactOpts, _ki uint64, _ti uint8, money *big.Int) (*types.Transaction, error) {
	return _IKmanage.contract.Transact(opts, "withdraw", _ki, _ti, money)
}

// Withdraw is a paid mutator transaction binding the contract method 0x259c6d5e.
//
// Solidity: function withdraw(uint64 _ki, uint8 _ti, uint256 money) returns(uint256)
func (_IKmanage *IKmanageSession) Withdraw(_ki uint64, _ti uint8, money *big.Int) (*types.Transaction, error) {
	return _IKmanage.Contract.Withdraw(&_IKmanage.TransactOpts, _ki, _ti, money)
}

// Withdraw is a paid mutator transaction binding the contract method 0x259c6d5e.
//
// Solidity: function withdraw(uint64 _ki, uint8 _ti, uint256 money) returns(uint256)
func (_IKmanage *IKmanageTransactorSession) Withdraw(_ki uint64, _ti uint8, money *big.Int) (*types.Transaction, error) {
	return _IKmanage.Contract.Withdraw(&_IKmanage.TransactOpts, _ki, _ti, money)
}

// IKmanageAddCntIterator is returned from FilterAddCnt and is used to iterate over the raw logs and unpacked data for AddCnt events raised by the IKmanage contract.
type IKmanageAddCntIterator struct {
	Event *IKmanageAddCnt // Event containing the contract specifics and raw log

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
func (it *IKmanageAddCntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IKmanageAddCnt)
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
		it.Event = new(IKmanageAddCnt)
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
func (it *IKmanageAddCntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IKmanageAddCntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IKmanageAddCnt represents a AddCnt event raised by the IKmanage contract.
type IKmanageAddCnt struct {
	Ki  uint64
	Cnt uint64
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAddCnt is a free log retrieval operation binding the contract event 0x5372d6aad551334f508508499c71755ebc6cde46b83fc1f944f1f0ae33cbb4c4.
//
// Solidity: event AddCnt(uint64 indexed ki, uint64 cnt)
func (_IKmanage *IKmanageFilterer) FilterAddCnt(opts *bind.FilterOpts, ki []uint64) (*IKmanageAddCntIterator, error) {

	var kiRule []interface{}
	for _, kiItem := range ki {
		kiRule = append(kiRule, kiItem)
	}

	logs, sub, err := _IKmanage.contract.FilterLogs(opts, "AddCnt", kiRule)
	if err != nil {
		return nil, err
	}
	return &IKmanageAddCntIterator{contract: _IKmanage.contract, event: "AddCnt", logs: logs, sub: sub}, nil
}

// WatchAddCnt is a free log subscription operation binding the contract event 0x5372d6aad551334f508508499c71755ebc6cde46b83fc1f944f1f0ae33cbb4c4.
//
// Solidity: event AddCnt(uint64 indexed ki, uint64 cnt)
func (_IKmanage *IKmanageFilterer) WatchAddCnt(opts *bind.WatchOpts, sink chan<- *IKmanageAddCnt, ki []uint64) (event.Subscription, error) {

	var kiRule []interface{}
	for _, kiItem := range ki {
		kiRule = append(kiRule, kiItem)
	}

	logs, sub, err := _IKmanage.contract.WatchLogs(opts, "AddCnt", kiRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IKmanageAddCnt)
				if err := _IKmanage.contract.UnpackLog(event, "AddCnt", log); err != nil {
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

// ParseAddCnt is a log parse operation binding the contract event 0x5372d6aad551334f508508499c71755ebc6cde46b83fc1f944f1f0ae33cbb4c4.
//
// Solidity: event AddCnt(uint64 indexed ki, uint64 cnt)
func (_IKmanage *IKmanageFilterer) ParseAddCnt(log types.Log) (*IKmanageAddCnt, error) {
	event := new(IKmanageAddCnt)
	if err := _IKmanage.contract.UnpackLog(event, "AddCnt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IKmanageAddProfitIterator is returned from FilterAddProfit and is used to iterate over the raw logs and unpacked data for AddProfit events raised by the IKmanage contract.
type IKmanageAddProfitIterator struct {
	Event *IKmanageAddProfit // Event containing the contract specifics and raw log

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
func (it *IKmanageAddProfitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IKmanageAddProfit)
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
		it.Event = new(IKmanageAddProfit)
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
func (it *IKmanageAddProfitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IKmanageAddProfitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IKmanageAddProfit represents a AddProfit event raised by the IKmanage contract.
type IKmanageAddProfit struct {
	Ti    uint8
	Money *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAddProfit is a free log retrieval operation binding the contract event 0xe5fd8ec20bdfeb1fadd32ec3786545a1db18d74f952effda5c6cd50af7c2e9e8.
//
// Solidity: event AddProfit(uint8 indexed ti, uint256 money)
func (_IKmanage *IKmanageFilterer) FilterAddProfit(opts *bind.FilterOpts, ti []uint8) (*IKmanageAddProfitIterator, error) {

	var tiRule []interface{}
	for _, tiItem := range ti {
		tiRule = append(tiRule, tiItem)
	}

	logs, sub, err := _IKmanage.contract.FilterLogs(opts, "AddProfit", tiRule)
	if err != nil {
		return nil, err
	}
	return &IKmanageAddProfitIterator{contract: _IKmanage.contract, event: "AddProfit", logs: logs, sub: sub}, nil
}

// WatchAddProfit is a free log subscription operation binding the contract event 0xe5fd8ec20bdfeb1fadd32ec3786545a1db18d74f952effda5c6cd50af7c2e9e8.
//
// Solidity: event AddProfit(uint8 indexed ti, uint256 money)
func (_IKmanage *IKmanageFilterer) WatchAddProfit(opts *bind.WatchOpts, sink chan<- *IKmanageAddProfit, ti []uint8) (event.Subscription, error) {

	var tiRule []interface{}
	for _, tiItem := range ti {
		tiRule = append(tiRule, tiItem)
	}

	logs, sub, err := _IKmanage.contract.WatchLogs(opts, "AddProfit", tiRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IKmanageAddProfit)
				if err := _IKmanage.contract.UnpackLog(event, "AddProfit", log); err != nil {
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

// ParseAddProfit is a log parse operation binding the contract event 0xe5fd8ec20bdfeb1fadd32ec3786545a1db18d74f952effda5c6cd50af7c2e9e8.
//
// Solidity: event AddProfit(uint8 indexed ti, uint256 money)
func (_IKmanage *IKmanageFilterer) ParseAddProfit(log types.Log) (*IKmanageAddProfit, error) {
	event := new(IKmanageAddProfit)
	if err := _IKmanage.contract.UnpackLog(event, "AddProfit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IKmanageGetterABI is the input ABI used to generate the binding from.
const IKmanageGetterABI = "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_ki\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"}],\"name\":\"getK\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getKCnt\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"}],\"name\":\"getPf\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"manageRate\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IKmanageGetterFuncSigs maps the 4-byte function signature to its string representation.
var IKmanageGetterFuncSigs = map[string]string{
	"fc3ba0ad": "balanceOf(uint64,uint8)",
	"dc0e7f45": "getK(uint64)",
	"39f2db96": "getKCnt()",
	"4f3c2eab": "getPf(uint8)",
	"6d23f6c8": "manageRate()",
}

// IKmanageGetter is an auto generated Go binding around an Ethereum contract.
type IKmanageGetter struct {
	IKmanageGetterCaller     // Read-only binding to the contract
	IKmanageGetterTransactor // Write-only binding to the contract
	IKmanageGetterFilterer   // Log filterer for contract events
}

// IKmanageGetterCaller is an auto generated read-only Go binding around an Ethereum contract.
type IKmanageGetterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IKmanageGetterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IKmanageGetterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IKmanageGetterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IKmanageGetterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IKmanageGetterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IKmanageGetterSession struct {
	Contract     *IKmanageGetter   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IKmanageGetterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IKmanageGetterCallerSession struct {
	Contract *IKmanageGetterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IKmanageGetterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IKmanageGetterTransactorSession struct {
	Contract     *IKmanageGetterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IKmanageGetterRaw is an auto generated low-level Go binding around an Ethereum contract.
type IKmanageGetterRaw struct {
	Contract *IKmanageGetter // Generic contract binding to access the raw methods on
}

// IKmanageGetterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IKmanageGetterCallerRaw struct {
	Contract *IKmanageGetterCaller // Generic read-only contract binding to access the raw methods on
}

// IKmanageGetterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IKmanageGetterTransactorRaw struct {
	Contract *IKmanageGetterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIKmanageGetter creates a new instance of IKmanageGetter, bound to a specific deployed contract.
func NewIKmanageGetter(address common.Address, backend bind.ContractBackend) (*IKmanageGetter, error) {
	contract, err := bindIKmanageGetter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IKmanageGetter{IKmanageGetterCaller: IKmanageGetterCaller{contract: contract}, IKmanageGetterTransactor: IKmanageGetterTransactor{contract: contract}, IKmanageGetterFilterer: IKmanageGetterFilterer{contract: contract}}, nil
}

// NewIKmanageGetterCaller creates a new read-only instance of IKmanageGetter, bound to a specific deployed contract.
func NewIKmanageGetterCaller(address common.Address, caller bind.ContractCaller) (*IKmanageGetterCaller, error) {
	contract, err := bindIKmanageGetter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IKmanageGetterCaller{contract: contract}, nil
}

// NewIKmanageGetterTransactor creates a new write-only instance of IKmanageGetter, bound to a specific deployed contract.
func NewIKmanageGetterTransactor(address common.Address, transactor bind.ContractTransactor) (*IKmanageGetterTransactor, error) {
	contract, err := bindIKmanageGetter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IKmanageGetterTransactor{contract: contract}, nil
}

// NewIKmanageGetterFilterer creates a new log filterer instance of IKmanageGetter, bound to a specific deployed contract.
func NewIKmanageGetterFilterer(address common.Address, filterer bind.ContractFilterer) (*IKmanageGetterFilterer, error) {
	contract, err := bindIKmanageGetter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IKmanageGetterFilterer{contract: contract}, nil
}

// bindIKmanageGetter binds a generic wrapper to an already deployed contract.
func bindIKmanageGetter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IKmanageGetterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IKmanageGetter *IKmanageGetterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IKmanageGetter.Contract.IKmanageGetterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IKmanageGetter *IKmanageGetterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IKmanageGetter.Contract.IKmanageGetterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IKmanageGetter *IKmanageGetterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IKmanageGetter.Contract.IKmanageGetterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IKmanageGetter *IKmanageGetterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IKmanageGetter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IKmanageGetter *IKmanageGetterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IKmanageGetter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IKmanageGetter *IKmanageGetterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IKmanageGetter.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0xfc3ba0ad.
//
// Solidity: function balanceOf(uint64 _ki, uint8 _ti) view returns(uint256, uint256)
func (_IKmanageGetter *IKmanageGetterCaller) BalanceOf(opts *bind.CallOpts, _ki uint64, _ti uint8) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _IKmanageGetter.contract.Call(opts, &out, "balanceOf", _ki, _ti)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// BalanceOf is a free data retrieval call binding the contract method 0xfc3ba0ad.
//
// Solidity: function balanceOf(uint64 _ki, uint8 _ti) view returns(uint256, uint256)
func (_IKmanageGetter *IKmanageGetterSession) BalanceOf(_ki uint64, _ti uint8) (*big.Int, *big.Int, error) {
	return _IKmanageGetter.Contract.BalanceOf(&_IKmanageGetter.CallOpts, _ki, _ti)
}

// BalanceOf is a free data retrieval call binding the contract method 0xfc3ba0ad.
//
// Solidity: function balanceOf(uint64 _ki, uint8 _ti) view returns(uint256, uint256)
func (_IKmanageGetter *IKmanageGetterCallerSession) BalanceOf(_ki uint64, _ti uint8) (*big.Int, *big.Int, error) {
	return _IKmanageGetter.Contract.BalanceOf(&_IKmanageGetter.CallOpts, _ki, _ti)
}

// GetK is a free data retrieval call binding the contract method 0xdc0e7f45.
//
// Solidity: function getK(uint64 _i) view returns(uint64)
func (_IKmanageGetter *IKmanageGetterCaller) GetK(opts *bind.CallOpts, _i uint64) (uint64, error) {
	var out []interface{}
	err := _IKmanageGetter.contract.Call(opts, &out, "getK", _i)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetK is a free data retrieval call binding the contract method 0xdc0e7f45.
//
// Solidity: function getK(uint64 _i) view returns(uint64)
func (_IKmanageGetter *IKmanageGetterSession) GetK(_i uint64) (uint64, error) {
	return _IKmanageGetter.Contract.GetK(&_IKmanageGetter.CallOpts, _i)
}

// GetK is a free data retrieval call binding the contract method 0xdc0e7f45.
//
// Solidity: function getK(uint64 _i) view returns(uint64)
func (_IKmanageGetter *IKmanageGetterCallerSession) GetK(_i uint64) (uint64, error) {
	return _IKmanageGetter.Contract.GetK(&_IKmanageGetter.CallOpts, _i)
}

// GetKCnt is a free data retrieval call binding the contract method 0x39f2db96.
//
// Solidity: function getKCnt() view returns(uint8)
func (_IKmanageGetter *IKmanageGetterCaller) GetKCnt(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _IKmanageGetter.contract.Call(opts, &out, "getKCnt")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetKCnt is a free data retrieval call binding the contract method 0x39f2db96.
//
// Solidity: function getKCnt() view returns(uint8)
func (_IKmanageGetter *IKmanageGetterSession) GetKCnt() (uint8, error) {
	return _IKmanageGetter.Contract.GetKCnt(&_IKmanageGetter.CallOpts)
}

// GetKCnt is a free data retrieval call binding the contract method 0x39f2db96.
//
// Solidity: function getKCnt() view returns(uint8)
func (_IKmanageGetter *IKmanageGetterCallerSession) GetKCnt() (uint8, error) {
	return _IKmanageGetter.Contract.GetKCnt(&_IKmanageGetter.CallOpts)
}

// GetPf is a free data retrieval call binding the contract method 0x4f3c2eab.
//
// Solidity: function getPf(uint8 _ti) view returns(uint64, uint256)
func (_IKmanageGetter *IKmanageGetterCaller) GetPf(opts *bind.CallOpts, _ti uint8) (uint64, *big.Int, error) {
	var out []interface{}
	err := _IKmanageGetter.contract.Call(opts, &out, "getPf", _ti)

	if err != nil {
		return *new(uint64), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetPf is a free data retrieval call binding the contract method 0x4f3c2eab.
//
// Solidity: function getPf(uint8 _ti) view returns(uint64, uint256)
func (_IKmanageGetter *IKmanageGetterSession) GetPf(_ti uint8) (uint64, *big.Int, error) {
	return _IKmanageGetter.Contract.GetPf(&_IKmanageGetter.CallOpts, _ti)
}

// GetPf is a free data retrieval call binding the contract method 0x4f3c2eab.
//
// Solidity: function getPf(uint8 _ti) view returns(uint64, uint256)
func (_IKmanageGetter *IKmanageGetterCallerSession) GetPf(_ti uint8) (uint64, *big.Int, error) {
	return _IKmanageGetter.Contract.GetPf(&_IKmanageGetter.CallOpts, _ti)
}

// ManageRate is a free data retrieval call binding the contract method 0x6d23f6c8.
//
// Solidity: function manageRate() view returns(uint8)
func (_IKmanageGetter *IKmanageGetterCaller) ManageRate(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _IKmanageGetter.contract.Call(opts, &out, "manageRate")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ManageRate is a free data retrieval call binding the contract method 0x6d23f6c8.
//
// Solidity: function manageRate() view returns(uint8)
func (_IKmanageGetter *IKmanageGetterSession) ManageRate() (uint8, error) {
	return _IKmanageGetter.Contract.ManageRate(&_IKmanageGetter.CallOpts)
}

// ManageRate is a free data retrieval call binding the contract method 0x6d23f6c8.
//
// Solidity: function manageRate() view returns(uint8)
func (_IKmanageGetter *IKmanageGetterCallerSession) ManageRate() (uint8, error) {
	return _IKmanageGetter.Contract.ManageRate(&_IKmanageGetter.CallOpts)
}

// IKmanageSetterABI is the input ABI used to generate the binding from.
const IKmanageSetterABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"ki\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"cnt\",\"type\":\"uint64\"}],\"name\":\"AddCnt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"ti\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"}],\"name\":\"AddProfit\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_ki\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_cnt\",\"type\":\"uint64\"}],\"name\":\"addCnt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_ki\",\"type\":\"uint64\"}],\"name\":\"addKeeper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_money\",\"type\":\"uint256\"}],\"name\":\"addProfit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_ki\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IKmanageSetterFuncSigs maps the 4-byte function signature to its string representation.
var IKmanageSetterFuncSigs = map[string]string{
	"024130e4": "addCnt(uint64,uint64)",
	"50cbb46f": "addKeeper(uint64)",
	"55d3d7ef": "addProfit(uint8,uint256)",
	"259c6d5e": "withdraw(uint64,uint8,uint256)",
}

// IKmanageSetter is an auto generated Go binding around an Ethereum contract.
type IKmanageSetter struct {
	IKmanageSetterCaller     // Read-only binding to the contract
	IKmanageSetterTransactor // Write-only binding to the contract
	IKmanageSetterFilterer   // Log filterer for contract events
}

// IKmanageSetterCaller is an auto generated read-only Go binding around an Ethereum contract.
type IKmanageSetterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IKmanageSetterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IKmanageSetterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IKmanageSetterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IKmanageSetterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IKmanageSetterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IKmanageSetterSession struct {
	Contract     *IKmanageSetter   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IKmanageSetterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IKmanageSetterCallerSession struct {
	Contract *IKmanageSetterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IKmanageSetterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IKmanageSetterTransactorSession struct {
	Contract     *IKmanageSetterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IKmanageSetterRaw is an auto generated low-level Go binding around an Ethereum contract.
type IKmanageSetterRaw struct {
	Contract *IKmanageSetter // Generic contract binding to access the raw methods on
}

// IKmanageSetterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IKmanageSetterCallerRaw struct {
	Contract *IKmanageSetterCaller // Generic read-only contract binding to access the raw methods on
}

// IKmanageSetterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IKmanageSetterTransactorRaw struct {
	Contract *IKmanageSetterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIKmanageSetter creates a new instance of IKmanageSetter, bound to a specific deployed contract.
func NewIKmanageSetter(address common.Address, backend bind.ContractBackend) (*IKmanageSetter, error) {
	contract, err := bindIKmanageSetter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IKmanageSetter{IKmanageSetterCaller: IKmanageSetterCaller{contract: contract}, IKmanageSetterTransactor: IKmanageSetterTransactor{contract: contract}, IKmanageSetterFilterer: IKmanageSetterFilterer{contract: contract}}, nil
}

// NewIKmanageSetterCaller creates a new read-only instance of IKmanageSetter, bound to a specific deployed contract.
func NewIKmanageSetterCaller(address common.Address, caller bind.ContractCaller) (*IKmanageSetterCaller, error) {
	contract, err := bindIKmanageSetter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IKmanageSetterCaller{contract: contract}, nil
}

// NewIKmanageSetterTransactor creates a new write-only instance of IKmanageSetter, bound to a specific deployed contract.
func NewIKmanageSetterTransactor(address common.Address, transactor bind.ContractTransactor) (*IKmanageSetterTransactor, error) {
	contract, err := bindIKmanageSetter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IKmanageSetterTransactor{contract: contract}, nil
}

// NewIKmanageSetterFilterer creates a new log filterer instance of IKmanageSetter, bound to a specific deployed contract.
func NewIKmanageSetterFilterer(address common.Address, filterer bind.ContractFilterer) (*IKmanageSetterFilterer, error) {
	contract, err := bindIKmanageSetter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IKmanageSetterFilterer{contract: contract}, nil
}

// bindIKmanageSetter binds a generic wrapper to an already deployed contract.
func bindIKmanageSetter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IKmanageSetterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IKmanageSetter *IKmanageSetterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IKmanageSetter.Contract.IKmanageSetterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IKmanageSetter *IKmanageSetterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IKmanageSetter.Contract.IKmanageSetterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IKmanageSetter *IKmanageSetterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IKmanageSetter.Contract.IKmanageSetterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IKmanageSetter *IKmanageSetterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IKmanageSetter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IKmanageSetter *IKmanageSetterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IKmanageSetter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IKmanageSetter *IKmanageSetterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IKmanageSetter.Contract.contract.Transact(opts, method, params...)
}

// AddCnt is a paid mutator transaction binding the contract method 0x024130e4.
//
// Solidity: function addCnt(uint64 _ki, uint64 _cnt) returns()
func (_IKmanageSetter *IKmanageSetterTransactor) AddCnt(opts *bind.TransactOpts, _ki uint64, _cnt uint64) (*types.Transaction, error) {
	return _IKmanageSetter.contract.Transact(opts, "addCnt", _ki, _cnt)
}

// AddCnt is a paid mutator transaction binding the contract method 0x024130e4.
//
// Solidity: function addCnt(uint64 _ki, uint64 _cnt) returns()
func (_IKmanageSetter *IKmanageSetterSession) AddCnt(_ki uint64, _cnt uint64) (*types.Transaction, error) {
	return _IKmanageSetter.Contract.AddCnt(&_IKmanageSetter.TransactOpts, _ki, _cnt)
}

// AddCnt is a paid mutator transaction binding the contract method 0x024130e4.
//
// Solidity: function addCnt(uint64 _ki, uint64 _cnt) returns()
func (_IKmanageSetter *IKmanageSetterTransactorSession) AddCnt(_ki uint64, _cnt uint64) (*types.Transaction, error) {
	return _IKmanageSetter.Contract.AddCnt(&_IKmanageSetter.TransactOpts, _ki, _cnt)
}

// AddKeeper is a paid mutator transaction binding the contract method 0x50cbb46f.
//
// Solidity: function addKeeper(uint64 _ki) returns()
func (_IKmanageSetter *IKmanageSetterTransactor) AddKeeper(opts *bind.TransactOpts, _ki uint64) (*types.Transaction, error) {
	return _IKmanageSetter.contract.Transact(opts, "addKeeper", _ki)
}

// AddKeeper is a paid mutator transaction binding the contract method 0x50cbb46f.
//
// Solidity: function addKeeper(uint64 _ki) returns()
func (_IKmanageSetter *IKmanageSetterSession) AddKeeper(_ki uint64) (*types.Transaction, error) {
	return _IKmanageSetter.Contract.AddKeeper(&_IKmanageSetter.TransactOpts, _ki)
}

// AddKeeper is a paid mutator transaction binding the contract method 0x50cbb46f.
//
// Solidity: function addKeeper(uint64 _ki) returns()
func (_IKmanageSetter *IKmanageSetterTransactorSession) AddKeeper(_ki uint64) (*types.Transaction, error) {
	return _IKmanageSetter.Contract.AddKeeper(&_IKmanageSetter.TransactOpts, _ki)
}

// AddProfit is a paid mutator transaction binding the contract method 0x55d3d7ef.
//
// Solidity: function addProfit(uint8 _ti, uint256 _money) returns()
func (_IKmanageSetter *IKmanageSetterTransactor) AddProfit(opts *bind.TransactOpts, _ti uint8, _money *big.Int) (*types.Transaction, error) {
	return _IKmanageSetter.contract.Transact(opts, "addProfit", _ti, _money)
}

// AddProfit is a paid mutator transaction binding the contract method 0x55d3d7ef.
//
// Solidity: function addProfit(uint8 _ti, uint256 _money) returns()
func (_IKmanageSetter *IKmanageSetterSession) AddProfit(_ti uint8, _money *big.Int) (*types.Transaction, error) {
	return _IKmanageSetter.Contract.AddProfit(&_IKmanageSetter.TransactOpts, _ti, _money)
}

// AddProfit is a paid mutator transaction binding the contract method 0x55d3d7ef.
//
// Solidity: function addProfit(uint8 _ti, uint256 _money) returns()
func (_IKmanageSetter *IKmanageSetterTransactorSession) AddProfit(_ti uint8, _money *big.Int) (*types.Transaction, error) {
	return _IKmanageSetter.Contract.AddProfit(&_IKmanageSetter.TransactOpts, _ti, _money)
}

// Withdraw is a paid mutator transaction binding the contract method 0x259c6d5e.
//
// Solidity: function withdraw(uint64 _ki, uint8 _ti, uint256 money) returns(uint256)
func (_IKmanageSetter *IKmanageSetterTransactor) Withdraw(opts *bind.TransactOpts, _ki uint64, _ti uint8, money *big.Int) (*types.Transaction, error) {
	return _IKmanageSetter.contract.Transact(opts, "withdraw", _ki, _ti, money)
}

// Withdraw is a paid mutator transaction binding the contract method 0x259c6d5e.
//
// Solidity: function withdraw(uint64 _ki, uint8 _ti, uint256 money) returns(uint256)
func (_IKmanageSetter *IKmanageSetterSession) Withdraw(_ki uint64, _ti uint8, money *big.Int) (*types.Transaction, error) {
	return _IKmanageSetter.Contract.Withdraw(&_IKmanageSetter.TransactOpts, _ki, _ti, money)
}

// Withdraw is a paid mutator transaction binding the contract method 0x259c6d5e.
//
// Solidity: function withdraw(uint64 _ki, uint8 _ti, uint256 money) returns(uint256)
func (_IKmanageSetter *IKmanageSetterTransactorSession) Withdraw(_ki uint64, _ti uint8, money *big.Int) (*types.Transaction, error) {
	return _IKmanageSetter.Contract.Withdraw(&_IKmanageSetter.TransactOpts, _ki, _ti, money)
}

// IKmanageSetterAddCntIterator is returned from FilterAddCnt and is used to iterate over the raw logs and unpacked data for AddCnt events raised by the IKmanageSetter contract.
type IKmanageSetterAddCntIterator struct {
	Event *IKmanageSetterAddCnt // Event containing the contract specifics and raw log

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
func (it *IKmanageSetterAddCntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IKmanageSetterAddCnt)
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
		it.Event = new(IKmanageSetterAddCnt)
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
func (it *IKmanageSetterAddCntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IKmanageSetterAddCntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IKmanageSetterAddCnt represents a AddCnt event raised by the IKmanageSetter contract.
type IKmanageSetterAddCnt struct {
	Ki  uint64
	Cnt uint64
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAddCnt is a free log retrieval operation binding the contract event 0x5372d6aad551334f508508499c71755ebc6cde46b83fc1f944f1f0ae33cbb4c4.
//
// Solidity: event AddCnt(uint64 indexed ki, uint64 cnt)
func (_IKmanageSetter *IKmanageSetterFilterer) FilterAddCnt(opts *bind.FilterOpts, ki []uint64) (*IKmanageSetterAddCntIterator, error) {

	var kiRule []interface{}
	for _, kiItem := range ki {
		kiRule = append(kiRule, kiItem)
	}

	logs, sub, err := _IKmanageSetter.contract.FilterLogs(opts, "AddCnt", kiRule)
	if err != nil {
		return nil, err
	}
	return &IKmanageSetterAddCntIterator{contract: _IKmanageSetter.contract, event: "AddCnt", logs: logs, sub: sub}, nil
}

// WatchAddCnt is a free log subscription operation binding the contract event 0x5372d6aad551334f508508499c71755ebc6cde46b83fc1f944f1f0ae33cbb4c4.
//
// Solidity: event AddCnt(uint64 indexed ki, uint64 cnt)
func (_IKmanageSetter *IKmanageSetterFilterer) WatchAddCnt(opts *bind.WatchOpts, sink chan<- *IKmanageSetterAddCnt, ki []uint64) (event.Subscription, error) {

	var kiRule []interface{}
	for _, kiItem := range ki {
		kiRule = append(kiRule, kiItem)
	}

	logs, sub, err := _IKmanageSetter.contract.WatchLogs(opts, "AddCnt", kiRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IKmanageSetterAddCnt)
				if err := _IKmanageSetter.contract.UnpackLog(event, "AddCnt", log); err != nil {
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

// ParseAddCnt is a log parse operation binding the contract event 0x5372d6aad551334f508508499c71755ebc6cde46b83fc1f944f1f0ae33cbb4c4.
//
// Solidity: event AddCnt(uint64 indexed ki, uint64 cnt)
func (_IKmanageSetter *IKmanageSetterFilterer) ParseAddCnt(log types.Log) (*IKmanageSetterAddCnt, error) {
	event := new(IKmanageSetterAddCnt)
	if err := _IKmanageSetter.contract.UnpackLog(event, "AddCnt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IKmanageSetterAddProfitIterator is returned from FilterAddProfit and is used to iterate over the raw logs and unpacked data for AddProfit events raised by the IKmanageSetter contract.
type IKmanageSetterAddProfitIterator struct {
	Event *IKmanageSetterAddProfit // Event containing the contract specifics and raw log

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
func (it *IKmanageSetterAddProfitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IKmanageSetterAddProfit)
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
		it.Event = new(IKmanageSetterAddProfit)
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
func (it *IKmanageSetterAddProfitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IKmanageSetterAddProfitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IKmanageSetterAddProfit represents a AddProfit event raised by the IKmanageSetter contract.
type IKmanageSetterAddProfit struct {
	Ti    uint8
	Money *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAddProfit is a free log retrieval operation binding the contract event 0xe5fd8ec20bdfeb1fadd32ec3786545a1db18d74f952effda5c6cd50af7c2e9e8.
//
// Solidity: event AddProfit(uint8 indexed ti, uint256 money)
func (_IKmanageSetter *IKmanageSetterFilterer) FilterAddProfit(opts *bind.FilterOpts, ti []uint8) (*IKmanageSetterAddProfitIterator, error) {

	var tiRule []interface{}
	for _, tiItem := range ti {
		tiRule = append(tiRule, tiItem)
	}

	logs, sub, err := _IKmanageSetter.contract.FilterLogs(opts, "AddProfit", tiRule)
	if err != nil {
		return nil, err
	}
	return &IKmanageSetterAddProfitIterator{contract: _IKmanageSetter.contract, event: "AddProfit", logs: logs, sub: sub}, nil
}

// WatchAddProfit is a free log subscription operation binding the contract event 0xe5fd8ec20bdfeb1fadd32ec3786545a1db18d74f952effda5c6cd50af7c2e9e8.
//
// Solidity: event AddProfit(uint8 indexed ti, uint256 money)
func (_IKmanageSetter *IKmanageSetterFilterer) WatchAddProfit(opts *bind.WatchOpts, sink chan<- *IKmanageSetterAddProfit, ti []uint8) (event.Subscription, error) {

	var tiRule []interface{}
	for _, tiItem := range ti {
		tiRule = append(tiRule, tiItem)
	}

	logs, sub, err := _IKmanageSetter.contract.WatchLogs(opts, "AddProfit", tiRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IKmanageSetterAddProfit)
				if err := _IKmanageSetter.contract.UnpackLog(event, "AddProfit", log); err != nil {
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

// ParseAddProfit is a log parse operation binding the contract event 0xe5fd8ec20bdfeb1fadd32ec3786545a1db18d74f952effda5c6cd50af7c2e9e8.
//
// Solidity: event AddProfit(uint8 indexed ti, uint256 money)
func (_IKmanageSetter *IKmanageSetterFilterer) ParseAddProfit(log types.Log) (*IKmanageSetterAddProfit, error) {
	event := new(IKmanageSetterAddProfit)
	if err := _IKmanageSetter.contract.UnpackLog(event, "AddProfit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOwnerABI is the input ABI used to generate the binding from.
const IOwnerABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isSet\",\"type\":\"bool\"}],\"name\":\"AddOwner\",\"type\":\"event\"}]"

// IOwner is an auto generated Go binding around an Ethereum contract.
type IOwner struct {
	IOwnerCaller     // Read-only binding to the contract
	IOwnerTransactor // Write-only binding to the contract
	IOwnerFilterer   // Log filterer for contract events
}

// IOwnerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IOwnerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOwnerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IOwnerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOwnerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IOwnerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOwnerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IOwnerSession struct {
	Contract     *IOwner           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IOwnerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IOwnerCallerSession struct {
	Contract *IOwnerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IOwnerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IOwnerTransactorSession struct {
	Contract     *IOwnerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IOwnerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IOwnerRaw struct {
	Contract *IOwner // Generic contract binding to access the raw methods on
}

// IOwnerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IOwnerCallerRaw struct {
	Contract *IOwnerCaller // Generic read-only contract binding to access the raw methods on
}

// IOwnerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IOwnerTransactorRaw struct {
	Contract *IOwnerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIOwner creates a new instance of IOwner, bound to a specific deployed contract.
func NewIOwner(address common.Address, backend bind.ContractBackend) (*IOwner, error) {
	contract, err := bindIOwner(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IOwner{IOwnerCaller: IOwnerCaller{contract: contract}, IOwnerTransactor: IOwnerTransactor{contract: contract}, IOwnerFilterer: IOwnerFilterer{contract: contract}}, nil
}

// NewIOwnerCaller creates a new read-only instance of IOwner, bound to a specific deployed contract.
func NewIOwnerCaller(address common.Address, caller bind.ContractCaller) (*IOwnerCaller, error) {
	contract, err := bindIOwner(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IOwnerCaller{contract: contract}, nil
}

// NewIOwnerTransactor creates a new write-only instance of IOwner, bound to a specific deployed contract.
func NewIOwnerTransactor(address common.Address, transactor bind.ContractTransactor) (*IOwnerTransactor, error) {
	contract, err := bindIOwner(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IOwnerTransactor{contract: contract}, nil
}

// NewIOwnerFilterer creates a new log filterer instance of IOwner, bound to a specific deployed contract.
func NewIOwnerFilterer(address common.Address, filterer bind.ContractFilterer) (*IOwnerFilterer, error) {
	contract, err := bindIOwner(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IOwnerFilterer{contract: contract}, nil
}

// bindIOwner binds a generic wrapper to an already deployed contract.
func bindIOwner(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IOwnerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOwner *IOwnerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOwner.Contract.IOwnerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOwner *IOwnerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOwner.Contract.IOwnerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOwner *IOwnerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOwner.Contract.IOwnerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOwner *IOwnerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOwner.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOwner *IOwnerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOwner.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOwner *IOwnerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOwner.Contract.contract.Transact(opts, method, params...)
}

// IOwnerAddOwnerIterator is returned from FilterAddOwner and is used to iterate over the raw logs and unpacked data for AddOwner events raised by the IOwner contract.
type IOwnerAddOwnerIterator struct {
	Event *IOwnerAddOwner // Event containing the contract specifics and raw log

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
func (it *IOwnerAddOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOwnerAddOwner)
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
		it.Event = new(IOwnerAddOwner)
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
func (it *IOwnerAddOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOwnerAddOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOwnerAddOwner represents a AddOwner event raised by the IOwner contract.
type IOwnerAddOwner struct {
	From  common.Address
	IsSet bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAddOwner is a free log retrieval operation binding the contract event 0x938b2a24c98e4e542127ffa74a91e48942c2bddccae3b6d75f82bfda7bbc0807.
//
// Solidity: event AddOwner(address from, bool isSet)
func (_IOwner *IOwnerFilterer) FilterAddOwner(opts *bind.FilterOpts) (*IOwnerAddOwnerIterator, error) {

	logs, sub, err := _IOwner.contract.FilterLogs(opts, "AddOwner")
	if err != nil {
		return nil, err
	}
	return &IOwnerAddOwnerIterator{contract: _IOwner.contract, event: "AddOwner", logs: logs, sub: sub}, nil
}

// WatchAddOwner is a free log subscription operation binding the contract event 0x938b2a24c98e4e542127ffa74a91e48942c2bddccae3b6d75f82bfda7bbc0807.
//
// Solidity: event AddOwner(address from, bool isSet)
func (_IOwner *IOwnerFilterer) WatchAddOwner(opts *bind.WatchOpts, sink chan<- *IOwnerAddOwner) (event.Subscription, error) {

	logs, sub, err := _IOwner.contract.WatchLogs(opts, "AddOwner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOwnerAddOwner)
				if err := _IOwner.contract.UnpackLog(event, "AddOwner", log); err != nil {
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

// ParseAddOwner is a log parse operation binding the contract event 0x938b2a24c98e4e542127ffa74a91e48942c2bddccae3b6d75f82bfda7bbc0807.
//
// Solidity: event AddOwner(address from, bool isSet)
func (_IOwner *IOwnerFilterer) ParseAddOwner(log types.Log) (*IOwnerAddOwner, error) {
	event := new(IOwnerAddOwner)
	if err := _IOwner.contract.UnpackLog(event, "AddOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPoolABI is the input ABI used to generate the binding from.
const IPoolABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"}],\"name\":\"Inflow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"}],\"name\":\"Outflow\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"}],\"name\":\"inflow\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"}],\"name\":\"outflow\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// IPoolFuncSigs maps the 4-byte function signature to its string representation.
var IPoolFuncSigs = map[string]string{
	"368007fe": "inflow(address,address,uint256)",
	"530345bb": "outflow(address,address,uint256)",
}

// IPool is an auto generated Go binding around an Ethereum contract.
type IPool struct {
	IPoolCaller     // Read-only binding to the contract
	IPoolTransactor // Write-only binding to the contract
	IPoolFilterer   // Log filterer for contract events
}

// IPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type IPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPoolSession struct {
	Contract     *IPool            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPoolCallerSession struct {
	Contract *IPoolCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPoolTransactorSession struct {
	Contract     *IPoolTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type IPoolRaw struct {
	Contract *IPool // Generic contract binding to access the raw methods on
}

// IPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPoolCallerRaw struct {
	Contract *IPoolCaller // Generic read-only contract binding to access the raw methods on
}

// IPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPoolTransactorRaw struct {
	Contract *IPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIPool creates a new instance of IPool, bound to a specific deployed contract.
func NewIPool(address common.Address, backend bind.ContractBackend) (*IPool, error) {
	contract, err := bindIPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPool{IPoolCaller: IPoolCaller{contract: contract}, IPoolTransactor: IPoolTransactor{contract: contract}, IPoolFilterer: IPoolFilterer{contract: contract}}, nil
}

// NewIPoolCaller creates a new read-only instance of IPool, bound to a specific deployed contract.
func NewIPoolCaller(address common.Address, caller bind.ContractCaller) (*IPoolCaller, error) {
	contract, err := bindIPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPoolCaller{contract: contract}, nil
}

// NewIPoolTransactor creates a new write-only instance of IPool, bound to a specific deployed contract.
func NewIPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*IPoolTransactor, error) {
	contract, err := bindIPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPoolTransactor{contract: contract}, nil
}

// NewIPoolFilterer creates a new log filterer instance of IPool, bound to a specific deployed contract.
func NewIPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*IPoolFilterer, error) {
	contract, err := bindIPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPoolFilterer{contract: contract}, nil
}

// bindIPool binds a generic wrapper to an already deployed contract.
func bindIPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IPoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPool *IPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPool.Contract.IPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPool *IPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPool.Contract.IPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPool *IPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPool.Contract.IPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPool *IPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPool *IPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPool *IPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPool.Contract.contract.Transact(opts, method, params...)
}

// Inflow is a paid mutator transaction binding the contract method 0x368007fe.
//
// Solidity: function inflow(address tAddr, address from, uint256 money) payable returns()
func (_IPool *IPoolTransactor) Inflow(opts *bind.TransactOpts, tAddr common.Address, from common.Address, money *big.Int) (*types.Transaction, error) {
	return _IPool.contract.Transact(opts, "inflow", tAddr, from, money)
}

// Inflow is a paid mutator transaction binding the contract method 0x368007fe.
//
// Solidity: function inflow(address tAddr, address from, uint256 money) payable returns()
func (_IPool *IPoolSession) Inflow(tAddr common.Address, from common.Address, money *big.Int) (*types.Transaction, error) {
	return _IPool.Contract.Inflow(&_IPool.TransactOpts, tAddr, from, money)
}

// Inflow is a paid mutator transaction binding the contract method 0x368007fe.
//
// Solidity: function inflow(address tAddr, address from, uint256 money) payable returns()
func (_IPool *IPoolTransactorSession) Inflow(tAddr common.Address, from common.Address, money *big.Int) (*types.Transaction, error) {
	return _IPool.Contract.Inflow(&_IPool.TransactOpts, tAddr, from, money)
}

// Outflow is a paid mutator transaction binding the contract method 0x530345bb.
//
// Solidity: function outflow(address tAddr, address to, uint256 money) payable returns()
func (_IPool *IPoolTransactor) Outflow(opts *bind.TransactOpts, tAddr common.Address, to common.Address, money *big.Int) (*types.Transaction, error) {
	return _IPool.contract.Transact(opts, "outflow", tAddr, to, money)
}

// Outflow is a paid mutator transaction binding the contract method 0x530345bb.
//
// Solidity: function outflow(address tAddr, address to, uint256 money) payable returns()
func (_IPool *IPoolSession) Outflow(tAddr common.Address, to common.Address, money *big.Int) (*types.Transaction, error) {
	return _IPool.Contract.Outflow(&_IPool.TransactOpts, tAddr, to, money)
}

// Outflow is a paid mutator transaction binding the contract method 0x530345bb.
//
// Solidity: function outflow(address tAddr, address to, uint256 money) payable returns()
func (_IPool *IPoolTransactorSession) Outflow(tAddr common.Address, to common.Address, money *big.Int) (*types.Transaction, error) {
	return _IPool.Contract.Outflow(&_IPool.TransactOpts, tAddr, to, money)
}

// IPoolInflowIterator is returned from FilterInflow and is used to iterate over the raw logs and unpacked data for Inflow events raised by the IPool contract.
type IPoolInflowIterator struct {
	Event *IPoolInflow // Event containing the contract specifics and raw log

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
func (it *IPoolInflowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPoolInflow)
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
		it.Event = new(IPoolInflow)
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
func (it *IPoolInflowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPoolInflowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPoolInflow represents a Inflow event raised by the IPool contract.
type IPoolInflow struct {
	From  common.Address
	TAddr common.Address
	Money *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterInflow is a free log retrieval operation binding the contract event 0xde52ea03a3979fafeaf8ea9d7fe6b3ddc6a95e9e8c0922562db3a047c0d72578.
//
// Solidity: event Inflow(address indexed from, address tAddr, uint256 money)
func (_IPool *IPoolFilterer) FilterInflow(opts *bind.FilterOpts, from []common.Address) (*IPoolInflowIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _IPool.contract.FilterLogs(opts, "Inflow", fromRule)
	if err != nil {
		return nil, err
	}
	return &IPoolInflowIterator{contract: _IPool.contract, event: "Inflow", logs: logs, sub: sub}, nil
}

// WatchInflow is a free log subscription operation binding the contract event 0xde52ea03a3979fafeaf8ea9d7fe6b3ddc6a95e9e8c0922562db3a047c0d72578.
//
// Solidity: event Inflow(address indexed from, address tAddr, uint256 money)
func (_IPool *IPoolFilterer) WatchInflow(opts *bind.WatchOpts, sink chan<- *IPoolInflow, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _IPool.contract.WatchLogs(opts, "Inflow", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPoolInflow)
				if err := _IPool.contract.UnpackLog(event, "Inflow", log); err != nil {
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

// ParseInflow is a log parse operation binding the contract event 0xde52ea03a3979fafeaf8ea9d7fe6b3ddc6a95e9e8c0922562db3a047c0d72578.
//
// Solidity: event Inflow(address indexed from, address tAddr, uint256 money)
func (_IPool *IPoolFilterer) ParseInflow(log types.Log) (*IPoolInflow, error) {
	event := new(IPoolInflow)
	if err := _IPool.contract.UnpackLog(event, "Inflow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPoolOutflowIterator is returned from FilterOutflow and is used to iterate over the raw logs and unpacked data for Outflow events raised by the IPool contract.
type IPoolOutflowIterator struct {
	Event *IPoolOutflow // Event containing the contract specifics and raw log

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
func (it *IPoolOutflowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPoolOutflow)
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
		it.Event = new(IPoolOutflow)
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
func (it *IPoolOutflowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPoolOutflowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPoolOutflow represents a Outflow event raised by the IPool contract.
type IPoolOutflow struct {
	To    common.Address
	TAddr common.Address
	Money *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterOutflow is a free log retrieval operation binding the contract event 0x4eb2adb2eba0bb5546cc2e8d62ae0c32e7b6fb40567d53be433eb4b17f5f996e.
//
// Solidity: event Outflow(address indexed to, address tAddr, uint256 money)
func (_IPool *IPoolFilterer) FilterOutflow(opts *bind.FilterOpts, to []common.Address) (*IPoolOutflowIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IPool.contract.FilterLogs(opts, "Outflow", toRule)
	if err != nil {
		return nil, err
	}
	return &IPoolOutflowIterator{contract: _IPool.contract, event: "Outflow", logs: logs, sub: sub}, nil
}

// WatchOutflow is a free log subscription operation binding the contract event 0x4eb2adb2eba0bb5546cc2e8d62ae0c32e7b6fb40567d53be433eb4b17f5f996e.
//
// Solidity: event Outflow(address indexed to, address tAddr, uint256 money)
func (_IPool *IPoolFilterer) WatchOutflow(opts *bind.WatchOpts, sink chan<- *IPoolOutflow, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IPool.contract.WatchLogs(opts, "Outflow", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPoolOutflow)
				if err := _IPool.contract.UnpackLog(event, "Outflow", log); err != nil {
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

// ParseOutflow is a log parse operation binding the contract event 0x4eb2adb2eba0bb5546cc2e8d62ae0c32e7b6fb40567d53be433eb4b17f5f996e.
//
// Solidity: event Outflow(address indexed to, address tAddr, uint256 money)
func (_IPool *IPoolFilterer) ParseOutflow(log types.Log) (*IPoolOutflow, error) {
	event := new(IPoolOutflow)
	if err := _IPool.contract.UnpackLog(event, "Outflow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KmanageABI is the input ABI used to generate the binding from.
const KmanageABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"mr\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"ki\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"cnt\",\"type\":\"uint64\"}],\"name\":\"AddCnt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isSet\",\"type\":\"bool\"}],\"name\":\"AddOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"ti\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"}],\"name\":\"AddProfit\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isSet\",\"type\":\"bool\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_ki\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cnt\",\"type\":\"uint64\"}],\"name\":\"addCnt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_ki\",\"type\":\"uint64\"}],\"name\":\"addKeeper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_money\",\"type\":\"uint256\"}],\"name\":\"addProfit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"auth\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_ki\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"}],\"name\":\"getK\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getKCnt\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"}],\"name\":\"getPf\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"manageRate\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"owners\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_ki\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// KmanageFuncSigs maps the 4-byte function signature to its string representation.
var KmanageFuncSigs = map[string]string{
	"4bf1b134": "add(address,bool,bytes[5])",
	"024130e4": "addCnt(uint64,uint64)",
	"50cbb46f": "addKeeper(uint64)",
	"55d3d7ef": "addProfit(uint8,uint256)",
	"de9375f2": "auth()",
	"fc3ba0ad": "balanceOf(uint64,uint8)",
	"dc0e7f45": "getK(uint64)",
	"39f2db96": "getKCnt()",
	"4f3c2eab": "getPf(uint8)",
	"6d23f6c8": "manageRate()",
	"022914a7": "owners(address)",
	"54fd4d50": "version()",
	"259c6d5e": "withdraw(uint64,uint8,uint256)",
}

// KmanageBin is the compiled bytecode used for deploying new contracts.
var KmanageBin = "0x608060405260018054600160a01b600160f81b03191665049d4002000160a11b17905534801561002e57600080fd5b5060405161125f38038061125f83398101604081905261004d91610081565b6001805460ff909216600160b01b02600162ff000160a01b03199092166001600160a01b03909316929092171790556100cd565b6000806040838503121561009457600080fd5b82516001600160a01b03811681146100ab57600080fd5b602084015190925060ff811681146100c257600080fd5b809150509250929050565b611183806100dc6000396000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c806350cbb46f1161008c5780636d23f6c8116100665780636d23f6c81461020d578063dc0e7f4514610221578063de9375f21461024c578063fc3ba0ad1461027757600080fd5b806350cbb46f146101bf57806354fd4d50146101d257806355d3d7ef146101fa57600080fd5b8063022914a7146100d4578063024130e41461010c578063259c6d5e1461012157806339f2db96146101425780634bf1b134146101585780634f3c2eab1461016b575b600080fd5b6100f76100e2366004610c7f565b60006020819052908152604090205460ff1681565b60405190151581526020015b60405180910390f35b61011f61011a366004610cb1565b61029f565b005b61013461012f366004610cf5565b61041f565b604051908152602001610103565b6003545b60405160ff9091168152602001610103565b61011f610166366004610d9f565b61080b565b6101a0610179366004610eb4565b60ff16600090815260056020526040902080546001909101546001600160401b0390911691565b604080516001600160401b039093168352602083019190915201610103565b61011f6101cd366004610ecf565b610984565b6001546101e790600160a01b900461ffff1681565b60405161ffff9091168152602001610103565b61011f610208366004610eea565b610ab6565b60015461014690600160b01b900460ff1681565b61023461022f366004610ecf565b610b44565b6040516001600160401b039091168152602001610103565b60015461025f906001600160a01b031681565b6040516001600160a01b039091168152602001610103565b61028a610285366004610f14565b610b92565b60408051928352602083019190915201610103565b3360009081526020819052604090205460ff166102d75760405162461bcd60e51b81526004016102ce90610f3e565b60405180910390fd5b6001600160401b03808316600090815260046020526040902054166103335760405162461bcd60e51b815260206004820152601260248201527139a7b93232b91d34b636329031b0b63632b960711b60448201526064016102ce565b6001600160401b0380831660009081526004602052604081208054849391929161035f91859116610f77565b92506101000a8154816001600160401b0302191690836001600160401b0316021790555080600260008282829054906101000a90046001600160401b03166103a79190610f77565b92506101000a8154816001600160401b0302191690836001600160401b03160217905550816001600160401b03167f5372d6aad551334f508508499c71755ebc6cde46b83fc1f944f1f0ae33cbb4c48260405161041391906001600160401b0391909116815260200190565b60405180910390a25050565b3360009081526020819052604081205460ff1661044e5760405162461bcd60e51b81526004016102ce90610f3e565b6001600160401b03808516600090815260046020526040812054909116900361047957506000610804565b60015460ff841660009081526005602052604090205442916001600160401b03600160b81b9091048116916104af911683610fa2565b6001600160401b031611156107905760ff84166000908152600560205260408120805467ffffffffffffffff19166001600160401b03938416178155600254600190910154919283926105059290911690610fe0565b905060005b600354811015610771576000600460006003848154811061052d5761052d610ff4565b600091825260208083206004830401546001600160401b0360086003909416939093026101000a900482168452830193909352604090910190205461057391168461100a565b90506002600460006003858154811061058e5761058e610ff4565b600091825260208083206004830401546001600160401b0360086003909416939093026101000a90048216845283019390935260409091019020546105d4929116611029565b6105df906001610f77565b60046000600385815481106105f6576105f6610ff4565b90600052602060002090600491828204019190066008029054906101000a90046001600160401b03166001600160401b03166001600160401b0316815260200190815260200160002060006101000a8154816001600160401b0302191690836001600160401b03160217905550600460006003848154811061067a5761067a610ff4565b600091825260208083206004830401546001600160401b0360086003909416939093026101000a90048216845283019390935260409091019020546106c0911685610f77565b93508060066000600385815481106106da576106da610ff4565b6000918252602080832060048304015460039092166008026101000a9091046001600160401b03168352828101939093526040918201812060ff8c1682529092528120805490919061072d90849061104f565b909155505060ff871660009081526005602052604081206001018054839290610757908490611067565b9091555082915061076990508161107e565b91505061050a565b50506002805467ffffffffffffffff19166001600160401b0383161790555b6001600160401b038516600090815260066020908152604080832060ff88168452909152902054808411156107c3578093505b6001600160401b038616600090815260066020908152604080832060ff89168452909152812080548692906107f9908490611067565b909155508493505050505b9392505050565b823b60008190036108535760405162461bcd60e51b81526020600482015260126024820152713732b2b21031b7b73a3930b1ba1030b2323960711b60448201526064016102ce565b6040516bffffffffffffffffffffffff1930606090811b821660208401526218591960ea1b603484015286901b16603782015283151560f81b604b820152600090604c0160408051601f1981840301815290829052805160209091012060015463a96bba9d60e01b83529092506001600160a01b03169063a96bba9d906108e09084908790600401611097565b600060405180830381600087803b1580156108fa57600080fd5b505af115801561090e573d6000803e3d6000fd5b5050604080516001600160a01b038916815287151560208201527f938b2a24c98e4e542127ffa74a91e48942c2bddccae3b6d75f82bfda7bbc0807935001905060405180910390a15050506001600160a01b03919091166000908152602081905260409020805460ff1916911515919091179055565b3360009081526020819052604090205460ff166109b35760405162461bcd60e51b81526004016102ce90610f3e565b6001600160401b038082166000908152600460205260409020541615610a055760405162461bcd60e51b81526020600482015260076024820152661ac8195e1a5cdd60ca1b60448201526064016102ce565b600380546001818101835560048083047fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b0180546001600160401b03808816600896909716959095026101000a86810290860219909116179055600093845260205260408320805467ffffffffffffffff191690911790556002805490911691610a8e83611127565b91906101000a8154816001600160401b0302191690836001600160401b031602179055505050565b3360009081526020819052604090205460ff16610ae55760405162461bcd60e51b81526004016102ce90610f3e565b60ff821660009081526005602052604081206001018054839290610b0a90849061104f565b909155505060405181815260ff8316907fe5fd8ec20bdfeb1fadd32ec3786545a1db18d74f952effda5c6cd50af7c2e9e890602001610413565b60006003826001600160401b031681548110610b6257610b62610ff4565b90600052602060002090600491828204019190066008029054906101000a90046001600160401b03169050919050565b6001600160401b03808316600081815260066020908152604080832060ff871680855290835281842054948452600483528184205460025491855260059093529083206001015492948594938593849390831692610bf292911690610fe0565b610bfc919061100a565b60015460ff88166000908152600560205260409020549192506001600160401b03600160b81b909104811691610c33911642611067565b1115610c4a57610c43818461104f565b9250610c57565b610c54818361104f565b91505b50909590945092505050565b80356001600160a01b0381168114610c7a57600080fd5b919050565b600060208284031215610c9157600080fd5b61080482610c63565b80356001600160401b0381168114610c7a57600080fd5b60008060408385031215610cc457600080fd5b610ccd83610c9a565b9150610cdb60208401610c9a565b90509250929050565b803560ff81168114610c7a57600080fd5b600080600060608486031215610d0a57600080fd5b610d1384610c9a565b9250610d2160208501610ce4565b9150604084013590509250925092565b634e487b7160e01b600052604160045260246000fd5b60405160a081016001600160401b0381118282101715610d6957610d69610d31565b60405290565b604051601f8201601f191681016001600160401b0381118282101715610d9757610d97610d31565b604052919050565b600080600060608486031215610db457600080fd5b610dbd84610c63565b92506020808501358015158114610dd357600080fd5b925060408501356001600160401b0380821115610def57600080fd5b8187019150601f8881840112610e0457600080fd5b610e0c610d47565b8060a085018b811115610e1e57600080fd5b855b81811015610ea257803586811115610e385760008081fd5b87018581018e13610e495760008081fd5b803587811115610e5b57610e5b610d31565b610e6c818801601f19168b01610d6f565b8181528f8b838501011115610e815760008081fd5b818b84018c83013760009181018b0191909152855250928701928701610e20565b50508096505050505050509250925092565b600060208284031215610ec657600080fd5b61080482610ce4565b600060208284031215610ee157600080fd5b61080482610c9a565b60008060408385031215610efd57600080fd5b610f0683610ce4565b946020939093013593505050565b60008060408385031215610f2757600080fd5b610f3083610c9a565b9150610cdb60208401610ce4565b6020808252600990820152683737ba1037bbb732b960b91b604082015260600190565b634e487b7160e01b600052601160045260246000fd5b60006001600160401b03808316818516808303821115610f9957610f99610f61565b01949350505050565b60006001600160401b0383811690831681811015610fc257610fc2610f61565b039392505050565b634e487b7160e01b600052601260045260246000fd5b600082610fef57610fef610fca565b500490565b634e487b7160e01b600052603260045260246000fd5b600081600019048311821515161561102457611024610f61565b500290565b60006001600160401b038084168061104357611043610fca565b92169190910492915050565b6000821982111561106257611062610f61565b500190565b60008282101561107957611079610f61565b500390565b60006001820161109057611090610f61565b5060010190565b8281526040602080830182905260009160e08401919084018584805b600581101561111957878603603f1901845282518051808852835b818110156110e95782810188015189820189015287016110ce565b818111156110f9578488838b0101525b50601f01601f1916969096018501955092840192918401916001016110b3565b509398975050505050505050565b60006001600160401b0380831681810361114357611143610f61565b600101939250505056fea2646970667358221220d61b4c43fbc2afc6384ec21a987e1f29503e04a5de72681e2e2579554deeec9b64736f6c634300080e0033"

// DeployKmanage deploys a new Ethereum contract, binding an instance of Kmanage to it.
func DeployKmanage(auth *bind.TransactOpts, backend bind.ContractBackend, _a common.Address, mr uint8) (common.Address, *types.Transaction, *Kmanage, error) {
	parsed, err := abi.JSON(strings.NewReader(KmanageABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(KmanageBin), backend, _a, mr)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Kmanage{KmanageCaller: KmanageCaller{contract: contract}, KmanageTransactor: KmanageTransactor{contract: contract}, KmanageFilterer: KmanageFilterer{contract: contract}}, nil
}

// Kmanage is an auto generated Go binding around an Ethereum contract.
type Kmanage struct {
	KmanageCaller     // Read-only binding to the contract
	KmanageTransactor // Write-only binding to the contract
	KmanageFilterer   // Log filterer for contract events
}

// KmanageCaller is an auto generated read-only Go binding around an Ethereum contract.
type KmanageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KmanageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KmanageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KmanageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KmanageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KmanageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KmanageSession struct {
	Contract     *Kmanage          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KmanageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KmanageCallerSession struct {
	Contract *KmanageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// KmanageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KmanageTransactorSession struct {
	Contract     *KmanageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// KmanageRaw is an auto generated low-level Go binding around an Ethereum contract.
type KmanageRaw struct {
	Contract *Kmanage // Generic contract binding to access the raw methods on
}

// KmanageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KmanageCallerRaw struct {
	Contract *KmanageCaller // Generic read-only contract binding to access the raw methods on
}

// KmanageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KmanageTransactorRaw struct {
	Contract *KmanageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKmanage creates a new instance of Kmanage, bound to a specific deployed contract.
func NewKmanage(address common.Address, backend bind.ContractBackend) (*Kmanage, error) {
	contract, err := bindKmanage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Kmanage{KmanageCaller: KmanageCaller{contract: contract}, KmanageTransactor: KmanageTransactor{contract: contract}, KmanageFilterer: KmanageFilterer{contract: contract}}, nil
}

// NewKmanageCaller creates a new read-only instance of Kmanage, bound to a specific deployed contract.
func NewKmanageCaller(address common.Address, caller bind.ContractCaller) (*KmanageCaller, error) {
	contract, err := bindKmanage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KmanageCaller{contract: contract}, nil
}

// NewKmanageTransactor creates a new write-only instance of Kmanage, bound to a specific deployed contract.
func NewKmanageTransactor(address common.Address, transactor bind.ContractTransactor) (*KmanageTransactor, error) {
	contract, err := bindKmanage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KmanageTransactor{contract: contract}, nil
}

// NewKmanageFilterer creates a new log filterer instance of Kmanage, bound to a specific deployed contract.
func NewKmanageFilterer(address common.Address, filterer bind.ContractFilterer) (*KmanageFilterer, error) {
	contract, err := bindKmanage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KmanageFilterer{contract: contract}, nil
}

// bindKmanage binds a generic wrapper to an already deployed contract.
func bindKmanage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KmanageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Kmanage *KmanageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Kmanage.Contract.KmanageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Kmanage *KmanageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Kmanage.Contract.KmanageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Kmanage *KmanageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Kmanage.Contract.KmanageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Kmanage *KmanageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Kmanage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Kmanage *KmanageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Kmanage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Kmanage *KmanageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Kmanage.Contract.contract.Transact(opts, method, params...)
}

// Auth is a free data retrieval call binding the contract method 0xde9375f2.
//
// Solidity: function auth() view returns(address)
func (_Kmanage *KmanageCaller) Auth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Kmanage.contract.Call(opts, &out, "auth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Auth is a free data retrieval call binding the contract method 0xde9375f2.
//
// Solidity: function auth() view returns(address)
func (_Kmanage *KmanageSession) Auth() (common.Address, error) {
	return _Kmanage.Contract.Auth(&_Kmanage.CallOpts)
}

// Auth is a free data retrieval call binding the contract method 0xde9375f2.
//
// Solidity: function auth() view returns(address)
func (_Kmanage *KmanageCallerSession) Auth() (common.Address, error) {
	return _Kmanage.Contract.Auth(&_Kmanage.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0xfc3ba0ad.
//
// Solidity: function balanceOf(uint64 _ki, uint8 _ti) view returns(uint256, uint256)
func (_Kmanage *KmanageCaller) BalanceOf(opts *bind.CallOpts, _ki uint64, _ti uint8) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Kmanage.contract.Call(opts, &out, "balanceOf", _ki, _ti)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// BalanceOf is a free data retrieval call binding the contract method 0xfc3ba0ad.
//
// Solidity: function balanceOf(uint64 _ki, uint8 _ti) view returns(uint256, uint256)
func (_Kmanage *KmanageSession) BalanceOf(_ki uint64, _ti uint8) (*big.Int, *big.Int, error) {
	return _Kmanage.Contract.BalanceOf(&_Kmanage.CallOpts, _ki, _ti)
}

// BalanceOf is a free data retrieval call binding the contract method 0xfc3ba0ad.
//
// Solidity: function balanceOf(uint64 _ki, uint8 _ti) view returns(uint256, uint256)
func (_Kmanage *KmanageCallerSession) BalanceOf(_ki uint64, _ti uint8) (*big.Int, *big.Int, error) {
	return _Kmanage.Contract.BalanceOf(&_Kmanage.CallOpts, _ki, _ti)
}

// GetK is a free data retrieval call binding the contract method 0xdc0e7f45.
//
// Solidity: function getK(uint64 _i) view returns(uint64)
func (_Kmanage *KmanageCaller) GetK(opts *bind.CallOpts, _i uint64) (uint64, error) {
	var out []interface{}
	err := _Kmanage.contract.Call(opts, &out, "getK", _i)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetK is a free data retrieval call binding the contract method 0xdc0e7f45.
//
// Solidity: function getK(uint64 _i) view returns(uint64)
func (_Kmanage *KmanageSession) GetK(_i uint64) (uint64, error) {
	return _Kmanage.Contract.GetK(&_Kmanage.CallOpts, _i)
}

// GetK is a free data retrieval call binding the contract method 0xdc0e7f45.
//
// Solidity: function getK(uint64 _i) view returns(uint64)
func (_Kmanage *KmanageCallerSession) GetK(_i uint64) (uint64, error) {
	return _Kmanage.Contract.GetK(&_Kmanage.CallOpts, _i)
}

// GetKCnt is a free data retrieval call binding the contract method 0x39f2db96.
//
// Solidity: function getKCnt() view returns(uint8)
func (_Kmanage *KmanageCaller) GetKCnt(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Kmanage.contract.Call(opts, &out, "getKCnt")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetKCnt is a free data retrieval call binding the contract method 0x39f2db96.
//
// Solidity: function getKCnt() view returns(uint8)
func (_Kmanage *KmanageSession) GetKCnt() (uint8, error) {
	return _Kmanage.Contract.GetKCnt(&_Kmanage.CallOpts)
}

// GetKCnt is a free data retrieval call binding the contract method 0x39f2db96.
//
// Solidity: function getKCnt() view returns(uint8)
func (_Kmanage *KmanageCallerSession) GetKCnt() (uint8, error) {
	return _Kmanage.Contract.GetKCnt(&_Kmanage.CallOpts)
}

// GetPf is a free data retrieval call binding the contract method 0x4f3c2eab.
//
// Solidity: function getPf(uint8 _ti) view returns(uint64, uint256)
func (_Kmanage *KmanageCaller) GetPf(opts *bind.CallOpts, _ti uint8) (uint64, *big.Int, error) {
	var out []interface{}
	err := _Kmanage.contract.Call(opts, &out, "getPf", _ti)

	if err != nil {
		return *new(uint64), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetPf is a free data retrieval call binding the contract method 0x4f3c2eab.
//
// Solidity: function getPf(uint8 _ti) view returns(uint64, uint256)
func (_Kmanage *KmanageSession) GetPf(_ti uint8) (uint64, *big.Int, error) {
	return _Kmanage.Contract.GetPf(&_Kmanage.CallOpts, _ti)
}

// GetPf is a free data retrieval call binding the contract method 0x4f3c2eab.
//
// Solidity: function getPf(uint8 _ti) view returns(uint64, uint256)
func (_Kmanage *KmanageCallerSession) GetPf(_ti uint8) (uint64, *big.Int, error) {
	return _Kmanage.Contract.GetPf(&_Kmanage.CallOpts, _ti)
}

// ManageRate is a free data retrieval call binding the contract method 0x6d23f6c8.
//
// Solidity: function manageRate() view returns(uint8)
func (_Kmanage *KmanageCaller) ManageRate(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Kmanage.contract.Call(opts, &out, "manageRate")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ManageRate is a free data retrieval call binding the contract method 0x6d23f6c8.
//
// Solidity: function manageRate() view returns(uint8)
func (_Kmanage *KmanageSession) ManageRate() (uint8, error) {
	return _Kmanage.Contract.ManageRate(&_Kmanage.CallOpts)
}

// ManageRate is a free data retrieval call binding the contract method 0x6d23f6c8.
//
// Solidity: function manageRate() view returns(uint8)
func (_Kmanage *KmanageCallerSession) ManageRate() (uint8, error) {
	return _Kmanage.Contract.ManageRate(&_Kmanage.CallOpts)
}

// Owners is a free data retrieval call binding the contract method 0x022914a7.
//
// Solidity: function owners(address ) view returns(bool)
func (_Kmanage *KmanageCaller) Owners(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Kmanage.contract.Call(opts, &out, "owners", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Owners is a free data retrieval call binding the contract method 0x022914a7.
//
// Solidity: function owners(address ) view returns(bool)
func (_Kmanage *KmanageSession) Owners(arg0 common.Address) (bool, error) {
	return _Kmanage.Contract.Owners(&_Kmanage.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x022914a7.
//
// Solidity: function owners(address ) view returns(bool)
func (_Kmanage *KmanageCallerSession) Owners(arg0 common.Address) (bool, error) {
	return _Kmanage.Contract.Owners(&_Kmanage.CallOpts, arg0)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint16)
func (_Kmanage *KmanageCaller) Version(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Kmanage.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint16)
func (_Kmanage *KmanageSession) Version() (uint16, error) {
	return _Kmanage.Contract.Version(&_Kmanage.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint16)
func (_Kmanage *KmanageCallerSession) Version() (uint16, error) {
	return _Kmanage.Contract.Version(&_Kmanage.CallOpts)
}

// Add is a paid mutator transaction binding the contract method 0x4bf1b134.
//
// Solidity: function add(address _a, bool isSet, bytes[5] signs) returns()
func (_Kmanage *KmanageTransactor) Add(opts *bind.TransactOpts, _a common.Address, isSet bool, signs [5][]byte) (*types.Transaction, error) {
	return _Kmanage.contract.Transact(opts, "add", _a, isSet, signs)
}

// Add is a paid mutator transaction binding the contract method 0x4bf1b134.
//
// Solidity: function add(address _a, bool isSet, bytes[5] signs) returns()
func (_Kmanage *KmanageSession) Add(_a common.Address, isSet bool, signs [5][]byte) (*types.Transaction, error) {
	return _Kmanage.Contract.Add(&_Kmanage.TransactOpts, _a, isSet, signs)
}

// Add is a paid mutator transaction binding the contract method 0x4bf1b134.
//
// Solidity: function add(address _a, bool isSet, bytes[5] signs) returns()
func (_Kmanage *KmanageTransactorSession) Add(_a common.Address, isSet bool, signs [5][]byte) (*types.Transaction, error) {
	return _Kmanage.Contract.Add(&_Kmanage.TransactOpts, _a, isSet, signs)
}

// AddCnt is a paid mutator transaction binding the contract method 0x024130e4.
//
// Solidity: function addCnt(uint64 _ki, uint64 cnt) returns()
func (_Kmanage *KmanageTransactor) AddCnt(opts *bind.TransactOpts, _ki uint64, cnt uint64) (*types.Transaction, error) {
	return _Kmanage.contract.Transact(opts, "addCnt", _ki, cnt)
}

// AddCnt is a paid mutator transaction binding the contract method 0x024130e4.
//
// Solidity: function addCnt(uint64 _ki, uint64 cnt) returns()
func (_Kmanage *KmanageSession) AddCnt(_ki uint64, cnt uint64) (*types.Transaction, error) {
	return _Kmanage.Contract.AddCnt(&_Kmanage.TransactOpts, _ki, cnt)
}

// AddCnt is a paid mutator transaction binding the contract method 0x024130e4.
//
// Solidity: function addCnt(uint64 _ki, uint64 cnt) returns()
func (_Kmanage *KmanageTransactorSession) AddCnt(_ki uint64, cnt uint64) (*types.Transaction, error) {
	return _Kmanage.Contract.AddCnt(&_Kmanage.TransactOpts, _ki, cnt)
}

// AddKeeper is a paid mutator transaction binding the contract method 0x50cbb46f.
//
// Solidity: function addKeeper(uint64 _ki) returns()
func (_Kmanage *KmanageTransactor) AddKeeper(opts *bind.TransactOpts, _ki uint64) (*types.Transaction, error) {
	return _Kmanage.contract.Transact(opts, "addKeeper", _ki)
}

// AddKeeper is a paid mutator transaction binding the contract method 0x50cbb46f.
//
// Solidity: function addKeeper(uint64 _ki) returns()
func (_Kmanage *KmanageSession) AddKeeper(_ki uint64) (*types.Transaction, error) {
	return _Kmanage.Contract.AddKeeper(&_Kmanage.TransactOpts, _ki)
}

// AddKeeper is a paid mutator transaction binding the contract method 0x50cbb46f.
//
// Solidity: function addKeeper(uint64 _ki) returns()
func (_Kmanage *KmanageTransactorSession) AddKeeper(_ki uint64) (*types.Transaction, error) {
	return _Kmanage.Contract.AddKeeper(&_Kmanage.TransactOpts, _ki)
}

// AddProfit is a paid mutator transaction binding the contract method 0x55d3d7ef.
//
// Solidity: function addProfit(uint8 _ti, uint256 _money) returns()
func (_Kmanage *KmanageTransactor) AddProfit(opts *bind.TransactOpts, _ti uint8, _money *big.Int) (*types.Transaction, error) {
	return _Kmanage.contract.Transact(opts, "addProfit", _ti, _money)
}

// AddProfit is a paid mutator transaction binding the contract method 0x55d3d7ef.
//
// Solidity: function addProfit(uint8 _ti, uint256 _money) returns()
func (_Kmanage *KmanageSession) AddProfit(_ti uint8, _money *big.Int) (*types.Transaction, error) {
	return _Kmanage.Contract.AddProfit(&_Kmanage.TransactOpts, _ti, _money)
}

// AddProfit is a paid mutator transaction binding the contract method 0x55d3d7ef.
//
// Solidity: function addProfit(uint8 _ti, uint256 _money) returns()
func (_Kmanage *KmanageTransactorSession) AddProfit(_ti uint8, _money *big.Int) (*types.Transaction, error) {
	return _Kmanage.Contract.AddProfit(&_Kmanage.TransactOpts, _ti, _money)
}

// Withdraw is a paid mutator transaction binding the contract method 0x259c6d5e.
//
// Solidity: function withdraw(uint64 _ki, uint8 _ti, uint256 amount) returns(uint256)
func (_Kmanage *KmanageTransactor) Withdraw(opts *bind.TransactOpts, _ki uint64, _ti uint8, amount *big.Int) (*types.Transaction, error) {
	return _Kmanage.contract.Transact(opts, "withdraw", _ki, _ti, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x259c6d5e.
//
// Solidity: function withdraw(uint64 _ki, uint8 _ti, uint256 amount) returns(uint256)
func (_Kmanage *KmanageSession) Withdraw(_ki uint64, _ti uint8, amount *big.Int) (*types.Transaction, error) {
	return _Kmanage.Contract.Withdraw(&_Kmanage.TransactOpts, _ki, _ti, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x259c6d5e.
//
// Solidity: function withdraw(uint64 _ki, uint8 _ti, uint256 amount) returns(uint256)
func (_Kmanage *KmanageTransactorSession) Withdraw(_ki uint64, _ti uint8, amount *big.Int) (*types.Transaction, error) {
	return _Kmanage.Contract.Withdraw(&_Kmanage.TransactOpts, _ki, _ti, amount)
}

// KmanageAddCntIterator is returned from FilterAddCnt and is used to iterate over the raw logs and unpacked data for AddCnt events raised by the Kmanage contract.
type KmanageAddCntIterator struct {
	Event *KmanageAddCnt // Event containing the contract specifics and raw log

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
func (it *KmanageAddCntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KmanageAddCnt)
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
		it.Event = new(KmanageAddCnt)
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
func (it *KmanageAddCntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KmanageAddCntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KmanageAddCnt represents a AddCnt event raised by the Kmanage contract.
type KmanageAddCnt struct {
	Ki  uint64
	Cnt uint64
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAddCnt is a free log retrieval operation binding the contract event 0x5372d6aad551334f508508499c71755ebc6cde46b83fc1f944f1f0ae33cbb4c4.
//
// Solidity: event AddCnt(uint64 indexed ki, uint64 cnt)
func (_Kmanage *KmanageFilterer) FilterAddCnt(opts *bind.FilterOpts, ki []uint64) (*KmanageAddCntIterator, error) {

	var kiRule []interface{}
	for _, kiItem := range ki {
		kiRule = append(kiRule, kiItem)
	}

	logs, sub, err := _Kmanage.contract.FilterLogs(opts, "AddCnt", kiRule)
	if err != nil {
		return nil, err
	}
	return &KmanageAddCntIterator{contract: _Kmanage.contract, event: "AddCnt", logs: logs, sub: sub}, nil
}

// WatchAddCnt is a free log subscription operation binding the contract event 0x5372d6aad551334f508508499c71755ebc6cde46b83fc1f944f1f0ae33cbb4c4.
//
// Solidity: event AddCnt(uint64 indexed ki, uint64 cnt)
func (_Kmanage *KmanageFilterer) WatchAddCnt(opts *bind.WatchOpts, sink chan<- *KmanageAddCnt, ki []uint64) (event.Subscription, error) {

	var kiRule []interface{}
	for _, kiItem := range ki {
		kiRule = append(kiRule, kiItem)
	}

	logs, sub, err := _Kmanage.contract.WatchLogs(opts, "AddCnt", kiRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KmanageAddCnt)
				if err := _Kmanage.contract.UnpackLog(event, "AddCnt", log); err != nil {
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

// ParseAddCnt is a log parse operation binding the contract event 0x5372d6aad551334f508508499c71755ebc6cde46b83fc1f944f1f0ae33cbb4c4.
//
// Solidity: event AddCnt(uint64 indexed ki, uint64 cnt)
func (_Kmanage *KmanageFilterer) ParseAddCnt(log types.Log) (*KmanageAddCnt, error) {
	event := new(KmanageAddCnt)
	if err := _Kmanage.contract.UnpackLog(event, "AddCnt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KmanageAddOwnerIterator is returned from FilterAddOwner and is used to iterate over the raw logs and unpacked data for AddOwner events raised by the Kmanage contract.
type KmanageAddOwnerIterator struct {
	Event *KmanageAddOwner // Event containing the contract specifics and raw log

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
func (it *KmanageAddOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KmanageAddOwner)
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
		it.Event = new(KmanageAddOwner)
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
func (it *KmanageAddOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KmanageAddOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KmanageAddOwner represents a AddOwner event raised by the Kmanage contract.
type KmanageAddOwner struct {
	From  common.Address
	IsSet bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAddOwner is a free log retrieval operation binding the contract event 0x938b2a24c98e4e542127ffa74a91e48942c2bddccae3b6d75f82bfda7bbc0807.
//
// Solidity: event AddOwner(address from, bool isSet)
func (_Kmanage *KmanageFilterer) FilterAddOwner(opts *bind.FilterOpts) (*KmanageAddOwnerIterator, error) {

	logs, sub, err := _Kmanage.contract.FilterLogs(opts, "AddOwner")
	if err != nil {
		return nil, err
	}
	return &KmanageAddOwnerIterator{contract: _Kmanage.contract, event: "AddOwner", logs: logs, sub: sub}, nil
}

// WatchAddOwner is a free log subscription operation binding the contract event 0x938b2a24c98e4e542127ffa74a91e48942c2bddccae3b6d75f82bfda7bbc0807.
//
// Solidity: event AddOwner(address from, bool isSet)
func (_Kmanage *KmanageFilterer) WatchAddOwner(opts *bind.WatchOpts, sink chan<- *KmanageAddOwner) (event.Subscription, error) {

	logs, sub, err := _Kmanage.contract.WatchLogs(opts, "AddOwner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KmanageAddOwner)
				if err := _Kmanage.contract.UnpackLog(event, "AddOwner", log); err != nil {
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

// ParseAddOwner is a log parse operation binding the contract event 0x938b2a24c98e4e542127ffa74a91e48942c2bddccae3b6d75f82bfda7bbc0807.
//
// Solidity: event AddOwner(address from, bool isSet)
func (_Kmanage *KmanageFilterer) ParseAddOwner(log types.Log) (*KmanageAddOwner, error) {
	event := new(KmanageAddOwner)
	if err := _Kmanage.contract.UnpackLog(event, "AddOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KmanageAddProfitIterator is returned from FilterAddProfit and is used to iterate over the raw logs and unpacked data for AddProfit events raised by the Kmanage contract.
type KmanageAddProfitIterator struct {
	Event *KmanageAddProfit // Event containing the contract specifics and raw log

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
func (it *KmanageAddProfitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KmanageAddProfit)
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
		it.Event = new(KmanageAddProfit)
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
func (it *KmanageAddProfitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KmanageAddProfitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KmanageAddProfit represents a AddProfit event raised by the Kmanage contract.
type KmanageAddProfit struct {
	Ti    uint8
	Money *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAddProfit is a free log retrieval operation binding the contract event 0xe5fd8ec20bdfeb1fadd32ec3786545a1db18d74f952effda5c6cd50af7c2e9e8.
//
// Solidity: event AddProfit(uint8 indexed ti, uint256 money)
func (_Kmanage *KmanageFilterer) FilterAddProfit(opts *bind.FilterOpts, ti []uint8) (*KmanageAddProfitIterator, error) {

	var tiRule []interface{}
	for _, tiItem := range ti {
		tiRule = append(tiRule, tiItem)
	}

	logs, sub, err := _Kmanage.contract.FilterLogs(opts, "AddProfit", tiRule)
	if err != nil {
		return nil, err
	}
	return &KmanageAddProfitIterator{contract: _Kmanage.contract, event: "AddProfit", logs: logs, sub: sub}, nil
}

// WatchAddProfit is a free log subscription operation binding the contract event 0xe5fd8ec20bdfeb1fadd32ec3786545a1db18d74f952effda5c6cd50af7c2e9e8.
//
// Solidity: event AddProfit(uint8 indexed ti, uint256 money)
func (_Kmanage *KmanageFilterer) WatchAddProfit(opts *bind.WatchOpts, sink chan<- *KmanageAddProfit, ti []uint8) (event.Subscription, error) {

	var tiRule []interface{}
	for _, tiItem := range ti {
		tiRule = append(tiRule, tiItem)
	}

	logs, sub, err := _Kmanage.contract.WatchLogs(opts, "AddProfit", tiRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KmanageAddProfit)
				if err := _Kmanage.contract.UnpackLog(event, "AddProfit", log); err != nil {
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

// ParseAddProfit is a log parse operation binding the contract event 0xe5fd8ec20bdfeb1fadd32ec3786545a1db18d74f952effda5c6cd50af7c2e9e8.
//
// Solidity: event AddProfit(uint8 indexed ti, uint256 money)
func (_Kmanage *KmanageFilterer) ParseAddProfit(log types.Log) (*KmanageAddProfit, error) {
	event := new(KmanageAddProfit)
	if err := _Kmanage.contract.UnpackLog(event, "AddProfit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwnerABI is the input ABI used to generate the binding from.
const OwnerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isSet\",\"type\":\"bool\"}],\"name\":\"AddOwner\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isSet\",\"type\":\"bool\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"auth\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"owners\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// OwnerFuncSigs maps the 4-byte function signature to its string representation.
var OwnerFuncSigs = map[string]string{
	"4bf1b134": "add(address,bool,bytes[5])",
	"de9375f2": "auth()",
	"022914a7": "owners(address)",
}

// OwnerBin is the compiled bytecode used for deploying new contracts.
var OwnerBin = "0x608060405234801561001057600080fd5b5060405161055838038061055883398101604081905261002f91610054565b600180546001600160a01b0319166001600160a01b0392909216919091179055610084565b60006020828403121561006657600080fd5b81516001600160a01b038116811461007d57600080fd5b9392505050565b6104c5806100936000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c8063022914a7146100465780634bf1b1341461007e578063de9375f214610093575b600080fd5b610069610054366004610257565b60006020819052908152604090205460ff1681565b60405190151581526020015b60405180910390f35b61009161008c3660046102e9565b6100be565b005b6001546100a6906001600160a01b031681565b6040516001600160a01b039091168152602001610075565b823b600081900361010a5760405162461bcd60e51b81526020600482015260126024820152713732b2b21031b7b73a3930b1ba1030b2323960711b604482015260640160405180910390fd5b6040516bffffffffffffffffffffffff1930606090811b821660208401526218591960ea1b603484015286901b16603782015283151560f81b604b820152600090604c0160408051601f1981840301815290829052805160209091012060015463a96bba9d60e01b83529092506001600160a01b03169063a96bba9d9061019790849087906004016103ff565b600060405180830381600087803b1580156101b157600080fd5b505af11580156101c5573d6000803e3d6000fd5b5050604080516001600160a01b038916815287151560208201527f938b2a24c98e4e542127ffa74a91e48942c2bddccae3b6d75f82bfda7bbc0807935001905060405180910390a15050506001600160a01b03919091166000908152602081905260409020805460ff1916911515919091179055565b80356001600160a01b038116811461025257600080fd5b919050565b60006020828403121561026957600080fd5b6102728261023b565b9392505050565b634e487b7160e01b600052604160045260246000fd5b60405160a0810167ffffffffffffffff811182821017156102b2576102b2610279565b60405290565b604051601f8201601f1916810167ffffffffffffffff811182821017156102e1576102e1610279565b604052919050565b6000806000606084860312156102fe57600080fd5b6103078461023b565b9250602080850135801515811461031d57600080fd5b9250604085013567ffffffffffffffff8082111561033a57600080fd5b8187019150601f888184011261034f57600080fd5b61035761028f565b8060a085018b81111561036957600080fd5b855b818110156103ed578035868111156103835760008081fd5b87018581018e136103945760008081fd5b8035878111156103a6576103a6610279565b6103b7818801601f19168b016102b8565b8181528f8b8385010111156103cc5760008081fd5b818b84018c83013760009181018b019190915285525092870192870161036b565b50508096505050505050509250925092565b8281526040602080830182905260009160e08401919084018584805b600581101561048157878603603f1901845282518051808852835b81811015610451578281018801518982018901528701610436565b81811115610461578488838b0101525b50601f01601f19169690960185019550928401929184019160010161041b565b50939897505050505050505056fea2646970667358221220e4109531dbabeb33ccbfa77f28019b7953ba57d32bea8b64099ee0dd97784c0464736f6c634300080e0033"

// DeployOwner deploys a new Ethereum contract, binding an instance of Owner to it.
func DeployOwner(auth *bind.TransactOpts, backend bind.ContractBackend, _a common.Address) (common.Address, *types.Transaction, *Owner, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OwnerBin), backend, _a)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Owner{OwnerCaller: OwnerCaller{contract: contract}, OwnerTransactor: OwnerTransactor{contract: contract}, OwnerFilterer: OwnerFilterer{contract: contract}}, nil
}

// Owner is an auto generated Go binding around an Ethereum contract.
type Owner struct {
	OwnerCaller     // Read-only binding to the contract
	OwnerTransactor // Write-only binding to the contract
	OwnerFilterer   // Log filterer for contract events
}

// OwnerCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnerSession struct {
	Contract     *Owner            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnerCallerSession struct {
	Contract *OwnerCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OwnerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnerTransactorSession struct {
	Contract     *OwnerTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnerRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnerRaw struct {
	Contract *Owner // Generic contract binding to access the raw methods on
}

// OwnerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnerCallerRaw struct {
	Contract *OwnerCaller // Generic read-only contract binding to access the raw methods on
}

// OwnerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnerTransactorRaw struct {
	Contract *OwnerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwner creates a new instance of Owner, bound to a specific deployed contract.
func NewOwner(address common.Address, backend bind.ContractBackend) (*Owner, error) {
	contract, err := bindOwner(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Owner{OwnerCaller: OwnerCaller{contract: contract}, OwnerTransactor: OwnerTransactor{contract: contract}, OwnerFilterer: OwnerFilterer{contract: contract}}, nil
}

// NewOwnerCaller creates a new read-only instance of Owner, bound to a specific deployed contract.
func NewOwnerCaller(address common.Address, caller bind.ContractCaller) (*OwnerCaller, error) {
	contract, err := bindOwner(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnerCaller{contract: contract}, nil
}

// NewOwnerTransactor creates a new write-only instance of Owner, bound to a specific deployed contract.
func NewOwnerTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnerTransactor, error) {
	contract, err := bindOwner(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnerTransactor{contract: contract}, nil
}

// NewOwnerFilterer creates a new log filterer instance of Owner, bound to a specific deployed contract.
func NewOwnerFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnerFilterer, error) {
	contract, err := bindOwner(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnerFilterer{contract: contract}, nil
}

// bindOwner binds a generic wrapper to an already deployed contract.
func bindOwner(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Owner *OwnerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Owner.Contract.OwnerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Owner *OwnerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Owner.Contract.OwnerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Owner *OwnerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Owner.Contract.OwnerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Owner *OwnerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Owner.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Owner *OwnerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Owner.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Owner *OwnerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Owner.Contract.contract.Transact(opts, method, params...)
}

// Auth is a free data retrieval call binding the contract method 0xde9375f2.
//
// Solidity: function auth() view returns(address)
func (_Owner *OwnerCaller) Auth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Owner.contract.Call(opts, &out, "auth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Auth is a free data retrieval call binding the contract method 0xde9375f2.
//
// Solidity: function auth() view returns(address)
func (_Owner *OwnerSession) Auth() (common.Address, error) {
	return _Owner.Contract.Auth(&_Owner.CallOpts)
}

// Auth is a free data retrieval call binding the contract method 0xde9375f2.
//
// Solidity: function auth() view returns(address)
func (_Owner *OwnerCallerSession) Auth() (common.Address, error) {
	return _Owner.Contract.Auth(&_Owner.CallOpts)
}

// Owners is a free data retrieval call binding the contract method 0x022914a7.
//
// Solidity: function owners(address ) view returns(bool)
func (_Owner *OwnerCaller) Owners(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Owner.contract.Call(opts, &out, "owners", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Owners is a free data retrieval call binding the contract method 0x022914a7.
//
// Solidity: function owners(address ) view returns(bool)
func (_Owner *OwnerSession) Owners(arg0 common.Address) (bool, error) {
	return _Owner.Contract.Owners(&_Owner.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x022914a7.
//
// Solidity: function owners(address ) view returns(bool)
func (_Owner *OwnerCallerSession) Owners(arg0 common.Address) (bool, error) {
	return _Owner.Contract.Owners(&_Owner.CallOpts, arg0)
}

// Add is a paid mutator transaction binding the contract method 0x4bf1b134.
//
// Solidity: function add(address _a, bool isSet, bytes[5] signs) returns()
func (_Owner *OwnerTransactor) Add(opts *bind.TransactOpts, _a common.Address, isSet bool, signs [5][]byte) (*types.Transaction, error) {
	return _Owner.contract.Transact(opts, "add", _a, isSet, signs)
}

// Add is a paid mutator transaction binding the contract method 0x4bf1b134.
//
// Solidity: function add(address _a, bool isSet, bytes[5] signs) returns()
func (_Owner *OwnerSession) Add(_a common.Address, isSet bool, signs [5][]byte) (*types.Transaction, error) {
	return _Owner.Contract.Add(&_Owner.TransactOpts, _a, isSet, signs)
}

// Add is a paid mutator transaction binding the contract method 0x4bf1b134.
//
// Solidity: function add(address _a, bool isSet, bytes[5] signs) returns()
func (_Owner *OwnerTransactorSession) Add(_a common.Address, isSet bool, signs [5][]byte) (*types.Transaction, error) {
	return _Owner.Contract.Add(&_Owner.TransactOpts, _a, isSet, signs)
}

// OwnerAddOwnerIterator is returned from FilterAddOwner and is used to iterate over the raw logs and unpacked data for AddOwner events raised by the Owner contract.
type OwnerAddOwnerIterator struct {
	Event *OwnerAddOwner // Event containing the contract specifics and raw log

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
func (it *OwnerAddOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnerAddOwner)
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
		it.Event = new(OwnerAddOwner)
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
func (it *OwnerAddOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnerAddOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnerAddOwner represents a AddOwner event raised by the Owner contract.
type OwnerAddOwner struct {
	From  common.Address
	IsSet bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAddOwner is a free log retrieval operation binding the contract event 0x938b2a24c98e4e542127ffa74a91e48942c2bddccae3b6d75f82bfda7bbc0807.
//
// Solidity: event AddOwner(address from, bool isSet)
func (_Owner *OwnerFilterer) FilterAddOwner(opts *bind.FilterOpts) (*OwnerAddOwnerIterator, error) {

	logs, sub, err := _Owner.contract.FilterLogs(opts, "AddOwner")
	if err != nil {
		return nil, err
	}
	return &OwnerAddOwnerIterator{contract: _Owner.contract, event: "AddOwner", logs: logs, sub: sub}, nil
}

// WatchAddOwner is a free log subscription operation binding the contract event 0x938b2a24c98e4e542127ffa74a91e48942c2bddccae3b6d75f82bfda7bbc0807.
//
// Solidity: event AddOwner(address from, bool isSet)
func (_Owner *OwnerFilterer) WatchAddOwner(opts *bind.WatchOpts, sink chan<- *OwnerAddOwner) (event.Subscription, error) {

	logs, sub, err := _Owner.contract.WatchLogs(opts, "AddOwner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnerAddOwner)
				if err := _Owner.contract.UnpackLog(event, "AddOwner", log); err != nil {
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

// ParseAddOwner is a log parse operation binding the contract event 0x938b2a24c98e4e542127ffa74a91e48942c2bddccae3b6d75f82bfda7bbc0807.
//
// Solidity: event AddOwner(address from, bool isSet)
func (_Owner *OwnerFilterer) ParseAddOwner(log types.Log) (*OwnerAddOwner, error) {
	event := new(OwnerAddOwner)
	if err := _Owner.contract.UnpackLog(event, "AddOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
