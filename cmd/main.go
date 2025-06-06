package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nintran52/my-tasks/config"
	"github.com/nintran52/my-tasks/delivery/http"
	"github.com/nintran52/my-tasks/repository/postgres"
	"github.com/nintran52/my-tasks/usecase"
)

func main() {
	app := fiber.New()
	db := config.InitDB()

	taskRepo := postgres.NewTaskPostgresRepo(db)
	taskUC := usecase.NewTaskUsecase(taskRepo)
	http.NewTaskHandler(app, taskUC)

	app.Listen(":3000")
}
