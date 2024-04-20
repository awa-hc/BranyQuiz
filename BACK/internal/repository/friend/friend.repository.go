package friend

import (
	"brainyquiz/internal/domain/entities"
	"context"
)

type FriendRepository interface {
	AddFriend(ctx context.Context, friend *entities.Friend) error
	RemoveFriend(ctx context.Context, friend *entities.Friend) error
	GetFriends(ctx context.Context, username string) ([]entities.User, error)
}
