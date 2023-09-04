// repository/repository.go
package repository

import (
	"api_tinggal_nikah/models"
	"errors"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *models.User) (*models.User, error)
	GetUserByID(id uint) (*models.User, error)
	BeforeCreateUser(user *models.User) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	// Add more methods as needed
}

type UserRepositoryImpl struct {
	tx *gorm.DB
}

func NewUserRepository(tx *gorm.DB) UserRepository {
	return &UserRepositoryImpl{tx}
}

func (ur *UserRepositoryImpl) CreateUser(user *models.User) (*models.User, error) {
	result := ur.tx.Create(&user)
	return user, result.Error
}

func (ur *UserRepositoryImpl) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	if err := ur.tx.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepositoryImpl) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := ur.tx.First(&user, "email = ?", email).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepositoryImpl) BeforeCreateUser(user *models.User) (*models.User, error) {
	result := ur.tx.First(&user, "email = ?", user.Email)

	if result.RowsAffected == 0 {
		return user, nil
	}

	return user, errors.New("user sudah ada")
}
