package service

import "github.com/dlccyes/receipt-processor/model"

type ReceiptService interface {
	SaveReceipt(receipt *model.Receipt) string
	GetReceipt(id string) (*model.Receipt, bool)
}

type PointService interface {
	CalculatePoints(receipt *model.Receipt) int64
}
