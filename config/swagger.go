package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
)

func AddSwaggerRoutes(app *fiber.App) {
	// setup swagger
	sw := app.Group("/swagger")

	sw.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	sw.Get("/*", swagger.HandlerDefault)
}
