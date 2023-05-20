package usecase

import (
	"bahno_bot/domain"
	"bahno_bot/models"
	"context"
	"time"
)

type substanceUseCase struct {
	substanceRepository domain.SubstanceRepository
	contextTimeout      time.Duration
}

func NewSubstanceUseCase(substanceRepository domain.SubstanceRepository, timeout time.Duration) substanceUseCase {
	return substanceUseCase{
		substanceRepository: substanceRepository,
		contextTimeout:      timeout,
	}
}
func (useCase substanceUseCase) GetSubstances(c context.Context) ([]models.Substance, error) {

	substances, err := useCase.substanceRepository.Fetch(c)
	if err != nil {
		return nil, err
	}
	return substances, nil
}
