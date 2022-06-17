// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package token

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

// ITokenABI is the input ABI used to generate the binding from.
const ITokenABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"tIndex\",\"type\":\"uint8\"}],\"name\":\"AddT\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"}],\"name\":\"BanT\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_t\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_mint\",\"type\":\"bool\"}],\"name\":\"addT\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_t\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isBan\",\"type\":\"bool\"}],\"name\":\"banT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"}],\"name\":\"checkT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"}],\"name\":\"getTA\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTCnt\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_t\",\"type\":\"address\"}],\"name\":\"getTI\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ITokenFuncSigs maps the 4-byte function signature to its string representation.
var ITokenFuncSigs = map[string]string{
	"ff4268a4": "addT(address,bool)",
	"83373b36": "banT(address,bool)",
	"81abb8fe": "checkT(uint8)",
	"8bb4a637": "getTA(uint8)",
	"7600b86a": "getTCnt()",
	"2df2685f": "getTI(address)",
}

// IToken is an auto generated Go binding around an Ethereum contract.
type IToken struct {
	ITokenCaller     // Read-only binding to the contract
	ITokenTransactor // Write-only binding to the contract
	ITokenFilterer   // Log filterer for contract events
}

// ITokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type ITokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ITokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ITokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ITokenSession struct {
	Contract     *IToken           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ITokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ITokenCallerSession struct {
	Contract *ITokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ITokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ITokenTransactorSession struct {
	Contract     *ITokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ITokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type ITokenRaw struct {
	Contract *IToken // Generic contract binding to access the raw methods on
}

// ITokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ITokenCallerRaw struct {
	Contract *ITokenCaller // Generic read-only contract binding to access the raw methods on
}

// ITokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ITokenTransactorRaw struct {
	Contract *ITokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIToken creates a new instance of IToken, bound to a specific deployed contract.
func NewIToken(address common.Address, backend bind.ContractBackend) (*IToken, error) {
	contract, err := bindIToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IToken{ITokenCaller: ITokenCaller{contract: contract}, ITokenTransactor: ITokenTransactor{contract: contract}, ITokenFilterer: ITokenFilterer{contract: contract}}, nil
}

// NewITokenCaller creates a new read-only instance of IToken, bound to a specific deployed contract.
func NewITokenCaller(address common.Address, caller bind.ContractCaller) (*ITokenCaller, error) {
	contract, err := bindIToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ITokenCaller{contract: contract}, nil
}

// NewITokenTransactor creates a new write-only instance of IToken, bound to a specific deployed contract.
func NewITokenTransactor(address common.Address, transactor bind.ContractTransactor) (*ITokenTransactor, error) {
	contract, err := bindIToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ITokenTransactor{contract: contract}, nil
}

// NewITokenFilterer creates a new log filterer instance of IToken, bound to a specific deployed contract.
func NewITokenFilterer(address common.Address, filterer bind.ContractFilterer) (*ITokenFilterer, error) {
	contract, err := bindIToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ITokenFilterer{contract: contract}, nil
}

// bindIToken binds a generic wrapper to an already deployed contract.
func bindIToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ITokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IToken *ITokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IToken.Contract.ITokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IToken *ITokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IToken.Contract.ITokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IToken *ITokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IToken.Contract.ITokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IToken *ITokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IToken *ITokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IToken *ITokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IToken.Contract.contract.Transact(opts, method, params...)
}

