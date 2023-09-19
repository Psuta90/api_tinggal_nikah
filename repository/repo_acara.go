package repository

import (
	"api_tinggal_nikah/models"

	"gorm.io/gorm"
)

type AcaraRepository interface {
	CreateAcara(acara *[]models.Acara) error
	UpdateAcara(acara *models.Acara, errChan chan error)
}

type AcaraRepositoryImpl struct {
	tx *gorm.DB
}

func NewAcaraRepository(tx *gorm.DB) AcaraRepository {
	return &AcaraRepositoryImpl{tx}
}

func (ar *AcaraRepositoryImpl) CreateAcara(acara *[]models.Acara) error {
	result := ar.tx.Create(&acara)
	return result.Error
}

func (ar *AcaraRepositoryImpl) UpdateAcara(acara *models.Acara, errChan chan error) {
	result := ar.tx.Model(&acara).Updates(&acara)

	if result.Error != nil {
		errChan <- result.Error
	}

	errChan <- nil

}
