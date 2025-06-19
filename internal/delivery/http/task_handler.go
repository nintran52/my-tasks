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

	app.Post("/tasks", common.AuthMiddleware, handler.Create)
	app.Get("/tasks", common.AuthMiddleware, handler.GetAll)
	app.Get("/tasks/:id", common.AuthMiddleware, handler.GetByID)
	app.Put("/tasks/:id", common.AuthMiddleware, handler.Update)
	app.Delete("/tasks/:id", common.AuthMiddleware, handler.Delete)
}

func (h *TaskHandler) Create(c *fiber.Ctx) error {
	var task domain.Task
	if err := c.BodyParser(&task); err != nil {
		return common.ResponseError(c, fiber.StatusBadRequest, "Invalid JSON")
	}

	if err := task.Validate(); err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	userID := c.Locals("user_id").(uint)
	task.CreatedBy = userID
	task.UpdatedBy = userID
	if err := h.Usecase.CreateTask(&task); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to create task")
	}

	return common.ResponseCreated(c, task.ID)
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

	tasks, total, err := h.Usecase.GetTasks(limit, offset, domain.Task{Title: titleParam, Done: doneParam == "true", CreatedBy: c.Locals("user_id").(uint)})
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to fetch tasks")
	}

	return common.ResponseSuccess(c, common.PaginatedResponse{
		Data:   tasks,
		Total:  total,
		Limit:  limit,
		Offset: offset,
	})
}

func (h *TaskHandler) GetByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	userID := c.Locals("user_id").(uint)

	task, err := h.Usecase.GetTaskByID(uint(id))
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Task not found")
	}

	if task.CreatedBy != userID {
		return fiber.NewError(fiber.StatusForbidden, "You do not have permission to access this task")
	}

	return common.ResponseSuccess(c, task)
}

func (h *TaskHandler) Update(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	userID := c.Locals("user_id").(uint)

	task, err := h.Usecase.GetTaskByID(uint(id))
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Task not found")
	}

	if task.CreatedBy != userID {
		return fiber.NewError(fiber.StatusForbidden, "You do not have permission to update this task")
	}

	if err := c.BodyParser(task); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid JSON")
	}
	if err := h.Usecase.UpdateTask(task); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to update task")
	}

	return common.ResponseSuccess(c, task)
}

func (h *TaskHandler) Delete(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	userID := c.Locals("user_id").(uint)

	task, err := h.Usecase.GetTaskByID(uint(id))
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Task not found")
	}

	if task.CreatedBy != userID {
		return fiber.NewError(fiber.StatusForbidden, "You do not have permission to delete this task")
	}

	if err := h.Usecase.DeleteTask(uint(id)); err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Task not found")
	}

	return common.ResponseNoContent(c)
}
