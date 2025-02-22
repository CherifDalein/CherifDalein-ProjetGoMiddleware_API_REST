package helpers

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./database.db") // Remplace par le chemin de ta base de données

	if err != nil {
		log.Fatal("Erreur lors de l'ouverture de la base de données:", err)
	}

	// Création de la table events si elle n'existe pas
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS events (
		id TEXT PRIMARY KEY,
		summary TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		start DATETIME NOT NULL,
		end DATETIME NOT NULL,
		uid TEXT NOT NULL,
		resource_id TEXT NOT NULL
	);`

	_, err = DB.Exec(createTableQuery)
	if err != nil {
		log.Fatal("Erreur lors de la création de la table events:", err)
	}

	fmt.Println("Base de données initialisée et table 'events' créée si elle n'existait pas.")
}
