package repo

import (
	"gorm.io/gorm"
	"generic.com/internal/models"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) FindByUsername(email string) (*models.User, error) {

	var user models.User
	if err := r.db.Where("username = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
