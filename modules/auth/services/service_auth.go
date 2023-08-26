package services

import (
	"api_tinggal_nikah/config"
	"api_tinggal_nikah/db"
	"api_tinggal_nikah/models"
	"api_tinggal_nikah/modules/dto"
	"api_tinggal_nikah/repository"
	"api_tinggal_nikah/utils"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type DataCallback struct {
	Token   string
	Message string
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
		Role:     models.Customer,
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
				Message: "gagal generate token",
			}, errors.New("gagal generate token")
		}

		return &DataCallback{
			Token:   tokenjwt,
			Message: "berhasil login",
		}, nil

	} else {

		claims.ID = data.ID
		claims.Email = data.Email
		claims.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Hour * 24))

		tokenjwt, err := utils.GenerateToken(claims)
		if err != nil {
			return &DataCallback{
				Token:   "",
				Message: "gagal generate token",
			}, errors.New("gagal generate token")
		}

		return &DataCallback{
			Token:   tokenjwt,
			Message: "berhasil login",
		}, nil
	}

}

func Login(userdto *dto.LoginDto) (*DataCallback, error) {

	conn := db.GetDB()

	userRepo := repository.NewUserRepository(conn)

	user, err := userRepo.GetUserByEmail(userdto.Email)
	if err != nil {
		return &DataCallback{Token: "", Message: "silahkan masukan email atau password yang terdaftar"}, errors.New("silahkan masukan email atau password yang terdaftar")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userdto.Password)); err != nil {
		return &DataCallback{Token: "", Message: "silahkan masukan email atau password yang terdaftar"}, errors.New("silahkan masukan email atau password yang terdaftar")
	}

	claims := &config.JwtCustomClaims{
		ID:    user.ID,
		Name:  user.FullName,
		Email: user.Email,
	}

	token, err := utils.GenerateToken(claims)
	if err != nil {
		return &DataCallback{}, errors.New("gagal create token silahkan coba login kembali")
	}

	var data = &DataCallback{
		Token:   token,
		Message: "berhasil login",
	}

	return data, nil

}
