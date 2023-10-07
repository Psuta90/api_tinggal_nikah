package services

import (
	"api_tinggal_nikah/apps/user/dto"
	"api_tinggal_nikah/db"
	"api_tinggal_nikah/models"
	"api_tinggal_nikah/repository"
	"api_tinggal_nikah/utils"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

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
