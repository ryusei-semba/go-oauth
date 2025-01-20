package main

import (
	"fmt"
	"go-oauth/infrastructure/database"
	"go-oauth/interfaces"
	"go-oauth/usecase"
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

	// データベース接続を初期化
	db, err := database.NewDBConnection()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// データベースの初期化
	initializer := database.NewDBInitializer(db)
	if err := initializer.InitializeDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// トークンユースケースを初期化
	tokenUsecase, err := usecase.NewTokenUsecase(db)
	if err != nil {
		log.Fatalf("Failed to initialize token usecase: %v", err)
	}

	// Ginエンジンを初期化
	app := gin.Default()

	// ルーティングを設定
	interfaces.Route(app, tokenUsecase)

	// サーバーを起動
	if err := app.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
