package database

import "fmt"

// DBInitializer データベースの初期化を担当する構造体
type DBInitializer struct {
	conn *DBConnection
}

// NewDBInitializer DBInitializerのインスタンスを生成する
func NewDBInitializer(conn *DBConnection) *DBInitializer {
	return &DBInitializer{
		conn: conn,
	}
}

// InitializeDB データベースを初期化する
func (i *DBInitializer) InitializeDB() error {
	if err := i.createTokenTable(); err != nil {
		return fmt.Errorf("failed to initialize database: %w", err)
	}
	return nil
}

// createTokenTable トークンテーブルを作成する
func (i *DBInitializer) createTokenTable() error {
	query := `
		CREATE TABLE IF NOT EXISTS oauth_tokens (
			access_token VARCHAR PRIMARY KEY,
			refresh_token VARCHAR UNIQUE,
			client_id VARCHAR,
			user_id VARCHAR,
			expires_at TIMESTAMP,
			scope VARCHAR,
			created_at TIMESTAMP
		)
	`
	_, err := i.conn.DB.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create oauth_tokens table: %w", err)
	}
	return nil
}
