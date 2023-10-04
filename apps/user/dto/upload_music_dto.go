package dto

import "mime/multipart"

type UploadMusicDto struct {
	Files []*multipart.FileHeader `validate:"required,valid-music" `
}
