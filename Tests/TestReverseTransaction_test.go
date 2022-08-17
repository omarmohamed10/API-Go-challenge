package test

import (
	business "example/API-Go-challenge/Business"
	Data "example/API-Go-challenge/Data"
	"testing"
)

func TestReverse(t *testing.T) {

	transaction := Data.Transactions[0]
	transaction = business.ReverseTransaction(transaction)

	ExpectedAmount := 0

	if ExpectedAmount != int(transaction.Amount) {
		t.Fatalf("Error: The reversed transaction amount not equal 0")
	}

}
