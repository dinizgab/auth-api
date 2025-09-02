package users

import (
	"auth-api/internal/auth"
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type Usecase interface {
	CreateUser(ctx context.Context, user User) error
	Login(ctx context.Context, email string, password string) (auth.TokenPair, error)
}

type usersUsecaseImpl struct {
	repo        Repository
	authService auth.Service
}

func NewUsecase(repo Repository, authService auth.Service) Usecase {
	return &usersUsecaseImpl{
		repo:        repo,
		authService: authService,
	}
}

func (u *usersUsecaseImpl) CreateUser(ctx context.Context, user User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return fmt.Errorf("UsersUsecase.CreateUser: %w", err)
	}

	user.Password = string(hashedPassword)

	return u.repo.CreateUser(ctx, user)
}

func (u *usersUsecaseImpl) Login(ctx context.Context, email string, password string) (auth.TokenPair, error) {
	user, err := u.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return auth.TokenPair{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return auth.TokenPair{}, fmt.Errorf("UsersUsecase.Login: %w", err)
	}

	tokenPair, err := u.authService.GenerateTokenPair(user.ID)
	if err != nil {
		return auth.TokenPair{}, nil
	}

	return tokenPair, nil
}
