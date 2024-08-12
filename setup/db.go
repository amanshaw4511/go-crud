package setup

import (
	"crud/model"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database test.db : %s", err))
	}

	err = database.AutoMigrate(&model.Book{})

	if err != nil {
		panic(fmt.Sprintf("Failed during DB auto migration: %s", err))
	}

	return database
}
