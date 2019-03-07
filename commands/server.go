package commands

import (
	"github.com/iden3/gas-station/config"
	"github.com/iden3/gas-station/endpoint"
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
}

func cmdStart(c *cli.Context) error {
	if err := config.MustRead(c); err != nil {
		return err
	}

	ks, acc := config.LoadKeyStore()
	ethSrv := config.LoadWeb3(ks, &acc)
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
