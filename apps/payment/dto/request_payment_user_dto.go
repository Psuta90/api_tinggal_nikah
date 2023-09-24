package dto

import "github.com/google/uuid"

type RequestPaymentUserDto struct {
	Method     string       `json:"method"`
	Amount     int          `json:"amount"`
	PackageID  uuid.UUID    `json:"package_id"`
	OrderItems []OrdersItem `json:"order_items"`
}
