package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TemplateMaster struct {
	gorm.Model
	ID             uuid.UUID `gorm:"default:uuid_generate_v4();primaryKey"`
	Name           string
	Css            string
	TypeTemplateID uuid.UUID
	TemplateUser   TemplateUser `gorm:"foreignKey:TemplateID;references:ID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
