package services

import (
	"api_tinggal_nikah/apps/payment/dto"
	"api_tinggal_nikah/utils"
	"fmt"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/monaco-io/request"
)

func AddPaymentService(c echo.Context, orderDto *dto.RequestTripayDTO) error {
	//
	expiredTime := time.Now().Add(time.Duration(15) * time.Minute).Unix()
	signature := utils.GenerateSignatureTripay(orderDto.Merchantref, orderDto.Amount)

	orderDto.ExpiredTime = int(expiredTime)
	orderDto.Signature = signature

	req := request.Client{
		URL:    os.Getenv("BASE_URL_TRIPAY") + "/transaction/create",
		Method: "POST",
		Bearer: os.Getenv("API_KEY"),
		JSON:   orderDto,
	}

	var responseData interface{}
	res := req.Send().ScanJSON(&responseData)

	if !res.OK() {
		return utils.NewAPIResponse(c).Error(0, "gagal hit tripay", &responseData)
	}

	fmt.Println(res)

	return utils.NewAPIResponse(c).Success(0, "api for add payment", responseData)
}
