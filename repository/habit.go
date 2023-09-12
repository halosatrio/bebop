package repository

import (
	"database/sql"

	"github.com/halosatrio/bebop/models"
)

type HabitRepository struct {
	DB *sql.DB
}

func NewHabitRepository(DB *sql.DB) *HabitRepository {
	return &HabitRepository{DB: DB}
}

func (r *HabitRepository) CreateHabit(habit *models.Habit) error {
	query := `
		INSERT INTO bebop.habit (user_id, title, icon, color, is_active, start_date, daily_goal, weekly_goal)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, created_at, updated_at
	`

	return r.DB.QueryRow(query, habit.UserID, habit.Title, habit.Icon, habit.Color, habit.IsActive, habit.StartDate, habit.DailyGoal, habit.WeeklyGoal).Scan(&habit.ID, &habit.CreatedAt, &habit.UpdatedAt)
}
