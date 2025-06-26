package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	_ "github.com/nintran52/my-tasks/docs" // Import generated docs
	"github.com/nintran52/my-tasks/internal/common"
	"github.com/nintran52/my-tasks/internal/delivery/http"
	"github.com/nintran52/my-tasks/internal/repository/postgres"
	"github.com/nintran52/my-tasks/internal/usecase"
	"github.com/nintran52/my-tasks/pkg/cache"
	"github.com/nintran52/my-tasks/pkg/database"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

// @title My Tasks API
// @version 1.0
// @description A task management API with authentication
// @host localhost:3000
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: common.ErrorHandlerFiber,
	})
	app.Use(recover.New())
	app.Use(common.LoggingFiber)

	// Swagger route
	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	db := database.InitDB()
	cache.InitRedis()

	userRepo := postgres.NewUserPostgresRepo(db)
	userUC := usecase.NewUserUsecase(userRepo)
	http.NewUserHandler(app, userUC)

	taskRepo := postgres.NewTaskPostgresRepo(db)
	taskUC := usecase.NewTaskUsecase(taskRepo)
	http.NewTaskHandler(app, taskUC)

	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}
}
