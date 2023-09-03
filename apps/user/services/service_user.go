package services

import (
	"api_tinggal_nikah/apps/user/dto"
	"api_tinggal_nikah/db"
	"api_tinggal_nikah/models"
	"api_tinggal_nikah/repository"
	"strconv"

	"github.com/labstack/echo/v4"
)

func AddWeddingService(data *dto.AddWeddingJSON) (echo.Map, error) {
	conn := db.GetDB()
	AcaraRepo := repository.NewAcaraRepository(conn)

	// acaras := new([]models.Acara)
	var sliceAcara []models.Acara
	acaraChannel := make(chan models.Acara)
	for _, value := range data.Acara {
		orderID, errorParse := strconv.Atoi(value.Order)
		if errorParse != nil {
			return nil, errorParse
		}
		acaraEntity := models.Acara{
			Title:     value.Title,
			Place:     value.Place,
			Location:  value.Location,
			Orders:    uint(orderID),
			StartDate: value.StartDate,
			EndDate:   value.EndDate,
		}
		go func(valueAcara models.Acara) {
			acaraChannel <- valueAcara
		}(acaraEntity)
	}

	for i := 0; i < len(data.Acara); i++ {
		sliceAcara = append(sliceAcara, <-acaraChannel)
	}

	_, errorCreateBulkAcara := AcaraRepo.CreateAcara(sliceAcara)
	if errorCreateBulkAcara != nil {
		return echo.Map{"data": nil}, errorCreateBulkAcara
	}

	return echo.Map{"data": data}, nil
}
