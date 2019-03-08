package endpoint

import (
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"github.com/iden3/tx-forwarder/eth"
)

func fail(c *gin.Context, err error) {
	color.Red("error: " + err.Error())
	c.JSON(400, gin.H{
		"error": err.Error(),
	})
}

func handleGetInfo(c *gin.Context) {
	c.JSON(200, gin.H{})
}

func handlePostTx(c *gin.Context) {
	var tx eth.TxMsg
	c.BindJSON(&tx)

	// server side verification here

	ethTx, err := ethsrv.ForwardTx(tx)
	if err != nil {
		fail(c, err)
		return
	}
	c.JSON(200, gin.H{
		"ethTx": ethTx.Hash().Hex(),
	})
}
