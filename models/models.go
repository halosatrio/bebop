package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

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

type Tracker struct {
	Date      time.Time `json:"date"`
	HabitID   string    `json:"habit_id"`
	Count     int       `json:"count"`
	UpdatedAt time.Time `json:"updated_at"`
}
