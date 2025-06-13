package common

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandlerFiber(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	msg := "Internal Server Error"
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
		msg = e.Message
	}
	return c.Status(code).JSON(fiber.Map{
		"error":   http.StatusText(code),
		"message": msg,
	})
}
