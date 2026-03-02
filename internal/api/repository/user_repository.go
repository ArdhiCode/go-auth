package repository

import (
	"context"

	"github.com/ArdhiCode/go-auth/internal/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, tx *gorm.DB, user entity.User) (entity.User, error)
	Get(ctx context.Context, tx *gorm.DB, user entity.User) (*entity.User, error)
	Update(ctx context.Context, tx *gorm.DB, user entity.User) (entity.User, error)
	Delete(ctx context.Context, tx *gorm.DB, user entity.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(ctx context.Context, tx *gorm.DB, user entity.User) (entity.User, error) {
	db := r.db
	if tx != nil {
		db = tx
	}

	if err := db.WithContext(ctx).Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) Get(ctx context.Context, tx *gorm.DB, user entity.User) (*entity.User, error) {
	db := r.db
	if tx != nil {
		db = tx
	}

	var result entity.User
	if err := db.WithContext(ctx).First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *userRepository) Update(ctx context.Context, tx *gorm.DB, user entity.User) (entity.User, error) {
	db := r.db
	if tx != nil {
		db = tx
	}

	if err := db.WithContext(ctx).Save(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) Delete(ctx context.Context, tx *gorm.DB, user entity.User) error {
	db := r.db
	if tx != nil {
		db = tx
	}

	if err := db.WithContext(ctx).Delete(&user).Error; err != nil {
		return err
	}

	return nil
}
