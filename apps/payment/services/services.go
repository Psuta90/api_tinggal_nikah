package services

import (
	"api_tinggal_nikah/apps/payment/dto"
	"api_tinggal_nikah/config"
	"api_tinggal_nikah/db"
	"api_tinggal_nikah/models"
	"api_tinggal_nikah/repository"
	"api_tinggal_nikah/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/monaco-io/request"
)

func AddPaymentService(c echo.Context, orderDto *dto.RequestPaymentUserDto) error {

	conn := db.GetDB()
	UserTransactionRepo := repository.NewUserTransactionRepository(conn)

	payloadJwt := c.Get("JWT").(*jwt.Token).Claims.(*config.JwtCustomClaims)
	responseTripay := new(dto.ResponseRequestTripay)

	lastIdTransaction, err := UserTransactionRepo.GetLastID()
	if err != nil {
		return utils.NewAPIResponse(c).Error(0, "telah terjadi masalah pada sistem silahkan coba lagi", err)
	}

	order_id := fmt.Sprintf("order-%d", lastIdTransaction+1)

	bodyTripay := &dto.RequestTripayDTO{
		Method:        orderDto.Method,
		Merchantref:   order_id,
		Amount:        orderDto.Amount,
		CustomerName:  payloadJwt.Name,
		CustomerEmail: payloadJwt.Email,
		Orders:        orderDto.OrderItems,
		ReturnUrl:     "",
		ExpiredTime:   int(time.Now().Add(10 * time.Minute).Unix()),
		Signature:     utils.GenerateSignatureTripay(order_id, orderDto.Amount),
	}

	req := request.Client{
		URL:    os.Getenv("BASE_URL_TRIPAY") + "/transaction/create",
		Method: "POST",
		Bearer: os.Getenv("API_KEY"),
		JSON:   bodyTripay,
	}

	res := req.Send().ScanJSON(&responseTripay)

	resJsonTripay, err := utils.StructToMap(responseTripay)
	if err != nil {
		return utils.NewAPIResponse(c).Error(0, "gagal conver struct to map response tripay", err)
	}

	if !responseTripay.Success || res.Response().StatusCode != 200 {
		return utils.NewAPIResponse(c).Error(0, "gagal hit tripay", responseTripay)
	}

	UserTransaction := &models.UserTransaction{
		OrderID:           order_id,
		UserID:            payloadJwt.ID,
		ResponseTripay:    resJsonTripay,
		ExpiredOrder:      time.Unix(responseTripay.Data.ExpiredTime, 0),
		PackageCategoryID: orderDto.PackageID,
		Status:            models.PENDING,
	}

	if err := UserTransactionRepo.Create(UserTransaction); err != nil {
		return utils.NewAPIResponse(c).Error(0, "gagal insert database", err)
	}

	return utils.NewAPIResponse(c).Success(0, "Berhasil", responseTripay.Data)
}

func ListPaymentChannelService(c echo.Context) error {

	listChannelPayment := new(dto.ResponsePaymentChannelTripayDto)
	ttlRedisCached := 30 * 24 * time.Hour

	cachedData, err := db.RedisClient.Get(c.Request().Context(), "listChannel").Result()
	if err != nil {
		req := request.Client{
			URL:    os.Getenv("BASE_URL_TRIPAY") + "/merchant/payment-channel",
			Method: "GET",
			Bearer: os.Getenv("API_KEY"),
		}

		res := req.Send().ScanJSON(&listChannelPayment)

		if !res.OK() {
			log.Println(res.Error())
		}

		jsonStr, err := json.Marshal(listChannelPayment)
		if err != nil {
			return utils.NewAPIResponse(c).Error(http.StatusInternalServerError, err.Error(), nil)
		}

		if err := db.RedisClient.Set(c.Request().Context(), "listChannel", jsonStr, ttlRedisCached).Err(); err != nil {
			return utils.NewAPIResponse(c).Error(http.StatusInternalServerError, err.Error(), nil)
		}

		return utils.NewAPIResponse(c).Success(0, "success", listChannelPayment.Data)
	}

	if err := json.Unmarshal([]byte(cachedData), &listChannelPayment); err != nil {
		return utils.NewAPIResponse(c).Error(http.StatusInternalServerError, err.Error(), nil)
	}

	return utils.NewAPIResponse(c).Success(0, "success", listChannelPayment.Data)

}

func CallBackTripayService(c echo.Context, data *dto.ResponseCallbackTripayDto) error {

	fmt.Println("masuk servcie")

	conn := db.GetDB()
	UserTransactionRepo := repository.NewUserTransactionRepository(conn)

	userTransaction, err := UserTransactionRepo.FindOneByOrderID(data.MerchantRef)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"success": false,
			"message": "Order ID Not Found",
		})
	}

	resJsonCallbackStr, err := utils.StructToMap(data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"success": false,
			"message": "gagal stringfy json to save in database",
		})
	}

	userTransaction.Status = models.StatusPayment(data.Status)
	userTransaction.ResponseTripay = resJsonCallbackStr

	if err := UserTransactionRepo.UpdateByOrderID(userTransaction); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"success": true,
		"message": "Berhasil Update",
	})
}
