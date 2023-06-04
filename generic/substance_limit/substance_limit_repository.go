package substance_limit

import (
	"bahno_bot/generic/models"

	"gorm.io/gorm"
)

type substanceLimitRepository struct {
	database *gorm.DB
}

func NewSubstanceRepository(db *gorm.DB) SubstanceLimitRepository {
	return &substanceLimitRepository{
		database: db,
	}
}

func (ur *substanceLimitRepository) Get(userId, substanceId uint) (substanceLimit models.SubstanceLimit, err error) {
	err = ur.database.Where("user_id = ? AND substance_id = ?", userId, substanceId).First(&substanceLimit).Error
	return
}

func (ur *substanceLimitRepository) Create(substance models.SubstanceLimit) error {
	return ur.database.Create(&substance).Error
}
