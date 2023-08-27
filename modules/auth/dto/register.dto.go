package dto

type Register struct {
	FullName         string `json:"fullname" validate:"required"`
	Email            string `json:"email" validate:"required,email"`
	Password         string `json:"password" validate:"required"`
	Confirm_Password string `json:"confirm_password" validate:"required"`
}
