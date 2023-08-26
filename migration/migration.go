package migration

import "gorm.io/gorm"

type Migration interface {
	Up(db *gorm.DB) error
	Down(db *gorm.DB) error
}

var Migrations = []Migration{}
