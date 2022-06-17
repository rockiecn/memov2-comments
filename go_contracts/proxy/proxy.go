// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package proxy

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

// OrderIn is an auto generated low-level Go binding around an user-defined struct.
type OrderIn struct {
	UIndex uint64
	PIndex uint64
	Start  uint64
	End    uint64
	Size   uint64
	Nonce  uint64
	TIndex uint8
	Sprice *big.Int
}

// PWIn is an auto generated low-level Go binding around an user-defined struct.
type PWIn struct {
	PIndex uint64
	TIndex uint8
	Pay    *big.Int
	Lost   *big.Int
}

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

// IControl1ABI is the input ABI used to generate the binding from.
const IControl1ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"i\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"}],\"name\":\"Pledge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"i\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"ti\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lock\",\"type\":\"uint256\"}],\"name\":\"Unpledge\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"activate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_t\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_mint\",\"type\":\"bool\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"addT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"}],\"name\":\"addToGroup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_p\",\"type\":\"address\"}],\"name\":\"alterPayee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"_ban\",\"type\":\"bool\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"ban\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"_ban\",\"type\":\"bool\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"banG\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_t\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_ban\",\"type\":\"bool\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"banT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"}],\"name\":\"confirmPayee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_level\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_mr\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_dc\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_pc\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_kr\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_pr\",\"type\":\"uint256\"}],\"name\":\"createGroup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"_money\",\"type\":\"uint256\"}],\"name\":\"pledge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"}],\"name\":\"reAcc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_rtype\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_extra\",\"type\":\"bytes\"}],\"name\":\"reRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_money\",\"type\":\"uint256\"}],\"name\":\"unpledge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IControl1FuncSigs maps the 4-byte function signature to its string representation.
var IControl1FuncSigs = map[string]string{
	"843d1b49": "activate(uint64,bytes[5])",
	"68612348": "addT(address,bool,bytes[5])",
	"d9637231": "addToGroup(address,uint64)",
	"457a80aa": "alterPayee(address,address)",
	"df897f4f": "ban(uint64,bool,bytes[5])",
	"3e54a8b8": "banG(uint64,bool,bytes[5])",
	"7a743984": "banT(address,bool,bytes[5])",
	"51c10b28": "confirmPayee(address,uint64)",
	"832ec245": "createGroup(uint8,uint8,uint8,uint8,uint256,uint256)",
	"1a22e62e": "pledge(address,uint64,uint256)",
	"effcafce": "reAcc(address)",
	"9ccfe4ff": "reRole(address,uint8,bytes)",
	"0a866aed": "unpledge(address,uint64,uint8,uint256)",
}

// IControl1 is an auto generated Go binding around an Ethereum contract.
type IControl1 struct {
	IControl1Caller     // Read-only binding to the contract
	IControl1Transactor // Write-only binding to the contract
	IControl1Filterer   // Log filterer for contract events
}

