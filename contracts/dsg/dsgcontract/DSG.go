// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dsgcontract

import (
	"strings"

	"github.com/unification-com/mainchain/accounts/abi"
	"github.com/unification-com/mainchain/accounts/abi/bind"
	"github.com/unification-com/mainchain/common"
	"github.com/unification-com/mainchain/core/types"
)

// DsgcontractABI is the input ABI used to generate the binding from.
const DsgcontractABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"x1\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x2\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x3\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"rotate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// DsgcontractBin is the compiled bytecode used for deploying new contracts.
const DsgcontractBin = `608060405273a13a71dfe5cd57f9b67ec6a54ad2ae7537d7fc3b6000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073b47fd09f1d379ce2c9bff59d668cf0b7b994a2b7600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073ca29f1275470de81de6e1bb53a55228da676e752600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555034801561010e57600080fd5b5061036f8061011e6000396000f300608060405260043610610062576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063343943bd146100675780634ff13571146100be578063a3b7887314610115578063d992818d1461016c575b600080fd5b34801561007357600080fd5b5061007c610183565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156100ca57600080fd5b506100d36101a8565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561012157600080fd5b5061012a6101ce565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561017857600080fd5b506101816101f4565b005b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b73f30f951b0426f8bf37ac71967407081358df7a7b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561024257610341565b73a13a71dfe5cd57f9b67ec6a54ad2ae7537d7fc3b6000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073b47fd09f1d379ce2c9bff59d668cf0b7b994a2b7600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073d71ad3263666e03004479b30ca840a26eac5b763600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b5600a165627a7a72305820030badbcfef3d6b134e8ae318d57c183da62667b40ae9a25135cf533bf92e2020029`

// DeployDsgcontract deploys a new Ethereum contract, binding an instance of Dsgcontract to it.
func DeployDsgcontract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Dsgcontract, error) {
	parsed, err := abi.JSON(strings.NewReader(DsgcontractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DsgcontractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Dsgcontract{DsgcontractCaller: DsgcontractCaller{contract: contract}, DsgcontractTransactor: DsgcontractTransactor{contract: contract}, DsgcontractFilterer: DsgcontractFilterer{contract: contract}}, nil
}

// Dsgcontract is an auto generated Go binding around an Ethereum contract.
type Dsgcontract struct {
	DsgcontractCaller     // Read-only binding to the contract
	DsgcontractTransactor // Write-only binding to the contract
	DsgcontractFilterer   // Log filterer for contract events
}

// DsgcontractCaller is an auto generated read-only Go binding around an Ethereum contract.
type DsgcontractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DsgcontractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DsgcontractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DsgcontractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DsgcontractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DsgcontractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DsgcontractSession struct {
	Contract     *Dsgcontract      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DsgcontractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DsgcontractCallerSession struct {
	Contract *DsgcontractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// DsgcontractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DsgcontractTransactorSession struct {
	Contract     *DsgcontractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// DsgcontractRaw is an auto generated low-level Go binding around an Ethereum contract.
type DsgcontractRaw struct {
	Contract *Dsgcontract // Generic contract binding to access the raw methods on
}

// DsgcontractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DsgcontractCallerRaw struct {
	Contract *DsgcontractCaller // Generic read-only contract binding to access the raw methods on
}

// DsgcontractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DsgcontractTransactorRaw struct {
	Contract *DsgcontractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDsgcontract creates a new instance of Dsgcontract, bound to a specific deployed contract.
func NewDsgcontract(address common.Address, backend bind.ContractBackend) (*Dsgcontract, error) {
	contract, err := bindDsgcontract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Dsgcontract{DsgcontractCaller: DsgcontractCaller{contract: contract}, DsgcontractTransactor: DsgcontractTransactor{contract: contract}, DsgcontractFilterer: DsgcontractFilterer{contract: contract}}, nil
}

// NewDsgcontractCaller creates a new read-only instance of Dsgcontract, bound to a specific deployed contract.
func NewDsgcontractCaller(address common.Address, caller bind.ContractCaller) (*DsgcontractCaller, error) {
	contract, err := bindDsgcontract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DsgcontractCaller{contract: contract}, nil
}

// NewDsgcontractTransactor creates a new write-only instance of Dsgcontract, bound to a specific deployed contract.
func NewDsgcontractTransactor(address common.Address, transactor bind.ContractTransactor) (*DsgcontractTransactor, error) {
	contract, err := bindDsgcontract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DsgcontractTransactor{contract: contract}, nil
}

// NewDsgcontractFilterer creates a new log filterer instance of Dsgcontract, bound to a specific deployed contract.
func NewDsgcontractFilterer(address common.Address, filterer bind.ContractFilterer) (*DsgcontractFilterer, error) {
	contract, err := bindDsgcontract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DsgcontractFilterer{contract: contract}, nil
}

// bindDsgcontract binds a generic wrapper to an already deployed contract.
func bindDsgcontract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DsgcontractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dsgcontract *DsgcontractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Dsgcontract.Contract.DsgcontractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dsgcontract *DsgcontractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dsgcontract.Contract.DsgcontractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dsgcontract *DsgcontractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dsgcontract.Contract.DsgcontractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dsgcontract *DsgcontractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Dsgcontract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dsgcontract *DsgcontractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dsgcontract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dsgcontract *DsgcontractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dsgcontract.Contract.contract.Transact(opts, method, params...)
}

// X1 is a free data retrieval call binding the contract method 0x343943bd.
//
// Solidity: function x1() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X1(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x1")
	return *ret0, err
}

// X1 is a free data retrieval call binding the contract method 0x343943bd.
//
// Solidity: function x1() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X1() (common.Address, error) {
	return _Dsgcontract.Contract.X1(&_Dsgcontract.CallOpts)
}

