package service_impl

import (
	"github.com/dlccyes/receipt-processor/model"
	"github.com/dlccyes/receipt-processor/test"
	"github.com/stretchr/testify/assert"
)

func (s *Suite) TestCalculatePoints() {
	testCases := []struct {
		caseName       string
		receipt        *model.Receipt
		expectedPoints int64
	}{
		{
			caseName: "test case 1",
			receipt: &model.Receipt{
				Retailer:     "Target",
				PurchaseDate: test.MustParseDate("2022-01-01"),
				PurchaseTime: test.MustParseTime("13:01"),
				Items: []model.Item{
					{
						ShortDescription: "Mountain Dew 12PK",
						Price:            6.49,
					},
					{
						ShortDescription: "Emils Cheese Pizza",
						Price:            12.25,
					},
					{
						ShortDescription: "Knorr Creamy Chicken",
						Price:            1.26,
					},
					{
						ShortDescription: "Doritos Nacho Cheese",
						Price:            3.35,
					},
					{
						ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ",
						Price:            12.00,
					},
				},
				Total: 35.35,
			},
			expectedPoints: 28,
		},
		{
			caseName: "test case 2",
			receipt: &model.Receipt{
				Retailer:     "M&M Corner Market",
				PurchaseDate: test.MustParseDate("2022-03-20"),
				PurchaseTime: test.MustParseTime("14:33"),
				Items: []model.Item{
					{
						ShortDescription: "Gatorade",
						Price:            2.25,
					},
					{
						ShortDescription: "Gatorade",
						Price:            2.25,
					},
					{
						ShortDescription: "Gatorade",
						Price:            2.25,
					},
					{
						ShortDescription: "Gatorade",
						Price:            2.25,
					},
				},
				Total: 9.00,
			},
			expectedPoints: 109,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.caseName, func() {
			point := s.PointService.CalculatePoints(tc.receipt)
			assert.Equal(s.T(), tc.expectedPoints, point)
		})
	}
}
