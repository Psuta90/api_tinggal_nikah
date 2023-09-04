package repository

import (
	"api_tinggal_nikah/models"

	"gorm.io/gorm"
)

type MempelaiPriaRepository interface {
	CreateMempelaiPria(mempelaipria *models.MempelaiPria) error
}

type MempelaiPriaRepositoryImpl struct {
	tx *gorm.DB
}

func NewMempelaiPriaRepository(tx *gorm.DB) MempelaiPriaRepository {
	return &MempelaiPriaRepositoryImpl{tx}
}

func (mpr *MempelaiPriaRepositoryImpl) CreateMempelaiPria(mempelaipria *models.MempelaiPria) error {
	result := mpr.tx.Create(mempelaipria)
	return result.Error
}
