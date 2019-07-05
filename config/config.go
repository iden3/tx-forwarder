package config

import (
	"github.com/spf13/viper"
	"github.com/urfave/cli"
)

type Config struct {
	Server struct {
		ServiceApi string
		AdminApi   string
	}
	Web3 struct {
		Url string
	}
	KeyStore struct {
		Path        string
		Address     string
		Password    string
		KeyJsonPath string
	}
	Contracts struct {
		SampleContract       string
		DisableIdContract    string
		Mimc7Contract        string
		RootCommitsContract  string
		WhitelistContract    string
		FullVerifierContract string
	}
	Ids struct {
		Certifier string
		Storer    string
	}
}

var C Config

func MustRead(c *cli.Context) error {

	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath(".") // adding home directory as first search path

	if c.GlobalString("config") != "" {
		viper.SetConfigFile(c.GlobalString("config"))
	}

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	err := viper.Unmarshal(&C)
	if err != nil {
		return err
	}
	return nil
}
