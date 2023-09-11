package repository

import (
	"api_tinggal_nikah/models"

	"gorm.io/gorm"
)

type LoveStoryRepository interface {
	CreateLoveStory(lovestory *[]models.LoveStory) error
	UpdateLoveStory(lovestory *models.LoveStory, errChan chan error)
}

type LoveStoryRepositoryImpl struct {
	tx *gorm.DB
}

func NewLoveStoryRepository(tx *gorm.DB) LoveStoryRepository {
	return &LoveStoryRepositoryImpl{tx}
}

func (lr *LoveStoryRepositoryImpl) CreateLoveStory(lovestory *[]models.LoveStory) error {
	result := lr.tx.Create(lovestory)
	return result.Error

}

func (lr *LoveStoryRepositoryImpl) UpdateLoveStory(lovestory *models.LoveStory, errChan chan error) {
	result := lr.tx.Model(&lovestory).Updates(&lovestory)
	if result.Error != nil {
		errChan <- result.Error
	}

	errChan <- nil

}
