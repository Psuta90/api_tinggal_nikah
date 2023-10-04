package routes

import (
	"api_tinggal_nikah/apps/user/controller"

	"github.com/labstack/echo/v4"
)

func GlobalRoutes(r *echo.Group) {
	v1 := r.Group("/global")
	v1.GET("/alltemplates", controller.GetAllTemplates)
	v1.GET("/allpackages", controller.GetAllPackages)
	v1.PATCH("/updateRsvp/:id", controller.UpdateRsvp)
	v1.GET("/getGuest/:name", controller.GetGuessByName)
}
