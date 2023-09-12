package router

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/halosatrio/bebop/handlers"
	"github.com/halosatrio/bebop/middleware"
	"github.com/halosatrio/bebop/repository"
	"github.com/halosatrio/bebop/service"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	r := gin.Default()

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	habitRepo := repository.NewHabitRepository(db)
	habitService := service.NewHabitService(habitRepo)
	habitHandler := handlers.NewHabitHandler(habitService)

	r.GET("/test", handlers.Welcome)
	r.POST("/register", userHandler.Register)
	r.POST("/auth", userHandler.Authenticate)

	authenticated := r.Group("/")
	authenticated.Use(middleware.JWTAuth())
	authenticated.POST("/create-habit", habitHandler.CreateHabit)

	return r
}
