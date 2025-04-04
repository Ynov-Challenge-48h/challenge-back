package routes

import (
	"api_test/DB"
	"api_test/data"
	"api_test/pkg/get"
	"log"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// Setup defines all the routes for the application.
func Setup(router *gin.Engine, dataApiContainer *data.ApiContainer) {

	dbPath := filepath.Join(dataApiContainer.DataApi.Folder, dataApiContainer.DataApi.NameDatabase)

	allClients, err := DB.GetClients(dbPath)
	if err != nil {
		log.Fatal(err)
	}

	allIndividus, err := DB.GetIndividus(dbPath)
	if err != nil {
		log.Fatal(err)
	}

	// Define a GET route for retrieving all clients.
	router.GET("/clients", func(c *gin.Context) {
		get.GetClients(c, dataApiContainer, allClients)
	})

	// Define a GET route for retrieving all individus associated with a specific client.
	router.GET("/clients/:number_clients/individus", func(c *gin.Context) {
		get.GetAllIndividusByClient(c, dataApiContainer, allClients, allIndividus)
	})
}
