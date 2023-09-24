package controller

import (
	"api_tinggal_nikah/apps/payment/dto"
	"api_tinggal_nikah/apps/payment/services"
	"api_tinggal_nikah/utils"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Add(c echo.Context) error {
	orders := new(dto.RequestPaymentUserDto)

	if err := c.Bind(orders); err != nil {
		return utils.NewAPIResponse(c).Error(0, err.Error(), err)
	}

	return services.AddPaymentService(c, orders)

}

func ListPaymentChannel(c echo.Context) error {
	return services.ListPaymentChannelService(c)
}

func CallBackTripay(c echo.Context) error {
	data := new(dto.ResponseCallbackTripayDto)

	if err := c.Bind(&data); err != nil {
		fmt.Println(err.Error())
		return err
	}

	token := c.QueryParam("token")

	if err := utils.ValidateSignaturTripay(c, token); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return services.CallBackTripayService(c, data)
}

// func MakePayment() {

// 	subject := "payment"
// 	_, err := messagebroker.SubscribeMessage(subject, func(msg *nats.Msg) {
// 		log.Printf("Received message: %s", string(msg.Data))
// 	})

// 	if err != nil {
// 		log.Fatalf("Failed to subscribe to NATS: %v", err)

// 	}

// }
