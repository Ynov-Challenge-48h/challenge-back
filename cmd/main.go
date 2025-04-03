package main

import (
	"api_test/config"
	"api_test/data"
	"api_test/pkg/get"
	"api_test/pkg/post"

	"github.com/gin-gonic/gin"
)

// The main function serves as the entry point for the application.
func main() {
	// Initialize the data API container with configuration and data API instances.
	dataApiContainer := &data.ApiContainer{
		ApiConfig: config.ConfigApi(),     // Load API configuration settings.
		DataApi:   config.ConfigDataApi(), // Load data API settings.
	}

	// Create a new Gin router instance.
	router := gin.Default()

	// Define a GET route for retrieving all albums.
	router.GET("/albums", func(c *gin.Context) {
		get.GetAlbums(c, dataApiContainer) // Call the GetAlbums function with the context and data API container.
	})

	// Define a GET route for retrieving a specific album by ID.
	router.GET("/albums/:id", func(c *gin.Context) {
		get.GetAlbumByID(c, dataApiContainer) // Call the GetAlbumByID function with the context and data API container.
	})

	// Define a POST route for creating a new album.
	router.POST("/albums", func(c *gin.Context) {
		post.PostAlbums(c, dataApiContainer) // Call the PostAlbums function with the context and data API container.
	})

	// Start the Gin server on the specified host and port from the API configuration.
	router.Run(dataApiContainer.ApiConfig.Host + ":" + dataApiContainer.ApiConfig.Port)
}
