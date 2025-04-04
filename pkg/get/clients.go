package get

import (
	"api_test/data"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetClients responds with the list of all clients as JSON.
func GetClients(c *gin.Context, dataApiContainer *data.ApiContainer, allClients []data.Client) {
	c.IndentedJSON(http.StatusOK, allClients) // Respond with all clients in JSON format.
}
