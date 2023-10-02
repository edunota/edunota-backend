package hello_test

import (
	"testing"

	"github.com/edunota/backend/pkg/modules/hello"
	"github.com/stretchr/testify/assert"
)

func TestHelloProvider(t *testing.T) {

	t.Run("it should implement IHelloProvider", func(t *testing.T) {
		a := assert.New(t)
		p := hello.HelloProvider()
		a.Implements((*hello.IHelloProvider)(nil), p)
	})

	t.Run("it should return Hello world on SayHello()", func(t *testing.T) {
		a := assert.New(t)
		p := hello.HelloProvider()

		a.Equal(p.SayHello(), "Hello world")
	})

	t.Run("it should return Hello {name} on SayHelloToName()", func(t *testing.T) {
		a := assert.New(t)
		p := hello.HelloProvider()

		a.Equal(p.SayHelloToName("Jdoe"), "Hello Jdoe")
	})
}
