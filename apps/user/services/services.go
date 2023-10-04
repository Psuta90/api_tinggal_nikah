package services

import (
	"api_tinggal_nikah/apps/user/dto"
	"api_tinggal_nikah/config"
	"api_tinggal_nikah/db"
	"api_tinggal_nikah/models"
	"api_tinggal_nikah/repository"
	"api_tinggal_nikah/utils"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
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
	MusicUserRepo := repository.NewMusicUserReporsitory(conn)

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

	music := &models.MusicUser{
		UserID:        user_id,
		MusicMasterID: data.Music,
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

	if err := MusicUserRepo.Add(music); err != nil {
		conn.Rollback()
		fmt.Println(err)
		return utils.NewAPIResponse(c).FailedInsertDB(0, "gagal pada saat insert music", nil)
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

func GetWeddingService(c echo.Context) error {

	conn := db.GetDB()
	user_id := c.Get("JWT").(*jwt.Token).Claims.(*config.JwtCustomClaims).ID

	UserRepo := repository.NewUserRepository(conn)

	data, err := UserRepo.GetWeddingUser(user_id)
	if err != nil {
		return utils.NewAPIResponse(c).Error(0, "error get data", data)
	}

	nm, err := utils.StructToMap(data)
	if err != nil {
		return err
	}

	datares := lo.OmitByKeys(nm, []string{"Password", "Role"})

	return utils.NewAPIResponse(c).Success(0, "Berhasil Mendapatkan Data", &datares)
}

func GetAllTemplatesService(c echo.Context) error {

	conn := db.GetDB()
	TypeTemplateRepo := repository.NewTemplateTypeRepository(conn)

	data, err := TypeTemplateRepo.GetAllTemplateType()
	if err != nil {
		return utils.NewAPIResponse(c).Error(0, "data tidak di temukan", nil)
	}

	datares := lo.Map(data, func(typeTemplate models.TypeTemplate, index int) map[string]interface{} {

		OmitedTemplateMaster := lo.Map(typeTemplate.TemplateMaster, func(tm models.TemplateMaster, index2 int) map[string]interface{} {

			nm, _ := utils.StructToMap(tm)
			omit := lo.OmitByKeys(nm, []string{"TemplateUser"})
			return omit
		})

		data := echo.Map{
			"ID":             typeTemplate.ID,
			"Name":           typeTemplate.Name,
			"TemplateMaster": OmitedTemplateMaster,
		}

		return data
	})

	return utils.NewAPIResponse(c).Success(0, "success", datares)

}

func AddPackagesService(c echo.Context, data *dto.AddPackagesDto) error {

	conn := db.GetDB()

	packages := &models.Package{
		GuestSize:         data.GuestSize,
		GallerySize:       data.GallerySize,
		VideoSize:         data.VideoSize,
		RSVP:              data.RSVP,
		LocationLink:      data.LocationLink,
		Story:             data.Story,
		GiftDigital:       data.GiftDigital,
		Music:             data.Music,
		PackageCategoryID: data.PackageCategoryID,
	}

	PackageRepo := repository.NewPackagesRepository(conn)

	if err := PackageRepo.CreatePackage(packages); err != nil {
		return utils.NewAPIResponse(c).Error(0, "gagal create package", err)
	}

	return utils.NewAPIResponse(c).Success(0, "success create packages ", data)
}

func AddPackagesCategoryService(c echo.Context, data *dto.AddPackagesCategorysDto) error {

	conn := db.GetDB()

	package_category := &models.PackageCategory{
		Name:               data.Name,
		Price:              data.Price,
		DiscountPercentage: data.DiscountPercentage,
		ActiveDays:         data.ActiveDays,
	}

	PackageCategoryRepo := repository.NewPackageCategoryRepository(conn)

	if err := PackageCategoryRepo.CreatePackageCategory(package_category); err != nil {
		return utils.NewAPIResponse(c).Error(0, "gagal create package category", err)
	}

	return utils.NewAPIResponse(c).Success(0, "success create packages category", data)
}

func UpdatePackagesService(c echo.Context, data *dto.UpdatePackagesDto) error {
	conn := db.GetDB()

	packages := &models.Package{
		ID:                data.ID,
		GuestSize:         data.GuestSize,
		GallerySize:       data.GallerySize,
		VideoSize:         data.VideoSize,
		RSVP:              data.RSVP,
		LocationLink:      data.LocationLink,
		Story:             data.Story,
		GiftDigital:       data.GiftDigital,
		Music:             data.Music,
		PackageCategoryID: data.PackageCategoryID,
	}

	PackageRepo := repository.NewPackagesRepository(conn)

	if err := PackageRepo.UpdatePackage(packages); err != nil {
		return utils.NewAPIResponse(c).Error(0, "gagal update package", err)
	}

	return utils.NewAPIResponse(c).Success(0, "success update packages", data)
}

func UpdatePackagesCategoryService(c echo.Context, data *dto.UpdatePackagesCategorysDto) error {

	conn := db.GetDB()

	PackageCategory := &models.PackageCategory{
		ID:                 data.ID,
		Name:               data.Name,
		Price:              data.Price,
		DiscountPercentage: data.DiscountPercentage,
		ActiveDays:         data.ActiveDays,
	}

	PackageCategoryRepo := repository.NewPackageCategoryRepository(conn)

	if err := PackageCategoryRepo.UpdatePackageCategory(PackageCategory); err != nil {
		return utils.NewAPIResponse(c).Error(0, "gagal update package", err)
	}

	return utils.NewAPIResponse(c).Success(0, "success update packages category", data)

}

func DeletePackagesService(c echo.Context, id uuid.UUID) error {

	conn := db.GetDB()

	packages := &models.Package{
		ID: id,
	}

	PackageRepo := repository.NewPackagesRepository(conn)

	if err := PackageRepo.DeletePackage(packages); err != nil {
		return utils.NewAPIResponse(c).Error(0, "gagal deleted package", err)
	}

	return utils.NewAPIResponse(c).Success(0, "success delete packages", id)
}

func DeletePackageCategoryService(c echo.Context, id uuid.UUID) error {

	conn := db.GetDB()

	PackageCategory := &models.PackageCategory{
		ID: id,
	}

	PackageCategoryRepo := repository.NewPackageCategoryRepository(conn)

	if err := PackageCategoryRepo.DeletePackageCategory(PackageCategory); err != nil {
		return utils.NewAPIResponse(c).Error(0, "gagal deleted package category", err)
	}

	return utils.NewAPIResponse(c).Success(0, "success delete packages", id)
}

func GetAllPackagesServices(c echo.Context) error {

	conn := db.GetDB()
	PackageCategoryRepo := repository.NewPackageCategoryRepository(conn)

	data, err := PackageCategoryRepo.GetAllPackageCategory()
	if err != nil {
		return utils.NewAPIResponse(c).Error(0, "gagal mendapatkan data package category", err)
	}

	return utils.NewAPIResponse(c).Success(0, "api untuk get all packages", data)
}

func AddTypeTemplateServices(c echo.Context, data *dto.AddTypeTemplateDto) error {
	conn := db.GetDB()
	TypeTemplateRepo := repository.NewTemplateTypeRepository(conn)

	typeTemplate := &models.TypeTemplate{
		Name: data.Name,
	}

	if err := TypeTemplateRepo.AddTypeTemplate(typeTemplate); err != nil {
		return utils.NewAPIResponse(c).Error(0, "gagal insert template type", err)
	}

	return utils.NewAPIResponse(c).Success(0, "success add template type", nil)
}

func UpdateTypeTemplateServices(c echo.Context, data *dto.UpdateTypeTemplateDto) error {
	conn := db.GetDB()
	TypeTemplateRepo := repository.NewTemplateTypeRepository(conn)

	typeTemplate := &models.TypeTemplate{
		Name: data.Name,
		ID:   data.ID,
	}

	if err := TypeTemplateRepo.UpdateTypeTemplate(typeTemplate); err != nil {
		return utils.NewAPIResponse(c).Error(0, "gagal insert template type", err)
	}

	return utils.NewAPIResponse(c).Success(0, "success add template type", nil)
}

func AddTemplateMasterService(c echo.Context, data *dto.AddTemplateMasterDto) error {

	conn := db.GetDB()
	TemplateMasterRepo := repository.NewTemplateMasterRepository(conn)

	TemplateMaster := &models.TemplateMaster{
		Name:           data.Name,
		Css:            data.Css,
		TypeTemplateID: data.TypeTemplateID,
	}

	if err := TemplateMasterRepo.CreateTemplateMaster(TemplateMaster); err != nil {
		return utils.NewAPIResponse(c).Error(0, "gagal insert template master", err)
	}

	return utils.NewAPIResponse(c).Success(0, "berhasil menambahkan add template master", nil)
}

func UpdateTemplateMasterServices(c echo.Context, data *dto.UpdateTemplateMasterDto) error {
	conn := db.GetDB()
	TemplateMasterRepo := repository.NewTemplateMasterRepository(conn)
	TemplateMaster := &models.TemplateMaster{
		ID:             data.ID,
		Name:           data.Name,
		Css:            data.Css,
		TypeTemplateID: data.TypeTemplateID,
	}

	if err := TemplateMasterRepo.UpdateTemplateMaster(TemplateMaster); err != nil {
		return utils.NewAPIResponse(c).Error(0, "gagal insert template type", err)
	}

	return utils.NewAPIResponse(c).Success(0, "success add template type", nil)
}

func GetUserPackageService(c echo.Context) error {

	user_id := c.Get("JWT").(*jwt.Token).Claims.(*config.JwtCustomClaims).ID

	conn := db.GetDB()
	UserPackaRepo := repository.NewUserPackageRepository(conn)

	data, err := UserPackaRepo.GetByUserID(user_id)
	if err != nil {
		return utils.NewAPIResponse(c).Error(0, "gagal get userpackage by user id", err)
	}

	var datares []interface{}

	for _, items := range data {
		if time.Now().Before(items.EndDate) {
			nm, _ := utils.StructToMap(items)
			omit := lo.OmitByKeys(nm, []string{"UserTransaction"})

			datares = append(datares, omit)
		}
	}

	return utils.NewAPIResponse(c).Success(0, "success", datares)
}

func AddMusicMasterService(c echo.Context, data *dto.AddMusicMasterDto) error {

	conn := db.GetDB()
	MusicMasterRepo := repository.NewMusicMasterRepository(conn)
	MusicMaster := []models.MusicMaster{}

	MusicMasterChan := make(chan models.MusicMaster)

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	path := filepath.Join(cwd, "temp_music")

	//check folder if not exist then create folder image_lp
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.Mkdir(path, 0755); err != nil {
			fmt.Println(err)
		}
	}

	go func() {
		for _, file := range data.Files {

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

			MusicMasterEntity := &models.MusicMaster{
				Path: destinationPath,
				Name: file.Filename,
			}

			MusicMasterChan <- *MusicMasterEntity

		}

		defer close(MusicMasterChan)

	}()

	for dataMusicMasterChan := range MusicMasterChan {
		MusicMaster = append(MusicMaster, dataMusicMasterChan)
	}

	if err := MusicMasterRepo.Add(&MusicMaster); err != nil {
		return utils.NewAPIResponse(c).Error(0, "gagal insert music", err)
	}

	return utils.NewAPIResponse(c).Success(0, "Behasil Insert Music", MusicMaster)
}

func UpdateMusicMasterServices(c echo.Context, data *dto.UpdateMusic) error {

	conn := db.GetDB().Begin()
	MusicMasterRepo := repository.NewMusicMasterRepository(conn)

	MusicChan := make(chan error)

	go func() {
		if data.Music != nil {

			cwd, err := os.Getwd()
			if err != nil {
				fmt.Println(err)
			}

			path := filepath.Join(cwd, "temp_music")

			for _, value := range data.Music {

				MusicMasterEntity := &models.MusicMaster{
					ID:   value.ID,
					Name: value.Name,
					Path: filepath.Join(path, value.Filename),
				}

				MusicMasterRepo.Update(MusicMasterEntity, MusicChan)
			}

			defer close(MusicChan)
		}
	}()

	for errMusicChan := range MusicChan {
		if errMusicChan != nil {
			conn.Rollback()
			return utils.NewAPIResponse(c).Error(0, "gagal update music", errMusicChan)
		}
	}

	if err := conn.Commit().Error; err != nil {
		return utils.NewAPIResponse(c).Error(0, "Failed to commit update music", nil)
	}

	return utils.NewAPIResponse(c).Success(0, "berhasil Update Music", data)
}

func UploadMusicService(c echo.Context, data []*multipart.FileHeader) error {

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	path := filepath.Join(cwd, "temp_music")

	for _, file := range data {

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

	}
	return utils.NewAPIResponse(c).Success(0, "upload music berhasil", data)
}

func GetAllMusicService(c echo.Context) error {
	conn := db.GetDB()
	MusicMasterRepo := repository.NewMusicMasterRepository(conn)

	data, err := MusicMasterRepo.FindAll()
	if err != nil {
		return utils.NewAPIResponse(c).Error(0, "gagal mendapatkan data music", err)
	}

	return utils.NewAPIResponse(c).Success(0, "succsess get data music", data)

}

func UpdateRsvpServices(c echo.Context, data *dto.UpdateRsvpGuestBookDto) error {
	conn := db.GetDB()
	GuestBookRepo := repository.NewGuestBookRepository(conn)

	GuestBookChan := make(chan error)

	guestBok := &models.GuestBook{
		ID:               data.ID,
		Attendance:       data.Attendences,
		MessageFromGuess: data.MessageFromGuest,
	}

	go func() {
		GuestBookRepo.UpdateGuestBook(guestBok, GuestBookChan)
		defer close(GuestBookChan)
	}()

	for errGuestBook := range GuestBookChan {
		if errGuestBook != nil {
			return utils.NewAPIResponse(c).Error(0, errGuestBook.Error(), errGuestBook)
		}
	}

	return utils.NewAPIResponse(c).Success(0, "api untuk rsvp", data)
}

func GetGuestServices(c echo.Context, name string) error {
	conn := db.GetDB().Begin()
	GuestBookRepo := repository.NewGuestBookRepository(conn)

	data, err := GuestBookRepo.FindByNameGuestBook(name)
	if err != nil {
		return utils.NewAPIResponse(c).Error(0, "gagal mendapatkan guestbook by name", err)
	}

	return utils.NewAPIResponse(c).Success(0, "success", data)
}
