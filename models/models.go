package models

import (
	"time"
)

type Tracker struct {
	Date      time.Time `json:"date"`
	HabitID   string    `json:"habit_id"`
	Count     int       `json:"count"`
	UpdatedAt time.Time `json:"updated_at"`
}
