package repository

import (
	"api_tinggal_nikah/models"

	"gorm.io/gorm"
)

type PackagesRepository interface {
	CreatePackage(packages *models.Package) error
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
