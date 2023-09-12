package service

import (
	"github.com/halosatrio/bebop/models"
	"github.com/halosatrio/bebop/repository"
)

type HabitService struct {
	repo *repository.HabitRepository
}

func NewHabitService(r *repository.HabitRepository) *HabitService {
	return &HabitService{repo: r}
}

func (s *HabitService) CreateHabit(habit *models.Habit) error {
	return s.repo.CreateHabit(habit)
}
