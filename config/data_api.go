package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

func ConfigDataApi() {
	dbFileName := "DB/database.db"

	// Open a connection to the SQLite database
	db, err := sql.Open("sqlite3", dbFileName)
	if err == nil {
		db.Close()
		return
	}
	
	// Create the SQLite database file
	file, err := os.Create(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	file.Close() // Close the file after creation

	db, err = sql.Open("sqlite3", dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()


	// Create Client table with UUID field
	clientTable := `
	CREATE TABLE IF NOT EXISTS Client (
		uuid TEXT NOT NULL UNIQUE PRIMARY KEY,
		numero_client INTEGER,
		nom_client TEXT NOT NULL
	);`
	if _, err := db.Exec(clientTable); err != nil {
		log.Fatalf("Error creating Client table: %v", err)
	}

	// Create Individu table with UUID field
	individuTable := `
	CREATE TABLE IF NOT EXISTS Individu (
		uuid TEXT NOT NULL UNIQUE PRIMARY KEY,
		id_individu INTEGER PRIMARY KEY,
		Nom TEXT NOT NULL,
		Prenom TEXT NOT NULL,
		date_naissance TEXT NOT NULL,
		date_fin_validite_CNI TEXT NOT NULL,
		numero_CNI TEXT NOT NULL UNIQUE
	);`
	if _, err := db.Exec(individuTable); err != nil {
		log.Fatalf("Error creating Individu table: %v", err)
	}

	// Create Logging table
	loggingTable := `
	CREATE TABLE IF NOT EXISTS Logging (
		uuid TEXT PRIMARY KEY,
		username TEXT NOT NULL UNIQUE,
		hashed_password TEXT NOT NULL
	);`
	if _, err := db.Exec(loggingTable); err != nil {
		log.Fatalf("Error creating Logging table: %v", err)
	}

	fmt.Println("Database and tables created successfully!")
}
