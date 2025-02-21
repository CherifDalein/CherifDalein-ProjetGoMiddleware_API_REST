package helpers

import (
	"database/sql"
	"log"
)

import (
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "config.db")
	if err != nil {
		log.Fatal(err)
	}

	// Création des tables si elles n'existent pas
	createTableSQL := `
    CREATE TABLE IF NOT EXISTS resources (
        id TEXT PRIMARY KEY,
        name TEXT NOT NULL,
        ucaID TEXT NOT NULL
    );
    CREATE TABLE IF NOT EXISTS alerts (
        id TEXT PRIMARY KEY,
        resource_id TEXT,
        email TEXT NOT NULL,
        FOREIGN KEY(resource_id) REFERENCES resources(id)
    );
    `
	_, err = DB.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
}
