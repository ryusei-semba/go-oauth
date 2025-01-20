package database

import (
	"database/sql"
	"fmt"
	"go-oauth/domain/token"
)

// TokenRepository トークンの永続化を担当するリポジトリの実装
type TokenRepository struct {
	conn *DBConnection
}

// NewTokenRepository TokenRepositoryのインスタンスを生成する
func NewTokenRepository(conn *DBConnection) token.Repository {
	return &TokenRepository{
		conn: conn,
	}
}

// Save トークンを保存する
func (r *TokenRepository) Save(token *token.Token) error {
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
		token.CreatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to save token: %w", err)
	}
	return nil
}

// FindByAccessToken アクセストークンからトークンを取得する
func (r *TokenRepository) FindByAccessToken(accessToken string) (*token.Token, error) {
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

	var t token.Token
	err := r.conn.DB.QueryRow(query, accessToken).Scan(
		&t.AccessToken,
		&t.RefreshToken,
		&t.ClientID,
		&t.UserID,
		&t.ExpiresAt,
		&t.Scope,
		&t.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("token not found")
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get token: %w", err)
	}

	return &t, nil
}

// FindByRefreshToken リフレッシュトークンからトークンを取得する
func (r *TokenRepository) FindByRefreshToken(refreshToken string) (*token.Token, error) {
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
		WHERE refresh_token = ?
	`

	var t token.Token
	err := r.conn.DB.QueryRow(query, refreshToken).Scan(
		&t.AccessToken,
		&t.RefreshToken,
		&t.ClientID,
		&t.UserID,
		&t.ExpiresAt,
		&t.Scope,
		&t.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("token not found")
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get token: %w", err)
	}

	return &t, nil
}

// Delete トークンを削除する
func (r *TokenRepository) Delete(accessToken string) error {
	query := `DELETE FROM oauth_tokens WHERE access_token = ?`
	_, err := r.conn.DB.Exec(query, accessToken)
	if err != nil {
		return fmt.Errorf("failed to delete token: %w", err)
	}
	return nil
}
