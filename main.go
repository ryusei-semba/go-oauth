package main

import "github.com/gin-gonic/gin"

// main function
func main() {
	// create a new gin app
	app := gin.Default()

	// healthcheck endpoint
	// curl -X GET http://localhost:8080/healthcheck
	app.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(200, gin.H{})
	})

	// start the server
	err := app.Run(":8080")
	if err != nil {
		return
	}
}
