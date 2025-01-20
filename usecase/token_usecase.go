package usecase

import (
	"fmt"
	"go-oauth/domain/token"
	"go-oauth/infrastructure/database"
)

// TokenUsecase トークン関連のユースケース
type TokenUsecase struct {
	tokenService *token.Service
}

// NewTokenUsecase TokenUsecaseのインスタンスを生成する
func NewTokenUsecase(db *database.DBConnection) (*TokenUsecase, error) {
	// データベースを初期化
	initializer := database.NewDBInitializer(db)
	if err := initializer.InitializeDB(); err != nil {
		return nil, fmt.Errorf("failed to initialize database: %w", err)
	}

	// リポジトリとサービスを初期化
	repo := database.NewTokenRepository(db)
	service := token.NewService(repo)

	return &TokenUsecase{
		tokenService: service,
	}, nil
}

// RequestToken トークンを要求する
func (u *TokenUsecase) RequestToken() (*token.TokenResponse, error) {
	// TODO: クライアント認証とユーザー認証の実装
	clientID := "sample_client_id"
	userID := "sample_user_id"
	scope := "read write"

	// トークンを発行
	response, err := u.tokenService.IssueToken(clientID, userID, scope)
	if err != nil {
		return nil, fmt.Errorf("failed to issue token: %w", err)
	}

	return response, nil
}

// RefreshToken トークンを更新する
func (u *TokenUsecase) RefreshToken(refreshToken string) (*token.TokenResponse, error) {
	response, err := u.tokenService.RefreshToken(refreshToken)
	if err != nil {
		return nil, fmt.Errorf("failed to refresh token: %w", err)
	}

	return response, nil
}

// RevokeToken トークンを無効化する
func (u *TokenUsecase) RevokeToken(accessToken string) error {
	if err := u.tokenService.RevokeToken(accessToken); err != nil {
		return fmt.Errorf("failed to revoke token: %w", err)
	}

	return nil
}

// ValidateToken トークンを検証する
func (u *TokenUsecase) ValidateToken(accessToken string) (*token.Token, error) {
	validToken, err := u.tokenService.ValidateToken(accessToken)
	if err != nil {
		return nil, fmt.Errorf("failed to validate token: %w", err)
	}

	return validToken, nil
}
