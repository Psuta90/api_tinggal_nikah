package migration

import (
	"github.com/Psuta90/api_tinggal_nikah/models"

	"gorm.io/gorm"
)

type CreatePackageCategory struct{}

func (m *CreatePackageCategory) Up(db *gorm.DB) error {
	// Check if the table already exists
	if db.Migrator().HasTable(&models.PackageCategory{}) {
		// Drop the table before migrating
		if err := db.Migrator().DropTable(&models.PackageCategory{}); err != nil {
			return err
		}
	}

	// Migrate the table
	return db.AutoMigrate(&models.PackageCategory{})
}

func (m *CreatePackageCategory) Down(db *gorm.DB) error {
	// Rollback by dropping the table
	return db.Migrator().DropTable(&models.PackageCategory{})
}
