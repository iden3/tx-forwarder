package eth

import (
	"context"
	"encoding/hex"
	"errors"
	"io/ioutil"
	"math"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	common3 "github.com/iden3/go-iden3/common"
	fullverifiercontract "github.com/iden3/tx-forwarder/eth/contracts/checkfullcircuit"
	disableid "github.com/iden3/tx-forwarder/eth/contracts/disableid"
	iden3helperscontract "github.com/iden3/tx-forwarder/eth/contracts/iden3helpers"
	mimc7contract "github.com/iden3/tx-forwarder/eth/contracts/mimc7"
	rootcommitscontract "github.com/iden3/tx-forwarder/eth/contracts/rootcommits"
	samplecontract "github.com/iden3/tx-forwarder/eth/contracts/samplecontract"
	verifiercontract "github.com/iden3/tx-forwarder/eth/contracts/verifier"
	whitelistcontract "github.com/iden3/tx-forwarder/eth/contracts/whitelist"
	log "github.com/sirupsen/logrus"
)

type EthService struct {
	ks                          *keystore.KeyStore
	acc                         *accounts.Account
	client                      *ethclient.Client
	SampleContract              *samplecontract.SampleContract
	FullVerifierContract        *fullverifiercontract.CheckFullCircuit
	WhitelistContract           *whitelistcontract.Whitelist
	Mimc7Contract               *disableid.DisableId
	Iden3HelpersContract        *disableid.DisableId
	DisableIdContract           *disableid.DisableId
	sampleContractAddress       common.Address
	rootcommitsContractAddress  common.Address
	whitelistContractAddress    common.Address
	zkpverifierContractAddress  common.Address
	mimc7ContractAddress        common.Address
	iden3helpersContractAddress common.Address
	disableidContractAddress    common.Address
	KeyStore                    struct {
		Path     string
		Password string
	}
}

func (ethSrv *EthService) Client() *ethclient.Client {
	return ethSrv.client
}

func (ethSrv *EthService) Account() *accounts.Account {
	return ethSrv.acc
}

func (ethSrv *EthService) SampleContractAddress() common.Address {
	return ethSrv.sampleContractAddress
}
func (ethSrv *EthService) FullVerifierContractAddress() common.Address {
	return ethSrv.zkpverifierContractAddress
}
func (ethSrv *EthService) Mimc7ContractAddress() common.Address {
	return ethSrv.mimc7ContractAddress
}
func (ethSrv *EthService) Iden3HelpersContractAddress() common.Address {
	return ethSrv.iden3helpersContractAddress
}
func (ethSrv *EthService) DisableIdContractAddress() common.Address {
	return ethSrv.disableidContractAddress
}

func NewEthService(url string, ks *keystore.KeyStore, acc *accounts.Account, keystorePath, password string) *EthService {
	client, err := ethclient.Dial(url)
	if err != nil {
		log.Error("Can not open connection to web3 (config.Web3.Url: " + url + ")\n" + err.Error() + "\n")
		os.Exit(0)
	}
	log.Info("Connection to web3 server opened")

	service := &EthService{
		ks:     ks,
		acc:    acc,
		client: client,
		KeyStore: struct {
			Path     string
			Password string
		}{
			Path:     keystorePath,
			Password: password,
		},
	}

	return service
}

func (ethSrv *EthService) GetBalance(address common.Address) (*big.Float, error) {
	balance, err := ethSrv.client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		return nil, err
	}
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethBalance := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	return ethBalance, nil
}

func (ethSrv *EthService) DeploySampleContract() (*common.Address, *types.Transaction, error) {
	auth, err := ethSrv.GetAuth()
	if err != nil {
		return nil, nil, err
	}

	address, tx, _, err := samplecontract.DeploySampleContract(auth, ethSrv.client)
	if err != nil {
		return nil, nil, err
	}
	ethSrv.sampleContractAddress = address

	return &address, tx, nil
}

func (ethSrv *EthService) DeployFullVerifierContract(smRootCommitsAddr, smWhitelistAddr common.Address, idCertifier, idStorer [31]byte) (*common.Address, *types.Transaction, *common.Address, *types.Transaction, error) {
	auth, err := ethSrv.GetAuth()
	if err != nil {
		return nil, nil, nil, nil, err
	}
	// deploy smVerifier
	smVerifierAddr, smVerifierTx, _, err := verifiercontract.DeployVerifier(auth, ethSrv.client)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	address, tx, _, err := fullverifiercontract.DeployCheckFullCircuit(auth, ethSrv.client, smVerifierAddr, smRootCommitsAddr, smWhitelistAddr, idCertifier, idStorer)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	ethSrv.zkpverifierContractAddress = address

	return &address, tx, &smVerifierAddr, smVerifierTx, nil
}

func (ethSrv *EthService) DeployMimc7Contract() (*common.Address, *types.Transaction, error) {
	auth, err := ethSrv.GetAuth()
	if err != nil {
		return nil, nil, err
	}

	address, tx, _, err := mimc7contract.DeployMimc7(auth, ethSrv.client)
	if err != nil {
		return nil, nil, err
	}
	ethSrv.mimc7ContractAddress = address

	return &address, tx, nil
}

