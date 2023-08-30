package migration

import (
	"github.com/Psuta90/api_tinggal_nikah/models"
	"gorm.io/gorm"
)

type CreateCasbinRule struct{}

func (m *CreateCasbinRule) Up(db *gorm.DB) error {
	// Check if the table already exists
	if db.Migrator().HasTable(&models.CasbinRule{}) {
		// Drop the table before migrating
		if err := db.Migrator().DropTable(&models.CasbinRule{}); err != nil {
			return err
		}
	}

	// Migrate the table
	return db.AutoMigrate(&models.CasbinRule{})
}

func (m *CreateCasbinRule) Down(db *gorm.DB) error {
	// Rollback by dropping the table
	return db.Migrator().DropTable(&models.CasbinRule{})
}
