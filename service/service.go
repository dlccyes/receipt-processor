package service

import "github.com/dlccyes/receipt-processor/model"

type ReceiptService interface {
	SaveReceipt(receipt *model.Receipt) int64
	GetReceipt(id int64) (*model.Receipt, bool)
}

type PointService interface {
	CalculatePoints(receipt *model.Receipt) int64
}
