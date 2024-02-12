package plugins

import (
	"strings"

	"github.com/atsuiest/gapigate/config"
	"github.com/atsuiest/gapigate/model"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type JwtPlugin struct {
	Secret     string
	Validation *model.Validation
}

func (jwtp *JwtPlugin) Process(ctx *fiber.Ctx) error {
	return nil
}

func CheckClaims(ctx *fiber.Ctx, validation model.Validation) bool {
	tokenString := strings.TrimPrefix(ctx.Get("Authorization"), "Bearer ")
	if tokenString == "" {
		return false
	}
	secret := config.GlobalConf.JwtSecret
	tokenByte, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
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
	ctx.Set("X-User", claims["user"].(string))
	ctx.Set("X-Site", claims["site"].(string))
	ctx.Set("X-Roles", claims["roles"].(string))
	return true
}
