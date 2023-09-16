package main

import (
	rest "github.com/edunota/backend/internal/app/rest"
)

func main() {

	rest.New().Run(":8082")
}
