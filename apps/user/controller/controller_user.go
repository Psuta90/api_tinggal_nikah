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

	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusOK, err.Error())
	}

	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	filesHalamanUtama := form.File["gallery_halaman_utama"]
	filesGalleryPhoto := form.File["gallery_photo"]

	mps := new(dto.DataMempelai)
	acs := new([]dto.DataAcara)
	lss := new([]dto.DataLoveStory)
	gds := new([]dto.DataGiftDigital)
	gbs := new([]dto.DataGuestBook)

	if err := json.Unmarshal([]byte(data.MempelaiJSONSTR), &mps); err != nil {
		fmt.Println("Error:", err)
		return utils.NewAPIResponse(c).Error(0, err.Error(), nil)
	}

	if err := json.Unmarshal([]byte(data.AcaraJSONSTR), &acs); err != nil {
		fmt.Println("Error:", err)
		return utils.NewAPIResponse(c).Error(0, err.Error(), nil)
	}

	if err := json.Unmarshal([]byte(data.LoveStoryJSONSTR), &lss); err != nil {
		fmt.Println("Error:", err)
		return utils.NewAPIResponse(c).Error(0, err.Error(), nil)
	}

	if err := json.Unmarshal([]byte(data.GiftDigitalJSONSTR), &gds); err != nil {
		fmt.Println("Error:", err)
		return utils.NewAPIResponse(c).Error(0, err.Error(), nil)
	}

	if err := json.Unmarshal([]byte(data.GusetBookJSONSTR), &gbs); err != nil {
		fmt.Println("Error:", err)
		return utils.NewAPIResponse(c).Error(0, err.Error(), nil)
	}

	wjs := &dto.AddWeddingJSON{
		Acara:               *acs,
		Mempelai:            *mps,
		LoveStory:           *lss,
		GiftDigital:         *gds,
		GuestBook:           *gbs,
		Subdomain:           data.Subdomain,
		PremiumDomain:       data.PremiumDomain,
		HalamanUtamaGallery: filesHalamanUtama,
		GalleryPhotos:       filesGalleryPhoto,
	}

	if err := utils.Validation(c, wjs); err != nil {
		return err
	}

	return services.AddWeddingService(c, wjs)

}

func UpdateWedding(c echo.Context) error {

	data := dto.UpdateWeddingDto{}

	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusOK, err.Error())
	}

	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	filesHalamanUtama := form.File["gallery_halaman_utama"]
	filesGalleryPhoto := form.File["gallery_photo"]

	uhu := new([]dto.UpdateDataGalleryPhotos)
	ugp := new([]dto.UpdateDataGalleryPhotos)
	umps := new(dto.UpdateDataMempelai)
	uacs := new([]dto.UpdateDataAcara)
	ulss := new([]dto.UpdateDataLoveStory)
	ugds := new([]dto.UpdateDataGiftDigital)
	ugbs := new([]dto.UpdateDataGuestBook)

	if err := json.Unmarshal([]byte(data.WeddingDto.MempelaiJSONSTR), &umps); err != nil {
		fmt.Println("Error:", err)
		return utils.NewAPIResponse(c).Error(0, err.Error()+"in Mempelai", nil)
	}

	if err := json.Unmarshal([]byte(data.WeddingDto.AcaraJSONSTR), &uacs); err != nil {
		fmt.Println("Error:", err)
		return utils.NewAPIResponse(c).Error(0, err.Error()+"in Acara", nil)
	}

	if err := json.Unmarshal([]byte(data.WeddingDto.LoveStoryJSONSTR), &ulss); err != nil {
		fmt.Println("Error:", err)
		return utils.NewAPIResponse(c).Error(0, err.Error()+"in LoveStory", nil)
	}

	if err := json.Unmarshal([]byte(data.WeddingDto.GiftDigitalJSONSTR), &ugds); err != nil {
		fmt.Println("Error:", err)
		return utils.NewAPIResponse(c).Error(0, err.Error()+"in GiftDigital", nil)
	}

	if err := json.Unmarshal([]byte(data.WeddingDto.GusetBookJSONSTR), &ugbs); err != nil {
		fmt.Println("Error:", err)
		return utils.NewAPIResponse(c).Error(0, err.Error()+"in GuestBook", nil)
	}

	if err := json.Unmarshal([]byte(data.HalamanUtamaJSONSTR), &uhu); err != nil {
		fmt.Println("Error:", err)
		return utils.NewAPIResponse(c).Error(0, err.Error()+"in HalamanUtamaJson", nil)
	}

	if err := json.Unmarshal([]byte(data.GalleryPhotoJSONSTR), &ugp); err != nil {
		fmt.Println("Error:", err)
		return utils.NewAPIResponse(c).Error(0, err.Error()+"in GalleryPhotoJson", nil)
	}

	uwjs := &dto.UpdateWeddingJSON{
		HalamanUtamaJSON:    *uhu,
		GalleryPhotoJSON:    *ugp,
		Acara:               *uacs,
		Mempelai:            *umps,
		LoveStory:           *ulss,
		GiftDigital:         *ugds,
		GuestBook:           *ugbs,
		Subdomain:           data.WeddingDto.Subdomain,
		PremiumDomain:       data.WeddingDto.PremiumDomain,
		HalamanUtamaGallery: filesHalamanUtama,
		GalleryPhotos:       filesGalleryPhoto,
	}

	if err := utils.Validation(c, uwjs); err != nil {
		return err
	}

	return services.UpdateWeddingService(c, uwjs)
}
