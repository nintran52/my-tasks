package http

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/nintran52/my-tasks/internal/common"
	"github.com/nintran52/my-tasks/internal/domain"
	"github.com/nintran52/my-tasks/internal/usecase"
)

type TaskHandler struct {
	Usecase *usecase.TaskUsecase
}

func NewTaskHandler(app *fiber.App, uc *usecase.TaskUsecase) {
	handler := &TaskHandler{Usecase: uc}

	app.Post("/tasks", handler.Create)
	app.Get("/tasks", handler.GetAll)
	app.Get("/tasks/:id", handler.GetByID)
	app.Put("/tasks/:id", handler.Update)
	app.Delete("/tasks/:id", handler.Delete)
}

func (h *TaskHandler) Create(c *fiber.Ctx) error {
	var task domain.Task
	if err := c.BodyParser(&task); err != nil {
		return common.RespondError(c, fiber.StatusBadRequest, "Invalid JSON")
	}

	if err := task.Validate(); err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	if err := h.Usecase.CreateTask(&task); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create task"})
	}
	return c.Status(fiber.StatusCreated).JSON(task)
}

func (h *TaskHandler) GetAll(c *fiber.Ctx) error {
	limitParam := c.Query("limit", "10")
	offsetParam := c.Query("offset", "0")
	titleParam := c.Query("title") // example: shopping
	doneParam := c.Query("done")   // true or false

	limit, _ := strconv.Atoi(limitParam)
	offset, _ := strconv.Atoi(offsetParam)

	if limit <= 0 || limit > 100 {
		limit = 10
	}

	tasks, total, err := h.Usecase.GetTasks(limit, offset, domain.Task{Title: titleParam, Done: doneParam == "true"})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch tasks"})
	}
	return c.JSON(common.PaginatedResponse{
		Data:       tasks,
		Total:      total,
		Limit:      limit,
		Offset:     offset,
		NextOffset: offset + limit,
	})
}

func (h *TaskHandler) GetByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	task, err := h.Usecase.GetTaskByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task not found"})
	}
	return c.JSON(task)
}

func (h *TaskHandler) Update(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	task, err := h.Usecase.GetTaskByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task not found"})
	}
	if err := c.BodyParser(task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid JSON"})
	}
	if err := h.Usecase.UpdateTask(task); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update task"})
	}
	return c.JSON(task)
}

func (h *TaskHandler) Delete(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	if err := h.Usecase.DeleteTask(uint(id)); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task not found"})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
