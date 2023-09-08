package delivery

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/halosatrio/bebop/models"
	"github.com/halosatrio/bebop/service/repository"
	"github.com/halosatrio/bebop/service/utils"
)

// InitRouter initializes the Gin router
func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/v1/test", TestHandler)

	router.POST("/v1/register", RegisterHandler)
	router.POST("/v1/login", LoginHandler)

	return router
}

// TestHandler handles the /v1/test route
func TestHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome to v1 test!",
	})
}

func RegisterHandler(c *gin.Context, db *sql.DB) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}

	user.Password = hashedPassword
	err = repository.CreateUser(db, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
}

func LoginHandler(c *gin.Context, db *sql.DB) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existingUser, err := repository.GetUserByEmail(db, user.Email)
	if err != nil || !utils.CheckPasswordHash(user.Password, existingUser.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	token, err := utils.GenerateJWT(existingUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating JWT"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
