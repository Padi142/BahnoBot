package repository

import (
	"bahno_bot/domain"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type recordRepository struct {
	database   mongo.Database
	collection string
}

func (ur *recordRepository) Create(c context.Context, userId string, record domain.Record) error {
	collection := ur.database.Collection("users/" + userId + "/" + ur.collection)

	_, err := collection.InsertOne(c, record)

	return err
}

func (ur *recordRepository) Fetch(c context.Context, userId string, record domain.Record) ([]domain.Record, error) {
	collection := ur.database.Collection("users/" + userId + "/" + ur.collection)

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
	collection := ur.database.Collection("users/" + userId + "/" + ur.collection)

	opts := options.FindOne().SetSort(bson.D{{"createdAt", -1}})

	// Find the newest document
	var record domain.Record
	err := collection.FindOne(context.Background(), bson.D{}, opts).Decode(&record)
	if err != nil {
		log.Fatal(err)
	}

	return record, err
}
