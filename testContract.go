// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// TestContractABI is the input ABI used to generate the binding from.
const TestContractABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"num1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"num2\",\"type\":\"uint256\"}],\"name\":\"add\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"num1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"num2\",\"type\":\"uint256\"}],\"name\":\"mult\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"store\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// TestContractFuncSigs maps the 4-byte function signature to its string representation.
var TestContractFuncSigs = map[string]string{
	"771602f7": "add(uint256,uint256)",
	"9aa727f6": "mult(uint256,uint256)",
	"131a0680": "store(string)",
}

// TestContractBin is the compiled bytecode used for deploying new contracts.
var TestContractBin = "0x608060405234801561001057600080fd5b5061030e806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c8063131a068014610046578063771602f71461005b5780639aa727f614610084575b600080fd5b610059610054366004610166565b610097565b005b61006e610069366004610210565b6100ae565b60405161007b9190610231565b60405180910390f35b61006e610092366004610210565b6100c1565b80516100aa9060019060208401906100cd565b5050565b60006100ba828461023a565b9392505050565b60006100ba8284610252565b8280546100d990610271565b90600052602060002090601f0160209004810192826100fb5760008555610141565b82601f1061011457805160ff1916838001178555610141565b82800160010185558215610141579182015b82811115610141578251825591602001919060010190610126565b5061014d929150610151565b5090565b5b8082111561014d5760008155600101610152565b600060208284031215610177578081fd5b813567ffffffffffffffff8082111561018e578283fd5b818401915084601f8301126101a1578283fd5b8135818111156101b3576101b36102c2565b604051601f8201601f19908116603f011681019083821181831017156101db576101db6102c2565b816040528281528760208487010111156101f3578586fd5b826020860160208301379182016020019490945295945050505050565b60008060408385031215610222578081fd5b50508035926020909101359150565b90815260200190565b6000821982111561024d5761024d6102ac565b500190565b600081600019048311821515161561026c5761026c6102ac565b500290565b60028104600182168061028557607f821691505b602082108114156102a657634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052604160045260246000fdfea26469706673582212203aa934f402320a031974fff231a464090687c9a7e74ca4e37aba352bc243150a64736f6c63430008010033"

// DeployTestContract deploys a new Ethereum contract, binding an instance of TestContract to it.
func DeployTestContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TestContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TestContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestContract{TestContractCaller: TestContractCaller{contract: contract}, TestContractTransactor: TestContractTransactor{contract: contract}, TestContractFilterer: TestContractFilterer{contract: contract}}, nil
}

// TestContract is an auto generated Go binding around an Ethereum contract.
type TestContract struct {
	TestContractCaller     // Read-only binding to the contract
	TestContractTransactor // Write-only binding to the contract
	TestContractFilterer   // Log filterer for contract events
}

// TestContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestContractSession struct {
	Contract     *TestContract     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestContractCallerSession struct {
	Contract *TestContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// TestContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestContractTransactorSession struct {
	Contract     *TestContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TestContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestContractRaw struct {
	Contract *TestContract // Generic contract binding to access the raw methods on
}

// TestContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestContractCallerRaw struct {
	Contract *TestContractCaller // Generic read-only contract binding to access the raw methods on
}

// TestContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestContractTransactorRaw struct {
	Contract *TestContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestContract creates a new instance of TestContract, bound to a specific deployed contract.
func NewTestContract(address common.Address, backend bind.ContractBackend) (*TestContract, error) {
	contract, err := bindTestContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestContract{TestContractCaller: TestContractCaller{contract: contract}, TestContractTransactor: TestContractTransactor{contract: contract}, TestContractFilterer: TestContractFilterer{contract: contract}}, nil
}

// NewTestContractCaller creates a new read-only instance of TestContract, bound to a specific deployed contract.
func NewTestContractCaller(address common.Address, caller bind.ContractCaller) (*TestContractCaller, error) {
	contract, err := bindTestContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestContractCaller{contract: contract}, nil
}

