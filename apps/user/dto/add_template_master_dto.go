package dto

import "github.com/google/uuid"

type AddTemplateMasterDto struct {
	Name           string    `json:"name" validate:"required"`
	Css            string    `json:"css" `
	TypeTemplateID uuid.UUID `json:"id_template" validate:"required"`
}
