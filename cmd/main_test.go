package main

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TSBootstrap struct {
	suite.Suite
}

type MockRest struct {
	mock.Mock
}

func (m MockRest) Listen(port string) error {
	args := m.Called(port)
	return args.Error(0)
}
func (m MockRest) Shutdown() error {
	args := m.Called()
	return args.Error(0)
}
func TestTSBootstrap(t *testing.T) {
	suite.Run(t, &TSBootstrap{})
}

func (s *TSBootstrap) TestListen() {

	mr := MockRest{}
	mr.On("Listen", ":8082").Return(nil)
	Bootstrap(mr)

}
