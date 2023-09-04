package repository

import (
	"api_tinggal_nikah/models"

	"gorm.io/gorm"
)

type GalleryPhotosRepository interface {
	CreateGalleryPhotos(gallery *[]models.GalleryPhotos) error
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
