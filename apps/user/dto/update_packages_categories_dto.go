package dto

import "github.com/google/uuid"

type UpdatePackagesCategorysDto struct {
	ID                 uuid.UUID `param:"id" validate:"required"`
	Name               string    `json:"name" `
	Price              int       `json:"price" `
	DiscountPercentage int       `json:"discount_precentage" `
	ActiveDays         int       `json:"active_days" `
}
