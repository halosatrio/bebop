package service

import (
	"errors"

	"github.com/halosatrio/bebop/models"
	"github.com/halosatrio/bebop/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(r *repository.UserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) Register(registerReq models.AuthRequset) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerReq.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	registerReq.Password = string(hashedPassword)
	return s.repo.Store(registerReq)
}

func (s *UserService) Authenticate(email, password string) (*models.User, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}
