package service_impl

import (
	"github.com/dlccyes/receipt-processor/service"
	"github.com/dlccyes/receipt-processor/utils/di"
	"go.uber.org/dig"
)

func Bind(container *dig.Container) {
	di.MustProvide(container, func() service.ReceiptService {
		return &receiptServiceImpl{}
	})
	di.MustProvide(container, func() service.PointService {
		return &pointServiceImpl{}
	})
}
