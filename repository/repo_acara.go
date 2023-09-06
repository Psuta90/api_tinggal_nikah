package repository

import (
	"api_tinggal_nikah/models"

	"gorm.io/gorm"
)

type AcaraRepository interface {
	CreateAcara(acara *[]models.Acara) error
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

// func (ar *AcaraRepositoryImpl) UpdateAcara(acara *models.Acara) error {
// 	result := ar.tx.Model(models.Acara{}).Where("id = ?", acara.ID)
// 	return result.Error
// }
