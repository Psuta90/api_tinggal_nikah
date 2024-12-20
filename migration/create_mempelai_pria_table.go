package migration

import (
	"api_tinggal_nikah/models"

	"gorm.io/gorm"
)

type CreateMempelaiPriaTable struct{}

func (m *CreateMempelaiPriaTable) Up(db *gorm.DB) error {
	// Check if the table already exists
	if db.Migrator().HasTable(&models.MempelaiPria{}) {
		// Drop the table before migrating
		if err := db.Migrator().DropTable(&models.MempelaiPria{}); err != nil {
			return err
		}
	}

	// Migrate the table
	return db.AutoMigrate(&models.MempelaiPria{})
}

func (m *CreateMempelaiPriaTable) Down(db *gorm.DB) error {
	// Rollback by dropping the table
	return db.Migrator().DropTable(&models.MempelaiPria{})
}
