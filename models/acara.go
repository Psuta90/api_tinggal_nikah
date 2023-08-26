package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Acara struct {
	gorm.Model
	ID        uuid.UUID `gorm:"default:uuid_generate_v4();primaryKey"`
	Title     string
	StartDate time.Time
	EndDate   time.Time
	Location  string
	Orders    uint
	UserID    string
}
