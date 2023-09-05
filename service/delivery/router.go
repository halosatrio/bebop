package delivery

import (
	"github.com/gin-gonic/gin"
)

// InitRouter initializes the Gin router
func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/v1/test", TestHandler)

	return router
}

// TestHandler handles the /v1/test route
func TestHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome to v1 test!",
	})
}
