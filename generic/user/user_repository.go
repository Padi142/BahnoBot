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
	return ur.database.Create(user).Error
}

func (ur *userRepository) GetAll() (users []models.User, err error) {
	err = ur.database.Preload("PreferredSubstance").Find(&users).Error

	return
}

func (ur *userRepository) GetUser(id uint) (user *models.User, err error) {
	err = ur.database.Preload("PreferredSubstance").First(&user, id).Error

	return
}
func (ur *userRepository) GetUserByDiscordId(id string) (user *models.User, err error) {
	err = ur.database.Preload("PreferredSubstance").Where("discord_id = ?", id).First(&user).Error

	return
}

func (ur *userRepository) SetPreferredSubstance(userId, substanceId uint) error {
	return ur.database.Model(&models.User{}).Where("id = ?", userId).Update("preferred_substance_id", substanceId).Error
}

func (ur *userRepository) GetUserRecords(userId uint) (records []models.Record, err error) {
	err = ur.database.Where("user_id = ?", userId).Find(&records).Error

	return
}

func (ur *userRepository) GetUserLastRecord(userId uint) (record *models.Record, err error) {
	err = ur.database.Where("user_id = ?", userId).Last(&record).Error

	return
}
