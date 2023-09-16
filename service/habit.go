package service

import (
	"github.com/google/uuid"
	"github.com/halosatrio/bebop/models"
	"github.com/halosatrio/bebop/repository"
)

type HabitService struct {
	repo *repository.HabitRepository
}

func NewHabitService(r *repository.HabitRepository) *HabitService {
	return &HabitService{repo: r}
}

func (s *HabitService) CreateHabit(h *models.Habit) error {
	return s.repo.CreateHabit(h)
}

func (s *HabitService) GetHabitsByUserID(userID uuid.UUID) ([]models.Habit, error) {
	return s.repo.FindByUserID(userID)
}
