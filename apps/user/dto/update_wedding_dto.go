package dto

import (
	"mime/multipart"

	"github.com/google/uuid"
)

type UpdateWeddingDto struct {
	WeddingDto          AddWeddingDto
	HalamanUtamaJSONSTR string `form:"halaman_utama" json:"halaman_utama"`
	GalleryPhotoJSONSTR string `form:"gallery" json:"gallery"`
}

type UpdateWeddingJSON struct {
	Mempelai            UpdateDataMempelai      `validate:"required"`
	Acara               []UpdateDataAcara       `validate:"required,dive,required"`
	LoveStory           []UpdateDataLoveStory   `validate:"required,dive,required"`
	GiftDigital         []UpdateDataGiftDigital `validate:"required,dive,required"`
	GuestBook           []UpdateDataGuestBook   `validate:"required,dive,required"`
	HalamanUtamaGallery []*multipart.FileHeader `validate:"required,valid-image" `
	GalleryPhotos       []*multipart.FileHeader `validate:"required,valid-image"`
	Subdomain           string                  `validate:"required"`
	PremiumDomain       string
	HalamanUtamaJSON    []UpdateDataGalleryPhotos
	GalleryPhotoJSON    []UpdateDataGalleryPhotos
}

type UpdateDataGalleryPhotos struct {
	ID       uuid.UUID `json:"id" validate:"required"`
	Namefile string    `json:"namefile" validate:"required"`
	Order    int       `json:"order" validate:"required"`
}

type UpdateDataMempelai struct {
	MempelaiPria   UpdateMempelai `json:"mempelai_pria" validate:"required"`
	MempelaiWanita UpdateMempelai `json:"mempelai_wanita" validate:"required"`
}

type UpdateDataAcara struct {
	ID uuid.UUID `json:"id" validate:"required"`
	DataAcara
}

type UpdateDataLoveStory struct {
	ID uuid.UUID `json:"id" validate:"required"`
	DataLoveStory
}

type UpdateDataGiftDigital struct {
	ID uuid.UUID `json:"id" validate:"required"`
	DataGiftDigital
}

type UpdateDataGuestBook struct {
	Group          string            `json:"group" validate:"required"`
	Order          int               `json:"order" validate:"required"`
	DatasGuestBook []UpdateGuestBook `json:"data_guestbook" validate:"required,dive,required"`
}

type UpdateGuestBook struct {
	ID uuid.UUID `json:"id" validate:"required"`
	GuestBook
}
type UpdateMempelai struct {
	ID uuid.UUID `json:"id" validate:"required"`
	Mempelai
}
