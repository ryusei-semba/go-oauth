package token

import "fmt"

// Service トークン関連のドメインサービス
type Service struct {
	repository Repository
}

// NewService トークンサービスのインスタンスを生成する
func NewService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

// IssueToken 新しいトークンを発行する
func (s *Service) IssueToken(clientID, userID, scope string) (*TokenResponse, error) {
	// 新しいトークンを生成
	token, err := NewToken(clientID, userID, scope)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	// トークンを保存
	if err := s.repository.Save(token); err != nil {
		return nil, fmt.Errorf("failed to save token: %w", err)
	}

	// レスポンスを生成
	response := token.ToResponse()
	return &response, nil
}

// ValidateToken トークンを検証する
func (s *Service) ValidateToken(accessToken string) (*Token, error) {
	// トークンを取得
	token, err := s.repository.FindByAccessToken(accessToken)
	if err != nil {
		return nil, fmt.Errorf("failed to find token: %w", err)
	}

	// 有効期限を検証
	if token.IsExpired() {
		return nil, fmt.Errorf("token is expired")
	}

	return token, nil
}

// RefreshToken リフレッシュトークンを使用して新しいトークンを発行する
func (s *Service) RefreshToken(refreshToken string) (*TokenResponse, error) {
	// リフレッシュトークンからトークンを取得
	oldToken, err := s.repository.FindByRefreshToken(refreshToken)
	if err != nil {
		return nil, fmt.Errorf("failed to find token by refresh token: %w", err)
	}

	// 古いトークンを削除
	if err := s.repository.Delete(oldToken.AccessToken); err != nil {
		return nil, fmt.Errorf("failed to delete old token: %w", err)
	}

	// 新しいトークンを発行
	return s.IssueToken(oldToken.ClientID, oldToken.UserID, oldToken.Scope)
}

// RevokeToken トークンを無効化する
func (s *Service) RevokeToken(accessToken string) error {
	if err := s.repository.Delete(accessToken); err != nil {
		return fmt.Errorf("failed to revoke token: %w", err)
	}
	return nil
}
