package controller

import (
	"api_tinggal_nikah/apps/user/dto"
	"api_tinggal_nikah/apps/user/services"
	"api_tinggal_nikah/utils"

	"github.com/labstack/echo/v4"
)

func AddPackages(c echo.Context) error {
	packages := new(dto.AddPackagesDto)
	if err := c.Bind(packages); err != nil {
		return utils.NewAPIResponse(c).Error(0, err.Error(), err)
	}

	if err := c.Validate(packages); err != nil {
		return utils.NewAPIResponse(c).Error(0, err.Error(), err)
	}

	return services.AddPackagesService(c, packages)
}

func AddPackagesCategorys(c echo.Context) error {

	packagesCategory := new(dto.AddPackagesCategorysDto)
	if err := c.Bind(packagesCategory); err != nil {
		return utils.NewAPIResponse(c).Error(0, err.Error(), err)
	}

	if err := c.Validate(packagesCategory); err != nil {
		return utils.NewAPIResponse(c).Error(0, err.Error(), err)
	}

	return services.AddPackagesCategoryService(c, packagesCategory)
}

func UpdatePackages(c echo.Context) error {
	updatePackages := new(dto.UpdatePackagesDto)

	if err := c.Bind(updatePackages); err != nil {
		return utils.NewAPIResponse(c).Error(0, err.Error(), err)
	}

	if err := utils.Validation(c, updatePackages); err != nil {
		return err
	}

	return services.UpdatePackagesService(c, updatePackages)
}

func UpdatePackagesCategorys(c echo.Context) error {
	updatepackagescategory := new(dto.UpdatePackagesCategorysDto)

	if err := c.Bind(updatepackagescategory); err != nil {
		return utils.NewAPIResponse(c).Error(0, err.Error(), err)
	}

	if err := utils.Validation(c, updatepackagescategory); err != nil {
		return err
	}

	return services.UpdatePackagesCategoryService(c, updatepackagescategory)
}

func DeletePackages(c echo.Context) error {

	ID := new(dto.DeleteDto)

	if err := c.Bind(ID); err != nil {
		return utils.NewAPIResponse(c).Error(0, err.Error(), err)
	}

	if err := utils.Validation(c, ID); err != nil {
		return err
	}

	return services.DeletePackagesService(c, ID.ID)
}

func DeletePackageCategory(c echo.Context) error {
	ID := new(dto.DeleteDto)

	if err := c.Bind(ID); err != nil {
		return utils.NewAPIResponse(c).Error(0, err.Error(), err)
	}

	if err := utils.Validation(c, ID); err != nil {
		return err
	}

	return services.DeletePackageCategoryService(c, ID.ID)

}

func AddTypeTemplate(c echo.Context) error {
	typeTemplate := new(dto.AddTypeTemplateDto)
	if err := c.Bind(typeTemplate); err != nil {
		return utils.NewAPIResponse(c).Error(0, err.Error(), err)
	}

	if err := utils.Validation(c, typeTemplate); err != nil {
		return err
	}

	return services.AddTypeTemplateServices(c, typeTemplate)
}

func UpdateTypeTemplate(c echo.Context) error {
	typeTemplate := new(dto.UpdateTypeTemplateDto)
	if err := c.Bind(typeTemplate); err != nil {
		return utils.NewAPIResponse(c).Error(0, err.Error(), err)
	}

	if err := utils.Validation(c, typeTemplate); err != nil {
		return err
	}

	return services.UpdateTypeTemplateServices(c, typeTemplate)
}

func AddTemplateMaster(c echo.Context) error {
	TemplateMaster := new(dto.AddTemplateMasterDto)

	if err := c.Bind(TemplateMaster); err != nil {
		return utils.NewAPIResponse(c).Error(0, "gagal bind to struct", err)
	}

	if err := utils.Validation(c, TemplateMaster); err != nil {
		return err
	}

	return services.AddTemplateMasterService(c, TemplateMaster)

}

func UpdateTemplateMaster(c echo.Context) error {
	TemplateMaster := new(dto.UpdateTemplateMasterDto)

	if err := c.Bind(TemplateMaster); err != nil {
		return utils.NewAPIResponse(c).Error(0, "gagal bind to struct", err)
	}

	if err := utils.Validation(c, TemplateMaster); err != nil {
		return err
	}

	return services.UpdateTemplateMasterServices(c, TemplateMaster)

}
