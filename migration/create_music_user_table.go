package migration

import (
	"api_tinggal_nikah/models"

	"gorm.io/gorm"
)

type CreateMusicUserTable struct{}

func (m *CreateMusicUserTable) Up(db *gorm.DB) error {
	// Check if the table already exists
	if db.Migrator().HasTable(&models.MusicUser{}) {
		// Drop the table before migrating
		if err := db.Migrator().DropTable(&models.MusicUser{}); err != nil {
			return err
		}
	}

	// Migrate the table
	return db.AutoMigrate(&models.MusicUser{})
}

func (m *CreateMusicUserTable) Down(db *gorm.DB) error {
	// Rollback by dropping the table
	return db.Migrator().DropTable(&models.MusicUser{})
}
