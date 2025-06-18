package token

import (
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	secretKey = "your-secret-key" // Replace with your actual secret key
)

func GenerateToken(userID uint, duration time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(duration).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}
