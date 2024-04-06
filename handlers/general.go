package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/halosatrio/bebop/utils"
)

func Welcome(c *gin.Context) {
	utils.SuccessResponseWithMessage(c, http.StatusOK, "Welcome to Bebop!")
}
