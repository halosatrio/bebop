package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/halosatrio/bebop/models"
)

var jwtKey = []byte("your_secret_key") // This should be moved to an environment variable or config file

func GenerateJWT(user *models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = user.ID
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // token will expire after 24 hours

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
