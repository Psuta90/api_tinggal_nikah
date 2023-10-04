package dto

import "github.com/google/uuid"

type UpdateRsvpGuestBookDto struct {
	ID               uuid.UUID `param:"id"`
	Attendences      bool      `json:"attendences"`
	MessageFromGuest string    `json:"message_from_guest"`
}
