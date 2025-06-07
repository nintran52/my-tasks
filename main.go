package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nintran52/my-tasks/adapter/http"
	"github.com/nintran52/my-tasks/adapter/postgres"
	"github.com/nintran52/my-tasks/database"
	service "github.com/nintran52/my-tasks/usecase"
)

func main() {
	db := database.InitPostgres()
	repo := postgres.NewTaskRepository(db)
	svc := service.NewTaskService(repo)
	h := http.NewTaskHandler(svc)

	app := fiber.New()
	h.RegisterRoutes(app)
	app.Listen(":3000")
}
