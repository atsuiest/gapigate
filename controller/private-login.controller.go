package controller

import (
	"time"

	"github.com/atsuiest/gapigate/tools"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func PrivateLogin(ctx *fiber.Ctx) (string, error) {
	// Create the Claims
	claims := jwt.MapClaims{
		"system":    "private",
		"user":      "root",
		"role":      "administrator",
		"publicKey": uuid.New().String(),
		"exp":       time.Now().Add(time.Minute * 5).Unix(),
	}

	return tools.GeneratePublicToken(&claims)
}
