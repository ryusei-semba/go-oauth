package interfaces

import (
	"go-oauth/usecase"

	"github.com/gin-gonic/gin"
)

func HealthcheckHandler(c *gin.Context) {
	c.JSON(200, gin.H{})
}

func TokenHandler(c *gin.Context) {
	res, err := usecase.RequestToken()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, res)
	}
}
