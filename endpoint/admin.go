package endpoint

import "github.com/gin-gonic/gin"

func handleInfo(c *gin.Context) {
	balance, err := ethsrv.GetBalance(ethsrv.Account().Address)
	if err != nil {
		fail(c, err)
		return
	}
	c.JSON(200, gin.H{
		"ethBalance":                 balance.String(),
		"serverAddress":              ethsrv.Account().Address,
		"sampleContractAddress":      ethsrv.SampleContractAddress().String(),
		"zkpverifierContractAddress": ethsrv.FullVerifierContractAddress().String(),
	})
}
