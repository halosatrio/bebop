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
		utils.ErrorResponseDataNull(c, http.StatusBadRequest, "[handler][user] "+err.Error())
		return
	}

	err := h.registerHandler(reqRegister)
	if err != nil {
		utils.ErrorResponseDataNull(c, http.StatusInternalServerError, "[handler][user] Failed to register user.")
		return
	}

	utils.SuccessResponseWithMessage(c, http.StatusOK, "Registration successful!")
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
		utils.ErrorResponseDataNull(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.authenticateHandler(reqAuth.Email, reqAuth.Password)
	if err != nil || user == nil {
		utils.ErrorResponseDataNull(c, http.StatusUnauthorized, "[handler][user] Authentication failed.")
		return
	}

	token, err := utils.GenerateJWT(user)
	if err != nil {
		utils.ErrorResponseDataNull(c, http.StatusInternalServerError, "[handler][user] Failed to generate token.")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success",
		"token":   token,
	})
}
