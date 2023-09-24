package dto

import "github.com/google/uuid"

type DeleteDto struct {
	ID uuid.UUID `param:"id" validate:"required"`
}
