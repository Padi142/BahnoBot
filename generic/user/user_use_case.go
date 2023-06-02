package user

import (
	"context"
	"time"
	"bahno_bot/generic/models"
)

type UseCase struct {
	userRepository UserRepository
	contextTimeout time.Duration
}

func NewUserUseCase(userRepository UserRepository, timeout time.Duration) UseCase {
	return UseCase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}
func (useCase UseCase) CreateUser(c context.Context, user models.User) error {

	err := useCase.userRepository.Create(context.Background(), &user)
	if err != nil {
		return err
	}
	return nil
}

func (useCase UseCase) GetProfileByID(c context.Context, userID string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(c, useCase.contextTimeout)
	defer cancel()

	user, err := useCase.userRepository.GetUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, nil
	}

	return &models.User{Name: user.Name, ID: user.ID, UserId: user.UserId, PreferredSubstance: user.PreferredSubstance}, nil
}

func (useCase UseCase) GetOrCreateUser(c context.Context, userID string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(c, useCase.contextTimeout)
	defer cancel()

	user, err := useCase.userRepository.GetUser(ctx, userID)
	if err == nil {
		return &models.User{Name: user.Name, ID: user.ID, UserId: user.UserId, PreferredSubstance: user.PreferredSubstance}, nil
	}

	err = useCase.userRepository.Create(ctx, &models.User{UserId: userID})
	if err != nil {
		return nil, err
	}

	user, err = useCase.userRepository.GetUser(ctx, userID)
	if err != nil {
		return nil, err
	}
	return &models.User{Name: user.Name, ID: user.ID, UserId: user.UserId, PreferredSubstance: user.PreferredSubstance}, nil
}

func (useCase UseCase) SetPreferredSubstance(c context.Context, userId, newSubstance string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(c, useCase.contextTimeout)
	defer cancel()

	err := useCase.userRepository.SetPreferredSubstance(context.Background(), userId, newSubstance)

	if err != nil {
		return nil, err
	}

	user, err := useCase.userRepository.GetUser(ctx, userId)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, nil
	}

	return &models.User{Name: user.Name, ID: user.ID, UserId: user.UserId, PreferredSubstance: user.PreferredSubstance}, nil
}
