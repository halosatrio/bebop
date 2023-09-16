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

	private := r.Group("/")
	private.Use(middleware.JWTAuth())
	private.POST("/create-habit", habitHandler.CreateHabit)
	private.GET("/habits/:user_id", habitHandler.GetHabits)

	// r.POST("/create-habit", middleware.JWTAuth(), habitHandler.CreateHabit)
	// r.GET("/habits/:user_id", middleware.JWTAuth(), habitHandler.GetHabits)

	return r
}
