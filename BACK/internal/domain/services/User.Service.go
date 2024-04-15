package services

import (
	"brainyquiz/internal/domain/entities"
	"brainyquiz/internal/repository/user"
	"context"
)

type UserService interface {
	CreateUser(ctx context.Context, user *entities.User) error
}

type userService struct {
	UserRepository user.UserRepository
}

func NewUserService(userRepository user.UserRepository) *userService {
	return &userService{UserRepository: userRepository}
}

func (s *userService) CreateUser(ctx context.Context, user *entities.User) error {
	if err := user.Validate(); err != nil {
		return err
	}
	if err := s.UserRepository.CreateUser(ctx, user); err != nil {
		return err
	}

	return nil
}
