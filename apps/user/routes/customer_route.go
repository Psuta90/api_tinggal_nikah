package routes

import (
	"api_tinggal_nikah/apps/user/controller"
	"api_tinggal_nikah/config"
	"api_tinggal_nikah/middleware"

	"github.com/casbin/casbin/v2"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func CustomerRoutes(r *echo.Group, enforcer *casbin.Enforcer) {
	v1 := r.Group("/customer")

	v1.Use(echojwt.WithConfig(config.ConfigJwt()))
	v1.Use(middleware.CasbinMiddleware(enforcer))

	v1.POST("/addWedding", controller.AddWedding)
	v1.PATCH("/updateWedding", controller.UpdateWedding)
	v1.POST("/upload", controller.UploadFile)
	v1.GET("/getWedding", controller.GetWedding)
	v1.DELETE("/deleteWedding", controller.DeleteWedding)
	v1.GET("/userHasPackage", controller.GetUserPackage)
	v1.GET("/getMusic", controller.GetUserMusic)
	// v1.POST("/testnats", controller.TestNats)
}
