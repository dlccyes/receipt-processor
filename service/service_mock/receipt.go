package service_mock

import (
	"github.com/dlccyes/receipt-processor/model"
	"github.com/stretchr/testify/mock"
)

type MockReceiptService struct {
	mock.Mock
}

func (m *MockReceiptService) GetReceipt(id int64) (*model.Receipt, bool) {
	args := m.Called(id)
	return args.Get(0).(*model.Receipt), args.Get(1).(bool)
}

func (m *MockReceiptService) SaveReceipt(receipt *model.Receipt) int64 {
	args := m.Called(receipt)
	return args.Get(0).(int64)
}
