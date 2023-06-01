package user

import (
	"bahno_bot/generic/models"
	"context"
)

type UserRepository interface {
	Create(user *models.User) error
	GetAll() ([]models.User, error)
	GetUser(id uint) (*models.User, error)
	SetPreferredSubstance(userId, substanceId uint) error
}
