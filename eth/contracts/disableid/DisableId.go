// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package DisableId

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

// DisableIdABI is the input ABI used to generate the binding from.
const DisableIdABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes31\"}],\"name\":\"disableList\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\"},{\"name\":\"rsv\",\"type\":\"bytes\"}],\"name\":\"checkSig\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"mtp\",\"type\":\"bytes\"},{\"name\":\"id\",\"type\":\"bytes31\"},{\"name\":\"ethAddress\",\"type\":\"address\"},{\"name\":\"msgHash\",\"type\":\"bytes32\"},{\"name\":\"rsv\",\"type\":\"bytes\"}],\"name\":\"disableIdentity\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes31\"}],\"name\":\"getRootFromId\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes27\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_mimcContractAddr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes31\"}],\"name\":\"disableIdEvent\",\"type\":\"event\"}]"

// DisableIdBin is the compiled bytecode used for deploying new contracts.
const DisableIdBin = `0x608060405234801561001057600080fd5b50604051610b47380380610b478339818101604052602081101561003357600080fd5b5051600080546001600160a01b039092166001600160a01b0319909216919091179055610ae2806100656000396000f3fe608060405234801561001057600080fd5b506004361061004b5760003560e01c8062dc8cdc1461005057806301b0452c1461008557806349e9d2a31461014e578063ad05a8d21461029c575b600080fd5b6100716004803603602081101561006657600080fd5b503560ff19166102d8565b604080519115158252519081900360200190f35b6101326004803603604081101561009b57600080fd5b813591908101906040810160208201356401000000008111156100bd57600080fd5b8201836020820111156100cf57600080fd5b803590602001918460018302840111640100000000831117156100f157600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506102ed945050505050565b604080516001600160a01b039092168252519081900360200190f35b61029a600480360360a081101561016457600080fd5b81019060208101813564010000000081111561017f57600080fd5b82018360208201111561019157600080fd5b803590602001918460018302840111640100000000831117156101b357600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929560ff19853516956001600160a01b036020870135169560408101359550919350915060808101906060013564010000000081111561022557600080fd5b82018360208201111561023757600080fd5b8035906020019184600183028401116401000000008311171561025957600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061036c945050505050565b005b6102bd600480360360208110156102b257600080fd5b503560ff19166104d9565b6040805164ffffffffff199092168252519081900360200190f35b60016020526000908152604090205460ff1681565b602081810151604080840151606080860151835160008082528188018087528a905291821a81860181905292810186905260808101849052935190959293919260019260a080820193601f1981019281900390910190855afa158015610357573d6000803e3d6000fd5b5050604051601f190151979650505050505050565b826001600160a01b031661038083836102ed565b6001600160a01b0316146103d1576040805162461bcd60e51b815260206004820152601360248201527214da59db985d1d5c99481b9bdd081d985b1a59606a1b604482015290519081900360640190fd5b6103d9610a11565b6103e2846104e5565b90506103ec610a11565b6103f582610520565b905060008061040383610578565b90925090506000610413896104d9565b9050610423818b8585608c6105b1565b1515600114610479576040805162461bcd60e51b815260206004820152601b60248201527f4d65726b6c6520747265652070726f6f66206e6f742076616c69640000000000604482015290519081900360640190fd5b60ff19808a1660008181526001602081815260409283902080549095169091179093558051918252517fad7c8149a76b6522635bb10cf23aa364e427591e066be8eab4474491eb2c479d929181900390910190a150505050505050505050565b60101b62ffffff191690565b6104ed610a11565b600960c01b8152600060208201819052606092831b6bffffffffffffffffffffffff191660408301529181019190915290565b610528610a11565b815160c01c6060828101828152602085015160a01c6bffffffff000000000000000016909217909152604080840151821c8184018181529290940151901c63ffffffff60a01b1690921790915290565b600080610595836040015160001c846060015160001c60006107b1565b835160208501519193506105aa9160006107b1565b9050915091565b60006105bb610a38565b6105c48661082b565b90506000816000015180156105da575081602001515b15610643576000805b811580156105f057508581105b15610618578088901c600116818560800151901c6001161860ff1691506001810190506105e3565b8161062a5760009450505050506107a8565b61063e84608001518560a0015160016107b1565b925050505b6060826040015160ff16604051908082528060200260200182016040528015610676578160200160208202803883390190505b509050600080805b856040015160ff168110156107075760608601516001600160f01b0316811c6001908116908114156106e35760008360200260400160ff169050808d01519450848684815181106106cb57fe5b602002602001018181525050600184019350506106fe565b60008583815181106106f157fe5b6020026020010181815250505b5060010161067e565b5084516000906107225761071d8a8a60016107b1565b610724565b845b9050600080600188604001510360ff1690505b85818151811061074357fe5b6020908102919091010151915060018c821c8116148061076e57610769848460006107b1565b61077a565b61077a838560006107b1565b9350816107875750610791565b5060001901610737565b505064ffffffffff198c8116911614955050505050505b95945050505050565b604080516002808252606080830184526000939092919060208301908038833901905050905084816000815181106107e557fe5b60200260200101818152505083816001815181106107ff57fe5b602002602001018181525050826108205761081b8160006108ba565b6107a8565b6107a88160016108ba565b610833610a38565b61083b610a6d565b610844836108cd565b90506000610851826108f6565b60018082168114855281811c81161460208501529050610870826108f6565b60ff16604084015261088182610908565b6001600160f01b031660608401526020830151156108b35783518401601f198101519051608085019190915260a08401525b5050919050565b60006108c6838361091a565b9392505050565b6108d5610a6d565b50604080518082019091526020828101825282518301810190820152919050565b80518051600190910190915260f81c90565b80518051601e90910190915260101c90565b6000817f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001825b8551811015610a0757600054865183916001600160a01b03169063d15ca1099089908590811061096c57fe5b6020026020010151866040518363ffffffff1660e01b8152600401808381526020018281526020019250505060206040518083038186803b1580156109b057600080fd5b505afa1580156109c4573d6000803e3d6000fd5b505050506040513d60208110156109da57600080fd5b505187518890849081106109ea57fe5b6020026020010151850101816109fc57fe5b069250600101610940565b5090949350505050565b60408051608081018252600080825260208201819052918101829052606081019190915290565b6040805160c081018252600080825260208201819052918101829052606081018290526080810182905260a081019190915290565b60405180604001604052806000815260200160008152509056fea265627a7a7230582063c58d842febdca4ce60cb42f2292f3dc491b21a11a70cbae7bd3d83e04e7fed64736f6c637828302e352e31302d646576656c6f702e323031392e362e31332b636f6d6d69742e36363839373262620058`

