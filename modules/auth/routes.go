package auth

import (
	"api_tinggal_nikah/modules/auth/controller"

	"github.com/labstack/echo/v4"
)

func Routes(r *echo.Group) {
	v1 := r.Group("/auth")

	v1.GET("/loginWithGoogle", controller.LoginWithGoogle)
	v1.GET("/callbackAuthGoogle", controller.CallbackAuthGoogle)
	v1.POST("/login", controller.Login)
	v1.POST("/register", controller.Register)

}
