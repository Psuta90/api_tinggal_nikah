package controller

import (
	"api_tinggal_nikah/apps/user/dto"
	"api_tinggal_nikah/apps/user/services"
	"api_tinggal_nikah/utils"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AddWedding(c echo.Context) error {

	data := dto.AddWeddingDto{}

	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	filesHalamanUtama := form.File["gallery_halaman_utama"]
	filesGalleryPhoto := form.File["gallery_photo"]

	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusOK, err.Error())
	}

	mps := new(dto.DataMempelai)
	acs := new([]dto.DataAcara)
	lss := new([]dto.DataLoveStory)
	gds := new([]dto.DataGiftDigital)
	gbs := new([]dto.DataGuestBook)

	if err := json.Unmarshal([]byte(data.MempelaiJSONSTR), &mps); err != nil {
		fmt.Println("Error:", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := json.Unmarshal([]byte(data.AcaraJSONSTR), &acs); err != nil {
		fmt.Println("Error:", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := json.Unmarshal([]byte(data.LoveStoryJSONSTR), &lss); err != nil {
		fmt.Println("Error:", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := json.Unmarshal([]byte(data.GiftDigitalJSONSTR), &gds); err != nil {
		fmt.Println("Error:", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := json.Unmarshal([]byte(data.GusetBookJSONSTR), &gbs); err != nil {
		fmt.Println("Error:", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	wjs := &dto.AddWeddingJSON{
		Acara:         *acs,
		Mempelai:      *mps,
		LoveStory:     *lss,
		GiftDigital:   *gds,
		GuestBook:     *gbs,
		Subdomain:     data.Subdomain,
		PremiumDomain: data.PremiumDomain,
	}

	wjs.HalamanUtamaGallery = append(wjs.HalamanUtamaGallery, filesHalamanUtama...)
	wjs.GalleryPhotos = append(wjs.GalleryPhotos, filesGalleryPhoto...)

	if err := utils.Validation(c, wjs); err != nil {
		return err
	}

	ServiceAddWedding, err := services.AddWeddingService(wjs)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	return c.JSON(http.StatusOK, ServiceAddWedding)

}
