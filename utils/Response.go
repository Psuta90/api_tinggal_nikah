package utils

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Response interface {
	Success(c echo.Context, status int, message string, data interface{}) error
	Error(c echo.Context, status int, message string, data interface{}) error
	FailedInsertDB(c echo.Context, status int, message string, data interface{}) error
}

type DetailResponse struct {
	Status  int
	Message string
	Data    interface{}
}

type APIResponse struct {
	c echo.Context
}

func NewAPIResponse(c echo.Context) *APIResponse {
	return &APIResponse{c}
}

func (c *APIResponse) Success(status int, message string, data interface{}) error {
	dataResponse := new(DetailResponse)

	if status != 0 {
		dataResponse.Status = status
	} else {
		dataResponse.Status = http.StatusOK
	}

	if message != "" {
		dataResponse.Message = message
	} else {
		dataResponse.Message = ""
	}

	if data != nil {
		dataResponse.Data = data
	} else {
		dataResponse.Data = echo.Map{}
	}

	return c.c.JSON(dataResponse.Status, &dataResponse)

}

func (c *APIResponse) Error(status int, message string, data interface{}) error {
	dataResponse := new(DetailResponse)

	if status != 0 {
		dataResponse.Status = status
	} else {
		dataResponse.Status = http.StatusBadRequest
	}

	if message != "" {
		dataResponse.Message = message
	} else {
		dataResponse.Message = ""
	}

	if data != nil {
		dataResponse.Data = data
	} else {
		dataResponse.Data = data
	}

	return c.c.JSON(dataResponse.Status, &dataResponse)
}

func (c *APIResponse) FailedInsertDB(status int, message string, data interface{}) error {
	dataResponse := new(DetailResponse)

	if status != 0 {
		dataResponse.Status = status
	} else {
		dataResponse.Status = http.StatusUnprocessableEntity
	}

	if message != "" {
		dataResponse.Message = message
	} else {
		dataResponse.Message = ""
	}

	if data != nil {
		dataResponse.Data = data
	} else {
		dataResponse.Data = data
	}

	return c.c.JSON(dataResponse.Status, &dataResponse)
}
