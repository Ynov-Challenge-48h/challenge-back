package DB

import (
	"api_test/data"
	"api_test/pkg/managefiles"
	"database/sql"
	"fmt"
	"log"
	"path/filepath"

	"github.com/google/uuid"
	_ "modernc.org/sqlite"
)

// ManageDB is responsible for managing the database lifecycle, including checking for the existence of the database file,
// creating it if it doesn't exist, and initializing the database.
func ManageDB(dataApiContainer *data.ApiContainer) {
	fmt.Println("Managing DB...")

	dir := dataApiContainer.DataApi.Folder
	file := dataApiContainer.DataApi.NameDatabase

	// Check if the database file exists; if not, create it.
	if !managefiles.FileExists(dir, file) {
		fmt.Println("The file does not exist, creating...")
		if err := managefiles.CreateFile(dir, file); err != nil {
			log.Fatalf("Error creating the file: %v", err)
		}
		fmt.Println("File created successfully.")
		// Construct the full path to the database file and initialize the database.
		path := filepath.Join(dir, file)
		InitDB(path)
	}
}

// InitDB initializes the database by opening a connection, creating necessary tables, and populating them with initial data.
func InitDB(dbPath string) {
	fmt.Println("Initializing the database...")

	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	createTables(db)
	fmt.Println("Database and tables created successfully!")
	PopulateDatabase(db)
}

// createTables creates the necessary tables in the database if they do not already exist.
func createTables(db *sql.DB) {
	tables := []string{
		`CREATE TABLE IF NOT EXISTS Client (
            uuid TEXT NOT NULL UNIQUE PRIMARY KEY,
            client_number INTEGER,
            client_name VARCHAR NOT NULL
        );`,
		`CREATE TABLE IF NOT EXISTS Individu (
            uuid TEXT NOT NULL UNIQUE PRIMARY KEY,
            individual_id INTEGER,
            last_name VARCHAR NOT NULL,
            first_name VARCHAR NOT NULL,
            birth_date DATE,
            cni_expiry_date DATE,
            cni_number TEXT UNIQUE
        );`,
		`CREATE TABLE IF NOT EXISTS Logging (
            uuid TEXT NOT NULL PRIMARY KEY,
            username VARCHAR NOT NULL UNIQUE,
            password VARCHAR NOT NULL
        );`,
	}

	// Execute each table creation statement.
	for _, table := range tables {
		if _, err := db.Exec(table); err != nil {
			log.Fatalf("Error creating table: %v", err)
		}
	}
}

// PopulateDatabase inserts initial data into the database tables.
func PopulateDatabase(db *sql.DB) {
	// Insert data into the Client table with real UUIDs.
	clientInsert := `INSERT INTO Client (uuid, client_number, client_name) VALUES (?, ?, ?);`
	clients := []struct {
		clientNumber int
		clientName   string
	}{
		{101, "Client A"},
		{102, "Client B"},
		{103, "Client C"},
	}

	// Insert each client into the database.
	for i, client := range clients {
		uuid := uuid.New().String()
		if _, err := db.Exec(clientInsert, uuid, client.clientNumber, client.clientName); err != nil {
			log.Fatalf("Error inserting data into Client table: %v", err)
		}
		log.Printf("Inserted Client %d with UUID: %s\n", i+1, uuid)
	}

	// Insert data into the Individu table with real UUIDs.
	individuInsert := `INSERT INTO Individu (uuid, individual_id, last_name, first_name) VALUES (?, ?, ?, ?);`
	individus := []struct {
		individualId int
		lastName     string
		firstName    string
	}{
		{102, "John", "Doe"},
		{102, "Jane", "Smith"},
		{103, "Jane", "Smith"},
		{104, "Alice", "Johnson"},
	}

	// Insert each individual into the database.
	for _, individu := range individus {
		uuid := uuid.New().String()
		if _, err := db.Exec(individuInsert, uuid, individu.individualId, individu.lastName, individu.firstName); err != nil {
			log.Fatalf("Error inserting data into Individu table: %v", err)
		}
		log.Printf("Inserted Individu with UUID: %s\n", uuid)
	}

	// Insert data into the Logging table with real UUIDs.
	loggingInsert := `INSERT INTO Logging (uuid, username, password) VALUES (?, ?, ?);`
	loggings := []struct {
		username       string
		hashedPassword string
	}{
		{"user1", "hashedpassword1"}, // fill with the real hash
	}

	// Insert each logging entry into the database.
	for _, logging := range loggings {
		uuid := uuid.New().String()
		if _, err := db.Exec(loggingInsert, uuid, logging.username, logging.hashedPassword); err != nil {
			log.Fatalf("Error inserting data into Logging table: %v", err)
		}
		log.Printf("Inserted Logging with UUID: %s\n", uuid)
	}

	log.Println("Pre-generated content inserted successfully!")
}

// GetClients retrieves all clients from the database.
func GetClients(dbPath string) ([]data.Client, error) {
	// Open a connection to the database.
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	// Define the query to select all clients.
	query := `SELECT uuid, client_number, client_name FROM Client;`

	// Execute the query.
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query clients: %v", err)
	}
	defer rows.Close()

	// Iterate over the result set and populate the clients slice.
	var clients []data.Client
	for rows.Next() {
		var client data.Client
		if err := rows.Scan(&client.UUID, &client.ClientNumber, &client.ClientName); err != nil {
			return nil, fmt.Errorf("failed to scan client: %v", err)
		}
		clients = append(clients, client)
	}

	// Check for errors from iterating over rows.
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over clients: %v", err)
	}

	return clients, nil
}

// GetIndividus retrieves all individus from the database.
func GetIndividus(dbPath string) ([]data.Individu, error) {
	// Open a connection to the database.
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	// Define the query to select all individuals.
	query := `SELECT uuid, individual_id, last_name, first_name FROM Individu;`

	// Execute the query.
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query individuals: %v", err)
	}
	defer rows.Close()

	// Iterate over the result set and populate the individuals slice.
	var individus []data.Individu
	for rows.Next() {
		var individu data.Individu
		if err := rows.Scan(&individu.UUID, &individu.IndividualID, &individu.LastName, &individu.FirstName); err != nil {
			return nil, fmt.Errorf("failed to scan individual: %v", err)
		}
		individus = append(individus, individu)
	}

	// Check for errors from iterating over rows.
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over individuals: %v", err)
	}

	return individus, nil
}
