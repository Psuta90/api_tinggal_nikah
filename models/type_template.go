package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TypeTemplate struct {
	gorm.Model
	ID             uuid.UUID `gorm:"default:uuid_generate_v4();primaryKey"`
	Name           string
	TemplateMaster []TemplateMaster `gorm:"foreignKey:TypeTemplateID;references:ID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
