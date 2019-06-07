package commands

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/iden3/tx-forwarder/config"
	"github.com/iden3/tx-forwarder/endpoint"
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
		Name:    "deploysample",
		Aliases: []string{},
		Usage:   "deploy sample contract",
		Action:  cmdDeploySampleContract,
	},
	{
		Name:    "deployzkpverifier",
		Aliases: []string{},
		Usage:   "deploy zkpverifier contract",
		Action:  cmdDeployZKPVerifierContract,
	},
}

func cmdStart(c *cli.Context) error {
	if err := config.MustRead(c); err != nil {
		return err
	}

	ks, acc := config.LoadKeyStore()
	ethSrv := config.LoadWeb3(ks, &acc)

	sampleContractAddr := common.HexToAddress(config.C.Contracts.SampleContract)
	zkpverifierContractAddr := common.HexToAddress(config.C.Contracts.SampleContract)
	ethSrv.LoadSampleContract(sampleContractAddr)
	ethSrv.LoadZKPVerifierContract(zkpverifierContractAddr)

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

	err := ethSrv.DeploySampleContract()

	return err
}
func cmdDeployZKPVerifierContract(c *cli.Context) error {
	if err := config.MustRead(c); err != nil {
		return err
	}

	ks, acc := config.LoadKeyStore()
	ethSrv := config.LoadWeb3(ks, &acc)

	err := ethSrv.DeployZKPVerifierContract()

	return err
}
