package get

import (
	"api_test/data"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllIndividusByClient retrieves all individus associated with a specific client.
func GetAllIndividusByClient(c *gin.Context, dataApiContainer *data.ApiContainer, allClients []data.Client, allIndividus []data.Individu) {
	UUIDClient := c.Param("number_clients")

	clientNumber, err := GetClientNumberByUUID(allClients, UUIDClient)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Filter the individus based on the client number.
	var clientIndividus []data.Individu
	for _, individu := range allIndividus {
		// Assuming there is a field in Individu that links it to a client, e.g., ClientNumber
		if individu.IndividualID == clientNumber {
			clientIndividus = append(clientIndividus, individu)
		}
	}

	// Respond with the list of individus for the specified client.
	if len(clientIndividus) > 0 {
		c.IndentedJSON(http.StatusOK, clientIndividus)
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "no individus found for the specified client"})
	}
}

// GetClientNumberByUUID retrieves the client number associated with a given UUID.
func GetClientNumberByUUID(clients []data.Client, inputUUID string) (int, error) {
	for _, client := range clients {
		if client.UUID == inputUUID {
			return client.ClientNumber, nil
		}
	}
	return 0, fmt.Errorf("client not found for UUID: %s", inputUUID)
}
