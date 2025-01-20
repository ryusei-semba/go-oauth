package database

import (
	"database/sql"
	"sync"

	_ "github.com/marcboeker/go-duckdb"
)

var (
	instance *sql.DB
	once     sync.Once
)

// DBConnection データベース接続を管理する構造体
type DBConnection struct {
	DB *sql.DB
}

// NewDBConnection DBConnectionのインスタンスを生成する
func NewDBConnection() (*DBConnection, error) {
	var err error
	once.Do(func() {
		instance, err = sql.Open("duckdb", "my_database.duckdb")
	})

	if err != nil {
		return nil, err
	}

	return &DBConnection{
		DB: instance,
	}, nil
}

// Close データベース接続を閉じる
func (c *DBConnection) Close() error {
	if c.DB != nil {
		return c.DB.Close()
	}
	return nil
}
