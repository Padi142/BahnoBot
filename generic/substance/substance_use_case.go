package substance

import (
	"bahno_bot/generic/models"
)

type UseCase struct {
	substanceRepository SubstanceRepository
}

func NewSubstanceUseCase(substanceRepository SubstanceRepository) UseCase {
	return UseCase{
		substanceRepository: substanceRepository,
	}
}
func (useCase UseCase) GetSubstances() ([]models.Substance, error) {
	substances, err := useCase.substanceRepository.GetAll()

	if err != nil {
		return nil, err
	}

	return substances, nil
}

func (useCase UseCase) GetSubstance(id uint) (models.Substance, error) {
	return useCase.substanceRepository.Get(id)
}
