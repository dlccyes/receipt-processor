package api

import (
	"github.com/dlccyes/receipt-processor/handler"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

type Router struct {
	dig.In
	Handler handler.Handler
}

func (r *Router) Run(addr string) error {
	router := gin.Default()

	router.POST("/receipts/process", r.Handler.ProcessReceipt)
	router.GET("/receipts/:id/points", r.Handler.GetPoints)

	return router.Run(addr)
}
