package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CreditCards  []CreditCard `gorm:"foreignKey:UserNumber;references:MemberNumber"`

type User struct {
	gorm.Model
	ID              uuid.UUID `gorm:"default:uuid_generate_v4();primaryKey"`
	FullName        string
	Email           string
	Password        string
	Role            RoleStatus        `gorm:"type:role_status"`
	Acara           []Acara           `gorm:"foreignKey:UserID;references:ID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	GalleryPhotos   []GalleryPhotos   `gorm:"foreignKey:UserID;references:ID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	LoveStory       []LoveStory       `gorm:"foreignKey:UserID;references:ID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	GuestBook       []GuestBook       `gorm:"foreignKey:UserID;references:ID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	MempelaiPria    MempelaiPria      `gorm:"foreignKey:UserID;references:ID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	MempelaiWanita  MempelaiWanita    `gorm:"foreignKey:UserID;references:ID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	GiftDigital     []GiftDigital     `gorm:"foreignKey:UserID;references:ID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Domain          Domain            `gorm:"foreignKey:UserID;references:ID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TemplateUser    TemplateUser      `gorm:"foreignKey:UserID;references:ID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserTransaction []UserTransaction `gorm:"foreignKey:UserID;references:ID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserPackage     []UserPackage     `gorm:"foreignKey:UserID;references:ID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	MusicUser       []MusicUser       `gorm:"foreignKey:UserID;references:ID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type RoleStatus string

const (
	Admin    RoleStatus = "admin"
	Customer RoleStatus = "customer"
)
