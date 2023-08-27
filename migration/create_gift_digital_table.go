package migration

import (
	"api_tinggal_nikah/models"

	"gorm.io/gorm"
)

type CreateGiftDigital struct{}

func (m *CreateGiftDigital) Up(db *gorm.DB) error {
	// Check if the table already exists
	if db.Migrator().HasTable(&models.GiftDigital{}) {
		// Drop the table before migrating
		if err := db.Migrator().DropTable(&models.GiftDigital{}); err != nil {
			return err
		}
	}

	// Migrate the table
	return db.AutoMigrate(&models.GiftDigital{})
}

func (m *CreateGiftDigital) Down(db *gorm.DB) error {
	// Rollback by dropping the table
	return db.Migrator().DropTable(&models.GiftDigital{})
}
