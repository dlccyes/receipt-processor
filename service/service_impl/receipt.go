package service_impl

import (
	"github.com/dlccyes/receipt-processor/model"
	"github.com/google/uuid"
)

type receiptServiceImpl struct{}

var receipts = make(map[string]*model.Receipt)

func (*receiptServiceImpl) SaveReceipt(receipt *model.Receipt) string {
	id := uuid.New().String()
	receipts[id] = receipt
	return id
}

func (*receiptServiceImpl) GetReceipt(id string) (*model.Receipt, bool) {
	if receipt, exists := receipts[id]; exists {
		return receipt, true
	}
	return nil, false
}
