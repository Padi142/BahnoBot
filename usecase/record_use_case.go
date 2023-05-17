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

func CreateRecordUseCase(recordRepository domain.RecordRepository, record domain.Record, userId string) error {
	err := recordRepository.Create(context.Background(), userId, record)
	if err != nil {
		return err
	}
	return nil
}
