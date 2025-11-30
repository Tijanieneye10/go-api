package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	Sqlite *sql.DB
}

func NewDB() *DB {
	db, err := sql.Open("sqlite3", "./database.sqlite")

	if err != nil {
		panic(err)
	}

	sqlStmt := `
    CREATE TABLE IF NOT EXISTS todos (
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        task TEXT,
        status BOOLEAN
    );
    `
	_, err = db.Exec(sqlStmt)

	if err != nil {
		panic(err)
	}

	return &DB{db}
}
