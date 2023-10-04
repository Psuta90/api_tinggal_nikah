package utils

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {

	if err := cv.validator.Struct(i); err != nil {
		var errorMessage string

		for _, err := range err.(validator.ValidationErrors) {
			switch err.Tag() {
			case "required":
				// Pesan kesalahan kustom untuk tag "required"
				errorMessage = fmt.Sprintf("%s : is required", err.StructNamespace())
			case "email":
				// Pesan kesalahan default untuk tag "email"
				errorMessage = fmt.Sprintf("%s : is not a valid email address", err.StructNamespace())
			case "valid-image":
				// Pesan kesalahan kustom untuk tag "required"
				errorMessage = fmt.Sprintf("%s : is not a valid image check size less then 5mb or format image", err.StructNamespace())
			case "valid-music":
				errorMessage = fmt.Sprintf("%s : is not a valid music check size less then 5mb or format music is must mp3", err.StructNamespace())
			}

		}
		return echo.NewHTTPError(http.StatusBadRequest, errorMessage)
	}

	return nil

}

func NewCustomValidator() echo.Validator {
	v := validator.New()
	v.RegisterValidation("valid-image", ValidImage)
	v.RegisterValidation("valid-music", ValidMusic)
	cv := &CustomValidator{validator: v}

	return cv
}

// ini untuk bisa mengirimkan beberapa struct untuk di validate sekaligus
func Validation(c echo.Context, structs ...interface{}) error {

	for _, s := range structs {
		if err := c.Validate(s); err != nil {
			return err
		}
	}

	return nil
}

// custom validation
func ValidImage(fl validator.FieldLevel) bool {
	files := fl.Field().Interface().([]*multipart.FileHeader)

	if len(files) == 0 {
		return false // Minimal satu file harus diunggah
	}
	maxSize := int64(1048576) * 5 // 1mb

	for _, file := range files {
		if !isValidFormatImage(file) {
			return false
		}

		if file.Size >= maxSize {
			return false
		}
	}

	return true
}

func ValidMusic(fl validator.FieldLevel) bool {
	files := fl.Field().Interface().([]*multipart.FileHeader)

	if len(files) == 0 {
		return false // Minimal satu file harus diunggah
	}
	maxSize := int64(1048576) * 5 // 1mb

	for _, file := range files {
		if !isValidFormatMusic(file) {
			return false
		}

		if file.Size >= maxSize {
			return false
		}
	}

	return true
}

func isValidFormatImage(file *multipart.FileHeader) bool {
	ext := strings.ToLower(filepath.Ext(file.Filename))

	// Tentukan ekstensi yang diizinkan (misalnya, jpg, jpeg, dan png)
	allowedExtensions := []string{".jpg", ".jpeg", ".png"}

	// Lakukan pengecekan apakah ekstensi file ada dalam daftar yang diizinkan
	for _, allowedExt := range allowedExtensions {
		if ext == allowedExt {
			return true
		}
	}

	// Jika ekstensi file tidak ada dalam daftar yang diizinkan, maka return false
	return false
}

func isValidFormatMusic(file *multipart.FileHeader) bool {
	ext := strings.ToLower(filepath.Ext(file.Filename))

	// Tentukan ekstensi yang diizinkan (misalnya, jpg, jpeg, dan png)
	allowedExtensions := []string{".mp3"}

	// Lakukan pengecekan apakah ekstensi file ada dalam daftar yang diizinkan
	for _, allowedExt := range allowedExtensions {
		if ext == allowedExt {
			return true
		}
	}

	// Jika ekstensi file tidak ada dalam daftar yang diizinkan, maka return false
	return false
}

// end custom validation
