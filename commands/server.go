package commands

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/iden3/go-iden3/core"
	"github.com/iden3/tx-forwarder/config"
	"github.com/iden3/tx-forwarder/endpoint"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var ServerCommands = []cli.Command{
	{
		Name:    "start",
		Aliases: []string{},
		Usage:   "start the server",
		Action:  cmdStart,
	},
	{
		Name:    "stop",
		Aliases: []string{},
		Usage:   "stops the server",
		Action:  cmdStop,
	},
	{
		Name:    "info",
		Aliases: []string{},
		Usage:   "server status",
		Action:  cmdInfo,
	},
	{
		Name:  "deploy",
		Usage: "deploy smart contracts",
		Subcommands: []cli.Command{
			{
				Name:    "sample",
				Aliases: []string{},
				Usage:   "deploy sample contract",
				Action:  cmdDeploySampleContract,
			},
			{
				Name:    "rootcommits",
				Aliases: []string{},
				Usage:   "deploy rootcommits contract",
				Action:  cmdDeployRootCommitsContract,
			},
			{
				Name:    "whitelist",
				Aliases: []string{},
				Usage:   "deploy whitelist contract",
				Action:  cmdDeployWhitelistContract,
			},
			{
				Name:    "fullverifier",
				Aliases: []string{},
				Usage:   "deploy fullverifier contract",
				Action:  cmdDeployFullVerifierContract,
			},
			{
				Name:    "disableid",
				Aliases: []string{},
				Usage:   "deploy disableid contract",
				Action:  cmdDeployDisableIdContract,
			},
		},
	},
}

func cmdStart(c *cli.Context) error {
	if err := config.MustRead(c); err != nil {
		return err
	}

	ks, acc := config.LoadKeyStore()
	ethSrv := config.LoadWeb3(ks, &acc)

	sampleContractAddr := common.HexToAddress(config.C.Contracts.SampleContract)
	zkpverifierContractAddr := common.HexToAddress(config.C.Contracts.FullVerifierContract)
	whitelistContractAddr := common.HexToAddress(config.C.Contracts.WhitelistContract)
	disableidContractAddr := common.HexToAddress(config.C.Contracts.DisableIdContract)

	ethSrv.LoadSampleContract(sampleContractAddr)
	ethSrv.LoadWhitelistContract(whitelistContractAddr)
	ethSrv.LoadFullVerifierContract(zkpverifierContractAddr)
	ethSrv.LoadDisableIdContract(disableidContractAddr)

	endpoint.Serve(ethSrv)

	return nil
}

func cmdStop(c *cli.Context) error {
	if err := config.MustRead(c); err != nil {
		return err
	}
	return nil
}

func cmdInfo(c *cli.Context) error {
	if err := config.MustRead(c); err != nil {
		return err
	}
	fmt.Println(c)
	return nil
}

func cmdDeploySampleContract(c *cli.Context) error {
	if err := config.MustRead(c); err != nil {
		return err
	}

	ks, acc := config.LoadKeyStore()
	ethSrv := config.LoadWeb3(ks, &acc)

	contractAddress, tx, err := ethSrv.DeploySampleContract()
	if err != nil {
		return err
	}
	log.Info("sample contract deployed at address: " + contractAddress.Hex())
	log.Info("deployment transaction: " + tx.Hash().Hex())

	return nil
}

func cmdDeployRootCommitsContract(c *cli.Context) error {
	if err := config.MustRead(c); err != nil {
		return err
	}

	ks, acc := config.LoadKeyStore()
	ethSrv := config.LoadWeb3(ks, &acc)

	mimc7Address := common.HexToAddress(config.C.Contracts.Mimc7Contract)

	contractAddress, tx, err := ethSrv.DeployRootCommitsContract(mimc7Address)
	if err != nil {
		return err
	}
	log.Info("rootcommits contract deployed at address: " + contractAddress.Hex())
	log.Info("deployment transaction: " + tx.Hash().Hex())

	return nil
}
func cmdDeployWhitelistContract(c *cli.Context) error {
	if err := config.MustRead(c); err != nil {
		return err
	}

	ks, acc := config.LoadKeyStore()
	ethSrv := config.LoadWeb3(ks, &acc)

	contractAddress, tx, err := ethSrv.DeployWhitelistContract()
	if err != nil {
		return err
	}
	log.Info("whitelist contract deployed at address: " + contractAddress.Hex())
	log.Info("deployment transaction: " + tx.Hash().Hex())

	return nil
}

func cmdDeployFullVerifierContract(c *cli.Context) error {
	if err := config.MustRead(c); err != nil {
		return err
	}

	ks, acc := config.LoadKeyStore()
	ethSrv := config.LoadWeb3(ks, &acc)

	rootcommitsContract := common.HexToAddress(config.C.Contracts.RootCommitsContract)
	whitelistContract := common.HexToAddress(config.C.Contracts.WhitelistContract)
	idCertifier, err := core.IDFromString(config.C.Ids.Certifier)
	if err != nil {
		return err
	}
	idStorer, err := core.IDFromString(config.C.Ids.Storer)
	if err != nil {
		return err
	}

	contractAddress, tx, smVerifierAddr, smVerifierTx, err := ethSrv.DeployFullVerifierContract(rootcommitsContract, whitelistContract, idCertifier, idStorer)
	if err != nil {
		return err
	}
	log.Info("zkverifier contract deployed at address: " + smVerifierAddr.Hex())
	log.Info("deployment transaction: " + smVerifierTx.Hash().Hex())
	log.Info("fullverifier contract deployed at address: " + contractAddress.Hex())
	log.Info("deployment transaction: " + tx.Hash().Hex())

	return nil
}

func cmdDeployDisableIdContract(c *cli.Context) error {
	if err := config.MustRead(c); err != nil {
		return err
	}

	ks, acc := config.LoadKeyStore()
	ethSrv := config.LoadWeb3(ks, &acc)

	log.Info("deploying smart contracts: Mimc7, Iden3Helpers, DisableId")

	var contractAddress *common.Address
	var tx *types.Transaction
	var err error
	if config.C.Contracts.Mimc7Contract == "" {
		contractAddress, tx, err = ethSrv.DeployMimc7Contract()
		if err != nil {
			return err
		}
		log.Info("Mimc7 contract deployed at address: " + contractAddress.Hex())
		log.Info("deployment transaction: " + tx.Hash().Hex())
	} else {
		parsedAddr := common.HexToAddress(config.C.Contracts.Mimc7Contract)
		contractAddress = &parsedAddr
	}

	contractAddress, tx, err = ethSrv.DeployIden3HelpersContract(*contractAddress)
	if err != nil {
		return err
	}
	log.Info("Iden3Helpers contract deployed at address: " + contractAddress.Hex())
	log.Info("deployment transaction: " + tx.Hash().Hex())

	contractAddress, tx, err = ethSrv.DeployDisableIdContract(*contractAddress)
	if err != nil {
		return err
	}
	log.Info("DisableId contract deployed at address: " + contractAddress.Hex())
	log.Info("deployment transaction: " + tx.Hash().Hex())

	return nil
}
