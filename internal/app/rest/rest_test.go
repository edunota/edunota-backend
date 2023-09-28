package rest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMesssage(t *testing.T) {
	message := GetMessage()
	assert.Equal(t, message, "Hello worlds")
}
