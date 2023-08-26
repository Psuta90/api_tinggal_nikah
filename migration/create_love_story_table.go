package migration

import (
	"api_tinggal_nikah/models"

	"gorm.io/gorm"
)

type CreateLoveStoryTable struct{}

func (m *CreateLoveStoryTable) Up(db *gorm.DB) error {
	// Check if the table already exists
	if db.Migrator().HasTable(&models.LoveStory{}) {
		// Drop the table before migrating
		if err := db.Migrator().DropTable(&models.LoveStory{}); err != nil {
			return err
		}
	}

	// Migrate the table
	return db.AutoMigrate(&models.LoveStory{})
}

func (m *CreateLoveStoryTable) Down(db *gorm.DB) error {
	// Rollback by dropping the table
	return db.Migrator().DropTable(&models.LoveStory{})
}
