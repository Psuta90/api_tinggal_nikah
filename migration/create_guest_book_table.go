package migration

import (
	"api_tinggal_nikah/models"

	"gorm.io/gorm"
)

type CreateGuestBook struct{}

func (m *CreateGuestBook) Up(db *gorm.DB) error {
	// Check if the table already exists
	if db.Migrator().HasTable(&models.GuestBook{}) {
		// Drop the table before migrating
		if err := db.Migrator().DropTable(&models.GuestBook{}); err != nil {
			return err
		}
	}

	// Migrate the table
	return db.AutoMigrate(&models.GuestBook{})
}

func (m *CreateGuestBook) Down(db *gorm.DB) error {
	// Rollback by dropping the table
	return db.Migrator().DropTable(&models.GuestBook{})
}
