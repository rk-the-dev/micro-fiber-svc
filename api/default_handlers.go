package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func Ping(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "pong"})
}
func Health(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "server is running"})
}
