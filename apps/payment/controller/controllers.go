package controller

import (
	"api_tinggal_nikah/apps/payment/dto"
	"api_tinggal_nikah/apps/payment/services"
	"api_tinggal_nikah/utils"

	"github.com/labstack/echo/v4"
)

func Add(c echo.Context) error {
	orders := new(dto.RequestTripayDTO)

	if err := c.Bind(orders); err != nil {
		return utils.NewAPIResponse(c).Error(0, err.Error(), err)
	}

	return services.AddPaymentService(c, orders)

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
