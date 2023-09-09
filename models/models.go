package models

import (
	"time"
)

type User struct {
	ID        string    `db:"id" 					json:"id"`
	Email     string    `db:"email" 			json:"email"`
	Password  string    `db:"password" 		json:"password"`
	CreatedAt time.Time `db:"created_at" 	json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" 	json:"updated_at"`
}

type Habit struct {
	ID         string    `db:"id" 					json:"id"`
	UserID     string    `db:"user_id" 			json:"user_id"`
	Title      string    `db:"title" 				json:"title"`
	Icon       string    `db:"icon" 				json:"icon"`
	Color      string    `db:"color" 				json:"color"`
	IsActive   bool      `db:"is_active" 		json:"is_active"`
	StartDate  time.Time `db:"start_date" 	json:"start_date"`
	DailyGoal  int       `db:"daily_goal" 	json:"daily_goal"`
	WeeklyGoal int       `db:"weekly_goal" 	json:"weekly_goal"`
	CreatedAt  time.Time `db:"created_at" 	json:"created_at"`
	UpdatedAt  time.Time `db:"updated_at" 	json:"updated_at"`
}

type Tracker struct {
	Date      time.Time `db:"date" 				json:"date"`
	HabitID   string    `db:"habit_id" 		json:"habit_id"`
	Count     int       `db:"count" 			json:"count"`
	UpdatedAt time.Time `db:"updated_at" 	json:"updated_at"`
}
