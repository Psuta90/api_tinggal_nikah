package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GiftDigital struct {
	gorm.Model
	ID           uuid.UUID `gorm:"default:uuid_generate_v4();primaryKey"`
	UserID       uuid.UUID
	NoRekening   uint
	PaymentType  string
	NameRekening string
	Orders       int
}
