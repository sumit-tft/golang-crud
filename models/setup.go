package models

import (
	envaccess "crudgoblog/access"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := envaccess.GoDotEnvVariable("CONNECTION_STRING")
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Post{})

	DB = database
}
