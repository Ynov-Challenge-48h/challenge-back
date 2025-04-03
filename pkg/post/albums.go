package post

import (
	"api_test/data"
	"net/http"

	"github.com/gin-gonic/gin"
)

// postAlbums adds an album from JSON received in the request body.
func PostAlbums(c *gin.Context, dataApiContainer *data.ApiContainer) {
	var newAlbum data.Album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	dataApiContainer.DataApi.Albums = append(dataApiContainer.DataApi.Albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
