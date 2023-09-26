package dto

import "github.com/google/uuid"

type UpdateTemplateMasterDto struct {
	ID uuid.UUID `param:"id"`
	AddTemplateMasterDto
}
