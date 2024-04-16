package user

import (
	"brainyquiz/internal/domain/entities"
	"context"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *entities.User) error
	GetUserByEmail(ctx context.Context, email string) (*entities.User, error)
	GetUserByUserName(ctx context.Context, userName string) (*entities.User, error)
}
