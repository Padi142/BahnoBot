package repository

import (
	"bahno_bot/domain"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"sort"
)

type recordRepository struct {
	database   mongo.Database
	collection string
}

func NewRecordRepository(db mongo.Database, collection string) domain.RecordRepository {
	return &recordRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *recordRepository) Create(c context.Context, userId string, record domain.Record) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"user_id": userId}

	update := bson.M{
		"$push": bson.M{"records": record},
	}

	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	return err
}

func (ur *recordRepository) Fetch(c context.Context, userId string) ([]domain.Record, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "userId", Value: userId}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var records []domain.Record

	err = cursor.All(c, &records)
	if records == nil {
		return []domain.Record{}, err
	}

	return records, err
}

func (ur *recordRepository) GetLastRecord(c context.Context, userId string) (domain.Record, error) {
	collection := ur.database.Collection(ur.collection)

	type UserArray struct {
		Records []*domain.Record `bson:"records"`
	}
	var userArray UserArray
	var record domain.Record

	err := collection.FindOne(c, bson.M{"user_id": userId}).Decode(&userArray)
	if err != nil {
		return domain.Record{}, err
	}

	sort.Slice(userArray.Records, func(i, j int) bool {
		return userArray.Records[i].Time.Before(userArray.Records[j].Time)
	})

	record = *userArray.Records[0]

	return record, err
}
