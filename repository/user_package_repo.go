package repository

import (
	"api_tinggal_nikah/models"

	"gorm.io/gorm"
)

type UserPackageRepository interface {
	Create(UserPackage *models.UserPackage) error
}

type UserPackageRepositoryImpl struct {
	tx *gorm.DB
}

func NewUserPackageRepository(tx *gorm.DB) UserPackageRepository {
	return &UserPackageRepositoryImpl{tx}
}

func (ut *UserPackageRepositoryImpl) GetLastID() (int64, error) {
	var TotalRows int64
	result := ut.tx.Model(&models.UserTransaction{}).Count(&TotalRows)

	return TotalRows, result.Error
}

func (ut *UserPackageRepositoryImpl) Create(UserPackage *models.UserPackage) error {
	result := ut.tx.Create(UserPackage)

	return result.Error
}
