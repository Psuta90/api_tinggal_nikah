package services

import (
	"api_tinggal_nikah/config"
	"api_tinggal_nikah/db"
	"api_tinggal_nikah/models"
	"api_tinggal_nikah/repository"
	"api_tinggal_nikah/utils"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type DataCallback struct {
	Token   string
	message string
}

func ServiceCallbackAuthGoogle(code string, c echo.Context) (*DataCallback, error) {
	token, err := config.GoogleOauthConfig.Exchange(c.Request().Context(), code)
	if err != nil {
		return &DataCallback{}, err
	}

	client := config.GoogleOauthConfig.Client(c.Request().Context(), token)

	userinfo, err := config.GetUserInfo(client)
	if err != nil {
		return &DataCallback{}, err
	}

	user := &models.User{
		Email:    userinfo.EmailAddresses[0].Value,
		FullName: userinfo.Names[0].DisplayName,
	}

	claims := &config.JwtCustomClaims{}

	conn := db.GetDB()

	userRepo := repository.NewUserRepository(conn)

	userRepo.BeforeCreateUser(user)

	if data, err := userRepo.CreateUser(user); err != nil {

		claims.ID = user.ID
		claims.Email = user.Email
		claims.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Hour * 24))

		tokenjwt, err := utils.GenerateToken(claims)
		if err != nil {
			return &DataCallback{
				Token:   "",
				message: "gagal generate token",
			}, errors.New("gagal generate token")
		}

		return &DataCallback{
			Token:   tokenjwt,
			message: "berhasil login",
		}, nil

	} else {

		claims.ID = data.ID
		claims.Email = data.Email
		claims.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Hour * 24))

		tokenjwt, err := utils.GenerateToken(claims)
		if err != nil {
			return &DataCallback{
				Token:   "",
				message: "gagal generate token",
			}, errors.New("gagal generate token")
		}

		return &DataCallback{
			Token:   tokenjwt,
			message: "berhasil login",
		}, nil
	}

}
