package middleware

import (
	"api_tinggal_nikah/config"
	"fmt"
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func CasbinMiddleware(enforcer *casbin.Enforcer) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			payloadJwt := c.Get("JWT").(*jwt.Token).Claims.(*config.JwtCustomClaims)
			obj := c.Request().URL.Path
			act := c.Request().Method

			if allowed, err := enforcer.Enforce(string(payloadJwt.Role), obj, act); err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusInternalServerError, "Internal Server Error")
			} else if allowed {
				return next(c)
			} else {
				return c.JSON(http.StatusForbidden, "Unauthorized")
			}
		}
	}
}
