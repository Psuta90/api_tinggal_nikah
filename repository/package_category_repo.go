package repository

import (
	"api_tinggal_nikah/models"

	"gorm.io/gorm"
)

type PackageCategoryRepository interface {
	CreatePackageCategory(packagecategory *models.PackageCategory) error
	UpdatePackageCategory(packagecategory *models.PackageCategory) error
	DeletePackageCategory(packagecategory *models.PackageCategory) error
	GetAllPackageCategory() ([]models.PackageCategory, error)
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

func (pc *PackageCategoryRepositoryImpl) UpdatePackageCategory(packagecategory *models.PackageCategory) error {
	result := pc.tx.Model(&packagecategory).Updates(packagecategory)
	return result.Error
}

func (pc *PackageCategoryRepositoryImpl) DeletePackageCategory(packagecategory *models.PackageCategory) error {
	result := pc.tx.Unscoped().Delete(&packagecategory)
	return result.Error
}

func (pc *PackageCategoryRepositoryImpl) GetAllPackageCategory() ([]models.PackageCategory, error) {
	packagecategory := new([]models.PackageCategory)
	result := pc.tx.Preload("Package").Find(packagecategory)
	return *packagecategory, result.Error
}
