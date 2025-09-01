package users

import "auth-api/internal/database"

type UsersRepository interface {
}

type usersRepositoryImpl struct {
	db database.Database
}

func NewRepository(db database.Database) UsersRepository {
	return &usersRepositoryImpl{
		db: db,
	}
}
