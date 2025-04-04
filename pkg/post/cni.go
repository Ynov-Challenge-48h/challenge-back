package post

import (
	"api_test/DB"
	"api_test/data"
	"api_test/pkg/tools"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// WIP / AddCNIdata processes an image associated with an individu and updates the individu's information.
// It retrieves the individu's UUID from the request parameters, analyzes the image to extract data,
// and updates the individu's record in the database.
//
// Parameters:
//   - c: The Gin context, which provides request and response handling.
//   - dbpath: The path to the database file.
//   - allIndividus: A slice of Individu objects representing all individus in the database.
//   - file: The path to the image file to be analyzed.
func AddCNIdata(c *gin.Context, dbpath string, allIndividus []data.Individu, file string) {
	UUIDindividu := c.Param("individu_uuid")

	// Find the correct individu with the specified UUID.
	var IndividuToUpdate data.Individu
	for _, individu := range allIndividus {
		// Assuming there is a field in Individu that links it to a client, e.g., ClientNumber
		if individu.UUID == UUIDindividu {
			IndividuToUpdate = individu
			break
		}
	}

	// Analyze the image to extract data.
	data, err := tools.AnalyzeImage(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to analyze image"})
		return
	}

	fmt.Println(data)

	// Update the individu's information in the database.
	DB.SetIndividus(dbpath, IndividuToUpdate)

	// Respond with a success status.
	c.IndentedJSON(http.StatusOK, true)
}
