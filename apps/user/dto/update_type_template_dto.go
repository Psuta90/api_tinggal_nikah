package dto

import "github.com/google/uuid"

type UpdateTypeTemplateDto struct {
	ID uuid.UUID `param:"id"`
	AddTypeTemplateDto
}
