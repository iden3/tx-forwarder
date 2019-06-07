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
const CheckFullCircuitABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"a\",\"type\":\"uint256[2]\"},{\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"name\":\"c\",\"type\":\"uint256[2]\"},{\"name\":\"input\",\"type\":\"uint256[11]\"}],\"name\":\"verify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"smVerifier\",\"type\":\"address\"},{\"name\":\"smRoots\",\"type\":\"address\"},{\"name\":\"smWhitelist\",\"type\":\"address\"},{\"name\":\"_addRelay\",\"type\":\"address\"},{\"name\":\"_addCertifier\",\"type\":\"address\"},{\"name\":\"_addStorer\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"registerId\",\"type\":\"event\"}]"

// CheckFullCircuitBin is the compiled bytecode used for deploying new contracts.
const CheckFullCircuitBin = `0x608060405234801561001057600080fd5b50604051610739380380610739833981810160405260c081101561003357600080fd5b508051602082015160408301516060840151608085015160a090950151600080546001600160a01b03199081166001600160a01b0396871617825560018054821697871697909717909655600280548716948616949094179093556003805486169285169290921790915560048054851695841695909517909455600580549093169190931617905561066d9081906100cc90396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063edf0689414610030575b600080fd5b61012d600480360361026081101561004757600080fd5b6040805180820182529183019291818301918390600290839083908082843760009201829052506040805180820190915293969594608081019493509150600290835b828210156100c8576040805180820182529080840286019060029083908390808284376000920191909152505050815260019091019060200161008a565b5050604080518082018252939695948181019493509150600290839083908082843760009201919091525050604080516101608181019092529295949381810193925090600b90839083908082843760009201919091525091945061012f9350505050565b005b80516020820151604083015161014084015160009061014d906105ef565b600054600354604080516303ce7bb760e11b81526001600160a01b039283166004820152905193945091169163079cf76e91602480820192602092909190829003018186803b15801561019f57600080fd5b505afa1580156101b3573d6000803e3d6000fd5b505050506040513d60208110156101c957600080fd5b5051841461021e576040805162461bcd60e51b815260206004820152601960248201527f526f6f742072656c617920646f6573206e6f74206d6174636800000000000000604482015290519081900360640190fd5b60005460048054604080516303ce7bb760e11b81526001600160a01b03928316938101939093525192169163079cf76e91602480820192602092909190829003018186803b15801561026f57600080fd5b505afa158015610283573d6000803e3d6000fd5b505050506040513d602081101561029957600080fd5b505183146102ee576040805162461bcd60e51b815260206004820152601d60248201527f526f6f742063657274696669657220646f6573206e6f74206d61746368000000604482015290519081900360640190fd5b600054600554604080516303ce7bb760e11b81526001600160a01b0392831660048201529051919092169163079cf76e916024808301926020929190829003018186803b15801561033e57600080fd5b505afa158015610352573d6000803e3d6000fd5b505050506040513d602081101561036857600080fd5b505182146103bd576040805162461bcd60e51b815260206004820152601a60248201527f526f6f742073746f72657220646f6573206e6f74206d61746368000000000000604482015290519081900360640190fd5b6001546040805163b9c6ea8760e01b81526001600160a01b039092169163b9c6ea87918b918b918b918b916004909101908190869080838360005b838110156104105781810151838201526020016103f8565b505050509050018460026000925b8184101561045e5760208402830151604080838360005b8381101561044d578181015183820152602001610435565b50505050905001926001019261041e565b9250505083600260200280838360005b8381101561048657818101518382015260200161046e565b5050505090500182600b60200280838360005b838110156104b1578181015183820152602001610499565b5050505090500194505050505060206040518083038186803b1580156104d657600080fd5b505afa1580156104ea573d6000803e3d6000fd5b505050506040513d602081101561050057600080fd5b505115156001146105425760405162461bcd60e51b81526004018080602001828103825260218152602001806105f36021913960400191505060405180910390fd5b60025460408051632210724360e11b81526001600160a01b03848116600483015291519190921691634420e48691602480830192600092919082900301818387803b15801561059057600080fd5b505af11580156105a4573d6000803e3d6000fd5b5050604080516001600160a01b038516815290517f2035c625189f8d36c7b18ef5cff7e68829acc76264c8372ae47038d503d5b9ba9350908190036020019150a15050505050505050565b9056fe5a65726f2d6b776f776c656467652070726f6f66206973206e6f742076616c6964a265627a7a7230582004744a9c9cff9156333483874da98dd0a27f99f87e29f6d1adfa787d77f8669164736f6c637827302e352e31302d646576656c6f702e323031392e362e362b636f6d6d69742e61343038376434620057`

