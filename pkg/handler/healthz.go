package handler

import "github.com/gofiber/fiber/v2"

func HealthZHandler(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}
