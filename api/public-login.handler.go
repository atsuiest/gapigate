package api

import (
	"github.com/atsuiest/gapigate/model"
	"github.com/atsuiest/gapigate/tools"
	"github.com/gofiber/fiber/v2"
)

func PublicLoginHandler(ctx *fiber.Ctx) error {
	res := &model.Response{
		Code:    "KO500",
		Data:    nil,
		Error:   "Unavailable",
		Message: "Service temporary unavailable",
	}
	token, err := tools.GeneratePublicToken(ctx)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(&res)
	}
	res.Code = "OK200"
	res.Data = &model.PublicToken{
		Token:     token,
		PublicKey: "",
	}
	res.Message = ""
	res.Error = ""
	return ctx.Status(fiber.StatusOK).JSON(&res)
}
