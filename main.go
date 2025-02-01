package main

import (
	"github.com/dlccyes/receipt-processor/api"
	"github.com/dlccyes/receipt-processor/handler"
	"github.com/dlccyes/receipt-processor/service/service_impl"
	"github.com/dlccyes/receipt-processor/utils/di"
	"go.uber.org/dig"
)

func main() {
	container := dig.New()
	service_impl.Bind(container)
	di.MustProvide(container, handler.NewHandler)
	di.MustProvide(container, api.NewRouter)
	di.MustInvoke(container, func(router *api.Router) {
		if err := router.Run(":8080"); err != nil {
			panic(err)
		}
	})
}
