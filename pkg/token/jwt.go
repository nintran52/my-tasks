package token

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/nintran52/my-tasks/internal/common"
)

func GenerateToken(userID uint, duration time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(duration).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(common.SecretKey))
}
