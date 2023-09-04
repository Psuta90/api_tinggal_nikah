package repository

import (
	"api_tinggal_nikah/models"

	"gorm.io/gorm"
)

type LoveStoryRepository interface {
	CreateLoveStory(lovestory *[]models.LoveStory) error
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
