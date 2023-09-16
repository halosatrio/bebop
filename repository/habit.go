package repository

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
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
		INSERT INTO bebop.habits (user_id, title, icon, color, is_active, start_date, daily_goal, weekly_goal, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`

	_, err := r.DB.Exec(query, habit.UserID, habit.Title, habit.Icon, habit.Color, habit.IsActive, habit.StartDate, habit.DailyGoal, habit.WeeklyGoal, time.Now(), time.Now())
	return err
}

func (r *HabitRepository) FindByUserID(userID uuid.UUID) ([]models.Habit, error) {
	habits := []models.Habit{}
	query := `
		SELECT id, user_id, title, icon, color, is_active, start_date, daily_goal, weekly_goal, created_at, updated_at
		FROM bebop.habits WHERE user_id = $1
	`

	rows, err := r.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var habit models.Habit
		err = rows.Scan(&habit.ID, &habit.UserID, &habit.Title, &habit.Icon, &habit.Color, &habit.IsActive,
			&habit.StartDate, &habit.DailyGoal, &habit.WeeklyGoal, &habit.CreatedAt, &habit.UpdatedAt)
		if err != nil {
			return nil, err
		}
		habits = append(habits, habit)
	}
	return habits, nil
}
