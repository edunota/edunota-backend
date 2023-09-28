package rest

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockedContext struct {
	mock.Mock
}

func TestGetMesssage(t *testing.T) {
	message := GetMessage()
	assert.Equal(t, message, "Hello worlds")
}
