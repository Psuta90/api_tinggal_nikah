package repository

import (
	"api_tinggal_nikah/models"

	"gorm.io/gorm"
)

type TemplateMasterRepository interface {
	GetAllTemplateMaster() ([]models.TemplateMaster, error)
}

type TemplateMasterRepositoryImpl struct {
	tx *gorm.DB
}

func NewTemplateMasterRepository(tx *gorm.DB) TemplateMasterRepository {
	return &TemplateMasterRepositoryImpl{tx}
}

func (tm *TemplateMasterRepositoryImpl) GetAllTemplateMaster() ([]models.TemplateMaster, error) {
	templatemasters := new([]models.TemplateMaster)
	result := tm.tx.
		Find(&templatemasters)
	if result.Error != nil {
		return nil, result.Error
	}

	return *templatemasters, nil
}

// func (tu *TemplateUserRepositoryImpl) CreateTemplateMaster(templateuser *models.TemplateUser) error {
// 	result := tu.tx.Create(templateuser)
// 	return result.Error
// }

// func (tu *TemplateUserRepositoryImpl) UpdateTemplateMaster(templateuser *models.TemplateUser) error {
// 	result := tu.tx.Model(&templateuser).Where("id = ?", templateuser.ID).Updates(templateuser)
// 	return result.Error
// }
