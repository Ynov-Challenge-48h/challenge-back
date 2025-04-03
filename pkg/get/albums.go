package get

import (
	"api_test/data"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func GetAlbumByID(c *gin.Context, dataApiContainer *data.ApiContainer) {
	id := c.Param("id") // Retrieve the album ID from the request parameters.

	// Loop through the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range dataApiContainer.DataApi.Albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a) // Respond with the album in JSON format if found.
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"}) // Respond with a 404 error if not found.
}

// GetAlbums responds with the list of all albums as JSON.
func GetAlbums(c *gin.Context, dataApiContainer *data.ApiContainer) {
	c.IndentedJSON(http.StatusOK, dataApiContainer.DataApi.Albums) // Respond with all albums in JSON format.
}
