package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"fmt"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}
	fmt.Print("Connected to DB...")
	DB = database
	fmt.Print("DB Details : ", DB)
}
