package user

import (
	"brainyquiz/internal/domain/entities"
	"context"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type gormUserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *gormUserRepository {
	return &gormUserRepository{db: db}
}

func (r *gormUserRepository) CreateUser(ctx context.Context, user *entities.User) error {
	if err := user.Validate(); err != nil {
		return err
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return r.db.Create(user).Error
}

func (r *gormUserRepository) GetUserByEmail(ctx context.Context, email string) (*entities.User, error) {
	var user entities.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *gormUserRepository) GetUserByUserName(ctx context.Context, userName string) (*entities.User, error) {
	var user entities.User
	err := r.db.Where("user_name = ?", userName).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
