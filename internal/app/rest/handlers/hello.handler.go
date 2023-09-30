package handlers

import (
	"github.com/edunota/backend/pkg/modules/hello"
	"github.com/gofiber/fiber/v2"
)

type IHeloHandler interface {
	SayHello(ctx *fiber.Ctx) error
	SayHelloByName(ctx *fiber.Ctx) error
}

type helloHandler struct {
	providers map[string]interface{}
}

func HelloHandler(router fiber.Router, helloProvider hello.IHelloProvider) {

	h := helloHandler{
		providers: map[string]interface{}{
			"hello": helloProvider,
		},
	}
	router.Get("", h.SayHello)
	router.Get(":name", h.SayHelloByName)

}

func (h *helloHandler) SayHello(ctx *fiber.Ctx) error {

	helloProvider := h.providers["hello"].(hello.IHelloProvider)

	data := helloProvider.SayHello()
	body := hello.SayHelloSerializer{
		Status: 200,
		Data:   data,
	}
	return ctx.JSON(&body)
}

func (h *helloHandler) SayHelloByName(ctx *fiber.Ctx) error {

	dto := hello.HelloToNameDto{}
	ctx.ParamsParser(&dto)
	helloProvider := h.providers["hello"].(hello.IHelloProvider)

	return ctx.JSON(&hello.SayHelloSerializer{
		Status: 200,
		Data:   helloProvider.SayHelloToName(dto.Name),
	})
}
