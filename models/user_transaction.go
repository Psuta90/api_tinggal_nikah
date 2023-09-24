package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserTransaction struct {
	gorm.Model
	ID                uuid.UUID `gorm:"default:uuid_generate_v4();primaryKey"`
	OrderID           string
	UserID            uuid.UUID
	ResponseTripay    JSONMap `gorm:"type:jsonb"`
	ExpiredOrder      time.Time
	PackageCategoryID uuid.UUID
	Status            StatusPayment   `gorm:"type:status_payment"`
	PackageCategory   PackageCategory `gorm:"foreignKey:PackageCategoryID; references:ID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type StatusPayment string

const (
	PAID    StatusPayment = "PAID"
	FAILED  StatusPayment = "FAILED"
	EXPIRED StatusPayment = "EXPIRED"
	REFUND  StatusPayment = "REFUND"
	PENDING StatusPayment = "PENDING"
)

type JSONMap map[string]interface{}

func (j JSONMap) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *JSONMap) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("Scan: expected []byte, got %T", value)
	}
	var m map[string]interface{}
	if err := json.Unmarshal(bytes, &m); err != nil {
		return err
	}
	*j = m
	return nil
}
