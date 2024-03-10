package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/qPyth/mobydev-internship-admin/internal/domain"
	domain2 "github.com/qPyth/mobydev-internship-admin/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

var adminRole = "admin"

type UserStorage interface {
	GetUser(ctx context.Context, email string) (domain2.User, error)
	CreateUser(ctx context.Context, email string, passHash []byte, role string) error
}

type TokenManager interface {
	NewToken(userID uint, role string) (string, error)
}

type UserService struct {
	tokenManager TokenManager
	userStorage  UserStorage
}

func NewUser(manager TokenManager, storage UserStorage) *UserService {
	return &UserService{
		tokenManager: manager,
		userStorage:  storage,
	}
}

func (s *UserService) RegisterAdmin(ctx context.Context, email, password string) error {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	if err = s.userStorage.CreateUser(ctx, email, hashPass, adminRole); err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

func (s *UserService) LoginAdmin(ctx context.Context, email, password string) (string, error) {
	user, err := s.userStorage.GetUser(ctx, email)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			return "", domain2.ErrInvalidCredentials
		}
		return "", fmt.Errorf("failed to get user: %w", err)
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		if errors.Is(bcrypt.ErrMismatchedHashAndPassword, err) {
			return "", domain2.ErrInvalidCredentials
		}
		return "", fmt.Errorf("failed to compare password: %w", err)
	}

	token, err := s.tokenManager.NewToken(uint(user.ID), user.Role)
	if err != nil {
		return "", fmt.Errorf("failed to create token: %w", err)
	}

	return token, nil
}
