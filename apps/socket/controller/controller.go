package controller

import (
	"api_tinggal_nikah/apps/socket/services"

	"github.com/labstack/echo/v4"
)

func SubscribePaymentStatus(c echo.Context) error {
	order_id := c.Param("order_id")

	c.Response().Header().Set("Content-Type", "text/event-stream")
	c.Response().Header().Set("Cache-Control", "no-cache")
	c.Response().Header().Set("Connection", "keep-alive")

	return services.SubscribePaymentStatusServices(c, order_id)
}
