package rest

import (
	"reflect"
	"testing"

	"github.com/edunota/backend/pkg/modules/hello"
	"github.com/gofiber/fiber/v2"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	helloProvider := hello.HelloProvider()
	r := New(helloProvider)

	a := assert.New(t)

	reflected := reflect.ValueOf(r)
	app := reflect.Indirect(reflected).FieldByName("App").Interface().(*fiber.App)
	count := app.HandlersCount()

	a.Implements((*IRest)(nil), r)
	a.GreaterOrEqual(count, uint32(1))
	a.Nil(r.Shutdown())

}
