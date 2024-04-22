package services

import (
	"brainyquiz/internal/domain/entities"
	"brainyquiz/internal/repository/friend"
	"context"
)

type FriendService struct {
	FriendRepository friend.FriendRepository
}

func NewFriendService(friendRepository friend.FriendRepository) *FriendService {
	return &FriendService{FriendRepository: friendRepository}
}

func (s *FriendService) AddFriend(ctx context.Context, request entities.Friend) error {
	if err := request.Validate(); err != nil {
		return err
	}
	return s.FriendRepository.AddFriend(ctx, &request)
}

func (s *FriendService) RemoveFriend(ctx context.Context, request entities.Friend) error {
	return s.FriendRepository.RemoveFriend(ctx, &request)
}

func (s *FriendService) GetFriends(ctx context.Context, username string) ([]entities.User, error) {
	return s.FriendRepository.GetFriends(ctx, username)
}
