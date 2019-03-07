// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package SampleContract

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// SampleContractABI is the input ABI used to generate the binding from.
const SampleContractABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allowed\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"addAllowed\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// SampleContractBin is the compiled bytecode used for deploying new contracts.
const SampleContractBin = `608060405234801561001057600080fd5b506101c0806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063214296e81461003b578063cb8523c6146100a9575b600080fd5b6100676004803603602081101561005157600080fd5b81019080803590602001909291905050506100ed565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6100eb600480360360208110156100bf57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061012b565b005b6000818154811015156100fc57fe5b906000526020600020016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008190806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505056fea165627a7a72305820e4716a4585195c39f8bddb8305568e5084677113fc70e6665a92ac16a5ad8b050029`

// DeploySampleContract deploys a new Ethereum contract, binding an instance of SampleContract to it.
func DeploySampleContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SampleContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SampleContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SampleContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SampleContract{SampleContractCaller: SampleContractCaller{contract: contract}, SampleContractTransactor: SampleContractTransactor{contract: contract}, SampleContractFilterer: SampleContractFilterer{contract: contract}}, nil
}

// SampleContract is an auto generated Go binding around an Ethereum contract.
type SampleContract struct {
	SampleContractCaller     // Read-only binding to the contract
	SampleContractTransactor // Write-only binding to the contract
	SampleContractFilterer   // Log filterer for contract events
}

// SampleContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type SampleContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SampleContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SampleContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SampleContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SampleContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SampleContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SampleContractSession struct {
	Contract     *SampleContract   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SampleContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SampleContractCallerSession struct {
	Contract *SampleContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// SampleContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SampleContractTransactorSession struct {
	Contract     *SampleContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// SampleContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type SampleContractRaw struct {
	Contract *SampleContract // Generic contract binding to access the raw methods on
}

// SampleContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SampleContractCallerRaw struct {
	Contract *SampleContractCaller // Generic read-only contract binding to access the raw methods on
}

// SampleContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SampleContractTransactorRaw struct {
	Contract *SampleContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSampleContract creates a new instance of SampleContract, bound to a specific deployed contract.
func NewSampleContract(address common.Address, backend bind.ContractBackend) (*SampleContract, error) {
	contract, err := bindSampleContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SampleContract{SampleContractCaller: SampleContractCaller{contract: contract}, SampleContractTransactor: SampleContractTransactor{contract: contract}, SampleContractFilterer: SampleContractFilterer{contract: contract}}, nil
}

// NewSampleContractCaller creates a new read-only instance of SampleContract, bound to a specific deployed contract.
func NewSampleContractCaller(address common.Address, caller bind.ContractCaller) (*SampleContractCaller, error) {
	contract, err := bindSampleContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SampleContractCaller{contract: contract}, nil
}

// NewSampleContractTransactor creates a new write-only instance of SampleContract, bound to a specific deployed contract.
func NewSampleContractTransactor(address common.Address, transactor bind.ContractTransactor) (*SampleContractTransactor, error) {
	contract, err := bindSampleContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SampleContractTransactor{contract: contract}, nil
}

// NewSampleContractFilterer creates a new log filterer instance of SampleContract, bound to a specific deployed contract.
func NewSampleContractFilterer(address common.Address, filterer bind.ContractFilterer) (*SampleContractFilterer, error) {
	contract, err := bindSampleContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SampleContractFilterer{contract: contract}, nil
}

// bindSampleContract binds a generic wrapper to an already deployed contract.
func bindSampleContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SampleContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SampleContract *SampleContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SampleContract.Contract.SampleContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SampleContract *SampleContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SampleContract.Contract.SampleContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SampleContract *SampleContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SampleContract.Contract.SampleContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SampleContract *SampleContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SampleContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SampleContract *SampleContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SampleContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SampleContract *SampleContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SampleContract.Contract.contract.Transact(opts, method, params...)
}

// Allowed is a free data retrieval call binding the contract method 0x214296e8.
//
// Solidity: function allowed( uint256) constant returns(address)
func (_SampleContract *SampleContractCaller) Allowed(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SampleContract.contract.Call(opts, out, "allowed", arg0)
	return *ret0, err
}

// Allowed is a free data retrieval call binding the contract method 0x214296e8.
//
// Solidity: function allowed( uint256) constant returns(address)
func (_SampleContract *SampleContractSession) Allowed(arg0 *big.Int) (common.Address, error) {
	return _SampleContract.Contract.Allowed(&_SampleContract.CallOpts, arg0)
}

// Allowed is a free data retrieval call binding the contract method 0x214296e8.
//
// Solidity: function allowed( uint256) constant returns(address)
func (_SampleContract *SampleContractCallerSession) Allowed(arg0 *big.Int) (common.Address, error) {
	return _SampleContract.Contract.Allowed(&_SampleContract.CallOpts, arg0)
}

// AddAllowed is a paid mutator transaction binding the contract method 0xcb8523c6.
//
// Solidity: function addAllowed(_addr address) returns()
func (_SampleContract *SampleContractTransactor) AddAllowed(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _SampleContract.contract.Transact(opts, "addAllowed", _addr)
}

// AddAllowed is a paid mutator transaction binding the contract method 0xcb8523c6.
//
// Solidity: function addAllowed(_addr address) returns()
func (_SampleContract *SampleContractSession) AddAllowed(_addr common.Address) (*types.Transaction, error) {
	return _SampleContract.Contract.AddAllowed(&_SampleContract.TransactOpts, _addr)
}

// AddAllowed is a paid mutator transaction binding the contract method 0xcb8523c6.
//
// Solidity: function addAllowed(_addr address) returns()
func (_SampleContract *SampleContractTransactorSession) AddAllowed(_addr common.Address) (*types.Transaction, error) {
	return _SampleContract.Contract.AddAllowed(&_SampleContract.TransactOpts, _addr)
}
