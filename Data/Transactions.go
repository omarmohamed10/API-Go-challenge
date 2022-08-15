package Data

import (
	Entities "example/API-Go-challenge/Entities"
	"time"

	"github.com/google/uuid"
)

var Transactions = []Entities.Transaction{

	{ID: uuid.New().String(), Amount: 45.44, CardId: uuid.New().String(), Currency: "usd",
		CreatedAT: time.Now().UTC(), MerchantId: uuid.New().String(), CompanyId: uuid.New().String(), MerchantName: "Amazon"},

	{ID: uuid.New().String(), Amount: 125.00, CardId: uuid.New().String(), Currency: "usd",
		CreatedAT: time.Now().UTC(), MerchantId: uuid.New().String(), CompanyId: uuid.New().String(), MerchantName: "Talabat"},

	{ID: uuid.New().String(), Amount: 12446.77, CardId: uuid.New().String(), Currency: "usd",
		CreatedAT: time.Now().UTC(), MerchantId: uuid.New().String(), CompanyId: uuid.New().String(), MerchantName: "Apple"},
}
