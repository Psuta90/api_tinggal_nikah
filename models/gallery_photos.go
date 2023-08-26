package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GalleryPhotos struct {
	gorm.Model
	ID             uuid.UUID `gorm:"default:uuid_generate_v4();primaryKey"`
	Path           string
	Orders         int
	IsGallery      bool
	IsHalamanUtama bool
	UserID         string
}
