package database

import (
	"bahno_bot/generic/models"
	"fmt"
	"gorm.io/gorm/logger"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase(host, user, password, dbname string, port uint) (db *gorm.DB) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Europe/Prague", host, user, password, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})

	if err != nil {
		log.Printf("Couldn't connect to database (host=%s).", host)
		return
	}

	err = db.AutoMigrate(&models.Substance{}, &models.User{}, &models.Record{})

	if err != nil {
		log.Printf(err.Error())
		return
	}

	return
}
