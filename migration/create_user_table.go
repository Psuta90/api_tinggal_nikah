package migration

import (
	"github.com/Psuta90/api_tinggal_nikah/models"

	"gorm.io/gorm"
)

type CreateUsersTable struct{}

func (m *CreateUsersTable) Up(db *gorm.DB) error {
	// Check if the table already exists
	if db.Migrator().HasTable(&models.User{}) {
		// Drop the table before migrating
		if err := db.Migrator().DropTable(&models.User{}); err != nil {
			return err
		}
	}

	// Migrate the table
	return db.AutoMigrate(&models.User{})
}

func (m *CreateUsersTable) Down(db *gorm.DB) error {
	// Rollback by dropping the table
	return db.Migrator().DropTable(&models.User{})
}
