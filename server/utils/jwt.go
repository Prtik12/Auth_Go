package utils

import (
	"time"

	"github.com/Prtik12/auth_Golang/models"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("my_secret_key")

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateJWT(user models.User) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &Claims{
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