// X1 is a free data retrieval call binding the contract method 0x343943bd.
//
// Solidity: function x1() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X1() (common.Address, error) {
	return _Dsgcontract.Contract.X1(&_Dsgcontract.CallOpts)
}

// X2 is a free data retrieval call binding the contract method 0x4ff13571.
//
// Solidity: function x2() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X2(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x2")
	return *ret0, err
}

// X2 is a free data retrieval call binding the contract method 0x4ff13571.
//
// Solidity: function x2() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X2() (common.Address, error) {
	return _Dsgcontract.Contract.X2(&_Dsgcontract.CallOpts)
}

// X2 is a free data retrieval call binding the contract method 0x4ff13571.
//
// Solidity: function x2() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X2() (common.Address, error) {
	return _Dsgcontract.Contract.X2(&_Dsgcontract.CallOpts)
}

// X3 is a free data retrieval call binding the contract method 0xa3b78873.
//
// Solidity: function x3() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X3(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x3")
	return *ret0, err
}

// X3 is a free data retrieval call binding the contract method 0xa3b78873.
//
// Solidity: function x3() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X3() (common.Address, error) {
	return _Dsgcontract.Contract.X3(&_Dsgcontract.CallOpts)
}

// X3 is a free data retrieval call binding the contract method 0xa3b78873.
//
// Solidity: function x3() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X3() (common.Address, error) {
	return _Dsgcontract.Contract.X3(&_Dsgcontract.CallOpts)
}

// Rotate is a paid mutator transaction binding the contract method 0xd992818d.
//
// Solidity: function rotate() returns()
func (_Dsgcontract *DsgcontractTransactor) Rotate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dsgcontract.contract.Transact(opts, "rotate")
}

// Rotate is a paid mutator transaction binding the contract method 0xd992818d.
//
// Solidity: function rotate() returns()
func (_Dsgcontract *DsgcontractSession) Rotate() (*types.Transaction, error) {
	return _Dsgcontract.Contract.Rotate(&_Dsgcontract.TransactOpts)
}

// Rotate is a paid mutator transaction binding the contract method 0xd992818d.
//
// Solidity: function rotate() returns()
func (_Dsgcontract *DsgcontractTransactorSession) Rotate() (*types.Transaction, error) {
	return _Dsgcontract.Contract.Rotate(&_Dsgcontract.TransactOpts)
}
