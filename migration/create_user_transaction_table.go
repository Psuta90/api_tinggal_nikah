package migration

import (
	"api_tinggal_nikah/models"

	"gorm.io/gorm"
)

type CreateUsersTransactionTable struct{}

func (m *CreateUsersTransactionTable) Up(db *gorm.DB) error {
	// Check if the table already exists
	if db.Migrator().HasTable(&models.UserTransaction{}) {
		// Drop the table before migrating
		if err := db.Migrator().DropTable(&models.UserTransaction{}); err != nil {
			return err
		}
	}

	// Migrate the table
	return db.AutoMigrate(&models.UserTransaction{})
}

func (m *CreateUsersTransactionTable) Down(db *gorm.DB) error {
	// Rollback by dropping the table
	return db.Migrator().DropTable(&models.UserTransaction{})
}
