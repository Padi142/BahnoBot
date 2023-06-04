package record

import (
	"bahno_bot/generic/models"
)

type RecordRepository interface {
	Create(record models.Record) error
	GetAll(userId uint) (records []models.Record, err error)
	GetLast(userId uint) (records models.Record, err error)
}
