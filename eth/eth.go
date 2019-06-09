package eth

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
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
	samplecontract "github.com/iden3/tx-forwarder/eth/contracts/samplecontract"
	zkpverifier "github.com/iden3/tx-forwarder/eth/contracts/zkpverifier"
	log "github.com/sirupsen/logrus"
)

type EthService struct {
	ks                         *keystore.KeyStore
	acc                        *accounts.Account
	client                     *ethclient.Client
	SampleContract             *samplecontract.SampleContract
	ZKPVerifierContract        *zkpverifier.CheckFullCircuit
	sampleContractAddress      common.Address
	zkpverifierContractAddress common.Address
	KeyStore                   struct {
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

func (ethSrv *EthService) ZKPVerifierContractAddress() common.Address {
	return ethSrv.zkpverifierContractAddress
}
func (ethSrv *EthService) SampleContractAddress() common.Address {
	return ethSrv.sampleContractAddress
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

func (ethSrv *EthService) DeploySampleContract() error {
	auth, err := ethSrv.GetAuth()
	if err != nil {
		return err
	}

	address, tx, _, err := samplecontract.DeploySampleContract(auth, ethSrv.client)
	if err != nil {
		return err
	}
	ethSrv.sampleContractAddress = address

	log.Info("sample contract deployed at address: " + address.Hex())
	log.Info("deployment transaction: " + tx.Hash().Hex())

	return nil
}

func (ethSrv *EthService) DeployZKPVerifierContract() error {
	auth, err := ethSrv.GetAuth()
	if err != nil {
		return err
	}

	address, tx, _, err := zkpverifier.DeployVerifier(auth, ethSrv.client)
	if err != nil {
		return err
	}
	ethSrv.zkpverifierContractAddress = address

	log.Info("zkpverifier contract deployed at address: " + address.Hex())
	log.Info("deployment transaction: " + tx.Hash().Hex())
	return nil
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

func (ethSrv *EthService) LoadZKPVerifierContract(contractAddr common.Address) {
	instance, err := zkpverifier.NewCheckFullCircuit(contractAddr, ethSrv.client)
	if err != nil {
		log.Error(err.Error())
	}
	ethSrv.zkpverifierContractAddress = contractAddr
	ethSrv.ZKPVerifierContract = instance
	log.Info("ZKPVerifier contract with address " + contractAddr.String() + " loaded")
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

	fmt.Println(d.Data)
	data, err := hex.DecodeString(d.Data[:2])
	if err != nil {
		return nil, err
	}
	fmt.Println(data)
	var data32 [32]byte
	copy(data32[:], data)

	ethTx, err := ethSrv.SampleContract.AddAllowed(auth, common.HexToAddress(d.Addr), data32)
	if err != nil {
		return nil, err
	}
	return ethTx, nil
}

type ZKPVerifierCallData struct {
	A      [2]string    `json:"a"`
	B      [2][2]string `json:"b"`
	C      [2]string    `json:"c"`
	Inputs [11]string   `json:"inputs"`
}

func (ethSrv *EthService) ForwardTxToZKPVerifierContract(d ZKPVerifierCallData) (*types.Transaction, error) {
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

	// zkp verifierinputs
	var a [2]*big.Int
	var b [2][2]*big.Int
	var c [2]*big.Int
	var inputs [11]*big.Int

	// a
	a0, ok := new(big.Int).SetString(d.A[0], 0)
	if !ok {
		fmt.Println(d.A[0])
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

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(inputs)

	ethTx, err := ethSrv.ZKPVerifierContract.Verify(auth, a, b, c, inputs)
	if err != nil {
		return nil, err
	}
	return ethTx, nil
}
