package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/iden3/tx-forwarder/eth"
	log "github.com/sirupsen/logrus"
)

const (
	passwdPrefix = "passwd:"
	filePrefix   = "file:"
)

func Assert(msg string, err error) {
	if err != nil {
		log.Error(msg, " ", err.Error())
		os.Exit(1)
	}
}

func LoadKeyStore() (*keystore.KeyStore, accounts.Account) {
	var err error
	var passwd string

	// Load keystore
	ks := keystore.NewKeyStore(C.KeyStore.Path, keystore.StandardScryptN, keystore.StandardScryptP)

	// Password can be prefixed by two options
	//   file: <path to file containing the password>
	//   passwd: raw password
	// if is not prefixed by any of those, file: is used
	if strings.HasPrefix(C.KeyStore.Password, passwdPrefix) {
		passwd = C.KeyStore.Password[len(passwdPrefix):]
	} else {
		filename := C.KeyStore.Password
		if strings.HasPrefix(filename, filePrefix) {
			filename = C.KeyStore.Password[len(filePrefix):]
		}
		passwdbytes, err := ioutil.ReadFile(filename)
		Assert("Cannot read password ", err)
		passwd = string(passwdbytes)
	}

	acc, err := ks.Find(accounts.Account{
		Address: common.HexToAddress(C.KeyStore.Address),
	})
	Assert("Cannot find keystore account", err)

	Assert("Cannot unlock account", ks.Unlock(acc, string(passwd)))
	log.WithField("acc", acc.Address.Hex()).Info("Keystore and account unlocked successfully")

	return ks, acc
}

func LoadWeb3(ks *keystore.KeyStore, acc *accounts.Account) *eth.EthService {
	// Create geth client
	url := C.Web3.Url
	hidden := strings.HasPrefix(url, "hidden:")
	if hidden {
		url = url[len("hidden:"):]
	}
	ethsrv := eth.NewEthService(url, ks, acc, C.KeyStore.KeyJsonPath, C.KeyStore.Password)
	if hidden {
		log.WithField("url", "(hidden)").Info("Connection to web3 server opened")
	} else {
		log.WithField("url", C.Web3.Url).Info("Connection to web3 server opened")
	}
	balance, err := ethsrv.GetBalance(acc.Address)
	if err != nil {
		fmt.Println("error getting balance")
		os.Exit(0)
	}
	log.Info("Current balance " + balance.String() + " ETH")
	return ethsrv
}
