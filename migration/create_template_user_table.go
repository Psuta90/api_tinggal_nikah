package migration

import (
	"api_tinggal_nikah/models"

	"gorm.io/gorm"
)

type TemplateUser struct{}

func (m *TemplateUser) Up(db *gorm.DB) error {
	// Check if the table already exists
	if db.Migrator().HasTable(&models.TemplateUser{}) {
		// Drop the table before migrating
		if err := db.Migrator().DropTable(&models.TemplateUser{}); err != nil {
			return err
		}
	}

	// Migrate the table
	return db.AutoMigrate(&models.TemplateUser{})
}

func (m *TemplateUser) Down(db *gorm.DB) error {
	// Rollback by dropping the table
	return db.Migrator().DropTable(&models.TemplateUser{})
}
