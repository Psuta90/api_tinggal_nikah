package controller

import (
	"api_tinggal_nikah/apps/user/services"

	"github.com/labstack/echo/v4"
)

func GetAllTemplates(c echo.Context) error {
	return services.GetAllTemplatesService(c)
}
