package services

import (
	"api_tinggal_nikah/apps/user/dto"

	"github.com/labstack/echo/v4"
)

func AddWeddingService(data *dto.AddWeddingJSON) (echo.Map, error) {
	// conn := db.GetDB()
	// AcaraRepo := repository.NewAcaraRepository(conn)

	// acaras := new([]models.Acara)

	return echo.Map{"data": data}, nil
}
