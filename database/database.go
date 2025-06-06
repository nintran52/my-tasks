package database

import (
	"fmt"
	"log"

	"github.com/nintran52/my-tasks/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=localhost user=postgres password=password dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}
	db.AutoMigrate(&model.User{}, &model.Task{})
	DB = db
	fmt.Println("Database connected and migrated")
}
