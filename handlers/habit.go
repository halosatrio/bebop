package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/halosatrio/bebop/models"
	"github.com/halosatrio/bebop/service"
)

type HabitHandler struct {
	service *service.HabitService
}

func NewHabitHandler(s *service.HabitService) *HabitHandler {
	return &HabitHandler{service: s}
}

func (h *HabitHandler) CreateHabit(c *gin.Context) {
	var habit models.Habit

	if err := c.ShouldBindJSON(&habit); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Set the user ID from JWT middleware
	habit.UserID = c.MustGet("user_id").(uuid.UUID)

	// fmt.Println("user_id", c.MustGet("user_id"))
	// fmt.Println("email", c.MustGet("email"))

	err := h.service.CreateHabit(&habit)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create habit"})
		return
	}

	c.JSON(201, habit)
}
