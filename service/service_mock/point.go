package service_mock

import (
	"github.com/dlccyes/receipt-processor/model"
	"github.com/stretchr/testify/mock"
)

type MockPointService struct {
	mock.Mock
}

func (m *MockPointService) CalculatePoints(receipt *model.Receipt) int64 {
	args := m.Called(receipt)
	return args.Get(0).(int64)
}