func (ethSrv *EthService) DeployRootCommitsContract(mimc7Addr common.Address) (*common.Address, *types.Transaction, error) {
	auth, err := ethSrv.GetAuth()
	if err != nil {
		return nil, nil, err
	}

	address, tx, _, err := rootcommitscontract.DeployRootCommits(auth, ethSrv.client, mimc7Addr)
	if err != nil {
		return nil, nil, err
	}
	ethSrv.rootcommitsContractAddress = address

	return &address, tx, nil
}

func (ethSrv *EthService) DeployWhitelistContract() (*common.Address, *types.Transaction, error) {
	auth, err := ethSrv.GetAuth()
	if err != nil {
		return nil, nil, err
	}

	address, tx, _, err := whitelistcontract.DeployWhitelist(auth, ethSrv.client)
	if err != nil {
		return nil, nil, err
	}
	ethSrv.whitelistContractAddress = address

	return &address, tx, nil
}

func (ethSrv *EthService) DeployIden3HelpersContract(mimc7Address common.Address) (*common.Address, *types.Transaction, error) {
	auth, err := ethSrv.GetAuth()
	if err != nil {
		return nil, nil, err
	}

	address, tx, _, err := iden3helperscontract.DeployIden3Helpers(auth, ethSrv.client, mimc7Address)
	if err != nil {
		return nil, nil, err
	}
	ethSrv.iden3helpersContractAddress = address

	return &address, tx, nil
}

func (ethSrv *EthService) DeployDisableIdContract(iden3HelpersAddress common.Address) (*common.Address, *types.Transaction, error) {
	auth, err := ethSrv.GetAuth()
	if err != nil {
		return nil, nil, err
	}

	address, tx, _, err := disableid.DeployDisableId(auth, ethSrv.client, iden3HelpersAddress)
	if err != nil {
		return nil, nil, err
	}
	ethSrv.disableidContractAddress = address

	return &address, tx, nil
}

func (ethSrv *EthService) LoadSampleContract(contractAddr common.Address) {
	instance, err := samplecontract.NewSampleContract(contractAddr, ethSrv.client)
	if err != nil {
		log.Error(err.Error())
	}
	ethSrv.sampleContractAddress = contractAddr
	ethSrv.SampleContract = instance
	log.Info("Sample contract with address " + contractAddr.String() + " loaded")
}

func (ethSrv *EthService) LoadWhitelistContract(contractAddr common.Address) {
	instance, err := whitelistcontract.NewWhitelist(contractAddr, ethSrv.client)
	if err != nil {
		log.Error(err.Error())
	}
	ethSrv.whitelistContractAddress = contractAddr
	ethSrv.WhitelistContract = instance
	log.Info("Whitelist contract with address " + contractAddr.String() + " loaded")
}

func (ethSrv *EthService) LoadFullVerifierContract(contractAddr common.Address) {
	instance, err := fullverifiercontract.NewCheckFullCircuit(contractAddr, ethSrv.client)
	if err != nil {
		log.Error(err.Error())
	}
	ethSrv.zkpverifierContractAddress = contractAddr
	ethSrv.FullVerifierContract = instance
	log.Info("FullVerifier contract with address " + contractAddr.String() + " loaded")
}

func (ethSrv *EthService) LoadDisableIdContract(contractAddr common.Address) {
	instance, err := disableid.NewDisableId(contractAddr, ethSrv.client)
	if err != nil {
		log.Error(err.Error())
	}
	ethSrv.disableidContractAddress = contractAddr
	ethSrv.DisableIdContract = instance
	log.Info("DisableId contract with address " + contractAddr.String() + " loaded")
}

func (ethSrv *EthService) GetAuth() (*bind.TransactOpts, error) {
	file, err := os.Open(ethSrv.KeyStore.Path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	b, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	passw, err := ioutil.ReadFile(ethSrv.KeyStore.Password)
	if err != nil {
		return nil, err
	}
	auth, err := bind.NewTransactor(strings.NewReader(string(b)), string(passw))
	return auth, err
}

func (ethSrv *EthService) GetTx(txHash common.Hash) (*types.Transaction, *types.Receipt, bool, error) {
	receipt, err := ethSrv.client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		log.Error(err)
		// return nil, nil, false, err
	}
	tx, isPending, err := ethSrv.client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Error(err)
		return nil, nil, false, err
	}
	return tx, receipt, isPending, nil
}

func (ethSrv *EthService) GetIdInWhitelist(ethAddr common.Address) (bool, error) {
	res, err := ethSrv.WhitelistContract.IsRegistered(nil, ethAddr)
	if err != nil {
		return false, err
	}

	return res, nil
}

type SampleCallData struct {
	Addr string `json:"addr"`
	Data string `json:"dataHex"`
}

