package user

import (
	"context"
	"bahno_bot/generic/models"
)


type UserRepository interface {
	Create(c context.Context, user *models.User) error
	GetAll(c context.Context) ([]models.User, error)
	GetUser(c context.Context, id string) (*models.User, error)
	SetPreferredSubstance(c context.Context, userId, substanceId uint) error
}
