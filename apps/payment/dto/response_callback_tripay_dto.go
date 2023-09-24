package dto

type ResponseCallbackTripayDto struct {
	Reference         string      `json:"reference"`
	MerchantRef       string      `json:"merchant_ref"`
	PaymentMethod     string      `json:"payment_method"`
	PaymentMethodCode string      `json:"payment_method_code"`
	TotalAmount       int         `json:"total_amount"`
	FeeMerchant       int         `json:"fee_merchant"`
	FeeCustomer       int         `json:"fee_customer"`
	TotalFee          int         `json:"total_fee"`
	AmountReceived    int         `json:"amount_received"`
	IsClosedPayment   int         `json:"is_closed_payment"`
	Status            string      `json:"status"`
	PaidAt            int         `json:"paid_at"`
	Note              interface{} `json:"note"`
}