func (ethSrv *EthService) ForwardTxToSampleContract(d SampleCallData) (*types.Transaction, error) {
	nonce, err := ethSrv.Client().PendingNonceAt(context.Background(), ethSrv.acc.Address)
	if err != nil {
		return nil, err
	}

	gasPrice, err := ethSrv.Client().SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	auth, err := ethSrv.GetAuth()
	if err != nil {
		return nil, err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	data, err := hex.DecodeString(d.Data[:2])
	if err != nil {
		return nil, err
	}
	var data32 [32]byte
	copy(data32[:], data)

	ethTx, err := ethSrv.SampleContract.AddAllowed(auth, common.HexToAddress(d.Addr), data32)
	if err != nil {
		return nil, err
	}
	return ethTx, nil
}

type FullVerifierCallData struct {
	A      [2]string    `json:"a"`
	B      [2][2]string `json:"b"`
	C      [2]string    `json:"c"`
	Inputs [10]string   `json:"inputs"`
}

func (ethSrv *EthService) ForwardTxToFullVerifierContract(d FullVerifierCallData) (*types.Transaction, error) {
	nonce, err := ethSrv.Client().PendingNonceAt(context.Background(), ethSrv.acc.Address)
	if err != nil {
		return nil, err
	}

	gasPrice, err := ethSrv.Client().SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	auth, err := ethSrv.GetAuth()
	if err != nil {
		return nil, err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice

	// zkp verifierinputs
	var a [2]*big.Int
	var b [2][2]*big.Int
	var c [2]*big.Int
	var inputs [10]*big.Int

	// a
	a0, ok := new(big.Int).SetString(d.A[0], 0)
	if !ok {
		return nil, errors.New("error parsing hex to bigint a[0]")
	}
	a1, ok := new(big.Int).SetString(d.A[1], 0)
	if !ok {
		return nil, errors.New("error parsing hex to bigint a[1]")
	}
	a = [2]*big.Int{a0, a1}

	// b
	b00, ok := new(big.Int).SetString(d.B[0][0], 0)
	if !ok {
		return nil, errors.New("error parsing hex to bigint b[0][0]")
	}
	b01, ok := new(big.Int).SetString(d.B[0][1], 0)
	if !ok {
		return nil, errors.New("error parsing hex to bigint b[0][1]")
	}
	b10, ok := new(big.Int).SetString(d.B[1][0], 0)
	if !ok {
		return nil, errors.New("error parsing hex to bigint b[1][0]")
	}
	b11, ok := new(big.Int).SetString(d.B[1][1], 0)
	if !ok {
		return nil, errors.New("error parsing hex to bigint b[1][1]")
	}
	b0 := [2]*big.Int{b00, b01}
	b1 := [2]*big.Int{b10, b11}
	b = [2][2]*big.Int{b0, b1}

	// c
	c0, ok := new(big.Int).SetString(d.C[0], 0)
	if !ok {
		return nil, errors.New("error parsing hex to bigint c[0]")
	}
	c1, ok := new(big.Int).SetString(d.C[1], 0)
	if !ok {
		return nil, errors.New("error parsing hex to bigint c[1]")
	}
	c = [2]*big.Int{c0, c1}

	// inputs
	for i, inpStr := range d.Inputs {
		inp, ok := new(big.Int).SetString(inpStr, 0)
		if !ok {
			return nil, errors.New("error parsing hex to bigint inputs[" + strconv.Itoa(i) + "]")
		}
		inputs[i] = inp
	}

	ethTx, err := ethSrv.FullVerifierContract.Verify(auth, a, b, c, inputs)
	if err != nil {
		return nil, err
	}
	return ethTx, nil
}

type DisableIdCallData struct {
	// all in hex format, from js
	Mtp        string
	Id         string
	EthAddress string
	MsgHash    string
	Rsv        string
}

func (ethSrv *EthService) ForwardTxToDisableIdContract(d DisableIdCallData) (*types.Transaction, error) {
	nonce, err := ethSrv.Client().PendingNonceAt(context.Background(), ethSrv.acc.Address)
	if err != nil {
		return nil, err
	}

	gasPrice, err := ethSrv.Client().SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	auth, err := ethSrv.GetAuth()
	if err != nil {
		return nil, err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	// disableid inputs
	var id [31]byte
	var mtp []byte
	var ethAddress common.Address
	var msgHash [32]byte
	var rsv []byte

	idBytes, err := common3.HexDecode(d.Id)
	if err != nil {
		return nil, err
	}
	copy(id[:], idBytes[:31])
	mtp, err = common3.HexDecode(d.Mtp)
	if err != nil {
		return nil, err
	}
	ethAddress = common.HexToAddress(d.EthAddress)
	msgHashBytes, err := common3.HexDecode(d.MsgHash)
	if err != nil {
		return nil, err
	}
	copy(msgHash[:], msgHashBytes[:32])
	rsv, err = common3.HexDecode(d.Rsv)
	if err != nil {
		return nil, err
	}

	ethTx, err := ethSrv.DisableIdContract.DisableIdentity(auth, mtp, id, ethAddress, msgHash, rsv)
	if err != nil {
		return nil, err
	}
	return ethTx, nil
}
