package dto

import (
	"github.com/google/uuid"
)

type UpdateMusic struct {
	Music []MusicMaster `json:"music"`
}

type MusicMaster struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Filename string    `json:"filename"`
}
