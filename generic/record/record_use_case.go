package record

import (
	"bahno_bot/generic/models"
)

type UseCase struct {
	recordRepository RecordRepository
}

func NewRecordUseCase(recordRepository RecordRepository) UseCase {
	return UseCase{
		recordRepository: recordRepository,
	}
}

func CreateRecordUseCase(recordRepository RecordRepository, record models.Record) error {
	err := recordRepository.Create(record)

	if err != nil {
		return err
	}

	return nil
}

func (useCase UseCase) CreateNewRecord(userId uint, record models.Record) (*models.Record, error) {
	err := useCase.recordRepository.Create(record)

	if err != nil {
		return nil, err
	}

	return &record, nil
}

func (useCase UseCase) GetAllRecords() ([]models.Record, error) {
	return useCase.GetAllRecords()
}
