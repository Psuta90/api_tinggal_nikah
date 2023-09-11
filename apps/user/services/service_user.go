package services

import (
	"api_tinggal_nikah/apps/user/dto"
	"api_tinggal_nikah/config"
	"api_tinggal_nikah/db"
	"api_tinggal_nikah/models"
	repository "api_tinggal_nikah/repository/wedding"
	"api_tinggal_nikah/utils"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func AddWeddingService(c echo.Context, data *dto.AddWeddingJSON) error {

	user_id := c.Get("JWT").(*jwt.Token).Claims.(*config.JwtCustomClaims).ID

	conn := db.GetDB().Begin()
	if conn.Error != nil {
		return utils.NewAPIResponse(c).Error(0, "Failed to start a transaction add wedding", nil)
	}

	AcaraRepo := repository.NewAcaraRepository(conn)
	LoveStoryRepo := repository.NewLoveStoryRepository(conn)
	GifDigitalRepo := repository.NewGiftDigitalRepository(conn)
	GuestBookRepo := repository.NewGuestBookRepository(conn)
	MepelaiPriaRepo := repository.NewMempelaiPriaRepository(conn)
	MempelaiWanitaRepo := repository.NewMempelaiWanitaRepository(conn)
	DomainRepo := repository.NewDomainRepository(conn)
	GalleryRepo := repository.NewGalleryPhotosRepository(conn)
	TemplateUserRepo := repository.NewTemplateUserRepository(conn)

	acaras := []models.Acara{}
	lovestorys := []models.LoveStory{}
	giftdigitals := []models.GiftDigital{}
	guestbook := []models.GuestBook{}
	gallery := []models.GalleryPhotos{}

	mempelaiPria := &models.MempelaiPria{
		NameAlias:  data.Mempelai.MempelaiPria.NameAlias,
		FullName:   data.Mempelai.MempelaiPria.Fullname,
		NameFather: data.Mempelai.MempelaiPria.NameFather,
		NameMother: data.Mempelai.MempelaiPria.NameMother,
		IsLeft:     data.Mempelai.MempelaiPria.IsLeft,
		UserID:     user_id,
	}

	mempelaiWanita := &models.MempelaiWanita{
		NameAlias:  data.Mempelai.MempelaiWanita.NameAlias,
		FullName:   data.Mempelai.MempelaiWanita.Fullname,
		NameFather: data.Mempelai.MempelaiWanita.NameFather,
		NameMother: data.Mempelai.MempelaiWanita.NameMother,
		IsLeft:     data.Mempelai.MempelaiWanita.IsLeft,
		UserID:     user_id,
	}

	domain := &models.Domain{
		Subdomain:     data.Subdomain,
		PremiumDomain: data.PremiumDomain,
		UserID:        user_id,
	}

	template := &models.TemplateUser{
		TemplateID: data.Template,
		UserID:     user_id,
	}

	AcaraChannel := make(chan models.Acara)
	LoveStoryChannel := make(chan models.LoveStory)
	GiftDigitalsChannel := make(chan models.GiftDigital)
	GuestBookChannel := make(chan models.GuestBook)
	GalleryChannel := make(chan models.GalleryPhotos)

	go func() {
		for _, value := range data.Acara {

			acaraEntity := models.Acara{
				Title:     value.Title,
				StartDate: value.StartDate,
				EndDate:   value.EndDate,
				Location:  value.Location,
				Place:     value.Place,
				Orders:    value.Order,
				UserID:    user_id,
			}

			AcaraChannel <- acaraEntity

		}
		defer close(AcaraChannel)

	}()
	go func() {
		for _, value := range data.LoveStory {

			lovestoryEntity := models.LoveStory{
				Title:    value.Title,
				Location: value.Location,
				Story:    value.Story,
				Orders:   value.Order,
				UserID:   user_id,
			}

			LoveStoryChannel <- lovestoryEntity

		}

		defer close(LoveStoryChannel)
	}()

	go func() {
		for _, value := range data.GiftDigital {

			no_rekening, err := strconv.Atoi(value.NoRekening)
			if err != nil {
				fmt.Println(err)
			}

			GiftDigitalEntity := models.GiftDigital{
				NoRekening:   uint(no_rekening),
				NameRekening: value.NameRekening,
				PaymentType:  value.PaymentType,
				Orders:       value.Order,
				UserID:       user_id,
			}

			GiftDigitalsChannel <- GiftDigitalEntity

		}

		defer close(GiftDigitalsChannel)
	}()

	go func() {
		for _, value := range data.GuestBook {

			for _, value2 := range value.DatasGuestBook {

				guest_bookEntity := models.GuestBook{
					Group:   value.Group,
					Name:    value2.Name,
					Phone:   value2.Phone,
					Message: value2.Message,
					Orders:  value.Order,
					UserID:  user_id,
				}

				GuestBookChannel <- guest_bookEntity
			}

		}

		defer close(GuestBookChannel)
	}()

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	path := filepath.Join(cwd, "temp_image")

	//check folder if not exist then create folder image_lp
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.Mkdir(path, 0755); err != nil {
			fmt.Println(err)
		}
	}

	go func() {
		for index, file := range data.HalamanUtamaGallery {

			dbpath := filepath.Join("temp_image", file.Filename)
			destinationPath := filepath.Join(path, file.Filename)

			// Source
			src, err := file.Open()
			if err != nil {
				fmt.Println(err)
			}
			defer src.Close()

			// Destination
			dst, err := os.Create(destinationPath)
			if err != nil {
				fmt.Println(err)
			}
			defer dst.Close()

			// Copy
			if _, err = io.Copy(dst, src); err != nil {
				fmt.Println(err)
			}

			galleryEntity := models.GalleryPhotos{
				Path:           dbpath,
				Orders:         index,
				IsGallery:      false,
				IsHalamanUtama: true,
				UserID:         user_id,
			}

			GalleryChannel <- galleryEntity

		}

		for index, file := range data.GalleryPhotos {

			dbpath := filepath.Join("temp_image", file.Filename)
			destinationPath := filepath.Join(path, file.Filename)

			// Source
			src, err := file.Open()
			if err != nil {
				fmt.Println(err)
			}
			defer src.Close()

			// Destination
			dst, err := os.Create(destinationPath)
			if err != nil {
				fmt.Println(err)
			}
			defer dst.Close()

			// Copy
			if _, err = io.Copy(dst, src); err != nil {
				fmt.Println(err)
			}

			galleryEntity := models.GalleryPhotos{
				Path:           dbpath,
				Orders:         index,
				IsGallery:      true,
				IsHalamanUtama: false,
				UserID:         user_id,
			}

			GalleryChannel <- galleryEntity

		}

		defer close(GalleryChannel)
	}()

	for dataAcaraChannel := range AcaraChannel {
		acaras = append(acaras, dataAcaraChannel)
	}

	for dataLoveStoryChannel := range LoveStoryChannel {
		lovestorys = append(lovestorys, dataLoveStoryChannel)
	}

	for dataGiftDigitalsChannel := range GiftDigitalsChannel {
		giftdigitals = append(giftdigitals, dataGiftDigitalsChannel)
	}

	for dataGuestBookChannel := range GuestBookChannel {
		guestbook = append(guestbook, dataGuestBookChannel)
	}

	for dataGalleryPhotosChannel := range GalleryChannel {
		gallery = append(gallery, dataGalleryPhotosChannel)
	}

	if err := AcaraRepo.CreateAcara(&acaras); err != nil {
		conn.Rollback()
		fmt.Println(err)
		return utils.NewAPIResponse(c).FailedInsertDB(0, "gagal pada saat insert data acara", nil)
	}

	if err := LoveStoryRepo.CreateLoveStory(&lovestorys); err != nil {
		conn.Rollback()
		fmt.Println(err)
		return utils.NewAPIResponse(c).FailedInsertDB(0, "gagal pada saat insert data lovestorys", nil)
	}

	if err := GifDigitalRepo.CreateGiftDigital(&giftdigitals); err != nil {
		conn.Rollback()
		fmt.Println(err)
		return utils.NewAPIResponse(c).FailedInsertDB(0, "gagal pada saat insert data giftdigitals", nil)
	}

	if err := GuestBookRepo.CreateGuestBook(&guestbook); err != nil {
		conn.Rollback()
		fmt.Println(err)
		return utils.NewAPIResponse(c).FailedInsertDB(0, "gagal pada saat insert data guestbook", nil)
	}

	if err := MepelaiPriaRepo.CreateMempelaiPria(mempelaiPria); err != nil {
		conn.Rollback()
		fmt.Println(err)
		return utils.NewAPIResponse(c).FailedInsertDB(0, "gagal pada saat insert data mempelaiPria", nil)
	}

	if err := MempelaiWanitaRepo.CreateMempelaiWanita(mempelaiWanita); err != nil {
		conn.Rollback()
		fmt.Println(err)
		return utils.NewAPIResponse(c).FailedInsertDB(0, "gagal pada saat insert data mempelaiWanita", nil)
	}

	if err := DomainRepo.CreateDomain(domain); err != nil {
		conn.Rollback()
		fmt.Println(err)
		return utils.NewAPIResponse(c).FailedInsertDB(0, "gagal pada saat insert data domain", nil)
	}

	if err := TemplateUserRepo.CreateTemplateUser(template); err != nil {
		conn.Rollback()
		fmt.Println(err)
		return utils.NewAPIResponse(c).FailedInsertDB(0, "gagal pada saat insert template", nil)
	}

	if err := GalleryRepo.CreateGalleryPhotos(&gallery); err != nil {
		conn.Rollback()
		fmt.Println(err)
		for _, value := range gallery {
			err := os.Remove(filepath.Join(cwd, value.Path))
			if err != nil {
				fmt.Println("error delete files", err.Error())
			}
		}

		return utils.NewAPIResponse(c).FailedInsertDB(0, "gagal pada saat insert data acara", err)

	}

	if err := conn.Commit().Error; err != nil {
		return utils.NewAPIResponse(c).Error(0, "Failed to commit transaction add wedding", nil)
	}

	return utils.NewAPIResponse(c).Success(0, "berhasil melakukan insert data", nil)
}

