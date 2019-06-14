package endpoint

import (
	"github.com/gin-gonic/gin"
	"github.com/iden3/tx-forwarder/eth"
	log "github.com/sirupsen/logrus"
)

func fail(c *gin.Context, err error) {
	log.Error("error: " + err.Error())
	c.JSON(400, gin.H{
		"error": err.Error(),
	})
}

func handleGetInfo(c *gin.Context) {
	c.JSON(200, gin.H{})
}

func handlePostTxSampleContract(c *gin.Context) {
	var d eth.SampleCallData
	err := c.BindJSON(&d)
	if err != nil {
		fail(c, err)
		return
	}

	ethTx, err := ethsrv.ForwardTxToSampleContract(d)
	if err != nil {
		fail(c, err)
		return
	}
	c.JSON(200, gin.H{
		"ethTx": ethTx.Hash().Hex(),
	})
}

func handlePostTxZKPVerifier(c *gin.Context) {
	var d eth.ZKPVerifierCallData
	err := c.BindJSON(&d)
	if err != nil {
		fail(c, err)
		return
	}

	ethTx, err := ethsrv.ForwardTxToZKPVerifierContract(d)
	if err != nil {
		fail(c, err)
		return
	}
	c.JSON(200, gin.H{
		"ethTx": ethTx.Hash().Hex(),
	})
}

func handlePostTxDisableId(c *gin.Context) {
	var d eth.DisableIdCallData
	err := c.BindJSON(&d)
	if err != nil {
		fail(c, err)
		return
	}

	ethTx, err := ethsrv.ForwardTxToDisableIdContract(d)
	if err != nil {
		fail(c, err)
		return
	}
	c.JSON(200, gin.H{
		"ethTx": ethTx.Hash().Hex(),
	})
}
