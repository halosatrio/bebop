package handlers

import (
	"fmt"

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
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := h.service.Register(user)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to register user."})
		return
	}
	c.JSON(200, gin.H{"message": "Registration successful."})
}

func (h *UserHandler) Authenticate(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	isValid, err := h.service.Authenticate(user.Email, user.Password)
	if err != nil || !isValid {
		c.JSON(401, gin.H{"error": "Authentication failed."})
		return
	}
	fmt.Print("handler", user)

	token, err := utils.GenerateJWT(&user)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate token."})
		return
	}
	c.JSON(200, gin.H{"token": token})
}
