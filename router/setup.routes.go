package router

import (
	"github.com/atsuiest/gapigate/config"
	"github.com/atsuiest/gapigate/plugins"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

func SetupRoutes(app *fiber.App) {
	for _, v := range config.GlobalConf.Endpoints {
		g1 := app.Group(v.Base)
		for _, endpoint := range v.Backend {
			if endpoint.Plugin.JwtEnabled {
				validation := config.ValidationsMap["jwt|"+endpoint.Plugin.JwtName]
				g1.Add(endpoint.Method, endpoint.Pattern, func(c *fiber.Ctx) error {
					ok := plugins.CheckClaims(c, validation)
					if ok {
						return c.Next()
					} else {
						return plugins.DenyAccess(c)
					}
				}, proxy.Forward(endpoint.Target.URL))
			} else {
				g1.Add(endpoint.Method, endpoint.Pattern, proxy.Forward(endpoint.Target.URL))
			}
		}
	}
}
