package dto

import "github.com/google/uuid"

type AddPackagesDto struct {
	GuestSize         int       `json:"guest_size" validate:"required"`
	GallerySize       int       `json:"gallery_size" validate:"required"`
	VideoSize         int       `json:"video_size" validate:"required"`
	RSVP              bool      `json:"rsvp" validate:"required"`
	LocationLink      bool      `json:"location_link" validate:"required"`
	Story             bool      `json:"story" validate:"required"`
	GiftDigital       bool      `json:"gift_digital" validate:"required"`
	Music             bool      `json:"music" validate:"required"`
	PackageCategoryID uuid.UUID `json:"package_category_id" validate:"required"`
}
