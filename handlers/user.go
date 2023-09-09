// handlers/register.go
package handlers

import (
	"database/sql"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
	dbs "github.com/halosatrio/bebop/db"
	"github.com/halosatrio/bebop/models"
)

type RegistrationInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

var db *sql.DB

func Register(c *gin.Context) {
	var input RegistrationInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to hash password"})
		return
	}

	newUser := models.User{
		Email:    input.Email,
		Password: string(hashedPassword),
	}

	if err := dbs.StoreUser(db, newUser); err != nil {
		c.JSON(500, gin.H{"error": "Failed to store user"})
		return
	}

	c.JSON(201, gin.H{"message": "Registration successful"})
}
