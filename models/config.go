package models

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"

	databaseConfig "app_translator/config"
)

type Config struct {
	ID             int
	Language       string
	BaseFile       string
	TranslatedFile string
	GroqToken      string
}

type ConfigRequest struct {
	Language  string `json:"language"`
	GroqToken string `json:"groqToken"`
}

type Languages struct {
	Code     string `json:"code"`
	Language string `json:"language"`
}

func GetConfig(db *sql.DB) Config {

	var config Config

	err := db.QueryRow("SELECT * FROM config LIMIT 1").
		Scan(&config.ID, &config.Language, &config.BaseFile, &config.TranslatedFile, &config.GroqToken)

	if err != nil {
		log.Fatal(err)
	}

	return config

	// return Config{
	// 	ID:             1,
	// 	Language:       "pt_br",
	// 	InputFile:      "files/input/base.file",
	// 	BaseFile:       "files/input/base.file",
	// 	TranslatedFile: "files/input/translated.file",
	// 	GroqToken:      "token here",
	// }
}

func SetConfig(config ConfigRequest) {

	DB := databaseConfig.DatabaseConnect()

	var err error
	// DB, err = sql.Open("sqlite3", "./config/app.db") // Open a connection to the SQLite database file named app.db
	// if err != nil {
	// 	log.Fatal(err) // Log an error and stop the program if the database can't be opened
	// }

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

	sqlStmt := `UPDATE config set language="` + config.Language + `",groqToken="` + config.GroqToken + `" WHERE id = 1`

	// sqlStmt := `INSERT INTO config ( language, inputFile, translatedFile, groqToken)
	// values ("` + config.Language + `", "files/input/base.file", "files/input/translated.file", "` + config.GroqToken + `");`

	_, err = DB.Exec(sqlStmt)
	if err != nil {
		log.Fatalf("Error insert into table: %q: %s\n", err, sqlStmt)
	}

}

func SetBaseFileConfig(fileName string) {
	DB := databaseConfig.DatabaseConnect()

	var err error

	sqlStmt := `UPDATE config set inputFile="` + fileName + `" WHERE id = 1`

	_, err = DB.Exec(sqlStmt)
	if err != nil {
		log.Fatalf("Error insert into table: %q: %s\n", err, sqlStmt)
	}

}
