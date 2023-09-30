package handlers

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/edunota/backend/pkg/modules/hello"
	mockhello "github.com/edunota/backend/pkg/modules/hello/mock"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
)

func TestSayHello(t *testing.T) {
	app := fiber.New()
	mockHelloProvider := new(mockhello.MockHelloProvider)

	mockHelloProvider.On("SayHello").Return("Hello world")

	fc := fasthttp.RequestCtx{}
	fc.URI().SetPath("/hello")

	ctx := app.AcquireCtx(&fc)

	h := helloHandler{
		providers: map[string]interface{}{
			"hello": mockHelloProvider,
		},
	}

	h.SayHello(ctx)
	resp := hello.SayHelloSerializer{}
	json.Unmarshal(ctx.Response().Body(), &resp)

	assert.Equal(t, hello.SayHelloSerializer{
		Status: 200,
		Data:   "Hello world",
	}, resp)
}

func filter[T interface{}](sl []T, prediction func(item T, index int) bool) []T {
	filtered := sl[:0]
	for in, i := range sl {
		if prediction(i, in) {
			filtered = append(filtered, i)
		}
	}
	return filtered
}

func IfIn[T comparable](slice []T, item T) bool {
	for _, it := range slice {
		if item == it {
			return true
		}
	}
	return false
}

func Map[T interface{}, R any](slice []T, prediction func(item T, index int) R) *[]R {
	mapped := make([]R, 0)
	for i, it := range slice {
		mapped = append(mapped, prediction(it, i))
	}
	return &mapped
}

func TestHelloConstructor(t *testing.T) {

	a := fiber.New()
	p := hello.HelloProvider()
	r := a.Group("/hello")

	HelloHandler(r, p)

	expRoutes := &[]string{"/hello", "/hello/:name"}
	fl := filter[fiber.Route](a.GetRoutes(), func(item fiber.Route, index int) bool {
		return item.Method != "HEAD" &&
			strings.HasPrefix(item.Path, "/hello") &&
			IfIn[string](*expRoutes, item.Path)
	})

	fm := Map[fiber.Route, string](fl, func(item fiber.Route, index int) string {
		return item.Path
	})

	assert.Equal(t, expRoutes, fm)
}
