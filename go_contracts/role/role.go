// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package role

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

// RoleOut is an auto generated low-level Go binding around an user-defined struct.
type RoleOut struct {
	IsBan    bool
	IsActive bool
	RType    uint8
	Index    uint64
	GIndex   uint64
	Owner    common.Address
	Next     common.Address
	Extra    []byte
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

// IRoleABI is the input ABI used to generate the binding from.
const IRoleABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"AlterPayee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"ConfirmPayee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"}],\"name\":\"ReAcc\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"rType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"}],\"name\":\"ReRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"gi\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"}],\"name\":\"SetG\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"}],\"name\":\"activate\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_index\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"_p\",\"type\":\"address\"}],\"name\":\"alterPayee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"_ban\",\"type\":\"bool\"}],\"name\":\"ban\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_rType\",\"type\":\"uint8\"}],\"name\":\"checkIG\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_index\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"_p\",\"type\":\"address\"}],\"name\":\"confirmPayee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getACnt\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"}],\"name\":\"getAcc\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"}],\"name\":\"getAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"}],\"name\":\"getRInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isBan\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"rType\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gIndex\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"next\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extra\",\"type\":\"bytes\"}],\"internalType\":\"structRoleOut\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"}],\"name\":\"reAcc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_rtype\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"extra\",\"type\":\"bytes\"}],\"name\":\"reRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"}],\"name\":\"setG\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IRoleFuncSigs maps the 4-byte function signature to its string representation.
var IRoleFuncSigs = map[string]string{
	"fc7e4746": "activate(uint64)",
	"2915824a": "alterPayee(uint64,address)",
	"0f78c61a": "ban(uint64,bool)",
	"7738515f": "checkIG(uint64,uint8)",
	"5bd04c3f": "confirmPayee(uint64,address)",
	"7264a551": "getACnt()",
	"caca4a06": "getAcc(address)",
	"9332aa6e": "getAddr(uint64)",
	"441abace": "getRInfo(address)",
	"effcafce": "reAcc(address)",
	"b9f6a8ca": "reRole(uint64,uint8,bytes)",
	"616bfd1f": "setG(uint64,uint64)",
}

// IRole is an auto generated Go binding around an Ethereum contract.
type IRole struct {
	IRoleCaller     // Read-only binding to the contract
	IRoleTransactor // Write-only binding to the contract
	IRoleFilterer   // Log filterer for contract events
}

