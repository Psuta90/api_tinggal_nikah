package user

import (
	"api_tinggal_nikah/config"
	"api_tinggal_nikah/middleware"
	"net/http"

	"github.com/casbin/casbin/v2"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func Routes(r *echo.Group, enforcer *casbin.Enforcer) {
	v1 := r.Group("/customer")

	v1.Use(echojwt.WithConfig(config.ConfigJwt()))
	v1.Use(middleware.CasbinMiddleware(enforcer))
	v1.GET("/listuser", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	v1.POST("/postuser", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
}
