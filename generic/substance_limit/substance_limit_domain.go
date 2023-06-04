package substance_limit

import (
	"bahno_bot/generic/models"
)

type SubstanceLimitRepository interface {
	Get(userId, substanceId uint) (models.SubstanceLimit, error)
	Create(models.SubstanceLimit) error
}
