package interfaces

import (
	"go-oauth/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Route ルーティングを設定する
func Route(app *gin.Engine, tokenUsecase *usecase.TokenUsecase) {
	oauth := app.Group("/oauth")
	{
		// トークン発行エンドポイント
		oauth.POST("/token", func(c *gin.Context) {
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
		oauth.GET("/validate", func(c *gin.Context) {
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
		oauth.POST("/token/refresh", func(c *gin.Context) {
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
		oauth.POST("/token/revoke", func(c *gin.Context) {
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
	}
}
