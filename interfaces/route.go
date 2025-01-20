package interfaces

import (
	"github.com/gin-gonic/gin"
)

func Route(app *gin.Engine) {
	app.GET("/healthcheck", HealthcheckHandler)
	app.POST("/token", TokenHandler)
}
