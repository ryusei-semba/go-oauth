package token

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

const (
	// AccessTokenLength アクセストークンの長さ（バイト）
	AccessTokenLength = 32
	// RefreshTokenLength リフレッシュトークンの長さ（バイト）
	RefreshTokenLength = 32
)

// Generator トークン生成のインターフェース
type Generator interface {
	GenerateAccessToken() (string, error)
	GenerateRefreshToken() (string, error)
}

// DefaultGenerator デフォルトのトークン生成実装
type DefaultGenerator struct{}

// NewGenerator Generatorのインスタンスを生成する
func NewGenerator() Generator {
	return &DefaultGenerator{}
}

// GenerateAccessToken アクセストークンを生成する
func (g *DefaultGenerator) GenerateAccessToken() (string, error) {
	return g.generateSecureToken(AccessTokenLength)
}

// GenerateRefreshToken リフレッシュトークンを生成する
func (g *DefaultGenerator) GenerateRefreshToken() (string, error) {
	return g.generateSecureToken(RefreshTokenLength)
}

// generateSecureToken 暗号学的に安全なランダムトークンを生成する
func (g *DefaultGenerator) generateSecureToken(length int) (string, error) {
	// 指定された長さのバイト列を生成
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", fmt.Errorf("failed to generate random bytes: %w", err)
	}

	// Base64エンコード（URLセーフ）
	token := base64.URLEncoding.EncodeToString(bytes)

	return token, nil
}
