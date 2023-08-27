package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Package struct {
	gorm.Model
	ID                uuid.UUID `gorm:"default:uuid_generate_v4();primaryKey"`
	GuestSize         int
	GallerySize       int
	VideoSize         int
	RSVP              bool
	LocationLink      bool
	Story             bool
	GiftDigital       bool
	Music             bool
	PackageCategoryID string
}
