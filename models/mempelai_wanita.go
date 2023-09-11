package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MempelaiWanita struct {
	gorm.Model
	ID         uuid.UUID `gorm:"default:uuid_generate_v4();primaryKey"`
	NameAlias  string    `json:"namealias"`
	FullName   string    `json:"fullname"`
	NameFather string    `json:"namefather"`
	NameMother string    `json:"namemother"`
	IsLeft     bool
	UserID     uuid.UUID `gorm:"unique"`
}
