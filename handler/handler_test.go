package handler

import (
	"net/http/httptest"
	"testing"

	"github.com/dlccyes/receipt-processor/service/service_mock"
	"github.com/dlccyes/receipt-processor/test"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite

	c       *gin.Context
	w       *httptest.ResponseRecorder
	handler Handler
	mocks   struct {
		PointService   *service_mock.MockPointService
		ReceiptService *service_mock.MockReceiptService
	}
}

func (s *Suite) SetupTest() {
	s.c, s.w = test.SetupHttpTest()
	s.mocks.PointService = &service_mock.MockPointService{}
	s.mocks.ReceiptService = &service_mock.MockReceiptService{}
	s.handler = Handler{
		ReceiptService: s.mocks.ReceiptService,
		PointService:   s.mocks.PointService,
	}
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}
