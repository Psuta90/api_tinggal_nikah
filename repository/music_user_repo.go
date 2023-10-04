package repository

import (
	"api_tinggal_nikah/models"

	"gorm.io/gorm"
)

type MusicUserRepository interface {
	Add(musicUser *models.MusicUser) error
}

type MusicUserRepositoryImpl struct {
	tx *gorm.DB
}

func NewMusicUserReporsitory(tx *gorm.DB) MusicUserRepository {
	return &MusicUserRepositoryImpl{tx}
}

func (mu *MusicUserRepositoryImpl) Add(musicUser *models.MusicUser) error {
	result := mu.tx.Create(musicUser)
	return result.Error
}