// NewTestContractTransactor creates a new write-only instance of TestContract, bound to a specific deployed contract.
func NewTestContractTransactor(address common.Address, transactor bind.ContractTransactor) (*TestContractTransactor, error) {
	contract, err := bindTestContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestContractTransactor{contract: contract}, nil
}

// NewTestContractFilterer creates a new log filterer instance of TestContract, bound to a specific deployed contract.
func NewTestContractFilterer(address common.Address, filterer bind.ContractFilterer) (*TestContractFilterer, error) {
	contract, err := bindTestContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestContractFilterer{contract: contract}, nil
}

// bindTestContract binds a generic wrapper to an already deployed contract.
func bindTestContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestContract *TestContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestContract.Contract.TestContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestContract *TestContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestContract.Contract.TestContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestContract *TestContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestContract.Contract.TestContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestContract *TestContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestContract *TestContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestContract *TestContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestContract.Contract.contract.Transact(opts, method, params...)
}

// Add is a free data retrieval call binding the contract method 0x771602f7.
//
// Solidity: function add(uint256 num1, uint256 num2) pure returns(uint256)
func (_TestContract *TestContractCaller) Add(opts *bind.CallOpts, num1 *big.Int, num2 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TestContract.contract.Call(opts, &out, "add", num1, num2)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Add is a free data retrieval call binding the contract method 0x771602f7.
//
// Solidity: function add(uint256 num1, uint256 num2) pure returns(uint256)
func (_TestContract *TestContractSession) Add(num1 *big.Int, num2 *big.Int) (*big.Int, error) {
	return _TestContract.Contract.Add(&_TestContract.CallOpts, num1, num2)
}

// Add is a free data retrieval call binding the contract method 0x771602f7.
//
// Solidity: function add(uint256 num1, uint256 num2) pure returns(uint256)
func (_TestContract *TestContractCallerSession) Add(num1 *big.Int, num2 *big.Int) (*big.Int, error) {
	return _TestContract.Contract.Add(&_TestContract.CallOpts, num1, num2)
}

// Mult is a free data retrieval call binding the contract method 0x9aa727f6.
//
// Solidity: function mult(uint256 num1, uint256 num2) pure returns(uint256)
func (_TestContract *TestContractCaller) Mult(opts *bind.CallOpts, num1 *big.Int, num2 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TestContract.contract.Call(opts, &out, "mult", num1, num2)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Mult is a free data retrieval call binding the contract method 0x9aa727f6.
//
// Solidity: function mult(uint256 num1, uint256 num2) pure returns(uint256)
func (_TestContract *TestContractSession) Mult(num1 *big.Int, num2 *big.Int) (*big.Int, error) {
	return _TestContract.Contract.Mult(&_TestContract.CallOpts, num1, num2)
}

// Mult is a free data retrieval call binding the contract method 0x9aa727f6.
//
// Solidity: function mult(uint256 num1, uint256 num2) pure returns(uint256)
func (_TestContract *TestContractCallerSession) Mult(num1 *big.Int, num2 *big.Int) (*big.Int, error) {
	return _TestContract.Contract.Mult(&_TestContract.CallOpts, num1, num2)
}

// Store is a paid mutator transaction binding the contract method 0x131a0680.
//
// Solidity: function store(string str) returns()
func (_TestContract *TestContractTransactor) Store(opts *bind.TransactOpts, str string) (*types.Transaction, error) {
	return _TestContract.contract.Transact(opts, "store", str)
}

// Store is a paid mutator transaction binding the contract method 0x131a0680.
//
// Solidity: function store(string str) returns()
func (_TestContract *TestContractSession) Store(str string) (*types.Transaction, error) {
	return _TestContract.Contract.Store(&_TestContract.TransactOpts, str)
}

// Store is a paid mutator transaction binding the contract method 0x131a0680.
//
// Solidity: function store(string str) returns()
func (_TestContract *TestContractTransactorSession) Store(str string) (*types.Transaction, error) {
	return _TestContract.Contract.Store(&_TestContract.TransactOpts, str)
}
