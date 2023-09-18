package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/halosatrio/bebop/models"
	"github.com/halosatrio/bebop/repository"
)

type HabitHandler struct {
	repo *repository.HabitRepository
}

func NewHabitHandler(r *repository.HabitRepository) *HabitHandler {
	return &HabitHandler{repo: r}
}

func (r *HabitHandler) createHabitRepo(h *models.Habit) error {
	return r.repo.CreateHabit(h)
}

func (h *HabitHandler) CreateHabit(c *gin.Context) {
	// Decode JWT and get user ID and email
	userID, _ := c.MustGet("user_id").(string)

	var req models.CreateHabitRequest
	var habit models.Habit

	fmt.Print("req start date", req.StartDate)

	if req.StartDate.IsZero() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "[handler][habit] Key: 'StartDate' Error:Field validation for 'StartDate' cannot be null"})
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "[handler][habit] " + err.Error()})
		return
	}

	habit.UserID = uuid.MustParse(userID)
	habit.IsActive = true
	habit.Color = req.Color
	habit.StartDate = req.StartDate
	habit.DailyGoal = req.DailyGoal
	habit.Icon = req.Icon
	habit.Title = req.Title
	habit.WeeklyGoal = req.WeeklyGoal

	err := h.createHabitRepo(&habit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "[handler][habit] Failed to create habit.", "message": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Habit successfully created!"})
}

func (s *HabitHandler) getHabitsByUserID(userID uuid.UUID) ([]models.Habit, error) {
	return s.repo.FindByUserID(userID)
}

func (h *HabitHandler) GetHabits(c *gin.Context) {
	userID, _ := c.MustGet("user_id").(string)

	habits, err := h.getHabitsByUserID(uuid.MustParse(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "[handler][habit] Failed to fetch habits."})
		return
	}

	c.JSON(http.StatusOK, habits)
}
