package dto

type AddPackagesCategorysDto struct {
	Name               string `json:"name" validate:"required"`
	Price              int    `json:"price" validate:"required"`
	DiscountPercentage int    `json:"discountprecentage" `
	ActiveDays         int    `json:"active_days" validate:"required"`
}
