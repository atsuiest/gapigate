package tools

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func GeneratePublicToken(claims *jwt.MapClaims) (string, error) {
	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return t, nil
}

func ValidateClaims(token string, key string, value string) bool {
	return false
}
