package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/halosatrio/bebop/models"
	"github.com/halosatrio/bebop/repository"
	"github.com/halosatrio/bebop/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	repo *repository.UserRepository
}

func NewUserHandler(r *repository.UserRepository) *UserHandler {
	return &UserHandler{repo: r}
}

func (s *UserHandler) registerHandler(registerReq models.AuthRequset) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerReq.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	registerReq.Password = string(hashedPassword)
	return s.repo.Store(registerReq)
}

func (h *UserHandler) Register(c *gin.Context) {
	var reqRegister models.AuthRequset
	if err := c.ShouldBindJSON(&reqRegister); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "[handler][user] " + err.Error()})
		return
	}

	err := h.registerHandler(reqRegister)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "[handler][user] Failed to register user."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Registration successful."})
}

func (s *UserHandler) authenticateHandler(email, password string) (*models.User, error) {
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

func (h *UserHandler) Authenticate(c *gin.Context) {
	var reqAuth models.AuthRequset
	if err := c.ShouldBindJSON(&reqAuth); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.authenticateHandler(reqAuth.Email, reqAuth.Password)
	if err != nil || user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "[handler][user] Authentication failed."})
		return
	}

	token, err := utils.GenerateJWT(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "[handler][user] Failed to generate token."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}
