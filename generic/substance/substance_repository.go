package substance

import (
	"bahno_bot/generic/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type substanceRepository struct {
	database   mongo.Database
	collection string
}

func NewSubstanceRepository(db mongo.Database, collection string) SubstanceRepository {
	return &substanceRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *substanceRepository) Fetch(c context.Context) ([]models.Substance, error) {
	collection := ur.database.Collection(ur.collection)

	cursor, err := collection.Find(c, bson.D{})

	if err != nil {
		return nil, err
	}

	var substances []models.Substance

	err = cursor.All(c, &substances)
	if substances == nil {
		return []models.Substance{}, err
	}

	return substances, err
}
