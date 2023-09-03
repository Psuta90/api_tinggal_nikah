package repository

import (
	"api_tinggal_nikah/models"

	"gorm.io/gorm"
)

type AcaraRepository interface {
	CreateAcara(acara *models.Acara) (*models.Acara, error)
}

type AcaraRepositoryImpl struct {
	db *gorm.DB
}

func NewAcaraRepository(db *gorm.DB) AcaraRepository {
	return &AcaraRepositoryImpl{db}
}

func (ar *AcaraRepositoryImpl) CreateAcara(acara *models.Acara) (*models.Acara, error) {
	result := ar.db.Create(&acara)

	return acara, result.Error
}
