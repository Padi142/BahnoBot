package record

import (
	"context"
	"bahno_bot/generic/models"
)



type RecordRepository interface {
	Create(c context.Context, userId string, record models.Record) error
	GetAll(c context.Context, userId string) ([]models.Record, error)
	GetLast(c context.Context, userId string) (models.Record, error)
}
