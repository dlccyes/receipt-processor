package handler

import (
	"net/http"

	"github.com/dlccyes/receipt-processor/model"
	"github.com/dlccyes/receipt-processor/test"
	"github.com/stretchr/testify/assert"
)

func (s *Suite) TestProcessReceipt() {
	req := &ProcessReceiptReq{
		Retailer:     "test seller",
		PurchaseDate: "2022-01-02",
		PurchaseTime: "18:01",
		Items: []ProcessReceiptReqItem{
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
	expectedReceipt := &model.Receipt{
		Retailer:     "test seller",
		PurchaseDate: test.MustParseDate("2022-01-02"),
		PurchaseTime: test.MustParseTime("18:01"),
		Items: []model.Item{
			{
				ShortDescription: "test",
				Price:            5.00,
			},
			{
				ShortDescription: "test",
				Price:            5.00,
			},
		},
		Total: 10.00,
	}
	s.mocks.ReceiptService.On("SaveReceipt", expectedReceipt).Return(int64(5))

	test.SetCtxRequestBody(s.c, req)
	s.handler.ProcessReceipt(s.c)
	assert.Equal(s.T(), http.StatusOK, s.w.Code)
	assert.Equal(s.T(), "{\"id\":\"5\"}", s.w.Body.String())
}

func (s *Suite) TestProcessReceipt_IncompleteReceipt() {
	testCases := []struct {
		caseName string
		receipt  *ProcessReceiptReq
	}{
		{
			caseName: "nil receipt",
			receipt:  nil,
		},
		{
			caseName: "empty receipt",
			receipt:  &ProcessReceiptReq{},
		},
		{
			caseName: "no retailer",
			receipt: &ProcessReceiptReq{
				PurchaseDate: "2022-01-02",
				PurchaseTime: "18:01",
				Items: []ProcessReceiptReqItem{
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
			receipt: &ProcessReceiptReq{
				Retailer:     "test seller",
				PurchaseTime: "18:01",
				Items: []ProcessReceiptReqItem{
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
			receipt: &ProcessReceiptReq{
				Retailer:     "test seller",
				PurchaseDate: "2022-01-02",
				Items: []ProcessReceiptReqItem{
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
			receipt: &ProcessReceiptReq{
				Retailer:     "test seller",
				PurchaseDate: "2022-01-02",
				PurchaseTime: "18:01",
				Items:        []ProcessReceiptReqItem{},
				Total:        "10.00",
			},
		},
		{
			caseName: "no total",
			receipt: &ProcessReceiptReq{
				Retailer:     "test seller",
				PurchaseDate: "2022-01-02",
				PurchaseTime: "18:01",
				Items: []ProcessReceiptReqItem{
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
		receipt  *ProcessReceiptReq
	}{
		{
			caseName: "invalid retailer",
			receipt: &ProcessReceiptReq{
				Retailer:     "test seller @",
				PurchaseDate: "2022-01-02",
				PurchaseTime: "18:01",
				Items: []ProcessReceiptReqItem{
					{
						ShortDescription: "test",
						Price:            "5.00",
					},
				},
				Total: "5.00",
			},
		},
		{
			caseName: "invalid purchase date",
			receipt: &ProcessReceiptReq{
				Retailer:     "test seller @",
				PurchaseDate: "20220102",
				PurchaseTime: "18:01",
				Items: []ProcessReceiptReqItem{
					{
						ShortDescription: "test",
						Price:            "5.00",
					},
				},
				Total: "5.00",
			},
		},
		{
			caseName: "invalid purchase time",
			receipt: &ProcessReceiptReq{
				Retailer:     "test seller @",
				PurchaseDate: "2022-01-02",
				PurchaseTime: "1801",
				Items: []ProcessReceiptReqItem{
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
			receipt: &ProcessReceiptReq{
				Retailer:     "test seller @",
				PurchaseDate: "2022-01-02",
				PurchaseTime: "18:01",
				Items: []ProcessReceiptReqItem{
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
			receipt: &ProcessReceiptReq{
				Retailer:     "test seller @",
				PurchaseDate: "2022-01-02",
				PurchaseTime: "18:01",
				Items: []ProcessReceiptReqItem{
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
