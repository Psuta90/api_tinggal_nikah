package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
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
