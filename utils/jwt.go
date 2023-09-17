package utils

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/halosatrio/bebop/models"
)

func GenerateJWT(user *models.User) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   user.ID,
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(secret))
}
