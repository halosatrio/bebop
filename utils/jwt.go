package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/halosatrio/bebop/models"
)

// This should be moved to an environment variable or config file
var jwtKey = []byte("your_secret_key")

func GenerateJWT(user *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   user.ID,
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString(jwtKey)
}

func ValidateJWT(tokenStr string) (int, error) {
	claims := &jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if claims, ok := token.Claims.(*jwt.MapClaims); ok && token.Valid {
		userID := (*claims)["sub"].(int)
		return userID, nil
	}

	return 0, err
}
