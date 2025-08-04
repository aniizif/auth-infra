package services

import (
	"errors"

	"github.com/aniizif/stack-mate/auth-service/internal/models"
	"github.com/aniizif/stack-mate/auth-service/internal/repository"
	"github.com/aniizif/stack-mate/auth-service/pkg/hash"
	"github.com/aniizif/stack-mate/auth-service/pkg/jwt"
)

type AuthService struct {
	repo *repository.UserRepository
}

func NewAuthService(repo *repository.UserRepository) *AuthService {
	return &AuthService{repo}
}

func (s *AuthService) Login(email, password string) (string, error) {
	user, err := s.repo.GetByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if !hash.Compare(password, user.PasswordHash) {
		return "", errors.New("invalid credentials")
	}

	return jwt.GenerateToken(user.ID)
}

func (s *AuthService) Register(email, password string) (*models.User, string, error) {
	if user, _ := s.repo.GetByEmail(email); user != nil {
		return nil, "", errors.New("email already exists")
	}

	hashedPassword, err := hash.Password(password)
	if err != nil {
		return nil, "", err
	}

	user := &models.User{
		Email:        email,
		PasswordHash: hashedPassword,
	}

	if err := s.repo.CreateUser(user); err != nil {
		return nil, "", err
	}

	token, err := jwt.GenerateToken(user.ID)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}
