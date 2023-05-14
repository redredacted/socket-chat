package router

import "github.com/gofiber/fiber/v2"

func healthCheck(c *fiber.Ctx) error {
	return c.SendString("OK")
}
