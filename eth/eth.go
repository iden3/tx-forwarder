package eth

import (
	"context"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	contracts "github.com/iden3/tx-forwarder/eth/contracts"
	log "github.com/sirupsen/logrus"
)

type EthService struct {
	ks              *keystore.KeyStore
	acc             *accounts.Account
	client          *ethclient.Client
	SampleContract  *contracts.SampleContract
	contractAddress common.Address
	KeyStore        struct {
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

func (ethSrv *EthService) ContractAddress() common.Address {
	return ethSrv.contractAddress
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

func (ethSrv *EthService) DeployContract() error {
	auth, err := ethSrv.GetAuth()
	if err != nil {
		return err
	}

	address, tx, _, err := contracts.DeploySampleContract(auth, ethSrv.client)
	if err != nil {
		return err
	}
	ethSrv.contractAddress = address

	log.Info("sample contract deployed at address: " + address.Hex())
	log.Info("deployment transaction: " + tx.Hash().Hex())
	return nil
}

func (ethSrv *EthService) LoadContract(contractAddr common.Address) {
	instance, err := contracts.NewSampleContract(contractAddr, ethSrv.client)
	if err != nil {
		log.Error(err.Error())
	}
	ethSrv.contractAddress = contractAddr
	ethSrv.SampleContract = instance
	log.Info("Contract with address " + contractAddr.String() + " loaded")
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

type TxMsg struct {
	Addr string `json:"addr"`
	Data string `json:"dataHex"`
}

func (ethSrv *EthService) ForwardTx(tx TxMsg) (*types.Transaction, error) {
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

	fmt.Println(tx.Data)
	data, err := hex.DecodeString(tx.Data[:2])
	if err != nil {
		return nil, err
	}
	fmt.Println(data)
	var data32 [32]byte
	copy(data32[:], data)

	ethTx, err := ethSrv.SampleContract.AddAllowed(auth, common.HexToAddress(tx.Addr), data32)
	if err != nil {
		return nil, err
	}
	return ethTx, nil
}
