package router

import (
	"github.com/gofiber/fiber/v2"
	"go-mongo-starter/handlers"
)

func authMiddleware(c *fiber.Ctx) error {
	if _, err := handlers.ExtractClaims(c); err != nil {
		return c.Status(401).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Unauthorized",
		})
	}

	return c.Next()
}
