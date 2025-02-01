package handler

import (
	"github.com/dlccyes/receipt-processor/service"
	"go.uber.org/dig"
)

type Handler struct {
	dig.In

	ReceiptService service.ReceiptService
	PointService   service.PointService
}
