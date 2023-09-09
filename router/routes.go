package router

import (
	"github.com/gin-gonic/gin"
	"github.com/halosatrio/bebop/handlers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/test", handlers.Welcome)

	return r
}
