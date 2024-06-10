package db

import (
	"task-manager-be/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
    database, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})

    if err != nil {
        panic("failed to connect to database")
    }

    err = database.AutoMigrate(&models.Task{})

    if err != nil {
        return
    }

    DB = database
}
