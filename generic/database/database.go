package database

import (
	"bahno_bot/generic/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase(host, user, password, dbname string, port uint) (db *gorm.DB) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Europe/Prague", host, user, password, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Printf("Couldn't connect to database (host=%s).", host)
		return
	}

	db.AutoMigrate(&models.Substance{}, &models.User{}, &models.Record{})
	// db.AutoMigrate(&models.User{})
	// db.AutoMigrate(&models.Record{})

	return
}
