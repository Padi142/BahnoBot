package substance

import (
	"bahno_bot/generic/models"
	"context"

	"gorm.io/gorm"
)

type substanceRepository struct {
	database  gorm.DB 
}

func NewSubstanceRepository(db gorm.DB) SubstanceRepository {
	return &substanceRepository{
		database:   db,
	}
}

func (ur *substanceRepository) GetAll(c context.Context) (substances []models.Substance, err error) {
	ur.database.Find(&substances)

	return
}
