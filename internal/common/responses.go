package common

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type PaginatedResponse struct {
	Data   any   `json:"data"`
	Total  int64 `json:"total"`
	Limit  int   `json:"limit"`
	Offset int   `json:"offset"`
}

func ResponseError(c *fiber.Ctx, code int, msg string) error {
	return c.Status(code).JSON(fiber.Map{
		"error":   http.StatusText(code),
		"message": msg,
	})
}

func ResponseCreated(c *fiber.Ctx, id uint) error {
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"id": id,
	})
}

func ResponseSuccess(c *fiber.Ctx, data any) error {
	if data == nil {
		return c.SendStatus(http.StatusNoContent)
	}
	return c.Status(http.StatusOK).JSON(data)
}

func ResponseNoContent(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusNoContent)
}
