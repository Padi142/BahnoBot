package record

import (
	"bahno_bot/generic/models"
)

type RecordRepository interface {
	Create(record models.Record) error
	GetAll(userId uint) (records []models.Record, err error)
	GetAllPaged(userId uint, page, pageSize int) (records []models.Record, count int64, err error)
	GetLast(userId uint) (records models.Record, err error)
}
