package services

import (
	"brainyquiz/internal/domain/entities"
	"brainyquiz/internal/repository/user"
	"context"
)

type UserService struct {
	UserRepository user.UserRepository
}

func NewUserService(userRepository user.UserRepository) *UserService {
	return &UserService{UserRepository: userRepository}
}

func (s *UserService) CreateUser(ctx context.Context, user *entities.User) error {
	if err := user.Validate(); err != nil {
		return err
	}
	if err := s.UserRepository.CreateUser(ctx, user); err != nil {
		return err
	}

	return nil
}

func (s *UserService) GetUserByEmail(ctx context.Context, email string) (*entities.User, error) {
	user, err := s.UserRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) GetUserByUserName(ctx context.Context, userName string) (*entities.User, error) {
	user, err := s.UserRepository.GetUserByUserName(ctx, userName)
	if err != nil {
		return nil, err
	}
	return user, nil
}
