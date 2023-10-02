package rest

import (
	"github.com/edunota/backend/internal/app/rest/handlers"
	"github.com/edunota/backend/pkg/modules/hello"
	"github.com/gofiber/fiber/v2"
)

type IRest interface {
	Listen(port string) error
	Shutdown() error
}
type Rest struct {
	App *fiber.App
}

func (r *Rest) Listen(addr string) error {
	return r.App.Listen(addr)
}

func (r *Rest) Shutdown() error {
	return r.App.Shutdown()
}
func New(helloProvider hello.IHelloProvider) IRest {

	app := fiber.New()

	// handlers
	handlers.HelloHandler(app.Group("/hello"), helloProvider)

	return &Rest{
		App: app,
	}
}
