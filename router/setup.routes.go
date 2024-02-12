package router

import (
	"github.com/atsuiest/gapigate/api"
	"github.com/atsuiest/gapigate/config"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	for _, v := range config.GlobalConf.Endpoints {
		g1 := app.Group(v.Base)
		for _, backend := range v.Backend {
			g1.All(backend.Pattern, func(ctx *fiber.Ctx) error {
				return api.RouteHandler(ctx)
			})
		}
	}
}
