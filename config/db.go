package config

import (
	"fmt"
	"log"

	"github.com/nintran52/my-tasks/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase() *gorm.DB {
	dsn := "host=localhost user=postgres password=password dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	db.AutoMigrate(&model.User{}, &model.Task{})
	fmt.Println("Database connected and migrated")
	return db
}
