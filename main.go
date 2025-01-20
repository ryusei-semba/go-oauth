package main

import (
	"fmt"
	"go-oauth/infrastructure/database"
	"go-oauth/usecase"
	"log"
	"net/http"
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

	// トークン発行エンドポイント
	app.POST("/oauth/token", func(c *gin.Context) {
		response, err := tokenUsecase.RequestToken()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, response)
	})

	// トークン検証エンドポイント
	app.GET("/oauth/validate", func(c *gin.Context) {
		accessToken := c.GetHeader("Authorization")
		if accessToken == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "access token is required",
			})
			return
		}

		// Bearer プレフィックスを除去
		if len(accessToken) > 7 && accessToken[:7] == "Bearer " {
			accessToken = accessToken[7:]
		}

		token, err := tokenUsecase.ValidateToken(accessToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"valid": true,
			"token": token,
		})
	})

	// トークン更新エンドポイント
	app.POST("/oauth/token/refresh", func(c *gin.Context) {
		var req struct {
			RefreshToken string `json:"refresh_token" binding:"required"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "refresh token is required",
			})
			return
		}

		response, err := tokenUsecase.RefreshToken(req.RefreshToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, response)
	})

	// トークン無効化エンドポイント
	app.POST("/oauth/token/revoke", func(c *gin.Context) {
		accessToken := c.GetHeader("Authorization")
		if accessToken == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "access token is required",
			})
			return
		}

		// Bearer プレフィックスを除去
		if len(accessToken) > 7 && accessToken[:7] == "Bearer " {
			accessToken = accessToken[7:]
		}

		if err := tokenUsecase.RevokeToken(accessToken); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "token revoked",
		})
	})

	if err := app.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
