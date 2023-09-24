package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"

	"os"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

// 	$privateKey   = 'ytf6ooi2gmlNPfpchd94jDOk8hRWOu';
// $merchantCode = 'T0001';
// $merchantRef  = 'INV55567';
// $amount       = 1500000;

func GenerateSignatureTripay(orderID string, amount int) string {
	hasher := hmac.New(sha256.New, []byte(os.Getenv("PRIVATE_KEY")))
	hasher.Write([]byte(os.Getenv("MERCHANT_CODE") + orderID + fmt.Sprint(amount)))
	signature := hex.EncodeToString(hasher.Sum(nil))

	fmt.Println(signature)

	return signature

}

func ValidateSignaturTripay(c echo.Context, token string) error {

	callbackSignature := c.Request().Header.Get("X-Callback-Signature")
	callbackEvent := c.Request().Header.Get("X-Callback-Event")

	fmt.Println("token : ", token)
	fmt.Println("signature : ", os.Getenv("SIGNATURE_CALLBACK"))
	fmt.Println(callbackSignature)

	if err := bcrypt.CompareHashAndPassword([]byte(token), []byte(os.Getenv("SIGNATURE_CALLBACK"))); err != nil {
		fmt.Println(err.Error())
		return errors.New("invalid signature")
	}

	if callbackEvent != "payment_status" {
		fmt.Println("masuk error payment status")
		return errors.New("invalid signature")
	}

	if callbackSignature == "" {
		fmt.Println("masuk error callback empty")
		return errors.New("invalid signature")
	}

	return nil

	// return signature == callbackSignature
}
