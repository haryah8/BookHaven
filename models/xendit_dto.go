package models

type TopUpRequestXenditDto struct {
	ExternalID string `json:"external_id"`
	Amount     uint   `json:"amount"`
}

type TopUpResponseXenditDto struct {
	Status     string `json:"status"`
	Amount     string `json:"amount"`
	InvoiceURL string `json:"invoice_url"`
	ExpiryDate string `json:"expiry_date"`
}

type TopUpErrorResponseXenditDto struct {
	ErrorCode    string `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}

type TopUpCallbackRequestDto struct {
	Status     string `json:"status"`
	Amount     int    `json:"amount"`
	ExternalId string `json:"external_id"`
}
