package service

import (
	"github.com/dlccyes/receipt-processor/model"
	"github.com/google/uuid"
)

var receipts = make(map[string]*model.Receipt)

//TODO: interface & dependency injection

func SaveReceipt(receipt *model.Receipt) string {
	id := uuid.New().String()
	receipts[id] = receipt
	return id
}

func GetReceipt(id string) (*model.Receipt, bool) {
	if receipt, exists := receipts[id]; exists {
		return receipt, true
	}
	return nil, false
}
