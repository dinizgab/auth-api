package users

import (
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type Usecase interface {
	CreateUser(ctx context.Context, user User) error
}

type usersUsecaseImpl struct {
	repo Repository
}

func NewUsecase(repo Repository) Usecase {
	return &usersUsecaseImpl{
		repo: repo,
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
