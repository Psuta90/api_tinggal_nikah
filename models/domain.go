package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Domain struct {
	gorm.Model
	ID            uuid.UUID `gorm:"default:uuid_generate_v4();primaryKey"`
	Subdomain     string
	PremiumDomain string
	UserID        uuid.UUID `gorm:"unique"`
}
