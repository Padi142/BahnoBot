package domain

import (
	"bahno_bot/models"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Record struct {
	ID        primitive.ObjectID `bson:"_id"`
	Substance models.Substance   `bson:"substance"`
	Time      time.Time          `bson:"time"`
	CreatedAt time.Time          `bson:"createdAt"`
}

type RecordRepository interface {
	Create(c context.Context, userId string, record Record) error
	Fetch(c context.Context, userId string) ([]Record, error)
	GetLastRecord(c context.Context, userId string) (Record, error)
}
