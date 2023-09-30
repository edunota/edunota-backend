package hello

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/edunota/backend/pkg/modules/hello"
	"github.com/stretchr/testify/assert"
)

func TestHelloToNameDto(t *testing.T) {
	dto := hello.HelloToNameDto{
		Name: "jdoe",
	}

	a := assert.New(t)

	dtojson, err := json.Marshal(&dto)

	field, ok := reflect.TypeOf(&dto).Elem().FieldByName("Name")

	a.Equal(ok, true)
	a.Equal(string(field.Tag.Get("param")), "name")
	a.Nil(err)
	a.Equal(string(dtojson), "{\"Name\":\"jdoe\"}")
}
