package migration

import (
	"api_tinggal_nikah/models"

	"gorm.io/gorm"
)

type CreatePackageTable struct{}

func (m *CreatePackageTable) Up(db *gorm.DB) error {
	// Check if the table already exists
	if db.Migrator().HasTable(&models.Package{}) {
		// Drop the table before migrating
		if err := db.Migrator().DropTable(&models.Package{}); err != nil {
			return err
		}
	}

	// Migrate the table
	return db.AutoMigrate(&models.Package{})
}

func (m *CreatePackageTable) Down(db *gorm.DB) error {
	// Rollback by dropping the table
	return db.Migrator().DropTable(&models.Package{})
}
