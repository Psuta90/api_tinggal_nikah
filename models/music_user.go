package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MusicUser struct {
	gorm.Model
	ID            uuid.UUID `gorm:"default:uuid_generate_v4();primaryKey"`
	UserID        uuid.UUID `gorm:"unique"`
	MusicMasterID uuid.UUID
	MusicMaster   MusicMaster `gorm:"foreignKey:MusicMasterID; references:ID"`
}
