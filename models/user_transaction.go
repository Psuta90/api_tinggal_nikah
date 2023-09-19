package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserTransaction struct {
	gorm.Model
	ID      uuid.UUID
	OrderID string
}
