package mock

import (
	"github.com/stretchr/testify/mock"
)

type MockHelloProvider struct {
	mock.Mock
}

func (m *MockHelloProvider) SayHello() string {

	args := m.Called()

	r := args.String(0)

	return r
}

func (m *MockHelloProvider) SayHelloToName(name string) string {

	args := m.Called(name)

	s := args.String(0)
	return s
}
