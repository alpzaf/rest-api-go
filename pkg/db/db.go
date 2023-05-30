package db

import (
	"log"
	"rest-go/pkg/pkg/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dbURL := "postgres://admin:password@localhost:5432/gocrud"

    db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

    if err != nil {
        log.Fatalln(err)
    }

    db.AutoMigrate(&models.User{})

    return db
}
