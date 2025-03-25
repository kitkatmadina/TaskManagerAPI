package db

import (
	"log"

	"github.com/kitkatmadina/TaskManagerAPI/models"

	"github.com/glebarez/sqlite" // ✅ use CGO-free sqlite driver
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	var err error
	DB, err = gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Database connected successfully!")

	err = DB.AutoMigrate(&models.Task{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Database migrated successfully!")
}
