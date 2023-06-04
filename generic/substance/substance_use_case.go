package substance

import (
	"bahno_bot/generic/models"
	"gorm.io/gorm"
)

type UseCase struct {
	substanceRepository SubstanceRepository
}

func NewSubstanceUseCase(db *gorm.DB) UseCase {
	return UseCase{
		substanceRepository: NewSubstanceRepository(db),
	}
}
func (useCase UseCase) GetSubstances() ([]models.Substance, error) {
	substances, err := useCase.substanceRepository.GetAll()

	if err != nil {
		return nil, err
	}

	return substances, nil
}
