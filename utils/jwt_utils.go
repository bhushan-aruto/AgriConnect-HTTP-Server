package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(userId, userName, userEmail string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":        userId,
		"user_name": userName,
		"email":     userEmail,
		"expiry":    time.Now().Add(365 * 24 * time.Hour).Unix(),
	})
	tokenString, err := token.SignedString([]byte("hello@123456"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
