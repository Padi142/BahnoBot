package domain

import (
	"bahno_bot/models"
	"context"
)

type SubstanceRepository interface {
	Fetch(c context.Context) ([]models.Substance, error)
}
