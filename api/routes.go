package api

import (
	"github.com/dlccyes/receipt-processor/handler"
	"github.com/gin-gonic/gin"
)

type Router struct {
	Handler *handler.Handler
}

func NewRouter(handler *handler.Handler) *Router {
	return &Router{
		Handler: handler,
	}
}

func (r *Router) Run(addr string) error {
	router := gin.Default()

	router.POST("/receipts/process", r.Handler.ProcessReceipt)
	router.GET("/receipts/:id/points", r.Handler.GetPoints)

	return router.Run(addr)
}
