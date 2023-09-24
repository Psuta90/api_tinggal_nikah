package dto

type ResponseRequestTripay struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    struct {
		Reference            string      `json:"reference"`
		MerchantRef          string      `json:"merchant_ref"`
		PaymentSelectionType string      `json:"payment_selection_type"`
		PaymentMethod        string      `json:"payment_method"`
		PaymentName          string      `json:"payment_name"`
		CustomerName         string      `json:"customer_name"`
		CustomerEmail        string      `json:"customer_email"`
		CustomerPhone        string      `json:"customer_phone"`
		CallbackURL          string      `json:"callback_url"`
		ReturnURL            string      `json:"return_url"`
		Amount               int         `json:"amount"`
		FeeMerchant          int         `json:"fee_merchant"`
		FeeCustomer          int         `json:"fee_customer"`
		TotalFee             int         `json:"total_fee"`
		AmountReceived       int         `json:"amount_received"`
		PayCode              string      `json:"pay_code"`
		PayURL               interface{} `json:"pay_url"`
		CheckoutURL          string      `json:"checkout_url"`
		Status               string      `json:"status"`
		ExpiredTime          int64       `json:"expired_time"`
		OrderItems           []struct {
			Sku        string `json:"sku"`
			Name       string `json:"name"`
			Price      int    `json:"price"`
			Quantity   int    `json:"quantity"`
			Subtotal   int    `json:"subtotal"`
			ProductURL string `json:"product_url"`
			ImageURL   string `json:"image_url"`
		} `json:"order_items"`
		Instructions []struct {
			Title string   `json:"title"`
			Steps []string `json:"steps"`
		} `json:"instructions"`
		QrString interface{} `json:"qr_string"`
		QrURL    interface{} `json:"qr_url"`
	} `json:"data"`
}
