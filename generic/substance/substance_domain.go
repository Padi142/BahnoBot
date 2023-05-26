package substance

import (
	"bahno_bot/generic/models"
	"context"
)

type SubstanceRepository interface {
	Fetch(c context.Context) ([]models.Substance, error)
}
