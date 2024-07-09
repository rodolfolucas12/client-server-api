package repository

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "cotacao.db")
	if err != nil {
		log.Fatal("Falha ao conectar ao banco de dados: ", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Falha ao conectar ao banco de dados: ", err)
		return nil, err
	}
	log.Println("Conectado ao banco de dados")
	createTable(db)

	return db, nil
}

func createTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS dollar_price (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		price TEXT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);`

	if _, err := db.Exec(query); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func InsertDollarPrice(price string) error {
	db, err := ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	query := `INSERT INTO dollar_price (price, created_at) VALUES (?, datetime('now'))`
	if _, err := db.ExecContext(ctx, query, price); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
