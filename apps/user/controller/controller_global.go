package controller

import (
	"api_tinggal_nikah/apps/user/dto"
	"api_tinggal_nikah/apps/user/services"
	"api_tinggal_nikah/utils"

	"github.com/labstack/echo/v4"
)

func GetAllTemplates(c echo.Context) error {
	return services.GetAllTemplatesService(c)
}

func GetAllPackages(c echo.Context) error {
	return services.GetAllPackagesServices(c)
}

func UpdateRsvp(c echo.Context) error {

	guestbook := new(dto.UpdateRsvpGuestBookDto)
	if err := c.Bind(guestbook); err != nil {
		return utils.NewAPIResponse(c).Error(0, err.Error(), err)
	}

	return services.UpdateRsvpServices(c, guestbook)
}

func GetGuessByName(c echo.Context) error {
	name := c.Param("name")
	return services.GetGuestServices(c, name)
}
