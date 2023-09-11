package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TemplateUser struct {
	gorm.Model
	ID         uuid.UUID `gorm:"default:uuid_generate_v4();primaryKey"`
	TemplateID uuid.UUID
	UserID     uuid.UUID `gorm:"unique"`
}
