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
