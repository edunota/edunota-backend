package mock

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/mock"
)

type MockRouter struct {
	mock.Mock
	fiber.Router
}

// Add implements fiber.Router.
func (m *MockRouter) Add(method string, path string, handlers ...func(*fiber.Ctx) error) fiber.Router {
	panic("unimplemented")
}

// All implements fiber.Router.
func (m *MockRouter) All(path string, handlers ...func(*fiber.Ctx) error) fiber.Router {
	panic("unimplemented")
}

// Connect implements fiber.Router.
func (m *MockRouter) Connect(path string, handlers ...func(*fiber.Ctx) error) fiber.Router {
	panic("unimplemented")
}

// Delete implements fiber.Router.
func (m *MockRouter) Delete(path string, handlers ...func(*fiber.Ctx) error) fiber.Router {
	panic("unimplemented")
}

// Get implements fiber.Router.
func (m *MockRouter) Get(path string, handlers ...func(*fiber.Ctx) error) fiber.Router {

	var args mock.Arguments
	if len(handlers) > 1 {
		args = m.Called(path, handlers)
	} else {
		args = m.Called(path, handlers)
	}

	return args.Get(0).(fiber.Router)
}

// Group implements fiber.Router.
func (m *MockRouter) Group(prefix string, handlers ...func(*fiber.Ctx) error) fiber.Router {
	panic("unimplemented")
}

// Head implements fiber.Router.
func (m *MockRouter) Head(path string, handlers ...func(*fiber.Ctx) error) fiber.Router {
	panic("unimplemented")
}

// Mount implements fiber.Router.
func (m *MockRouter) Mount(prefix string, fiber *fiber.App) fiber.Router {
	panic("unimplemented")
}

// Name implements fiber.Router.
func (m *MockRouter) Name(name string) fiber.Router {
	panic("unimplemented")
}

// Options implements fiber.Router.
func (m *MockRouter) Options(path string, handlers ...func(*fiber.Ctx) error) fiber.Router {
	panic("unimplemented")
}

// Patch implements fiber.Router.
func (m *MockRouter) Patch(path string, handlers ...func(*fiber.Ctx) error) fiber.Router {
	panic("unimplemented")
}

// Post implements fiber.Router.
func (m *MockRouter) Post(path string, handlers ...func(*fiber.Ctx) error) fiber.Router {
	panic("unimplemented")
}

// Put implements fiber.Router.
func (m *MockRouter) Put(path string, handlers ...func(*fiber.Ctx) error) fiber.Router {
	panic("unimplemented")
}

// Route implements fiber.Router.
func (m *MockRouter) Route(prefix string, fn func(router fiber.Router), name ...string) fiber.Router {
	panic("unimplemented")
}

// Static implements fiber.Router.
func (m *MockRouter) Static(prefix string, root string, config ...fiber.Static) fiber.Router {
	panic("unimplemented")
}

// Trace implements fiber.Router.
func (m *MockRouter) Trace(path string, handlers ...func(*fiber.Ctx) error) fiber.Router {
	panic("unimplemented")
}

// Use implements fiber.Router.
func (m *MockRouter) Use(args ...interface{}) fiber.Router {
	panic("unimplemented")
}
