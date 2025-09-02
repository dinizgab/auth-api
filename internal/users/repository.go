package users

import (
	"auth-api/internal/database"
	"context"
	"fmt"

	_ "embed"
)

type Repository interface {
	CreateUser(ctx context.Context, user User) error
}

type usersRepositoryImpl struct {
	db database.Database
}

var (
	//go:embed sql/create_user.sql
	CreateUserQuery string
)

func NewRepository(db database.Database) Repository {
	return &usersRepositoryImpl{
		db: db,
	}
}

func (r *usersRepositoryImpl) CreateUser(ctx context.Context, user User) error {
	_, err := r.db.Exec(ctx, CreateUserQuery, user.Username, user.Email, user.Password)
	if err != nil {
		return fmt.Errorf("UserRepository.CreateUser: %w", err)
	}

	return nil
}
