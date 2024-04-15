package initializers

import (
	"brainyquiz/internal/domain/entities"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	dbURL := os.Getenv("DB_URL")
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}
	err = db.AutoMigrate(&entities.User{})
	if err != nil {
		panic("Failed to migrate database")
	}
	return db, nil
}
