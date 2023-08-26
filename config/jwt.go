package config

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type JwtCustomClaims struct {
	ID    uuid.UUID
	Name  string
	Email string
	jwt.RegisteredClaims
}

func ConfigJwt() echojwt.Config {
	JWTCONFIG := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(JwtCustomClaims)
		},
		SigningKey: []byte(os.Getenv("JWT_SECRET_KEY")),
		ContextKey: "JWT",
	}

	return JWTCONFIG
}
