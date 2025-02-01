package main

import (
	"github.com/dlccyes/receipt-processor/api"
	"github.com/dlccyes/receipt-processor/service/service_impl"
	"github.com/dlccyes/receipt-processor/utils/di"
	"go.uber.org/dig"
)

func main() {
	container := dig.New()
	service_impl.Bind(container)
	di.MustInvoke(container, func(router api.Router) {
		if err := router.Init().Run(":8080"); err != nil {
			panic(err)
		}
	})
}
