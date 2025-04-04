package post

import (
	"api_test/DB"
	"api_test/data"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context, dataApiContainer *data.ApiContainer, dbPath string) {

	var InputLogin data.InputLogin

	// Bind the JSON input
	if err := c.ShouldBindJSON(&InputLogin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Retrieve the client from the DB (this should be adapted à ton schéma de BDD)
	user, err := DB.GetClientByUsername(dbPath, InputLogin.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Client not found"})
		return
	}

	// Compare the provided password with the hashed password in the DB
	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(InputLogin.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	// If password matches
	c.JSON(http.StatusOK, gin.H{"message": "Password is correct"})
}
