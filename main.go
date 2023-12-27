package main

import (
	"os"

	"github.com/atsuiest/gapigate/api"
	"github.com/atsuiest/gapigate/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

func main() {
	os.Setenv("JWT_SECRET", "WARNING_FOR_TEST_ONLY") // Warning: Use this secret only for testing purposes

	app := fiber.New()
	app.Get("/public/login", api.PublicLoginHandler)
	for _, v := range config.GlobalConf.Endpoints {
		g1 := app.Group(v.Base)
		for _, endpoint := range v.Backend {
			g1.Add(endpoint.Method, endpoint.Pattern, proxy.Forward(endpoint.Target.URL))
		}
	}
	app.Listen(":3000")
}
