package data

// ApiConfig holds the configuration for the API server, including host and port.
type ApiConfig struct {
	Host string // The hostname where the application will run.
	Port string // The port number where the application will listen for requests.
}

// DataAPI contains the methods for interacting with data.
type DataApi struct {
	Folder       string // The path to the data folder.
	NameDatabase string // The name of the data database.
}

// Client represents a client in the system.
type Client struct {
	UUID         string `json:"uuid"`
	ClientNumber int    `json:"client_number"`
	ClientName   string `json:"client_name"`
}

// Individu represents an individual in the system.
type Individu struct {
	UUID         string `json:"uuid"`
	IndividualID int    `json:"individu_number"`
	LastName     string `json:"individu_name"`
	FirstName    string `json:"first_name"`
}

// ApiContainer aggregates the API configuration and data API.
type ApiContainer struct {
	ApiConfig *ApiConfig // Pointer to the ApiConfig struct for server configuration.
	DataApi   *DataApi   // Pointer to the DataApi struct for data interaction.
}
