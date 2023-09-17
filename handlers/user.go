package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/halosatrio/bebop/models"
	"github.com/halosatrio/bebop/service"
	"github.com/halosatrio/bebop/utils"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "[handler][user] " + err.Error()})
		return
	}

	err := h.service.Register(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "[handler][user] Failed to register user."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Registration successful."})
}

func (h *UserHandler) Authenticate(c *gin.Context) {
	var reqUser models.User
	if err := c.ShouldBindJSON(&reqUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.Authenticate(reqUser.Email, reqUser.Password)
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
