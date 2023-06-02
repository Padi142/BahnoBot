package user

import (
	"context"

	"bahno_bot/generic/models"
	"gorm.io/gorm"
)

type userRepository struct {
	database   gorm.DB
}

func NewUserRepository(db gorm.DB) UserRepository {
	return &userRepository {
		database: db,
	}
}

func (ur *userRepository) Create(c context.Context, user *models.User) error {
	result := ur.database.Create(user);

	return result.Error
}

func (ur *userRepository) Fetch(c context.Context) (users []models.User, err error) {
	ur.database.Find(&users)

	return
}

func (ur *userRepository) GetByUserID(c context.Context, id string) (user *models.User, err error) {
	ur.database.First(user, id)

	return 

	
}
func (ur *userRepository) SetPreferredSubstance(c context.Context, userId, newSubstance uint) error {
	// ur.database.Model(&)
	return nil
}
