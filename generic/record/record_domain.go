package record

import (
	"bahno_bot/generic/models"
)



type RecordRepository interface {
	Create(record models.Record) error
	GetAll(userId uint) ([]models.Record, error)
	GetLast(userId uint) (models.Record, error)
}