// DeployCheckFullCircuit deploys a new Ethereum contract, binding an instance of CheckFullCircuit to it.
func DeployCheckFullCircuit(auth *bind.TransactOpts, backend bind.ContractBackend, smVerifier common.Address, smRoots common.Address, smWhitelist common.Address, _addRelay common.Address, _addCertifier common.Address, _addStorer common.Address) (common.Address, *types.Transaction, *CheckFullCircuit, error) {
	parsed, err := abi.JSON(strings.NewReader(CheckFullCircuitABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CheckFullCircuitBin), backend, smVerifier, smRoots, smWhitelist, _addRelay, _addCertifier, _addStorer)
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

// Verify is a paid mutator transaction binding the contract method 0xedf06894.
//
// Solidity: function verify(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[11] input) returns()
func (_CheckFullCircuit *CheckFullCircuitTransactor) Verify(opts *bind.TransactOpts, a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [11]*big.Int) (*types.Transaction, error) {
	return _CheckFullCircuit.contract.Transact(opts, "verify", a, b, c, input)
}

// Verify is a paid mutator transaction binding the contract method 0xedf06894.
//
// Solidity: function verify(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[11] input) returns()
func (_CheckFullCircuit *CheckFullCircuitSession) Verify(a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [11]*big.Int) (*types.Transaction, error) {
	return _CheckFullCircuit.Contract.Verify(&_CheckFullCircuit.TransactOpts, a, b, c, input)
}

// Verify is a paid mutator transaction binding the contract method 0xedf06894.
//
// Solidity: function verify(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[11] input) returns()
func (_CheckFullCircuit *CheckFullCircuitTransactorSession) Verify(a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [11]*big.Int) (*types.Transaction, error) {
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
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRegisterId is a free log retrieval operation binding the contract event 0x2035c625189f8d36c7b18ef5cff7e68829acc76264c8372ae47038d503d5b9ba.
//
// Solidity: event registerId(address addr)
func (_CheckFullCircuit *CheckFullCircuitFilterer) FilterRegisterId(opts *bind.FilterOpts) (*CheckFullCircuitRegisterIdIterator, error) {

	logs, sub, err := _CheckFullCircuit.contract.FilterLogs(opts, "registerId")
	if err != nil {
		return nil, err
	}
	return &CheckFullCircuitRegisterIdIterator{contract: _CheckFullCircuit.contract, event: "registerId", logs: logs, sub: sub}, nil
}

// WatchRegisterId is a free log subscription operation binding the contract event 0x2035c625189f8d36c7b18ef5cff7e68829acc76264c8372ae47038d503d5b9ba.
//
// Solidity: event registerId(address addr)
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

// PairingABI is the input ABI used to generate the binding from.
const PairingABI = "[]"

// PairingBin is the compiled bytecode used for deploying new contracts.
const PairingBin = `0x607a6023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a7230582015f89bea7e57a402406c766b799b92469807d5486ead7f53a742d6ee3c629aaa64736f6c637827302e352e31302d646576656c6f702e323031392e362e362b636f6d6d69742e61343038376434620057`

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
const RootCommitsABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getRoot\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_blockN\",\"type\":\"uint64\"}],\"name\":\"getRootByBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_timestamp\",\"type\":\"uint64\"}],\"name\":\"getRootByTime\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_root\",\"type\":\"bytes32\"}],\"name\":\"setRoot\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"blockN\",\"type\":\"uint64\"},{\"indexed\":false,\"name\":\"timestamp\",\"type\":\"uint64\"},{\"indexed\":false,\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"RootUpdated\",\"type\":\"event\"}]"

// RootCommitsBin is the compiled bytecode used for deploying new contracts.
const RootCommitsBin = `0x608060405234801561001057600080fd5b50610935806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c8063079cf76e1461005157806318c9990b14610089578063ab481d23146100be578063dab5f340146100f3575b600080fd5b6100776004803603602081101561006757600080fd5b50356001600160a01b0316610112565b60408051918252519081900360200190f35b6100776004803603604081101561009f57600080fd5b5080356001600160a01b031690602001356001600160401b031661017a565b610077600480360360408110156100d457600080fd5b5080356001600160a01b031690602001356001600160401b0316610479565b6101106004803603602081101561010957600080fd5b5035610755565b005b6001600160a01b0381166000908152600160205260408120546101385750600054610175565b6001600160a01b03821660009081526001602052604090208054600019810190811061016057fe5b90600052602060002090600202016001015490505b919050565b600043826001600160401b0316106101ce576040805162461bcd60e51b8152602060048201526012602482015271195c9c939bd19d5d1d5c99505b1b1bddd95960721b604482015290519081900360640190fd5b6001600160a01b0383166000908152600160205260409020546101f45750600054610473565b6001600160a01b03831660009081526001602052604081208054600019810190811061021c57fe5b60009182526020909120600290910201546001600160401b0390811691508316811015610285576001600160a01b03841660009081526001602052604090208054600019810190811061026b57fe5b906000526020600020906002020160010154915050610473565b6001600160a01b038416600090815260016020526040812054600019015b80821161046a576001600160a01b03861660009081526001602052604090208054600284840104916001600160401b03881691839081106102e057fe5b60009182526020909120600290910201546001600160401b03161415610341576001600160a01b038716600090815260016020526040902080548290811061032457fe5b906000526020600020906002020160010154945050505050610473565b6001600160a01b038716600090815260016020526040902080548290811061036557fe5b60009182526020909120600290910201546001600160401b039081169087161180156103e0575060016000886001600160a01b03166001600160a01b0316815260200190815260200160002081600101815481106103bf57fe5b60009182526020909120600290910201546001600160401b03908116908716105b15610409576001600160a01b038716600090815260016020526040902080548290811061032457fe5b6001600160a01b038716600090815260016020526040902080548290811061042d57fe5b60009182526020909120600290910201546001600160401b03908116908716111561045d57806001019250610464565b6001810391505b506102a3565b60005493505050505b92915050565b600042826001600160401b0316106104cd576040805162461bcd60e51b8152602060048201526012602482015271195c9c939bd19d5d1d5c99505b1b1bddd95960721b604482015290519081900360640190fd5b6001600160a01b0383166000908152600160205260409020546104f35750600054610473565b6001600160a01b03831660009081526001602052604081208054600019810190811061051b57fe5b60009182526020909120600290910201546001600160401b03600160401b909104811691508316811015610571576001600160a01b03841660009081526001602052604090208054600019810190811061026b57fe5b6001600160a01b038416600090815260016020526040812054600019015b80821161046a576001600160a01b03861660009081526001602052604090208054600284840104916001600160401b03881691839081106105cc57fe5b6000918252602090912060029091020154600160401b90046001600160401b03161415610617576001600160a01b038716600090815260016020526040902080548290811061032457fe5b6001600160a01b038716600090815260016020526040902080548290811061063b57fe5b60009182526020909120600290910201546001600160401b03600160401b90910481169087161180156106c4575060016000886001600160a01b03166001600160a01b03168152602001908152602001600020816001018154811061069c57fe5b60009182526020909120600290910201546001600160401b03600160401b9091048116908716105b156106ed576001600160a01b038716600090815260016020526040902080548290811061032457fe5b6001600160a01b038716600090815260016020526040902080548290811061071157fe5b60009182526020909120600290910201546001600160401b03600160401b909104811690871611156107485780600101925061074f565b6001810391505b5061058f565b33600090815260016020526040902054156107e3573360009081526001602052604090208054439190600019810190811061078c57fe5b60009182526020909120600290910201546001600160401b031614156107e35760405162461bcd60e51b81526004018080602001828103825260218152602001806108bb6021913960400191505060405180910390fd5b3360008181526001602081815260408084208151606080820184526001600160401b034381168084524282168488018181528588018d81528754808c018955978c529a89902095516002909702909501805495518416600160401b026fffffffffffffffff0000000000000000199790941667ffffffffffffffff19909616959095179590951691909117835596519190950155815195865291850193909352838301528201839052517f1d3d3c252bc04ff83e8069c09279b4d8acb5264996a680b6d7dcf20db2b7417b9181900360800190a15056fe6e6f206d756c7469706c652073657420696e207468652073616d6520626c6f636ba265627a7a72305820409b7eda04a4bea9dda45c45b7dfce81dbee7d71139e451fa990dfd8134a408764736f6c637827302e352e31302d646576656c6f702e323031392e362e362b636f6d6d69742e61343038376434620057`

// DeployRootCommits deploys a new Ethereum contract, binding an instance of RootCommits to it.
func DeployRootCommits(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RootCommits, error) {
	parsed, err := abi.JSON(strings.NewReader(RootCommitsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RootCommitsBin), backend)
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

// GetRoot is a free data retrieval call binding the contract method 0x079cf76e.
//
// Solidity: function getRoot(address _address) constant returns(bytes32)
func (_RootCommits *RootCommitsCaller) GetRoot(opts *bind.CallOpts, _address common.Address) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _RootCommits.contract.Call(opts, out, "getRoot", _address)
	return *ret0, err
}

// GetRoot is a free data retrieval call binding the contract method 0x079cf76e.
//
// Solidity: function getRoot(address _address) constant returns(bytes32)
func (_RootCommits *RootCommitsSession) GetRoot(_address common.Address) ([32]byte, error) {
	return _RootCommits.Contract.GetRoot(&_RootCommits.CallOpts, _address)
}

// GetRoot is a free data retrieval call binding the contract method 0x079cf76e.
//
// Solidity: function getRoot(address _address) constant returns(bytes32)
func (_RootCommits *RootCommitsCallerSession) GetRoot(_address common.Address) ([32]byte, error) {
	return _RootCommits.Contract.GetRoot(&_RootCommits.CallOpts, _address)
}

// GetRootByBlock is a free data retrieval call binding the contract method 0x18c9990b.
//
// Solidity: function getRootByBlock(address _address, uint64 _blockN) constant returns(bytes32)
func (_RootCommits *RootCommitsCaller) GetRootByBlock(opts *bind.CallOpts, _address common.Address, _blockN uint64) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _RootCommits.contract.Call(opts, out, "getRootByBlock", _address, _blockN)
	return *ret0, err
}

// GetRootByBlock is a free data retrieval call binding the contract method 0x18c9990b.
//
// Solidity: function getRootByBlock(address _address, uint64 _blockN) constant returns(bytes32)
func (_RootCommits *RootCommitsSession) GetRootByBlock(_address common.Address, _blockN uint64) ([32]byte, error) {
	return _RootCommits.Contract.GetRootByBlock(&_RootCommits.CallOpts, _address, _blockN)
}

// GetRootByBlock is a free data retrieval call binding the contract method 0x18c9990b.
//
// Solidity: function getRootByBlock(address _address, uint64 _blockN) constant returns(bytes32)
func (_RootCommits *RootCommitsCallerSession) GetRootByBlock(_address common.Address, _blockN uint64) ([32]byte, error) {
	return _RootCommits.Contract.GetRootByBlock(&_RootCommits.CallOpts, _address, _blockN)
}

// GetRootByTime is a free data retrieval call binding the contract method 0xab481d23.
//
// Solidity: function getRootByTime(address _address, uint64 _timestamp) constant returns(bytes32)
func (_RootCommits *RootCommitsCaller) GetRootByTime(opts *bind.CallOpts, _address common.Address, _timestamp uint64) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _RootCommits.contract.Call(opts, out, "getRootByTime", _address, _timestamp)
	return *ret0, err
}

