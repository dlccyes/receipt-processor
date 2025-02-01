package handler

import (
	"github.com/dlccyes/receipt-processor/service"
)

type Handler struct {
	ReceiptService service.ReceiptService
	PointService   service.PointService
}

func NewHandler(
	receiptService service.ReceiptService,
	pointService service.PointService,
) *Handler {
	return &Handler{
		ReceiptService: receiptService,
		PointService:   pointService,
	}
}
