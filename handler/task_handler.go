package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nintran52/my-tasks/model"
	"github.com/nintran52/my-tasks/repository"
	"github.com/nintran52/my-tasks/service"
	"gorm.io/gorm"
)

type TaskHandler struct {
	service *service.TaskService
}

func NewTaskHandler(db *gorm.DB) *TaskHandler {
	repo := repository.NewTaskRepository(db)
	service := service.NewTaskService(repo)
	return &TaskHandler{service: service}
}

func (h *TaskHandler) CreateTask(c *fiber.Ctx) error {
	task := new(model.Task)
	if err := c.BodyParser(task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid JSON"})
	}
	if err := h.service.Create(task); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create task"})
	}
	return c.Status(fiber.StatusCreated).JSON(task)
}

func (h *TaskHandler) GetAllTasks(c *fiber.Ctx) error {
	tasks, err := h.service.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch tasks"})
	}
	return c.JSON(tasks)
}

func (h *TaskHandler) GetTaskByID(c *fiber.Ctx) error {
	id := c.Params("id")
	task, err := h.service.GetByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task not found"})
	}
	return c.JSON(task)
}

func (h *TaskHandler) UpdateTask(c *fiber.Ctx) error {
	id := c.Params("id")
	task, err := h.service.GetByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task not found"})
	}
	update := new(model.Task)
	if err := c.BodyParser(update); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid JSON"})
	}
	task.Title = update.Title
	task.Done = update.Done
	if err := h.service.Update(task); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update task"})
	}
	return c.JSON(task)
}

func (h *TaskHandler) DeleteTask(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.service.Delete(id); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task not found"})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
