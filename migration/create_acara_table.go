package migration

import (
	"api_tinggal_nikah/models"

	"gorm.io/gorm"
)

type CreateAcaraTable struct{}

func (m *CreateAcaraTable) Up(db *gorm.DB) error {
	// Check if the table already exists
	if db.Migrator().HasTable(&models.Acara{}) {
		// Drop the table before migrating
		if err := db.Migrator().DropTable(&models.Acara{}); err != nil {
			return err
		}
	}

	// Migrate the table
	return db.AutoMigrate(&models.Acara{})
}

func (m *CreateAcaraTable) Down(db *gorm.DB) error {
	// Rollback by dropping the table
	return db.Migrator().DropTable(&models.Acara{})
}
