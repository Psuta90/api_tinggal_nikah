package dto

import "mime/multipart"

type AddMusicMasterDto struct {
	Files []*multipart.FileHeader `validate:"required,valid-music" `
}
