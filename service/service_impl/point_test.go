package service_impl

import (
	"github.com/dlccyes/receipt-processor/model"
	"github.com/dlccyes/receipt-processor/test"
	"github.com/stretchr/testify/assert"
)

func (s *Suite) TestCalculatePoints() {
	point := s.PointService.CalculatePoints(&model.Receipt{
		Retailer:     "a1234!()/;[]",                                 // #1 +5
		PurchaseDate: test.MustParseTime("2022-01-01", "2006-01-02"), // #6 +6
		PurchaseTime: test.MustParseTime("14:01", "15:04"),           // #7 +10
		Items: []model.Item{ // #4 +5
			{
				ShortDescription: "123", // #5 +1
				Price:            6.00,
			},
			{
				ShortDescription: " 12345",
				Price:            6.00,
			},
		},
		Total: 12.00, // #2 +50 #3 +25
	})
	assert.Equal(s.T(), int64(102), point)
}
