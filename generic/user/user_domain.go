package user

import (
	"context"
	"bahno_bot/generic/models"
)


type UserRepository interface {
	Create(c context.Context, user *models.User) error
	Fetch(c context.Context) ([]models.User, error)
	GetByUserID(c context.Context, id string) (*models.User, error)
	SetPreferredSubstance(c context.Context, id, newSubstance string) error
}
