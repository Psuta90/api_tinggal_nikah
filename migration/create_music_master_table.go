package migration

import (
	"api_tinggal_nikah/models"

	"gorm.io/gorm"
)

type CreateMusicMasterTable struct{}

func (m *CreateMusicMasterTable) Up(db *gorm.DB) error {
	// Check if the table already exists
	if db.Migrator().HasTable(&models.MusicMaster{}) {
		// Drop the table before migrating
		if err := db.Migrator().DropTable(&models.MusicMaster{}); err != nil {
			return err
		}
	}

	// Migrate the table
	return db.AutoMigrate(&models.MusicMaster{})
}

func (m *CreateMusicMasterTable) Down(db *gorm.DB) error {
	// Rollback by dropping the table
	return db.Migrator().DropTable(&models.MusicMaster{})
}
