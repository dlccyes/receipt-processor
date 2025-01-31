package main

import (
	"github.com/dlccyes/receipt-processor/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/receipts/process", handler.ProcessReceipt)
	r.GET("/receipts/:id/points", handler.GetPoints)

	r.Run(":8080")
}
