package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CreditCards  []CreditCard `gorm:"foreignKey:UserNumber;references:MemberNumber"`

type User struct {
	gorm.Model
	ID             uuid.UUID `gorm:"default:uuid_generate_v4();primaryKey"`
	FullName       string
	Email          string
	Password       string
	Role           RoleStatus      `gorm:"type:role_status"`
	Acara          []Acara         `gorm:"foreignKey:UserID;references:ID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	GalleryPhotos  []GalleryPhotos `gorm:"foreignKey:UserID;references:ID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	LoveStory      []LoveStory     `gorm:"foreignKey:UserID;references:ID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	GuestBook      []GuestBook     `gorm:"foreignKey:UserID;references:ID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	MempelaiPria   MempelaiPria    `gorm:"foreignKey:UserID;references:ID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	MempelaiWanita MempelaiWanita  `gorm:"foreignKey:UserID;references:ID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	GiftDigital    []GiftDigital   `gorm:"foreignKey:UserID;references:ID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Domain         Domain          `gorm:"foreignKey:UserID;references:ID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type RoleStatus string

const (
	Admin    RoleStatus = "admin"
	Customer RoleStatus = "customer"
)
