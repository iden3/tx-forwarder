// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ZKPVerifier

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

// CheckFullCircuitABI is the input ABI used to generate the binding from.
const CheckFullCircuitABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"a\",\"type\":\"uint256[2]\"},{\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"name\":\"c\",\"type\":\"uint256[2]\"},{\"name\":\"input\",\"type\":\"uint256[10]\"}],\"name\":\"verify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"inputsKYC\",\"outputs\":[{\"name\":\"userEncPubKeyX\",\"type\":\"uint256\"},{\"name\":\"userEncPubKeyY\",\"type\":\"uint256\"},{\"name\":\"encUuidX\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"smVerifier\",\"type\":\"address\"},{\"name\":\"smRoots\",\"type\":\"address\"},{\"name\":\"smWhitelist\",\"type\":\"address\"},{\"name\":\"_idCertifier\",\"type\":\"bytes31\"},{\"name\":\"_idStorer\",\"type\":\"bytes31\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"userEncPubKeyX\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"userEncPubKeyY\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"encUuidX\",\"type\":\"uint256\"}],\"name\":\"registerId\",\"type\":\"event\"}]"

// CheckFullCircuitBin is the compiled bytecode used for deploying new contracts.
const CheckFullCircuitBin = `0x60806040527f274dbce8d15179969bc0d49fa725bddf9de555e0ba6a693c6adb52fc9ee7a82c6005557f05ce98c61b05f47fe2eae9a542bd99f6b2e78246231640b54595febfd51eb8536006557f1bddd2ee218e2a2eaf8dc4ed626280424a5424197f4e68ff030b21b86a6d8a8f6007557f0b1e3150c37f3349941416396465c569ed015a11bf4d2d98d632dac8f6f17f446008553480156100a057600080fd5b506040516109f43803806109f4833981810160405260a08110156100c357600080fd5b508051602082015160408301516060840151608090940151600180546001600160a01b039586166001600160a01b03199182161790915560008054948616948216949094179093556002805494909216939092169290921790915560038054600893841c7fff0000000000000000000000000000000000000000000000000000000000000091821617909155600480549290931c91161790556108898061016b6000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063636db9a71461003b5780639698a1721461013a575b600080fd5b610138600480360361024081101561005257600080fd5b6040805180820182529183019291818301918390600290839083908082843760009201829052506040805180820190915293969594608081019493509150600290835b828210156100d35760408051808201825290808402860190600290839083908082843760009201919091525050508152600190910190602001610095565b5050604080518082018252939695948181019493509150600290839083908082843760009201919091525050604080516101408181019092529295949381810193925090600a90839083908082843760009201919091525091945061017e9350505050565b005b6101606004803603602081101561015057600080fd5b50356001600160a01b0316610726565b60408051938452602084019290925282820152519081900360600190f35b6005546040820151146101c25760405162461bcd60e51b815260040180806020018281038252602c815260200180610804602c913960400191505060405180910390fd5b6006546060820151146102065760405162461bcd60e51b815260040180806020018281038252602c8152602001806107a2602c913960400191505060405180910390fd5b60075460808201511461024a5760405162461bcd60e51b81526004018080602001828103825260368152602001806107ce6036913960400191505060405180910390fd5b60085460a08201511461028e5760405162461bcd60e51b815260040180806020018281038252603681526020018061074b6036913960400191505060405180910390fd5b805160208201516101208301516000906102a790610747565b6000546003546040805163fead90d760e01b815260089290921b60ff19166004830152519293506001600160a01b039091169163fead90d791602480820192602092909190829003018186803b15801561030057600080fd5b505afa158015610314573d6000803e3d6000fd5b505050506040513d602081101561032a57600080fd5b5051831461037f576040805162461bcd60e51b815260206004820152601d60248201527f526f6f742063657274696669657220646f6573206e6f74206d61746368000000604482015290519081900360640190fd5b600054600480546040805163fead90d760e01b815260089290921b60ff19169282019290925290516001600160a01b039092169163fead90d791602480820192602092909190829003018186803b1580156103d957600080fd5b505afa1580156103ed573d6000803e3d6000fd5b505050506040513d602081101561040357600080fd5b50518214610458576040805162461bcd60e51b815260206004820152601a60248201527f526f6f742073746f72657220646f6573206e6f74206d61746368000000000000604482015290519081900360640190fd5b600154604080516379ddb87b60e11b81526001600160a01b039092169163f3bb70f6918a918a918a918a916004909101908190869080838360005b838110156104ab578181015183820152602001610493565b505050509050018460026000925b818410156104f95760208402830151604080838360005b838110156104e85781810151838201526020016104d0565b5050505090500192600101926104b9565b9250505083600260200280838360005b83811015610521578181015183820152602001610509565b5050505090500182600a60200280838360005b8381101561054c578181015183820152602001610534565b5050505090500194505050505060206040518083038186803b15801561057157600080fd5b505afa158015610585573d6000803e3d6000fd5b505050506040513d602081101561059b57600080fd5b505115156001146105dd5760405162461bcd60e51b81526004018080602001828103825260218152602001806107816021913960400191505060405180910390fd5b60025460408051632210724360e11b81526001600160a01b03848116600483015291519190921691634420e48691602480830192600092919082900301818387803b15801561062b57600080fd5b505af115801561063f573d6000803e3d6000fd5b505050506040518060600160405280856006600a811061065b57fe5b60200201518152602001856007600a811061067257fe5b60200201518152602001856008600a811061068957fe5b602090810291909101519091526001600160a01b038316600081815260098352604090819020845181558484015160018201559381015160029094019390935560c087015160e088015161010089015185519384529383019190915281840152606081019190915290517f4b762c5f9176ec51200505065b864bb94492e8957a741f45e31b832875395f999181900360800190a150505050505050565b60096020526000908152604090208054600182015460029092015490919083565b9056fe5920636f6f7264696e617465206c617720656e666f7263656d656e74207075626c6963206b657920646f6573206e6f74206d617463685a65726f2d6b776f776c656467652070726f6f66206973206e6f742076616c69645920636f6f7264696e6174652067656e65726174696f6e20706f696e7420646f6573206e6f74206d617463685820636f6f7264696e617465206c617720656e666f7263656d656e74207075626c6963206b657920646f6573206e6f74206d617463685820636f6f7264696e6174652067656e65726174696f6e20706f696e7420646f6573206e6f74206d61746368a265627a7a72315820b5e6e9fee0ae2aed466470b29adf64d3009c6aff5012067dd3196cdf766ed18064736f6c637827302e352e31312d646576656c6f702e323031392e372e332b636f6d6d69742e62383337373035320057`

// DeployCheckFullCircuit deploys a new Ethereum contract, binding an instance of CheckFullCircuit to it.
func DeployCheckFullCircuit(auth *bind.TransactOpts, backend bind.ContractBackend, smVerifier common.Address, smRoots common.Address, smWhitelist common.Address, _idCertifier [31]byte, _idStorer [31]byte) (common.Address, *types.Transaction, *CheckFullCircuit, error) {
	parsed, err := abi.JSON(strings.NewReader(CheckFullCircuitABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CheckFullCircuitBin), backend, smVerifier, smRoots, smWhitelist, _idCertifier, _idStorer)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CheckFullCircuit{CheckFullCircuitCaller: CheckFullCircuitCaller{contract: contract}, CheckFullCircuitTransactor: CheckFullCircuitTransactor{contract: contract}, CheckFullCircuitFilterer: CheckFullCircuitFilterer{contract: contract}}, nil
}

// CheckFullCircuit is an auto generated Go binding around an Ethereum contract.
type CheckFullCircuit struct {
	CheckFullCircuitCaller     // Read-only binding to the contract
	CheckFullCircuitTransactor // Write-only binding to the contract
	CheckFullCircuitFilterer   // Log filterer for contract events
}

// CheckFullCircuitCaller is an auto generated read-only Go binding around an Ethereum contract.
type CheckFullCircuitCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CheckFullCircuitTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CheckFullCircuitTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CheckFullCircuitFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CheckFullCircuitFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CheckFullCircuitSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CheckFullCircuitSession struct {
	Contract     *CheckFullCircuit // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CheckFullCircuitCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CheckFullCircuitCallerSession struct {
	Contract *CheckFullCircuitCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// CheckFullCircuitTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CheckFullCircuitTransactorSession struct {
	Contract     *CheckFullCircuitTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// CheckFullCircuitRaw is an auto generated low-level Go binding around an Ethereum contract.
type CheckFullCircuitRaw struct {
	Contract *CheckFullCircuit // Generic contract binding to access the raw methods on
}

// CheckFullCircuitCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CheckFullCircuitCallerRaw struct {
	Contract *CheckFullCircuitCaller // Generic read-only contract binding to access the raw methods on
}

// CheckFullCircuitTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CheckFullCircuitTransactorRaw struct {
	Contract *CheckFullCircuitTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCheckFullCircuit creates a new instance of CheckFullCircuit, bound to a specific deployed contract.
func NewCheckFullCircuit(address common.Address, backend bind.ContractBackend) (*CheckFullCircuit, error) {
	contract, err := bindCheckFullCircuit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CheckFullCircuit{CheckFullCircuitCaller: CheckFullCircuitCaller{contract: contract}, CheckFullCircuitTransactor: CheckFullCircuitTransactor{contract: contract}, CheckFullCircuitFilterer: CheckFullCircuitFilterer{contract: contract}}, nil
}

// NewCheckFullCircuitCaller creates a new read-only instance of CheckFullCircuit, bound to a specific deployed contract.
func NewCheckFullCircuitCaller(address common.Address, caller bind.ContractCaller) (*CheckFullCircuitCaller, error) {
	contract, err := bindCheckFullCircuit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CheckFullCircuitCaller{contract: contract}, nil
}

// NewCheckFullCircuitTransactor creates a new write-only instance of CheckFullCircuit, bound to a specific deployed contract.
func NewCheckFullCircuitTransactor(address common.Address, transactor bind.ContractTransactor) (*CheckFullCircuitTransactor, error) {
	contract, err := bindCheckFullCircuit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CheckFullCircuitTransactor{contract: contract}, nil
}

// NewCheckFullCircuitFilterer creates a new log filterer instance of CheckFullCircuit, bound to a specific deployed contract.
func NewCheckFullCircuitFilterer(address common.Address, filterer bind.ContractFilterer) (*CheckFullCircuitFilterer, error) {
	contract, err := bindCheckFullCircuit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CheckFullCircuitFilterer{contract: contract}, nil
}

// bindCheckFullCircuit binds a generic wrapper to an already deployed contract.
func bindCheckFullCircuit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CheckFullCircuitABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CheckFullCircuit *CheckFullCircuitRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CheckFullCircuit.Contract.CheckFullCircuitCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CheckFullCircuit *CheckFullCircuitRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CheckFullCircuit.Contract.CheckFullCircuitTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CheckFullCircuit *CheckFullCircuitRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CheckFullCircuit.Contract.CheckFullCircuitTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CheckFullCircuit *CheckFullCircuitCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CheckFullCircuit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CheckFullCircuit *CheckFullCircuitTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CheckFullCircuit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CheckFullCircuit *CheckFullCircuitTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CheckFullCircuit.Contract.contract.Transact(opts, method, params...)
}

// InputsKYC is a free data retrieval call binding the contract method 0x9698a172.
//
// Solidity: function inputsKYC(address ) constant returns(uint256 userEncPubKeyX, uint256 userEncPubKeyY, uint256 encUuidX)
func (_CheckFullCircuit *CheckFullCircuitCaller) InputsKYC(opts *bind.CallOpts, arg0 common.Address) (struct {
	UserEncPubKeyX *big.Int
	UserEncPubKeyY *big.Int
	EncUuidX       *big.Int
}, error) {
	ret := new(struct {
		UserEncPubKeyX *big.Int
		UserEncPubKeyY *big.Int
		EncUuidX       *big.Int
	})
	out := ret
	err := _CheckFullCircuit.contract.Call(opts, out, "inputsKYC", arg0)
	return *ret, err
}

// InputsKYC is a free data retrieval call binding the contract method 0x9698a172.
//
// Solidity: function inputsKYC(address ) constant returns(uint256 userEncPubKeyX, uint256 userEncPubKeyY, uint256 encUuidX)
func (_CheckFullCircuit *CheckFullCircuitSession) InputsKYC(arg0 common.Address) (struct {
	UserEncPubKeyX *big.Int
	UserEncPubKeyY *big.Int
	EncUuidX       *big.Int
}, error) {
	return _CheckFullCircuit.Contract.InputsKYC(&_CheckFullCircuit.CallOpts, arg0)
}

// InputsKYC is a free data retrieval call binding the contract method 0x9698a172.
//
// Solidity: function inputsKYC(address ) constant returns(uint256 userEncPubKeyX, uint256 userEncPubKeyY, uint256 encUuidX)
func (_CheckFullCircuit *CheckFullCircuitCallerSession) InputsKYC(arg0 common.Address) (struct {
	UserEncPubKeyX *big.Int
	UserEncPubKeyY *big.Int
	EncUuidX       *big.Int
}, error) {
	return _CheckFullCircuit.Contract.InputsKYC(&_CheckFullCircuit.CallOpts, arg0)
}

// Verify is a paid mutator transaction binding the contract method 0x636db9a7.
//
// Solidity: function verify(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[10] input) returns()
func (_CheckFullCircuit *CheckFullCircuitTransactor) Verify(opts *bind.TransactOpts, a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [10]*big.Int) (*types.Transaction, error) {
	return _CheckFullCircuit.contract.Transact(opts, "verify", a, b, c, input)
}

// Verify is a paid mutator transaction binding the contract method 0x636db9a7.
//
// Solidity: function verify(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[10] input) returns()
func (_CheckFullCircuit *CheckFullCircuitSession) Verify(a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [10]*big.Int) (*types.Transaction, error) {
	return _CheckFullCircuit.Contract.Verify(&_CheckFullCircuit.TransactOpts, a, b, c, input)
}

// Verify is a paid mutator transaction binding the contract method 0x636db9a7.
//
// Solidity: function verify(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[10] input) returns()
func (_CheckFullCircuit *CheckFullCircuitTransactorSession) Verify(a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [10]*big.Int) (*types.Transaction, error) {
	return _CheckFullCircuit.Contract.Verify(&_CheckFullCircuit.TransactOpts, a, b, c, input)
}

