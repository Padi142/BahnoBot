package substance

import (
	"bahno_bot/generic/models"
	"context"
)

type SubstanceRepository interface {
	GetAll(c context.Context) ([]models.Substance, error)
}
