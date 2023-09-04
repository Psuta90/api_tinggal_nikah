package dto

import (
	"mime/multipart"
	"time"
)

type AddWeddingDto struct {
	MempelaiJSONSTR    string `form:"mempelai" json:"mempelai"`
	AcaraJSONSTR       string `form:"acara" json:"acara"`
	LoveStoryJSONSTR   string `form:"lovestory" json:"lovestory"`
	GiftDigitalJSONSTR string `form:"gift_digital" json:"gift_digital"`
	GusetBookJSONSTR   string `form:"guest_book" json:"guest_book"`
	Subdomain          string `form:"subdomain"`
	PremiumDomain      string `form:"premiumdomain"`
}

type AddWeddingJSON struct {
	Mempelai            DataMempelai            `validate:"required"`
	Acara               []DataAcara             `validate:"required,dive,required"`
	LoveStory           []DataLoveStory         `validate:"required,dive,required"`
	GiftDigital         []DataGiftDigital       `validate:"required,dive,required"`
	GuestBook           []DataGuestBook         `validate:"required,dive,required"`
	HalamanUtamaGallery []*multipart.FileHeader `validate:"required,valid-image" `
	GalleryPhotos       []*multipart.FileHeader `validate:"required,valid-image"`
	Subdomain           string                  `validate:"required"`
	PremiumDomain       string
}

type DataMempelai struct {
	MempelaiPria   Mempelai `json:"mempelai_pria" validate:"required"`
	MempelaiWanita Mempelai `json:"mempelai_wanita" validate:"required"`
}

type DataAcara struct {
	Title     string    `json:"title" validate:"required"`
	StartDate time.Time `json:"start_date" validate:"required"`
	EndDate   time.Time `json:"end_date" validate:"required"`
	Location  string    `json:"location" validate:"required"`
	Place     string    `json:"place" validate:"required"`
	Order     int       `json:"order" validate:"required"`
}

type DataLoveStory struct {
	Title    string `json:"title" validate:"required"`
	Location string `json:"location" validate:"required"`
	Story    string `json:"story" validate:"required"`
	Order    int    `json:"order" validate:"required"`
}

type DataGiftDigital struct {
	NoRekening   string `json:"no_rekening" validate:"required"`
	PaymentType  string `json:"payment_type" validate:"required"`
	NameRekening string `json:"nama_rekening" validate:"required"`
	Order        int    `json:"order" validate:"required"`
}

type DataGuestBook struct {
	Group          string      `json:"group" validate:"required"`
	Order          int         `json:"order" validate:"required"`
	DatasGuestBook []GuestBook `json:"data_guestbook" validate:"required,dive,required"`
}

type GuestBook struct {
	Name    string `json:"name" validate:"required"`
	Phone   string
	Message string
}
type Mempelai struct {
	IsLeft     bool   `json:"is_left" `
	NameAlias  string `json:"namealias" validate:"required"`
	Fullname   string `json:"fullname" validate:"required"`
	NameFather string `json:"namefather" validate:"required"`
	NameMother string `json:"namemother" validate:"required"`
}