// IRoleCaller is an auto generated read-only Go binding around an Ethereum contract.
type IRoleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRoleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IRoleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRoleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IRoleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRoleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IRoleSession struct {
	Contract     *IRole            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IRoleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IRoleCallerSession struct {
	Contract *IRoleCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IRoleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IRoleTransactorSession struct {
	Contract     *IRoleTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IRoleRaw is an auto generated low-level Go binding around an Ethereum contract.
type IRoleRaw struct {
	Contract *IRole // Generic contract binding to access the raw methods on
}

// IRoleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IRoleCallerRaw struct {
	Contract *IRoleCaller // Generic read-only contract binding to access the raw methods on
}

// IRoleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IRoleTransactorRaw struct {
	Contract *IRoleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIRole creates a new instance of IRole, bound to a specific deployed contract.
func NewIRole(address common.Address, backend bind.ContractBackend) (*IRole, error) {
	contract, err := bindIRole(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IRole{IRoleCaller: IRoleCaller{contract: contract}, IRoleTransactor: IRoleTransactor{contract: contract}, IRoleFilterer: IRoleFilterer{contract: contract}}, nil
}

// NewIRoleCaller creates a new read-only instance of IRole, bound to a specific deployed contract.
func NewIRoleCaller(address common.Address, caller bind.ContractCaller) (*IRoleCaller, error) {
	contract, err := bindIRole(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IRoleCaller{contract: contract}, nil
}

// NewIRoleTransactor creates a new write-only instance of IRole, bound to a specific deployed contract.
func NewIRoleTransactor(address common.Address, transactor bind.ContractTransactor) (*IRoleTransactor, error) {
	contract, err := bindIRole(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IRoleTransactor{contract: contract}, nil
}

// NewIRoleFilterer creates a new log filterer instance of IRole, bound to a specific deployed contract.
func NewIRoleFilterer(address common.Address, filterer bind.ContractFilterer) (*IRoleFilterer, error) {
	contract, err := bindIRole(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IRoleFilterer{contract: contract}, nil
}

// bindIRole binds a generic wrapper to an already deployed contract.
func bindIRole(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IRoleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRole *IRoleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRole.Contract.IRoleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRole *IRoleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRole.Contract.IRoleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRole *IRoleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRole.Contract.IRoleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRole *IRoleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRole.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRole *IRoleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRole.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRole *IRoleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRole.Contract.contract.Transact(opts, method, params...)
}

// CheckIG is a free data retrieval call binding the contract method 0x7738515f.
//
// Solidity: function checkIG(uint64 _i, uint8 _rType) view returns(address, address, uint64, uint8)
func (_IRole *IRoleCaller) CheckIG(opts *bind.CallOpts, _i uint64, _rType uint8) (common.Address, common.Address, uint64, uint8, error) {
	var out []interface{}
	err := _IRole.contract.Call(opts, &out, "checkIG", _i, _rType)

	if err != nil {
		return *new(common.Address), *new(common.Address), *new(uint64), *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	out2 := *abi.ConvertType(out[2], new(uint64)).(*uint64)
	out3 := *abi.ConvertType(out[3], new(uint8)).(*uint8)

	return out0, out1, out2, out3, err

}

// CheckIG is a free data retrieval call binding the contract method 0x7738515f.
//
// Solidity: function checkIG(uint64 _i, uint8 _rType) view returns(address, address, uint64, uint8)
func (_IRole *IRoleSession) CheckIG(_i uint64, _rType uint8) (common.Address, common.Address, uint64, uint8, error) {
	return _IRole.Contract.CheckIG(&_IRole.CallOpts, _i, _rType)
}

// CheckIG is a free data retrieval call binding the contract method 0x7738515f.
//
// Solidity: function checkIG(uint64 _i, uint8 _rType) view returns(address, address, uint64, uint8)
func (_IRole *IRoleCallerSession) CheckIG(_i uint64, _rType uint8) (common.Address, common.Address, uint64, uint8, error) {
	return _IRole.Contract.CheckIG(&_IRole.CallOpts, _i, _rType)
}

// GetACnt is a free data retrieval call binding the contract method 0x7264a551.
//
// Solidity: function getACnt() view returns(uint64)
func (_IRole *IRoleCaller) GetACnt(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _IRole.contract.Call(opts, &out, "getACnt")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetACnt is a free data retrieval call binding the contract method 0x7264a551.
//
// Solidity: function getACnt() view returns(uint64)
func (_IRole *IRoleSession) GetACnt() (uint64, error) {
	return _IRole.Contract.GetACnt(&_IRole.CallOpts)
}

// GetACnt is a free data retrieval call binding the contract method 0x7264a551.
//
// Solidity: function getACnt() view returns(uint64)
func (_IRole *IRoleCallerSession) GetACnt() (uint64, error) {
	return _IRole.Contract.GetACnt(&_IRole.CallOpts)
}

// GetAcc is a free data retrieval call binding the contract method 0xcaca4a06.
//
// Solidity: function getAcc(address _a) view returns(uint64)
func (_IRole *IRoleCaller) GetAcc(opts *bind.CallOpts, _a common.Address) (uint64, error) {
	var out []interface{}
	err := _IRole.contract.Call(opts, &out, "getAcc", _a)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetAcc is a free data retrieval call binding the contract method 0xcaca4a06.
//
// Solidity: function getAcc(address _a) view returns(uint64)
func (_IRole *IRoleSession) GetAcc(_a common.Address) (uint64, error) {
	return _IRole.Contract.GetAcc(&_IRole.CallOpts, _a)
}

// GetAcc is a free data retrieval call binding the contract method 0xcaca4a06.
//
// Solidity: function getAcc(address _a) view returns(uint64)
func (_IRole *IRoleCallerSession) GetAcc(_a common.Address) (uint64, error) {
	return _IRole.Contract.GetAcc(&_IRole.CallOpts, _a)
}

// GetAddr is a free data retrieval call binding the contract method 0x9332aa6e.
//
// Solidity: function getAddr(uint64 _i) view returns(address)
func (_IRole *IRoleCaller) GetAddr(opts *bind.CallOpts, _i uint64) (common.Address, error) {
	var out []interface{}
	err := _IRole.contract.Call(opts, &out, "getAddr", _i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddr is a free data retrieval call binding the contract method 0x9332aa6e.
//
// Solidity: function getAddr(uint64 _i) view returns(address)
func (_IRole *IRoleSession) GetAddr(_i uint64) (common.Address, error) {
	return _IRole.Contract.GetAddr(&_IRole.CallOpts, _i)
}

// GetAddr is a free data retrieval call binding the contract method 0x9332aa6e.
//
// Solidity: function getAddr(uint64 _i) view returns(address)
func (_IRole *IRoleCallerSession) GetAddr(_i uint64) (common.Address, error) {
	return _IRole.Contract.GetAddr(&_IRole.CallOpts, _i)
}

// GetRInfo is a free data retrieval call binding the contract method 0x441abace.
//
// Solidity: function getRInfo(address _a) view returns((bool,bool,uint8,uint64,uint64,address,address,bytes))
func (_IRole *IRoleCaller) GetRInfo(opts *bind.CallOpts, _a common.Address) (RoleOut, error) {
	var out []interface{}
	err := _IRole.contract.Call(opts, &out, "getRInfo", _a)

	if err != nil {
		return *new(RoleOut), err
	}

	out0 := *abi.ConvertType(out[0], new(RoleOut)).(*RoleOut)

	return out0, err

}

// GetRInfo is a free data retrieval call binding the contract method 0x441abace.
//
// Solidity: function getRInfo(address _a) view returns((bool,bool,uint8,uint64,uint64,address,address,bytes))
func (_IRole *IRoleSession) GetRInfo(_a common.Address) (RoleOut, error) {
	return _IRole.Contract.GetRInfo(&_IRole.CallOpts, _a)
}

// GetRInfo is a free data retrieval call binding the contract method 0x441abace.
//
// Solidity: function getRInfo(address _a) view returns((bool,bool,uint8,uint64,uint64,address,address,bytes))
func (_IRole *IRoleCallerSession) GetRInfo(_a common.Address) (RoleOut, error) {
	return _IRole.Contract.GetRInfo(&_IRole.CallOpts, _a)
}

// Activate is a paid mutator transaction binding the contract method 0xfc7e4746.
//
// Solidity: function activate(uint64 _i) returns(uint64)
func (_IRole *IRoleTransactor) Activate(opts *bind.TransactOpts, _i uint64) (*types.Transaction, error) {
	return _IRole.contract.Transact(opts, "activate", _i)
}

// Activate is a paid mutator transaction binding the contract method 0xfc7e4746.
//
// Solidity: function activate(uint64 _i) returns(uint64)
func (_IRole *IRoleSession) Activate(_i uint64) (*types.Transaction, error) {
	return _IRole.Contract.Activate(&_IRole.TransactOpts, _i)
}

// Activate is a paid mutator transaction binding the contract method 0xfc7e4746.
//
// Solidity: function activate(uint64 _i) returns(uint64)
func (_IRole *IRoleTransactorSession) Activate(_i uint64) (*types.Transaction, error) {
	return _IRole.Contract.Activate(&_IRole.TransactOpts, _i)
}

// AlterPayee is a paid mutator transaction binding the contract method 0x2915824a.
//
// Solidity: function alterPayee(uint64 _index, address _p) returns()
func (_IRole *IRoleTransactor) AlterPayee(opts *bind.TransactOpts, _index uint64, _p common.Address) (*types.Transaction, error) {
	return _IRole.contract.Transact(opts, "alterPayee", _index, _p)
}

// AlterPayee is a paid mutator transaction binding the contract method 0x2915824a.
//
// Solidity: function alterPayee(uint64 _index, address _p) returns()
func (_IRole *IRoleSession) AlterPayee(_index uint64, _p common.Address) (*types.Transaction, error) {
	return _IRole.Contract.AlterPayee(&_IRole.TransactOpts, _index, _p)
}

// AlterPayee is a paid mutator transaction binding the contract method 0x2915824a.
//
// Solidity: function alterPayee(uint64 _index, address _p) returns()
func (_IRole *IRoleTransactorSession) AlterPayee(_index uint64, _p common.Address) (*types.Transaction, error) {
	return _IRole.Contract.AlterPayee(&_IRole.TransactOpts, _index, _p)
}

// Ban is a paid mutator transaction binding the contract method 0x0f78c61a.
//
// Solidity: function ban(uint64 _i, bool _ban) returns()
func (_IRole *IRoleTransactor) Ban(opts *bind.TransactOpts, _i uint64, _ban bool) (*types.Transaction, error) {
	return _IRole.contract.Transact(opts, "ban", _i, _ban)
}

// Ban is a paid mutator transaction binding the contract method 0x0f78c61a.
//
// Solidity: function ban(uint64 _i, bool _ban) returns()
func (_IRole *IRoleSession) Ban(_i uint64, _ban bool) (*types.Transaction, error) {
	return _IRole.Contract.Ban(&_IRole.TransactOpts, _i, _ban)
}

// Ban is a paid mutator transaction binding the contract method 0x0f78c61a.
//
// Solidity: function ban(uint64 _i, bool _ban) returns()
func (_IRole *IRoleTransactorSession) Ban(_i uint64, _ban bool) (*types.Transaction, error) {
	return _IRole.Contract.Ban(&_IRole.TransactOpts, _i, _ban)
}

// ConfirmPayee is a paid mutator transaction binding the contract method 0x5bd04c3f.
//
// Solidity: function confirmPayee(uint64 _index, address _p) returns()
func (_IRole *IRoleTransactor) ConfirmPayee(opts *bind.TransactOpts, _index uint64, _p common.Address) (*types.Transaction, error) {
	return _IRole.contract.Transact(opts, "confirmPayee", _index, _p)
}

// ConfirmPayee is a paid mutator transaction binding the contract method 0x5bd04c3f.
//
// Solidity: function confirmPayee(uint64 _index, address _p) returns()
func (_IRole *IRoleSession) ConfirmPayee(_index uint64, _p common.Address) (*types.Transaction, error) {
	return _IRole.Contract.ConfirmPayee(&_IRole.TransactOpts, _index, _p)
}

// ConfirmPayee is a paid mutator transaction binding the contract method 0x5bd04c3f.
//
// Solidity: function confirmPayee(uint64 _index, address _p) returns()
func (_IRole *IRoleTransactorSession) ConfirmPayee(_index uint64, _p common.Address) (*types.Transaction, error) {
	return _IRole.Contract.ConfirmPayee(&_IRole.TransactOpts, _index, _p)
}

// ReAcc is a paid mutator transaction binding the contract method 0xeffcafce.
//
// Solidity: function reAcc(address _a) returns()
func (_IRole *IRoleTransactor) ReAcc(opts *bind.TransactOpts, _a common.Address) (*types.Transaction, error) {
	return _IRole.contract.Transact(opts, "reAcc", _a)
}

// ReAcc is a paid mutator transaction binding the contract method 0xeffcafce.
//
// Solidity: function reAcc(address _a) returns()
func (_IRole *IRoleSession) ReAcc(_a common.Address) (*types.Transaction, error) {
	return _IRole.Contract.ReAcc(&_IRole.TransactOpts, _a)
}

// ReAcc is a paid mutator transaction binding the contract method 0xeffcafce.
//
// Solidity: function reAcc(address _a) returns()
func (_IRole *IRoleTransactorSession) ReAcc(_a common.Address) (*types.Transaction, error) {
	return _IRole.Contract.ReAcc(&_IRole.TransactOpts, _a)
}

// ReRole is a paid mutator transaction binding the contract method 0xb9f6a8ca.
//
// Solidity: function reRole(uint64 _i, uint8 _rtype, bytes extra) returns()
func (_IRole *IRoleTransactor) ReRole(opts *bind.TransactOpts, _i uint64, _rtype uint8, extra []byte) (*types.Transaction, error) {
	return _IRole.contract.Transact(opts, "reRole", _i, _rtype, extra)
}

// ReRole is a paid mutator transaction binding the contract method 0xb9f6a8ca.
//
// Solidity: function reRole(uint64 _i, uint8 _rtype, bytes extra) returns()
func (_IRole *IRoleSession) ReRole(_i uint64, _rtype uint8, extra []byte) (*types.Transaction, error) {
	return _IRole.Contract.ReRole(&_IRole.TransactOpts, _i, _rtype, extra)
}

// ReRole is a paid mutator transaction binding the contract method 0xb9f6a8ca.
//
// Solidity: function reRole(uint64 _i, uint8 _rtype, bytes extra) returns()
func (_IRole *IRoleTransactorSession) ReRole(_i uint64, _rtype uint8, extra []byte) (*types.Transaction, error) {
	return _IRole.Contract.ReRole(&_IRole.TransactOpts, _i, _rtype, extra)
}

// SetG is a paid mutator transaction binding the contract method 0x616bfd1f.
//
// Solidity: function setG(uint64 _i, uint64 _gi) returns()
func (_IRole *IRoleTransactor) SetG(opts *bind.TransactOpts, _i uint64, _gi uint64) (*types.Transaction, error) {
	return _IRole.contract.Transact(opts, "setG", _i, _gi)
}

// SetG is a paid mutator transaction binding the contract method 0x616bfd1f.
//
// Solidity: function setG(uint64 _i, uint64 _gi) returns()
func (_IRole *IRoleSession) SetG(_i uint64, _gi uint64) (*types.Transaction, error) {
	return _IRole.Contract.SetG(&_IRole.TransactOpts, _i, _gi)
}

// SetG is a paid mutator transaction binding the contract method 0x616bfd1f.
//
// Solidity: function setG(uint64 _i, uint64 _gi) returns()
func (_IRole *IRoleTransactorSession) SetG(_i uint64, _gi uint64) (*types.Transaction, error) {
	return _IRole.Contract.SetG(&_IRole.TransactOpts, _i, _gi)
}

// IRoleAlterPayeeIterator is returned from FilterAlterPayee and is used to iterate over the raw logs and unpacked data for AlterPayee events raised by the IRole contract.
type IRoleAlterPayeeIterator struct {
	Event *IRoleAlterPayee // Event containing the contract specifics and raw log

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
func (it *IRoleAlterPayeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IRoleAlterPayee)
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
		it.Event = new(IRoleAlterPayee)
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
func (it *IRoleAlterPayeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IRoleAlterPayeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IRoleAlterPayee represents a AlterPayee event raised by the IRole contract.
type IRoleAlterPayee struct {
	Payee common.Address
	Addr  common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAlterPayee is a free log retrieval operation binding the contract event 0x3b4058e93b2a019add56609a3c2463a7b3d8476abf30c934cfcdb9f2158b5e68.
//
// Solidity: event AlterPayee(address indexed payee, address addr)
func (_IRole *IRoleFilterer) FilterAlterPayee(opts *bind.FilterOpts, payee []common.Address) (*IRoleAlterPayeeIterator, error) {

	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}

	logs, sub, err := _IRole.contract.FilterLogs(opts, "AlterPayee", payeeRule)
	if err != nil {
		return nil, err
	}
	return &IRoleAlterPayeeIterator{contract: _IRole.contract, event: "AlterPayee", logs: logs, sub: sub}, nil
}

// WatchAlterPayee is a free log subscription operation binding the contract event 0x3b4058e93b2a019add56609a3c2463a7b3d8476abf30c934cfcdb9f2158b5e68.
//
// Solidity: event AlterPayee(address indexed payee, address addr)
func (_IRole *IRoleFilterer) WatchAlterPayee(opts *bind.WatchOpts, sink chan<- *IRoleAlterPayee, payee []common.Address) (event.Subscription, error) {

	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}

	logs, sub, err := _IRole.contract.WatchLogs(opts, "AlterPayee", payeeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IRoleAlterPayee)
				if err := _IRole.contract.UnpackLog(event, "AlterPayee", log); err != nil {
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

// ParseAlterPayee is a log parse operation binding the contract event 0x3b4058e93b2a019add56609a3c2463a7b3d8476abf30c934cfcdb9f2158b5e68.
//
// Solidity: event AlterPayee(address indexed payee, address addr)
func (_IRole *IRoleFilterer) ParseAlterPayee(log types.Log) (*IRoleAlterPayee, error) {
	event := new(IRoleAlterPayee)
	if err := _IRole.contract.UnpackLog(event, "AlterPayee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IRoleConfirmPayeeIterator is returned from FilterConfirmPayee and is used to iterate over the raw logs and unpacked data for ConfirmPayee events raised by the IRole contract.
type IRoleConfirmPayeeIterator struct {
	Event *IRoleConfirmPayee // Event containing the contract specifics and raw log

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
func (it *IRoleConfirmPayeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IRoleConfirmPayee)
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
		it.Event = new(IRoleConfirmPayee)
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
func (it *IRoleConfirmPayeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IRoleConfirmPayeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IRoleConfirmPayee represents a ConfirmPayee event raised by the IRole contract.
type IRoleConfirmPayee struct {
	Payee common.Address
	Addr  common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterConfirmPayee is a free log retrieval operation binding the contract event 0x4a21d873c33448e8dfb90cdbfdb12f4aed5f5e2abde5525215b1655d519784dc.
//
// Solidity: event ConfirmPayee(address indexed payee, address addr)
func (_IRole *IRoleFilterer) FilterConfirmPayee(opts *bind.FilterOpts, payee []common.Address) (*IRoleConfirmPayeeIterator, error) {

	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}

	logs, sub, err := _IRole.contract.FilterLogs(opts, "ConfirmPayee", payeeRule)
	if err != nil {
		return nil, err
	}
	return &IRoleConfirmPayeeIterator{contract: _IRole.contract, event: "ConfirmPayee", logs: logs, sub: sub}, nil
}

// WatchConfirmPayee is a free log subscription operation binding the contract event 0x4a21d873c33448e8dfb90cdbfdb12f4aed5f5e2abde5525215b1655d519784dc.
//
// Solidity: event ConfirmPayee(address indexed payee, address addr)
func (_IRole *IRoleFilterer) WatchConfirmPayee(opts *bind.WatchOpts, sink chan<- *IRoleConfirmPayee, payee []common.Address) (event.Subscription, error) {

	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}

	logs, sub, err := _IRole.contract.WatchLogs(opts, "ConfirmPayee", payeeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IRoleConfirmPayee)
				if err := _IRole.contract.UnpackLog(event, "ConfirmPayee", log); err != nil {
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

// ParseConfirmPayee is a log parse operation binding the contract event 0x4a21d873c33448e8dfb90cdbfdb12f4aed5f5e2abde5525215b1655d519784dc.
//
// Solidity: event ConfirmPayee(address indexed payee, address addr)
func (_IRole *IRoleFilterer) ParseConfirmPayee(log types.Log) (*IRoleConfirmPayee, error) {
	event := new(IRoleConfirmPayee)
	if err := _IRole.contract.UnpackLog(event, "ConfirmPayee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IRoleReAccIterator is returned from FilterReAcc and is used to iterate over the raw logs and unpacked data for ReAcc events raised by the IRole contract.
type IRoleReAccIterator struct {
	Event *IRoleReAcc // Event containing the contract specifics and raw log

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
func (it *IRoleReAccIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IRoleReAcc)
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
		it.Event = new(IRoleReAcc)
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
func (it *IRoleReAccIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IRoleReAccIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IRoleReAcc represents a ReAcc event raised by the IRole contract.
type IRoleReAcc struct {
	Addr  common.Address
	Index uint64
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterReAcc is a free log retrieval operation binding the contract event 0xe22ccfc11bbbb1d3350012248b26ad616e32a8a13eece408dd14f1953fe24752.
//
// Solidity: event ReAcc(address addr, uint64 index)
func (_IRole *IRoleFilterer) FilterReAcc(opts *bind.FilterOpts) (*IRoleReAccIterator, error) {

	logs, sub, err := _IRole.contract.FilterLogs(opts, "ReAcc")
	if err != nil {
		return nil, err
	}
	return &IRoleReAccIterator{contract: _IRole.contract, event: "ReAcc", logs: logs, sub: sub}, nil
}

// WatchReAcc is a free log subscription operation binding the contract event 0xe22ccfc11bbbb1d3350012248b26ad616e32a8a13eece408dd14f1953fe24752.
//
// Solidity: event ReAcc(address addr, uint64 index)
func (_IRole *IRoleFilterer) WatchReAcc(opts *bind.WatchOpts, sink chan<- *IRoleReAcc) (event.Subscription, error) {

	logs, sub, err := _IRole.contract.WatchLogs(opts, "ReAcc")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IRoleReAcc)
				if err := _IRole.contract.UnpackLog(event, "ReAcc", log); err != nil {
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

// ParseReAcc is a log parse operation binding the contract event 0xe22ccfc11bbbb1d3350012248b26ad616e32a8a13eece408dd14f1953fe24752.
//
// Solidity: event ReAcc(address addr, uint64 index)
func (_IRole *IRoleFilterer) ParseReAcc(log types.Log) (*IRoleReAcc, error) {
	event := new(IRoleReAcc)
	if err := _IRole.contract.UnpackLog(event, "ReAcc", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IRoleReRoleIterator is returned from FilterReRole and is used to iterate over the raw logs and unpacked data for ReRole events raised by the IRole contract.
type IRoleReRoleIterator struct {
	Event *IRoleReRole // Event containing the contract specifics and raw log

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
func (it *IRoleReRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IRoleReRole)
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
		it.Event = new(IRoleReRole)
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
func (it *IRoleReRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IRoleReRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IRoleReRole represents a ReRole event raised by the IRole contract.
type IRoleReRole struct {
	RType uint8
	Index uint64
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterReRole is a free log retrieval operation binding the contract event 0x5550fbab402c3b014955818e5b6e3cefef20641aa6bed43ae83cd5d3258b8922.
//
// Solidity: event ReRole(uint8 indexed rType, uint64 index)
func (_IRole *IRoleFilterer) FilterReRole(opts *bind.FilterOpts, rType []uint8) (*IRoleReRoleIterator, error) {

	var rTypeRule []interface{}
	for _, rTypeItem := range rType {
		rTypeRule = append(rTypeRule, rTypeItem)
	}

	logs, sub, err := _IRole.contract.FilterLogs(opts, "ReRole", rTypeRule)
	if err != nil {
		return nil, err
	}
	return &IRoleReRoleIterator{contract: _IRole.contract, event: "ReRole", logs: logs, sub: sub}, nil
}

// WatchReRole is a free log subscription operation binding the contract event 0x5550fbab402c3b014955818e5b6e3cefef20641aa6bed43ae83cd5d3258b8922.
//
// Solidity: event ReRole(uint8 indexed rType, uint64 index)
func (_IRole *IRoleFilterer) WatchReRole(opts *bind.WatchOpts, sink chan<- *IRoleReRole, rType []uint8) (event.Subscription, error) {

	var rTypeRule []interface{}
	for _, rTypeItem := range rType {
		rTypeRule = append(rTypeRule, rTypeItem)
	}

	logs, sub, err := _IRole.contract.WatchLogs(opts, "ReRole", rTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IRoleReRole)
				if err := _IRole.contract.UnpackLog(event, "ReRole", log); err != nil {
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

// ParseReRole is a log parse operation binding the contract event 0x5550fbab402c3b014955818e5b6e3cefef20641aa6bed43ae83cd5d3258b8922.
//
// Solidity: event ReRole(uint8 indexed rType, uint64 index)
func (_IRole *IRoleFilterer) ParseReRole(log types.Log) (*IRoleReRole, error) {
	event := new(IRoleReRole)
	if err := _IRole.contract.UnpackLog(event, "ReRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IRoleSetGIterator is returned from FilterSetG and is used to iterate over the raw logs and unpacked data for SetG events raised by the IRole contract.
type IRoleSetGIterator struct {
	Event *IRoleSetG // Event containing the contract specifics and raw log

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
func (it *IRoleSetGIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IRoleSetG)
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
		it.Event = new(IRoleSetG)
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
func (it *IRoleSetGIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IRoleSetGIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IRoleSetG represents a SetG event raised by the IRole contract.
type IRoleSetG struct {
	Gi    uint64
	Index uint64
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSetG is a free log retrieval operation binding the contract event 0xe7329265e8e4e2fc8e892a75f1809767d09c472bb2eef16f889eceb5583291a9.
//
// Solidity: event SetG(uint64 indexed gi, uint64 index)
func (_IRole *IRoleFilterer) FilterSetG(opts *bind.FilterOpts, gi []uint64) (*IRoleSetGIterator, error) {

	var giRule []interface{}
	for _, giItem := range gi {
		giRule = append(giRule, giItem)
	}

	logs, sub, err := _IRole.contract.FilterLogs(opts, "SetG", giRule)
	if err != nil {
		return nil, err
	}
	return &IRoleSetGIterator{contract: _IRole.contract, event: "SetG", logs: logs, sub: sub}, nil
}

// WatchSetG is a free log subscription operation binding the contract event 0xe7329265e8e4e2fc8e892a75f1809767d09c472bb2eef16f889eceb5583291a9.
//
// Solidity: event SetG(uint64 indexed gi, uint64 index)
func (_IRole *IRoleFilterer) WatchSetG(opts *bind.WatchOpts, sink chan<- *IRoleSetG, gi []uint64) (event.Subscription, error) {

	var giRule []interface{}
	for _, giItem := range gi {
		giRule = append(giRule, giItem)
	}

	logs, sub, err := _IRole.contract.WatchLogs(opts, "SetG", giRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IRoleSetG)
				if err := _IRole.contract.UnpackLog(event, "SetG", log); err != nil {
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

// ParseSetG is a log parse operation binding the contract event 0xe7329265e8e4e2fc8e892a75f1809767d09c472bb2eef16f889eceb5583291a9.
//
// Solidity: event SetG(uint64 indexed gi, uint64 index)
func (_IRole *IRoleFilterer) ParseSetG(log types.Log) (*IRoleSetG, error) {
	event := new(IRoleSetG)
	if err := _IRole.contract.UnpackLog(event, "SetG", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IRoleGetterABI is the input ABI used to generate the binding from.
const IRoleGetterABI = "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_rType\",\"type\":\"uint8\"}],\"name\":\"checkIG\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getACnt\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"}],\"name\":\"getAcc\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"}],\"name\":\"getAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"}],\"name\":\"getRInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isBan\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"rType\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gIndex\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"next\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extra\",\"type\":\"bytes\"}],\"internalType\":\"structRoleOut\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IRoleGetterFuncSigs maps the 4-byte function signature to its string representation.
var IRoleGetterFuncSigs = map[string]string{
	"7738515f": "checkIG(uint64,uint8)",
	"7264a551": "getACnt()",
	"caca4a06": "getAcc(address)",
	"9332aa6e": "getAddr(uint64)",
	"441abace": "getRInfo(address)",
}

// IRoleGetter is an auto generated Go binding around an Ethereum contract.
type IRoleGetter struct {
	IRoleGetterCaller     // Read-only binding to the contract
	IRoleGetterTransactor // Write-only binding to the contract
	IRoleGetterFilterer   // Log filterer for contract events
}

// IRoleGetterCaller is an auto generated read-only Go binding around an Ethereum contract.
type IRoleGetterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRoleGetterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IRoleGetterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRoleGetterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IRoleGetterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRoleGetterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IRoleGetterSession struct {
	Contract     *IRoleGetter      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IRoleGetterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IRoleGetterCallerSession struct {
	Contract *IRoleGetterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// IRoleGetterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IRoleGetterTransactorSession struct {
	Contract     *IRoleGetterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IRoleGetterRaw is an auto generated low-level Go binding around an Ethereum contract.
type IRoleGetterRaw struct {
	Contract *IRoleGetter // Generic contract binding to access the raw methods on
}

// IRoleGetterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IRoleGetterCallerRaw struct {
	Contract *IRoleGetterCaller // Generic read-only contract binding to access the raw methods on
}

// IRoleGetterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IRoleGetterTransactorRaw struct {
	Contract *IRoleGetterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIRoleGetter creates a new instance of IRoleGetter, bound to a specific deployed contract.
func NewIRoleGetter(address common.Address, backend bind.ContractBackend) (*IRoleGetter, error) {
	contract, err := bindIRoleGetter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IRoleGetter{IRoleGetterCaller: IRoleGetterCaller{contract: contract}, IRoleGetterTransactor: IRoleGetterTransactor{contract: contract}, IRoleGetterFilterer: IRoleGetterFilterer{contract: contract}}, nil
}

// NewIRoleGetterCaller creates a new read-only instance of IRoleGetter, bound to a specific deployed contract.
func NewIRoleGetterCaller(address common.Address, caller bind.ContractCaller) (*IRoleGetterCaller, error) {
	contract, err := bindIRoleGetter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IRoleGetterCaller{contract: contract}, nil
}

// NewIRoleGetterTransactor creates a new write-only instance of IRoleGetter, bound to a specific deployed contract.
func NewIRoleGetterTransactor(address common.Address, transactor bind.ContractTransactor) (*IRoleGetterTransactor, error) {
	contract, err := bindIRoleGetter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IRoleGetterTransactor{contract: contract}, nil
}

// NewIRoleGetterFilterer creates a new log filterer instance of IRoleGetter, bound to a specific deployed contract.
func NewIRoleGetterFilterer(address common.Address, filterer bind.ContractFilterer) (*IRoleGetterFilterer, error) {
	contract, err := bindIRoleGetter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IRoleGetterFilterer{contract: contract}, nil
}

// bindIRoleGetter binds a generic wrapper to an already deployed contract.
func bindIRoleGetter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IRoleGetterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRoleGetter *IRoleGetterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRoleGetter.Contract.IRoleGetterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRoleGetter *IRoleGetterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRoleGetter.Contract.IRoleGetterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRoleGetter *IRoleGetterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRoleGetter.Contract.IRoleGetterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRoleGetter *IRoleGetterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRoleGetter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRoleGetter *IRoleGetterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRoleGetter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRoleGetter *IRoleGetterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRoleGetter.Contract.contract.Transact(opts, method, params...)
}

// CheckIG is a free data retrieval call binding the contract method 0x7738515f.
//
// Solidity: function checkIG(uint64 _i, uint8 _rType) view returns(address, address, uint64, uint8)
func (_IRoleGetter *IRoleGetterCaller) CheckIG(opts *bind.CallOpts, _i uint64, _rType uint8) (common.Address, common.Address, uint64, uint8, error) {
	var out []interface{}
	err := _IRoleGetter.contract.Call(opts, &out, "checkIG", _i, _rType)

	if err != nil {
		return *new(common.Address), *new(common.Address), *new(uint64), *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	out2 := *abi.ConvertType(out[2], new(uint64)).(*uint64)
	out3 := *abi.ConvertType(out[3], new(uint8)).(*uint8)

	return out0, out1, out2, out3, err

}

// CheckIG is a free data retrieval call binding the contract method 0x7738515f.
//
// Solidity: function checkIG(uint64 _i, uint8 _rType) view returns(address, address, uint64, uint8)
func (_IRoleGetter *IRoleGetterSession) CheckIG(_i uint64, _rType uint8) (common.Address, common.Address, uint64, uint8, error) {
	return _IRoleGetter.Contract.CheckIG(&_IRoleGetter.CallOpts, _i, _rType)
}

// CheckIG is a free data retrieval call binding the contract method 0x7738515f.
//
// Solidity: function checkIG(uint64 _i, uint8 _rType) view returns(address, address, uint64, uint8)
func (_IRoleGetter *IRoleGetterCallerSession) CheckIG(_i uint64, _rType uint8) (common.Address, common.Address, uint64, uint8, error) {
	return _IRoleGetter.Contract.CheckIG(&_IRoleGetter.CallOpts, _i, _rType)
}

// GetACnt is a free data retrieval call binding the contract method 0x7264a551.
//
// Solidity: function getACnt() view returns(uint64)
func (_IRoleGetter *IRoleGetterCaller) GetACnt(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _IRoleGetter.contract.Call(opts, &out, "getACnt")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetACnt is a free data retrieval call binding the contract method 0x7264a551.
//
// Solidity: function getACnt() view returns(uint64)
func (_IRoleGetter *IRoleGetterSession) GetACnt() (uint64, error) {
	return _IRoleGetter.Contract.GetACnt(&_IRoleGetter.CallOpts)
}

// GetACnt is a free data retrieval call binding the contract method 0x7264a551.
//
// Solidity: function getACnt() view returns(uint64)
func (_IRoleGetter *IRoleGetterCallerSession) GetACnt() (uint64, error) {
	return _IRoleGetter.Contract.GetACnt(&_IRoleGetter.CallOpts)
}

// GetAcc is a free data retrieval call binding the contract method 0xcaca4a06.
//
// Solidity: function getAcc(address _a) view returns(uint64)
func (_IRoleGetter *IRoleGetterCaller) GetAcc(opts *bind.CallOpts, _a common.Address) (uint64, error) {
	var out []interface{}
	err := _IRoleGetter.contract.Call(opts, &out, "getAcc", _a)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetAcc is a free data retrieval call binding the contract method 0xcaca4a06.
//
// Solidity: function getAcc(address _a) view returns(uint64)
func (_IRoleGetter *IRoleGetterSession) GetAcc(_a common.Address) (uint64, error) {
	return _IRoleGetter.Contract.GetAcc(&_IRoleGetter.CallOpts, _a)
}

// GetAcc is a free data retrieval call binding the contract method 0xcaca4a06.
//
// Solidity: function getAcc(address _a) view returns(uint64)
func (_IRoleGetter *IRoleGetterCallerSession) GetAcc(_a common.Address) (uint64, error) {
	return _IRoleGetter.Contract.GetAcc(&_IRoleGetter.CallOpts, _a)
}

// GetAddr is a free data retrieval call binding the contract method 0x9332aa6e.
//
// Solidity: function getAddr(uint64 _i) view returns(address)
func (_IRoleGetter *IRoleGetterCaller) GetAddr(opts *bind.CallOpts, _i uint64) (common.Address, error) {
	var out []interface{}
	err := _IRoleGetter.contract.Call(opts, &out, "getAddr", _i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddr is a free data retrieval call binding the contract method 0x9332aa6e.
//
// Solidity: function getAddr(uint64 _i) view returns(address)
func (_IRoleGetter *IRoleGetterSession) GetAddr(_i uint64) (common.Address, error) {
	return _IRoleGetter.Contract.GetAddr(&_IRoleGetter.CallOpts, _i)
}

// GetAddr is a free data retrieval call binding the contract method 0x9332aa6e.
//
// Solidity: function getAddr(uint64 _i) view returns(address)
func (_IRoleGetter *IRoleGetterCallerSession) GetAddr(_i uint64) (common.Address, error) {
	return _IRoleGetter.Contract.GetAddr(&_IRoleGetter.CallOpts, _i)
}

// GetRInfo is a free data retrieval call binding the contract method 0x441abace.
//
// Solidity: function getRInfo(address _a) view returns((bool,bool,uint8,uint64,uint64,address,address,bytes))
func (_IRoleGetter *IRoleGetterCaller) GetRInfo(opts *bind.CallOpts, _a common.Address) (RoleOut, error) {
	var out []interface{}
	err := _IRoleGetter.contract.Call(opts, &out, "getRInfo", _a)

	if err != nil {
		return *new(RoleOut), err
	}

	out0 := *abi.ConvertType(out[0], new(RoleOut)).(*RoleOut)

	return out0, err

}

// GetRInfo is a free data retrieval call binding the contract method 0x441abace.
//
// Solidity: function getRInfo(address _a) view returns((bool,bool,uint8,uint64,uint64,address,address,bytes))
func (_IRoleGetter *IRoleGetterSession) GetRInfo(_a common.Address) (RoleOut, error) {
	return _IRoleGetter.Contract.GetRInfo(&_IRoleGetter.CallOpts, _a)
}

// GetRInfo is a free data retrieval call binding the contract method 0x441abace.
//
// Solidity: function getRInfo(address _a) view returns((bool,bool,uint8,uint64,uint64,address,address,bytes))
func (_IRoleGetter *IRoleGetterCallerSession) GetRInfo(_a common.Address) (RoleOut, error) {
	return _IRoleGetter.Contract.GetRInfo(&_IRoleGetter.CallOpts, _a)
}

// IRoleSetterABI is the input ABI used to generate the binding from.
const IRoleSetterABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"AlterPayee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"ConfirmPayee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"}],\"name\":\"ReAcc\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"rType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"}],\"name\":\"ReRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"gi\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"}],\"name\":\"SetG\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"}],\"name\":\"activate\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_index\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"_p\",\"type\":\"address\"}],\"name\":\"alterPayee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"_ban\",\"type\":\"bool\"}],\"name\":\"ban\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_index\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"_p\",\"type\":\"address\"}],\"name\":\"confirmPayee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"}],\"name\":\"reAcc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_rtype\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"extra\",\"type\":\"bytes\"}],\"name\":\"reRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"}],\"name\":\"setG\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IRoleSetterFuncSigs maps the 4-byte function signature to its string representation.
var IRoleSetterFuncSigs = map[string]string{
	"fc7e4746": "activate(uint64)",
	"2915824a": "alterPayee(uint64,address)",
	"0f78c61a": "ban(uint64,bool)",
	"5bd04c3f": "confirmPayee(uint64,address)",
	"effcafce": "reAcc(address)",
	"b9f6a8ca": "reRole(uint64,uint8,bytes)",
	"616bfd1f": "setG(uint64,uint64)",
}

// IRoleSetter is an auto generated Go binding around an Ethereum contract.
type IRoleSetter struct {
	IRoleSetterCaller     // Read-only binding to the contract
	IRoleSetterTransactor // Write-only binding to the contract
	IRoleSetterFilterer   // Log filterer for contract events
}

// IRoleSetterCaller is an auto generated read-only Go binding around an Ethereum contract.
type IRoleSetterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRoleSetterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IRoleSetterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRoleSetterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IRoleSetterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRoleSetterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IRoleSetterSession struct {
	Contract     *IRoleSetter      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IRoleSetterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IRoleSetterCallerSession struct {
	Contract *IRoleSetterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// IRoleSetterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IRoleSetterTransactorSession struct {
	Contract     *IRoleSetterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IRoleSetterRaw is an auto generated low-level Go binding around an Ethereum contract.
type IRoleSetterRaw struct {
	Contract *IRoleSetter // Generic contract binding to access the raw methods on
}

// IRoleSetterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IRoleSetterCallerRaw struct {
	Contract *IRoleSetterCaller // Generic read-only contract binding to access the raw methods on
}

// IRoleSetterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IRoleSetterTransactorRaw struct {
	Contract *IRoleSetterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIRoleSetter creates a new instance of IRoleSetter, bound to a specific deployed contract.
func NewIRoleSetter(address common.Address, backend bind.ContractBackend) (*IRoleSetter, error) {
	contract, err := bindIRoleSetter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IRoleSetter{IRoleSetterCaller: IRoleSetterCaller{contract: contract}, IRoleSetterTransactor: IRoleSetterTransactor{contract: contract}, IRoleSetterFilterer: IRoleSetterFilterer{contract: contract}}, nil
}

// NewIRoleSetterCaller creates a new read-only instance of IRoleSetter, bound to a specific deployed contract.
func NewIRoleSetterCaller(address common.Address, caller bind.ContractCaller) (*IRoleSetterCaller, error) {
	contract, err := bindIRoleSetter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IRoleSetterCaller{contract: contract}, nil
}

// NewIRoleSetterTransactor creates a new write-only instance of IRoleSetter, bound to a specific deployed contract.
func NewIRoleSetterTransactor(address common.Address, transactor bind.ContractTransactor) (*IRoleSetterTransactor, error) {
	contract, err := bindIRoleSetter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IRoleSetterTransactor{contract: contract}, nil
}

// NewIRoleSetterFilterer creates a new log filterer instance of IRoleSetter, bound to a specific deployed contract.
func NewIRoleSetterFilterer(address common.Address, filterer bind.ContractFilterer) (*IRoleSetterFilterer, error) {
	contract, err := bindIRoleSetter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IRoleSetterFilterer{contract: contract}, nil
}

// bindIRoleSetter binds a generic wrapper to an already deployed contract.
func bindIRoleSetter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IRoleSetterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRoleSetter *IRoleSetterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRoleSetter.Contract.IRoleSetterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRoleSetter *IRoleSetterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRoleSetter.Contract.IRoleSetterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRoleSetter *IRoleSetterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRoleSetter.Contract.IRoleSetterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRoleSetter *IRoleSetterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRoleSetter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRoleSetter *IRoleSetterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRoleSetter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRoleSetter *IRoleSetterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRoleSetter.Contract.contract.Transact(opts, method, params...)
}

// Activate is a paid mutator transaction binding the contract method 0xfc7e4746.
//
// Solidity: function activate(uint64 _i) returns(uint64)
func (_IRoleSetter *IRoleSetterTransactor) Activate(opts *bind.TransactOpts, _i uint64) (*types.Transaction, error) {
	return _IRoleSetter.contract.Transact(opts, "activate", _i)
}

// Activate is a paid mutator transaction binding the contract method 0xfc7e4746.
//
// Solidity: function activate(uint64 _i) returns(uint64)
func (_IRoleSetter *IRoleSetterSession) Activate(_i uint64) (*types.Transaction, error) {
	return _IRoleSetter.Contract.Activate(&_IRoleSetter.TransactOpts, _i)
}

// Activate is a paid mutator transaction binding the contract method 0xfc7e4746.
//
// Solidity: function activate(uint64 _i) returns(uint64)
func (_IRoleSetter *IRoleSetterTransactorSession) Activate(_i uint64) (*types.Transaction, error) {
	return _IRoleSetter.Contract.Activate(&_IRoleSetter.TransactOpts, _i)
}

// AlterPayee is a paid mutator transaction binding the contract method 0x2915824a.
//
// Solidity: function alterPayee(uint64 _index, address _p) returns()
func (_IRoleSetter *IRoleSetterTransactor) AlterPayee(opts *bind.TransactOpts, _index uint64, _p common.Address) (*types.Transaction, error) {
	return _IRoleSetter.contract.Transact(opts, "alterPayee", _index, _p)
}

// AlterPayee is a paid mutator transaction binding the contract method 0x2915824a.
//
// Solidity: function alterPayee(uint64 _index, address _p) returns()
func (_IRoleSetter *IRoleSetterSession) AlterPayee(_index uint64, _p common.Address) (*types.Transaction, error) {
	return _IRoleSetter.Contract.AlterPayee(&_IRoleSetter.TransactOpts, _index, _p)
}

// AlterPayee is a paid mutator transaction binding the contract method 0x2915824a.
//
// Solidity: function alterPayee(uint64 _index, address _p) returns()
func (_IRoleSetter *IRoleSetterTransactorSession) AlterPayee(_index uint64, _p common.Address) (*types.Transaction, error) {
	return _IRoleSetter.Contract.AlterPayee(&_IRoleSetter.TransactOpts, _index, _p)
}

// Ban is a paid mutator transaction binding the contract method 0x0f78c61a.
//
// Solidity: function ban(uint64 _i, bool _ban) returns()
func (_IRoleSetter *IRoleSetterTransactor) Ban(opts *bind.TransactOpts, _i uint64, _ban bool) (*types.Transaction, error) {
	return _IRoleSetter.contract.Transact(opts, "ban", _i, _ban)
}

// Ban is a paid mutator transaction binding the contract method 0x0f78c61a.
//
// Solidity: function ban(uint64 _i, bool _ban) returns()
func (_IRoleSetter *IRoleSetterSession) Ban(_i uint64, _ban bool) (*types.Transaction, error) {
	return _IRoleSetter.Contract.Ban(&_IRoleSetter.TransactOpts, _i, _ban)
}

// Ban is a paid mutator transaction binding the contract method 0x0f78c61a.
//
// Solidity: function ban(uint64 _i, bool _ban) returns()
func (_IRoleSetter *IRoleSetterTransactorSession) Ban(_i uint64, _ban bool) (*types.Transaction, error) {
	return _IRoleSetter.Contract.Ban(&_IRoleSetter.TransactOpts, _i, _ban)
}

// ConfirmPayee is a paid mutator transaction binding the contract method 0x5bd04c3f.
//
// Solidity: function confirmPayee(uint64 _index, address _p) returns()
func (_IRoleSetter *IRoleSetterTransactor) ConfirmPayee(opts *bind.TransactOpts, _index uint64, _p common.Address) (*types.Transaction, error) {
	return _IRoleSetter.contract.Transact(opts, "confirmPayee", _index, _p)
}

// ConfirmPayee is a paid mutator transaction binding the contract method 0x5bd04c3f.
//
// Solidity: function confirmPayee(uint64 _index, address _p) returns()
func (_IRoleSetter *IRoleSetterSession) ConfirmPayee(_index uint64, _p common.Address) (*types.Transaction, error) {
	return _IRoleSetter.Contract.ConfirmPayee(&_IRoleSetter.TransactOpts, _index, _p)
}

// ConfirmPayee is a paid mutator transaction binding the contract method 0x5bd04c3f.
//
// Solidity: function confirmPayee(uint64 _index, address _p) returns()
func (_IRoleSetter *IRoleSetterTransactorSession) ConfirmPayee(_index uint64, _p common.Address) (*types.Transaction, error) {
	return _IRoleSetter.Contract.ConfirmPayee(&_IRoleSetter.TransactOpts, _index, _p)
}

// ReAcc is a paid mutator transaction binding the contract method 0xeffcafce.
//
// Solidity: function reAcc(address _a) returns()
func (_IRoleSetter *IRoleSetterTransactor) ReAcc(opts *bind.TransactOpts, _a common.Address) (*types.Transaction, error) {
	return _IRoleSetter.contract.Transact(opts, "reAcc", _a)
}

// ReAcc is a paid mutator transaction binding the contract method 0xeffcafce.
//
// Solidity: function reAcc(address _a) returns()
func (_IRoleSetter *IRoleSetterSession) ReAcc(_a common.Address) (*types.Transaction, error) {
	return _IRoleSetter.Contract.ReAcc(&_IRoleSetter.TransactOpts, _a)
}

// ReAcc is a paid mutator transaction binding the contract method 0xeffcafce.
//
// Solidity: function reAcc(address _a) returns()
func (_IRoleSetter *IRoleSetterTransactorSession) ReAcc(_a common.Address) (*types.Transaction, error) {
	return _IRoleSetter.Contract.ReAcc(&_IRoleSetter.TransactOpts, _a)
}

// ReRole is a paid mutator transaction binding the contract method 0xb9f6a8ca.
//
// Solidity: function reRole(uint64 _i, uint8 _rtype, bytes extra) returns()
func (_IRoleSetter *IRoleSetterTransactor) ReRole(opts *bind.TransactOpts, _i uint64, _rtype uint8, extra []byte) (*types.Transaction, error) {
	return _IRoleSetter.contract.Transact(opts, "reRole", _i, _rtype, extra)
}

// ReRole is a paid mutator transaction binding the contract method 0xb9f6a8ca.
//
// Solidity: function reRole(uint64 _i, uint8 _rtype, bytes extra) returns()
func (_IRoleSetter *IRoleSetterSession) ReRole(_i uint64, _rtype uint8, extra []byte) (*types.Transaction, error) {
	return _IRoleSetter.Contract.ReRole(&_IRoleSetter.TransactOpts, _i, _rtype, extra)
}

// ReRole is a paid mutator transaction binding the contract method 0xb9f6a8ca.
//
// Solidity: function reRole(uint64 _i, uint8 _rtype, bytes extra) returns()
func (_IRoleSetter *IRoleSetterTransactorSession) ReRole(_i uint64, _rtype uint8, extra []byte) (*types.Transaction, error) {
	return _IRoleSetter.Contract.ReRole(&_IRoleSetter.TransactOpts, _i, _rtype, extra)
}

// SetG is a paid mutator transaction binding the contract method 0x616bfd1f.
//
// Solidity: function setG(uint64 _i, uint64 _gi) returns()
func (_IRoleSetter *IRoleSetterTransactor) SetG(opts *bind.TransactOpts, _i uint64, _gi uint64) (*types.Transaction, error) {
	return _IRoleSetter.contract.Transact(opts, "setG", _i, _gi)
}

// SetG is a paid mutator transaction binding the contract method 0x616bfd1f.
//
// Solidity: function setG(uint64 _i, uint64 _gi) returns()
func (_IRoleSetter *IRoleSetterSession) SetG(_i uint64, _gi uint64) (*types.Transaction, error) {
	return _IRoleSetter.Contract.SetG(&_IRoleSetter.TransactOpts, _i, _gi)
}

// SetG is a paid mutator transaction binding the contract method 0x616bfd1f.
//
// Solidity: function setG(uint64 _i, uint64 _gi) returns()
func (_IRoleSetter *IRoleSetterTransactorSession) SetG(_i uint64, _gi uint64) (*types.Transaction, error) {
	return _IRoleSetter.Contract.SetG(&_IRoleSetter.TransactOpts, _i, _gi)
}

// IRoleSetterAlterPayeeIterator is returned from FilterAlterPayee and is used to iterate over the raw logs and unpacked data for AlterPayee events raised by the IRoleSetter contract.
type IRoleSetterAlterPayeeIterator struct {
	Event *IRoleSetterAlterPayee // Event containing the contract specifics and raw log

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
func (it *IRoleSetterAlterPayeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IRoleSetterAlterPayee)
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
		it.Event = new(IRoleSetterAlterPayee)
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
func (it *IRoleSetterAlterPayeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IRoleSetterAlterPayeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IRoleSetterAlterPayee represents a AlterPayee event raised by the IRoleSetter contract.
type IRoleSetterAlterPayee struct {
	Payee common.Address
	Addr  common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAlterPayee is a free log retrieval operation binding the contract event 0x3b4058e93b2a019add56609a3c2463a7b3d8476abf30c934cfcdb9f2158b5e68.
//
// Solidity: event AlterPayee(address indexed payee, address addr)
func (_IRoleSetter *IRoleSetterFilterer) FilterAlterPayee(opts *bind.FilterOpts, payee []common.Address) (*IRoleSetterAlterPayeeIterator, error) {

	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}

	logs, sub, err := _IRoleSetter.contract.FilterLogs(opts, "AlterPayee", payeeRule)
	if err != nil {
		return nil, err
	}
	return &IRoleSetterAlterPayeeIterator{contract: _IRoleSetter.contract, event: "AlterPayee", logs: logs, sub: sub}, nil
}

// WatchAlterPayee is a free log subscription operation binding the contract event 0x3b4058e93b2a019add56609a3c2463a7b3d8476abf30c934cfcdb9f2158b5e68.
//
// Solidity: event AlterPayee(address indexed payee, address addr)
func (_IRoleSetter *IRoleSetterFilterer) WatchAlterPayee(opts *bind.WatchOpts, sink chan<- *IRoleSetterAlterPayee, payee []common.Address) (event.Subscription, error) {

	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}

	logs, sub, err := _IRoleSetter.contract.WatchLogs(opts, "AlterPayee", payeeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IRoleSetterAlterPayee)
				if err := _IRoleSetter.contract.UnpackLog(event, "AlterPayee", log); err != nil {
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

// ParseAlterPayee is a log parse operation binding the contract event 0x3b4058e93b2a019add56609a3c2463a7b3d8476abf30c934cfcdb9f2158b5e68.
//
// Solidity: event AlterPayee(address indexed payee, address addr)
func (_IRoleSetter *IRoleSetterFilterer) ParseAlterPayee(log types.Log) (*IRoleSetterAlterPayee, error) {
	event := new(IRoleSetterAlterPayee)
	if err := _IRoleSetter.contract.UnpackLog(event, "AlterPayee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IRoleSetterConfirmPayeeIterator is returned from FilterConfirmPayee and is used to iterate over the raw logs and unpacked data for ConfirmPayee events raised by the IRoleSetter contract.
type IRoleSetterConfirmPayeeIterator struct {
	Event *IRoleSetterConfirmPayee // Event containing the contract specifics and raw log

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
func (it *IRoleSetterConfirmPayeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IRoleSetterConfirmPayee)
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
		it.Event = new(IRoleSetterConfirmPayee)
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
func (it *IRoleSetterConfirmPayeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IRoleSetterConfirmPayeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IRoleSetterConfirmPayee represents a ConfirmPayee event raised by the IRoleSetter contract.
type IRoleSetterConfirmPayee struct {
	Payee common.Address
	Addr  common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterConfirmPayee is a free log retrieval operation binding the contract event 0x4a21d873c33448e8dfb90cdbfdb12f4aed5f5e2abde5525215b1655d519784dc.
//
// Solidity: event ConfirmPayee(address indexed payee, address addr)
func (_IRoleSetter *IRoleSetterFilterer) FilterConfirmPayee(opts *bind.FilterOpts, payee []common.Address) (*IRoleSetterConfirmPayeeIterator, error) {

	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}

	logs, sub, err := _IRoleSetter.contract.FilterLogs(opts, "ConfirmPayee", payeeRule)
	if err != nil {
		return nil, err
	}
	return &IRoleSetterConfirmPayeeIterator{contract: _IRoleSetter.contract, event: "ConfirmPayee", logs: logs, sub: sub}, nil
}

// WatchConfirmPayee is a free log subscription operation binding the contract event 0x4a21d873c33448e8dfb90cdbfdb12f4aed5f5e2abde5525215b1655d519784dc.
//
// Solidity: event ConfirmPayee(address indexed payee, address addr)
func (_IRoleSetter *IRoleSetterFilterer) WatchConfirmPayee(opts *bind.WatchOpts, sink chan<- *IRoleSetterConfirmPayee, payee []common.Address) (event.Subscription, error) {

	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}

	logs, sub, err := _IRoleSetter.contract.WatchLogs(opts, "ConfirmPayee", payeeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IRoleSetterConfirmPayee)
				if err := _IRoleSetter.contract.UnpackLog(event, "ConfirmPayee", log); err != nil {
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

// ParseConfirmPayee is a log parse operation binding the contract event 0x4a21d873c33448e8dfb90cdbfdb12f4aed5f5e2abde5525215b1655d519784dc.
//
// Solidity: event ConfirmPayee(address indexed payee, address addr)
func (_IRoleSetter *IRoleSetterFilterer) ParseConfirmPayee(log types.Log) (*IRoleSetterConfirmPayee, error) {
	event := new(IRoleSetterConfirmPayee)
	if err := _IRoleSetter.contract.UnpackLog(event, "ConfirmPayee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IRoleSetterReAccIterator is returned from FilterReAcc and is used to iterate over the raw logs and unpacked data for ReAcc events raised by the IRoleSetter contract.
type IRoleSetterReAccIterator struct {
	Event *IRoleSetterReAcc // Event containing the contract specifics and raw log

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
func (it *IRoleSetterReAccIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IRoleSetterReAcc)
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
		it.Event = new(IRoleSetterReAcc)
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
func (it *IRoleSetterReAccIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IRoleSetterReAccIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IRoleSetterReAcc represents a ReAcc event raised by the IRoleSetter contract.
type IRoleSetterReAcc struct {
	Addr  common.Address
	Index uint64
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterReAcc is a free log retrieval operation binding the contract event 0xe22ccfc11bbbb1d3350012248b26ad616e32a8a13eece408dd14f1953fe24752.
//
// Solidity: event ReAcc(address addr, uint64 index)
func (_IRoleSetter *IRoleSetterFilterer) FilterReAcc(opts *bind.FilterOpts) (*IRoleSetterReAccIterator, error) {

	logs, sub, err := _IRoleSetter.contract.FilterLogs(opts, "ReAcc")
	if err != nil {
		return nil, err
	}
	return &IRoleSetterReAccIterator{contract: _IRoleSetter.contract, event: "ReAcc", logs: logs, sub: sub}, nil
}

// WatchReAcc is a free log subscription operation binding the contract event 0xe22ccfc11bbbb1d3350012248b26ad616e32a8a13eece408dd14f1953fe24752.
//
// Solidity: event ReAcc(address addr, uint64 index)
func (_IRoleSetter *IRoleSetterFilterer) WatchReAcc(opts *bind.WatchOpts, sink chan<- *IRoleSetterReAcc) (event.Subscription, error) {

	logs, sub, err := _IRoleSetter.contract.WatchLogs(opts, "ReAcc")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IRoleSetterReAcc)
				if err := _IRoleSetter.contract.UnpackLog(event, "ReAcc", log); err != nil {
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

// ParseReAcc is a log parse operation binding the contract event 0xe22ccfc11bbbb1d3350012248b26ad616e32a8a13eece408dd14f1953fe24752.
//
// Solidity: event ReAcc(address addr, uint64 index)
func (_IRoleSetter *IRoleSetterFilterer) ParseReAcc(log types.Log) (*IRoleSetterReAcc, error) {
	event := new(IRoleSetterReAcc)
	if err := _IRoleSetter.contract.UnpackLog(event, "ReAcc", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IRoleSetterReRoleIterator is returned from FilterReRole and is used to iterate over the raw logs and unpacked data for ReRole events raised by the IRoleSetter contract.
type IRoleSetterReRoleIterator struct {
	Event *IRoleSetterReRole // Event containing the contract specifics and raw log

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
func (it *IRoleSetterReRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IRoleSetterReRole)
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
		it.Event = new(IRoleSetterReRole)
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
func (it *IRoleSetterReRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IRoleSetterReRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IRoleSetterReRole represents a ReRole event raised by the IRoleSetter contract.
type IRoleSetterReRole struct {
	RType uint8
	Index uint64
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterReRole is a free log retrieval operation binding the contract event 0x5550fbab402c3b014955818e5b6e3cefef20641aa6bed43ae83cd5d3258b8922.
//
// Solidity: event ReRole(uint8 indexed rType, uint64 index)
func (_IRoleSetter *IRoleSetterFilterer) FilterReRole(opts *bind.FilterOpts, rType []uint8) (*IRoleSetterReRoleIterator, error) {

	var rTypeRule []interface{}
	for _, rTypeItem := range rType {
		rTypeRule = append(rTypeRule, rTypeItem)
	}

	logs, sub, err := _IRoleSetter.contract.FilterLogs(opts, "ReRole", rTypeRule)
	if err != nil {
		return nil, err
	}
	return &IRoleSetterReRoleIterator{contract: _IRoleSetter.contract, event: "ReRole", logs: logs, sub: sub}, nil
}

// WatchReRole is a free log subscription operation binding the contract event 0x5550fbab402c3b014955818e5b6e3cefef20641aa6bed43ae83cd5d3258b8922.
//
// Solidity: event ReRole(uint8 indexed rType, uint64 index)
func (_IRoleSetter *IRoleSetterFilterer) WatchReRole(opts *bind.WatchOpts, sink chan<- *IRoleSetterReRole, rType []uint8) (event.Subscription, error) {

	var rTypeRule []interface{}
	for _, rTypeItem := range rType {
		rTypeRule = append(rTypeRule, rTypeItem)
	}

	logs, sub, err := _IRoleSetter.contract.WatchLogs(opts, "ReRole", rTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IRoleSetterReRole)
				if err := _IRoleSetter.contract.UnpackLog(event, "ReRole", log); err != nil {
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

// ParseReRole is a log parse operation binding the contract event 0x5550fbab402c3b014955818e5b6e3cefef20641aa6bed43ae83cd5d3258b8922.
//
// Solidity: event ReRole(uint8 indexed rType, uint64 index)
func (_IRoleSetter *IRoleSetterFilterer) ParseReRole(log types.Log) (*IRoleSetterReRole, error) {
	event := new(IRoleSetterReRole)
	if err := _IRoleSetter.contract.UnpackLog(event, "ReRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IRoleSetterSetGIterator is returned from FilterSetG and is used to iterate over the raw logs and unpacked data for SetG events raised by the IRoleSetter contract.
type IRoleSetterSetGIterator struct {
	Event *IRoleSetterSetG // Event containing the contract specifics and raw log

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
func (it *IRoleSetterSetGIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IRoleSetterSetG)
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
		it.Event = new(IRoleSetterSetG)
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
func (it *IRoleSetterSetGIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IRoleSetterSetGIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IRoleSetterSetG represents a SetG event raised by the IRoleSetter contract.
type IRoleSetterSetG struct {
	Gi    uint64
	Index uint64
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSetG is a free log retrieval operation binding the contract event 0xe7329265e8e4e2fc8e892a75f1809767d09c472bb2eef16f889eceb5583291a9.
//
// Solidity: event SetG(uint64 indexed gi, uint64 index)
func (_IRoleSetter *IRoleSetterFilterer) FilterSetG(opts *bind.FilterOpts, gi []uint64) (*IRoleSetterSetGIterator, error) {

	var giRule []interface{}
	for _, giItem := range gi {
		giRule = append(giRule, giItem)
	}

	logs, sub, err := _IRoleSetter.contract.FilterLogs(opts, "SetG", giRule)
	if err != nil {
		return nil, err
	}
	return &IRoleSetterSetGIterator{contract: _IRoleSetter.contract, event: "SetG", logs: logs, sub: sub}, nil
}

// WatchSetG is a free log subscription operation binding the contract event 0xe7329265e8e4e2fc8e892a75f1809767d09c472bb2eef16f889eceb5583291a9.
//
// Solidity: event SetG(uint64 indexed gi, uint64 index)
func (_IRoleSetter *IRoleSetterFilterer) WatchSetG(opts *bind.WatchOpts, sink chan<- *IRoleSetterSetG, gi []uint64) (event.Subscription, error) {

	var giRule []interface{}
	for _, giItem := range gi {
		giRule = append(giRule, giItem)
	}

	logs, sub, err := _IRoleSetter.contract.WatchLogs(opts, "SetG", giRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IRoleSetterSetG)
				if err := _IRoleSetter.contract.UnpackLog(event, "SetG", log); err != nil {
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

// ParseSetG is a log parse operation binding the contract event 0xe7329265e8e4e2fc8e892a75f1809767d09c472bb2eef16f889eceb5583291a9.
//
// Solidity: event SetG(uint64 indexed gi, uint64 index)
func (_IRoleSetter *IRoleSetterFilterer) ParseSetG(log types.Log) (*IRoleSetterSetG, error) {
	event := new(IRoleSetterSetG)
	if err := _IRoleSetter.contract.UnpackLog(event, "SetG", log); err != nil {
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

// RoleABI is the input ABI used to generate the binding from.
const RoleABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"f\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isSet\",\"type\":\"bool\"}],\"name\":\"AddOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"AlterPayee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"ConfirmPayee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"}],\"name\":\"ReAcc\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"rType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"}],\"name\":\"ReRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"gi\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"}],\"name\":\"SetG\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_index\",\"type\":\"uint64\"}],\"name\":\"activate\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isSet\",\"type\":\"bool\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_index\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"_p\",\"type\":\"address\"}],\"name\":\"alterPayee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"auth\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_index\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"_banned\",\"type\":\"bool\"}],\"name\":\"ban\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_index\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_rType\",\"type\":\"uint8\"}],\"name\":\"checkIG\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_index\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"_p\",\"type\":\"address\"}],\"name\":\"confirmPayee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getACnt\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"}],\"name\":\"getAcc\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"}],\"name\":\"getAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"acc\",\"type\":\"address\"}],\"name\":\"getRInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isBan\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"rType\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gIndex\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"next\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extra\",\"type\":\"bytes\"}],\"internalType\":\"structRoleOut\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"owners\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"reAcc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_index\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_rType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_extra\",\"type\":\"bytes\"}],\"name\":\"reRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_index\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"}],\"name\":\"setG\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// RoleFuncSigs maps the 4-byte function signature to its string representation.
var RoleFuncSigs = map[string]string{
	"fc7e4746": "activate(uint64)",
	"4bf1b134": "add(address,bool,bytes[5])",
	"2915824a": "alterPayee(uint64,address)",
	"de9375f2": "auth()",
	"0f78c61a": "ban(uint64,bool)",
	"7738515f": "checkIG(uint64,uint8)",
	"5bd04c3f": "confirmPayee(uint64,address)",
	"7264a551": "getACnt()",
	"caca4a06": "getAcc(address)",
	"9332aa6e": "getAddr(uint64)",
	"441abace": "getRInfo(address)",
	"022914a7": "owners(address)",
	"effcafce": "reAcc(address)",
	"b9f6a8ca": "reRole(uint64,uint8,bytes)",
	"616bfd1f": "setG(uint64,uint64)",
	"54fd4d50": "version()",
}

// RoleBin is the compiled bytecode used for deploying new contracts.
var RoleBin = "0x60806040526001805461ffff60a01b1916600160a11b17905534801561002457600080fd5b506040516200175b3803806200175b83398101604081905261004591610113565b600180546001600160a01b038481166001600160a01b0319928316178355600280548085019091557f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace018054918516918316821790556000818152600360209081526040808320909501805490941683179093558351918252918101919091527fe22ccfc11bbbb1d3350012248b26ad616e32a8a13eece408dd14f1953fe24752910160405180910390a15050610146565b80516001600160a01b038116811461010e57600080fd5b919050565b6000806040838503121561012657600080fd5b61012f836100f7565b915061013d602084016100f7565b90509250929050565b61160580620001566000396000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c80637264a55111610097578063caca4a0611610066578063caca4a061461028f578063de9375f2146102c8578063effcafce146102db578063fc7e4746146102ee57600080fd5b80637264a551146101e65780637738515f146102025780639332aa6e14610251578063b9f6a8ca1461027c57600080fd5b80634bf1b134116100d35780634bf1b1341461018557806354fd4d50146101985780635bd04c3f146101c0578063616bfd1f146101d357600080fd5b8063022914a7146101055780630f78c61a1461013d5780632915824a14610152578063441abace14610165575b600080fd5b610128610113366004611109565b60006020819052908152604090205460ff1681565b60405190151581526020015b60405180910390f35b61015061014b366004611152565b610301565b005b610150610160366004611185565b61038d565b610178610173366004611109565b610452565b60405161013491906111fc565b61015061019336600461136c565b610591565b6001546101ad90600160a01b900461ffff1681565b60405161ffff9091168152602001610134565b6101506101ce366004611185565b61070a565b6101506101e1366004611428565b610848565b6002545b6040516001600160401b039091168152602001610134565b610215610210366004611463565b610955565b604080516001600160a01b0395861681529490931660208501526001600160401b039091169183019190915260ff166060820152608001610134565b61026461025f36600461148d565b610acd565b6040516001600160a01b039091168152602001610134565b61015061028a3660046114a8565b610b06565b6101ea61029d366004611109565b6001600160a01b0316600090815260036020526040902054630100000090046001600160401b031690565b600154610264906001600160a01b031681565b6101506102e9366004611109565b610d07565b6101ea6102fc36600461148d565b610e99565b3360009081526020819052604090205460ff166103395760405162461bcd60e51b815260040161033090611505565b60405180910390fd5b60006002836001600160401b03168154811061035757610357611528565b6000918252602080832091909101546001600160a01b031682526003905260409020805460ff1916921515929092179091555050565b3360009081526020819052604090205460ff166103bc5760405162461bcd60e51b815260040161033090611505565b60006002836001600160401b0316815481106103da576103da611528565b6000918252602080832091909101546001600160a01b039081168084526003835260409384902060020180546001600160a01b03191692871692831790559251838152929350917f3b4058e93b2a019add56609a3c2463a7b3d8476abf30c934cfcdb9f2158b5e6891015b60405180910390a2505050565b61045a611013565b610462611013565b6040805161010080820183526001600160a01b03808716600081815260036020818152878320805460ff80821615158a5297810488161515838a0152620100008104909716988801989098526001600160401b036301000000870481166060890152600160581b90960490951660808701526001870154841660a0870152600287015490931660c0860152529081905291909101805460e0830191906105079061153e565b80601f01602080910402602001604051908101604052809291908181526020018280546105339061153e565b80156105805780601f1061055557610100808354040283529160200191610580565b820191906000526020600020905b81548152906001019060200180831161056357829003601f168201915b505050919092525090949350505050565b823b60008190036105d95760405162461bcd60e51b81526020600482015260126024820152713732b2b21031b7b73a3930b1ba1030b2323960711b6044820152606401610330565b6040516bffffffffffffffffffffffff1930606090811b821660208401526218591960ea1b603484015286901b16603782015283151560f81b604b820152600090604c0160408051601f1981840301815290829052805160209091012060015463a96bba9d60e01b83529092506001600160a01b03169063a96bba9d906106669084908790600401611578565b600060405180830381600087803b15801561068057600080fd5b505af1158015610694573d6000803e3d6000fd5b5050604080516001600160a01b038916815287151560208201527f938b2a24c98e4e542127ffa74a91e48942c2bddccae3b6d75f82bfda7bbc0807935001905060405180910390a15050506001600160a01b03919091166000908152602081905260409020805460ff1916911515919091179055565b3360009081526020819052604090205460ff166107395760405162461bcd60e51b815260040161033090611505565b60006002836001600160401b03168154811061075757610757611528565b60009182526020808320909101546001600160a01b0316808352600390915260409091205490915060ff161580156107ae57506001600160a01b038181166000908152600360205260409020600201548116908316145b6107e75760405162461bcd60e51b815260206004820152600a60248201526931b0b63632b91032b93960b11b6044820152606401610330565b6001600160a01b0381811660008181526003602090815260409182902060010180546001600160a01b031916948716948517905590519182527f4a21d873c33448e8dfb90cdbfdb12f4aed5f5e2abde5525215b1655d519784dc9101610445565b3360009081526020819052604090205460ff166108775760405162461bcd60e51b815260040161033090611505565b60006002836001600160401b03168154811061089557610895611528565b6000918252602080832091909101546001600160a01b03168083526003918290526040909220805467ffffffffffffffff60581b1916600160581b6001600160401b0387160217908190559192506201000090910460ff1614610918576001600160a01b0381166000908152600360205260409020805461ff0019166101001790555b6040516001600160401b0384811682528316907fe7329265e8e4e2fc8e892a75f1809767d09c472bb2eef16f889eceb5583291a990602001610445565b60008060008060006002876001600160401b03168154811061097957610979611528565b60009182526020808320909101546001600160a01b0316808352600390915260409091205490915060ff16156109de5760405162461bcd60e51b815260206004820152600a6024820152691858d8c818985b9b995960b21b6044820152606401610330565b60ff861615610a83576001600160a01b038116600090815260036020526040902054610100900460ff168015610a3757506001600160a01b03811660009081526003602052604090205460ff8781166201000090920416145b610a835760405162461bcd60e51b815260206004820152601a60248201527f616363207479706520657272206f72206e6f74206163746976650000000000006044820152606401610330565b6001600160a01b03818116600090815260036020526040902060018101549054929991169750600160581b82046001600160401b031696506201000090910460ff16945092505050565b60006002826001600160401b031681548110610aeb57610aeb611528565b6000918252602090912001546001600160a01b031692915050565b3360009081526020819052604090205460ff16610b355760405162461bcd60e51b815260040161033090611505565b60006002846001600160401b031681548110610b5357610b53611528565b6000918252602090912001546001600160a01b0316905060ff8316610ba75760405162461bcd60e51b815260206004820152600a602482015269686173207265526f6c6560b01b6044820152606401610330565b6001600160a01b03811660009081526003602052604090205462010000900460ff16158015610bfc57506001600160a01b038116600090815260036020526040902054600160581b90046001600160401b0316155b8015610c2157506001600160a01b03811660009081526003602052604090205460ff16155b8015610c4b57506001600160a01b038116600090815260036020526040902054610100900460ff16155b610c7c5760405162461bcd60e51b8152602060048201526002602482015261414560f01b6044820152606401610330565b6001600160a01b0381166000908152600360208181526040909220805462ff000019166201000060ff8816021781558451610cbf93919092019190850190611059565b506040516001600160401b038516815260ff8416907f5550fbab402c3b014955818e5b6e3cefef20641aa6bed43ae83cd5d3258b89229060200160405180910390a250505050565b3360009081526020819052604090205460ff16610d365760405162461bcd60e51b815260040161033090611505565b6001600160a01b038116600090815260036020526040902054630100000090046001600160401b0316158015610d9257506001600160a01b038116600090815260036020526040902054600160581b90046001600160401b0316155b610dca5760405162461bcd60e51b815260206004820152600960248201526868617320726541636360b81b6044820152606401610330565b600280546001600160a01b038316600081815260036020908152604080832080546affffffffffffffff000000191663010000006001600160401b038816908102919091178255600191820180546001600160a01b0319908116881790915588549283018955979094527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace0180549096168417909555845192835282015290917fe22ccfc11bbbb1d3350012248b26ad616e32a8a13eece408dd14f1953fe24752910160405180910390a15050565b3360009081526020819052604081205460ff16610ec85760405162461bcd60e51b815260040161033090611505565b60006002836001600160401b031681548110610ee657610ee6611528565b60009182526020808320909101546001600160a01b0316808352600390915260409091205490915060ff61010090910416158015610f3d57506001600160a01b03811660009081526003602052604090205460ff16155b8015610f6b57506001600160a01b03811660009081526003602081905260409091205462010000900460ff16145b8015610f9e57506001600160a01b038116600090815260036020526040902054600160581b90046001600160401b031615155b610fd45760405162461bcd60e51b815260206004820152600760248201526630b1b19032b93960c91b6044820152606401610330565b6001600160a01b03166000908152600360205260409020805461ff00191661010017908190556001600160401b03600160581b9091041690505b919050565b604080516101008101825260008082526020820181905291810182905260608082018390526080820183905260a0820183905260c082019290925260e081019190915290565b8280546110659061153e565b90600052602060002090601f01602090048101928261108757600085556110cd565b82601f106110a057805160ff19168380011785556110cd565b828001600101855582156110cd579182015b828111156110cd5782518255916020019190600101906110b2565b506110d99291506110dd565b5090565b5b808211156110d957600081556001016110de565b80356001600160a01b038116811461100e57600080fd5b60006020828403121561111b57600080fd5b611124826110f2565b9392505050565b80356001600160401b038116811461100e57600080fd5b8035801515811461100e57600080fd5b6000806040838503121561116557600080fd5b61116e8361112b565b915061117c60208401611142565b90509250929050565b6000806040838503121561119857600080fd5b6111a18361112b565b915061117c602084016110f2565b6000815180845260005b818110156111d5576020818501810151868301820152016111b9565b818111156111e7576000602083870101525b50601f01601f19169290920160200192915050565b602081528151151560208201526020820151151560408201526000604083015161122b606084018260ff169052565b5060608301516001600160401b03811660808401525060808301516001600160401b03811660a08401525060a08301516001600160a01b03811660c08401525060c08301516001600160a01b03811660e08401525060e08301516101008381015261129a6101208401826111af565b949350505050565b634e487b7160e01b600052604160045260246000fd5b60405160a081016001600160401b03811182821017156112da576112da6112a2565b60405290565b600082601f8301126112f157600080fd5b81356001600160401b038082111561130b5761130b6112a2565b604051601f8301601f19908116603f01168101908282118183101715611333576113336112a2565b8160405283815286602085880101111561134c57600080fd5b836020870160208301376000602085830101528094505050505092915050565b60008060006060848603121561138157600080fd5b61138a846110f2565b92506020611399818601611142565b925060408501356001600160401b03808211156113b557600080fd5b818701915087601f8301126113c957600080fd5b6113d16112b8565b8060a084018a8111156113e357600080fd5b845b81811015611417578035858111156113fd5760008081fd5b6114098d8289016112e0565b8552509286019286016113e5565b505080955050505050509250925092565b6000806040838503121561143b57600080fd5b6114448361112b565b915061117c6020840161112b565b803560ff8116811461100e57600080fd5b6000806040838503121561147657600080fd5b61147f8361112b565b915061117c60208401611452565b60006020828403121561149f57600080fd5b6111248261112b565b6000806000606084860312156114bd57600080fd5b6114c68461112b565b92506114d460208501611452565b915060408401356001600160401b038111156114ef57600080fd5b6114fb868287016112e0565b9150509250925092565b6020808252600990820152683737ba1037bbb732b960b91b604082015260600190565b634e487b7160e01b600052603260045260246000fd5b600181811c9082168061155257607f821691505b60208210810361157257634e487b7160e01b600052602260045260246000fd5b50919050565b8281526040602080830182905260009160e084019190840185845b60058110156115c257603f198786030183526115b08583516111af565b94509183019190830190600101611593565b509297965050505050505056fea26469706673582212201fd3f1e883a1e667eda1ac1f797e72501ec9ade123f016486421752d4921668e64736f6c634300080e0033"

// DeployRole deploys a new Ethereum contract, binding an instance of Role to it.
func DeployRole(auth *bind.TransactOpts, backend bind.ContractBackend, _a common.Address, f common.Address) (common.Address, *types.Transaction, *Role, error) {
	parsed, err := abi.JSON(strings.NewReader(RoleABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RoleBin), backend, _a, f)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Role{RoleCaller: RoleCaller{contract: contract}, RoleTransactor: RoleTransactor{contract: contract}, RoleFilterer: RoleFilterer{contract: contract}}, nil
}

// Role is an auto generated Go binding around an Ethereum contract.
type Role struct {
	RoleCaller     // Read-only binding to the contract
	RoleTransactor // Write-only binding to the contract
	RoleFilterer   // Log filterer for contract events
}

// RoleCaller is an auto generated read-only Go binding around an Ethereum contract.
type RoleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RoleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RoleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RoleSession struct {
	Contract     *Role             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RoleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RoleCallerSession struct {
	Contract *RoleCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RoleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RoleTransactorSession struct {
	Contract     *RoleTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RoleRaw is an auto generated low-level Go binding around an Ethereum contract.
type RoleRaw struct {
	Contract *Role // Generic contract binding to access the raw methods on
}

// RoleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RoleCallerRaw struct {
	Contract *RoleCaller // Generic read-only contract binding to access the raw methods on
}

// RoleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RoleTransactorRaw struct {
	Contract *RoleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRole creates a new instance of Role, bound to a specific deployed contract.
func NewRole(address common.Address, backend bind.ContractBackend) (*Role, error) {
	contract, err := bindRole(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Role{RoleCaller: RoleCaller{contract: contract}, RoleTransactor: RoleTransactor{contract: contract}, RoleFilterer: RoleFilterer{contract: contract}}, nil
}

// NewRoleCaller creates a new read-only instance of Role, bound to a specific deployed contract.
func NewRoleCaller(address common.Address, caller bind.ContractCaller) (*RoleCaller, error) {
	contract, err := bindRole(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RoleCaller{contract: contract}, nil
}

// NewRoleTransactor creates a new write-only instance of Role, bound to a specific deployed contract.
func NewRoleTransactor(address common.Address, transactor bind.ContractTransactor) (*RoleTransactor, error) {
	contract, err := bindRole(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RoleTransactor{contract: contract}, nil
}

// NewRoleFilterer creates a new log filterer instance of Role, bound to a specific deployed contract.
func NewRoleFilterer(address common.Address, filterer bind.ContractFilterer) (*RoleFilterer, error) {
	contract, err := bindRole(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RoleFilterer{contract: contract}, nil
}

// bindRole binds a generic wrapper to an already deployed contract.
func bindRole(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RoleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Role *RoleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Role.Contract.RoleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Role *RoleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Role.Contract.RoleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Role *RoleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Role.Contract.RoleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Role *RoleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Role.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Role *RoleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Role.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Role *RoleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Role.Contract.contract.Transact(opts, method, params...)
}

// Auth is a free data retrieval call binding the contract method 0xde9375f2.
//
// Solidity: function auth() view returns(address)
func (_Role *RoleCaller) Auth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Role.contract.Call(opts, &out, "auth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Auth is a free data retrieval call binding the contract method 0xde9375f2.
//
// Solidity: function auth() view returns(address)
func (_Role *RoleSession) Auth() (common.Address, error) {
	return _Role.Contract.Auth(&_Role.CallOpts)
}

// Auth is a free data retrieval call binding the contract method 0xde9375f2.
//
// Solidity: function auth() view returns(address)
func (_Role *RoleCallerSession) Auth() (common.Address, error) {
	return _Role.Contract.Auth(&_Role.CallOpts)
}

// CheckIG is a free data retrieval call binding the contract method 0x7738515f.
//
// Solidity: function checkIG(uint64 _index, uint8 _rType) view returns(address, address, uint64, uint8)
func (_Role *RoleCaller) CheckIG(opts *bind.CallOpts, _index uint64, _rType uint8) (common.Address, common.Address, uint64, uint8, error) {
	var out []interface{}
	err := _Role.contract.Call(opts, &out, "checkIG", _index, _rType)

	if err != nil {
		return *new(common.Address), *new(common.Address), *new(uint64), *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	out2 := *abi.ConvertType(out[2], new(uint64)).(*uint64)
	out3 := *abi.ConvertType(out[3], new(uint8)).(*uint8)

	return out0, out1, out2, out3, err

}

// CheckIG is a free data retrieval call binding the contract method 0x7738515f.
//
// Solidity: function checkIG(uint64 _index, uint8 _rType) view returns(address, address, uint64, uint8)
func (_Role *RoleSession) CheckIG(_index uint64, _rType uint8) (common.Address, common.Address, uint64, uint8, error) {
	return _Role.Contract.CheckIG(&_Role.CallOpts, _index, _rType)
}

// CheckIG is a free data retrieval call binding the contract method 0x7738515f.
//
// Solidity: function checkIG(uint64 _index, uint8 _rType) view returns(address, address, uint64, uint8)
func (_Role *RoleCallerSession) CheckIG(_index uint64, _rType uint8) (common.Address, common.Address, uint64, uint8, error) {
	return _Role.Contract.CheckIG(&_Role.CallOpts, _index, _rType)
}

// GetACnt is a free data retrieval call binding the contract method 0x7264a551.
//
// Solidity: function getACnt() view returns(uint64)
func (_Role *RoleCaller) GetACnt(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Role.contract.Call(opts, &out, "getACnt")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetACnt is a free data retrieval call binding the contract method 0x7264a551.
//
// Solidity: function getACnt() view returns(uint64)
func (_Role *RoleSession) GetACnt() (uint64, error) {
	return _Role.Contract.GetACnt(&_Role.CallOpts)
}

// GetACnt is a free data retrieval call binding the contract method 0x7264a551.
//
// Solidity: function getACnt() view returns(uint64)
func (_Role *RoleCallerSession) GetACnt() (uint64, error) {
	return _Role.Contract.GetACnt(&_Role.CallOpts)
}

// GetAcc is a free data retrieval call binding the contract method 0xcaca4a06.
//
// Solidity: function getAcc(address _a) view returns(uint64)
func (_Role *RoleCaller) GetAcc(opts *bind.CallOpts, _a common.Address) (uint64, error) {
	var out []interface{}
	err := _Role.contract.Call(opts, &out, "getAcc", _a)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetAcc is a free data retrieval call binding the contract method 0xcaca4a06.
//
// Solidity: function getAcc(address _a) view returns(uint64)
func (_Role *RoleSession) GetAcc(_a common.Address) (uint64, error) {
	return _Role.Contract.GetAcc(&_Role.CallOpts, _a)
}

// GetAcc is a free data retrieval call binding the contract method 0xcaca4a06.
//
// Solidity: function getAcc(address _a) view returns(uint64)
func (_Role *RoleCallerSession) GetAcc(_a common.Address) (uint64, error) {
	return _Role.Contract.GetAcc(&_Role.CallOpts, _a)
}

// GetAddr is a free data retrieval call binding the contract method 0x9332aa6e.
//
// Solidity: function getAddr(uint64 _i) view returns(address)
func (_Role *RoleCaller) GetAddr(opts *bind.CallOpts, _i uint64) (common.Address, error) {
	var out []interface{}
	err := _Role.contract.Call(opts, &out, "getAddr", _i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddr is a free data retrieval call binding the contract method 0x9332aa6e.
//
// Solidity: function getAddr(uint64 _i) view returns(address)
func (_Role *RoleSession) GetAddr(_i uint64) (common.Address, error) {
	return _Role.Contract.GetAddr(&_Role.CallOpts, _i)
}

// GetAddr is a free data retrieval call binding the contract method 0x9332aa6e.
//
// Solidity: function getAddr(uint64 _i) view returns(address)
func (_Role *RoleCallerSession) GetAddr(_i uint64) (common.Address, error) {
	return _Role.Contract.GetAddr(&_Role.CallOpts, _i)
}

// GetRInfo is a free data retrieval call binding the contract method 0x441abace.
//
// Solidity: function getRInfo(address acc) view returns((bool,bool,uint8,uint64,uint64,address,address,bytes))
func (_Role *RoleCaller) GetRInfo(opts *bind.CallOpts, acc common.Address) (RoleOut, error) {
	var out []interface{}
	err := _Role.contract.Call(opts, &out, "getRInfo", acc)

	if err != nil {
		return *new(RoleOut), err
	}

	out0 := *abi.ConvertType(out[0], new(RoleOut)).(*RoleOut)

	return out0, err

}

// GetRInfo is a free data retrieval call binding the contract method 0x441abace.
//
// Solidity: function getRInfo(address acc) view returns((bool,bool,uint8,uint64,uint64,address,address,bytes))
func (_Role *RoleSession) GetRInfo(acc common.Address) (RoleOut, error) {
	return _Role.Contract.GetRInfo(&_Role.CallOpts, acc)
}

// GetRInfo is a free data retrieval call binding the contract method 0x441abace.
//
// Solidity: function getRInfo(address acc) view returns((bool,bool,uint8,uint64,uint64,address,address,bytes))
func (_Role *RoleCallerSession) GetRInfo(acc common.Address) (RoleOut, error) {
	return _Role.Contract.GetRInfo(&_Role.CallOpts, acc)
}

// Owners is a free data retrieval call binding the contract method 0x022914a7.
//
// Solidity: function owners(address ) view returns(bool)
func (_Role *RoleCaller) Owners(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Role.contract.Call(opts, &out, "owners", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Owners is a free data retrieval call binding the contract method 0x022914a7.
//
// Solidity: function owners(address ) view returns(bool)
func (_Role *RoleSession) Owners(arg0 common.Address) (bool, error) {
	return _Role.Contract.Owners(&_Role.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x022914a7.
//
// Solidity: function owners(address ) view returns(bool)
func (_Role *RoleCallerSession) Owners(arg0 common.Address) (bool, error) {
	return _Role.Contract.Owners(&_Role.CallOpts, arg0)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint16)
func (_Role *RoleCaller) Version(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Role.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint16)
func (_Role *RoleSession) Version() (uint16, error) {
	return _Role.Contract.Version(&_Role.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint16)
func (_Role *RoleCallerSession) Version() (uint16, error) {
	return _Role.Contract.Version(&_Role.CallOpts)
}

// Activate is a paid mutator transaction binding the contract method 0xfc7e4746.
//
// Solidity: function activate(uint64 _index) returns(uint64)
func (_Role *RoleTransactor) Activate(opts *bind.TransactOpts, _index uint64) (*types.Transaction, error) {
	return _Role.contract.Transact(opts, "activate", _index)
}

// Activate is a paid mutator transaction binding the contract method 0xfc7e4746.
//
// Solidity: function activate(uint64 _index) returns(uint64)
func (_Role *RoleSession) Activate(_index uint64) (*types.Transaction, error) {
	return _Role.Contract.Activate(&_Role.TransactOpts, _index)
}

// Activate is a paid mutator transaction binding the contract method 0xfc7e4746.
//
// Solidity: function activate(uint64 _index) returns(uint64)
func (_Role *RoleTransactorSession) Activate(_index uint64) (*types.Transaction, error) {
	return _Role.Contract.Activate(&_Role.TransactOpts, _index)
}

// Add is a paid mutator transaction binding the contract method 0x4bf1b134.
//
// Solidity: function add(address _a, bool isSet, bytes[5] signs) returns()
func (_Role *RoleTransactor) Add(opts *bind.TransactOpts, _a common.Address, isSet bool, signs [5][]byte) (*types.Transaction, error) {
	return _Role.contract.Transact(opts, "add", _a, isSet, signs)
}

// Add is a paid mutator transaction binding the contract method 0x4bf1b134.
//
// Solidity: function add(address _a, bool isSet, bytes[5] signs) returns()
func (_Role *RoleSession) Add(_a common.Address, isSet bool, signs [5][]byte) (*types.Transaction, error) {
	return _Role.Contract.Add(&_Role.TransactOpts, _a, isSet, signs)
}

// Add is a paid mutator transaction binding the contract method 0x4bf1b134.
//
// Solidity: function add(address _a, bool isSet, bytes[5] signs) returns()
func (_Role *RoleTransactorSession) Add(_a common.Address, isSet bool, signs [5][]byte) (*types.Transaction, error) {
	return _Role.Contract.Add(&_Role.TransactOpts, _a, isSet, signs)
}

// AlterPayee is a paid mutator transaction binding the contract method 0x2915824a.
//
// Solidity: function alterPayee(uint64 _index, address _p) returns()
func (_Role *RoleTransactor) AlterPayee(opts *bind.TransactOpts, _index uint64, _p common.Address) (*types.Transaction, error) {
	return _Role.contract.Transact(opts, "alterPayee", _index, _p)
}

// AlterPayee is a paid mutator transaction binding the contract method 0x2915824a.
//
// Solidity: function alterPayee(uint64 _index, address _p) returns()
func (_Role *RoleSession) AlterPayee(_index uint64, _p common.Address) (*types.Transaction, error) {
	return _Role.Contract.AlterPayee(&_Role.TransactOpts, _index, _p)
}

// AlterPayee is a paid mutator transaction binding the contract method 0x2915824a.
//
// Solidity: function alterPayee(uint64 _index, address _p) returns()
func (_Role *RoleTransactorSession) AlterPayee(_index uint64, _p common.Address) (*types.Transaction, error) {
	return _Role.Contract.AlterPayee(&_Role.TransactOpts, _index, _p)
}

// Ban is a paid mutator transaction binding the contract method 0x0f78c61a.
//
// Solidity: function ban(uint64 _index, bool _banned) returns()
func (_Role *RoleTransactor) Ban(opts *bind.TransactOpts, _index uint64, _banned bool) (*types.Transaction, error) {
	return _Role.contract.Transact(opts, "ban", _index, _banned)
}

// Ban is a paid mutator transaction binding the contract method 0x0f78c61a.
//
// Solidity: function ban(uint64 _index, bool _banned) returns()
func (_Role *RoleSession) Ban(_index uint64, _banned bool) (*types.Transaction, error) {
	return _Role.Contract.Ban(&_Role.TransactOpts, _index, _banned)
}

// Ban is a paid mutator transaction binding the contract method 0x0f78c61a.
//
// Solidity: function ban(uint64 _index, bool _banned) returns()
func (_Role *RoleTransactorSession) Ban(_index uint64, _banned bool) (*types.Transaction, error) {
	return _Role.Contract.Ban(&_Role.TransactOpts, _index, _banned)
}

// ConfirmPayee is a paid mutator transaction binding the contract method 0x5bd04c3f.
//
// Solidity: function confirmPayee(uint64 _index, address _p) returns()
func (_Role *RoleTransactor) ConfirmPayee(opts *bind.TransactOpts, _index uint64, _p common.Address) (*types.Transaction, error) {
	return _Role.contract.Transact(opts, "confirmPayee", _index, _p)
}

// ConfirmPayee is a paid mutator transaction binding the contract method 0x5bd04c3f.
//
// Solidity: function confirmPayee(uint64 _index, address _p) returns()
func (_Role *RoleSession) ConfirmPayee(_index uint64, _p common.Address) (*types.Transaction, error) {
	return _Role.Contract.ConfirmPayee(&_Role.TransactOpts, _index, _p)
}

// ConfirmPayee is a paid mutator transaction binding the contract method 0x5bd04c3f.
//
// Solidity: function confirmPayee(uint64 _index, address _p) returns()
func (_Role *RoleTransactorSession) ConfirmPayee(_index uint64, _p common.Address) (*types.Transaction, error) {
	return _Role.Contract.ConfirmPayee(&_Role.TransactOpts, _index, _p)
}

// ReAcc is a paid mutator transaction binding the contract method 0xeffcafce.
//
// Solidity: function reAcc(address a) returns()
func (_Role *RoleTransactor) ReAcc(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _Role.contract.Transact(opts, "reAcc", a)
}

// ReAcc is a paid mutator transaction binding the contract method 0xeffcafce.
//
// Solidity: function reAcc(address a) returns()
func (_Role *RoleSession) ReAcc(a common.Address) (*types.Transaction, error) {
	return _Role.Contract.ReAcc(&_Role.TransactOpts, a)
}

// ReAcc is a paid mutator transaction binding the contract method 0xeffcafce.
//
// Solidity: function reAcc(address a) returns()
func (_Role *RoleTransactorSession) ReAcc(a common.Address) (*types.Transaction, error) {
	return _Role.Contract.ReAcc(&_Role.TransactOpts, a)
}

// ReRole is a paid mutator transaction binding the contract method 0xb9f6a8ca.
//
// Solidity: function reRole(uint64 _index, uint8 _rType, bytes _extra) returns()
func (_Role *RoleTransactor) ReRole(opts *bind.TransactOpts, _index uint64, _rType uint8, _extra []byte) (*types.Transaction, error) {
	return _Role.contract.Transact(opts, "reRole", _index, _rType, _extra)
}

// ReRole is a paid mutator transaction binding the contract method 0xb9f6a8ca.
//
// Solidity: function reRole(uint64 _index, uint8 _rType, bytes _extra) returns()
func (_Role *RoleSession) ReRole(_index uint64, _rType uint8, _extra []byte) (*types.Transaction, error) {
	return _Role.Contract.ReRole(&_Role.TransactOpts, _index, _rType, _extra)
}

// ReRole is a paid mutator transaction binding the contract method 0xb9f6a8ca.
//
// Solidity: function reRole(uint64 _index, uint8 _rType, bytes _extra) returns()
func (_Role *RoleTransactorSession) ReRole(_index uint64, _rType uint8, _extra []byte) (*types.Transaction, error) {
	return _Role.Contract.ReRole(&_Role.TransactOpts, _index, _rType, _extra)
}

// SetG is a paid mutator transaction binding the contract method 0x616bfd1f.
//
// Solidity: function setG(uint64 _index, uint64 _gi) returns()
func (_Role *RoleTransactor) SetG(opts *bind.TransactOpts, _index uint64, _gi uint64) (*types.Transaction, error) {
	return _Role.contract.Transact(opts, "setG", _index, _gi)
}

// SetG is a paid mutator transaction binding the contract method 0x616bfd1f.
//
// Solidity: function setG(uint64 _index, uint64 _gi) returns()
func (_Role *RoleSession) SetG(_index uint64, _gi uint64) (*types.Transaction, error) {
	return _Role.Contract.SetG(&_Role.TransactOpts, _index, _gi)
}

// SetG is a paid mutator transaction binding the contract method 0x616bfd1f.
//
// Solidity: function setG(uint64 _index, uint64 _gi) returns()
func (_Role *RoleTransactorSession) SetG(_index uint64, _gi uint64) (*types.Transaction, error) {
	return _Role.Contract.SetG(&_Role.TransactOpts, _index, _gi)
}

// RoleAddOwnerIterator is returned from FilterAddOwner and is used to iterate over the raw logs and unpacked data for AddOwner events raised by the Role contract.
type RoleAddOwnerIterator struct {
	Event *RoleAddOwner // Event containing the contract specifics and raw log

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
func (it *RoleAddOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoleAddOwner)
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
		it.Event = new(RoleAddOwner)
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
func (it *RoleAddOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoleAddOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoleAddOwner represents a AddOwner event raised by the Role contract.
type RoleAddOwner struct {
	From  common.Address
	IsSet bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAddOwner is a free log retrieval operation binding the contract event 0x938b2a24c98e4e542127ffa74a91e48942c2bddccae3b6d75f82bfda7bbc0807.
//
// Solidity: event AddOwner(address from, bool isSet)
func (_Role *RoleFilterer) FilterAddOwner(opts *bind.FilterOpts) (*RoleAddOwnerIterator, error) {

	logs, sub, err := _Role.contract.FilterLogs(opts, "AddOwner")
	if err != nil {
		return nil, err
	}
	return &RoleAddOwnerIterator{contract: _Role.contract, event: "AddOwner", logs: logs, sub: sub}, nil
}

// WatchAddOwner is a free log subscription operation binding the contract event 0x938b2a24c98e4e542127ffa74a91e48942c2bddccae3b6d75f82bfda7bbc0807.
//
// Solidity: event AddOwner(address from, bool isSet)
func (_Role *RoleFilterer) WatchAddOwner(opts *bind.WatchOpts, sink chan<- *RoleAddOwner) (event.Subscription, error) {

	logs, sub, err := _Role.contract.WatchLogs(opts, "AddOwner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoleAddOwner)
				if err := _Role.contract.UnpackLog(event, "AddOwner", log); err != nil {
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
func (_Role *RoleFilterer) ParseAddOwner(log types.Log) (*RoleAddOwner, error) {
	event := new(RoleAddOwner)
	if err := _Role.contract.UnpackLog(event, "AddOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoleAlterPayeeIterator is returned from FilterAlterPayee and is used to iterate over the raw logs and unpacked data for AlterPayee events raised by the Role contract.
type RoleAlterPayeeIterator struct {
	Event *RoleAlterPayee // Event containing the contract specifics and raw log

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
func (it *RoleAlterPayeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoleAlterPayee)
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
		it.Event = new(RoleAlterPayee)
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
func (it *RoleAlterPayeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoleAlterPayeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoleAlterPayee represents a AlterPayee event raised by the Role contract.
type RoleAlterPayee struct {
	Payee common.Address
	Addr  common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAlterPayee is a free log retrieval operation binding the contract event 0x3b4058e93b2a019add56609a3c2463a7b3d8476abf30c934cfcdb9f2158b5e68.
//
// Solidity: event AlterPayee(address indexed payee, address addr)
func (_Role *RoleFilterer) FilterAlterPayee(opts *bind.FilterOpts, payee []common.Address) (*RoleAlterPayeeIterator, error) {

	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}

	logs, sub, err := _Role.contract.FilterLogs(opts, "AlterPayee", payeeRule)
	if err != nil {
		return nil, err
	}
	return &RoleAlterPayeeIterator{contract: _Role.contract, event: "AlterPayee", logs: logs, sub: sub}, nil
}

// WatchAlterPayee is a free log subscription operation binding the contract event 0x3b4058e93b2a019add56609a3c2463a7b3d8476abf30c934cfcdb9f2158b5e68.
//
// Solidity: event AlterPayee(address indexed payee, address addr)
func (_Role *RoleFilterer) WatchAlterPayee(opts *bind.WatchOpts, sink chan<- *RoleAlterPayee, payee []common.Address) (event.Subscription, error) {

	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}

	logs, sub, err := _Role.contract.WatchLogs(opts, "AlterPayee", payeeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoleAlterPayee)
				if err := _Role.contract.UnpackLog(event, "AlterPayee", log); err != nil {
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

// ParseAlterPayee is a log parse operation binding the contract event 0x3b4058e93b2a019add56609a3c2463a7b3d8476abf30c934cfcdb9f2158b5e68.
//
// Solidity: event AlterPayee(address indexed payee, address addr)
func (_Role *RoleFilterer) ParseAlterPayee(log types.Log) (*RoleAlterPayee, error) {
	event := new(RoleAlterPayee)
	if err := _Role.contract.UnpackLog(event, "AlterPayee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoleConfirmPayeeIterator is returned from FilterConfirmPayee and is used to iterate over the raw logs and unpacked data for ConfirmPayee events raised by the Role contract.
type RoleConfirmPayeeIterator struct {
	Event *RoleConfirmPayee // Event containing the contract specifics and raw log

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
func (it *RoleConfirmPayeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoleConfirmPayee)
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
		it.Event = new(RoleConfirmPayee)
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
func (it *RoleConfirmPayeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoleConfirmPayeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoleConfirmPayee represents a ConfirmPayee event raised by the Role contract.
type RoleConfirmPayee struct {
	Payee common.Address
	Addr  common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterConfirmPayee is a free log retrieval operation binding the contract event 0x4a21d873c33448e8dfb90cdbfdb12f4aed5f5e2abde5525215b1655d519784dc.
//
// Solidity: event ConfirmPayee(address indexed payee, address addr)
func (_Role *RoleFilterer) FilterConfirmPayee(opts *bind.FilterOpts, payee []common.Address) (*RoleConfirmPayeeIterator, error) {

	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}

	logs, sub, err := _Role.contract.FilterLogs(opts, "ConfirmPayee", payeeRule)
	if err != nil {
		return nil, err
	}
	return &RoleConfirmPayeeIterator{contract: _Role.contract, event: "ConfirmPayee", logs: logs, sub: sub}, nil
}

// WatchConfirmPayee is a free log subscription operation binding the contract event 0x4a21d873c33448e8dfb90cdbfdb12f4aed5f5e2abde5525215b1655d519784dc.
//
// Solidity: event ConfirmPayee(address indexed payee, address addr)
func (_Role *RoleFilterer) WatchConfirmPayee(opts *bind.WatchOpts, sink chan<- *RoleConfirmPayee, payee []common.Address) (event.Subscription, error) {

	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}

	logs, sub, err := _Role.contract.WatchLogs(opts, "ConfirmPayee", payeeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoleConfirmPayee)
				if err := _Role.contract.UnpackLog(event, "ConfirmPayee", log); err != nil {
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

// ParseConfirmPayee is a log parse operation binding the contract event 0x4a21d873c33448e8dfb90cdbfdb12f4aed5f5e2abde5525215b1655d519784dc.
//
// Solidity: event ConfirmPayee(address indexed payee, address addr)
func (_Role *RoleFilterer) ParseConfirmPayee(log types.Log) (*RoleConfirmPayee, error) {
	event := new(RoleConfirmPayee)
	if err := _Role.contract.UnpackLog(event, "ConfirmPayee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoleReAccIterator is returned from FilterReAcc and is used to iterate over the raw logs and unpacked data for ReAcc events raised by the Role contract.
type RoleReAccIterator struct {
	Event *RoleReAcc // Event containing the contract specifics and raw log

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
func (it *RoleReAccIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoleReAcc)
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
		it.Event = new(RoleReAcc)
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
func (it *RoleReAccIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoleReAccIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoleReAcc represents a ReAcc event raised by the Role contract.
type RoleReAcc struct {
	Addr  common.Address
	Index uint64
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterReAcc is a free log retrieval operation binding the contract event 0xe22ccfc11bbbb1d3350012248b26ad616e32a8a13eece408dd14f1953fe24752.
//
// Solidity: event ReAcc(address addr, uint64 index)
func (_Role *RoleFilterer) FilterReAcc(opts *bind.FilterOpts) (*RoleReAccIterator, error) {

	logs, sub, err := _Role.contract.FilterLogs(opts, "ReAcc")
	if err != nil {
		return nil, err
	}
	return &RoleReAccIterator{contract: _Role.contract, event: "ReAcc", logs: logs, sub: sub}, nil
}

// WatchReAcc is a free log subscription operation binding the contract event 0xe22ccfc11bbbb1d3350012248b26ad616e32a8a13eece408dd14f1953fe24752.
//
// Solidity: event ReAcc(address addr, uint64 index)
func (_Role *RoleFilterer) WatchReAcc(opts *bind.WatchOpts, sink chan<- *RoleReAcc) (event.Subscription, error) {

	logs, sub, err := _Role.contract.WatchLogs(opts, "ReAcc")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoleReAcc)
				if err := _Role.contract.UnpackLog(event, "ReAcc", log); err != nil {
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

// ParseReAcc is a log parse operation binding the contract event 0xe22ccfc11bbbb1d3350012248b26ad616e32a8a13eece408dd14f1953fe24752.
//
// Solidity: event ReAcc(address addr, uint64 index)
func (_Role *RoleFilterer) ParseReAcc(log types.Log) (*RoleReAcc, error) {
	event := new(RoleReAcc)
	if err := _Role.contract.UnpackLog(event, "ReAcc", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoleReRoleIterator is returned from FilterReRole and is used to iterate over the raw logs and unpacked data for ReRole events raised by the Role contract.
type RoleReRoleIterator struct {
	Event *RoleReRole // Event containing the contract specifics and raw log

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
func (it *RoleReRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoleReRole)
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
		it.Event = new(RoleReRole)
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
func (it *RoleReRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoleReRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoleReRole represents a ReRole event raised by the Role contract.
type RoleReRole struct {
	RType uint8
	Index uint64
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterReRole is a free log retrieval operation binding the contract event 0x5550fbab402c3b014955818e5b6e3cefef20641aa6bed43ae83cd5d3258b8922.
//
// Solidity: event ReRole(uint8 indexed rType, uint64 index)
func (_Role *RoleFilterer) FilterReRole(opts *bind.FilterOpts, rType []uint8) (*RoleReRoleIterator, error) {

	var rTypeRule []interface{}
	for _, rTypeItem := range rType {
		rTypeRule = append(rTypeRule, rTypeItem)
	}

	logs, sub, err := _Role.contract.FilterLogs(opts, "ReRole", rTypeRule)
	if err != nil {
		return nil, err
	}
	return &RoleReRoleIterator{contract: _Role.contract, event: "ReRole", logs: logs, sub: sub}, nil
}

// WatchReRole is a free log subscription operation binding the contract event 0x5550fbab402c3b014955818e5b6e3cefef20641aa6bed43ae83cd5d3258b8922.
//
// Solidity: event ReRole(uint8 indexed rType, uint64 index)
func (_Role *RoleFilterer) WatchReRole(opts *bind.WatchOpts, sink chan<- *RoleReRole, rType []uint8) (event.Subscription, error) {

	var rTypeRule []interface{}
	for _, rTypeItem := range rType {
		rTypeRule = append(rTypeRule, rTypeItem)
	}

	logs, sub, err := _Role.contract.WatchLogs(opts, "ReRole", rTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoleReRole)
				if err := _Role.contract.UnpackLog(event, "ReRole", log); err != nil {
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

// ParseReRole is a log parse operation binding the contract event 0x5550fbab402c3b014955818e5b6e3cefef20641aa6bed43ae83cd5d3258b8922.
//
// Solidity: event ReRole(uint8 indexed rType, uint64 index)
func (_Role *RoleFilterer) ParseReRole(log types.Log) (*RoleReRole, error) {
	event := new(RoleReRole)
	if err := _Role.contract.UnpackLog(event, "ReRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoleSetGIterator is returned from FilterSetG and is used to iterate over the raw logs and unpacked data for SetG events raised by the Role contract.
type RoleSetGIterator struct {
	Event *RoleSetG // Event containing the contract specifics and raw log

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
func (it *RoleSetGIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoleSetG)
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
		it.Event = new(RoleSetG)
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
func (it *RoleSetGIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoleSetGIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoleSetG represents a SetG event raised by the Role contract.
type RoleSetG struct {
	Gi    uint64
	Index uint64
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSetG is a free log retrieval operation binding the contract event 0xe7329265e8e4e2fc8e892a75f1809767d09c472bb2eef16f889eceb5583291a9.
//
// Solidity: event SetG(uint64 indexed gi, uint64 index)
func (_Role *RoleFilterer) FilterSetG(opts *bind.FilterOpts, gi []uint64) (*RoleSetGIterator, error) {

	var giRule []interface{}
	for _, giItem := range gi {
		giRule = append(giRule, giItem)
	}

	logs, sub, err := _Role.contract.FilterLogs(opts, "SetG", giRule)
	if err != nil {
		return nil, err
	}
	return &RoleSetGIterator{contract: _Role.contract, event: "SetG", logs: logs, sub: sub}, nil
}

// WatchSetG is a free log subscription operation binding the contract event 0xe7329265e8e4e2fc8e892a75f1809767d09c472bb2eef16f889eceb5583291a9.
//
// Solidity: event SetG(uint64 indexed gi, uint64 index)
func (_Role *RoleFilterer) WatchSetG(opts *bind.WatchOpts, sink chan<- *RoleSetG, gi []uint64) (event.Subscription, error) {

	var giRule []interface{}
	for _, giItem := range gi {
		giRule = append(giRule, giItem)
	}

	logs, sub, err := _Role.contract.WatchLogs(opts, "SetG", giRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoleSetG)
				if err := _Role.contract.UnpackLog(event, "SetG", log); err != nil {
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

// ParseSetG is a log parse operation binding the contract event 0xe7329265e8e4e2fc8e892a75f1809767d09c472bb2eef16f889eceb5583291a9.
//
// Solidity: event SetG(uint64 indexed gi, uint64 index)
func (_Role *RoleFilterer) ParseSetG(log types.Log) (*RoleSetG, error) {
	event := new(RoleSetG)
	if err := _Role.contract.UnpackLog(event, "SetG", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
