package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// main function
func main() {
	// .env ファイルから環境変数を読み込む
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	clientId := os.Getenv("CLIENT_ID")
	fmt.Println("Client ID:", clientId)

	// create a new gin app
	app := gin.Default()

	// healthcheck endpoint
	// curl -X GET http://localhost:8080/healthcheck
	app.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(200, gin.H{})
	})

	// start the server
	err = app.Run(":8080")
	if err != nil {
		return
	}
}
