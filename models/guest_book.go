package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GuestBook struct {
	gorm.Model
	ID           uuid.UUID `gorm:"default:uuid_generate_v4();primaryKey"`
	UserID       string
	Group        string
	Name         string
	Phone        string
	LinkWhatsapp string
	Message      string
}
