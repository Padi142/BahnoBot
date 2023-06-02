package record

import (
	"bahno_bot/generic/models"
	"context"
	"gorm.io/gorm"
)

type recordRepository struct {
	database  gorm.DB 
	collection string
}

func NewRecordRepository(db gorm.DB, collection string) RecordRepository {
	return &recordRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *recordRepository) Create(c context.Context, userId string, record models.Record) ( err error) {
	result :=  ur.database.Create(&record);
	err = result.Error

	return
}

func (ur *recordRepository) Fetch(c context.Context, userId string) (records []models.Record, err error) {
	// if err := ur.database.Model(&models.User{}).
	// Where("id = ?", userId).
	// Joins("JOIN records ON records.user_id = users.id"). 
	// Find(&records). 
	// Error; err != nil {
	// 	return nil, err

	// }

	// return records, err
	return nil, nil
}

func (ur *recordRepository) GetLastRecord(c context.Context, userId string) (record models.Record, err error) {
	ur.database.Last(&record);

	return 
}
