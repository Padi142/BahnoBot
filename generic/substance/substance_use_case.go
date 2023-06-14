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
	return useCase.substanceRepository.GetAll()
}

func (useCase UseCase) GetSubstance(id uint) (models.Substance, error) {
	return useCase.substanceRepository.Get(id)
}

func (useCase UseCase) GetSubstanceByValue(value string) (models.Substance, error) {
	return useCase.substanceRepository.GetByValue(value)
}
