package repository

import (
	"api_tinggal_nikah/models"

	"gorm.io/gorm"
)

type PackageCategoryRepository interface {
	CreatePackageCategory(packagecategory *models.PackageCategory) error
	// UpdatePackageCategory(packagecategory *models.PackageCategory, errChan chan error)
}

type PackageCategoryRepositoryImpl struct {
	tx *gorm.DB
}

func NewPackageCategoryRepository(tx *gorm.DB) PackageCategoryRepository {
	return &PackageCategoryRepositoryImpl{tx}
}

func (pc *PackageCategoryRepositoryImpl) CreatePackageCategory(packagecategory *models.PackageCategory) error {
	result := pc.tx.Create(packagecategory)
	return result.Error
}

// func (pc *PackageCategoryRepositoryImpl) UpdateGuestBook(guestbook *models.GuestBook, errChan chan error) {
// 	result := gr.tx.Model(&guestbook).Updates(&guestbook)
// 	if result.Error != nil {
// 		errChan <- result.Error
// 	}

// 	errChan <- nil
// }
