package controller

import (
	"api_tinggal_nikah/config"
	"api_tinggal_nikah/modules/auth/services"
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
	return c.JSON(http.StatusOK, "berhasil login")
}
