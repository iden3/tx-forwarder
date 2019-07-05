module github.com/iden3/tx-forwarder

go 1.12

require (
	github.com/ethereum/go-ethereum v1.8.29-0.20190620093831-25c3282cf126
	github.com/gin-contrib/cors v1.3.0
	github.com/gin-gonic/gin v1.4.0
	github.com/iden3/go-iden3 v0.0.5
	github.com/mgutz/ansi v0.0.0-20170206155736-9520e82c474b // indirect
	github.com/mgutz/logxi v0.0.0-20161027140823-aebf8a7d67ab
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/viper v1.4.0
	github.com/urfave/cli v1.20.0
)

replace github.com/iden3/go-iden3 => ../go-iden3
