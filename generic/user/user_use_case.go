package user

import (
	"bahno_bot/generic/models"
	"bahno_bot/generic/record"
)

type UseCase struct {
	userRepository UserRepository;
	recordRepostiroy record.RecordRepository;
}

func NewUserUseCase(userRepository UserRepository, recordRepository record.RecordRepository) UseCase {
	return UseCase{
		userRepository: userRepository,
		recordRepostiroy: recordRepository,
	}
}

func (useCase UseCase) GetUsers() ([]models.User, error) {
	return useCase.userRepository.GetAll()
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

func (useCase UseCase) GetUserRecords(userId uint) ([]models.Record, error) {
	return useCase.recordRepostiroy.GetAll(userId)
}

func (useCase UseCase) GetLastUserRecord(userId uint) (models.Record, error) {
	return useCase.recordRepostiroy.GetLast(userId)
}