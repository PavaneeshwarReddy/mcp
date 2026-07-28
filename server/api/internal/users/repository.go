package users

import "gorm.io/gorm"

type Repository interface {
	ExistsByUsername(username string) (bool, error)
	Create(user User) error
}

type UserRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &UserRepository{db: db}
}

func (r *UserRepository) ExistsByUsername(username string) (bool, error) {

	var count int64
	if err := r.db.Model(&User{}).Where("name = ?", username).Count(&count).Error; err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}

func (r *UserRepository) Create(user User) error {
	return r.db.Create(&user).Error
}
