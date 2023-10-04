package repository

import (
	"api_tinggal_nikah/models"

	"gorm.io/gorm"
)

type MusicMasterRepository interface {
	Add(musicMaster *[]models.MusicMaster) error
	Update(musicMaster *models.MusicMaster, errChan chan error)
	FindAll() (*models.MusicMaster, error)
}

type MusicMasterRepositoryImpl struct {
	tx *gorm.DB
}

func NewMusicMasterRepository(tx *gorm.DB) MusicMasterRepository {
	return &MusicMasterRepositoryImpl{tx}
}

func (mm *MusicMasterRepositoryImpl) Add(musicMaster *[]models.MusicMaster) error {
	result := mm.tx.Create(musicMaster)

	return result.Error
}

func (mm *MusicMasterRepositoryImpl) Update(musicMaster *models.MusicMaster, errChan chan error) {
	result := mm.tx.Model(&musicMaster).Updates(&musicMaster)
	if result.Error != nil {
		errChan <- result.Error
	}

	errChan <- nil
}

func (mm *MusicMasterRepositoryImpl) FindAll() (*models.MusicMaster, error) {
	musicMaster := new(models.MusicMaster)
	result := mm.tx.Find(&musicMaster)

	return musicMaster, result.Error
}