// CheckT is a free data retrieval call binding the contract method 0x81abb8fe.
//
// Solidity: function checkT(uint8 _ti) view returns(address, bool)
func (_IToken *ITokenCaller) CheckT(opts *bind.CallOpts, _ti uint8) (common.Address, bool, error) {
	var out []interface{}
	err := _IToken.contract.Call(opts, &out, "checkT", _ti)

	if err != nil {
		return *new(common.Address), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// CheckT is a free data retrieval call binding the contract method 0x81abb8fe.
//
// Solidity: function checkT(uint8 _ti) view returns(address, bool)
func (_IToken *ITokenSession) CheckT(_ti uint8) (common.Address, bool, error) {
	return _IToken.Contract.CheckT(&_IToken.CallOpts, _ti)
}

// CheckT is a free data retrieval call binding the contract method 0x81abb8fe.
//
// Solidity: function checkT(uint8 _ti) view returns(address, bool)
func (_IToken *ITokenCallerSession) CheckT(_ti uint8) (common.Address, bool, error) {
	return _IToken.Contract.CheckT(&_IToken.CallOpts, _ti)
}

// GetTA is a free data retrieval call binding the contract method 0x8bb4a637.
//
// Solidity: function getTA(uint8 _ti) view returns(address)
func (_IToken *ITokenCaller) GetTA(opts *bind.CallOpts, _ti uint8) (common.Address, error) {
	var out []interface{}
	err := _IToken.contract.Call(opts, &out, "getTA", _ti)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTA is a free data retrieval call binding the contract method 0x8bb4a637.
//
// Solidity: function getTA(uint8 _ti) view returns(address)
func (_IToken *ITokenSession) GetTA(_ti uint8) (common.Address, error) {
	return _IToken.Contract.GetTA(&_IToken.CallOpts, _ti)
}

// GetTA is a free data retrieval call binding the contract method 0x8bb4a637.
//
// Solidity: function getTA(uint8 _ti) view returns(address)
func (_IToken *ITokenCallerSession) GetTA(_ti uint8) (common.Address, error) {
	return _IToken.Contract.GetTA(&_IToken.CallOpts, _ti)
}

// GetTCnt is a free data retrieval call binding the contract method 0x7600b86a.
//
// Solidity: function getTCnt() view returns(uint8)
func (_IToken *ITokenCaller) GetTCnt(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _IToken.contract.Call(opts, &out, "getTCnt")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetTCnt is a free data retrieval call binding the contract method 0x7600b86a.
//
// Solidity: function getTCnt() view returns(uint8)
func (_IToken *ITokenSession) GetTCnt() (uint8, error) {
	return _IToken.Contract.GetTCnt(&_IToken.CallOpts)
}

// GetTCnt is a free data retrieval call binding the contract method 0x7600b86a.
//
// Solidity: function getTCnt() view returns(uint8)
func (_IToken *ITokenCallerSession) GetTCnt() (uint8, error) {
	return _IToken.Contract.GetTCnt(&_IToken.CallOpts)
}

// GetTI is a free data retrieval call binding the contract method 0x2df2685f.
//
// Solidity: function getTI(address _t) view returns(uint8, bool, bool)
func (_IToken *ITokenCaller) GetTI(opts *bind.CallOpts, _t common.Address) (uint8, bool, bool, error) {
	var out []interface{}
	err := _IToken.contract.Call(opts, &out, "getTI", _t)

	if err != nil {
		return *new(uint8), *new(bool), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)
	out2 := *abi.ConvertType(out[2], new(bool)).(*bool)

	return out0, out1, out2, err

}

// GetTI is a free data retrieval call binding the contract method 0x2df2685f.
//
// Solidity: function getTI(address _t) view returns(uint8, bool, bool)
func (_IToken *ITokenSession) GetTI(_t common.Address) (uint8, bool, bool, error) {
	return _IToken.Contract.GetTI(&_IToken.CallOpts, _t)
}

// GetTI is a free data retrieval call binding the contract method 0x2df2685f.
//
// Solidity: function getTI(address _t) view returns(uint8, bool, bool)
func (_IToken *ITokenCallerSession) GetTI(_t common.Address) (uint8, bool, bool, error) {
	return _IToken.Contract.GetTI(&_IToken.CallOpts, _t)
}

// AddT is a paid mutator transaction binding the contract method 0xff4268a4.
//
// Solidity: function addT(address _t, bool _mint) returns(uint8)
func (_IToken *ITokenTransactor) AddT(opts *bind.TransactOpts, _t common.Address, _mint bool) (*types.Transaction, error) {
	return _IToken.contract.Transact(opts, "addT", _t, _mint)
}

// AddT is a paid mutator transaction binding the contract method 0xff4268a4.
//
// Solidity: function addT(address _t, bool _mint) returns(uint8)
func (_IToken *ITokenSession) AddT(_t common.Address, _mint bool) (*types.Transaction, error) {
	return _IToken.Contract.AddT(&_IToken.TransactOpts, _t, _mint)
}

// AddT is a paid mutator transaction binding the contract method 0xff4268a4.
//
// Solidity: function addT(address _t, bool _mint) returns(uint8)
func (_IToken *ITokenTransactorSession) AddT(_t common.Address, _mint bool) (*types.Transaction, error) {
	return _IToken.Contract.AddT(&_IToken.TransactOpts, _t, _mint)
}

// BanT is a paid mutator transaction binding the contract method 0x83373b36.
//
// Solidity: function banT(address _t, bool _isBan) returns()
func (_IToken *ITokenTransactor) BanT(opts *bind.TransactOpts, _t common.Address, _isBan bool) (*types.Transaction, error) {
	return _IToken.contract.Transact(opts, "banT", _t, _isBan)
}

// BanT is a paid mutator transaction binding the contract method 0x83373b36.
//
// Solidity: function banT(address _t, bool _isBan) returns()
func (_IToken *ITokenSession) BanT(_t common.Address, _isBan bool) (*types.Transaction, error) {
	return _IToken.Contract.BanT(&_IToken.TransactOpts, _t, _isBan)
}

// BanT is a paid mutator transaction binding the contract method 0x83373b36.
//
// Solidity: function banT(address _t, bool _isBan) returns()
func (_IToken *ITokenTransactorSession) BanT(_t common.Address, _isBan bool) (*types.Transaction, error) {
	return _IToken.Contract.BanT(&_IToken.TransactOpts, _t, _isBan)
}

// ITokenAddTIterator is returned from FilterAddT and is used to iterate over the raw logs and unpacked data for AddT events raised by the IToken contract.
type ITokenAddTIterator struct {
	Event *ITokenAddT // Event containing the contract specifics and raw log

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
func (it *ITokenAddTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITokenAddT)
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
		it.Event = new(ITokenAddT)
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
func (it *ITokenAddTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITokenAddTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITokenAddT represents a AddT event raised by the IToken contract.
type ITokenAddT struct {
	T      common.Address
	TIndex uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAddT is a free log retrieval operation binding the contract event 0xf10191767d4ea3fc26b057d307336c7e8df8880bb07ddaebe4e853db6b5fd936.
//
// Solidity: event AddT(address t, uint8 tIndex)
func (_IToken *ITokenFilterer) FilterAddT(opts *bind.FilterOpts) (*ITokenAddTIterator, error) {

	logs, sub, err := _IToken.contract.FilterLogs(opts, "AddT")
	if err != nil {
		return nil, err
	}
	return &ITokenAddTIterator{contract: _IToken.contract, event: "AddT", logs: logs, sub: sub}, nil
}

// WatchAddT is a free log subscription operation binding the contract event 0xf10191767d4ea3fc26b057d307336c7e8df8880bb07ddaebe4e853db6b5fd936.
//
// Solidity: event AddT(address t, uint8 tIndex)
func (_IToken *ITokenFilterer) WatchAddT(opts *bind.WatchOpts, sink chan<- *ITokenAddT) (event.Subscription, error) {

	logs, sub, err := _IToken.contract.WatchLogs(opts, "AddT")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITokenAddT)
				if err := _IToken.contract.UnpackLog(event, "AddT", log); err != nil {
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

// ParseAddT is a log parse operation binding the contract event 0xf10191767d4ea3fc26b057d307336c7e8df8880bb07ddaebe4e853db6b5fd936.
//
// Solidity: event AddT(address t, uint8 tIndex)
func (_IToken *ITokenFilterer) ParseAddT(log types.Log) (*ITokenAddT, error) {
	event := new(ITokenAddT)
	if err := _IToken.contract.UnpackLog(event, "AddT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITokenBanTIterator is returned from FilterBanT and is used to iterate over the raw logs and unpacked data for BanT events raised by the IToken contract.
type ITokenBanTIterator struct {
	Event *ITokenBanT // Event containing the contract specifics and raw log

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
func (it *ITokenBanTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITokenBanT)
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
		it.Event = new(ITokenBanT)
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
func (it *ITokenBanTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITokenBanTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITokenBanT represents a BanT event raised by the IToken contract.
type ITokenBanT struct {
	T   common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterBanT is a free log retrieval operation binding the contract event 0xdeb67222898c33fcb1fcdd95f5ef10a63c58b43d070d9534ce894fb04bb9b0d8.
//
// Solidity: event BanT(address t)
func (_IToken *ITokenFilterer) FilterBanT(opts *bind.FilterOpts) (*ITokenBanTIterator, error) {

	logs, sub, err := _IToken.contract.FilterLogs(opts, "BanT")
	if err != nil {
		return nil, err
	}
	return &ITokenBanTIterator{contract: _IToken.contract, event: "BanT", logs: logs, sub: sub}, nil
}

// WatchBanT is a free log subscription operation binding the contract event 0xdeb67222898c33fcb1fcdd95f5ef10a63c58b43d070d9534ce894fb04bb9b0d8.
//
// Solidity: event BanT(address t)
func (_IToken *ITokenFilterer) WatchBanT(opts *bind.WatchOpts, sink chan<- *ITokenBanT) (event.Subscription, error) {

	logs, sub, err := _IToken.contract.WatchLogs(opts, "BanT")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITokenBanT)
				if err := _IToken.contract.UnpackLog(event, "BanT", log); err != nil {
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

// ParseBanT is a log parse operation binding the contract event 0xdeb67222898c33fcb1fcdd95f5ef10a63c58b43d070d9534ce894fb04bb9b0d8.
//
// Solidity: event BanT(address t)
func (_IToken *ITokenFilterer) ParseBanT(log types.Log) (*ITokenBanT, error) {
	event := new(ITokenBanT)
	if err := _IToken.contract.UnpackLog(event, "BanT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITokenGetterABI is the input ABI used to generate the binding from.
const ITokenGetterABI = "[{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"}],\"name\":\"checkT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"}],\"name\":\"getTA\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTCnt\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_t\",\"type\":\"address\"}],\"name\":\"getTI\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ITokenGetterFuncSigs maps the 4-byte function signature to its string representation.
var ITokenGetterFuncSigs = map[string]string{
	"81abb8fe": "checkT(uint8)",
	"8bb4a637": "getTA(uint8)",
	"7600b86a": "getTCnt()",
	"2df2685f": "getTI(address)",
}

// ITokenGetter is an auto generated Go binding around an Ethereum contract.
type ITokenGetter struct {
	ITokenGetterCaller     // Read-only binding to the contract
	ITokenGetterTransactor // Write-only binding to the contract
	ITokenGetterFilterer   // Log filterer for contract events
}

// ITokenGetterCaller is an auto generated read-only Go binding around an Ethereum contract.
type ITokenGetterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITokenGetterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ITokenGetterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITokenGetterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ITokenGetterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITokenGetterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ITokenGetterSession struct {
	Contract     *ITokenGetter     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ITokenGetterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ITokenGetterCallerSession struct {
	Contract *ITokenGetterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ITokenGetterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ITokenGetterTransactorSession struct {
	Contract     *ITokenGetterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ITokenGetterRaw is an auto generated low-level Go binding around an Ethereum contract.
type ITokenGetterRaw struct {
	Contract *ITokenGetter // Generic contract binding to access the raw methods on
}

// ITokenGetterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ITokenGetterCallerRaw struct {
	Contract *ITokenGetterCaller // Generic read-only contract binding to access the raw methods on
}

// ITokenGetterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ITokenGetterTransactorRaw struct {
	Contract *ITokenGetterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewITokenGetter creates a new instance of ITokenGetter, bound to a specific deployed contract.
func NewITokenGetter(address common.Address, backend bind.ContractBackend) (*ITokenGetter, error) {
	contract, err := bindITokenGetter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ITokenGetter{ITokenGetterCaller: ITokenGetterCaller{contract: contract}, ITokenGetterTransactor: ITokenGetterTransactor{contract: contract}, ITokenGetterFilterer: ITokenGetterFilterer{contract: contract}}, nil
}

// NewITokenGetterCaller creates a new read-only instance of ITokenGetter, bound to a specific deployed contract.
func NewITokenGetterCaller(address common.Address, caller bind.ContractCaller) (*ITokenGetterCaller, error) {
	contract, err := bindITokenGetter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ITokenGetterCaller{contract: contract}, nil
}

// NewITokenGetterTransactor creates a new write-only instance of ITokenGetter, bound to a specific deployed contract.
func NewITokenGetterTransactor(address common.Address, transactor bind.ContractTransactor) (*ITokenGetterTransactor, error) {
	contract, err := bindITokenGetter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ITokenGetterTransactor{contract: contract}, nil
}

// NewITokenGetterFilterer creates a new log filterer instance of ITokenGetter, bound to a specific deployed contract.
func NewITokenGetterFilterer(address common.Address, filterer bind.ContractFilterer) (*ITokenGetterFilterer, error) {
	contract, err := bindITokenGetter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ITokenGetterFilterer{contract: contract}, nil
}

// bindITokenGetter binds a generic wrapper to an already deployed contract.
func bindITokenGetter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ITokenGetterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITokenGetter *ITokenGetterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITokenGetter.Contract.ITokenGetterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITokenGetter *ITokenGetterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITokenGetter.Contract.ITokenGetterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITokenGetter *ITokenGetterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITokenGetter.Contract.ITokenGetterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITokenGetter *ITokenGetterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITokenGetter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITokenGetter *ITokenGetterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITokenGetter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITokenGetter *ITokenGetterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITokenGetter.Contract.contract.Transact(opts, method, params...)
}

// CheckT is a free data retrieval call binding the contract method 0x81abb8fe.
//
// Solidity: function checkT(uint8 _ti) view returns(address, bool)
func (_ITokenGetter *ITokenGetterCaller) CheckT(opts *bind.CallOpts, _ti uint8) (common.Address, bool, error) {
	var out []interface{}
	err := _ITokenGetter.contract.Call(opts, &out, "checkT", _ti)

	if err != nil {
		return *new(common.Address), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// CheckT is a free data retrieval call binding the contract method 0x81abb8fe.
//
// Solidity: function checkT(uint8 _ti) view returns(address, bool)
func (_ITokenGetter *ITokenGetterSession) CheckT(_ti uint8) (common.Address, bool, error) {
	return _ITokenGetter.Contract.CheckT(&_ITokenGetter.CallOpts, _ti)
}

// CheckT is a free data retrieval call binding the contract method 0x81abb8fe.
//
// Solidity: function checkT(uint8 _ti) view returns(address, bool)
func (_ITokenGetter *ITokenGetterCallerSession) CheckT(_ti uint8) (common.Address, bool, error) {
	return _ITokenGetter.Contract.CheckT(&_ITokenGetter.CallOpts, _ti)
}

// GetTA is a free data retrieval call binding the contract method 0x8bb4a637.
//
// Solidity: function getTA(uint8 _ti) view returns(address)
func (_ITokenGetter *ITokenGetterCaller) GetTA(opts *bind.CallOpts, _ti uint8) (common.Address, error) {
	var out []interface{}
	err := _ITokenGetter.contract.Call(opts, &out, "getTA", _ti)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTA is a free data retrieval call binding the contract method 0x8bb4a637.
//
// Solidity: function getTA(uint8 _ti) view returns(address)
func (_ITokenGetter *ITokenGetterSession) GetTA(_ti uint8) (common.Address, error) {
	return _ITokenGetter.Contract.GetTA(&_ITokenGetter.CallOpts, _ti)
}

// GetTA is a free data retrieval call binding the contract method 0x8bb4a637.
//
// Solidity: function getTA(uint8 _ti) view returns(address)
func (_ITokenGetter *ITokenGetterCallerSession) GetTA(_ti uint8) (common.Address, error) {
	return _ITokenGetter.Contract.GetTA(&_ITokenGetter.CallOpts, _ti)
}

// GetTCnt is a free data retrieval call binding the contract method 0x7600b86a.
//
// Solidity: function getTCnt() view returns(uint8)
func (_ITokenGetter *ITokenGetterCaller) GetTCnt(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ITokenGetter.contract.Call(opts, &out, "getTCnt")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetTCnt is a free data retrieval call binding the contract method 0x7600b86a.
//
// Solidity: function getTCnt() view returns(uint8)
func (_ITokenGetter *ITokenGetterSession) GetTCnt() (uint8, error) {
	return _ITokenGetter.Contract.GetTCnt(&_ITokenGetter.CallOpts)
}

// GetTCnt is a free data retrieval call binding the contract method 0x7600b86a.
//
// Solidity: function getTCnt() view returns(uint8)
func (_ITokenGetter *ITokenGetterCallerSession) GetTCnt() (uint8, error) {
	return _ITokenGetter.Contract.GetTCnt(&_ITokenGetter.CallOpts)
}

// GetTI is a free data retrieval call binding the contract method 0x2df2685f.
//
// Solidity: function getTI(address _t) view returns(uint8, bool, bool)
func (_ITokenGetter *ITokenGetterCaller) GetTI(opts *bind.CallOpts, _t common.Address) (uint8, bool, bool, error) {
	var out []interface{}
	err := _ITokenGetter.contract.Call(opts, &out, "getTI", _t)

	if err != nil {
		return *new(uint8), *new(bool), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)
	out2 := *abi.ConvertType(out[2], new(bool)).(*bool)

	return out0, out1, out2, err

}

// GetTI is a free data retrieval call binding the contract method 0x2df2685f.
//
// Solidity: function getTI(address _t) view returns(uint8, bool, bool)
func (_ITokenGetter *ITokenGetterSession) GetTI(_t common.Address) (uint8, bool, bool, error) {
	return _ITokenGetter.Contract.GetTI(&_ITokenGetter.CallOpts, _t)
}

// GetTI is a free data retrieval call binding the contract method 0x2df2685f.
//
// Solidity: function getTI(address _t) view returns(uint8, bool, bool)
func (_ITokenGetter *ITokenGetterCallerSession) GetTI(_t common.Address) (uint8, bool, bool, error) {
	return _ITokenGetter.Contract.GetTI(&_ITokenGetter.CallOpts, _t)
}

// ITokenSetterABI is the input ABI used to generate the binding from.
const ITokenSetterABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"tIndex\",\"type\":\"uint8\"}],\"name\":\"AddT\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"}],\"name\":\"BanT\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_t\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_mint\",\"type\":\"bool\"}],\"name\":\"addT\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_t\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isBan\",\"type\":\"bool\"}],\"name\":\"banT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ITokenSetterFuncSigs maps the 4-byte function signature to its string representation.
var ITokenSetterFuncSigs = map[string]string{
	"ff4268a4": "addT(address,bool)",
	"83373b36": "banT(address,bool)",
}

// ITokenSetter is an auto generated Go binding around an Ethereum contract.
type ITokenSetter struct {
	ITokenSetterCaller     // Read-only binding to the contract
	ITokenSetterTransactor // Write-only binding to the contract
	ITokenSetterFilterer   // Log filterer for contract events
}

// ITokenSetterCaller is an auto generated read-only Go binding around an Ethereum contract.
type ITokenSetterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITokenSetterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ITokenSetterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITokenSetterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ITokenSetterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITokenSetterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ITokenSetterSession struct {
	Contract     *ITokenSetter     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ITokenSetterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ITokenSetterCallerSession struct {
	Contract *ITokenSetterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ITokenSetterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ITokenSetterTransactorSession struct {
	Contract     *ITokenSetterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ITokenSetterRaw is an auto generated low-level Go binding around an Ethereum contract.
type ITokenSetterRaw struct {
	Contract *ITokenSetter // Generic contract binding to access the raw methods on
}

// ITokenSetterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ITokenSetterCallerRaw struct {
	Contract *ITokenSetterCaller // Generic read-only contract binding to access the raw methods on
}

// ITokenSetterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ITokenSetterTransactorRaw struct {
	Contract *ITokenSetterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewITokenSetter creates a new instance of ITokenSetter, bound to a specific deployed contract.
func NewITokenSetter(address common.Address, backend bind.ContractBackend) (*ITokenSetter, error) {
	contract, err := bindITokenSetter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ITokenSetter{ITokenSetterCaller: ITokenSetterCaller{contract: contract}, ITokenSetterTransactor: ITokenSetterTransactor{contract: contract}, ITokenSetterFilterer: ITokenSetterFilterer{contract: contract}}, nil
}

// NewITokenSetterCaller creates a new read-only instance of ITokenSetter, bound to a specific deployed contract.
func NewITokenSetterCaller(address common.Address, caller bind.ContractCaller) (*ITokenSetterCaller, error) {
	contract, err := bindITokenSetter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ITokenSetterCaller{contract: contract}, nil
}

// NewITokenSetterTransactor creates a new write-only instance of ITokenSetter, bound to a specific deployed contract.
func NewITokenSetterTransactor(address common.Address, transactor bind.ContractTransactor) (*ITokenSetterTransactor, error) {
	contract, err := bindITokenSetter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ITokenSetterTransactor{contract: contract}, nil
}

// NewITokenSetterFilterer creates a new log filterer instance of ITokenSetter, bound to a specific deployed contract.
func NewITokenSetterFilterer(address common.Address, filterer bind.ContractFilterer) (*ITokenSetterFilterer, error) {
	contract, err := bindITokenSetter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ITokenSetterFilterer{contract: contract}, nil
}

// bindITokenSetter binds a generic wrapper to an already deployed contract.
func bindITokenSetter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ITokenSetterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITokenSetter *ITokenSetterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITokenSetter.Contract.ITokenSetterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITokenSetter *ITokenSetterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITokenSetter.Contract.ITokenSetterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITokenSetter *ITokenSetterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITokenSetter.Contract.ITokenSetterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITokenSetter *ITokenSetterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITokenSetter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITokenSetter *ITokenSetterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITokenSetter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITokenSetter *ITokenSetterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITokenSetter.Contract.contract.Transact(opts, method, params...)
}

// AddT is a paid mutator transaction binding the contract method 0xff4268a4.
//
// Solidity: function addT(address _t, bool _mint) returns(uint8)
func (_ITokenSetter *ITokenSetterTransactor) AddT(opts *bind.TransactOpts, _t common.Address, _mint bool) (*types.Transaction, error) {
	return _ITokenSetter.contract.Transact(opts, "addT", _t, _mint)
}

// AddT is a paid mutator transaction binding the contract method 0xff4268a4.
//
// Solidity: function addT(address _t, bool _mint) returns(uint8)
func (_ITokenSetter *ITokenSetterSession) AddT(_t common.Address, _mint bool) (*types.Transaction, error) {
	return _ITokenSetter.Contract.AddT(&_ITokenSetter.TransactOpts, _t, _mint)
}

// AddT is a paid mutator transaction binding the contract method 0xff4268a4.
//
// Solidity: function addT(address _t, bool _mint) returns(uint8)
func (_ITokenSetter *ITokenSetterTransactorSession) AddT(_t common.Address, _mint bool) (*types.Transaction, error) {
	return _ITokenSetter.Contract.AddT(&_ITokenSetter.TransactOpts, _t, _mint)
}

// BanT is a paid mutator transaction binding the contract method 0x83373b36.
//
// Solidity: function banT(address _t, bool _isBan) returns()
func (_ITokenSetter *ITokenSetterTransactor) BanT(opts *bind.TransactOpts, _t common.Address, _isBan bool) (*types.Transaction, error) {
	return _ITokenSetter.contract.Transact(opts, "banT", _t, _isBan)
}

// BanT is a paid mutator transaction binding the contract method 0x83373b36.
//
// Solidity: function banT(address _t, bool _isBan) returns()
func (_ITokenSetter *ITokenSetterSession) BanT(_t common.Address, _isBan bool) (*types.Transaction, error) {
	return _ITokenSetter.Contract.BanT(&_ITokenSetter.TransactOpts, _t, _isBan)
}

// BanT is a paid mutator transaction binding the contract method 0x83373b36.
//
// Solidity: function banT(address _t, bool _isBan) returns()
func (_ITokenSetter *ITokenSetterTransactorSession) BanT(_t common.Address, _isBan bool) (*types.Transaction, error) {
	return _ITokenSetter.Contract.BanT(&_ITokenSetter.TransactOpts, _t, _isBan)
}

// ITokenSetterAddTIterator is returned from FilterAddT and is used to iterate over the raw logs and unpacked data for AddT events raised by the ITokenSetter contract.
type ITokenSetterAddTIterator struct {
	Event *ITokenSetterAddT // Event containing the contract specifics and raw log

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
func (it *ITokenSetterAddTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITokenSetterAddT)
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
		it.Event = new(ITokenSetterAddT)
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
func (it *ITokenSetterAddTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITokenSetterAddTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITokenSetterAddT represents a AddT event raised by the ITokenSetter contract.
type ITokenSetterAddT struct {
	T      common.Address
	TIndex uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAddT is a free log retrieval operation binding the contract event 0xf10191767d4ea3fc26b057d307336c7e8df8880bb07ddaebe4e853db6b5fd936.
//
// Solidity: event AddT(address t, uint8 tIndex)
func (_ITokenSetter *ITokenSetterFilterer) FilterAddT(opts *bind.FilterOpts) (*ITokenSetterAddTIterator, error) {

	logs, sub, err := _ITokenSetter.contract.FilterLogs(opts, "AddT")
	if err != nil {
		return nil, err
	}
	return &ITokenSetterAddTIterator{contract: _ITokenSetter.contract, event: "AddT", logs: logs, sub: sub}, nil
}

// WatchAddT is a free log subscription operation binding the contract event 0xf10191767d4ea3fc26b057d307336c7e8df8880bb07ddaebe4e853db6b5fd936.
//
// Solidity: event AddT(address t, uint8 tIndex)
func (_ITokenSetter *ITokenSetterFilterer) WatchAddT(opts *bind.WatchOpts, sink chan<- *ITokenSetterAddT) (event.Subscription, error) {

	logs, sub, err := _ITokenSetter.contract.WatchLogs(opts, "AddT")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITokenSetterAddT)
				if err := _ITokenSetter.contract.UnpackLog(event, "AddT", log); err != nil {
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

// ParseAddT is a log parse operation binding the contract event 0xf10191767d4ea3fc26b057d307336c7e8df8880bb07ddaebe4e853db6b5fd936.
//
// Solidity: event AddT(address t, uint8 tIndex)
func (_ITokenSetter *ITokenSetterFilterer) ParseAddT(log types.Log) (*ITokenSetterAddT, error) {
	event := new(ITokenSetterAddT)
	if err := _ITokenSetter.contract.UnpackLog(event, "AddT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITokenSetterBanTIterator is returned from FilterBanT and is used to iterate over the raw logs and unpacked data for BanT events raised by the ITokenSetter contract.
type ITokenSetterBanTIterator struct {
	Event *ITokenSetterBanT // Event containing the contract specifics and raw log

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
func (it *ITokenSetterBanTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITokenSetterBanT)
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
		it.Event = new(ITokenSetterBanT)
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
func (it *ITokenSetterBanTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITokenSetterBanTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITokenSetterBanT represents a BanT event raised by the ITokenSetter contract.
type ITokenSetterBanT struct {
	T   common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterBanT is a free log retrieval operation binding the contract event 0xdeb67222898c33fcb1fcdd95f5ef10a63c58b43d070d9534ce894fb04bb9b0d8.
//
// Solidity: event BanT(address t)
func (_ITokenSetter *ITokenSetterFilterer) FilterBanT(opts *bind.FilterOpts) (*ITokenSetterBanTIterator, error) {

	logs, sub, err := _ITokenSetter.contract.FilterLogs(opts, "BanT")
	if err != nil {
		return nil, err
	}
	return &ITokenSetterBanTIterator{contract: _ITokenSetter.contract, event: "BanT", logs: logs, sub: sub}, nil
}

// WatchBanT is a free log subscription operation binding the contract event 0xdeb67222898c33fcb1fcdd95f5ef10a63c58b43d070d9534ce894fb04bb9b0d8.
//
// Solidity: event BanT(address t)
func (_ITokenSetter *ITokenSetterFilterer) WatchBanT(opts *bind.WatchOpts, sink chan<- *ITokenSetterBanT) (event.Subscription, error) {

	logs, sub, err := _ITokenSetter.contract.WatchLogs(opts, "BanT")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITokenSetterBanT)
				if err := _ITokenSetter.contract.UnpackLog(event, "BanT", log); err != nil {
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

// ParseBanT is a log parse operation binding the contract event 0xdeb67222898c33fcb1fcdd95f5ef10a63c58b43d070d9534ce894fb04bb9b0d8.
//
// Solidity: event BanT(address t)
func (_ITokenSetter *ITokenSetterFilterer) ParseBanT(log types.Log) (*ITokenSetterBanT, error) {
	event := new(ITokenSetterBanT)
	if err := _ITokenSetter.contract.UnpackLog(event, "BanT", log); err != nil {
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

// TokenABI is the input ABI used to generate the binding from.
const TokenABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isSet\",\"type\":\"bool\"}],\"name\":\"AddOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"tIndex\",\"type\":\"uint8\"}],\"name\":\"AddT\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"}],\"name\":\"BanT\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isSet\",\"type\":\"bool\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_t\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_mint\",\"type\":\"bool\"}],\"name\":\"addT\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"auth\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_t\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_ban\",\"type\":\"bool\"}],\"name\":\"banT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"tIndex\",\"type\":\"uint8\"}],\"name\":\"checkT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"tIndex\",\"type\":\"uint8\"}],\"name\":\"getTA\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTCnt\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"}],\"name\":\"getTI\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"owners\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// TokenFuncSigs maps the 4-byte function signature to its string representation.
var TokenFuncSigs = map[string]string{
	"4bf1b134": "add(address,bool,bytes[5])",
	"ff4268a4": "addT(address,bool)",
	"de9375f2": "auth()",
	"83373b36": "banT(address,bool)",
	"81abb8fe": "checkT(uint8)",
	"8bb4a637": "getTA(uint8)",
	"7600b86a": "getTCnt()",
	"2df2685f": "getTI(address)",
	"022914a7": "owners(address)",
	"54fd4d50": "version()",
}

// TokenBin is the compiled bytecode used for deploying new contracts.
var TokenBin = "0x60806040526001805461ffff60a01b1916600160a11b17905534801561002457600080fd5b50604051610b6a380380610b6a83398101604081905261004391610068565b600180546001600160a01b0319166001600160a01b0392909216919091179055610098565b60006020828403121561007a57600080fd5b81516001600160a01b038116811461009157600080fd5b9392505050565b610ac3806100a76000396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c806381abb8fe1161006657806381abb8fe1461016257806383373b36146101945780638bb4a637146101a7578063de9375f2146101d2578063ff4268a4146101e557600080fd5b8063022914a7146100a35780632df2685f146100db5780634bf1b1341461010f57806354fd4d50146101245780637600b86a1461014c575b600080fd5b6100c66100b13660046107e0565b60006020819052908152604090205460ff1681565b60405190151581526020015b60405180910390f35b6100ee6100e93660046107e0565b6101f8565b6040805160ff909416845291151560208401521515908201526060016100d2565b61012261011d366004610882565b610294565b005b60015461013990600160a01b900461ffff1681565b60405161ffff90911681526020016100d2565b6002545b60405160ff90911681526020016100d2565b610175610170366004610991565b610412565b604080516001600160a01b0390931683529015156020830152016100d2565b6101226101a23660046109b4565b61051a565b6101ba6101b5366004610991565b6105c6565b6040516001600160a01b0390911681526020016100d2565b6001546101ba906001600160a01b031681565b6101506101f33660046109b4565b6105f9565b6001600160a01b03811660009081526003602052604081205460028054839283926201000090910460ff169182908110610234576102346109e7565b6000918252602090912001546001600160a01b0390811690861603610282576001600160a01b03851660009081526003602052604090205490935060ff61010082048116935016905061028d565b925060019150600090505b9193909250565b823b60008190036102e15760405162461bcd60e51b81526020600482015260126024820152713732b2b21031b7b73a3930b1ba1030b2323960711b60448201526064015b60405180910390fd5b6040516bffffffffffffffffffffffff1930606090811b821660208401526218591960ea1b603484015286901b16603782015283151560f81b604b820152600090604c0160408051601f1981840301815290829052805160209091012060015463a96bba9d60e01b83529092506001600160a01b03169063a96bba9d9061036e90849087906004016109fd565b600060405180830381600087803b15801561038857600080fd5b505af115801561039c573d6000803e3d6000fd5b5050604080516001600160a01b038916815287151560208201527f938b2a24c98e4e542127ffa74a91e48942c2bddccae3b6d75f82bfda7bbc0807935001905060405180910390a15050506001600160a01b03919091166000908152602081905260409020805460ff1916911515919091179055565b6000806003600060028560ff168154811061042f5761042f6109e7565b60009182526020808320909101546001600160a01b0316835282019290925260400190205460ff61010090910416156104995760405162461bcd60e51b815260206004820152600c60248201526b1d1bdad95b8818985b9b995960a21b60448201526064016102d8565b60028360ff16815481106104af576104af6109e7565b9060005260206000200160009054906101000a90046001600160a01b03166003600060028660ff16815481106104e7576104e76109e7565b60009182526020808320909101546001600160a01b03168352820192909252604001902054909460ff9091169350915050565b3360009081526020819052604090205460ff166105655760405162461bcd60e51b81526020600482015260096024820152683737ba1037bbb732b960b91b60448201526064016102d8565b6001600160a01b038216600081815260036020908152604091829020805461ff0019166101008615150217905590519182527fdeb67222898c33fcb1fcdd95f5ef10a63c58b43d070d9534ce894fb04bb9b0d8910160405180910390a15050565b600060028260ff16815481106105de576105de6109e7565b6000918252602090912001546001600160a01b031692915050565b3360009081526020819052604081205460ff166106445760405162461bcd60e51b81526020600482015260096024820152683737ba1037bbb732b960b91b60448201526064016102d8565b61064d8361073f565b156106865760405162461bcd60e51b81526020600482015260096024820152683430b9903a37b5b2b760b91b60448201526064016102d8565b60028054600181019091557f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace810180546001600160a01b0319166001600160a01b038616908117909155600081815260036020908152604091829020805462ff00ff19166201000060ff871690810260ff191691909117881515179091558251938452908301527ff10191767d4ea3fc26b057d307336c7e8df8880bb07ddaebe4e853db6b5fd936910160405180910390a19392505050565b6001600160a01b03811660009081526003602052604081205462010000900460ff161561076e57506001919050565b600254158015906107af5750816001600160a01b03166002600081548110610798576107986109e7565b6000918252602090912001546001600160a01b0316145b156107bc57506001919050565b506000919050565b80356001600160a01b03811681146107db57600080fd5b919050565b6000602082840312156107f257600080fd5b6107fb826107c4565b9392505050565b803580151581146107db57600080fd5b634e487b7160e01b600052604160045260246000fd5b60405160a0810167ffffffffffffffff8111828210171561084b5761084b610812565b60405290565b604051601f8201601f1916810167ffffffffffffffff8111828210171561087a5761087a610812565b604052919050565b60008060006060848603121561089757600080fd5b6108a0846107c4565b925060206108af818601610802565b9250604085013567ffffffffffffffff808211156108cc57600080fd5b8187019150601f88818401126108e157600080fd5b6108e9610828565b8060a085018b8111156108fb57600080fd5b855b8181101561097f578035868111156109155760008081fd5b87018581018e136109265760008081fd5b80358781111561093857610938610812565b610949818801601f19168b01610851565b8181528f8b83850101111561095e5760008081fd5b818b84018c83013760009181018b01919091528552509287019287016108fd565b50508096505050505050509250925092565b6000602082840312156109a357600080fd5b813560ff811681146107fb57600080fd5b600080604083850312156109c757600080fd5b6109d0836107c4565b91506109de60208401610802565b90509250929050565b634e487b7160e01b600052603260045260246000fd5b8281526040602080830182905260009160e08401919084018584805b6005811015610a7f57878603603f1901845282518051808852835b81811015610a4f578281018801518982018901528701610a34565b81811115610a5f578488838b0101525b50601f01601f191696909601850195509284019291840191600101610a19565b50939897505050505050505056fea2646970667358221220678e026bd4d71226544b8fc0c31d0487e5bb5e9570e9f002beb0619fa7f0959664736f6c634300080e0033"

// DeployToken deploys a new Ethereum contract, binding an instance of Token to it.
func DeployToken(auth *bind.TransactOpts, backend bind.ContractBackend, _a common.Address) (common.Address, *types.Transaction, *Token, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TokenBin), backend, _a)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}, TokenFilterer: TokenFilterer{contract: contract}}, nil
}

// Token is an auto generated Go binding around an Ethereum contract.
type Token struct {
	TokenCaller     // Read-only binding to the contract
	TokenTransactor // Write-only binding to the contract
	TokenFilterer   // Log filterer for contract events
}

// TokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenSession struct {
	Contract     *Token            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenCallerSession struct {
	Contract *TokenCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenTransactorSession struct {
	Contract     *TokenTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenRaw struct {
	Contract *Token // Generic contract binding to access the raw methods on
}

// TokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenCallerRaw struct {
	Contract *TokenCaller // Generic read-only contract binding to access the raw methods on
}

// TokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenTransactorRaw struct {
	Contract *TokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewToken creates a new instance of Token, bound to a specific deployed contract.
func NewToken(address common.Address, backend bind.ContractBackend) (*Token, error) {
	contract, err := bindToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}, TokenFilterer: TokenFilterer{contract: contract}}, nil
}

// NewTokenCaller creates a new read-only instance of Token, bound to a specific deployed contract.
func NewTokenCaller(address common.Address, caller bind.ContractCaller) (*TokenCaller, error) {
	contract, err := bindToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenCaller{contract: contract}, nil
}

// NewTokenTransactor creates a new write-only instance of Token, bound to a specific deployed contract.
func NewTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenTransactor, error) {
	contract, err := bindToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenTransactor{contract: contract}, nil
}

// NewTokenFilterer creates a new log filterer instance of Token, bound to a specific deployed contract.
func NewTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenFilterer, error) {
	contract, err := bindToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenFilterer{contract: contract}, nil
}

// bindToken binds a generic wrapper to an already deployed contract.
func bindToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Token.Contract.TokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Token.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.contract.Transact(opts, method, params...)
}

// Auth is a free data retrieval call binding the contract method 0xde9375f2.
//
// Solidity: function auth() view returns(address)
func (_Token *TokenCaller) Auth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "auth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Auth is a free data retrieval call binding the contract method 0xde9375f2.
//
// Solidity: function auth() view returns(address)
func (_Token *TokenSession) Auth() (common.Address, error) {
	return _Token.Contract.Auth(&_Token.CallOpts)
}

// Auth is a free data retrieval call binding the contract method 0xde9375f2.
//
// Solidity: function auth() view returns(address)
func (_Token *TokenCallerSession) Auth() (common.Address, error) {
	return _Token.Contract.Auth(&_Token.CallOpts)
}

// CheckT is a free data retrieval call binding the contract method 0x81abb8fe.
//
// Solidity: function checkT(uint8 tIndex) view returns(address, bool)
func (_Token *TokenCaller) CheckT(opts *bind.CallOpts, tIndex uint8) (common.Address, bool, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "checkT", tIndex)

	if err != nil {
		return *new(common.Address), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// CheckT is a free data retrieval call binding the contract method 0x81abb8fe.
//
// Solidity: function checkT(uint8 tIndex) view returns(address, bool)
func (_Token *TokenSession) CheckT(tIndex uint8) (common.Address, bool, error) {
	return _Token.Contract.CheckT(&_Token.CallOpts, tIndex)
}

// CheckT is a free data retrieval call binding the contract method 0x81abb8fe.
//
// Solidity: function checkT(uint8 tIndex) view returns(address, bool)
func (_Token *TokenCallerSession) CheckT(tIndex uint8) (common.Address, bool, error) {
	return _Token.Contract.CheckT(&_Token.CallOpts, tIndex)
}

// GetTA is a free data retrieval call binding the contract method 0x8bb4a637.
//
// Solidity: function getTA(uint8 tIndex) view returns(address)
func (_Token *TokenCaller) GetTA(opts *bind.CallOpts, tIndex uint8) (common.Address, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "getTA", tIndex)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTA is a free data retrieval call binding the contract method 0x8bb4a637.
//
// Solidity: function getTA(uint8 tIndex) view returns(address)
func (_Token *TokenSession) GetTA(tIndex uint8) (common.Address, error) {
	return _Token.Contract.GetTA(&_Token.CallOpts, tIndex)
}

// GetTA is a free data retrieval call binding the contract method 0x8bb4a637.
//
// Solidity: function getTA(uint8 tIndex) view returns(address)
func (_Token *TokenCallerSession) GetTA(tIndex uint8) (common.Address, error) {
	return _Token.Contract.GetTA(&_Token.CallOpts, tIndex)
}

// GetTCnt is a free data retrieval call binding the contract method 0x7600b86a.
//
// Solidity: function getTCnt() view returns(uint8)
func (_Token *TokenCaller) GetTCnt(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "getTCnt")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetTCnt is a free data retrieval call binding the contract method 0x7600b86a.
//
// Solidity: function getTCnt() view returns(uint8)
func (_Token *TokenSession) GetTCnt() (uint8, error) {
	return _Token.Contract.GetTCnt(&_Token.CallOpts)
}

// GetTCnt is a free data retrieval call binding the contract method 0x7600b86a.
//
// Solidity: function getTCnt() view returns(uint8)
func (_Token *TokenCallerSession) GetTCnt() (uint8, error) {
	return _Token.Contract.GetTCnt(&_Token.CallOpts)
}

// GetTI is a free data retrieval call binding the contract method 0x2df2685f.
//
// Solidity: function getTI(address t) view returns(uint8, bool, bool)
func (_Token *TokenCaller) GetTI(opts *bind.CallOpts, t common.Address) (uint8, bool, bool, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "getTI", t)

	if err != nil {
		return *new(uint8), *new(bool), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)
	out2 := *abi.ConvertType(out[2], new(bool)).(*bool)

	return out0, out1, out2, err

}

// GetTI is a free data retrieval call binding the contract method 0x2df2685f.
//
// Solidity: function getTI(address t) view returns(uint8, bool, bool)
func (_Token *TokenSession) GetTI(t common.Address) (uint8, bool, bool, error) {
	return _Token.Contract.GetTI(&_Token.CallOpts, t)
}

// GetTI is a free data retrieval call binding the contract method 0x2df2685f.
//
// Solidity: function getTI(address t) view returns(uint8, bool, bool)
func (_Token *TokenCallerSession) GetTI(t common.Address) (uint8, bool, bool, error) {
	return _Token.Contract.GetTI(&_Token.CallOpts, t)
}

// Owners is a free data retrieval call binding the contract method 0x022914a7.
//
// Solidity: function owners(address ) view returns(bool)
func (_Token *TokenCaller) Owners(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "owners", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Owners is a free data retrieval call binding the contract method 0x022914a7.
//
// Solidity: function owners(address ) view returns(bool)
func (_Token *TokenSession) Owners(arg0 common.Address) (bool, error) {
	return _Token.Contract.Owners(&_Token.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x022914a7.
//
// Solidity: function owners(address ) view returns(bool)
func (_Token *TokenCallerSession) Owners(arg0 common.Address) (bool, error) {
	return _Token.Contract.Owners(&_Token.CallOpts, arg0)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint16)
func (_Token *TokenCaller) Version(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint16)
func (_Token *TokenSession) Version() (uint16, error) {
	return _Token.Contract.Version(&_Token.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint16)
func (_Token *TokenCallerSession) Version() (uint16, error) {
	return _Token.Contract.Version(&_Token.CallOpts)
}

// Add is a paid mutator transaction binding the contract method 0x4bf1b134.
//
// Solidity: function add(address _a, bool isSet, bytes[5] signs) returns()
func (_Token *TokenTransactor) Add(opts *bind.TransactOpts, _a common.Address, isSet bool, signs [5][]byte) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "add", _a, isSet, signs)
}

// Add is a paid mutator transaction binding the contract method 0x4bf1b134.
//
// Solidity: function add(address _a, bool isSet, bytes[5] signs) returns()
func (_Token *TokenSession) Add(_a common.Address, isSet bool, signs [5][]byte) (*types.Transaction, error) {
	return _Token.Contract.Add(&_Token.TransactOpts, _a, isSet, signs)
}

// Add is a paid mutator transaction binding the contract method 0x4bf1b134.
//
// Solidity: function add(address _a, bool isSet, bytes[5] signs) returns()
func (_Token *TokenTransactorSession) Add(_a common.Address, isSet bool, signs [5][]byte) (*types.Transaction, error) {
	return _Token.Contract.Add(&_Token.TransactOpts, _a, isSet, signs)
}

// AddT is a paid mutator transaction binding the contract method 0xff4268a4.
//
// Solidity: function addT(address _t, bool _mint) returns(uint8)
func (_Token *TokenTransactor) AddT(opts *bind.TransactOpts, _t common.Address, _mint bool) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "addT", _t, _mint)
}

// AddT is a paid mutator transaction binding the contract method 0xff4268a4.
//
// Solidity: function addT(address _t, bool _mint) returns(uint8)
func (_Token *TokenSession) AddT(_t common.Address, _mint bool) (*types.Transaction, error) {
	return _Token.Contract.AddT(&_Token.TransactOpts, _t, _mint)
}

// AddT is a paid mutator transaction binding the contract method 0xff4268a4.
//
// Solidity: function addT(address _t, bool _mint) returns(uint8)
func (_Token *TokenTransactorSession) AddT(_t common.Address, _mint bool) (*types.Transaction, error) {
	return _Token.Contract.AddT(&_Token.TransactOpts, _t, _mint)
}

// BanT is a paid mutator transaction binding the contract method 0x83373b36.
//
// Solidity: function banT(address _t, bool _ban) returns()
func (_Token *TokenTransactor) BanT(opts *bind.TransactOpts, _t common.Address, _ban bool) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "banT", _t, _ban)
}

// BanT is a paid mutator transaction binding the contract method 0x83373b36.
//
// Solidity: function banT(address _t, bool _ban) returns()
func (_Token *TokenSession) BanT(_t common.Address, _ban bool) (*types.Transaction, error) {
	return _Token.Contract.BanT(&_Token.TransactOpts, _t, _ban)
}

// BanT is a paid mutator transaction binding the contract method 0x83373b36.
//
// Solidity: function banT(address _t, bool _ban) returns()
func (_Token *TokenTransactorSession) BanT(_t common.Address, _ban bool) (*types.Transaction, error) {
	return _Token.Contract.BanT(&_Token.TransactOpts, _t, _ban)
}

// TokenAddOwnerIterator is returned from FilterAddOwner and is used to iterate over the raw logs and unpacked data for AddOwner events raised by the Token contract.
type TokenAddOwnerIterator struct {
	Event *TokenAddOwner // Event containing the contract specifics and raw log

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
func (it *TokenAddOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenAddOwner)
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
		it.Event = new(TokenAddOwner)
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
func (it *TokenAddOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenAddOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenAddOwner represents a AddOwner event raised by the Token contract.
type TokenAddOwner struct {
	From  common.Address
	IsSet bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAddOwner is a free log retrieval operation binding the contract event 0x938b2a24c98e4e542127ffa74a91e48942c2bddccae3b6d75f82bfda7bbc0807.
//
// Solidity: event AddOwner(address from, bool isSet)
func (_Token *TokenFilterer) FilterAddOwner(opts *bind.FilterOpts) (*TokenAddOwnerIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "AddOwner")
	if err != nil {
		return nil, err
	}
	return &TokenAddOwnerIterator{contract: _Token.contract, event: "AddOwner", logs: logs, sub: sub}, nil
}

// WatchAddOwner is a free log subscription operation binding the contract event 0x938b2a24c98e4e542127ffa74a91e48942c2bddccae3b6d75f82bfda7bbc0807.
//
// Solidity: event AddOwner(address from, bool isSet)
func (_Token *TokenFilterer) WatchAddOwner(opts *bind.WatchOpts, sink chan<- *TokenAddOwner) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "AddOwner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenAddOwner)
				if err := _Token.contract.UnpackLog(event, "AddOwner", log); err != nil {
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
func (_Token *TokenFilterer) ParseAddOwner(log types.Log) (*TokenAddOwner, error) {
	event := new(TokenAddOwner)
	if err := _Token.contract.UnpackLog(event, "AddOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenAddTIterator is returned from FilterAddT and is used to iterate over the raw logs and unpacked data for AddT events raised by the Token contract.
type TokenAddTIterator struct {
	Event *TokenAddT // Event containing the contract specifics and raw log

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
func (it *TokenAddTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenAddT)
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
		it.Event = new(TokenAddT)
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
func (it *TokenAddTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenAddTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenAddT represents a AddT event raised by the Token contract.
type TokenAddT struct {
	T      common.Address
	TIndex uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAddT is a free log retrieval operation binding the contract event 0xf10191767d4ea3fc26b057d307336c7e8df8880bb07ddaebe4e853db6b5fd936.
//
// Solidity: event AddT(address t, uint8 tIndex)
func (_Token *TokenFilterer) FilterAddT(opts *bind.FilterOpts) (*TokenAddTIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "AddT")
	if err != nil {
		return nil, err
	}
	return &TokenAddTIterator{contract: _Token.contract, event: "AddT", logs: logs, sub: sub}, nil
}

// WatchAddT is a free log subscription operation binding the contract event 0xf10191767d4ea3fc26b057d307336c7e8df8880bb07ddaebe4e853db6b5fd936.
//
// Solidity: event AddT(address t, uint8 tIndex)
func (_Token *TokenFilterer) WatchAddT(opts *bind.WatchOpts, sink chan<- *TokenAddT) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "AddT")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenAddT)
				if err := _Token.contract.UnpackLog(event, "AddT", log); err != nil {
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

// ParseAddT is a log parse operation binding the contract event 0xf10191767d4ea3fc26b057d307336c7e8df8880bb07ddaebe4e853db6b5fd936.
//
// Solidity: event AddT(address t, uint8 tIndex)
func (_Token *TokenFilterer) ParseAddT(log types.Log) (*TokenAddT, error) {
	event := new(TokenAddT)
	if err := _Token.contract.UnpackLog(event, "AddT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenBanTIterator is returned from FilterBanT and is used to iterate over the raw logs and unpacked data for BanT events raised by the Token contract.
type TokenBanTIterator struct {
	Event *TokenBanT // Event containing the contract specifics and raw log

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
func (it *TokenBanTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenBanT)
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
		it.Event = new(TokenBanT)
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
func (it *TokenBanTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenBanTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenBanT represents a BanT event raised by the Token contract.
type TokenBanT struct {
	T   common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterBanT is a free log retrieval operation binding the contract event 0xdeb67222898c33fcb1fcdd95f5ef10a63c58b43d070d9534ce894fb04bb9b0d8.
//
// Solidity: event BanT(address t)
func (_Token *TokenFilterer) FilterBanT(opts *bind.FilterOpts) (*TokenBanTIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "BanT")
	if err != nil {
		return nil, err
	}
	return &TokenBanTIterator{contract: _Token.contract, event: "BanT", logs: logs, sub: sub}, nil
}

// WatchBanT is a free log subscription operation binding the contract event 0xdeb67222898c33fcb1fcdd95f5ef10a63c58b43d070d9534ce894fb04bb9b0d8.
//
// Solidity: event BanT(address t)
func (_Token *TokenFilterer) WatchBanT(opts *bind.WatchOpts, sink chan<- *TokenBanT) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "BanT")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenBanT)
				if err := _Token.contract.UnpackLog(event, "BanT", log); err != nil {
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

// ParseBanT is a log parse operation binding the contract event 0xdeb67222898c33fcb1fcdd95f5ef10a63c58b43d070d9534ce894fb04bb9b0d8.
//
// Solidity: event BanT(address t)
func (_Token *TokenFilterer) ParseBanT(log types.Log) (*TokenBanT, error) {
	event := new(TokenBanT)
	if err := _Token.contract.UnpackLog(event, "BanT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
