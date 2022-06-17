// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package access

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

// AccessABI is the input ABI used to generate the binding from.
const AccessABI = "[{\"inputs\":[{\"internalType\":\"address[5]\",\"name\":\"_addrs\",\"type\":\"address[5]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"access\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"controls\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"floor\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isSet\",\"type\":\"bool\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// AccessFuncSigs maps the 4-byte function signature to its string representation.
var AccessFuncSigs = map[string]string{
	"6fae3d76": "access(address)",
	"da191d88": "controls(uint256)",
	"40695363": "floor()",
	"affed0e0": "nonce()",
	"464b1c45": "set(address,bool,bytes[5])",
	"54fd4d50": "version()",
}

// AccessBin is the compiled bytecode used for deploying new contracts.
var AccessBin = "0x60806040526000805461ffff1916600217905534801561001e57600080fd5b5060405161081438038061081483398101604081905261003d916100da565b61004a6001826005610051565b505061016c565b8260058101928215610099579160200282015b8281111561009957825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190610064565b506100a59291506100a9565b5090565b5b808211156100a557600081556001016100aa565b80516001600160a01b03811681146100d557600080fd5b919050565b600060a082840312156100ec57600080fd5b82601f8301126100fb57600080fd5b60405160a081016001600160401b038111828210171561012b57634e487b7160e01b600052604160045260246000fd5b6040528060a084018581111561014057600080fd5b845b8181101561016157610153816100be565b835260209283019201610142565b509195945050505050565b6106998061017b6000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c80634069536314610067578063464b1c451461008657806354fd4d501461009b5780636fae3d76146100bc578063affed0e0146100ef578063da191d8814610106575b600080fd5b61006f600381565b60405160ff90911681526020015b60405180910390f35b610099610094366004610412565b610131565b005b6000546100a99061ffff1681565b60405161ffff909116815260200161007d565b6100df6100ca36600461052a565b60076020526000908152604090205460ff1681565b604051901515815260200161007d565b6100f860065481565b60405190815260200161007d565b61011961011436600461054e565b61036a565b6040516001600160a01b03909116815260200161007d565b823b60008190036101895760405162461bcd60e51b815260206004820152601d60248201527f6163636573732061646472206e6f7420636f6e7472616374206164647200000060448201526064015b60405180910390fd5b6006546040516bffffffffffffffffffffffff1930606090811b821660208401526034830193909352621cd95d60ea1b60548301529186901b909116605782015283151560f81b606b820152600090606c016040516020818303038152906040528051906020012090506000805b60058110156102dc576001816005811061021357610213610567565b01546001600160a01b031673__$6230a6feddd2d01b0a9ffb5c5e188a1065$__6319045a258588856005811061024b5761024b610567565b60200201516040518363ffffffff1660e01b815260040161026d92919061057d565b602060405180830381865af415801561028a573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102ae91906105da565b6001600160a01b0316036102ca576102c760018361060d565b91505b806102d481610632565b9150506101f7565b50600360ff821610156103265760405162461bcd60e51b815260206004820152601260248201527106a40e6d2cedce640dcdee840cadcdeeaced60731b6044820152606401610180565b6001600160a01b0386166000908152600760205260408120805460ff1916871515179055600680546001929061035d90849061064b565b9091555050505050505050565b6001816005811061037a57600080fd5b01546001600160a01b0316905081565b6001600160a01b038116811461039f57600080fd5b50565b634e487b7160e01b600052604160045260246000fd5b60405160a0810167ffffffffffffffff811182821017156103db576103db6103a2565b60405290565b604051601f8201601f1916810167ffffffffffffffff8111828210171561040a5761040a6103a2565b604052919050565b60008060006060848603121561042757600080fd5b83356104328161038a565b9250602084810135801515811461044857600080fd5b9250604085013567ffffffffffffffff8082111561046557600080fd5b8187019150601f888184011261047a57600080fd5b6104826103b8565b8060a085018b81111561049457600080fd5b855b81811015610518578035868111156104ae5760008081fd5b87018581018e136104bf5760008081fd5b8035878111156104d1576104d16103a2565b6104e2818801601f19168b016103e1565b8181528f8b8385010111156104f75760008081fd5b818b84018c83013760009181018b0191909152855250928701928701610496565b50508096505050505050509250925092565b60006020828403121561053c57600080fd5b81356105478161038a565b9392505050565b60006020828403121561056057600080fd5b5035919050565b634e487b7160e01b600052603260045260246000fd5b82815260006020604081840152835180604085015260005b818110156105b157858101830151858201606001528201610595565b818111156105c3576000606083870101525b50601f01601f191692909201606001949350505050565b6000602082840312156105ec57600080fd5b81516105478161038a565b634e487b7160e01b600052601160045260246000fd5b600060ff821660ff84168060ff0382111561062a5761062a6105f7565b019392505050565b600060018201610644576106446105f7565b5060010190565b6000821982111561065e5761065e6105f7565b50019056fea2646970667358221220f0b39cb3330f44011dfeca9a19af3accf1a21deb4f6b904d71d110e1b1b14cd164736f6c634300080e0033"

