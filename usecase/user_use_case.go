package usecase

import (
	"bahno_bot/domain"
	"context"
	"time"
)

type userUseCase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewUserUseCase(userRepository domain.UserRepository, timeout time.Duration) userUseCase {
	return userUseCase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}
func (useCase userUseCase) CreateUser(c context.Context, user domain.User) error {

	err := useCase.userRepository.Create(context.Background(), &user)
	if err != nil {
		return err
	}
	return nil
}

func (useCase userUseCase) GetProfileByID(c context.Context, userID string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(c, useCase.contextTimeout)
	defer cancel()

	user, err := useCase.userRepository.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, nil
	}

	return &domain.User{Name: user.Name, ID: user.ID, UserId: user.UserId, PreferredSubstance: user.PreferredSubstance}, nil
}

func (useCase userUseCase) GetOrCreateUser(c context.Context, userID string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(c, useCase.contextTimeout)
	defer cancel()

	user, err := useCase.userRepository.GetByUserID(ctx, userID)
	if err == nil {
		return &domain.User{Name: user.Name, ID: user.ID, UserId: user.UserId, PreferredSubstance: user.PreferredSubstance}, nil
	}

	err = useCase.userRepository.Create(ctx, &domain.User{UserId: userID})
	if err != nil {
		return nil, err
	}

	user, err = useCase.userRepository.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return &domain.User{Name: user.Name, ID: user.ID, UserId: user.UserId, PreferredSubstance: user.PreferredSubstance}, nil
}

func (useCase userUseCase) SetPreferredSubstance(c context.Context, userId, newSubstance string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(c, useCase.contextTimeout)
	defer cancel()

	err := useCase.userRepository.SetPreferredSubstance(context.Background(), userId, newSubstance)

	if err != nil {
		return nil, err
	}

	user, err := useCase.userRepository.GetByUserID(ctx, userId)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, nil
	}

	return &domain.User{Name: user.Name, ID: user.ID, UserId: user.UserId, PreferredSubstance: user.PreferredSubstance}, nil
}
