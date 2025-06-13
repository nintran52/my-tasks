package common

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func RespondError(c *fiber.Ctx, code int, msg string) error {
	return c.Status(code).JSON(fiber.Map{
		"error":   http.StatusText(code),
		"message": msg,
	})
}
