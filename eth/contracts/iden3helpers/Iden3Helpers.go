// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Iden3Helpers

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

// Iden3HelpersABI is the input ABI used to generate the binding from.
const Iden3HelpersABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\"},{\"name\":\"rsv\",\"type\":\"bytes\"}],\"name\":\"checkSig\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes31\"}],\"name\":\"getRootFromId\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes27\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_mimcContractAddr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// Iden3HelpersBin is the compiled bytecode used for deploying new contracts.
const Iden3HelpersBin = `0x608060405234801561001057600080fd5b5060405161028b38038061028b8339818101604052602081101561003357600080fd5b5051600080546001600160a01b039092166001600160a01b0319909216919091179055610226806100656000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c806301b0452c1461003b578063ad05a8d214610104575b600080fd5b6100e86004803603604081101561005157600080fd5b8135919081019060408101602082013564010000000081111561007357600080fd5b82018360208201111561008557600080fd5b803590602001918460018302840111640100000000831117156100a757600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610140945050505050565b604080516001600160a01b039092168252519081900360200190f35b6101256004803603602081101561011a57600080fd5b503560ff19166101bf565b6040805164ffffffffff199092168252519081900360200190f35b602081810151604080840151606080860151835160008082528188018087528a905291821a81860181905292810186905260808101849052935190959293919260019260a080820193601f1981019281900390910190855afa1580156101aa573d6000803e3d6000fd5b5050604051601f190151979650505050505050565b60101b62ffffff19169056fea265627a7a723058205e0a0751b58d0516f5de3d5755106ae8c51bfed662ae66d970b85182c0c3a03564736f6c637828302e352e31302d646576656c6f702e323031392e362e31332b636f6d6d69742e36363839373262620058`

