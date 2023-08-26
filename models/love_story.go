package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LoveStory struct {
	gorm.Model
	ID       uuid.UUID `gorm:"default:uuid_generate_v4();primaryKey"`
	Title    string
	Location string
	Story    string `gorm:"type:text"`
	Orders   uint
	UserID   string
}
