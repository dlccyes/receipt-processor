package handler

import (
	"net/http"

	"github.com/dlccyes/receipt-processor/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func (s *Suite) TestGetPoints() {
	mockReceipt := &model.Receipt{
		Retailer: "test seller",
	}
	s.mocks.ReceiptService.On("GetReceipt", int64(1)).Return(mockReceipt, true)
	s.mocks.PointService.On("CalculatePoints", mockReceipt).Return(int64(10), nil)

	s.c.Params = gin.Params{gin.Param{Key: "id", Value: "1"}}
	s.handler.GetPoints(s.c)
	assert.Equal(s.T(), http.StatusOK, s.w.Code)
	assert.Equal(s.T(), "{\"points\":10}", s.w.Body.String())
}

func (s *Suite) TestGetPoints_InvalidID() {
	s.c.Params = gin.Params{gin.Param{Key: "id", Value: "abc"}}
	s.handler.GetPoints(s.c)
	assert.Equal(s.T(), http.StatusBadRequest, s.w.Code)
}

func (s *Suite) TestGetPoints_IDNotFound() {
	s.mocks.ReceiptService.On("GetReceipt", int64(1)).Return((*model.Receipt)(nil), false)
	s.c.Params = gin.Params{gin.Param{Key: "id", Value: "1"}}
	s.handler.GetPoints(s.c)
	assert.Equal(s.T(), http.StatusNotFound, s.w.Code)
}
