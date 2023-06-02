package record

import (
	"context"
	"bahno_bot/generic/models"
)



type RecordRepository interface {
	Create(c context.Context, userId string, record models.Record) error
	Fetch(c context.Context, userId string) ([]models.Record, error)
	GetLastRecord(c context.Context, userId string) (models.Record, error)
}
