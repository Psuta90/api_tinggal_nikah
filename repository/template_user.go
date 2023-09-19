package repository

import (
	"api_tinggal_nikah/models"

	"gorm.io/gorm"
)

type TemplateUserRepository interface {
	CreateTemplateUser(templateuser *models.TemplateUser) error
	UpdateTemplateUser(templateuser *models.TemplateUser) error
}

type TemplateUserRepositoryImpl struct {
	tx *gorm.DB
}

func NewTemplateUserRepository(tx *gorm.DB) TemplateUserRepository {
	return &TemplateUserRepositoryImpl{tx}
}

func (tu *TemplateUserRepositoryImpl) CreateTemplateUser(templateuser *models.TemplateUser) error {
	result := tu.tx.Create(templateuser)
	return result.Error
}

func (tu *TemplateUserRepositoryImpl) UpdateTemplateUser(templateuser *models.TemplateUser) error {
	result := tu.tx.Model(&templateuser).Where("id = ?", templateuser.ID).Updates(templateuser)
	return result.Error
}
