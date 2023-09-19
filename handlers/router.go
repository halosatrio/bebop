package handlers

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/halosatrio/bebop/repository"
	"github.com/halosatrio/bebop/utils"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	r := gin.Default()

	userRepo := repository.NewUserRepository(db)
	userHandler := NewUserHandler(userRepo)

	habitRepo := repository.NewHabitRepository(db)
	habitHandler := NewHabitHandler(habitRepo)

	r.GET("/test", Welcome)
	r.POST("/register", userHandler.Register)
	r.POST("/auth", userHandler.Authenticate)

	private := r.Group("/")
	private.Use(utils.JWTAuth())
	private.POST("/create-habit", habitHandler.CreateHabit)
	private.GET("/habits", habitHandler.GetHabits)
	private.GET("/habit/:id", habitHandler.GetHabit)

	// r.POST("/create-habit", middleware.JWTAuth(), habitHandler.CreateHabit)
	// r.GET("/habits/:user_id", middleware.JWTAuth(), habitHandler.GetHabits)

	return r
}
