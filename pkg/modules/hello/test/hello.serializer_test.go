package hello_test

import (
	"encoding/json"
	"testing"

	"github.com/edunota/backend/pkg/modules/hello"
	"github.com/stretchr/testify/assert"
)

func TestHelloSerializer(t *testing.T) {

	a := assert.New(t)
	serializer := hello.SayHelloSerializer{
		Status: 200,
		Data:   "Hello world",
	}

	data, error := json.Marshal(&serializer)

	a.Nil(error)

	a.Equal(string(data), "{\"status\":200,\"data\":\"Hello world\"}")

}
