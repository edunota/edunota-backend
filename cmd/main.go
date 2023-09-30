package main

import (
	"github.com/edunota/backend/internal/app/rest"
	"github.com/edunota/backend/pkg/modules/hello"
)

func main() {
	helloProvider := hello.HelloProvider()
	Bootstrap(rest.New(helloProvider))
}

func Bootstrap(rest rest.IRest) {
	rest.Listen(":8082")
}
