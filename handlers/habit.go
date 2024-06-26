package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/halosatrio/bebop/models"
	"github.com/halosatrio/bebop/repository"
	"github.com/halosatrio/bebop/utils"
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

	if serr := c.ShouldBindJSON(&req); serr != nil {
		utils.ErrorResponseDataNull(c, http.StatusBadRequest, "[handler][habit] "+serr.Error())
		return
	}

	parsedDate, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		utils.ErrorResponseDataNull(c, http.StatusBadRequest, "[handler][habit] Invalid start_date format. It should be 'YYYY-MM-DD'.")
		return
	}

	habit.UserID = uuid.MustParse(userID)
	habit.IsActive = true
	habit.Color = req.Color
	habit.StartDate = parsedDate
	habit.DailyGoal = req.DailyGoal
	habit.Icon = req.Icon
	habit.Title = req.Title
	habit.WeeklyGoal = req.WeeklyGoal

	errx := h.createHabitRepo(&habit)
	if errx != nil {
		utils.ErrorResponseDataNull(c, http.StatusInternalServerError, "[handler][habit] Failed to create habit. "+errx.Error())
		return
	}

	utils.SuccessResponseWithMessage(c, http.StatusOK, "Habit is successfully created!")
}

func (r *HabitHandler) getHabitsByUserID(userID uuid.UUID) ([]models.Habit, error) {
	return r.repo.FindByUserID(userID)
}

func (h *HabitHandler) GetHabits(c *gin.Context) {
	userID, _ := c.MustGet("user_id").(string)

	habits, err := h.getHabitsByUserID(uuid.MustParse(userID))
	if err != nil {
		utils.ErrorResponseDataNull(c, http.StatusInternalServerError, "[handler][habit] Failed to fetch habits.")
		return
	}

	utils.SuccessResponseWithData(c, http.StatusOK, habits)
}

func (r *HabitHandler) getHabitByID(habitID uuid.UUID) (*models.Habit, error) {
	return r.repo.FindHabitByID(habitID)
}

func (h *HabitHandler) GetHabit(c *gin.Context) {
	habitID := c.Param("id")

	habit, err := h.getHabitByID(uuid.MustParse(habitID))
	if err != nil {
		utils.ErrorResponseDataNull(c, http.StatusInternalServerError, "Failed to fetch the habit.")
		return
	}

	utils.SuccessResponseWithData(c, http.StatusOK, habit)
}
