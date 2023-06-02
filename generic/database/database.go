package database

import (
	"fmt"
	"log"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"

)

func NewDatabase(host, user, password, dbname string, port uint) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Europe/Prague", host, user, password, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		// TODO	
		log.Println("Big problem.")
	}

	return db
}

func CloseDBConnection(client *gorm.DB) {
	if client == nil {
		return
	}

	// TODO

	log.Println("Connection to DB closed.")
}
