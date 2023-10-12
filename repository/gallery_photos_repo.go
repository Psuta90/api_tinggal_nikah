package repository

import (
	"api_tinggal_nikah/models"

	"gorm.io/gorm"
)

type GalleryPhotosRepository interface {
	CreateGalleryPhotos(gallery *[]models.GalleryPhotos) error
	UpdateGalleryPhotos(gallery models.GalleryPhotos, errChan chan error)
	GetAllGalleryPhotos() ([]models.GalleryPhotos, error)
}

type GalleryPhotosRepositoryImpl struct {
	tx *gorm.DB
}

func NewGalleryPhotosRepository(tx *gorm.DB) GalleryPhotosRepository {
	return &GalleryPhotosRepositoryImpl{tx}
}

func (gp *GalleryPhotosRepositoryImpl) CreateGalleryPhotos(gallery *[]models.GalleryPhotos) error {
	result := gp.tx.Create(gallery)
	return result.Error
}

func (gp *GalleryPhotosRepositoryImpl) UpdateGalleryPhotos(gallery models.GalleryPhotos, errChan chan error) {

	result := gp.tx.Model(&gallery).Updates(&gallery)
	if result.Error != nil {
		errChan <- result.Error
	}

	errChan <- nil
}

func (gp *GalleryPhotosRepositoryImpl) GetAllGalleryPhotos() ([]models.GalleryPhotos, error) {
	Gallery := new([]models.GalleryPhotos)
	result := gp.tx.Find(&Gallery)
	return *Gallery, result.Error
}
