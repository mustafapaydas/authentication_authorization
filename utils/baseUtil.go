package utils

import (
	"github.com/gofiber/fiber/v2"
)

func NotFound(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).SendString("Endpoint not found")
}

func UnprocessableEntity(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusUnprocessableEntity).JSON(message)
}
func InternalServerError(c *fiber.Ctx) error {
	return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
}
func InternalServerErrorMessage(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusInternalServerError).SendString(message)
}
func BadRequest(c *fiber.Ctx) error {
	return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
}
