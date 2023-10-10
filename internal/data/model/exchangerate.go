package model

type ExchangeRates struct {
	Data []ExchangeRate `json:"data,omitempty"`
}

type ExchangeRate struct {
	ExchangeRate string `json:"exchange_rate,omitempty"`
	RecordDate   string `json:"record_date,omitempty"`
}
