package migration

import (
	"api_tinggal_nikah/models"

	"gorm.io/gorm"
)

type TemplateType struct{}

func (m *TemplateType) Up(db *gorm.DB) error {
	// Check if the table already exists
	if db.Migrator().HasTable(&models.TypeTemplate{}) {
		// Drop the table before migrating
		if err := db.Migrator().DropTable(&models.TypeTemplate{}); err != nil {
			return err
		}
	}

	// Migrate the table
	return db.AutoMigrate(&models.TypeTemplate{})
}

func (m *TemplateType) Down(db *gorm.DB) error {
	// Rollback by dropping the table
	return db.Migrator().DropTable(&models.TypeTemplate{})
}
