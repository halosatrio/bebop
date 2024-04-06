package utils

import (
	"github.com/gin-gonic/gin"
)

func SuccessResponseWithData(c *gin.Context, status int, data interface{}) {
	c.JSON(status, gin.H{
		"status":  status,
		"message": "Success",
		"data":    data,
	})
}

func SuccessResponseWithMessage(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{
		"status":  status,
		"message": message,
	})
}

func ErrorResponseDataNull(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{
		"status":  status,
		"message": "Failed",
		"data":    nil,
		"error":   message,
	})
}
