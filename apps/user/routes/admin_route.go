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
	v1.PATCH("/updatePackages/:id", controller.UpdatePackages)
	v1.PATCH("/updatePackagesCategorys/:id", controller.UpdatePackagesCategorys)
	v1.DELETE("/deletePackages/:id", controller.DeletePackages)
	v1.DELETE("/deletePackagesCategorys/:id", controller.DeletePackageCategory)
	v1.POST("/addtypeTemplate", controller.AddTypeTemplate)
	v1.PATCH("/updatetypeTemplate/:id", controller.UpdateTypeTemplate)
	v1.POST("/addTemplateMaster", controller.AddTemplateMaster)
	v1.PATCH("/updateTemplateMaster/:id", controller.UpdateTemplateMaster)
}
