package main

import (
	"fmt"
	"go-oauth/interfaces"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	clientId := os.Getenv("CLIENT_ID")
	fmt.Println("Client ID:", clientId)

	app := gin.Default()

	interfaces.Route(app)

	err = app.Run(":8080")
	if err != nil {
		return
	}
}
