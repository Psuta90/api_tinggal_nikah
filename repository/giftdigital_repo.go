package repository

import (
	"api_tinggal_nikah/models"

	"gorm.io/gorm"
)

type GiftDigitalRepository interface {
	CreateGiftDigital(GiftDigital *[]models.GiftDigital) error
	UpdateGiftDigital(GiftDigital *models.GiftDigital, errChan chan error)
}

type GiftDigitalRepositoryImpl struct {
	tx *gorm.DB
}

func NewGiftDigitalRepository(tx *gorm.DB) GiftDigitalRepository {
	return &GiftDigitalRepositoryImpl{tx}
}

func (gdr *GiftDigitalRepositoryImpl) CreateGiftDigital(GiftDigital *[]models.GiftDigital) error {
	result := gdr.tx.Create(GiftDigital)
	return result.Error

}

func (gdr *GiftDigitalRepositoryImpl) UpdateGiftDigital(GiftDigital *models.GiftDigital, errChan chan error) {
	result := gdr.tx.Model(&GiftDigital).Updates(&GiftDigital)
	if result.Error != nil {
		errChan <- result.Error
	}

}
