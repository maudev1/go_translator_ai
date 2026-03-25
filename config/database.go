package config

import (
	"database/sql"
	"log"
)

func DatabaseConnect() *sql.DB {
	db, err := sql.Open("sqlite3", "config/app.db")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
