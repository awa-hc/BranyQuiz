package services

import (
	"brainyquiz/internal/repository/auth"
	"context"
)

type AuthService struct {
	AuthRepository auth.AuthRepository
}

func NewAuthService(authRepository auth.AuthRepository) *AuthService {
	return &AuthService{AuthRepository: authRepository}
}

func (s *AuthService) LoginWithEmail(ctx context.Context, email string, password string) (string, error) {
	token, err := s.AuthRepository.LoginWithEmail(ctx, email, password)
	if err != nil {
		return "", err
	}
	return token, nil
}
