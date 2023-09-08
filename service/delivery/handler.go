package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (bx bebopHandler) Welcome(ctx *gin.Context) {
	ctx.Status(http.StatusOK)
	ctx.JSON(200, gin.H{
		"message": "Welcome to v1 test!",
	})
}
