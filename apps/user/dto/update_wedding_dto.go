package dto

import (
	"time"

	"github.com/google/uuid"
)

type UpdateWeddingDto struct {
	Mempelai    UpdateDataMempelai      `json:"Mempelai" `
	Acara       []UpdateDataAcara       `json:"Acara"`
	LoveStory   []UpdateDataLoveStory   `json:"LoveStory"`
	GiftDigital []UpdateDataGiftDigital `json:"GiftDigital"`
	GuestBook   []UpdateDataGuestBook   `json:"GuestBook"`
	Domain      UpdateDomain            `json:"Domain"`
	Gallery     []UpdateGallery         `json:"Gallery"`
}

type UpdateDataMempelai struct {
	MempelaiPria   UpdateMempelai `json:"mempelai_pria" `
	MempelaiWanita UpdateMempelai `json:"mempelai_wanita" `
}

type UpdateDataAcara struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title" `
	StartDate time.Time `json:"start_date" `
	EndDate   time.Time `json:"end_date" `
	Location  string    `json:"location" `
	Place     string    `json:"place" `
	Order     int       `json:"order" `
}

type UpdateDataLoveStory struct {
	ID       uuid.UUID `json:"id"`
	Title    string    `json:"title" `
	Location string    `json:"location" `
	Story    string    `json:"story" `
	Order    int       `json:"order" `
}

type UpdateDataGiftDigital struct {
	ID           uuid.UUID `json:"id"`
	NoRekening   string    `json:"no_rekening" `
	PaymentType  string    `json:"payment_type" `
	NameRekening string    `json:"nama_rekening" `
	Order        int       `json:"order" `
}

type UpdateDataGuestBook struct {
	Group          string            `json:"group" `
	Order          int               `json:"order" `
	DatasGuestBook []UpdateGuestBook `json:"data_guestbook"`
}

type UpdateGuestBook struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name" `
	Phone   string
	Message string
}
type UpdateMempelai struct {
	ID         uuid.UUID `json:"id"`
	IsLeft     bool      `json:"is_left" `
	NameAlias  string    `json:"namealias" `
	Fullname   string    `json:"fullname" `
	NameFather string    `json:"namefather" `
	NameMother string    `json:"namemother" `
}

type UpdateGallery struct {
	ID             uuid.UUID `json:"id"`
	Filename       string    `json:"filename"`
	Order          int       `json:"order"`
	IsGallery      bool      `json:"is_gallery"`
	IsHalamanUtama bool      `json:"is_halaman_utama"`
}

type UpdateDomain struct {
	ID            uuid.UUID `json:"id"`
	Subdomain     string    `json:"Subdomain"`
	PremiumDomain string    `json:"PremiumDomain"`
}
