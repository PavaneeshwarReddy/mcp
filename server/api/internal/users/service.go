package users

type Service struct {
	repo *Repository
}

func NewService(repository *Repository) Service {
	return Service{
		repo: repository,
	}
}
