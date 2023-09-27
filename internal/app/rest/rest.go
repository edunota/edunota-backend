package rest

import (
	"github.com/edunota/backend/internal/app/rest/handlers"
	"github.com/gofiber/fiber/v2"
)

type IRest interface {
	Listen(port string) error
}
type rest struct {
	*fiber.App
}

func New( /*inject dependencies here*/ ) *rest {

	app := fiber.New()

	app.Get("/hello", handlers.GetHello)
	return &rest{
		App: app,
	}
}
