package api

import (
	"github.com/atsuiest/gapigate/controller"
	"github.com/atsuiest/gapigate/model"
	"github.com/gofiber/fiber/v2"
)

func PublicLoginHandler(ctx *fiber.Ctx) error {
	token, err := controller.PublicLogin(ctx)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(&RES500)
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
