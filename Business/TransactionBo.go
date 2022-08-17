package business

import (
	Entities "example/API-Go-challenge/Entities"
)

func ReverseTransaction(transaction Entities.Transaction) Entities.Transaction {

	transaction.Amount = 0
	return transaction

}
