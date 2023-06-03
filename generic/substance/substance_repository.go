package substance

import (
	"bahno_bot/generic/models"
	"gorm.io/gorm"
)

type substanceRepository struct {
	database *gorm.DB 
}

func NewSubstanceRepository(db *gorm.DB) SubstanceRepository {
	return &substanceRepository{
		database:   db,
	}
}

func (ur *substanceRepository) GetAll() (substances []models.Substance, err error) {
	ur.database.Find(&substances)

	return
}

func (ur *substanceRepository) Create(substance models.Substance) error {
	return ur.database.Create(&substance).Error
}
