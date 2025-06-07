package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nintran52/my-tasks/config"
	"github.com/nintran52/my-tasks/handler"
)

func main() {
	db := config.InitDatabase()
	h := handler.NewTaskHandler(db)

	app := fiber.New()

	app.Post("/tasks", h.CreateTask)
	app.Get("/tasks", h.GetAllTasks)
	app.Get("/tasks/:id", h.GetTaskByID)
	app.Put("/tasks/:id", h.UpdateTask)
	app.Delete("/tasks/:id", h.DeleteTask)

	app.Listen(":3000")
}
