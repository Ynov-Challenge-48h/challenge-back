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
		client_number INTEGER,
		client_name VARCHAR NOT NULL
	);`
	if _, err := db.Exec(clientTable); err != nil {
		log.Fatalf("Error creating Client table: %v", err)
	}

	// Create Individu table with UUID field
	individuTable := `
	CREATE TABLE IF NOT EXISTS Individu (
		uuid TEXT NOT NULL UNIQUE PRIMARY KEY,
		individual_id INTEGER,
		last_name VARCHAR NOT NULL,
		first_name VARCHAR NOT NULL,
		birth_date DATE,
		cni_expiry_date DATE,
		cni_number TEXT UNIQUE
	);`
	if _, err := db.Exec(individuTable); err != nil {
		log.Fatalf("Error creating Individu table: %v", err)
	}

	// Create Logging table
	loggingTable := `
	CREATE TABLE IF NOT EXISTS Logging (
		uuid TEXT NOT NULL PRIMARY KEY,
		username VARCHAR NOT NULL UNIQUE,
		password VARCHAR NOT NULL
	);`
	if _, err := db.Exec(loggingTable); err != nil {
		log.Fatalf("Error creating Logging table: %v", err)
	}

	fmt.Println("Database and tables created successfully!")
	PopulateDatabase(dbFileName)
}

func PopulateDatabase(dbFileName string) {
	// Open the SQLite database
	db, err := sql.Open("sqlite3", dbFileName)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	// Insert data into Client table with real UUIDs
	clientInsert := `
	INSERT INTO Client (uuid, numero_client, nom_client) VALUES (?, ?, ?);`
	for i, client := range []struct {
		numeroClient int
		nomClient    string
	}{
		{101, "Client A"},
		{102, "Client B"},
		{103, "Client C"},
	} {
		uuid := uuid.New().String() // Generate a new UUID
		if _, err := db.Exec(clientInsert, uuid, client.numeroClient, client.nomClient); err != nil {
			log.Fatalf("Error inserting data into Client table: %v", err)
		}
		log.Printf("Inserted Client %d with UUID: %s\n", i+1, uuid)
	}

	// Insert data into Individu table with real UUIDs
	individuInsert := `
	INSERT INTO Individu (uuid, id_individu, Nom, Prenom) VALUES (?, ?, ?, ?);`
	for _, individu := range []struct {
		idIndividu int
		nom        string
		prenom     string
	}{
		{102, "John", "Doe"},
		{102, "Jane", "Smith"},
		{103, "Jane", "Smith"},
		{104, "Alice", "Johnson"},
	} {
		uuid := uuid.New().String() // Generate a new UUID
		if _, err := db.Exec(individuInsert, uuid, individu.idIndividu, individu.nom, individu.prenom); err != nil {
			log.Fatalf("Error inserting data into Individu table: %v", err)
		}
		log.Printf("Inserted Individu with UUID: %s\n", uuid)
	}

	// Insert data into Logging table with real UUIDs
	loggingInsert := `
	INSERT INTO Logging (uuid, username, hashed_password) VALUES (?, ?, ?);`
	for _, logging := range []struct {
		username       string
		hashedPassword string
	}{
		{"user1", "hashedpassword1"}, // fill with the real hash
	} {
		uuid := uuid.New().String() // Generate a new UUID
		if _, err := db.Exec(loggingInsert, uuid, logging.username, logging.hashedPassword); err != nil {
			log.Fatalf("Error inserting data into Logging table: %v", err)
		}
		log.Printf("Inserted Logging with UUID: %s\n", uuid)
	}

	log.Println("Pre-generated content inserted successfully!")
}
