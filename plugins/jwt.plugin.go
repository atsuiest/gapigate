package plugins

import (
	"fmt"
	"os"
	"strings"

	"github.com/atsuiest/gapigate/model"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type JwtPlugin struct {
	Secret     string
	Validation *model.Validation
}

func NewJwtPlugin(secret *model.Secret, validation *model.Validation) Plugin {
	plugin := JwtPlugin{
		Secret:     "",
		Validation: validation,
	}
	return &plugin
}

func (jwtp *JwtPlugin) Process(ctx *fiber.Ctx) error {
	return nil
}

// Verify Claims
func DenyAccess(ctx *fiber.Ctx) error {
	res := model.Response{
		Code:    "E403",
		Error:   "Access Denied",
		Message: "The provided access token has no access to requested path",
	}
	return ctx.Status(fiber.StatusForbidden).JSON(&res)
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT", "data": nil})
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT", "data": nil})
}

func CheckClaims(ctx *fiber.Ctx, validation model.Validation) bool {
	var tokenString string

	tokenString = strings.TrimPrefix(ctx.Get("Authorization"), "Bearer ")
	if tokenString == "" {
		return false
	}
	secret := []byte(os.Getenv("JWT_SECRET"))
	tokenByte, err := jwt.Parse(tokenString, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", jwtToken.Header["alg"])
		}
		return secret, nil
	})
	if tokenByte == nil && err != nil {
		return false
	}
	claims, ok := tokenByte.Claims.(jwt.MapClaims)
	if !ok || !tokenByte.Valid {
		return false
	}
	valid := true
	for _, v := range validation.Claims {
		if !valid {
			return false
		}
		valid = claims[v.Key] != nil && (v.Value == "" || claims[v.Key] == v.Value)
	}
	return true
}
