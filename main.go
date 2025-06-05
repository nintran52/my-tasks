package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// User model
type User struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Email     string `gorm:"uniqueIndex"`
	Password  string
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

// Task model
type Task struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

func initDatabase() {
	dsn := "host=localhost user=postgres password=password dbname=postgres port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	// Migrate schema
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Task{})
	fmt.Println("Database connection established & migrated")
}

func main() {
	app := fiber.New()
	initDatabase()

	// Create task
	app.Post("/tasks", func(c *fiber.Ctx) error {
		var task Task
		if err := c.BodyParser(&task); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid JSON"})
		}
		if err := db.Create(&task).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create task"})
		}
		return c.Status(fiber.StatusCreated).JSON(task)
	})

	// Get all tasks
	app.Get("/tasks", func(c *fiber.Ctx) error {
		var tasks []Task
		if err := db.Find(&tasks).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch tasks"})
		}
		return c.JSON(tasks)
	})

	// Get task by ID
	app.Get("/tasks/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var task Task
		if err := db.First(&task, id).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task not found"})
		}
		return c.JSON(task)
	})

	// Update task
	app.Put("/tasks/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var task Task
		if err := db.First(&task, id).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task not found"})
		}
		var update Task
		if err := c.BodyParser(&update); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid JSON"})
		}
		task.Title = update.Title
		task.Done = update.Done
		if err := db.Save(&task).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update task"})
		}
		return c.JSON(task)
	})

	// Delete task
	app.Delete("/tasks/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		if err := db.Delete(&Task{}, id).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task not found"})
		}
		return c.SendStatus(fiber.StatusNoContent)
	})

	log.Fatal(app.Listen(":3000"))
}
