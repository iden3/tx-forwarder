// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package verifier

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

// PairingABI is the input ABI used to generate the binding from.
const PairingABI = "[]"

// PairingBin is the compiled bytecode used for deploying new contracts.
const PairingBin = `0x607a6023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a7231582075c403b1ee8c12ce9fd64ae4b9f42d0e38d4cf60c99f44bff146a05c8302d5d164736f6c637827302e352e31312d646576656c6f702e323031392e372e342b636f6d6d69742e30313965633633660057`

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

// VerifierABI is the input ABI used to generate the binding from.
const VerifierABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"a\",\"type\":\"uint256[2]\"},{\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"name\":\"c\",\"type\":\"uint256[2]\"},{\"name\":\"input\",\"type\":\"uint256[10]\"}],\"name\":\"verifyProof\",\"outputs\":[{\"name\":\"r\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// VerifierBin is the compiled bytecode used for deploying new contracts.
const VerifierBin = `0x608060405234801561001057600080fd5b506110e6806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063f3bb70f614610030575b600080fd5b61012d600480360361024081101561004757600080fd5b6040805180820182529183019291818301918390600290839083908082843760009201829052506040805180820190915293969594608081019493509150600290835b828210156100c8576040805180820182529080840286019060029083908390808284376000920191909152505050815260019091019060200161008a565b5050604080518082018252939695948181019493509150600290839083908082843760009201919091525050604080516101408181019092529295949381810193925090600a9083908390808284376000920191909152509194506101419350505050565b604080519115158252519081900360200190f35b600061014b610f5f565b6040805180820182528751815260208089015181830152908352815160808101835287515181840190815288518301516060808401919091529082528351808501855289840180515182525184015181850152828401528483019190915282518084018452875181528783015181840152848401528251600a80825261016082019094529092909182016101408038833901905050905060005b600a81101561021e578481600a81106101fa57fe5b602002015182828151811061020b57fe5b60209081029190910101526001016101e5565b506102298183610247565b6102385760019250505061023f565b6000925050505b949350505050565b6000610251610f92565b610259610358565b905080608001515184516001011461027057600080fd5b610278610fda565b50604080518082019091526000808252602082018190525b85518110156102e3576102d9826102d4856080015184600101815181106102b357fe5b60200260200101518985815181106102c757fe5b6020026020010151610ad5565b610b2d565b9150600101610290565b506103068183608001516000815181106102f957fe5b6020026020010151610b2d565b905061033c6103188560000151610b7b565b8560200151846000015185602001518587604001518a604001518960600151610c07565b61034b57600192505050610352565b6000925050505b92915050565b610360610f92565b6040805180820182527f2614b958e596d9690f5d3217af600404ea44eb7dd3aa13e36fdd7a87700ee4f881527f05ff6d4734079449005efcfedb3aee7576c590bd639d2cce21f65886b8634faa6020808301919091529083528151608080820184527f01720d524f0fc778bbde14c7b822e64949bf00279f12bbcfc50705824a64dad88285019081527f0669e94bdbe1ea368aa1ab27fee5dc0c5a373bf517ce52d8b662545beafff4ab606080850191909152908352845180860186527f0b28c4152fe2becbc836985f8ad909d777bbece88d496e07f12b90f627c006ba81527f1521166f34a9a17825359aa95ea8048b3511ec2cb6b5d785790be73fc6bba1e3818601528385015285840192909252835180820185527f199d5edac0e49cff949ef3d575f165a0eaaf375055cc688080a4c59347a6d71f8186019081527f0e4e11831095dc06701e66848e73b561de0b884174c5450c6f20979311351de3828501528152845180860186527f1fba33820ddec168c9ca62558e898f5fdce2b41eee44545fa6e11fd294aa69b781527f204d6f8abc7dd042ed457bdcee4e75b1cce9ae5314bd82db5e96e771e17754d6818601528185015285850152835190810184527f0c3f7a9eab611dea28bf032973a27402eb22c3ec6836ced8c87415bcec0d949d8185019081527f295ee9899095c28912781a515d23ae2bb0ec6158675342b28bcbf8559a15c1bf828401528152835180850185527f1f1aaeb4d95088adfeec589e094d56fd7490e9f5e326c3ce39adb9419e2fb8dc81527f026c18a2d69f2346717cc9da2fce9454ecf9f4d7cc054ac417b8b43a54c56b708185015281840152908401528151600b8082526101808201909352919082015b6105e4610fda565b8152602001906001900390816105dc57505060808201908152604080518082019091527f14cad51fa47c563e7e295318e071cec1ce02028b5b157681b7093a6fa31aa91e81527f2f1518dd4e4564fec533dcfb6df43794214e9d5a5a3b0c76a5da1b426056d0c160208201529051805160009061065d57fe5b602002602001018190525060405180604001604052807f08316ff5fa7c62df4e7f84b3a1d6dfc18dffc72f1e5146005424591ececa426581526020017f0b4e3e9f78525e412fa03a75cdfd28e3c6aaf3ffdda151e444e3de3910a9814481525081608001516001815181106106ce57fe5b602002602001018190525060405180604001604052807f092254d5ce850404029f14b8b7c96890185f19edf3ad490aa5872fdd29b41d0881526020017f168eed5e0b4ac1fc0baa328477427807e985dbe9a9145a5e4335c6d6d0205ec4815250816080015160028151811061073f57fe5b602002602001018190525060405180604001604052807f019ddba6b4f6877b647b3dc0fdc6c5628539a38c4e7187a38246bdc4c47551f981526020017f14f3662a55eb0ad31223324899d0c7a03825556bf82f1d66d2a7fe85f1be613381525081608001516003815181106107b057fe5b602002602001018190525060405180604001604052807f08fd17f1714f6db6c44d8a619fce65a9b78cd3ad9f570428d26de6764b49a5b081526020017f2d070656c2816ac138c7fb8b92c1b82d08312d3284c2f61fbc10090aa5e1186e815250816080015160048151811061082157fe5b602002602001018190525060405180604001604052807f2ad2142b5edbde99bb8d08f37726ea44b323f1f07a89d25d2ef17ee9df09055181526020017f076681650542fa47eba4632baad88513eba0758461e06f8177a48604fc46227f815250816080015160058151811061089257fe5b602002602001018190525060405180604001604052807f0ebd0b4acdb1b18e08e648c44c47b4731d78c4ccad42542929a1084d8eaeaf7581526020017f2c36a990b92285a65c8a0594daa8bb71638bc242461486c49639ea0c54ebcdea815250816080015160068151811061090357fe5b602002602001018190525060405180604001604052807f21cde01401e40af1736ff000851c39e71e4b6852275a7e2f0ec046b5aef281fc81526020017f24ccc84adb6bf7148e2ad3dba77a621c477440070da9765c7d8617c985a30118815250816080015160078151811061097457fe5b602002602001018190525060405180604001604052807f25736209042c8507eca3d150689827b0e34dba68bd0883f8f64af67c193ea89981526020017f2ad83c650ff617ba946895317bfbb48be6bfba646972346872c48d5d23f7b62a81525081608001516008815181106109e557fe5b602002602001018190525060405180604001604052807f16221514886880eaaf20a8f31dc9e90cec1360dee0be7874eb167f791ad511a481526020017f2b11ee944484273b77aec337d5e8b507f7dbb3baaed3ffcb4c0c837dc5f3efeb8152508160800151600981518110610a5657fe5b602002602001018190525060405180604001604052807f046febb49880dc8a06c67985a3ce3e5bd41275496b413d4930448b11300ced9681526020017f304b0670f913beb293ffa356274c1853645e3e005a613815e996a6284c2445458152508160800151600a81518110610ac757fe5b602002602001018190525090565b610add610fda565b610ae5610ff4565b835181526020808501519082015260408101839052600060608360808460076107d05a03fa9050808015610b1857610b1a565bfe5b5080610b2557600080fd5b505092915050565b610b35610fda565b610b3d611012565b8351815260208085015181830152835160408301528301516060808301919091526000908360c08460066107d05a03fa9050808015610b1857610b1a565b610b83610fda565b81517f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4790158015610bb657506020830151155b15610bd65750506040805180820190915260008082526020820152610c02565b60405180604001604052808460000151815260200182856020015181610bf857fe5b0683038152509150505b919050565b60408051600480825260a0820190925260009160609190816020015b610c2b610fda565b815260200190600190039081610c2357505060408051600480825260a0820190925291925060609190602082015b610c61611030565b815260200190600190039081610c595790505090508a82600081518110610c8457fe5b60200260200101819052508882600181518110610c9d57fe5b60200260200101819052508682600281518110610cb657fe5b60200260200101819052508482600381518110610ccf57fe5b60200260200101819052508981600081518110610ce857fe5b60200260200101819052508781600181518110610d0157fe5b60200260200101819052508581600281518110610d1a57fe5b60200260200101819052508381600381518110610d3357fe5b6020026020010181905250610d488282610d57565b9b9a5050505050505050505050565b60008151835114610d6757600080fd5b8251604080516006830280825260c084028201602001909252606090828015610d9a578160200160208202803883390190505b50905060005b83811015610f1f57868181518110610db457fe5b602002602001015160000151828260060260000181518110610dd257fe5b602002602001018181525050868181518110610dea57fe5b602002602001015160200151828260060260010181518110610e0857fe5b602002602001018181525050858181518110610e2057fe5b602090810291909101015151518251839060026006850201908110610e4157fe5b602002602001018181525050858181518110610e5957fe5b60209081029190910101515160016020020151828260060260030181518110610e7e57fe5b602002602001018181525050858181518110610e9657fe5b602002602001015160200151600060028110610eae57fe5b6020020151828260060260040181518110610ec557fe5b602002602001018181525050858181518110610edd57fe5b602002602001015160200151600160028110610ef557fe5b6020020151828260060260050181518110610f0c57fe5b6020908102919091010152600101610da0565b50610f28611050565b6000602082602086026020860160086107d05a03fa9050808015610b18575080610f5157600080fd5b505115159695505050505050565b604051806101000160405280610f73610fda565b8152602001610f80611030565b8152602001610f8d610fda565b905290565b604051806101e00160405280610fa6610fda565b8152602001610fb3611030565b8152602001610fc0611030565b8152602001610fcd611030565b8152602001606081525090565b604051806040016040528060008152602001600081525090565b60405180606001604052806003906020820280388339509192915050565b60405180608001604052806004906020820280388339509192915050565b604051806080016040528061104361106e565b8152602001610f8d61106e565b60405180602001604052806001906020820280388339509192915050565b6040518060400160405280600290602082028038833950919291505056fea265627a7a72315820f17ffb597961a6196816a3d8b6e881305c001905b5ee7fb23e6afab2d58de4e464736f6c637827302e352e31312d646576656c6f702e323031392e372e342b636f6d6d69742e30313965633633660057`

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
