package service_impl

import (
	"github.com/dlccyes/receipt-processor/model"
)

type receiptServiceImpl struct{}

var (
	id       = int64(0)
	receipts = make(map[int64]*model.Receipt)
)

func (*receiptServiceImpl) GetReceipt(id int64) (*model.Receipt, bool) {
	if receipt, exists := receipts[id]; exists {
		return receipt, true
	}
	return nil, false
}

func (*receiptServiceImpl) SaveReceipt(receipt *model.Receipt) int64 {
	id += 1
	receipts[id] = receipt
	return id
}
