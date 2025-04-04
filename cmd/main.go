package main

import (
	"api_test/DB"
	"api_test/config"
	"api_test/data"
	"api_test/pkg/routes"

	"github.com/gin-gonic/gin"
)

// The main function serves as the entry point for the application.
func main() {
	// Initialize the data API container with configuration and data API instances.
	dataApiContainer := &data.ApiContainer{
		ApiConfig: config.ConfigApi(),     // Load API configuration settings.
		DataApi:   config.ConfigDataApi(), // Load data API settings.
	}

	// Initialize the database connection with the provided data API container.
	DB.ManageDB(dataApiContainer)

	// Create a new Gin router instance.
	router := gin.Default()

	// Set up the routes using the routes package.
	routes.Setup(router, dataApiContainer)

	// Start the Gin server on the specified host and port from the API configuration.
	router.Run(dataApiContainer.ApiConfig.Host + ":" + dataApiContainer.ApiConfig.Port)
}
