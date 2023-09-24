package dto

type ResponsePaymentChannelTripayDto struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    []struct {
		Group      string `json:"group"`
		Code       string `json:"code"`
		Name       string `json:"name"`
		Type       string `json:"type"`
		MinimumFee int    `json:"minimum_fee"`
		MaximumFee int    `json:"maximum_fee"`
		IconURL    string `json:"icon_url"`
		Active     bool   `json:"active"`
	} `json:"data"`
}
