package data

// ApiConfig holds the configuration for the API server, including host and port.
type ApiConfig struct {
	Host string // The hostname where the application will run.
	Port string // The port number where the application will listen for requests.
}

// Album represents a music album with its details.
type Album struct {
	ID     string  `json:"id"`     // Unique identifier for the album.
	Title  string  `json:"title"`  // Title of the album.
	Artist string  `json:"artist"` // Artist of the album.
	Price  float64 `json:"price"`  // Price of the album.
}

// DataApi contains a collection of albums.
type DataApi struct {
	Albums []Album // A slice of Album structs representing the available albums.
}

// ApiContainer aggregates the API configuration and data API.
type ApiContainer struct {
	ApiConfig *ApiConfig // Pointer to the ApiConfig struct for server configuration.
	DataApi   *DataApi   // Pointer to the DataApi struct containing album data.
}
