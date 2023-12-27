package plugins

import "github.com/gofiber/fiber/v2"

type Plugin interface {
	Process(ctx *fiber.Ctx) error
}
