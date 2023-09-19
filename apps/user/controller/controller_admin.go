package controller

import (
	"api_tinggal_nikah/apps/user/dto"
	"api_tinggal_nikah/apps/user/services"
	"api_tinggal_nikah/utils"

	"github.com/labstack/echo/v4"
)

func AddPackages(c echo.Context) error {
	packages := new(dto.AddPackagesDto)
	if err := c.Bind(&packages); err != nil {
		return utils.NewAPIResponse(c).Error(0, err.Error(), err)
	}

	if err := c.Validate(packages); err != nil {
		return utils.NewAPIResponse(c).Error(0, err.Error(), err)
	}

	return services.AddPackagesService(c, packages)
}

func AddPackagesCategorys(c echo.Context) error {

	packagesCategory := new(dto.AddPackagesCategorysDto)
	if err := c.Bind(&packagesCategory); err != nil {
		return utils.NewAPIResponse(c).Error(0, err.Error(), err)
	}

	if err := c.Validate(packagesCategory); err != nil {
		return utils.NewAPIResponse(c).Error(0, err.Error(), err)
	}

	return services.AddPackagesCategoryService(c, packagesCategory)
}
