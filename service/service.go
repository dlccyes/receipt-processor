package service

import "github.com/dlccyes/receipt-processor/model"

type ReceiptService interface {
	GetReceipt(id int64) (*model.Receipt, bool)
	SaveReceipt(receipt *model.Receipt) int64
}

type PointService interface {
	CalculatePoints(receipt *model.Receipt) int64
}
