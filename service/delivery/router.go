package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/halosatrio/bebop/service"
)

type bebopHandler struct {
	generalUseCase service.GeneralUseCase
}

func InitRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func CreateHandler(
	router *gin.Engine,
	generalUseCase service.GeneralUseCase,
) {
	handler := &bebopHandler{
		generalUseCase: generalUseCase,
	}

	router.GET("/test", handler.Welcome)
}
