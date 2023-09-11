package migration

import (
	"api_tinggal_nikah/models"

	"gorm.io/gorm"
)

type TemplateMasterTable struct{}

func (m *TemplateMasterTable) Up(db *gorm.DB) error {
	// Check if the table already exists
	if db.Migrator().HasTable(&models.TemplateMaster{}) {
		// Drop the table before migrating
		if err := db.Migrator().DropTable(&models.TemplateMaster{}); err != nil {
			return err
		}
	}

	// Migrate the table
	return db.AutoMigrate(&models.TemplateMaster{})
}

func (m *TemplateMasterTable) Down(db *gorm.DB) error {
	// Rollback by dropping the table
	return db.Migrator().DropTable(&models.TemplateMaster{})
}
