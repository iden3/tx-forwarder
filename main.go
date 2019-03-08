package main

import (
	"os"

	"github.com/iden3/tx-forwarder/commands"
	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = "tx-forwarder"
	app.Version = "0.0.1-alpha"
	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "config"},
	}

	app.Commands = []cli.Command{}
	app.Commands = append(app.Commands, commands.ServerCommands...)

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
