package main

import "github.com/edunota/backend/internal/app/rest"

func main() {

	Bootstrap(rest.New())
}

func Bootstrap(rest rest.IRest) {
	if err := rest.Listen(":8082"); err != nil {
		panic(err)
	}
}
