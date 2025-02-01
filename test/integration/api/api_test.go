package api

import (
	"testing"

	"github.com/dlccyes/receipt-processor/api"
	"github.com/dlccyes/receipt-processor/service/service_impl"
	"github.com/dlccyes/receipt-processor/utils/di"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"go.uber.org/dig"
)

type Suite struct {
	suite.Suite

	container *dig.Container
	router    *gin.Engine
}

func (s *Suite) SetupTest() {
	s.container = dig.New()
	service_impl.Bind(s.container)
	di.MustInvoke(s.container, func(router api.Router) {
		s.router = router.Init()
	})
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}
