package endpoint

import (
	"github.com/gin-gonic/gin"
)

func handleGetInfo(c *gin.Context) {
	c.JSON(200, gin.H{})
}
