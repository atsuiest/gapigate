package main

import (
	"github.com/atsuiest/gapigate/api"
	"github.com/atsuiest/gapigate/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	//os.Setenv("JWT_SECRET", "JwtSecretKey") // Warning: Use this secret only for testing purposes

	app := fiber.New()
	// Initialize default config
	app.Use(logger.New())

	// Or extend your config for customization
	// Logging remote IP and Port
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))
	app.Get("/public/login", api.PublicLoginHandler)
	app.Get("/private/login", api.PrivateLoginHandler)
	router.SetupRoutes(app)
	app.Listen("127.0.0.1:3000") // Using ":3000" may cause Security prompts on Windows
}
