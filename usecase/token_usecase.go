package usecase

import (
	"fmt"
	"go-oauth/infrastructure/database"
	"time"
)

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
}

func RequestToken() (TokenResponse, error) {
	// データベース接続を取得
	conn, err := database.NewDBConnection()
	if err != nil {
		return TokenResponse{}, fmt.Errorf("failed to connect to database: %w", err)
	}
	defer conn.Close()

	// トークンリポジトリを初期化
	repo := database.NewTokenRepository(conn)

	// トークンテーブルを初期化
	if err := repo.InitTokenTable(); err != nil {
		return TokenResponse{}, fmt.Errorf("failed to init token table: %w", err)
	}

	// トークンの有効期限を設定（1時間）
	expiresIn := 3600
	expiresAt := time.Now().Add(time.Duration(expiresIn) * time.Second)

	// トークンを保存
	token := struct {
		AccessToken  string
		RefreshToken string
		ClientID     string
		UserID       string
		ExpiresAt    time.Time
		Scope        string
	}{
		AccessToken:  "sample_access_token",  // 本来はランダムな文字列を生成
		RefreshToken: "sample_refresh_token", // 本来はランダムな文字列を生成
		ClientID:     "sample_client_id",
		UserID:       "sample_user_id",
		ExpiresAt:    expiresAt,
		Scope:        "read write",
	}

	if err := repo.SaveToken(token); err != nil {
		return TokenResponse{}, fmt.Errorf("failed to save token: %w", err)
	}

	return TokenResponse{
		AccessToken: token.AccessToken,
		TokenType:   "Bearer",
		ExpiresIn:   expiresIn,
		Scope:       token.Scope,
	}, nil
}
