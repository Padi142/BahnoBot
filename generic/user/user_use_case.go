package user

import (
	"bahno_bot/generic/models"
)

type UseCase struct {
	userRepository UserRepository
}

func NewUserUseCase(userRepository UserRepository) UseCase {
	return UseCase{
		userRepository: userRepository,
	}
}
func (useCase UseCase) CreateUser(user models.User) error {

	err := useCase.userRepository.Create(&user)
	if err != nil {
		return err
	}
	return nil
}

func (useCase UseCase) GetProfileByID(userID uint) (*models.User, error) {
	user, err := useCase.userRepository.GetUser(userID)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, nil
	}

	return user, nil
}

func (useCase UseCase) GetProfileByDiscordID(userID string) (*models.User, error) {
	user, err := useCase.userRepository.GetUserByDiscordId(userID)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, nil
	}

	return user, nil
}

func (useCase UseCase) GetOrCreateUser(userID uint) (*models.User, error) {
	user, err := useCase.userRepository.GetUser(userID)
	if err == nil {
		return user, nil
	}

	err = useCase.userRepository.Create(&models.User{ID: userID})
	if err != nil {
		return nil, err
	}

	user, err = useCase.userRepository.GetUser(userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (useCase UseCase) GetOrCreateDiscordUser(discordId string) (*models.User, error) {
	user, err := useCase.userRepository.GetUserByDiscordId(discordId)
	if err == nil {
		return user, nil
	}

	err = useCase.userRepository.Create(&models.User{DiscordID: discordId})
	if err != nil {
		return nil, err
	}

	user, err = useCase.userRepository.GetUserByDiscordId(discordId)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (useCase UseCase) SetPreferredSubstance(userId, substanceId uint) (*models.User, error) {
	err := useCase.userRepository.SetPreferredSubstance(userId, substanceId)

	if err != nil {
		return nil, err
	}

	user, err := useCase.userRepository.GetUser(userId)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, nil
	}

	return user, nil
}
