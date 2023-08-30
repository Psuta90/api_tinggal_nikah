package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func AddWedding(c echo.Context) error {

	return c.JSON(http.StatusOK, "endpoint for add user")
}
