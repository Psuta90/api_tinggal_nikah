package dto

type AddTypeTemplateDto struct {
	Name string `json:"name" validate:"required"`
}
