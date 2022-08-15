package main

import (
	"github.com/gin-gonic/gin"

	Controller "example/API-Go-challenge/Controllers"
)

func main() {

	router := gin.Default()
	router.GET("/Transactions", Controller.GetAllTransactions)
	router.GET("/Transaction/:id", Controller.TransactionById)
	router.POST("/AddTransaction", Controller.CreateTransaction)
	router.PATCH("/ReverseTransaction", Controller.ReverseTransaction)
	router.Run("localhost:8080")

}
