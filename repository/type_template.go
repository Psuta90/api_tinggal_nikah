package repository

import (
	"api_tinggal_nikah/models"

	"gorm.io/gorm"
)

type TypeTemplateRepository interface {
	GetAllTemplateType() ([]models.TypeTemplate, error)
	AddTypeTemplate(typeTemplate *models.TypeTemplate) error
	UpdateTypeTemplate(typeTemplate *models.TypeTemplate) error
}

type TypeTemplateRepositoryImpl struct {
	tx *gorm.DB
}

func NewTemplateTypeRepository(tx *gorm.DB) TypeTemplateRepository {
	return &TypeTemplateRepositoryImpl{tx}
}

func (tm *TypeTemplateRepositoryImpl) GetAllTemplateType() ([]models.TypeTemplate, error) {
	typetemplates := new([]models.TypeTemplate)
	result := tm.tx.
		Preload("TemplateMaster").
		Find(&typetemplates)
	if result.Error != nil {
		return nil, result.Error
	}

	return *typetemplates, nil
}

func (tm *TypeTemplateRepositoryImpl) AddTypeTemplate(typeTemplate *models.TypeTemplate) error {
	result := tm.tx.Create(typeTemplate)
	return result.Error

}

func (tm *TypeTemplateRepositoryImpl) UpdateTypeTemplate(typeTemplate *models.TypeTemplate) error {
	result := tm.tx.Model(&typeTemplate).Updates(&typeTemplate)
	return result.Error
}
