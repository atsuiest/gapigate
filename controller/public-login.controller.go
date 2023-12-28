package controller

import (
	"time"

	"github.com/atsuiest/gapigate/tools"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func PublicLogin(ctx *fiber.Ctx) (string, error) {
	// Create the Claims
	claims := jwt.MapClaims{
		"system":    "public",
		"publicKey": uuid.New().String(),
		"exp":       time.Now().Add(time.Minute * 5).Unix(),
	}

	return tools.GeneratePublicToken(&claims)
}
