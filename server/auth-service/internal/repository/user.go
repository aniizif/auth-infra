package repository

import (
	"gorm.io/gorm"

	"github.com/aniizif/stack-mate/auth-service/internal/models"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	return r.db.Create(&user).Error
}

func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	result := r.db.Where(&models.User{Email: email}).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
