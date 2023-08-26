package controller

import (
	"api_tinggal_nikah/config"
	"api_tinggal_nikah/modules/auth/services"
	"api_tinggal_nikah/modules/dto"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
)

func LoginWithGoogle(c echo.Context) error {
	url := config.GoogleOauthConfig.AuthCodeURL("", oauth2.AccessTypeOffline)
	return c.Redirect(http.StatusSeeOther, url)
}

func CallbackAuthGoogle(c echo.Context) error {
	code := c.QueryParam("code")

	data, err := services.ServiceCallbackAuthGoogle(code, c)
	if err != nil {
		return c.JSON(http.StatusOK, err)
	}

	return c.JSON(http.StatusOK, data)
}

func Login(c echo.Context) error {
	user := new(dto.LoginDto)

	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"code":    http.StatusBadRequest,
			"message": "invalid request",
		})
	}

	if err := c.Validate(user); err != nil {
		return err
	}

	data, err := services.Login(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, data)
	}

	return c.JSON(http.StatusOK, "berhasil login")
}
