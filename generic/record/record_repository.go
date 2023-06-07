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

func (ur *recordRepository) Create(record models.Record) (err error) {
	result := ur.database.Create(&record)
	err = result.Error

	return
}

func (ur *recordRepository) GetAll(userId uint) (records []models.Record, err error) {
	ur.database.Where("user_id = ?", userId).Preload("Substance").Find(&records)

	return
}

func (ur *recordRepository) GetAllPaged(userId uint, page, pageSize int) (records []models.Record, count int64, err error) {
	ur.database.Where("user_id = ?", userId).Preload("Substance").Limit(pageSize).Offset(page * pageSize).Find(&records).Count(&count)

	return
}

//func (ur *recordRepository) GetRecordsCount(userId uint) (count int64, err error) {
//	ur.database.Where("user_id = ?", userId).Find(records).Count(&count)
//
//	return
//}

func (ur *recordRepository) GetLast(userId uint) (record models.Record, err error) {
	ur.database.Preload("Substance").Last(&record, userId)

	return
}
