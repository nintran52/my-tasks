package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nintran52/my-tasks/controller"
	"github.com/nintran52/my-tasks/database"
)

func main() {
	database.ConnectDB()
	db := database.DB
	app := fiber.New()

	taskController := controller.NewTaskController(db)
	taskController.RegisterRoutes(app)

	app.Listen(":3000")
}
