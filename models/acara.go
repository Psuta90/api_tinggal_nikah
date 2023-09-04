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
	Place     string
	Location  string
	Orders    int
	UserID    uuid.UUID
}
