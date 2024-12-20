package repository

import (
	"api_tinggal_nikah/models"

	"gorm.io/gorm"
)

type DomainRepository interface {
	CreateDomain(domain *models.Domain) error
	UpdateDomain(domain *models.Domain) error
}

type DomainRepositoryImpl struct {
	tx *gorm.DB
}

func NewDomainRepository(tx *gorm.DB) DomainRepository {
	return &DomainRepositoryImpl{tx}
}

func (dr *DomainRepositoryImpl) CreateDomain(domain *models.Domain) error {
	result := dr.tx.Create(domain)

	return result.Error
}

func (dr *DomainRepositoryImpl) UpdateDomain(domain *models.Domain) error {
	result := dr.tx.Model(&domain).Where("id = ?", domain.ID).Updates(domain)
	return result.Error
}
