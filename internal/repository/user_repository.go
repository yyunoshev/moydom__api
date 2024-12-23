package repository

import (
	"gorm.io/gorm"
	"moydom_api/internal/domain"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindByID(id uint) (domain.User, error) {
	var user domain.User
	if err := r.db.First(&user, id).Error; err != nil {
		return domain.User{}, err
	}
	return user, nil
}
func (r *UserRepository) GetByUsername(username string) (domain.User, error) {
	var user domain.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		if err.Error() == "record not found" {
			return domain.User{}, nil
		}
		return domain.User{}, err
	}
	return user, err
}
func (r *UserRepository) Create(user domain.User) (domain.User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return domain.User{}, err
	}
	return user, nil
}
