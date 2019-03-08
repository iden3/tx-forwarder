package commands

import (
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
		Name:    "deploy",
		Aliases: []string{},
		Usage:   "deploy contract",
		Action:  cmdDeployContract,
	},
}

func cmdStart(c *cli.Context) error {
	if err := config.MustRead(c); err != nil {
		return err
	}

	ks, acc := config.LoadKeyStore()
	ethSrv := config.LoadWeb3(ks, &acc)

	contractAddr := common.HexToAddress(config.C.Contracts.SampleContract)
	ethSrv.LoadContract(contractAddr)

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
	return nil
}
func cmdDeployContract(c *cli.Context) error {
	if err := config.MustRead(c); err != nil {
		return err
	}

	ks, acc := config.LoadKeyStore()
	ethSrv := config.LoadWeb3(ks, &acc)

	err := ethSrv.DeployContract()

	return err
}
