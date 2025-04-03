package data

// ApiConfig holds the configuration for the API server, including host and port.
type ApiConfig struct {
	Host string // The hostname where the application will run.
	Port string // The port number where the application will listen for requests.
}

// ApiContainer aggregates the API configuration and data API.
type ApiContainer struct {
	ApiConfig *ApiConfig // Pointer to the ApiConfig struct for server configuration.
}
