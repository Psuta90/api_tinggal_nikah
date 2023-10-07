package services

import (
	"api_tinggal_nikah/db"
	"api_tinggal_nikah/models"
	"api_tinggal_nikah/repository"
	"api_tinggal_nikah/utils"

	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
)

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

func GetAllPackagesServices(c echo.Context) error {

	conn := db.GetDB()
	PackageCategoryRepo := repository.NewPackageCategoryRepository(conn)

	data, err := PackageCategoryRepo.GetAllPackageCategory()
	if err != nil {
		return utils.NewAPIResponse(c).Error(0, "gagal mendapatkan data package category", err)
	}

	return utils.NewAPIResponse(c).Success(0, "api untuk get all packages", data)
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
