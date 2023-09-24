package repository

import (
	"api_tinggal_nikah/models"
	"fmt"

	"gorm.io/gorm"
)

type UserTransactionRepository interface {
	GetLastID() (int64, error)
	Create(usertransaction *models.UserTransaction) error
	FindOneByOrderID(order_id string) (*models.UserTransaction, error)
	UpdateByOrderID(usertransaction *models.UserTransaction) error
}

type UserTransactionRepositoryImpl struct {
	tx *gorm.DB
}

func NewUserTransactionRepository(tx *gorm.DB) UserTransactionRepository {
	return &UserTransactionRepositoryImpl{tx}
}

func (ut *UserTransactionRepositoryImpl) GetLastID() (int64, error) {
	var TotalRows int64
	result := ut.tx.Model(&models.UserTransaction{}).Count(&TotalRows)

	return TotalRows, result.Error
}

func (ut *UserTransactionRepositoryImpl) Create(usertransaction *models.UserTransaction) error {
	result := ut.tx.Create(usertransaction)

	return result.Error
}

func (ut *UserTransactionRepositoryImpl) FindOneByOrderID(order_id string) (*models.UserTransaction, error) {
	userTransaction := new(models.UserTransaction)

	result := ut.tx.First(&userTransaction, "order_id = ?", order_id)

	return userTransaction, result.Error

}

func (ut *UserTransactionRepositoryImpl) UpdateByOrderID(usertransaction *models.UserTransaction) error {

	result := ut.tx.Model(&usertransaction).Where("order_id = ?", usertransaction.OrderID).Updates(usertransaction)

	fmt.Println(result.Error)

	return result.Error

}
