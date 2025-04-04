package post

import (
	"api_test/data"
	"fmt"
	"net/http"
	"api_test/DB"

	"github.com/gin-gonic/gin"
)

// get the pictures and individu uuid and update information about the individus with the pictures
func AddCNIdata(c *gin.Context, dbpath string, allIndividus []data.Individu) {
	UUIDindividu := c.Param("individu_uuid")

	// Find the good individu whith the corect uuid.
	var IndividuToUpdate data.Individu
	for _, individu := range allIndividus {
		// Assuming there is a field in Individu that links it to a client, e.g., ClientNumber
		if individu.UUID == UUIDindividu {
			 IndividuToUpdate = individu
			break
		}
	}
	fmt.Print(IndividuToUpdate)

	DB.SetIndividus(dbpath,IndividuToUpdate)

	c.IndentedJSON(http.StatusOK, true)
}