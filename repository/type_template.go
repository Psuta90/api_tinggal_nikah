package repository

import (
	"api_tinggal_nikah/models"

	"gorm.io/gorm"
)

type TypeTemplateRepository interface {
	GetAllTemplateType() ([]models.TypeTemplate, error)
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
