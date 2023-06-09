package user

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userRepository struct {
	database   mongo.Database
	collection string
}

func NewUserRepository(db mongo.Database, collection string) UserRepository {
	return &userRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *userRepository) Create(c context.Context, user *User) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, user)

	return err
}

func (ur *userRepository) Fetch(c context.Context) ([]User, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var users []User

	err = cursor.All(c, &users)
	if users == nil {
		return []User{}, err
	}

	return users, err
}

func (ur *userRepository) GetByUserID(c context.Context, id string) (*User, error) {
	collection := ur.database.Collection(ur.collection)

	var user User

	err := collection.FindOne(c, bson.M{"user_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (ur *userRepository) SetPreferredSubstance(c context.Context, userId, newSubstance string) error {
	collection := ur.database.Collection(ur.collection)

	res, err := collection.UpdateOne(c, bson.M{"user_id": userId}, bson.M{"$set": bson.M{"preferred_substance": newSubstance}})
	if err != nil {
		return err
	}
	if res.ModifiedCount != 1 {
		return err
	}
	return err
}
