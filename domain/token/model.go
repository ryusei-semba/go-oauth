package token

import (
	"fmt"
	"time"
)

// Token OAuthトークンを表すエンティティ
type Token struct {
	AccessToken  string
	RefreshToken string
	ClientID     string
	UserID       string
	ExpiresAt    time.Time
	Scope        string
	CreatedAt    time.Time
}

// TokenResponse OAuth2.0仕様に基づくトークンレスポンス
type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	Scope        string `json:"scope"`
	RefreshToken string `json:"refresh_token"`
}

// NewToken 新しいトークンを生成する
func NewToken(clientID, userID, scope string) (*Token, error) {
	generator := NewGenerator()

	accessToken, err := generator.GenerateAccessToken()
	if err != nil {
		return nil, fmt.Errorf("failed to generate access token: %w", err)
	}

	refreshToken, err := generator.GenerateRefreshToken()
	if err != nil {
		return nil, fmt.Errorf("failed to generate refresh token: %w", err)
	}

	now := time.Now()
	expiresAt := now.Add(1 * time.Hour) // トークンの有効期限は1時間

	return &Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ClientID:     clientID,
		UserID:       userID,
		ExpiresAt:    expiresAt,
		Scope:        scope,
		CreatedAt:    now,
	}, nil
}

// ToResponse TokenResponseに変換する
func (t *Token) ToResponse() TokenResponse {
	return TokenResponse{
		AccessToken:  t.AccessToken,
		TokenType:    "Bearer",
		ExpiresIn:    int(time.Until(t.ExpiresAt).Seconds()),
		Scope:        t.Scope,
		RefreshToken: t.RefreshToken,
	}
}

// IsExpired トークンが有効期限切れかどうかを判定する
func (t *Token) IsExpired() bool {
	return time.Now().After(t.ExpiresAt)
}

// ValidateScope スコープが有効かどうかを検証する
func (t *Token) ValidateScope(requiredScope string) bool {
	// TODO: スコープの検証ロジックを実装
	// 現在は単純な文字列比較だが、より複雑なスコープ検証が必要になる可能性がある
	return t.Scope == requiredScope
}
