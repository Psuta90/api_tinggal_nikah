package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PackageCategory struct {
	gorm.Model
	ID                 uuid.UUID `gorm:"default:uuid_generate_v4();primaryKey"`
	Name               string
	Price              int
	DiscountPercentage int
	ActiveDays         int
	Package            Package `gorm:"foreignKey:PackageCategoryID;references:ID"`
}
