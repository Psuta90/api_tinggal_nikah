package models

import "gorm.io/gorm"

type CasbinRule struct {
	gorm.Model
	PType string `gorm:"size:100;index"`
	V0    string `gorm:"size:100"`
	V1    string `gorm:"size:100"`
	V2    string `gorm:"size:100"`
	V3    string `gorm:"size:100"`
	V4    string `gorm:"size:100"`
	V5    string `gorm:"size:100"`
}
