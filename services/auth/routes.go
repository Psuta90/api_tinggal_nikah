package auth

import (
	"api_tinggal_nikah/services/auth/controller"

	"github.com/labstack/echo/v4"
)

func Routes(r *echo.Group) {
	v1 := r.Group("/auth")

	v1.POST("/login", controller.Login)

}
