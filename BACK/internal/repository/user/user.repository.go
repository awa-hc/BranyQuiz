package user

import (
	"brainyquiz/internal/domain/entities"
	"context"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *entities.User) error
}
