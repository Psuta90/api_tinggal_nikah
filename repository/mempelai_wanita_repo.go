package repository

import (
	"api_tinggal_nikah/models"

	"gorm.io/gorm"
)

type MempelaiWanitaRepository interface {
	CreateMempelaiWanita(mempelaiWanita *models.MempelaiWanita) error
}

type MempelaiWanitaRepositoryImpl struct {
	tx *gorm.DB
}

func NewMempelaiWanitaRepository(tx *gorm.DB) MempelaiWanitaRepository {
	return &MempelaiWanitaRepositoryImpl{tx}
}

func (mpr *MempelaiWanitaRepositoryImpl) CreateMempelaiWanita(mempelaiWanita *models.MempelaiWanita) error {
	result := mpr.tx.Create(mempelaiWanita)
	return result.Error
}
