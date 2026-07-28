package users

import "rest-mcp/internal/shared"

type Service struct {
	repo Repository
}

func NewService(repository Repository) Service {
	return Service{
		repo: repository,
	}
}

func (s *Service) CreateUser(data RegisterUserRequest) error {
	exists, err := s.repo.ExistsByUsername(data.Username)
	if err != nil {
		return err
	}
	if exists {
		return ErrUserAlreadyExisits
	}

	hashedPassword, err := shared.HashPassword(data.Password)
	if err != nil {
		return err
	}

	user := User{
		Username: data.Username,
		Email:    data.Email,
		Password: hashedPassword,
		Age:      data.Age,
	}

	return s.repo.Create(user)
}