// DeployDisableId deploys a new Ethereum contract, binding an instance of DisableId to it.
func DeployDisableId(auth *bind.TransactOpts, backend bind.ContractBackend, _mimcContractAddr common.Address) (common.Address, *types.Transaction, *DisableId, error) {
	parsed, err := abi.JSON(strings.NewReader(DisableIdABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DisableIdBin), backend, _mimcContractAddr)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DisableId{DisableIdCaller: DisableIdCaller{contract: contract}, DisableIdTransactor: DisableIdTransactor{contract: contract}, DisableIdFilterer: DisableIdFilterer{contract: contract}}, nil
}

// DisableId is an auto generated Go binding around an Ethereum contract.
type DisableId struct {
	DisableIdCaller     // Read-only binding to the contract
	DisableIdTransactor // Write-only binding to the contract
	DisableIdFilterer   // Log filterer for contract events
}

// DisableIdCaller is an auto generated read-only Go binding around an Ethereum contract.
type DisableIdCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DisableIdTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DisableIdTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DisableIdFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DisableIdFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DisableIdSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DisableIdSession struct {
	Contract     *DisableId        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DisableIdCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DisableIdCallerSession struct {
	Contract *DisableIdCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// DisableIdTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DisableIdTransactorSession struct {
	Contract     *DisableIdTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// DisableIdRaw is an auto generated low-level Go binding around an Ethereum contract.
type DisableIdRaw struct {
	Contract *DisableId // Generic contract binding to access the raw methods on
}

// DisableIdCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DisableIdCallerRaw struct {
	Contract *DisableIdCaller // Generic read-only contract binding to access the raw methods on
}

// DisableIdTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DisableIdTransactorRaw struct {
	Contract *DisableIdTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDisableId creates a new instance of DisableId, bound to a specific deployed contract.
func NewDisableId(address common.Address, backend bind.ContractBackend) (*DisableId, error) {
	contract, err := bindDisableId(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DisableId{DisableIdCaller: DisableIdCaller{contract: contract}, DisableIdTransactor: DisableIdTransactor{contract: contract}, DisableIdFilterer: DisableIdFilterer{contract: contract}}, nil
}

// NewDisableIdCaller creates a new read-only instance of DisableId, bound to a specific deployed contract.
func NewDisableIdCaller(address common.Address, caller bind.ContractCaller) (*DisableIdCaller, error) {
	contract, err := bindDisableId(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DisableIdCaller{contract: contract}, nil
}

// NewDisableIdTransactor creates a new write-only instance of DisableId, bound to a specific deployed contract.
func NewDisableIdTransactor(address common.Address, transactor bind.ContractTransactor) (*DisableIdTransactor, error) {
	contract, err := bindDisableId(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DisableIdTransactor{contract: contract}, nil
}

// NewDisableIdFilterer creates a new log filterer instance of DisableId, bound to a specific deployed contract.
func NewDisableIdFilterer(address common.Address, filterer bind.ContractFilterer) (*DisableIdFilterer, error) {
	contract, err := bindDisableId(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DisableIdFilterer{contract: contract}, nil
}

// bindDisableId binds a generic wrapper to an already deployed contract.
func bindDisableId(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DisableIdABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DisableId *DisableIdRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DisableId.Contract.DisableIdCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DisableId *DisableIdRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DisableId.Contract.DisableIdTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DisableId *DisableIdRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DisableId.Contract.DisableIdTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DisableId *DisableIdCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DisableId.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DisableId *DisableIdTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DisableId.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DisableId *DisableIdTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DisableId.Contract.contract.Transact(opts, method, params...)
}

// CheckSig is a free data retrieval call binding the contract method 0x01b0452c.
//
// Solidity: function checkSig(bytes32 msgHash, bytes rsv) constant returns(address)
func (_DisableId *DisableIdCaller) CheckSig(opts *bind.CallOpts, msgHash [32]byte, rsv []byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DisableId.contract.Call(opts, out, "checkSig", msgHash, rsv)
	return *ret0, err
}

// CheckSig is a free data retrieval call binding the contract method 0x01b0452c.
//
// Solidity: function checkSig(bytes32 msgHash, bytes rsv) constant returns(address)
func (_DisableId *DisableIdSession) CheckSig(msgHash [32]byte, rsv []byte) (common.Address, error) {
	return _DisableId.Contract.CheckSig(&_DisableId.CallOpts, msgHash, rsv)
}

// CheckSig is a free data retrieval call binding the contract method 0x01b0452c.
//
// Solidity: function checkSig(bytes32 msgHash, bytes rsv) constant returns(address)
func (_DisableId *DisableIdCallerSession) CheckSig(msgHash [32]byte, rsv []byte) (common.Address, error) {
	return _DisableId.Contract.CheckSig(&_DisableId.CallOpts, msgHash, rsv)
}

// DisableList is a free data retrieval call binding the contract method 0x00dc8cdc.
//
// Solidity: function disableList(bytes31 ) constant returns(bool)
func (_DisableId *DisableIdCaller) DisableList(opts *bind.CallOpts, arg0 [31]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DisableId.contract.Call(opts, out, "disableList", arg0)
	return *ret0, err
}

// DisableList is a free data retrieval call binding the contract method 0x00dc8cdc.
//
// Solidity: function disableList(bytes31 ) constant returns(bool)
func (_DisableId *DisableIdSession) DisableList(arg0 [31]byte) (bool, error) {
	return _DisableId.Contract.DisableList(&_DisableId.CallOpts, arg0)
}

// DisableList is a free data retrieval call binding the contract method 0x00dc8cdc.
//
// Solidity: function disableList(bytes31 ) constant returns(bool)
func (_DisableId *DisableIdCallerSession) DisableList(arg0 [31]byte) (bool, error) {
	return _DisableId.Contract.DisableList(&_DisableId.CallOpts, arg0)
}

// GetRootFromId is a free data retrieval call binding the contract method 0xad05a8d2.
//
// Solidity: function getRootFromId(bytes31 id) constant returns(bytes27)
func (_DisableId *DisableIdCaller) GetRootFromId(opts *bind.CallOpts, id [31]byte) ([27]byte, error) {
	var (
		ret0 = new([27]byte)
	)
	out := ret0
	err := _DisableId.contract.Call(opts, out, "getRootFromId", id)
	return *ret0, err
}

// GetRootFromId is a free data retrieval call binding the contract method 0xad05a8d2.
//
// Solidity: function getRootFromId(bytes31 id) constant returns(bytes27)
func (_DisableId *DisableIdSession) GetRootFromId(id [31]byte) ([27]byte, error) {
	return _DisableId.Contract.GetRootFromId(&_DisableId.CallOpts, id)
}

// GetRootFromId is a free data retrieval call binding the contract method 0xad05a8d2.
//
// Solidity: function getRootFromId(bytes31 id) constant returns(bytes27)
func (_DisableId *DisableIdCallerSession) GetRootFromId(id [31]byte) ([27]byte, error) {
	return _DisableId.Contract.GetRootFromId(&_DisableId.CallOpts, id)
}

// DisableIdentity is a paid mutator transaction binding the contract method 0x49e9d2a3.
//
// Solidity: function disableIdentity(bytes mtp, bytes31 id, address ethAddress, bytes32 msgHash, bytes rsv) returns()
func (_DisableId *DisableIdTransactor) DisableIdentity(opts *bind.TransactOpts, mtp []byte, id [31]byte, ethAddress common.Address, msgHash [32]byte, rsv []byte) (*types.Transaction, error) {
	return _DisableId.contract.Transact(opts, "disableIdentity", mtp, id, ethAddress, msgHash, rsv)
}

// DisableIdentity is a paid mutator transaction binding the contract method 0x49e9d2a3.
//
// Solidity: function disableIdentity(bytes mtp, bytes31 id, address ethAddress, bytes32 msgHash, bytes rsv) returns()
func (_DisableId *DisableIdSession) DisableIdentity(mtp []byte, id [31]byte, ethAddress common.Address, msgHash [32]byte, rsv []byte) (*types.Transaction, error) {
	return _DisableId.Contract.DisableIdentity(&_DisableId.TransactOpts, mtp, id, ethAddress, msgHash, rsv)
}

// DisableIdentity is a paid mutator transaction binding the contract method 0x49e9d2a3.
//
// Solidity: function disableIdentity(bytes mtp, bytes31 id, address ethAddress, bytes32 msgHash, bytes rsv) returns()
func (_DisableId *DisableIdTransactorSession) DisableIdentity(mtp []byte, id [31]byte, ethAddress common.Address, msgHash [32]byte, rsv []byte) (*types.Transaction, error) {
	return _DisableId.Contract.DisableIdentity(&_DisableId.TransactOpts, mtp, id, ethAddress, msgHash, rsv)
}

// DisableIdDisableIdEventIterator is returned from FilterDisableIdEvent and is used to iterate over the raw logs and unpacked data for DisableIdEvent events raised by the DisableId contract.
type DisableIdDisableIdEventIterator struct {
	Event *DisableIdDisableIdEvent // Event containing the contract specifics and raw log

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
func (it *DisableIdDisableIdEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DisableIdDisableIdEvent)
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
		it.Event = new(DisableIdDisableIdEvent)
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
func (it *DisableIdDisableIdEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DisableIdDisableIdEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DisableIdDisableIdEvent represents a DisableIdEvent event raised by the DisableId contract.
type DisableIdDisableIdEvent struct {
	Id  [31]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDisableIdEvent is a free log retrieval operation binding the contract event 0xad7c8149a76b6522635bb10cf23aa364e427591e066be8eab4474491eb2c479d.
//
// Solidity: event disableIdEvent(bytes31 id)
func (_DisableId *DisableIdFilterer) FilterDisableIdEvent(opts *bind.FilterOpts) (*DisableIdDisableIdEventIterator, error) {

	logs, sub, err := _DisableId.contract.FilterLogs(opts, "disableIdEvent")
	if err != nil {
		return nil, err
	}
	return &DisableIdDisableIdEventIterator{contract: _DisableId.contract, event: "disableIdEvent", logs: logs, sub: sub}, nil
}

// WatchDisableIdEvent is a free log subscription operation binding the contract event 0xad7c8149a76b6522635bb10cf23aa364e427591e066be8eab4474491eb2c479d.
//
// Solidity: event disableIdEvent(bytes31 id)
func (_DisableId *DisableIdFilterer) WatchDisableIdEvent(opts *bind.WatchOpts, sink chan<- *DisableIdDisableIdEvent) (event.Subscription, error) {

	logs, sub, err := _DisableId.contract.WatchLogs(opts, "disableIdEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DisableIdDisableIdEvent)
				if err := _DisableId.contract.UnpackLog(event, "disableIdEvent", log); err != nil {
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
const Iden3HelpersBin = `0x608060405234801561001057600080fd5b5060405161028b38038061028b8339818101604052602081101561003357600080fd5b5051600080546001600160a01b039092166001600160a01b0319909216919091179055610226806100656000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c806301b0452c1461003b578063ad05a8d214610104575b600080fd5b6100e86004803603604081101561005157600080fd5b8135919081019060408101602082013564010000000081111561007357600080fd5b82018360208201111561008557600080fd5b803590602001918460018302840111640100000000831117156100a757600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610140945050505050565b604080516001600160a01b039092168252519081900360200190f35b6101256004803603602081101561011a57600080fd5b503560ff19166101bf565b6040805164ffffffffff199092168252519081900360200190f35b602081810151604080840151606080860151835160008082528188018087528a905291821a81860181905292810186905260808101849052935190959293919260019260a080820193601f1981019281900390910190855afa1580156101aa573d6000803e3d6000fd5b5050604051601f190151979650505050505050565b60101b62ffffff19169056fea265627a7a723058202f80c47be1dfe23a612a18c4e97269979d0bf2b2c0d349b63aecfb703baf69ee64736f6c637828302e352e31302d646576656c6f702e323031392e362e31332b636f6d6d69742e36363839373262620058`

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
const MemoryBin = `0x607b6023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723058200a4ef102adac847d46e8170ef3b7321c3b837ecf43c3aec583e29cfb8553852b64736f6c637828302e352e31302d646576656c6f702e323031392e362e31332b636f6d6d69742e36363839373262620058`

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
const MimcUnitBin = `0x6080604052348015600f57600080fd5b5060c28061001e6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c8063d15ca10914602d575b600080fd5b604d60048036036040811015604157600080fd5b5080359060200135605f565b60408051918252519081900360200190f35b60009291505056fea265627a7a7230582097a4555aba7643d76c574bdd06dd960280bb6f5e858652950339d08819702dbf64736f6c637828302e352e31302d646576656c6f702e323031392e362e31332b636f6d6d69742e36363839373262620058`

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
