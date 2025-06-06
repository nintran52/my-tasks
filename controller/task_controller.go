package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nintran52/my-tasks/model"
	"gorm.io/gorm"
)

type TaskController struct {
	db *gorm.DB
}

func NewTaskController(db *gorm.DB) *TaskController {
	return &TaskController{db: db}
}

func (tc *TaskController) RegisterRoutes(app *fiber.App) {
	app.Post("/tasks", tc.Create)
	app.Get("/tasks", tc.GetAll)
	app.Get("/tasks/:id", tc.GetByID)
	app.Put("/tasks/:id", tc.Update)
	app.Delete("/tasks/:id", tc.Delete)
}

func (tc *TaskController) Create(c *fiber.Ctx) error {
	task := new(model.Task)
	if err := c.BodyParser(task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid JSON"})
	}
	if err := tc.db.Create(task).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create task"})
	}
	return c.Status(fiber.StatusCreated).JSON(task)
}

func (tc *TaskController) GetAll(c *fiber.Ctx) error {
	var tasks []model.Task
	if err := tc.db.Find(&tasks).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch tasks"})
	}
	return c.JSON(tasks)
}

func (tc *TaskController) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var task model.Task
	if err := tc.db.First(&task, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task not found"})
	}
	return c.JSON(task)
}

func (tc *TaskController) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	var task model.Task
	if err := tc.db.First(&task, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task not found"})
	}
	var update model.Task
	if err := c.BodyParser(&update); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid JSON"})
	}
	task.Title = update.Title
	task.Done = update.Done
	if err := tc.db.Save(&task).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update task"})
	}
	return c.JSON(task)
}

func (tc *TaskController) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := tc.db.Delete(&model.Task{}, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task not found"})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
