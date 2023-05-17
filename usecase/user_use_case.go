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

func NewUserUseCase(userRepository domain.UserRepository, user domain.User) error {

	err := userRepository.Create(context.Background(), &user)
	if err != nil {
		return err
	}
	return nil
}

func (useCase *userUseCase) GetProfileByID(c context.Context, userID string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(c, useCase.contextTimeout)
	defer cancel()

	user, err := useCase.userRepository.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &domain.User{Name: user.Name, ID: user.ID, UserId: user.UserId, PreferredSubstance: user.PreferredSubstance}, nil
}

func (useCase *userUseCase) GetOrCreateUser(c context.Context, userID string) (*domain.User, error) {
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
