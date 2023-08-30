package migration

import (
	"api_tinggal_nikah/models"

	"gorm.io/gorm"
)

type CreateGalleryPhotosTable struct{}

func (m *CreateGalleryPhotosTable) Up(db *gorm.DB) error {
	// Check if the table already exists
	if db.Migrator().HasTable(&models.GalleryPhotos{}) {
		// Drop the table before migrating
		if err := db.Migrator().DropTable(&models.GalleryPhotos{}); err != nil {
			return err
		}
	}

	// Migrate the table
	return db.AutoMigrate(&models.GalleryPhotos{})
}

func (m *CreateGalleryPhotosTable) Down(db *gorm.DB) error {
	// Rollback by dropping the table
	return db.Migrator().DropTable(&models.GalleryPhotos{})
}