// DeployIden3Helpers deploys a new Ethereum contract, binding an instance of Iden3Helpers to it.
func DeployIden3Helpers(auth *bind.TransactOpts, backend bind.ContractBackend, _mimcContractAddr common.Address) (common.Address, *types.Transaction, *Iden3Helpers, error) {
	parsed, err := abi.JSON(strings.NewReader(Iden3HelpersABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(Iden3HelpersBin), backend, _mimcContractAddr)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Iden3Helpers{Iden3HelpersCaller: Iden3HelpersCaller{contract: contract}, Iden3HelpersTransactor: Iden3HelpersTransactor{contract: contract}, Iden3HelpersFilterer: Iden3HelpersFilterer{contract: contract}}, nil
}

// Iden3Helpers is an auto generated Go binding around an Ethereum contract.
type Iden3Helpers struct {
	Iden3HelpersCaller     // Read-only binding to the contract
	Iden3HelpersTransactor // Write-only binding to the contract
	Iden3HelpersFilterer   // Log filterer for contract events
}

// Iden3HelpersCaller is an auto generated read-only Go binding around an Ethereum contract.
type Iden3HelpersCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Iden3HelpersTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Iden3HelpersTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Iden3HelpersFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Iden3HelpersFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Iden3HelpersSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Iden3HelpersSession struct {
	Contract     *Iden3Helpers     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Iden3HelpersCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Iden3HelpersCallerSession struct {
	Contract *Iden3HelpersCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// Iden3HelpersTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Iden3HelpersTransactorSession struct {
	Contract     *Iden3HelpersTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// Iden3HelpersRaw is an auto generated low-level Go binding around an Ethereum contract.
type Iden3HelpersRaw struct {
	Contract *Iden3Helpers // Generic contract binding to access the raw methods on
}

// Iden3HelpersCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Iden3HelpersCallerRaw struct {
	Contract *Iden3HelpersCaller // Generic read-only contract binding to access the raw methods on
}

// Iden3HelpersTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Iden3HelpersTransactorRaw struct {
	Contract *Iden3HelpersTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIden3Helpers creates a new instance of Iden3Helpers, bound to a specific deployed contract.
func NewIden3Helpers(address common.Address, backend bind.ContractBackend) (*Iden3Helpers, error) {
	contract, err := bindIden3Helpers(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Iden3Helpers{Iden3HelpersCaller: Iden3HelpersCaller{contract: contract}, Iden3HelpersTransactor: Iden3HelpersTransactor{contract: contract}, Iden3HelpersFilterer: Iden3HelpersFilterer{contract: contract}}, nil
}

// NewIden3HelpersCaller creates a new read-only instance of Iden3Helpers, bound to a specific deployed contract.
func NewIden3HelpersCaller(address common.Address, caller bind.ContractCaller) (*Iden3HelpersCaller, error) {
	contract, err := bindIden3Helpers(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Iden3HelpersCaller{contract: contract}, nil
}

// NewIden3HelpersTransactor creates a new write-only instance of Iden3Helpers, bound to a specific deployed contract.
func NewIden3HelpersTransactor(address common.Address, transactor bind.ContractTransactor) (*Iden3HelpersTransactor, error) {
	contract, err := bindIden3Helpers(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Iden3HelpersTransactor{contract: contract}, nil
}

// NewIden3HelpersFilterer creates a new log filterer instance of Iden3Helpers, bound to a specific deployed contract.
func NewIden3HelpersFilterer(address common.Address, filterer bind.ContractFilterer) (*Iden3HelpersFilterer, error) {
	contract, err := bindIden3Helpers(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Iden3HelpersFilterer{contract: contract}, nil
}

// bindIden3Helpers binds a generic wrapper to an already deployed contract.
func bindIden3Helpers(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Iden3HelpersABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iden3Helpers *Iden3HelpersRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Iden3Helpers.Contract.Iden3HelpersCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iden3Helpers *Iden3HelpersRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iden3Helpers.Contract.Iden3HelpersTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iden3Helpers *Iden3HelpersRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iden3Helpers.Contract.Iden3HelpersTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iden3Helpers *Iden3HelpersCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Iden3Helpers.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iden3Helpers *Iden3HelpersTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iden3Helpers.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iden3Helpers *Iden3HelpersTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iden3Helpers.Contract.contract.Transact(opts, method, params...)
}

// CheckSig is a free data retrieval call binding the contract method 0x01b0452c.
//
// Solidity: function checkSig(bytes32 msgHash, bytes rsv) constant returns(address)
func (_Iden3Helpers *Iden3HelpersCaller) CheckSig(opts *bind.CallOpts, msgHash [32]byte, rsv []byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Iden3Helpers.contract.Call(opts, out, "checkSig", msgHash, rsv)
	return *ret0, err
}

// CheckSig is a free data retrieval call binding the contract method 0x01b0452c.
//
// Solidity: function checkSig(bytes32 msgHash, bytes rsv) constant returns(address)
func (_Iden3Helpers *Iden3HelpersSession) CheckSig(msgHash [32]byte, rsv []byte) (common.Address, error) {
	return _Iden3Helpers.Contract.CheckSig(&_Iden3Helpers.CallOpts, msgHash, rsv)
}

// CheckSig is a free data retrieval call binding the contract method 0x01b0452c.
//
// Solidity: function checkSig(bytes32 msgHash, bytes rsv) constant returns(address)
func (_Iden3Helpers *Iden3HelpersCallerSession) CheckSig(msgHash [32]byte, rsv []byte) (common.Address, error) {
	return _Iden3Helpers.Contract.CheckSig(&_Iden3Helpers.CallOpts, msgHash, rsv)
}

// GetRootFromId is a free data retrieval call binding the contract method 0xad05a8d2.
//
// Solidity: function getRootFromId(bytes31 id) constant returns(bytes27)
func (_Iden3Helpers *Iden3HelpersCaller) GetRootFromId(opts *bind.CallOpts, id [31]byte) ([27]byte, error) {
	var (
		ret0 = new([27]byte)
	)
	out := ret0
	err := _Iden3Helpers.contract.Call(opts, out, "getRootFromId", id)
	return *ret0, err
}

// GetRootFromId is a free data retrieval call binding the contract method 0xad05a8d2.
//
// Solidity: function getRootFromId(bytes31 id) constant returns(bytes27)
func (_Iden3Helpers *Iden3HelpersSession) GetRootFromId(id [31]byte) ([27]byte, error) {
	return _Iden3Helpers.Contract.GetRootFromId(&_Iden3Helpers.CallOpts, id)
}

// GetRootFromId is a free data retrieval call binding the contract method 0xad05a8d2.
//
// Solidity: function getRootFromId(bytes31 id) constant returns(bytes27)
func (_Iden3Helpers *Iden3HelpersCallerSession) GetRootFromId(id [31]byte) ([27]byte, error) {
	return _Iden3Helpers.Contract.GetRootFromId(&_Iden3Helpers.CallOpts, id)
}

// MemoryABI is the input ABI used to generate the binding from.
const MemoryABI = "[]"

// MemoryBin is the compiled bytecode used for deploying new contracts.
const MemoryBin = `0x607b6023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723058207d716ce03ca9fef6e6bcc1cde4840b90f442bb019e0fe5e761facba88198f90864736f6c637828302e352e31302d646576656c6f702e323031392e362e31332b636f6d6d69742e36363839373262620058`

// DeployMemory deploys a new Ethereum contract, binding an instance of Memory to it.
func DeployMemory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Memory, error) {
	parsed, err := abi.JSON(strings.NewReader(MemoryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MemoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Memory{MemoryCaller: MemoryCaller{contract: contract}, MemoryTransactor: MemoryTransactor{contract: contract}, MemoryFilterer: MemoryFilterer{contract: contract}}, nil
}

// Memory is an auto generated Go binding around an Ethereum contract.
type Memory struct {
	MemoryCaller     // Read-only binding to the contract
	MemoryTransactor // Write-only binding to the contract
	MemoryFilterer   // Log filterer for contract events
}

// MemoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type MemoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MemoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MemoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MemoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MemoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MemorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MemorySession struct {
	Contract     *Memory           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MemoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MemoryCallerSession struct {
	Contract *MemoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MemoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MemoryTransactorSession struct {
	Contract     *MemoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MemoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type MemoryRaw struct {
	Contract *Memory // Generic contract binding to access the raw methods on
}

// MemoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MemoryCallerRaw struct {
	Contract *MemoryCaller // Generic read-only contract binding to access the raw methods on
}

// MemoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MemoryTransactorRaw struct {
	Contract *MemoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMemory creates a new instance of Memory, bound to a specific deployed contract.
func NewMemory(address common.Address, backend bind.ContractBackend) (*Memory, error) {
	contract, err := bindMemory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Memory{MemoryCaller: MemoryCaller{contract: contract}, MemoryTransactor: MemoryTransactor{contract: contract}, MemoryFilterer: MemoryFilterer{contract: contract}}, nil
}

// NewMemoryCaller creates a new read-only instance of Memory, bound to a specific deployed contract.
func NewMemoryCaller(address common.Address, caller bind.ContractCaller) (*MemoryCaller, error) {
	contract, err := bindMemory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MemoryCaller{contract: contract}, nil
}

// NewMemoryTransactor creates a new write-only instance of Memory, bound to a specific deployed contract.
func NewMemoryTransactor(address common.Address, transactor bind.ContractTransactor) (*MemoryTransactor, error) {
	contract, err := bindMemory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MemoryTransactor{contract: contract}, nil
}

// NewMemoryFilterer creates a new log filterer instance of Memory, bound to a specific deployed contract.
func NewMemoryFilterer(address common.Address, filterer bind.ContractFilterer) (*MemoryFilterer, error) {
	contract, err := bindMemory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MemoryFilterer{contract: contract}, nil
}

// bindMemory binds a generic wrapper to an already deployed contract.
func bindMemory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MemoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Memory *MemoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Memory.Contract.MemoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Memory *MemoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Memory.Contract.MemoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Memory *MemoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Memory.Contract.MemoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Memory *MemoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Memory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Memory *MemoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Memory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Memory *MemoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Memory.Contract.contract.Transact(opts, method, params...)
}

// MimcUnitABI is the input ABI used to generate the binding from.
const MimcUnitABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"MiMCpe7\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// MimcUnitBin is the compiled bytecode used for deploying new contracts.
const MimcUnitBin = `0x6080604052348015600f57600080fd5b5060c28061001e6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c8063d15ca10914602d575b600080fd5b604d60048036036040811015604157600080fd5b5080359060200135605f565b60408051918252519081900360200190f35b60009291505056fea265627a7a7230582008a3090af2e01383eebe5d49bea9331133f336967239511652acae23fdb65d8e64736f6c637828302e352e31302d646576656c6f702e323031392e362e31332b636f6d6d69742e36363839373262620058`

// DeployMimcUnit deploys a new Ethereum contract, binding an instance of MimcUnit to it.
func DeployMimcUnit(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MimcUnit, error) {
	parsed, err := abi.JSON(strings.NewReader(MimcUnitABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MimcUnitBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MimcUnit{MimcUnitCaller: MimcUnitCaller{contract: contract}, MimcUnitTransactor: MimcUnitTransactor{contract: contract}, MimcUnitFilterer: MimcUnitFilterer{contract: contract}}, nil
}

// MimcUnit is an auto generated Go binding around an Ethereum contract.
type MimcUnit struct {
	MimcUnitCaller     // Read-only binding to the contract
	MimcUnitTransactor // Write-only binding to the contract
	MimcUnitFilterer   // Log filterer for contract events
}

// MimcUnitCaller is an auto generated read-only Go binding around an Ethereum contract.
type MimcUnitCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MimcUnitTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MimcUnitTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MimcUnitFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MimcUnitFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MimcUnitSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MimcUnitSession struct {
	Contract     *MimcUnit         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MimcUnitCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MimcUnitCallerSession struct {
	Contract *MimcUnitCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// MimcUnitTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MimcUnitTransactorSession struct {
	Contract     *MimcUnitTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// MimcUnitRaw is an auto generated low-level Go binding around an Ethereum contract.
type MimcUnitRaw struct {
	Contract *MimcUnit // Generic contract binding to access the raw methods on
}

// MimcUnitCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MimcUnitCallerRaw struct {
	Contract *MimcUnitCaller // Generic read-only contract binding to access the raw methods on
}

// MimcUnitTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MimcUnitTransactorRaw struct {
	Contract *MimcUnitTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMimcUnit creates a new instance of MimcUnit, bound to a specific deployed contract.
func NewMimcUnit(address common.Address, backend bind.ContractBackend) (*MimcUnit, error) {
	contract, err := bindMimcUnit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MimcUnit{MimcUnitCaller: MimcUnitCaller{contract: contract}, MimcUnitTransactor: MimcUnitTransactor{contract: contract}, MimcUnitFilterer: MimcUnitFilterer{contract: contract}}, nil
}

// NewMimcUnitCaller creates a new read-only instance of MimcUnit, bound to a specific deployed contract.
func NewMimcUnitCaller(address common.Address, caller bind.ContractCaller) (*MimcUnitCaller, error) {
	contract, err := bindMimcUnit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MimcUnitCaller{contract: contract}, nil
}

// NewMimcUnitTransactor creates a new write-only instance of MimcUnit, bound to a specific deployed contract.
func NewMimcUnitTransactor(address common.Address, transactor bind.ContractTransactor) (*MimcUnitTransactor, error) {
	contract, err := bindMimcUnit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MimcUnitTransactor{contract: contract}, nil
}

// NewMimcUnitFilterer creates a new log filterer instance of MimcUnit, bound to a specific deployed contract.
func NewMimcUnitFilterer(address common.Address, filterer bind.ContractFilterer) (*MimcUnitFilterer, error) {
	contract, err := bindMimcUnit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MimcUnitFilterer{contract: contract}, nil
}

// bindMimcUnit binds a generic wrapper to an already deployed contract.
func bindMimcUnit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MimcUnitABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MimcUnit *MimcUnitRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MimcUnit.Contract.MimcUnitCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MimcUnit *MimcUnitRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MimcUnit.Contract.MimcUnitTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MimcUnit *MimcUnitRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MimcUnit.Contract.MimcUnitTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MimcUnit *MimcUnitCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MimcUnit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MimcUnit *MimcUnitTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MimcUnit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MimcUnit *MimcUnitTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MimcUnit.Contract.contract.Transact(opts, method, params...)
}

// MiMCpe7 is a free data retrieval call binding the contract method 0xd15ca109.
//
// Solidity: function MiMCpe7(uint256 , uint256 ) constant returns(uint256)
func (_MimcUnit *MimcUnitCaller) MiMCpe7(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MimcUnit.contract.Call(opts, out, "MiMCpe7", arg0, arg1)
	return *ret0, err
}

// MiMCpe7 is a free data retrieval call binding the contract method 0xd15ca109.
//
// Solidity: function MiMCpe7(uint256 , uint256 ) constant returns(uint256)
func (_MimcUnit *MimcUnitSession) MiMCpe7(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _MimcUnit.Contract.MiMCpe7(&_MimcUnit.CallOpts, arg0, arg1)
}

// MiMCpe7 is a free data retrieval call binding the contract method 0xd15ca109.
//
// Solidity: function MiMCpe7(uint256 , uint256 ) constant returns(uint256)
func (_MimcUnit *MimcUnitCallerSession) MiMCpe7(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _MimcUnit.Contract.MiMCpe7(&_MimcUnit.CallOpts, arg0, arg1)
}
