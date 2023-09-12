package service

import (
	"errors"
	"fmt"

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

func (s *UserService) Register(user models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	return s.repo.Store(user)
}

func (s *UserService) Authenticate(email, password string) (bool, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return false, errors.New("invalid credentials")
	}

	return true, nil
}