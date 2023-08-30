package migration

import (
	"api_tinggal_nikah/models"

	"gorm.io/gorm"
)

type CreateDomainTable struct{}

func (m *CreateDomainTable) Up(db *gorm.DB) error {
	// Check if the table already exists
	if db.Migrator().HasTable(&models.Domain{}) {
		// Drop the table before migrating
		if err := db.Migrator().DropTable(&models.Domain{}); err != nil {
			return err
		}
	}

	// Migrate the table
	return db.AutoMigrate(&models.Domain{})
}

func (m *CreateDomainTable) Down(db *gorm.DB) error {
	// Rollback by dropping the table
	return db.Migrator().DropTable(&models.Domain{})
}
