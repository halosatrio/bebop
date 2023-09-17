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
	var reqRegister models.AuthRequset
	if err := c.ShouldBindJSON(&reqRegister); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "[handler][user] " + err.Error()})
		return
	}

	err := h.service.Register(reqRegister)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "[handler][user] Failed to register user."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Registration successful."})
}

func (h *UserHandler) Authenticate(c *gin.Context) {
	var reqAuth models.AuthRequset
	if err := c.ShouldBindJSON(&reqAuth); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.Authenticate(reqAuth.Email, reqAuth.Password)
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
