package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MusicMaster struct {
	gorm.Model
	ID   uuid.UUID `gorm:"default:uuid_generate_v4();primaryKey"`
	Path string
	Name string
}
