package user

import (
	"bahno_bot/generic/models"
	"gorm.io/gorm"
)

type userRepository struct {
	database *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		database: db,
	}
}

func (ur *userRepository) Create(user *models.User) error {
	result := ur.database.Create(user)

	return result.Error
}

func (ur *userRepository) GetAll() (users []models.User, err error) {
	ur.database.Preload("PreferredSubstance").Find(&users)

	return
}

func (ur *userRepository) GetUser(id uint) (user *models.User, err error) {
	ur.database.Preload("PreferredSubstance").First(&user, id)

	return
}
func (ur *userRepository) GetUserByDiscordId(id string) (user *models.User, err error) {
	err = ur.database.Preload("PreferredSubstance").Where("discord_id = ?", id).First(&user).Error

	return
}

func (ur *userRepository) SetPreferredSubstance(userId, substanceId uint) error {
	ur.database.Model(&models.User{}).Where("id = ?", userId).Update("preferred_substance_id", substanceId)

	return nil
}