// GetRootByTime is a free data retrieval call binding the contract method 0xab481d23.
//
// Solidity: function getRootByTime(address _address, uint64 _timestamp) constant returns(bytes32)
func (_RootCommits *RootCommitsSession) GetRootByTime(_address common.Address, _timestamp uint64) ([32]byte, error) {
	return _RootCommits.Contract.GetRootByTime(&_RootCommits.CallOpts, _address, _timestamp)
}

// GetRootByTime is a free data retrieval call binding the contract method 0xab481d23.
//
// Solidity: function getRootByTime(address _address, uint64 _timestamp) constant returns(bytes32)
func (_RootCommits *RootCommitsCallerSession) GetRootByTime(_address common.Address, _timestamp uint64) ([32]byte, error) {
	return _RootCommits.Contract.GetRootByTime(&_RootCommits.CallOpts, _address, _timestamp)
}

// SetRoot is a paid mutator transaction binding the contract method 0xdab5f340.
//
// Solidity: function setRoot(bytes32 _root) returns()
func (_RootCommits *RootCommitsTransactor) SetRoot(opts *bind.TransactOpts, _root [32]byte) (*types.Transaction, error) {
	return _RootCommits.contract.Transact(opts, "setRoot", _root)
}

// SetRoot is a paid mutator transaction binding the contract method 0xdab5f340.
//
// Solidity: function setRoot(bytes32 _root) returns()
func (_RootCommits *RootCommitsSession) SetRoot(_root [32]byte) (*types.Transaction, error) {
	return _RootCommits.Contract.SetRoot(&_RootCommits.TransactOpts, _root)
}

