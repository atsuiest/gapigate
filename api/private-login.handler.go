package api

import (
	"github.com/atsuiest/gapigate/controller"
	"github.com/atsuiest/gapigate/model"
	"github.com/gofiber/fiber/v2"
)

func PrivateLoginHandler(ctx *fiber.Ctx) error {
	token, err := controller.PrivateLogin(ctx)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(&res500)
	}
	return ctx.Status(fiber.StatusOK).JSON(&model.Response{
		Code: "OK200",
		Data: &model.PublicToken{
			Token:     token,
			PublicKey: "",
		},
		Error:   "",
		Message: "",
	})
}
