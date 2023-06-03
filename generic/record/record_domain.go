package record

import (
	"bahno_bot/generic/models"
)

type RecordRepository interface {
	Create(record models.Record) error
	GetAll() ([]models.Record, error)
}
