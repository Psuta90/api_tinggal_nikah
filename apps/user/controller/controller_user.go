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
		return c.JSON(http.StatusBadRequest, err.Error())
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
		Template:            data.Template,
	}

	if err := utils.Validation(c, wjs); err != nil {
		return err
	}

	return services.AddWeddingService(c, wjs)

}

func UpdateWedding(c echo.Context) error {

	data := dto.UpdateWeddingDto{}

	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return services.UdateWeddingService(c, &data)

}

func UploadFile(c echo.Context) error {
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	filesHalamanUtama := form.File["gallery_halaman_utama"]
	filesGalleryPhoto := form.File["gallery_photo"]

	data := &dto.UploadFileDto{
		HalamanUtamaGallery: filesHalamanUtama,
		GalleryPhotos:       filesGalleryPhoto,
	}

	if err := utils.Validation(c, data); err != nil {
		return err
	}

	return services.UploadFileService(c, data)
}

func GetWedding(c echo.Context) error {
	return services.GetWeddingService(c)
}

func DeleteWedding(c echo.Context) error {
	return utils.NewAPIResponse(c).Success(0, "route untuk delete", nil)
}

func GetUserPackage(c echo.Context) error {
	//
	return services.GetUserPackageService(c)
}

// func TestNats(c echo.Context) error {
// 	message := "Hello, NATS JetStream!"

// 	if err := messagebroker.NatsConn.Publish("payment", []byte(message)); err != nil {
// 		return c.String(http.StatusInternalServerError, "Failed to publish message to NATS JetStream")
// 	}

// 	return c.String(http.StatusOK, "Message published to NATS JetStream")
// }
