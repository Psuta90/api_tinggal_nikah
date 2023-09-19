package routes

import (
	"api_tinggal_nikah/apps/user/controller"
	"api_tinggal_nikah/config"
	"api_tinggal_nikah/middleware"

	"github.com/casbin/casbin/v2"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func AdminRoutes(r *echo.Group, enforcer *casbin.Enforcer) {
	v1 := r.Group("/admin")

	v1.Use(echojwt.WithConfig(config.ConfigJwt()))
	v1.Use(middleware.CasbinMiddleware(enforcer))

	v1.POST("/addPackages", controller.AddPackages)
	v1.POST("/addPackagesCategorys", controller.AddPackagesCategorys)
}
