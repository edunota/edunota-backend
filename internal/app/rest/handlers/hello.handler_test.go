package handlers

import (
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/valyala/fasthttp"
)

type MockCtx struct {
	mock.Mock
}

func (m *MockCtx) SendString(s string) error {
	args := m.Called(s)

	return args.Error(0)
}

func TestHello(t *testing.T) {
	app := fiber.New()
	c := app.AcquireCtx(&fasthttp.RequestCtx{})

	GetHello(c)

	b := c.Context().Response.Body()

	assert.Equal(t, string(b), "Hello world")
}
