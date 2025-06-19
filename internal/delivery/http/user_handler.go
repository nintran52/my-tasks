package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nintran52/my-tasks/internal/common"
	"github.com/nintran52/my-tasks/internal/domain"
	"github.com/nintran52/my-tasks/internal/usecase"
)

type UserHandler struct {
	Usecase *usecase.UserUsecase
}

func NewUserHandler(app *fiber.App, uc *usecase.UserUsecase) {
	handler := &UserHandler{Usecase: uc}

	app.Post("/v1/users/register", handler.Create)
}

func (h *UserHandler) Create(c *fiber.Ctx) error {
	var user domain.User
	if err := c.BodyParser(&user); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid JSON")
	}

	if err := user.Validate(); err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	if err := h.Usecase.CreateUser(&user); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return common.ResponseCreated(c, user.ID)
}
