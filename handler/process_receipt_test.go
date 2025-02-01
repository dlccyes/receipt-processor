package handler

import (
	"net/http"

	"github.com/dlccyes/receipt-processor/model"
	"github.com/dlccyes/receipt-processor/test"
	"github.com/stretchr/testify/assert"
)

func (s *Suite) TestProcessReceipt() {
	receipt := &model.Receipt{
		Retailer:     "test seller",
		PurchaseDate: "2022-01-02",
		PurchaseTime: "18:01",
		Items: []model.Item{
			{
				ShortDescription: "test",
				Price:            "5.00",
			},
			{
				ShortDescription: "test",
				Price:            "5.00",
			},
		},
		Total: "10.00",
	}
	s.mocks.ReceiptService.On("SaveReceipt", receipt).Return(int64(5))

	test.SetCtxRequestBody(s.c, receipt)
	s.handler.ProcessReceipt(s.c)
	assert.Equal(s.T(), http.StatusOK, s.w.Code)
	assert.Equal(s.T(), "{\"id\":\"5\"}", s.w.Body.String())
}

func (s *Suite) TestProcessReceipt_IncompleteReceipt() {
	testCases := []struct {
		caseName string
		receipt  *model.Receipt
	}{
		{
			caseName: "nil receipt",
			receipt:  nil,
		},
		{
			caseName: "empty receipt",
			receipt:  &model.Receipt{},
		},
		{
			caseName: "no retailer",
			receipt: &model.Receipt{
				PurchaseDate: "2022-01-02",
				PurchaseTime: "18:01",
				Items: []model.Item{
					{
						ShortDescription: "test",
						Price:            "5.00",
					},
				},
				Total: "5.00",
			},
		},
		{
			caseName: "no purchase date",
			receipt: &model.Receipt{
				Retailer:     "test seller",
				PurchaseTime: "18:01",
				Items: []model.Item{
					{
						ShortDescription: "test",
						Price:            "5.00",
					},
					{
						ShortDescription: "test",
						Price:            "5.00",
					},
				},
				Total: "10.00",
			},
		},
		{
			caseName: "no purchase time",
			receipt: &model.Receipt{
				Retailer:     "test seller",
				PurchaseDate: "2022-01-02",
				Items: []model.Item{
					{
						ShortDescription: "test",
						Price:            "5.00",
					},
					{
						ShortDescription: "test",
						Price:            "5.00",
					},
				},
				Total: "10.00",
			},
		},
		{
			caseName: "no items",
			receipt: &model.Receipt{
				Retailer:     "test seller",
				PurchaseDate: "2022-01-02",
				PurchaseTime: "18:01",
				Items:        []model.Item{},
				Total:        "10.00",
			},
		},
		{
			caseName: "no total",
			receipt: &model.Receipt{
				Retailer:     "test seller",
				PurchaseDate: "2022-01-02",
				PurchaseTime: "18:01",
				Items: []model.Item{
					{
						ShortDescription: "test",
						Price:            "5.00",
					},
					{
						ShortDescription: "test",
						Price:            "5.00",
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		test.SetCtxRequestBody(s.c, tc.receipt)
		s.handler.ProcessReceipt(s.c)
		assert.Equal(s.T(), http.StatusBadRequest, s.w.Code)
	}
}

func (s *Suite) TestProcessReceipt_InvalidReceiptFormat() {
	testCases := []struct {
		caseName string
		receipt  *model.Receipt
	}{
		{
			caseName: "invalid retailer",
			receipt: &model.Receipt{
				Retailer:     "test seller @",
				PurchaseDate: "2022-01-02",
				PurchaseTime: "18:01",
				Items: []model.Item{
					{
						ShortDescription: "test",
						Price:            "5.00",
					},
				},
				Total: "5.00",
			},
		},
		{
			caseName: "invalid total",
			receipt: &model.Receipt{
				Retailer:     "test seller @",
				PurchaseDate: "2022-01-02",
				PurchaseTime: "18:01",
				Items: []model.Item{
					{
						ShortDescription: "test",
						Price:            "5.00",
					},
				},
				Total: "abc",
			},
		},
		{
			caseName: "invalid item price",
			receipt: &model.Receipt{
				Retailer:     "test seller @",
				PurchaseDate: "2022-01-02",
				PurchaseTime: "18:01",
				Items: []model.Item{
					{
						ShortDescription: "test",
						Price:            "abc",
					},
				},
				Total: "5.00",
			},
		},
	}

	for _, tc := range testCases {
		test.SetCtxRequestBody(s.c, tc.receipt)
		s.handler.ProcessReceipt(s.c)
		assert.Equal(s.T(), http.StatusBadRequest, s.w.Code)
	}
}
