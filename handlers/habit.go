package handlers

import (
	"time"

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
	// Decode JWT and get user ID and email
	userID, _ := c.MustGet("user_id").(string)

	var habit models.Habit
	if err := c.ShouldBindJSON(&habit); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	habit.UserID = uuid.MustParse(userID) // Set the user ID from the JWT
	habit.IsActive = true                 // Assuming the habit should be active when created
	habit.StartDate = time.Now()          // Assuming start date is the creation date

	err := h.service.CreateHabit(&habit)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create habit.", "message": err})
		return
	}

	c.JSON(200, gin.H{"message": "Habit successfully created!"})
}

func (h *HabitHandler) GetHabits(c *gin.Context) {
	userID, _ := c.MustGet("user_id").(string)

	habits, err := h.service.GetHabitsByUserID(uuid.MustParse(userID))
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch habits."})
		return
	}

	c.JSON(200, habits)
}
