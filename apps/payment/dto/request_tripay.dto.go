package dto

type RequestTripayDTO struct {
	Method        string       `json:"method"`
	Merchantref   string       `json:"merchant_ref"`
	Amount        int          `json:"amount"`
	CustomerName  string       `json:"customer_name"`
	CustomerEmail string       `json:"customer_email"`
	CustomerPhone string       `json:"customer_phone"`
	Orders        []OrdersItem `json:"order_items"`
	ReturnUrl     string       `json:"return_url"`
	ExpiredTime   int          `json:"expired_time"`
	Signature     string       `json:"signature"`
}

type OrdersItem struct {
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
}