// DeployAccess deploys a new Ethereum contract, binding an instance of Access to it.
func DeployAccess(auth *bind.TransactOpts, backend bind.ContractBackend, _addrs [5]common.Address) (common.Address, *types.Transaction, *Access, error) {
	parsed, err := abi.JSON(strings.NewReader(AccessABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	recoverAddr, _, _, _ := DeployRecover(auth, backend)
	AccessBin = strings.Replace(AccessBin, "__$6230a6feddd2d01b0a9ffb5c5e188a1065$__", recoverAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AccessBin), backend, _addrs)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Access{AccessCaller: AccessCaller{contract: contract}, AccessTransactor: AccessTransactor{contract: contract}, AccessFilterer: AccessFilterer{contract: contract}}, nil
}

// Access is an auto generated Go binding around an Ethereum contract.
type Access struct {
	AccessCaller     // Read-only binding to the contract
	AccessTransactor // Write-only binding to the contract
	AccessFilterer   // Log filterer for contract events
}

// AccessCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccessCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccessTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccessFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccessSession struct {
	Contract     *Access           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AccessCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccessCallerSession struct {
	Contract *AccessCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AccessTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccessTransactorSession struct {
	Contract     *AccessTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AccessRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccessRaw struct {
	Contract *Access // Generic contract binding to access the raw methods on
}

// AccessCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccessCallerRaw struct {
	Contract *AccessCaller // Generic read-only contract binding to access the raw methods on
}

// AccessTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccessTransactorRaw struct {
	Contract *AccessTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccess creates a new instance of Access, bound to a specific deployed contract.
func NewAccess(address common.Address, backend bind.ContractBackend) (*Access, error) {
	contract, err := bindAccess(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Access{AccessCaller: AccessCaller{contract: contract}, AccessTransactor: AccessTransactor{contract: contract}, AccessFilterer: AccessFilterer{contract: contract}}, nil
}

// NewAccessCaller creates a new read-only instance of Access, bound to a specific deployed contract.
func NewAccessCaller(address common.Address, caller bind.ContractCaller) (*AccessCaller, error) {
	contract, err := bindAccess(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccessCaller{contract: contract}, nil
}

// NewAccessTransactor creates a new write-only instance of Access, bound to a specific deployed contract.
func NewAccessTransactor(address common.Address, transactor bind.ContractTransactor) (*AccessTransactor, error) {
	contract, err := bindAccess(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccessTransactor{contract: contract}, nil
}

// NewAccessFilterer creates a new log filterer instance of Access, bound to a specific deployed contract.
func NewAccessFilterer(address common.Address, filterer bind.ContractFilterer) (*AccessFilterer, error) {
	contract, err := bindAccess(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccessFilterer{contract: contract}, nil
}

// bindAccess binds a generic wrapper to an already deployed contract.
func bindAccess(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AccessABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Access *AccessRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Access.Contract.AccessCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Access *AccessRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Access.Contract.AccessTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Access *AccessRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Access.Contract.AccessTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Access *AccessCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Access.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Access *AccessTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Access.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Access *AccessTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Access.Contract.contract.Transact(opts, method, params...)
}

// Access is a free data retrieval call binding the contract method 0x6fae3d76.
//
// Solidity: function access(address ) view returns(bool)
func (_Access *AccessCaller) Access(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Access.contract.Call(opts, &out, "access", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Access is a free data retrieval call binding the contract method 0x6fae3d76.
//
// Solidity: function access(address ) view returns(bool)
func (_Access *AccessSession) Access(arg0 common.Address) (bool, error) {
	return _Access.Contract.Access(&_Access.CallOpts, arg0)
}

// Access is a free data retrieval call binding the contract method 0x6fae3d76.
//
// Solidity: function access(address ) view returns(bool)
func (_Access *AccessCallerSession) Access(arg0 common.Address) (bool, error) {
	return _Access.Contract.Access(&_Access.CallOpts, arg0)
}

// Controls is a free data retrieval call binding the contract method 0xda191d88.
//
// Solidity: function controls(uint256 ) view returns(address)
func (_Access *AccessCaller) Controls(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Access.contract.Call(opts, &out, "controls", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Controls is a free data retrieval call binding the contract method 0xda191d88.
//
// Solidity: function controls(uint256 ) view returns(address)
func (_Access *AccessSession) Controls(arg0 *big.Int) (common.Address, error) {
	return _Access.Contract.Controls(&_Access.CallOpts, arg0)
}

// Controls is a free data retrieval call binding the contract method 0xda191d88.
//
// Solidity: function controls(uint256 ) view returns(address)
func (_Access *AccessCallerSession) Controls(arg0 *big.Int) (common.Address, error) {
	return _Access.Contract.Controls(&_Access.CallOpts, arg0)
}

// Floor is a free data retrieval call binding the contract method 0x40695363.
//
// Solidity: function floor() view returns(uint8)
func (_Access *AccessCaller) Floor(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Access.contract.Call(opts, &out, "floor")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Floor is a free data retrieval call binding the contract method 0x40695363.
//
// Solidity: function floor() view returns(uint8)
func (_Access *AccessSession) Floor() (uint8, error) {
	return _Access.Contract.Floor(&_Access.CallOpts)
}

// Floor is a free data retrieval call binding the contract method 0x40695363.
//
// Solidity: function floor() view returns(uint8)
func (_Access *AccessCallerSession) Floor() (uint8, error) {
	return _Access.Contract.Floor(&_Access.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_Access *AccessCaller) Nonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Access.contract.Call(opts, &out, "nonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_Access *AccessSession) Nonce() (*big.Int, error) {
	return _Access.Contract.Nonce(&_Access.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_Access *AccessCallerSession) Nonce() (*big.Int, error) {
	return _Access.Contract.Nonce(&_Access.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint16)
func (_Access *AccessCaller) Version(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Access.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint16)
func (_Access *AccessSession) Version() (uint16, error) {
	return _Access.Contract.Version(&_Access.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint16)
func (_Access *AccessCallerSession) Version() (uint16, error) {
	return _Access.Contract.Version(&_Access.CallOpts)
}

// Set is a paid mutator transaction binding the contract method 0x464b1c45.
//
// Solidity: function set(address account, bool isSet, bytes[5] signs) returns()
func (_Access *AccessTransactor) Set(opts *bind.TransactOpts, account common.Address, isSet bool, signs [5][]byte) (*types.Transaction, error) {
	return _Access.contract.Transact(opts, "set", account, isSet, signs)
}

// Set is a paid mutator transaction binding the contract method 0x464b1c45.
//
// Solidity: function set(address account, bool isSet, bytes[5] signs) returns()
func (_Access *AccessSession) Set(account common.Address, isSet bool, signs [5][]byte) (*types.Transaction, error) {
	return _Access.Contract.Set(&_Access.TransactOpts, account, isSet, signs)
}

// Set is a paid mutator transaction binding the contract method 0x464b1c45.
//
// Solidity: function set(address account, bool isSet, bytes[5] signs) returns()
func (_Access *AccessTransactorSession) Set(account common.Address, isSet bool, signs [5][]byte) (*types.Transaction, error) {
	return _Access.Contract.Set(&_Access.TransactOpts, account, isSet, signs)
}

// IAccessABI is the input ABI used to generate the binding from.
const IAccessABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"}],\"name\":\"access\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_set\",\"type\":\"bool\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IAccessFuncSigs maps the 4-byte function signature to its string representation.
var IAccessFuncSigs = map[string]string{
	"6fae3d76": "access(address)",
	"464b1c45": "set(address,bool,bytes[5])",
}

// IAccess is an auto generated Go binding around an Ethereum contract.
type IAccess struct {
	IAccessCaller     // Read-only binding to the contract
	IAccessTransactor // Write-only binding to the contract
	IAccessFilterer   // Log filterer for contract events
}

// IAccessCaller is an auto generated read-only Go binding around an Ethereum contract.
type IAccessCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAccessTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IAccessTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAccessFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAccessFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAccessSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IAccessSession struct {
	Contract     *IAccess          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAccessCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IAccessCallerSession struct {
	Contract *IAccessCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IAccessTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IAccessTransactorSession struct {
	Contract     *IAccessTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IAccessRaw is an auto generated low-level Go binding around an Ethereum contract.
type IAccessRaw struct {
	Contract *IAccess // Generic contract binding to access the raw methods on
}

// IAccessCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IAccessCallerRaw struct {
	Contract *IAccessCaller // Generic read-only contract binding to access the raw methods on
}

// IAccessTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IAccessTransactorRaw struct {
	Contract *IAccessTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIAccess creates a new instance of IAccess, bound to a specific deployed contract.
func NewIAccess(address common.Address, backend bind.ContractBackend) (*IAccess, error) {
	contract, err := bindIAccess(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAccess{IAccessCaller: IAccessCaller{contract: contract}, IAccessTransactor: IAccessTransactor{contract: contract}, IAccessFilterer: IAccessFilterer{contract: contract}}, nil
}

// NewIAccessCaller creates a new read-only instance of IAccess, bound to a specific deployed contract.
func NewIAccessCaller(address common.Address, caller bind.ContractCaller) (*IAccessCaller, error) {
	contract, err := bindIAccess(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAccessCaller{contract: contract}, nil
}

// NewIAccessTransactor creates a new write-only instance of IAccess, bound to a specific deployed contract.
func NewIAccessTransactor(address common.Address, transactor bind.ContractTransactor) (*IAccessTransactor, error) {
	contract, err := bindIAccess(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAccessTransactor{contract: contract}, nil
}

// NewIAccessFilterer creates a new log filterer instance of IAccess, bound to a specific deployed contract.
func NewIAccessFilterer(address common.Address, filterer bind.ContractFilterer) (*IAccessFilterer, error) {
	contract, err := bindIAccess(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAccessFilterer{contract: contract}, nil
}

// bindIAccess binds a generic wrapper to an already deployed contract.
func bindIAccess(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IAccessABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAccess *IAccessRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAccess.Contract.IAccessCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAccess *IAccessRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAccess.Contract.IAccessTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAccess *IAccessRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAccess.Contract.IAccessTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAccess *IAccessCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAccess.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAccess *IAccessTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAccess.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAccess *IAccessTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAccess.Contract.contract.Transact(opts, method, params...)
}

// Access is a paid mutator transaction binding the contract method 0x6fae3d76.
//
// Solidity: function access(address _a) returns(bool)
func (_IAccess *IAccessTransactor) Access(opts *bind.TransactOpts, _a common.Address) (*types.Transaction, error) {
	return _IAccess.contract.Transact(opts, "access", _a)
}

// Access is a paid mutator transaction binding the contract method 0x6fae3d76.
//
// Solidity: function access(address _a) returns(bool)
func (_IAccess *IAccessSession) Access(_a common.Address) (*types.Transaction, error) {
	return _IAccess.Contract.Access(&_IAccess.TransactOpts, _a)
}

// Access is a paid mutator transaction binding the contract method 0x6fae3d76.
//
// Solidity: function access(address _a) returns(bool)
func (_IAccess *IAccessTransactorSession) Access(_a common.Address) (*types.Transaction, error) {
	return _IAccess.Contract.Access(&_IAccess.TransactOpts, _a)
}

// Set is a paid mutator transaction binding the contract method 0x464b1c45.
//
// Solidity: function set(address _a, bool _set, bytes[5] signs) returns()
func (_IAccess *IAccessTransactor) Set(opts *bind.TransactOpts, _a common.Address, _set bool, signs [5][]byte) (*types.Transaction, error) {
	return _IAccess.contract.Transact(opts, "set", _a, _set, signs)
}

// Set is a paid mutator transaction binding the contract method 0x464b1c45.
//
// Solidity: function set(address _a, bool _set, bytes[5] signs) returns()
func (_IAccess *IAccessSession) Set(_a common.Address, _set bool, signs [5][]byte) (*types.Transaction, error) {
	return _IAccess.Contract.Set(&_IAccess.TransactOpts, _a, _set, signs)
}

// Set is a paid mutator transaction binding the contract method 0x464b1c45.
//
// Solidity: function set(address _a, bool _set, bytes[5] signs) returns()
func (_IAccess *IAccessTransactorSession) Set(_a common.Address, _set bool, signs [5][]byte) (*types.Transaction, error) {
	return _IAccess.Contract.Set(&_IAccess.TransactOpts, _a, _set, signs)
}

// RecoverABI is the input ABI used to generate the binding from.
const RecoverABI = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"recover\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// RecoverFuncSigs maps the 4-byte function signature to its string representation.
var RecoverFuncSigs = map[string]string{
	"19045a25": "recover(bytes32,bytes)",
}

// RecoverBin is the compiled bytecode used for deploying new contracts.
var RecoverBin = "0x6102a861003a600b82828239805160001a60731461002d57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100355760003560e01c806319045a251461003a575b600080fd5b61004d610048366004610184565b610069565b6040516001600160a01b03909116815260200160405180910390f35b6000815160411461007c57506000610168565b60208201516040830151606084015160001a7f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08211156100c25760009350505050610168565b601b8160ff1610156100dc576100d981601b61023f565b90505b8060ff16601b141580156100f457508060ff16601c14155b156101055760009350505050610168565b60408051600081526020810180835288905260ff831691810191909152606081018490526080810183905260019060a0016020604051602081039080840390855afa158015610158573d6000803e3d6000fd5b5050506020604051035193505050505b92915050565b634e487b7160e01b600052604160045260246000fd5b6000806040838503121561019757600080fd5b82359150602083013567ffffffffffffffff808211156101b657600080fd5b818501915085601f8301126101ca57600080fd5b8135818111156101dc576101dc61016e565b604051601f8201601f19908116603f011681019083821181831017156102045761020461016e565b8160405282815288602084870101111561021d57600080fd5b8260208601602083013760006020848301015280955050505050509250929050565b600060ff821660ff84168060ff0382111561026a57634e487b7160e01b600052601160045260246000fd5b01939250505056fea26469706673582212200d88ed66bdb84578aad3e4db894d00f01a13c6b71be38cba41e6c2a7a75b243364736f6c634300080e0033"

// DeployRecover deploys a new Ethereum contract, binding an instance of Recover to it.
func DeployRecover(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Recover, error) {
	parsed, err := abi.JSON(strings.NewReader(RecoverABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RecoverBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Recover{RecoverCaller: RecoverCaller{contract: contract}, RecoverTransactor: RecoverTransactor{contract: contract}, RecoverFilterer: RecoverFilterer{contract: contract}}, nil
}

// Recover is an auto generated Go binding around an Ethereum contract.
type Recover struct {
	RecoverCaller     // Read-only binding to the contract
	RecoverTransactor // Write-only binding to the contract
	RecoverFilterer   // Log filterer for contract events
}

// RecoverCaller is an auto generated read-only Go binding around an Ethereum contract.
type RecoverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RecoverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RecoverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RecoverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RecoverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RecoverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RecoverSession struct {
	Contract     *Recover          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RecoverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RecoverCallerSession struct {
	Contract *RecoverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// RecoverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RecoverTransactorSession struct {
	Contract     *RecoverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// RecoverRaw is an auto generated low-level Go binding around an Ethereum contract.
type RecoverRaw struct {
	Contract *Recover // Generic contract binding to access the raw methods on
}

// RecoverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RecoverCallerRaw struct {
	Contract *RecoverCaller // Generic read-only contract binding to access the raw methods on
}

// RecoverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RecoverTransactorRaw struct {
	Contract *RecoverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRecover creates a new instance of Recover, bound to a specific deployed contract.
func NewRecover(address common.Address, backend bind.ContractBackend) (*Recover, error) {
	contract, err := bindRecover(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Recover{RecoverCaller: RecoverCaller{contract: contract}, RecoverTransactor: RecoverTransactor{contract: contract}, RecoverFilterer: RecoverFilterer{contract: contract}}, nil
}

// NewRecoverCaller creates a new read-only instance of Recover, bound to a specific deployed contract.
func NewRecoverCaller(address common.Address, caller bind.ContractCaller) (*RecoverCaller, error) {
	contract, err := bindRecover(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RecoverCaller{contract: contract}, nil
}

// NewRecoverTransactor creates a new write-only instance of Recover, bound to a specific deployed contract.
func NewRecoverTransactor(address common.Address, transactor bind.ContractTransactor) (*RecoverTransactor, error) {
	contract, err := bindRecover(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RecoverTransactor{contract: contract}, nil
}

// NewRecoverFilterer creates a new log filterer instance of Recover, bound to a specific deployed contract.
func NewRecoverFilterer(address common.Address, filterer bind.ContractFilterer) (*RecoverFilterer, error) {
	contract, err := bindRecover(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RecoverFilterer{contract: contract}, nil
}

// bindRecover binds a generic wrapper to an already deployed contract.
func bindRecover(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RecoverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Recover *RecoverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Recover.Contract.RecoverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Recover *RecoverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Recover.Contract.RecoverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Recover *RecoverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Recover.Contract.RecoverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Recover *RecoverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Recover.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Recover *RecoverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Recover.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Recover *RecoverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Recover.Contract.contract.Transact(opts, method, params...)
}

// Recover is a free data retrieval call binding the contract method 0x19045a25.
//
// Solidity: function recover(bytes32 hash, bytes signature) pure returns(address)
func (_Recover *RecoverCaller) Recover(opts *bind.CallOpts, hash [32]byte, signature []byte) (common.Address, error) {
	var out []interface{}
	err := _Recover.contract.Call(opts, &out, "recover", hash, signature)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Recover is a free data retrieval call binding the contract method 0x19045a25.
//
// Solidity: function recover(bytes32 hash, bytes signature) pure returns(address)
func (_Recover *RecoverSession) Recover(hash [32]byte, signature []byte) (common.Address, error) {
	return _Recover.Contract.Recover(&_Recover.CallOpts, hash, signature)
}

// Recover is a free data retrieval call binding the contract method 0x19045a25.
//
// Solidity: function recover(bytes32 hash, bytes signature) pure returns(address)
func (_Recover *RecoverCallerSession) Recover(hash [32]byte, signature []byte) (common.Address, error) {
	return _Recover.Contract.Recover(&_Recover.CallOpts, hash, signature)
}
