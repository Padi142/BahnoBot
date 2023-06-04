package substance_limit

import (
	"bahno_bot/generic/models"
)

type UseCase struct {
	substanceLimitRepository SubstanceLimitRepository
}

func NewSubstanceUseCase(substanceRepository SubstanceLimitRepository) UseCase {
	return UseCase{
		substanceLimitRepository: substanceRepository,
	}
}

func (useCase UseCase) GetSubstanceLimit(userId, substanceId uint) (models.SubstanceLimit, error) {
	return useCase.substanceLimitRepository.Get(userId, substanceId)
}

func (useCase UseCase) CreateSubstanceLimit(substanceLimit models.SubstanceLimit) error {
	return useCase.substanceLimitRepository.Create(substanceLimit)
}
