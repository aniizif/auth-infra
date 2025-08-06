package repository

import (
	"github.com/aniizif/stack-mate/auth-service/internal/metrics"
	"gorm.io/gorm"
	"time"

	"github.com/aniizif/stack-mate/auth-service/internal/models"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	start := time.Now()
	err := r.db.Create(user).Error
	duration := time.Since(start).Seconds()

	metrics.DBQueryDuration.WithLabelValues("INSERT", "user").Observe(duration)

	return err
}

func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	start := time.Now()

	var user models.User
	result := r.db.Where(&models.User{Email: email}).First(&user)

	duration := time.Since(start).Seconds()
	metrics.DBQueryDuration.WithLabelValues("SELECT", "user").Observe(duration)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
