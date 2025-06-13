package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/nintran52/my-tasks/internal/common"
	"github.com/nintran52/my-tasks/internal/delivery/http"
	"github.com/nintran52/my-tasks/internal/repository/postgres"
	"github.com/nintran52/my-tasks/internal/usecase"
	"github.com/nintran52/my-tasks/pkg/database"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: common.ErrorHandlerFiber,
	})
	app.Use(recover.New())
	db := database.InitDB()

	taskRepo := postgres.NewTaskPostgresRepo(db)
	taskUC := usecase.NewTaskUsecase(taskRepo)
	http.NewTaskHandler(app, taskUC)

	app.Listen(":3000")
}
