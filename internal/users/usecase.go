package users

type UsersUsecase interface {
}

type usersUsecaseImpl struct {
	repo UsersRepository
}

func NewUsecase(repo UsersRepository) UsersUsecase {
	return &usersUsecaseImpl{
		repo: repo,
	}
}
