package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "secured"

func Generate(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Minute * 30).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}
