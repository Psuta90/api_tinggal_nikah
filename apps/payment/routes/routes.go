package routes

import (
	"api_tinggal_nikah/apps/payment/controller"

	"github.com/labstack/echo/v4"
)

func Routes(r *echo.Group) {
	v1 := r.Group("/Payment")

	v1.POST("/add", controller.Add)
}
