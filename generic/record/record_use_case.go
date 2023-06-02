package record

import (
	"context"
	"time"
	"bahno_bot/generic/models"
)

type UseCase struct {
	recordRepository RecordRepository
	contextTimeout   time.Duration
}

func NewRecordUseCase(recordRepository RecordRepository, timeout time.Duration) UseCase {
	return UseCase{
		recordRepository: recordRepository,
		contextTimeout:   timeout,
	}
}

func CreateRecordUseCase(recordRepository RecordRepository, record models.Record, userId string) error {
	err := recordRepository.Create(context.Background(), userId, record)

	if err != nil {
		return err
	}
	return nil
}
func (useCase UseCase) GetLatestRecord(c context.Context, userId string) (*models.Record, error) {

	record, err := useCase.recordRepository.GetLast(c, userId)

	if err != nil {
		return nil, err
	}

	return &record, nil
}

func (useCase UseCase) CreateNewRecord(c context.Context, userId string, record models.Record) (*models.Record, error) {
	err := useCase.recordRepository.Create(c, userId, record)
	
	if err != nil {
		return nil, err
	}

	newRecord, err := useCase.recordRepository.GetLast(c, userId)

	if err != nil {
		return nil, err
	}

	return &newRecord, nil
}
