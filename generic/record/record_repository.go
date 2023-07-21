package record

import (
	"bahno_bot/generic/models"
	"time"

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
	ur.database.Where("user_id = ?", userId).Order("created_at DESC").Preload("Substance").Preload("User").Find(&records)

	return
}

func (ur *recordRepository) GetAllPaged(userId uint, page, pageSize int) (records []models.Record, count int64, err error) {
	ur.database.Where("user_id = ?", userId).Order("created_at DESC").Preload("Substance").Preload("User").Limit(pageSize).Offset(page * pageSize).Find(&records)

	return
}

func (ur *recordRepository) GetAllPagedForSubstance(userId, substanceId uint, page, pageSize int) (records []models.Record, count int64, err error) {
	ur.database.Where("user_id = ?", userId).Where("substance_id = ?", substanceId).Order("created_at DESC").Preload("Substance").Preload("User").Limit(pageSize).Offset(page * pageSize).Find(&records)

	return
}

func (ur *recordRepository) GetAllInTimePeriod(userId uint, time time.Time) (records []models.Record, err error) {
	ur.database.Where("user_id = ?", userId).Where("created_at >= ?", time).Preload("Substance").Preload("User").Find(&records)

	return
}

func (ur *recordRepository) GetLeaderboardInTimePeriod(time time.Time) (leaderboard []models.LeaderboardOccurrence, err error) {
	err = ur.database.Table("records").Select("user_id, COUNT(*) as occurrence").
		Where("created_at >= ?", time).
		Group("user_id").
		Order("occurrence DESC").
		Preload("Substance").
		Preload("User").
		Find(&leaderboard).Error

	return
}

func (ur *recordRepository) GetLeaderboardForSubstanceInTimePeriod(substanceId uint, time time.Time) (leaderboard []models.LeaderboardOccurrence, err error) {
	err = ur.database.Table("records").Select("user_id, COUNT(*) as occurrence").
		Where("substance_id = ?", substanceId).
		Where("created_at >= ?", time).
		Group("user_id").
		Order("occurrence DESC").
		Preload("Substance").
		Preload("User").
		Find(&leaderboard).Error

	return
}

//func (ur *recordRepository) GetRecordsCount(userId uint) (count int64, err error) {
//	ur.database.Where("user_id = ?", userId).Find(records).Count(&count)
//
//	return
//}

func (ur *recordRepository) GetLast(userId uint) (record models.Record, err error) {
	err = ur.database.
		Preload("Substance").
		Where("user_id = ?", userId).
		Order("created_at DESC").
		Limit(1).
		Find(&record).
		Error

	return
}

func (ur *recordRepository) GetLastForSubstance(substanceId, userId uint) (record models.Record, err error) {
	res := ur.database.
		Preload("Substance").
		Where("user_id = ?", userId).
		Where("substance_id = ?", substanceId).
		Order("created_at DESC").
		Limit(1).
		Find(&record)

	err = res.Error

	if res.RowsAffected == 0 {
		err = gorm.ErrRecordNotFound
	}

	return
}
