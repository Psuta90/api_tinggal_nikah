package dto

import "github.com/google/uuid"

type UpdatePackagesDto struct {
	ID                uuid.UUID `param:"id" validate:"required"`
	GuestSize         int       `json:"guest_size" `
	GallerySize       int       `json:"gallery_size" `
	VideoSize         int       `json:"video_size" `
	RSVP              bool      `json:"rsvp" `
	LocationLink      bool      `json:"location_link" `
	Story             bool      `json:"story" `
	GiftDigital       bool      `json:"gift_digital" `
	Music             bool      `json:"music" `
	PackageCategoryID uuid.UUID `json:"package_category_id" `
}
