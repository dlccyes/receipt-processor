package service_impl

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite

	PointService   *pointServiceImpl
	ReceiptService *receiptServiceImpl
}

func (s *Suite) SetupTest() {
	s.PointService = &pointServiceImpl{}
	s.ReceiptService = &receiptServiceImpl{}
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}
