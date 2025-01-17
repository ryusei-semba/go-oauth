package usecase

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/marcboeker/go-duckdb"
	"golang.org/x/exp/rand"
)

type TokenResponse struct {
	Token string `json:"token"`
}

func RequestToken() (TokenResponse, error) {
	// TODO Openをinit関数に切り出す
	db, err := sql.Open("duckdb", "my_database.duckdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// テーブルが存在しない場合は作成する
	_, err = db.Exec(`CREATE TABLE people (id INTEGER, name VARCHAR)`)
	if err != nil {
		log.Fatal(err)
	}

	// ランダムID,Nameを生成してデータをdbにInsertする
	for i := 0; i < 10; i++ {
		id := rand.Intn(100)
		name := fmt.Sprintf("name%d", id)
		_, err = db.Exec(`INSERT INTO people VALUES (?, ?)`, id, name)
		if err != nil {
			log.Fatal(err)
		}
	}

	// _, err = db.Exec(`INSERT INTO people VALUES (42, 'John')`)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	var (
		id   int
		name string
	)
	// 全件取得
	rows, err := db.Query(`SELECT id, name FROM people`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("id: %d, name: %s\n", id, name)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	// テーブルが存在する場合はテーブルを削除する
	_, err = db.Exec(`DROP TABLE people`)
	if err != nil {
		log.Fatal(err)
	}

	return TokenResponse{
		"tokenSample",
	}, nil
}
