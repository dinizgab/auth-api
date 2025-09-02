package users

import (
	"auth-api/internal/database"
	"context"
	"fmt"

	_ "embed"
)

type Repository interface {
	CreateUser(ctx context.Context, user User) error
	GetUserByEmail(ctx context.Context, email string) (User, error)
}

type usersRepositoryImpl struct {
	db database.Database
}

var (
	//go:embed sql/create_user.sql
	CreateUserQuery string
	//go:embed sql/get_user_by_email.sql
	GetUserByEmailQuery string
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

func (r *usersRepositoryImpl) GetUserByEmail(ctx context.Context, email string) (User, error) {
	var user User
	err := r.db.QueryRow(ctx, GetUserByEmailQuery, email).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return User{}, fmt.Errorf("UserRepository.GetUserByEmail: %w", err)
	}

	return user, nil
}
