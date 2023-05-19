package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID                 primitive.ObjectID `bson:"_id"`
	Name               string             `bson:"name"`
	PreferredSubstance string             `bson:"preferred_substance"`
	UserId             string             `bson:"user_id"`
}

type UserRepository interface {
	Create(c context.Context, user *User) error
	Fetch(c context.Context) ([]User, error)
	GetByUserID(c context.Context, id string) (*User, error)
	SetPreferredSubstance(c context.Context, id, newSubstance string) error
}
