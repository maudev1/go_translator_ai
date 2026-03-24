package models

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

type Config struct {
	ID             int
	Language       string
	InputFile      string
	BaseFile       string
	TranslatedFile string
	GroqToken      string
}

func GetConfig() Config {

	SetConfig()

	return Config{
		ID:             1,
		Language:       "pt_br",
		InputFile:      "files/input/base.file",
		BaseFile:       "files/input/base.file",
		TranslatedFile: "files/input/translated.file",
		GroqToken:      "token here",
	}
}

func SetConfig() {

	var err error
	DB, err = sql.Open("sqlite3", "./config/app.db") // Open a connection to the SQLite database file named app.db
	if err != nil {
		log.Fatal(err) // Log an error and stop the program if the database can't be opened
	}

	// SQL statement to create the todos table if it doesn't exist
	// sqlStmt := `CREATE TABLE IF NOT EXISTS config (
	// id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	// language TEXT,
	// inputFile TEXT,
	// TranslatedFile TEXT
	// );`

	// _, err = DB.Exec(sqlStmt)
	// if err != nil {
	// 	log.Fatalf("Error creating table: %q: %s\n", err, sqlStmt)
	// }

	sqlStmt := `INSERT INTO config ( language, inputFile, translatedFile) 
	values ("pt_br", "files/input/base.file", "files/input/translated.file");`

	_, err = DB.Exec(sqlStmt)
	if err != nil {
		log.Fatalf("Error insert into table: %q: %s\n", err, sqlStmt)
	}

}
