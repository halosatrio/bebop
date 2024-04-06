package utils

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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

func ErrorResponseUnauthorizedJwt(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{
		"status":  status,
		"message": "Failed",
		"data":    nil,
		"error":   "[middleware][jwt] " + message,
	})
}

// this function identify as middleware
// should be in different package: middleware
func JWTAuth() gin.HandlerFunc {
	secret := os.Getenv("JWT_SECRET")

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			ErrorResponseUnauthorizedJwt(c, http.StatusUnauthorized, "[middleware][jwt] Authorization header prefix is not provided")
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			ErrorResponseUnauthorizedJwt(c, http.StatusUnauthorized, "[middleware][jwt] Invalid or missing Bearer token")
			c.Abort()
			return
		}

		tokenString := parts[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

		if err != nil {
			ErrorResponseUnauthorizedJwt(c, http.StatusUnauthorized, "[middleware][jwt] Invalid token")
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			ErrorResponseUnauthorizedJwt(c, http.StatusUnauthorized, "[middleware][jwt] Invalid token")
			c.Abort()
			return
		}

		c.Set("user_id", claims["sub"].(string))
		c.Set("email", claims["email"].(string))
		c.Next()
	}
}