// IControl1Caller is an auto generated read-only Go binding around an Ethereum contract.
type IControl1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IControl1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IControl1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IControl1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IControl1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IControl1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IControl1Session struct {
	Contract     *IControl1        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IControl1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IControl1CallerSession struct {
	Contract *IControl1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IControl1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IControl1TransactorSession struct {
	Contract     *IControl1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IControl1Raw is an auto generated low-level Go binding around an Ethereum contract.
type IControl1Raw struct {
	Contract *IControl1 // Generic contract binding to access the raw methods on
}

// IControl1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IControl1CallerRaw struct {
	Contract *IControl1Caller // Generic read-only contract binding to access the raw methods on
}

// IControl1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IControl1TransactorRaw struct {
	Contract *IControl1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIControl1 creates a new instance of IControl1, bound to a specific deployed contract.
func NewIControl1(address common.Address, backend bind.ContractBackend) (*IControl1, error) {
	contract, err := bindIControl1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IControl1{IControl1Caller: IControl1Caller{contract: contract}, IControl1Transactor: IControl1Transactor{contract: contract}, IControl1Filterer: IControl1Filterer{contract: contract}}, nil
}

// NewIControl1Caller creates a new read-only instance of IControl1, bound to a specific deployed contract.
func NewIControl1Caller(address common.Address, caller bind.ContractCaller) (*IControl1Caller, error) {
	contract, err := bindIControl1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IControl1Caller{contract: contract}, nil
}

// NewIControl1Transactor creates a new write-only instance of IControl1, bound to a specific deployed contract.
func NewIControl1Transactor(address common.Address, transactor bind.ContractTransactor) (*IControl1Transactor, error) {
	contract, err := bindIControl1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IControl1Transactor{contract: contract}, nil
}

// NewIControl1Filterer creates a new log filterer instance of IControl1, bound to a specific deployed contract.
func NewIControl1Filterer(address common.Address, filterer bind.ContractFilterer) (*IControl1Filterer, error) {
	contract, err := bindIControl1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IControl1Filterer{contract: contract}, nil
}

// bindIControl1 binds a generic wrapper to an already deployed contract.
func bindIControl1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IControl1ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IControl1 *IControl1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IControl1.Contract.IControl1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IControl1 *IControl1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IControl1.Contract.IControl1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IControl1 *IControl1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IControl1.Contract.IControl1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IControl1 *IControl1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IControl1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IControl1 *IControl1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IControl1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IControl1 *IControl1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IControl1.Contract.contract.Transact(opts, method, params...)
}

// Activate is a paid mutator transaction binding the contract method 0x843d1b49.
//
// Solidity: function activate(uint64 _i, bytes[5] signs) returns()
func (_IControl1 *IControl1Transactor) Activate(opts *bind.TransactOpts, _i uint64, signs [5][]byte) (*types.Transaction, error) {
	return _IControl1.contract.Transact(opts, "activate", _i, signs)
}

// Activate is a paid mutator transaction binding the contract method 0x843d1b49.
//
// Solidity: function activate(uint64 _i, bytes[5] signs) returns()
func (_IControl1 *IControl1Session) Activate(_i uint64, signs [5][]byte) (*types.Transaction, error) {
	return _IControl1.Contract.Activate(&_IControl1.TransactOpts, _i, signs)
}

// Activate is a paid mutator transaction binding the contract method 0x843d1b49.
//
// Solidity: function activate(uint64 _i, bytes[5] signs) returns()
func (_IControl1 *IControl1TransactorSession) Activate(_i uint64, signs [5][]byte) (*types.Transaction, error) {
	return _IControl1.Contract.Activate(&_IControl1.TransactOpts, _i, signs)
}

// AddT is a paid mutator transaction binding the contract method 0x68612348.
//
// Solidity: function addT(address _t, bool _mint, bytes[5] signs) returns()
func (_IControl1 *IControl1Transactor) AddT(opts *bind.TransactOpts, _t common.Address, _mint bool, signs [5][]byte) (*types.Transaction, error) {
	return _IControl1.contract.Transact(opts, "addT", _t, _mint, signs)
}

// AddT is a paid mutator transaction binding the contract method 0x68612348.
//
// Solidity: function addT(address _t, bool _mint, bytes[5] signs) returns()
func (_IControl1 *IControl1Session) AddT(_t common.Address, _mint bool, signs [5][]byte) (*types.Transaction, error) {
	return _IControl1.Contract.AddT(&_IControl1.TransactOpts, _t, _mint, signs)
}

// AddT is a paid mutator transaction binding the contract method 0x68612348.
//
// Solidity: function addT(address _t, bool _mint, bytes[5] signs) returns()
func (_IControl1 *IControl1TransactorSession) AddT(_t common.Address, _mint bool, signs [5][]byte) (*types.Transaction, error) {
	return _IControl1.Contract.AddT(&_IControl1.TransactOpts, _t, _mint, signs)
}

// AddToGroup is a paid mutator transaction binding the contract method 0xd9637231.
//
// Solidity: function addToGroup(address _a, uint64 _gi) returns()
func (_IControl1 *IControl1Transactor) AddToGroup(opts *bind.TransactOpts, _a common.Address, _gi uint64) (*types.Transaction, error) {
	return _IControl1.contract.Transact(opts, "addToGroup", _a, _gi)
}

// AddToGroup is a paid mutator transaction binding the contract method 0xd9637231.
//
// Solidity: function addToGroup(address _a, uint64 _gi) returns()
func (_IControl1 *IControl1Session) AddToGroup(_a common.Address, _gi uint64) (*types.Transaction, error) {
	return _IControl1.Contract.AddToGroup(&_IControl1.TransactOpts, _a, _gi)
}

// AddToGroup is a paid mutator transaction binding the contract method 0xd9637231.
//
// Solidity: function addToGroup(address _a, uint64 _gi) returns()
func (_IControl1 *IControl1TransactorSession) AddToGroup(_a common.Address, _gi uint64) (*types.Transaction, error) {
	return _IControl1.Contract.AddToGroup(&_IControl1.TransactOpts, _a, _gi)
}

// AlterPayee is a paid mutator transaction binding the contract method 0x457a80aa.
//
// Solidity: function alterPayee(address _a, address _p) returns()
func (_IControl1 *IControl1Transactor) AlterPayee(opts *bind.TransactOpts, _a common.Address, _p common.Address) (*types.Transaction, error) {
	return _IControl1.contract.Transact(opts, "alterPayee", _a, _p)
}

// AlterPayee is a paid mutator transaction binding the contract method 0x457a80aa.
//
// Solidity: function alterPayee(address _a, address _p) returns()
func (_IControl1 *IControl1Session) AlterPayee(_a common.Address, _p common.Address) (*types.Transaction, error) {
	return _IControl1.Contract.AlterPayee(&_IControl1.TransactOpts, _a, _p)
}

// AlterPayee is a paid mutator transaction binding the contract method 0x457a80aa.
//
// Solidity: function alterPayee(address _a, address _p) returns()
func (_IControl1 *IControl1TransactorSession) AlterPayee(_a common.Address, _p common.Address) (*types.Transaction, error) {
	return _IControl1.Contract.AlterPayee(&_IControl1.TransactOpts, _a, _p)
}

// Ban is a paid mutator transaction binding the contract method 0xdf897f4f.
//
// Solidity: function ban(uint64 _i, bool _ban, bytes[5] signs) returns()
func (_IControl1 *IControl1Transactor) Ban(opts *bind.TransactOpts, _i uint64, _ban bool, signs [5][]byte) (*types.Transaction, error) {
	return _IControl1.contract.Transact(opts, "ban", _i, _ban, signs)
}

// Ban is a paid mutator transaction binding the contract method 0xdf897f4f.
//
// Solidity: function ban(uint64 _i, bool _ban, bytes[5] signs) returns()
func (_IControl1 *IControl1Session) Ban(_i uint64, _ban bool, signs [5][]byte) (*types.Transaction, error) {
	return _IControl1.Contract.Ban(&_IControl1.TransactOpts, _i, _ban, signs)
}

// Ban is a paid mutator transaction binding the contract method 0xdf897f4f.
//
// Solidity: function ban(uint64 _i, bool _ban, bytes[5] signs) returns()
func (_IControl1 *IControl1TransactorSession) Ban(_i uint64, _ban bool, signs [5][]byte) (*types.Transaction, error) {
	return _IControl1.Contract.Ban(&_IControl1.TransactOpts, _i, _ban, signs)
}

// BanG is a paid mutator transaction binding the contract method 0x3e54a8b8.
//
// Solidity: function banG(uint64 _gi, bool _ban, bytes[5] signs) returns()
func (_IControl1 *IControl1Transactor) BanG(opts *bind.TransactOpts, _gi uint64, _ban bool, signs [5][]byte) (*types.Transaction, error) {
	return _IControl1.contract.Transact(opts, "banG", _gi, _ban, signs)
}

// BanG is a paid mutator transaction binding the contract method 0x3e54a8b8.
//
// Solidity: function banG(uint64 _gi, bool _ban, bytes[5] signs) returns()
func (_IControl1 *IControl1Session) BanG(_gi uint64, _ban bool, signs [5][]byte) (*types.Transaction, error) {
	return _IControl1.Contract.BanG(&_IControl1.TransactOpts, _gi, _ban, signs)
}

// BanG is a paid mutator transaction binding the contract method 0x3e54a8b8.
//
// Solidity: function banG(uint64 _gi, bool _ban, bytes[5] signs) returns()
func (_IControl1 *IControl1TransactorSession) BanG(_gi uint64, _ban bool, signs [5][]byte) (*types.Transaction, error) {
	return _IControl1.Contract.BanG(&_IControl1.TransactOpts, _gi, _ban, signs)
}

// BanT is a paid mutator transaction binding the contract method 0x7a743984.
//
// Solidity: function banT(address _t, bool _ban, bytes[5] signs) returns()
func (_IControl1 *IControl1Transactor) BanT(opts *bind.TransactOpts, _t common.Address, _ban bool, signs [5][]byte) (*types.Transaction, error) {
	return _IControl1.contract.Transact(opts, "banT", _t, _ban, signs)
}

// BanT is a paid mutator transaction binding the contract method 0x7a743984.
//
// Solidity: function banT(address _t, bool _ban, bytes[5] signs) returns()
func (_IControl1 *IControl1Session) BanT(_t common.Address, _ban bool, signs [5][]byte) (*types.Transaction, error) {
	return _IControl1.Contract.BanT(&_IControl1.TransactOpts, _t, _ban, signs)
}

// BanT is a paid mutator transaction binding the contract method 0x7a743984.
//
// Solidity: function banT(address _t, bool _ban, bytes[5] signs) returns()
func (_IControl1 *IControl1TransactorSession) BanT(_t common.Address, _ban bool, signs [5][]byte) (*types.Transaction, error) {
	return _IControl1.Contract.BanT(&_IControl1.TransactOpts, _t, _ban, signs)
}

// ConfirmPayee is a paid mutator transaction binding the contract method 0x51c10b28.
//
// Solidity: function confirmPayee(address _a, uint64 _i) returns()
func (_IControl1 *IControl1Transactor) ConfirmPayee(opts *bind.TransactOpts, _a common.Address, _i uint64) (*types.Transaction, error) {
	return _IControl1.contract.Transact(opts, "confirmPayee", _a, _i)
}

// ConfirmPayee is a paid mutator transaction binding the contract method 0x51c10b28.
//
// Solidity: function confirmPayee(address _a, uint64 _i) returns()
func (_IControl1 *IControl1Session) ConfirmPayee(_a common.Address, _i uint64) (*types.Transaction, error) {
	return _IControl1.Contract.ConfirmPayee(&_IControl1.TransactOpts, _a, _i)
}

// ConfirmPayee is a paid mutator transaction binding the contract method 0x51c10b28.
//
// Solidity: function confirmPayee(address _a, uint64 _i) returns()
func (_IControl1 *IControl1TransactorSession) ConfirmPayee(_a common.Address, _i uint64) (*types.Transaction, error) {
	return _IControl1.Contract.ConfirmPayee(&_IControl1.TransactOpts, _a, _i)
}

// CreateGroup is a paid mutator transaction binding the contract method 0x832ec245.
//
// Solidity: function createGroup(uint8 _level, uint8 _mr, uint8 _dc, uint8 _pc, uint256 _kr, uint256 _pr) returns()
func (_IControl1 *IControl1Transactor) CreateGroup(opts *bind.TransactOpts, _level uint8, _mr uint8, _dc uint8, _pc uint8, _kr *big.Int, _pr *big.Int) (*types.Transaction, error) {
	return _IControl1.contract.Transact(opts, "createGroup", _level, _mr, _dc, _pc, _kr, _pr)
}

// CreateGroup is a paid mutator transaction binding the contract method 0x832ec245.
//
// Solidity: function createGroup(uint8 _level, uint8 _mr, uint8 _dc, uint8 _pc, uint256 _kr, uint256 _pr) returns()
func (_IControl1 *IControl1Session) CreateGroup(_level uint8, _mr uint8, _dc uint8, _pc uint8, _kr *big.Int, _pr *big.Int) (*types.Transaction, error) {
	return _IControl1.Contract.CreateGroup(&_IControl1.TransactOpts, _level, _mr, _dc, _pc, _kr, _pr)
}

// CreateGroup is a paid mutator transaction binding the contract method 0x832ec245.
//
// Solidity: function createGroup(uint8 _level, uint8 _mr, uint8 _dc, uint8 _pc, uint256 _kr, uint256 _pr) returns()
func (_IControl1 *IControl1TransactorSession) CreateGroup(_level uint8, _mr uint8, _dc uint8, _pc uint8, _kr *big.Int, _pr *big.Int) (*types.Transaction, error) {
	return _IControl1.Contract.CreateGroup(&_IControl1.TransactOpts, _level, _mr, _dc, _pc, _kr, _pr)
}

// Pledge is a paid mutator transaction binding the contract method 0x1a22e62e.
//
// Solidity: function pledge(address _a, uint64 _i, uint256 _money) returns()
func (_IControl1 *IControl1Transactor) Pledge(opts *bind.TransactOpts, _a common.Address, _i uint64, _money *big.Int) (*types.Transaction, error) {
	return _IControl1.contract.Transact(opts, "pledge", _a, _i, _money)
}

// Pledge is a paid mutator transaction binding the contract method 0x1a22e62e.
//
// Solidity: function pledge(address _a, uint64 _i, uint256 _money) returns()
func (_IControl1 *IControl1Session) Pledge(_a common.Address, _i uint64, _money *big.Int) (*types.Transaction, error) {
	return _IControl1.Contract.Pledge(&_IControl1.TransactOpts, _a, _i, _money)
}

// Pledge is a paid mutator transaction binding the contract method 0x1a22e62e.
//
// Solidity: function pledge(address _a, uint64 _i, uint256 _money) returns()
func (_IControl1 *IControl1TransactorSession) Pledge(_a common.Address, _i uint64, _money *big.Int) (*types.Transaction, error) {
	return _IControl1.Contract.Pledge(&_IControl1.TransactOpts, _a, _i, _money)
}

// ReAcc is a paid mutator transaction binding the contract method 0xeffcafce.
//
// Solidity: function reAcc(address _a) returns()
func (_IControl1 *IControl1Transactor) ReAcc(opts *bind.TransactOpts, _a common.Address) (*types.Transaction, error) {
	return _IControl1.contract.Transact(opts, "reAcc", _a)
}

// ReAcc is a paid mutator transaction binding the contract method 0xeffcafce.
//
// Solidity: function reAcc(address _a) returns()
func (_IControl1 *IControl1Session) ReAcc(_a common.Address) (*types.Transaction, error) {
	return _IControl1.Contract.ReAcc(&_IControl1.TransactOpts, _a)
}

// ReAcc is a paid mutator transaction binding the contract method 0xeffcafce.
//
// Solidity: function reAcc(address _a) returns()
func (_IControl1 *IControl1TransactorSession) ReAcc(_a common.Address) (*types.Transaction, error) {
	return _IControl1.Contract.ReAcc(&_IControl1.TransactOpts, _a)
}

// ReRole is a paid mutator transaction binding the contract method 0x9ccfe4ff.
//
// Solidity: function reRole(address _a, uint8 _rtype, bytes _extra) returns()
func (_IControl1 *IControl1Transactor) ReRole(opts *bind.TransactOpts, _a common.Address, _rtype uint8, _extra []byte) (*types.Transaction, error) {
	return _IControl1.contract.Transact(opts, "reRole", _a, _rtype, _extra)
}

// ReRole is a paid mutator transaction binding the contract method 0x9ccfe4ff.
//
// Solidity: function reRole(address _a, uint8 _rtype, bytes _extra) returns()
func (_IControl1 *IControl1Session) ReRole(_a common.Address, _rtype uint8, _extra []byte) (*types.Transaction, error) {
	return _IControl1.Contract.ReRole(&_IControl1.TransactOpts, _a, _rtype, _extra)
}

// ReRole is a paid mutator transaction binding the contract method 0x9ccfe4ff.
//
// Solidity: function reRole(address _a, uint8 _rtype, bytes _extra) returns()
func (_IControl1 *IControl1TransactorSession) ReRole(_a common.Address, _rtype uint8, _extra []byte) (*types.Transaction, error) {
	return _IControl1.Contract.ReRole(&_IControl1.TransactOpts, _a, _rtype, _extra)
}

// Unpledge is a paid mutator transaction binding the contract method 0x0a866aed.
//
// Solidity: function unpledge(address _a, uint64 _i, uint8 _ti, uint256 _money) returns()
func (_IControl1 *IControl1Transactor) Unpledge(opts *bind.TransactOpts, _a common.Address, _i uint64, _ti uint8, _money *big.Int) (*types.Transaction, error) {
	return _IControl1.contract.Transact(opts, "unpledge", _a, _i, _ti, _money)
}

// Unpledge is a paid mutator transaction binding the contract method 0x0a866aed.
//
// Solidity: function unpledge(address _a, uint64 _i, uint8 _ti, uint256 _money) returns()
func (_IControl1 *IControl1Session) Unpledge(_a common.Address, _i uint64, _ti uint8, _money *big.Int) (*types.Transaction, error) {
	return _IControl1.Contract.Unpledge(&_IControl1.TransactOpts, _a, _i, _ti, _money)
}

// Unpledge is a paid mutator transaction binding the contract method 0x0a866aed.
//
// Solidity: function unpledge(address _a, uint64 _i, uint8 _ti, uint256 _money) returns()
func (_IControl1 *IControl1TransactorSession) Unpledge(_a common.Address, _i uint64, _ti uint8, _money *big.Int) (*types.Transaction, error) {
	return _IControl1.Contract.Unpledge(&_IControl1.TransactOpts, _a, _i, _ti, _money)
}

// IControl1PledgeIterator is returned from FilterPledge and is used to iterate over the raw logs and unpacked data for Pledge events raised by the IControl1 contract.
type IControl1PledgeIterator struct {
	Event *IControl1Pledge // Event containing the contract specifics and raw log

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
func (it *IControl1PledgeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IControl1Pledge)
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
		it.Event = new(IControl1Pledge)
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
func (it *IControl1PledgeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IControl1PledgeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IControl1Pledge represents a Pledge event raised by the IControl1 contract.
type IControl1Pledge struct {
	Payer common.Address
	I     uint64
	Money *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPledge is a free log retrieval operation binding the contract event 0x09b6ee23f9f6fb394aa2974c56e3d3bef9d90a07cb6ed69bea1b1ca84a958ad9.
//
// Solidity: event Pledge(address indexed payer, uint64 indexed i, uint256 money)
func (_IControl1 *IControl1Filterer) FilterPledge(opts *bind.FilterOpts, payer []common.Address, i []uint64) (*IControl1PledgeIterator, error) {

	var payerRule []interface{}
	for _, payerItem := range payer {
		payerRule = append(payerRule, payerItem)
	}
	var iRule []interface{}
	for _, iItem := range i {
		iRule = append(iRule, iItem)
	}

	logs, sub, err := _IControl1.contract.FilterLogs(opts, "Pledge", payerRule, iRule)
	if err != nil {
		return nil, err
	}
	return &IControl1PledgeIterator{contract: _IControl1.contract, event: "Pledge", logs: logs, sub: sub}, nil
}

// WatchPledge is a free log subscription operation binding the contract event 0x09b6ee23f9f6fb394aa2974c56e3d3bef9d90a07cb6ed69bea1b1ca84a958ad9.
//
// Solidity: event Pledge(address indexed payer, uint64 indexed i, uint256 money)
func (_IControl1 *IControl1Filterer) WatchPledge(opts *bind.WatchOpts, sink chan<- *IControl1Pledge, payer []common.Address, i []uint64) (event.Subscription, error) {

	var payerRule []interface{}
	for _, payerItem := range payer {
		payerRule = append(payerRule, payerItem)
	}
	var iRule []interface{}
	for _, iItem := range i {
		iRule = append(iRule, iItem)
	}

	logs, sub, err := _IControl1.contract.WatchLogs(opts, "Pledge", payerRule, iRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IControl1Pledge)
				if err := _IControl1.contract.UnpackLog(event, "Pledge", log); err != nil {
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

// ParsePledge is a log parse operation binding the contract event 0x09b6ee23f9f6fb394aa2974c56e3d3bef9d90a07cb6ed69bea1b1ca84a958ad9.
//
// Solidity: event Pledge(address indexed payer, uint64 indexed i, uint256 money)
func (_IControl1 *IControl1Filterer) ParsePledge(log types.Log) (*IControl1Pledge, error) {
	event := new(IControl1Pledge)
	if err := _IControl1.contract.UnpackLog(event, "Pledge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IControl1UnpledgeIterator is returned from FilterUnpledge and is used to iterate over the raw logs and unpacked data for Unpledge events raised by the IControl1 contract.
type IControl1UnpledgeIterator struct {
	Event *IControl1Unpledge // Event containing the contract specifics and raw log

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
func (it *IControl1UnpledgeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IControl1Unpledge)
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
		it.Event = new(IControl1Unpledge)
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
func (it *IControl1UnpledgeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IControl1UnpledgeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IControl1Unpledge represents a Unpledge event raised by the IControl1 contract.
type IControl1Unpledge struct {
	Payee common.Address
	I     uint64
	Ti    uint8
	Money *big.Int
	Lock  *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUnpledge is a free log retrieval operation binding the contract event 0xbccb4889fc3c63ae9fc95055fbaa744519cd0a8117b6e6602bac77fdffe62456.
//
// Solidity: event Unpledge(address indexed payee, uint64 indexed i, uint8 indexed ti, uint256 money, uint256 lock)
func (_IControl1 *IControl1Filterer) FilterUnpledge(opts *bind.FilterOpts, payee []common.Address, i []uint64, ti []uint8) (*IControl1UnpledgeIterator, error) {

	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}
	var iRule []interface{}
	for _, iItem := range i {
		iRule = append(iRule, iItem)
	}
	var tiRule []interface{}
	for _, tiItem := range ti {
		tiRule = append(tiRule, tiItem)
	}

	logs, sub, err := _IControl1.contract.FilterLogs(opts, "Unpledge", payeeRule, iRule, tiRule)
	if err != nil {
		return nil, err
	}
	return &IControl1UnpledgeIterator{contract: _IControl1.contract, event: "Unpledge", logs: logs, sub: sub}, nil
}

// WatchUnpledge is a free log subscription operation binding the contract event 0xbccb4889fc3c63ae9fc95055fbaa744519cd0a8117b6e6602bac77fdffe62456.
//
// Solidity: event Unpledge(address indexed payee, uint64 indexed i, uint8 indexed ti, uint256 money, uint256 lock)
func (_IControl1 *IControl1Filterer) WatchUnpledge(opts *bind.WatchOpts, sink chan<- *IControl1Unpledge, payee []common.Address, i []uint64, ti []uint8) (event.Subscription, error) {

	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}
	var iRule []interface{}
	for _, iItem := range i {
		iRule = append(iRule, iItem)
	}
	var tiRule []interface{}
	for _, tiItem := range ti {
		tiRule = append(tiRule, tiItem)
	}

	logs, sub, err := _IControl1.contract.WatchLogs(opts, "Unpledge", payeeRule, iRule, tiRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IControl1Unpledge)
				if err := _IControl1.contract.UnpackLog(event, "Unpledge", log); err != nil {
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

// ParseUnpledge is a log parse operation binding the contract event 0xbccb4889fc3c63ae9fc95055fbaa744519cd0a8117b6e6602bac77fdffe62456.
//
// Solidity: event Unpledge(address indexed payee, uint64 indexed i, uint8 indexed ti, uint256 money, uint256 lock)
func (_IControl1 *IControl1Filterer) ParseUnpledge(log types.Log) (*IControl1Unpledge, error) {
	event := new(IControl1Unpledge)
	if err := _IControl1.contract.UnpackLog(event, "Unpledge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IControl2ABI is the input ABI used to generate the binding from.
const IControl2ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"pi\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"ti\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pay\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lost\",\"type\":\"uint256\"}],\"name\":\"ProWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"i\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"isLock\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"ti\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"}],\"name\":\"Recharge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"i\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"ti\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actualMoney\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"uIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"tIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"sprice\",\"type\":\"uint256\"}],\"internalType\":\"structOrderIn\",\"name\":\"_oi\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"uSign\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"pSign\",\"type\":\"bytes\"}],\"name\":\"addOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"uIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"tIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"sprice\",\"type\":\"uint256\"}],\"internalType\":\"structOrderIn\",\"name\":\"_oi\",\"type\":\"tuple\"},{\"internalType\":\"uint64[]\",\"name\":\"_kis\",\"type\":\"uint64[]\"},{\"internalType\":\"bytes[]\",\"name\":\"ksigns\",\"type\":\"bytes[]\"}],\"name\":\"addRepair\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"pIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"tIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"pay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lost\",\"type\":\"uint256\"}],\"internalType\":\"structPWIn\",\"name\":\"_ps\",\"type\":\"tuple\"},{\"internalType\":\"uint64[]\",\"name\":\"_kis\",\"type\":\"uint64[]\"},{\"internalType\":\"bytes[]\",\"name\":\"ksigns\",\"type\":\"bytes[]\"}],\"name\":\"proWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isLock\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_money\",\"type\":\"uint256\"}],\"name\":\"recharge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"uIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"tIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"sprice\",\"type\":\"uint256\"}],\"internalType\":\"structOrderIn\",\"name\":\"_oi\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"uSign\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"pSign\",\"type\":\"bytes\"}],\"name\":\"subOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_money\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IControl2FuncSigs maps the 4-byte function signature to its string representation.
var IControl2FuncSigs = map[string]string{
	"ae9d0b40": "addOrder(address,(uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256),bytes,bytes)",
	"af99c59a": "addRepair(address,(uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256),uint64[],bytes[])",
	"54aa6642": "proWithdraw(address,(uint64,uint8,uint256,uint256),uint64[],bytes[])",
	"f661f9e3": "recharge(address,uint64,uint8,bool,uint256)",
	"42f45166": "subOrder(address,(uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256),bytes,bytes)",
	"9b4c757a": "withdraw(address,uint64,uint8,uint256)",
}

// IControl2 is an auto generated Go binding around an Ethereum contract.
type IControl2 struct {
	IControl2Caller     // Read-only binding to the contract
	IControl2Transactor // Write-only binding to the contract
	IControl2Filterer   // Log filterer for contract events
}

// IControl2Caller is an auto generated read-only Go binding around an Ethereum contract.
type IControl2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IControl2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IControl2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IControl2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IControl2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IControl2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IControl2Session struct {
	Contract     *IControl2        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IControl2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IControl2CallerSession struct {
	Contract *IControl2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IControl2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IControl2TransactorSession struct {
	Contract     *IControl2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IControl2Raw is an auto generated low-level Go binding around an Ethereum contract.
type IControl2Raw struct {
	Contract *IControl2 // Generic contract binding to access the raw methods on
}

// IControl2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IControl2CallerRaw struct {
	Contract *IControl2Caller // Generic read-only contract binding to access the raw methods on
}

// IControl2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IControl2TransactorRaw struct {
	Contract *IControl2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIControl2 creates a new instance of IControl2, bound to a specific deployed contract.
func NewIControl2(address common.Address, backend bind.ContractBackend) (*IControl2, error) {
	contract, err := bindIControl2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IControl2{IControl2Caller: IControl2Caller{contract: contract}, IControl2Transactor: IControl2Transactor{contract: contract}, IControl2Filterer: IControl2Filterer{contract: contract}}, nil
}

// NewIControl2Caller creates a new read-only instance of IControl2, bound to a specific deployed contract.
func NewIControl2Caller(address common.Address, caller bind.ContractCaller) (*IControl2Caller, error) {
	contract, err := bindIControl2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IControl2Caller{contract: contract}, nil
}

// NewIControl2Transactor creates a new write-only instance of IControl2, bound to a specific deployed contract.
func NewIControl2Transactor(address common.Address, transactor bind.ContractTransactor) (*IControl2Transactor, error) {
	contract, err := bindIControl2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IControl2Transactor{contract: contract}, nil
}

// NewIControl2Filterer creates a new log filterer instance of IControl2, bound to a specific deployed contract.
func NewIControl2Filterer(address common.Address, filterer bind.ContractFilterer) (*IControl2Filterer, error) {
	contract, err := bindIControl2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IControl2Filterer{contract: contract}, nil
}

// bindIControl2 binds a generic wrapper to an already deployed contract.
func bindIControl2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IControl2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IControl2 *IControl2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IControl2.Contract.IControl2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IControl2 *IControl2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IControl2.Contract.IControl2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IControl2 *IControl2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IControl2.Contract.IControl2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IControl2 *IControl2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IControl2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IControl2 *IControl2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IControl2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IControl2 *IControl2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IControl2.Contract.contract.Transact(opts, method, params...)
}

// AddOrder is a paid mutator transaction binding the contract method 0xae9d0b40.
//
// Solidity: function addOrder(address _a, (uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) _oi, bytes uSign, bytes pSign) returns()
func (_IControl2 *IControl2Transactor) AddOrder(opts *bind.TransactOpts, _a common.Address, _oi OrderIn, uSign []byte, pSign []byte) (*types.Transaction, error) {
	return _IControl2.contract.Transact(opts, "addOrder", _a, _oi, uSign, pSign)
}

// AddOrder is a paid mutator transaction binding the contract method 0xae9d0b40.
//
// Solidity: function addOrder(address _a, (uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) _oi, bytes uSign, bytes pSign) returns()
func (_IControl2 *IControl2Session) AddOrder(_a common.Address, _oi OrderIn, uSign []byte, pSign []byte) (*types.Transaction, error) {
	return _IControl2.Contract.AddOrder(&_IControl2.TransactOpts, _a, _oi, uSign, pSign)
}

// AddOrder is a paid mutator transaction binding the contract method 0xae9d0b40.
//
// Solidity: function addOrder(address _a, (uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) _oi, bytes uSign, bytes pSign) returns()
func (_IControl2 *IControl2TransactorSession) AddOrder(_a common.Address, _oi OrderIn, uSign []byte, pSign []byte) (*types.Transaction, error) {
	return _IControl2.Contract.AddOrder(&_IControl2.TransactOpts, _a, _oi, uSign, pSign)
}

// AddRepair is a paid mutator transaction binding the contract method 0xaf99c59a.
//
// Solidity: function addRepair(address _a, (uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) _oi, uint64[] _kis, bytes[] ksigns) returns()
func (_IControl2 *IControl2Transactor) AddRepair(opts *bind.TransactOpts, _a common.Address, _oi OrderIn, _kis []uint64, ksigns [][]byte) (*types.Transaction, error) {
	return _IControl2.contract.Transact(opts, "addRepair", _a, _oi, _kis, ksigns)
}

// AddRepair is a paid mutator transaction binding the contract method 0xaf99c59a.
//
// Solidity: function addRepair(address _a, (uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) _oi, uint64[] _kis, bytes[] ksigns) returns()
func (_IControl2 *IControl2Session) AddRepair(_a common.Address, _oi OrderIn, _kis []uint64, ksigns [][]byte) (*types.Transaction, error) {
	return _IControl2.Contract.AddRepair(&_IControl2.TransactOpts, _a, _oi, _kis, ksigns)
}

// AddRepair is a paid mutator transaction binding the contract method 0xaf99c59a.
//
// Solidity: function addRepair(address _a, (uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) _oi, uint64[] _kis, bytes[] ksigns) returns()
func (_IControl2 *IControl2TransactorSession) AddRepair(_a common.Address, _oi OrderIn, _kis []uint64, ksigns [][]byte) (*types.Transaction, error) {
	return _IControl2.Contract.AddRepair(&_IControl2.TransactOpts, _a, _oi, _kis, ksigns)
}

// ProWithdraw is a paid mutator transaction binding the contract method 0x54aa6642.
//
// Solidity: function proWithdraw(address _a, (uint64,uint8,uint256,uint256) _ps, uint64[] _kis, bytes[] ksigns) returns()
func (_IControl2 *IControl2Transactor) ProWithdraw(opts *bind.TransactOpts, _a common.Address, _ps PWIn, _kis []uint64, ksigns [][]byte) (*types.Transaction, error) {
	return _IControl2.contract.Transact(opts, "proWithdraw", _a, _ps, _kis, ksigns)
}

// ProWithdraw is a paid mutator transaction binding the contract method 0x54aa6642.
//
// Solidity: function proWithdraw(address _a, (uint64,uint8,uint256,uint256) _ps, uint64[] _kis, bytes[] ksigns) returns()
func (_IControl2 *IControl2Session) ProWithdraw(_a common.Address, _ps PWIn, _kis []uint64, ksigns [][]byte) (*types.Transaction, error) {
	return _IControl2.Contract.ProWithdraw(&_IControl2.TransactOpts, _a, _ps, _kis, ksigns)
}

// ProWithdraw is a paid mutator transaction binding the contract method 0x54aa6642.
//
// Solidity: function proWithdraw(address _a, (uint64,uint8,uint256,uint256) _ps, uint64[] _kis, bytes[] ksigns) returns()
func (_IControl2 *IControl2TransactorSession) ProWithdraw(_a common.Address, _ps PWIn, _kis []uint64, ksigns [][]byte) (*types.Transaction, error) {
	return _IControl2.Contract.ProWithdraw(&_IControl2.TransactOpts, _a, _ps, _kis, ksigns)
}

// Recharge is a paid mutator transaction binding the contract method 0xf661f9e3.
//
// Solidity: function recharge(address _a, uint64 _i, uint8 _ti, bool isLock, uint256 _money) returns()
func (_IControl2 *IControl2Transactor) Recharge(opts *bind.TransactOpts, _a common.Address, _i uint64, _ti uint8, isLock bool, _money *big.Int) (*types.Transaction, error) {
	return _IControl2.contract.Transact(opts, "recharge", _a, _i, _ti, isLock, _money)
}

// Recharge is a paid mutator transaction binding the contract method 0xf661f9e3.
//
// Solidity: function recharge(address _a, uint64 _i, uint8 _ti, bool isLock, uint256 _money) returns()
func (_IControl2 *IControl2Session) Recharge(_a common.Address, _i uint64, _ti uint8, isLock bool, _money *big.Int) (*types.Transaction, error) {
	return _IControl2.Contract.Recharge(&_IControl2.TransactOpts, _a, _i, _ti, isLock, _money)
}

// Recharge is a paid mutator transaction binding the contract method 0xf661f9e3.
//
// Solidity: function recharge(address _a, uint64 _i, uint8 _ti, bool isLock, uint256 _money) returns()
func (_IControl2 *IControl2TransactorSession) Recharge(_a common.Address, _i uint64, _ti uint8, isLock bool, _money *big.Int) (*types.Transaction, error) {
	return _IControl2.Contract.Recharge(&_IControl2.TransactOpts, _a, _i, _ti, isLock, _money)
}

// SubOrder is a paid mutator transaction binding the contract method 0x42f45166.
//
// Solidity: function subOrder(address _a, (uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) _oi, bytes uSign, bytes pSign) returns()
func (_IControl2 *IControl2Transactor) SubOrder(opts *bind.TransactOpts, _a common.Address, _oi OrderIn, uSign []byte, pSign []byte) (*types.Transaction, error) {
	return _IControl2.contract.Transact(opts, "subOrder", _a, _oi, uSign, pSign)
}

// SubOrder is a paid mutator transaction binding the contract method 0x42f45166.
//
// Solidity: function subOrder(address _a, (uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) _oi, bytes uSign, bytes pSign) returns()
func (_IControl2 *IControl2Session) SubOrder(_a common.Address, _oi OrderIn, uSign []byte, pSign []byte) (*types.Transaction, error) {
	return _IControl2.Contract.SubOrder(&_IControl2.TransactOpts, _a, _oi, uSign, pSign)
}

// SubOrder is a paid mutator transaction binding the contract method 0x42f45166.
//
// Solidity: function subOrder(address _a, (uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) _oi, bytes uSign, bytes pSign) returns()
func (_IControl2 *IControl2TransactorSession) SubOrder(_a common.Address, _oi OrderIn, uSign []byte, pSign []byte) (*types.Transaction, error) {
	return _IControl2.Contract.SubOrder(&_IControl2.TransactOpts, _a, _oi, uSign, pSign)
}

// Withdraw is a paid mutator transaction binding the contract method 0x9b4c757a.
//
// Solidity: function withdraw(address _a, uint64 _i, uint8 _ti, uint256 _money) returns()
func (_IControl2 *IControl2Transactor) Withdraw(opts *bind.TransactOpts, _a common.Address, _i uint64, _ti uint8, _money *big.Int) (*types.Transaction, error) {
	return _IControl2.contract.Transact(opts, "withdraw", _a, _i, _ti, _money)
}

// Withdraw is a paid mutator transaction binding the contract method 0x9b4c757a.
//
// Solidity: function withdraw(address _a, uint64 _i, uint8 _ti, uint256 _money) returns()
func (_IControl2 *IControl2Session) Withdraw(_a common.Address, _i uint64, _ti uint8, _money *big.Int) (*types.Transaction, error) {
	return _IControl2.Contract.Withdraw(&_IControl2.TransactOpts, _a, _i, _ti, _money)
}

// Withdraw is a paid mutator transaction binding the contract method 0x9b4c757a.
//
// Solidity: function withdraw(address _a, uint64 _i, uint8 _ti, uint256 _money) returns()
func (_IControl2 *IControl2TransactorSession) Withdraw(_a common.Address, _i uint64, _ti uint8, _money *big.Int) (*types.Transaction, error) {
	return _IControl2.Contract.Withdraw(&_IControl2.TransactOpts, _a, _i, _ti, _money)
}

// IControl2ProWithdrawIterator is returned from FilterProWithdraw and is used to iterate over the raw logs and unpacked data for ProWithdraw events raised by the IControl2 contract.
type IControl2ProWithdrawIterator struct {
	Event *IControl2ProWithdraw // Event containing the contract specifics and raw log

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
func (it *IControl2ProWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IControl2ProWithdraw)
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
		it.Event = new(IControl2ProWithdraw)
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
func (it *IControl2ProWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IControl2ProWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IControl2ProWithdraw represents a ProWithdraw event raised by the IControl2 contract.
type IControl2ProWithdraw struct {
	A    common.Address
	Pi   uint64
	Ti   uint8
	Pay  *big.Int
	Lost *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterProWithdraw is a free log retrieval operation binding the contract event 0x3ea20fac35aae529e59d08cd5e2b0481108ccd275b50215084c05039245766db.
//
// Solidity: event ProWithdraw(address indexed a, uint64 indexed pi, uint8 ti, uint256 pay, uint256 lost)
func (_IControl2 *IControl2Filterer) FilterProWithdraw(opts *bind.FilterOpts, a []common.Address, pi []uint64) (*IControl2ProWithdrawIterator, error) {

	var aRule []interface{}
	for _, aItem := range a {
		aRule = append(aRule, aItem)
	}
	var piRule []interface{}
	for _, piItem := range pi {
		piRule = append(piRule, piItem)
	}

	logs, sub, err := _IControl2.contract.FilterLogs(opts, "ProWithdraw", aRule, piRule)
	if err != nil {
		return nil, err
	}
	return &IControl2ProWithdrawIterator{contract: _IControl2.contract, event: "ProWithdraw", logs: logs, sub: sub}, nil
}

// WatchProWithdraw is a free log subscription operation binding the contract event 0x3ea20fac35aae529e59d08cd5e2b0481108ccd275b50215084c05039245766db.
//
// Solidity: event ProWithdraw(address indexed a, uint64 indexed pi, uint8 ti, uint256 pay, uint256 lost)
func (_IControl2 *IControl2Filterer) WatchProWithdraw(opts *bind.WatchOpts, sink chan<- *IControl2ProWithdraw, a []common.Address, pi []uint64) (event.Subscription, error) {

	var aRule []interface{}
	for _, aItem := range a {
		aRule = append(aRule, aItem)
	}
	var piRule []interface{}
	for _, piItem := range pi {
		piRule = append(piRule, piItem)
	}

	logs, sub, err := _IControl2.contract.WatchLogs(opts, "ProWithdraw", aRule, piRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IControl2ProWithdraw)
				if err := _IControl2.contract.UnpackLog(event, "ProWithdraw", log); err != nil {
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

// ParseProWithdraw is a log parse operation binding the contract event 0x3ea20fac35aae529e59d08cd5e2b0481108ccd275b50215084c05039245766db.
//
// Solidity: event ProWithdraw(address indexed a, uint64 indexed pi, uint8 ti, uint256 pay, uint256 lost)
func (_IControl2 *IControl2Filterer) ParseProWithdraw(log types.Log) (*IControl2ProWithdraw, error) {
	event := new(IControl2ProWithdraw)
	if err := _IControl2.contract.UnpackLog(event, "ProWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IControl2RechargeIterator is returned from FilterRecharge and is used to iterate over the raw logs and unpacked data for Recharge events raised by the IControl2 contract.
type IControl2RechargeIterator struct {
	Event *IControl2Recharge // Event containing the contract specifics and raw log

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
func (it *IControl2RechargeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IControl2Recharge)
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
		it.Event = new(IControl2Recharge)
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
func (it *IControl2RechargeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IControl2RechargeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IControl2Recharge represents a Recharge event raised by the IControl2 contract.
type IControl2Recharge struct {
	Payer  common.Address
	I      uint64
	IsLock bool
	Ti     uint8
	Money  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRecharge is a free log retrieval operation binding the contract event 0x544832b39fc44fc05cca7aa2c9f113bada9a6316b096b6cefcb302e433454a17.
//
// Solidity: event Recharge(address indexed payer, uint64 indexed i, bool indexed isLock, uint8 ti, uint256 money)
func (_IControl2 *IControl2Filterer) FilterRecharge(opts *bind.FilterOpts, payer []common.Address, i []uint64, isLock []bool) (*IControl2RechargeIterator, error) {

	var payerRule []interface{}
	for _, payerItem := range payer {
		payerRule = append(payerRule, payerItem)
	}
	var iRule []interface{}
	for _, iItem := range i {
		iRule = append(iRule, iItem)
	}
	var isLockRule []interface{}
	for _, isLockItem := range isLock {
		isLockRule = append(isLockRule, isLockItem)
	}

	logs, sub, err := _IControl2.contract.FilterLogs(opts, "Recharge", payerRule, iRule, isLockRule)
	if err != nil {
		return nil, err
	}
	return &IControl2RechargeIterator{contract: _IControl2.contract, event: "Recharge", logs: logs, sub: sub}, nil
}

// WatchRecharge is a free log subscription operation binding the contract event 0x544832b39fc44fc05cca7aa2c9f113bada9a6316b096b6cefcb302e433454a17.
//
// Solidity: event Recharge(address indexed payer, uint64 indexed i, bool indexed isLock, uint8 ti, uint256 money)
func (_IControl2 *IControl2Filterer) WatchRecharge(opts *bind.WatchOpts, sink chan<- *IControl2Recharge, payer []common.Address, i []uint64, isLock []bool) (event.Subscription, error) {

	var payerRule []interface{}
	for _, payerItem := range payer {
		payerRule = append(payerRule, payerItem)
	}
	var iRule []interface{}
	for _, iItem := range i {
		iRule = append(iRule, iItem)
	}
	var isLockRule []interface{}
	for _, isLockItem := range isLock {
		isLockRule = append(isLockRule, isLockItem)
	}

	logs, sub, err := _IControl2.contract.WatchLogs(opts, "Recharge", payerRule, iRule, isLockRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IControl2Recharge)
				if err := _IControl2.contract.UnpackLog(event, "Recharge", log); err != nil {
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

// ParseRecharge is a log parse operation binding the contract event 0x544832b39fc44fc05cca7aa2c9f113bada9a6316b096b6cefcb302e433454a17.
//
// Solidity: event Recharge(address indexed payer, uint64 indexed i, bool indexed isLock, uint8 ti, uint256 money)
func (_IControl2 *IControl2Filterer) ParseRecharge(log types.Log) (*IControl2Recharge, error) {
	event := new(IControl2Recharge)
	if err := _IControl2.contract.UnpackLog(event, "Recharge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IControl2WithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the IControl2 contract.
type IControl2WithdrawIterator struct {
	Event *IControl2Withdraw // Event containing the contract specifics and raw log

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
func (it *IControl2WithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IControl2Withdraw)
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
		it.Event = new(IControl2Withdraw)
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
func (it *IControl2WithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IControl2WithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IControl2Withdraw represents a Withdraw event raised by the IControl2 contract.
type IControl2Withdraw struct {
	Addr        common.Address
	I           uint64
	Ti          uint8
	Money       *big.Int
	ActualMoney *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x8cf9f6665801a7a41feb765d72a4472b5f7d327ed18c77c6e81d2bb5622cd8c0.
//
// Solidity: event Withdraw(address indexed addr, uint64 indexed i, uint8 ti, uint256 money, uint256 actualMoney)
func (_IControl2 *IControl2Filterer) FilterWithdraw(opts *bind.FilterOpts, addr []common.Address, i []uint64) (*IControl2WithdrawIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var iRule []interface{}
	for _, iItem := range i {
		iRule = append(iRule, iItem)
	}

	logs, sub, err := _IControl2.contract.FilterLogs(opts, "Withdraw", addrRule, iRule)
	if err != nil {
		return nil, err
	}
	return &IControl2WithdrawIterator{contract: _IControl2.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x8cf9f6665801a7a41feb765d72a4472b5f7d327ed18c77c6e81d2bb5622cd8c0.
//
// Solidity: event Withdraw(address indexed addr, uint64 indexed i, uint8 ti, uint256 money, uint256 actualMoney)
func (_IControl2 *IControl2Filterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *IControl2Withdraw, addr []common.Address, i []uint64) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var iRule []interface{}
	for _, iItem := range i {
		iRule = append(iRule, iItem)
	}

	logs, sub, err := _IControl2.contract.WatchLogs(opts, "Withdraw", addrRule, iRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IControl2Withdraw)
				if err := _IControl2.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x8cf9f6665801a7a41feb765d72a4472b5f7d327ed18c77c6e81d2bb5622cd8c0.
//
// Solidity: event Withdraw(address indexed addr, uint64 indexed i, uint8 ti, uint256 money, uint256 actualMoney)
func (_IControl2 *IControl2Filterer) ParseWithdraw(log types.Log) (*IControl2Withdraw, error) {
	event := new(IControl2Withdraw)
	if err := _IControl2.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IFileSysSetterABI is the input ABI used to generate the binding from.
const IFileSysSetterABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"ti\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"ui\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"pi\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sprice\",\"type\":\"uint256\"}],\"name\":\"AddOrder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"ti\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"ui\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"pi\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sprice\",\"type\":\"uint256\"}],\"name\":\"AddRepair\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"ti\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"ui\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"pi\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sprice\",\"type\":\"uint256\"}],\"name\":\"SubOrder\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"uIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"tIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"sprice\",\"type\":\"uint256\"}],\"internalType\":\"structOrderIn\",\"name\":\"ps\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_mr\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"}],\"name\":\"addOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"uIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"tIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"sprice\",\"type\":\"uint256\"}],\"internalType\":\"structOrderIn\",\"name\":\"ps\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"}],\"name\":\"addRepair\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"pIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"tIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"pay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lost\",\"type\":\"uint256\"}],\"internalType\":\"structPWIn\",\"name\":\"ps\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"}],\"name\":\"proWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isLock\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"}],\"name\":\"recharge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"uIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"tIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"sprice\",\"type\":\"uint256\"}],\"internalType\":\"structOrderIn\",\"name\":\"ps\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"}],\"name\":\"subOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IFileSysSetterFuncSigs maps the 4-byte function signature to its string representation.
var IFileSysSetterFuncSigs = map[string]string{
	"89cdc743": "addOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256),uint256,uint64)",
	"80faaf88": "addRepair((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256),uint64)",
	"a4703e16": "proWithdraw((uint64,uint8,uint256,uint256),uint64)",
	"24d11d40": "recharge(uint64,uint8,bool,uint256)",
	"248d02a0": "subOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256),uint64)",
	"259c6d5e": "withdraw(uint64,uint8,uint256)",
}

// IFileSysSetter is an auto generated Go binding around an Ethereum contract.
type IFileSysSetter struct {
	IFileSysSetterCaller     // Read-only binding to the contract
	IFileSysSetterTransactor // Write-only binding to the contract
	IFileSysSetterFilterer   // Log filterer for contract events
}

// IFileSysSetterCaller is an auto generated read-only Go binding around an Ethereum contract.
type IFileSysSetterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFileSysSetterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IFileSysSetterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFileSysSetterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IFileSysSetterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFileSysSetterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IFileSysSetterSession struct {
	Contract     *IFileSysSetter   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IFileSysSetterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IFileSysSetterCallerSession struct {
	Contract *IFileSysSetterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IFileSysSetterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IFileSysSetterTransactorSession struct {
	Contract     *IFileSysSetterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IFileSysSetterRaw is an auto generated low-level Go binding around an Ethereum contract.
type IFileSysSetterRaw struct {
	Contract *IFileSysSetter // Generic contract binding to access the raw methods on
}

// IFileSysSetterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IFileSysSetterCallerRaw struct {
	Contract *IFileSysSetterCaller // Generic read-only contract binding to access the raw methods on
}

// IFileSysSetterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IFileSysSetterTransactorRaw struct {
	Contract *IFileSysSetterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIFileSysSetter creates a new instance of IFileSysSetter, bound to a specific deployed contract.
func NewIFileSysSetter(address common.Address, backend bind.ContractBackend) (*IFileSysSetter, error) {
	contract, err := bindIFileSysSetter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IFileSysSetter{IFileSysSetterCaller: IFileSysSetterCaller{contract: contract}, IFileSysSetterTransactor: IFileSysSetterTransactor{contract: contract}, IFileSysSetterFilterer: IFileSysSetterFilterer{contract: contract}}, nil
}

// NewIFileSysSetterCaller creates a new read-only instance of IFileSysSetter, bound to a specific deployed contract.
func NewIFileSysSetterCaller(address common.Address, caller bind.ContractCaller) (*IFileSysSetterCaller, error) {
	contract, err := bindIFileSysSetter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IFileSysSetterCaller{contract: contract}, nil
}

// NewIFileSysSetterTransactor creates a new write-only instance of IFileSysSetter, bound to a specific deployed contract.
func NewIFileSysSetterTransactor(address common.Address, transactor bind.ContractTransactor) (*IFileSysSetterTransactor, error) {
	contract, err := bindIFileSysSetter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IFileSysSetterTransactor{contract: contract}, nil
}

// NewIFileSysSetterFilterer creates a new log filterer instance of IFileSysSetter, bound to a specific deployed contract.
func NewIFileSysSetterFilterer(address common.Address, filterer bind.ContractFilterer) (*IFileSysSetterFilterer, error) {
	contract, err := bindIFileSysSetter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IFileSysSetterFilterer{contract: contract}, nil
}

// bindIFileSysSetter binds a generic wrapper to an already deployed contract.
func bindIFileSysSetter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IFileSysSetterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFileSysSetter *IFileSysSetterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IFileSysSetter.Contract.IFileSysSetterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFileSysSetter *IFileSysSetterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFileSysSetter.Contract.IFileSysSetterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFileSysSetter *IFileSysSetterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFileSysSetter.Contract.IFileSysSetterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFileSysSetter *IFileSysSetterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IFileSysSetter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFileSysSetter *IFileSysSetterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFileSysSetter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFileSysSetter *IFileSysSetterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFileSysSetter.Contract.contract.Transact(opts, method, params...)
}

// AddOrder is a paid mutator transaction binding the contract method 0x89cdc743.
//
// Solidity: function addOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) ps, uint256 _mr, uint64 _gi) returns()
func (_IFileSysSetter *IFileSysSetterTransactor) AddOrder(opts *bind.TransactOpts, ps OrderIn, _mr *big.Int, _gi uint64) (*types.Transaction, error) {
	return _IFileSysSetter.contract.Transact(opts, "addOrder", ps, _mr, _gi)
}

// AddOrder is a paid mutator transaction binding the contract method 0x89cdc743.
//
// Solidity: function addOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) ps, uint256 _mr, uint64 _gi) returns()
func (_IFileSysSetter *IFileSysSetterSession) AddOrder(ps OrderIn, _mr *big.Int, _gi uint64) (*types.Transaction, error) {
	return _IFileSysSetter.Contract.AddOrder(&_IFileSysSetter.TransactOpts, ps, _mr, _gi)
}

// AddOrder is a paid mutator transaction binding the contract method 0x89cdc743.
//
// Solidity: function addOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) ps, uint256 _mr, uint64 _gi) returns()
func (_IFileSysSetter *IFileSysSetterTransactorSession) AddOrder(ps OrderIn, _mr *big.Int, _gi uint64) (*types.Transaction, error) {
	return _IFileSysSetter.Contract.AddOrder(&_IFileSysSetter.TransactOpts, ps, _mr, _gi)
}

// AddRepair is a paid mutator transaction binding the contract method 0x80faaf88.
//
// Solidity: function addRepair((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) ps, uint64 _gi) returns()
func (_IFileSysSetter *IFileSysSetterTransactor) AddRepair(opts *bind.TransactOpts, ps OrderIn, _gi uint64) (*types.Transaction, error) {
	return _IFileSysSetter.contract.Transact(opts, "addRepair", ps, _gi)
}

// AddRepair is a paid mutator transaction binding the contract method 0x80faaf88.
//
// Solidity: function addRepair((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) ps, uint64 _gi) returns()
func (_IFileSysSetter *IFileSysSetterSession) AddRepair(ps OrderIn, _gi uint64) (*types.Transaction, error) {
	return _IFileSysSetter.Contract.AddRepair(&_IFileSysSetter.TransactOpts, ps, _gi)
}

// AddRepair is a paid mutator transaction binding the contract method 0x80faaf88.
//
// Solidity: function addRepair((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) ps, uint64 _gi) returns()
func (_IFileSysSetter *IFileSysSetterTransactorSession) AddRepair(ps OrderIn, _gi uint64) (*types.Transaction, error) {
	return _IFileSysSetter.Contract.AddRepair(&_IFileSysSetter.TransactOpts, ps, _gi)
}

// ProWithdraw is a paid mutator transaction binding the contract method 0xa4703e16.
//
// Solidity: function proWithdraw((uint64,uint8,uint256,uint256) ps, uint64 _gi) returns(uint256)
func (_IFileSysSetter *IFileSysSetterTransactor) ProWithdraw(opts *bind.TransactOpts, ps PWIn, _gi uint64) (*types.Transaction, error) {
	return _IFileSysSetter.contract.Transact(opts, "proWithdraw", ps, _gi)
}

// ProWithdraw is a paid mutator transaction binding the contract method 0xa4703e16.
//
// Solidity: function proWithdraw((uint64,uint8,uint256,uint256) ps, uint64 _gi) returns(uint256)
func (_IFileSysSetter *IFileSysSetterSession) ProWithdraw(ps PWIn, _gi uint64) (*types.Transaction, error) {
	return _IFileSysSetter.Contract.ProWithdraw(&_IFileSysSetter.TransactOpts, ps, _gi)
}

// ProWithdraw is a paid mutator transaction binding the contract method 0xa4703e16.
//
// Solidity: function proWithdraw((uint64,uint8,uint256,uint256) ps, uint64 _gi) returns(uint256)
func (_IFileSysSetter *IFileSysSetterTransactorSession) ProWithdraw(ps PWIn, _gi uint64) (*types.Transaction, error) {
	return _IFileSysSetter.Contract.ProWithdraw(&_IFileSysSetter.TransactOpts, ps, _gi)
}

// Recharge is a paid mutator transaction binding the contract method 0x24d11d40.
//
// Solidity: function recharge(uint64 _i, uint8 _ti, bool isLock, uint256 money) returns()
func (_IFileSysSetter *IFileSysSetterTransactor) Recharge(opts *bind.TransactOpts, _i uint64, _ti uint8, isLock bool, money *big.Int) (*types.Transaction, error) {
	return _IFileSysSetter.contract.Transact(opts, "recharge", _i, _ti, isLock, money)
}

// Recharge is a paid mutator transaction binding the contract method 0x24d11d40.
//
// Solidity: function recharge(uint64 _i, uint8 _ti, bool isLock, uint256 money) returns()
func (_IFileSysSetter *IFileSysSetterSession) Recharge(_i uint64, _ti uint8, isLock bool, money *big.Int) (*types.Transaction, error) {
	return _IFileSysSetter.Contract.Recharge(&_IFileSysSetter.TransactOpts, _i, _ti, isLock, money)
}

// Recharge is a paid mutator transaction binding the contract method 0x24d11d40.
//
// Solidity: function recharge(uint64 _i, uint8 _ti, bool isLock, uint256 money) returns()
func (_IFileSysSetter *IFileSysSetterTransactorSession) Recharge(_i uint64, _ti uint8, isLock bool, money *big.Int) (*types.Transaction, error) {
	return _IFileSysSetter.Contract.Recharge(&_IFileSysSetter.TransactOpts, _i, _ti, isLock, money)
}

// SubOrder is a paid mutator transaction binding the contract method 0x248d02a0.
//
// Solidity: function subOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) ps, uint64 _gi) returns()
func (_IFileSysSetter *IFileSysSetterTransactor) SubOrder(opts *bind.TransactOpts, ps OrderIn, _gi uint64) (*types.Transaction, error) {
	return _IFileSysSetter.contract.Transact(opts, "subOrder", ps, _gi)
}

// SubOrder is a paid mutator transaction binding the contract method 0x248d02a0.
//
// Solidity: function subOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) ps, uint64 _gi) returns()
func (_IFileSysSetter *IFileSysSetterSession) SubOrder(ps OrderIn, _gi uint64) (*types.Transaction, error) {
	return _IFileSysSetter.Contract.SubOrder(&_IFileSysSetter.TransactOpts, ps, _gi)
}

// SubOrder is a paid mutator transaction binding the contract method 0x248d02a0.
//
// Solidity: function subOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) ps, uint64 _gi) returns()
func (_IFileSysSetter *IFileSysSetterTransactorSession) SubOrder(ps OrderIn, _gi uint64) (*types.Transaction, error) {
	return _IFileSysSetter.Contract.SubOrder(&_IFileSysSetter.TransactOpts, ps, _gi)
}

// Withdraw is a paid mutator transaction binding the contract method 0x259c6d5e.
//
// Solidity: function withdraw(uint64 _i, uint8 _ti, uint256 money) returns(uint256)
func (_IFileSysSetter *IFileSysSetterTransactor) Withdraw(opts *bind.TransactOpts, _i uint64, _ti uint8, money *big.Int) (*types.Transaction, error) {
	return _IFileSysSetter.contract.Transact(opts, "withdraw", _i, _ti, money)
}

// Withdraw is a paid mutator transaction binding the contract method 0x259c6d5e.
//
// Solidity: function withdraw(uint64 _i, uint8 _ti, uint256 money) returns(uint256)
func (_IFileSysSetter *IFileSysSetterSession) Withdraw(_i uint64, _ti uint8, money *big.Int) (*types.Transaction, error) {
	return _IFileSysSetter.Contract.Withdraw(&_IFileSysSetter.TransactOpts, _i, _ti, money)
}

// Withdraw is a paid mutator transaction binding the contract method 0x259c6d5e.
//
// Solidity: function withdraw(uint64 _i, uint8 _ti, uint256 money) returns(uint256)
func (_IFileSysSetter *IFileSysSetterTransactorSession) Withdraw(_i uint64, _ti uint8, money *big.Int) (*types.Transaction, error) {
	return _IFileSysSetter.Contract.Withdraw(&_IFileSysSetter.TransactOpts, _i, _ti, money)
}

// IFileSysSetterAddOrderIterator is returned from FilterAddOrder and is used to iterate over the raw logs and unpacked data for AddOrder events raised by the IFileSysSetter contract.
type IFileSysSetterAddOrderIterator struct {
	Event *IFileSysSetterAddOrder // Event containing the contract specifics and raw log

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
func (it *IFileSysSetterAddOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IFileSysSetterAddOrder)
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
		it.Event = new(IFileSysSetterAddOrder)
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
func (it *IFileSysSetterAddOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IFileSysSetterAddOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IFileSysSetterAddOrder represents a AddOrder event raised by the IFileSysSetter contract.
type IFileSysSetterAddOrder struct {
	Ti     uint8
	Ui     uint64
	Pi     uint64
	Start  uint64
	End    uint64
	Size   uint64
	Sprice *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAddOrder is a free log retrieval operation binding the contract event 0x9ad187def591246eb8fff84534926005f62988636593745516c0ac08fa086314.
//
// Solidity: event AddOrder(uint8 indexed ti, uint64 indexed ui, uint64 indexed pi, uint64 start, uint64 end, uint64 size, uint256 sprice)
func (_IFileSysSetter *IFileSysSetterFilterer) FilterAddOrder(opts *bind.FilterOpts, ti []uint8, ui []uint64, pi []uint64) (*IFileSysSetterAddOrderIterator, error) {

	var tiRule []interface{}
	for _, tiItem := range ti {
		tiRule = append(tiRule, tiItem)
	}
	var uiRule []interface{}
	for _, uiItem := range ui {
		uiRule = append(uiRule, uiItem)
	}
	var piRule []interface{}
	for _, piItem := range pi {
		piRule = append(piRule, piItem)
	}

	logs, sub, err := _IFileSysSetter.contract.FilterLogs(opts, "AddOrder", tiRule, uiRule, piRule)
	if err != nil {
		return nil, err
	}
	return &IFileSysSetterAddOrderIterator{contract: _IFileSysSetter.contract, event: "AddOrder", logs: logs, sub: sub}, nil
}

// WatchAddOrder is a free log subscription operation binding the contract event 0x9ad187def591246eb8fff84534926005f62988636593745516c0ac08fa086314.
//
// Solidity: event AddOrder(uint8 indexed ti, uint64 indexed ui, uint64 indexed pi, uint64 start, uint64 end, uint64 size, uint256 sprice)
func (_IFileSysSetter *IFileSysSetterFilterer) WatchAddOrder(opts *bind.WatchOpts, sink chan<- *IFileSysSetterAddOrder, ti []uint8, ui []uint64, pi []uint64) (event.Subscription, error) {

	var tiRule []interface{}
	for _, tiItem := range ti {
		tiRule = append(tiRule, tiItem)
	}
	var uiRule []interface{}
	for _, uiItem := range ui {
		uiRule = append(uiRule, uiItem)
	}
	var piRule []interface{}
	for _, piItem := range pi {
		piRule = append(piRule, piItem)
	}

	logs, sub, err := _IFileSysSetter.contract.WatchLogs(opts, "AddOrder", tiRule, uiRule, piRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IFileSysSetterAddOrder)
				if err := _IFileSysSetter.contract.UnpackLog(event, "AddOrder", log); err != nil {
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

// ParseAddOrder is a log parse operation binding the contract event 0x9ad187def591246eb8fff84534926005f62988636593745516c0ac08fa086314.
//
// Solidity: event AddOrder(uint8 indexed ti, uint64 indexed ui, uint64 indexed pi, uint64 start, uint64 end, uint64 size, uint256 sprice)
func (_IFileSysSetter *IFileSysSetterFilterer) ParseAddOrder(log types.Log) (*IFileSysSetterAddOrder, error) {
	event := new(IFileSysSetterAddOrder)
	if err := _IFileSysSetter.contract.UnpackLog(event, "AddOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IFileSysSetterAddRepairIterator is returned from FilterAddRepair and is used to iterate over the raw logs and unpacked data for AddRepair events raised by the IFileSysSetter contract.
type IFileSysSetterAddRepairIterator struct {
	Event *IFileSysSetterAddRepair // Event containing the contract specifics and raw log

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
func (it *IFileSysSetterAddRepairIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IFileSysSetterAddRepair)
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
		it.Event = new(IFileSysSetterAddRepair)
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
func (it *IFileSysSetterAddRepairIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IFileSysSetterAddRepairIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IFileSysSetterAddRepair represents a AddRepair event raised by the IFileSysSetter contract.
type IFileSysSetterAddRepair struct {
	Ti     uint8
	Ui     uint64
	Pi     uint64
	Start  uint64
	End    uint64
	Size   uint64
	Sprice *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAddRepair is a free log retrieval operation binding the contract event 0x82f944cf0134663ff19eea526226f1cd8d5746a71cc837983659b8cbb2674e48.
//
// Solidity: event AddRepair(uint8 indexed ti, uint64 indexed ui, uint64 indexed pi, uint64 start, uint64 end, uint64 size, uint256 sprice)
func (_IFileSysSetter *IFileSysSetterFilterer) FilterAddRepair(opts *bind.FilterOpts, ti []uint8, ui []uint64, pi []uint64) (*IFileSysSetterAddRepairIterator, error) {

	var tiRule []interface{}
	for _, tiItem := range ti {
		tiRule = append(tiRule, tiItem)
	}
	var uiRule []interface{}
	for _, uiItem := range ui {
		uiRule = append(uiRule, uiItem)
	}
	var piRule []interface{}
	for _, piItem := range pi {
		piRule = append(piRule, piItem)
	}

	logs, sub, err := _IFileSysSetter.contract.FilterLogs(opts, "AddRepair", tiRule, uiRule, piRule)
	if err != nil {
		return nil, err
	}
	return &IFileSysSetterAddRepairIterator{contract: _IFileSysSetter.contract, event: "AddRepair", logs: logs, sub: sub}, nil
}

// WatchAddRepair is a free log subscription operation binding the contract event 0x82f944cf0134663ff19eea526226f1cd8d5746a71cc837983659b8cbb2674e48.
//
// Solidity: event AddRepair(uint8 indexed ti, uint64 indexed ui, uint64 indexed pi, uint64 start, uint64 end, uint64 size, uint256 sprice)
func (_IFileSysSetter *IFileSysSetterFilterer) WatchAddRepair(opts *bind.WatchOpts, sink chan<- *IFileSysSetterAddRepair, ti []uint8, ui []uint64, pi []uint64) (event.Subscription, error) {

	var tiRule []interface{}
	for _, tiItem := range ti {
		tiRule = append(tiRule, tiItem)
	}
	var uiRule []interface{}
	for _, uiItem := range ui {
		uiRule = append(uiRule, uiItem)
	}
	var piRule []interface{}
	for _, piItem := range pi {
		piRule = append(piRule, piItem)
	}

	logs, sub, err := _IFileSysSetter.contract.WatchLogs(opts, "AddRepair", tiRule, uiRule, piRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IFileSysSetterAddRepair)
				if err := _IFileSysSetter.contract.UnpackLog(event, "AddRepair", log); err != nil {
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

// ParseAddRepair is a log parse operation binding the contract event 0x82f944cf0134663ff19eea526226f1cd8d5746a71cc837983659b8cbb2674e48.
//
// Solidity: event AddRepair(uint8 indexed ti, uint64 indexed ui, uint64 indexed pi, uint64 start, uint64 end, uint64 size, uint256 sprice)
func (_IFileSysSetter *IFileSysSetterFilterer) ParseAddRepair(log types.Log) (*IFileSysSetterAddRepair, error) {
	event := new(IFileSysSetterAddRepair)
	if err := _IFileSysSetter.contract.UnpackLog(event, "AddRepair", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IFileSysSetterSubOrderIterator is returned from FilterSubOrder and is used to iterate over the raw logs and unpacked data for SubOrder events raised by the IFileSysSetter contract.
type IFileSysSetterSubOrderIterator struct {
	Event *IFileSysSetterSubOrder // Event containing the contract specifics and raw log

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
func (it *IFileSysSetterSubOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IFileSysSetterSubOrder)
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
		it.Event = new(IFileSysSetterSubOrder)
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
func (it *IFileSysSetterSubOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IFileSysSetterSubOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IFileSysSetterSubOrder represents a SubOrder event raised by the IFileSysSetter contract.
type IFileSysSetterSubOrder struct {
	Ti     uint8
	Ui     uint64
	Pi     uint64
	Start  uint64
	End    uint64
	Size   uint64
	Sprice *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSubOrder is a free log retrieval operation binding the contract event 0xe42ef64dc7dfe9121f13026a70df328cf6687348fdf8f71745f8e40fa5e52b82.
//
// Solidity: event SubOrder(uint8 indexed ti, uint64 indexed ui, uint64 indexed pi, uint64 start, uint64 end, uint64 size, uint256 sprice)
func (_IFileSysSetter *IFileSysSetterFilterer) FilterSubOrder(opts *bind.FilterOpts, ti []uint8, ui []uint64, pi []uint64) (*IFileSysSetterSubOrderIterator, error) {

	var tiRule []interface{}
	for _, tiItem := range ti {
		tiRule = append(tiRule, tiItem)
	}
	var uiRule []interface{}
	for _, uiItem := range ui {
		uiRule = append(uiRule, uiItem)
	}
	var piRule []interface{}
	for _, piItem := range pi {
		piRule = append(piRule, piItem)
	}

	logs, sub, err := _IFileSysSetter.contract.FilterLogs(opts, "SubOrder", tiRule, uiRule, piRule)
	if err != nil {
		return nil, err
	}
	return &IFileSysSetterSubOrderIterator{contract: _IFileSysSetter.contract, event: "SubOrder", logs: logs, sub: sub}, nil
}

// WatchSubOrder is a free log subscription operation binding the contract event 0xe42ef64dc7dfe9121f13026a70df328cf6687348fdf8f71745f8e40fa5e52b82.
//
// Solidity: event SubOrder(uint8 indexed ti, uint64 indexed ui, uint64 indexed pi, uint64 start, uint64 end, uint64 size, uint256 sprice)
func (_IFileSysSetter *IFileSysSetterFilterer) WatchSubOrder(opts *bind.WatchOpts, sink chan<- *IFileSysSetterSubOrder, ti []uint8, ui []uint64, pi []uint64) (event.Subscription, error) {

	var tiRule []interface{}
	for _, tiItem := range ti {
		tiRule = append(tiRule, tiItem)
	}
	var uiRule []interface{}
	for _, uiItem := range ui {
		uiRule = append(uiRule, uiItem)
	}
	var piRule []interface{}
	for _, piItem := range pi {
		piRule = append(piRule, piItem)
	}

	logs, sub, err := _IFileSysSetter.contract.WatchLogs(opts, "SubOrder", tiRule, uiRule, piRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IFileSysSetterSubOrder)
				if err := _IFileSysSetter.contract.UnpackLog(event, "SubOrder", log); err != nil {
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

// ParseSubOrder is a log parse operation binding the contract event 0xe42ef64dc7dfe9121f13026a70df328cf6687348fdf8f71745f8e40fa5e52b82.
//
// Solidity: event SubOrder(uint8 indexed ti, uint64 indexed ui, uint64 indexed pi, uint64 start, uint64 end, uint64 size, uint256 sprice)
func (_IFileSysSetter *IFileSysSetterFilterer) ParseSubOrder(log types.Log) (*IFileSysSetterSubOrder, error) {
	event := new(IFileSysSetterSubOrder)
	if err := _IFileSysSetter.contract.UnpackLog(event, "SubOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IInstanceABI is the input ABI used to generate the binding from.
const IInstanceABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"_type\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Alter\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_type\",\"type\":\"uint8\"}],\"name\":\"instances\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IInstanceFuncSigs maps the 4-byte function signature to its string representation.
var IInstanceFuncSigs = map[string]string{
	"3ec7d5b9": "instances(uint8)",
}

// IInstance is an auto generated Go binding around an Ethereum contract.
type IInstance struct {
	IInstanceCaller     // Read-only binding to the contract
	IInstanceTransactor // Write-only binding to the contract
	IInstanceFilterer   // Log filterer for contract events
}

// IInstanceCaller is an auto generated read-only Go binding around an Ethereum contract.
type IInstanceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInstanceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IInstanceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInstanceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IInstanceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInstanceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IInstanceSession struct {
	Contract     *IInstance        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IInstanceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IInstanceCallerSession struct {
	Contract *IInstanceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IInstanceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IInstanceTransactorSession struct {
	Contract     *IInstanceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IInstanceRaw is an auto generated low-level Go binding around an Ethereum contract.
type IInstanceRaw struct {
	Contract *IInstance // Generic contract binding to access the raw methods on
}

// IInstanceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IInstanceCallerRaw struct {
	Contract *IInstanceCaller // Generic read-only contract binding to access the raw methods on
}

// IInstanceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IInstanceTransactorRaw struct {
	Contract *IInstanceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIInstance creates a new instance of IInstance, bound to a specific deployed contract.
func NewIInstance(address common.Address, backend bind.ContractBackend) (*IInstance, error) {
	contract, err := bindIInstance(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IInstance{IInstanceCaller: IInstanceCaller{contract: contract}, IInstanceTransactor: IInstanceTransactor{contract: contract}, IInstanceFilterer: IInstanceFilterer{contract: contract}}, nil
}

// NewIInstanceCaller creates a new read-only instance of IInstance, bound to a specific deployed contract.
func NewIInstanceCaller(address common.Address, caller bind.ContractCaller) (*IInstanceCaller, error) {
	contract, err := bindIInstance(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IInstanceCaller{contract: contract}, nil
}

// NewIInstanceTransactor creates a new write-only instance of IInstance, bound to a specific deployed contract.
func NewIInstanceTransactor(address common.Address, transactor bind.ContractTransactor) (*IInstanceTransactor, error) {
	contract, err := bindIInstance(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IInstanceTransactor{contract: contract}, nil
}

// NewIInstanceFilterer creates a new log filterer instance of IInstance, bound to a specific deployed contract.
func NewIInstanceFilterer(address common.Address, filterer bind.ContractFilterer) (*IInstanceFilterer, error) {
	contract, err := bindIInstance(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IInstanceFilterer{contract: contract}, nil
}

// bindIInstance binds a generic wrapper to an already deployed contract.
func bindIInstance(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IInstanceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IInstance *IInstanceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IInstance.Contract.IInstanceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IInstance *IInstanceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IInstance.Contract.IInstanceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IInstance *IInstanceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IInstance.Contract.IInstanceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IInstance *IInstanceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IInstance.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IInstance *IInstanceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IInstance.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IInstance *IInstanceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IInstance.Contract.contract.Transact(opts, method, params...)
}

// Instances is a free data retrieval call binding the contract method 0x3ec7d5b9.
//
// Solidity: function instances(uint8 _type) view returns(address)
func (_IInstance *IInstanceCaller) Instances(opts *bind.CallOpts, _type uint8) (common.Address, error) {
	var out []interface{}
	err := _IInstance.contract.Call(opts, &out, "instances", _type)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Instances is a free data retrieval call binding the contract method 0x3ec7d5b9.
//
// Solidity: function instances(uint8 _type) view returns(address)
func (_IInstance *IInstanceSession) Instances(_type uint8) (common.Address, error) {
	return _IInstance.Contract.Instances(&_IInstance.CallOpts, _type)
}

// Instances is a free data retrieval call binding the contract method 0x3ec7d5b9.
//
// Solidity: function instances(uint8 _type) view returns(address)
func (_IInstance *IInstanceCallerSession) Instances(_type uint8) (common.Address, error) {
	return _IInstance.Contract.Instances(&_IInstance.CallOpts, _type)
}

// IInstanceAlterIterator is returned from FilterAlter and is used to iterate over the raw logs and unpacked data for Alter events raised by the IInstance contract.
type IInstanceAlterIterator struct {
	Event *IInstanceAlter // Event containing the contract specifics and raw log

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
func (it *IInstanceAlterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IInstanceAlter)
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
		it.Event = new(IInstanceAlter)
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
func (it *IInstanceAlterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IInstanceAlterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IInstanceAlter represents a Alter event raised by the IInstance contract.
type IInstanceAlter struct {
	Type uint8
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAlter is a free log retrieval operation binding the contract event 0x69da8aaa18d0d64ebdd4d982e4d32ac4aaab1fe0c1034b19846e51a61fcc0f02.
//
// Solidity: event Alter(uint8 _type, address from, address to)
func (_IInstance *IInstanceFilterer) FilterAlter(opts *bind.FilterOpts) (*IInstanceAlterIterator, error) {

	logs, sub, err := _IInstance.contract.FilterLogs(opts, "Alter")
	if err != nil {
		return nil, err
	}
	return &IInstanceAlterIterator{contract: _IInstance.contract, event: "Alter", logs: logs, sub: sub}, nil
}

// WatchAlter is a free log subscription operation binding the contract event 0x69da8aaa18d0d64ebdd4d982e4d32ac4aaab1fe0c1034b19846e51a61fcc0f02.
//
// Solidity: event Alter(uint8 _type, address from, address to)
func (_IInstance *IInstanceFilterer) WatchAlter(opts *bind.WatchOpts, sink chan<- *IInstanceAlter) (event.Subscription, error) {

	logs, sub, err := _IInstance.contract.WatchLogs(opts, "Alter")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IInstanceAlter)
				if err := _IInstance.contract.UnpackLog(event, "Alter", log); err != nil {
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

// ParseAlter is a log parse operation binding the contract event 0x69da8aaa18d0d64ebdd4d982e4d32ac4aaab1fe0c1034b19846e51a61fcc0f02.
//
// Solidity: event Alter(uint8 _type, address from, address to)
func (_IInstance *IInstanceFilterer) ParseAlter(log types.Log) (*IInstanceAlter, error) {
	event := new(IInstanceAlter)
	if err := _IInstance.contract.UnpackLog(event, "Alter", log); err != nil {
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

// IProxyABI is the input ABI used to generate the binding from.
const IProxyABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"i\",\"type\":\"uint64\"}],\"name\":\"Activate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"mint\",\"type\":\"bool\"}],\"name\":\"AddT\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"i\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"ban\",\"type\":\"bool\"}],\"name\":\"Ban\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"gi\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"ban\",\"type\":\"bool\"}],\"name\":\"BanG\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"ban\",\"type\":\"bool\"}],\"name\":\"BanT\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"activate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"uIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"tIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"sprice\",\"type\":\"uint256\"}],\"internalType\":\"structOrderIn\",\"name\":\"_oi\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"uSign\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"pSign\",\"type\":\"bytes\"}],\"name\":\"addOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"uIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"tIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"sprice\",\"type\":\"uint256\"}],\"internalType\":\"structOrderIn\",\"name\":\"_oi\",\"type\":\"tuple\"},{\"internalType\":\"uint64[]\",\"name\":\"_kis\",\"type\":\"uint64[]\"},{\"internalType\":\"bytes[]\",\"name\":\"ksigns\",\"type\":\"bytes[]\"}],\"name\":\"addRepair\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_t\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_mint\",\"type\":\"bool\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"addT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"}],\"name\":\"addToGroup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_p\",\"type\":\"address\"}],\"name\":\"alterPayee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"_ban\",\"type\":\"bool\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"ban\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"_ban\",\"type\":\"bool\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"banG\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_t\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_ban\",\"type\":\"bool\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"banT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"}],\"name\":\"confirmPayee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_level\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_mr\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_dc\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_pc\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_kr\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_pr\",\"type\":\"uint256\"}],\"name\":\"createGroup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"_money\",\"type\":\"uint256\"}],\"name\":\"pledge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"pIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"tIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"pay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lost\",\"type\":\"uint256\"}],\"internalType\":\"structPWIn\",\"name\":\"_ps\",\"type\":\"tuple\"},{\"internalType\":\"uint64[]\",\"name\":\"_kis\",\"type\":\"uint64[]\"},{\"internalType\":\"bytes[]\",\"name\":\"ksigns\",\"type\":\"bytes[]\"}],\"name\":\"proWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reAcc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_rtype\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_extra\",\"type\":\"bytes\"}],\"name\":\"reRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isLock\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_money\",\"type\":\"uint256\"}],\"name\":\"recharge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"uIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"tIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"sprice\",\"type\":\"uint256\"}],\"internalType\":\"structOrderIn\",\"name\":\"_oi\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"uSign\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"pSign\",\"type\":\"bytes\"}],\"name\":\"subOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_money\",\"type\":\"uint256\"}],\"name\":\"unpledge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_money\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IProxyFuncSigs maps the 4-byte function signature to its string representation.
var IProxyFuncSigs = map[string]string{
	"843d1b49": "activate(uint64,bytes[5])",
	"c05bdcd0": "addOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256),bytes,bytes)",
	"97f2352c": "addRepair((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256),uint64[],bytes[])",
	"68612348": "addT(address,bool,bytes[5])",
	"a246de42": "addToGroup(uint64)",
	"556e71c2": "alterPayee(address)",
	"df897f4f": "ban(uint64,bool,bytes[5])",
	"3e54a8b8": "banG(uint64,bool,bytes[5])",
	"7a743984": "banT(address,bool,bytes[5])",
	"310aa0be": "confirmPayee(uint64)",
	"832ec245": "createGroup(uint8,uint8,uint8,uint8,uint256,uint256)",
	"d23f7482": "pledge(uint64,uint256)",
	"e182d54c": "proWithdraw((uint64,uint8,uint256,uint256),uint64[],bytes[])",
	"6526250f": "reAcc()",
	"898e00f8": "reRole(uint8,bytes)",
	"24d11d40": "recharge(uint64,uint8,bool,uint256)",
	"153ede07": "subOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256),bytes,bytes)",
	"60985756": "unpledge(uint64,uint8,uint256)",
	"259c6d5e": "withdraw(uint64,uint8,uint256)",
}

// IProxy is an auto generated Go binding around an Ethereum contract.
type IProxy struct {
	IProxyCaller     // Read-only binding to the contract
	IProxyTransactor // Write-only binding to the contract
	IProxyFilterer   // Log filterer for contract events
}

// IProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type IProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IProxySession struct {
	Contract     *IProxy           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IProxyCallerSession struct {
	Contract *IProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IProxyTransactorSession struct {
	Contract     *IProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type IProxyRaw struct {
	Contract *IProxy // Generic contract binding to access the raw methods on
}

// IProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IProxyCallerRaw struct {
	Contract *IProxyCaller // Generic read-only contract binding to access the raw methods on
}

// IProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IProxyTransactorRaw struct {
	Contract *IProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIProxy creates a new instance of IProxy, bound to a specific deployed contract.
func NewIProxy(address common.Address, backend bind.ContractBackend) (*IProxy, error) {
	contract, err := bindIProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IProxy{IProxyCaller: IProxyCaller{contract: contract}, IProxyTransactor: IProxyTransactor{contract: contract}, IProxyFilterer: IProxyFilterer{contract: contract}}, nil
}

// NewIProxyCaller creates a new read-only instance of IProxy, bound to a specific deployed contract.
func NewIProxyCaller(address common.Address, caller bind.ContractCaller) (*IProxyCaller, error) {
	contract, err := bindIProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IProxyCaller{contract: contract}, nil
}

// NewIProxyTransactor creates a new write-only instance of IProxy, bound to a specific deployed contract.
func NewIProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*IProxyTransactor, error) {
	contract, err := bindIProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IProxyTransactor{contract: contract}, nil
}

// NewIProxyFilterer creates a new log filterer instance of IProxy, bound to a specific deployed contract.
func NewIProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*IProxyFilterer, error) {
	contract, err := bindIProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IProxyFilterer{contract: contract}, nil
}

// bindIProxy binds a generic wrapper to an already deployed contract.
func bindIProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IProxy *IProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IProxy.Contract.IProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IProxy *IProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IProxy.Contract.IProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IProxy *IProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IProxy.Contract.IProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IProxy *IProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IProxy *IProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IProxy *IProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IProxy.Contract.contract.Transact(opts, method, params...)
}

// Activate is a paid mutator transaction binding the contract method 0x843d1b49.
//
// Solidity: function activate(uint64 _i, bytes[5] signs) returns()
func (_IProxy *IProxyTransactor) Activate(opts *bind.TransactOpts, _i uint64, signs [5][]byte) (*types.Transaction, error) {
	return _IProxy.contract.Transact(opts, "activate", _i, signs)
}

// Activate is a paid mutator transaction binding the contract method 0x843d1b49.
//
// Solidity: function activate(uint64 _i, bytes[5] signs) returns()
func (_IProxy *IProxySession) Activate(_i uint64, signs [5][]byte) (*types.Transaction, error) {
	return _IProxy.Contract.Activate(&_IProxy.TransactOpts, _i, signs)
}

// Activate is a paid mutator transaction binding the contract method 0x843d1b49.
//
// Solidity: function activate(uint64 _i, bytes[5] signs) returns()
func (_IProxy *IProxyTransactorSession) Activate(_i uint64, signs [5][]byte) (*types.Transaction, error) {
	return _IProxy.Contract.Activate(&_IProxy.TransactOpts, _i, signs)
}

// AddOrder is a paid mutator transaction binding the contract method 0xc05bdcd0.
//
// Solidity: function addOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) _oi, bytes uSign, bytes pSign) returns()
func (_IProxy *IProxyTransactor) AddOrder(opts *bind.TransactOpts, _oi OrderIn, uSign []byte, pSign []byte) (*types.Transaction, error) {
	return _IProxy.contract.Transact(opts, "addOrder", _oi, uSign, pSign)
}

// AddOrder is a paid mutator transaction binding the contract method 0xc05bdcd0.
//
// Solidity: function addOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) _oi, bytes uSign, bytes pSign) returns()
func (_IProxy *IProxySession) AddOrder(_oi OrderIn, uSign []byte, pSign []byte) (*types.Transaction, error) {
	return _IProxy.Contract.AddOrder(&_IProxy.TransactOpts, _oi, uSign, pSign)
}

// AddOrder is a paid mutator transaction binding the contract method 0xc05bdcd0.
//
// Solidity: function addOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) _oi, bytes uSign, bytes pSign) returns()
func (_IProxy *IProxyTransactorSession) AddOrder(_oi OrderIn, uSign []byte, pSign []byte) (*types.Transaction, error) {
	return _IProxy.Contract.AddOrder(&_IProxy.TransactOpts, _oi, uSign, pSign)
}

// AddRepair is a paid mutator transaction binding the contract method 0x97f2352c.
//
// Solidity: function addRepair((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) _oi, uint64[] _kis, bytes[] ksigns) returns()
func (_IProxy *IProxyTransactor) AddRepair(opts *bind.TransactOpts, _oi OrderIn, _kis []uint64, ksigns [][]byte) (*types.Transaction, error) {
	return _IProxy.contract.Transact(opts, "addRepair", _oi, _kis, ksigns)
}

// AddRepair is a paid mutator transaction binding the contract method 0x97f2352c.
//
// Solidity: function addRepair((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) _oi, uint64[] _kis, bytes[] ksigns) returns()
func (_IProxy *IProxySession) AddRepair(_oi OrderIn, _kis []uint64, ksigns [][]byte) (*types.Transaction, error) {
	return _IProxy.Contract.AddRepair(&_IProxy.TransactOpts, _oi, _kis, ksigns)
}

// AddRepair is a paid mutator transaction binding the contract method 0x97f2352c.
//
// Solidity: function addRepair((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) _oi, uint64[] _kis, bytes[] ksigns) returns()
func (_IProxy *IProxyTransactorSession) AddRepair(_oi OrderIn, _kis []uint64, ksigns [][]byte) (*types.Transaction, error) {
	return _IProxy.Contract.AddRepair(&_IProxy.TransactOpts, _oi, _kis, ksigns)
}

// AddT is a paid mutator transaction binding the contract method 0x68612348.
//
// Solidity: function addT(address _t, bool _mint, bytes[5] signs) returns()
func (_IProxy *IProxyTransactor) AddT(opts *bind.TransactOpts, _t common.Address, _mint bool, signs [5][]byte) (*types.Transaction, error) {
	return _IProxy.contract.Transact(opts, "addT", _t, _mint, signs)
}

// AddT is a paid mutator transaction binding the contract method 0x68612348.
//
// Solidity: function addT(address _t, bool _mint, bytes[5] signs) returns()
func (_IProxy *IProxySession) AddT(_t common.Address, _mint bool, signs [5][]byte) (*types.Transaction, error) {
	return _IProxy.Contract.AddT(&_IProxy.TransactOpts, _t, _mint, signs)
}

// AddT is a paid mutator transaction binding the contract method 0x68612348.
//
// Solidity: function addT(address _t, bool _mint, bytes[5] signs) returns()
func (_IProxy *IProxyTransactorSession) AddT(_t common.Address, _mint bool, signs [5][]byte) (*types.Transaction, error) {
	return _IProxy.Contract.AddT(&_IProxy.TransactOpts, _t, _mint, signs)
}

// AddToGroup is a paid mutator transaction binding the contract method 0xa246de42.
//
// Solidity: function addToGroup(uint64 _gi) returns()
func (_IProxy *IProxyTransactor) AddToGroup(opts *bind.TransactOpts, _gi uint64) (*types.Transaction, error) {
	return _IProxy.contract.Transact(opts, "addToGroup", _gi)
}

// AddToGroup is a paid mutator transaction binding the contract method 0xa246de42.
//
// Solidity: function addToGroup(uint64 _gi) returns()
func (_IProxy *IProxySession) AddToGroup(_gi uint64) (*types.Transaction, error) {
	return _IProxy.Contract.AddToGroup(&_IProxy.TransactOpts, _gi)
}

// AddToGroup is a paid mutator transaction binding the contract method 0xa246de42.
//
// Solidity: function addToGroup(uint64 _gi) returns()
func (_IProxy *IProxyTransactorSession) AddToGroup(_gi uint64) (*types.Transaction, error) {
	return _IProxy.Contract.AddToGroup(&_IProxy.TransactOpts, _gi)
}

// AlterPayee is a paid mutator transaction binding the contract method 0x556e71c2.
//
// Solidity: function alterPayee(address _p) returns()
func (_IProxy *IProxyTransactor) AlterPayee(opts *bind.TransactOpts, _p common.Address) (*types.Transaction, error) {
	return _IProxy.contract.Transact(opts, "alterPayee", _p)
}

// AlterPayee is a paid mutator transaction binding the contract method 0x556e71c2.
//
// Solidity: function alterPayee(address _p) returns()
func (_IProxy *IProxySession) AlterPayee(_p common.Address) (*types.Transaction, error) {
	return _IProxy.Contract.AlterPayee(&_IProxy.TransactOpts, _p)
}

// AlterPayee is a paid mutator transaction binding the contract method 0x556e71c2.
//
// Solidity: function alterPayee(address _p) returns()
func (_IProxy *IProxyTransactorSession) AlterPayee(_p common.Address) (*types.Transaction, error) {
	return _IProxy.Contract.AlterPayee(&_IProxy.TransactOpts, _p)
}

// Ban is a paid mutator transaction binding the contract method 0xdf897f4f.
//
// Solidity: function ban(uint64 _i, bool _ban, bytes[5] signs) returns()
func (_IProxy *IProxyTransactor) Ban(opts *bind.TransactOpts, _i uint64, _ban bool, signs [5][]byte) (*types.Transaction, error) {
	return _IProxy.contract.Transact(opts, "ban", _i, _ban, signs)
}

// Ban is a paid mutator transaction binding the contract method 0xdf897f4f.
//
// Solidity: function ban(uint64 _i, bool _ban, bytes[5] signs) returns()
func (_IProxy *IProxySession) Ban(_i uint64, _ban bool, signs [5][]byte) (*types.Transaction, error) {
	return _IProxy.Contract.Ban(&_IProxy.TransactOpts, _i, _ban, signs)
}

// Ban is a paid mutator transaction binding the contract method 0xdf897f4f.
//
// Solidity: function ban(uint64 _i, bool _ban, bytes[5] signs) returns()
func (_IProxy *IProxyTransactorSession) Ban(_i uint64, _ban bool, signs [5][]byte) (*types.Transaction, error) {
	return _IProxy.Contract.Ban(&_IProxy.TransactOpts, _i, _ban, signs)
}

// BanG is a paid mutator transaction binding the contract method 0x3e54a8b8.
//
// Solidity: function banG(uint64 _gi, bool _ban, bytes[5] signs) returns()
func (_IProxy *IProxyTransactor) BanG(opts *bind.TransactOpts, _gi uint64, _ban bool, signs [5][]byte) (*types.Transaction, error) {
	return _IProxy.contract.Transact(opts, "banG", _gi, _ban, signs)
}

// BanG is a paid mutator transaction binding the contract method 0x3e54a8b8.
//
// Solidity: function banG(uint64 _gi, bool _ban, bytes[5] signs) returns()
func (_IProxy *IProxySession) BanG(_gi uint64, _ban bool, signs [5][]byte) (*types.Transaction, error) {
	return _IProxy.Contract.BanG(&_IProxy.TransactOpts, _gi, _ban, signs)
}

// BanG is a paid mutator transaction binding the contract method 0x3e54a8b8.
//
// Solidity: function banG(uint64 _gi, bool _ban, bytes[5] signs) returns()
func (_IProxy *IProxyTransactorSession) BanG(_gi uint64, _ban bool, signs [5][]byte) (*types.Transaction, error) {
	return _IProxy.Contract.BanG(&_IProxy.TransactOpts, _gi, _ban, signs)
}

// BanT is a paid mutator transaction binding the contract method 0x7a743984.
//
// Solidity: function banT(address _t, bool _ban, bytes[5] signs) returns()
func (_IProxy *IProxyTransactor) BanT(opts *bind.TransactOpts, _t common.Address, _ban bool, signs [5][]byte) (*types.Transaction, error) {
	return _IProxy.contract.Transact(opts, "banT", _t, _ban, signs)
}

// BanT is a paid mutator transaction binding the contract method 0x7a743984.
//
// Solidity: function banT(address _t, bool _ban, bytes[5] signs) returns()
func (_IProxy *IProxySession) BanT(_t common.Address, _ban bool, signs [5][]byte) (*types.Transaction, error) {
	return _IProxy.Contract.BanT(&_IProxy.TransactOpts, _t, _ban, signs)
}

// BanT is a paid mutator transaction binding the contract method 0x7a743984.
//
// Solidity: function banT(address _t, bool _ban, bytes[5] signs) returns()
func (_IProxy *IProxyTransactorSession) BanT(_t common.Address, _ban bool, signs [5][]byte) (*types.Transaction, error) {
	return _IProxy.Contract.BanT(&_IProxy.TransactOpts, _t, _ban, signs)
}

// ConfirmPayee is a paid mutator transaction binding the contract method 0x310aa0be.
//
// Solidity: function confirmPayee(uint64 _i) returns()
func (_IProxy *IProxyTransactor) ConfirmPayee(opts *bind.TransactOpts, _i uint64) (*types.Transaction, error) {
	return _IProxy.contract.Transact(opts, "confirmPayee", _i)
}

// ConfirmPayee is a paid mutator transaction binding the contract method 0x310aa0be.
//
// Solidity: function confirmPayee(uint64 _i) returns()
func (_IProxy *IProxySession) ConfirmPayee(_i uint64) (*types.Transaction, error) {
	return _IProxy.Contract.ConfirmPayee(&_IProxy.TransactOpts, _i)
}

// ConfirmPayee is a paid mutator transaction binding the contract method 0x310aa0be.
//
// Solidity: function confirmPayee(uint64 _i) returns()
func (_IProxy *IProxyTransactorSession) ConfirmPayee(_i uint64) (*types.Transaction, error) {
	return _IProxy.Contract.ConfirmPayee(&_IProxy.TransactOpts, _i)
}

// CreateGroup is a paid mutator transaction binding the contract method 0x832ec245.
//
// Solidity: function createGroup(uint8 _level, uint8 _mr, uint8 _dc, uint8 _pc, uint256 _kr, uint256 _pr) returns()
func (_IProxy *IProxyTransactor) CreateGroup(opts *bind.TransactOpts, _level uint8, _mr uint8, _dc uint8, _pc uint8, _kr *big.Int, _pr *big.Int) (*types.Transaction, error) {
	return _IProxy.contract.Transact(opts, "createGroup", _level, _mr, _dc, _pc, _kr, _pr)
}

// CreateGroup is a paid mutator transaction binding the contract method 0x832ec245.
//
// Solidity: function createGroup(uint8 _level, uint8 _mr, uint8 _dc, uint8 _pc, uint256 _kr, uint256 _pr) returns()
func (_IProxy *IProxySession) CreateGroup(_level uint8, _mr uint8, _dc uint8, _pc uint8, _kr *big.Int, _pr *big.Int) (*types.Transaction, error) {
	return _IProxy.Contract.CreateGroup(&_IProxy.TransactOpts, _level, _mr, _dc, _pc, _kr, _pr)
}

// CreateGroup is a paid mutator transaction binding the contract method 0x832ec245.
//
// Solidity: function createGroup(uint8 _level, uint8 _mr, uint8 _dc, uint8 _pc, uint256 _kr, uint256 _pr) returns()
func (_IProxy *IProxyTransactorSession) CreateGroup(_level uint8, _mr uint8, _dc uint8, _pc uint8, _kr *big.Int, _pr *big.Int) (*types.Transaction, error) {
	return _IProxy.Contract.CreateGroup(&_IProxy.TransactOpts, _level, _mr, _dc, _pc, _kr, _pr)
}

// Pledge is a paid mutator transaction binding the contract method 0xd23f7482.
//
// Solidity: function pledge(uint64 _i, uint256 _money) returns()
func (_IProxy *IProxyTransactor) Pledge(opts *bind.TransactOpts, _i uint64, _money *big.Int) (*types.Transaction, error) {
	return _IProxy.contract.Transact(opts, "pledge", _i, _money)
}

// Pledge is a paid mutator transaction binding the contract method 0xd23f7482.
//
// Solidity: function pledge(uint64 _i, uint256 _money) returns()
func (_IProxy *IProxySession) Pledge(_i uint64, _money *big.Int) (*types.Transaction, error) {
	return _IProxy.Contract.Pledge(&_IProxy.TransactOpts, _i, _money)
}

// Pledge is a paid mutator transaction binding the contract method 0xd23f7482.
//
// Solidity: function pledge(uint64 _i, uint256 _money) returns()
func (_IProxy *IProxyTransactorSession) Pledge(_i uint64, _money *big.Int) (*types.Transaction, error) {
	return _IProxy.Contract.Pledge(&_IProxy.TransactOpts, _i, _money)
}

// ProWithdraw is a paid mutator transaction binding the contract method 0xe182d54c.
//
// Solidity: function proWithdraw((uint64,uint8,uint256,uint256) _ps, uint64[] _kis, bytes[] ksigns) returns()
func (_IProxy *IProxyTransactor) ProWithdraw(opts *bind.TransactOpts, _ps PWIn, _kis []uint64, ksigns [][]byte) (*types.Transaction, error) {
	return _IProxy.contract.Transact(opts, "proWithdraw", _ps, _kis, ksigns)
}

// ProWithdraw is a paid mutator transaction binding the contract method 0xe182d54c.
//
// Solidity: function proWithdraw((uint64,uint8,uint256,uint256) _ps, uint64[] _kis, bytes[] ksigns) returns()
func (_IProxy *IProxySession) ProWithdraw(_ps PWIn, _kis []uint64, ksigns [][]byte) (*types.Transaction, error) {
	return _IProxy.Contract.ProWithdraw(&_IProxy.TransactOpts, _ps, _kis, ksigns)
}

// ProWithdraw is a paid mutator transaction binding the contract method 0xe182d54c.
//
// Solidity: function proWithdraw((uint64,uint8,uint256,uint256) _ps, uint64[] _kis, bytes[] ksigns) returns()
func (_IProxy *IProxyTransactorSession) ProWithdraw(_ps PWIn, _kis []uint64, ksigns [][]byte) (*types.Transaction, error) {
	return _IProxy.Contract.ProWithdraw(&_IProxy.TransactOpts, _ps, _kis, ksigns)
}

// ReAcc is a paid mutator transaction binding the contract method 0x6526250f.
//
// Solidity: function reAcc() returns()
func (_IProxy *IProxyTransactor) ReAcc(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IProxy.contract.Transact(opts, "reAcc")
}

// ReAcc is a paid mutator transaction binding the contract method 0x6526250f.
//
// Solidity: function reAcc() returns()
func (_IProxy *IProxySession) ReAcc() (*types.Transaction, error) {
	return _IProxy.Contract.ReAcc(&_IProxy.TransactOpts)
}

// ReAcc is a paid mutator transaction binding the contract method 0x6526250f.
//
// Solidity: function reAcc() returns()
func (_IProxy *IProxyTransactorSession) ReAcc() (*types.Transaction, error) {
	return _IProxy.Contract.ReAcc(&_IProxy.TransactOpts)
}

// ReRole is a paid mutator transaction binding the contract method 0x898e00f8.
//
// Solidity: function reRole(uint8 _rtype, bytes _extra) returns()
func (_IProxy *IProxyTransactor) ReRole(opts *bind.TransactOpts, _rtype uint8, _extra []byte) (*types.Transaction, error) {
	return _IProxy.contract.Transact(opts, "reRole", _rtype, _extra)
}

// ReRole is a paid mutator transaction binding the contract method 0x898e00f8.
//
// Solidity: function reRole(uint8 _rtype, bytes _extra) returns()
func (_IProxy *IProxySession) ReRole(_rtype uint8, _extra []byte) (*types.Transaction, error) {
	return _IProxy.Contract.ReRole(&_IProxy.TransactOpts, _rtype, _extra)
}

// ReRole is a paid mutator transaction binding the contract method 0x898e00f8.
//
// Solidity: function reRole(uint8 _rtype, bytes _extra) returns()
func (_IProxy *IProxyTransactorSession) ReRole(_rtype uint8, _extra []byte) (*types.Transaction, error) {
	return _IProxy.Contract.ReRole(&_IProxy.TransactOpts, _rtype, _extra)
}

// Recharge is a paid mutator transaction binding the contract method 0x24d11d40.
//
// Solidity: function recharge(uint64 _i, uint8 _ti, bool isLock, uint256 _money) returns()
func (_IProxy *IProxyTransactor) Recharge(opts *bind.TransactOpts, _i uint64, _ti uint8, isLock bool, _money *big.Int) (*types.Transaction, error) {
	return _IProxy.contract.Transact(opts, "recharge", _i, _ti, isLock, _money)
}

// Recharge is a paid mutator transaction binding the contract method 0x24d11d40.
//
// Solidity: function recharge(uint64 _i, uint8 _ti, bool isLock, uint256 _money) returns()
func (_IProxy *IProxySession) Recharge(_i uint64, _ti uint8, isLock bool, _money *big.Int) (*types.Transaction, error) {
	return _IProxy.Contract.Recharge(&_IProxy.TransactOpts, _i, _ti, isLock, _money)
}

// Recharge is a paid mutator transaction binding the contract method 0x24d11d40.
//
// Solidity: function recharge(uint64 _i, uint8 _ti, bool isLock, uint256 _money) returns()
func (_IProxy *IProxyTransactorSession) Recharge(_i uint64, _ti uint8, isLock bool, _money *big.Int) (*types.Transaction, error) {
	return _IProxy.Contract.Recharge(&_IProxy.TransactOpts, _i, _ti, isLock, _money)
}

// SubOrder is a paid mutator transaction binding the contract method 0x153ede07.
//
// Solidity: function subOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) _oi, bytes uSign, bytes pSign) returns()
func (_IProxy *IProxyTransactor) SubOrder(opts *bind.TransactOpts, _oi OrderIn, uSign []byte, pSign []byte) (*types.Transaction, error) {
	return _IProxy.contract.Transact(opts, "subOrder", _oi, uSign, pSign)
}

// SubOrder is a paid mutator transaction binding the contract method 0x153ede07.
//
// Solidity: function subOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) _oi, bytes uSign, bytes pSign) returns()
func (_IProxy *IProxySession) SubOrder(_oi OrderIn, uSign []byte, pSign []byte) (*types.Transaction, error) {
	return _IProxy.Contract.SubOrder(&_IProxy.TransactOpts, _oi, uSign, pSign)
}

// SubOrder is a paid mutator transaction binding the contract method 0x153ede07.
//
// Solidity: function subOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) _oi, bytes uSign, bytes pSign) returns()
func (_IProxy *IProxyTransactorSession) SubOrder(_oi OrderIn, uSign []byte, pSign []byte) (*types.Transaction, error) {
	return _IProxy.Contract.SubOrder(&_IProxy.TransactOpts, _oi, uSign, pSign)
}

// Unpledge is a paid mutator transaction binding the contract method 0x60985756.
//
// Solidity: function unpledge(uint64 _i, uint8 _ti, uint256 _money) returns()
func (_IProxy *IProxyTransactor) Unpledge(opts *bind.TransactOpts, _i uint64, _ti uint8, _money *big.Int) (*types.Transaction, error) {
	return _IProxy.contract.Transact(opts, "unpledge", _i, _ti, _money)
}

// Unpledge is a paid mutator transaction binding the contract method 0x60985756.
//
// Solidity: function unpledge(uint64 _i, uint8 _ti, uint256 _money) returns()
func (_IProxy *IProxySession) Unpledge(_i uint64, _ti uint8, _money *big.Int) (*types.Transaction, error) {
	return _IProxy.Contract.Unpledge(&_IProxy.TransactOpts, _i, _ti, _money)
}

// Unpledge is a paid mutator transaction binding the contract method 0x60985756.
//
// Solidity: function unpledge(uint64 _i, uint8 _ti, uint256 _money) returns()
func (_IProxy *IProxyTransactorSession) Unpledge(_i uint64, _ti uint8, _money *big.Int) (*types.Transaction, error) {
	return _IProxy.Contract.Unpledge(&_IProxy.TransactOpts, _i, _ti, _money)
}

// Withdraw is a paid mutator transaction binding the contract method 0x259c6d5e.
//
// Solidity: function withdraw(uint64 _i, uint8 _ti, uint256 _money) returns()
func (_IProxy *IProxyTransactor) Withdraw(opts *bind.TransactOpts, _i uint64, _ti uint8, _money *big.Int) (*types.Transaction, error) {
	return _IProxy.contract.Transact(opts, "withdraw", _i, _ti, _money)
}

// Withdraw is a paid mutator transaction binding the contract method 0x259c6d5e.
//
// Solidity: function withdraw(uint64 _i, uint8 _ti, uint256 _money) returns()
func (_IProxy *IProxySession) Withdraw(_i uint64, _ti uint8, _money *big.Int) (*types.Transaction, error) {
	return _IProxy.Contract.Withdraw(&_IProxy.TransactOpts, _i, _ti, _money)
}

// Withdraw is a paid mutator transaction binding the contract method 0x259c6d5e.
//
// Solidity: function withdraw(uint64 _i, uint8 _ti, uint256 _money) returns()
func (_IProxy *IProxyTransactorSession) Withdraw(_i uint64, _ti uint8, _money *big.Int) (*types.Transaction, error) {
	return _IProxy.Contract.Withdraw(&_IProxy.TransactOpts, _i, _ti, _money)
}

// IProxyActivateIterator is returned from FilterActivate and is used to iterate over the raw logs and unpacked data for Activate events raised by the IProxy contract.
type IProxyActivateIterator struct {
	Event *IProxyActivate // Event containing the contract specifics and raw log

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
func (it *IProxyActivateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IProxyActivate)
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
		it.Event = new(IProxyActivate)
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
func (it *IProxyActivateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IProxyActivateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IProxyActivate represents a Activate event raised by the IProxy contract.
type IProxyActivate struct {
	I   uint64
	Raw types.Log // Blockchain specific contextual infos
}

// FilterActivate is a free log retrieval operation binding the contract event 0x452a42fb51ebab9c8ebcbb3d536371397d861478788df33e9d2c9ae57765431c.
//
// Solidity: event Activate(uint64 i)
func (_IProxy *IProxyFilterer) FilterActivate(opts *bind.FilterOpts) (*IProxyActivateIterator, error) {

	logs, sub, err := _IProxy.contract.FilterLogs(opts, "Activate")
	if err != nil {
		return nil, err
	}
	return &IProxyActivateIterator{contract: _IProxy.contract, event: "Activate", logs: logs, sub: sub}, nil
}

// WatchActivate is a free log subscription operation binding the contract event 0x452a42fb51ebab9c8ebcbb3d536371397d861478788df33e9d2c9ae57765431c.
//
// Solidity: event Activate(uint64 i)
func (_IProxy *IProxyFilterer) WatchActivate(opts *bind.WatchOpts, sink chan<- *IProxyActivate) (event.Subscription, error) {

	logs, sub, err := _IProxy.contract.WatchLogs(opts, "Activate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IProxyActivate)
				if err := _IProxy.contract.UnpackLog(event, "Activate", log); err != nil {
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

// ParseActivate is a log parse operation binding the contract event 0x452a42fb51ebab9c8ebcbb3d536371397d861478788df33e9d2c9ae57765431c.
//
// Solidity: event Activate(uint64 i)
func (_IProxy *IProxyFilterer) ParseActivate(log types.Log) (*IProxyActivate, error) {
	event := new(IProxyActivate)
	if err := _IProxy.contract.UnpackLog(event, "Activate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IProxyAddTIterator is returned from FilterAddT and is used to iterate over the raw logs and unpacked data for AddT events raised by the IProxy contract.
type IProxyAddTIterator struct {
	Event *IProxyAddT // Event containing the contract specifics and raw log

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
func (it *IProxyAddTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IProxyAddT)
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
		it.Event = new(IProxyAddT)
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
func (it *IProxyAddTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IProxyAddTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IProxyAddT represents a AddT event raised by the IProxy contract.
type IProxyAddT struct {
	T    common.Address
	Mint bool
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAddT is a free log retrieval operation binding the contract event 0x4954a00722fe171d5c0079b423ebc3b3984d83c24170fdb068dfbd979d250412.
//
// Solidity: event AddT(address t, bool mint)
func (_IProxy *IProxyFilterer) FilterAddT(opts *bind.FilterOpts) (*IProxyAddTIterator, error) {

	logs, sub, err := _IProxy.contract.FilterLogs(opts, "AddT")
	if err != nil {
		return nil, err
	}
	return &IProxyAddTIterator{contract: _IProxy.contract, event: "AddT", logs: logs, sub: sub}, nil
}

// WatchAddT is a free log subscription operation binding the contract event 0x4954a00722fe171d5c0079b423ebc3b3984d83c24170fdb068dfbd979d250412.
//
// Solidity: event AddT(address t, bool mint)
func (_IProxy *IProxyFilterer) WatchAddT(opts *bind.WatchOpts, sink chan<- *IProxyAddT) (event.Subscription, error) {

	logs, sub, err := _IProxy.contract.WatchLogs(opts, "AddT")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IProxyAddT)
				if err := _IProxy.contract.UnpackLog(event, "AddT", log); err != nil {
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

// ParseAddT is a log parse operation binding the contract event 0x4954a00722fe171d5c0079b423ebc3b3984d83c24170fdb068dfbd979d250412.
//
// Solidity: event AddT(address t, bool mint)
func (_IProxy *IProxyFilterer) ParseAddT(log types.Log) (*IProxyAddT, error) {
	event := new(IProxyAddT)
	if err := _IProxy.contract.UnpackLog(event, "AddT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IProxyBanIterator is returned from FilterBan and is used to iterate over the raw logs and unpacked data for Ban events raised by the IProxy contract.
type IProxyBanIterator struct {
	Event *IProxyBan // Event containing the contract specifics and raw log

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
func (it *IProxyBanIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IProxyBan)
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
		it.Event = new(IProxyBan)
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
func (it *IProxyBanIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IProxyBanIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IProxyBan represents a Ban event raised by the IProxy contract.
type IProxyBan struct {
	I   uint64
	Ban bool
	Raw types.Log // Blockchain specific contextual infos
}

// FilterBan is a free log retrieval operation binding the contract event 0xb5b8b6cac24ab10b3b6a9c41f09df98b5763635f844abf1feb84ae9909902201.
//
// Solidity: event Ban(uint64 i, bool ban)
func (_IProxy *IProxyFilterer) FilterBan(opts *bind.FilterOpts) (*IProxyBanIterator, error) {

	logs, sub, err := _IProxy.contract.FilterLogs(opts, "Ban")
	if err != nil {
		return nil, err
	}
	return &IProxyBanIterator{contract: _IProxy.contract, event: "Ban", logs: logs, sub: sub}, nil
}

// WatchBan is a free log subscription operation binding the contract event 0xb5b8b6cac24ab10b3b6a9c41f09df98b5763635f844abf1feb84ae9909902201.
//
// Solidity: event Ban(uint64 i, bool ban)
func (_IProxy *IProxyFilterer) WatchBan(opts *bind.WatchOpts, sink chan<- *IProxyBan) (event.Subscription, error) {

	logs, sub, err := _IProxy.contract.WatchLogs(opts, "Ban")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IProxyBan)
				if err := _IProxy.contract.UnpackLog(event, "Ban", log); err != nil {
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

// ParseBan is a log parse operation binding the contract event 0xb5b8b6cac24ab10b3b6a9c41f09df98b5763635f844abf1feb84ae9909902201.
//
// Solidity: event Ban(uint64 i, bool ban)
func (_IProxy *IProxyFilterer) ParseBan(log types.Log) (*IProxyBan, error) {
	event := new(IProxyBan)
	if err := _IProxy.contract.UnpackLog(event, "Ban", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IProxyBanGIterator is returned from FilterBanG and is used to iterate over the raw logs and unpacked data for BanG events raised by the IProxy contract.
type IProxyBanGIterator struct {
	Event *IProxyBanG // Event containing the contract specifics and raw log

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
func (it *IProxyBanGIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IProxyBanG)
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
		it.Event = new(IProxyBanG)
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
func (it *IProxyBanGIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IProxyBanGIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IProxyBanG represents a BanG event raised by the IProxy contract.
type IProxyBanG struct {
	Gi  uint64
	Ban bool
	Raw types.Log // Blockchain specific contextual infos
}

// FilterBanG is a free log retrieval operation binding the contract event 0x976eddb294ef382c026f36e87244e52b5ce40d48ce2020101d563021441aceec.
//
// Solidity: event BanG(uint64 gi, bool ban)
func (_IProxy *IProxyFilterer) FilterBanG(opts *bind.FilterOpts) (*IProxyBanGIterator, error) {

	logs, sub, err := _IProxy.contract.FilterLogs(opts, "BanG")
	if err != nil {
		return nil, err
	}
	return &IProxyBanGIterator{contract: _IProxy.contract, event: "BanG", logs: logs, sub: sub}, nil
}

// WatchBanG is a free log subscription operation binding the contract event 0x976eddb294ef382c026f36e87244e52b5ce40d48ce2020101d563021441aceec.
//
// Solidity: event BanG(uint64 gi, bool ban)
func (_IProxy *IProxyFilterer) WatchBanG(opts *bind.WatchOpts, sink chan<- *IProxyBanG) (event.Subscription, error) {

	logs, sub, err := _IProxy.contract.WatchLogs(opts, "BanG")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IProxyBanG)
				if err := _IProxy.contract.UnpackLog(event, "BanG", log); err != nil {
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

// ParseBanG is a log parse operation binding the contract event 0x976eddb294ef382c026f36e87244e52b5ce40d48ce2020101d563021441aceec.
//
// Solidity: event BanG(uint64 gi, bool ban)
func (_IProxy *IProxyFilterer) ParseBanG(log types.Log) (*IProxyBanG, error) {
	event := new(IProxyBanG)
	if err := _IProxy.contract.UnpackLog(event, "BanG", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IProxyBanTIterator is returned from FilterBanT and is used to iterate over the raw logs and unpacked data for BanT events raised by the IProxy contract.
type IProxyBanTIterator struct {
	Event *IProxyBanT // Event containing the contract specifics and raw log

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
func (it *IProxyBanTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IProxyBanT)
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
		it.Event = new(IProxyBanT)
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
func (it *IProxyBanTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IProxyBanTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IProxyBanT represents a BanT event raised by the IProxy contract.
type IProxyBanT struct {
	T   common.Address
	Ban bool
	Raw types.Log // Blockchain specific contextual infos
}

// FilterBanT is a free log retrieval operation binding the contract event 0xffaa091f00488e7d9934b8efc3cbe385e3bb89957862cb891e428474b69b6bc9.
//
// Solidity: event BanT(address t, bool ban)
func (_IProxy *IProxyFilterer) FilterBanT(opts *bind.FilterOpts) (*IProxyBanTIterator, error) {

	logs, sub, err := _IProxy.contract.FilterLogs(opts, "BanT")
	if err != nil {
		return nil, err
	}
	return &IProxyBanTIterator{contract: _IProxy.contract, event: "BanT", logs: logs, sub: sub}, nil
}

// WatchBanT is a free log subscription operation binding the contract event 0xffaa091f00488e7d9934b8efc3cbe385e3bb89957862cb891e428474b69b6bc9.
//
// Solidity: event BanT(address t, bool ban)
func (_IProxy *IProxyFilterer) WatchBanT(opts *bind.WatchOpts, sink chan<- *IProxyBanT) (event.Subscription, error) {

	logs, sub, err := _IProxy.contract.WatchLogs(opts, "BanT")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IProxyBanT)
				if err := _IProxy.contract.UnpackLog(event, "BanT", log); err != nil {
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

// ParseBanT is a log parse operation binding the contract event 0xffaa091f00488e7d9934b8efc3cbe385e3bb89957862cb891e428474b69b6bc9.
//
// Solidity: event BanT(address t, bool ban)
func (_IProxy *IProxyFilterer) ParseBanT(log types.Log) (*IProxyBanT, error) {
	event := new(IProxyBanT)
	if err := _IProxy.contract.UnpackLog(event, "BanT", log); err != nil {
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

// ProxyABI is the input ABI used to generate the binding from.
const ProxyABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_inst\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"i\",\"type\":\"uint64\"}],\"name\":\"Activate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isSet\",\"type\":\"bool\"}],\"name\":\"AddOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"mint\",\"type\":\"bool\"}],\"name\":\"AddT\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"i\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"ban\",\"type\":\"bool\"}],\"name\":\"Ban\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"gi\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"ban\",\"type\":\"bool\"}],\"name\":\"BanG\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"ban\",\"type\":\"bool\"}],\"name\":\"BanT\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"activate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isSet\",\"type\":\"bool\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"uIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"tIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"sprice\",\"type\":\"uint256\"}],\"internalType\":\"structOrderIn\",\"name\":\"_oi\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"uSign\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"pSign\",\"type\":\"bytes\"}],\"name\":\"addOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"uIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"tIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"sprice\",\"type\":\"uint256\"}],\"internalType\":\"structOrderIn\",\"name\":\"_oi\",\"type\":\"tuple\"},{\"internalType\":\"uint64[]\",\"name\":\"_kis\",\"type\":\"uint64[]\"},{\"internalType\":\"bytes[]\",\"name\":\"ksigns\",\"type\":\"bytes[]\"}],\"name\":\"addRepair\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_t\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_mint\",\"type\":\"bool\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"addT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"}],\"name\":\"addToGroup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_p\",\"type\":\"address\"}],\"name\":\"alterPayee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"auth\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"_ban\",\"type\":\"bool\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"ban\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"_ban\",\"type\":\"bool\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"banG\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_t\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isBan\",\"type\":\"bool\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"banT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"}],\"name\":\"confirmPayee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_level\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_mr\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_dc\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_pc\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_kr\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_pr\",\"type\":\"uint256\"}],\"name\":\"createGroup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inst\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"owners\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"_money\",\"type\":\"uint256\"}],\"name\":\"pledge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"pIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"tIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"pay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lost\",\"type\":\"uint256\"}],\"internalType\":\"structPWIn\",\"name\":\"_ps\",\"type\":\"tuple\"},{\"internalType\":\"uint64[]\",\"name\":\"_kis\",\"type\":\"uint64[]\"},{\"internalType\":\"bytes[]\",\"name\":\"ksigns\",\"type\":\"bytes[]\"}],\"name\":\"proWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reAcc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_rtype\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_extra\",\"type\":\"bytes\"}],\"name\":\"reRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isLock\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_money\",\"type\":\"uint256\"}],\"name\":\"recharge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"uIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"tIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"sprice\",\"type\":\"uint256\"}],\"internalType\":\"structOrderIn\",\"name\":\"_oi\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"uSign\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"pSign\",\"type\":\"bytes\"}],\"name\":\"subOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_money\",\"type\":\"uint256\"}],\"name\":\"unpledge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_money\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// ProxyFuncSigs maps the 4-byte function signature to its string representation.
var ProxyFuncSigs = map[string]string{
	"843d1b49": "activate(uint64,bytes[5])",
	"4bf1b134": "add(address,bool,bytes[5])",
	"c05bdcd0": "addOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256),bytes,bytes)",
	"97f2352c": "addRepair((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256),uint64[],bytes[])",
	"68612348": "addT(address,bool,bytes[5])",
	"a246de42": "addToGroup(uint64)",
	"556e71c2": "alterPayee(address)",
	"de9375f2": "auth()",
	"df897f4f": "ban(uint64,bool,bytes[5])",
	"3e54a8b8": "banG(uint64,bool,bytes[5])",
	"7a743984": "banT(address,bool,bytes[5])",
	"310aa0be": "confirmPayee(uint64)",
	"832ec245": "createGroup(uint8,uint8,uint8,uint8,uint256,uint256)",
	"bd6b2222": "inst()",
	"022914a7": "owners(address)",
	"d23f7482": "pledge(uint64,uint256)",
	"e182d54c": "proWithdraw((uint64,uint8,uint256,uint256),uint64[],bytes[])",
	"6526250f": "reAcc()",
	"898e00f8": "reRole(uint8,bytes)",
	"24d11d40": "recharge(uint64,uint8,bool,uint256)",
	"153ede07": "subOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256),bytes,bytes)",
	"60985756": "unpledge(uint64,uint8,uint256)",
	"54fd4d50": "version()",
	"259c6d5e": "withdraw(uint64,uint8,uint256)",
}

// ProxyBin is the compiled bytecode used for deploying new contracts.
var ProxyBin = "0x60806040526001805461ffff60a01b1916600160a11b1790553480156200002557600080fd5b50604051620021fb380380620021fb833981016040819052620000489162000097565b600180546001600160a01b039384166001600160a01b03199182161790915560028054929093169116179055620000cf565b80516001600160a01b03811681146200009257600080fd5b919050565b60008060408385031215620000ab57600080fd5b620000b6836200007a565b9150620000c6602084016200007a565b90509250929050565b61211c80620000df6000396000f3fe60806040526004361061014f5760003560e01c80637a743984116100b6578063bd6b22221161006f578063bd6b2222146103cc578063c05bdcd014610404578063d23f748214610424578063de9375f214610444578063df897f4f14610464578063e182d54c1461048457600080fd5b80637a7439841461030c578063832ec2451461032c578063843d1b491461034c578063898e00f81461036c57806397f2352c1461038c578063a246de42146103ac57600080fd5b80634bf1b134116101085780634bf1b1341461024257806354fd4d5014610262578063556e71c21461029757806360985756146102b75780636526250f146102d757806368612348146102ec57600080fd5b8063022914a71461015b578063153ede07146101a057806324d11d40146101c2578063259c6d5e146101e2578063310aa0be146102025780633e54a8b81461022257600080fd5b3661015657005b600080fd5b34801561016757600080fd5b5061018b610176366004611611565b60006020819052908152604090205460ff1681565b60405190151581526020015b60405180910390f35b3480156101ac57600080fd5b506101c06101bb3660046117ff565b6104a4565b005b3480156101ce57600080fd5b506101c06101dd366004611886565b610579565b3480156101ee57600080fd5b506101c06101fd3660046118d1565b61066e565b34801561020e57600080fd5b506101c061021d36600461190d565b610728565b34801561022e57600080fd5b506101c061023d3660046119b5565b610803565b34801561024e57600080fd5b506101c061025d366004611a08565b61091b565b34801561026e57600080fd5b5060015461028490600160a01b900461ffff1681565b60405161ffff9091168152602001610197565b3480156102a357600080fd5b506101c06102b2366004611611565b610a98565b3480156102c357600080fd5b506101c06102d23660046118d1565b610b3b565b3480156102e357600080fd5b506101c0610bf5565b3480156102f857600080fd5b506101c0610307366004611a08565b610cbf565b34801561031857600080fd5b506101c0610327366004611a08565b610dce565b34801561033857600080fd5b506101c0610347366004611a28565b610edd565b34801561035857600080fd5b506101c0610367366004611a8e565b610fd6565b34801561037857600080fd5b506101c0610387366004611adb565b6110e3565b34801561039857600080fd5b506101c06103a7366004611c32565b6111b5565b3480156103b857600080fd5b506101c06103c736600461190d565b611253565b3480156103d857600080fd5b506002546103ec906001600160a01b031681565b6040516001600160a01b039091168152602001610197565b34801561041057600080fd5b506101c061041f3660046117ff565b6112fd565b34801561043057600080fd5b506101c061043f366004611c9f565b61139b565b34801561045057600080fd5b506001546103ec906001600160a01b031681565b34801561047057600080fd5b506101c061047f3660046119b5565b61144c565b34801561049057600080fd5b506101c061049f366004611cc9565b61155b565b600254604051633ec7d5b960e01b8152606660048201526001600160a01b0390911690633ec7d5b990602401602060405180830381865afa1580156104ed573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105119190611d66565b6001600160a01b03166342f45166338585856040518563ffffffff1660e01b81526004016105429493929190611e3b565b600060405180830381600087803b15801561055c57600080fd5b505af1158015610570573d6000803e3d6000fd5b50505050505050565b600254604051633ec7d5b960e01b8152606660048201526001600160a01b0390911690633ec7d5b990602401602060405180830381865afa1580156105c2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105e69190611d66565b60405163f661f9e360e01b81523360048201526001600160401b038616602482015260ff851660448201528315156064820152608481018390526001600160a01b03919091169063f661f9e39060a401600060405180830381600087803b15801561065057600080fd5b505af1158015610664573d6000803e3d6000fd5b5050505050505050565b600254604051633ec7d5b960e01b8152606660048201526001600160a01b0390911690633ec7d5b990602401602060405180830381865afa1580156106b7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106db9190611d66565b604051634da63abd60e11b81523360048201526001600160401b038516602482015260ff84166044820152606481018390526001600160a01b039190911690639b4c757a90608401610542565b600254604051633ec7d5b960e01b8152606560048201526001600160a01b0390911690633ec7d5b990602401602060405180830381865afa158015610771573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107959190611d66565b604051630a38216560e31b81523360048201526001600160401b03831660248201526001600160a01b0391909116906351c10b28906044015b600060405180830381600087803b1580156107e857600080fd5b505af11580156107fc573d6000803e3d6000fd5b5050505050565b600254604051633ec7d5b960e01b8152606560048201526001600160a01b0390911690633ec7d5b990602401602060405180830381865afa15801561084c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108709190611d66565b6001600160a01b0316633e54a8b88484846040518463ffffffff1660e01b815260040161089f93929190611ed1565b600060405180830381600087803b1580156108b957600080fd5b505af11580156108cd573d6000803e3d6000fd5b5050604080516001600160401b038716815285151560208201527f976eddb294ef382c026f36e87244e52b5ce40d48ce2020101d563021441aceec93500190505b60405180910390a1505050565b823b60008190036109675760405162461bcd60e51b81526020600482015260126024820152713732b2b21031b7b73a3930b1ba1030b2323960711b604482015260640160405180910390fd5b6040516bffffffffffffffffffffffff1930606090811b821660208401526218591960ea1b603484015286901b16603782015283151560f81b604b820152600090604c0160408051601f1981840301815290829052805160209091012060015463a96bba9d60e01b83529092506001600160a01b03169063a96bba9d906109f49084908790600401611f04565b600060405180830381600087803b158015610a0e57600080fd5b505af1158015610a22573d6000803e3d6000fd5b5050604080516001600160a01b038916815287151560208201527f938b2a24c98e4e542127ffa74a91e48942c2bddccae3b6d75f82bfda7bbc0807935001905060405180910390a15050506001600160a01b03919091166000908152602081905260409020805460ff1916911515919091179055565b600254604051633ec7d5b960e01b8152606560048201526001600160a01b0390911690633ec7d5b990602401602060405180830381865afa158015610ae1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b059190611d66565b6040516322bd405560e11b81523360048201526001600160a01b038381166024830152919091169063457a80aa906044016107ce565b600254604051633ec7d5b960e01b8152606560048201526001600160a01b0390911690633ec7d5b990602401602060405180830381865afa158015610b84573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ba89190611d66565b604051630a866aed60e01b81523360048201526001600160401b038516602482015260ff84166044820152606481018390526001600160a01b039190911690630a866aed90608401610542565b600254604051633ec7d5b960e01b8152606560048201526001600160a01b0390911690633ec7d5b990602401602060405180830381865afa158015610c3e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c629190611d66565b6040516377fe57e760e11b81523360048201526001600160a01b03919091169063effcafce90602401600060405180830381600087803b158015610ca557600080fd5b505af1158015610cb9573d6000803e3d6000fd5b50505050565b600254604051633ec7d5b960e01b8152606560048201526001600160a01b0390911690633ec7d5b990602401602060405180830381865afa158015610d08573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d2c9190611d66565b6001600160a01b031663686123488484846040518463ffffffff1660e01b8152600401610d5b93929190611f25565b600060405180830381600087803b158015610d7557600080fd5b505af1158015610d89573d6000803e3d6000fd5b5050604080516001600160a01b038716815285151560208201527f4954a00722fe171d5c0079b423ebc3b3984d83c24170fdb068dfbd979d250412935001905061090e565b600254604051633ec7d5b960e01b8152606560048201526001600160a01b0390911690633ec7d5b990602401602060405180830381865afa158015610e17573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e3b9190611d66565b6001600160a01b0316637a7439848484846040518463ffffffff1660e01b8152600401610e6a93929190611f25565b600060405180830381600087803b158015610e8457600080fd5b505af1158015610e98573d6000803e3d6000fd5b5050604080516001600160a01b038716815285151560208201527fffaa091f00488e7d9934b8efc3cbe385e3bb89957862cb891e428474b69b6bc9935001905061090e565b600254604051633ec7d5b960e01b8152606560048201526001600160a01b0390911690633ec7d5b990602401602060405180830381865afa158015610f26573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f4a9190611d66565b60405163832ec24560e01b815260ff808916600483015280881660248301528087166044830152851660648201526084810184905260a481018390526001600160a01b03919091169063832ec2459060c401600060405180830381600087803b158015610fb657600080fd5b505af1158015610fca573d6000803e3d6000fd5b50505050505050505050565b600254604051633ec7d5b960e01b8152606560048201526001600160a01b0390911690633ec7d5b990602401602060405180830381865afa15801561101f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110439190611d66565b6001600160a01b031663843d1b4983836040518363ffffffff1660e01b8152600401611070929190611f51565b600060405180830381600087803b15801561108a57600080fd5b505af115801561109e573d6000803e3d6000fd5b50506040516001600160401b03851681527f452a42fb51ebab9c8ebcbb3d536371397d861478788df33e9d2c9ae57765431c9250602001905060405180910390a15050565b600254604051633ec7d5b960e01b8152606560048201526001600160a01b0390911690633ec7d5b990602401602060405180830381865afa15801561112c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111509190611d66565b6001600160a01b0316639ccfe4ff3384846040518463ffffffff1660e01b815260040161117f93929190611f73565b600060405180830381600087803b15801561119957600080fd5b505af11580156111ad573d6000803e3d6000fd5b505050505050565b600254604051633ec7d5b960e01b8152606660048201526001600160a01b0390911690633ec7d5b990602401602060405180830381865afa1580156111fe573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112229190611d66565b6001600160a01b031663af99c59a338585856040518563ffffffff1660e01b8152600401610542949392919061203c565b600254604051633ec7d5b960e01b8152606560048201526001600160a01b0390911690633ec7d5b990602401602060405180830381865afa15801561129c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112c09190611d66565b60405163d963723160e01b81523360048201526001600160401b03831660248201526001600160a01b03919091169063d9637231906044016107ce565b600254604051633ec7d5b960e01b8152606660048201526001600160a01b0390911690633ec7d5b990602401602060405180830381865afa158015611346573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061136a9190611d66565b6001600160a01b031663ae9d0b40338585856040518563ffffffff1660e01b81526004016105429493929190611e3b565b600254604051633ec7d5b960e01b8152606560048201526001600160a01b0390911690633ec7d5b990602401602060405180830381865afa1580156113e4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114089190611d66565b604051630d11731760e11b81523360048201526001600160401b0384166024820152604481018390526001600160a01b039190911690631a22e62e9060640161117f565b600254604051633ec7d5b960e01b8152606560048201526001600160a01b0390911690633ec7d5b990602401602060405180830381865afa158015611495573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114b99190611d66565b6001600160a01b031663df897f4f8484846040518463ffffffff1660e01b81526004016114e893929190611ed1565b600060405180830381600087803b15801561150257600080fd5b505af1158015611516573d6000803e3d6000fd5b5050604080516001600160401b038716815285151560208201527fb5b8b6cac24ab10b3b6a9c41f09df98b5763635f844abf1feb84ae9909902201935001905061090e565b600254604051633ec7d5b960e01b8152606660048201526001600160a01b0390911690633ec7d5b990602401602060405180830381865afa1580156115a4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115c89190611d66565b6001600160a01b03166354aa6642338585856040518563ffffffff1660e01b81526004016105429493929190612082565b6001600160a01b038116811461160e57600080fd5b50565b60006020828403121561162357600080fd5b813561162e816115f9565b9392505050565b634e487b7160e01b600052604160045260246000fd5b604051608081016001600160401b038111828210171561166d5761166d611635565b60405290565b604051601f8201601f191681016001600160401b038111828210171561169b5761169b611635565b604052919050565b80356001600160401b03811681146116ba57600080fd5b919050565b803560ff811681146116ba57600080fd5b60006101008083850312156116e457600080fd5b604051908101906001600160401b038211818310171561170657611706611635565b81604052809250611716846116a3565b8152611724602085016116a3565b6020820152611735604085016116a3565b6040820152611746606085016116a3565b6060820152611757608085016116a3565b608082015261176860a085016116a3565b60a082015261177960c085016116bf565b60c082015260e084013560e0820152505092915050565b600082601f8301126117a157600080fd5b81356001600160401b038111156117ba576117ba611635565b6117cd601f8201601f1916602001611673565b8181528460208386010111156117e257600080fd5b816020850160208301376000918101602001919091529392505050565b6000806000610140848603121561181557600080fd5b61181f85856116d0565b92506101008401356001600160401b038082111561183c57600080fd5b61184887838801611790565b935061012086013591508082111561185f57600080fd5b5061186c86828701611790565b9150509250925092565b803580151581146116ba57600080fd5b6000806000806080858703121561189c57600080fd5b6118a5856116a3565b93506118b3602086016116bf565b92506118c160408601611876565b9396929550929360600135925050565b6000806000606084860312156118e657600080fd5b6118ef846116a3565b92506118fd602085016116bf565b9150604084013590509250925092565b60006020828403121561191f57600080fd5b61162e826116a3565b600082601f83011261193957600080fd5b60405160a081016001600160401b03828210818311171561195c5761195c611635565b8160405282915060a085018681111561197457600080fd5b855b818110156119a95780358381111561198e5760008081fd5b61199a89828a01611790565b85525060209384019301611976565b50929695505050505050565b6000806000606084860312156119ca57600080fd5b6119d3846116a3565b92506119e160208501611876565b915060408401356001600160401b038111156119fc57600080fd5b61186c86828701611928565b600080600060608486031215611a1d57600080fd5b83356119d3816115f9565b60008060008060008060c08789031215611a4157600080fd5b611a4a876116bf565b9550611a58602088016116bf565b9450611a66604088016116bf565b9350611a74606088016116bf565b92506080870135915060a087013590509295509295509295565b60008060408385031215611aa157600080fd5b611aaa836116a3565b915060208301356001600160401b03811115611ac557600080fd5b611ad185828601611928565b9150509250929050565b60008060408385031215611aee57600080fd5b611af7836116bf565b915060208301356001600160401b03811115611b1257600080fd5b611ad185828601611790565b60006001600160401b03821115611b3757611b37611635565b5060051b60200190565b600082601f830112611b5257600080fd5b81356020611b67611b6283611b1e565b611673565b82815260059290921b84018101918181019086841115611b8657600080fd5b8286015b84811015611ba857611b9b816116a3565b8352918301918301611b8a565b509695505050505050565b600082601f830112611bc457600080fd5b81356020611bd4611b6283611b1e565b82815260059290921b84018101918181019086841115611bf357600080fd5b8286015b84811015611ba85780356001600160401b03811115611c165760008081fd5b611c248986838b0101611790565b845250918301918301611bf7565b60008060006101408486031215611c4857600080fd5b611c5285856116d0565b92506101008401356001600160401b0380821115611c6f57600080fd5b611c7b87838801611b41565b9350610120860135915080821115611c9257600080fd5b5061186c86828701611bb3565b60008060408385031215611cb257600080fd5b611cbb836116a3565b946020939093013593505050565b600080600083850360c0811215611cdf57600080fd5b6080811215611ced57600080fd5b50611cf661164b565b611cff856116a3565b8152611d0d602086016116bf565b602082015260408501356040820152606085013560608201528093505060808401356001600160401b0380821115611d4457600080fd5b611d5087838801611b41565b935060a0860135915080821115611c9257600080fd5b600060208284031215611d7857600080fd5b815161162e816115f9565b6001600160401b038082511683528060208301511660208401528060408301511660408401528060608301511660608401528060808301511660808401528060a08301511660a08401525060c0810151611de260c084018260ff169052565b5060e090810151910152565b6000815180845260005b81811015611e1457602081850181015186830182015201611df8565b81811115611e26576000602083870101525b50601f01601f19169290920160200192915050565b6001600160a01b03851681526000610160611e596020840187611d83565b80610120840152611e6c81840186611dee565b9050828103610140840152611e818185611dee565b979650505050505050565b60008260a081018360005b6005811015611ec6578383038752611eb0838351611dee565b6020978801979093509190910190600101611e97565b509095945050505050565b6001600160401b03841681528215156020820152606060408201526000611efb6060830184611e8c565b95945050505050565b828152604060208201526000611f1d6040830184611e8c565b949350505050565b6001600160a01b03841681528215156020820152606060408201819052600090611efb90830184611e8c565b6001600160401b0383168152604060208201526000611f1d6040830184611e8c565b6001600160a01b038416815260ff83166020820152606060408201819052600090611efb90830184611dee565b600081518084526020808501945080840160005b83811015611fd95781516001600160401b031687529582019590820190600101611fb4565b509495945050505050565b600082825180855260208086019550808260051b84010181860160005b8481101561202f57601f1986840301895261201d838351611dee565b98840198925090830190600101612001565b5090979650505050505050565b6001600160a01b0385168152600061016061205a6020840187611d83565b8061012084015261206d81840186611fa0565b9050828103610140840152611e818185611fe4565b60018060a01b03851681526001600160401b03845116602082015260ff6020850151166040820152604084015160608201526060840151608082015260e060a082015260006120d460e0830185611fa0565b82810360c0840152611e818185611fe456fea2646970667358221220efabbbf585c03463f65fd3e92eda3cc7237999034b16f66bb31285e1de5ccc2c64736f6c634300080e0033"

// DeployProxy deploys a new Ethereum contract, binding an instance of Proxy to it.
func DeployProxy(auth *bind.TransactOpts, backend bind.ContractBackend, _a common.Address, _inst common.Address) (common.Address, *types.Transaction, *Proxy, error) {
	parsed, err := abi.JSON(strings.NewReader(ProxyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ProxyBin), backend, _a, _inst)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Proxy{ProxyCaller: ProxyCaller{contract: contract}, ProxyTransactor: ProxyTransactor{contract: contract}, ProxyFilterer: ProxyFilterer{contract: contract}}, nil
}

// Proxy is an auto generated Go binding around an Ethereum contract.
type Proxy struct {
	ProxyCaller     // Read-only binding to the contract
	ProxyTransactor // Write-only binding to the contract
	ProxyFilterer   // Log filterer for contract events
}

// ProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProxySession struct {
	Contract     *Proxy            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProxyCallerSession struct {
	Contract *ProxyCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProxyTransactorSession struct {
	Contract     *ProxyTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProxyRaw struct {
	Contract *Proxy // Generic contract binding to access the raw methods on
}

// ProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProxyCallerRaw struct {
	Contract *ProxyCaller // Generic read-only contract binding to access the raw methods on
}

// ProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProxyTransactorRaw struct {
	Contract *ProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProxy creates a new instance of Proxy, bound to a specific deployed contract.
func NewProxy(address common.Address, backend bind.ContractBackend) (*Proxy, error) {
	contract, err := bindProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Proxy{ProxyCaller: ProxyCaller{contract: contract}, ProxyTransactor: ProxyTransactor{contract: contract}, ProxyFilterer: ProxyFilterer{contract: contract}}, nil
}

// NewProxyCaller creates a new read-only instance of Proxy, bound to a specific deployed contract.
func NewProxyCaller(address common.Address, caller bind.ContractCaller) (*ProxyCaller, error) {
	contract, err := bindProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyCaller{contract: contract}, nil
}

// NewProxyTransactor creates a new write-only instance of Proxy, bound to a specific deployed contract.
func NewProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*ProxyTransactor, error) {
	contract, err := bindProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyTransactor{contract: contract}, nil
}

// NewProxyFilterer creates a new log filterer instance of Proxy, bound to a specific deployed contract.
func NewProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*ProxyFilterer, error) {
	contract, err := bindProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProxyFilterer{contract: contract}, nil
}

// bindProxy binds a generic wrapper to an already deployed contract.
func bindProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Proxy *ProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Proxy.Contract.ProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Proxy *ProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proxy.Contract.ProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Proxy *ProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Proxy.Contract.ProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Proxy *ProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Proxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Proxy *ProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Proxy *ProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Proxy.Contract.contract.Transact(opts, method, params...)
}

// Auth is a free data retrieval call binding the contract method 0xde9375f2.
//
// Solidity: function auth() view returns(address)
func (_Proxy *ProxyCaller) Auth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Proxy.contract.Call(opts, &out, "auth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Auth is a free data retrieval call binding the contract method 0xde9375f2.
//
// Solidity: function auth() view returns(address)
func (_Proxy *ProxySession) Auth() (common.Address, error) {
	return _Proxy.Contract.Auth(&_Proxy.CallOpts)
}

// Auth is a free data retrieval call binding the contract method 0xde9375f2.
//
// Solidity: function auth() view returns(address)
func (_Proxy *ProxyCallerSession) Auth() (common.Address, error) {
	return _Proxy.Contract.Auth(&_Proxy.CallOpts)
}

// Inst is a free data retrieval call binding the contract method 0xbd6b2222.
//
// Solidity: function inst() view returns(address)
func (_Proxy *ProxyCaller) Inst(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Proxy.contract.Call(opts, &out, "inst")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Inst is a free data retrieval call binding the contract method 0xbd6b2222.
//
// Solidity: function inst() view returns(address)
func (_Proxy *ProxySession) Inst() (common.Address, error) {
	return _Proxy.Contract.Inst(&_Proxy.CallOpts)
}

// Inst is a free data retrieval call binding the contract method 0xbd6b2222.
//
// Solidity: function inst() view returns(address)
func (_Proxy *ProxyCallerSession) Inst() (common.Address, error) {
	return _Proxy.Contract.Inst(&_Proxy.CallOpts)
}

// Owners is a free data retrieval call binding the contract method 0x022914a7.
//
// Solidity: function owners(address ) view returns(bool)
func (_Proxy *ProxyCaller) Owners(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Proxy.contract.Call(opts, &out, "owners", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Owners is a free data retrieval call binding the contract method 0x022914a7.
//
// Solidity: function owners(address ) view returns(bool)
func (_Proxy *ProxySession) Owners(arg0 common.Address) (bool, error) {
	return _Proxy.Contract.Owners(&_Proxy.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x022914a7.
//
// Solidity: function owners(address ) view returns(bool)
func (_Proxy *ProxyCallerSession) Owners(arg0 common.Address) (bool, error) {
	return _Proxy.Contract.Owners(&_Proxy.CallOpts, arg0)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint16)
func (_Proxy *ProxyCaller) Version(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Proxy.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint16)
func (_Proxy *ProxySession) Version() (uint16, error) {
	return _Proxy.Contract.Version(&_Proxy.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint16)
func (_Proxy *ProxyCallerSession) Version() (uint16, error) {
	return _Proxy.Contract.Version(&_Proxy.CallOpts)
}

// Activate is a paid mutator transaction binding the contract method 0x843d1b49.
//
// Solidity: function activate(uint64 _i, bytes[5] signs) returns()
func (_Proxy *ProxyTransactor) Activate(opts *bind.TransactOpts, _i uint64, signs [5][]byte) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "activate", _i, signs)
}

// Activate is a paid mutator transaction binding the contract method 0x843d1b49.
//
// Solidity: function activate(uint64 _i, bytes[5] signs) returns()
func (_Proxy *ProxySession) Activate(_i uint64, signs [5][]byte) (*types.Transaction, error) {
	return _Proxy.Contract.Activate(&_Proxy.TransactOpts, _i, signs)
}

// Activate is a paid mutator transaction binding the contract method 0x843d1b49.
//
// Solidity: function activate(uint64 _i, bytes[5] signs) returns()
func (_Proxy *ProxyTransactorSession) Activate(_i uint64, signs [5][]byte) (*types.Transaction, error) {
	return _Proxy.Contract.Activate(&_Proxy.TransactOpts, _i, signs)
}

// Add is a paid mutator transaction binding the contract method 0x4bf1b134.
//
// Solidity: function add(address _a, bool isSet, bytes[5] signs) returns()
func (_Proxy *ProxyTransactor) Add(opts *bind.TransactOpts, _a common.Address, isSet bool, signs [5][]byte) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "add", _a, isSet, signs)
}

// Add is a paid mutator transaction binding the contract method 0x4bf1b134.
//
// Solidity: function add(address _a, bool isSet, bytes[5] signs) returns()
func (_Proxy *ProxySession) Add(_a common.Address, isSet bool, signs [5][]byte) (*types.Transaction, error) {
	return _Proxy.Contract.Add(&_Proxy.TransactOpts, _a, isSet, signs)
}

// Add is a paid mutator transaction binding the contract method 0x4bf1b134.
//
// Solidity: function add(address _a, bool isSet, bytes[5] signs) returns()
func (_Proxy *ProxyTransactorSession) Add(_a common.Address, isSet bool, signs [5][]byte) (*types.Transaction, error) {
	return _Proxy.Contract.Add(&_Proxy.TransactOpts, _a, isSet, signs)
}

// AddOrder is a paid mutator transaction binding the contract method 0xc05bdcd0.
//
// Solidity: function addOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) _oi, bytes uSign, bytes pSign) returns()
func (_Proxy *ProxyTransactor) AddOrder(opts *bind.TransactOpts, _oi OrderIn, uSign []byte, pSign []byte) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "addOrder", _oi, uSign, pSign)
}

// AddOrder is a paid mutator transaction binding the contract method 0xc05bdcd0.
//
// Solidity: function addOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) _oi, bytes uSign, bytes pSign) returns()
func (_Proxy *ProxySession) AddOrder(_oi OrderIn, uSign []byte, pSign []byte) (*types.Transaction, error) {
	return _Proxy.Contract.AddOrder(&_Proxy.TransactOpts, _oi, uSign, pSign)
}

// AddOrder is a paid mutator transaction binding the contract method 0xc05bdcd0.
//
// Solidity: function addOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) _oi, bytes uSign, bytes pSign) returns()
func (_Proxy *ProxyTransactorSession) AddOrder(_oi OrderIn, uSign []byte, pSign []byte) (*types.Transaction, error) {
	return _Proxy.Contract.AddOrder(&_Proxy.TransactOpts, _oi, uSign, pSign)
}

// AddRepair is a paid mutator transaction binding the contract method 0x97f2352c.
//
// Solidity: function addRepair((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) _oi, uint64[] _kis, bytes[] ksigns) returns()
func (_Proxy *ProxyTransactor) AddRepair(opts *bind.TransactOpts, _oi OrderIn, _kis []uint64, ksigns [][]byte) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "addRepair", _oi, _kis, ksigns)
}

// AddRepair is a paid mutator transaction binding the contract method 0x97f2352c.
//
// Solidity: function addRepair((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) _oi, uint64[] _kis, bytes[] ksigns) returns()
func (_Proxy *ProxySession) AddRepair(_oi OrderIn, _kis []uint64, ksigns [][]byte) (*types.Transaction, error) {
	return _Proxy.Contract.AddRepair(&_Proxy.TransactOpts, _oi, _kis, ksigns)
}

// AddRepair is a paid mutator transaction binding the contract method 0x97f2352c.
//
// Solidity: function addRepair((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) _oi, uint64[] _kis, bytes[] ksigns) returns()
func (_Proxy *ProxyTransactorSession) AddRepair(_oi OrderIn, _kis []uint64, ksigns [][]byte) (*types.Transaction, error) {
	return _Proxy.Contract.AddRepair(&_Proxy.TransactOpts, _oi, _kis, ksigns)
}

// AddT is a paid mutator transaction binding the contract method 0x68612348.
//
// Solidity: function addT(address _t, bool _mint, bytes[5] signs) returns()
func (_Proxy *ProxyTransactor) AddT(opts *bind.TransactOpts, _t common.Address, _mint bool, signs [5][]byte) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "addT", _t, _mint, signs)
}

// AddT is a paid mutator transaction binding the contract method 0x68612348.
//
// Solidity: function addT(address _t, bool _mint, bytes[5] signs) returns()
func (_Proxy *ProxySession) AddT(_t common.Address, _mint bool, signs [5][]byte) (*types.Transaction, error) {
	return _Proxy.Contract.AddT(&_Proxy.TransactOpts, _t, _mint, signs)
}

// AddT is a paid mutator transaction binding the contract method 0x68612348.
//
// Solidity: function addT(address _t, bool _mint, bytes[5] signs) returns()
func (_Proxy *ProxyTransactorSession) AddT(_t common.Address, _mint bool, signs [5][]byte) (*types.Transaction, error) {
	return _Proxy.Contract.AddT(&_Proxy.TransactOpts, _t, _mint, signs)
}

// AddToGroup is a paid mutator transaction binding the contract method 0xa246de42.
//
// Solidity: function addToGroup(uint64 _gi) returns()
func (_Proxy *ProxyTransactor) AddToGroup(opts *bind.TransactOpts, _gi uint64) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "addToGroup", _gi)
}

// AddToGroup is a paid mutator transaction binding the contract method 0xa246de42.
//
// Solidity: function addToGroup(uint64 _gi) returns()
func (_Proxy *ProxySession) AddToGroup(_gi uint64) (*types.Transaction, error) {
	return _Proxy.Contract.AddToGroup(&_Proxy.TransactOpts, _gi)
}

// AddToGroup is a paid mutator transaction binding the contract method 0xa246de42.
//
// Solidity: function addToGroup(uint64 _gi) returns()
func (_Proxy *ProxyTransactorSession) AddToGroup(_gi uint64) (*types.Transaction, error) {
	return _Proxy.Contract.AddToGroup(&_Proxy.TransactOpts, _gi)
}

// AlterPayee is a paid mutator transaction binding the contract method 0x556e71c2.
//
// Solidity: function alterPayee(address _p) returns()
func (_Proxy *ProxyTransactor) AlterPayee(opts *bind.TransactOpts, _p common.Address) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "alterPayee", _p)
}

// AlterPayee is a paid mutator transaction binding the contract method 0x556e71c2.
//
// Solidity: function alterPayee(address _p) returns()
func (_Proxy *ProxySession) AlterPayee(_p common.Address) (*types.Transaction, error) {
	return _Proxy.Contract.AlterPayee(&_Proxy.TransactOpts, _p)
}

// AlterPayee is a paid mutator transaction binding the contract method 0x556e71c2.
//
// Solidity: function alterPayee(address _p) returns()
func (_Proxy *ProxyTransactorSession) AlterPayee(_p common.Address) (*types.Transaction, error) {
	return _Proxy.Contract.AlterPayee(&_Proxy.TransactOpts, _p)
}

// Ban is a paid mutator transaction binding the contract method 0xdf897f4f.
//
// Solidity: function ban(uint64 _i, bool _ban, bytes[5] signs) returns()
func (_Proxy *ProxyTransactor) Ban(opts *bind.TransactOpts, _i uint64, _ban bool, signs [5][]byte) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "ban", _i, _ban, signs)
}

// Ban is a paid mutator transaction binding the contract method 0xdf897f4f.
//
// Solidity: function ban(uint64 _i, bool _ban, bytes[5] signs) returns()
func (_Proxy *ProxySession) Ban(_i uint64, _ban bool, signs [5][]byte) (*types.Transaction, error) {
	return _Proxy.Contract.Ban(&_Proxy.TransactOpts, _i, _ban, signs)
}

// Ban is a paid mutator transaction binding the contract method 0xdf897f4f.
//
// Solidity: function ban(uint64 _i, bool _ban, bytes[5] signs) returns()
func (_Proxy *ProxyTransactorSession) Ban(_i uint64, _ban bool, signs [5][]byte) (*types.Transaction, error) {
	return _Proxy.Contract.Ban(&_Proxy.TransactOpts, _i, _ban, signs)
}

// BanG is a paid mutator transaction binding the contract method 0x3e54a8b8.
//
// Solidity: function banG(uint64 _gi, bool _ban, bytes[5] signs) returns()
func (_Proxy *ProxyTransactor) BanG(opts *bind.TransactOpts, _gi uint64, _ban bool, signs [5][]byte) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "banG", _gi, _ban, signs)
}

// BanG is a paid mutator transaction binding the contract method 0x3e54a8b8.
//
// Solidity: function banG(uint64 _gi, bool _ban, bytes[5] signs) returns()
func (_Proxy *ProxySession) BanG(_gi uint64, _ban bool, signs [5][]byte) (*types.Transaction, error) {
	return _Proxy.Contract.BanG(&_Proxy.TransactOpts, _gi, _ban, signs)
}

// BanG is a paid mutator transaction binding the contract method 0x3e54a8b8.
//
// Solidity: function banG(uint64 _gi, bool _ban, bytes[5] signs) returns()
func (_Proxy *ProxyTransactorSession) BanG(_gi uint64, _ban bool, signs [5][]byte) (*types.Transaction, error) {
	return _Proxy.Contract.BanG(&_Proxy.TransactOpts, _gi, _ban, signs)
}

// BanT is a paid mutator transaction binding the contract method 0x7a743984.
//
// Solidity: function banT(address _t, bool _isBan, bytes[5] signs) returns()
func (_Proxy *ProxyTransactor) BanT(opts *bind.TransactOpts, _t common.Address, _isBan bool, signs [5][]byte) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "banT", _t, _isBan, signs)
}

// BanT is a paid mutator transaction binding the contract method 0x7a743984.
//
// Solidity: function banT(address _t, bool _isBan, bytes[5] signs) returns()
func (_Proxy *ProxySession) BanT(_t common.Address, _isBan bool, signs [5][]byte) (*types.Transaction, error) {
	return _Proxy.Contract.BanT(&_Proxy.TransactOpts, _t, _isBan, signs)
}

// BanT is a paid mutator transaction binding the contract method 0x7a743984.
//
// Solidity: function banT(address _t, bool _isBan, bytes[5] signs) returns()
func (_Proxy *ProxyTransactorSession) BanT(_t common.Address, _isBan bool, signs [5][]byte) (*types.Transaction, error) {
	return _Proxy.Contract.BanT(&_Proxy.TransactOpts, _t, _isBan, signs)
}

// ConfirmPayee is a paid mutator transaction binding the contract method 0x310aa0be.
//
// Solidity: function confirmPayee(uint64 _i) returns()
func (_Proxy *ProxyTransactor) ConfirmPayee(opts *bind.TransactOpts, _i uint64) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "confirmPayee", _i)
}

// ConfirmPayee is a paid mutator transaction binding the contract method 0x310aa0be.
//
// Solidity: function confirmPayee(uint64 _i) returns()
func (_Proxy *ProxySession) ConfirmPayee(_i uint64) (*types.Transaction, error) {
	return _Proxy.Contract.ConfirmPayee(&_Proxy.TransactOpts, _i)
}

// ConfirmPayee is a paid mutator transaction binding the contract method 0x310aa0be.
//
// Solidity: function confirmPayee(uint64 _i) returns()
func (_Proxy *ProxyTransactorSession) ConfirmPayee(_i uint64) (*types.Transaction, error) {
	return _Proxy.Contract.ConfirmPayee(&_Proxy.TransactOpts, _i)
}

// CreateGroup is a paid mutator transaction binding the contract method 0x832ec245.
//
// Solidity: function createGroup(uint8 _level, uint8 _mr, uint8 _dc, uint8 _pc, uint256 _kr, uint256 _pr) returns()
func (_Proxy *ProxyTransactor) CreateGroup(opts *bind.TransactOpts, _level uint8, _mr uint8, _dc uint8, _pc uint8, _kr *big.Int, _pr *big.Int) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "createGroup", _level, _mr, _dc, _pc, _kr, _pr)
}

// CreateGroup is a paid mutator transaction binding the contract method 0x832ec245.
//
// Solidity: function createGroup(uint8 _level, uint8 _mr, uint8 _dc, uint8 _pc, uint256 _kr, uint256 _pr) returns()
func (_Proxy *ProxySession) CreateGroup(_level uint8, _mr uint8, _dc uint8, _pc uint8, _kr *big.Int, _pr *big.Int) (*types.Transaction, error) {
	return _Proxy.Contract.CreateGroup(&_Proxy.TransactOpts, _level, _mr, _dc, _pc, _kr, _pr)
}

// CreateGroup is a paid mutator transaction binding the contract method 0x832ec245.
//
// Solidity: function createGroup(uint8 _level, uint8 _mr, uint8 _dc, uint8 _pc, uint256 _kr, uint256 _pr) returns()
func (_Proxy *ProxyTransactorSession) CreateGroup(_level uint8, _mr uint8, _dc uint8, _pc uint8, _kr *big.Int, _pr *big.Int) (*types.Transaction, error) {
	return _Proxy.Contract.CreateGroup(&_Proxy.TransactOpts, _level, _mr, _dc, _pc, _kr, _pr)
}

// Pledge is a paid mutator transaction binding the contract method 0xd23f7482.
//
// Solidity: function pledge(uint64 _i, uint256 _money) returns()
func (_Proxy *ProxyTransactor) Pledge(opts *bind.TransactOpts, _i uint64, _money *big.Int) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "pledge", _i, _money)
}

// Pledge is a paid mutator transaction binding the contract method 0xd23f7482.
//
// Solidity: function pledge(uint64 _i, uint256 _money) returns()
func (_Proxy *ProxySession) Pledge(_i uint64, _money *big.Int) (*types.Transaction, error) {
	return _Proxy.Contract.Pledge(&_Proxy.TransactOpts, _i, _money)
}

// Pledge is a paid mutator transaction binding the contract method 0xd23f7482.
//
// Solidity: function pledge(uint64 _i, uint256 _money) returns()
func (_Proxy *ProxyTransactorSession) Pledge(_i uint64, _money *big.Int) (*types.Transaction, error) {
	return _Proxy.Contract.Pledge(&_Proxy.TransactOpts, _i, _money)
}

// ProWithdraw is a paid mutator transaction binding the contract method 0xe182d54c.
//
// Solidity: function proWithdraw((uint64,uint8,uint256,uint256) _ps, uint64[] _kis, bytes[] ksigns) returns()
func (_Proxy *ProxyTransactor) ProWithdraw(opts *bind.TransactOpts, _ps PWIn, _kis []uint64, ksigns [][]byte) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "proWithdraw", _ps, _kis, ksigns)
}

// ProWithdraw is a paid mutator transaction binding the contract method 0xe182d54c.
//
// Solidity: function proWithdraw((uint64,uint8,uint256,uint256) _ps, uint64[] _kis, bytes[] ksigns) returns()
func (_Proxy *ProxySession) ProWithdraw(_ps PWIn, _kis []uint64, ksigns [][]byte) (*types.Transaction, error) {
	return _Proxy.Contract.ProWithdraw(&_Proxy.TransactOpts, _ps, _kis, ksigns)
}

// ProWithdraw is a paid mutator transaction binding the contract method 0xe182d54c.
//
// Solidity: function proWithdraw((uint64,uint8,uint256,uint256) _ps, uint64[] _kis, bytes[] ksigns) returns()
func (_Proxy *ProxyTransactorSession) ProWithdraw(_ps PWIn, _kis []uint64, ksigns [][]byte) (*types.Transaction, error) {
	return _Proxy.Contract.ProWithdraw(&_Proxy.TransactOpts, _ps, _kis, ksigns)
}

// ReAcc is a paid mutator transaction binding the contract method 0x6526250f.
//
// Solidity: function reAcc() returns()
func (_Proxy *ProxyTransactor) ReAcc(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "reAcc")
}

// ReAcc is a paid mutator transaction binding the contract method 0x6526250f.
//
// Solidity: function reAcc() returns()
func (_Proxy *ProxySession) ReAcc() (*types.Transaction, error) {
	return _Proxy.Contract.ReAcc(&_Proxy.TransactOpts)
}

// ReAcc is a paid mutator transaction binding the contract method 0x6526250f.
//
// Solidity: function reAcc() returns()
func (_Proxy *ProxyTransactorSession) ReAcc() (*types.Transaction, error) {
	return _Proxy.Contract.ReAcc(&_Proxy.TransactOpts)
}

// ReRole is a paid mutator transaction binding the contract method 0x898e00f8.
//
// Solidity: function reRole(uint8 _rtype, bytes _extra) returns()
func (_Proxy *ProxyTransactor) ReRole(opts *bind.TransactOpts, _rtype uint8, _extra []byte) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "reRole", _rtype, _extra)
}

// ReRole is a paid mutator transaction binding the contract method 0x898e00f8.
//
// Solidity: function reRole(uint8 _rtype, bytes _extra) returns()
func (_Proxy *ProxySession) ReRole(_rtype uint8, _extra []byte) (*types.Transaction, error) {
	return _Proxy.Contract.ReRole(&_Proxy.TransactOpts, _rtype, _extra)
}

// ReRole is a paid mutator transaction binding the contract method 0x898e00f8.
//
// Solidity: function reRole(uint8 _rtype, bytes _extra) returns()
func (_Proxy *ProxyTransactorSession) ReRole(_rtype uint8, _extra []byte) (*types.Transaction, error) {
	return _Proxy.Contract.ReRole(&_Proxy.TransactOpts, _rtype, _extra)
}

// Recharge is a paid mutator transaction binding the contract method 0x24d11d40.
//
// Solidity: function recharge(uint64 _i, uint8 _ti, bool isLock, uint256 _money) returns()
func (_Proxy *ProxyTransactor) Recharge(opts *bind.TransactOpts, _i uint64, _ti uint8, isLock bool, _money *big.Int) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "recharge", _i, _ti, isLock, _money)
}

// Recharge is a paid mutator transaction binding the contract method 0x24d11d40.
//
// Solidity: function recharge(uint64 _i, uint8 _ti, bool isLock, uint256 _money) returns()
func (_Proxy *ProxySession) Recharge(_i uint64, _ti uint8, isLock bool, _money *big.Int) (*types.Transaction, error) {
	return _Proxy.Contract.Recharge(&_Proxy.TransactOpts, _i, _ti, isLock, _money)
}

// Recharge is a paid mutator transaction binding the contract method 0x24d11d40.
//
// Solidity: function recharge(uint64 _i, uint8 _ti, bool isLock, uint256 _money) returns()
func (_Proxy *ProxyTransactorSession) Recharge(_i uint64, _ti uint8, isLock bool, _money *big.Int) (*types.Transaction, error) {
	return _Proxy.Contract.Recharge(&_Proxy.TransactOpts, _i, _ti, isLock, _money)
}

// SubOrder is a paid mutator transaction binding the contract method 0x153ede07.
//
// Solidity: function subOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) _oi, bytes uSign, bytes pSign) returns()
func (_Proxy *ProxyTransactor) SubOrder(opts *bind.TransactOpts, _oi OrderIn, uSign []byte, pSign []byte) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "subOrder", _oi, uSign, pSign)
}

// SubOrder is a paid mutator transaction binding the contract method 0x153ede07.
//
// Solidity: function subOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) _oi, bytes uSign, bytes pSign) returns()
func (_Proxy *ProxySession) SubOrder(_oi OrderIn, uSign []byte, pSign []byte) (*types.Transaction, error) {
	return _Proxy.Contract.SubOrder(&_Proxy.TransactOpts, _oi, uSign, pSign)
}

// SubOrder is a paid mutator transaction binding the contract method 0x153ede07.
//
// Solidity: function subOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) _oi, bytes uSign, bytes pSign) returns()
func (_Proxy *ProxyTransactorSession) SubOrder(_oi OrderIn, uSign []byte, pSign []byte) (*types.Transaction, error) {
	return _Proxy.Contract.SubOrder(&_Proxy.TransactOpts, _oi, uSign, pSign)
}

// Unpledge is a paid mutator transaction binding the contract method 0x60985756.
//
// Solidity: function unpledge(uint64 _i, uint8 _ti, uint256 _money) returns()
func (_Proxy *ProxyTransactor) Unpledge(opts *bind.TransactOpts, _i uint64, _ti uint8, _money *big.Int) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "unpledge", _i, _ti, _money)
}

// Unpledge is a paid mutator transaction binding the contract method 0x60985756.
//
// Solidity: function unpledge(uint64 _i, uint8 _ti, uint256 _money) returns()
func (_Proxy *ProxySession) Unpledge(_i uint64, _ti uint8, _money *big.Int) (*types.Transaction, error) {
	return _Proxy.Contract.Unpledge(&_Proxy.TransactOpts, _i, _ti, _money)
}

// Unpledge is a paid mutator transaction binding the contract method 0x60985756.
//
// Solidity: function unpledge(uint64 _i, uint8 _ti, uint256 _money) returns()
func (_Proxy *ProxyTransactorSession) Unpledge(_i uint64, _ti uint8, _money *big.Int) (*types.Transaction, error) {
	return _Proxy.Contract.Unpledge(&_Proxy.TransactOpts, _i, _ti, _money)
}

// Withdraw is a paid mutator transaction binding the contract method 0x259c6d5e.
//
// Solidity: function withdraw(uint64 _i, uint8 _ti, uint256 _money) returns()
func (_Proxy *ProxyTransactor) Withdraw(opts *bind.TransactOpts, _i uint64, _ti uint8, _money *big.Int) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "withdraw", _i, _ti, _money)
}

// Withdraw is a paid mutator transaction binding the contract method 0x259c6d5e.
//
// Solidity: function withdraw(uint64 _i, uint8 _ti, uint256 _money) returns()
func (_Proxy *ProxySession) Withdraw(_i uint64, _ti uint8, _money *big.Int) (*types.Transaction, error) {
	return _Proxy.Contract.Withdraw(&_Proxy.TransactOpts, _i, _ti, _money)
}

// Withdraw is a paid mutator transaction binding the contract method 0x259c6d5e.
//
// Solidity: function withdraw(uint64 _i, uint8 _ti, uint256 _money) returns()
func (_Proxy *ProxyTransactorSession) Withdraw(_i uint64, _ti uint8, _money *big.Int) (*types.Transaction, error) {
	return _Proxy.Contract.Withdraw(&_Proxy.TransactOpts, _i, _ti, _money)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Proxy *ProxyTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proxy.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Proxy *ProxySession) Receive() (*types.Transaction, error) {
	return _Proxy.Contract.Receive(&_Proxy.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Proxy *ProxyTransactorSession) Receive() (*types.Transaction, error) {
	return _Proxy.Contract.Receive(&_Proxy.TransactOpts)
}

// ProxyActivateIterator is returned from FilterActivate and is used to iterate over the raw logs and unpacked data for Activate events raised by the Proxy contract.
type ProxyActivateIterator struct {
	Event *ProxyActivate // Event containing the contract specifics and raw log

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
func (it *ProxyActivateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProxyActivate)
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
		it.Event = new(ProxyActivate)
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
func (it *ProxyActivateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProxyActivateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProxyActivate represents a Activate event raised by the Proxy contract.
type ProxyActivate struct {
	I   uint64
	Raw types.Log // Blockchain specific contextual infos
}

// FilterActivate is a free log retrieval operation binding the contract event 0x452a42fb51ebab9c8ebcbb3d536371397d861478788df33e9d2c9ae57765431c.
//
// Solidity: event Activate(uint64 i)
func (_Proxy *ProxyFilterer) FilterActivate(opts *bind.FilterOpts) (*ProxyActivateIterator, error) {

	logs, sub, err := _Proxy.contract.FilterLogs(opts, "Activate")
	if err != nil {
		return nil, err
	}
	return &ProxyActivateIterator{contract: _Proxy.contract, event: "Activate", logs: logs, sub: sub}, nil
}

// WatchActivate is a free log subscription operation binding the contract event 0x452a42fb51ebab9c8ebcbb3d536371397d861478788df33e9d2c9ae57765431c.
//
// Solidity: event Activate(uint64 i)
func (_Proxy *ProxyFilterer) WatchActivate(opts *bind.WatchOpts, sink chan<- *ProxyActivate) (event.Subscription, error) {

	logs, sub, err := _Proxy.contract.WatchLogs(opts, "Activate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProxyActivate)
				if err := _Proxy.contract.UnpackLog(event, "Activate", log); err != nil {
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

// ParseActivate is a log parse operation binding the contract event 0x452a42fb51ebab9c8ebcbb3d536371397d861478788df33e9d2c9ae57765431c.
//
// Solidity: event Activate(uint64 i)
func (_Proxy *ProxyFilterer) ParseActivate(log types.Log) (*ProxyActivate, error) {
	event := new(ProxyActivate)
	if err := _Proxy.contract.UnpackLog(event, "Activate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProxyAddOwnerIterator is returned from FilterAddOwner and is used to iterate over the raw logs and unpacked data for AddOwner events raised by the Proxy contract.
type ProxyAddOwnerIterator struct {
	Event *ProxyAddOwner // Event containing the contract specifics and raw log

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
func (it *ProxyAddOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProxyAddOwner)
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
		it.Event = new(ProxyAddOwner)
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
func (it *ProxyAddOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProxyAddOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProxyAddOwner represents a AddOwner event raised by the Proxy contract.
type ProxyAddOwner struct {
	From  common.Address
	IsSet bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAddOwner is a free log retrieval operation binding the contract event 0x938b2a24c98e4e542127ffa74a91e48942c2bddccae3b6d75f82bfda7bbc0807.
//
// Solidity: event AddOwner(address from, bool isSet)
func (_Proxy *ProxyFilterer) FilterAddOwner(opts *bind.FilterOpts) (*ProxyAddOwnerIterator, error) {

	logs, sub, err := _Proxy.contract.FilterLogs(opts, "AddOwner")
	if err != nil {
		return nil, err
	}
	return &ProxyAddOwnerIterator{contract: _Proxy.contract, event: "AddOwner", logs: logs, sub: sub}, nil
}

// WatchAddOwner is a free log subscription operation binding the contract event 0x938b2a24c98e4e542127ffa74a91e48942c2bddccae3b6d75f82bfda7bbc0807.
//
// Solidity: event AddOwner(address from, bool isSet)
func (_Proxy *ProxyFilterer) WatchAddOwner(opts *bind.WatchOpts, sink chan<- *ProxyAddOwner) (event.Subscription, error) {

	logs, sub, err := _Proxy.contract.WatchLogs(opts, "AddOwner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProxyAddOwner)
				if err := _Proxy.contract.UnpackLog(event, "AddOwner", log); err != nil {
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
func (_Proxy *ProxyFilterer) ParseAddOwner(log types.Log) (*ProxyAddOwner, error) {
	event := new(ProxyAddOwner)
	if err := _Proxy.contract.UnpackLog(event, "AddOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProxyAddTIterator is returned from FilterAddT and is used to iterate over the raw logs and unpacked data for AddT events raised by the Proxy contract.
type ProxyAddTIterator struct {
	Event *ProxyAddT // Event containing the contract specifics and raw log

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
func (it *ProxyAddTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProxyAddT)
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
		it.Event = new(ProxyAddT)
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
func (it *ProxyAddTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProxyAddTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProxyAddT represents a AddT event raised by the Proxy contract.
type ProxyAddT struct {
	T    common.Address
	Mint bool
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAddT is a free log retrieval operation binding the contract event 0x4954a00722fe171d5c0079b423ebc3b3984d83c24170fdb068dfbd979d250412.
//
// Solidity: event AddT(address t, bool mint)
func (_Proxy *ProxyFilterer) FilterAddT(opts *bind.FilterOpts) (*ProxyAddTIterator, error) {

	logs, sub, err := _Proxy.contract.FilterLogs(opts, "AddT")
	if err != nil {
		return nil, err
	}
	return &ProxyAddTIterator{contract: _Proxy.contract, event: "AddT", logs: logs, sub: sub}, nil
}

// WatchAddT is a free log subscription operation binding the contract event 0x4954a00722fe171d5c0079b423ebc3b3984d83c24170fdb068dfbd979d250412.
//
// Solidity: event AddT(address t, bool mint)
func (_Proxy *ProxyFilterer) WatchAddT(opts *bind.WatchOpts, sink chan<- *ProxyAddT) (event.Subscription, error) {

	logs, sub, err := _Proxy.contract.WatchLogs(opts, "AddT")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProxyAddT)
				if err := _Proxy.contract.UnpackLog(event, "AddT", log); err != nil {
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

// ParseAddT is a log parse operation binding the contract event 0x4954a00722fe171d5c0079b423ebc3b3984d83c24170fdb068dfbd979d250412.
//
// Solidity: event AddT(address t, bool mint)
func (_Proxy *ProxyFilterer) ParseAddT(log types.Log) (*ProxyAddT, error) {
	event := new(ProxyAddT)
	if err := _Proxy.contract.UnpackLog(event, "AddT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProxyBanIterator is returned from FilterBan and is used to iterate over the raw logs and unpacked data for Ban events raised by the Proxy contract.
type ProxyBanIterator struct {
	Event *ProxyBan // Event containing the contract specifics and raw log

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
func (it *ProxyBanIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProxyBan)
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
		it.Event = new(ProxyBan)
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
func (it *ProxyBanIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProxyBanIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProxyBan represents a Ban event raised by the Proxy contract.
type ProxyBan struct {
	I   uint64
	Ban bool
	Raw types.Log // Blockchain specific contextual infos
}

// FilterBan is a free log retrieval operation binding the contract event 0xb5b8b6cac24ab10b3b6a9c41f09df98b5763635f844abf1feb84ae9909902201.
//
// Solidity: event Ban(uint64 i, bool ban)
func (_Proxy *ProxyFilterer) FilterBan(opts *bind.FilterOpts) (*ProxyBanIterator, error) {

	logs, sub, err := _Proxy.contract.FilterLogs(opts, "Ban")
	if err != nil {
		return nil, err
	}
	return &ProxyBanIterator{contract: _Proxy.contract, event: "Ban", logs: logs, sub: sub}, nil
}

// WatchBan is a free log subscription operation binding the contract event 0xb5b8b6cac24ab10b3b6a9c41f09df98b5763635f844abf1feb84ae9909902201.
//
// Solidity: event Ban(uint64 i, bool ban)
func (_Proxy *ProxyFilterer) WatchBan(opts *bind.WatchOpts, sink chan<- *ProxyBan) (event.Subscription, error) {

	logs, sub, err := _Proxy.contract.WatchLogs(opts, "Ban")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProxyBan)
				if err := _Proxy.contract.UnpackLog(event, "Ban", log); err != nil {
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

// ParseBan is a log parse operation binding the contract event 0xb5b8b6cac24ab10b3b6a9c41f09df98b5763635f844abf1feb84ae9909902201.
//
// Solidity: event Ban(uint64 i, bool ban)
func (_Proxy *ProxyFilterer) ParseBan(log types.Log) (*ProxyBan, error) {
	event := new(ProxyBan)
	if err := _Proxy.contract.UnpackLog(event, "Ban", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProxyBanGIterator is returned from FilterBanG and is used to iterate over the raw logs and unpacked data for BanG events raised by the Proxy contract.
type ProxyBanGIterator struct {
	Event *ProxyBanG // Event containing the contract specifics and raw log

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
func (it *ProxyBanGIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProxyBanG)
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
		it.Event = new(ProxyBanG)
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
func (it *ProxyBanGIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProxyBanGIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProxyBanG represents a BanG event raised by the Proxy contract.
type ProxyBanG struct {
	Gi  uint64
	Ban bool
	Raw types.Log // Blockchain specific contextual infos
}

// FilterBanG is a free log retrieval operation binding the contract event 0x976eddb294ef382c026f36e87244e52b5ce40d48ce2020101d563021441aceec.
//
// Solidity: event BanG(uint64 gi, bool ban)
func (_Proxy *ProxyFilterer) FilterBanG(opts *bind.FilterOpts) (*ProxyBanGIterator, error) {

	logs, sub, err := _Proxy.contract.FilterLogs(opts, "BanG")
	if err != nil {
		return nil, err
	}
	return &ProxyBanGIterator{contract: _Proxy.contract, event: "BanG", logs: logs, sub: sub}, nil
}

// WatchBanG is a free log subscription operation binding the contract event 0x976eddb294ef382c026f36e87244e52b5ce40d48ce2020101d563021441aceec.
//
// Solidity: event BanG(uint64 gi, bool ban)
func (_Proxy *ProxyFilterer) WatchBanG(opts *bind.WatchOpts, sink chan<- *ProxyBanG) (event.Subscription, error) {

	logs, sub, err := _Proxy.contract.WatchLogs(opts, "BanG")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProxyBanG)
				if err := _Proxy.contract.UnpackLog(event, "BanG", log); err != nil {
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

// ParseBanG is a log parse operation binding the contract event 0x976eddb294ef382c026f36e87244e52b5ce40d48ce2020101d563021441aceec.
//
// Solidity: event BanG(uint64 gi, bool ban)
func (_Proxy *ProxyFilterer) ParseBanG(log types.Log) (*ProxyBanG, error) {
	event := new(ProxyBanG)
	if err := _Proxy.contract.UnpackLog(event, "BanG", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProxyBanTIterator is returned from FilterBanT and is used to iterate over the raw logs and unpacked data for BanT events raised by the Proxy contract.
type ProxyBanTIterator struct {
	Event *ProxyBanT // Event containing the contract specifics and raw log

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
func (it *ProxyBanTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProxyBanT)
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
		it.Event = new(ProxyBanT)
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
func (it *ProxyBanTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProxyBanTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProxyBanT represents a BanT event raised by the Proxy contract.
type ProxyBanT struct {
	T   common.Address
	Ban bool
	Raw types.Log // Blockchain specific contextual infos
}

// FilterBanT is a free log retrieval operation binding the contract event 0xffaa091f00488e7d9934b8efc3cbe385e3bb89957862cb891e428474b69b6bc9.
//
// Solidity: event BanT(address t, bool ban)
func (_Proxy *ProxyFilterer) FilterBanT(opts *bind.FilterOpts) (*ProxyBanTIterator, error) {

	logs, sub, err := _Proxy.contract.FilterLogs(opts, "BanT")
	if err != nil {
		return nil, err
	}
	return &ProxyBanTIterator{contract: _Proxy.contract, event: "BanT", logs: logs, sub: sub}, nil
}

// WatchBanT is a free log subscription operation binding the contract event 0xffaa091f00488e7d9934b8efc3cbe385e3bb89957862cb891e428474b69b6bc9.
//
// Solidity: event BanT(address t, bool ban)
func (_Proxy *ProxyFilterer) WatchBanT(opts *bind.WatchOpts, sink chan<- *ProxyBanT) (event.Subscription, error) {

	logs, sub, err := _Proxy.contract.WatchLogs(opts, "BanT")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProxyBanT)
				if err := _Proxy.contract.UnpackLog(event, "BanT", log); err != nil {
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

// ParseBanT is a log parse operation binding the contract event 0xffaa091f00488e7d9934b8efc3cbe385e3bb89957862cb891e428474b69b6bc9.
//
// Solidity: event BanT(address t, bool ban)
func (_Proxy *ProxyFilterer) ParseBanT(log types.Log) (*ProxyBanT, error) {
	event := new(ProxyBanT)
	if err := _Proxy.contract.UnpackLog(event, "BanT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
