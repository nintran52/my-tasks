package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/nintran52/my-tasks/internal/domain/model"
	"github.com/nintran52/my-tasks/internal/domain/port"
)

type TaskHandler struct {
	service port.TaskService
}

func NewTaskHandler(service port.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

func (h *TaskHandler) RegisterRoutes(app *fiber.App) {
	app.Post("/tasks", h.CreateTask)
	app.Get("/tasks", h.GetTasks)
	app.Get("/tasks/:id", h.GetTask)
	app.Put("/tasks/:id", h.UpdateTask)
	app.Delete("/tasks/:id", h.DeleteTask)
}

func (h *TaskHandler) CreateTask(c *fiber.Ctx) error {
	task := new(model.Task)
	if err := c.BodyParser(task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid JSON"})
	}
	if err := h.service.CreateTask(task); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(task)
}

func (h *TaskHandler) GetTasks(c *fiber.Ctx) error {
	tasks, err := h.service.GetTasks()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(tasks)
}

func (h *TaskHandler) GetTask(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	task, err := h.service.GetTask(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "not found"})
	}
	return c.JSON(task)
}

func (h *TaskHandler) UpdateTask(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	task := new(model.Task)
	if err := c.BodyParser(task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid JSON"})
	}
	task.ID = uint(id)
	if err := h.service.UpdateTask(task); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(task)
}

func (h *TaskHandler) DeleteTask(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	if err := h.service.DeleteTask(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
