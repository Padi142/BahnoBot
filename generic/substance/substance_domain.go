package substance

import (
	"bahno_bot/generic/models"
)

type SubstanceRepository interface {
	GetAll() ([]models.Substance, error)
	Get(id uint) (models.Substance, error)
	Create(models.Substance) error
}
