package repository

import (
	"api_tinggal_nikah/models"

	"gorm.io/gorm"
)

type PackagesRepository interface {
	CreatePackage(packages *models.Package) error
	UpdatePackage(packages *models.Package) error
	DeletePackage(packages *models.Package) error
}

type PackagesRepositoryImpl struct {
	tx *gorm.DB
}

func NewPackagesRepository(tx *gorm.DB) PackagesRepository {
	return &PackagesRepositoryImpl{tx}
}

func (p *PackagesRepositoryImpl) CreatePackage(packages *models.Package) error {
	result := p.tx.Create(packages)
	return result.Error
}

func (p *PackagesRepositoryImpl) UpdatePackage(packages *models.Package) error {
	result := p.tx.Model(&packages).Updates(packages)
	return result.Error
}

func (p *PackagesRepositoryImpl) DeletePackage(packages *models.Package) error {
	result := p.tx.Unscoped().Delete(&packages)
	return result.Error
}
