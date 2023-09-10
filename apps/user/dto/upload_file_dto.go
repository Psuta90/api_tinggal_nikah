package dto

import "mime/multipart"

type UploadFileDto struct {
	HalamanUtamaGallery []*multipart.FileHeader `validate:"required,valid-image" `
	GalleryPhotos       []*multipart.FileHeader `validate:"required,valid-image"`
}