// CheckFullCircuitRegisterIdIterator is returned from FilterRegisterId and is used to iterate over the raw logs and unpacked data for RegisterId events raised by the CheckFullCircuit contract.
type CheckFullCircuitRegisterIdIterator struct {
	Event *CheckFullCircuitRegisterId // Event containing the contract specifics and raw log

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
func (it *CheckFullCircuitRegisterIdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CheckFullCircuitRegisterId)
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
		it.Event = new(CheckFullCircuitRegisterId)
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
func (it *CheckFullCircuitRegisterIdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CheckFullCircuitRegisterIdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CheckFullCircuitRegisterId represents a RegisterId event raised by the CheckFullCircuit contract.
type CheckFullCircuitRegisterId struct {
	Addr           common.Address
	UserEncPubKeyX *big.Int
	UserEncPubKeyY *big.Int
	EncUuidX       *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRegisterId is a free log retrieval operation binding the contract event 0x4b762c5f9176ec51200505065b864bb94492e8957a741f45e31b832875395f99.
//
// Solidity: event registerId(address addr, uint256 userEncPubKeyX, uint256 userEncPubKeyY, uint256 encUuidX)
func (_CheckFullCircuit *CheckFullCircuitFilterer) FilterRegisterId(opts *bind.FilterOpts) (*CheckFullCircuitRegisterIdIterator, error) {

	logs, sub, err := _CheckFullCircuit.contract.FilterLogs(opts, "registerId")
	if err != nil {
		return nil, err
	}
	return &CheckFullCircuitRegisterIdIterator{contract: _CheckFullCircuit.contract, event: "registerId", logs: logs, sub: sub}, nil
}

// WatchRegisterId is a free log subscription operation binding the contract event 0x4b762c5f9176ec51200505065b864bb94492e8957a741f45e31b832875395f99.
//
// Solidity: event registerId(address addr, uint256 userEncPubKeyX, uint256 userEncPubKeyY, uint256 encUuidX)
func (_CheckFullCircuit *CheckFullCircuitFilterer) WatchRegisterId(opts *bind.WatchOpts, sink chan<- *CheckFullCircuitRegisterId) (event.Subscription, error) {

	logs, sub, err := _CheckFullCircuit.contract.WatchLogs(opts, "registerId")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CheckFullCircuitRegisterId)
				if err := _CheckFullCircuit.contract.UnpackLog(event, "registerId", log); err != nil {
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

// Iden3HelpersABI is the input ABI used to generate the binding from.
const Iden3HelpersABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\"},{\"name\":\"rsv\",\"type\":\"bytes\"}],\"name\":\"checkSig\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes31\"}],\"name\":\"getRootFromId\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes27\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_mimcContractAddr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// Iden3HelpersBin is the compiled bytecode used for deploying new contracts.
const Iden3HelpersBin = `0x608060405234801561001057600080fd5b5060405161028a38038061028a8339818101604052602081101561003357600080fd5b5051600080546001600160a01b039092166001600160a01b0319909216919091179055610225806100656000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c806301b0452c1461003b578063ad05a8d214610104575b600080fd5b6100e86004803603604081101561005157600080fd5b8135919081019060408101602082013564010000000081111561007357600080fd5b82018360208201111561008557600080fd5b803590602001918460018302840111640100000000831117156100a757600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610140945050505050565b604080516001600160a01b039092168252519081900360200190f35b6101256004803603602081101561011a57600080fd5b503560ff19166101bf565b6040805164ffffffffff199092168252519081900360200190f35b602081810151604080840151606080860151835160008082528188018087528a905291821a81860181905292810186905260808101849052935190959293919260019260a080820193601f1981019281900390910190855afa1580156101aa573d6000803e3d6000fd5b5050604051601f190151979650505050505050565b60101b62ffffff19169056fea265627a7a72315820c1f255a492be8557dc33d4d341e1b2f6f102706757b75a1e74708b4082b3a27164736f6c637827302e352e31312d646576656c6f702e323031392e372e332b636f6d6d69742e62383337373035320057`

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
const MemoryBin = `0x607a6023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158201e614cd2ee03f6184836e343dc1fa444157aa41ac743b5c8f81502e0890ea8c964736f6c637827302e352e31312d646576656c6f702e323031392e372e332b636f6d6d69742e62383337373035320057`

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
const MimcUnitBin = `0x6080604052348015600f57600080fd5b5060c18061001e6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c8063d15ca10914602d575b600080fd5b604d60048036036040811015604157600080fd5b5080359060200135605f565b60408051918252519081900360200190f35b60009291505056fea265627a7a7231582041aebd24b0bc315a2f510d55a61bc713a15fe278de93029e3fd617bae8bc1cd764736f6c637827302e352e31312d646576656c6f702e323031392e372e332b636f6d6d69742e62383337373035320057`

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

// PairingABI is the input ABI used to generate the binding from.
const PairingABI = "[]"

// PairingBin is the compiled bytecode used for deploying new contracts.
const PairingBin = `0x607a6023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a7231582035644aabf99891c3330143141fc8860cde4c3ac2e075f6a43b93d5435078e90c64736f6c637827302e352e31312d646576656c6f702e323031392e372e332b636f6d6d69742e62383337373035320057`

// DeployPairing deploys a new Ethereum contract, binding an instance of Pairing to it.
func DeployPairing(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Pairing, error) {
	parsed, err := abi.JSON(strings.NewReader(PairingABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PairingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Pairing{PairingCaller: PairingCaller{contract: contract}, PairingTransactor: PairingTransactor{contract: contract}, PairingFilterer: PairingFilterer{contract: contract}}, nil
}

// Pairing is an auto generated Go binding around an Ethereum contract.
type Pairing struct {
	PairingCaller     // Read-only binding to the contract
	PairingTransactor // Write-only binding to the contract
	PairingFilterer   // Log filterer for contract events
}

// PairingCaller is an auto generated read-only Go binding around an Ethereum contract.
type PairingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PairingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PairingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PairingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PairingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PairingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PairingSession struct {
	Contract     *Pairing          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PairingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PairingCallerSession struct {
	Contract *PairingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// PairingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PairingTransactorSession struct {
	Contract     *PairingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// PairingRaw is an auto generated low-level Go binding around an Ethereum contract.
type PairingRaw struct {
	Contract *Pairing // Generic contract binding to access the raw methods on
}

// PairingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PairingCallerRaw struct {
	Contract *PairingCaller // Generic read-only contract binding to access the raw methods on
}

// PairingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PairingTransactorRaw struct {
	Contract *PairingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPairing creates a new instance of Pairing, bound to a specific deployed contract.
func NewPairing(address common.Address, backend bind.ContractBackend) (*Pairing, error) {
	contract, err := bindPairing(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pairing{PairingCaller: PairingCaller{contract: contract}, PairingTransactor: PairingTransactor{contract: contract}, PairingFilterer: PairingFilterer{contract: contract}}, nil
}

// NewPairingCaller creates a new read-only instance of Pairing, bound to a specific deployed contract.
func NewPairingCaller(address common.Address, caller bind.ContractCaller) (*PairingCaller, error) {
	contract, err := bindPairing(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PairingCaller{contract: contract}, nil
}

// NewPairingTransactor creates a new write-only instance of Pairing, bound to a specific deployed contract.
func NewPairingTransactor(address common.Address, transactor bind.ContractTransactor) (*PairingTransactor, error) {
	contract, err := bindPairing(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PairingTransactor{contract: contract}, nil
}

// NewPairingFilterer creates a new log filterer instance of Pairing, bound to a specific deployed contract.
func NewPairingFilterer(address common.Address, filterer bind.ContractFilterer) (*PairingFilterer, error) {
	contract, err := bindPairing(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PairingFilterer{contract: contract}, nil
}

// bindPairing binds a generic wrapper to an already deployed contract.
func bindPairing(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PairingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pairing *PairingRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Pairing.Contract.PairingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pairing *PairingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pairing.Contract.PairingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pairing *PairingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pairing.Contract.PairingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pairing *PairingCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Pairing.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pairing *PairingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pairing.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pairing *PairingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pairing.Contract.contract.Transact(opts, method, params...)
}

// RootCommitsABI is the input ABI used to generate the binding from.
const RootCommitsABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\"},{\"name\":\"rsv\",\"type\":\"bytes\"}],\"name\":\"checkSig\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes31\"},{\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"getRootByTime\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes31\"}],\"name\":\"getRootFromId\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes27\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes31\"},{\"name\":\"blockN\",\"type\":\"uint64\"}],\"name\":\"getRootByBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newRoot\",\"type\":\"bytes32\"},{\"name\":\"id\",\"type\":\"bytes31\"},{\"name\":\"mtp\",\"type\":\"bytes\"}],\"name\":\"setRoot\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes31\"}],\"name\":\"getRoot\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_mimcContractAddr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes31\"},{\"indexed\":false,\"name\":\"blockN\",\"type\":\"uint64\"},{\"indexed\":false,\"name\":\"timestamp\",\"type\":\"uint64\"},{\"indexed\":false,\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"RootUpdated\",\"type\":\"event\"}]"

// RootCommitsBin is the compiled bytecode used for deploying new contracts.
const RootCommitsBin = `0x608060405234801561001057600080fd5b5060405161122e38038061122e8339818101604052602081101561003357600080fd5b5051600080546001600160a01b039092166001600160a01b03199092169190911790556111c9806100656000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c806301b0452c146100675780634175dae514610130578063ad05a8d214610172578063b816ff6f146101ae578063e0681acd146101de578063fead90d714610296575b600080fd5b6101146004803603604081101561007d57600080fd5b8135919081019060408101602082013564010000000081111561009f57600080fd5b8201836020820111156100b157600080fd5b803590602001918460018302840111640100000000831117156100d357600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506102b7945050505050565b604080516001600160a01b039092168252519081900360200190f35b6101606004803603604081101561014657600080fd5b50803560ff191690602001356001600160401b0316610337565b60408051918252519081900360200190f35b6101936004803603602081101561018857600080fd5b503560ff191661061c565b6040805164ffffffffff199092168252519081900360200190f35b610160600480360360408110156101c457600080fd5b50803560ff191690602001356001600160401b031661062d565b610294600480360360608110156101f457600080fd5b81359160ff196020820135169181019060608101604082013564010000000081111561021f57600080fd5b82018360208201111561023157600080fd5b8035906020019184600183028401116401000000008311171561025357600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506108ab945050505050565b005b610160600480360360208110156102ac57600080fd5b503560ff1916610b48565b602081810151604080840151606080860151835160008082528188018087528a905291821a81860181905292810186905260808101849052935190959293919260019260a080820193601f1981019281900390910190855afa158015610321573d6000803e3d6000fd5b5050506020604051035193505050505b92915050565b600042826001600160401b03161061038b576040805162461bcd60e51b8152602060048201526012602482015271195c9c939bd19d5d1d5c99505b1b1bddd95960721b604482015290519081900360640190fd5b60ff1983166000908152600160205260409020546103ac5750600254610331565b60ff1983166000908152600160205260408120805460001981019081106103cf57fe5b60009182526020909120600290910201546001600160401b03600160401b90910481169150831681101561043a5760ff19841660009081526001602052604090208054600019810190811061042057fe5b906000526020600020906002020160010154915050610331565b60ff198416600090815260016020526040812054600019015b80821161060f5760ff19861660009081526001602052604090208054600284840104916001600160401b038816918390811061048b57fe5b6000918252602090912060029091020154600160401b90046001600160401b031614156104ee5760ff19871660009081526001602052604090208054829081106104d157fe5b906000526020600020906002020160010154945050505050610331565b60ff198716600090815260016020526040902080548290811061050d57fe5b60009182526020909120600290910201546001600160401b03600160401b9091048116908716118015610588575060ff198716600090815260016020819052604090912080549091830190811061056057fe5b60009182526020909120600290910201546001600160401b03600160401b9091048116908716105b156105ac5760ff19871660009081526001602052604090208054829081106104d157fe5b60ff19871660009081526001602052604090208054829081106105cb57fe5b60009182526020909120600290910201546001600160401b03600160401b9091048116908716111561060257806001019250610609565b6001810391505b50610453565b5050600254949350505050565b62ffffff19601082901b165b919050565b600043826001600160401b031610610681576040805162461bcd60e51b8152602060048201526012602482015271195c9c939bd19d5d1d5c99505b1b1bddd95960721b604482015290519081900360640190fd5b60ff1983166000908152600160205260409020546106a25750600254610331565b60ff1983166000908152600160205260408120805460001981019081106106c557fe5b60009182526020909120600290910201546001600160401b039081169150831681101561070f5760ff19841660009081526001602052604090208054600019810190811061042057fe5b60ff198416600090815260016020526040812054600019015b80821161060f5760ff19861660009081526001602052604090208054600284840104916001600160401b038816918390811061076057fe5b60009182526020909120600290910201546001600160401b0316141561079f5760ff19871660009081526001602052604090208054829081106104d157fe5b60ff19871660009081526001602052604090208054829081106107be57fe5b60009182526020909120600290910201546001600160401b0390811690871611801561082b575060ff198716600090815260016020819052604090912080549091830190811061080a57fe5b60009182526020909120600290910201546001600160401b03908116908716105b1561084f5760ff19871660009081526001602052604090208054829081106104d157fe5b60ff198716600090815260016020526040902080548290811061086e57fe5b60009182526020909120600290910201546001600160401b03908116908716111561089e578060010192506108a5565b6001810391505b50610728565b6108b36110d8565b6108bc33610ba5565b90506108c66110d8565b6108cf82610be3565b90506000806108dd83610c3b565b909250905060006108ed8761061c565b90506108fd81878585608c610c74565b1515600114610953576040805162461bcd60e51b815260206004820152601b60248201527f4d65726b6c6520747265652070726f6f66206e6f742076616c69640000000000604482015290519081900360640190fd5b60ff198716600090815260016020526040902054156109e95760ff19871660009081526001602052604090208054439190600019810190811061099257fe5b60009182526020909120600290910201546001600160401b031614156109e95760405162461bcd60e51b815260040180806020018281038252602181526020018061114f6021913960400191505060405180910390fd5b600160008860ff191660ff191681526020019081526020016000206040518060600160405280436001600160401b03168152602001426001600160401b031681526020018a8152509080600181540180825580915050906001820390600052602060002090600202016000909192909190915060008201518160000160006101000a8154816001600160401b0302191690836001600160401b0316021790555060208201518160000160086101000a8154816001600160401b0302191690836001600160401b03160217905550604082015181600101555050507fc26526f78af19d7325e758deea00a04b0db26e055a576222266dfe241755a57f8743428b604051808560ff191660ff19168152602001846001600160401b03166001600160401b03168152602001836001600160401b03166001600160401b0316815260200182815260200194505050505060405180910390a15050505050505050565b60ff198116600090815260016020526040812054610b695750600254610628565b60ff198216600090815260016020526040902080546000198101908110610b8c57fe5b9060005260206000209060020201600101549050919050565b610bad6110d8565b600960c01b815260006020820152606091821b6bffffffffffffffffffffffff19166040820152600360e01b9181019190915290565b610beb6110d8565b815160c01c6060828101828152602085015160a01c6bffffffff000000000000000016909217909152604080840151821c8184018181529290940151901c63ffffffff60a01b1690921790915290565b600080610c58836040015160001c846060015160001c6000610e78565b83516020850151919350610c6d916000610e78565b9050915091565b6000610c7e6110ff565b610c8786610ef2565b9050600081600001518015610c9d575081602001515b15610d06576000805b81158015610cb357508581105b15610cdb578088901c600116818560800151901c6001161860ff169150600181019050610ca6565b81610ced576000945050505050610e6f565b610d0184608001518560a001516001610e78565b925050505b6060826040015160ff16604051908082528060200260200182016040528015610d39578160200160208202803883390190505b509050600080805b856040015160ff16811015610dca5760608601516001600160f01b0316811c600190811690811415610da65760008360200260400160ff169050808d0151945084868481518110610d8e57fe5b60200260200101818152505060018401935050610dc1565b6000858381518110610db457fe5b6020026020010181815250505b50600101610d41565b508451600090610de557610de08a8a6001610e78565b610de7565b845b9050600080600188604001510360ff1690505b858181518110610e0657fe5b6020908102919091010151915060018c821c81161480610e3157610e2c84846000610e78565b610e3d565b610e3d83856000610e78565b935081610e4a5750610e54565b5060001901610dfa565b505060281b64ffffffffff19908116908c1614955050505050505b95945050505050565b60408051600280825260608083018452600093909291906020830190803883390190505090508481600081518110610eac57fe5b6020026020010181815250508381600181518110610ec657fe5b60200260200101818152505082610ee757610ee2816000610f81565b610e6f565b610e6f816001610f81565b610efa6110ff565b610f02611134565b610f0b83610f94565b90506000610f1882610fbd565b60018082168114855281811c81161460208501529050610f3782610fbd565b60ff166040840152610f4882610fcf565b6001600160f01b03166060840152602083015115610f7a5783518401601f198101519051608085019190915260a08401525b5050919050565b6000610f8d8383610fe1565b9392505050565b610f9c611134565b50604080518082019091526020828101825282518301810190820152919050565b80518051600190910190915260f81c90565b80518051601e90910190915260101c90565b6000817f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001825b85518110156110ce57600054865183916001600160a01b03169063d15ca1099089908590811061103357fe5b6020026020010151866040518363ffffffff1660e01b8152600401808381526020018281526020019250505060206040518083038186803b15801561107757600080fd5b505afa15801561108b573d6000803e3d6000fd5b505050506040513d60208110156110a157600080fd5b505187518890849081106110b157fe5b6020026020010151850101816110c357fe5b069250600101611007565b5090949350505050565b60408051608081018252600080825260208201819052918101829052606081019190915290565b6040805160c081018252600080825260208201819052918101829052606081018290526080810182905260a081019190915290565b60405180604001604052806000815260200160008152509056fe6e6f206d756c7469706c652073657420696e207468652073616d6520626c6f636ba265627a7a72315820801b6cb6ca6cf5b9d21203270d7763349693c2d0fa817f9665997104f191401a64736f6c637827302e352e31312d646576656c6f702e323031392e372e332b636f6d6d69742e62383337373035320057`

// DeployRootCommits deploys a new Ethereum contract, binding an instance of RootCommits to it.
func DeployRootCommits(auth *bind.TransactOpts, backend bind.ContractBackend, _mimcContractAddr common.Address) (common.Address, *types.Transaction, *RootCommits, error) {
	parsed, err := abi.JSON(strings.NewReader(RootCommitsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RootCommitsBin), backend, _mimcContractAddr)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RootCommits{RootCommitsCaller: RootCommitsCaller{contract: contract}, RootCommitsTransactor: RootCommitsTransactor{contract: contract}, RootCommitsFilterer: RootCommitsFilterer{contract: contract}}, nil
}

// RootCommits is an auto generated Go binding around an Ethereum contract.
type RootCommits struct {
	RootCommitsCaller     // Read-only binding to the contract
	RootCommitsTransactor // Write-only binding to the contract
	RootCommitsFilterer   // Log filterer for contract events
}

// RootCommitsCaller is an auto generated read-only Go binding around an Ethereum contract.
type RootCommitsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootCommitsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RootCommitsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootCommitsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RootCommitsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootCommitsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RootCommitsSession struct {
	Contract     *RootCommits      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RootCommitsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RootCommitsCallerSession struct {
	Contract *RootCommitsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// RootCommitsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RootCommitsTransactorSession struct {
	Contract     *RootCommitsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// RootCommitsRaw is an auto generated low-level Go binding around an Ethereum contract.
type RootCommitsRaw struct {
	Contract *RootCommits // Generic contract binding to access the raw methods on
}

// RootCommitsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RootCommitsCallerRaw struct {
	Contract *RootCommitsCaller // Generic read-only contract binding to access the raw methods on
}

// RootCommitsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RootCommitsTransactorRaw struct {
	Contract *RootCommitsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRootCommits creates a new instance of RootCommits, bound to a specific deployed contract.
func NewRootCommits(address common.Address, backend bind.ContractBackend) (*RootCommits, error) {
	contract, err := bindRootCommits(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RootCommits{RootCommitsCaller: RootCommitsCaller{contract: contract}, RootCommitsTransactor: RootCommitsTransactor{contract: contract}, RootCommitsFilterer: RootCommitsFilterer{contract: contract}}, nil
}

// NewRootCommitsCaller creates a new read-only instance of RootCommits, bound to a specific deployed contract.
func NewRootCommitsCaller(address common.Address, caller bind.ContractCaller) (*RootCommitsCaller, error) {
	contract, err := bindRootCommits(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RootCommitsCaller{contract: contract}, nil
}

// NewRootCommitsTransactor creates a new write-only instance of RootCommits, bound to a specific deployed contract.
func NewRootCommitsTransactor(address common.Address, transactor bind.ContractTransactor) (*RootCommitsTransactor, error) {
	contract, err := bindRootCommits(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RootCommitsTransactor{contract: contract}, nil
}

// NewRootCommitsFilterer creates a new log filterer instance of RootCommits, bound to a specific deployed contract.
func NewRootCommitsFilterer(address common.Address, filterer bind.ContractFilterer) (*RootCommitsFilterer, error) {
	contract, err := bindRootCommits(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RootCommitsFilterer{contract: contract}, nil
}

// bindRootCommits binds a generic wrapper to an already deployed contract.
func bindRootCommits(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RootCommitsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RootCommits *RootCommitsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RootCommits.Contract.RootCommitsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RootCommits *RootCommitsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootCommits.Contract.RootCommitsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RootCommits *RootCommitsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RootCommits.Contract.RootCommitsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RootCommits *RootCommitsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RootCommits.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RootCommits *RootCommitsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootCommits.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RootCommits *RootCommitsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RootCommits.Contract.contract.Transact(opts, method, params...)
}

// CheckSig is a free data retrieval call binding the contract method 0x01b0452c.
//
// Solidity: function checkSig(bytes32 msgHash, bytes rsv) constant returns(address)
func (_RootCommits *RootCommitsCaller) CheckSig(opts *bind.CallOpts, msgHash [32]byte, rsv []byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RootCommits.contract.Call(opts, out, "checkSig", msgHash, rsv)
	return *ret0, err
}

// CheckSig is a free data retrieval call binding the contract method 0x01b0452c.
//
// Solidity: function checkSig(bytes32 msgHash, bytes rsv) constant returns(address)
func (_RootCommits *RootCommitsSession) CheckSig(msgHash [32]byte, rsv []byte) (common.Address, error) {
	return _RootCommits.Contract.CheckSig(&_RootCommits.CallOpts, msgHash, rsv)
}

// CheckSig is a free data retrieval call binding the contract method 0x01b0452c.
//
// Solidity: function checkSig(bytes32 msgHash, bytes rsv) constant returns(address)
func (_RootCommits *RootCommitsCallerSession) CheckSig(msgHash [32]byte, rsv []byte) (common.Address, error) {
	return _RootCommits.Contract.CheckSig(&_RootCommits.CallOpts, msgHash, rsv)
}

// GetRoot is a free data retrieval call binding the contract method 0xfead90d7.
//
// Solidity: function getRoot(bytes31 id) constant returns(bytes32)
func (_RootCommits *RootCommitsCaller) GetRoot(opts *bind.CallOpts, id [31]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _RootCommits.contract.Call(opts, out, "getRoot", id)
	return *ret0, err
}

// GetRoot is a free data retrieval call binding the contract method 0xfead90d7.
//
// Solidity: function getRoot(bytes31 id) constant returns(bytes32)
func (_RootCommits *RootCommitsSession) GetRoot(id [31]byte) ([32]byte, error) {
	return _RootCommits.Contract.GetRoot(&_RootCommits.CallOpts, id)
}

// GetRoot is a free data retrieval call binding the contract method 0xfead90d7.
//
// Solidity: function getRoot(bytes31 id) constant returns(bytes32)
func (_RootCommits *RootCommitsCallerSession) GetRoot(id [31]byte) ([32]byte, error) {
	return _RootCommits.Contract.GetRoot(&_RootCommits.CallOpts, id)
}

// GetRootByBlock is a free data retrieval call binding the contract method 0xb816ff6f.
//
// Solidity: function getRootByBlock(bytes31 id, uint64 blockN) constant returns(bytes32)
func (_RootCommits *RootCommitsCaller) GetRootByBlock(opts *bind.CallOpts, id [31]byte, blockN uint64) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _RootCommits.contract.Call(opts, out, "getRootByBlock", id, blockN)
	return *ret0, err
}

// GetRootByBlock is a free data retrieval call binding the contract method 0xb816ff6f.
//
// Solidity: function getRootByBlock(bytes31 id, uint64 blockN) constant returns(bytes32)
func (_RootCommits *RootCommitsSession) GetRootByBlock(id [31]byte, blockN uint64) ([32]byte, error) {
	return _RootCommits.Contract.GetRootByBlock(&_RootCommits.CallOpts, id, blockN)
}

// GetRootByBlock is a free data retrieval call binding the contract method 0xb816ff6f.
//
// Solidity: function getRootByBlock(bytes31 id, uint64 blockN) constant returns(bytes32)
func (_RootCommits *RootCommitsCallerSession) GetRootByBlock(id [31]byte, blockN uint64) ([32]byte, error) {
	return _RootCommits.Contract.GetRootByBlock(&_RootCommits.CallOpts, id, blockN)
}

// GetRootByTime is a free data retrieval call binding the contract method 0x4175dae5.
//
// Solidity: function getRootByTime(bytes31 id, uint64 timestamp) constant returns(bytes32)
func (_RootCommits *RootCommitsCaller) GetRootByTime(opts *bind.CallOpts, id [31]byte, timestamp uint64) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _RootCommits.contract.Call(opts, out, "getRootByTime", id, timestamp)
	return *ret0, err
}

// GetRootByTime is a free data retrieval call binding the contract method 0x4175dae5.
//
// Solidity: function getRootByTime(bytes31 id, uint64 timestamp) constant returns(bytes32)
func (_RootCommits *RootCommitsSession) GetRootByTime(id [31]byte, timestamp uint64) ([32]byte, error) {
	return _RootCommits.Contract.GetRootByTime(&_RootCommits.CallOpts, id, timestamp)
}

// GetRootByTime is a free data retrieval call binding the contract method 0x4175dae5.
//
// Solidity: function getRootByTime(bytes31 id, uint64 timestamp) constant returns(bytes32)
func (_RootCommits *RootCommitsCallerSession) GetRootByTime(id [31]byte, timestamp uint64) ([32]byte, error) {
	return _RootCommits.Contract.GetRootByTime(&_RootCommits.CallOpts, id, timestamp)
}

// GetRootFromId is a free data retrieval call binding the contract method 0xad05a8d2.
//
// Solidity: function getRootFromId(bytes31 id) constant returns(bytes27)
func (_RootCommits *RootCommitsCaller) GetRootFromId(opts *bind.CallOpts, id [31]byte) ([27]byte, error) {
	var (
		ret0 = new([27]byte)
	)
	out := ret0
	err := _RootCommits.contract.Call(opts, out, "getRootFromId", id)
	return *ret0, err
}

// GetRootFromId is a free data retrieval call binding the contract method 0xad05a8d2.
//
// Solidity: function getRootFromId(bytes31 id) constant returns(bytes27)
func (_RootCommits *RootCommitsSession) GetRootFromId(id [31]byte) ([27]byte, error) {
	return _RootCommits.Contract.GetRootFromId(&_RootCommits.CallOpts, id)
}

// GetRootFromId is a free data retrieval call binding the contract method 0xad05a8d2.
//
// Solidity: function getRootFromId(bytes31 id) constant returns(bytes27)
func (_RootCommits *RootCommitsCallerSession) GetRootFromId(id [31]byte) ([27]byte, error) {
	return _RootCommits.Contract.GetRootFromId(&_RootCommits.CallOpts, id)
}

// SetRoot is a paid mutator transaction binding the contract method 0xe0681acd.
//
// Solidity: function setRoot(bytes32 newRoot, bytes31 id, bytes mtp) returns()
func (_RootCommits *RootCommitsTransactor) SetRoot(opts *bind.TransactOpts, newRoot [32]byte, id [31]byte, mtp []byte) (*types.Transaction, error) {
	return _RootCommits.contract.Transact(opts, "setRoot", newRoot, id, mtp)
}

// SetRoot is a paid mutator transaction binding the contract method 0xe0681acd.
//
// Solidity: function setRoot(bytes32 newRoot, bytes31 id, bytes mtp) returns()
func (_RootCommits *RootCommitsSession) SetRoot(newRoot [32]byte, id [31]byte, mtp []byte) (*types.Transaction, error) {
	return _RootCommits.Contract.SetRoot(&_RootCommits.TransactOpts, newRoot, id, mtp)
}

// SetRoot is a paid mutator transaction binding the contract method 0xe0681acd.
//
// Solidity: function setRoot(bytes32 newRoot, bytes31 id, bytes mtp) returns()
func (_RootCommits *RootCommitsTransactorSession) SetRoot(newRoot [32]byte, id [31]byte, mtp []byte) (*types.Transaction, error) {
	return _RootCommits.Contract.SetRoot(&_RootCommits.TransactOpts, newRoot, id, mtp)
}

// RootCommitsRootUpdatedIterator is returned from FilterRootUpdated and is used to iterate over the raw logs and unpacked data for RootUpdated events raised by the RootCommits contract.
type RootCommitsRootUpdatedIterator struct {
	Event *RootCommitsRootUpdated // Event containing the contract specifics and raw log

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
func (it *RootCommitsRootUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootCommitsRootUpdated)
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
		it.Event = new(RootCommitsRootUpdated)
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
func (it *RootCommitsRootUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootCommitsRootUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootCommitsRootUpdated represents a RootUpdated event raised by the RootCommits contract.
type RootCommitsRootUpdated struct {
	Id        [31]byte
	BlockN    uint64
	Timestamp uint64
	Root      [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRootUpdated is a free log retrieval operation binding the contract event 0xc26526f78af19d7325e758deea00a04b0db26e055a576222266dfe241755a57f.
//
// Solidity: event RootUpdated(bytes31 id, uint64 blockN, uint64 timestamp, bytes32 root)
func (_RootCommits *RootCommitsFilterer) FilterRootUpdated(opts *bind.FilterOpts) (*RootCommitsRootUpdatedIterator, error) {

	logs, sub, err := _RootCommits.contract.FilterLogs(opts, "RootUpdated")
	if err != nil {
		return nil, err
	}
	return &RootCommitsRootUpdatedIterator{contract: _RootCommits.contract, event: "RootUpdated", logs: logs, sub: sub}, nil
}

// WatchRootUpdated is a free log subscription operation binding the contract event 0xc26526f78af19d7325e758deea00a04b0db26e055a576222266dfe241755a57f.
//
// Solidity: event RootUpdated(bytes31 id, uint64 blockN, uint64 timestamp, bytes32 root)
func (_RootCommits *RootCommitsFilterer) WatchRootUpdated(opts *bind.WatchOpts, sink chan<- *RootCommitsRootUpdated) (event.Subscription, error) {

	logs, sub, err := _RootCommits.contract.WatchLogs(opts, "RootUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootCommitsRootUpdated)
				if err := _RootCommits.contract.UnpackLog(event, "RootUpdated", log); err != nil {
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

// VerifierABI is the input ABI used to generate the binding from.
const VerifierABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"a\",\"type\":\"uint256[2]\"},{\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"name\":\"c\",\"type\":\"uint256[2]\"},{\"name\":\"input\",\"type\":\"uint256[10]\"}],\"name\":\"verifyProof\",\"outputs\":[{\"name\":\"r\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// VerifierBin is the compiled bytecode used for deploying new contracts.
const VerifierBin = `0x608060405234801561001057600080fd5b506110e6806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063f3bb70f614610030575b600080fd5b61012d600480360361024081101561004757600080fd5b6040805180820182529183019291818301918390600290839083908082843760009201829052506040805180820190915293969594608081019493509150600290835b828210156100c8576040805180820182529080840286019060029083908390808284376000920191909152505050815260019091019060200161008a565b5050604080518082018252939695948181019493509150600290839083908082843760009201919091525050604080516101408181019092529295949381810193925090600a9083908390808284376000920191909152509194506101419350505050565b604080519115158252519081900360200190f35b600061014b610f5f565b6040805180820182528751815260208089015181830152908352815160808101835287515181840190815288518301516060808401919091529082528351808501855289840180515182525184015181850152828401528483019190915282518084018452875181528783015181840152848401528251600a80825261016082019094529092909182016101408038833901905050905060005b600a81101561021e578481600a81106101fa57fe5b602002015182828151811061020b57fe5b60209081029190910101526001016101e5565b506102298183610247565b6102385760019250505061023f565b6000925050505b949350505050565b6000610251610f92565b610259610358565b905080608001515184516001011461027057600080fd5b610278610fda565b50604080518082019091526000808252602082018190525b85518110156102e3576102d9826102d4856080015184600101815181106102b357fe5b60200260200101518985815181106102c757fe5b6020026020010151610ad5565b610b2d565b9150600101610290565b506103068183608001516000815181106102f957fe5b6020026020010151610b2d565b905061033c6103188560000151610b7b565b8560200151846000015185602001518587604001518a604001518960600151610c07565b61034b57600192505050610352565b6000925050505b92915050565b610360610f92565b6040805180820182527f2614b958e596d9690f5d3217af600404ea44eb7dd3aa13e36fdd7a87700ee4f881527f05ff6d4734079449005efcfedb3aee7576c590bd639d2cce21f65886b8634faa6020808301919091529083528151608080820184527f01720d524f0fc778bbde14c7b822e64949bf00279f12bbcfc50705824a64dad88285019081527f0669e94bdbe1ea368aa1ab27fee5dc0c5a373bf517ce52d8b662545beafff4ab606080850191909152908352845180860186527f0b28c4152fe2becbc836985f8ad909d777bbece88d496e07f12b90f627c006ba81527f1521166f34a9a17825359aa95ea8048b3511ec2cb6b5d785790be73fc6bba1e3818601528385015285840192909252835180820185527f199d5edac0e49cff949ef3d575f165a0eaaf375055cc688080a4c59347a6d71f8186019081527f0e4e11831095dc06701e66848e73b561de0b884174c5450c6f20979311351de3828501528152845180860186527f1fba33820ddec168c9ca62558e898f5fdce2b41eee44545fa6e11fd294aa69b781527f204d6f8abc7dd042ed457bdcee4e75b1cce9ae5314bd82db5e96e771e17754d6818601528185015285850152835190810184527f0c3f7a9eab611dea28bf032973a27402eb22c3ec6836ced8c87415bcec0d949d8185019081527f295ee9899095c28912781a515d23ae2bb0ec6158675342b28bcbf8559a15c1bf828401528152835180850185527f1f1aaeb4d95088adfeec589e094d56fd7490e9f5e326c3ce39adb9419e2fb8dc81527f026c18a2d69f2346717cc9da2fce9454ecf9f4d7cc054ac417b8b43a54c56b708185015281840152908401528151600b8082526101808201909352919082015b6105e4610fda565b8152602001906001900390816105dc57505060808201908152604080518082019091527f14cad51fa47c563e7e295318e071cec1ce02028b5b157681b7093a6fa31aa91e81527f2f1518dd4e4564fec533dcfb6df43794214e9d5a5a3b0c76a5da1b426056d0c160208201529051805160009061065d57fe5b602002602001018190525060405180604001604052807f08316ff5fa7c62df4e7f84b3a1d6dfc18dffc72f1e5146005424591ececa426581526020017f0b4e3e9f78525e412fa03a75cdfd28e3c6aaf3ffdda151e444e3de3910a9814481525081608001516001815181106106ce57fe5b602002602001018190525060405180604001604052807f092254d5ce850404029f14b8b7c96890185f19edf3ad490aa5872fdd29b41d0881526020017f168eed5e0b4ac1fc0baa328477427807e985dbe9a9145a5e4335c6d6d0205ec4815250816080015160028151811061073f57fe5b602002602001018190525060405180604001604052807f019ddba6b4f6877b647b3dc0fdc6c5628539a38c4e7187a38246bdc4c47551f981526020017f14f3662a55eb0ad31223324899d0c7a03825556bf82f1d66d2a7fe85f1be613381525081608001516003815181106107b057fe5b602002602001018190525060405180604001604052807f08fd17f1714f6db6c44d8a619fce65a9b78cd3ad9f570428d26de6764b49a5b081526020017f2d070656c2816ac138c7fb8b92c1b82d08312d3284c2f61fbc10090aa5e1186e815250816080015160048151811061082157fe5b602002602001018190525060405180604001604052807f2ad2142b5edbde99bb8d08f37726ea44b323f1f07a89d25d2ef17ee9df09055181526020017f076681650542fa47eba4632baad88513eba0758461e06f8177a48604fc46227f815250816080015160058151811061089257fe5b602002602001018190525060405180604001604052807f0ebd0b4acdb1b18e08e648c44c47b4731d78c4ccad42542929a1084d8eaeaf7581526020017f2c36a990b92285a65c8a0594daa8bb71638bc242461486c49639ea0c54ebcdea815250816080015160068151811061090357fe5b602002602001018190525060405180604001604052807f21cde01401e40af1736ff000851c39e71e4b6852275a7e2f0ec046b5aef281fc81526020017f24ccc84adb6bf7148e2ad3dba77a621c477440070da9765c7d8617c985a30118815250816080015160078151811061097457fe5b602002602001018190525060405180604001604052807f25736209042c8507eca3d150689827b0e34dba68bd0883f8f64af67c193ea89981526020017f2ad83c650ff617ba946895317bfbb48be6bfba646972346872c48d5d23f7b62a81525081608001516008815181106109e557fe5b602002602001018190525060405180604001604052807f16221514886880eaaf20a8f31dc9e90cec1360dee0be7874eb167f791ad511a481526020017f2b11ee944484273b77aec337d5e8b507f7dbb3baaed3ffcb4c0c837dc5f3efeb8152508160800151600981518110610a5657fe5b602002602001018190525060405180604001604052807f046febb49880dc8a06c67985a3ce3e5bd41275496b413d4930448b11300ced9681526020017f304b0670f913beb293ffa356274c1853645e3e005a613815e996a6284c2445458152508160800151600a81518110610ac757fe5b602002602001018190525090565b610add610fda565b610ae5610ff4565b835181526020808501519082015260408101839052600060608360808460076107d05a03fa9050808015610b1857610b1a565bfe5b5080610b2557600080fd5b505092915050565b610b35610fda565b610b3d611012565b8351815260208085015181830152835160408301528301516060808301919091526000908360c08460066107d05a03fa9050808015610b1857610b1a565b610b83610fda565b81517f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4790158015610bb657506020830151155b15610bd65750506040805180820190915260008082526020820152610c02565b60405180604001604052808460000151815260200182856020015181610bf857fe5b0683038152509150505b919050565b60408051600480825260a0820190925260009160609190816020015b610c2b610fda565b815260200190600190039081610c2357505060408051600480825260a0820190925291925060609190602082015b610c61611030565b815260200190600190039081610c595790505090508a82600081518110610c8457fe5b60200260200101819052508882600181518110610c9d57fe5b60200260200101819052508682600281518110610cb657fe5b60200260200101819052508482600381518110610ccf57fe5b60200260200101819052508981600081518110610ce857fe5b60200260200101819052508781600181518110610d0157fe5b60200260200101819052508581600281518110610d1a57fe5b60200260200101819052508381600381518110610d3357fe5b6020026020010181905250610d488282610d57565b9b9a5050505050505050505050565b60008151835114610d6757600080fd5b8251604080516006830280825260c084028201602001909252606090828015610d9a578160200160208202803883390190505b50905060005b83811015610f1f57868181518110610db457fe5b602002602001015160000151828260060260000181518110610dd257fe5b602002602001018181525050868181518110610dea57fe5b602002602001015160200151828260060260010181518110610e0857fe5b602002602001018181525050858181518110610e2057fe5b602090810291909101015151518251839060026006850201908110610e4157fe5b602002602001018181525050858181518110610e5957fe5b60209081029190910101515160016020020151828260060260030181518110610e7e57fe5b602002602001018181525050858181518110610e9657fe5b602002602001015160200151600060028110610eae57fe5b6020020151828260060260040181518110610ec557fe5b602002602001018181525050858181518110610edd57fe5b602002602001015160200151600160028110610ef557fe5b6020020151828260060260050181518110610f0c57fe5b6020908102919091010152600101610da0565b50610f28611050565b6000602082602086026020860160086107d05a03fa9050808015610b18575080610f5157600080fd5b505115159695505050505050565b604051806101000160405280610f73610fda565b8152602001610f80611030565b8152602001610f8d610fda565b905290565b604051806101e00160405280610fa6610fda565b8152602001610fb3611030565b8152602001610fc0611030565b8152602001610fcd611030565b8152602001606081525090565b604051806040016040528060008152602001600081525090565b60405180606001604052806003906020820280388339509192915050565b60405180608001604052806004906020820280388339509192915050565b604051806080016040528061104361106e565b8152602001610f8d61106e565b60405180602001604052806001906020820280388339509192915050565b6040518060400160405280600290602082028038833950919291505056fea265627a7a72315820fd706750724b88ba7a5e249dbb5674549a4cffb16955d95eec8f513f77899c1164736f6c637827302e352e31312d646576656c6f702e323031392e372e332b636f6d6d69742e62383337373035320057`

// DeployVerifier deploys a new Ethereum contract, binding an instance of Verifier to it.
func DeployVerifier(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Verifier, error) {
	parsed, err := abi.JSON(strings.NewReader(VerifierABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VerifierBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Verifier{VerifierCaller: VerifierCaller{contract: contract}, VerifierTransactor: VerifierTransactor{contract: contract}, VerifierFilterer: VerifierFilterer{contract: contract}}, nil
}

// Verifier is an auto generated Go binding around an Ethereum contract.
type Verifier struct {
	VerifierCaller     // Read-only binding to the contract
	VerifierTransactor // Write-only binding to the contract
	VerifierFilterer   // Log filterer for contract events
}

// VerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type VerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VerifierSession struct {
	Contract     *Verifier         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VerifierCallerSession struct {
	Contract *VerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// VerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VerifierTransactorSession struct {
	Contract     *VerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// VerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type VerifierRaw struct {
	Contract *Verifier // Generic contract binding to access the raw methods on
}

// VerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VerifierCallerRaw struct {
	Contract *VerifierCaller // Generic read-only contract binding to access the raw methods on
}

// VerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VerifierTransactorRaw struct {
	Contract *VerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVerifier creates a new instance of Verifier, bound to a specific deployed contract.
func NewVerifier(address common.Address, backend bind.ContractBackend) (*Verifier, error) {
	contract, err := bindVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Verifier{VerifierCaller: VerifierCaller{contract: contract}, VerifierTransactor: VerifierTransactor{contract: contract}, VerifierFilterer: VerifierFilterer{contract: contract}}, nil
}

// NewVerifierCaller creates a new read-only instance of Verifier, bound to a specific deployed contract.
func NewVerifierCaller(address common.Address, caller bind.ContractCaller) (*VerifierCaller, error) {
	contract, err := bindVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VerifierCaller{contract: contract}, nil
}

// NewVerifierTransactor creates a new write-only instance of Verifier, bound to a specific deployed contract.
func NewVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*VerifierTransactor, error) {
	contract, err := bindVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VerifierTransactor{contract: contract}, nil
}

// NewVerifierFilterer creates a new log filterer instance of Verifier, bound to a specific deployed contract.
func NewVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*VerifierFilterer, error) {
	contract, err := bindVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VerifierFilterer{contract: contract}, nil
}

// bindVerifier binds a generic wrapper to an already deployed contract.
func bindVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VerifierABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Verifier *VerifierRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Verifier.Contract.VerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Verifier *VerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Verifier.Contract.VerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Verifier *VerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Verifier.Contract.VerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Verifier *VerifierCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Verifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Verifier *VerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Verifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Verifier *VerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Verifier.Contract.contract.Transact(opts, method, params...)
}

// VerifyProof is a free data retrieval call binding the contract method 0xf3bb70f6.
//
// Solidity: function verifyProof(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[10] input) constant returns(bool r)
func (_Verifier *VerifierCaller) VerifyProof(opts *bind.CallOpts, a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [10]*big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Verifier.contract.Call(opts, out, "verifyProof", a, b, c, input)
	return *ret0, err
}

// VerifyProof is a free data retrieval call binding the contract method 0xf3bb70f6.
//
// Solidity: function verifyProof(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[10] input) constant returns(bool r)
func (_Verifier *VerifierSession) VerifyProof(a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [10]*big.Int) (bool, error) {
	return _Verifier.Contract.VerifyProof(&_Verifier.CallOpts, a, b, c, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0xf3bb70f6.
//
// Solidity: function verifyProof(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[10] input) constant returns(bool r)
func (_Verifier *VerifierCallerSession) VerifyProof(a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [10]*big.Int) (bool, error) {
	return _Verifier.Contract.VerifyProof(&_Verifier.CallOpts, a, b, c, input)
}

// WhitelistABI is the input ABI used to generate the binding from.
const WhitelistABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"register\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isRegistered\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"Register\",\"type\":\"event\"}]"

// WhitelistBin is the compiled bytecode used for deploying new contracts.
const WhitelistBin = `0x608060405234801561001057600080fd5b5061016e806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80634420e4861461003b578063c3c5a54714610063575b600080fd5b6100616004803603602081101561005157600080fd5b50356001600160a01b031661009d565b005b6100896004803603602081101561007957600080fd5b50356001600160a01b03166100f6565b604080519115158252519081900360200190f35b6001600160a01b03811660008181526020818152604091829020805460ff19166001179055815192835290517feeda149c76076b34d4b9d8896c2f7efc0d33d1c7b53ea3c5db490d64613f603a9281900390910190a150565b6001600160a01b031660009081526020819052604090205460ff169056fea265627a7a7231582016503f0f6829ff7c3f8871a55759f6a156ba66f16620acc5052547fde16fb9ac64736f6c637827302e352e31312d646576656c6f702e323031392e372e332b636f6d6d69742e62383337373035320057`

// DeployWhitelist deploys a new Ethereum contract, binding an instance of Whitelist to it.
func DeployWhitelist(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Whitelist, error) {
	parsed, err := abi.JSON(strings.NewReader(WhitelistABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(WhitelistBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Whitelist{WhitelistCaller: WhitelistCaller{contract: contract}, WhitelistTransactor: WhitelistTransactor{contract: contract}, WhitelistFilterer: WhitelistFilterer{contract: contract}}, nil
}

// Whitelist is an auto generated Go binding around an Ethereum contract.
type Whitelist struct {
	WhitelistCaller     // Read-only binding to the contract
	WhitelistTransactor // Write-only binding to the contract
	WhitelistFilterer   // Log filterer for contract events
}

// WhitelistCaller is an auto generated read-only Go binding around an Ethereum contract.
type WhitelistCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WhitelistTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WhitelistTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WhitelistFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WhitelistFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WhitelistSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WhitelistSession struct {
	Contract     *Whitelist        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WhitelistCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WhitelistCallerSession struct {
	Contract *WhitelistCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// WhitelistTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WhitelistTransactorSession struct {
	Contract     *WhitelistTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// WhitelistRaw is an auto generated low-level Go binding around an Ethereum contract.
type WhitelistRaw struct {
	Contract *Whitelist // Generic contract binding to access the raw methods on
}

// WhitelistCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WhitelistCallerRaw struct {
	Contract *WhitelistCaller // Generic read-only contract binding to access the raw methods on
}

// WhitelistTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WhitelistTransactorRaw struct {
	Contract *WhitelistTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWhitelist creates a new instance of Whitelist, bound to a specific deployed contract.
func NewWhitelist(address common.Address, backend bind.ContractBackend) (*Whitelist, error) {
	contract, err := bindWhitelist(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Whitelist{WhitelistCaller: WhitelistCaller{contract: contract}, WhitelistTransactor: WhitelistTransactor{contract: contract}, WhitelistFilterer: WhitelistFilterer{contract: contract}}, nil
}

// NewWhitelistCaller creates a new read-only instance of Whitelist, bound to a specific deployed contract.
func NewWhitelistCaller(address common.Address, caller bind.ContractCaller) (*WhitelistCaller, error) {
	contract, err := bindWhitelist(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WhitelistCaller{contract: contract}, nil
}

// NewWhitelistTransactor creates a new write-only instance of Whitelist, bound to a specific deployed contract.
func NewWhitelistTransactor(address common.Address, transactor bind.ContractTransactor) (*WhitelistTransactor, error) {
	contract, err := bindWhitelist(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WhitelistTransactor{contract: contract}, nil
}

// NewWhitelistFilterer creates a new log filterer instance of Whitelist, bound to a specific deployed contract.
func NewWhitelistFilterer(address common.Address, filterer bind.ContractFilterer) (*WhitelistFilterer, error) {
	contract, err := bindWhitelist(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WhitelistFilterer{contract: contract}, nil
}

// bindWhitelist binds a generic wrapper to an already deployed contract.
func bindWhitelist(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WhitelistABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Whitelist *WhitelistRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Whitelist.Contract.WhitelistCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Whitelist *WhitelistRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Whitelist.Contract.WhitelistTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Whitelist *WhitelistRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Whitelist.Contract.WhitelistTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Whitelist *WhitelistCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Whitelist.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Whitelist *WhitelistTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Whitelist.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Whitelist *WhitelistTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Whitelist.Contract.contract.Transact(opts, method, params...)
}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(address addr) constant returns(bool)
func (_Whitelist *WhitelistCaller) IsRegistered(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Whitelist.contract.Call(opts, out, "isRegistered", addr)
	return *ret0, err
}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(address addr) constant returns(bool)
func (_Whitelist *WhitelistSession) IsRegistered(addr common.Address) (bool, error) {
	return _Whitelist.Contract.IsRegistered(&_Whitelist.CallOpts, addr)
}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(address addr) constant returns(bool)
func (_Whitelist *WhitelistCallerSession) IsRegistered(addr common.Address) (bool, error) {
	return _Whitelist.Contract.IsRegistered(&_Whitelist.CallOpts, addr)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address addr) returns()
func (_Whitelist *WhitelistTransactor) Register(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Whitelist.contract.Transact(opts, "register", addr)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address addr) returns()
func (_Whitelist *WhitelistSession) Register(addr common.Address) (*types.Transaction, error) {
	return _Whitelist.Contract.Register(&_Whitelist.TransactOpts, addr)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address addr) returns()
func (_Whitelist *WhitelistTransactorSession) Register(addr common.Address) (*types.Transaction, error) {
	return _Whitelist.Contract.Register(&_Whitelist.TransactOpts, addr)
}

// WhitelistRegisterIterator is returned from FilterRegister and is used to iterate over the raw logs and unpacked data for Register events raised by the Whitelist contract.
type WhitelistRegisterIterator struct {
	Event *WhitelistRegister // Event containing the contract specifics and raw log

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
func (it *WhitelistRegisterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WhitelistRegister)
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
		it.Event = new(WhitelistRegister)
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
func (it *WhitelistRegisterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WhitelistRegisterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WhitelistRegister represents a Register event raised by the Whitelist contract.
type WhitelistRegister struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRegister is a free log retrieval operation binding the contract event 0xeeda149c76076b34d4b9d8896c2f7efc0d33d1c7b53ea3c5db490d64613f603a.
//
// Solidity: event Register(address addr)
func (_Whitelist *WhitelistFilterer) FilterRegister(opts *bind.FilterOpts) (*WhitelistRegisterIterator, error) {

	logs, sub, err := _Whitelist.contract.FilterLogs(opts, "Register")
	if err != nil {
		return nil, err
	}
	return &WhitelistRegisterIterator{contract: _Whitelist.contract, event: "Register", logs: logs, sub: sub}, nil
}

// WatchRegister is a free log subscription operation binding the contract event 0xeeda149c76076b34d4b9d8896c2f7efc0d33d1c7b53ea3c5db490d64613f603a.
//
// Solidity: event Register(address addr)
func (_Whitelist *WhitelistFilterer) WatchRegister(opts *bind.WatchOpts, sink chan<- *WhitelistRegister) (event.Subscription, error) {

	logs, sub, err := _Whitelist.contract.WatchLogs(opts, "Register")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WhitelistRegister)
				if err := _Whitelist.contract.UnpackLog(event, "Register", log); err != nil {
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
