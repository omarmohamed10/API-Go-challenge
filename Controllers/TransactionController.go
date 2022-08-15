package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	Data "example/API-Go-challenge/Data"
	Entities "example/API-Go-challenge/Entities"
	business "example/API-Go-challenge/business"
)

func GetAllTransactions(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Data.Transactions)
}

func CreateTransaction(c *gin.Context) {

	var NewTransaction Entities.Transaction
	err := c.BindJSON(&NewTransaction)
	if err != nil {
		return
	}
	Data.Transactions = append(Data.Transactions, NewTransaction)
	c.IndentedJSON(http.StatusCreated, NewTransaction)

}

func TransactionById(c *gin.Context) {

	id := c.Param("id")
	transaction, err := GetTransactionById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Transaction Not Found"})
		return
	}
	c.IndentedJSON(http.StatusOK, transaction)
}

func GetTransactionById(id string) (*Entities.Transaction, error) {
	for _, transaction := range Data.Transactions {
		if transaction.ID == id {
			return &transaction, nil
		}
	}
	return nil, errors.New("transaction not found")
}

func ReverseTransaction(c *gin.Context) {

	id, err := c.GetQuery("id")

	if !err {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing the id query parameter"})
		return
	}

	transaction, ok := GetTransactionById(id)
	if ok != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "transaction not found"})
		return
	}
	transaction = business.ReverseTransaction(*transaction)

	c.IndentedJSON(http.StatusOK, transaction)

}
