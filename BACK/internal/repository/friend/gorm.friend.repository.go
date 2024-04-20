package friend

import (
	"brainyquiz/internal/domain/entities"
	"context"

	"gorm.io/gorm"
)

type gormFriendRepository struct {
	db *gorm.DB
}

func NewFriendRepository(db *gorm.DB) *gormFriendRepository {
	return &gormFriendRepository{db: db}
}

func (r *gormFriendRepository) AddFriend(ctx context.Context, friend *entities.Friend) error {

	if err := friend.Validate(); err != nil {
		return err
	}

	return r.db.Create(friend).Error

}

func (r *gormFriendRepository) RemoveFriend(ctx context.Context, friend *entities.Friend) error {

	var f entities.Friend
	if err := r.db.Where("user_id = ? AND friend_id = ?", friend.UserID, friend.FriendID).First(&f).Error; err != nil {
		return err
	}
	return r.db.Delete(&f).Error
}

func (r *gormFriendRepository) GetFriends(ctx context.Context, username string) ([]entities.User, error) {

	var friends []entities.User
	err := r.db.Table("users").
		Select("users.id, users.user_name, users.email").
		Joins("JOIN friends ON users.id = friends.friend_id").
		Joins("JOIN users u ON u.id = friends.user_id").
		Where("u.user_name = ?", username).
		Scan(&friends).Error
	if err != nil {
		return nil, err
	}
	return friends, nil

}
