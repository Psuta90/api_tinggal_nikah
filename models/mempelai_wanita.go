package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MempelaiWanita struct {
	gorm.Model
	ID         uuid.UUID `gorm:"default:uuid_generate_v4();primaryKey"`
	NameAlias  string
	FullName   string
	NameFather string
	NameMother string
	UserID     uuid.UUID
}
