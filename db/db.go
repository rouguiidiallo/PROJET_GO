package db

import (
	"database/sql"
	"log"
	_ "github.com/glebarez/go-sqlite" // Import pour le driver SQLite
)

var DB *sql.DB

// Initialise la base de données SQLite
func InitDB() {
	var err error
	DB, err = sql.Open("sqlite", "./restaurant_orders.db")
	if err != nil {
		log.Fatal("Erreur lors de l'ouverture de la base de données:", err)
	}

	// Créer la table si elle n'existe pas
	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS orders (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			customerName TEXT,
			item TEXT,
			quantity INTEGER,
			status TEXT
		);
	`)
	if err != nil {
		log.Fatal("Erreur lors de la création de la table:", err)
	}
}

// Fermer la base de données
func CloseDB() {
	if err := DB.Close(); err != nil {
		log.Fatal("Erreur lors de la fermeture de la base de données:", err)
	}
}
