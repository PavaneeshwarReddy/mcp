package users

import "gorm.io/gorm"

type Repository interface {
	Exists()
}

type UserRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Exists() {

}