func UploadFileService(c echo.Context, data *dto.UploadFileDto) error {

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	path := filepath.Join(cwd, "temp_image")

	//check folder if not exist then create folder image_lp
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.Mkdir(path, 0755); err != nil {
			fmt.Println(err)
		}
	}

	for _, file := range data.GalleryPhotos {

		destinationPath := filepath.Join(path, file.Filename)

		// Source
		src, err := file.Open()
		if err != nil {
			return utils.NewAPIResponse(c).Error(0, "", err)
		}
		defer src.Close()

		// Destination
		dst, err := os.Create(destinationPath)
		if err != nil {
			return utils.NewAPIResponse(c).Error(0, "", err)
		}
		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			return utils.NewAPIResponse(c).Error(0, "", err)
		}

	}

	for _, file := range data.HalamanUtamaGallery {

		destinationPath := filepath.Join(path, file.Filename)

		// Source
		src, err := file.Open()
		if err != nil {
			return utils.NewAPIResponse(c).Error(0, "", err)
		}
		defer src.Close()

		// Destination
		dst, err := os.Create(destinationPath)
		if err != nil {
			return utils.NewAPIResponse(c).Error(0, "", err)
		}
		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			return utils.NewAPIResponse(c).Error(0, "", err)
		}

	}

	return utils.NewAPIResponse(c).Success(0, "berhasil upload data", nil)
}

