package service_impl

import (
	"github.com/dlccyes/receipt-processor/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func (s *Suite) TestCalculatePoints() {
	point, err := s.PointService.CalculatePoints(&model.Receipt{
		Retailer:     "a1234!()/;[]", // #1 +5
		PurchaseDate: "2022-01-01",   // #6 +6
		PurchaseTime: "14:01",        // #7 +10
		Items: []model.Item{ // #4 +5
			{
				ShortDescription: "123", // #5 +1
				Price:            "6.00",
			},
			{
				ShortDescription: " 12345",
				Price:            "6.00",
			},
		},
		Total: "12.00", // #2 +50 #3 +25
	})
	require.NoError(s.T(), err)
	assert.Equal(s.T(), int64(102), point)
}
