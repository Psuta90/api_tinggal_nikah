package repository

import (
	"api_tinggal_nikah/models"

	"gorm.io/gorm"
)

type GuestBookRepository interface {
	CreateGuestBook(guestbook *[]models.GuestBook) error
	UpdateGuestBook(guestbook *models.GuestBook, errChan chan error)
}

type GuestBookRepositoryImpl struct {
	tx *gorm.DB
}

func NewGuestBookRepository(tx *gorm.DB) GuestBookRepository {
	return &GuestBookRepositoryImpl{tx}
}

func (gr *GuestBookRepositoryImpl) CreateGuestBook(guestbook *[]models.GuestBook) error {
	result := gr.tx.Create(guestbook)
	return result.Error
}

func (gr *GuestBookRepositoryImpl) UpdateGuestBook(guestbook *models.GuestBook, errChan chan error) {
	result := gr.tx.Model(&guestbook).Updates(&guestbook)
	if result.Error != nil {
		errChan <- result.Error
	}

	errChan <- nil
}