func UdateWeddingService(c echo.Context, data *dto.UpdateWeddingDto) error {

	conn := db.GetDB().Begin()
	DomainRepo := repository.NewDomainRepository(conn)
	MempelaiPriaRepo := repository.NewMempelaiPriaRepository(conn)
	MempelaiWanitaRepo := repository.NewMempelaiWanitaRepository(conn)
	GalleryPhotosRepo := repository.NewGalleryPhotosRepository(conn)
	AcaraRepo := repository.NewAcaraRepository(conn)
	LoveStoryRepo := repository.NewLoveStoryRepository(conn)
	GifDigitalRepo := repository.NewGiftDigitalRepository(conn)
	GuestBookRepo := repository.NewGuestBookRepository(conn)
	TemplateUserRepo := repository.NewTemplateUserRepository(conn)

	MempelaiPria := &models.MempelaiPria{
		ID:         data.Mempelai.MempelaiPria.ID,
		NameAlias:  data.Mempelai.MempelaiPria.NameAlias,
		FullName:   data.Mempelai.MempelaiPria.Fullname,
		NameFather: data.Mempelai.MempelaiPria.NameFather,
		NameMother: data.Mempelai.MempelaiPria.NameMother,
		IsLeft:     data.Mempelai.MempelaiPria.IsLeft,
	}

	MempelaiWanita := &models.MempelaiWanita{
		ID:         data.Mempelai.MempelaiWanita.ID,
		NameAlias:  data.Mempelai.MempelaiWanita.NameAlias,
		FullName:   data.Mempelai.MempelaiWanita.Fullname,
		NameFather: data.Mempelai.MempelaiWanita.NameFather,
		NameMother: data.Mempelai.MempelaiWanita.NameMother,
		IsLeft:     data.Mempelai.MempelaiWanita.IsLeft,
	}

	domain := &models.Domain{
		ID:            data.Domain.ID,
		Subdomain:     data.Domain.Subdomain,
		PremiumDomain: data.Domain.PremiumDomain,
	}

	template := &models.TemplateUser{
		ID:         data.Template.ID,
		TemplateID: data.Template.TemplateID,
	}

	gallery := make(chan error)
	acara := make(chan error)
	lovestory := make(chan error)
	giftdigital := make(chan error)
	guestbook := make(chan error)

	go func() {
		if data.Gallery != nil {
			for _, value := range data.Gallery {

				cwd, err := os.Getwd()
				if err != nil {
					fmt.Println(err)
				}

				path := filepath.Join(cwd, "temp_image")

				GalleryEntity := &models.GalleryPhotos{
					ID:             value.ID,
					Path:           filepath.Join(path, value.Filename),
					Orders:         value.Order,
					IsGallery:      value.IsGallery,
					IsHalamanUtama: value.IsHalamanUtama,
				}

				GalleryPhotosRepo.UpdateGalleryPhotos(*GalleryEntity, gallery)

			}
		}

		defer close(gallery)
	}()

	go func() {
		if data.Acara != nil {
			for _, value := range data.Acara {

				AcaraEntity := &models.Acara{
					ID:        value.ID,
					Title:     value.Title,
					StartDate: value.StartDate,
					EndDate:   value.EndDate,
					Place:     value.Place,
					Location:  value.Location,
					Orders:    value.Order,
				}

				AcaraRepo.UpdateAcara(AcaraEntity, acara)

			}
		}
		defer close(acara)
	}()

	go func() {
		if data.LoveStory != nil {
			for _, value := range data.LoveStory {

				LoveStoryEntity := &models.LoveStory{
					ID:       value.ID,
					Title:    value.Title,
					Location: value.Location,
					Story:    value.Story,
					Orders:   value.Order,
				}

				LoveStoryRepo.UpdateLoveStory(LoveStoryEntity, lovestory)

			}
		}

		defer close(lovestory)
	}()

	go func() {
		if data.GiftDigital != nil {
			for _, value := range data.GiftDigital {
				no_rekening, err := strconv.Atoi(value.NoRekening)
				if err != nil {
					fmt.Println(err)
				}

				GiftDigitalEntity := &models.GiftDigital{
					ID:           value.ID,
					NoRekening:   uint(no_rekening),
					PaymentType:  value.PaymentType,
					NameRekening: value.NameRekening,
					Orders:       value.Order,
				}

				GifDigitalRepo.UpdateGiftDigital(GiftDigitalEntity, giftdigital)

			}
		}

		defer close(giftdigital)
	}()

	go func() {

		if data.GuestBook != nil {
			for _, value := range data.GuestBook {

				for _, value2 := range value.DatasGuestBook {

					guest_bookEntity := models.GuestBook{
						ID:      value2.ID,
						Group:   value.Group,
						Name:    value2.Name,
						Phone:   value2.Phone,
						Message: value2.Message,
						Orders:  value.Order,
					}

					GuestBookRepo.UpdateGuestBook(&guest_bookEntity, guestbook)

				}

			}
		}

		defer close(guestbook)

	}()

	for errGallery := range gallery {
		if errGallery != nil {
			conn.Rollback()
			return utils.NewAPIResponse(c).Error(0, "gagal update HalamanUtama Atau Album Photo", errGallery.Error())
		}
	}

	for errAcara := range acara {
		if errAcara != nil {
			conn.Rollback()
			return utils.NewAPIResponse(c).Error(0, "gagal update acara", errAcara.Error())
		}
	}

	for errLovestory := range lovestory {
		if errLovestory != nil {
			conn.Rollback()
			return utils.NewAPIResponse(c).Error(0, "gagal update lovestory", errLovestory.Error())
		}
	}

	for errGiftDigital := range giftdigital {
		if errGiftDigital != nil {
			conn.Rollback()
			return utils.NewAPIResponse(c).Error(0, "gagal update giftDigital", errGiftDigital.Error())
		}
	}

	for errGuestBook := range guestbook {
		if errGuestBook != nil {
			conn.Rollback()
			return utils.NewAPIResponse(c).Error(0, "gagal update GuestBook", errGuestBook.Error())
		}
	}

	if data.Template != (dto.UpdateTemplateUser{}) {
		if err := TemplateUserRepo.UpdateTemplateUser(template); err != nil {
			conn.Rollback()
			return utils.NewAPIResponse(c).Error(0, "gagal update Template", err)
		}
	}

	if data.Mempelai.MempelaiPria != (dto.UpdateMempelai{}) {
		if err := MempelaiPriaRepo.UpdateMempelaiPria(MempelaiPria); err != nil {
			conn.Rollback()
			return utils.NewAPIResponse(c).Error(0, "gagal update MempelaiPria", err)
		}
	}

	if data.Mempelai.MempelaiWanita != (dto.UpdateMempelai{}) {
		if err := MempelaiWanitaRepo.UpdateMempelaiWanita(MempelaiWanita); err != nil {
			conn.Rollback()
			return utils.NewAPIResponse(c).Error(0, "gagal update MempelaiWanita", err)
		}
	}

	if data.Domain != (dto.UpdateDomain{}) {
		if err := DomainRepo.UpdateDomain(domain); err != nil {
			conn.Rollback()
			return utils.NewAPIResponse(c).Error(0, "gagal update domain", err)
		}
	}

	if err := conn.Commit().Error; err != nil {
		return utils.NewAPIResponse(c).Error(0, "Failed to commit transaction update wedding", nil)
	}

	return utils.NewAPIResponse(c).Success(0, "", data)

}
