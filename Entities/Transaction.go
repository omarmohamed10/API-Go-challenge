package Entities

import (
	"time"
)

type Transaction struct {
	ID           string    `json:"id"`
	Amount       float32   `json:"amount"`
	CardId       string    `json:"cardId"`
	Currency     string    `json:"currency"`
	CreatedAT    time.Time `json:"date"`
	MerchantId   string    `json:"merchantId"`
	CompanyId    string    `json:"companyId"`
	MerchantName string    `json:"merchantName"`
}
