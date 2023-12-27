package tools

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func GeneratePublicToken(ctx *fiber.Ctx) (string, error) {
	// Create the Claims
	claims := jwt.MapClaims{
		"name":  "John Doe",
		"admin": true,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
		// return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return t, nil
	// return ctx.JSON(fiber.Map{"token": t})
}

func ValidateClaims(token string, key string, value string) bool {
	return false
}
