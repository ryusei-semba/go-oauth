package database

import (
	"fmt"
	"time"
)

// TokenRepository トークン関連のデータベース操作を行うリポジトリ
type TokenRepository struct {
	conn *DBConnection
}

// NewTokenRepository TokenRepositoryのインスタンスを生成する
func NewTokenRepository(conn *DBConnection) *TokenRepository {
	return &TokenRepository{
		conn: conn,
	}
}

// InitTokenTable トークンテーブルを初期化する
func (r *TokenRepository) InitTokenTable() error {
	query := `
		CREATE TABLE IF NOT EXISTS oauth_tokens (
			access_token VARCHAR PRIMARY KEY,
			refresh_token VARCHAR,
			client_id VARCHAR,
			user_id VARCHAR,
			expires_at TIMESTAMP,
			scope VARCHAR,
			created_at TIMESTAMP
		)
	`
	_, err := r.conn.DB.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create oauth_tokens table: %w", err)
	}
	return nil
}

// SaveToken アクセストークンを保存する
func (r *TokenRepository) SaveToken(token struct {
	AccessToken  string
	RefreshToken string
	ClientID     string
	UserID       string
	ExpiresAt    time.Time
	Scope        string
}) error {
	query := `
		INSERT INTO oauth_tokens (
			access_token,
			refresh_token,
			client_id,
			user_id,
			expires_at,
			scope,
			created_at
		) VALUES (?, ?, ?, ?, ?, ?, ?)
	`
	_, err := r.conn.DB.Exec(
		query,
		token.AccessToken,
		token.RefreshToken,
		token.ClientID,
		token.UserID,
		token.ExpiresAt,
		token.Scope,
		time.Now(),
	)
	if err != nil {
		return fmt.Errorf("failed to save token: %w", err)
	}
	return nil
}

// GetToken アクセストークンを取得する
func (r *TokenRepository) GetToken(accessToken string) (*struct {
	AccessToken  string
	RefreshToken string
	ClientID     string
	UserID       string
	ExpiresAt    time.Time
	Scope        string
	CreatedAt    time.Time
}, error) {
	query := `
		SELECT
			access_token,
			refresh_token,
			client_id,
			user_id,
			expires_at,
			scope,
			created_at
		FROM oauth_tokens
		WHERE access_token = ?
	`

	var token struct {
		AccessToken  string
		RefreshToken string
		ClientID     string
		UserID       string
		ExpiresAt    time.Time
		Scope        string
		CreatedAt    time.Time
	}

	err := r.conn.DB.QueryRow(query, accessToken).Scan(
		&token.AccessToken,
		&token.RefreshToken,
		&token.ClientID,
		&token.UserID,
		&token.ExpiresAt,
		&token.Scope,
		&token.CreatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get token: %w", err)
	}

	return &token, nil
}

// DeleteToken アクセストークンを削除する
func (r *TokenRepository) DeleteToken(accessToken string) error {
	query := `DELETE FROM oauth_tokens WHERE access_token = ?`
	_, err := r.conn.DB.Exec(query, accessToken)
	if err != nil {
		return fmt.Errorf("failed to delete token: %w", err)
	}
	return nil
}
