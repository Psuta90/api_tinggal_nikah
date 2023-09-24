package migration

import (
	"api_tinggal_nikah/models"

	"gorm.io/gorm"
)

type CreateUsersPackage struct{}

func (m *CreateUsersPackage) Up(db *gorm.DB) error {
	// Check if the table already exists
	if db.Migrator().HasTable(&models.UserPackage{}) {
		// Drop the table before migrating
		if err := db.Migrator().DropTable(&models.UserPackage{}); err != nil {
			return err
		}
	}

	// Migrate the table
	return db.AutoMigrate(&models.UserPackage{})
}

func (m *CreateUsersPackage) Down(db *gorm.DB) error {
	// Rollback by dropping the table
	return db.Migrator().DropTable(&models.UserPackage{})
}