// SetRoot is a paid mutator transaction binding the contract method 0xdab5f340.
//
// Solidity: function setRoot(bytes32 _root) returns()
func (_RootCommits *RootCommitsTransactorSession) SetRoot(_root [32]byte) (*types.Transaction, error) {
	return _RootCommits.Contract.SetRoot(&_RootCommits.TransactOpts, _root)
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
	Addr      common.Address
	BlockN    uint64
	Timestamp uint64
	Root      [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRootUpdated is a free log retrieval operation binding the contract event 0x1d3d3c252bc04ff83e8069c09279b4d8acb5264996a680b6d7dcf20db2b7417b.
//
// Solidity: event RootUpdated(address addr, uint64 blockN, uint64 timestamp, bytes32 root)
func (_RootCommits *RootCommitsFilterer) FilterRootUpdated(opts *bind.FilterOpts) (*RootCommitsRootUpdatedIterator, error) {

	logs, sub, err := _RootCommits.contract.FilterLogs(opts, "RootUpdated")
	if err != nil {
		return nil, err
	}
	return &RootCommitsRootUpdatedIterator{contract: _RootCommits.contract, event: "RootUpdated", logs: logs, sub: sub}, nil
}

// WatchRootUpdated is a free log subscription operation binding the contract event 0x1d3d3c252bc04ff83e8069c09279b4d8acb5264996a680b6d7dcf20db2b7417b.
//
// Solidity: event RootUpdated(address addr, uint64 blockN, uint64 timestamp, bytes32 root)
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
const VerifierABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"a\",\"type\":\"uint256[2]\"},{\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"name\":\"c\",\"type\":\"uint256[2]\"},{\"name\":\"input\",\"type\":\"uint256[11]\"}],\"name\":\"verifyProof\",\"outputs\":[{\"name\":\"r\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// VerifierBin is the compiled bytecode used for deploying new contracts.
const VerifierBin = `0x608060405234801561001057600080fd5b50611154806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063b9c6ea8714610030575b600080fd5b61012d600480360361026081101561004757600080fd5b6040805180820182529183019291818301918390600290839083908082843760009201829052506040805180820190915293969594608081019493509150600290835b828210156100c8576040805180820182529080840286019060029083908390808284376000920191909152505050815260019091019060200161008a565b5050604080518082018252939695948181019493509150600290839083908082843760009201919091525050604080516101608181019092529295949381810193925090600b9083908390808284376000920191909152509194506101419350505050565b604080519115158252519081900360200190f35b600061014b610fcd565b6040805180820182528751815260208089015181830152908352815160808101835287515181840190815288518301516060808401919091529082528351808501855289840180515182525184015181850152828401528483019190915282518084018452875181528783015181840152848401528251600b80825261018082019094529092909182016101608038833901905050905060005b600b81101561021e578481600b81106101fa57fe5b602002015182828151811061020b57fe5b60209081029190910101526001016101e5565b506102298183610247565b6102385760019250505061023f565b6000925050505b949350505050565b6000610251611000565b610259610358565b905080608001515184516001011461027057600080fd5b610278611048565b50604080518082019091526000808252602082018190525b85518110156102e3576102d9826102d4856080015184600101815181106102b357fe5b60200260200101518985815181106102c757fe5b6020026020010151610b43565b610b9b565b9150600101610290565b506103068183608001516000815181106102f957fe5b6020026020010151610b9b565b905061033c6103188560000151610be9565b8560200151846000015185602001518587604001518a604001518960600151610c75565b61034b57600192505050610352565b6000925050505b92915050565b610360611000565b6040805180820182527f1f81053e1ea1c571aa5a1e0e33cc432f279680cc3c92693d06266f836098079681527f09019a5363c87612e243961622fc13910e9d22295189cfb03eac36674aa7d6646020808301919091529083528151608080820184527f2307626b89b10c5fd4be907c88d5bef1adaf4efe8d0a0a63ef76f74d817e99eb8285019081527f11fc5b82dd7e583bba01bd3b4d8ef1d28c7225ef61fe940d7bd5892b0800c70c606080850191909152908352845180860186527f2ed23243cf9d2d1ade1d825d69bde0ae751392f53c5f4d447d8de1857eeaf35081527f1da2e415bef73f43a6f06da1f5caf146c688b9a38eb9016349ea127afceaf976818601528385015285840192909252835180820185527f1d6b2dd26bc821a9282f1e314aca18344d16e39bcfc9ad83552695d3b2a02d6c8186019081527f0f3681a769f002e13e66a21ea4c49fbd82d88fada6021bf1719ca397a789a421828501528152845180860186527f2d43be907a31eb491ce4567c2dab9b18d056060a105a5690c58e45ac1541496781527eeb8844fa2c3f765762970d74b440167688f8e2b6876edcc49f700569cc0abb818601528185015285850152835190810184527f2ed212362fb8d740ff38d9590e6b5db6aedbc024500dbc4f481b556046fe62e58185019081527f2130c8c3b1ed3a2fcfd4abde05743a6d1dd93c70f94e513a80f100bdee248837828401528152835180850185527f21abc5ea1e9d7b8b22cc40bea1aef61b95a1f3ec80bb67827badb5891654e03c81527f08deef3ad488559d2cebc359ea54c0de755416e362bb227ccd2b42ff2257c03d8185015281840152908401528151600c8082526101a08201909352919082015b6105e3611048565b8152602001906001900390816105db57505060808201908152604080518082019091527f2fa5aaca8a92144ac5e793bd55d38b5cc753848c303bb3884522ef8600169c7681527f2544647a817c84d3fae76aa8d43ff6467a85fa68e94cc9623b0c3131a973392860208201529051805160009061065c57fe5b602002602001018190525060405180604001604052807f014b094b08be69a199d95a7180bc0f7b5eed3db42e4fe4d5d177008f15d096b381526020017f20b937488a33cf9355185518119a4b37c2b45432fd450a236adaa3eefe044b1081525081608001516001815181106106cd57fe5b602002602001018190525060405180604001604052807f0f406ae141a6e27246127829e328183e61fcf207db242f68f2df94269ad07c5381526020017f2ca5b6f6aa167364410f4b2691af7aae14f81fb2ed101710cf754f21b8acb44f815250816080015160028151811061073e57fe5b602002602001018190525060405180604001604052807f1932c7cd2c16d61290e27599f4a9982032fd8515ffcb0eecf75cd10778b192b381526020017f23da605d5f2d83795e862ed21a5195b32f03d7a9618afbe0f30e648f915cb07c81525081608001516003815181106107af57fe5b602002602001018190525060405180604001604052807f1cfbc7eb9c69f30dd44af5894596269fd27d450ac09d02ef5596f4e53006c66881526020017eded045b9cc2725e0ba1b332425ef10ac5d9f9bde35ae6eb33547b827296f20815250816080015160048151811061081f57fe5b602002602001018190525060405180604001604052807f0514b594067e2219ae595c2cf79dc29682ef8088deb1b4fcab3bad21891e7e9c81526020017f21fc107aa14f84b96713cc67f931e9452fd395279fef6333923ca0d040c0cec3815250816080015160058151811061089057fe5b602002602001018190525060405180604001604052807f0d5da04a55f9ed59a997181258a75bfdf13d38756294bd566b055729c1117c5381526020017f285a06883b5b8dd3f7209fd55344ba3c8cdb68542a807625d76673e7a381cbf2815250816080015160068151811061090157fe5b602002602001018190525060405180604001604052807e99885111f3bc475ed956c1439572c4675e67c2e76a1b242aa972a83e9a377f81526020017f141c04dc02375c5e09491781c997a339805f3542ecaeeb439a729bea1808ea5c815250816080015160078151811061097157fe5b602002602001018190525060405180604001604052807f1c4125e9cf817aef314fdeea23a61d48643fb2c6585b5c847ea95324c9a5453981526020017f0feae0f34634b392a777111e530e87a34808d4a7d76891b117e8c06c7324e6d081525081608001516008815181106109e257fe5b602002602001018190525060405180604001604052807f1886ece3808aed97ff4b93eb1134584b1760da2126e71e2c754da3bb53ac3e1381526020017f1e7479cc64b569c08d8a4215fa7e9e1d26baf36a5f32b66388689ee9771f5be38152508160800151600981518110610a5357fe5b602002602001018190525060405180604001604052807f2647f6c402e995c19736b387ce6e388765980a8cb154185c0c7d307128162c7581526020017f03ee7b60dce357f080f35004b4722eab36937912045feafa3ceea077c4401b008152508160800151600a81518110610ac457fe5b602002602001018190525060405180604001604052807f27be6584b7902a87a6a53010cdb848d0d7ed34234131dfba6cfc37196d9bbb1781526020017f08bd9f51448d4e2470bfcc70afa3f8bc8e643d58055597fc5fe8155e73a7be748152508160800151600b81518110610b3557fe5b602002602001018190525090565b610b4b611048565b610b53611062565b835181526020808501519082015260408101839052600060608360808460076107d05a03fa9050808015610b8657610b88565bfe5b5080610b9357600080fd5b505092915050565b610ba3611048565b610bab611080565b8351815260208085015181830152835160408301528301516060808301919091526000908360c08460066107d05a03fa9050808015610b8657610b88565b610bf1611048565b81517f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4790158015610c2457506020830151155b15610c445750506040805180820190915260008082526020820152610c70565b60405180604001604052808460000151815260200182856020015181610c6657fe5b0683038152509150505b919050565b60408051600480825260a0820190925260009160609190816020015b610c99611048565b815260200190600190039081610c9157505060408051600480825260a0820190925291925060609190602082015b610ccf61109e565b815260200190600190039081610cc75790505090508a82600081518110610cf257fe5b60200260200101819052508882600181518110610d0b57fe5b60200260200101819052508682600281518110610d2457fe5b60200260200101819052508482600381518110610d3d57fe5b60200260200101819052508981600081518110610d5657fe5b60200260200101819052508781600181518110610d6f57fe5b60200260200101819052508581600281518110610d8857fe5b60200260200101819052508381600381518110610da157fe5b6020026020010181905250610db68282610dc5565b9b9a5050505050505050505050565b60008151835114610dd557600080fd5b8251604080516006830280825260c084028201602001909252606090828015610e08578160200160208202803883390190505b50905060005b83811015610f8d57868181518110610e2257fe5b602002602001015160000151828260060260000181518110610e4057fe5b602002602001018181525050868181518110610e5857fe5b602002602001015160200151828260060260010181518110610e7657fe5b602002602001018181525050858181518110610e8e57fe5b602090810291909101015151518251839060026006850201908110610eaf57fe5b602002602001018181525050858181518110610ec757fe5b60209081029190910101515160016020020151828260060260030181518110610eec57fe5b602002602001018181525050858181518110610f0457fe5b602002602001015160200151600060028110610f1c57fe5b6020020151828260060260040181518110610f3357fe5b602002602001018181525050858181518110610f4b57fe5b602002602001015160200151600160028110610f6357fe5b6020020151828260060260050181518110610f7a57fe5b6020908102919091010152600101610e0e565b50610f966110be565b6000602082602086026020860160086107d05a03fa9050808015610b86575080610fbf57600080fd5b505115159695505050505050565b604051806101000160405280610fe1611048565b8152602001610fee61109e565b8152602001610ffb611048565b905290565b604051806101e00160405280611014611048565b815260200161102161109e565b815260200161102e61109e565b815260200161103b61109e565b8152602001606081525090565b604051806040016040528060008152602001600081525090565b60405180606001604052806003906020820280388339509192915050565b60405180608001604052806004906020820280388339509192915050565b60405180608001604052806110b16110dc565b8152602001610ffb6110dc565b60405180602001604052806001906020820280388339509192915050565b6040518060400160405280600290602082028038833950919291505056fea265627a7a72305820616e79d2ab2a09149dfab7adad2d5281e6a826d1d116eaaf66f38d4a3c60475e64736f6c637827302e352e31302d646576656c6f702e323031392e362e362b636f6d6d69742e61343038376434620057`

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

// VerifyProof is a free data retrieval call binding the contract method 0xb9c6ea87.
//
// Solidity: function verifyProof(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[11] input) constant returns(bool r)
func (_Verifier *VerifierCaller) VerifyProof(opts *bind.CallOpts, a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [11]*big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Verifier.contract.Call(opts, out, "verifyProof", a, b, c, input)
	return *ret0, err
}

// VerifyProof is a free data retrieval call binding the contract method 0xb9c6ea87.
//
// Solidity: function verifyProof(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[11] input) constant returns(bool r)
func (_Verifier *VerifierSession) VerifyProof(a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [11]*big.Int) (bool, error) {
	return _Verifier.Contract.VerifyProof(&_Verifier.CallOpts, a, b, c, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0xb9c6ea87.
//
// Solidity: function verifyProof(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[11] input) constant returns(bool r)
func (_Verifier *VerifierCallerSession) VerifyProof(a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [11]*big.Int) (bool, error) {
	return _Verifier.Contract.VerifyProof(&_Verifier.CallOpts, a, b, c, input)
}

// WhitelistABI is the input ABI used to generate the binding from.
const WhitelistABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"register\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isRegistered\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"Register\",\"type\":\"event\"}]"

// WhitelistBin is the compiled bytecode used for deploying new contracts.
const WhitelistBin = `0x608060405234801561001057600080fd5b5061016e806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80634420e4861461003b578063c3c5a54714610063575b600080fd5b6100616004803603602081101561005157600080fd5b50356001600160a01b031661009d565b005b6100896004803603602081101561007957600080fd5b50356001600160a01b03166100f6565b604080519115158252519081900360200190f35b6001600160a01b03811660008181526020818152604091829020805460ff19166001179055815192835290517feeda149c76076b34d4b9d8896c2f7efc0d33d1c7b53ea3c5db490d64613f603a9281900390910190a150565b6001600160a01b031660009081526020819052604090205460ff169056fea265627a7a72305820efc936a92bc178a6b5203ecb4fa0d1bc1d3263aab054eb432f26a0842a89051664736f6c637827302e352e31302d646576656c6f702e323031392e362e362b636f6d6d69742e61343038376434620057`

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
