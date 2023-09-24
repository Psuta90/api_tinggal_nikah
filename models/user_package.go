package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserPackage struct {
	gorm.Model
	ID                uuid.UUID `gorm:"default:uuid_generate_v4();primaryKey"`
	UserTransactionID uuid.UUID
	PackageCategoryID uuid.UUID
	UserID            uuid.UUID
	StartDate         time.Time
	EndDate           time.Time
	IsActive          bool
	PackageCategory   PackageCategory `gorm:"foreignKey:PackageCategoryID; references:ID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserTransaction   UserTransaction `gorm:"foreignKey:UserTransactionID; references:ID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
