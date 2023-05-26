package substance

import (
	"bahno_bot/generic/models"
	"context"
	"time"
)

type UseCase struct {
	substanceRepository SubstanceRepository
	contextTimeout      time.Duration
}

func NewSubstanceUseCase(substanceRepository SubstanceRepository, timeout time.Duration) UseCase {
	return UseCase{
		substanceRepository: substanceRepository,
		contextTimeout:      timeout,
	}
}
func (useCase UseCase) GetSubstances(c context.Context) ([]models.Substance, error) {

	substances, err := useCase.substanceRepository.Fetch(c)
	if err != nil {
		return nil, err
	}
	return substances, nil
}
