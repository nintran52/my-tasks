package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nintran52/my-tasks/internal/adapter/handler"
	"github.com/nintran52/my-tasks/internal/adapter/repository"
	"github.com/nintran52/my-tasks/internal/application/service"
	"github.com/nintran52/my-tasks/pkg/database"
)

func main() {
	db := database.InitPostgres()
	repo := repository.NewTaskRepository(db)
	svc := service.NewTaskService(repo)
	h := handler.NewTaskHandler(svc)

	app := fiber.New()
	h.RegisterRoutes(app)
	app.Listen(":3000")
}
