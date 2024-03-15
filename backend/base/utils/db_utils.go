package utils

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDb() *gorm.DB {
	var DATABASE_URL string = os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(DATABASE_URL), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(fmt.Sprintf("Error connecting to database: %s", err.Error()))
	}

	return db
}
