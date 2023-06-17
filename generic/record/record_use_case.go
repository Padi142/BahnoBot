package record

import (
	"bahno_bot/generic/models"
	"errors"

	"gorm.io/gorm"
)

type UseCase struct {
	recordRepository RecordRepository
}

func NewRecordUseCase(db *gorm.DB) UseCase {
	return UseCase{
		recordRepository: NewRecordRepository(db),
	}
}

func (us UseCase) CreateRecordUseCase(recordRepository RecordRepository, record models.Record) error {
	err := recordRepository.Create(record)

	if err != nil {
		return err
	}

	return nil
}

func (us UseCase) CreateNewRecord(userId uint, record models.Record) (*models.Record, error) {
	err := us.recordRepository.Create(record)

	if err != nil {
		return nil, err
	}

	return &record, nil
}

func (us UseCase) GetAllRecords(userId uint) ([]models.Record, error) {
	return us.recordRepository.GetAll(userId)
}

func (us UseCase) GetPagedRecords(userId uint, page, pageSize int) ([]models.Record, int64, error) {
	return us.recordRepository.GetAllPaged(userId, page, pageSize)
}
func (us UseCase) GetLastRecord(userId uint) (models.Record, error) {
	return us.recordRepository.GetLast(userId)
}

func (us UseCase) GetLastRecordForSubstance(substanceId, userId uint) (models.Record, error) {
	return us.recordRepository.GetLastForSubstance(substanceId, userId)
}

func (us UseCase) GetLeaderboardTime(durationString string) ([]models.LeaderboardOccurrence, error) {
	duration := models.GetTimeDuration(durationString)
	if duration == "" {
		return nil, errors.New("duration not valid")
	}
	date := models.GetTimeValue(duration)

	return us.recordRepository.GetLeaderboardInTimePeriod(date)
}

func (us UseCase) GetLeaderboardTimeForSubstance(durationString string, substance uint) ([]models.LeaderboardOccurrence, error) {
	duration := models.GetTimeDuration(durationString)
	if duration == "" {
		return nil, errors.New("duration not valid")
	}
	date := models.GetTimeValue(duration)

	return us.recordRepository.GetLeaderboardForSubstanceInTimePeriod(substance, date)
}
