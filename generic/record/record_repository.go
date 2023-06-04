package record

import (
	"bahno_bot/generic/models"

	"gorm.io/gorm"
)

type recordRepository struct {
	database *gorm.DB
}

func NewRecordRepository(db *gorm.DB) RecordRepository {
	return &recordRepository{
		database: db,
	}
}

func (ur *recordRepository) Create(record models.Record) error {
	return ur.database.Create(&record).Error
}

func (ur *recordRepository) GetAll() (records []models.Record, err error) {
	err = ur.database.Find(&records).Error

	return
}
