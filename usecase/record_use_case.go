package usecase

import (
	"bahno_bot/domain"
	"context"
	"time"
)

type recordUseCase struct {
	recordRepository domain.RecordRepository
	contextTimeout   time.Duration
}

func NewRecordUseCase(recordRepository domain.RecordRepository, timeout time.Duration) recordUseCase {
	return recordUseCase{
		recordRepository: recordRepository,
		contextTimeout:   timeout,
	}
}

func CreateRecordUseCase(recordRepository domain.RecordRepository, record domain.Record, userId string) error {
	err := recordRepository.Create(context.Background(), userId, record)
	if err != nil {
		return err
	}
	return nil
}
func (useCase recordUseCase) GetLatestRecord(c context.Context, userId string) (*domain.Record, error) {

	record, err := useCase.recordRepository.GetLastRecord(c, userId)
	if err != nil {
		return nil, err
	}

	return &record, nil
}

func (useCase recordUseCase) CreateNewRecord(c context.Context, userId string, record domain.Record) (*domain.Record, error) {

	err := useCase.recordRepository.Create(c, userId, record)
	if err != nil {
		return nil, err
	}

	newRecord, err := useCase.recordRepository.GetLastRecord(c, userId)
	if err != nil {
		return nil, err
	}

	return &newRecord, nil
}
