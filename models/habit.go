package models

import (
	"time"

	"github.com/google/uuid"
)

type Habit struct {
	ID         uuid.UUID `json:"id"`
	UserID     uuid.UUID `json:"user_id"`
	Title      string    `json:"title"`
	Icon       string    `json:"icon"`
	Color      string    `json:"color"`
	IsActive   bool      `json:"is_active"`
	StartDate  time.Time `json:"start_date"`
	DailyGoal  int       `json:"daily_goal"`
	WeeklyGoal int       `json:"weekly_goal"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type CreateHabitRequest struct {
	Title      string `json:"title" binding:"required"`
	Icon       string `json:"icon" binding:"required"`
	Color      string `json:"color" binding:"required"`
	StartDate  string `json:"start_date" binding:"required"`
	DailyGoal  int    `json:"daily_goal" binding:"required"`
	WeeklyGoal int    `json:"weekly_goal" binding:"required"`
}
