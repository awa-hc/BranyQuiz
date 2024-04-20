package entities

import "gorm.io/gorm"

type Friend struct {
	gorm.Model
	UserID   uint
	FriendID uint
}

func (f *Friend) Validate() error {
	if f.UserID == 0 {
		return &ValidationError{Field: "UserID", Message: "UserID is required"}
	}
	if f.FriendID == 0 {
		return &ValidationError{Field: "FriendID", Message: "FriendID is required"}
	}
	return nil
}
