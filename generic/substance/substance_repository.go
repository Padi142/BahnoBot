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
		database: db,
	}
}

func (ur *substanceRepository) GetAll() (substances []models.Substance, err error) {
	err = ur.database.Find(&substances).Error
	return
}

func (ur *substanceRepository) Get(id uint) (substance models.Substance, err error) {
	err = ur.database.First(&substance, id).Error
	return
}

func (ur *substanceRepository) GetByValue(value string) (substance models.Substance, err error) {
	err = ur.database.Where("value = ?", value).First(&substance).Error
	return
}

func (ur *substanceRepository) Create(substance models.Substance) error {
	return ur.database.Create(&substance).Error
}
